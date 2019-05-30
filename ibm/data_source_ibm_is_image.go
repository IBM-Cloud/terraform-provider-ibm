package ibm

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	"github.ibm.com/Bluemix/riaas-go-client/clients/compute"
)

func dataSourceIBMISImage() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceIBMISImageRead,

		Schema: map[string]*schema.Schema{

			"name": {
				Type:     schema.TypeString,
				Required: true,
			},

			"visibility": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validateAllowedStringValue([]string{"public", "private"}),
			},

			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"os": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"architecture": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"crn": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceIBMISImageRead(d *schema.ResourceData, meta interface{}) error {
	sess, err := meta.(ClientSession).ISSession()
	if err != nil {
		return err
	}
	imageC := compute.NewImageClient(sess)
	var visibility string
	if v, ok := d.GetOk("visibility"); ok {
		visibility = v.(string)
	}
	images, _, err := imageC.ListWithFilter("", visibility, "")
	if err != nil {
		return err
	}
	name := d.Get("name").(string)
	for _, image := range images {
		if image.Name == name {
			d.SetId(image.ID.String())
			d.Set("status", image.Status)
			d.Set("name", image.Name)
			d.Set("visibility", image.Visibility)
			d.Set("os", image.OperatingSystem.Name)
			d.Set("architecture", image.Architecture)
			d.Set("crn", image.Crn)
			return nil
		}
	}
	return fmt.Errorf("No Image found with name %s", name)
}
