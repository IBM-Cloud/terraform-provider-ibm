// Copyright IBM Corp. 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package toolchain

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.ibm.com/org-ids/toolchain-go-sdk/toolchainv2"
)

func DataSourceIBMCdToolchainToolSecuritycompliance() *schema.Resource {
	return &schema.Resource{
		ReadContext: DataSourceIBMCdToolchainToolSecuritycomplianceRead,

		Schema: map[string]*schema.Schema{
			"toolchain_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "ID of the toolchain.",
			},
			"integration_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "ID of the tool integration bound to the toolchain.",
			},
			"resource_group_id": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Resource group where tool integration can be found.",
			},
			"crn": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Tool integration CRN.",
			},
			"toolchain_crn": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "CRN of toolchain which the integration is bound to.",
			},
			"href": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "URI representing the tool integration.",
			},
			"referent": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Information on URIs to access this resource through the UI or API.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ui_href": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "URI representing the this resource through the UI.",
						},
						"api_href": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "URI representing the this resource through an API.",
						},
					},
				},
			},
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Tool integration name.",
			},
			"updated_at": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Latest tool integration update timestamp.",
			},
			"parameters": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Parameters to be used to create the integration.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Give this tool integration a name, for example: my-security-compliance.",
						},
						"evidence_repo_name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "To collect and store evidence for all tasks performed, a Git repository is required as an evidence locker.",
						},
						"trigger_scan": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Enabling trigger validation scans provides details for a pipeline task to trigger a scan.",
						},
						"location": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"evidence_namespace": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The kind of pipeline evidence to be displayed in Security and Compliance Center for this toolchain. The evidence locker will be searched for CD (Continuous Deployment) pipeline evidence, or for CC (Continuous Compliance) pipeline evidence.",
						},
						"api_key": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Sensitive:   true,
							Description: "The IBM Cloud API key is used to access the Security and Compliance API. You can obtain your API key with 'ibmcloud iam api-key-create' or via the console at https://cloud.ibm.com/iam#/apikeys by clicking **Create API key** (Each API key only can be viewed once).",
						},
						"scope": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Select an existing scope name to narrow the focus of the validation scan. [Learn more.](https://cloud.ibm.com/docs/security-compliance?topic=security-compliance-scopes) ![](https://cloud.ibm.com/media/docs/images/icons/launch-glyph.svg).",
						},
						"profile": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Select an existing profile, where a profile is a collection of security controls. [Learn more.](https://cloud.ibm.com/docs/security-compliance?topic=security-compliance-profiles) ![](https://cloud.ibm.com/media/docs/images/icons/launch-glyph.svg).",
						},
						"trigger_info": &schema.Schema{
							Type:        schema.TypeMap,
							Computed:    true,
							Description: "trigger_info.",
						},
					},
				},
			},
			"state": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Current configuration state of the tool integration.",
			},
		},
	}
}

func DataSourceIBMCdToolchainToolSecuritycomplianceRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	toolchainClient, err := meta.(conns.ClientSession).ToolchainV2()
	if err != nil {
		return diag.FromErr(err)
	}

	getIntegrationByIDOptions := &toolchainv2.GetIntegrationByIDOptions{}

	getIntegrationByIDOptions.SetToolchainID(d.Get("toolchain_id").(string))
	getIntegrationByIDOptions.SetIntegrationID(d.Get("integration_id").(string))

	getIntegrationByIDResponse, response, err := toolchainClient.GetIntegrationByIDWithContext(context, getIntegrationByIDOptions)
	if err != nil {
		log.Printf("[DEBUG] GetIntegrationByIDWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("GetIntegrationByIDWithContext failed %s\n%s", err, response))
	}

	if *getIntegrationByIDResponse.ToolID != "security_compliance" {
		return diag.FromErr(fmt.Errorf("Retrieved tool is not the correct type: %s", *getIntegrationByIDResponse.ToolID))
	}

	d.SetId(fmt.Sprintf("%s/%s", *getIntegrationByIDOptions.ToolchainID, *getIntegrationByIDOptions.IntegrationID))

	if err = d.Set("resource_group_id", getIntegrationByIDResponse.ResourceGroupID); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting resource_group_id: %s", err))
	}

	if err = d.Set("crn", getIntegrationByIDResponse.CRN); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting crn: %s", err))
	}

	if err = d.Set("toolchain_crn", getIntegrationByIDResponse.ToolchainCRN); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting toolchain_crn: %s", err))
	}

	if err = d.Set("href", getIntegrationByIDResponse.Href); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting href: %s", err))
	}

	referent := []map[string]interface{}{}
	if getIntegrationByIDResponse.Referent != nil {
		modelMap, err := DataSourceIBMCdToolchainToolSecuritycomplianceToolIntegrationReferentToMap(getIntegrationByIDResponse.Referent)
		if err != nil {
			return diag.FromErr(err)
		}
		referent = append(referent, modelMap)
	}
	if err = d.Set("referent", referent); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting referent %s", err))
	}

	if err = d.Set("name", getIntegrationByIDResponse.Name); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting name: %s", err))
	}

	if err = d.Set("updated_at", flex.DateTimeToString(getIntegrationByIDResponse.UpdatedAt)); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting updated_at: %s", err))
	}

	parameters := []map[string]interface{}{}
	if getIntegrationByIDResponse.Parameters != nil {
		remapFields := map[string]string{
			"api_key": "api-key",
		}
		modelMap := GetParametersFromRead(getIntegrationByIDResponse.Parameters, DataSourceIBMCdToolchainToolSecuritycompliance(), remapFields)
		parameters = append(parameters, modelMap)
	}
	if err = d.Set("parameters", parameters); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting parameters %s", err))
	}

	if err = d.Set("state", getIntegrationByIDResponse.State); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting state: %s", err))
	}

	return nil
}

func DataSourceIBMCdToolchainToolSecuritycomplianceToolIntegrationReferentToMap(model *toolchainv2.ToolIntegrationReferent) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.UIHref != nil {
		modelMap["ui_href"] = *model.UIHref
	}
	if model.APIHref != nil {
		modelMap["api_href"] = *model.APIHref
	}
	return modelMap, nil
}
