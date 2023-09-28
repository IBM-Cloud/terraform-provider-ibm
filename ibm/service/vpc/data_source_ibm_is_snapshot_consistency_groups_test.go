// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIBMIsSnapshotConsistencyGroupsDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsSnapshotConsistencyGroupsDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_snapshot_consistency_groups.is_snapshot_consistency_groups", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_snapshot_consistency_groups.is_snapshot_consistency_groups", "first.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_snapshot_consistency_groups.is_snapshot_consistency_groups", "limit"),
					resource.TestCheckResourceAttrSet("data.ibm_is_snapshot_consistency_groups.is_snapshot_consistency_groups", "snapshot_consistency_groups.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_snapshot_consistency_groups.is_snapshot_consistency_groups", "total_count"),
				),
			},
		},
	})
}

func testAccCheckIBMIsSnapshotConsistencyGroupsDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		data "ibm_is_snapshot_consistency_groups" "is_snapshot_consistency_groups_instance" {
			resource_group.id = "resource_group.id"
			name = "name"
			sort = "name"
			backup_policy_plan.id = "backup_policy_plan.id"
		}
	`)
}
