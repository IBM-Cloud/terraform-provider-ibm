// Copyright IBM Corp. 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package power_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccIBMPISPPPlacementGroupDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPISPPPlacementGroupDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_pi_spp_placement_group.testacc_ds_spp_placement_group", "id"),
				),
			},
		},
	})
}

func testAccCheckIBMPISPPPlacementGroupDataSourceConfig() string {
	return fmt.Sprintf(`
		data "ibm_pi_spp_placement_group" "testacc_ds_spp_placement_group" {
			pi_spp_placement_group_id = "%s"
			pi_cloud_instance_id = "%s"
		}`, acc.Pi_spp_placement_group_id, acc.Pi_cloud_instance_id)
}
