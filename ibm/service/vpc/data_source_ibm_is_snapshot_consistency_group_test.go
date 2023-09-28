// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIBMIsSnapshotConsistencyGroupDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsSnapshotConsistencyGroupDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_snapshot_consistency_group.is_snapshot_consistency_group", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_snapshot_consistency_group.is_snapshot_consistency_group", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_is_snapshot_consistency_group.is_snapshot_consistency_group", "crn"),
					resource.TestCheckResourceAttrSet("data.ibm_is_snapshot_consistency_group.is_snapshot_consistency_group", "delete_snapshots_on_delete"),
					resource.TestCheckResourceAttrSet("data.ibm_is_snapshot_consistency_group.is_snapshot_consistency_group", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_snapshot_consistency_group.is_snapshot_consistency_group", "lifecycle_state"),
					resource.TestCheckResourceAttrSet("data.ibm_is_snapshot_consistency_group.is_snapshot_consistency_group", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_is_snapshot_consistency_group.is_snapshot_consistency_group", "resource_group.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_snapshot_consistency_group.is_snapshot_consistency_group", "resource_type"),
					resource.TestCheckResourceAttrSet("data.ibm_is_snapshot_consistency_group.is_snapshot_consistency_group", "snapshots.#"),
				),
			},
		},
	})
}

func testAccCheckIBMIsSnapshotConsistencyGroupDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		data "ibm_is_snapshot_consistency_group" "is_snapshot_consistency_group" {
			identifier = "r134-628982dd-bb79-4f7c-ac64-d1b70b1064e8"
	  	}
	`)
}
