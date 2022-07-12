// Copyright IBM Corp. 2021, 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
	"time"

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
	vpcname := fmt.Sprintf("tf-vpc-%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf-instnace-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tf-subnet-%d", acctest.RandIntRange(10, 100))
	sshname := fmt.Sprintf("tf-ssh-%d", acctest.RandIntRange(10, 100))
	volname := fmt.Sprintf("tf-vol-%d", acctest.RandIntRange(10, 100))
	bakupPolicyName := fmt.Sprintf("tfbakuppolicyname%d", acctest.RandIntRange(10, 100))
	bakupPolicyPlanName := fmt.Sprintf("tfbakuppolicyplanname%d", acctest.RandIntRange(10, 100))
	bakupPolicyPlanNameUpdate := fmt.Sprintf("tfbakuppolicyplannameupdate%d", acctest.RandIntRange(10, 100))
	cronSpec := strings.TrimSpace(strconv.Itoa(time.Now().UTC().Minute()) + " " + strconv.Itoa(time.Now().UTC().Hour()) + " " + "*" + " " + "*" + " " + "*")
	cronSpecUpdate := strings.TrimSpace(strconv.Itoa(time.Now().UTC().Minute()+1) + " " + strconv.Itoa(time.Now().UTC().Hour()) + " " + "*" + " " + "*" + " " + "*")
	if strings.TrimSpace(strconv.Itoa(time.Now().UTC().Minute())) == "61" {
		cronSpecUpdate = strings.TrimSpace(("1") + " " + strconv.Itoa(time.Now().UTC().Hour()+1) + " " + "*" + " " + "*" + " " + "*")
	}
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIsBackupPolicyPlanDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsBackupPolicyPlanConfigBasic(bakupPolicyName, vpcname, subnetname, sshname, volname, name, cronSpec, bakupPolicyPlanName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIsBackupPolicyPlanExists("ibm_is_backup_policy_plan.is_backup_policy_plan.0", conf),
					resource.TestCheckResourceAttr("ibm_is_backup_policy_plan.is_backup_policy_plan.0", "name", bakupPolicyPlanName+"-1"),
					resource.TestCheckResourceAttr("ibm_is_backup_policy_plan.is_backup_policy_plan.0", "cron_spec", cronSpec),
					resource.TestCheckResourceAttrSet("ibm_is_backup_policy_plan.is_backup_policy_plan.0", "backup_policy_id"),
					resource.TestCheckResourceAttrSet("ibm_is_backup_policy_plan.is_backup_policy_plan.0", "backup_policy_plan_id"),
					resource.TestCheckResourceAttrSet("ibm_is_backup_policy_plan.is_backup_policy_plan.0", "active"),
					resource.TestCheckResourceAttrSet("ibm_is_backup_policy_plan.is_backup_policy_plan.0", "copy_user_tags"),
					resource.TestCheckResourceAttrSet("ibm_is_backup_policy_plan.is_backup_policy_plan.0", "deletion_trigger.#"),
					resource.TestCheckResourceAttrSet("ibm_is_backup_policy_plan.is_backup_policy_plan.0", "created_at"),
					resource.TestCheckResourceAttrSet("ibm_is_backup_policy_plan.is_backup_policy_plan.0", "href"),
					resource.TestCheckResourceAttrSet("ibm_is_backup_policy_plan.is_backup_policy_plan.0", "lifecycle_state"),
					resource.TestCheckResourceAttrSet("ibm_is_backup_policy_plan.is_backup_policy_plan.0", "resource_type"),
					resource.TestCheckResourceAttrSet("ibm_is_backup_policy_plan.is_backup_policy_plan.0", "version"),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMIsBackupPolicyPlanConfigBasic(bakupPolicyName, vpcname, subnetname, sshname, volname, name, cronSpecUpdate, bakupPolicyPlanNameUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIsBackupPolicyPlanExists("ibm_is_backup_policy_plan.is_backup_policy_plan.0", conf),
					resource.TestCheckResourceAttr("ibm_is_backup_policy_plan.is_backup_policy_plan.0", "name", bakupPolicyPlanNameUpdate+"-1"),
					resource.TestCheckResourceAttr("ibm_is_backup_policy_plan.is_backup_policy_plan.0", "cron_spec", cronSpecUpdate),
					resource.TestCheckResourceAttrSet("ibm_is_backup_policy_plan.is_backup_policy_plan.0", "backup_policy_id"),
					resource.TestCheckResourceAttrSet("ibm_is_backup_policy_plan.is_backup_policy_plan.0", "backup_policy_plan_id"),
					resource.TestCheckResourceAttrSet("ibm_is_backup_policy_plan.is_backup_policy_plan.0", "active"),
					resource.TestCheckResourceAttrSet("ibm_is_backup_policy_plan.is_backup_policy_plan.0", "copy_user_tags"),
					resource.TestCheckResourceAttrSet("ibm_is_backup_policy_plan.is_backup_policy_plan.0", "deletion_trigger.#"),
					resource.TestCheckResourceAttrSet("ibm_is_backup_policy_plan.is_backup_policy_plan.0", "created_at"),
					resource.TestCheckResourceAttrSet("ibm_is_backup_policy_plan.is_backup_policy_plan.0", "href"),
					resource.TestCheckResourceAttrSet("ibm_is_backup_policy_plan.is_backup_policy_plan.0", "lifecycle_state"),
					resource.TestCheckResourceAttrSet("ibm_is_backup_policy_plan.is_backup_policy_plan.0", "resource_type"),
					resource.TestCheckResourceAttrSet("ibm_is_backup_policy_plan.is_backup_policy_plan.0", "version"),
				),
			},
		},
	})
}

