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

func TestAccIBMDLOfferingSpeedsDataSource_basic(t *testing.T) {
	node1 := "data.ibm_dl_offering_speeds.test1"
	node2 := "data.ibm_dl_offering_speeds.test2"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMDLOfferingSpeedsDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(node1, "offering_speeds.0.link_speed"),
					resource.TestCheckResourceAttrSet(node2, "offering_speeds.0.link_speed"),
				),
			},
		},
	})
}

func testAccCheckIBMDLOfferingSpeedsDataSourceConfig() string {
	return fmt.Sprintf(`
	data "ibm_dl_offering_speeds" "test1" {
		offering_type = "dedicated"
	}

	data "ibm_dl_offering_speeds" "test2" {
		offering_type = "connect"
  	}
	`)
}
