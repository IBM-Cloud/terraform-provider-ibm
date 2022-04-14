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
	"github.com/IBM/go-sdk-core/v5/core"
	"github.ibm.com/org-ids/toolchain-go-sdk/toolchainv2"
)

func ResourceIBMToolchainToolGit() *schema.Resource {
	return &schema.Resource{
		CreateContext: ResourceIBMToolchainToolGitCreate,
		ReadContext:   ResourceIBMToolchainToolGitRead,
		UpdateContext: ResourceIBMToolchainToolGitUpdate,
		DeleteContext: ResourceIBMToolchainToolGitDelete,
		Importer:      &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"toolchain_id": &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validate.InvokeValidator("ibm_toolchain_tool_git", "toolchain_id"),
				Description:  "ID of the toolchain to bind integration to.",
			},
			"git_provider": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The Git provider.",
			},
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Name of tool integration.",
			},
			"parameters": &schema.Schema{
				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enable_traceability": &schema.Schema{
							Type:     schema.TypeBool,
							Optional: true,
						},
						"has_issues": &schema.Schema{
							Type:     schema.TypeBool,
							Optional: true,
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
						"type": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"private_repo": &schema.Schema{
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Select this check box to make this repository private.",
						},
						"git_id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"owner_id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
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
							Optional: true,
							ForceNew: true,
						},
						"private_repo": &schema.Schema{
							Type:        schema.TypeBool,
							Optional:    true,
							Default:     false,
							ForceNew:    true,
							Description: "Select this check box to make this repository private.",
						},
						"git_id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
						"owner_id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
					},
				},
			},
			"resource_group_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"crn": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"toolchain_crn": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"href": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"referent": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ui_href": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"api_href": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"updated_at": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func ResourceIBMToolchainToolGitValidator() *validate.ResourceValidator {
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
	)

	resourceValidator := validate.ResourceValidator{ResourceName: "ibm_toolchain_tool_git", Schema: validateSchema}
	return &resourceValidator
}

func ResourceIBMToolchainToolGitCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	toolchainClient, err := meta.(conns.ClientSession).ToolchainV2()
	if err != nil {
		return diag.FromErr(err)
	}

	postIntegrationOptions := &toolchainv2.PostIntegrationOptions{}

	postIntegrationOptions.SetToolchainID(d.Get("toolchain_id").(string))
	postIntegrationOptions.SetToolID(d.Get("git_provider").(string))
	if _, ok := d.GetOk("name"); ok {
		postIntegrationOptions.SetName(d.Get("name").(string))
	}

	modelMapParam := make(map[string]interface{})
	if _, ok := d.GetOk("parameters"); ok {
		modelMapParam = d.Get("parameters.0").(map[string]interface{})
	}
	parameters, err := ResourceIBMToolchainToolGitMapToParametersCreate(d.Get("initialization.0").(map[string]interface{}), modelMapParam)
	if err != nil {
		return diag.FromErr(err)
	}
	if d.Get("git_provider").(string) == "github_integrated" {
		parameters["legal"] = true
	}
	postIntegrationOptions.SetParameters(parameters)

	postIntegrationResponse, response, err := toolchainClient.PostIntegrationWithContext(context, postIntegrationOptions)
	if err != nil {
		log.Printf("[DEBUG] PostIntegrationWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("PostIntegrationWithContext failed %s\n%s", err, response))
	}

	d.SetId(fmt.Sprintf("%s/%s", *postIntegrationOptions.ToolchainID, *postIntegrationResponse.ID))

	return ResourceIBMToolchainToolGitRead(context, d, meta)
}

func ResourceIBMToolchainToolGitRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	// TODO: handle argument of type Initialization
	// TODO: handle argument of type map[string]interface{}
	if err = d.Set("toolchain_id", getIntegrationByIDResponse.ToolchainID); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting toolchain_id: %s", err))
	}
	if err = d.Set("name", getIntegrationByIDResponse.Name); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting name: %s", err))
	}
	if getIntegrationByIDResponse.Parameters != nil {
		parametersMap, err := ResourceIBMToolchainToolGitParametersToMap(getIntegrationByIDResponse.Parameters)
		if err != nil {
			return diag.FromErr(err)
		}
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
	referentMap, err := ResourceIBMToolchainToolGitGetIntegrationByIDResponseReferentToMap(getIntegrationByIDResponse.Referent)
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

	return nil
}

func ResourceIBMToolchainToolGitUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	toolchainClient, err := meta.(conns.ClientSession).ToolchainV2()
	if err != nil {
		return diag.FromErr(err)
	}

	patchToolIntegrationOptions := &toolchainv2.PatchToolIntegrationOptions{}

	parts, err := flex.SepIdParts(d.Id(), "/")
	if err != nil {
		return diag.FromErr(err)
	}

	patchToolIntegrationOptions.SetToolchainID(parts[0])
	patchToolIntegrationOptions.SetIntegrationID(parts[1])
	patchToolIntegrationOptions.SetToolID(d.Get("git_provider").(string))

	hasChange := false

	if d.HasChange("toolchain_id") {
		return diag.FromErr(fmt.Errorf("Cannot update resource property \"%s\" with the ForceNew annotation."+
			" The resource must be re-created to update this property.", "toolchain_id"))
	}
	if d.HasChange("git_provider") {
		return diag.FromErr(fmt.Errorf("Cannot update resource property \"%s\" with the ForceNew annotation."+
			" The resource must be re-created to update this property.", "git_provider"))
	}
	if d.HasChange("name") {
		patchToolIntegrationOptions.SetName(d.Get("name").(string))
		hasChange = true
	}
	if d.HasChange("parameters") {
		parameters, err := ResourceIBMToolchainToolGitMapToParametersUpdate(d.Get("parameters.0").(map[string]interface{}))
		if err != nil {
			return diag.FromErr(err)
		}
		if d.Get("git_provider").(string) == "github_integrated" {
			parameters["legal"] = true
		}
		patchToolIntegrationOptions.SetParameters(parameters)
		hasChange = true
	}

	if hasChange {
		_, response, err := toolchainClient.PatchToolIntegrationWithContext(context, patchToolIntegrationOptions)
		if err != nil {
			log.Printf("[DEBUG] PatchToolIntegrationWithContext failed %s\n%s", err, response)
			return diag.FromErr(fmt.Errorf("PatchToolIntegrationWithContext failed %s\n%s", err, response))
		}
	}

	return ResourceIBMToolchainToolGitRead(context, d, meta)
}

func ResourceIBMToolchainToolGitDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	toolchainClient, err := meta.(conns.ClientSession).ToolchainV2()
	if err != nil {
		return diag.FromErr(err)
	}

	deleteToolIntegrationOptions := &toolchainv2.DeleteToolIntegrationOptions{}

	parts, err := flex.SepIdParts(d.Id(), "/")
	if err != nil {
		return diag.FromErr(err)
	}

	deleteToolIntegrationOptions.SetToolchainID(parts[0])
	deleteToolIntegrationOptions.SetIntegrationID(parts[1])

	response, err := toolchainClient.DeleteToolIntegrationWithContext(context, deleteToolIntegrationOptions)
	if err != nil {
		log.Printf("[DEBUG] DeleteToolIntegrationWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("DeleteToolIntegrationWithContext failed %s\n%s", err, response))
	}

	d.SetId("")

	return nil
}

