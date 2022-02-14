// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"fmt"
	"strings"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMIsBackupPolicyPlanDataSourceBasic(t *testing.T) {
	vpcname := fmt.Sprintf("tf-vpc-%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf-instnace-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tf-subnet-%d", acctest.RandIntRange(10, 100))
	publicKey := strings.TrimSpace(`
ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCKVmnMOlHKcZK8tpt3MP1lqOLAcqcJzhsvJcjscgVERRN7/9484SOBJ3HSKxxNG5JN8owAjy5f9yYwcUg+JaUVuytn5Pv3aeYROHGGg+5G346xaq3DAwX6Y5ykr2fvjObgncQBnuU5KHWCECO/4h8uWuwh/kfniXPVjFToc+gnkqA+3RKpAecZhFXwfalQ9mMuYGFxn+fwn8cYEApsJbsEmb0iJwPiZ5hjFC8wREuiTlhPHDgkBLOiycd20op2nXzDbHfCHInquEe/gYxEitALONxm0swBOwJZwlTDOB7C6y2dzlrtxr1L59m7pCkWI4EtTRLvleehBoj3u7jB4usR
`)
	sshname := fmt.Sprintf("tf-ssh-%d", acctest.RandIntRange(10, 100))
	volname := fmt.Sprintf("tf-vol-%d", acctest.RandIntRange(10, 100))
	bakupPolicyName := fmt.Sprintf("tfbakuppolicyname%d", acctest.RandIntRange(10, 100))
	bakupPolicyPlanName := fmt.Sprintf("tfbakuppolicyplanname%d", acctest.RandIntRange(10, 100))
	cronSpec := fmt.Sprintf("tf_cron_spec_%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsBackupPolicyPlanDataSourceConfigBasic(bakupPolicyName, vpcname, subnetname, sshname, publicKey, volname, name, cronSpec, bakupPolicyPlanName),
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

func testAccCheckIBMIsBackupPolicyPlanDataSourceConfigBasic(backupPolicyName, vpcname, subnetname, sshname, publicKey, volName, name, bakupPolicyPlanName, cronSpec string) string {

	return testAccCheckIBMIsBackupPolicyPlanConfigBasic(backupPolicyName, vpcname, subnetname, sshname, publicKey, volName, name, cronSpec, bakupPolicyPlanName) + fmt.Sprintf(`
		data "ibm_is_backup_policy_plan" "is_backup_policy_plan" {
			backup_policy_id = ibm_is_backup_policy.is_backup_policy.id
			identifier = ibm_is_backup_policy_plan.is_backup_policy_plan.backup_policy_plan_id
		}
	`)
}

// func testAccCheckIBMIsBackupPolicyPlanNameDataSourceConfigBasic() string {
// 	return fmt.Sprintf(`
// 		data "ibm_is_backup_policy_plan" "is_backup_policy_plan" {
// 			backup_policy_id = "r134-59921a4a-bd6e-4b0c-88e4-2dff44c9142f"
// 			name = "my-backup-policy"
// 		}
// 	`)
// }
