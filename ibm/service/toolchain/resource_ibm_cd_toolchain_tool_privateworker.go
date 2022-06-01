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

func ResourceIBMCdToolchainToolPrivateworker() *schema.Resource {
	return &schema.Resource{
		CreateContext: ResourceIBMCdToolchainToolPrivateworkerCreate,
		ReadContext:   ResourceIBMCdToolchainToolPrivateworkerRead,
		UpdateContext: ResourceIBMCdToolchainToolPrivateworkerUpdate,
		DeleteContext: ResourceIBMCdToolchainToolPrivateworkerDelete,
		Importer:      &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"toolchain_id": &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validate.InvokeValidator("ibm_cd_toolchain_tool_privateworker", "toolchain_id"),
				Description:  "ID of the toolchain to bind integration to.",
			},
			"name": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validate.InvokeValidator("ibm_cd_toolchain_tool_privateworker", "name"),
				Description:  "Name of tool integration.",
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
							Description: "Enter a name for this tool integration. For example, my-private-worker. This name is displayed on your toolchain.",
						},
						"worker_queue_credentials": &schema.Schema{
							Type:             schema.TypeString,
							Required:         true,
							DiffSuppressFunc: flex.SuppressHashedRawSecret,
							Sensitive:        true,
							Description:      "Use a secret from the secrets store, or create a service ID API key that is used by the private worker to authenticate access to the work queue.",
						},
						"worker_queue_identifier": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
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

func ResourceIBMCdToolchainToolPrivateworkerValidator() *validate.ResourceValidator {
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

	resourceValidator := validate.ResourceValidator{ResourceName: "ibm_cd_toolchain_tool_privateworker", Schema: validateSchema}
	return &resourceValidator
}

func ResourceIBMCdToolchainToolPrivateworkerCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	toolchainClient, err := meta.(conns.ClientSession).ToolchainV2()
	if err != nil {
		return diag.FromErr(err)
	}

	createIntegrationOptions := &toolchainv2.CreateIntegrationOptions{}

	createIntegrationOptions.SetToolchainID(d.Get("toolchain_id").(string))
	createIntegrationOptions.SetToolID("private_worker")
	if _, ok := d.GetOk("name"); ok {
		createIntegrationOptions.SetName(d.Get("name").(string))
	}
	if _, ok := d.GetOk("parameters"); ok {
		remapFields := map[string]string{
			"worker_queue_credentials": "workerQueueCredentials",
			"worker_queue_identifier":  "workerQueueIdentifier",
		}
		parametersModel := GetParametersForCreate(d, ResourceIBMCdToolchainToolPrivateworker(), remapFields)
		createIntegrationOptions.SetParameters(parametersModel)
	}

	postIntegrationResponse, response, err := toolchainClient.CreateIntegrationWithContext(context, createIntegrationOptions)
	if err != nil {
		log.Printf("[DEBUG] CreateIntegrationWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("CreateIntegrationWithContext failed %s\n%s", err, response))
	}

	d.SetId(fmt.Sprintf("%s/%s", *createIntegrationOptions.ToolchainID, *postIntegrationResponse.ID))

	return ResourceIBMCdToolchainToolPrivateworkerRead(context, d, meta)
}

func ResourceIBMCdToolchainToolPrivateworkerRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
			"worker_queue_credentials": "workerQueueCredentials",
			"worker_queue_identifier":  "workerQueueIdentifier",
		}
		parametersMap := GetParametersFromRead(getIntegrationByIDResponse.Parameters, ResourceIBMCdToolchainToolPrivateworker(), remapFields)
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
	referentMap, err := ResourceIBMCdToolchainToolPrivateworkerToolIntegrationReferentToMap(getIntegrationByIDResponse.Referent)
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

func ResourceIBMCdToolchainToolPrivateworkerUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
	updateIntegrationOptions.SetToolID("private_worker")

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
		remapFields := map[string]string{
			"worker_queue_credentials": "workerQueueCredentials",
			"worker_queue_identifier":  "workerQueueIdentifier",
		}
		parameters := GetParametersForUpdate(d, ResourceIBMCdToolchainToolPrivateworker(), remapFields)
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

	return ResourceIBMCdToolchainToolPrivateworkerRead(context, d, meta)
}

func ResourceIBMCdToolchainToolPrivateworkerDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

func ResourceIBMCdToolchainToolPrivateworkerToolIntegrationReferentToMap(model *toolchainv2.ToolIntegrationReferent) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.UIHref != nil {
		modelMap["ui_href"] = model.UIHref
	}
	if model.APIHref != nil {
		modelMap["api_href"] = model.APIHref
	}
	return modelMap, nil
}
