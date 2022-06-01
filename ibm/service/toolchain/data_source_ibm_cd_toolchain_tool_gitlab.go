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

func DataSourceIBMCdToolchainToolGitlab() *schema.Resource {
	return &schema.Resource{
		ReadContext: DataSourceIBMCdToolchainToolGitlabRead,

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
			"id": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Tool integration ID.",
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
						"git_id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"title": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "e.g. My GitLab Enterprise Server.",
						},
						"api_root_url": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "e.g. https://gitlab.example.com/api/v4.",
						},
						"default_branch": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "e.g. main.",
						},
						"root_url": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "e.g. https://gitlab.example.com.",
						},
						"access_token": &schema.Schema{
							Type:      schema.TypeString,
							Computed:  true,
							Sensitive: true,
						},
						"blind_connection": &schema.Schema{
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Select this checkbox only if the server is not addressable on the public internet. IBM Cloud will not be able to validate the connection details you provide. Certain functionality that requires API access to the git server will be disabled. Delivery pipeline will only work using a private worker that has network access to the git server.",
						},
						"owner_id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"repo_name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"repo_url": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Type the URL of the repository that you are linking to.",
						},
						"source_repo_url": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Type the URL of the repository that you are forking or cloning.",
						},
						"token_url": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Integration token URL.",
						},
						"type": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"private_repo": &schema.Schema{
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Select this check box to make this repository private.",
						},
						"has_issues": &schema.Schema{
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Select this check box to enable GitLab Issues for lightweight issue tracking.",
						},
						"enable_traceability": &schema.Schema{
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Select this check box to track the deployment of code changes by creating tags, labels and comments on commits, pull requests and referenced issues.",
						},
						"authorized": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"integration_owner": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Select the user which git operations will be performed as.",
						},
						"auth_type": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"api_token": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Sensitive:   true,
							Description: "Personal Access Token.",
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

func DataSourceIBMCdToolchainToolGitlabRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	if *getIntegrationByIDResponse.ToolID != "gitlab" {
		return diag.FromErr(fmt.Errorf("Retrieved tool is not the correct type: %s", *getIntegrationByIDResponse.ToolID))
	}

	d.SetId(*getIntegrationByIDResponse.ID)

	if err = d.Set("id", getIntegrationByIDResponse.ID); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting id: %s", err))
	}

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
		modelMap, err := DataSourceIBMCdToolchainToolGitlabToolIntegrationReferentToMap(getIntegrationByIDResponse.Referent)
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
		modelMap := GetParametersFromRead(getIntegrationByIDResponse.Parameters, DataSourceIBMCdToolchainToolGitlab(), nil)
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

func DataSourceIBMCdToolchainToolGitlabToolIntegrationReferentToMap(model *toolchainv2.ToolIntegrationReferent) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.UIHref != nil {
		modelMap["ui_href"] = *model.UIHref
	}
	if model.APIHref != nil {
		modelMap["api_href"] = *model.APIHref
	}
	return modelMap, nil
}
