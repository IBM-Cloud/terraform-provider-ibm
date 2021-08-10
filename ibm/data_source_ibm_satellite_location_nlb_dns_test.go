// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMSatelliteLocationNLBDNSListBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMSatelliteLocationNLBDNSListConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_satellite_location_nlb_dns.dns_list", "id"),
				),
			},
		},
	})
}

func testAccCheckIBMSatelliteLocationNLBDNSListConfig() string {
	return `
	data ibm_satellite_location_nlb_dns dns_list {
		location ="cp-location"
	}
`
}
