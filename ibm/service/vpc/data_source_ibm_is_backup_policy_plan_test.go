// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
	"time"

	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMIsBackupPolicyPlanDataSourceBasic(t *testing.T) {
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
				Config: testAccCheckIBMIsBackupPolicyPlanDataSourceConfigBasic(bakupPolicyName, vpcname, subnetname, sshname, volname, name, cronSpec, bakupPolicyPlanName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_backup_policy_plan.is_backup_policy_plan", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_backup_policy_plan.is_backup_policy_plan", "backup_policy_id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_backup_policy_plan.is_backup_policy_plan", "active"),
					resource.TestCheckResourceAttrSet("data.ibm_is_backup_policy_plan.is_backup_policy_plan", "attach_user_tags.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_backup_policy_plan.is_backup_policy_plan", "copy_user_tags"),
					resource.TestCheckResourceAttrSet("data.ibm_is_backup_policy_plan.is_backup_policy_plan", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_is_backup_policy_plan.is_backup_policy_plan", "cron_spec"),
					resource.TestCheckResourceAttrSet("data.ibm_is_backup_policy_plan.is_backup_policy_plan", "deletion_trigger.0.delete_after"),
					resource.TestCheckResourceAttrSet("data.ibm_is_backup_policy_plan.is_backup_policy_plan", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_backup_policy_plan.is_backup_policy_plan", "lifecycle_state"),
					resource.TestCheckResourceAttrSet("data.ibm_is_backup_policy_plan.is_backup_policy_plan", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_is_backup_policy_plan.is_backup_policy_plan", "resource_type"),
					resource.TestCheckResourceAttrSet("data.ibm_is_backup_policy_plan.is_backup_policy_plan", "remote_region_policy.0.delete_over_count"),
					resource.TestCheckResourceAttrSet("data.ibm_is_backup_policy_plan.is_backup_policy_plan", "remote_region_policy.0.encryption_key"),
					resource.TestCheckResourceAttrSet("data.ibm_is_backup_policy_plan.is_backup_policy_plan", "remote_region_policy.0.region"),
				),
			},
		},
	})
}
func TestAccIBMIsBackupPolicyPlanDataSourceClonesBasic(t *testing.T) {
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
				Config: testAccCheckIBMIsBackupPolicyPlanDataSourceConfigClonesBasic(bakupPolicyName, vpcname, subnetname, sshname, volname, name, cronSpec, bakupPolicyPlanName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_backup_policy_plan.is_backup_policy_plan", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_backup_policy_plan.is_backup_policy_plan", "backup_policy_id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_backup_policy_plan.is_backup_policy_plan", "active"),
					resource.TestCheckResourceAttrSet("data.ibm_is_backup_policy_plan.is_backup_policy_plan", "attach_user_tags.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_backup_policy_plan.is_backup_policy_plan", "copy_user_tags"),
					resource.TestCheckResourceAttrSet("data.ibm_is_backup_policy_plan.is_backup_policy_plan", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_is_backup_policy_plan.is_backup_policy_plan", "cron_spec"),
					resource.TestCheckResourceAttrSet("data.ibm_is_backup_policy_plan.is_backup_policy_plan", "deletion_trigger.0.delete_after"),
					resource.TestCheckResourceAttrSet("data.ibm_is_backup_policy_plan.is_backup_policy_plan", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_backup_policy_plan.is_backup_policy_plan", "lifecycle_state"),
					resource.TestCheckResourceAttrSet("data.ibm_is_backup_policy_plan.is_backup_policy_plan", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_is_backup_policy_plan.is_backup_policy_plan", "resource_type"),
					resource.TestCheckResourceAttrSet("data.ibm_is_backup_policy_plan.is_backup_policy_plan", "clone_policy.0.max_snapshots"),
					resource.TestCheckResourceAttrSet("data.ibm_is_backup_policy_plan.is_backup_policy_plan", "clone_policy.0.zones.#"),
					resource.TestCheckResourceAttr("data.ibm_is_backup_policy_plan.is_backup_policy_plan", "clone_policy.0.zones.0", acc.ISZoneName),
					resource.TestCheckResourceAttr("data.ibm_is_backup_policy_plan.is_backup_policy_plan", "clone_policy.0.zones.1", acc.ISZoneName2),
				),
			},
		},
	})
}

