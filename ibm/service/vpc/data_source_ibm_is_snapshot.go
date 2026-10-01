// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc

import (
	"context"
	"fmt"
	"log"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/validate"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceSnapshot() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMISSnapshotRead,

		Schema: map[string]*schema.Schema{
			"identifier": {
				Type:         schema.TypeString,
				Optional:     true,
				ExactlyOneOf: []string{isSnapshotName, "identifier"},
				Description:  "Snapshot identifier",
				ValidateFunc: validate.InvokeDataSourceValidator("ibm_is_snapshot", "identifier"),
			},

			isSnapshotCopies: {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The copies of this snapshot in other regions.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"crn": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The CRN for the copied snapshot.",
						},
						"deleted": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "If present, this property indicates the referenced resource has been deleted, and provides some supplementary information.",
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
							Description: "The URL for the copied snapshot.",
						},
						"id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The unique identifier for the copied snapshot.",
						},
						"name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name for the copied snapshot. The name is unique across all snapshots in the copied snapshot's native region.",
						},
						"remote": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "If present, this property indicates the referenced resource is remote to this region,and identifies the native region.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"href": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The URL for this region.",
									},
									"name": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The globally unique name for this region.",
									},
								},
							},
						},
						"resource_type": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The resource type.",
						},
					},
				},
			},

			isSnapshotName: {
				Type:         schema.TypeString,
				Optional:     true,
				ExactlyOneOf: []string{isSnapshotName, "identifier"},
				ValidateFunc: validate.InvokeDataSourceValidator("ibm_is_snapshot", isSnapshotName),
				Description:  "Snapshot name",
			},

			"service_tags": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The [service tags](https://cloud.ibm.com/apidocs/tagging#types-of-tags) prefixed with `is.snapshot:` associated with this snapshot.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},

			isSnapshotResourceGroup: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Resource group info",
			},

			isSnapshotSourceVolume: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Snapshot source volume id",
			},
			isSnapshotSourceSnapshot: {
				Type:        schema.TypeList,
				Optional:    true,
				Computed:    true,
				Description: "If present, the source snapshot this snapshot was created from.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"crn": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "The CRN of the source snapshot.",
						},
						"deleted": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "If present, this property indicates the referenced resource has been deleted, and providessome supplementary information.",
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
							Description: "The URL for the source snapshot.",
						},
						"id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The unique identifier for the source snapshot.",
						},
						"name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name for the source snapshot. The name is unique across all snapshots in the source snapshot's native region.",
						},
						"remote": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "If present, this property indicates the referenced resource is remote to this region,and identifies the native region.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"href": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The URL for this region.",
									},
									"name": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The globally unique name for this region.",
									},
								},
							},
						},
						"resource_type": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The resource type.",
						},
					},
				},
			},
			isSnapshotConsistencyGroup: {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The snapshot consistency group which created this snapshot.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"crn": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The CRN of this snapshot consistency group.",
						},
						"deleted": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "If present, this property indicates the referenced resource has been deleted, and providessome supplementary information.",
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
							Description: "The URL for the snapshot consistency group.",
						},
						"id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The unique identifier for the snapshot consistency group.",
						},
						"name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name for the snapshot consistency group. The name is unique across all snapshot consistency groups in the region.",
						},
						"resource_type": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The resource type.",
						},
					},
				},
			},
			isSnapshotSourceImage: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "If present, the image id from which the data on this volume was most directly provisioned.",
			},

			isSnapshotAccessTags: {
				Type:        schema.TypeSet,
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Set:         flex.ResourceIBMVPCHash,
				Description: "List of access tags",
			},

			isSnapshotOperatingSystem: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The globally unique name for the operating system included in this image",
			},

			isSnapshotBootable: {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Indicates if a boot volume attachment can be created with a volume created from this snapshot",
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
			isSnapshotCatalogOffering: {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The catalog offering inherited from the snapshot's source. If a virtual server instance is provisioned with a source_snapshot specifying this snapshot, the virtual server instance will use this snapshot's catalog offering, including its pricing plan.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						isSnapshotCatalogOfferingPlanCrn: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The CRN for this catalog offering version's billing plan",
						},
						"deleted": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "If present, this property indicates the referenced resource has been deleted and provides some supplementary information.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"more_info": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Link to documentation about deleted resources.",
									},
								},
							},
						},
						isSnapshotCatalogOfferingVersionCrn: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The CRN for this version of a catalog offering",
						},
					},
				},
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
			"allowed_use": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The usage constraints to match against the requested instance or bare metal server properties to determine compatibility. Can only be specified for bootable snapshots.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"bare_metal_server": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The expression that must be satisfied by the properties of a bare metal server provisioned using the image data in this snapshot.",
						},
						"instance": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The expression that must be satisfied by the properties of a virtual server instance provisioned using this snapshot.",
						},
						"api_version": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The API version with which to evaluate the expressions.",
						},
					},
				},
			},
			"software_attachments": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The software attachments for this snapshot.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"deleted": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "If present, this property indicates the referenced resource has been deleted, and providessome supplementary information.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"more_info": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "A link to documentation about deleted resources.",
									},
								},
							},
						},
						"href": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The URL for this snapshot software attachment.",
						},
						"id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The unique identifier for this snapshot software attachment.",
						},
						"name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name for this snapshot software attachment. The name is unique across all software attachments for the snapshot.",
						},
						"resource_type": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The resource type.",
						},
					},
				},
			},
		},
	}
}

