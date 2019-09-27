package ibm

import (
	"github.com/IBM-Cloud/power-go-client/helpers"
	"github.com/IBM-Cloud/power-go-client/power/models"
	"github.com/hashicorp/go-uuid"
	"github.com/hashicorp/terraform/helper/schema"

	//"fmt"
	"github.com/IBM-Cloud/power-go-client/clients/instance"
	"github.com/hashicorp/terraform/helper/validation"
)

func dataSourceIBMPIVolumes() *schema.Resource {

	return &schema.Resource{
		Read: dataSourceIBMPIVolumesRead,
		Schema: map[string]*schema.Schema{

			helpers.PIInstanceName: {
				Type:         schema.TypeString,
				Required:     true,
				Description:  "Volume Name to be used for pvminstances",
				ValidateFunc: validation.NoZeroValues,
			},

			helpers.PICloudInstanceId: {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.NoZeroValues,
			},

			//Computed Attributes

			"instance_volumes": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"volume_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"volume_size": {
							Type:     schema.TypeFloat,
							Computed: true,
						},
						"volume_href": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"volume_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"volume_state": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"volume_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"volume_shareable": {
							Type:     schema.TypeBool,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceIBMPIVolumesRead(d *schema.ResourceData, meta interface{}) error {

	sess, err := meta.(ClientSession).IBMPISession()
	if err != nil {
		return err
	}

	powerinstanceid := d.Get(helpers.PICloudInstanceId).(string)
	volumeC := instance.NewIBMPIVolumeClient(sess, powerinstanceid)
	volumedata, err := volumeC.GetAll(d.Get(helpers.PIInstanceName).(string), powerinstanceid)

	if err != nil {
		return err
	}

	var clientgenU, _ = uuid.GenerateUUID()
	d.SetId(clientgenU)

	d.Set("instance_volumes", flattenVolumesInInstances(volumedata.Volumes))

	return nil

}

func flattenVolumesInInstances(list []*models.VolumeReference) []map[string]interface{} {
	result := make([]map[string]interface{}, 0, len(list))
	for _, i := range list {
		l := map[string]interface{}{
			"volume_id":         *i.VolumeID,
			"volume_state":      *i.State,
			"volume_href":       *i.Href,
			"volume_name":       *i.Name,
			"volume_size_in_gb": *i.Size,
			"volume_type":       *i.DiskType,
			"volume_shareable":  *i.Shareable,
		}

		result = append(result, l)
	}
	return result
}
