// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package satellite_test

import (
	"fmt"
	"testing"

	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIbmSatelliteLinkDataSourceBasic(t *testing.T) {
	locationID := fmt.Sprintf("tf-satellite-loc-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIbmSatelliteLinkDataSourceConfigBasic(locationID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_satellite_link.satellite_link", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_satellite_link.satellite_link", "location"),
					resource.TestCheckResourceAttrSet("data.ibm_satellite_link.satellite_link", "ws_endpoint"),
					resource.TestCheckResourceAttrSet("data.ibm_satellite_link.satellite_link", "crn"),
					resource.TestCheckResourceAttrSet("data.ibm_satellite_link.satellite_link", "description"),
					resource.TestCheckResourceAttrSet("data.ibm_satellite_link.satellite_link", "satellite_link_host"),
					resource.TestCheckResourceAttrSet("data.ibm_satellite_link.satellite_link", "status"),
					resource.TestCheckResourceAttrSet("data.ibm_satellite_link.satellite_link", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_satellite_link.satellite_link", "last_change"),
					resource.TestCheckResourceAttrSet("data.ibm_satellite_link.satellite_link", "performance.#"),
				),
			},
		},
	})
}

func testAccCheckIbmSatelliteLinkDataSourceConfigBasic(locationID string) string {
	return fmt.Sprintf(`
		data "ibm_satellite_link" "satellite_link" {
			location = "%s"
		}
	`, locationID)
}
