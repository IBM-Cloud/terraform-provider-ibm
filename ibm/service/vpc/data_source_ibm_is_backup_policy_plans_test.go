// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMIsBackupPolicyPlansDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsBackupPolicyPlansDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_backup_policy_plans.is_backup_policy_plans", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_backup_policy_plans.is_backup_policy_plans", "backup_policy_id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_backup_policy_plans.is_backup_policy_plans", "plans.#"),
				),
			},
		},
	})
}

func testAccCheckIBMIsBackupPolicyPlansDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		data "ibm_is_backup_policy_plans" "is_backup_policy_plans" {
			backup_policy_id = "r134-59921a4a-bd6e-4b0c-88e4-2dff44c9142f"
			name = "my-backup-policy"
		}
	`)
}
