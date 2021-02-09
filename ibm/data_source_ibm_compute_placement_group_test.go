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

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccIBMComputePlacementGroupDataSource_Basic(t *testing.T) {

	group1 := fmt.Sprintf("%s%s", "tfuatpgrp", acctest.RandString(10))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMComputePlacementGroupDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMComputePlacementGroupdsConfig(group1),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_compute_placement_group.placementGroup", "name", group1),
					resource.TestCheckResourceAttr(
						"ibm_compute_placement_group.placementGroup", "rule", "SPREAD"),
					resource.TestCheckResourceAttr(
						"data.ibm_compute_placement_group.placementGroupds", "name", group1),
					resource.TestCheckResourceAttr(
						"data.ibm_compute_placement_group.placementGroupds", "rule", "SPREAD"),
					resource.TestCheckResourceAttr(
						"data.ibm_compute_placement_group.placementGroupds", "datacenter", "dal05"),
					resource.TestCheckResourceAttr(
						"data.ibm_compute_placement_group.placementGroupds", "pod", "pod01"),
				),
			},
		},
	})
}

func testAccCheckIBMComputePlacementGroupdsConfig(name string) string {
	return fmt.Sprintf(`
resource "ibm_compute_placement_group" "placementGroup" {
    name = "%s"
	datacenter = "dal05"
	pod = "pod01"
}

data "ibm_compute_placement_group" "placementGroupds" {
    name = "${ibm_compute_placement_group.placementGroup.name}"
}
`, name)
}
