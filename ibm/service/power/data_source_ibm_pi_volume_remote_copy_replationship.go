// Copyright IBM Corp. 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package power

import (
	//"fmt"

	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	"github.com/IBM-Cloud/power-go-client/clients/instance"
	"github.com/IBM-Cloud/power-go-client/helpers"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
)

func DataSourceIBMPIVolumeRemoteCopyRelationship() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMPIVolumeRemoteCopyRelationshipsReads,
		Schema: map[string]*schema.Schema{
			helpers.PIVolumeName: {
				Type:         schema.TypeString,
				Required:     true,
				Description:  "Volume name",
				ValidateFunc: validation.NoZeroValues,
			},
			helpers.PICloudInstanceId: {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.NoZeroValues,
			},

			// Computed Attributes
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
			"cycle_period_seconds": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Indicates the minimum period in seconds between multiple cycles",
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
	}
}

func dataSourceIBMPIVolumeRemoteCopyRelationshipsReads(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := meta.(conns.ClientSession).IBMPISession()
	if err != nil {
		return diag.FromErr(err)
	}

	cloudInstanceID := d.Get(helpers.PICloudInstanceId).(string)
	volClient := instance.NewIBMPIVolumeClient(ctx, sess, cloudInstanceID)
	volData, err := volClient.GetVolumeRemoteCopyRelationships(d.Get(helpers.PIVolumeName).(string))
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(volData.ID)
	d.Set("aux_changed_volume_name", volData.AuxChangedVolumeName)
	d.Set("aux_volume_name", volData.AuxVolumeName)
	d.Set("consistency_group_name", volData.ConsistencyGroupName)
	d.Set("copy_type", volData.CopyType)
	d.Set("cycling_mode", volData.CyclingMode)
	d.Set("cycle_period_seconds", volData.CyclePeriodSeconds)
	d.Set("freeze_time", volData.FreezeTime.String())
	d.Set("master_changed_volume_name", volData.MasterChangedVolumeName)
	d.Set("master_volume_name", volData.MasterVolumeName)
	d.Set("name", volData.Name)
	d.Set("primary_role", volData.PrimaryRole)
	d.Set("progress", volData.Progress)
	d.Set("remote_copy_id", volData.RemoteCopyID)
	d.Set("state", volData.State)
	d.Set("sync", volData.Sync)

	return nil
}
