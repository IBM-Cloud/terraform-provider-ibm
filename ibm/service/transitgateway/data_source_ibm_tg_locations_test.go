// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package transitgateway_test

import (
	"testing"

	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMTransitGatewaysLocationsDataSource_basic(t *testing.T) {
	resName := "data.ibm_tg_locations.test_tg_locations"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMTransitGatewaysLocationsDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resName, "locations.0.name"),
					resource.TestCheckResourceAttrSet(resName, "locations.0.type"),
					resource.TestCheckResourceAttrSet(resName, "locations.0.billing_location"),
				),
			},
		},
	})
}

func testAccCheckIBMTransitGatewaysLocationsDataSourceConfig() string {
	// status filter defaults to empty
	return `
	data "ibm_tg_locations" "test_tg_locations" {
		}   `
}
