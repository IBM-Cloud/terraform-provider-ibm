// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"fmt"
	"strings"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/IBM/vpc-go-sdk/vpcv1"
)

func TestAccIBMIsBackupPolicyBasic(t *testing.T) {
	vpcname := fmt.Sprintf("tf-vpc-%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf-instnace-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tf-subnet-%d", acctest.RandIntRange(10, 100))
	publicKey := strings.TrimSpace(`
ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCKVmnMOlHKcZK8tpt3MP1lqOLAcqcJzhsvJcjscgVERRN7/9484SOBJ3HSKxxNG5JN8owAjy5f9yYwcUg+JaUVuytn5Pv3aeYROHGGg+5G346xaq3DAwX6Y5ykr2fvjObgncQBnuU5KHWCECO/4h8uWuwh/kfniXPVjFToc+gnkqA+3RKpAecZhFXwfalQ9mMuYGFxn+fwn8cYEApsJbsEmb0iJwPiZ5hjFC8wREuiTlhPHDgkBLOiycd20op2nXzDbHfCHInquEe/gYxEitALONxm0swBOwJZwlTDOB7C6y2dzlrtxr1L59m7pCkWI4EtTRLvleehBoj3u7jB4usR
`)
	sshname := fmt.Sprintf("tf-ssh-%d", acctest.RandIntRange(10, 100))
	volname := fmt.Sprintf("tf-vol-%d", acctest.RandIntRange(10, 100))
	backupPolicyName := fmt.Sprintf("tfbakuppolicyname%d", acctest.RandIntRange(10, 100))
	backupPolicyNameUpdate := fmt.Sprintf("tfbakuppolicyname%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIsBackupPolicyDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsBackupPolicyConfigBasic(backupPolicyName, vpcname, subnetname, sshname, publicKey, volname, name),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_is_backup_policy.is_backup_policy", "name", backupPolicyName),
					resource.TestCheckResourceAttrSet("ibm_is_backup_policy.is_backup_policy", "match_resource_types.#"),
					resource.TestCheckResourceAttrSet("ibm_is_backup_policy.is_backup_policy", "match_user_tags.#"),
					resource.TestCheckResourceAttrSet("ibm_is_backup_policy.is_backup_policy", "resource_group.#"),
					resource.TestCheckResourceAttrSet("ibm_is_backup_policy.is_backup_policy", "created_at"),
					resource.TestCheckResourceAttrSet("ibm_is_backup_policy.is_backup_policy", "crn"),
					resource.TestCheckResourceAttrSet("ibm_is_backup_policy.is_backup_policy", "href"),
					resource.TestCheckResourceAttrSet("ibm_is_backup_policy.is_backup_policy", "lifecycle_state"),
					resource.TestCheckResourceAttrSet("ibm_is_backup_policy.is_backup_policy", "resource_type"),
					resource.TestCheckResourceAttrSet("ibm_is_backup_policy.is_backup_policy", "version"),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMIsBackupPolicyConfigBasic(backupPolicyNameUpdate, vpcname, subnetname, sshname, publicKey, volname, name),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_is_backup_policy.is_backup_policy", "name", backupPolicyNameUpdate),
					resource.TestCheckResourceAttrSet("ibm_is_backup_policy.is_backup_policy", "match_resource_types.#"),
					resource.TestCheckResourceAttrSet("ibm_is_backup_policy.is_backup_policy", "match_user_tags.#"),
					resource.TestCheckResourceAttrSet("ibm_is_backup_policy.is_backup_policy", "resource_group.#"),
					resource.TestCheckResourceAttrSet("ibm_is_backup_policy.is_backup_policy", "created_at"),
					resource.TestCheckResourceAttrSet("ibm_is_backup_policy.is_backup_policy", "crn"),
					resource.TestCheckResourceAttrSet("ibm_is_backup_policy.is_backup_policy", "href"),
					resource.TestCheckResourceAttrSet("ibm_is_backup_policy.is_backup_policy", "lifecycle_state"),
					resource.TestCheckResourceAttrSet("ibm_is_backup_policy.is_backup_policy", "resource_type"),
					resource.TestCheckResourceAttrSet("ibm_is_backup_policy.is_backup_policy", "version"),
				),
			},
		},
	})
}

func testAccCheckIBMIsBackupPolicyConfigBasic(backupPolicyName string, vpcname, subnetname, sshname, publicKey, volName, name string) string {
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
		keys    = ["r134-0391ca6e-2437-429c-8d58-1ac186c53555"]
		volumes = [ibm_is_volume.storage.id]
	  }

	  resource "ibm_is_backup_policy" "is_backup_policy" {
		depends_on  = [ibm_is_instance.testacc_instance]
		match_user_tags = ["tag-0"]
		name            = "%s"
	}`, vpcname, subnetname, acc.ISZoneName, acc.ISCIDR, volName, acc.ISZoneName, name, "r134-8070fddf-88bf-4c58-a0f4-05a8306af951", acc.InstanceProfileName, acc.ISZoneName, backupPolicyName)
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
