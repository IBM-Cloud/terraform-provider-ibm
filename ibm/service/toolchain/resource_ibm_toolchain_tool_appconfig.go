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

func ResourceIBMToolchainToolAppconfig() *schema.Resource {
	return &schema.Resource{
		CreateContext: ResourceIBMToolchainToolAppconfigCreate,
		ReadContext:   ResourceIBMToolchainToolAppconfigRead,
		UpdateContext: ResourceIBMToolchainToolAppconfigUpdate,
		DeleteContext: ResourceIBMToolchainToolAppconfigDelete,
		Importer:      &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"toolchain_id": &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validate.InvokeValidator("ibm_toolchain_tool_appconfig", "toolchain_id"),
				Description:  "ID of the toolchain to bind integration to.",
			},
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Name of tool integration.",
			},
			"parameters": &schema.Schema{
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Description: "Parameters to be used to create the integration.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:        schema.TypeString,
							Required:    true,
							Description: "Type a name for this tool integration, for example: my-appconfig. This name displays on your toolchain.",
						},
						"region": &schema.Schema{
							Type:        schema.TypeString,
							Required:    true,
							Description: "Region.",
						},
						"resource_group": &schema.Schema{
							Type:        schema.TypeString,
							Required:    true,
							Description: "Resource group.",
						},
						"instance_name": &schema.Schema{
							Type:        schema.TypeString,
							Required:    true,
							Description: "The name of your App Configuration instance. You should choose an entry from the list provided based on the selected region and resource group. e.g: App Configuration-01.",
						},
						"environment_name": &schema.Schema{
							Type:        schema.TypeString,
							Required:    true,
							Description: "App Configuration environment.",
						},
						"collection_name": &schema.Schema{
							Type:        schema.TypeString,
							Required:    true,
							Description: "App Configuration collection.",
						},
						"integration_status": &schema.Schema{
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
			"instance_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func ResourceIBMToolchainToolAppconfigValidator() *validate.ResourceValidator {
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

	resourceValidator := validate.ResourceValidator{ResourceName: "ibm_toolchain_tool_appconfig", Schema: validateSchema}
	return &resourceValidator
}

func ResourceIBMToolchainToolAppconfigCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	toolchainClient, err := meta.(conns.ClientSession).ToolchainV2()
	if err != nil {
		return diag.FromErr(err)
	}

	postIntegrationOptions := &toolchainv2.PostIntegrationOptions{}

	postIntegrationOptions.SetToolchainID(d.Get("toolchain_id").(string))
	postIntegrationOptions.SetToolID("appconfig")
	if _, ok := d.GetOk("name"); ok {
		postIntegrationOptions.SetName(d.Get("name").(string))
	}
	if _, ok := d.GetOk("parameters"); ok {
		remapFields := map[string]string{
			"resource_group":     "resource-group",
			"instance_name":      "instance-name",
			"environment_name":   "environment-name",
			"collection_name":    "collection-name",
			"integration_status": "integration-status",
		}
		parametersModel := GetParametersForCreate(d, ResourceIBMToolchainToolAppconfig(), remapFields)
		postIntegrationOptions.SetParameters(parametersModel)
	}

	postIntegrationResponse, response, err := toolchainClient.PostIntegrationWithContext(context, postIntegrationOptions)
	if err != nil {
		log.Printf("[DEBUG] PostIntegrationWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("PostIntegrationWithContext failed %s\n%s", err, response))
	}

	d.SetId(fmt.Sprintf("%s/%s", *postIntegrationOptions.ToolchainID, *postIntegrationResponse.ID))

	return ResourceIBMToolchainToolAppconfigRead(context, d, meta)
}

func ResourceIBMToolchainToolAppconfigRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
		remapFields := map[string]string{
			"resource_group":     "resource-group",
			"instance_name":      "instance-name",
			"environment_name":   "environment-name",
			"collection_name":    "collection-name",
			"integration_status": "integration-status",
		}
		parametersMap := GetParametersFromRead(getIntegrationByIDResponse.Parameters, ResourceIBMToolchainToolAppconfig(), remapFields)
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
	referentMap, err := ResourceIBMToolchainToolAppconfigGetIntegrationByIDResponseReferentToMap(getIntegrationByIDResponse.Referent)
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
	if err = d.Set("instance_id", getIntegrationByIDResponse.ID); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting instance_id: %s", err))
	}

	return nil
}

func ResourceIBMToolchainToolAppconfigUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
	patchToolIntegrationOptions.SetToolID("appconfig")

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
		remapFields := map[string]string{
			"resource_group":     "resource-group",
			"instance_name":      "instance-name",
			"environment_name":   "environment-name",
			"collection_name":    "collection-name",
			"integration_status": "integration-status",
		}
		parameters := GetParametersForUpdate(d, ResourceIBMToolchainToolAppconfig(), remapFields)
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

	return ResourceIBMToolchainToolAppconfigRead(context, d, meta)
}

func ResourceIBMToolchainToolAppconfigDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

func ResourceIBMToolchainToolAppconfigGetIntegrationByIDResponseReferentToMap(model *toolchainv2.GetIntegrationByIDResponseReferent) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.UIHref != nil {
		modelMap["ui_href"] = model.UIHref
	}
	if model.APIHref != nil {
		modelMap["api_href"] = model.APIHref
	}
	return modelMap, nil
}
