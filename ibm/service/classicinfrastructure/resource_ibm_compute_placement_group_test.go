// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package classicinfrastructure_test

import (
	"fmt"
	"strconv"
	"strings"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/softlayer/softlayer-go/datatypes"
	"github.com/softlayer/softlayer-go/services"
)

func TestAccIBMComputePlacementGroup_Basic(t *testing.T) {
	var group datatypes.Virtual_PlacementGroup

	group1 := fmt.Sprintf("%s%s", "tfuatpgrp", acctest.RandString(10))
	group2 := fmt.Sprintf("%s%s", "tfuatpgrp", acctest.RandString(10))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMComputePlacementGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMComputePlacementGroupConfig(group1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMComputePlacementGroupExists("ibm_compute_placement_group.placementGroup", &group),
					resource.TestCheckResourceAttr(
						"ibm_compute_placement_group.placementGroup", "name", group1),
					resource.TestCheckResourceAttr(
						"ibm_compute_placement_group.placementGroup", "rule", "SPREAD"),
					resource.TestCheckResourceAttr(
						"ibm_compute_placement_group.placementGroup", "datacenter", "dal05"),
					resource.TestCheckResourceAttr(
						"ibm_compute_placement_group.placementGroup", "pod", "pod01"),
				),
			},

			{
				Config: testAccCheckIBMComputePlacementGroupUpdate(group2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMComputePlacementGroupExists("ibm_compute_placement_group.placementGroup", &group),
					resource.TestCheckResourceAttr(
						"ibm_compute_placement_group.placementGroup", "name", group2),
					resource.TestCheckResourceAttr(
						"ibm_compute_placement_group.placementGroup", "rule", "SPREAD"),
					resource.TestCheckResourceAttr(
						"ibm_compute_placement_group.placementGroup", "datacenter", "dal05"),
					resource.TestCheckResourceAttr(
						"ibm_compute_placement_group.placementGroup", "pod", "pod01"),
				),
			},
		},
	})
}

func TestAccIBMComputePlacementGroupWithTag(t *testing.T) {
	var group datatypes.Virtual_PlacementGroup

	group1 := fmt.Sprintf("%s%s", "tfuatpgrp", acctest.RandString(10))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMComputePlacementGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMComputePlacementGroupWithTag(group1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMComputePlacementGroupExists("ibm_compute_placement_group.placementGroup", &group),
					resource.TestCheckResourceAttr(
						"ibm_compute_placement_group.placementGroup", "name", group1),
					resource.TestCheckResourceAttr(
						"ibm_compute_placement_group.placementGroup", "tags.#", "2"),
					resource.TestCheckResourceAttr(
						"ibm_compute_placement_group.placementGroup", "datacenter", "lon02"),
					resource.TestCheckResourceAttr(
						"ibm_compute_placement_group.placementGroup", "pod", "pod01"),
				),
			},

			{
				Config: testAccCheckIBMComputePlacementGroupWithUpdatedTag(group1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMComputePlacementGroupExists("ibm_compute_placement_group.placementGroup", &group),
					resource.TestCheckResourceAttr(
						"ibm_compute_placement_group.placementGroup", "name", group1),
					resource.TestCheckResourceAttr(
						"ibm_compute_placement_group.placementGroup", "tags.#", "3"),
					resource.TestCheckResourceAttr(
						"ibm_compute_placement_group.placementGroup", "datacenter", "lon02"),
					resource.TestCheckResourceAttr(
						"ibm_compute_placement_group.placementGroup", "pod", "pod01"),
				),
			},
		},
	})
}

func TestAccIBMComputePlacementGroupImport(t *testing.T) {
	var group datatypes.Virtual_PlacementGroup

	group1 := fmt.Sprintf("%s%s", "tfuatpgrp", acctest.RandString(10))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMComputePlacementGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMComputePlacementGroupConfig(group1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMComputePlacementGroupExists("ibm_compute_placement_group.placementGroup", &group),
					resource.TestCheckResourceAttr(
						"ibm_compute_placement_group.placementGroup", "name", group1),
					resource.TestCheckResourceAttr(
						"ibm_compute_placement_group.placementGroup", "datacenter", "dal05"),
					resource.TestCheckResourceAttr(
						"ibm_compute_placement_group.placementGroup", "pod", "pod01"),
				),
			},

			{
				ResourceName:      "ibm_compute_placement_group.placementGroup",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIBMComputePlacementGroupDestroy(s *terraform.State) error {
	service := services.GetVirtualPlacementGroupService(acc.TestAccProvider.Meta().(conns.ClientSession).SoftLayerSession())

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_compute_placement_group" {
			continue
		}

		pgrpId, _ := strconv.Atoi(rs.Primary.ID)

		// Try to find the provisioning pgrp
		_, err := service.Id(pgrpId).GetObject()

		if err == nil {
			return fmt.Errorf("Placement group still exists: %s", rs.Primary.ID)
		} else if !strings.Contains(err.Error(), "404") {
			return fmt.Errorf("[ERROR] Error waiting for placement group (%s) to be destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}

func testAccCheckIBMComputePlacementGroupExists(n string, pgrp *datatypes.Virtual_PlacementGroup) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("[ERROR] Not  found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("[ERROR] No Record ID is set")
		}

		pgrpId, _ := strconv.Atoi(rs.Primary.ID)

		service := services.GetVirtualPlacementGroupService(acc.TestAccProvider.Meta().(conns.ClientSession).SoftLayerSession())
		foundpgrp, err := service.Id(pgrpId).GetObject()

		if err != nil {
			return err
		}

		if strconv.Itoa(int(*foundpgrp.Id)) != rs.Primary.ID {
			return fmt.Errorf("Record not found")
		}

		*pgrp = foundpgrp

		return nil
	}
}

func testAccCheckIBMComputePlacementGroupConfig(name string) string {
	return fmt.Sprintf(`
resource "ibm_compute_placement_group" "placementGroup" {
    name = "%s"
	datacenter = "dal05"
	pod = "pod01"
}`, name)
}

func testAccCheckIBMComputePlacementGroupUpdate(name string) string {
	return fmt.Sprintf(`
resource "ibm_compute_placement_group" "placementGroup" {
    name = "%s"
	datacenter = "dal05"
	pod = "pod01"
}`, name)
}

func testAccCheckIBMComputePlacementGroupWithTag(name string) string {
	return fmt.Sprintf(`
resource "ibm_compute_placement_group" "placementGroup" {
    name = "%s"
	datacenter = "lon02"
	pod = "pod01"
	tags = ["one", "two"]
}`, name)
}

func testAccCheckIBMComputePlacementGroupWithUpdatedTag(name string) string {
	return fmt.Sprintf(`
resource "ibm_compute_placement_group" "placementGroup" {
    name = "%s"
	datacenter = "lon02"
	pod = "pod01"
	tags = ["one", "two", "three"]
}`, name)
}
