// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc

import (
	"context"
	"fmt"
	"log"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceSnapshotClone() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMISSnapshotCloneRead,

		Schema: map[string]*schema.Schema{
			isSnapshot: {
				Type:     schema.TypeString,
				Required: true,
			},

			isSnapshotCloneAvailable: {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Indicates whether this snapshot clone is available for use.",
			},

			isSnapshotCloneCreatedAt: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The date and time that this snapshot clone was created.",
			},

			isSnapshotCloneZone: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The zone this snapshot clone resides in.",
			},
		},
	}
}

func dataSourceIBMISSnapshotCloneRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	id := d.Get(isSnapshot).(string)
	zone := d.Get(isSnapshotCloneZone).(string)
	err := getSnapshotClone(context, d, meta, id, zone)
	if err != nil {
		return err
	}
	return nil
}

func getSnapshotClone(context context.Context, d *schema.ResourceData, meta interface{}, id, zone string) diag.Diagnostics {
	sess, err := meta.(conns.ClientSession).VpcV1API()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_ibm_is_snapshot_clone", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	getSnapshotCloneOptions := &vpcv1.GetSnapshotCloneOptions{
		ID:       &id,
		ZoneName: &zone,
	}

	clone, _, err := sess.GetSnapshotCloneWithContext(context, getSnapshotCloneOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetSnapshotCloneWithContext failed: %s", err.Error()), "(Data) ibm_ibm_is_snapshot_clone", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	if clone != nil && clone.Zone != nil {
		d.SetId(*clone.Zone.Name)
		if err = d.Set("zone", *clone.Zone.Name); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting zone: %s", err), "(Data) ibm_ibm_is_snapshot_clone", "read", "set-zone").GetDiag()
		}
		if err = d.Set("available", *clone.Available); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting available: %s", err), "(Data) ibm_ibm_is_snapshot_clone", "read", "set-available").GetDiag()
		}
		if clone.CreatedAt != nil {
			if err = d.Set("created_at", flex.DateTimeToString(clone.CreatedAt)); err != nil {
				return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting created_at: %s", err), "(Data) ibm_ibm_is_snapshot_clone", "read", "set-created_at").GetDiag()
			}
		}
	} else {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("No snapshot(%s) clone(%s) found ", id, zone), "(Data) ibm_ibm_is_snapshot_clone", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	return nil
}
