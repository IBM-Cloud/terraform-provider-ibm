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
	"github.com/IBM/go-sdk-core/v5/core"
	"github.ibm.com/BackupAndRecovery/brs-migration-orchestrator/brsmigrationv2"
)

func DataSourceIbmBrsMigrationVolume() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIbmBrsMigrationVolumeRead,

		Schema: map[string]*schema.Schema{
			"migration_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The migration project ID (mgr-{uuid4} format).",
			},
			"volume_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The migration service volume ID (vol-{uuid4} format).",
			},
			"env": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Infrastructure environment this volume belongs to.",
			},
			"storage_type": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
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
							Computed:    true,
							Description: "Raw infrastructure volume ID. Classic: numeric string (e.g. \"98765432\"). VPC: UUID (e.g. r134-abcdef01-…).",
						},
						"name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Human-readable name of the volume as set in IBM Cloud.",
						},
						"capacity_gib": &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Provisioned capacity in gibibytes (GiB).",
						},
						"iops": &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Maximum I/O operations per second. VPC block: from VPC API (required). Classic: from Endurance/Performance tier. Absent for file/san/local volumes.",
						},
						"profile": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Volume profile or tier. Classic: e.g. Endurance, Performance. VPC: e.g. general-purpose, 5iops-tier.",
						},
						"lifecycle_state": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Normalised lifecycle state of the volume. All infrastructure-specific states are mapped to this canonical set before storage. Classic `ready` → `stable`; Classic `provisioning` → `pending`. VPC block `available` → `stable`; `pending_deletion` → `deleting`; `unusable` → `failed`. VPC file share states (`stable`, `pending`, `updating`, `deleting`, `suspended`, `waiting`, `failed`) map directly.",
						},
						"encryption": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Encryption type for the volume. Classic: e.g. `aes256`. VPC block: `provider_managed` or `user_managed`. VPC file shares: `provider_managed` or `user_managed`.",
						},
						"throughput_mbps": &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Maximum throughput in Mbps. VPC block: from the VPC API. Classic and VPC file shares: absent if not applicable.",
						},
						"source_paths": &schema.Schema{
							Type:        schema.TypeList,
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
										Computed:    true,
										Description: "VPC ID this mount path belongs to. Present for VPC file volumes only; absent for Classic file volumes.",
									},
									"mount_target_id": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "VPC file share mount target ID (share_mount_target_id from the VPC API). Present for VPC file volumes only; absent for Classic file volumes. Useful for lifecycle operations such as mount target teardown during workload completion.",
									},
								},
							},
						},
						"datacenter": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Classic datacenter slug where this volume is provisioned (e.g. dal10).",
						},
						"iscsi_target_ips": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "iSCSI portal IPs for this volume. Classic block (iSCSI) volumes only; absent for Classic file, SAN, and local volumes. The orchestrator passes all IPs to discover.sh; the script tries each in sequence until an active session is found.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"region": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "VPC region (e.g. us-south).",
						},
						"zone": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "VPC zone (e.g. us-south-1).",
						},
						"crn": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "IBM Cloud CRN for this VPC volume.",
						},
						"resource_group_id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "ID of the IBM Cloud resource group this volume belongs to.",
						},
						"replication_role": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Replication role of the VPC file share. `none` means replication is not configured. `source` means this share is the replication source. `replica` means this share is the replication target (replica). File shares only.",
						},
						"access_control_mode": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Access control mode for the VPC file share. Determines how mount target access is governed (`security_group` or `vpc`). Set to `none` for VPC block volumes and Classic volumes where this concept does not apply.",
						},
						"availability_mode": &schema.Schema{
							Type:        schema.TypeString,
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
							Computed:    true,
							Description: "OS-level mount path of the volume on this host (e.g. /mnt/data).",
						},
						"type": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The filesystem type of this volume on the host (e.g. ext4, xfs, nfs4).",
						},
						"block_device": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Block device path on this host. Present for block/local volumes; empty for file/NFS.",
						},
						"device_id": &schema.Schema{
							Type:        schema.TypeString,
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
		},
	}
}

func dataSourceIbmBrsMigrationVolumeRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	brsMigrationClient, err := meta.(conns.ClientSession).BrsMigrationV2()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_brs_migration_volume", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	getVolumeOptions := &brsmigrationv2.GetVolumeOptions{}

	getVolumeOptions.SetMigrationID(d.Get("migration_id").(string))
	getVolumeOptions.SetVolumeID(d.Get("volume_id").(string))

	volume, _, err := brsMigrationClient.GetVolumeWithContext(context, getVolumeOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetVolumeWithContext failed: %s", err.Error()), "(Data) ibm_brs_migration_volume", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(fmt.Sprintf("%s/%s", *getVolumeOptions.MigrationID, *getVolumeOptions.VolumeID))

	if err = d.Set("env", volume.Env); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting env: %s", err), "(Data) ibm_brs_migration_volume", "read", "set-env").GetDiag()
	}

	if err = d.Set("storage_type", volume.StorageType); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting storage_type: %s", err), "(Data) ibm_brs_migration_volume", "read", "set-storage_type").GetDiag()
	}

	if err = d.Set("attachment_state", volume.AttachmentState); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting attachment_state: %s", err), "(Data) ibm_brs_migration_volume", "read", "set-attachment_state").GetDiag()
	}

	storage := []map[string]interface{}{}
	storageMap, err := DataSourceIbmBrsMigrationVolumeVolumeStorageToMap(volume.Storage)
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_brs_migration_volume", "read", "storage-to-map").GetDiag()
	}
	storage = append(storage, storageMap)
	if err = d.Set("storage", storage); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting storage: %s", err), "(Data) ibm_brs_migration_volume", "read", "set-storage").GetDiag()
	}

	hostAttachments := []map[string]interface{}{}
	for _, hostAttachmentsItem := range volume.HostAttachments {
		hostAttachmentsItemMap, err := DataSourceIbmBrsMigrationVolumeHostAttachmentToMap(&hostAttachmentsItem) // #nosec G601
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_brs_migration_volume", "read", "host_attachments-to-map").GetDiag()
		}
		hostAttachments = append(hostAttachments, hostAttachmentsItemMap)
	}
	if err = d.Set("host_attachments", hostAttachments); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting host_attachments: %s", err), "(Data) ibm_brs_migration_volume", "read", "set-host_attachments").GetDiag()
	}

	if err = d.Set("migrated", volume.Migrated); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting migrated: %s", err), "(Data) ibm_brs_migration_volume", "read", "set-migrated").GetDiag()
	}

	if !core.IsNil(volume.WorkloadID) {
		if err = d.Set("workload_id", volume.WorkloadID); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting workload_id: %s", err), "(Data) ibm_brs_migration_volume", "read", "set-workload_id").GetDiag()
		}
	}

	if err = d.Set("registered_at", flex.DateTimeToString(volume.RegisteredAt)); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting registered_at: %s", err), "(Data) ibm_brs_migration_volume", "read", "set-registered_at").GetDiag()
	}

	return nil
}

func DataSourceIbmBrsMigrationVolumeVolumeStorageToMap(model brsmigrationv2.VolumeStorageIntf) (map[string]interface{}, error) {
	if _, ok := model.(*brsmigrationv2.VolumeStorageClassicVolumeStorageDetails); ok {
		return DataSourceIbmBrsMigrationVolumeVolumeStorageClassicVolumeStorageDetailsToMap(model.(*brsmigrationv2.VolumeStorageClassicVolumeStorageDetails))
	} else if _, ok := model.(*brsmigrationv2.VolumeStorageVPCVolumeStorageDetails); ok {
		return DataSourceIbmBrsMigrationVolumeVolumeStorageVPCVolumeStorageDetailsToMap(model.(*brsmigrationv2.VolumeStorageVPCVolumeStorageDetails))
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
				sourcePathsItemMap, err := DataSourceIbmBrsMigrationVolumeSourcePathToMap(&sourcePathsItem) // #nosec G601
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

func DataSourceIbmBrsMigrationVolumeSourcePathToMap(model *brsmigrationv2.SourcePath) (map[string]interface{}, error) {
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

func DataSourceIbmBrsMigrationVolumeVolumeStorageClassicVolumeStorageDetailsToMap(model *brsmigrationv2.VolumeStorageClassicVolumeStorageDetails) (map[string]interface{}, error) {
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
		sourcePathsItemMap, err := DataSourceIbmBrsMigrationVolumeSourcePathToMap(&sourcePathsItem) // #nosec G601
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

func DataSourceIbmBrsMigrationVolumeVolumeStorageVPCVolumeStorageDetailsToMap(model *brsmigrationv2.VolumeStorageVPCVolumeStorageDetails) (map[string]interface{}, error) {
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
		sourcePathsItemMap, err := DataSourceIbmBrsMigrationVolumeSourcePathToMap(&sourcePathsItem) // #nosec G601
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

func DataSourceIbmBrsMigrationVolumeHostAttachmentToMap(model *brsmigrationv2.HostAttachment) (map[string]interface{}, error) {
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
