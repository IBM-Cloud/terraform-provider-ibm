// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	// "github.ibm.com/ibmcloud/vpc-go-sdk/vpcv1"
	"github.com/IBM/vpc-go-sdk/vpcv1"
)

func dataSourceIBMIsVolumes() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMIsVolumesRead,

		Schema: map[string]*schema.Schema{
			"first": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "A link to the first page of resources.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"href": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The URL for a page of resources.",
						},
					},
				},
			},
			"limit": &schema.Schema{
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The maximum number of resources that can be returned by the request.",
			},
			"next": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "A link to the next page of resources. This property is present for all pagesexcept the last page.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"href": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The URL for a page of resources.",
						},
					},
				},
			},
			"volumes": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Collection of volumes.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"active": &schema.Schema{
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Indicates whether a running virtual server instance has an attachment to this volume.",
						},
						"bandwidth": &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The maximum bandwidth (in megabits per second) for the volume.",
						},
						"busy": &schema.Schema{
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Indicates whether this volume is performing an operation that must be serialized. If an operation specifies that it requires serialization, the operation will fail unless this property is `false`.",
						},
						"capacity": &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The capacity to use for the volume (in gigabytes). The specified minimum and maximum capacity values for creating or updating volumes may expand in the future.",
						},
						"created_at": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The date and time that the volume was created.",
						},
						"crn": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The CRN for this volume.",
						},
						"encryption": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The type of encryption used on the volume.",
						},
						"encryption_key": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The root key used to wrap the data encryption key for the volume.This property will be present for volumes with an `encryption` type of`user_managed`.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"crn": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The CRN of the [Key Protect Root Key](https://cloud.ibm.com/docs/key-protect?topic=key-protect-getting-started-tutorial) or [Hyper Protect Crypto Service Root Key](https://cloud.ibm.com/docs/hs-crypto?topic=hs-crypto-get-started) for this resource.",
									},
								},
							},
						},
						"href": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The URL for this volume.",
						},
						"id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The unique identifier for this volume.",
						},
						"iops": &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The maximum I/O operations per second (IOPS) to use for the volume. Applicable only to volumes using a profile `family` of `custom`.",
						},
						"name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The unique user-defined name for this volume.",
						},
						"operating_system": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The operating system associated with this volume. If absent, this volume was notcreated from an image, or the image did not include an operating system.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"href": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The URL for this operating system.",
									},
									"name": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The globally unique name for this operating system.",
									},
								},
							},
						},
						"profile": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The profile this volume uses.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"href": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The URL for this volume profile.",
									},
									"name": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The globally unique name for this volume profile.",
									},
								},
							},
						},
						"resource_group": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The resource group for this volume.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"href": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The URL for this resource group.",
									},
									"id": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The unique identifier for this resource group.",
									},
									"name": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The user-defined name for this resource group.",
									},
								},
							},
						},
						"source_image": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The image from which this volume was created (this may be[deleted](https://cloud.ibm.com/apidocs/vpc#deleted-resources)).If absent, this volume was not created from an image.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"crn": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The CRN for this image.",
									},
									"deleted": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "If present, this property indicates the referenced resource has been deleted and providessome supplementary information.",
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
										Description: "The URL for this image.",
									},
									"id": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The unique identifier for this image.",
									},
									"name": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The user-defined or system-provided name for this image.",
									},
								},
							},
						},
						"source_snapshot": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The snapshot from which this volume was cloned.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"crn": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The CRN for this snapshot.",
									},
									"deleted": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "If present, this property indicates the referenced resource has been deleted and providessome supplementary information.",
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
										Description: "The URL for this snapshot.",
									},
									"id": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The unique identifier for this snapshot.",
									},
									"name": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The user-defined name for this snapshot.",
									},
									"resource_type": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The resource type.",
									},
								},
							},
						},
						"status": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The status of the volume.The enumerated values for this property will expand in the future. When processing this property, check for and log unknown values. Optionally halt processing and surface the error, or bypass the volume on which the unexpected property value was encountered.",
						},
						"status_reasons": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The reasons for the current status (if any).The enumerated reason code values for this property will expand in the future. When processing this property, check for and log unknown values. Optionally halt processing and surface the error, or bypass the resource on which the unexpected reason code was encountered.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"code": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "A snake case string succinctly identifying the status reason.",
									},
									"message": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "An explanation of the status reason.",
									},
									"more_info": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Link to documentation about this status reason.",
									},
								},
							},
						},
						"volume_attachments": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The volume attachments for this volume.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"delete_volume_on_instance_delete": &schema.Schema{
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "If set to true, when deleting the instance the volume will also be deleted.",
									},
									"deleted": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "If present, this property indicates the referenced resource has been deleted and providessome supplementary information.",
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
									"device": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Information about how the volume is exposed to the instance operating system.This property may be absent if the volume attachment's `status` is not `attached`.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"id": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "A unique identifier for the device which is exposed to the instance operating system.",
												},
											},
										},
									},
									"href": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The URL for this volume attachment.",
									},
									"id": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The unique identifier for this volume attachment.",
									},
									"instance": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "The attached instance.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"crn": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The CRN for this virtual server instance.",
												},
												"deleted": &schema.Schema{
													Type:        schema.TypeList,
													Computed:    true,
													Description: "If present, this property indicates the referenced resource has been deleted and providessome supplementary information.",
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
													Description: "The URL for this virtual server instance.",
												},
												"id": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The unique identifier for this virtual server instance.",
												},
												"name": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The user-defined name for this virtual server instance (and default system hostname).",
												},
											},
										},
									},
									"name": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The user-defined name for this volume attachment.",
									},
									"type": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The type of volume attachment.",
									},
								},
							},
						},
						"zone": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The zone this volume resides in.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"href": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The URL for this zone.",
									},
									"name": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The globally unique name for this zone.",
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

func dataSourceIBMIsVolumesRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	vpcClient, err := meta.(ClientSession).VpcV1API()
	if err != nil {
		return diag.FromErr(err)
	}

	listVolumesOptions := &vpcv1.ListVolumesOptions{}

	volumeCollection, response, err := vpcClient.ListVolumesWithContext(context, listVolumesOptions)
	if err != nil {
		log.Printf("[DEBUG] ListVolumesWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("ListVolumesWithContext failed %s\n%s", err, response))
	}

	d.SetId(dataSourceIBMIsVolumesID(d))

	if volumeCollection.First != nil {
		err = d.Set("first", dataSourceVolumeCollectionFlattenFirst(*volumeCollection.First))
		if err != nil {
			return diag.FromErr(fmt.Errorf("Error setting first %s", err))
		}
	}
	if err = d.Set("limit", intValue(volumeCollection.Limit)); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting limit: %s", err))
	}

	if volumeCollection.Next != nil {
		err = d.Set("next", dataSourceVolumeCollectionFlattenNext(*volumeCollection.Next))
		if err != nil {
			return diag.FromErr(fmt.Errorf("Error setting next %s", err))
		}
	}

	if volumeCollection.Volumes != nil {
		err = d.Set("volumes", dataSourceVolumeCollectionFlattenVolumes(volumeCollection.Volumes))
		if err != nil {
			return diag.FromErr(fmt.Errorf("Error setting volumes %s", err))
		}
	}

	return nil
}

// dataSourceIBMIsVolumesID returns a reasonable ID for the list.
func dataSourceIBMIsVolumesID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}

func dataSourceVolumeCollectionFlattenFirst(result vpcv1.VolumeCollectionFirst) (finalList []map[string]interface{}) {
	finalList = []map[string]interface{}{}
	finalMap := dataSourceVolumeCollectionFirstToMap(result)
	finalList = append(finalList, finalMap)

	return finalList
}

func dataSourceVolumeCollectionFirstToMap(firstItem vpcv1.VolumeCollectionFirst) (firstMap map[string]interface{}) {
	firstMap = map[string]interface{}{}

	if firstItem.Href != nil {
		firstMap["href"] = firstItem.Href
	}

	return firstMap
}

