// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/IBM/vpc-go-sdk/vpcv1"
)

func TestAccIBMIsBackupPolicyPlanBasic(t *testing.T) {
	var conf vpcv1.BackupPolicyPlan
	backupPolicyID := fmt.Sprintf("tf_backup_policy_id_%d", acctest.RandIntRange(10, 100))
	cronSpec := fmt.Sprintf("tf_cron_spec_%d", acctest.RandIntRange(10, 100))
	cronSpecUpdate := fmt.Sprintf("tf_cron_spec_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIsBackupPolicyPlanDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsBackupPolicyPlanConfigBasic(backupPolicyID, cronSpec),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIsBackupPolicyPlanExists("ibm_is_backup_policy_plan.is_backup_policy_plan", conf),
					resource.TestCheckResourceAttr("ibm_is_backup_policy_plan.is_backup_policy_plan", "backup_policy_id", backupPolicyID),
					resource.TestCheckResourceAttr("ibm_is_backup_policy_plan.is_backup_policy_plan", "cron_spec", cronSpec),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMIsBackupPolicyPlanConfigBasic(backupPolicyID, cronSpecUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_is_backup_policy_plan.is_backup_policy_plan", "backup_policy_id", backupPolicyID),
					resource.TestCheckResourceAttr("ibm_is_backup_policy_plan.is_backup_policy_plan", "cron_spec", cronSpecUpdate),
				),
			},
		},
	})
}

func TestAccIBMIsBackupPolicyPlanAllArgs(t *testing.T) {
	var conf vpcv1.BackupPolicyPlan
	backupPolicyID := fmt.Sprintf("tf_backup_policy_id_%d", acctest.RandIntRange(10, 100))
	cronSpec := fmt.Sprintf("tf_cron_spec_%d", acctest.RandIntRange(10, 100))
	active := "true"
	copyUserTags := "true"
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	cronSpecUpdate := fmt.Sprintf("tf_cron_spec_%d", acctest.RandIntRange(10, 100))
	activeUpdate := "false"
	copyUserTagsUpdate := "false"
	nameUpdate := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIsBackupPolicyPlanDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsBackupPolicyPlanConfig(backupPolicyID, cronSpec, active, copyUserTags, name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIsBackupPolicyPlanExists("ibm_is_backup_policy_plan.is_backup_policy_plan", conf),
					resource.TestCheckResourceAttr("ibm_is_backup_policy_plan.is_backup_policy_plan", "backup_policy_id", backupPolicyID),
					resource.TestCheckResourceAttr("ibm_is_backup_policy_plan.is_backup_policy_plan", "cron_spec", cronSpec),
					resource.TestCheckResourceAttr("ibm_is_backup_policy_plan.is_backup_policy_plan", "active", active),
					resource.TestCheckResourceAttr("ibm_is_backup_policy_plan.is_backup_policy_plan", "copy_user_tags", copyUserTags),
					resource.TestCheckResourceAttr("ibm_is_backup_policy_plan.is_backup_policy_plan", "name", name),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMIsBackupPolicyPlanConfig(backupPolicyID, cronSpecUpdate, activeUpdate, copyUserTagsUpdate, nameUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_is_backup_policy_plan.is_backup_policy_plan", "backup_policy_id", backupPolicyID),
					resource.TestCheckResourceAttr("ibm_is_backup_policy_plan.is_backup_policy_plan", "cron_spec", cronSpecUpdate),
					resource.TestCheckResourceAttr("ibm_is_backup_policy_plan.is_backup_policy_plan", "active", activeUpdate),
					resource.TestCheckResourceAttr("ibm_is_backup_policy_plan.is_backup_policy_plan", "copy_user_tags", copyUserTagsUpdate),
					resource.TestCheckResourceAttr("ibm_is_backup_policy_plan.is_backup_policy_plan", "name", nameUpdate),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_is_backup_policy_plan.is_backup_policy_plan",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIBMIsBackupPolicyPlanConfigBasic(backupPolicyID string, cronSpec string) string {
	return fmt.Sprintf(`

		resource "ibm_is_backup_policy_plan" "is_backup_policy_plan" {
			backup_policy_id = "%s"
			cron_spec = "%s"
		}
	`, backupPolicyID, cronSpec)
}

func testAccCheckIBMIsBackupPolicyPlanConfig(backupPolicyID string, cronSpec string, active string, copyUserTags string, name string) string {
	return fmt.Sprintf(`

		resource "ibm_is_backup_policy_plan" "is_backup_policy_plan" {
			backup_policy_id = "%s"
			cron_spec = "%s"
			active = %s
			attach_user_tags = "FIXME"
			clone_policy {
				max_snapshots = 1
				zones {
					name = "us-south-1"
					href = "https://us-south.iaas.cloud.ibm.com/v1/regions/us-south/zones/us-south-1"
				}
			}
			copy_user_tags = %s
			deletion_trigger {
				delete_after = 20
				delete_over_count = 20
			}
			name = "%s"
		}
	`, backupPolicyID, cronSpec, active, copyUserTags, name)
}

func testAccCheckIBMIsBackupPolicyPlanExists(n string, obj vpcv1.BackupPolicyPlan) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		vpcClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).VpcV1API()
		if err != nil {
			return err
		}

		getBackupPolicyPlanOptions := &vpcv1.GetBackupPolicyPlanOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getBackupPolicyPlanOptions.SetBackupPolicyID(parts[0])
		getBackupPolicyPlanOptions.SetID(parts[1])

		backupPolicyPlan, _, err := vpcClient.GetBackupPolicyPlan(getBackupPolicyPlanOptions)
		if err != nil {
			return err
		}

		obj = *backupPolicyPlan
		return nil
	}
}

func testAccCheckIBMIsBackupPolicyPlanDestroy(s *terraform.State) error {
	vpcClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).VpcV1API()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_is_backup_policy_plan" {
			continue
		}

		getBackupPolicyPlanOptions := &vpcv1.GetBackupPolicyPlanOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getBackupPolicyPlanOptions.SetBackupPolicyID(parts[0])
		getBackupPolicyPlanOptions.SetID(parts[1])

		// Try to find the key
		_, response, err := vpcClient.GetBackupPolicyPlan(getBackupPolicyPlanOptions)

		if err == nil {
			return fmt.Errorf("BackupPolicyPlan still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for BackupPolicyPlan (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
