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

func DataSourceIBMPIVolumeGroupDetails() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMPIVolumeGroupDetailsRead,
		Schema: map[string]*schema.Schema{
			PIVolumeGroupName: {
				Type:         schema.TypeString,
				Required:     true,
				Description:  "Name of the volume group",
				ValidateFunc: validation.NoZeroValues,
			},
			helpers.PICloudInstanceId: {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.NoZeroValues,
			},

			// Computed Attributes
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"replication_status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"consistency_group_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"status_description": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"errors": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"key": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"message": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"vol_ids": {
										Type:     schema.TypeList,
										Computed: true,
										Elem:     &schema.Schema{Type: schema.TypeString},
									},
								},
							},
						},
					},
				},
			},
			"volume_ids": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func dataSourceIBMPIVolumeGroupDetailsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := meta.(conns.ClientSession).IBMPISession()
	if err != nil {
		return diag.FromErr(err)
	}

	cloudInstanceID := d.Get(helpers.PICloudInstanceId).(string)
	vgClient := instance.NewIBMPIVolumeGroupClient(ctx, sess, cloudInstanceID)
	vgData, err := vgClient.GetDetails(d.Get(PIVolumeGroupName).(string))
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(*vgData.ID)
	d.Set("status", vgData.Status)
	d.Set("consistency_group_name", vgData.ConsistencyGroupName)
	d.Set("replication_status", vgData.ReplicationStatus)
	d.Set("volume_ids", vgData.VolumeIDs)
	if vgData.StatusDescription != nil {
		d.Set("status_description", flattenVolumeGroupStatusDescription(vgData.StatusDescription.Errors))
	}

	return nil
}
