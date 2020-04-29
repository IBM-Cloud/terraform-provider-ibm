package ibm

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.ibm.com/ibmcloud/vpc-go-sdk/vpcclassicv1"
	"github.ibm.com/ibmcloud/vpc-go-sdk/vpcv1"
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
	userDetails, err := meta.(ClientSession).BluemixUserDetails()
	if err != nil {
		return err
	}
	regionName := d.Get(isZoneRegion).(string)
	zoneName := d.Get(isZoneName).(string)
	if userDetails.generation == 1 {
		err := classicZoneGet(d, meta, regionName, zoneName)
		if err != nil {
			return err
		}
	} else {
		err := zoneGet(d, meta, regionName, zoneName)
		if err != nil {
			return err
		}
	}
	return nil
}

func classicZoneGet(d *schema.ResourceData, meta interface{}, regionName, zoneName string) error {
	sess, err := classicVpcClient(meta)
	if err != nil {
		return err
	}
	getZoneOptions := &vpcclassicv1.GetZoneOptions{
		RegionName: &regionName,
		ZoneName:   &zoneName,
	}
	zone, _, err := sess.GetZone(getZoneOptions)
	if err != nil {
		return err
	}
	// For lack of anything better, compose our id from region name + zone name.
	id := fmt.Sprintf("%s.%s", *zone.Region.Name, *zone.Name)
	d.SetId(id)
	d.Set(isZoneName, *zone.Name)
	d.Set(isZoneRegion, *zone.Region.Name)
	d.Set(isZoneStatus, *zone.Status)
	return nil
}

func zoneGet(d *schema.ResourceData, meta interface{}, regionName, zoneName string) error {
	sess, err := vpcClient(meta)
	if err != nil {
		return err
	}
	getZoneOptions := &vpcv1.GetZoneOptions{
		RegionName: &regionName,
		ZoneName:   &zoneName,
	}
	zone, _, err := sess.GetZone(getZoneOptions)
	if err != nil {
		return err
	}
	// For lack of anything better, compose our id from region name + zone name.
	id := fmt.Sprintf("%s.%s", *zone.Region.Name, *zone.Name)
	d.SetId(id)
	d.Set(isZoneName, *zone.Name)
	d.Set(isZoneRegion, *zone.Region.Name)
	d.Set(isZoneStatus, *zone.Status)
	return nil
}
