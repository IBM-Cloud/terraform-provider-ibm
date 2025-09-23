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

func TestAccIBMIsBackupPolicyDataSourceBasic(t *testing.T) {
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
				Config: testAccCheckIBMIsBackupPolicyDataSourceConfigBasic(bakupPolicyName, vpcname, subnetname, sshname, volname, name, cronSpec, bakupPolicyPlanName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_backup_policy.is_backup_policy", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_is_backup_policy.is_backup_policy", "crn"),
					resource.TestCheckResourceAttrSet("data.ibm_is_backup_policy.is_backup_policy", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_backup_policy.is_backup_policy", "lifecycle_state"),
					resource.TestCheckResourceAttrSet("data.ibm_is_backup_policy.is_backup_policy", "match_resource_types.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_backup_policy.is_backup_policy", "match_user_tags.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_backup_policy.is_backup_policy", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_is_backup_policy.is_backup_policy", "plans.#"),
					// resource.TestCheckResourceAttrSet("data.ibm_is_backup_policy.is_backup_policy", "plans.0.href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_backup_policy.is_backup_policy", "plans.0.id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_backup_policy.is_backup_policy", "plans.0.name"),
					// resource.TestCheckResourceAttrSet("data.ibm_is_backup_policy.is_backup_policy", "plans.0.resource_type"),
					resource.TestCheckResourceAttrSet("data.ibm_is_backup_policy.is_backup_policy", "resource_group.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_backup_policy.is_backup_policy", "resource_group.0.href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_backup_policy.is_backup_policy", "resource_group.0.id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_backup_policy.is_backup_policy", "resource_group.0.name"),
					resource.TestCheckResourceAttrSet("data.ibm_is_backup_policy.is_backup_policy", "resource_type"),
				),
			},
		},
	})
}

func testAccCheckIBMIsBackupPolicyDataSourceConfigBasic(backupPolicyName, vpcname, subnetname, sshname, volName, name, cronSpec, bakupPolicyPlanName string) string {

	return testAccCheckIBMIsBackupPolicyPlanConfigBasic(backupPolicyName, vpcname, subnetname, sshname, volName, name, cronSpec, bakupPolicyPlanName) + fmt.Sprintf(`
		data "ibm_is_backup_policy" "is_backup_policy" {
			depends_on  = [ibm_is_backup_policy_plan.is_backup_policy_plan]
			identifier = ibm_is_backup_policy.is_backup_policy.id
		}
	`)
}
