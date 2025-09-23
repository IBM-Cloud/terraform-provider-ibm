// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package classicinfrastructure_test

import (
	"fmt"
	"strconv"
	"strings"
	"testing"

	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/conns"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/softlayer/softlayer-go/datatypes"
	"github.com/softlayer/softlayer-go/services"
)

func TestAccIBMComputeReservedCapacity_Basic(t *testing.T) {
	var group datatypes.Virtual_ReservedCapacityGroup

	group1 := fmt.Sprintf("%s%s", "tfuatreservedcapacity", acctest.RandString(10))
	group2 := fmt.Sprintf("%s%s", "tfuatreservedcapacity", acctest.RandString(10))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		//CheckDestroy: testAccCheckIBMComputeReservedCapacityDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMComputeReservedCapacityConfig(group1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMComputeReservedCapacityExists("ibm_compute_reserved_capacity.reservedCapacity", &group),
					resource.TestCheckResourceAttr(
						"ibm_compute_reserved_capacity.reservedCapacity", "name", group1),
					resource.TestCheckResourceAttr(
						"ibm_compute_reserved_capacity.reservedCapacity", "flavor", "B1_2X4_1_YEAR_TERM"),
					resource.TestCheckResourceAttr(
						"ibm_compute_reserved_capacity.reservedCapacity", "datacenter", "lon02"),
					resource.TestCheckResourceAttr(
						"ibm_compute_reserved_capacity.reservedCapacity", "pod", "pod01"),
				),
			},

			{
				Config: testAccCheckIBMComputeReservedCapacityUpdate(group2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMComputeReservedCapacityExists("ibm_compute_reserved_capacity.reservedCapacity", &group),
					resource.TestCheckResourceAttr(
						"ibm_compute_reserved_capacity.reservedCapacity", "name", group2),
					resource.TestCheckResourceAttr(
						"ibm_compute_reserved_capacity.reservedCapacity", "flavor", "B1_2X4_1_YEAR_TERM"),
					resource.TestCheckResourceAttr(
						"ibm_compute_reserved_capacity.reservedCapacity", "datacenter", "lon02"),
					resource.TestCheckResourceAttr(
						"ibm_compute_reserved_capacity.reservedCapacity", "pod", "pod01"),
				),
			},

			{
				ResourceName:      "ibm_compute_reserved_capacity.reservedCapacity",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIBMComputeReservedCapacityDestroy(s *terraform.State) error {
	service := services.GetVirtualReservedCapacityGroupService(acc.TestAccProvider.Meta().(conns.ClientSession).SoftLayerSession())

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_compute_reserved_capacity" {
			continue
		}

		reservedcapacityId, _ := strconv.Atoi(rs.Primary.ID)

		// Try to find the provisioning reservedcapacity
		_, err := service.Id(reservedcapacityId).GetObject()

		if err == nil {
			return fmt.Errorf("Reserved Capacity still exists: %s", rs.Primary.ID)
		} else if !strings.Contains(err.Error(), "404") {
			return fmt.Errorf("[ERROR] Error waiting for reserved capacity (%s) to be destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}

func testAccCheckIBMComputeReservedCapacityExists(n string, reservedcapacity *datatypes.Virtual_ReservedCapacityGroup) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("[ERROR] Not  found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("[ERROR] No Record ID is set")
		}

		reservedcapacityId, _ := strconv.Atoi(rs.Primary.ID)

		service := services.GetVirtualReservedCapacityGroupService(acc.TestAccProvider.Meta().(conns.ClientSession).SoftLayerSession())
		foundreservedcapacity, err := service.Id(reservedcapacityId).GetObject()

		if err != nil {
			return err
		}

		if strconv.Itoa(int(*foundreservedcapacity.Id)) != rs.Primary.ID {
			return fmt.Errorf("Record not found")
		}

		*reservedcapacity = foundreservedcapacity

		return nil
	}
}

func testAccCheckIBMComputeReservedCapacityConfig(name string) string {
	return fmt.Sprintf(`
resource "ibm_compute_reserved_capacity" "reservedCapacity" {
	datacenter = "lon02"
    pod = "pod01"
    instances = 6
    name = "%s"
    flavor = "B1_2X4_1_YEAR_TERM"
}`, name)
}

func testAccCheckIBMComputeReservedCapacityUpdate(name string) string {
	return fmt.Sprintf(`
resource "ibm_compute_reserved_capacity" "reservedCapacity" {
    datacenter = "lon02"
    pod = "pod01"
    instances = 6
    name = "%s"
    flavor = "B1_2X4_1_YEAR_TERM"
}`, name)
}
