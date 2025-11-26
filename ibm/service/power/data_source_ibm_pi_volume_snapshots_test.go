// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package power_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccIBMPIVolumeSnapshotsDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPIVolumeSnapshotsDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_pi_volume_snapshots.snapshots", "id"),
				),
			},
		},
	})
}

func testAccCheckIBMPIVolumeSnapshotsDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		data "ibm_pi_volume_snapshots" "snapshots" {
			pi_cloud_instance_id = "%s"
		}
	`, acc.Pi_cloud_instance_id)
}
