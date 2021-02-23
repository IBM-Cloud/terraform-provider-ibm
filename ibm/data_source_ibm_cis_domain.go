/* IBM Confidential
*  Object Code Only Source Materials
*  5747-SM3
*  (c) Copyright IBM Corp. 2017,2021
*
*  The source code for this program is not published or otherwise divested
*  of its trade secrets, irrespective of what has been deposited with the
*  U.S. Copyright Office. */

package ibm

import (
	"fmt"
	"log"

	"github.com/IBM/go-sdk-core/v4/core"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceIBMCISDomain() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceIBMCISDomainRead,

		Schema: map[string]*schema.Schema{
			cisID: {
				Type:        schema.TypeString,
				Description: "CIS object id",
				Required:    true,
			},
			cisDomain: {
				Type:        schema.TypeString,
				Description: "CISzone - Domain",
				Required:    true,
			},
			cisDomainPaused: {
				Type:     schema.TypeBool,
				Computed: true,
			},
			cisDomainStatus: {
				Type:     schema.TypeString,
				Computed: true,
			},
			cisDomainNameServers: {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			cisDomainOriginalNameServers: {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			cisDomainID: {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceIBMCISDomainRead(d *schema.ResourceData, meta interface{}) error {
	var zoneFound bool
	cisClient, err := meta.(ClientSession).CisZonesV1ClientSession()
	if err != nil {
		return err
	}

	crn := d.Get(cisID).(string)
	cisClient.Crn = core.StringPtr(crn)
	zoneName := d.Get(cisDomain).(string)

	opt := cisClient.NewListZonesOptions()
	opt.SetPage(1)       // list all zones in one page
	opt.SetPerPage(1000) // maximum allowed limit is 1000 per page
	zones, resp, err := cisClient.ListZones(opt)
	if err != nil {
		log.Printf("dataSourcCISdomainRead - ListZones Failed %s\n", resp)
		return err
	}

	for _, zone := range zones.Result {
		if *zone.Name == zoneName {
			d.SetId(convertCisToTfTwoVar(*zone.ID, crn))
			d.Set(cisID, crn)
			d.Set(cisDomain, *zone.Name)
			d.Set(cisDomainStatus, *zone.Status)
			d.Set(cisDomainPaused, *zone.Paused)
			d.Set(cisDomainNameServers, zone.NameServers)
			d.Set(cisDomainOriginalNameServers, zone.OriginalNameServers)
			d.Set(cisDomainID, *zone.ID)
			zoneFound = true
		}
	}

	if zoneFound == false {
		return fmt.Errorf("Given zone does not exist. Please specify correct domain")
	}

	return nil
}
