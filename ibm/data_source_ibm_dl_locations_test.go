// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"testing"
)

func TestAccIBMDLLocationsDataSource_basic(t *testing.T) {
	node := "data.ibm_dl_locations.test_dl_locations"
	offeringType := "dedicated"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMDLOfferingLocationsDataSourceConfig(offeringType),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(node, "locations.#"),
				),
			},
		},
	})
}

func testAccCheckIBMDLOfferingLocationsDataSourceConfig(offeringType string) string {
	return fmt.Sprintf(`
	   data "ibm_dl_locations" "test_dl_locations"{
		offering_type = "%s"
	 }

	  `, offeringType)
}
