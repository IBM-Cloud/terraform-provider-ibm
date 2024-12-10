// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package power_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccIBMPIInstanceSnapshotsDataSource_basic(t *testing.T) {
	snapshotResData := "data.ibm_pi_instance_snapshots.testacc_ds_snapshots"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPIInstanceSnapshotsDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(snapshotResData, "id"),
				),
			},
		},
	})
}

func testAccCheckIBMPIInstanceSnapshotsDataSourceConfig() string {
	return fmt.Sprintf(`
		data "ibm_pi_instance_snapshots" "testacc_ds_snapshots" {
			pi_cloud_instance_id = "%s"
		}`, acc.Pi_cloud_instance_id)
}
