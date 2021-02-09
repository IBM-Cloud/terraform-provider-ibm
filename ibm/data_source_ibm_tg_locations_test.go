/**
 * © Copyright IBM Corporation 2020. All Rights Reserved.
 *
 * Licensed under the Mozilla Public License, version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at https://mozilla.org/MPL/2.0/
 */

package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccIBMTransitGatewaysLocationsDataSource_basic(t *testing.T) {
	resName := "data.ibm_tg_locations.test_tg_locations"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
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
	return fmt.Sprintf(`
	data "ibm_tg_locations" "test_tg_locations" {
		}   `)
}
