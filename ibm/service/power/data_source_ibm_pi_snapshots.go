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

// Attributes and Arguments defined in data_source_ibm_pi_snapshot.go
func DataSourceIBMPISnapshots() *schema.Resource {

	return &schema.Resource{
		ReadContext: dataSourceIBMPISnapshotsRead,
		Schema: map[string]*schema.Schema{
			PICloudInstanceID: {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.NoZeroValues,
			},

			//Computed Attributes
			Snapshots: {
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

func dataSourceIBMPISnapshotsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := meta.(conns.ClientSession).IBMPISession()
	if err != nil {
		return diag.FromErr(err)
	}

	cloudInstanceID := d.Get(PICloudInstanceID).(string)
	snapshot := instance.NewIBMPISnapshotClient(ctx, sess, cloudInstanceID)
	snapshotData, err := snapshot.GetAll()
	if err != nil {
		return diag.FromErr(err)
	}

	var clientgenU, _ = uuid.GenerateUUID()
	d.SetId(clientgenU)
	d.Set(Snapshots, flattenSnapshotsInstances(snapshotData.Snapshots))

	return nil
}

func flattenSnapshotsInstances(list []*models.Snapshot) []map[string]interface{} {
	log.Printf("Calling the flattensnapshotsinstances call with list %d", len(list))
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
