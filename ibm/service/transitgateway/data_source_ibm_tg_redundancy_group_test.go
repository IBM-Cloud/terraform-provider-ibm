// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package transitgateway_test

import (
	"regexp"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMTransitGatewayRedundancyGroupDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMTransitGatewayRedundancyGroupDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(
						"data.ibm_tg_redundancy_group.test_rg", "id"),
					resource.TestCheckResourceAttr(
						"data.ibm_tg_redundancy_group.test_rg", "name", "rg3-test"),
					resource.TestCheckResourceAttrSet(
						"data.ibm_tg_redundancy_group.test_rg", "created_at"),
				),
			},
		},
	})
}

// TestAccIBMTransitGatewayRedundancyGroupDataSource_notFound verifies that looking up
// a redundancy group by a name that does not exist returns a descriptive error.
func TestAccIBMTransitGatewayRedundancyGroupDataSource_notFound(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config:      testAccCheckIBMTransitGatewayRedundancyGroupDataSourceNotFoundConfig(),
				ExpectError: regexp.MustCompile(`Couldn't find any redundancy group with the specified name`),
			},
		},
	})
}

func testAccCheckIBMTransitGatewayRedundancyGroupDataSourceConfig() string {
	return `
data "ibm_tg_redundancy_group" "test_rg" {
  name = "rg3-test"
}`
}

func testAccCheckIBMTransitGatewayRedundancyGroupDataSourceNotFoundConfig() string {
	return `
data "ibm_tg_redundancy_group" "test_rg_missing" {
  name = "this-redundancy-group-does-not-exist"
}`
}
