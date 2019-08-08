package ibm

import (
	"github.com/hashicorp/go-uuid"
	"github.com/hashicorp/terraform/helper/schema"
	//"fmt"
	"github.com/hashicorp/terraform/helper/validation"
	"github.ibm.com/Bluemix/power-go-client/clients/instance"
)

func dataSourceIBMPowerVolume() *schema.Resource {

	return &schema.Resource{
		Read: dataSourceIBMPowerVolumesRead,
		Schema: map[string]*schema.Schema{

			"volumename": {
				Type:         schema.TypeString,
				Required:     true,
				Description:  "Volume Name to be used for pvminstances",
				ValidateFunc: validation.NoZeroValues,
			},
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

func dataSourceIBMPowerVolumesRead(d *schema.ResourceData, meta interface{}) error {

	sess, err := meta.(ClientSession).PowerSession()
	if err != nil {
		return err
	}

	volumeC := instance.NewPowerVolumeClient(sess)
	volumedata, err := volumeC.Get(d.Get("volumename").(string))

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
	//return fmt.Errorf("No Image found with name %s", imagedata.)

}
