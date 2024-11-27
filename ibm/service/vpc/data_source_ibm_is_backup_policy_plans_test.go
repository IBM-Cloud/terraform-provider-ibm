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
	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
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
					resource.TestCheckResourceAttrSet("data.ibm_is_backup_policy_plans.is_backup_policy_plans", "plans.0.remote_region_policy.0.delete_over_count"),
					resource.TestCheckResourceAttrSet("data.ibm_is_backup_policy_plans.is_backup_policy_plans", "plans.0.remote_region_policy.0.encryption_key"),
					resource.TestCheckResourceAttrSet("data.ibm_is_backup_policy_plans.is_backup_policy_plans", "plans.0.remote_region_policy.0.region"),
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

func TestAccIBMIsBackupPolicyPlansDataSourceRemoteCopyPolicy(t *testing.T) {
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
					resource.TestCheckResourceAttrSet("data.ibm_is_backup_policy_plans.is_backup_policy_plans", "plans.0.remote_region_policy.0.delete_over_count"),
					resource.TestCheckResourceAttrSet("data.ibm_is_backup_policy_plans.is_backup_policy_plans", "plans.0.remote_region_policy.0.encryption_key"),
					resource.TestCheckResourceAttrSet("data.ibm_is_backup_policy_plans.is_backup_policy_plans", "plans.0.remote_region_policy.0.region"),
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

func testAccCheckIBMIsBackupPolicyPlansDataSourceConfigRemoteCopyPolicies(backupPolicyName, vpcname, subnetname, sshname, volName, name, bakupPolicyPlanName string) string {
	return fmt.Sprintf(`
	resource "ibm_is_vpc" "testacc_vpc" {
		name = "%s"
	  }

	  resource "ibm_is_subnet" "testacc_subnet" {
		name            = "%s"
		vpc             = ibm_is_vpc.testacc_vpc.id
		zone            = "%s"
		ipv4_cidr_block = "%s"
	  }

	  resource "ibm_is_ssh_key" "testacc_sshkey" {
		name       = "%s"
		public_key = file("~/.ssh/id_rsa.pub")
	  }

	  resource "ibm_is_volume" "storage" {
		name    = "%s"
		profile = "10iops-tier"
		zone    = "%s"
		# capacity= 200
		tags 	= ["tag-0"]
	  }

	  resource "ibm_is_instance" "testacc_instance" {
		name    = "%s"
		image   = "%s"
		profile = "%s"
		primary_network_interface {
		  subnet = ibm_is_subnet.testacc_subnet.id
		}
		vpc     = ibm_is_vpc.testacc_vpc.id
		zone    = "%s"
		keys    = [ibm_is_ssh_key.testacc_sshkey.id]
		volumes = [ibm_is_volume.storage.id]
	  }

	  resource "ibm_is_backup_policy" "is_backup_policy" {
		depends_on  		= [ibm_is_instance.testacc_instance]
		match_user_tags 	= ["tag-0"]
		name            	= "%s"
		remote_region_policy {
			delete_over_count 		= 1
			encryption_key 			= "%s"
			region 					= "us-south"
		}
	}`, vpcname, subnetname, acc.ISZoneName, acc.ISCIDR, sshname, volName, acc.ISZoneName, name, acc.IsImage, acc.InstanceProfileName, acc.ISZoneName, backupPolicyName, acc.BaasEncryptionkeyCRN)
}
