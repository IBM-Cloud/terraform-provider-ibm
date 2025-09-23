// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package directlink_test

import (
	"testing"

	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMDLOfferingSpeedsDataSource_basic(t *testing.T) {
	node1 := "data.ibm_dl_offering_speeds.test1"
	node2 := "data.ibm_dl_offering_speeds.test2"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMDLOfferingSpeedsDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(node1, "offering_speeds.0.capabilities.#"),
					resource.TestCheckResourceAttrSet(node1, "offering_speeds.0.link_speed"),
					resource.TestCheckResourceAttrSet(node2, "offering_speeds.0.capabilities.#"),
					resource.TestCheckResourceAttrSet(node2, "offering_speeds.0.link_speed"),
				),
			},
		},
	})
}

func testAccCheckIBMDLOfferingSpeedsDataSourceConfig() string {
	return `
	data "ibm_dl_offering_speeds" "test1" {
		offering_type = "dedicated"
	}

	data "ibm_dl_offering_speeds" "test2" {
		offering_type = "connect"
  	}
	`
}
