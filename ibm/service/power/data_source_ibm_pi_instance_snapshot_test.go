// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package power_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMPIInstanceSnapshotDataSource_basic(t *testing.T) {
	snapshotResData := "data.ibm_pi_instance_snapshot.testacc_ds_snapshot"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPIInstanceSnapshotDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(snapshotResData, "id"),
				),
			},
		},
	})
}

func testAccCheckIBMPIInstanceSnapshotDataSourceConfig() string {
	return fmt.Sprintf(`
		data "ibm_pi_instance_snapshot" "testacc_ds_snapshot" {
			pi_cloud_instance_id = "%s"
			pi_snapshot_id = "%s"
		}`, acc.Pi_cloud_instance_id, acc.Pi_snapshot_id)
}