func TestAccIBMIsBackupPolicyPlanImport(t *testing.T) {
	var conf vpcv1.BackupPolicyPlan
	vpcname := fmt.Sprintf("tf-vpc-%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf-instnace-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tf-subnet-%d", acctest.RandIntRange(10, 100))
	sshname := fmt.Sprintf("tf-ssh-%d", acctest.RandIntRange(10, 100))
	volname := fmt.Sprintf("tf-vol-%d", acctest.RandIntRange(10, 100))
	bakupPolicyName := fmt.Sprintf("tfbakuppolicynameimport%d", acctest.RandIntRange(10, 100))
	bakupPolicyPlanName := fmt.Sprintf("tfbakuppolicyplannameimport%d", acctest.RandIntRange(10, 100))
	// bakupPolicyPlanNameUpdate := fmt.Sprintf("tfbakuppolicyplannameupdate%d", acctest.RandIntRange(10, 100))
	cronSpec := strings.TrimSpace(strconv.Itoa(time.Now().UTC().Minute()) + " " + strconv.Itoa(time.Now().UTC().Hour()) + " " + "*" + " " + "*" + " " + "*")
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIsBackupPolicyPlanDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsBackupPolicyPlanConfigImport(bakupPolicyName, vpcname, subnetname, sshname, volname, name, cronSpec, bakupPolicyPlanName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIsBackupPolicyPlanExists("ibm_is_backup_policy_plan.is_backup_policy_plan_import", conf),
					resource.TestCheckResourceAttr("ibm_is_backup_policy_plan.is_backup_policy_plan_import", "name", bakupPolicyPlanName+"-import"),
					resource.TestCheckResourceAttr("ibm_is_backup_policy_plan.is_backup_policy_plan_import", "cron_spec", cronSpec),
					resource.TestCheckResourceAttrSet("ibm_is_backup_policy_plan.is_backup_policy_plan_import", "backup_policy_id"),
					resource.TestCheckResourceAttrSet("ibm_is_backup_policy_plan.is_backup_policy_plan_import", "backup_policy_plan_id"),
					resource.TestCheckResourceAttrSet("ibm_is_backup_policy_plan.is_backup_policy_plan_import", "active"),
					resource.TestCheckResourceAttrSet("ibm_is_backup_policy_plan.is_backup_policy_plan_import", "copy_user_tags"),
					resource.TestCheckResourceAttrSet("ibm_is_backup_policy_plan.is_backup_policy_plan_import", "deletion_trigger.#"),
					resource.TestCheckResourceAttrSet("ibm_is_backup_policy_plan.is_backup_policy_plan_import", "created_at"),
					resource.TestCheckResourceAttrSet("ibm_is_backup_policy_plan.is_backup_policy_plan_import", "href"),
					resource.TestCheckResourceAttrSet("ibm_is_backup_policy_plan.is_backup_policy_plan_import", "lifecycle_state"),
					resource.TestCheckResourceAttrSet("ibm_is_backup_policy_plan.is_backup_policy_plan_import", "resource_type"),
					resource.TestCheckResourceAttrSet("ibm_is_backup_policy_plan.is_backup_policy_plan_import", "version"),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_is_backup_policy_plan.is_backup_policy_plan_import",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIBMIsBackupPolicyPlanConfigBasic(backupPolicyName, vpcname, subnetname, sshname, volName, name, cronSpec, bakupPolicyPlanName string) string {

	return testAccCheckIBMIsBackupPolicyConfigBasic(backupPolicyName, vpcname, subnetname, sshname, volName, name) + fmt.Sprintf(`
		resource "ibm_is_backup_policy_plan" "is_backup_policy_plan" {
			count  = 2
			backup_policy_id = ibm_is_backup_policy.is_backup_policy.id
			name = "%s-${count.index+1}"
			cron_spec = "%s"
		}
	`, bakupPolicyPlanName, cronSpec)
}

func testAccCheckIBMIsBackupPolicyPlanConfigImport(backupPolicyName, vpcname, subnetname, sshname, volName, name, cronSpec, bakupPolicyPlanName string) string {

	return testAccCheckIBMIsBackupPolicyConfigBasic(backupPolicyName, vpcname, subnetname, sshname, volName, name) + fmt.Sprintf(`
		resource "ibm_is_backup_policy_plan" "is_backup_policy_plan_import" {
			backup_policy_id = ibm_is_backup_policy.is_backup_policy.id
			name = "%s-import"
			cron_spec = "%s"
		}
	`, bakupPolicyPlanName, cronSpec)
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
