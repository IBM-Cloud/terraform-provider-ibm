// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.114.3-943fbc81-20260603-173645
*/

package brsmigration

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
	"github.ibm.com/BackupAndRecovery/brs-migration-orchestrator/brsmigrationv2"
)

func ResourceIbmBrsMigrationVolume() *schema.Resource {
	return &schema.Resource{
		CreateContext:   resourceIbmBrsMigrationVolumeCreate,
		ReadContext:     resourceIbmBrsMigrationVolumeRead,
		UpdateContext:   resourceIbmBrsMigrationVolumeUpdate,
		DeleteContext:   resourceIbmBrsMigrationVolumeDelete,
		Importer: &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"migration_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				ValidateFunc: validate.InvokeValidator("ibm_brs_migration_volume", "migration_id"),
				Description: "The migration project ID (mgr-{uuid4} format).",
			},
			"env": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				ValidateFunc: validate.InvokeValidator("ibm_brs_migration_volume", "env"),
				Description: "Infrastructure environment this volume belongs to.",
			},
			"storage_type": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				ValidateFunc: validate.InvokeValidator("ibm_brs_migration_volume", "storage_type"),
				Description: "Storage type of the volume.",
			},
			"attachment_state": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Migration-service-computed attachment status. Set to `attached` when `host_attachments` is non-empty, `unattached` when empty.",
			},
			"storage": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Enriched storage details.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"global_identifier": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "Raw infrastructure volume ID. Classic: numeric string (e.g. \"98765432\"). VPC: UUID (e.g. r134-abcdef01-…).",
						},
						"name": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "Human-readable name of the volume as set in IBM Cloud.",
						},
						"capacity_gib": &schema.Schema{
							Type:        schema.TypeInt,
							Optional:    true,
							Computed:    true,
							Description: "Provisioned capacity in gibibytes (GiB).",
						},
						"iops": &schema.Schema{
							Type:        schema.TypeInt,
							Optional:    true,
							Computed:    true,
							Description: "Maximum I/O operations per second. VPC block: from VPC API (required). Classic: from Endurance/Performance tier. Absent for file/san/local volumes.",
						},
						"profile": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "Volume profile or tier. Classic: e.g. Endurance, Performance. VPC: e.g. general-purpose, 5iops-tier.",
						},
						"lifecycle_state": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "Normalised lifecycle state of the volume. All infrastructure-specific states are mapped to this canonical set before storage. Classic `ready` → `stable`; Classic `provisioning` → `pending`. VPC block `available` → `stable`; `pending_deletion` → `deleting`; `unusable` → `failed`. VPC file share states (`stable`, `pending`, `updating`, `deleting`, `suspended`, `waiting`, `failed`) map directly.",
						},
						"encryption": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "Encryption type for the volume. Classic: e.g. `aes256`. VPC block: `provider_managed` or `user_managed`. VPC file shares: `provider_managed` or `user_managed`.",
						},
						"throughput_mbps": &schema.Schema{
							Type:        schema.TypeInt,
							Optional:    true,
							Computed:    true,
							Description: "Maximum throughput in Mbps. VPC block: from the VPC API. Classic and VPC file shares: absent if not applicable.",
						},
						"source_paths": &schema.Schema{
							Type:        schema.TypeList,
							Optional:    true,
							Computed:    true,
							Description: "NFS mount paths for this volume. Present for file volumes only (classic/file and vpc/file); absent for block, SAN, and local volumes. Classic file: one or more export paths, no vpc_id or mount_target_id. VPC file: one entry per VPC mount target, each with vpc_id and mount_target_id. The orchestrator passes all entries to discover.sh; the script tries each in sequence until an active mount is found.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"path": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "NFS mount path (host:/export format).",
									},
									"vpc_id": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
										Description: "VPC ID this mount path belongs to. Present for VPC file volumes only; absent for Classic file volumes.",
									},
									"mount_target_id": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
										Description: "VPC file share mount target ID (share_mount_target_id from the VPC API). Present for VPC file volumes only; absent for Classic file volumes. Useful for lifecycle operations such as mount target teardown during workload completion.",
									},
								},
							},
						},
						"datacenter": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "Classic datacenter slug where this volume is provisioned (e.g. dal10).",
						},
						"iscsi_target_ips": &schema.Schema{
							Type:        schema.TypeList,
							Optional:    true,
							Computed:    true,
							Description: "iSCSI portal IPs for this volume. Classic block (iSCSI) volumes only; absent for Classic file, SAN, and local volumes. The orchestrator passes all IPs to discover.sh; the script tries each in sequence until an active session is found.",
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
						"region": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "VPC region (e.g. us-south).",
						},
						"zone": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "VPC zone (e.g. us-south-1).",
						},
						"crn": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "IBM Cloud CRN for this VPC volume.",
						},
						"resource_group_id": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "ID of the IBM Cloud resource group this volume belongs to.",
						},
						"replication_role": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "Replication role of the VPC file share. `none` means replication is not configured. `source` means this share is the replication source. `replica` means this share is the replication target (replica). File shares only.",
						},
						"access_control_mode": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "Access control mode for the VPC file share. Determines how mount target access is governed (`security_group` or `vpc`). Set to `none` for VPC block volumes and Classic volumes where this concept does not apply.",
						},
						"availability_mode": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "Availability zone scope of the VPC file share. `regional` means the share is accessible across all zones in its region. `zonal` means the share is pinned to a single zone (profile `dp2`). Set to `none` for VPC block volumes and Classic volumes where this concept does not apply.",
						},
					},
				},
			},
			"host_attachments": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Per-host attachment records for this volume.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"host_id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Migration service host ID (host-* prefix).",
						},
						"mount_path": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "OS-level mount path of the volume on this host (e.g. /mnt/data).",
						},
						"type": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "The filesystem type of this volume on the host (e.g. ext4, xfs, nfs4).",
						},
						"block_device": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "Block device path on this host. Present for block/local volumes; empty for file/NFS.",
						},
						"device_id": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "The storage device identifier for this volume on the host. For VPC block volumes, this is the volume attachment device ID. For Classic SAN and local volumes, this is the numeric device slot index. Present for VPC block, Classic SAN, and Classic local volumes; absent for Classic iSCSI block and all file volumes.",
						},
					},
				},
			},
			"migrated": &schema.Schema{
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Set to true when the workload covering this volume is completed.",
			},
			"workload_id": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "ID of the workload this volume is associated with.",
			},
			"registered_at": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Timestamp when this volume was registered in the Migration API.",
			},
			"volume_id": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Migration service volume ID (vol-{uuid4} format).",
			},
		},
	}
}

