// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package power

import (
	"context"

	"github.com/IBM-Cloud/power-go-client/power/models"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/hashicorp/go-uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	//"fmt"
	"github.com/IBM-Cloud/power-go-client/clients/instance"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/*
Datasource to get the list of images that are available when a power instance is created
*/

// Attributes and Arguments defined in data_source_ibm_pi_image.go
func DataSourceIBMPIImages() *schema.Resource {

	return &schema.Resource{
		ReadContext: dataSourceIBMPIImagesAllRead,
		Schema: map[string]*schema.Schema{

			PICloudInstanceID: {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.NoZeroValues,
			},

			// Computed Attributes

			Images: {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						ImagesID: {
							Type:     schema.TypeString,
							Computed: true,
						},
						ImageName: {
							Type:     schema.TypeString,
							Computed: true,
						},
						ImageHref: {
							Type:     schema.TypeString,
							Computed: true,
						},
						ImageState: {
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
				},
			},
		},
	}
}

func dataSourceIBMPIImagesAllRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := meta.(conns.ClientSession).IBMPISession()
	if err != nil {
		return diag.FromErr(err)
	}

	cloudInstanceID := d.Get(PICloudInstanceID).(string)

	imageC := instance.NewIBMPIImageClient(ctx, sess, cloudInstanceID)
	imagedata, err := imageC.GetAll()
	if err != nil {
		return diag.FromErr(err)
	}

	var clientgenU, _ = uuid.GenerateUUID()
	d.SetId(clientgenU)
	d.Set(Images, flattenStockImages(imagedata.Images))

	return nil

}

func flattenStockImages(list []*models.ImageReference) []map[string]interface{} {
	result := make([]map[string]interface{}, 0, len(list))
	for _, i := range list {

		l := map[string]interface{}{
			ImagesID:         *i.ImageID,
			ImageState:       *i.State,
			ImageHref:        *i.Href,
			ImageName:        *i.Name,
			ImageStorageType: *i.StorageType,
			ImageStoragePool: *i.StoragePool,
			ImageType:        i.Specifications.ImageType,
		}

		result = append(result, l)

	}
	return result
}
