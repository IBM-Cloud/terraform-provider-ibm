// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
	"time"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMIsBackupPolicyPlansDataSourceBasic(t *testing.T) {
	vpcname := fmt.Sprintf("tf-vpc-%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf-instnace-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tf-subnet-%d", acctest.RandIntRange(10, 100))
	sshname := fmt.Sprintf("tf-ssh-%d", acctest.RandIntRange(10, 100))
	volname := fmt.Sprintf("tf-vol-%d", acctest.RandIntRange(10, 100))
	bakupPolicyName := fmt.Sprintf("tfbakuppolicyname%d", acctest.RandIntRange(10, 100))
	bakupPolicyPlanName := fmt.Sprintf("tfbakuppolicyplanname%d", acctest.RandIntRange(10, 100))
	cronSpec := strings.TrimSpace(strconv.Itoa(time.Now().UTC().Minute()) + " " + strconv.Itoa(time.Now().UTC().Hour()) + " " + "*" + " " + "*" + " " + "*")
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsBackupPolicyPlansDataSourceConfigBasic(bakupPolicyName, vpcname, subnetname, sshname, volname, name, cronSpec, bakupPolicyPlanName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_backup_policy_plans.is_backup_policy_plans", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_backup_policy_plans.is_backup_policy_plans", "backup_policy_id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_backup_policy_plans.is_backup_policy_plans", "plans.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_backup_policy_plans.is_backup_policy_plans", "plans.0.active"),
					resource.TestCheckResourceAttrSet("data.ibm_is_backup_policy_plans.is_backup_policy_plans", "plans.0.attach_user_tags.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_backup_policy_plans.is_backup_policy_plans", "plans.0.copy_user_tags"),
					resource.TestCheckResourceAttrSet("data.ibm_is_backup_policy_plans.is_backup_policy_plans", "plans.0.created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_is_backup_policy_plans.is_backup_policy_plans", "plans.0.cron_spec"),
					resource.TestCheckResourceAttrSet("data.ibm_is_backup_policy_plans.is_backup_policy_plans", "plans.0.deletion_trigger.0.delete_after"),
					resource.TestCheckResourceAttrSet("data.ibm_is_backup_policy_plans.is_backup_policy_plans", "plans.0.href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_backup_policy_plans.is_backup_policy_plans", "plans.0.id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_backup_policy_plans.is_backup_policy_plans", "plans.0.lifecycle_state"),
					resource.TestCheckResourceAttrSet("data.ibm_is_backup_policy_plans.is_backup_policy_plans", "plans.0.name"),
					resource.TestCheckResourceAttrSet("data.ibm_is_backup_policy_plans.is_backup_policy_plans", "plans.0.resource_type"),
				),
			},
		},
	})
}
func TestAccIBMIsBackupPolicyPlansDataSourceClonesBasic(t *testing.T) {
	vpcname := fmt.Sprintf("tf-vpc-%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf-instnace-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tf-subnet-%d", acctest.RandIntRange(10, 100))
	sshname := fmt.Sprintf("tf-ssh-%d", acctest.RandIntRange(10, 100))
	volname := fmt.Sprintf("tf-vol-%d", acctest.RandIntRange(10, 100))
	bakupPolicyName := fmt.Sprintf("tfbakuppolicyname%d", acctest.RandIntRange(10, 100))
	bakupPolicyPlanName := fmt.Sprintf("tfbakuppolicyplanname%d", acctest.RandIntRange(10, 100))
	cronSpec := strings.TrimSpace(strconv.Itoa(time.Now().UTC().Minute()) + " " + strconv.Itoa(time.Now().UTC().Hour()) + " " + "*" + " " + "*" + " " + "*")
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsBackupPolicyPlansDataSourceConfigClonesBasic(bakupPolicyName, vpcname, subnetname, sshname, volname, name, cronSpec, bakupPolicyPlanName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_backup_policy_plans.is_backup_policy_plans", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_backup_policy_plans.is_backup_policy_plans", "backup_policy_id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_backup_policy_plans.is_backup_policy_plans", "plans.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_backup_policy_plans.is_backup_policy_plans", "plans.0.active"),
					resource.TestCheckResourceAttrSet("data.ibm_is_backup_policy_plans.is_backup_policy_plans", "plans.0.attach_user_tags.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_backup_policy_plans.is_backup_policy_plans", "plans.0.copy_user_tags"),
					resource.TestCheckResourceAttrSet("data.ibm_is_backup_policy_plans.is_backup_policy_plans", "plans.0.created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_is_backup_policy_plans.is_backup_policy_plans", "plans.0.cron_spec"),
					resource.TestCheckResourceAttrSet("data.ibm_is_backup_policy_plans.is_backup_policy_plans", "plans.0.deletion_trigger.0.delete_after"),
					resource.TestCheckResourceAttrSet("data.ibm_is_backup_policy_plans.is_backup_policy_plans", "plans.0.href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_backup_policy_plans.is_backup_policy_plans", "plans.0.id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_backup_policy_plans.is_backup_policy_plans", "plans.0.lifecycle_state"),
					resource.TestCheckResourceAttrSet("data.ibm_is_backup_policy_plans.is_backup_policy_plans", "plans.0.name"),
					resource.TestCheckResourceAttrSet("data.ibm_is_backup_policy_plans.is_backup_policy_plans", "plans.0.resource_type"),
					resource.TestCheckResourceAttrSet("data.ibm_is_backup_policy_plans.is_backup_policy_plans", "plans.0.clone_policy.0.max_snapshots"),
					resource.TestCheckResourceAttrSet("data.ibm_is_backup_policy_plans.is_backup_policy_plans", "plans.0.clone_policy.0.zones.#"),
					resource.TestCheckResourceAttr("data.ibm_is_backup_policy_plans.is_backup_policy_plans", "plans.0.clone_policy.0.zones.0", acc.ISZoneName),
					resource.TestCheckResourceAttr("data.ibm_is_backup_policy_plans.is_backup_policy_plans", "plans.0.clone_policy.0.zones.1", acc.ISZoneName2),
				),
			},
		},
	})
}

func testAccCheckIBMIsBackupPolicyPlansDataSourceConfigBasic(backupPolicyName, vpcname, subnetname, sshname, volName, name, cronSpec, bakupPolicyPlanName string) string {

	return testAccCheckIBMIsBackupPolicyPlanConfigBasic(backupPolicyName, vpcname, subnetname, sshname, volName, name, cronSpec, bakupPolicyPlanName) + fmt.Sprintf(`
		data "ibm_is_backup_policy_plans" "is_backup_policy_plans" {
			depends_on  = [ibm_is_backup_policy_plan.is_backup_policy_plan]
			backup_policy_id = ibm_is_backup_policy.is_backup_policy.id
			name = "%s"
		}
	`, bakupPolicyPlanName+"-1")
}
func testAccCheckIBMIsBackupPolicyPlansDataSourceConfigClonesBasic(backupPolicyName, vpcname, subnetname, sshname, volName, name, cronSpec, bakupPolicyPlanName string) string {

	return testAccCheckIBMIsBackupPolicyPlanConfigClonesBasic(backupPolicyName, vpcname, subnetname, sshname, volName, name, cronSpec, bakupPolicyPlanName) + fmt.Sprintf(`
		data "ibm_is_backup_policy_plans" "is_backup_policy_plans" {
			depends_on  = [ibm_is_backup_policy_plan.is_backup_policy_plan]
			backup_policy_id = ibm_is_backup_policy.is_backup_policy.id
			name = "%s"
		}
	`, bakupPolicyPlanName+"-1")
}