func dataSourceVolumeCollectionFlattenNext(result vpcv1.VolumeCollectionNext) (finalList []map[string]interface{}) {
	finalList = []map[string]interface{}{}
	finalMap := dataSourceVolumeCollectionNextToMap(result)
	finalList = append(finalList, finalMap)

	return finalList
}

func dataSourceVolumeCollectionNextToMap(nextItem vpcv1.VolumeCollectionNext) (nextMap map[string]interface{}) {
	nextMap = map[string]interface{}{}

	if nextItem.Href != nil {
		nextMap["href"] = nextItem.Href
	}

	return nextMap
}

func dataSourceVolumeCollectionFlattenVolumes(result []vpcv1.Volume) (volumes []map[string]interface{}) {
	for _, volumesItem := range result {
		volumes = append(volumes, dataSourceVolumeCollectionVolumesToMap(volumesItem))
	}

	return volumes
}

func dataSourceVolumeCollectionVolumesToMap(volumesItem vpcv1.Volume) (volumesMap map[string]interface{}) {
	volumesMap = map[string]interface{}{}

	if volumesItem.Active != nil {
		volumesMap["active"] = volumesItem.Active
	}
	if volumesItem.Bandwidth != nil {
		volumesMap["bandwidth"] = volumesItem.Bandwidth
	}
	if volumesItem.Busy != nil {
		volumesMap["busy"] = volumesItem.Busy
	}
	if volumesItem.Capacity != nil {
		volumesMap["capacity"] = volumesItem.Capacity
	}
	if volumesItem.CreatedAt != nil {
		volumesMap["created_at"] = volumesItem.CreatedAt.String()
	}
	if volumesItem.CRN != nil {
		volumesMap["crn"] = volumesItem.CRN
	}
	if volumesItem.Encryption != nil {
		volumesMap["encryption"] = volumesItem.Encryption
	}
	if volumesItem.EncryptionKey != nil {
		encryptionKeyList := []map[string]interface{}{}
		encryptionKeyMap := dataSourceVolumeCollectionVolumesEncryptionKeyToMap(*volumesItem.EncryptionKey)
		encryptionKeyList = append(encryptionKeyList, encryptionKeyMap)
		volumesMap["encryption_key"] = encryptionKeyList
	}
	if volumesItem.Href != nil {
		volumesMap["href"] = volumesItem.Href
	}
	if volumesItem.ID != nil {
		volumesMap["id"] = volumesItem.ID
	}
	if volumesItem.Iops != nil {
		volumesMap["iops"] = volumesItem.Iops
	}
	if volumesItem.Name != nil {
		volumesMap["name"] = volumesItem.Name
	}
	if volumesItem.OperatingSystem != nil {
		operatingSystemList := []map[string]interface{}{}
		operatingSystemMap := dataSourceVolumeCollectionVolumesOperatingSystemToMap(*volumesItem.OperatingSystem)
		operatingSystemList = append(operatingSystemList, operatingSystemMap)
		volumesMap["operating_system"] = operatingSystemList
	}
	if volumesItem.Profile != nil {
		profileList := []map[string]interface{}{}
		profileMap := dataSourceVolumeCollectionVolumesProfileToMap(*volumesItem.Profile)
		profileList = append(profileList, profileMap)
		volumesMap["profile"] = profileList
	}
	if volumesItem.ResourceGroup != nil {
		resourceGroupList := []map[string]interface{}{}
		resourceGroupMap := dataSourceVolumeCollectionVolumesResourceGroupToMap(*volumesItem.ResourceGroup)
		resourceGroupList = append(resourceGroupList, resourceGroupMap)
		volumesMap["resource_group"] = resourceGroupList
	}
	if volumesItem.SourceImage != nil {
		sourceImageList := []map[string]interface{}{}
		sourceImageMap := dataSourceVolumeCollectionVolumesSourceImageToMap(*volumesItem.SourceImage)
		sourceImageList = append(sourceImageList, sourceImageMap)
		volumesMap["source_image"] = sourceImageList
	}
	if volumesItem.SourceSnapshot != nil {
		sourceSnapshotList := []map[string]interface{}{}
		sourceSnapshotMap := dataSourceVolumeCollectionVolumesSourceSnapshotToMap(*volumesItem.SourceSnapshot)
		sourceSnapshotList = append(sourceSnapshotList, sourceSnapshotMap)
		volumesMap["source_snapshot"] = sourceSnapshotList
	}
	if volumesItem.Status != nil {
		volumesMap["status"] = volumesItem.Status
	}
	if volumesItem.StatusReasons != nil {
		statusReasonsList := []map[string]interface{}{}
		for _, statusReasonsItem := range volumesItem.StatusReasons {
			statusReasonsList = append(statusReasonsList, dataSourceVolumeCollectionVolumesStatusReasonsToMap(statusReasonsItem))
		}
		volumesMap["status_reasons"] = statusReasonsList
	}
	if volumesItem.VolumeAttachments != nil {
		volumeAttachmentsList := []map[string]interface{}{}
		for _, volumeAttachmentsItem := range volumesItem.VolumeAttachments {
			volumeAttachmentsList = append(volumeAttachmentsList, dataSourceVolumeCollectionVolumesVolumeAttachmentsToMap(volumeAttachmentsItem))
		}
		volumesMap["volume_attachments"] = volumeAttachmentsList
	}
	if volumesItem.Zone != nil {
		zoneList := []map[string]interface{}{}
		zoneMap := dataSourceVolumeCollectionVolumesZoneToMap(*volumesItem.Zone)
		zoneList = append(zoneList, zoneMap)
		volumesMap["zone"] = zoneList
	}

	return volumesMap
}

