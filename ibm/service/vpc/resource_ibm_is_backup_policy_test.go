// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"fmt"
	"testing"

	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/conns"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/IBM/vpc-go-sdk/vpcv1"
)

func TestAccIBMIsBackupPolicyBasic(t *testing.T) {
	vpcname := fmt.Sprintf("tf-vpc-%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf-instnace-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tf-subnet-%d", acctest.RandIntRange(10, 100))
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
				Config: testAccCheckIBMIsBackupPolicyConfigBasic(backupPolicyName, vpcname, subnetname, sshname, volname, name),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_is_backup_policy.is_backup_policy", "name", backupPolicyName),
					resource.TestCheckResourceAttrSet("ibm_is_backup_policy.is_backup_policy", "match_resource_types.#"),
					resource.TestCheckResourceAttrSet("ibm_is_backup_policy.is_backup_policy", "match_user_tags.#"),
					resource.TestCheckResourceAttrSet("ibm_is_backup_policy.is_backup_policy", "resource_group"),
					resource.TestCheckResourceAttrSet("ibm_is_backup_policy.is_backup_policy", "created_at"),
					resource.TestCheckResourceAttrSet("ibm_is_backup_policy.is_backup_policy", "crn"),
					resource.TestCheckResourceAttrSet("ibm_is_backup_policy.is_backup_policy", "href"),
					resource.TestCheckResourceAttrSet("ibm_is_backup_policy.is_backup_policy", "lifecycle_state"),
					resource.TestCheckResourceAttrSet("ibm_is_backup_policy.is_backup_policy", "resource_type"),
					resource.TestCheckResourceAttrSet("ibm_is_backup_policy.is_backup_policy", "version"),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMIsBackupPolicyConfigBasic(backupPolicyNameUpdate, vpcname, subnetname, sshname, volname, name),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_is_backup_policy.is_backup_policy", "name", backupPolicyNameUpdate),
					resource.TestCheckResourceAttrSet("ibm_is_backup_policy.is_backup_policy", "match_resource_types.#"),
					resource.TestCheckResourceAttrSet("ibm_is_backup_policy.is_backup_policy", "match_user_tags.#"),
					resource.TestCheckResourceAttrSet("ibm_is_backup_policy.is_backup_policy", "resource_group"),
					resource.TestCheckResourceAttrSet("ibm_is_backup_policy.is_backup_policy", "created_at"),
					resource.TestCheckResourceAttrSet("ibm_is_backup_policy.is_backup_policy", "crn"),
					resource.TestCheckResourceAttrSet("ibm_is_backup_policy.is_backup_policy", "href"),
					resource.TestCheckResourceAttrSet("ibm_is_backup_policy.is_backup_policy", "lifecycle_state"),
					resource.TestCheckResourceAttrSet("ibm_is_backup_policy.is_backup_policy", "resource_type"),
					resource.TestCheckResourceAttrSet("ibm_is_backup_policy.is_backup_policy", "version"),
				),
			},
			{
				ResourceName:      "ibm_is_backup_policy.is_backup_policy",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccIBMIsBackupPolicyMatchResourceType(t *testing.T) {
	vpcname := fmt.Sprintf("tf-vpc-%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf-instnace-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tf-subnet-%d", acctest.RandIntRange(10, 100))
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
				Config: testAccCheckIBMIsBackupPolicyConfigMatchResourceType(backupPolicyName, vpcname, subnetname, sshname, volname, name),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_is_backup_policy.is_backup_policy", "name", backupPolicyName),
					resource.TestCheckResourceAttrSet("ibm_is_backup_policy.is_backup_policy", "match_resource_types.#"),
					resource.TestCheckResourceAttrSet("ibm_is_backup_policy.is_backup_policy", "match_user_tags.#"),
					resource.TestCheckResourceAttrSet("ibm_is_backup_policy.is_backup_policy", "resource_group"),
					resource.TestCheckResourceAttrSet("ibm_is_backup_policy.is_backup_policy", "created_at"),
					resource.TestCheckResourceAttrSet("ibm_is_backup_policy.is_backup_policy", "crn"),
					resource.TestCheckResourceAttrSet("ibm_is_backup_policy.is_backup_policy", "href"),
					resource.TestCheckResourceAttrSet("ibm_is_backup_policy.is_backup_policy", "lifecycle_state"),
					resource.TestCheckResourceAttrSet("ibm_is_backup_policy.is_backup_policy", "resource_type"),
					resource.TestCheckResourceAttrSet("ibm_is_backup_policy.is_backup_policy", "version"),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMIsBackupPolicyConfigMatchResourceType(backupPolicyNameUpdate, vpcname, subnetname, sshname, volname, name),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_is_backup_policy.is_backup_policy", "name", backupPolicyNameUpdate),
					resource.TestCheckResourceAttrSet("ibm_is_backup_policy.is_backup_policy", "match_resource_types.#"),
					resource.TestCheckResourceAttrSet("ibm_is_backup_policy.is_backup_policy", "match_user_tags.#"),
					resource.TestCheckResourceAttrSet("ibm_is_backup_policy.is_backup_policy", "resource_group"),
					resource.TestCheckResourceAttrSet("ibm_is_backup_policy.is_backup_policy", "created_at"),
					resource.TestCheckResourceAttrSet("ibm_is_backup_policy.is_backup_policy", "crn"),
					resource.TestCheckResourceAttrSet("ibm_is_backup_policy.is_backup_policy", "href"),
					resource.TestCheckResourceAttrSet("ibm_is_backup_policy.is_backup_policy", "lifecycle_state"),
					resource.TestCheckResourceAttrSet("ibm_is_backup_policy.is_backup_policy", "resource_type"),
					resource.TestCheckResourceAttrSet("ibm_is_backup_policy.is_backup_policy", "version"),
				),
			},
			{
				ResourceName:      "ibm_is_backup_policy.is_backup_policy",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIBMIsBackupPolicyConfigBasic(backupPolicyName string, vpcname, subnetname, sshname, volName, name string) string {
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
		// public_key = file("../../test-fixtures/.ssh/id_rsa")
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
		depends_on  = [ibm_is_instance.testacc_instance]
		match_user_tags = ["tag-0"]
		name            = "%s"
	}`, vpcname, subnetname, acc.ISZoneName, acc.ISCIDR, sshname, volName, acc.ISZoneName, name, acc.IsImage, acc.InstanceProfileName, acc.ISZoneName, backupPolicyName)
}

func testAccCheckIBMIsBackupPolicyConfigMatchResourceType(backupPolicyName string, vpcname, subnetname, sshname, volName, name string) string {
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
		// public_key = file("../../test-fixtures/.ssh/id_rsa")
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
		depends_on  = [ibm_is_instance.testacc_instance]
		match_user_tags = ["tag-0"]
		match_resource_type = "volume"
		name            = "%s"
	}`, vpcname, subnetname, acc.ISZoneName, acc.ISCIDR, sshname, volName, acc.ISZoneName, name, acc.IsImage, acc.InstanceProfileName, acc.ISZoneName, backupPolicyName)
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

