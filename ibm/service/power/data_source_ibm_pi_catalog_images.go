// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package power

import (
	"context"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	"github.com/IBM-Cloud/power-go-client/clients/instance"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
)

/*
Datasource to get the list of images that are available when a power instance is created

*/
func DataSourceIBMPICatalogImages() *schema.Resource {

	return &schema.Resource{
		ReadContext: dataSourceIBMPICatalogImagesRead,
		Schema: map[string]*schema.Schema{

			PICloudInstanceID: {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.NoZeroValues,
			},
			CatalogImagesSAP: {
				Type:     schema.TypeBool,
				Optional: true,
			},
			CatalogImagesVTL: {
				Type:     schema.TypeBool,
				Optional: true,
			},
			CatalogImages: {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						ImageID: {
							Type:     schema.TypeString,
							Computed: true,
						},
						ImageName: {
							Type:     schema.TypeString,
							Computed: true,
						},
						ImageState: {
							Type:     schema.TypeString,
							Computed: true,
						},
						ImageDescription: {
							Type:     schema.TypeString,
							Computed: true,
						},
						ImageStorageType: {
							Type:     schema.TypeString,
							Computed: true,
						},
						ImageStoragePool: {
							Type:     schema.TypeString,
							Computed: true,
						},
						ImageCreationDate: {
							Type:     schema.TypeString,
							Computed: true,
						},
						ImageLastUpdateDate: {
							Type:     schema.TypeString,
							Computed: true,
						},
						ImageType: {
							Type:     schema.TypeString,
							Computed: true,
						},
						ImageContainerFormat: {
							Type:     schema.TypeString,
							Computed: true,
						},
						ImageDiskFormat: {
							Type:     schema.TypeString,
							Computed: true,
						},
						CatalogImageOperatingSystem: {
							Type:     schema.TypeString,
							Computed: true,
						},
						ImageHypervisorType: {
							Type:     schema.TypeString,
							Computed: true,
						},
						ImageArchitecture: {
							Type:     schema.TypeString,
							Computed: true,
						},
						ImageEndianness: {
							Type:     schema.TypeString,
							Computed: true,
						},
						ImageHref: {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceIBMPICatalogImagesRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := meta.(conns.ClientSession).IBMPISession()
	if err != nil {
		return diag.FromErr(err)
	}

	cloudInstanceID := d.Get(PICloudInstanceID).(string)
	includeSAP := false
	if s, ok := d.GetOk(CatalogImagesSAP); ok {
		includeSAP = s.(bool)
	}
	includeVTL := false
	if v, ok := d.GetOk(CatalogImagesVTL); ok {
		includeVTL = v.(bool)
	}
	imageC := instance.NewIBMPIImageClient(ctx, sess, cloudInstanceID)
	stockImages, err := imageC.GetAllStockImages(includeSAP, includeVTL)
	if err != nil {
		return diag.FromErr(err)
	}

	images := make([]map[string]interface{}, 0)
	for _, i := range stockImages.Images {
		image := make(map[string]interface{})
		image[ImageID] = *i.ImageID
		image[ImageName] = *i.Name
		if i.State != nil {
			image[ImageState] = *i.State
		}
		if i.Description != nil {
			image[ImageDescription] = *i.Description
		}
		if i.StorageType != nil {
			image[ImageStorageType] = *i.StorageType
		}
		if i.StoragePool != nil {
			image[ImageStoragePool] = *i.StoragePool
		}
		if i.CreationDate != nil {
			image[ImageCreationDate] = i.CreationDate.String()
		}
		if i.LastUpdateDate != nil {
			image[ImageLastUpdateDate] = i.LastUpdateDate.String()
		}
		if i.Href != nil {
			image[ImageHref] = *i.Href
		}
		if i.Specifications != nil {
			s := i.Specifications
			if &s.ImageType != nil {
				image[ImageType] = s.ImageType
			}
			if &s.ContainerFormat != nil {
				image[ImageContainerFormat] = s.ContainerFormat
			}
			if &s.DiskFormat != nil {
				image[ImageDiskFormat] = s.DiskFormat
			}
			if &s.OperatingSystem != nil {
				image[CatalogImageOperatingSystem] = s.OperatingSystem
			}
			if &s.HypervisorType != nil {
				image[ImageHypervisorType] = s.HypervisorType
			}
			if &s.Architecture != nil {
				image[ImageArchitecture] = s.Architecture
			}
			if &s.Endianness != nil {
				image[ImageEndianness] = s.Endianness
			}
		}
		images = append(images, image)
	}
	d.SetId(time.Now().UTC().String())
	d.Set(CatalogImages, images)
	return nil

}
