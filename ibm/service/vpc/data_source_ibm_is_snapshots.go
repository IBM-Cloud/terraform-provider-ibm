// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc

import (
	"fmt"
	"log"
	"time"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	isSnapshots  = "snapshots"
	isSnapshotId = "id"
)

func DataSourceSnapshots() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceIBMISSnapshotsRead,

		Schema: map[string]*schema.Schema{

			isSnapshotResourceGroup: {
				Type:        schema.TypeString,
				Description: "Filters the snapshot collection by resources group id",
				Optional:    true,
			},

			isSnapshotName: {
				Type:        schema.TypeString,
				Description: "Filters the snapshot collection by snapshot name",
				Optional:    true,
			},

			isSnapshotSourceImage: {
				Type:        schema.TypeString,
				Description: "Filters the snapshot collection by source image id",
				Optional:    true,
			},

			isSnapshotSourceVolume: {
				Type:        schema.TypeString,
				Description: "Filters the snapshot collection by source volume id",
				Optional:    true,
			},

			"backup_policy_plan_id": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Filters the collection to backup policy jobs with the backup plan with the specified identifier",
			},

			"tag": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Filters the collection to resources with the exact tag value",
			},

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
						isSnapshotSourceImage: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "If present, the image id from which the data on this volume was most directly provisioned.",
						},

						isSnapshotOperatingSystem: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The globally unique name for the operating system included in this image",
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
						isSnapshotEncryptionKey: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "A reference to the root key used to wrap the data encryption key for the source volume.",
						},
						isSnapshotHref: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "URL for the snapshot",
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

						isSnapshotClones: {
							Type:        schema.TypeSet,
							Computed:    true,
							Elem:        &schema.Schema{Type: schema.TypeString},
							Set:         schema.HashString,
							Description: "Zones for creating the snapshot clone",
						},

						isSnapshotCapturedAt: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The date and time that this snapshot was created",
						},

						isSnapshotUserTags: {
							Type:        schema.TypeSet,
							Computed:    true,
							Elem:        &schema.Schema{Type: schema.TypeString},
							Set:         flex.ResourceIBMVPCHash,
							Description: "User Tags for the snapshot",
						},

						isSnapshotAccessTags: {
							Type:        schema.TypeSet,
							Computed:    true,
							Elem:        &schema.Schema{Type: schema.TypeString},
							Set:         flex.ResourceIBMVPCHash,
							Description: "List of access tags",
						},

						isSnapshotBackupPolicyPlan: {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "If present, the backup policy plan which created this snapshot.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"deleted": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "If present, this property indicates the referenced resource has been deleted and provides some supplementary information.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"more_info": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Link to documentation about deleted resources.",
												},
											},
										},
									},
									"href": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The URL for this backup policy plan.",
									},
									"id": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The unique identifier for this backup policy plan.",
									},
									"name": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The unique user-defined name for this backup policy plan.",
									},
									"resource_type": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The type of resource referenced",
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

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
		if rgFilterOk, ok := d.GetOk(isSnapshotResourceGroup); ok {
			rgFilter := rgFilterOk.(string)
			listSnapshotOptions.ResourceGroupID = &rgFilter
		}
		if nameFilterOk, ok := d.GetOk(isSnapshotName); ok {
			nameFilter := nameFilterOk.(string)
			listSnapshotOptions.Name = &nameFilter
		}
		if sourceImageFilterOk, ok := d.GetOk(isSnapshotSourceImage); ok {
			sourceImageFilter := sourceImageFilterOk.(string)
			listSnapshotOptions.SourceImageID = &sourceImageFilter
		}
		if sourceVolumeFilterOk, ok := d.GetOk(isSnapshotSourceVolume); ok {
			sourceVolumeFilter := sourceVolumeFilterOk.(string)
			listSnapshotOptions.SourceVolumeID = &sourceVolumeFilter
		}
		if backupPolicyPlanIdFilterOk, ok := d.GetOk("backup_policy_plan_id"); ok {
			backupPolicyPlanIdFilter := backupPolicyPlanIdFilterOk.(string)
			listSnapshotOptions.BackupPolicyPlanID = &backupPolicyPlanIdFilter
		}
		if tagFilterOk, ok := d.GetOk("tag"); ok {
			tagFilter := tagFilterOk.(string)
			listSnapshotOptions.Tag = &tagFilter
		}

		snapshots, response, err := sess.ListSnapshots(listSnapshotOptions)
		if err != nil {
			return fmt.Errorf("[ERROR] Error fetching snapshots %s\n%s", err, response)
		}
		start = flex.GetNext(snapshots.Next)
		allrecs = append(allrecs, snapshots.Snapshots...)
		if start == "" {
			break
		}
	}

	snapshotsInfo := make([]map[string]interface{}, 0)
	for _, snapshot := range allrecs {
		l := map[string]interface{}{
			isSnapshotId:           *snapshot.ID,
			isSnapshotName:         *snapshot.Name,
			isSnapshotHref:         *snapshot.Href,
			isSnapshotCRN:          *snapshot.CRN,
			isSnapshotMinCapacity:  *snapshot.MinimumCapacity,
			isSnapshotSize:         *snapshot.Size,
			isSnapshotEncryption:   *snapshot.Encryption,
			isSnapshotLCState:      *snapshot.LifecycleState,
			isSnapshotResourceType: *snapshot.ResourceType,
			isSnapshotBootable:     *snapshot.Bootable,
		}
		if snapshot.EncryptionKey != nil && snapshot.EncryptionKey.CRN != nil {
			l[isSnapshotEncryptionKey] = *snapshot.EncryptionKey.CRN
		}
		if snapshot.CapturedAt != nil {
			l[isSnapshotCapturedAt] = (*snapshot.CapturedAt).String()
		}

		if snapshot.UserTags != nil {
			l[isSnapshotUserTags] = snapshot.UserTags
		}
		if snapshot.ResourceGroup != nil && snapshot.ResourceGroup.ID != nil {
			l[isSnapshotResourceGroup] = *snapshot.ResourceGroup.ID
		}
		if snapshot.SourceVolume != nil && snapshot.SourceVolume.ID != nil {
			l[isSnapshotSourceVolume] = *snapshot.SourceVolume.ID
		}
		if snapshot.SourceImage != nil && snapshot.SourceImage.ID != nil {
			l[isSnapshotSourceImage] = *snapshot.SourceImage.ID
		}
		if snapshot.OperatingSystem != nil && snapshot.OperatingSystem.Name != nil {
			l[isSnapshotOperatingSystem] = *snapshot.OperatingSystem.Name
		}
		var clones []string
		clones = make([]string, 0)
		if snapshot.Clones != nil {
			for _, clone := range snapshot.Clones {
				if clone.Zone != nil {
					clones = append(clones, *clone.Zone.Name)
				}
			}
		}
		l[isSnapshotClones] = flex.NewStringSet(schema.HashString, clones)
		backupPolicyPlanList := []map[string]interface{}{}
		if snapshot.BackupPolicyPlan != nil {
			backupPolicyPlan := map[string]interface{}{}
			if snapshot.BackupPolicyPlan.Deleted != nil {
				snapshotBackupPolicyPlanDeletedMap := map[string]interface{}{}
				snapshotBackupPolicyPlanDeletedMap["more_info"] = snapshot.BackupPolicyPlan.Deleted.MoreInfo
				backupPolicyPlan["deleted"] = []map[string]interface{}{snapshotBackupPolicyPlanDeletedMap}
			}
			backupPolicyPlan["href"] = snapshot.BackupPolicyPlan.Href
			backupPolicyPlan["id"] = snapshot.BackupPolicyPlan.ID
			backupPolicyPlan["name"] = snapshot.BackupPolicyPlan.Name
			backupPolicyPlan["resource_type"] = snapshot.BackupPolicyPlan.ResourceType
			backupPolicyPlanList = append(backupPolicyPlanList, backupPolicyPlan)
		}
		l[isSnapshotBackupPolicyPlan] = backupPolicyPlanList
		accesstags, err := flex.GetGlobalTagsUsingCRN(meta, *snapshot.CRN, "", isAccessTagType)
		if err != nil {
			log.Printf(
				"Error on get of resource snapshot (%s) access tags: %s", d.Id(), err)
		}
		l[isSnapshotAccessTags] = accesstags
		snapshotsInfo = append(snapshotsInfo, l)
	}
	d.SetId(dataSourceIBMISSnapshotsID(d))
	d.Set(isSnapshots, snapshotsInfo)
	return nil
}

// dataSourceIBMISSnapshotsID returns a reasonable ID for the snapshot list.
func dataSourceIBMISSnapshotsID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}
