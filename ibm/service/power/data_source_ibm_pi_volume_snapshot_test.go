// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package power_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccIBMPIVolumeSnapshotDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPIVolumeSnapshotDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_pi_volume_snapshot.snapshot_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_pi_volume_snapshot.snapshot_instance", "name"),
				),
			},
		},
	})
}

func testAccCheckIBMPIVolumeSnapshotDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		data "ibm_pi_volume_snapshot" "snapshot_instance" {
			pi_cloud_instance_id = "%s"
			pi_volume_snapshot_id = "%s"
		}
	`, acc.Pi_cloud_instance_id, acc.Pi_snapshot_id)
}
