// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package power_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMPIPlacementGroupDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
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
}`, acc.Pi_placement_group_name, acc.Pi_cloud_instance_id)

}
