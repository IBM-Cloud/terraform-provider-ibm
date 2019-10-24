package ibm

import (
	"time"

	"github.com/hashicorp/terraform/helper/schema"
	"github.ibm.com/Bluemix/riaas-go-client/clients/geography"
)

const (
	isZoneNames = "zones"
)

func dataSourceIBMISZones() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceIBMISZonesRead,

		Schema: map[string]*schema.Schema{

			isZoneRegion: {
				Type:     schema.TypeString,
				Required: true,
			},

			isZoneStatus: {
				Type:     schema.TypeString,
				Optional: true,
			},

			isZoneNames: {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func dataSourceIBMISZonesRead(d *schema.ResourceData, meta interface{}) error {
	sess, err := meta.(ClientSession).ISSession()
	if err != nil {
		return err
	}
	zoneC := geography.NewZoneClient(sess)
	availableZones, err := zoneC.List(d.Get(isZoneRegion).(string))
	if err != nil {
		return err
	}
	names := make([]string, 0)
	status := d.Get(isZoneStatus).(string)
	for _, zone := range availableZones {
		if status == "" || zone.Status == status {
			names = append(names, zone.Name)
		}
	}
	d.SetId(dataSourceIBMISZonesId(d))
	d.Set(isZoneNames, names)
	return nil
}

// dataSourceIBMISZonesId returns a reasonable ID for a zone list.
func dataSourceIBMISZonesId(d *schema.ResourceData) string {
	// Our zone list is not guaranteed to be stable because the content
	// of the list can vary between two calls if any of the following
	// events occur between calls:
	// - a zone is added to our region
	// - a zone is dropped from our region
	// - we are using the status filter and the status of one or more
	//   zones changes between calls.
	//
	// For simplicity we are using a timestamp for the required terraform id.
	// If we find through usage that this choice is too ephemeral for our users
	// then we can change this function to use a more stable id, perhaps
	// composed from a hash of the list contents. But, for now, a timestamp
	// is good enough.
	return time.Now().UTC().String()
}
