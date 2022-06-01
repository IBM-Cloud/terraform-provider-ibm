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
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/validate"
	"github.ibm.com/org-ids/toolchain-go-sdk/toolchainv2"
)

func ResourceIBMCdToolchainToolGitlab() *schema.Resource {
	return &schema.Resource{
		CreateContext: ResourceIBMCdToolchainToolGitlabCreate,
		ReadContext:   ResourceIBMCdToolchainToolGitlabRead,
		UpdateContext: ResourceIBMCdToolchainToolGitlabUpdate,
		DeleteContext: ResourceIBMCdToolchainToolGitlabDelete,
		Importer:      &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"toolchain_id": &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validate.InvokeValidator("ibm_cd_toolchain_tool_gitlab", "toolchain_id"),
				Description:  "ID of the toolchain to bind integration to.",
			},
			"name": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validate.InvokeValidator("ibm_cd_toolchain_tool_gitlab", "name"),
				Description:  "Name of tool integration.",
			},
			"parameters": &schema.Schema{
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Description: "Parameters to be used to create the integration.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"git_id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"title": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "e.g. My GitLab Enterprise Server.",
						},
						"api_root_url": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "e.g. https://gitlab.example.com/api/v4.",
						},
						"default_branch": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "e.g. main.",
						},
						"root_url": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "e.g. https://gitlab.example.com.",
						},
						"access_token": &schema.Schema{
							Type:             schema.TypeString,
							Optional:         true,
							Computed:         true,
							DiffSuppressFunc: flex.SuppressHashedRawSecret,
							Sensitive:        true,
						},
						"blind_connection": &schema.Schema{
							Type:        schema.TypeBool,
							Optional:    true,
							Computed:    true,
							Description: "Select this checkbox only if the server is not addressable on the public internet. IBM Cloud will not be able to validate the connection details you provide. Certain functionality that requires API access to the git server will be disabled. Delivery pipeline will only work using a private worker that has network access to the git server.",
						},
						"owner_id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"repo_name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"repo_url": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "Type the URL of the repository that you are linking to.",
						},
						"source_repo_url": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "Type the URL of the repository that you are forking or cloning.",
						},
						"token_url": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "Integration token URL.",
						},
						"type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"private_repo": &schema.Schema{
							Type:        schema.TypeBool,
							Optional:    true,
							Computed:    true,
							Description: "Select this check box to make this repository private.",
						},
						"has_issues": &schema.Schema{
							Type:        schema.TypeBool,
							Optional:    true,
							Default:     true,
							Description: "Select this check box to enable GitLab Issues for lightweight issue tracking.",
						},
						"enable_traceability": &schema.Schema{
							Type:        schema.TypeBool,
							Optional:    true,
							Default:     false,
							Description: "Select this check box to track the deployment of code changes by creating tags, labels and comments on commits, pull requests and referenced issues.",
						},
						"authorized": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"integration_owner": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "Select the user which git operations will be performed as.",
						},
						"auth_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"api_token": &schema.Schema{
							Type:             schema.TypeString,
							Optional:         true,
							Computed:         true,
							DiffSuppressFunc: flex.SuppressHashedRawSecret,
							Sensitive:        true,
							Description:      "Personal Access Token.",
						},
					},
				},
			},
			"initialization": &schema.Schema{
				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"repo_name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
						"repo_url": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
							Description: "Type the URL of the repository that you are linking to.",
						},
						"source_repo_url": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
							Description: "Type the URL of the repository that you are forking or cloning.",
						},
						"type": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"private_repo": &schema.Schema{
							Type:        schema.TypeBool,
							Optional:    true,
							Default:     true,
							ForceNew:    true,
							Description: "Select this check box to make this repository private.",
						},
					},
				},
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
							Optional:    true,
							Description: "URI representing the this resource through the UI.",
						},
						"api_href": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "URI representing the this resource through an API.",
						},
					},
				},
			},
			"updated_at": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Latest tool integration update timestamp.",
			},
			"state": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Current configuration state of the tool integration.",
			},
			"integration_id": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Tool integration ID.",
			},
		},
	}
}

func ResourceIBMCdToolchainToolGitlabValidator() *validate.ResourceValidator {
	validateSchema := make([]validate.ValidateSchema, 1)
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 "toolchain_id",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Required:                   true,
			Regexp:                     `^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[89abAB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$`,
			MinValueLength:             36,
			MaxValueLength:             36,
		},
		validate.ValidateSchema{
			Identifier:                 "name",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Optional:                   true,
			Regexp:                     `^([^\\x00-\\x7F]|[a-zA-Z0-9-._ ])+$`,
			MinValueLength:             0,
			MaxValueLength:             128,
		},
	)

	resourceValidator := validate.ResourceValidator{ResourceName: "ibm_cd_toolchain_tool_gitlab", Schema: validateSchema}
	return &resourceValidator
}

func ResourceIBMCdToolchainToolGitlabCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	toolchainClient, err := meta.(conns.ClientSession).ToolchainV2()
	if err != nil {
		return diag.FromErr(err)
	}

	createIntegrationOptions := &toolchainv2.CreateIntegrationOptions{}

	createIntegrationOptions.SetToolchainID(d.Get("toolchain_id").(string))
	createIntegrationOptions.SetToolID("gitlab")
	if _, ok := d.GetOk("name"); ok {
		createIntegrationOptions.SetName(d.Get("name").(string))
	}
	_, pok := d.GetOk("parameters")
	_, iok := d.GetOk("initialization")
	if pok || iok {
		parametersModel := GetParametersForCreate(d, ResourceIBMCdToolchainToolGitlab(), nil)
		createIntegrationOptions.SetParameters(parametersModel)
	}

	postIntegrationResponse, response, err := toolchainClient.CreateIntegrationWithContext(context, createIntegrationOptions)
	if err != nil {
		log.Printf("[DEBUG] CreateIntegrationWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("CreateIntegrationWithContext failed %s\n%s", err, response))
	}

	d.SetId(fmt.Sprintf("%s/%s", *createIntegrationOptions.ToolchainID, *postIntegrationResponse.ID))

	return ResourceIBMCdToolchainToolGitlabRead(context, d, meta)
}

func ResourceIBMCdToolchainToolGitlabRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	toolchainClient, err := meta.(conns.ClientSession).ToolchainV2()
	if err != nil {
		return diag.FromErr(err)
	}

	getIntegrationByIDOptions := &toolchainv2.GetIntegrationByIDOptions{}

	parts, err := flex.SepIdParts(d.Id(), "/")
	if err != nil {
		return diag.FromErr(err)
	}

	getIntegrationByIDOptions.SetToolchainID(parts[0])
	getIntegrationByIDOptions.SetIntegrationID(parts[1])

	getIntegrationByIDResponse, response, err := toolchainClient.GetIntegrationByIDWithContext(context, getIntegrationByIDOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		log.Printf("[DEBUG] GetIntegrationByIDWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("GetIntegrationByIDWithContext failed %s\n%s", err, response))
	}

	if err = d.Set("toolchain_id", getIntegrationByIDResponse.ToolchainID); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting toolchain_id: %s", err))
	}
	if err = d.Set("name", getIntegrationByIDResponse.Name); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting name: %s", err))
	}
	if getIntegrationByIDResponse.Parameters != nil {
		parametersMap := GetParametersFromRead(getIntegrationByIDResponse.Parameters, ResourceIBMCdToolchainToolGitlab(), nil)
		if err = d.Set("parameters", []map[string]interface{}{parametersMap}); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting parameters: %s", err))
		}
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
	referentMap, err := ResourceIBMCdToolchainToolGitlabToolIntegrationReferentToMap(getIntegrationByIDResponse.Referent)
	if err != nil {
		return diag.FromErr(err)
	}
	if err = d.Set("referent", []map[string]interface{}{referentMap}); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting referent: %s", err))
	}
	if err = d.Set("updated_at", flex.DateTimeToString(getIntegrationByIDResponse.UpdatedAt)); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting updated_at: %s", err))
	}
	if err = d.Set("state", getIntegrationByIDResponse.State); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting state: %s", err))
	}
	if err = d.Set("integration_id", getIntegrationByIDResponse.ID); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting integration_id: %s", err))
	}

	return nil
}

func ResourceIBMCdToolchainToolGitlabUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	toolchainClient, err := meta.(conns.ClientSession).ToolchainV2()
	if err != nil {
		return diag.FromErr(err)
	}

	updateIntegrationOptions := &toolchainv2.UpdateIntegrationOptions{}

	parts, err := flex.SepIdParts(d.Id(), "/")
	if err != nil {
		return diag.FromErr(err)
	}

	updateIntegrationOptions.SetToolchainID(parts[0])
	updateIntegrationOptions.SetIntegrationID(parts[1])
	updateIntegrationOptions.SetToolID("gitlab")

	hasChange := false

	if d.HasChange("toolchain_id") {
		return diag.FromErr(fmt.Errorf("Cannot update resource property \"%s\" with the ForceNew annotation."+
			" The resource must be re-created to update this property.", "toolchain_id"))
	}
	if d.HasChange("name") {
		updateIntegrationOptions.SetName(d.Get("name").(string))
		hasChange = true
	}
	if d.HasChange("parameters") {
		parameters := GetParametersForUpdate(d, ResourceIBMCdToolchainToolGitlab(), nil)
		updateIntegrationOptions.SetParameters(parameters)
		hasChange = true
	}

	if hasChange {
		response, err := toolchainClient.UpdateIntegrationWithContext(context, updateIntegrationOptions)
		if err != nil {
			log.Printf("[DEBUG] UpdateIntegrationWithContext failed %s\n%s", err, response)
			return diag.FromErr(fmt.Errorf("UpdateIntegrationWithContext failed %s\n%s", err, response))
		}
	}

	return ResourceIBMCdToolchainToolGitlabRead(context, d, meta)
}

func ResourceIBMCdToolchainToolGitlabDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	toolchainClient, err := meta.(conns.ClientSession).ToolchainV2()
	if err != nil {
		return diag.FromErr(err)
	}

	deleteIntegrationOptions := &toolchainv2.DeleteIntegrationOptions{}

	parts, err := flex.SepIdParts(d.Id(), "/")
	if err != nil {
		return diag.FromErr(err)
	}

	deleteIntegrationOptions.SetToolchainID(parts[0])
	deleteIntegrationOptions.SetIntegrationID(parts[1])

	response, err := toolchainClient.DeleteIntegrationWithContext(context, deleteIntegrationOptions)
	if err != nil {
		log.Printf("[DEBUG] DeleteIntegrationWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("DeleteIntegrationWithContext failed %s\n%s", err, response))
	}

	d.SetId("")

	return nil
}

func ResourceIBMCdToolchainToolGitlabToolIntegrationReferentToMap(model *toolchainv2.ToolIntegrationReferent) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.UIHref != nil {
		modelMap["ui_href"] = model.UIHref
	}
	if model.APIHref != nil {
		modelMap["api_href"] = model.APIHref
	}
	return modelMap, nil
}