func TestAccIBMIsBackupPolicyPlanDataSourceRemoteCopyPolicy(t *testing.T) {
	vpcname := fmt.Sprintf("tf-vpc-%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf-instnace-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tf-subnet-%d", acctest.RandIntRange(10, 100))
	sshname := fmt.Sprintf("tf-ssh-%d", acctest.RandIntRange(10, 100))
	volname := fmt.Sprintf("tf-vol-%d", acctest.RandIntRange(10, 100))
	bakupPolicyName := fmt.Sprintf("tfbakuppolicyname%d", acctest.RandIntRange(10, 100))
	bakupPolicyPlanName := fmt.Sprintf("tfbakuppolicyplanname%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsBackupPolicyPlansDataSourceConfigRemoteCopyPolicies(bakupPolicyName, vpcname, subnetname, sshname, volname, name, bakupPolicyPlanName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_backup_policy_plan.is_backup_policy_plan", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_backup_policy_plan.is_backup_policy_plan", "backup_policy_id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_backup_policy_plan.is_backup_policy_plan", "active"),
					resource.TestCheckResourceAttrSet("data.ibm_is_backup_policy_plan.is_backup_policy_plan", "attach_user_tags.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_backup_policy_plan.is_backup_policy_plan", "copy_user_tags"),
					resource.TestCheckResourceAttrSet("data.ibm_is_backup_policy_plan.is_backup_policy_plan", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_is_backup_policy_plan.is_backup_policy_plan", "cron_spec"),
					resource.TestCheckResourceAttrSet("data.ibm_is_backup_policy_plan.is_backup_policy_plan", "deletion_trigger.0.delete_after"),
					resource.TestCheckResourceAttrSet("data.ibm_is_backup_policy_plan.is_backup_policy_plan", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_backup_policy_plan.is_backup_policy_plan", "lifecycle_state"),
					resource.TestCheckResourceAttrSet("data.ibm_is_backup_policy_plan.is_backup_policy_plan", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_is_backup_policy_plan.is_backup_policy_plan", "resource_type"),
					resource.TestCheckResourceAttrSet("data.ibm_is_backup_policy_plan.is_backup_policy_plan", "remote_region_policy.0.delete_over_count"),
					resource.TestCheckResourceAttrSet("data.ibm_is_backup_policy_plan.is_backup_policy_plan", "remote_region_policy.0.encryption_key"),
					resource.TestCheckResourceAttrSet("data.ibm_is_backup_policy_plan.is_backup_policy_plan", "remote_region_policy.0.region"),
				),
			},
		},
	})
}

func testAccCheckIBMIsBackupPolicyPlanDataSourceConfigBasic(backupPolicyName, vpcname, subnetname, sshname, volName, name, cronSpec, bakupPolicyPlanName string) string {

	return testAccCheckIBMIsBackupPolicyPlanConfigBasic(backupPolicyName, vpcname, subnetname, sshname, volName, name, cronSpec, bakupPolicyPlanName) + fmt.Sprintf(`
		data "ibm_is_backup_policy_plan" "is_backup_policy_plan" {
			backup_policy_id = ibm_is_backup_policy.is_backup_policy.id
			identifier = ibm_is_backup_policy_plan.is_backup_policy_plan[0].backup_policy_plan_id
		}
	`)
}
func testAccCheckIBMIsBackupPolicyPlanDataSourceConfigClonesBasic(backupPolicyName, vpcname, subnetname, sshname, volName, name, cronSpec, bakupPolicyPlanName string) string {

	return testAccCheckIBMIsBackupPolicyPlanConfigClonesBasic(backupPolicyName, vpcname, subnetname, sshname, volName, name, cronSpec, bakupPolicyPlanName) + fmt.Sprintf(`
		data "ibm_is_backup_policy_plan" "is_backup_policy_plan" {
			backup_policy_id = ibm_is_backup_policy.is_backup_policy.id
			identifier = ibm_is_backup_policy_plan.is_backup_policy_plan[0].backup_policy_plan_id
		}
	`)
}