func dataSourceVolumeCollectionVolumesEncryptionKeyToMap(encryptionKeyItem vpcv1.EncryptionKeyReference) (encryptionKeyMap map[string]interface{}) {
	encryptionKeyMap = map[string]interface{}{}

	if encryptionKeyItem.CRN != nil {
		encryptionKeyMap["crn"] = encryptionKeyItem.CRN
	}

	return encryptionKeyMap
}

func dataSourceVolumeCollectionVolumesOperatingSystemToMap(operatingSystemItem vpcv1.OperatingSystemReference) (operatingSystemMap map[string]interface{}) {
	operatingSystemMap = map[string]interface{}{}

	if operatingSystemItem.Href != nil {
		operatingSystemMap["href"] = operatingSystemItem.Href
	}
	if operatingSystemItem.Name != nil {
		operatingSystemMap["name"] = operatingSystemItem.Name
	}

	return operatingSystemMap
}

func dataSourceVolumeCollectionVolumesProfileToMap(profileItem vpcv1.VolumeProfileReference) (profileMap map[string]interface{}) {
	profileMap = map[string]interface{}{}

	if profileItem.Href != nil {
		profileMap["href"] = profileItem.Href
	}
	if profileItem.Name != nil {
		profileMap["name"] = profileItem.Name
	}

	return profileMap
}

func dataSourceVolumeCollectionVolumesResourceGroupToMap(resourceGroupItem vpcv1.ResourceGroupReference) (resourceGroupMap map[string]interface{}) {
	resourceGroupMap = map[string]interface{}{}

	if resourceGroupItem.Href != nil {
		resourceGroupMap["href"] = resourceGroupItem.Href
	}
	if resourceGroupItem.ID != nil {
		resourceGroupMap["id"] = resourceGroupItem.ID
	}
	if resourceGroupItem.Name != nil {
		resourceGroupMap["name"] = resourceGroupItem.Name
	}

	return resourceGroupMap
}

