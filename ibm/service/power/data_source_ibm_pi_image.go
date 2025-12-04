// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package power

import (
	"context"
	"fmt"
	"log"

	"github.com/IBM-Cloud/power-go-client/clients/instance"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func DataSourceIBMPIImage() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMPIImagesRead,
		Schema: map[string]*schema.Schema{
			// Arguments
			Arg_CloudInstanceID: {
				Description:  "The GUID of the service instance associated with an account.",
				Required:     true,
				Type:         schema.TypeString,
				ValidateFunc: validation.NoZeroValues,
			},
			Arg_ImageID: {
				AtLeastOneOf:  []string{Arg_ImageID, Arg_ImageName},
				ConflictsWith: []string{Arg_ImageName},
				Description:   "The image ID.",
				Optional:      true,
				Type:          schema.TypeString,
				ValidateFunc:  validation.NoZeroValues,
			},
			Arg_ImageName: {
				AtLeastOneOf:  []string{Arg_ImageID, Arg_ImageName},
				ConflictsWith: []string{Arg_ImageID},
				Deprecated:    "The pi_image_name field is deprecated. Please use pi_image_id instead",
				Description:   "The name of the image.",
				Optional:      true,
				Type:          schema.TypeString,
				ValidateFunc:  validation.NoZeroValues,
			},

			// Attributes
			Attr_Architecture: {
				Computed:    true,
				Description: "The CPU architecture that the image is designed for.",
				Type:        schema.TypeString,
			},
			Attr_ContainerFormat: {
				Computed:    true,
				Description: "The container format.",
				Type:        schema.TypeString,
			},
			Attr_CRN: {
				Computed:    true,
				Description: "The CRN of this resource.",
				Type:        schema.TypeString,
			},
			Attr_DiskFormat: {
				Computed:    true,
				Description: "The disk format.",
				Type:        schema.TypeString,
			},
			Attr_Endianness: {
				Computed:    true,
				Description: "The endianness order.",
				Type:        schema.TypeString,
			},
			Attr_Hypervisor: {
				Computed:    true,
				Description: "Hypervision Type.",
				Type:        schema.TypeString,
			},
			Attr_ImageType: {
				Computed:    true,
				Description: "The identifier of this image type.",
				Type:        schema.TypeString,
			},
			Attr_Name: {
				Computed:    true,
				Description: "The name of an image.",
				Type:        schema.TypeString,
			},
			Attr_OperatingSystem: {
				Computed:    true,
				Description: "The operating system that is installed with the image.",
				Type:        schema.TypeString,
			},
			Attr_Shared: {
				Computed:    true,
				Description: "Indicates whether the image is shared.",
				Type:        schema.TypeBool,
			},
			Attr_Size: {
				Computed:    true,
				Description: "The size of the image in megabytes.",
				Type:        schema.TypeInt,
			},
			Attr_SourceChecksum: {
				Computed:    true,
				Description: "Checksum of the image.",
				Type:        schema.TypeString,
			},
			Attr_State: {
				Computed:    true,
				Description: "The state for this image. ",
				Type:        schema.TypeString,
			},
			Attr_StoragePool: {
				Computed:    true,
				Description: "Storage pool where image resides.",
				Type:        schema.TypeString,
			},
			Attr_StorageType: {
				Computed:    true,
				Description: "The storage type for this image.",
				Type:        schema.TypeString,
			},
			Attr_UserTags: {
				Computed:    true,
				Description: "List of user tags attached to the resource.",
				Elem:        &schema.Schema{Type: schema.TypeString},
				Set:         schema.HashString,
				Type:        schema.TypeSet,
			},
			Attr_Volumes: {
				Computed:    true,
				Description: "List of image volumes.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						Attr_Bootable: {
							Computed:    true,
							Description: "Indicates if the volume is boot capable.",
							Type:        schema.TypeBool,
						},
						Attr_Name: {
							Computed:    true,
							Description: "The volume name of the image.",
							Type:        schema.TypeString,
						},
						Attr_Size: {
							Computed:    true,
							Description: "The volume size of the image.",
							Type:        schema.TypeFloat,
						},
						Attr_VolumeID: {
							Computed:    true,
							Description: "The volume ID of the image.",
							Type:        schema.TypeString,
						},
					},
				},
				Type: schema.TypeList,
			},
		},
	}
}

func dataSourceIBMPIImagesRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	sess, err := meta.(conns.ClientSession).IBMPISession()
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("IBMPISession failed: %s", err.Error()), "(Data) ibm_pi_image", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	cloudInstanceID := d.Get(Arg_CloudInstanceID).(string)
	var imageID string
	if v, ok := d.GetOk(Arg_ImageID); ok {
		imageID = v.(string)
	} else if v, ok := d.GetOk(Arg_ImageName); ok {
		imageID = v.(string)
	}

	imageC := instance.NewIBMPIImageClient(ctx, sess, cloudInstanceID)
	imagedata, err := imageC.Get(imageID)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Get failed: %s", err.Error()), "(Data) ibm_pi_image", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(*imagedata.ImageID)
	d.Set(Attr_Architecture, imagedata.Specifications.Architecture)
	d.Set(Attr_ContainerFormat, imagedata.Specifications.ContainerFormat)
	if imagedata.Crn != "" {
		d.Set(Attr_CRN, imagedata.Crn)
		tags, err := flex.GetGlobalTagsUsingCRN(meta, string(imagedata.Crn), "", UserTagType)
		if err != nil {
			log.Printf("Error on get of pi image (%s) user_tags: %s", *imagedata.ImageID, err)
		}
		d.Set(Attr_UserTags, tags)
	}
	d.Set(Attr_DiskFormat, imagedata.Specifications.DiskFormat)
	d.Set(Attr_Endianness, imagedata.Specifications.Endianness)
	d.Set(Attr_Hypervisor, imagedata.Specifications.HypervisorType)
	d.Set(Attr_ImageType, imagedata.Specifications.ImageType)
	d.Set(Attr_Name, imagedata.Name)
	d.Set(Attr_OperatingSystem, imagedata.Specifications.OperatingSystem)
	d.Set(Attr_Shared, imagedata.Specifications.Shared)
	d.Set(Attr_Size, imagedata.Size)
	d.Set(Attr_SourceChecksum, imagedata.Specifications.SourceChecksum)
	d.Set(Attr_State, imagedata.State)
	d.Set(Attr_StoragePool, imagedata.StoragePool)
	d.Set(Attr_StorageType, imagedata.StorageType)
	volumeMap := []map[string]any{}
	if imagedata.Volumes != nil {
		for _, n := range imagedata.Volumes {
			if n != nil {
				v := map[string]any{
					Attr_Bootable: n.Bootable,
					Attr_Name:     n.Name,
					Attr_Size:     n.Size,
					Attr_VolumeID: n.VolumeID,
				}
				volumeMap = append(volumeMap, v)
			}
		}
	}
	d.Set(Attr_Volumes, volumeMap)
	return nil
}