func TestAccIBMIsBackupPolicyBasicWithScope(t *testing.T) {
	backupPolicyName := fmt.Sprintf("tfbakuppolicyname%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIsBackupPolicyDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsBackupPolicyConfigBasicWithScope(backupPolicyName, acc.EnterpriseCRN),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_is_backup_policy.is_backup_policy", "name", backupPolicyName),
					resource.TestCheckResourceAttrSet("ibm_is_backup_policy.is_backup_policy", "match_resource_types.#"),
					resource.TestCheckResourceAttrSet("ibm_is_backup_policy.is_backup_policy", "match_user_tags.#"),
					resource.TestCheckResourceAttrSet("ibm_is_backup_policy.is_backup_policy", "resource_group"),
					resource.TestCheckResourceAttrSet("ibm_is_backup_policy.is_backup_policy", "created_at"),
					resource.TestCheckResourceAttrSet("ibm_is_backup_policy.is_backup_policy", "crn"),
					resource.TestCheckResourceAttrSet("ibm_is_backup_policy.is_backup_policy", "href"),
					resource.TestCheckResourceAttrSet("ibm_is_backup_policy.is_backup_policy", "scope.#"),
					resource.TestCheckResourceAttrSet("ibm_is_backup_policy.is_backup_policy", "scope.0.id"),
					resource.TestCheckResourceAttrSet("ibm_is_backup_policy.is_backup_policy", "lifecycle_state"),
					resource.TestCheckResourceAttrSet("ibm_is_backup_policy.is_backup_policy", "resource_type"),
					resource.TestCheckResourceAttrSet("ibm_is_backup_policy.is_backup_policy", "version"),
				),
			},
			{
				ResourceName:      "ibm_is_backup_policy.is_backup_policy",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIBMIsBackupPolicyConfigBasicWithScope(backupPolicyName, entCrn string) string {
	return fmt.Sprintf(`
	  resource "ibm_is_backup_policy" "is_backup_policy" {
		match_user_tags = ["dev:test"]
		name            = "%s"
		scope {
			crn = "%s"
		}
	}`, backupPolicyName, entCrn)
}
