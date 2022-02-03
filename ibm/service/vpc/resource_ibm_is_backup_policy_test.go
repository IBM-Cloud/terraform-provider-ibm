// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/IBM/vpc-go-sdk/vpcv1"
)

func TestAccIBMIsBackupPolicyBasic(t *testing.T) {
	// var conf vpcv1.BackupPolicy
	bakupPolicyName := fmt.Sprintf("tfbakuppolicyname%d", acctest.RandIntRange(10, 100))
	// bakupPolicyNameUpdate := fmt.Sprintf("tfbakuppolicyname%d", acctest.RandIntRange(10, 100))
	bakupPolicyPlanName := fmt.Sprintf("tfbakuppolicyplanname%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIsBackupPolicyDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsBackupPolicyConfigBasic(bakupPolicyName, bakupPolicyPlanName),
				// Check: resource.ComposeAggregateTestCheckFunc(
				// 	testAccCheckIBMIsBackupPolicyExists("ibm_is_backup_policy.is_backup_policy", conf),
				// 	resource.TestCheckResourceAttr("ibm_is_backup_policy.is_backup_policy", "name", bakupPolicyName),
				// ),
			},
			// resource.TestStep{
			// 	Config: testAccCheckIBMIsBackupPolicyConfigBasic(bakupPolicyNameUpdate, bakupPolicyPlanName),
			// 	Check: resource.ComposeAggregateTestCheckFunc(
			// 		resource.TestCheckResourceAttr("ibm_is_backup_policy.is_backup_policy", "name", bakupPolicyNameUpdate),
			// 	),
			// },
		},
	})
}

func testAccCheckIBMIsBackupPolicyConfigBasic(name string, planName string) string {
	// return fmt.Sprintf(`

	// 	resource "ibm_is_backup_policy" "is_backup_policy" {
	// 		match_user_tags = ["barbq"]
	// 		name = "%s"
	// 		plans {
	// 			name = "%s"
	// 			cron_spec = "*/5 1,2,3 * * *"
	// 			delete_after = 40
	// 			copy_user_tags = true
	// 		}
	// 	}
	// `, name, planName)

	// return fmt.Sprintf(`	resource "ibm_is_backup_policy" "is_backup_policy" {
	// 	match_user_tags = ["sunithabackuppolicy"]
	// 	name            = "sunitha-backup-policy"
	// 	plans {
	// 	  name      = "sunitha-backup-policy-plan"
	// 	  cron_spec = "*/5 1,2,3 * * *"
	// 	  deletion_trigger {
	// 		delete_after      = 20
	// 		delete_over_count = 20
	// 	  }
	// 	  copy_user_tags = true
	// 	}
	//   }`)
	return fmt.Sprintf(`
	resource "ibm_is_backup_policy" "is_backup_policy" {
		match_user_tags = ["zs-backup-demo"]
		name            = "sunitha-backup-policy-checking"
		plans {
		  name      = "sunitha-backup-policy-plan"
		  cron_spec = "30 09 * * *" 
		  deletion_trigger {
			delete_after      = 20
			delete_over_count = 20
		  }
		  copy_user_tags = true
		}
	  }`)
}

func testAccCheckIBMIsBackupPolicyExists(n string, obj vpcv1.BackupPolicy) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		vpcClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).VpcV1API()
		if err != nil {
			return err
		}

		getBackupPolicyOptions := &vpcv1.GetBackupPolicyOptions{}

		getBackupPolicyOptions.SetID(rs.Primary.ID)

		backupPolicy, _, err := vpcClient.GetBackupPolicy(getBackupPolicyOptions)
		if err != nil {
			return err
		}

		obj = *backupPolicy
		return nil
	}
}

func testAccCheckIBMIsBackupPolicyDestroy(s *terraform.State) error {
	vpcClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).VpcV1API()

	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_is_backup_policy" {
			continue
		}

		getBackupPolicyOptions := &vpcv1.GetBackupPolicyOptions{}

		getBackupPolicyOptions.SetID(rs.Primary.ID)

		// Try to find the key
		_, response, err := vpcClient.GetBackupPolicy(getBackupPolicyOptions)

		if err == nil {
			return fmt.Errorf("BackupPolicy still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for BackupPolicy (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
