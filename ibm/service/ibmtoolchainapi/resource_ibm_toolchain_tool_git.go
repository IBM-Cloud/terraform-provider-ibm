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

func ResourceIbmToolchainToolGit() *schema.Resource {
	return &schema.Resource{
		CreateContext: ResourceIbmToolchainToolGitCreate,
		ReadContext:   ResourceIbmToolchainToolGitRead,
		UpdateContext: ResourceIbmToolchainToolGitUpdate,
		DeleteContext: ResourceIbmToolchainToolGitDelete,
		Importer:      &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"toolchain_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"git_provider": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"parameters": &schema.Schema{
				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"repo_url": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"action": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
						"legal": &schema.Schema{
							Type:     schema.TypeBool,
							Optional: true,
						},
						"enable_traceability": &schema.Schema{
							Type:     schema.TypeBool,
							Optional: true,
						},
						"has_issues": &schema.Schema{
							Type:     schema.TypeBool,
							Optional: true,
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

func ResourceIbmToolchainToolGitCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	ibmToolchainApiClient, err := meta.(conns.ClientSession).IbmToolchainApiV2()
	if err != nil {
		return diag.FromErr(err)
	}

	createServiceInstanceOptions := &ibmtoolchainapiv2.CreateServiceInstanceOptions{}

	createServiceInstanceOptions.SetToolchainID(d.Get("toolchain_id").(string))
	if _, ok := d.GetOk("git_provider"); ok {
		createServiceInstanceOptions.SetServiceID(d.Get("git_provider").(string))
	}
	if _, ok := d.GetOk("parameters"); ok {
		parameters, err := ResourceIbmToolchainToolGitMapToParameters(d.Get("parameters.0").(map[string]interface{}))
		if err != nil {
			return diag.FromErr(err)
		}
		createServiceInstanceOptions.SetParameters(parameters)
	}
	if _, ok := d.GetOk("parameters_references"); ok {
		// TODO: Add code to handle map container: ParametersReferences
	}
	if _, ok := d.GetOk("container"); ok {
		container, err := ResourceIbmToolchainToolGitMapToContainer(d.Get("container.0").(map[string]interface{}))
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

	return ResourceIbmToolchainToolGitRead(context, d, meta)
}

func ResourceIbmToolchainToolGitRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
		parametersMap, err := ResourceIbmToolchainToolGitParametersToMap(serviceResponse.Parameters)
		if err != nil {
			return diag.FromErr(err)
		}
		oldParams := d.Get("parameters.0").(map[string]interface{})
		parametersMap["action"] = oldParams["action"]
		if err = d.Set("parameters", []map[string]interface{}{parametersMap}); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting parameters: %s", err))
		}
	}
	if serviceResponse.Container != nil {
		containerMap, err := ResourceIbmToolchainToolGitContainerToMap(serviceResponse.Container)
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

func ResourceIbmToolchainToolGitUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	ibmToolchainApiClient, err := meta.(conns.ClientSession).IbmToolchainApiV2()
	if err != nil {
		return diag.FromErr(err)
	}

	patchServiceInstanceOptions := &ibmtoolchainapiv2.PatchServiceInstanceOptions{}

	patchServiceInstanceOptions.SetServiceInstanceID(d.Id())

	hasChange := false

	if d.HasChange("toolchain_id") {
		return diag.FromErr(fmt.Errorf("Cannot update resource property \"%s\" with the ForceNew annotation."+
			" The resource must be re-created to update this property.", "toolchain_id"))
	}
	if d.HasChange("git_provider") {
		return diag.FromErr(fmt.Errorf("Cannot update resource property \"%s\" with the ForceNew annotation."+
			" The resource must be re-created to update this property.", "git_provider"))
	}
	if d.HasChange("container") {
		return diag.FromErr(fmt.Errorf("Cannot update resource property \"%s\" with the ForceNew annotation."+
			" The resource must be re-created to update this property.", "container"))
	}
	if d.HasChange("parameters") {
		parameters, err := ResourceIbmToolchainToolGitMapToParameters(d.Get("parameters.0").(map[string]interface{}))
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

	return ResourceIbmToolchainToolGitRead(context, d, meta)
}

func ResourceIbmToolchainToolGitDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

func ResourceIbmToolchainToolGitMapToParameters(modelMap map[string]interface{}) (map[string]interface{}, error) {
	model := make(map[string]interface{})
	if modelMap["repo_url"] != nil {
		model["repo_url"] = core.StringPtr(modelMap["repo_url"].(string))
	}
	if modelMap["action"] != nil {
		model["type"] = core.StringPtr(modelMap["action"].(string))
	}
	if modelMap["legal"] != nil {
		model["legal"] = core.BoolPtr(modelMap["legal"].(bool))
	}
	if modelMap["enable_traceability"] != nil {
		model["enable_traceability"] = core.BoolPtr(modelMap["enable_traceability"].(bool))
	}
	if modelMap["has_issues"] != nil {
		model["has_issues"] = core.BoolPtr(modelMap["has_issues"].(bool))
	}
	return model, nil
}

func ResourceIbmToolchainToolGitMapToContainer(modelMap map[string]interface{}) (*ibmtoolchainapiv2.Container, error) {
	model := &ibmtoolchainapiv2.Container{}
	model.Guid = core.StringPtr(modelMap["guid"].(string))
	model.Type = core.StringPtr(modelMap["type"].(string))
	return model, nil
}

func ResourceIbmToolchainToolGitParametersToMap(model map[string]interface{}) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model["repo_url"] != nil {
		modelMap["repo_url"] = model["repo_url"]
	}
	if model["legal"] != nil {
		modelMap["legal"] = model["legal"]
	}
	if model["enable_traceability"] != nil {
		modelMap["enable_traceability"] = model["enable_traceability"]
	}
	if model["has_issues"] != nil {
		modelMap["has_issues"] = model["has_issues"]
	}
	return modelMap, nil
}

func ResourceIbmToolchainToolGitContainerToMap(model *ibmtoolchainapiv2.Container) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["guid"] = model.Guid
	modelMap["type"] = model.Type
	return modelMap, nil
}
