// Copyright IBM Corp. 2024,2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.99.0-d27cee72-20250129-204831
 */

package sdsaas

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
	"github.com/IBM/sds-go-sdk/sdsaasv1"
)

func ResourceIBMSdsVolume() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIBMSdsVolumeCreate,
		ReadContext:   resourceIBMSdsVolumeRead,
		UpdateContext: resourceIBMSdsVolumeUpdate,
		DeleteContext: resourceIBMSdsVolumeDelete,
		Importer:      &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"sds_endpoint": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The endpoint to use for operations",
			},
			"capacity": &schema.Schema{
				Type:         schema.TypeInt,
				Required:     true,
				ValidateFunc: validate.InvokeValidator("ibm_sds_volume", "capacity"),
				Description:  "The capacity of the volume (in gigabytes).",
			},
			"name": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validate.InvokeValidator("ibm_sds_volume", "name"),
				Description:  "Unique name of the host.",
			},
			"bandwidth": &schema.Schema{
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The maximum bandwidth (in megabits per second) for the volume.",
			},
			"created_at": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The date and time that the volume was created.",
			},
			"href": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The URL for this resource.",
			},
			"iops": &schema.Schema{
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Iops The maximum I/O operations per second (IOPS) for this volume.",
			},
			"resource_type": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The resource type of the volume.",
			},
			"status": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The status of the volume resource. The enumerated values for this property will expand in the future. When processing this property, check for and log unknown values. Optionally halt processing and surface the error, or bypass the resource on which the unexpected property value was encountered.",
			},
			"status_reasons": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The reasons for the current status (if any).",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"code": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "A snake case string succinctly identifying the status reason.",
						},
						"message": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "An explanation of the status reason.",
						},
						"more_info": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "Link to documentation about this status reason.",
						},
					},
				},
			},
			"volume_mappings": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "List of volume mappings for this volume.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"status": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The status of the volume mapping. The enumerated values for this property will expand in the future. When processing this property, check for and log unknown values. Optionally halt processing and surface the error, or bypass the resource on which the unexpected property value was encountered.",
						},
						"storage_identifier": &schema.Schema{
							Type:        schema.TypeList,
							Optional:    true,
							Computed:    true,
							Description: "Storage network and ID information associated with a volume/host mapping.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"subsystem_nqn": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The NVMe target subsystem NQN (NVMe Qualified Name) that can be used for doing NVMe connect by the initiator.",
									},
									"namespace_id": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "NVMe namespace ID that can be used to co-relate the discovered devices on host to the corresponding volume.",
									},
									"namespace_uuid": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The namespace UUID associated with a volume/host mapping.",
									},
									"gateways": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "List of NVMe gateways.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"ip_address": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Network information for volume/host mappings.",
												},
												"port": &schema.Schema{
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "Network information for volume/host mappings.",
												},
											},
										},
									},
								},
							},
						},
						"href": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The URL for this resource.",
						},
						"id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Unique identifier of the mapping.",
						},
						"volume": &schema.Schema{
							Type:        schema.TypeList,
							Optional:    true,
							Computed:    true,
							Description: "The volume reference.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Unique identifer of the host.",
									},
									"name": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Unique name of the host.",
									},
								},
							},
						},
						"host": &schema.Schema{
							Type:        schema.TypeList,
							Optional:    true,
							Computed:    true,
							Description: "Host mapping schema.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Unique identifer of the host.",
									},
									"name": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Unique name of the host.",
									},
									"nqn": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The NQN (NVMe Qualified Name) as configured on the initiator (compute/host) accessing the storage.",
									},
								},
							},
						},
						"subsystem_nqn": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "The NVMe target subsystem NQN (NVMe Qualified Name) that can be used for doing NVMe connect by the initiator.",
						},
						"namespace": &schema.Schema{
							Type:        schema.TypeList,
							Optional:    true,
							Computed:    true,
							Description: "The NVMe namespace properties for a given volume mapping.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": &schema.Schema{
										Type:        schema.TypeInt,
										Optional:    true,
										Computed:    true,
										Description: "NVMe namespace ID that can be used to co-relate the discovered devices on host to the corresponding volume.",
									},
									"uuid": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
										Description: "UUID of the NVMe namespace.",
									},
								},
							},
						},
						"gateways": &schema.Schema{
							Type:        schema.TypeList,
							Optional:    true,
							Computed:    true,
							Description: "List of NVMe gateways.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ip_address": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Network information for volume/host mappings.",
									},
									"port": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Network information for volume/host mappings.",
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func getSDSConfigClient(meta interface{}, endpoint string) (*sdsaasv1.SdsaasV1, error) {
	sdsconfigClient, err := meta.(conns.ClientSession).SdsaasV1()
	if err != nil {
		return nil, err
	}
	url := conns.EnvFallBack([]string{"IBMCLOUD_SDS_ENDPOINT"}, endpoint)
	sdsconfigClient.Service.Options.URL = url
	return sdsconfigClient, nil
}

