// Copyright IBM Corp. 2024 All Rights Reserved.
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
				Description:  "The ID of the volume clone task.",
				ValidateFunc: validation.NoZeroValues,
			},
			Arg_CloudInstanceID: {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.NoZeroValues,
				Description:  "The GUID of the service instance associated with an account.",
			},
			// Computed attributes
			"cloned_volumes": clonedVolumesSchema(),
			"failure_reason": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The reason the clone volumes task has failed.",
			},
			"percent_complete": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The completion percentage of the volume clone task.",
			},
			"status": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The status of the volume clone task.",
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
		d.Set("status", *volClone.Status)
	}
	d.Set("failure_reason", volClone.FailedReason)
	if volClone.PercentComplete != nil {
		d.Set("percent_complete", *volClone.PercentComplete)
	}
	d.Set("cloned_volumes", flattenClonedVolumes(volClone.ClonedVolumes))

	return nil
}
