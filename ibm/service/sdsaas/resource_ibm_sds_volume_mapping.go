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

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/validate"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/sds-go-sdk/sdsaasv1"
)

func ResourceIBMSdsVolumeMapping() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIBMSdsVolumeMappingCreate,
		ReadContext:   resourceIBMSdsVolumeMappingRead,
		UpdateContext: resourceIBMSdsVolumeMappingUpdate,
		DeleteContext: resourceIBMSdsVolumeMappingDelete,
		Importer:      &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"sds_endpoint": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The endpoint to use for operations",
			},
			"host_id": &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validate.InvokeValidator("ibm_sds_volume_mapping", "host_id"),
				Description:  "A unique host ID.",
			},
			"volume": &schema.Schema{
				Type:        schema.TypeList,
				MinItems:    1,
				MaxItems:    1,
				Required:    true,
				Description: "The volume reference.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:        schema.TypeString,
							Required:    true,
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
			"status": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The status of the volume mapping. The enumerated values for this property will expand in the future. When processing this property, check for and log unknown values. Optionally halt processing and surface the error, or bypass the resource on which the unexpected property value was encountered.",
			},
			"storage_identifier": &schema.Schema{
				Type:        schema.TypeList,
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
			"host": &schema.Schema{
				Type:        schema.TypeList,
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
				Computed:    true,
				Description: "The NVMe target subsystem NQN (NVMe Qualified Name) that can be used for doing NVMe connect by the initiator.",
			},
			"namespace": &schema.Schema{
				Type:        schema.TypeList,
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
			"volume_mapping_id": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Unique identifier of the mapping.",
			},
		},
	}
}

func ResourceIBMSdsVolumeMappingValidator() *validate.ResourceValidator {
	validateSchema := make([]validate.ValidateSchema, 0)
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 "host_id",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Required:                   true,
			Regexp:                     `^\S+$`,
			MinValueLength:             0,
			MaxValueLength:             200,
		},
	)

	resourceValidator := validate.ResourceValidator{ResourceName: "ibm_sds_volume_mapping", Schema: validateSchema}
	return &resourceValidator
}

func resourceIBMSdsVolumeMappingCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	endpoint := d.Get("sds_endpoint").(string)
	sdsaasClient, err := getSDSConfigClient(meta, endpoint)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sds_volume_mapping", "update", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	hostMappingCreateOptions := &sdsaasv1.HostMappingCreateOptions{}

	hostMappingCreateOptions.SetHostID(d.Get("host_id").(string))
	volumeModel, err := ResourceIBMSdsVolumeMappingMapToVolumeIdentity(d.Get("volume.0").(map[string]interface{}))
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sds_volume_mapping", "update", "parse-volume").GetDiag()
	}
	hostMappingCreateOptions.SetVolume(volumeModel)

	volumeMapping, _, err := sdsaasClient.HostMappingCreateWithContext(context, hostMappingCreateOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("HostMappingCreateWithContext failed: %s", err.Error()), "ibm_sds_volume_mapping", "update")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(fmt.Sprintf("%s/%s", *hostMappingCreateOptions.HostID, *volumeMapping.ID))

	return resourceIBMSdsVolumeMappingRead(context, d, meta)
}

func resourceIBMSdsVolumeMappingRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	endpoint := d.Get("sds_endpoint").(string)
	sdsaasClient, err := getSDSConfigClient(meta, endpoint)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sds_volume_mapping", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	hostMappingOptions := &sdsaasv1.HostMappingOptions{}

	parts, err := flex.SepIdParts(d.Id(), "/")
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sds_volume_mapping", "read", "sep-id-parts").GetDiag()
	}

	hostMappingOptions.SetHostID(parts[0])
	hostMappingOptions.SetVolumeMappingID(parts[1])

	volumeMapping, response, err := sdsaasClient.HostMappingWithContext(context, hostMappingOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("HostMappingWithContext failed: %s", err.Error()), "ibm_sds_volume_mapping", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	volumeMap, err := ResourceIBMSdsVolumeMappingVolumeReferenceToMap(volumeMapping.Volume)
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sds_volume_mapping", "read", "volume-to-map").GetDiag()
	}
	if err = d.Set("volume", []map[string]interface{}{volumeMap}); err != nil {
		err = fmt.Errorf("Error setting volume: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sds_volume_mapping", "read", "set-volume").GetDiag()
	}
	if err = d.Set("status", volumeMapping.Status); err != nil {
		err = fmt.Errorf("Error setting status: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sds_volume_mapping", "read", "set-status").GetDiag()
	}
	if !core.IsNil(volumeMapping.StorageIdentifier) {
		storageIdentifierMap, err := ResourceIBMSdsVolumeMappingStorageIdentifierToMap(volumeMapping.StorageIdentifier)
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sds_volume_mapping", "read", "storage_identifier-to-map").GetDiag()
		}
		if err = d.Set("storage_identifier", []map[string]interface{}{storageIdentifierMap}); err != nil {
			err = fmt.Errorf("Error setting storage_identifier: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sds_volume_mapping", "read", "set-storage_identifier").GetDiag()
		}
	}
	if err = d.Set("href", volumeMapping.Href); err != nil {
		err = fmt.Errorf("Error setting href: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sds_volume_mapping", "read", "set-href").GetDiag()
	}
	if !core.IsNil(volumeMapping.Host) {
		hostMap, err := ResourceIBMSdsVolumeMappingHostReferenceToMap(volumeMapping.Host)
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sds_volume_mapping", "read", "host-to-map").GetDiag()
		}
		if err = d.Set("host", []map[string]interface{}{hostMap}); err != nil {
			err = fmt.Errorf("Error setting host: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sds_volume_mapping", "read", "set-host").GetDiag()
		}
	}
	if !core.IsNil(volumeMapping.SubsystemNqn) {
		if err = d.Set("subsystem_nqn", volumeMapping.SubsystemNqn); err != nil {
			err = fmt.Errorf("Error setting subsystem_nqn: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sds_volume_mapping", "read", "set-subsystem_nqn").GetDiag()
		}
	}
	if !core.IsNil(volumeMapping.Namespace) {
		namespaceMap, err := ResourceIBMSdsVolumeMappingNamespaceToMap(volumeMapping.Namespace)
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sds_volume_mapping", "read", "namespace-to-map").GetDiag()
		}
		if err = d.Set("namespace", []map[string]interface{}{namespaceMap}); err != nil {
			err = fmt.Errorf("Error setting namespace: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sds_volume_mapping", "read", "set-namespace").GetDiag()
		}
	}
	if !core.IsNil(volumeMapping.Gateways) {
		gateways := []map[string]interface{}{}
		for _, gatewaysItem := range volumeMapping.Gateways {
			gatewaysItemMap, err := ResourceIBMSdsVolumeMappingGatewayToMap(&gatewaysItem) // #nosec G601
			if err != nil {
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sds_volume_mapping", "read", "gateways-to-map").GetDiag()
			}
			gateways = append(gateways, gatewaysItemMap)
		}
		if err = d.Set("gateways", gateways); err != nil {
			err = fmt.Errorf("Error setting gateways: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sds_volume_mapping", "read", "set-gateways").GetDiag()
		}
	}
	if err = d.Set("volume_mapping_id", volumeMapping.ID); err != nil {
		err = fmt.Errorf("Error setting volume_mapping_id: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sds_volume_mapping", "read", "set-volume_mapping_id").GetDiag()
	}

	return nil
}

func resourceIBMSdsVolumeMappingUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	endpoint := d.Get("sds_endpoint").(string)
	sdsaasClient, err := getSDSConfigClient(meta, endpoint)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sds_volume_mapping", "update", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	hostMappingCreateOptions := &sdsaasv1.HostMappingCreateOptions{}

	parts, err := flex.SepIdParts(d.Id(), "/")
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sds_volume_mapping", "update", "sep-id-parts").GetDiag()
	}

	hostMappingCreateOptions.SetHostID(parts[0])

	volumeModel, err := ResourceIBMSdsVolumeMappingMapToVolumeIdentity(d.Get("volume.0").(map[string]interface{}))
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sds_volume_mapping", "update", "parse-volume").GetDiag()
	}
	hostMappingCreateOptions.SetVolume(volumeModel)

	hasChange := false

	if d.HasChange("host_id") {
		errMsg := fmt.Sprintf("Cannot update resource property \"%s\" with the ForceNew annotation."+
			" The resource must be re-created to update this property.", "host_id")
		return flex.DiscriminatedTerraformErrorf(nil, errMsg, "ibm_sds_volume_mapping", "update", "host_id-forces-new").GetDiag()
	}
	if d.HasChange("volume") {
		volume, err := ResourceIBMSdsVolumeMappingMapToVolumeIdentity(d.Get("volume.0").(map[string]interface{}))
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sds_volume_mapping", "update", "parse-volume").GetDiag()
		}
		hostMappingCreateOptions.SetVolume(volume)
		hasChange = true
	}

	if hasChange {
		_, _, err = sdsaasClient.HostMappingCreateWithContext(context, hostMappingCreateOptions)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("HostMappingCreateWithContext failed: %s", err.Error()), "ibm_sds_volume_mapping", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
	}

	return resourceIBMSdsVolumeMappingRead(context, d, meta)
}

func resourceIBMSdsVolumeMappingDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	endpoint := d.Get("sds_endpoint").(string)
	sdsaasClient, err := getSDSConfigClient(meta, endpoint)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sds_volume_mapping", "delete", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	hostMappingDeleteOptions := &sdsaasv1.HostMappingDeleteOptions{}

	parts, err := flex.SepIdParts(d.Id(), "/")
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sds_volume_mapping", "delete", "sep-id-parts").GetDiag()
	}

	hostMappingDeleteOptions.SetHostID(parts[0])
	hostMappingDeleteOptions.SetVolumeMappingID(parts[1])

	_, err = sdsaasClient.HostMappingDeleteWithContext(context, hostMappingDeleteOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("HostMappingDeleteWithContext failed: %s", err.Error()), "ibm_sds_volume_mapping", "delete")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId("")

	return nil
}

func ResourceIBMSdsVolumeMappingMapToVolumeIdentity(modelMap map[string]interface{}) (*sdsaasv1.VolumeIdentity, error) {
	model := &sdsaasv1.VolumeIdentity{}
	model.ID = core.StringPtr(modelMap["id"].(string))
	return model, nil
}

func ResourceIBMSdsVolumeMappingVolumeReferenceToMap(model *sdsaasv1.VolumeReference) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["id"] = *model.ID
	modelMap["name"] = *model.Name
	return modelMap, nil
}

func ResourceIBMSdsVolumeMappingStorageIdentifierToMap(model *sdsaasv1.StorageIdentifier) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["subsystem_nqn"] = *model.SubsystemNqn
	modelMap["namespace_id"] = flex.IntValue(model.NamespaceID)
	modelMap["namespace_uuid"] = *model.NamespaceUUID
	gateways := []map[string]interface{}{}
	for _, gatewaysItem := range model.Gateways {
		gatewaysItemMap, err := ResourceIBMSdsVolumeMappingGatewayToMap(&gatewaysItem) // #nosec G601
		if err != nil {
			return modelMap, err
		}
		gateways = append(gateways, gatewaysItemMap)
	}
	modelMap["gateways"] = gateways
	return modelMap, nil
}

func ResourceIBMSdsVolumeMappingGatewayToMap(model *sdsaasv1.Gateway) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["ip_address"] = *model.IPAddress
	modelMap["port"] = flex.IntValue(model.Port)
	return modelMap, nil
}

func ResourceIBMSdsVolumeMappingHostReferenceToMap(model *sdsaasv1.HostReference) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["id"] = *model.ID
	modelMap["name"] = *model.Name
	modelMap["nqn"] = *model.Nqn
	return modelMap, nil
}

func ResourceIBMSdsVolumeMappingNamespaceToMap(model *sdsaasv1.Namespace) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ID != nil {
		modelMap["id"] = flex.IntValue(model.ID)
	}
	if model.UUID != nil {
		modelMap["uuid"] = *model.UUID
	}
	return modelMap, nil
}
