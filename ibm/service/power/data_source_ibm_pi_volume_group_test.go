// Copyright IBM Corp. 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package power_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccIBMPIVolumeGroupDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPIVolumeGroupDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_pi_volume_group.testacc_ds_volume_group", "id"),
				),
			},
		},
	})
}

func testAccCheckIBMPIVolumeGroupDataSourceConfig() string {
	return fmt.Sprintf(`
		data "ibm_pi_volume_group" "testacc_ds_volume_group" {
			pi_volume_group_id   = "%s"
			pi_cloud_instance_id = "%s"
		}`, acc.Pi_volume_group_id, acc.Pi_cloud_instance_id)
}
