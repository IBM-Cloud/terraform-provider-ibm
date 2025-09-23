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

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/sds-go-sdk/sdsaasv1"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/flex"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/validate"
)

func ResourceIBMSdsHost() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIBMSdsHostCreate,
		ReadContext:   resourceIBMSdsHostRead,
		UpdateContext: resourceIBMSdsHostUpdate,
		DeleteContext: resourceIBMSdsHostDelete,
		Importer:      &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"sds_endpoint": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The endpoint to use for operations",
			},
			"name": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validate.InvokeValidator("ibm_sds_host", "name"),
				Description:  "Unique name of the host.",
			},
			"nqn": &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validate.InvokeValidator("ibm_sds_host", "nqn"),
				Description:  "The NQN (NVMe Qualified Name) as configured on the initiator (compute/host) accessing the storage.",
			},
			"volume_mappings": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				Description: "The host-to-volume map.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"status": &schema.Schema{
							Type:        schema.TypeString,
							Required:    true,
							Description: "The status of the volume mapping. The enumerated values for this property will expand in the future. When processing this property, check for and log unknown values. Optionally halt processing and surface the error, or bypass the resource on which the unexpected property value was encountered.",
						},
						"storage_identifier": &schema.Schema{
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Description: "Storage network and ID information associated with a volume/host mapping.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"subsystem_nqn": &schema.Schema{
										Type:        schema.TypeString,
										Required:    true,
										Description: "The NVMe target subsystem NQN (NVMe Qualified Name) that can be used for doing NVMe connect by the initiator.",
									},
									"namespace_id": &schema.Schema{
										Type:        schema.TypeInt,
										Required:    true,
										Description: "NVMe namespace ID that can be used to co-relate the discovered devices on host to the corresponding volume.",
									},
									"namespace_uuid": &schema.Schema{
										Type:        schema.TypeString,
										Required:    true,
										Description: "The namespace UUID associated with a volume/host mapping.",
									},
									"gateways": &schema.Schema{
										Type:        schema.TypeList,
										Required:    true,
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
							Required:    true,
							Description: "The URL for this resource.",
						},
						"id": &schema.Schema{
							Type:        schema.TypeString,
							Required:    true,
							Description: "Unique identifier of the mapping.",
						},
						"volume": &schema.Schema{
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
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
						"host": &schema.Schema{
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Description: "Host mapping schema.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": &schema.Schema{
										Type:        schema.TypeString,
										Required:    true,
										Description: "Unique identifer of the host.",
									},
									"name": &schema.Schema{
										Type:        schema.TypeString,
										Required:    true,
										Description: "Unique name of the host.",
									},
									"nqn": &schema.Schema{
										Type:        schema.TypeString,
										Required:    true,
										Description: "The NQN (NVMe Qualified Name) as configured on the initiator (compute/host) accessing the storage.",
									},
								},
							},
						},
						"subsystem_nqn": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The NVMe target subsystem NQN (NVMe Qualified Name) that can be used for doing NVMe connect by the initiator.",
						},
						"namespace": &schema.Schema{
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Description: "The NVMe namespace properties for a given volume mapping.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": &schema.Schema{
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "NVMe namespace ID that can be used to co-relate the discovered devices on host to the corresponding volume.",
									},
									"uuid": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "UUID of the NVMe namespace.",
									},
								},
							},
						},
						"gateways": &schema.Schema{
							Type:        schema.TypeList,
							Optional:    true,
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
			"created_at": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The date and time when the resource was created.",
			},
			"href": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The URL for this resource.",
			},
		},
	}
}

func ResourceIBMSdsHostValidator() *validate.ResourceValidator {
	validateSchema := make([]validate.ValidateSchema, 0)
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 "name",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Optional:                   true,
			Regexp:                     `^-?([a-z]|[a-z][-a-z0-9]*[a-z0-9]|[0-9][-a-z0-9]*([a-z]|[-a-z][-a-z0-9]*[a-z0-9]))$`,
			MinValueLength:             1,
			MaxValueLength:             63,
		},
		validate.ValidateSchema{
			Identifier:                 "nqn",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Required:                   true,
			Regexp:                     `^nqn\.\d{4}-\d{2}\.[a-z0-9-]+(?:\.[a-z0-9-]+)*:[a-zA-Z0-9.\-:]+$`,
			MinValueLength:             16,
			MaxValueLength:             223,
		},
	)

	resourceValidator := validate.ResourceValidator{ResourceName: "ibm_sds_host", Schema: validateSchema}
	return &resourceValidator
}

func resourceIBMSdsHostCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	endpoint := d.Get("sds_endpoint").(string)
	sdsaasClient, err := getSDSConfigClient(meta, endpoint)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sds_host", "create", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	hostCreateOptions := &sdsaasv1.HostCreateOptions{}

	hostCreateOptions.SetNqn(d.Get("nqn").(string))
	if _, ok := d.GetOk("name"); ok {
		hostCreateOptions.SetName(d.Get("name").(string))
	}
	if _, ok := d.GetOk("volume_mappings"); ok {
		var volumeMappings []sdsaasv1.VolumeMappingPrototype
		for _, v := range d.Get("volume_mappings").([]interface{}) {
			value := v.(map[string]interface{})
			volumeMappingsItem, err := ResourceIBMSdsHostMapToVolumeMappingPrototype(value)
			if err != nil {
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sds_host", "create", "parse-volume_mappings").GetDiag()
			}
			volumeMappings = append(volumeMappings, *volumeMappingsItem)
		}
		hostCreateOptions.SetVolumeMappings(volumeMappings)
	}

	host, _, err := sdsaasClient.HostCreateWithContext(context, hostCreateOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("HostCreateWithContext failed: %s", err.Error()), "ibm_sds_host", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(*host.ID)

	return resourceIBMSdsHostRead(context, d, meta)
}

func resourceIBMSdsHostRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	endpoint := d.Get("sds_endpoint").(string)
	sdsaasClient, err := getSDSConfigClient(meta, endpoint)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sds_host", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	hostOptions := &sdsaasv1.HostOptions{}

	hostOptions.SetHostID(d.Id())

	host, response, err := sdsaasClient.HostWithContext(context, hostOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("HostWithContext failed: %s", err.Error()), "ibm_sds_host", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	if !core.IsNil(host.Name) {
		if err = d.Set("name", host.Name); err != nil {
			err = fmt.Errorf("Error setting name: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sds_host", "read", "set-name").GetDiag()
		}
	}
	if err = d.Set("nqn", host.Nqn); err != nil {
		err = fmt.Errorf("Error setting nqn: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sds_host", "read", "set-nqn").GetDiag()
	}
	if !core.IsNil(host.VolumeMappings) {
		volumeMappings := []map[string]interface{}{}
		for _, volumeMappingsItem := range host.VolumeMappings {
			volumeMappingsItemMap, err := ResourceIBMSdsHostVolumeMappingToMap(&volumeMappingsItem) // #nosec G601
			if err != nil {
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sds_host", "read", "volume_mappings-to-map").GetDiag()
			}
			volumeMappings = append(volumeMappings, volumeMappingsItemMap)
		}
		if err = d.Set("volume_mappings", volumeMappings); err != nil {
			err = fmt.Errorf("Error setting volume_mappings: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sds_host", "read", "set-volume_mappings").GetDiag()
		}
	}
	if err = d.Set("created_at", flex.DateTimeToString(host.CreatedAt)); err != nil {
		err = fmt.Errorf("Error setting created_at: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sds_host", "read", "set-created_at").GetDiag()
	}
	if err = d.Set("href", host.Href); err != nil {
		err = fmt.Errorf("Error setting href: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sds_host", "read", "set-href").GetDiag()
	}

	return nil
}

func resourceIBMSdsHostUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	endpoint := d.Get("sds_endpoint").(string)
	sdsaasClient, err := getSDSConfigClient(meta, endpoint)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sds_host", "update", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	hostUpdateOptions := &sdsaasv1.HostUpdateOptions{}

	hostUpdateOptions.SetHostID(d.Id())

	hasChange := false

	patchVals := &sdsaasv1.HostPatch{}
	if d.HasChange("name") {
		newName := d.Get("name").(string)
		patchVals.Name = &newName
		hasChange = true
	}

	if hasChange {
		// Fields with `nil` values are omitted from the generic map,
		// so we need to re-add them to support removing arguments
		// in merge-patch operations sent to the service.
		hostUpdateOptions.HostPatch = ResourceIBMSdsHostHostPatchAsPatch(patchVals, d)

		_, _, err = sdsaasClient.HostUpdateWithContext(context, hostUpdateOptions)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("HostUpdateWithContext failed: %s", err.Error()), "ibm_sds_host", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
	}

	return resourceIBMSdsHostRead(context, d, meta)
}

func resourceIBMSdsHostDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	endpoint := d.Get("sds_endpoint").(string)
	sdsaasClient, err := getSDSConfigClient(meta, endpoint)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sds_host", "delete", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	hostDeleteOptions := &sdsaasv1.HostDeleteOptions{}

	hostDeleteOptions.SetHostID(d.Id())

	_, err = sdsaasClient.HostDeleteWithContext(context, hostDeleteOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("HostDeleteWithContext failed: %s", err.Error()), "ibm_sds_host", "delete")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId("")

	return nil
}

