package ibm

import (
	"github.com/IBM-Cloud/power-go-client/helpers"
	"github.com/hashicorp/go-uuid"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	//"fmt"
	"github.com/IBM-Cloud/power-go-client/clients/instance"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func dataSourceIBMPIVolume() *schema.Resource {

	return &schema.Resource{
		Read: dataSourceIBMPIVolumeRead,
		Schema: map[string]*schema.Schema{

			helpers.PIVolumeName: {
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

			// Computed Attributes
			"volumeid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"size": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"shareable": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"bootable": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"creationdate": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"disktype": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceIBMPIVolumeRead(d *schema.ResourceData, meta interface{}) error {

	sess, err := meta.(ClientSession).IBMPISession()
	if err != nil {
		return err
	}

	powerinstanceid := d.Get(helpers.PICloudInstanceId).(string)
	volumeC := instance.NewIBMPIVolumeClient(sess, powerinstanceid)
	volumedata, err := volumeC.Get(d.Get(helpers.PIVolumeName).(string), powerinstanceid)

	if err != nil {
		return err
	}

	var clientgenU, _ = uuid.GenerateUUID()
	d.SetId(clientgenU)
	d.Set("volumeid", volumedata.VolumeID)
	d.Set("size", volumedata.Size)
	d.Set("disktype", volumedata.DiskType)
	d.Set("bootable", volumedata.Bootable)
	d.Set("state", volumedata.State)

	return nil

}
