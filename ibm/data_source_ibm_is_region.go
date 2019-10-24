package ibm

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.ibm.com/Bluemix/riaas-go-client/clients/geography"
)

const (
	isRegionEndpoint = "endpoint"
	isRegionName     = "name"
	isRegionStatus   = "status"
)

func dataSourceIBMISRegion() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceIBMISRegionRead,

		Schema: map[string]*schema.Schema{

			isRegionEndpoint: {
				Type:     schema.TypeString,
				Computed: true,
			},

			isRegionName: {
				Type:     schema.TypeString,
				Required: true,
			},

			isRegionStatus: {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceIBMISRegionRead(d *schema.ResourceData, meta interface{}) error {
	sess, err := meta.(ClientSession).ISSession()
	if err != nil {
		return err
	}
	regionC := geography.NewRegionClient(sess)
	region, err := regionC.Get(d.Get("name").(string))
	if err != nil {
		return err
	}
	d.SetId(region.Name)
	d.Set(isRegionEndpoint, region.Endpoint)
	d.Set(isRegionName, region.Name)
	d.Set(isRegionStatus, region.Status)
	return nil
}
