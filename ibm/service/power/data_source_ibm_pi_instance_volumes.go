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

	"github.com/IBM-Cloud/power-go-client/clients/instance"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

const (
	// Attributes
	PIInstanceVolumesName = "pi_volume_name"

	// Arguments
	InstanceVolumesBootVolumeID    = "boot_volume_id"
	InstanceVolumesInstanceVolumes = "instance_volumes"
	InstanceVolumesBootable        = "bootable"
	InstanceVolumesHref            = "href"
	InstanceVolumesID              = "id"
	InstanceVolumesName            = "name"
	InstanceVolumesPool            = "pool"
	InstanceVolumesShareable       = "shareable"
	InstanceVolumesSize            = "size"
	InstanceVolumesState           = "state"
	InstanceVolumesType            = "type"
)

func DataSourceIBMPIInstanceVolumes() *schema.Resource {

	return &schema.Resource{
		ReadContext: dataSourceIBMPIInstanceVolumesRead,
		Schema: map[string]*schema.Schema{
			PIInstanceVolumesName: {
				Type:         schema.TypeString,
				Required:     true,
				Description:  "Instance Name to be used for pvminstances",
				ValidateFunc: validation.NoZeroValues,
			},
			PICloudInstanceID: {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.NoZeroValues,
			},

			//Computed Attributes
			InstanceVolumesBootVolumeID: {
				Type:     schema.TypeString,
				Computed: true,
			},
			InstanceVolumesInstanceVolumes: {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						InstanceVolumesID: {
							Type:     schema.TypeString,
							Computed: true,
						},
						InstanceVolumesSize: {
							Type:     schema.TypeFloat,
							Computed: true,
						},
						InstanceVolumesHref: {
							Type:     schema.TypeString,
							Computed: true,
						},
						InstanceVolumesName: {
							Type:     schema.TypeString,
							Computed: true,
						},
						InstanceVolumesState: {
							Type:     schema.TypeString,
							Computed: true,
						},
						InstanceVolumesType: {
							Type:     schema.TypeString,
							Computed: true,
						},
						InstanceVolumesPool: {
							Type:     schema.TypeString,
							Computed: true,
						},
						InstanceVolumesShareable: {
							Type:     schema.TypeBool,
							Computed: true,
						},
						InstanceVolumesBootable: {
							Type:     schema.TypeBool,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceIBMPIInstanceVolumesRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := meta.(conns.ClientSession).IBMPISession()
	if err != nil {
		return diag.FromErr(err)
	}

	cloudInstanceID := d.Get(PICloudInstanceID).(string)

	volumeC := instance.NewIBMPIVolumeClient(ctx, sess, cloudInstanceID)
	volumedata, err := volumeC.GetAllInstanceVolumes(d.Get(PIInstanceVolumesName).(string))
	if err != nil {
		return diag.FromErr(err)
	}

	var clientgenU, _ = uuid.GenerateUUID()
	d.SetId(clientgenU)
	d.Set(InstanceVolumesBootVolumeID, *volumedata.Volumes[0].VolumeID)
	d.Set(InstanceVolumesInstanceVolumes, flattenVolumesInstances(volumedata.Volumes))

	return nil

}

func flattenVolumesInstances(list []*models.VolumeReference) []map[string]interface{} {
	result := make([]map[string]interface{}, 0, len(list))
	for _, i := range list {
		l := map[string]interface{}{
			InstanceVolumesID:        *i.VolumeID,
			InstanceVolumesState:     *i.State,
			InstanceVolumesHref:      *i.Href,
			InstanceVolumesName:      *i.Name,
			InstanceVolumesSize:      *i.Size,
			InstanceVolumesType:      *i.DiskType,
			InstanceVolumesPool:      i.VolumePool,
			InstanceVolumesShareable: *i.Shareable,
			InstanceVolumesBootable:  *i.Bootable,
		}

		result = append(result, l)
	}
	return result
}
