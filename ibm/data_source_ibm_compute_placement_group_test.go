package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
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