func ResourceIbmBrsMigrationVolumeValidator() *validate.ResourceValidator {
	validateSchema := make([]validate.ValidateSchema, 0)
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 "migration_id",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Required:                   true,
			Regexp:                     `^mgr-[0-9a-f-]{36}$`,
			MinValueLength:             40,
			MaxValueLength:             40,
		},
		validate.ValidateSchema{
			Identifier:                 "env",
			ValidateFunctionIdentifier: validate.ValidateAllowedStringValue,
			Type:                       validate.TypeString,
			Required:                   true,
			AllowedValues:              "classic, vpc",
		},
		validate.ValidateSchema{
			Identifier:                 "storage_type",
			ValidateFunctionIdentifier: validate.ValidateAllowedStringValue,
			Type:                       validate.TypeString,
			Required:                   true,
			AllowedValues:              "block, file, local, san",
		},
	)

	resourceValidator := validate.ResourceValidator{ResourceName: "ibm_brs_migration_volume", Schema: validateSchema}
	return &resourceValidator
}

func resourceIbmBrsMigrationVolumeCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	brsMigrationClient, err := meta.(conns.ClientSession).BrsMigrationV2()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_brs_migration_volume", "create", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	registerVolumeOptions := &brsmigrationv2.RegisterVolumeOptions{}

	registerVolumeOptions.SetMigrationID(d.Get("migration_id").(string))
	registerVolumeOptions.SetEnv(d.Get("env").(string))
	registerVolumeOptions.SetStorageType(d.Get("storage_type").(string))
	registerVolumeOptions.SetGlobalIdentifier(d.Get("global_identifier").(string))
	if _, ok := d.GetOk("datacenter"); ok {
		registerVolumeOptions.SetDatacenter(d.Get("datacenter").(string))
	}
	if _, ok := d.GetOk("region"); ok {
		registerVolumeOptions.SetRegion(d.Get("region").(string))
	}
	if _, ok := d.GetOk("zone"); ok {
		registerVolumeOptions.SetZone(d.Get("zone").(string))
	}

	volume, _, err := brsMigrationClient.RegisterVolumeWithContext(context, registerVolumeOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("RegisterVolumeWithContext failed: %s", err.Error()), "ibm_brs_migration_volume", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(fmt.Sprintf("%s/%s", *registerVolumeOptions.MigrationID, *volume.ID))

	return resourceIbmBrsMigrationVolumeRead(context, d, meta)
}

func resourceIbmBrsMigrationVolumeRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	brsMigrationClient, err := meta.(conns.ClientSession).BrsMigrationV2()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_brs_migration_volume", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	getVolumeOptions := &brsmigrationv2.GetVolumeOptions{}

	parts, err := flex.SepIdParts(d.Id(), "/")
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_brs_migration_volume", "read", "sep-id-parts").GetDiag()
	}

	getVolumeOptions.SetMigrationID(parts[0])
	getVolumeOptions.SetVolumeID(parts[1])

	volume, response, err := brsMigrationClient.GetVolumeWithContext(context, getVolumeOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetVolumeWithContext failed: %s", err.Error()), "ibm_brs_migration_volume", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	if err = d.Set("env", volume.Env); err != nil {
		err = fmt.Errorf("Error setting env: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_brs_migration_volume", "read", "set-env").GetDiag()
	}
	if err = d.Set("storage_type", volume.StorageType); err != nil {
		err = fmt.Errorf("Error setting storage_type: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_brs_migration_volume", "read", "set-storage_type").GetDiag()
	}
	if err = d.Set("attachment_state", volume.AttachmentState); err != nil {
		err = fmt.Errorf("Error setting attachment_state: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_brs_migration_volume", "read", "set-attachment_state").GetDiag()
	}
	storageMap, err := ResourceIbmBrsMigrationVolumeVolumeStorageToMap(volume.Storage)
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_brs_migration_volume", "read", "storage-to-map").GetDiag()
	}
	if err = d.Set("storage", []map[string]interface{}{storageMap}); err != nil {
		err = fmt.Errorf("Error setting storage: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_brs_migration_volume", "read", "set-storage").GetDiag()
	}
	hostAttachments := []map[string]interface{}{}
	for _, hostAttachmentsItem := range volume.HostAttachments {
		hostAttachmentsItemMap, err := ResourceIbmBrsMigrationVolumeHostAttachmentToMap(&hostAttachmentsItem) // #nosec G601
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_brs_migration_volume", "read", "host_attachments-to-map").GetDiag()
		}
		hostAttachments = append(hostAttachments, hostAttachmentsItemMap)
	}
	if err = d.Set("host_attachments", hostAttachments); err != nil {
		err = fmt.Errorf("Error setting host_attachments: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_brs_migration_volume", "read", "set-host_attachments").GetDiag()
	}
	if err = d.Set("migrated", volume.Migrated); err != nil {
		err = fmt.Errorf("Error setting migrated: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_brs_migration_volume", "read", "set-migrated").GetDiag()
	}
	if !core.IsNil(volume.WorkloadID) {
		if err = d.Set("workload_id", volume.WorkloadID); err != nil {
			err = fmt.Errorf("Error setting workload_id: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_brs_migration_volume", "read", "set-workload_id").GetDiag()
		}
	}
	if err = d.Set("registered_at", flex.DateTimeToString(volume.RegisteredAt)); err != nil {
		err = fmt.Errorf("Error setting registered_at: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_brs_migration_volume", "read", "set-registered_at").GetDiag()
	}
	if err = d.Set("volume_id", volume.ID); err != nil {
		err = fmt.Errorf("Error setting volume_id: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_brs_migration_volume", "read", "set-volume_id").GetDiag()
	}

	return nil
}

func resourceIbmBrsMigrationVolumeUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	brsMigrationClient, err := meta.(conns.ClientSession).BrsMigrationV2()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_brs_migration_volume", "update", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	updateVolumeOptions := &brsmigrationv2.UpdateVolumeOptions{}

	parts, err := flex.SepIdParts(d.Id(), "/")
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_brs_migration_volume", "update", "sep-id-parts").GetDiag()
	}

	updateVolumeOptions.SetMigrationID(parts[0])
	updateVolumeOptions.SetVolumeID(parts[1])

	hasChange := false

	if d.HasChange("migration_id") {
		errMsg := fmt.Sprintf("Cannot update resource property \"%s\" with the ForceNew annotation." +
			" The resource must be re-created to update this property.", "migration_id")
		return flex.DiscriminatedTerraformErrorf(nil, errMsg, "ibm_brs_migration_volume", "update", "migration_id-forces-new").GetDiag()
	}
	if d.HasChange("host_attachments") {
		var hostAttachments []brsmigrationv2.HostAttachment
		for _, v := range d.Get("host_attachments").([]interface{}) {
			value := v.(map[string]interface{})
			hostAttachmentsItem, err := ResourceIbmBrsMigrationVolumeMapToHostAttachment(value)
			if err != nil {
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_brs_migration_volume", "update", "parse-host_attachments").GetDiag()
			}
			hostAttachments = append(hostAttachments, *hostAttachmentsItem)
		}
		updateVolumeOptions.SetHostAttachments(hostAttachments)
		hasChange = true
	}

	if hasChange {
		_, _, err = brsMigrationClient.UpdateVolumeWithContext(context, updateVolumeOptions)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("UpdateVolumeWithContext failed: %s", err.Error()), "ibm_brs_migration_volume", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
	}

	return resourceIbmBrsMigrationVolumeRead(context, d, meta)
}

func resourceIbmBrsMigrationVolumeDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	brsMigrationClient, err := meta.(conns.ClientSession).BrsMigrationV2()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_brs_migration_volume", "delete", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	deleteVolumeOptions := &brsmigrationv2.DeleteVolumeOptions{}

	parts, err := flex.SepIdParts(d.Id(), "/")
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_brs_migration_volume", "delete", "sep-id-parts").GetDiag()
	}

	deleteVolumeOptions.SetMigrationID(parts[0])
	deleteVolumeOptions.SetVolumeID(parts[1])

	_, err = brsMigrationClient.DeleteVolumeWithContext(context, deleteVolumeOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("DeleteVolumeWithContext failed: %s", err.Error()), "ibm_brs_migration_volume", "delete")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId("")

	return nil
}

func ResourceIbmBrsMigrationVolumeMapToHostAttachment(modelMap map[string]interface{}) (*brsmigrationv2.HostAttachment, error) {
	model := &brsmigrationv2.HostAttachment{}
	model.HostID = core.StringPtr(modelMap["host_id"].(string))
	if modelMap["mount_path"] != nil && modelMap["mount_path"].(string) != "" {
		model.MountPath = core.StringPtr(modelMap["mount_path"].(string))
	}
	if modelMap["type"] != nil && modelMap["type"].(string) != "" {
		model.Type = core.StringPtr(modelMap["type"].(string))
	}
	if modelMap["block_device"] != nil && modelMap["block_device"].(string) != "" {
		model.BlockDevice = core.StringPtr(modelMap["block_device"].(string))
	}
	if modelMap["device_id"] != nil && modelMap["device_id"].(string) != "" {
		model.DeviceID = core.StringPtr(modelMap["device_id"].(string))
	}
	return model, nil
}

