package ibm

import (
	"github.com/hashicorp/go-uuid"
	"github.com/hashicorp/terraform/helper/schema"
	//"fmt"
	"github.com/hashicorp/terraform/helper/validation"
	"github.ibm.com/Bluemix/power-go-client/clients/instance"
)

func dataSourceIBMPowerImage() *schema.Resource {

	return &schema.Resource{
		Read: dataSourceIBMPowerImagesRead,
		Schema: map[string]*schema.Schema{

			"name": {
				Type:         schema.TypeString,
				Required:     true,
				Description:  "Imagename Name to be used for pvminstances",
				ValidateFunc: validation.NoZeroValues,
			},

			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"imageid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"size": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"architecture": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"operatingsystem": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"hypervisor": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceIBMPowerImagesRead(d *schema.ResourceData, meta interface{}) error {

	sess, err := meta.(ClientSession).PowerSession()
	if err != nil {
		return err
	}

	imageC := instance.NewPowerImageClient(sess)
	imagedata, err := imageC.Get(d.Get("name").(string))

	if err != nil {
		return err
	}

	var clientgenU, _ = uuid.GenerateUUID()
	d.SetId(clientgenU)
	d.Set("imageid", imagedata.ImageID)
	d.Set("state", imagedata.State)
	d.Set("size", imagedata.Size)
	d.Set("architecture", imagedata.Specifications.Architecture)
	d.Set("hypervisor", imagedata.Specifications.HypervisorType)

	return nil
	//return fmt.Errorf("No Image found with name %s", imagedata.)

}
