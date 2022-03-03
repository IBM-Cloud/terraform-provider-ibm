// Copyright IBM Corp. 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibmtoolchainapi

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.ibm.com/org-ids/toolchain-go-sdk/ibmtoolchainapiv2"
)

func ResourceIbmToolchainToolPipeline() *schema.Resource {
	return &schema.Resource{
		CreateContext: ResourceIbmToolchainToolPipelineCreate,
		ReadContext:   ResourceIbmToolchainToolPipelineRead,
		UpdateContext: ResourceIbmToolchainToolPipelineUpdate,
		DeleteContext: ResourceIbmToolchainToolPipelineDelete,
		Importer:      &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"toolchain_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"parameters": &schema.Schema{
				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
						"type": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"ui_pipeline": &schema.Schema{
							Type:        schema.TypeBool,
							Optional:    true,
							Default:     false,
							Description: "When this check box is selected, the applications that this pipeline deploys are shown in the View app menu on the toolchain page. This setting is best for UI apps that can be accessed from a browser.",
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
			"container": &schema.Schema{
				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"guid": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
						"type": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
			"dashboard_url": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The URL of a user-facing user interface for this instance of a service.",
			},
		},
	}
}

func ResourceIbmToolchainToolPipelineCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	ibmToolchainApiClient, err := meta.(conns.ClientSession).IbmToolchainApiV2()
	if err != nil {
		return diag.FromErr(err)
	}

	createServiceInstanceOptions := &ibmtoolchainapiv2.CreateServiceInstanceOptions{}

	createServiceInstanceOptions.SetServiceID("pipeline")
	createServiceInstanceOptions.SetToolchainID(d.Get("toolchain_id").(string))
	if _, ok := d.GetOk("parameters"); ok {
		parameters, err := ResourceIbmToolchainToolPipelineMapToParameters(d.Get("parameters.0").(map[string]interface{}))
		if err != nil {
			return diag.FromErr(err)
		}
		createServiceInstanceOptions.SetParameters(parameters)
	}
	if _, ok := d.GetOk("parameters_references"); ok {
		// TODO: Add code to handle map container: ParametersReferences
	}
	if _, ok := d.GetOk("container"); ok {
		container, err := ResourceIbmToolchainToolPipelineMapToContainer(d.Get("container.0").(map[string]interface{}))
		if err != nil {
			return diag.FromErr(err)
		}
		createServiceInstanceOptions.SetContainer(container)
	}

	serviceResponse, response, err := ibmToolchainApiClient.CreateServiceInstanceWithContext(context, createServiceInstanceOptions)
	if err != nil {
		log.Printf("[DEBUG] CreateServiceInstanceWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("CreateServiceInstanceWithContext failed %s\n%s", err, response))
	}

	d.SetId(*serviceResponse.InstanceID)

	return ResourceIbmToolchainToolPipelineRead(context, d, meta)
}

func ResourceIbmToolchainToolPipelineRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	ibmToolchainApiClient, err := meta.(conns.ClientSession).IbmToolchainApiV2()
	if err != nil {
		return diag.FromErr(err)
	}

	getServiceInstanceOptions := &ibmtoolchainapiv2.GetServiceInstanceOptions{}

	getServiceInstanceOptions.SetServiceInstanceID(d.Id())

	serviceResponse, response, err := ibmToolchainApiClient.GetServiceInstanceWithContext(context, getServiceInstanceOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		log.Printf("[DEBUG] GetServiceInstanceWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("GetServiceInstanceWithContext failed %s\n%s", err, response))
	}

	// TODO: handle argument of type map[string]interface{}
	if err = d.Set("toolchain_id", serviceResponse.ToolchainID); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting toolchain_id: %s", err))
	}
	if serviceResponse.Parameters != nil {
		parametersMap, err := ResourceIbmToolchainToolPipelineParametersToMap(serviceResponse.Parameters)
		if err != nil {
			return diag.FromErr(err)
		}
		if err = d.Set("parameters", []map[string]interface{}{parametersMap}); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting parameters: %s", err))
		}
	}
	if serviceResponse.Container != nil {
		containerMap, err := ResourceIbmToolchainToolPipelineContainerToMap(serviceResponse.Container)
		if err != nil {
			return diag.FromErr(err)
		}
		if err = d.Set("container", []map[string]interface{}{containerMap}); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting container: %s", err))
		}
	}
	if err = d.Set("dashboard_url", serviceResponse.DashboardURL); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting dashboard_url: %s", err))
	}

	return nil
}

func ResourceIbmToolchainToolPipelineUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	ibmToolchainApiClient, err := meta.(conns.ClientSession).IbmToolchainApiV2()
	if err != nil {
		return diag.FromErr(err)
	}

	patchServiceInstanceOptions := &ibmtoolchainapiv2.PatchServiceInstanceOptions{}

	patchServiceInstanceOptions.SetServiceID("pipeline")
	patchServiceInstanceOptions.SetServiceInstanceID(d.Id())

	hasChange := false

	if d.HasChange("toolchain_id") {
		return diag.FromErr(fmt.Errorf("Cannot update resource property \"%s\" with the ForceNew annotation."+
			" The resource must be re-created to update this property.", "toolchain_id"))
	}
	if d.HasChange("container") {
		return diag.FromErr(fmt.Errorf("Cannot update resource property \"%s\" with the ForceNew annotation."+
			" The resource must be re-created to update this property.", "container"))
	}
	if d.HasChange("parameters") {
		parameters, err := ResourceIbmToolchainToolPipelineMapToParameters(d.Get("parameters.0").(map[string]interface{}))
		if err != nil {
			return diag.FromErr(err)
		}
		patchServiceInstanceOptions.SetParameters(parameters)
		hasChange = true
	}
	if d.HasChange("parameters_references") {
		// TODO: handle ParametersReferences of type TypeMap -- not primitive, not model
		hasChange = true
	}

	if hasChange {
		response, err := ibmToolchainApiClient.PatchServiceInstanceWithContext(context, patchServiceInstanceOptions)
		if err != nil {
			log.Printf("[DEBUG] PatchServiceInstanceWithContext failed %s\n%s", err, response)
			return diag.FromErr(fmt.Errorf("PatchServiceInstanceWithContext failed %s\n%s", err, response))
		}
	}

	return ResourceIbmToolchainToolPipelineRead(context, d, meta)
}

func ResourceIbmToolchainToolPipelineDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	ibmToolchainApiClient, err := meta.(conns.ClientSession).IbmToolchainApiV2()
	if err != nil {
		return diag.FromErr(err)
	}

	deleteServiceInstanceOptions := &ibmtoolchainapiv2.DeleteServiceInstanceOptions{}

	deleteServiceInstanceOptions.SetServiceInstanceID(d.Id())

	response, err := ibmToolchainApiClient.DeleteServiceInstanceWithContext(context, deleteServiceInstanceOptions)
	if err != nil {
		log.Printf("[DEBUG] DeleteServiceInstanceWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("DeleteServiceInstanceWithContext failed %s\n%s", err, response))
	}

	d.SetId("")

	return nil
}

func ResourceIbmToolchainToolPipelineMapToParameters(modelMap map[string]interface{}) (map[string]interface{}, error) {
	model := make(map[string]interface{})
	model["name"] = core.StringPtr(modelMap["name"].(string))
	model["type"] = core.StringPtr(modelMap["type"].(string))
	if modelMap["ui_pipeline"] != nil {
		model["ui_pipeline"] = core.BoolPtr(modelMap["ui_pipeline"].(bool))
	}
	return model, nil
}

func ResourceIbmToolchainToolPipelineMapToContainer(modelMap map[string]interface{}) (*ibmtoolchainapiv2.Container, error) {
	model := &ibmtoolchainapiv2.Container{}
	model.Guid = core.StringPtr(modelMap["guid"].(string))
	model.Type = core.StringPtr(modelMap["type"].(string))
	return model, nil
}

func ResourceIbmToolchainToolPipelineParametersToMap(model map[string]interface{}) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["name"] = model["name"]
	modelMap["type"] = model["type"]
	if model["ui_pipeline"] != nil {
		modelMap["ui_pipeline"] = model["ui_pipeline"]
	}
	return modelMap, nil
}

func ResourceIbmToolchainToolPipelineContainerToMap(model *ibmtoolchainapiv2.Container) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["guid"] = model.Guid
	modelMap["type"] = model.Type
	return modelMap, nil
}