func DataSourceIBMISSnapshotValidator() *validate.ResourceValidator {
	validateSchema := make([]validate.ValidateSchema, 0)
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 "identifier",
			ValidateFunctionIdentifier: validate.ValidateNoZeroValues,
			Type:                       validate.TypeString})

	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 isSnapshotName,
			ValidateFunctionIdentifier: validate.ValidateNoZeroValues,
			Type:                       validate.TypeString})

	ibmISSnapshotDataSourceValidator := validate.ResourceValidator{ResourceName: "ibm_is_snapshot", Schema: validateSchema}
	return &ibmISSnapshotDataSourceValidator
}

func dataSourceIBMISSnapshotRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	name := d.Get(isSnapshotName).(string)
	id := d.Get("identifier").(string)
	err := snapshotGetByNameOrID(context, d, meta, name, id)
	if err != nil {
		return err
	}
	return nil
}

func snapshotGetByNameOrID(context context.Context, d *schema.ResourceData, meta interface{}, name, id string) diag.Diagnostics {
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_is_snapshot", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	if name != "" {
		start := ""
		allrecs := []vpcv1.Snapshot{}
		for {
			listSnapshotOptions := &vpcv1.ListSnapshotsOptions{
				Name: &name,
			}
			if start != "" {
				listSnapshotOptions.Start = &start
			}
			snapshots, _, err := sess.ListSnapshotsWithContext(context, listSnapshotOptions)
			if err != nil {
				tfErr := flex.TerraformErrorf(err, fmt.Sprintf("ListSnapshotsWithContext failed: %s", err.Error()), "(Data) ibm_is_snapshot", "read")
				log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
				return tfErr.GetDiag()
			}
			start = flex.GetNext(snapshots.Next)
			allrecs = append(allrecs, snapshots.Snapshots...)
			if start == "" {
				break
			}
		}

		for _, snapshot := range allrecs {
			if *snapshot.Name == name || *snapshot.ID == id {
				d.SetId(*snapshot.ID)
				if err = d.Set("name", snapshot.Name); err != nil {
					return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting name: %s", err), "(Data) ibm_is_snapshot", "read", "set-name").GetDiag()
				}
				if err = d.Set("href", snapshot.Href); err != nil {
					return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting href: %s", err), "(Data) ibm_is_snapshot", "read", "set-href").GetDiag()
				}
				if err = d.Set("crn", snapshot.CRN); err != nil {
					return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting crn: %s", err), "(Data) ibm_is_snapshot", "read", "set-crn").GetDiag()
				}
				if err = d.Set("minimum_capacity", flex.IntValue(snapshot.MinimumCapacity)); err != nil {
					return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting minimum_capacity: %s", err), "(Data) ibm_is_snapshot", "read", "set-minimum_capacity").GetDiag()
				}
				if err = d.Set("size", flex.IntValue(snapshot.Size)); err != nil {
					return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting size: %s", err), "(Data) ibm_is_snapshot", "read", "set-size").GetDiag()
				}
				if err = d.Set("encryption", snapshot.Encryption); err != nil {
					return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting encryption: %s", err), "(Data) ibm_is_snapshot", "read", "set-encryption").GetDiag()
				}
				if snapshot.EncryptionKey != nil && snapshot.EncryptionKey.CRN != nil {
					if err = d.Set("encryption_key", *snapshot.EncryptionKey.CRN); err != nil {
						return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting encryption_key: %s", err), "(Data) ibm_is_snapshot", "read", "set-encryption_key").GetDiag()
					}
				}
				if err = d.Set("lifecycle_state", snapshot.LifecycleState); err != nil {
					return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting lifecycle_state: %s", err), "(Data) ibm_is_snapshot", "read", "set-lifecycle_state").GetDiag()
				}
				if err = d.Set("resource_type", snapshot.ResourceType); err != nil {
					return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting resource_type: %s", err), "(Data) ibm_is_snapshot", "read", "set-resource_type").GetDiag()
				}
				if err = d.Set("bootable", snapshot.Bootable); err != nil {
					return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting bootable: %s", err), "(Data) ibm_is_snapshot", "read", "set-bootable").GetDiag()
				}

				// source snapshot
				sourceSnapshotList := []map[string]interface{}{}
				if snapshot.SourceSnapshot != nil {
					sourceSnapshot := map[string]interface{}{}
					sourceSnapshot["href"] = snapshot.SourceSnapshot.Href
					if snapshot.SourceSnapshot.Deleted != nil {
						snapshotSourceSnapshotDeletedMap := map[string]interface{}{}
						snapshotSourceSnapshotDeletedMap["more_info"] = snapshot.SourceSnapshot.Deleted.MoreInfo
						sourceSnapshot["deleted"] = []map[string]interface{}{snapshotSourceSnapshotDeletedMap}
					}
					sourceSnapshot["id"] = snapshot.SourceSnapshot.ID
					sourceSnapshot["name"] = snapshot.SourceSnapshot.Name

					sourceSnapshot["resource_type"] = snapshot.SourceSnapshot.ResourceType
					sourceSnapshotList = append(sourceSnapshotList, sourceSnapshot)
				}
				if err = d.Set("source_snapshot", sourceSnapshotList); err != nil {
					return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting source_snapshot: %s", err), "(Data) ibm_is_snapshot", "read", "set-source_snapshot").GetDiag()
				}

				// snapshot consistency group
				snapshotConsistencyGroupList := []map[string]interface{}{}
				if snapshot.SnapshotConsistencyGroup != nil {
					snapshotConsistencyGroup := map[string]interface{}{}
					snapshotConsistencyGroup["href"] = snapshot.SnapshotConsistencyGroup.Href
					snapshotConsistencyGroup["crn"] = snapshot.SnapshotConsistencyGroup.CRN
					if snapshot.SnapshotConsistencyGroup.Deleted != nil {
						snapshotConsistencyGroupDeletedMap := map[string]interface{}{}
						snapshotConsistencyGroupDeletedMap["more_info"] = snapshot.SnapshotConsistencyGroup.Deleted.MoreInfo
						snapshotConsistencyGroup["deleted"] = []map[string]interface{}{snapshotConsistencyGroupDeletedMap}
					}
					snapshotConsistencyGroup["id"] = snapshot.SnapshotConsistencyGroup.ID
					snapshotConsistencyGroup["name"] = snapshot.SnapshotConsistencyGroup.Name
					snapshotConsistencyGroup["resource_type"] = snapshot.SnapshotConsistencyGroup.ResourceType
					snapshotConsistencyGroupList = append(snapshotConsistencyGroupList, snapshotConsistencyGroup)
				}
				if err = d.Set("snapshot_consistency_group", snapshotConsistencyGroupList); err != nil {
					return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting snapshot_consistency_group: %s", err), "(Data) ibm_is_snapshot", "read", "set-snapshot_consistency_group").GetDiag()
				}

				// snapshot copies
				snapshotCopies := []map[string]interface{}{}
				if snapshot.Copies != nil {
					for _, copiesItem := range snapshot.Copies {
						copiesMap, err := dataSourceIBMIsSnapshotsSnapshotCopiesItemToMap(&copiesItem)
						if err != nil {
							return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_is_snapshot", "read", "copies-to-map").GetDiag()
						}
						snapshotCopies = append(snapshotCopies, copiesMap)
					}
				}
				if err = d.Set("copies", snapshotCopies); err != nil {
					return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting copies: %s", err), "(Data) ibm_is_snapshot", "read", "set-copies").GetDiag()
				}

				if snapshot.UserTags != nil {
					if err = d.Set("tags", snapshot.UserTags); err != nil {
						return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting tags: %s", err), "(Data) ibm_is_snapshot", "read", "set-tags").GetDiag()
					}
				}
				if snapshot.ResourceGroup != nil && snapshot.ResourceGroup.ID != nil {
					if err = d.Set("resource_group", *snapshot.ResourceGroup.ID); err != nil {
						return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting resource_group: %s", err), "(Data) ibm_is_snapshot", "read", "set-resource_group").GetDiag()
					}
				}
				if snapshot.SourceVolume != nil && snapshot.SourceVolume.ID != nil {
					if err = d.Set("source_volume", *snapshot.SourceVolume.ID); err != nil {
						return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting source_volume: %s", err), "(Data) ibm_is_snapshot", "read", "set-source_volume").GetDiag()
					}
				}
				if snapshot.SourceImage != nil && snapshot.SourceImage.ID != nil {
					if err = d.Set("source_image", *snapshot.SourceImage.ID); err != nil {
						return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting source_image: %s", err), "(Data) ibm_is_snapshot", "read", "set-source_image").GetDiag()
					}
				}
				if snapshot.OperatingSystem != nil && snapshot.OperatingSystem.Name != nil {
					if err = d.Set("operating_system", *snapshot.OperatingSystem.Name); err != nil {
						return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting operating_system: %s", err), "(Data) ibm_is_snapshot", "read", "set-operating_system").GetDiag()
					}
				}
				// catalog
				catalogList := make([]map[string]interface{}, 0)
				if snapshot.CatalogOffering != nil {
					versionCrn := ""
					if snapshot.CatalogOffering.Version != nil && snapshot.CatalogOffering.Version.CRN != nil {
						versionCrn = *snapshot.CatalogOffering.Version.CRN
					}
					catalogMap := map[string]interface{}{}
					if versionCrn != "" {
						catalogMap[isSnapshotCatalogOfferingVersionCrn] = versionCrn
					}
					if snapshot.CatalogOffering.Plan != nil {
						planCrn := ""
						if snapshot.CatalogOffering.Plan != nil && snapshot.CatalogOffering.Plan.CRN != nil {
							planCrn = *snapshot.CatalogOffering.Plan.CRN
						}
						if planCrn != "" {
							catalogMap[isSnapshotCatalogOfferingPlanCrn] = planCrn
						}
						if snapshot.CatalogOffering.Plan.Deleted != nil {
							deletedMap := resourceIbmIsSnapshotCatalogOfferingVersionPlanReferenceDeletedToMap(*snapshot.CatalogOffering.Plan.Deleted)
							catalogMap["deleted"] = []map[string]interface{}{deletedMap}
						}
					}
					catalogList = append(catalogList, catalogMap)
				}
				if err = d.Set("catalog_offering", catalogList); err != nil {
					return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting catalog_offering: %s", err), "(Data) ibm_is_snapshot", "read", "set-catalog_offering").GetDiag()
				}

				var clones []string
				clones = make([]string, 0)
				if snapshot.Clones != nil {
					for _, clone := range snapshot.Clones {
						if clone.Zone != nil && clone.Zone.Name != nil {
							clones = append(clones, *clone.Zone.Name)
						}
					}
				}
				if err = d.Set("clones", flex.NewStringSet(schema.HashString, clones)); err != nil {
					return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting clones: %s", err), "(Data) ibm_is_snapshot", "read", "set-clones").GetDiag()
				}
				if err = d.Set("service_tags", snapshot.ServiceTags); err != nil {
					return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting service_tags: %s", err), "(Data) ibm_is_snapshot", "read", "set-service_tags").GetDiag()
				}
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
				if err = d.Set("backup_policy_plan", backupPolicyPlanList); err != nil {
					return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting backup_policy_plan: %s", err), "(Data) ibm_is_snapshot", "read", "set-backup_policy_plan").GetDiag()
				}
				allowedUsed := []map[string]interface{}{}
				if snapshot.AllowedUse != nil {
					modelMap, err := DataSourceIBMIsSnapshotAllowedUseToMap(snapshot.AllowedUse)
					if err != nil {
						tfErr := flex.TerraformErrorf(err, err.Error(), "(Data) ibm_is_snapshot", "read")
						log.Println(tfErr.GetDiag())
					}
					allowedUsed = append(allowedUsed, modelMap)
				}
				if err = d.Set("allowed_use", allowedUsed); err != nil {
					return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting allowed_use: %s", err), "(Data) ibm_is_snapshot", "read", "set-allowed_use").GetDiag()
				}
				accesstags, err := flex.GetGlobalTagsUsingCRN(meta, *snapshot.CRN, "", isAccessTagType)
				if err != nil {
					log.Printf(
						"[ERROR] Error on get of resource snapshot (%s) access tags: %s", d.Id(), err)
				}
				if err = d.Set("access_tags", accesstags); err != nil {
					return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting access_tags: %s", err), "(Data) ibm_is_snapshot", "read", "set-access_tags").GetDiag()
				}
				return nil
			}
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("No snapshot found with name : %s", name), "(Data) ibm_is_snapshot", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	} else {
		getSnapshotOptions := &vpcv1.GetSnapshotOptions{
			ID: &id,
		}
		snapshot, _, err := sess.GetSnapshot(getSnapshotOptions)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetSnapshotWithContext failed: %s", err.Error()), "(Data) ibm_is_snapshot", "read")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
		d.SetId(*snapshot.ID)
		if err = d.Set("name", snapshot.Name); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting name: %s", err), "(Data) ibm_is_snapshot", "read", "set-name").GetDiag()
		}
		if err = d.Set("href", snapshot.Href); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting href: %s", err), "(Data) ibm_is_snapshot", "read", "set-href").GetDiag()
		}
		if err = d.Set("crn", snapshot.CRN); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting crn: %s", err), "(Data) ibm_is_snapshot", "read", "set-crn").GetDiag()
		}
		if err = d.Set("minimum_capacity", flex.IntValue(snapshot.MinimumCapacity)); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting minimum_capacity: %s", err), "(Data) ibm_is_snapshot", "read", "set-minimum_capacity").GetDiag()
		}
		if err = d.Set("size", flex.IntValue(snapshot.Size)); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting size: %s", err), "(Data) ibm_is_snapshot", "read", "set-size").GetDiag()
		}
		if err = d.Set("encryption", snapshot.Encryption); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting encryption: %s", err), "(Data) ibm_is_snapshot", "read", "set-encryption").GetDiag()
		}
		if err = d.Set("lifecycle_state", snapshot.LifecycleState); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting lifecycle_state: %s", err), "(Data) ibm_is_snapshot", "read", "set-lifecycle_state").GetDiag()
		}
		if err = d.Set("resource_type", snapshot.ResourceType); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting resource_type: %s", err), "(Data) ibm_is_snapshot", "read", "set-resource_type").GetDiag()
		}
		if err = d.Set("bootable", snapshot.Bootable); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting bootable: %s", err), "(Data) ibm_is_snapshot", "read", "set-bootable").GetDiag()
		}

		if snapshot.EncryptionKey != nil && snapshot.EncryptionKey.CRN != nil {
			if err = d.Set("encryption_key", snapshot.EncryptionKey.CRN); err != nil {
				return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting encryption_key: %s", err), "(Data) ibm_is_snapshot", "read", "set-encryption_key").GetDiag()
			}
		}
		if err = d.Set("service_tags", snapshot.ServiceTags); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting service_tags: %s", err), "(Data) ibm_is_snapshot", "read", "set-service_tags").GetDiag()
		}
		// source snapshot
		sourceSnapshotList := []map[string]interface{}{}
		if snapshot.SourceSnapshot != nil {
			sourceSnapshot := map[string]interface{}{}
			sourceSnapshot["href"] = snapshot.SourceSnapshot.Href
			sourceSnapshot["crn"] = snapshot.SourceSnapshot.CRN
			if snapshot.SourceSnapshot.Deleted != nil {
				snapshotSourceSnapshotDeletedMap := map[string]interface{}{}
				snapshotSourceSnapshotDeletedMap["more_info"] = snapshot.SourceSnapshot.Deleted.MoreInfo
				sourceSnapshot["deleted"] = []map[string]interface{}{snapshotSourceSnapshotDeletedMap}
			}
			sourceSnapshot["id"] = snapshot.SourceSnapshot.ID
			sourceSnapshot["name"] = snapshot.SourceSnapshot.Name
			sourceSnapshot["resource_type"] = snapshot.SourceSnapshot.ResourceType
			sourceSnapshotList = append(sourceSnapshotList, sourceSnapshot)
		}
		if err = d.Set("source_snapshot", sourceSnapshotList); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting source_snapshot: %s", err), "(Data) ibm_is_snapshot", "read", "set-source_snapshot").GetDiag()
		}
		// snapshot copies
		snapshotCopies := []map[string]interface{}{}
		if snapshot.Copies != nil {
			for _, copiesItem := range snapshot.Copies {
				copiesMap, err := dataSourceIBMIsSnapshotsSnapshotCopiesItemToMap(&copiesItem)
				if err != nil {
					return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_is_snapshot", "read", "copies-to-map").GetDiag()
				}
				snapshotCopies = append(snapshotCopies, copiesMap)
			}
		}
		if err = d.Set("copies", snapshotCopies); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting copies: %s", err), "(Data) ibm_is_snapshot", "read", "set-copies").GetDiag()
		}

		// software attachments
		softwareAttachments := []map[string]interface{}{}
		for _, softwareAttachmentsItem := range snapshot.SoftwareAttachments {
			softwareAttachmentsItemMap, err := DataSourceIBMIsSnapshotSnapshotSoftwareAttachmentReferenceToMap(&softwareAttachmentsItem) // #nosec G601
			if err != nil {
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_is_snapshot", "read", "software_attachments-to-map").GetDiag()
			}
			softwareAttachments = append(softwareAttachments, softwareAttachmentsItemMap)
		}
		if err = d.Set("software_attachments", softwareAttachments); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting software_attachments: %s", err), "(Data) ibm_is_snapshot", "read", "set-software_attachments").GetDiag()
		}

		if !core.IsNil(snapshot.CapturedAt) {
			if err = d.Set("captured_at", flex.DateTimeToString(snapshot.CapturedAt)); err != nil {
				return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting captured_at: %s", err), "(Data) ibm_is_snapshot", "read", "set-captured_at").GetDiag()
			}
		}
		if snapshot.UserTags != nil {
			if err = d.Set("tags", snapshot.UserTags); err != nil {
				return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting tags: %s", err), "(Data) ibm_is_snapshot", "read", "set-tags").GetDiag()
			}
		}
		if snapshot.ResourceGroup != nil && snapshot.ResourceGroup.ID != nil {
			if err = d.Set("resource_group", *snapshot.ResourceGroup.ID); err != nil {
				return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting resource_group: %s", err), "(Data) ibm_is_snapshot", "read", "set-resource_group").GetDiag()
			}
		}
		if snapshot.SourceVolume != nil && snapshot.SourceVolume.ID != nil {
			if err = d.Set("source_volume", *snapshot.SourceVolume.ID); err != nil {
				return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting source_volume: %s", err), "(Data) ibm_is_snapshot", "read", "set-source_volume").GetDiag()
			}
		}
		if snapshot.SourceImage != nil && snapshot.SourceImage.ID != nil {
			if err = d.Set("source_image", *snapshot.SourceImage.ID); err != nil {
				return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting source_image: %s", err), "(Data) ibm_is_snapshot", "read", "set-source_image").GetDiag()
			}
		}
		if snapshot.OperatingSystem != nil && snapshot.OperatingSystem.Name != nil {
			if err = d.Set("operating_system", *snapshot.OperatingSystem.Name); err != nil {
				return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting operating_system: %s", err), "(Data) ibm_is_snapshot", "read", "set-operating_system").GetDiag()
			}
		}
		var clones []string
		clones = make([]string, 0)
		if snapshot.Clones != nil {
			for _, clone := range snapshot.Clones {
				if clone.Zone != nil && clone.Zone.Name != nil {
					clones = append(clones, *clone.Zone.Name)
				}
			}
		}
		if err = d.Set("clones", flex.NewStringSet(schema.HashString, clones)); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting clones: %s", err), "(Data) ibm_is_snapshot", "read", "set-clones").GetDiag()
		}

		if snapshot.CatalogOffering != nil {
			versionCrn := ""
			if snapshot.CatalogOffering.Version != nil && snapshot.CatalogOffering.Version.CRN != nil {
				versionCrn = *snapshot.CatalogOffering.Version.CRN
			}
			catalogList := make([]map[string]interface{}, 0)
			catalogMap := map[string]interface{}{}
			if versionCrn != "" {
				catalogMap[isSnapshotCatalogOfferingVersionCrn] = versionCrn
			}
			if snapshot.CatalogOffering.Plan != nil {
				planCrn := ""
				if snapshot.CatalogOffering.Plan != nil && snapshot.CatalogOffering.Plan.CRN != nil {
					planCrn = *snapshot.CatalogOffering.Plan.CRN
				}
				if planCrn != "" {
					catalogMap[isSnapshotCatalogOfferingPlanCrn] = planCrn
				}
				if snapshot.CatalogOffering.Plan.Deleted != nil {
					deletedMap := resourceIbmIsSnapshotCatalogOfferingVersionPlanReferenceDeletedToMap(*snapshot.CatalogOffering.Plan.Deleted)
					catalogMap["deleted"] = []map[string]interface{}{deletedMap}
				}
			}
			catalogList = append(catalogList, catalogMap)
			if err = d.Set("catalog_offering", catalogList); err != nil {
				return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting catalog_offering: %s", err), "(Data) ibm_is_snapshot", "read", "set-catalog_offering").GetDiag()
			}
		}

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
		if err = d.Set("backup_policy_plan", backupPolicyPlanList); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting backup_policy_plan: %s", err), "(Data) ibm_is_snapshot", "read", "set-backup_policy_plan").GetDiag()
		}
		allowedUses := []map[string]interface{}{}
		if snapshot.AllowedUse != nil {
			modelMap, err := DataSourceIBMIsSnapshotAllowedUseToMap(snapshot.AllowedUse)
			if err != nil {
				tfErr := flex.TerraformErrorf(err, err.Error(), "(Data) ibm_is_snapshot", "read")
				log.Println(tfErr.GetDiag())
			}
			allowedUses = append(allowedUses, modelMap)
		}
		if err = d.Set("allowed_use", allowedUses); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting allowed_use: %s", err), "(Data) ibm_is_snapshot", "read", "set-allowed_use").GetDiag()
		}
		accesstags, err := flex.GetGlobalTagsUsingCRN(meta, *snapshot.CRN, "", isAccessTagType)
		if err != nil {
			log.Printf(
				"[ERROR] Error on get of resource snapshot (%s) access tags: %s", d.Id(), err)
		}
		if err = d.Set("access_tags", accesstags); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting access_tags: %s", err), "(Data) ibm_is_snapshot", "read", "set-access_tags").GetDiag()
		}
		return nil
	}
}

