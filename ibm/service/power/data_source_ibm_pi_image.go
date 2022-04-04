// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package power

import (
	"context"

	"github.com/IBM-Cloud/power-go-client/helpers"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	//"fmt"
	"github.com/IBM-Cloud/power-go-client/clients/instance"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func DataSourceIBMPIImage() *schema.Resource {

	return &schema.Resource{
		ReadContext: dataSourceIBMPIImagesRead,
		Schema: map[string]*schema.Schema{

			PIImageName: {
				Type:         schema.TypeString,
				Required:     true,
				Description:  "Imagename Name to be used for pvminstances",
				ValidateFunc: validation.NoZeroValues,
			},
			PICloudInstanceID: {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.NoZeroValues,
			},

			ImageState: {
				Type:     schema.TypeString,
				Computed: true,
			},
			ImageSize: {
				Type:     schema.TypeInt,
				Computed: true,
			},
			ImageArchitecture: {
				Type:     schema.TypeString,
				Computed: true,
			},
			ImageOperatingSystem: {
				Type:     schema.TypeString,
				Computed: true,
			},
			ImageHyperVisor: {
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
			ImageType: {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceIBMPIImagesRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := meta.(conns.ClientSession).IBMPISession()
	if err != nil {
		return diag.FromErr(err)
	}

	cloudInstanceID := d.Get(helpers.PICloudInstanceId).(string)

	imageC := instance.NewIBMPIImageClient(ctx, sess, cloudInstanceID)
	imagedata, err := imageC.Get(d.Get(helpers.PIImageName).(string))
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(*imagedata.ImageID)
	d.Set(ImageState, imagedata.State)
	d.Set(ImageSize, imagedata.Size)
	d.Set(ImageArchitecture, imagedata.Specifications.Architecture)
	d.Set(ImageHyperVisor, imagedata.Specifications.HypervisorType)
	d.Set(ImageOperatingSystem, imagedata.Specifications.OperatingSystem)
	d.Set(ImageStorageType, imagedata.StorageType)
	d.Set(ImageStoragePool, imagedata.StoragePool)
	d.Set(ImageType, imagedata.Specifications.ImageType)

	return nil

}
