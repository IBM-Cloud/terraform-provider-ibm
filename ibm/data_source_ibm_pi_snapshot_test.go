// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMPISnapshotDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPISnapshotDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_pi_pvm_snapshots.testacc_pi_snapshots", "id"),
				),
			},
		},
	})
}

func testAccCheckIBMPISnapshotDataSourceConfig() string {
	return fmt.Sprintf(`
	
data "ibm_pi_pvm_snapshots" "testacc_pi_snapshots" {
    pi_instance_name = "%s"
    pi_cloud_instance_id = "%s"
}`, pi_instance_name, pi_cloud_instance_id)

}
