// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMIsBackupPolicyJobDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsBackupPolicyJobDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_backup_policy_job.is_backup_policy_job", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_backup_policy_job.is_backup_policy_job", "backup_policy_id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_backup_policy_job.is_backup_policy_job", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_backup_policy_job.is_backup_policy_job", "auto_delete"),
					resource.TestCheckResourceAttrSet("data.ibm_is_backup_policy_job.is_backup_policy_job", "auto_delete_after"),
					resource.TestCheckResourceAttrSet("data.ibm_is_backup_policy_job.is_backup_policy_job", "backup_policy_plan.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_backup_policy_job.is_backup_policy_job", "completed_at"),
					resource.TestCheckResourceAttrSet("data.ibm_is_backup_policy_job.is_backup_policy_job", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_is_backup_policy_job.is_backup_policy_job", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_backup_policy_job.is_backup_policy_job", "job_type"),
					resource.TestCheckResourceAttrSet("data.ibm_is_backup_policy_job.is_backup_policy_job", "resource_type"),
					resource.TestCheckResourceAttrSet("data.ibm_is_backup_policy_job.is_backup_policy_job", "source_volume.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_backup_policy_job.is_backup_policy_job", "status"),
					resource.TestCheckResourceAttrSet("data.ibm_is_backup_policy_job.is_backup_policy_job", "status_reasons.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_backup_policy_job.is_backup_policy_job", "target_snapshot.#"),
				),
			},
		},
	})
}

func testAccCheckIBMIsBackupPolicyJobDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		data "ibm_is_backup_policy_job" "is_backup_policy_job" {
			backup_policy_id = "backup_policy_id"
			identifier = "id"
		}
	`)
}