func dataSourceVolumeCollectionVolumesSourceImageToMap(sourceImageItem vpcv1.ImageReference) (sourceImageMap map[string]interface{}) {
	sourceImageMap = map[string]interface{}{}

	if sourceImageItem.CRN != nil {
		sourceImageMap["crn"] = sourceImageItem.CRN
	}
	if sourceImageItem.Deleted != nil {
		deletedList := []map[string]interface{}{}
		deletedMap := dataSourceVolumeCollectionSourceImageDeletedToMap(*sourceImageItem.Deleted)
		deletedList = append(deletedList, deletedMap)
		sourceImageMap["deleted"] = deletedList
	}
	if sourceImageItem.Href != nil {
		sourceImageMap["href"] = sourceImageItem.Href
	}
	if sourceImageItem.ID != nil {
		sourceImageMap["id"] = sourceImageItem.ID
	}
	if sourceImageItem.Name != nil {
		sourceImageMap["name"] = sourceImageItem.Name
	}

	return sourceImageMap
}

func dataSourceVolumeCollectionSourceImageDeletedToMap(deletedItem vpcv1.ImageReferenceDeleted) (deletedMap map[string]interface{}) {
	deletedMap = map[string]interface{}{}

	if deletedItem.MoreInfo != nil {
		deletedMap["more_info"] = deletedItem.MoreInfo
	}

	return deletedMap
}

func dataSourceVolumeCollectionVolumesSourceSnapshotToMap(sourceSnapshotItem vpcv1.SnapshotReference) (sourceSnapshotMap map[string]interface{}) {
	sourceSnapshotMap = map[string]interface{}{}

	if sourceSnapshotItem.CRN != nil {
		sourceSnapshotMap["crn"] = sourceSnapshotItem.CRN
	}
	if sourceSnapshotItem.Deleted != nil {
		deletedList := []map[string]interface{}{}
		deletedMap := dataSourceVolumeCollectionSourceSnapshotDeletedToMap(*sourceSnapshotItem.Deleted)
		deletedList = append(deletedList, deletedMap)
		sourceSnapshotMap["deleted"] = deletedList
	}
	if sourceSnapshotItem.Href != nil {
		sourceSnapshotMap["href"] = sourceSnapshotItem.Href
	}
	if sourceSnapshotItem.ID != nil {
		sourceSnapshotMap["id"] = sourceSnapshotItem.ID
	}
	if sourceSnapshotItem.Name != nil {
		sourceSnapshotMap["name"] = sourceSnapshotItem.Name
	}
	if sourceSnapshotItem.ResourceType != nil {
		sourceSnapshotMap["resource_type"] = sourceSnapshotItem.ResourceType
	}

	return sourceSnapshotMap
}

func dataSourceVolumeCollectionSourceSnapshotDeletedToMap(deletedItem vpcv1.SnapshotReferenceDeleted) (deletedMap map[string]interface{}) {
	deletedMap = map[string]interface{}{}

	if deletedItem.MoreInfo != nil {
		deletedMap["more_info"] = deletedItem.MoreInfo
	}

	return deletedMap
}

func dataSourceVolumeCollectionVolumesStatusReasonsToMap(statusReasonsItem vpcv1.VolumeStatusReason) (statusReasonsMap map[string]interface{}) {
	statusReasonsMap = map[string]interface{}{}

	if statusReasonsItem.Code != nil {
		statusReasonsMap["code"] = statusReasonsItem.Code
	}
	if statusReasonsItem.Message != nil {
		statusReasonsMap["message"] = statusReasonsItem.Message
	}
	if statusReasonsItem.MoreInfo != nil {
		statusReasonsMap["more_info"] = statusReasonsItem.MoreInfo
	}

	return statusReasonsMap
}

