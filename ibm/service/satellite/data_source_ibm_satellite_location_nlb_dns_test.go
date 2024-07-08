// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package satellite_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMSatelliteLocationNLBDNSListBasic(t *testing.T) {
	name := fmt.Sprintf("tf-satellitelocation-%d", acctest.RandIntRange(10, 100))
	managed_from := "wdc04"
	physical_address := "test location address"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMSatelliteLocationNLBDNSListConfig(name, managed_from, physical_address),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_satellite_location_nlb_dns.dns_list", "id"),
				),
			},
		},
	})
}

func testAccCheckIBMSatelliteLocationNLBDNSListConfig(name, managed_from, physical_address string) string {
	return testAccCheckSatelliteLocationDataSource(name, managed_from, physical_address) + `
	data ibm_satellite_location_nlb_dns dns_list {
		location = ibm_satellite_location.location.id
	}
`
}
