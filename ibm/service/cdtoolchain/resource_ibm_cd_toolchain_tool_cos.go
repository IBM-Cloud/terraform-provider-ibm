// Copyright IBM Corp. 2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.96.0-d6dec9d7-20241008-212902
 */

package cdtoolchain

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/validate"
	"github.com/IBM/continuous-delivery-go-sdk/v2/cdtoolchainv2"
	"github.com/IBM/go-sdk-core/v5/core"
)

func ResourceIBMCdToolchainToolCos() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIBMCdToolchainToolCosCreate,
		ReadContext:   resourceIBMCdToolchainToolCosRead,
		UpdateContext: resourceIBMCdToolchainToolCosUpdate,
		DeleteContext: resourceIBMCdToolchainToolCosDelete,
		Importer:      &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"toolchain_id": &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validate.InvokeValidator("ibm_cd_toolchain_tool_cos", "toolchain_id"),
				Description:  "ID of the toolchain to bind the tool to.",
			},
			"name": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validate.InvokeValidator("ibm_cd_toolchain_tool_cos", "name"),
				Description:  "Name of the tool.",
			},
			"parameters": &schema.Schema{
				Type:        schema.TypeList,
				MinItems:    1,
				MaxItems:    1,
				Required:    true,
				Description: "Unique key-value pairs representing parameters to be used to create the tool. A list of parameters for each tool integration can be found in the <a href=\"https://cloud.ibm.com/docs/ContinuousDelivery?topic=ContinuousDelivery-integrations\">Configuring tool integrations page</a>.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:        schema.TypeString,
							Required:    true,
							Description: "The name used to identify this tool integration.",
						},
						"auth_type": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "The authentication type. Options are `apikey` IBM Cloud API Key or `hmac` HMAC (Hash Message Authentication Code). The default is `apikey`.",
						},
						"cos_api_key": &schema.Schema{
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: flex.SuppressHashedRawSecret,
							Sensitive:        true,
							Description:      "The IBM Cloud API key used to access the Cloud Object Storage service. Only relevant when using `apikey` as the `auth_type`.",
						},
						"instance_crn": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The CRN (Cloud Resource Name) of the IBM Cloud Object Storage service instance, only relevant when using `apikey` as the `auth_type`.",
						},
						"bucket_name": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The name of the Cloud Object Storage service bucket.",
						},
						"endpoint": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The [Cloud Object Storage endpoint](https://cloud.ibm.com/docs/cloud-object-storage?topic=cloud-object-storage-endpoints) in IBM Cloud or other endpoint. For example for IBM Cloud Object Storage: `s3.direct.us-south.cloud-object-storage.appdomain.cloud`.",
						},
						"hmac_access_key_id": &schema.Schema{
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: flex.SuppressHashedRawSecret,
							Sensitive:        true,
							Description:      "The HMAC Access Key ID which is part of an HMAC (Hash Message Authentication Code) credential set. HMAC is identified by a combination of an Access Key ID and a Secret Access Key. Only relevant when `auth_type` is set to `hmac`.",
						},
						"hmac_secret_access_key": &schema.Schema{
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: flex.SuppressHashedRawSecret,
							Sensitive:        true,
							Description:      "The HMAC Secret Access Key which is part of an HMAC (Hash Message Authentication Code) credential set. HMAC is identified by a combination of an Access Key ID and a Secret Access Key. Only relevant when `auth_type` is set to `hmac`.",
						},
					},
				},
			},
			"resource_group_id": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Resource group where the tool is located.",
			},
			"crn": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Tool CRN.",
			},
			"toolchain_crn": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "CRN of toolchain which the tool is bound to.",
			},
			"href": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "URI representing the tool.",
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
							Computed:    true,
							Description: "URI representing this resource through the UI.",
						},
						"api_href": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "URI representing this resource through an API.",
						},
					},
				},
			},
			"updated_at": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Latest tool update timestamp.",
			},
			"state": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Current configuration state of the tool.",
			},
			"tool_id": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Tool ID.",
			},
		},
	}
}

func ResourceIBMCdToolchainToolCosValidator() *validate.ResourceValidator {
	validateSchema := make([]validate.ValidateSchema, 0)
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
			Regexp:                     `^([^\x00-\x7F]|[a-zA-Z0-9-._ ])+$`,
			MinValueLength:             0,
			MaxValueLength:             128,
		},
	)

	resourceValidator := validate.ResourceValidator{ResourceName: "ibm_cd_toolchain_tool_cos", Schema: validateSchema}
	return &resourceValidator
}

func resourceIBMCdToolchainToolCosCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	cdToolchainClient, err := meta.(conns.ClientSession).CdToolchainV2()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_cd_toolchain_tool_cos", "create", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	createToolOptions := &cdtoolchainv2.CreateToolOptions{}

	createToolOptions.SetToolchainID(d.Get("toolchain_id").(string))
	createToolOptions.SetToolTypeID("cloudobjectstorage")
	parametersModel := GetParametersForCreate(d, ResourceIBMCdToolchainToolCos(), nil)
	createToolOptions.SetParameters(parametersModel)
	if _, ok := d.GetOk("name"); ok {
		createToolOptions.SetName(d.Get("name").(string))
	}

	toolchainToolPost, _, err := cdToolchainClient.CreateToolWithContext(context, createToolOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("CreateToolWithContext failed: %s", err.Error()), "ibm_cd_toolchain_tool_cos", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(fmt.Sprintf("%s/%s", *createToolOptions.ToolchainID, *toolchainToolPost.ID))

	return resourceIBMCdToolchainToolCosRead(context, d, meta)
}

func resourceIBMCdToolchainToolCosRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	cdToolchainClient, err := meta.(conns.ClientSession).CdToolchainV2()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_cd_toolchain_tool_cos", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	getToolByIDOptions := &cdtoolchainv2.GetToolByIDOptions{}

	parts, err := flex.SepIdParts(d.Id(), "/")
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_cd_toolchain_tool_cos", "read", "sep-id-parts").GetDiag()
	}

	getToolByIDOptions.SetToolchainID(parts[0])
	getToolByIDOptions.SetToolID(parts[1])

	var toolchainTool *cdtoolchainv2.ToolchainTool
	var response *core.DetailedResponse
	err = resource.RetryContext(context, 10*time.Second, func() *resource.RetryError {
		toolchainTool, response, err = cdToolchainClient.GetToolByIDWithContext(context, getToolByIDOptions)
		if err != nil || toolchainTool == nil {
			if response != nil && response.StatusCode == 404 {
				return resource.RetryableError(err)
			}
			return resource.NonRetryableError(err)
		}
		return nil
	})
	if conns.IsResourceTimeoutError(err) {
		toolchainTool, response, err = cdToolchainClient.GetToolByIDWithContext(context, getToolByIDOptions)
	}
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetToolByIDWithContext failed: %s", err.Error()), "ibm_cd_toolchain_tool_cos", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	if err = d.Set("toolchain_id", toolchainTool.ToolchainID); err != nil {
		err = fmt.Errorf("Error setting toolchain_id: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_cd_toolchain_tool_cos", "read", "set-toolchain_id").GetDiag()
	}
	if !core.IsNil(toolchainTool.Name) {
		if err = d.Set("name", toolchainTool.Name); err != nil {
			err = fmt.Errorf("Error setting name: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_cd_toolchain_tool_cos", "read", "set-name").GetDiag()
		}
	}
	parametersMap := GetParametersFromRead(toolchainTool.Parameters, ResourceIBMCdToolchainToolCos(), nil)
	if err = d.Set("parameters", []map[string]interface{}{parametersMap}); err != nil {
		err = fmt.Errorf("Error setting parameters: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_cd_toolchain_tool_cos", "read", "set-parameters").GetDiag()
	}
	if err = d.Set("resource_group_id", toolchainTool.ResourceGroupID); err != nil {
		err = fmt.Errorf("Error setting resource_group_id: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_cd_toolchain_tool_cos", "read", "set-resource_group_id").GetDiag()
	}
	if err = d.Set("crn", toolchainTool.CRN); err != nil {
		err = fmt.Errorf("Error setting crn: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_cd_toolchain_tool_cos", "read", "set-crn").GetDiag()
	}
	if err = d.Set("toolchain_crn", toolchainTool.ToolchainCRN); err != nil {
		err = fmt.Errorf("Error setting toolchain_crn: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_cd_toolchain_tool_cos", "read", "set-toolchain_crn").GetDiag()
	}
	if err = d.Set("href", toolchainTool.Href); err != nil {
		err = fmt.Errorf("Error setting href: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_cd_toolchain_tool_cos", "read", "set-href").GetDiag()
	}
	referentMap, err := ResourceIBMCdToolchainToolCosToolModelReferentToMap(toolchainTool.Referent)
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_cd_toolchain_tool_cos", "read", "referent-to-map").GetDiag()
	}
	if err = d.Set("referent", []map[string]interface{}{referentMap}); err != nil {
		err = fmt.Errorf("Error setting referent: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_cd_toolchain_tool_cos", "read", "set-referent").GetDiag()
	}
	if err = d.Set("updated_at", flex.DateTimeToString(toolchainTool.UpdatedAt)); err != nil {
		err = fmt.Errorf("Error setting updated_at: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_cd_toolchain_tool_cos", "read", "set-updated_at").GetDiag()
	}
	if err = d.Set("state", toolchainTool.State); err != nil {
		err = fmt.Errorf("Error setting state: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_cd_toolchain_tool_cos", "read", "set-state").GetDiag()
	}
	if err = d.Set("tool_id", toolchainTool.ID); err != nil {
		err = fmt.Errorf("Error setting tool_id: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_cd_toolchain_tool_cos", "read", "set-tool_id").GetDiag()
	}

	return nil
}

func resourceIBMCdToolchainToolCosUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	cdToolchainClient, err := meta.(conns.ClientSession).CdToolchainV2()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_cd_toolchain_tool_cos", "update", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	updateToolOptions := &cdtoolchainv2.UpdateToolOptions{}

	parts, err := flex.SepIdParts(d.Id(), "/")
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_cd_toolchain_tool_cos", "update", "sep-id-parts").GetDiag()
	}

	updateToolOptions.SetToolchainID(parts[0])
	updateToolOptions.SetToolID(parts[1])

	hasChange := false

	patchVals := &cdtoolchainv2.ToolchainToolPrototypePatch{}
	if d.HasChange("toolchain_id") {
		errMsg := fmt.Sprintf("Cannot update resource property \"%s\" with the ForceNew annotation."+
			" The resource must be re-created to update this property.", "toolchain_id")
		return flex.DiscriminatedTerraformErrorf(nil, errMsg, "ibm_cd_toolchain_tool_cos", "update", "toolchain_id-forces-new").GetDiag()
	}
	if d.HasChange("name") {
		newName := d.Get("name").(string)
		patchVals.Name = &newName
		hasChange = true
	}
	if d.HasChange("parameters") {
		parameters := GetParametersForUpdate(d, ResourceIBMCdToolchainToolCos(), nil)
		patchVals.Parameters = parameters
		hasChange = true
	}

	if hasChange {
		// Fields with `nil` values are omitted from the generic map,
		// so we need to re-add them to support removing arguments
		// in merge-patch operations sent to the service.
		updateToolOptions.ToolchainToolPrototypePatch = ResourceIBMCdToolchainToolCosToolchainToolPrototypePatchAsPatch(patchVals, d)

		_, _, err = cdToolchainClient.UpdateToolWithContext(context, updateToolOptions)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("UpdateToolWithContext failed: %s", err.Error()), "ibm_cd_toolchain_tool_cos", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
	}

	return resourceIBMCdToolchainToolCosRead(context, d, meta)
}

func resourceIBMCdToolchainToolCosDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	cdToolchainClient, err := meta.(conns.ClientSession).CdToolchainV2()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_cd_toolchain_tool_cos", "delete", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	deleteToolOptions := &cdtoolchainv2.DeleteToolOptions{}

	parts, err := flex.SepIdParts(d.Id(), "/")
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_cd_toolchain_tool_cos", "delete", "sep-id-parts").GetDiag()
	}

	deleteToolOptions.SetToolchainID(parts[0])
	deleteToolOptions.SetToolID(parts[1])

	_, err = cdToolchainClient.DeleteToolWithContext(context, deleteToolOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("DeleteToolWithContext failed: %s", err.Error()), "ibm_cd_toolchain_tool_cos", "delete")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId("")

	return nil
}

func ResourceIBMCdToolchainToolCosToolModelReferentToMap(model *cdtoolchainv2.ToolModelReferent) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.UIHref != nil {
		modelMap["ui_href"] = *model.UIHref
	}
	if model.APIHref != nil {
		modelMap["api_href"] = *model.APIHref
	}
	return modelMap, nil
}

func ResourceIBMCdToolchainToolCosToolchainToolPrototypePatchAsPatch(patchVals *cdtoolchainv2.ToolchainToolPrototypePatch, d *schema.ResourceData) map[string]interface{} {
	patch, _ := patchVals.AsPatch()
	var path string

	path = "name"
	if _, exists := d.GetOk(path); d.HasChange(path) && !exists {
		patch["name"] = nil
	}
	path = "tool_type_id"
	if _, exists := d.GetOk(path); d.HasChange(path) && !exists {
		patch["tool_type_id"] = nil
	}
	path = "parameters"
	if _, exists := d.GetOk(path); d.HasChange(path) && !exists {
		patch["parameters"] = nil
	} else if exists && patch["parameters"] != nil {
		ResourceIBMCdToolchainToolCosToolModelParametersAsPatch(patch["parameters"].(map[string]interface{}), d)
	}

	return patch
}

func ResourceIBMCdToolchainToolCosToolModelParametersAsPatch(patch map[string]interface{}, d *schema.ResourceData) {
	var path string

	path = "parameters.0.auth_type"
	if _, exists := d.GetOk(path); d.HasChange(path) && !exists {
		patch["auth_type"] = nil
	}
	path = "parameters.0.cos_api_key"
	if _, exists := d.GetOk(path); d.HasChange(path) && !exists {
		patch["cos_api_key"] = nil
	}
	path = "parameters.0.instance_crn"
	if _, exists := d.GetOk(path); d.HasChange(path) && !exists {
		patch["instance_crn"] = nil
	}
	path = "parameters.0.bucket_name"
	if _, exists := d.GetOk(path); d.HasChange(path) && !exists {
		patch["bucket_name"] = nil
	}
	path = "parameters.0.endpoint"
	if _, exists := d.GetOk(path); d.HasChange(path) && !exists {
		patch["endpoint"] = nil
	}
	path = "parameters.0.hmac_access_key_id"
	if _, exists := d.GetOk(path); d.HasChange(path) && !exists {
		patch["hmac_access_key_id"] = nil
	}
	path = "parameters.0.hmac_secret_access_key"
	if _, exists := d.GetOk(path); d.HasChange(path) && !exists {
		patch["hmac_secret_access_key"] = nil
	}
}
