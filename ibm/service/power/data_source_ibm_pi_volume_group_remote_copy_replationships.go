// Copyright IBM Corp. 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package power

import (
	//"fmt"

	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	"github.com/IBM-Cloud/power-go-client/clients/instance"
	"github.com/IBM-Cloud/power-go-client/helpers"
	"github.com/IBM-Cloud/power-go-client/power/models"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
)

func DataSourceIBMPIVolumeGroupRemoteCopyRelationships() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMPIVolumeGroupRemoteCopyRelationshipsReads,
		Schema: map[string]*schema.Schema{
			PIVolumeGroupName: {
				Type:         schema.TypeString,
				Required:     true,
				Description:  "Volume group name",
				ValidateFunc: validation.NoZeroValues,
			},
			helpers.PICloudInstanceId: {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.NoZeroValues,
			},

			// Computed Attributes
			"remote_copy_relationships": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "List of remote copy relationships",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"aux_changed_volume_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Name of the volume that is acting as the auxiliary change volume for the relationship",
						},
						"aux_volume_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Auxiliary volume name at storage host level",
						},
						"consistency_group_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Consistency Group Name if volume is a part of volume group",
						},
						"copy_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Indicates the copy type.",
						},
						"cycling_mode": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Indicates the type of cycling mode used.",
						},
						"freeze_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Freeze time of remote copy relationship",
						},
						"master_changed_volume_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Name of the volume that is acting as the master change volume for the relationship",
						},
						"master_volume_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Master volume name at storage host level",
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Remote copy relationship name",
						},
						"primary_role": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Indicates whether master/aux volume is playing the primary role",
						},
						"progress": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Indicates the relationship progress",
						},
						"remote_copy_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Remote copy relationship ID",
						},
						"state": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Indicates the relationship state",
						},
						"sync": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Indicates whether the relationship is synchronized",
						},
					},
				},
			},
		},
	}
}

func dataSourceIBMPIVolumeGroupRemoteCopyRelationshipsReads(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := meta.(conns.ClientSession).IBMPISession()
	if err != nil {
		return diag.FromErr(err)
	}

	cloudInstanceID := d.Get(helpers.PICloudInstanceId).(string)
	vgClient := instance.NewIBMPIVolumeGroupClient(ctx, sess, cloudInstanceID)
	vgData, err := vgClient.GetVolumeGroupRemoteCopyRelationships(d.Get(PIVolumeGroupName).(string))
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(vgData.ID)
	d.Set("remote_copy_relationships", flattenVolumeGroupRemoteCopyRelationships(vgData.RemoteCopyRelationships))

	return nil
}

func flattenVolumeGroupRemoteCopyRelationships(list []*models.RemoteCopyRelationship) []map[string]interface{} {
	log.Printf("Calling the flattenVolumeGroupRemoteCopyRelationships call with list %d", len(list))
	result := make([]map[string]interface{}, 0, len(list))
	for _, i := range list {
		l := map[string]interface{}{
			"aux_changed_volume_name":    i.AuxChangedVolumeName,
			"aux_volume_name":            i.AuxVolumeName,
			"consistency_group_name":     i.ConsistencyGroupName,
			"copy_type":                  i.CopyType,
			"cycling_mode":               i.CyclingMode,
			"freeze_time":                i.FreezeTime.String(),
			"master_changed_volume_name": i.MasterChangedVolumeName,
			"master_volume_name":         i.MasterVolumeName,
			"name":                       i.Name,
			"primary_role":               i.PrimaryRole,
			"progress":                   i.Progress,
			"remote_copy_id":             i.RemoteCopyID,
			"state":                      i.State,
			"sync":                       i.Sync,
		}

		result = append(result, l)
	}

	return result
}
