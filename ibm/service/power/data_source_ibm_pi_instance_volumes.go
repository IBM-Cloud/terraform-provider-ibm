// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package power

import (
	"context"
	"fmt"
	"log"

	"github.com/IBM-Cloud/power-go-client/clients/instance"
	"github.com/IBM-Cloud/power-go-client/power/models"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/hashicorp/go-uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func DataSourceIBMPIInstanceVolumes() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMPIInstanceVolumesRead,
		Schema: map[string]*schema.Schema{
			// Arguments
			Arg_CloudInstanceID: {
				Description:  "The GUID of the service instance associated with an account.",
				Required:     true,
				Type:         schema.TypeString,
				ValidateFunc: validation.NoZeroValues,
			},
			Arg_InstanceID: {
				AtLeastOneOf:  []string{Arg_InstanceID, Arg_InstanceName},
				ConflictsWith: []string{Arg_InstanceName},
				Description:   "The PVM instance ID.",
				Optional:      true,
				Type:          schema.TypeString,
				ValidateFunc:  validation.NoZeroValues,
			},
			Arg_InstanceName: {
				AtLeastOneOf:  []string{Arg_InstanceID, Arg_InstanceName},
				ConflictsWith: []string{Arg_InstanceID},
				Description:   "The name of the PVM instance.",
				Optional:      true,
				Type:          schema.TypeString,
				ValidateFunc:  validation.NoZeroValues,
			},

			// Attribute
			Attr_BootVolumeID: {
				Computed:    true,
				Description: "The unique identifier of the boot volume.",
				Type:        schema.TypeString,
			},
			Attr_InstanceVolumes: {
				Computed:    true,
				Description: "List of volumes attached to instance.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						Attr_Auxiliary: {
							Computed:    true,
							Description: "Indicates if the volume is auxiliary or not.",
							Type:        schema.TypeBool,
						},
						Attr_AuxiliaryVolumeName: {
							Computed:    true,
							Description: "The auxiliary volume name.",
							Type:        schema.TypeString,
						},
						Attr_Bootable: {
							Computed:    true,
							Description: "Indicates if the volume is boot capable.",
							Type:        schema.TypeBool,
						},
						Attr_ConsistencyGroupName: {
							Computed:    true,
							Description: "The name of consistency group at storage controller level.",
							Type:        schema.TypeString,
						},
						Attr_CreationDate: {
							Computed:    true,
							Description: "Date volume was created.",
							Type:        schema.TypeString,
						},
						Attr_CRN: {
							Computed:    true,
							Description: "The CRN of this resource.",
							Type:        schema.TypeString,
						},
						Attr_DeleteOnTermination: {
							Computed:    true,
							Description: "Indicates if the volume should be deleted when the server terminates.",
							Type:        schema.TypeBool,
						},
						Attr_FreezeTime: {
							Computed:    true,
							Description: "The freeze time of remote copy.",
							Type:        schema.TypeString,
						},
						Attr_GroupID: {
							Computed:    true,
							Description: "The volume group id in which the volume belongs.",
							Type:        schema.TypeString,
						},
						Attr_Href: {
							Computed:    true,
							Description: "The hyper link of the volume.",
							Type:        schema.TypeString,
						},
						Attr_ID: {
							Computed:    true,
							Description: "The unique identifier of the volume.",
							Type:        schema.TypeString,
						},
						Attr_IOThrottleRate: {
							Computed:    true,
							Description: "Amount of iops assigned to the volume",
							Type:        schema.TypeString,
						},
						Attr_LastUpdateDate: {
							Computed:    true,
							Description: "The last updated date of the volume.",
							Type:        schema.TypeString,
						},
						Attr_MasterVolumeName: {
							Computed:    true,
							Description: "Indicates master volume name",
							Type:        schema.TypeString,
						},
						Attr_MirroringState: {
							Computed:    true,
							Description: "Mirroring state for replication enabled volume",
							Type:        schema.TypeString,
						},
						Attr_Name: {
							Computed:    true,
							Description: "The name of the volume.",
							Type:        schema.TypeString,
						},
						Attr_OutOfBandDeleted: {
							Computed:    true,
							Description: "Indicates if the volume does not exist on storage controller.",
							Type:        schema.TypeBool,
						},
						Attr_Pool: {
							Computed:    true,
							Description: "Volume pool, name of storage pool where the volume is located.",
							Type:        schema.TypeString,
						},
						Attr_PrimaryRole: {
							Computed:    true,
							Description: "Indicates whether master/aux volume is playing the primary role.",
							Type:        schema.TypeString,
						},
						Attr_ReplicationEnabled: {
							Computed:    true,
							Description: "Indicates if the volume should be replication enabled or not.",
							Type:        schema.TypeBool,
						},
						Attr_ReplicationSites: {
							Computed:    true,
							Description: "List of replication sites for volume replication.",
							Elem:        &schema.Schema{Type: schema.TypeString},
							Type:        schema.TypeList,
						},
						Attr_ReplicationStatus: {
							Computed:    true,
							Description: "The replication status of the volume.",
							Type:        schema.TypeString,
						},
						Attr_ReplicationType: {
							Computed:    true,
							Description: "The replication type of the volume, 'metro' or 'global'.",
							Type:        schema.TypeString,
						},
						Attr_Shareable: {
							Computed:    true,
							Description: "Indicates if the volume is shareable between VMs.",
							Type:        schema.TypeBool,
						},
						Attr_Size: {
							Computed:    true,
							Description: "The size of this volume in GB.",
							Type:        schema.TypeFloat,
						},
						Attr_State: {
							Computed:    true,
							Description: "The state of the volume.",
							Type:        schema.TypeString,
						},
						Attr_Type: {
							Computed:    true,
							Description: "The disk type that is used for this volume.",
							Type:        schema.TypeString,
						},
						Attr_UserTags: {
							Computed:    true,
							Description: "List of user tags attached to the resource.",
							Elem:        &schema.Schema{Type: schema.TypeString},
							Set:         schema.HashString,
							Type:        schema.TypeSet,
						},
						Attr_VolumePool: {
							Computed:    true,
							Description: "Name of the storage pool where the volume is located.",
							Type:        schema.TypeString,
						},
						Attr_VolumeType: {
							Computed:    true,
							Description: "Name of storage template used to create the volume.",
							Type:        schema.TypeString,
						},
						Attr_WWN: {
							Computed:    true,
							Description: "The world wide name of the volume.",
							Type:        schema.TypeString,
						},
					},
				},
				Type: schema.TypeList,
			},
		},
	}
}

func dataSourceIBMPIInstanceVolumesRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	sess, err := meta.(conns.ClientSession).IBMPISession()
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("IBMPISession failed: %s", err.Error()), "(Data) ibm_pi_instance_volumes", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	cloudInstanceID := d.Get(Arg_CloudInstanceID).(string)
	var instanceID string
	if v, ok := d.GetOk(Arg_InstanceID); ok {
		instanceID = v.(string)
	} else if v, ok := d.GetOk(Arg_InstanceName); ok {
		instanceID = v.(string)
	}

	volumeC := instance.NewIBMPIVolumeClient(ctx, sess, cloudInstanceID)
	volumedata, err := volumeC.GetAllInstanceVolumes(instanceID)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf(" failed: %s", err.Error()), "(Data) ibm_pi_instance_volumes", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	var clientgenU, _ = uuid.GenerateUUID()
	d.SetId(clientgenU)
	if len(volumedata.Volumes) > 0 {
		d.Set(Attr_BootVolumeID, *volumedata.Volumes[0].VolumeID)
	}

	d.Set(Attr_InstanceVolumes, flattenVolumesInstances(volumedata.Volumes, meta))

	return nil
}

func flattenVolumesInstances(list []*models.VolumeReference, meta any) []map[string]any {
	result := make([]map[string]any, 0, len(list))
	for _, i := range list {
		l := map[string]any{
			Attr_Auxiliary:            *i.Auxiliary,
			Attr_AuxiliaryVolumeName:  i.AuxVolumeName,
			Attr_Bootable:             *i.Bootable,
			Attr_ConsistencyGroupName: i.ConsistencyGroupName,
			Attr_CreationDate:         i.CreationDate.String(),
			Attr_GroupID:              i.GroupID,
			Attr_Href:                 *i.Href,
			Attr_ID:                   *i.VolumeID,
			Attr_IOThrottleRate:       i.IoThrottleRate,
			Attr_LastUpdateDate:       i.LastUpdateDate.String(),
			Attr_MasterVolumeName:     i.MasterVolumeName,
			Attr_MirroringState:       i.MirroringState,
			Attr_Name:                 *i.Name,
			Attr_OutOfBandDeleted:     i.OutOfBandDeleted,
			Attr_Pool:                 i.VolumePool,
			Attr_PrimaryRole:          i.PrimaryRole,
			Attr_ReplicationEnabled:   i.ReplicationEnabled,
			Attr_ReplicationStatus:    i.ReplicationStatus,
			Attr_ReplicationType:      i.ReplicationType,
			Attr_Shareable:            *i.Shareable,
			Attr_Size:                 *i.Size,
			Attr_State:                *i.State,
			Attr_Type:                 *i.DiskType,
			Attr_VolumePool:           i.VolumePool,
			Attr_VolumeType:           i.VolumeType,
			Attr_WWN:                  *i.Wwn,
		}
		if i.Crn != "" {
			l[Attr_CRN] = i.Crn
			tags, err := flex.GetGlobalTagsUsingCRN(meta, string(i.Crn), "", UserTagType)
			if err != nil {
				log.Printf("Error on get of volume (%s) user_tags: %s", *i.VolumeID, err)
			}
			l[Attr_UserTags] = tags
		}
		if i.DeleteOnTermination != nil {
			l[Attr_DeleteOnTermination] = *i.DeleteOnTermination
		}
		if i.FreezeTime != nil {
			l[Attr_FreezeTime] = i.FreezeTime.String()
		}
		if len(i.ReplicationSites) > 0 {
			l[Attr_ReplicationSites] = i.ReplicationSites
		}
		result = append(result, l)
	}
	return result
}
