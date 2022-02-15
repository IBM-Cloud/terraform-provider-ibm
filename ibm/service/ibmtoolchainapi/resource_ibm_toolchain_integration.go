// Copyright IBM Corp. 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibmtoolchainapi

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.ibm.com/org-ids/toolchain-go-sdk/ibmtoolchainapiv2"
)

func ResourceIbmToolchainIntegration() *schema.Resource {
	return &schema.Resource{
		CreateContext: ResourceIbmToolchainIntegrationCreate,
		ReadContext:   ResourceIbmToolchainIntegrationRead,
		UpdateContext: ResourceIbmToolchainIntegrationUpdate,
		DeleteContext: ResourceIbmToolchainIntegrationDelete,
		Importer:      &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"service_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The unique short name of the service that should be provisioned.",
			},
			"toolchain_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"parameters": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
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

func ResourceIbmToolchainIntegrationCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	ibmToolchainApiClient, err := meta.(conns.ClientSession).IbmToolchainApiV2()
	if err != nil {
		return diag.FromErr(err)
	}

	createServiceInstanceOptions := &ibmtoolchainapiv2.CreateServiceInstanceOptions{}

	createServiceInstanceOptions.SetServiceID(d.Get("service_id").(string))
	createServiceInstanceOptions.SetToolchainID(d.Get("toolchain_id").(string))
	if s, ok := d.GetOk("parameters"); ok {
		params := map[string]interface{}{}
		json.Unmarshal([]byte(s.(string)), &params)
		createServiceInstanceOptions.SetParameters(params)
	}
	if _, ok := d.GetOk("parameters_references"); ok {
		// TODO: Add code to handle map container: ParametersReferences
	}
	if _, ok := d.GetOk("container"); ok {
		container, err := ResourceIbmToolchainIntegrationMapToContainer(d.Get("container.0").(map[string]interface{}))
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

	return ResourceIbmToolchainIntegrationRead(context, d, meta)
}

func ResourceIbmToolchainIntegrationRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
	if err = d.Set("service_id", serviceResponse.ServiceID); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting service_id: %s", err))
	}
	if err = d.Set("toolchain_id", serviceResponse.ToolchainID); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting toolchain_id: %s", err))
	}

	if serviceResponse.Parameters != nil {
		oldParamStr := d.Get("parameters")
		params := map[string]interface{}{}
		json.Unmarshal([]byte(oldParamStr.(string)), &params)
		newMap := make(map[string]interface{})
		for key, _ := range params {
			newMap[key] = serviceResponse.Parameters[key]
		}
		buf, _ := json.Marshal(newMap)
		if err = d.Set("parameters", string(buf)); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting parameters: %s", err))
		}
	}

	if serviceResponse.Container != nil {
		containerMap, err := ResourceIbmToolchainIntegrationContainerToMap(serviceResponse.Container)
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

func ResourceIbmToolchainIntegrationUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	ibmToolchainApiClient, err := meta.(conns.ClientSession).IbmToolchainApiV2()
	if err != nil {
		return diag.FromErr(err)
	}

	patchServiceInstanceOptions := &ibmtoolchainapiv2.PatchServiceInstanceOptions{}

	patchServiceInstanceOptions.SetServiceInstanceID(d.Id())

	hasChange := false

	if d.HasChange("service_id") {
		return diag.FromErr(fmt.Errorf("Cannot update resource property \"%s\" with the ForceNew annotation."+
			" The resource must be re-created to update this property.", "service_id"))
	}
	if d.HasChange("toolchain_id") {
		return diag.FromErr(fmt.Errorf("Cannot update resource property \"%s\" with the ForceNew annotation."+
			" The resource must be re-created to update this property.", "toolchain_id"))
	}
	if d.HasChange("container") {
		return diag.FromErr(fmt.Errorf("Cannot update resource property \"%s\" with the ForceNew annotation."+
			" The resource must be re-created to update this property.", "container"))
	}
	if d.HasChange("parameters") {
		if s, ok := d.GetOk("parameters"); ok {
			params := map[string]interface{}{}
			json.Unmarshal([]byte(s.(string)), &params)
			patchServiceInstanceOptions.SetParameters(params)
			hasChange = true
		}
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

	return ResourceIbmToolchainIntegrationRead(context, d, meta)
}

func ResourceIbmToolchainIntegrationDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

func ResourceIbmToolchainIntegrationMapToContainer(modelMap map[string]interface{}) (*ibmtoolchainapiv2.Container, error) {
	model := &ibmtoolchainapiv2.Container{}
	model.Guid = core.StringPtr(modelMap["guid"].(string))
	model.Type = core.StringPtr(modelMap["type"].(string))
	return model, nil
}

func ResourceIbmToolchainIntegrationContainerToMap(model *ibmtoolchainapiv2.Container) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["guid"] = model.Guid
	modelMap["type"] = model.Type
	return modelMap, nil
}
