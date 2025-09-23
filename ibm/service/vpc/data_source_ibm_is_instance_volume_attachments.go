// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/flex"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceIBMISInstanceVolumeAttachments() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMISInstanceVolumeAttachmentsRead,

		Schema: map[string]*schema.Schema{
			isInstanceId: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Instance id",
			},

			isInstanceVolumeAttachments: {
				Type:        schema.TypeList,
				Description: "List of volume attachments on an instance",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						isInstanceVolumeDeleteOnInstanceDelete: {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "If set to true, when deleting the instance the volume will also be deleted.",
						},
						isInstanceVolAttName: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The user-defined name for this volume attachment.",
						},
						"bandwidth": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The maximum bandwidth (in megabits per second) for the volume when attached to this instance. This may be lower than the volume bandwidth depending on the configuration of the instance.",
						},

						isInstanceVolumeAttDevice: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "A unique identifier for the device which is exposed to the instance operating system",
						},

						isInstanceVolumeAttHref: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The URL for this volume attachment",
						},

						isInstanceVolAttId: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The unique identifier for this volume attachment",
						},

						isInstanceVolumeAttStatus: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The status of this volume attachment, one of [ attached, attaching, deleting, detaching ]",
						},

						isInstanceVolumeAttType: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The type of volume attachment one of [ boot, data ]",
						},
						isInstanceVolumeAttVolumeReferenceId: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The unique identifier for this volume",
						},
						isInstanceVolumeAttVolumeReferenceName: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The unique user-defined name for this volume",
						},
						isInstanceVolumeAttVolumeReferenceCrn: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The CRN for this volume",
						},
						isInstanceVolumeAttVolumeReferenceDeleted: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Link to documentation about deleted resources",
						},
						isInstanceVolumeAttVolumeReferenceHref: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The URL for this volume",
						},
						isInstanceVolumeAttVolumeReference: {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The attached volume",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{},
							},
						},
					},
				},
			},
		},
	}
}

func dataSourceIBMISInstanceVolumeAttachmentsRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	instanceId := d.Get(isInstanceId).(string)

	err := instanceGetVolumeAttachments(context, d, meta, instanceId)
	if err != nil {
		return err
	}

	return nil
}

func instanceGetVolumeAttachments(context context.Context, d *schema.ResourceData, meta interface{}, instanceId string) diag.Diagnostics {
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_is_instance_volume_attachments", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	allrecs := []vpcv1.VolumeAttachment{}
	listInstanceVolumeAttOptions := &vpcv1.ListInstanceVolumeAttachmentsOptions{
		InstanceID: &instanceId,
	}
	volumeAtts, _, err := sess.ListInstanceVolumeAttachmentsWithContext(context, listInstanceVolumeAttOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("ListInstanceVolumeAttachmentsWithContext failed: %s", err.Error()), "(Data) ibm_is_instance_volume_attachments", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	allrecs = append(allrecs, volumeAtts.VolumeAttachments...)
	volAttList := make([]map[string]interface{}, 0)
	for _, volumeAtt := range allrecs {
		currentVolAtt := map[string]interface{}{}
		currentVolAtt[isInstanceVolAttName] = *volumeAtt.Name
		// bandwidth changes
		currentVolAtt["bandwidth"] = *volumeAtt.Bandwidth
		currentVolAtt[isInstanceVolumeDeleteOnInstanceDelete] = *volumeAtt.DeleteVolumeOnInstanceDelete
		if volumeAtt.Device != nil {
			currentVolAtt[isInstanceVolumeAttDevice] = volumeAtt.Device.ID
		}
		currentVolAtt[isInstanceVolumeAttHref] = *volumeAtt.Href
		currentVolAtt[isInstanceVolAttId] = *volumeAtt.ID
		currentVolAtt[isInstanceVolumeAttStatus] = *volumeAtt.Status
		currentVolAtt[isInstanceVolumeAttType] = *volumeAtt.Type

		if volumeAtt.Volume != nil {
			currentVolAtt[isInstanceVolumeAttVolumeReferenceId] = *volumeAtt.Volume.ID
			currentVolAtt[isInstanceVolumeAttVolumeReferenceName] = *volumeAtt.Volume.Name
			currentVolAtt[isInstanceVolumeAttVolumeReferenceCrn] = *volumeAtt.Volume.CRN
			if volumeAtt.Volume.Deleted != nil {
				currentVolAtt[isInstanceVolumeAttVolumeReferenceDeleted] = *volumeAtt.Volume.Deleted.MoreInfo
			}
			currentVolAtt[isInstanceVolumeAttVolumeReferenceHref] = *volumeAtt.Volume.Href
		}

		volAttList = append(volAttList, currentVolAtt)
	}
	d.SetId(dataSourceIBMISInstanceVolumeAttachmentsID(d))
	if err = d.Set("volume_attachments", volAttList); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting volume_attachments: %s", err), "(Data) ibm_is_instance_volume_attachments", "read", "set-volume_attachments").GetDiag()
	}
	return nil
}

// dataSourceIBMISInstanceVolumeAttachmentsID returns a reasonable ID for a Instance volume attachments list.
func dataSourceIBMISInstanceVolumeAttachmentsID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}
