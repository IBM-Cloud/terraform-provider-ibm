// Copyright IBM Corp. 2023, 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package power

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	st "github.com/IBM-Cloud/power-go-client/clients/instance"
	"github.com/IBM-Cloud/power-go-client/helpers"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
)

func DataSourceIBMPIVolumeClone() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMPIVolumeCloneRead,
		Schema: map[string]*schema.Schema{
			PIVolumeCloneTaskID: {
				Type:         schema.TypeString,
				Required:     true,
				Description:  "Clone task ID",
				ValidateFunc: validation.NoZeroValues,
			},
			helpers.PICloudInstanceId: {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.NoZeroValues,
			},
			// Computed attributes
			"cloned_volumes": clonedVolumesSchema(),
			"volume_clone_failure_reason": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The reason the clone volumes task has failed",
			},
			"volume_clone_percent_complete": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Clone task completion percentage",
			},
			"volume_clone_status": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Status of the clone volumes task",
			},
		},
	}
}

func dataSourceIBMPIVolumeCloneRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := meta.(conns.ClientSession).IBMPISession()
	if err != nil {
		return diag.FromErr(err)
	}

	cloudInstanceID := d.Get(helpers.PICloudInstanceId).(string)
	client := st.NewIBMPICloneVolumeClient(ctx, sess, cloudInstanceID)
	volClone, err := client.Get(d.Get(PIVolumeCloneTaskID).(string))
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(d.Get(PIVolumeCloneTaskID).(string))
	if volClone.Status != nil {
		d.Set("volume_clone_status", *volClone.Status)
	}
	d.Set("volume_clone_failure_reason", volClone.FailedReason)
	if volClone.PercentComplete != nil {
		d.Set("volume_clone_percent_complete", *volClone.PercentComplete)
	}
	d.Set("cloned_volumes", flattenClonedVolumes(volClone.ClonedVolumes))

	return nil
}
