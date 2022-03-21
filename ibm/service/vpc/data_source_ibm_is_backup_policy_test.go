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

func TestAccIBMIsBackupPolicyDataSourceBasic(t *testing.T) {
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
	cronSpec := strings.TrimSpace(strconv.Itoa(time.Now().UTC().Minute()+1) + " " + strconv.Itoa(time.Now().UTC().Second()) + " " + "*" + " " + "*" + " " + "*")
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsBackupPolicyDataSourceConfigBasic(bakupPolicyName, vpcname, subnetname, sshname, publicKey, volname, name, cronSpec, bakupPolicyPlanName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_backup_policy.is_backup_policy", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_is_backup_policy.is_backup_policy", "crn"),
					resource.TestCheckResourceAttrSet("data.ibm_is_backup_policy.is_backup_policy", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_backup_policy.is_backup_policy", "lifecycle_state"),
					resource.TestCheckResourceAttrSet("data.ibm_is_backup_policy.is_backup_policy", "match_resource_types.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_backup_policy.is_backup_policy", "match_user_tags.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_backup_policy.is_backup_policy", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_is_backup_policy.is_backup_policy", "plans.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_backup_policy.is_backup_policy", "resource_group.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_backup_policy.is_backup_policy", "resource_type"),
				),
			},
		},
	})
}

func testAccCheckIBMIsBackupPolicyDataSourceConfigBasic(backupPolicyName, vpcname, subnetname, sshname, publicKey, volName, name, bakupPolicyPlanName, cronSpec string) string {

	return testAccCheckIBMIsBackupPolicyPlanConfigBasic(backupPolicyName, vpcname, subnetname, sshname, publicKey, volName, name, cronSpec, bakupPolicyPlanName) + fmt.Sprintf(`
		data "ibm_is_backup_policy" "is_backup_policy" {
			identifier = ibm_is_backup_policy.is_backup_policy.id
		}
	`)
}
