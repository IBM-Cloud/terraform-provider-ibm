// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMIsBackupPolicyPlanDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsBackupPolicyPlanDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_backup_policy_plan.is_backup_policy_plan", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_backup_policy_plan.is_backup_policy_plan", "backup_policy_id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_backup_policy_plan.is_backup_policy_plan", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_backup_policy_plan.is_backup_policy_plan", "active"),
					// resource.TestCheckResourceAttrSet("data.ibm_is_backup_policy_plan.is_backup_policy_plan", "attach_user_tags.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_backup_policy_plan.is_backup_policy_plan", "copy_user_tags"),
					// resource.TestCheckResourceAttrSet("data.ibm_is_backup_policy_plan.is_backup_policy_plan", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_is_backup_policy_plan.is_backup_policy_plan", "cron_spec"),
					resource.TestCheckResourceAttrSet("data.ibm_is_backup_policy_plan.is_backup_policy_plan", "deletion_trigger.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_backup_policy_plan.is_backup_policy_plan", "href"),
					// resource.TestCheckResourceAttrSet("data.ibm_is_backup_policy_plan.is_backup_policy_plan", "lifecycle_state"),
					resource.TestCheckResourceAttrSet("data.ibm_is_backup_policy_plan.is_backup_policy_plan", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_is_backup_policy_plan.is_backup_policy_plan", "resource_type"),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMIsBackupPolicyPlanNameDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_backup_policy_plan.is_backup_policy_plan", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_backup_policy_plan.is_backup_policy_plan", "backup_policy_id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_backup_policy_plan.is_backup_policy_plan", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_backup_policy_plan.is_backup_policy_plan", "active"),
					// resource.TestCheckResourceAttrSet("data.ibm_is_backup_policy_plan.is_backup_policy_plan", "attach_user_tags.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_backup_policy_plan.is_backup_policy_plan", "copy_user_tags"),
					// resource.TestCheckResourceAttrSet("data.ibm_is_backup_policy_plan.is_backup_policy_plan", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_is_backup_policy_plan.is_backup_policy_plan", "cron_spec"),
					resource.TestCheckResourceAttrSet("data.ibm_is_backup_policy_plan.is_backup_policy_plan", "deletion_trigger.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_backup_policy_plan.is_backup_policy_plan", "href"),
					// resource.TestCheckResourceAttrSet("data.ibm_is_backup_policy_plan.is_backup_policy_plan", "lifecycle_state"),
					resource.TestCheckResourceAttrSet("data.ibm_is_backup_policy_plan.is_backup_policy_plan", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_is_backup_policy_plan.is_backup_policy_plan", "resource_type"),
				),
			},
		},
	})
}

func testAccCheckIBMIsBackupPolicyPlanDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		data "ibm_is_backup_policy_plan" "is_backup_policy_plan" {
			backup_policy_id = "r134-59921a4a-bd6e-4b0c-88e4-2dff44c9142f"
			identifier = "r134-edf26f05-5956-4f9c-9672-c9fc685c9b12"
		}
	`)
}

func testAccCheckIBMIsBackupPolicyPlanNameDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		data "ibm_is_backup_policy_plan" "is_backup_policy_plan" {
			backup_policy_id = "r134-59921a4a-bd6e-4b0c-88e4-2dff44c9142f"
			name = "my-backup-policy"
		}
	`)
}
