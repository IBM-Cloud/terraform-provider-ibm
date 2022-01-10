// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMPIPlacementGroupDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPIPlacementGroupDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_pi_placement_group.testacc_ds_placement_group", "id"),
				),
			},
		},
	})
}

func testAccCheckIBMPIPlacementGroupDataSourceConfig() string {
	return fmt.Sprintf(`
data "ibm_pi_placement_group" "testacc_ds_placement_group" {
	pi_placement_group_name = "%s"
    pi_cloud_instance_id = "%s"
}`, pi_placement_group_name, pi_cloud_instance_id)

}