func ResourceIBMSdsVolumeValidator() *validate.ResourceValidator {
	validateSchema := make([]validate.ValidateSchema, 0)
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 "capacity",
			ValidateFunctionIdentifier: validate.IntBetween,
			Type:                       validate.TypeInt,
			Required:                   true,
			MinValue:                   "1",
			MaxValue:                   "32000",
		},
		validate.ValidateSchema{
			Identifier:                 "name",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Optional:                   true,
			Regexp:                     `^.*$`,
			MinValueLength:             1,
			MaxValueLength:             63,
		},
	)

	resourceValidator := validate.ResourceValidator{ResourceName: "ibm_sds_volume", Schema: validateSchema}
	return &resourceValidator
}

func resourceIBMSdsVolumeCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	endpoint := d.Get("sds_endpoint").(string)
	sdsaasClient, err := getSDSConfigClient(meta, endpoint)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sds_volume", "create", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	volumeCreateOptions := &sdsaasv1.VolumeCreateOptions{}

	volumeCreateOptions.SetCapacity(int64(d.Get("capacity").(int)))
	if _, ok := d.GetOk("name"); ok {
		volumeCreateOptions.SetName(d.Get("name").(string))
	}

	volume, _, err := sdsaasClient.VolumeCreateWithContext(context, volumeCreateOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("VolumeCreateWithContext failed: %s", err.Error()), "ibm_sds_volume", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(*volume.ID)

	return resourceIBMSdsVolumeRead(context, d, meta)
}

func resourceIBMSdsVolumeRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	endpoint := d.Get("sds_endpoint").(string)
	sdsaasClient, err := getSDSConfigClient(meta, endpoint)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sds_volume", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	volumeOptions := &sdsaasv1.VolumeOptions{}

	volumeOptions.SetVolumeID(d.Id())

	volume, response, err := sdsaasClient.VolumeWithContext(context, volumeOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("VolumeWithContext failed: %s", err.Error()), "ibm_sds_volume", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	if err = d.Set("capacity", flex.IntValue(volume.Capacity)); err != nil {
		err = fmt.Errorf("Error setting capacity: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sds_volume", "read", "set-capacity").GetDiag()
	}
	if !core.IsNil(volume.Name) {
		if err = d.Set("name", volume.Name); err != nil {
			err = fmt.Errorf("Error setting name: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sds_volume", "read", "set-name").GetDiag()
		}
	}
	if err = d.Set("bandwidth", flex.IntValue(volume.Bandwidth)); err != nil {
		err = fmt.Errorf("Error setting bandwidth: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sds_volume", "read", "set-bandwidth").GetDiag()
	}
	if err = d.Set("created_at", flex.DateTimeToString(volume.CreatedAt)); err != nil {
		err = fmt.Errorf("Error setting created_at: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sds_volume", "read", "set-created_at").GetDiag()
	}
	if err = d.Set("href", volume.Href); err != nil {
		err = fmt.Errorf("Error setting href: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sds_volume", "read", "set-href").GetDiag()
	}
	if err = d.Set("iops", flex.IntValue(volume.Iops)); err != nil {
		err = fmt.Errorf("Error setting iops: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sds_volume", "read", "set-iops").GetDiag()
	}
	if err = d.Set("resource_type", volume.ResourceType); err != nil {
		err = fmt.Errorf("Error setting resource_type: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sds_volume", "read", "set-resource_type").GetDiag()
	}
	if !core.IsNil(volume.Status) {
		if err = d.Set("status", volume.Status); err != nil {
			err = fmt.Errorf("Error setting status: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sds_volume", "read", "set-status").GetDiag()
		}
	}
	if !core.IsNil(volume.StatusReasons) {
		statusReasons := []map[string]interface{}{}
		for _, statusReasonsItem := range volume.StatusReasons {
			statusReasonsItemMap, err := ResourceIBMSdsVolumeVolumeStatusReasonToMap(&statusReasonsItem) // #nosec G601
			if err != nil {
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sds_volume", "read", "status_reasons-to-map").GetDiag()
			}
			statusReasons = append(statusReasons, statusReasonsItemMap)
		}
		if err = d.Set("status_reasons", statusReasons); err != nil {
			err = fmt.Errorf("Error setting status_reasons: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sds_volume", "read", "set-status_reasons").GetDiag()
		}
	}
	volumeMappings := []map[string]interface{}{}
	for _, volumeMappingsItem := range volume.VolumeMappings {
		volumeMappingsItemMap, err := ResourceIBMSdsVolumeVolumeMappingToMap(&volumeMappingsItem) // #nosec G601
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sds_volume", "read", "volume_mappings-to-map").GetDiag()
		}
		volumeMappings = append(volumeMappings, volumeMappingsItemMap)
	}
	if err = d.Set("volume_mappings", volumeMappings); err != nil {
		err = fmt.Errorf("Error setting volume_mappings: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sds_volume", "read", "set-volume_mappings").GetDiag()
	}

	return nil
}

func resourceIBMSdsVolumeUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	endpoint := d.Get("sds_endpoint").(string)
	sdsaasClient, err := getSDSConfigClient(meta, endpoint)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sds_volume", "update", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	volumeUpdateOptions := &sdsaasv1.VolumeUpdateOptions{}

	volumeUpdateOptions.SetVolumeID(d.Id())

	hasChange := false

	patchVals := &sdsaasv1.VolumePatch{}
	if d.HasChange("capacity") {
		newCapacity := int64(d.Get("capacity").(int))
		patchVals.Capacity = &newCapacity
		hasChange = true
	}
	if d.HasChange("name") {
		newName := d.Get("name").(string)
		patchVals.Name = &newName
		hasChange = true
	}

	if hasChange {
		// Fields with `nil` values are omitted from the generic map,
		// so we need to re-add them to support removing arguments
		// in merge-patch operations sent to the service.
		volumeUpdateOptions.VolumePatch = ResourceIBMSdsVolumeVolumePatchAsPatch(patchVals, d)

		_, _, err = sdsaasClient.VolumeUpdateWithContext(context, volumeUpdateOptions)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("VolumeUpdateWithContext failed: %s", err.Error()), "ibm_sds_volume", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
	}

	return resourceIBMSdsVolumeRead(context, d, meta)
}

func resourceIBMSdsVolumeDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	endpoint := d.Get("sds_endpoint").(string)
	sdsaasClient, err := getSDSConfigClient(meta, endpoint)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sds_volume", "delete", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	volumeDeleteOptions := &sdsaasv1.VolumeDeleteOptions{}

	volumeDeleteOptions.SetVolumeID(d.Id())

	_, err = sdsaasClient.VolumeDeleteWithContext(context, volumeDeleteOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("VolumeDeleteWithContext failed: %s", err.Error()), "ibm_sds_volume", "delete")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId("")

	return nil
}

func ResourceIBMSdsVolumeVolumeStatusReasonToMap(model *sdsaasv1.VolumeStatusReason) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["code"] = *model.Code
	modelMap["message"] = *model.Message
	if model.MoreInfo != nil {
		modelMap["more_info"] = *model.MoreInfo
	}
	return modelMap, nil
}

