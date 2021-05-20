// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"time"

	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	isSnapshots  = "snapshots"
	isSnapshotId = "id"
)

func dataSourceSnapshots() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceIBMISSnapshotsRead,

		Schema: map[string]*schema.Schema{

			isSnapshots: {
				Type:        schema.TypeList,
				Description: "List of snapshots",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						isSnapshotId: {
							Type:     schema.TypeString,
							Computed: true,
						},

						isSnapshotName: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Snapshot name",
						},

						isSnapshotResourceGroup: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Resource group info",
						},

						isSnapshotSourceVolume: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Snapshot source volume",
						},

						isSnapshotLCState: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Snapshot lifecycle state",
						},
						isSnapshotCRN: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The crn of the resource",
						},
						isSnapshotEncryption: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Encryption type of the snapshot",
						},
						isSnapshotHref: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "URL for the snapshot",
						},

						isSnapshotDeletable: {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether snapshot is deletable",
						},

						isSnapshotBootable: {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Indicates if a boot volume attachment can be created with a volume created from this snapshot",
						},

						isSnapshotMinCapacity: {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Minimum capacity of the snapshot",
						},
						isSnapshotResourceType: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The resource type of the snapshot",
						},

						isSnapshotSize: {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The size of the snapshot",
						},
					},
				},
			},
		},
	}
}

// func dataSourceIBMISSnapshotValidator() *ResourceValidator {
// 	validateSchema := make([]ValidateSchema, 1)
// 	validateSchema = append(validateSchema,
// 		ValidateSchema{
// 			Identifier:                 isSnapshotName,
// 			ValidateFunctionIdentifier: ValidateNoZeroValues,
// 			Type:                       TypeString})

// 	ibmISSnapshotDataSourceValidator := ResourceValidator{ResourceName: "ibm_is_snapshot", Schema: validateSchema}
// 	return &ibmISSnapshotDataSourceValidator
// }

func dataSourceIBMISSnapshotsRead(d *schema.ResourceData, meta interface{}) error {
	err := getSnapshots(d, meta)
	if err != nil {
		return err
	}
	return nil
}

func getSnapshots(d *schema.ResourceData, meta interface{}) error {
	sess, err := vpcClient(meta)
	if err != nil {
		return err
	}
	start := ""
	allrecs := []vpcv1.Snapshot{}
	for {
		listSnapshotOptions := &vpcv1.ListSnapshotsOptions{}
		if start != "" {
			listSnapshotOptions.Start = &start
		}
		snapshots, response, err := sess.ListSnapshots(listSnapshotOptions)
		if err != nil {
			return fmt.Errorf("Error Fetching vpcs %s\n%s", err, response)
		}
		start = GetNext(snapshots.Next)
		allrecs = append(allrecs, snapshots.Snapshots...)
		if start == "" {
			break
		}
	}

	snapshotsInfo := make([]map[string]interface{}, 0)
	for _, snapshot := range allrecs {
		l := map[string]interface{}{
			isSnapshotId:            *snapshot.ID,
			isSnapshotName:          *snapshot.Name,
			isSnapshotDeletable:     *snapshot.Deletable,
			isSnapshotHref:          *snapshot.Href,
			isSnapshotCRN:           *snapshot.CRN,
			isSnapshotMinCapacity:   *snapshot.MinimumCapacity,
			isSnapshotSize:          *snapshot.Size,
			isSnapshotEncryption:    *snapshot.Encryption,
			isSnapshotLCState:       *snapshot.LifecycleState,
			isSnapshotResourceType:  *snapshot.ResourceType,
			isSnapshotBootable:      *snapshot.Bootable,
			isSnapshotResourceGroup: *snapshot.ResourceGroup.ID,
			isSnapshotSourceVolume:  *snapshot.SourceVolume.ID,
		}
		snapshotsInfo = append(snapshotsInfo, l)
	}
	d.SetId(dataSourceIBMISSnapshotsID(d))
	d.Set(isSnapshots, snapshotsInfo)
	return nil
}

// dataSourceIBMISSnapshotsID returns a reasonable ID for a snapshot list.
func dataSourceIBMISSnapshotsID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}