func resourceIbmIsSnapshotCatalogOfferingVersionPlanReferenceDeletedToMap(catalogOfferingVersionPlanReferenceDeleted vpcv1.Deleted) map[string]interface{} {
	catalogOfferingVersionPlanReferenceDeletedMap := map[string]interface{}{}

	catalogOfferingVersionPlanReferenceDeletedMap["more_info"] = catalogOfferingVersionPlanReferenceDeleted.MoreInfo

	return catalogOfferingVersionPlanReferenceDeletedMap
}

func DataSourceIBMIsSnapshotAllowedUseToMap(model *vpcv1.SnapshotAllowedUse) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.BareMetalServer != nil {
		modelMap["bare_metal_server"] = *model.BareMetalServer
	}
	if model.Instance != nil {
		modelMap["instance"] = *model.Instance
	}
	if model.ApiVersion != nil {
		modelMap["api_version"] = *model.ApiVersion
	}
	return modelMap, nil
}

func DataSourceIBMIsSnapshotSnapshotSoftwareAttachmentReferenceToMap(model *vpcv1.SnapshotSoftwareAttachmentReference) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Deleted != nil {
		deletedMap, err := DataSourceIBMIsSnapshotDeletedToMap(model.Deleted)
		if err != nil {
			return modelMap, err
		}
		modelMap["deleted"] = []map[string]interface{}{deletedMap}
	}
	modelMap["href"] = *model.Href
	modelMap["id"] = *model.ID
	modelMap["name"] = *model.Name
	modelMap["resource_type"] = *model.ResourceType
	return modelMap, nil
}

func DataSourceIBMIsSnapshotDeletedToMap(model *vpcv1.Deleted) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["more_info"] = *model.MoreInfo
	return modelMap, nil
}
