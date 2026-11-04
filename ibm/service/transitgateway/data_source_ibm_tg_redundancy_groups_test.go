// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package transitgateway_test

import (
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMTransitGatewayRedundancyGroupsDataSource_basic(t *testing.T) {
	resName := "data.ibm_tg_redundancy_groups.all"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				// List all redundancy groups
				Config: testAccCheckIBMTransitGatewayRedundancyGroupsDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resName, "redundancy_groups.#"),
					resource.TestCheckResourceAttrSet(resName, "redundancy_groups.0.id"),
					resource.TestCheckResourceAttrSet(resName, "redundancy_groups.0.name"),
					resource.TestCheckResourceAttrSet(resName, "redundancy_groups.0.created_at"),
				),
			},
			{
				// Filter by name
				Config: testAccCheckIBMTransitGatewayRedundancyGroupsDataSourceFilterConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"data.ibm_tg_redundancy_groups.filtered", "redundancy_groups.#", "1"),
					resource.TestCheckResourceAttr(
						"data.ibm_tg_redundancy_groups.filtered", "redundancy_groups.0.name", "rg3-test"),
					resource.TestCheckResourceAttrSet(
						"data.ibm_tg_redundancy_groups.filtered", "redundancy_groups.0.id"),
				),
			},
		},
	})
}

func testAccCheckIBMTransitGatewayRedundancyGroupsDataSourceConfig() string {
	return `
data "ibm_tg_redundancy_groups" "all" {
}`
}

func testAccCheckIBMTransitGatewayRedundancyGroupsDataSourceFilterConfig() string {
	return `
data "ibm_tg_redundancy_groups" "filtered" {
  name = "rg3-test"
}`
}
