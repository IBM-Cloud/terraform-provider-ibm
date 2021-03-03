// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccIBMPISnapshotsDataSource_basic(t *testing.T) {
	t.Skip()
	//name := "Trial "
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMPISnapshotsDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ibm_pi_pvminstance_snapshots.testacc_ds_snapshots", "pi_cloud_instance_id", pi_cloud_instance_id),
				),
			},
		},
	})
}

func testAccCheckIBMPISnapshotsDataSourceConfig() string {
	return fmt.Sprintf(`
	
data "ibm_pi_instance_snapshots" "testacc_ds_snapshots" {
    pi_cloud_instance_id = "%s"
}`, pi_cloud_instance_id)

}