func ResourceIBMSdsVolumeVolumeMappingToMap(model *sdsaasv1.VolumeMapping) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["status"] = *model.Status
	if model.StorageIdentifier != nil {
		storageIdentifierMap, err := ResourceIBMSdsVolumeStorageIdentifierToMap(model.StorageIdentifier)
		if err != nil {
			return modelMap, err
		}
		modelMap["storage_identifier"] = []map[string]interface{}{storageIdentifierMap}
	}
	modelMap["href"] = *model.Href
	modelMap["id"] = *model.ID
	if model.Volume != nil {
		volumeMap, err := ResourceIBMSdsVolumeVolumeReferenceToMap(model.Volume)
		if err != nil {
			return modelMap, err
		}
		modelMap["volume"] = []map[string]interface{}{volumeMap}
	}
	if model.Host != nil {
		hostMap, err := ResourceIBMSdsVolumeHostReferenceToMap(model.Host)
		if err != nil {
			return modelMap, err
		}
		modelMap["host"] = []map[string]interface{}{hostMap}
	}
	if model.SubsystemNqn != nil {
		modelMap["subsystem_nqn"] = *model.SubsystemNqn
	}
	if model.Namespace != nil {
		namespaceMap, err := ResourceIBMSdsVolumeNamespaceToMap(model.Namespace)
		if err != nil {
			return modelMap, err
		}
		modelMap["namespace"] = []map[string]interface{}{namespaceMap}
	}
	if model.Gateways != nil {
		gateways := []map[string]interface{}{}
		for _, gatewaysItem := range model.Gateways {
			gatewaysItemMap, err := ResourceIBMSdsVolumeGatewayToMap(&gatewaysItem) // #nosec G601
			if err != nil {
				return modelMap, err
			}
			gateways = append(gateways, gatewaysItemMap)
		}
		modelMap["gateways"] = gateways
	}
	return modelMap, nil
}

func ResourceIBMSdsVolumeStorageIdentifierToMap(model *sdsaasv1.StorageIdentifier) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["subsystem_nqn"] = *model.SubsystemNqn
	modelMap["namespace_id"] = flex.IntValue(model.NamespaceID)
	modelMap["namespace_uuid"] = *model.NamespaceUUID
	gateways := []map[string]interface{}{}
	for _, gatewaysItem := range model.Gateways {
		gatewaysItemMap, err := ResourceIBMSdsVolumeGatewayToMap(&gatewaysItem) // #nosec G601
		if err != nil {
			return modelMap, err
		}
		gateways = append(gateways, gatewaysItemMap)
	}
	modelMap["gateways"] = gateways
	return modelMap, nil
}

func ResourceIBMSdsVolumeGatewayToMap(model *sdsaasv1.Gateway) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["ip_address"] = *model.IPAddress
	modelMap["port"] = flex.IntValue(model.Port)
	return modelMap, nil
}

func ResourceIBMSdsVolumeVolumeReferenceToMap(model *sdsaasv1.VolumeReference) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["id"] = *model.ID
	modelMap["name"] = *model.Name
	return modelMap, nil
}

func ResourceIBMSdsVolumeHostReferenceToMap(model *sdsaasv1.HostReference) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["id"] = *model.ID
	modelMap["name"] = *model.Name
	modelMap["nqn"] = *model.Nqn
	return modelMap, nil
}

func ResourceIBMSdsVolumeNamespaceToMap(model *sdsaasv1.Namespace) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ID != nil {
		modelMap["id"] = flex.IntValue(model.ID)
	}
	if model.UUID != nil {
		modelMap["uuid"] = *model.UUID
	}
	return modelMap, nil
}

func ResourceIBMSdsVolumeVolumePatchAsPatch(patchVals *sdsaasv1.VolumePatch, d *schema.ResourceData) map[string]interface{} {
	patch, _ := patchVals.AsPatch()
	var path string

	path = "capacity"
	if _, exists := d.GetOk(path); d.HasChange(path) && !exists {
		patch["capacity"] = nil
	} else if !exists {
		delete(patch, "capacity")
	}
	path = "name"
	if _, exists := d.GetOk(path); d.HasChange(path) && !exists {
		patch["name"] = nil
	} else if !exists {
		delete(patch, "name")
	}

	return patch
}
