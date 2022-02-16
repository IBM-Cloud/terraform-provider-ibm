// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package power

import (
	//"fmt"

	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	"github.com/IBM-Cloud/power-go-client/clients/instance"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
)

const (
	// Arguments
	PIVolumeName                  = "pi_volume_name"
	PIVolumeAffinityInstance      = "pi_affinity_instance"
	PIVolumeAffinityPolicy        = "pi_affinity_policy"
	PIVolumeAffinityVolume        = "pi_affinity_volume"
	PIVolumeAntiAffinityInstances = "pi_anti_affinity_instances"
	PIVolumeAniAffinityVolumes    = "pi_anti_affinity_volumes"
	PIVolumePool                  = "pi_volume_pool"
	PIVolumeShareable             = "pi_volume_shareable"
	PIVolumeSize                  = "pi_volume_size"
	PIVolumeType                  = "pi_volume_type"
	PIVolumeAttachInstanceID      = "pi_instance_id"
	PIVolumeAttachVolumeID        = "pi_volume_id"

	// Attributes
	VolumeDiskType            = "disk_type"
	VolumeBootable            = "bootable"
	VolumeShareable           = "shareable"
	VolumeSize                = "size"
	VolumeState               = "state"
	VolumeVolumePool          = "volume_pool"
	VolumeWWN                 = "wwn"
	VolumeDeleteOnTermination = "delete_on_termination"
	VolumeID                  = "volume_id"
	VolumeStatus              = "volume_status"
	VolumeAttachStatus        = "status"
)

func DataSourceIBMPIVolume() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMPIVolumeRead,
		Schema: map[string]*schema.Schema{
			PIVolumeName: {
				Type:         schema.TypeString,
				Required:     true,
				Description:  "Volume Name to be used for pvminstances",
				ValidateFunc: validation.NoZeroValues,
			},
			PICloudInstanceID: {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.NoZeroValues,
			},

			// Computed Attributes
			VolumeState: {
				Type:     schema.TypeString,
				Computed: true,
			},
			VolumeSize: {
				Type:     schema.TypeInt,
				Computed: true,
			},
			VolumeShareable: {
				Type:     schema.TypeBool,
				Computed: true,
			},
			VolumeBootable: {
				Type:     schema.TypeBool,
				Computed: true,
			},
			VolumeDiskType: {
				Type:     schema.TypeString,
				Computed: true,
			},
			VolumeVolumePool: {
				Type:     schema.TypeString,
				Computed: true,
			},
			VolumeWWN: {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceIBMPIVolumeRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := meta.(conns.ClientSession).IBMPISession()
	if err != nil {
		return diag.FromErr(err)
	}

	cloudInstanceID := d.Get(PICloudInstanceID).(string)
	volumeC := instance.NewIBMPIVolumeClient(ctx, sess, cloudInstanceID)
	volumedata, err := volumeC.Get(d.Get(PIVolumeName).(string))
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(*volumedata.VolumeID)
	d.Set(VolumeSize, volumedata.Size)
	d.Set(VolumeState, volumedata.State)
	d.Set(VolumeShareable, volumedata.Shareable)
	d.Set(VolumeBootable, volumedata.Bootable)
	d.Set(VolumeDiskType, volumedata.DiskType)
	d.Set(VolumeVolumePool, volumedata.VolumePool)
	d.Set(VolumeWWN, volumedata.Wwn)

	return nil
}
