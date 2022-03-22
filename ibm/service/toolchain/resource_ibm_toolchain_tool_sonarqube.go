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

func ResourceIbmToolchainToolSonarqube() *schema.Resource {
	return &schema.Resource{
		CreateContext: ResourceIbmToolchainToolSonarqubeCreate,
		ReadContext:   ResourceIbmToolchainToolSonarqubeRead,
		UpdateContext: ResourceIbmToolchainToolSonarqubeUpdate,
		DeleteContext: ResourceIbmToolchainToolSonarqubeDelete,
		Importer:      &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"toolchain_id": &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validate.InvokeValidator("ibm_toolchain_tool_sonarqube", "toolchain_id"),
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
							Description: "Type a name for this tool integration, for example: my-sonarqube. This name displays on your toolchain.",
						},
						"dashboard_url": &schema.Schema{
							Type:        schema.TypeString,
							Required:    true,
							Description: "Type the URL of the SonarQube instance that you want to open when you click the SonarQube card in your toolchain.",
						},
						"user_login": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "If you are using an authentication token, leave this field empty.",
						},
						"user_password": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"blind_connection": &schema.Schema{
							Type:        schema.TypeBool,
							Optional:    true,
							Default:     false,
							Description: "Select this checkbox only if the server is not addressable on the public internet. IBM Cloud will not be able to validate the connection details you provide.",
						},
					},
				},
			},
			"parameters_references": &schema.Schema{
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "Decoded values used on provision in the broker that reference fields in the parameters.",
				Elem:        &schema.Schema{Type: schema.TypeString},
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
			"get_integration_by_id_response_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func ResourceIbmToolchainToolSonarqubeValidator() *validate.ResourceValidator {
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

	resourceValidator := validate.ResourceValidator{ResourceName: "ibm_toolchain_tool_sonarqube", Schema: validateSchema}
	return &resourceValidator
}

func ResourceIbmToolchainToolSonarqubeCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	toolchainClient, err := meta.(conns.ClientSession).ToolchainV2()
	if err != nil {
		return diag.FromErr(err)
	}

	postIntegrationOptions := &toolchainv2.PostIntegrationOptions{}

	postIntegrationOptions.SetToolchainID(d.Get("toolchain_id").(string))
	postIntegrationOptions.SetServiceID("sonarqube")
	if _, ok := d.GetOk("name"); ok {
		postIntegrationOptions.SetName(d.Get("name").(string))
	}
	if _, ok := d.GetOk("parameters"); ok {
		parametersModel, err := ResourceIbmToolchainToolSonarqubeMapToParameters(d.Get("parameters.0").(map[string]interface{}))
		if err != nil {
			return diag.FromErr(err)
		}
		postIntegrationOptions.SetParameters(parametersModel)
	}
	if _, ok := d.GetOk("parameters_references"); ok {
		// TODO: Add code to handle map container: ParametersReferences
	}

	postIntegrationResponse, response, err := toolchainClient.PostIntegrationWithContext(context, postIntegrationOptions)
	if err != nil {
		log.Printf("[DEBUG] PostIntegrationWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("PostIntegrationWithContext failed %s\n%s", err, response))
	}

	d.SetId(fmt.Sprintf("%s/%s", *postIntegrationOptions.ToolchainID, *postIntegrationResponse.ID))

	return ResourceIbmToolchainToolSonarqubeRead(context, d, meta)
}

func ResourceIbmToolchainToolSonarqubeRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	toolchainClient, err := meta.(conns.ClientSession).ToolchainV2()
	if err != nil {
		return diag.FromErr(err)
	}

	getIntegrationByIdOptions := &toolchainv2.GetIntegrationByIdOptions{}

	parts, err := flex.SepIdParts(d.Id(), "/")
	if err != nil {
		return diag.FromErr(err)
	}

	getIntegrationByIdOptions.SetToolchainID(parts[0])
	getIntegrationByIdOptions.SetIntegrationID(parts[1])

	getIntegrationByIdResponse, response, err := toolchainClient.GetIntegrationByIDWithContext(context, getIntegrationByIdOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		log.Printf("[DEBUG] GetIntegrationByIDWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("GetIntegrationByIDWithContext failed %s\n%s", err, response))
	}

	// TODO: handle argument of type map[string]interface{}
	if err = d.Set("toolchain_id", getIntegrationByIdResponse.ToolchainID); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting toolchain_id: %s", err))
	}
	if err = d.Set("name", getIntegrationByIdResponse.Name); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting name: %s", err))
	}
	if getIntegrationByIdResponse.Parameters != nil {
		parametersMap, err := ResourceIbmToolchainToolSonarqubeParametersToMap(getIntegrationByIdResponse.Parameters, d)
		if err != nil {
			return diag.FromErr(err)
		}
		if err = d.Set("parameters", []map[string]interface{}{parametersMap}); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting parameters: %s", err))
		}
	}
	if err = d.Set("resource_group_id", getIntegrationByIdResponse.ResourceGroupID); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting resource_group_id: %s", err))
	}
	if err = d.Set("crn", getIntegrationByIdResponse.Crn); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting crn: %s", err))
	}
	if err = d.Set("toolchain_crn", getIntegrationByIdResponse.ToolchainCrn); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting toolchain_crn: %s", err))
	}
	if err = d.Set("href", getIntegrationByIdResponse.Href); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting href: %s", err))
	}
	referentMap, err := ResourceIbmToolchainToolSonarqubeGetIntegrationByIdResponseReferentToMap(getIntegrationByIdResponse.Referent)
	if err != nil {
		return diag.FromErr(err)
	}
	if err = d.Set("referent", []map[string]interface{}{referentMap}); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting referent: %s", err))
	}
	if err = d.Set("updated_at", flex.DateTimeToString(getIntegrationByIdResponse.UpdatedAt)); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting updated_at: %s", err))
	}
	if err = d.Set("state", getIntegrationByIdResponse.State); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting state: %s", err))
	}
	if err = d.Set("get_integration_by_id_response_id", getIntegrationByIdResponse.ID); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting get_integration_by_id_response_id: %s", err))
	}

	return nil
}

func ResourceIbmToolchainToolSonarqubeUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
	patchToolIntegrationOptions.SetServiceID("sonarqube")

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
		parameters, err := ResourceIbmToolchainToolSonarqubeMapToParameters(d.Get("parameters.0").(map[string]interface{}))
		if err != nil {
			return diag.FromErr(err)
		}
		patchToolIntegrationOptions.SetParameters(parameters)
		hasChange = true
	}
	if d.HasChange("parameters_references") {
		// TODO: handle ParametersReferences of type TypeMap -- not primitive, not model
		hasChange = true
	}

	if hasChange {
		_, response, err := toolchainClient.PatchToolIntegrationWithContext(context, patchToolIntegrationOptions)
		if err != nil {
			log.Printf("[DEBUG] PatchToolIntegrationWithContext failed %s\n%s", err, response)
			return diag.FromErr(fmt.Errorf("PatchToolIntegrationWithContext failed %s\n%s", err, response))
		}
	}

	return ResourceIbmToolchainToolSonarqubeRead(context, d, meta)
}

func ResourceIbmToolchainToolSonarqubeDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

func ResourceIbmToolchainToolSonarqubeMapToParameters(modelMap map[string]interface{}) (map[string]interface{}, error) {
	model := make(map[string]interface{})
	model["name"] = core.StringPtr(modelMap["name"].(string))
	model["dashboard_url"] = core.StringPtr(modelMap["dashboard_url"].(string))
	if modelMap["user_login"] != nil {
		model["user_login"] = core.StringPtr(modelMap["user_login"].(string))
	}
	if modelMap["user_password"] != nil {
		model["user_password"] = core.StringPtr(modelMap["user_password"].(string))
	}
	if modelMap["blind_connection"] != nil {
		model["blind_connection"] = core.BoolPtr(modelMap["blind_connection"].(bool))
	}
	return model, nil
}

func ResourceIbmToolchainToolSonarqubeParametersToMap(model map[string]interface{}, d *schema.ResourceData) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["name"] = model["name"]
	modelMap["dashboard_url"] = model["dashboard_url"]
	if model["user_login"] != nil {
		if model["user_login"] == "****" {
			modelMap["user_login"] = d.Get("parameters.0.user_login").(string)
		} else {
			modelMap["user_login"] = model["user_login"]
		}
	}
	if model["user_password"] != nil {
		modelMap["user_password"] = model["user_password"]
	}
	if model["blind_connection"] != nil {
		modelMap["blind_connection"] = model["blind_connection"]
	}
	return modelMap, nil
}

func ResourceIbmToolchainToolSonarqubeGetIntegrationByIdResponseReferentToMap(model *toolchainv2.GetIntegrationByIdResponseReferent) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.UiHref != nil {
		modelMap["ui_href"] = model.UiHref
	}
	if model.ApiHref != nil {
		modelMap["api_href"] = model.ApiHref
	}
	return modelMap, nil
}
