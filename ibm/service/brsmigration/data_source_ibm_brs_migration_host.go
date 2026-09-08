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
	"github.com/IBM/ibm-brs-migration-sdk-go/brsmigrationv1"
)

func DataSourceIbmBrsMigrationHost() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIbmBrsMigrationHostRead,

		Schema: map[string]*schema.Schema{
			"migration_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The migration project ID (mgr-{uuid4} format).",
			},
			"host_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The migration service host ID (host-{uuid4} format).",
			},
			"type": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Whether the host is a Virtual Server Instance or bare metal server.",
			},
			"env": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Infrastructure environment this host belongs to.",
			},
			"compute": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Enriched compute details. Schema variant matches the sibling `env` field: `classic` → `ClassicComputeDetails`, `vpc` → `VPCComputeDetails`.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"status": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Current power/lifecycle status (union of VPC Virtual Server Instance and bare metal status enums).",
						},
						"os_family": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "OS family of the instance.",
						},
						"global_identifier": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "GUID that uniquely identifies this instance in the infrastructure.",
						},
						"throughput_mbps": &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Total network throughput in Mbps. Required on both VPC Virtual Server Instance and bare metal.",
						},
						"public_ips": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Public IP addresses on this host. For VPC this maps from `floating_ips` on the primary network interface. For Classic this maps from `primaryPublicIpAddress`. Empty array when none.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Display name or hostname of the instance.",
						},
						"os_type": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "OS image identifier as returned by the infrastructure API.",
						},
						"ip_address": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Primary IP address of the instance.",
						},
						"profile": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Instance profile name.",
						},
						"vcpu_count": &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Number of virtual CPUs (from vcpu.count on VPC Virtual Server Instance).",
						},
						"memory_gib": &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Memory in gibibytes (GiB).",
						},
						"image_id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Boot image ID. Optional — VPC Virtual Server Instance only, not present on bare metal.",
						},
						"datacenter": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Classic datacenter (e.g. dal10).",
						},
						"region": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "VPC region.",
						},
						"zone": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "VPC zone.",
						},
						"lifecycle_state": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Lifecycle state as returned by the VPC API. Present on both VPC Virtual Server Instance and bare metal.",
						},
						"health_state": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Health state as reported by the VPC API. Same enum on both Virtual Server Instance and bare metal.",
						},
						"cpu_architecture": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "CPU architecture of the instance.",
						},
						"subnet_id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "ID of the subnet the primary network interface is attached to.",
						},
						"security_groups": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Security group IDs on the primary network interface. Empty array when none.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"resource_group_id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "ID of the IBM Cloud resource group this instance belongs to.",
						},
						"crn": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "IBM Cloud Resource Name for this instance.",
						},
						"vpc_id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "ID of the VPC this instance belongs to.",
						},
						"vpc_name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Name of the VPC.",
						},
						"boot_volume_id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Migration service `vol-*` ID of the boot volume attachment.",
						},
					},
				},
			},
			"volume_attachments": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Per-volume attachment records for this host.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"volume_id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Migration service volume ID (vol-* prefix).",
						},
					},
				},
			},
			"migrated": &schema.Schema{
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Set to true when `POST /migrations/{migration_id}/workloads/{workload_id}/complete` is called for a workload that includes this host.",
			},
			"workload_id": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "ID of the workload this host is associated with. Null when the host has not been added to any workload yet.",
			},
			"registered_at": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Timestamp when this host was registered in the Migration API.",
			},
		},
	}
}

func dataSourceIbmBrsMigrationHostRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	brsMigrationClient, err := meta.(conns.ClientSession).BrsMigrationV1()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_brs_migration_host", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	getHostOptions := &brsmigrationv1.GetHostOptions{}

	getHostOptions.SetMigrationID(d.Get("migration_id").(string))
	getHostOptions.SetHostID(d.Get("host_id").(string))

	host, _, err := brsMigrationClient.GetHostWithContext(context, getHostOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetHostWithContext failed: %s", err.Error()), "(Data) ibm_brs_migration_host", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(fmt.Sprintf("%s/%s", *getHostOptions.MigrationID, *getHostOptions.HostID))

	if err = d.Set("type", host.Type); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting type: %s", err), "(Data) ibm_brs_migration_host", "read", "set-type").GetDiag()
	}

	if err = d.Set("env", host.Env); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting env: %s", err), "(Data) ibm_brs_migration_host", "read", "set-env").GetDiag()
	}

	compute := []map[string]interface{}{}
	computeMap, err := DataSourceIbmBrsMigrationHostHostComputeToMap(host.Compute)
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_brs_migration_host", "read", "compute-to-map").GetDiag()
	}
	compute = append(compute, computeMap)
	if err = d.Set("compute", compute); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting compute: %s", err), "(Data) ibm_brs_migration_host", "read", "set-compute").GetDiag()
	}

	volumeAttachments := []map[string]interface{}{}
	for _, volumeAttachmentsItem := range host.VolumeAttachments {
		volumeAttachmentsItemMap, err := DataSourceIbmBrsMigrationHostVolumeAttachmentToMap(&volumeAttachmentsItem) // #nosec G601
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_brs_migration_host", "read", "volume_attachments-to-map").GetDiag()
		}
		volumeAttachments = append(volumeAttachments, volumeAttachmentsItemMap)
	}
	if err = d.Set("volume_attachments", volumeAttachments); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting volume_attachments: %s", err), "(Data) ibm_brs_migration_host", "read", "set-volume_attachments").GetDiag()
	}

	if err = d.Set("migrated", host.Migrated); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting migrated: %s", err), "(Data) ibm_brs_migration_host", "read", "set-migrated").GetDiag()
	}

	if !core.IsNil(host.WorkloadID) {
		if err = d.Set("workload_id", host.WorkloadID); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting workload_id: %s", err), "(Data) ibm_brs_migration_host", "read", "set-workload_id").GetDiag()
		}
	}

	if err = d.Set("registered_at", flex.DateTimeToString(host.RegisteredAt)); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting registered_at: %s", err), "(Data) ibm_brs_migration_host", "read", "set-registered_at").GetDiag()
	}

	return nil
}

func DataSourceIbmBrsMigrationHostHostComputeToMap(model brsmigrationv1.HostComputeIntf) (map[string]interface{}, error) {
	if _, ok := model.(*brsmigrationv1.HostComputeClassicComputeDetails); ok {
		return DataSourceIbmBrsMigrationHostHostComputeClassicComputeDetailsToMap(model.(*brsmigrationv1.HostComputeClassicComputeDetails))
	} else if _, ok := model.(*brsmigrationv1.HostComputeVPCComputeDetails); ok {
		return DataSourceIbmBrsMigrationHostHostComputeVPCComputeDetailsToMap(model.(*brsmigrationv1.HostComputeVPCComputeDetails))
	} else if _, ok := model.(*brsmigrationv1.HostCompute); ok {
		modelMap := make(map[string]interface{})
		model := model.(*brsmigrationv1.HostCompute)
		if model.Status != nil {
			modelMap["status"] = *model.Status
		}
		if model.OsFamily != nil {
			modelMap["os_family"] = *model.OsFamily
		}
		if model.GlobalIdentifier != nil {
			modelMap["global_identifier"] = *model.GlobalIdentifier
		}
		if model.ThroughputMbps != nil {
			modelMap["throughput_mbps"] = flex.IntValue(model.ThroughputMbps)
		}
		if model.PublicIps != nil {
			modelMap["public_ips"] = model.PublicIps
		}
		if model.Name != nil {
			modelMap["name"] = *model.Name
		}
		if model.OsType != nil {
			modelMap["os_type"] = *model.OsType
		}
		if model.IpAddress != nil {
			modelMap["ip_address"] = *model.IpAddress
		}
		if model.Profile != nil {
			modelMap["profile"] = *model.Profile
		}
		if model.VcpuCount != nil {
			modelMap["vcpu_count"] = flex.IntValue(model.VcpuCount)
		}
		if model.MemoryGib != nil {
			modelMap["memory_gib"] = flex.IntValue(model.MemoryGib)
		}
		if model.ImageID != nil {
			modelMap["image_id"] = *model.ImageID
		}
		if model.Datacenter != nil {
			modelMap["datacenter"] = *model.Datacenter
		}
		if model.Region != nil {
			modelMap["region"] = *model.Region
		}
		if model.Zone != nil {
			modelMap["zone"] = *model.Zone
		}
		if model.LifecycleState != nil {
			modelMap["lifecycle_state"] = *model.LifecycleState
		}
		if model.HealthState != nil {
			modelMap["health_state"] = *model.HealthState
		}
		if model.CpuArchitecture != nil {
			modelMap["cpu_architecture"] = *model.CpuArchitecture
		}
		if model.SubnetID != nil {
			modelMap["subnet_id"] = *model.SubnetID
		}
		if model.SecurityGroups != nil {
			modelMap["security_groups"] = model.SecurityGroups
		}
		if model.ResourceGroupID != nil {
			modelMap["resource_group_id"] = *model.ResourceGroupID
		}
		if model.Crn != nil {
			modelMap["crn"] = *model.Crn
		}
		if model.VpcID != nil {
			modelMap["vpc_id"] = *model.VpcID
		}
		if model.VpcName != nil {
			modelMap["vpc_name"] = *model.VpcName
		}
		if model.BootVolumeID != nil {
			modelMap["boot_volume_id"] = *model.BootVolumeID
		}
		return modelMap, nil
	} else {
		return nil, fmt.Errorf("Unrecognized brsmigrationv1.HostComputeIntf subtype encountered")
	}
}

