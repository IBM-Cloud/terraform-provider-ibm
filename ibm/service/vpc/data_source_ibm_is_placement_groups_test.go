// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"fmt"
	"testing"

	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIbmIsPlacementGroupsDataSourceAllArgs(t *testing.T) {
	placementGroupStrategy := "host_spread"
	placementGroupName := fmt.Sprintf("tf-pg-name%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIbmIsPlacementGroupsDataSourceConfig(placementGroupStrategy, placementGroupName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_placement_groups.is_placement_groups", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_placement_groups.is_placement_groups", "placement_groups.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_placement_groups.is_placement_groups", "placement_groups.0.created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_is_placement_groups.is_placement_groups", "placement_groups.0.crn"),
					resource.TestCheckResourceAttrSet("data.ibm_is_placement_groups.is_placement_groups", "placement_groups.0.href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_placement_groups.is_placement_groups", "placement_groups.0.id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_placement_groups.is_placement_groups", "placement_groups.0.lifecycle_state"),
					resource.TestCheckResourceAttrSet("data.ibm_is_placement_groups.is_placement_groups", "placement_groups.0.name"),
					resource.TestCheckResourceAttrSet("data.ibm_is_placement_groups.is_placement_groups", "placement_groups.0.resource_type"),
					resource.TestCheckResourceAttrSet("data.ibm_is_placement_groups.is_placement_groups", "placement_groups.0.strategy"),
					resource.TestCheckResourceAttrSet("data.ibm_is_placement_groups.is_placement_groups", "total_count"),
				),
			},
		},
	})
}

func testAccCheckIbmIsPlacementGroupsDataSourceConfig(placementGroupStrategy string, placementGroupName string) string {
	return fmt.Sprintf(`
		data "ibm_resource_group" "default" {
			is_default=true
		}
		resource "ibm_is_placement_group" "is_placement_group" {
			strategy = "%s"
			name = "%s"
			resource_group = data.ibm_resource_group.default.id
		}

		data "ibm_is_placement_groups" "is_placement_groups" {
			depends_on = [
				ibm_is_placement_group.is_placement_group
			]
		}
	`, placementGroupStrategy, placementGroupName)
}