func ResourceIbmBrsMigrationVolumeVolumeStorageToMap(model brsmigrationv2.VolumeStorageIntf) (map[string]interface{}, error) {
	if _, ok := model.(*brsmigrationv2.VolumeStorageClassicVolumeStorageDetails); ok {
		return ResourceIbmBrsMigrationVolumeVolumeStorageClassicVolumeStorageDetailsToMap(model.(*brsmigrationv2.VolumeStorageClassicVolumeStorageDetails))
	} else if _, ok := model.(*brsmigrationv2.VolumeStorageVPCVolumeStorageDetails); ok {
		return ResourceIbmBrsMigrationVolumeVolumeStorageVPCVolumeStorageDetailsToMap(model.(*brsmigrationv2.VolumeStorageVPCVolumeStorageDetails))
	} else if _, ok := model.(*brsmigrationv2.VolumeStorage); ok {
		modelMap := make(map[string]interface{})
		model := model.(*brsmigrationv2.VolumeStorage)
		if model.GlobalIdentifier != nil {
			modelMap["global_identifier"] = *model.GlobalIdentifier
		}
		if model.Name != nil {
			modelMap["name"] = *model.Name
		}
		if model.CapacityGib != nil {
			modelMap["capacity_gib"] = flex.IntValue(model.CapacityGib)
		}
		if model.Iops != nil {
			modelMap["iops"] = flex.IntValue(model.Iops)
		}
		if model.Profile != nil {
			modelMap["profile"] = *model.Profile
		}
		if model.LifecycleState != nil {
			modelMap["lifecycle_state"] = *model.LifecycleState
		}
		if model.Encryption != nil {
			modelMap["encryption"] = *model.Encryption
		}
		if model.ThroughputMbps != nil {
			modelMap["throughput_mbps"] = flex.IntValue(model.ThroughputMbps)
		}
		if model.SourcePaths != nil {
			sourcePaths := []map[string]interface{}{}
			for _, sourcePathsItem := range model.SourcePaths {
				sourcePathsItemMap, err := ResourceIbmBrsMigrationVolumeSourcePathToMap(&sourcePathsItem) // #nosec G601
				if err != nil {
					return modelMap, err
				}
				sourcePaths = append(sourcePaths, sourcePathsItemMap)
			}
			modelMap["source_paths"] = sourcePaths
		}
		if model.Datacenter != nil {
			modelMap["datacenter"] = *model.Datacenter
		}
		if model.IscsiTargetIps != nil {
			modelMap["iscsi_target_ips"] = model.IscsiTargetIps
		}
		if model.Region != nil {
			modelMap["region"] = *model.Region
		}
		if model.Zone != nil {
			modelMap["zone"] = *model.Zone
		}
		if model.Crn != nil {
			modelMap["crn"] = *model.Crn
		}
		if model.ResourceGroupID != nil {
			modelMap["resource_group_id"] = *model.ResourceGroupID
		}
		if model.ReplicationRole != nil {
			modelMap["replication_role"] = *model.ReplicationRole
		}
		if model.AccessControlMode != nil {
			modelMap["access_control_mode"] = *model.AccessControlMode
		}
		if model.AvailabilityMode != nil {
			modelMap["availability_mode"] = *model.AvailabilityMode
		}
		return modelMap, nil
	} else {
		return nil, fmt.Errorf("Unrecognized brsmigrationv2.VolumeStorageIntf subtype encountered")
	}
}

func ResourceIbmBrsMigrationVolumeSourcePathToMap(model *brsmigrationv2.SourcePath) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["path"] = *model.Path
	if model.VpcID != nil {
		modelMap["vpc_id"] = *model.VpcID
	}
	if model.MountTargetID != nil {
		modelMap["mount_target_id"] = *model.MountTargetID
	}
	return modelMap, nil
}