func ResourceIBMToolchainToolGitMapToParametersCreate(modelMapInit map[string]interface{}, modelMapParam map[string]interface{}) (map[string]interface{}, error) {
	model := make(map[string]interface{})
	if modelMapInit["repo_name"] != nil {
		model["repo_name"] = core.StringPtr(modelMapInit["repo_name"].(string))
	}
	if modelMapInit["repo_url"] != nil {
		model["repo_url"] = core.StringPtr(modelMapInit["repo_url"].(string))
		log.Printf("[INFO] init %s", *model["repo_url"].(*string))
	}
	if modelMapInit["source_repo_url"] != nil {
		model["source_repo_url"] = core.StringPtr(modelMapInit["source_repo_url"].(string))
	}
	if modelMapInit["type"] != nil {
		model["type"] = core.StringPtr(modelMapInit["type"].(string))
	}
	if modelMapInit["private_repo"] != nil {
		model["private_repo"] = core.BoolPtr(modelMapInit["private_repo"].(bool))
	}
	if modelMapInit["git_id"] != nil {
		model["git_id"] = core.StringPtr(modelMapInit["git_id"].(string))
	}
	if modelMapInit["owner_id"] != nil {
		model["owner_id"] = core.StringPtr(modelMapInit["owner_id"].(string))
	}
	if modelMapParam["enable_traceability"] != nil {
		model["enable_traceability"] = core.BoolPtr(modelMapParam["enable_traceability"].(bool))
	}
	if modelMapParam["has_issues"] != nil {
		model["has_issues"] = core.BoolPtr(modelMapParam["has_issues"].(bool))
	}
	return model, nil
}

func ResourceIBMToolchainToolGitMapToParametersUpdate(modelMap map[string]interface{}) (map[string]interface{}, error) {
	model := make(map[string]interface{})
	if modelMap["enable_traceability"] != nil {
		model["enable_traceability"] = core.BoolPtr(modelMap["enable_traceability"].(bool))
	}
	if modelMap["has_issues"] != nil {
		model["has_issues"] = core.BoolPtr(modelMap["has_issues"].(bool))
	}
	if modelMap["repo_name"] != nil {
		model["repo_name"] = core.StringPtr(modelMap["repo_name"].(string))
	}
	if modelMap["repo_url"] != nil {
		model["repo_url"] = core.StringPtr(modelMap["repo_url"].(string))
	}
	if modelMap["source_repo_url"] != nil {
		model["source_repo_url"] = core.StringPtr(modelMap["source_repo_url"].(string))
	}
	if modelMap["type"] != nil {
		model["type"] = core.StringPtr(modelMap["type"].(string))
	}
	if modelMap["private_repo"] != nil {
		model["private_repo"] = core.BoolPtr(modelMap["private_repo"].(bool))
	}
	if modelMap["git_id"] != nil {
		model["git_id"] = core.StringPtr(modelMap["git_id"].(string))
	}
	if modelMap["owner_id"] != nil {
		model["owner_id"] = core.StringPtr(modelMap["owner_id"].(string))
	}
	return model, nil
}

func ResourceIBMToolchainToolGitParametersToMap(model map[string]interface{}) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model["enable_traceability"] != nil {
		modelMap["enable_traceability"] = model["enable_traceability"]
	}
	if model["has_issues"] != nil {
		modelMap["has_issues"] = model["has_issues"]
	}
	if model["repo_name"] != nil {
		modelMap["repo_name"] = model["repo_name"]
	}
	if model["repo_url"] != nil {
		modelMap["repo_url"] = model["repo_url"]
	}
	if model["source_repo_url"] != nil {
		modelMap["source_repo_url"] = model["source_repo_url"]
	}
	if model["type"] != nil {
		modelMap["type"] = model["type"]
	}
	if model["private_repo"] != nil {
		modelMap["private_repo"] = model["private_repo"]
	}
	if model["git_id"] != nil {
		modelMap["git_id"] = model["git_id"]
	}
	if model["owner_id"] != nil {
		modelMap["owner_id"] = model["owner_id"]
	}
	return modelMap, nil
}

func ResourceIBMToolchainToolGitGetIntegrationByIDResponseReferentToMap(model *toolchainv2.GetIntegrationByIDResponseReferent) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.UIHref != nil {
		modelMap["ui_href"] = model.UIHref
	}
	if model.APIHref != nil {
		modelMap["api_href"] = model.APIHref
	}
	return modelMap, nil
}
