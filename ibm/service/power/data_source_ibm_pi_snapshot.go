// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package power

import (
	"context"
	"log"

	"github.com/IBM-Cloud/power-go-client/clients/instance"
	"github.com/IBM-Cloud/power-go-client/power/models"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/hashicorp/go-uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

const (
	// Arguments
	PISnapshotInstanceName = "pi_instance_name"
	PISnapshotName         = "pi_snapshot_name"
	PISnapshotDescription  = "pi_description"
	PISnapshotVolumeIDs    = "pi_volume_ids"

	// Attributes
	Snapshots               = "instance_snapshots"
	SnapshotSnapshotID      = "snapshot_id"
	Snapshot                = "pvm_snapshots"
	SnapshotAction          = "action"
	SnapshotCreationDate    = "creation_date"
	SnapshotDescription     = "description"
	SnapshotID              = "id"
	SnapshotLastUpdatedDate = "last_updated_date"
	SnapshotName            = "name"
	SnapshotPercentComplete = "percent_complete"
	SnapshotStatus          = "status"
	SnapshotVolumeSnapshots = "volume_snapshots"
)

func DataSourceIBMPISnapshot() *schema.Resource {

	return &schema.Resource{
		ReadContext: dataSourceIBMPISnapshotRead,
		Schema: map[string]*schema.Schema{

			PICloudInstanceID: {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.NoZeroValues,
			},

			PISnapshotInstanceName: {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.NoZeroValues,
			},
			//Computed Attributes

			Snapshot: {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						SnapshotID: {
							Type:     schema.TypeString,
							Computed: true,
						},
						SnapshotName: {
							Type:     schema.TypeString,
							Computed: true,
						},
						SnapshotPercentComplete: {
							Type:     schema.TypeInt,
							Computed: true,
						},

						SnapshotDescription: {
							Type:     schema.TypeString,
							Computed: true,
						},
						SnapshotAction: {
							Type:     schema.TypeString,
							Computed: true,
						},
						SnapshotStatus: {
							Type:     schema.TypeString,
							Computed: true,
						},
						SnapshotCreationDate: {
							Type:     schema.TypeString,
							Computed: true,
						},
						SnapshotLastUpdatedDate: {
							Type:     schema.TypeString,
							Computed: true,
						},
						SnapshotVolumeSnapshots: {
							Type:     schema.TypeMap,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceIBMPISnapshotRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	sess, err := meta.(conns.ClientSession).IBMPISession()
	if err != nil {
		return diag.FromErr(err)
	}

	cloudInstanceID := d.Get(PICloudInstanceID).(string)
	powerinstancename := d.Get(PISnapshotInstanceName).(string)
	snapshot := instance.NewIBMPIInstanceClient(ctx, sess, cloudInstanceID)
	snapshotData, err := snapshot.GetSnapShotVM(powerinstancename)

	if err != nil {
		return diag.FromErr(err)
	}

	var clientgenU, _ = uuid.GenerateUUID()
	d.SetId(clientgenU)
	d.Set(Snapshot, flattenSnapshotInstances(snapshotData.Snapshots))

	return nil

}

func flattenSnapshotInstances(list []*models.Snapshot) []map[string]interface{} {
	log.Printf("Calling the flattensnapshotinstances call with list %d", len(list))
	result := make([]map[string]interface{}, 0, len(list))
	for _, i := range list {
		l := map[string]interface{}{
			SnapshotID:              *i.SnapshotID,
			SnapshotName:            *i.Name,
			SnapshotDescription:     i.Description,
			SnapshotCreationDate:    i.CreationDate.String(),
			SnapshotLastUpdatedDate: i.LastUpdateDate.String(),
			SnapshotAction:          i.Action,
			SnapshotPercentComplete: i.PercentComplete,
			SnapshotStatus:          i.Status,
			SnapshotVolumeSnapshots: i.VolumeSnapshots,
		}

		result = append(result, l)
	}

	return result
}
