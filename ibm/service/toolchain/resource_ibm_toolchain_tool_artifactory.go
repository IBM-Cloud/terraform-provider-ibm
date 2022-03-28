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

func ResourceIBMToolchainToolArtifactory() *schema.Resource {
	return &schema.Resource{
		CreateContext: ResourceIBMToolchainToolArtifactoryCreate,
		ReadContext:   ResourceIBMToolchainToolArtifactoryRead,
		UpdateContext: ResourceIBMToolchainToolArtifactoryUpdate,
		DeleteContext: ResourceIBMToolchainToolArtifactoryDelete,
		Importer:      &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"toolchain_id": &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validate.InvokeValidator("ibm_toolchain_tool_artifactory", "toolchain_id"),
				Description:  "ID of the toolchain to bind integration to.",
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
						"name": &schema.Schema{
							Type:        schema.TypeString,
							Required:    true,
							Description: "Type a name for this tool integration, for example: my-artifactory. This name displays on your toolchain.",
						},
						"dashboard_url": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Type the URL that you want to navigate to when you click the Artifactory integration tile.",
						},
						"type": &schema.Schema{
							Type:        schema.TypeString,
							Required:    true,
							Description: "Choose the type of repository for your Artifactory integration.",
						},
						"user_id": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Type the User ID or email for your Artifactory repository.",
						},
						"token": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Type the API key for your Artifactory repository.",
						},
						"release_url": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Type the URL for your Artifactory release repository.",
						},
						"mirror_url": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Type the URL for your Artifactory virtual repository, which is a repository that can see your private repositories and a cache of the public repositories.",
						},
						"snapshot_url": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Type the URL for your Artifactory snapshot repository.",
						},
						"repository_name": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Type the name of your artifactory repository where your docker images are located.",
						},
						"repository_url": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Type the URL of your artifactory repository where your docker images are located.",
						},
						"docker_config_json": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
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

func ResourceIBMToolchainToolArtifactoryValidator() *validate.ResourceValidator {
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

	resourceValidator := validate.ResourceValidator{ResourceName: "ibm_toolchain_tool_artifactory", Schema: validateSchema}
	return &resourceValidator
}

func ResourceIBMToolchainToolArtifactoryCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	toolchainClient, err := meta.(conns.ClientSession).ToolchainV2()
	if err != nil {
		return diag.FromErr(err)
	}

	postIntegrationOptions := &toolchainv2.PostIntegrationOptions{}

	postIntegrationOptions.SetToolchainID(d.Get("toolchain_id").(string))
	postIntegrationOptions.SetToolID("artifactory")
	if _, ok := d.GetOk("name"); ok {
		postIntegrationOptions.SetName(d.Get("name").(string))
	}
	if _, ok := d.GetOk("parameters"); ok {
		parametersModel, err := ResourceIBMToolchainToolArtifactoryMapToParameters(d.Get("parameters.0").(map[string]interface{}))
		if err != nil {
			return diag.FromErr(err)
		}
		postIntegrationOptions.SetParameters(parametersModel)
	}

	postIntegrationResponse, response, err := toolchainClient.PostIntegrationWithContext(context, postIntegrationOptions)
	if err != nil {
		log.Printf("[DEBUG] PostIntegrationWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("PostIntegrationWithContext failed %s\n%s", err, response))
	}

	d.SetId(fmt.Sprintf("%s/%s", *postIntegrationOptions.ToolchainID, *postIntegrationResponse.ID))

	return ResourceIBMToolchainToolArtifactoryRead(context, d, meta)
}

func ResourceIBMToolchainToolArtifactoryRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	// TODO: handle argument of type map[string]interface{}
	if err = d.Set("toolchain_id", getIntegrationByIDResponse.ToolchainID); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting toolchain_id: %s", err))
	}
	if err = d.Set("name", getIntegrationByIDResponse.Name); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting name: %s", err))
	}
	if getIntegrationByIDResponse.Parameters != nil {
		parametersMap, err := ResourceIBMToolchainToolArtifactoryParametersToMap(getIntegrationByIDResponse.Parameters, d)
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
	referentMap, err := ResourceIBMToolchainToolArtifactoryGetIntegrationByIDResponseReferentToMap(getIntegrationByIDResponse.Referent)
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

func ResourceIBMToolchainToolArtifactoryUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
	patchToolIntegrationOptions.SetToolID("artifactory")

	hasChange := false

	if d.HasChange("toolchain_id") {
		return diag.FromErr(fmt.Errorf("Cannot update resource property \"%s\" with the ForceNew annotation."+
			" The resource must be re-created to update this property.", "toolchain_id"))
	}
	if d.HasChange("name") {
		patchToolIntegrationOptions.SetName(d.Get("name").(string))
		hasChange = true
	}
	if d.HasChange("parameters") {
		parameters, err := ResourceIBMToolchainToolArtifactoryMapToParameters(d.Get("parameters.0").(map[string]interface{}))
		if err != nil {
			return diag.FromErr(err)
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

	return ResourceIBMToolchainToolArtifactoryRead(context, d, meta)
}

func ResourceIBMToolchainToolArtifactoryDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

func ResourceIBMToolchainToolArtifactoryMapToParameters(modelMap map[string]interface{}) (map[string]interface{}, error) {
	model := make(map[string]interface{})
	model["name"] = core.StringPtr(modelMap["name"].(string))
	if modelMap["dashboard_url"] != nil {
		model["dashboard_url"] = core.StringPtr(modelMap["dashboard_url"].(string))
	}
	model["type"] = core.StringPtr(modelMap["type"].(string))
	if modelMap["user_id"] != nil {
		model["user_id"] = core.StringPtr(modelMap["user_id"].(string))
	}
	if modelMap["token"] != nil {
		model["token"] = core.StringPtr(modelMap["token"].(string))
	}
	if modelMap["release_url"] != nil {
		model["release_url"] = core.StringPtr(modelMap["release_url"].(string))
	}
	if modelMap["mirror_url"] != nil {
		model["mirror_url"] = core.StringPtr(modelMap["mirror_url"].(string))
	}
	if modelMap["snapshot_url"] != nil {
		model["snapshot_url"] = core.StringPtr(modelMap["snapshot_url"].(string))
	}
	if modelMap["repository_name"] != nil {
		model["repository_name"] = core.StringPtr(modelMap["repository_name"].(string))
	}
	if modelMap["repository_url"] != nil {
		model["repository_url"] = core.StringPtr(modelMap["repository_url"].(string))
	}
	return model, nil
}

func ResourceIBMToolchainToolArtifactoryParametersToMap(model map[string]interface{}, d *schema.ResourceData) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["name"] = model["name"]
	if model["dashboard_url"] != nil {
		modelMap["dashboard_url"] = model["dashboard_url"]
	}
	modelMap["type"] = model["type"]
	if model["user_id"] != nil {
		modelMap["user_id"] = model["user_id"]
	}
	if model["token"] != nil {
		if model["token"] == "****" {
			modelMap["token"] = d.Get("parameters.0.token").(string)
		} else {
			modelMap["token"] = model["token"]
		}
	}
	if model["release_url"] != nil {
		modelMap["release_url"] = model["release_url"]
	}
	if model["mirror_url"] != nil {
		modelMap["mirror_url"] = model["mirror_url"]
	}
	if model["snapshot_url"] != nil {
		modelMap["snapshot_url"] = model["snapshot_url"]
	}
	if model["repository_name"] != nil {
		modelMap["repository_name"] = model["repository_name"]
	}
	if model["repository_url"] != nil {
		modelMap["repository_url"] = model["repository_url"]
	}
	return modelMap, nil
}

func ResourceIBMToolchainToolArtifactoryGetIntegrationByIDResponseReferentToMap(model *toolchainv2.GetIntegrationByIDResponseReferent) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.UIHref != nil {
		modelMap["ui_href"] = model.UIHref
	}
	if model.APIHref != nil {
		modelMap["api_href"] = model.APIHref
	}
	return modelMap, nil
}
