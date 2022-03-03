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

func ResourceIbmToolchainToolArtifactory() *schema.Resource {
	return &schema.Resource{
		CreateContext: ResourceIbmToolchainToolArtifactoryCreate,
		ReadContext:   ResourceIbmToolchainToolArtifactoryRead,
		UpdateContext: ResourceIbmToolchainToolArtifactoryUpdate,
		DeleteContext: ResourceIbmToolchainToolArtifactoryDelete,
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

func ResourceIbmToolchainToolArtifactoryCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	ibmToolchainApiClient, err := meta.(conns.ClientSession).IbmToolchainApiV2()
	if err != nil {
		return diag.FromErr(err)
	}

	createServiceInstanceOptions := &ibmtoolchainapiv2.CreateServiceInstanceOptions{}

	createServiceInstanceOptions.SetServiceID("artifactory")
	createServiceInstanceOptions.SetToolchainID(d.Get("toolchain_id").(string))
	if _, ok := d.GetOk("parameters"); ok {
		parameters, err := ResourceIbmToolchainToolArtifactoryMapToParameters(d.Get("parameters.0").(map[string]interface{}))
		if err != nil {
			return diag.FromErr(err)
		}
		createServiceInstanceOptions.SetParameters(parameters)
	}
	if _, ok := d.GetOk("parameters_references"); ok {
		// TODO: Add code to handle map container: ParametersReferences
	}
	if _, ok := d.GetOk("container"); ok {
		container, err := ResourceIbmToolchainToolArtifactoryMapToContainer(d.Get("container.0").(map[string]interface{}))
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

	return ResourceIbmToolchainToolArtifactoryRead(context, d, meta)
}

func ResourceIbmToolchainToolArtifactoryRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
		parametersMap, err := ResourceIbmToolchainToolArtifactoryParametersToMap(serviceResponse.Parameters)
		if err != nil {
			return diag.FromErr(err)
		}
		oldParams := d.Get("parameters.0").(map[string]interface{})
		if parametersMap["token"] == "****" {
			parametersMap["token"] = oldParams["token"]
		}

		if err = d.Set("parameters", []map[string]interface{}{parametersMap}); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting parameters: %s", err))
		}
	}
	if serviceResponse.Container != nil {
		containerMap, err := ResourceIbmToolchainToolArtifactoryContainerToMap(serviceResponse.Container)
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

func ResourceIbmToolchainToolArtifactoryUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	ibmToolchainApiClient, err := meta.(conns.ClientSession).IbmToolchainApiV2()
	if err != nil {
		return diag.FromErr(err)
	}

	patchServiceInstanceOptions := &ibmtoolchainapiv2.PatchServiceInstanceOptions{}

	patchServiceInstanceOptions.SetServiceID("artifactory")
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
		parameters, err := ResourceIbmToolchainToolArtifactoryMapToParameters(d.Get("parameters.0").(map[string]interface{}))
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

	return ResourceIbmToolchainToolArtifactoryRead(context, d, meta)
}

func ResourceIbmToolchainToolArtifactoryDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

func ResourceIbmToolchainToolArtifactoryMapToParameters(modelMap map[string]interface{}) (map[string]interface{}, error) {
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

func ResourceIbmToolchainToolArtifactoryMapToContainer(modelMap map[string]interface{}) (*ibmtoolchainapiv2.Container, error) {
	model := &ibmtoolchainapiv2.Container{}
	model.Guid = core.StringPtr(modelMap["guid"].(string))
	model.Type = core.StringPtr(modelMap["type"].(string))
	return model, nil
}

func ResourceIbmToolchainToolArtifactoryParametersToMap(model map[string]interface{}) (map[string]interface{}, error) {
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
		modelMap["token"] = model["token"]
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

func ResourceIbmToolchainToolArtifactoryContainerToMap(model *ibmtoolchainapiv2.Container) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["guid"] = model.Guid
	modelMap["type"] = model.Type
	return modelMap, nil
}
