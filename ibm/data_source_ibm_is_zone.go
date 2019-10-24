package ibm

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	"github.ibm.com/Bluemix/riaas-go-client/clients/geography"
)

const (
	isZoneName   = "name"
	isZoneRegion = "region"
	isZoneStatus = "status"
)

func dataSourceIBMISZone() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceIBMISZoneRead,

		Schema: map[string]*schema.Schema{

			isZoneName: {
				Type:     schema.TypeString,
				Required: true,
			},

			isZoneRegion: {
				Type:     schema.TypeString,
				Required: true,
			},

			isZoneStatus: {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceIBMISZoneRead(d *schema.ResourceData, meta interface{}) error {
	sess, err := meta.(ClientSession).ISSession()
	if err != nil {
		return err
	}
	zoneC := geography.NewZoneClient(sess)
	zone, err := zoneC.Get(d.Get(isZoneRegion).(string), d.Get(isZoneName).(string))
	if err != nil {
		return err
	}
	// For lack of anything better, compose our id from region name + zone name.
	id := fmt.Sprintf("%s.%s", zone.Region.Name, zone.Name)
	d.SetId(id)
	d.Set(isZoneName, zone.Name)
	d.Set(isZoneRegion, zone.Region.Name)
	d.Set(isZoneStatus, zone.Status)
	return nil
}
