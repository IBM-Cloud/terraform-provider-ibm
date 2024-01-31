// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package power

import (
	"context"

	"github.com/IBM-Cloud/power-go-client/clients/instance"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func DataSourceIBMPIInstanceSnapshot() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMPIInstanceSnapshotRead,
		Schema: map[string]*schema.Schema{
			// Arguments
			Arg_CloudInstanceID: {
				Description:  "The GUID of the service instance associated with an account.",
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.NoZeroValues,
			},
			Arg_SnapshotID: {
				Description:  "The unique identifier of the Power Systems Virtual Machine instance snapshot.",
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.NoZeroValues,
			},

			// Attributes
			Attr_Action: {
				Computed:    true,
				Description: "Action performed on the instance snapshot.",
				Type:        schema.TypeString,
			},
			Attr_CreationDate: {
				Computed:    true,
				Description: "Date of snapshot creation.",
				Type:        schema.TypeString,
			},
			Attr_Description: {
				Computed:    true,
				Description: "The description of the snapshot.",
				Type:        schema.TypeString,
			},
			Attr_LastUpdatedDate: {
				Computed:    true,
				Description: "Date of last update.",
				Type:        schema.TypeString,
			},
			Attr_Name: {
				Computed:    true,
				Description: "The name of the Power Systems Virtual Machine instance snapshot.",
				Type:        schema.TypeString,
			},
			Attr_PercentComplete: {
				Computed:    true,
				Description: "The snapshot completion percentage.",
				Type:        schema.TypeInt,
			},
			Attr_Status: {
				Computed:    true,
				Description: "The status of the Power Virtual Machine instance snapshot.",
				Type:        schema.TypeString,
			},
			Attr_VolumeSnapshots: {
				Computed:    true,
				Description: "A map of volume snapshots included in the Power Virtual Machine instance snapshot.",
				Type:        schema.TypeMap,
			},
		},
	}
}

func dataSourceIBMPIInstanceSnapshotRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := meta.(conns.ClientSession).IBMPISession()
	if err != nil {
		return diag.FromErr(err)
	}

	cloudInstanceID := d.Get(Arg_CloudInstanceID).(string)
	snapshot := instance.NewIBMPISnapshotClient(ctx, sess, cloudInstanceID)
	snapshotData, err := snapshot.Get(d.Get(Arg_SnapshotID).(string))
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(*snapshotData.SnapshotID)
	d.Set(Attr_Action, snapshotData.Action)
	d.Set(Attr_CreationDate, snapshotData.CreationDate.String())
	d.Set(Attr_Description, snapshotData.Description)
	d.Set(Attr_LastUpdatedDate, snapshotData.LastUpdateDate.String())
	d.Set(Attr_Name, snapshotData.Name)
	d.Set(Attr_PercentComplete, snapshotData.PercentComplete)
	d.Set(Attr_Status, snapshotData.Status)
	d.Set(Attr_VolumeSnapshots, snapshotData.VolumeSnapshots)
	return nil
}
