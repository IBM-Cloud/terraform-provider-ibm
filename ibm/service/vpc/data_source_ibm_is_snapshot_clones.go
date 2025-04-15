// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	isSnapshot               = "snapshot"
	isSnapshotClones         = "clones"
	isSnapshotCloneAvailable = "available"
	isSnapshotCloneCreatedAt = "created_at"
	isSnapshotCloneZone      = "zone"
)

func DataSourceSnapshotClones() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMISSnapshotClonesRead,

		Schema: map[string]*schema.Schema{
			isSnapshot: {
				Type:     schema.TypeString,
				Required: true,
			},

			isSnapshotClones: {
				Type:        schema.TypeList,
				Description: "List of snapshot clones",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
							Computed:    true,
							Description: "The zone this snapshot clone resides in.",
						},
					},
				},
			},
		},
	}
}

func dataSourceIBMISSnapshotClonesRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	id := d.Get(isSnapshot).(string)
	err := getSnapshotClones(context, d, meta, id)
	if err != nil {
		return err
	}
	return nil
}

func getSnapshotClones(context context.Context, d *schema.ResourceData, meta interface{}, id string) diag.Diagnostics {
	sess, err := meta.(conns.ClientSession).VpcV1API()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_ibm_is_snapshot_clones", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	listSnapshotClonesOptions := &vpcv1.ListSnapshotClonesOptions{
		ID: &id,
	}

	clonesCollection, _, err := sess.ListSnapshotClonesWithContext(context, listSnapshotClonesOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("ListSnapshotClonesWithContext failed: %s", err.Error()), "(Data) ibm_ibm_is_snapshot_clones", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	clones := clonesCollection.Clones

	clonesInfo := make([]map[string]interface{}, 0)
	for _, clone := range clones {
		l := map[string]interface{}{
			isSnapshotCloneAvailable: *clone.Available,
		}
		if clone.CreatedAt != nil {
			l[isSnapshotCloneCreatedAt] = flex.DateTimeToString(clone.CreatedAt)
		}
		if clone.Zone != nil {
			l[isSnapshotCloneZone] = *clone.Zone.Name
		}

		clonesInfo = append(clonesInfo, l)
	}
	d.SetId(dataSourceIBMISSnapshotClonesID(d))
	if err = d.Set("clones", clonesInfo); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting clones: %s", err), "(Data) ibm_ibm_is_snapshot_clones", "read", "set-clones").GetDiag()
	}
	return nil
}

// dataSourceIBMISSnapshotClonesID returns a reasonable ID for the clone list.
func dataSourceIBMISSnapshotClonesID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}