func dataSourceVolumeCollectionVolumesVolumeAttachmentsToMap(volumeAttachmentsItem vpcv1.VolumeAttachmentReferenceVolumeContext) (volumeAttachmentsMap map[string]interface{}) {
	volumeAttachmentsMap = map[string]interface{}{}

	if volumeAttachmentsItem.DeleteVolumeOnInstanceDelete != nil {
		volumeAttachmentsMap["delete_volume_on_instance_delete"] = volumeAttachmentsItem.DeleteVolumeOnInstanceDelete
	}
	if volumeAttachmentsItem.Deleted != nil {
		deletedList := []map[string]interface{}{}
		deletedMap := dataSourceVolumeCollectionVolumeAttachmentsDeletedToMap(*volumeAttachmentsItem.Deleted)
		deletedList = append(deletedList, deletedMap)
		volumeAttachmentsMap["deleted"] = deletedList
	}
	if volumeAttachmentsItem.Device != nil {
		deviceList := []map[string]interface{}{}
		deviceMap := dataSourceVolumeCollectionVolumeAttachmentsDeviceToMap(*volumeAttachmentsItem.Device)
		deviceList = append(deviceList, deviceMap)
		volumeAttachmentsMap["device"] = deviceList
	}
	if volumeAttachmentsItem.Href != nil {
		volumeAttachmentsMap["href"] = volumeAttachmentsItem.Href
	}
	if volumeAttachmentsItem.ID != nil {
		volumeAttachmentsMap["id"] = volumeAttachmentsItem.ID
	}
	if volumeAttachmentsItem.Instance != nil {
		instanceList := []map[string]interface{}{}
		instanceMap := dataSourceVolumeCollectionVolumeAttachmentsInstanceToMap(*volumeAttachmentsItem.Instance)
		instanceList = append(instanceList, instanceMap)
		volumeAttachmentsMap["instance"] = instanceList
	}
	if volumeAttachmentsItem.Name != nil {
		volumeAttachmentsMap["name"] = volumeAttachmentsItem.Name
	}
	if volumeAttachmentsItem.Type != nil {
		volumeAttachmentsMap["type"] = volumeAttachmentsItem.Type
	}

	return volumeAttachmentsMap
}

func dataSourceVolumeCollectionVolumeAttachmentsDeletedToMap(deletedItem vpcv1.VolumeAttachmentReferenceVolumeContextDeleted) (deletedMap map[string]interface{}) {
	deletedMap = map[string]interface{}{}

	if deletedItem.MoreInfo != nil {
		deletedMap["more_info"] = deletedItem.MoreInfo
	}

	return deletedMap
}

func dataSourceVolumeCollectionVolumeAttachmentsDeviceToMap(deviceItem vpcv1.VolumeAttachmentDevice) (deviceMap map[string]interface{}) {
	deviceMap = map[string]interface{}{}

	if deviceItem.ID != nil {
		deviceMap["id"] = deviceItem.ID
	}

	return deviceMap
}

func dataSourceVolumeCollectionVolumeAttachmentsInstanceToMap(instanceItem vpcv1.InstanceReference) (instanceMap map[string]interface{}) {
	instanceMap = map[string]interface{}{}

	if instanceItem.CRN != nil {
		instanceMap["crn"] = instanceItem.CRN
	}
	if instanceItem.Deleted != nil {
		deletedList := []map[string]interface{}{}
		deletedMap := dataSourceVolumeCollectionInstanceDeletedToMap(*instanceItem.Deleted)
		deletedList = append(deletedList, deletedMap)
		instanceMap["deleted"] = deletedList
	}
	if instanceItem.Href != nil {
		instanceMap["href"] = instanceItem.Href
	}
	if instanceItem.ID != nil {
		instanceMap["id"] = instanceItem.ID
	}
	if instanceItem.Name != nil {
		instanceMap["name"] = instanceItem.Name
	}

	return instanceMap
}

func dataSourceVolumeCollectionInstanceDeletedToMap(deletedItem vpcv1.InstanceReferenceDeleted) (deletedMap map[string]interface{}) {
	deletedMap = map[string]interface{}{}

	if deletedItem.MoreInfo != nil {
		deletedMap["more_info"] = deletedItem.MoreInfo
	}

	return deletedMap
}

func dataSourceVolumeCollectionVolumesZoneToMap(zoneItem vpcv1.ZoneReference) (zoneMap map[string]interface{}) {
	zoneMap = map[string]interface{}{}

	if zoneItem.Href != nil {
		zoneMap["href"] = zoneItem.Href
	}
	if zoneItem.Name != nil {
		zoneMap["name"] = zoneItem.Name
	}

	return zoneMap
}
