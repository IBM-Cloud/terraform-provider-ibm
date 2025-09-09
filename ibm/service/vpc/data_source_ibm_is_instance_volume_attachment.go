// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc

import (
	"context"
	"fmt"
	"log"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	isInstanceVolumeAttDevice                 = "device"
	isInstanceVolumeAttHref                   = "href"
	isInstanceVolumeAttStatus                 = "status"
	isInstanceVolumeAttType                   = "type"
	isInstanceVolumeAttVolumeReference        = "volume_reference"
	isInstanceVolumeAttVolumeReferenceCrn     = "volume_crn"
	isInstanceVolumeAttVolumeReferenceDeleted = "volume_deleted"
	isInstanceVolumeAttVolumeReferenceHref    = "volume_href"
	isInstanceVolumeAttVolumeReferenceId      = "volume_id"
	isInstanceVolumeAttVolumeReferenceName    = "volume_name"
)

func DataSourceIBMISInstanceVolumeAttachment() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMISInstanceVolumeAttachmentRead,

		Schema: map[string]*schema.Schema{

			isInstanceId: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Instance id",
			},
			isInstanceVolAttName: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The user-defined name for this volume attachment.",
			},

			isInstanceVolumeDeleteOnInstanceDelete: {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "If set to true, when deleting the instance the volume will also be deleted.",
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
			"bandwidth": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The maximum bandwidth (in megabits per second) for the volume when attached to this instance. This may be lower than the volume bandwidth depending on the configuration of the instance.",
			},

			isInstanceVolumeAttVolumeReference: {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The attached volume",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
					},
				},
			},
		},
	}
}

func dataSourceIBMISInstanceVolumeAttachmentRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	instanceId := d.Get(isInstanceId).(string)
	name := d.Get(isInstanceName).(string)
	err := instanceVolumeAttachmentGetByName(context, d, meta, instanceId, name)
	if err != nil {
		return err
	}
	return nil
}

func instanceVolumeAttachmentGetByName(context context.Context, d *schema.ResourceData, meta interface{}, instanceId, name string) diag.Diagnostics {
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_is_instance_volume_attachment", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	allrecs := []vpcv1.VolumeAttachment{}
	listInstanceVolumeAttOptions := &vpcv1.ListInstanceVolumeAttachmentsOptions{
		InstanceID: &instanceId,
	}
	volumeAtts, _, err := sess.ListInstanceVolumeAttachmentsWithContext(context, listInstanceVolumeAttOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("ListInstanceVolumeAttachmentsWithContext failed: %s", err.Error()), "(Data) ibm_is_instance_volume_attachment", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	allrecs = append(allrecs, volumeAtts.VolumeAttachments...)
	for _, volumeAttachment := range allrecs {
		if *volumeAttachment.Name == name {
			d.SetId(makeTerraformVolAttID(instanceId, *volumeAttachment.ID))
			if err = d.Set("name", volumeAttachment.Name); err != nil {
				return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting name: %s", err), "(Data) ibm_is_instance_volume_attachment", "read", "set-name").GetDiag()
			}
			// bandwidth changes
			d.Set("bandwidth", *volumeAttachment.Bandwidth)
			if err = d.Set("bandwidth", flex.IntValue(volumeAttachment.Bandwidth)); err != nil {
				return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting bandwidth: %s", err), "(Data) ibm_is_instance_volume_attachment", "read", "set-bandwidth").GetDiag()
			}
			if err = d.Set("delete_volume_on_instance_delete", volumeAttachment.DeleteVolumeOnInstanceDelete); err != nil {
				return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting delete_volume_on_instance_delete: %s", err), "(Data) ibm_is_instance_volume_attachment", "read", "set-delete_volume_on_instance_delete").GetDiag()
			}
			if volumeAttachment.Device != nil {
				if err = d.Set("device", *volumeAttachment.Device.ID); err != nil {
					return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting device: %s", err), "(Data) ibm_is_instance_volume_attachment", "read", "set-device").GetDiag()
				}
			}
			if err = d.Set("href", volumeAttachment.Href); err != nil {
				return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting href: %s", err), "(Data) ibm_is_instance_volume_attachment", "read", "set-href").GetDiag()
			}
			if err = d.Set("volume_attachment_id", *volumeAttachment.ID); err != nil {
				return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting volume_attachment_id: %s", err), "(Data) ibm_is_instance_volume_attachment", "read", "set-volume_attachment_id").GetDiag()
			}
			if err = d.Set("status", volumeAttachment.Status); err != nil {
				return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting status: %s", err), "(Data) ibm_is_instance_volume_attachment", "read", "set-status").GetDiag()
			}
			if err = d.Set("type", volumeAttachment.Type); err != nil {
				return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting type: %s", err), "(Data) ibm_is_instance_volume_attachment", "read", "set-type").GetDiag()
			}
			volList := make([]map[string]interface{}, 0)
			if volumeAttachment.Volume != nil {
				currentVol := map[string]interface{}{}
				currentVol[isInstanceVolumeAttVolumeReferenceId] = *volumeAttachment.Volume.ID
				currentVol[isInstanceVolumeAttVolumeReferenceName] = *volumeAttachment.Volume.Name
				currentVol[isInstanceVolumeAttVolumeReferenceCrn] = *volumeAttachment.Volume.CRN
				if volumeAttachment.Volume.Deleted != nil {
					currentVol[isInstanceVolumeAttVolumeReferenceDeleted] = *volumeAttachment.Volume.Deleted.MoreInfo
				}
				currentVol[isInstanceVolumeAttVolumeReferenceHref] = *volumeAttachment.Volume.Href
				volList = append(volList, currentVol)
			}
			if err = d.Set(isInstanceVolumeAttVolumeReference, volList); err != nil {
				return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting volume_reference: %s", err), "(Data) ibm_is_instance_volume_attachment", "read", "set-volume_reference").GetDiag()
			}
			return nil
		}
	}
	err = fmt.Errorf("No Instance volume attachment found with name %s on instance %s", name, instanceId)
	tfErr := flex.TerraformErrorf(err, fmt.Sprintf("ListInstanceVolumeAttachmentsWithContext failed: %s", err.Error()), "(Data) ibm_is_instance_volume_attachment", "read")
	log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
	return tfErr.GetDiag()
}