func ResourceIbmBrsMigrationVolumeVolumeStorageClassicVolumeStorageDetailsToMap(model *brsmigrationv2.VolumeStorageClassicVolumeStorageDetails) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["global_identifier"] = *model.GlobalIdentifier
	if model.Name != nil {
		modelMap["name"] = *model.Name
	}
	if model.CapacityGib != nil {
		modelMap["capacity_gib"] = flex.IntValue(model.CapacityGib)
	}
	if model.Iops != nil {
		modelMap["iops"] = flex.IntValue(model.Iops)
	}
	if model.Profile != nil {
		modelMap["profile"] = *model.Profile
	}
	modelMap["lifecycle_state"] = *model.LifecycleState
	if model.Encryption != nil {
		modelMap["encryption"] = *model.Encryption
	}
	if model.ThroughputMbps != nil {
		modelMap["throughput_mbps"] = flex.IntValue(model.ThroughputMbps)
	}
	sourcePaths := []map[string]interface{}{}
	for _, sourcePathsItem := range model.SourcePaths {
		sourcePathsItemMap, err := ResourceIbmBrsMigrationVolumeSourcePathToMap(&sourcePathsItem) // #nosec G601
		if err != nil {
			return modelMap, err
		}
		sourcePaths = append(sourcePaths, sourcePathsItemMap)
	}
	modelMap["source_paths"] = sourcePaths
	modelMap["datacenter"] = *model.Datacenter
	modelMap["iscsi_target_ips"] = model.IscsiTargetIps
	return modelMap, nil
}

func ResourceIbmBrsMigrationVolumeVolumeStorageVPCVolumeStorageDetailsToMap(model *brsmigrationv2.VolumeStorageVPCVolumeStorageDetails) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["global_identifier"] = *model.GlobalIdentifier
	if model.Name != nil {
		modelMap["name"] = *model.Name
	}
	if model.CapacityGib != nil {
		modelMap["capacity_gib"] = flex.IntValue(model.CapacityGib)
	}
	if model.Iops != nil {
		modelMap["iops"] = flex.IntValue(model.Iops)
	}
	if model.Profile != nil {
		modelMap["profile"] = *model.Profile
	}
	modelMap["lifecycle_state"] = *model.LifecycleState
	if model.Encryption != nil {
		modelMap["encryption"] = *model.Encryption
	}
	if model.ThroughputMbps != nil {
		modelMap["throughput_mbps"] = flex.IntValue(model.ThroughputMbps)
	}
	sourcePaths := []map[string]interface{}{}
	for _, sourcePathsItem := range model.SourcePaths {
		sourcePathsItemMap, err := ResourceIbmBrsMigrationVolumeSourcePathToMap(&sourcePathsItem) // #nosec G601
		if err != nil {
			return modelMap, err
		}
		sourcePaths = append(sourcePaths, sourcePathsItemMap)
	}
	modelMap["source_paths"] = sourcePaths
	modelMap["region"] = *model.Region
	modelMap["zone"] = *model.Zone
	if model.Crn != nil {
		modelMap["crn"] = *model.Crn
	}
	if model.ResourceGroupID != nil {
		modelMap["resource_group_id"] = *model.ResourceGroupID
	}
	modelMap["replication_role"] = *model.ReplicationRole
	modelMap["access_control_mode"] = *model.AccessControlMode
	modelMap["availability_mode"] = *model.AvailabilityMode
	return modelMap, nil
}

func ResourceIbmBrsMigrationVolumeHostAttachmentToMap(model *brsmigrationv2.HostAttachment) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["host_id"] = *model.HostID
	if model.MountPath != nil {
		modelMap["mount_path"] = *model.MountPath
	}
	if model.Type != nil {
		modelMap["type"] = *model.Type
	}
	if model.BlockDevice != nil {
		modelMap["block_device"] = *model.BlockDevice
	}
	if model.DeviceID != nil {
		modelMap["device_id"] = *model.DeviceID
	}
	return modelMap, nil
}