func ResourceIBMSdsHostMapToVolumeMappingPrototype(modelMap map[string]interface{}) (*sdsaasv1.VolumeMappingPrototype, error) {
	model := &sdsaasv1.VolumeMappingPrototype{}
	VolumeModel, err := ResourceIBMSdsHostMapToVolumeIdentity(modelMap["volume"].([]interface{})[0].(map[string]interface{}))
	if err != nil {
		return model, err
	}
	model.Volume = VolumeModel
	return model, nil
}

func ResourceIBMSdsHostMapToVolumeIdentity(modelMap map[string]interface{}) (*sdsaasv1.VolumeIdentity, error) {
	model := &sdsaasv1.VolumeIdentity{}
	model.ID = core.StringPtr(modelMap["id"].(string))
	return model, nil
}

func ResourceIBMSdsHostVolumeMappingToMap(model *sdsaasv1.VolumeMapping) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["status"] = *model.Status
	if model.StorageIdentifier != nil {
		storageIdentifierMap, err := ResourceIBMSdsHostStorageIdentifierToMap(model.StorageIdentifier)
		if err != nil {
			return modelMap, err
		}
		modelMap["storage_identifier"] = []map[string]interface{}{storageIdentifierMap}
	}
	modelMap["href"] = *model.Href
	modelMap["id"] = *model.ID
	if model.Volume != nil {
		volumeMap, err := ResourceIBMSdsHostVolumeReferenceToMap(model.Volume)
		if err != nil {
			return modelMap, err
		}
		modelMap["volume"] = []map[string]interface{}{volumeMap}
	}
	if model.Host != nil {
		hostMap, err := ResourceIBMSdsHostHostReferenceToMap(model.Host)
		if err != nil {
			return modelMap, err
		}
		modelMap["host"] = []map[string]interface{}{hostMap}
	}
	if model.SubsystemNqn != nil {
		modelMap["subsystem_nqn"] = *model.SubsystemNqn
	}
	if model.Namespace != nil {
		namespaceMap, err := ResourceIBMSdsHostNamespaceToMap(model.Namespace)
		if err != nil {
			return modelMap, err
		}
		modelMap["namespace"] = []map[string]interface{}{namespaceMap}
	}
	if model.Gateways != nil {
		gateways := []map[string]interface{}{}
		for _, gatewaysItem := range model.Gateways {
			gatewaysItemMap, err := ResourceIBMSdsHostGatewayToMap(&gatewaysItem) // #nosec G601
			if err != nil {
				return modelMap, err
			}
			gateways = append(gateways, gatewaysItemMap)
		}
		modelMap["gateways"] = gateways
	}
	return modelMap, nil
}

func ResourceIBMSdsHostStorageIdentifierToMap(model *sdsaasv1.StorageIdentifier) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["subsystem_nqn"] = *model.SubsystemNqn
	modelMap["namespace_id"] = flex.IntValue(model.NamespaceID)
	modelMap["namespace_uuid"] = *model.NamespaceUUID
	gateways := []map[string]interface{}{}
	for _, gatewaysItem := range model.Gateways {
		gatewaysItemMap, err := ResourceIBMSdsHostGatewayToMap(&gatewaysItem) // #nosec G601
		if err != nil {
			return modelMap, err
		}
		gateways = append(gateways, gatewaysItemMap)
	}
	modelMap["gateways"] = gateways
	return modelMap, nil
}

func ResourceIBMSdsHostGatewayToMap(model *sdsaasv1.Gateway) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["ip_address"] = *model.IPAddress
	modelMap["port"] = flex.IntValue(model.Port)
	return modelMap, nil
}

func ResourceIBMSdsHostVolumeReferenceToMap(model *sdsaasv1.VolumeReference) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["id"] = *model.ID
	modelMap["name"] = *model.Name
	return modelMap, nil
}

func ResourceIBMSdsHostHostReferenceToMap(model *sdsaasv1.HostReference) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["id"] = *model.ID
	modelMap["name"] = *model.Name
	modelMap["nqn"] = *model.Nqn
	return modelMap, nil
}

func ResourceIBMSdsHostNamespaceToMap(model *sdsaasv1.Namespace) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ID != nil {
		modelMap["id"] = flex.IntValue(model.ID)
	}
	if model.UUID != nil {
		modelMap["uuid"] = *model.UUID
	}
	return modelMap, nil
}

func ResourceIBMSdsHostHostPatchAsPatch(patchVals *sdsaasv1.HostPatch, d *schema.ResourceData) map[string]interface{} {
	patch, _ := patchVals.AsPatch()
	var path string

	path = "name"
	if _, exists := d.GetOk(path); d.HasChange(path) && !exists {
		patch["name"] = nil
	} else if !exists {
		delete(patch, "name")
	}

	return patch
}