func DataSourceIbmBrsMigrationHostHostComputeClassicComputeDetailsToMap(model *brsmigrationv1.HostComputeClassicComputeDetails) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["status"] = *model.Status
	modelMap["os_family"] = *model.OsFamily
	modelMap["global_identifier"] = *model.GlobalIdentifier
	modelMap["throughput_mbps"] = flex.IntValue(model.ThroughputMbps)
	modelMap["public_ips"] = model.PublicIps
	if model.Name != nil {
		modelMap["name"] = *model.Name
	}
	if model.OsType != nil {
		modelMap["os_type"] = *model.OsType
	}
	if model.IpAddress != nil {
		modelMap["ip_address"] = *model.IpAddress
	}
	if model.Profile != nil {
		modelMap["profile"] = *model.Profile
	}
	if model.VcpuCount != nil {
		modelMap["vcpu_count"] = flex.IntValue(model.VcpuCount)
	}
	if model.MemoryGib != nil {
		modelMap["memory_gib"] = flex.IntValue(model.MemoryGib)
	}
	if model.ImageID != nil {
		modelMap["image_id"] = *model.ImageID
	}
	modelMap["datacenter"] = *model.Datacenter
	return modelMap, nil
}

func DataSourceIbmBrsMigrationHostHostComputeVPCComputeDetailsToMap(model *brsmigrationv1.HostComputeVPCComputeDetails) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["status"] = *model.Status
	modelMap["os_family"] = *model.OsFamily
	modelMap["global_identifier"] = *model.GlobalIdentifier
	modelMap["throughput_mbps"] = flex.IntValue(model.ThroughputMbps)
	modelMap["public_ips"] = model.PublicIps
	if model.Name != nil {
		modelMap["name"] = *model.Name
	}
	if model.OsType != nil {
		modelMap["os_type"] = *model.OsType
	}
	if model.IpAddress != nil {
		modelMap["ip_address"] = *model.IpAddress
	}
	if model.Profile != nil {
		modelMap["profile"] = *model.Profile
	}
	if model.VcpuCount != nil {
		modelMap["vcpu_count"] = flex.IntValue(model.VcpuCount)
	}
	if model.MemoryGib != nil {
		modelMap["memory_gib"] = flex.IntValue(model.MemoryGib)
	}
	if model.ImageID != nil {
		modelMap["image_id"] = *model.ImageID
	}
	modelMap["region"] = *model.Region
	modelMap["zone"] = *model.Zone
	modelMap["lifecycle_state"] = *model.LifecycleState
	modelMap["health_state"] = *model.HealthState
	modelMap["cpu_architecture"] = *model.CpuArchitecture
	modelMap["subnet_id"] = *model.SubnetID
	modelMap["security_groups"] = model.SecurityGroups
	modelMap["resource_group_id"] = *model.ResourceGroupID
	if model.Crn != nil {
		modelMap["crn"] = *model.Crn
	}
	if model.VpcID != nil {
		modelMap["vpc_id"] = *model.VpcID
	}
	if model.VpcName != nil {
		modelMap["vpc_name"] = *model.VpcName
	}
	if model.BootVolumeID != nil {
		modelMap["boot_volume_id"] = *model.BootVolumeID
	}
	return modelMap, nil
}

func DataSourceIbmBrsMigrationHostVolumeAttachmentToMap(model *brsmigrationv1.VolumeAttachment) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["volume_id"] = *model.VolumeID
	return modelMap, nil
}
