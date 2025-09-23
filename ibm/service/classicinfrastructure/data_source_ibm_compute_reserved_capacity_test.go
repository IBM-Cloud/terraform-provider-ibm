// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package classicinfrastructure_test

import (
	"fmt"
	"testing"

	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMComputeReservedCapacityDataSource_Basic(t *testing.T) {

	group1 := fmt.Sprintf("%s%s", "tfuatreservedcapacity", acctest.RandString(10))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		//CheckDestroy: testAccCheckIBMComputeReservedCapacityDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMComputeReservedCapacitydsConfig(group1),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_compute_reserved_capacity.reservedcapacity", "name", group1),
					resource.TestCheckResourceAttr(
						"ibm_compute_reserved_capacity.reservedcapacity", "flavor", "B1_2X4_1_YEAR_TERM"),
					resource.TestCheckResourceAttr(
						"data.ibm_compute_reserved_capacity.reservedcapacityds", "name", group1),
					resource.TestCheckResourceAttr(
						"data.ibm_compute_reserved_capacity.reservedcapacityds", "flavor", "B1_2X4_1_YEAR_TERM"),
					resource.TestCheckResourceAttr(
						"data.ibm_compute_reserved_capacity.reservedcapacityds", "datacenter", "dal05"),
					resource.TestCheckResourceAttr(
						"data.ibm_compute_reserved_capacity.reservedcapacityds", "pod", "pod01"),
					resource.TestCheckResourceAttr(
						"data.ibm_compute_reserved_capacity.reservedcapacityds", "instances", "1"),
				),
			},
		},
	})
}

func testAccCheckIBMComputeReservedCapacitydsConfig(name string) string {
	return fmt.Sprintf(`
resource "ibm_compute_reserved_capacity" "reservedcapacity" {
    name = "%s"
	datacenter = "dal05"
	pod = "pod01"
	flavor = "B1_2X4_1_YEAR_TERM"
	instances = "1"
}

data "ibm_compute_reserved_capacity" "reservedcapacityds" {
	depends_on = [ibm_compute_reserved_capacity.reservedcapacity]
    name = "${ibm_compute_reserved_capacity.reservedcapacity.name}"
}
`, name)
}
