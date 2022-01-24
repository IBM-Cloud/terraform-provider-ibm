// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package classicinfrastructure

import (
	"fmt"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/softlayer/softlayer-go/services"
)

func DataSourceIBMDNSSecondary() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceIBMDNSSecondaryRead,

		Schema: map[string]*schema.Schema{

			"zone_name": {
				Description: "The name of the secondary",
				Type:        schema.TypeString,
				Required:    true,
			},

			"master_ip_address": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"transfer_frequency": {
				Type:     schema.TypeInt,
				Computed: true,
			},

			"status_id": {
				Type:     schema.TypeInt,
				Computed: true,
			},

			"status_text": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceIBMDNSSecondaryRead(d *schema.ResourceData, meta interface{}) error {
	sess := meta.(conns.ClientSession).SoftLayerSession()
	service := services.GetAccountService(sess)

	name := d.Get("zone_name").(string)

	names, err := service.
		Mask("id, masterIpAddress, transferFrequency, zoneName, statusId, statusText").
		GetSecondaryDomains()

	if err != nil {
		return fmt.Errorf("[ERROR] Error retrieving secondary zone: %s", err)
	}

	if len(names) == 0 {
		return fmt.Errorf("[ERROR] No secondary zone found with name: %s", name)
	}

	for _, zone := range names {
		if name == *zone.ZoneName {
			d.SetId(fmt.Sprintf("%d", *zone.Id))
			d.Set("master_ip_address", *zone.MasterIpAddress)
			d.Set("transfer_frequency", *zone.TransferFrequency)
			d.Set("zone_name", *zone.ZoneName)
			d.Set("status_id", *zone.StatusId)
			d.Set("status_text", *zone.StatusText)
			return nil

		}
	}
	return fmt.Errorf("[ERROR] No secondary zone found with name: %s", name)

}
