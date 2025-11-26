// Copyright IBM Corp. 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package power_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccIBMPISPPPlacementGroupsDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPISPPPlacementGroupsDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_pi_spp_placement_groups.test", "id"),
				),
			},
		},
	})
}

func testAccCheckIBMPISPPPlacementGroupsDataSourceConfig() string {
	return fmt.Sprintf(`
		data "ibm_pi_spp_placement_groups" "test" {
			pi_cloud_instance_id = "%s"
		}
	`, acc.Pi_cloud_instance_id)
}
