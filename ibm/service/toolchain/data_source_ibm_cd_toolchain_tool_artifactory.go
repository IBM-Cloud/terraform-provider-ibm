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

func DataSourceIBMCdToolchainToolArtifactory() *schema.Resource {
	return &schema.Resource{
		ReadContext: DataSourceIBMCdToolchainToolArtifactoryRead,

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
							Description: "Type a name for this tool integration, for example: my-artifactory. This name displays on your toolchain.",
						},
						"dashboard_url": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Type the URL that you want to navigate to when you click the Artifactory integration tile.",
						},
						"type": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Choose the type of repository for your Artifactory integration.",
						},
						"user_id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Type the User ID or email for your Artifactory repository.",
						},
						"token": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Sensitive:   true,
							Description: "Type the API key for your Artifactory repository.",
						},
						"release_url": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Type the URL for your Artifactory release repository.",
						},
						"mirror_url": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Type the URL for your Artifactory virtual repository, which is a repository that can see your private repositories and a cache of the public repositories.",
						},
						"snapshot_url": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Type the URL for your Artifactory snapshot repository.",
						},
						"repository_name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Type the name of your artifactory repository where your docker images are located.",
						},
						"repository_url": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Type the URL of your artifactory repository where your docker images are located.",
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

func DataSourceIBMCdToolchainToolArtifactoryRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	if *getIntegrationByIDResponse.ToolID != "artifactory" {
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
		modelMap, err := DataSourceIBMCdToolchainToolArtifactoryToolIntegrationReferentToMap(getIntegrationByIDResponse.Referent)
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
		modelMap := GetParametersFromRead(getIntegrationByIDResponse.Parameters, DataSourceIBMCdToolchainToolArtifactory(), nil)
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

func DataSourceIBMCdToolchainToolArtifactoryToolIntegrationReferentToMap(model *toolchainv2.ToolIntegrationReferent) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.UIHref != nil {
		modelMap["ui_href"] = *model.UIHref
	}
	if model.APIHref != nil {
		modelMap["api_href"] = *model.APIHref
	}
	return modelMap, nil
}
