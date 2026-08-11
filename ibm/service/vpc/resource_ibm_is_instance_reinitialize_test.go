// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"fmt"
	"strings"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMISInstanceReinitialize_basic(t *testing.T) {
	vpcname := fmt.Sprintf("tf-vpc-%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf-server-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tfip-subnet-%d", acctest.RandIntRange(10, 100))
	publicKey := strings.TrimSpace(`
ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCKVmnMOlHKcZK8tpt3MP1lqOLAcqcJzhsvJcjscgVERRN7/9484SOBJ3HSKxxNG5JN8owAjy5f9yYwcUg+JaUVuytn5Pv3aeYROHGGg+5G346xaq3DAwX6Y5ykr2fvjObgncQBnuU5KHWCECO/4h8uWuwh/kfniXPVjFToc+gnkqA+3RKpAecZhFXwfalQ9mMuYGFxn+fwn8cYEApsJbsEmb0iJwPiZ5hjFC8wREuiTlhPHDgkBLOiycd20op2nXzDbHfCHInquEe/gYxEitALONxm0swBOwJZwlTDOB7C6y2dzlrtxr1L59m7pCkWI4EtTRLvleehBoj3u7jB4usR
`)
	sshname := fmt.Sprintf("tf-sshname-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISInstanceReinitializeByImageConfig(vpcname, subnetname, sshname, publicKey, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_is_instance_reinitialize.test_reinit", "instance_id", name),
					resource.TestCheckResourceAttrSet(
						"ibm_is_instance_reinitialize.test_reinit", "status"),
				),
			},
		},
	})
}

func TestAccIBMISInstanceReinitializeWithVolumeAttachment(t *testing.T) {
	vpcname := fmt.Sprintf("tf-vpc-%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf-server-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tfip-subnet-%d", acctest.RandIntRange(10, 100))
	volumeName := fmt.Sprintf("tf-volume-%d", acctest.RandIntRange(10, 100))
	publicKey := strings.TrimSpace(`
ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCKVmnMOlHKcZK8tpt3MP1lqOLAcqcJzhsvJcjscgVERRN7/9484SOBJ3HSKxxNG5JN8owAjy5f9yYwcUg+JaUVuytn5Pv3aeYROHGGg+5G346xaq3DAwX6Y5ykr2fvjObgncQBnuU5KHWCECO/4h8uWuwh/kfniXPVjFToc+gnkqA+3RKpAecZhFXwfalQ9mMuYGFxn+fwn8cYEApsJbsEmb0iJwPiZ5hjFC8wREuiTlhPHDgkBLOiycd20op2nXzDbHfCHInquEe/gYxEitALONxm0swBOwJZwlTDOB7C6y2dzlrtxr1L59m7pCkWI4EtTRLvleehBoj3u7jB4usR
`)
	sshname := fmt.Sprintf("tf-sshname-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISInstanceReinitializeByVolumeAttachmentConfig(vpcname, subnetname, sshname, publicKey, name, volumeName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_is_instance_reinitialize.test_reinit_volume", "instance_id", name),
					resource.TestCheckResourceAttrSet(
						"ibm_is_instance_reinitialize.test_reinit_volume", "status"),
				),
			},
		},
	})
}

func TestAccIBMISInstanceReinitializeWithKeysAndUserData(t *testing.T) {
	vpcname := fmt.Sprintf("tf-vpc-%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf-server-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tfip-subnet-%d", acctest.RandIntRange(10, 100))
	publicKey := strings.TrimSpace(`
ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCKVmnMOlHKcZK8tpt3MP1lqOLAcqcJzhsvJcjscgVERRN7/9484SOBJ3HSKxxNG5JN8owAjy5f9yYwcUg+JaUVuytn5Pv3aeYROHGGg+5G346xaq3DAwX6Y5ykr2fvjObgncQBnuU5KHWCECO/4h8uWuwh/kfniXPVjFToc+gnkqA+3RKpAecZhFXwfalQ9mMuYGFxn+fwn8cYEApsJbsEmb0iJwPiZ5hjFC8wREuiTlhPHDgkBLOiycd20op2nXzDbHfCHInquEe/gYxEitALONxm0swBOwJZwlTDOB7C6y2dzlrtxr1L59m7pCkWI4EtTRLvleehBoj3u7jB4usR
`)
	sshname := fmt.Sprintf("tf-sshname-%d", acctest.RandIntRange(10, 100))
	sshname2 := fmt.Sprintf("tf-sshname2-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISInstanceReinitializeWithKeysAndUserDataConfig(vpcname, subnetname, sshname, sshname2, publicKey, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_is_instance_reinitialize.test_reinit_keys", "instance_id", name),
					resource.TestCheckResourceAttrSet(
						"ibm_is_instance_reinitialize.test_reinit_keys", "status"),
				),
			},
		},
	})
}

func TestAccIBMISInstanceReinitializeWithTrustedProfile(t *testing.T) {
	vpcname := fmt.Sprintf("tf-vpc-%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf-server-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tfip-subnet-%d", acctest.RandIntRange(10, 100))
	profileName := fmt.Sprintf("tf-trusted-profile-%d", acctest.RandIntRange(10, 100))
	publicKey := strings.TrimSpace(`
ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCKVmnMOlHKcZK8tpt3MP1lqOLAcqcJzhsvJcjscgVERRN7/9484SOBJ3HSKxxNG5JN8owAjy5f9yYwcUg+JaUVuytn5Pv3aeYROHGGg+5G346xaq3DAwX6Y5ykr2fvjObgncQBnuU5KHWCECO/4h8uWuwh/kfniXPVjFToc+gnkqA+3RKpAecZhFXwfalQ9mMuYGFxn+fwn8cYEApsJbsEmb0iJwPiZ5hjFC8wREuiTlhPHDgkBLOiycd20op2nXzDbHfCHInquEe/gYxEitALONxm0swBOwJZwlTDOB7C6y2dzlrtxr1L59m7pCkWI4EtTRLvleehBoj3u7jB4usR
`)
	sshname := fmt.Sprintf("tf-sshname-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISInstanceReinitializeWithTrustedProfileConfig(vpcname, subnetname, sshname, publicKey, name, profileName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_is_instance_reinitialize.test_reinit_profile", "instance_id", name),
					resource.TestCheckResourceAttrSet(
						"ibm_is_instance_reinitialize.test_reinit_profile", "status"),
				),
			},
		},
	})
}

func TestAccIBMISInstanceReinitializeImport(t *testing.T) {
	vpcname := fmt.Sprintf("tf-vpc-%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf-server-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tfip-subnet-%d", acctest.RandIntRange(10, 100))
	publicKey := strings.TrimSpace(`
ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCKVmnMOlHKcZK8tpt3MP1lqOLAcqcJzhsvJcjscgVERRN7/9484SOBJ3HSKxxNG5JN8owAjy5f9yYwcUg+JaUVuytn5Pv3aeYROHGGg+5G346xaq3DAwX6Y5ykr2fvjObgncQBnuU5KHWCECO/4h8uWuwh/kfniXPVjFToc+gnkqA+3RKpAecZhFXwfalQ9mMuYGFxn+fwn8cYEApsJbsEmb0iJwPiZ5hjFC8wREuiTlhPHDgkBLOiycd20op2nXzDbHfCHInquEe/gYxEitALONxm0swBOwJZwlTDOB7C6y2dzlrtxr1L59m7pCkWI4EtTRLvleehBoj3u7jB4usR
`)
	sshname := fmt.Sprintf("tf-sshname-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISInstanceReinitializeByImageConfig(vpcname, subnetname, sshname, publicKey, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_is_instance_reinitialize.test_reinit", "instance_id", name),
				),
			},
			{
				ResourceName:      "ibm_is_instance_reinitialize.test_reinit",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

// Test configuration for reinitialization by image
func testAccCheckIBMISInstanceReinitializeByImageConfig(vpcname, subnetname, sshname, publicKey, name string) string {
	return fmt.Sprintf(`
	data "ibm_is_images" "im_images" {
  
	}
	
	data "ibm_is_images" "ubuntu_images" {
		visibility = "public"
		operating_system = "ubuntu-20-04-amd64"
	}
	
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
		public_key = "%s"
	}
	
	resource "ibm_is_instance" "testacc_instance" {
		name    = "%s"
		image   = data.ibm_is_images.ubuntu_images.images[0].id
		profile = "bx2d-2x8"
		primary_network_interface {
			subnet = ibm_is_subnet.testacc_subnet.id
		}
		vpc  = ibm_is_vpc.testacc_vpc.id
		zone = "%s"
		keys = [ibm_is_ssh_key.testacc_sshkey.id]
	}
	
	resource "ibm_is_instance_reinitialize" "test_reinit" {
		depends_on = [ibm_is_instance.testacc_instance]
		instance_id = ibm_is_instance.testacc_instance.id
		image = data.ibm_is_images.im_images.images[4].id
	}
	`, vpcname, subnetname, acc.ISZoneName, acc.ISCIDR, sshname, publicKey, name, acc.ISZoneName)
}

// Test configuration for reinitialization by volume attachment
func testAccCheckIBMISInstanceReinitializeByVolumeAttachmentConfig(vpcname, subnetname, sshname, publicKey, name, volumeName string) string {
	return fmt.Sprintf(`
	data "ibm_is_images" "ubuntu_images" {
		visibility = "public"
		operating_system = "ubuntu-20-04-amd64"
	}
	
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
		public_key = "%s"
	}
	
	resource "ibm_is_volume" "testacc_volume" {
		name       = "%s"
		profile    = "general-purpose"
		zone       = "%s"
		capacity   = 100
	}
	
	resource "ibm_is_instance" "testacc_instance" {
		name    = "%s"
		image   = data.ibm_is_images.ubuntu_images.images[0].id
		profile = "bx2d-2x8"
		primary_network_interface {
			subnet = ibm_is_subnet.testacc_subnet.id
		}
		vpc  = ibm_is_vpc.testacc_vpc.id
		zone = "%s"
		keys = [ibm_is_ssh_key.testacc_sshkey.id]
	}
	
	resource "ibm_is_instance_reinitialize" "test_reinit_volume" {
		depends_on = [ibm_is_instance.testacc_instance, ibm_is_volume.testacc_volume]
		instance_id = ibm_is_instance.testacc_instance.id
		boot_volume_attachment {
			name = "reinit-boot-vol"
			volume {
				id = ibm_is_volume.testacc_volume.id
				name = "reinit-volume"
			}
			delete_volume_on_instance_delete = false
		}
	}
	`, vpcname, subnetname, acc.ISZoneName, acc.ISCIDR, sshname, publicKey, volumeName, acc.ISZoneName, name, acc.ISZoneName)
}

// Test configuration for reinitialization with custom keys and user data
func testAccCheckIBMISInstanceReinitializeWithKeysAndUserDataConfig(vpcname, subnetname, sshname, sshname2, publicKey, name string) string {
	return fmt.Sprintf(`
	data "ibm_is_images" "ubuntu_images" {
		visibility = "public"
		operating_system = "ubuntu-20-04-amd64"
	}
	
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
		public_key = "%s"
	}
	
	resource "ibm_is_ssh_key" "testacc_sshkey2" {
		name       = "%s"
		public_key = "%s"
	}
	
	resource "ibm_is_instance" "testacc_instance" {
		name    = "%s"
		image   = data.ibm_is_images.ubuntu_images.images[0].id
		profile = "bx2d-2x8"
		primary_network_interface {
			subnet = ibm_is_subnet.testacc_subnet.id
		}
		vpc  = ibm_is_vpc.testacc_vpc.id
		zone = "%s"
		keys = [ibm_is_ssh_key.testacc_sshkey.id]
	}
	
	resource "ibm_is_instance_reinitialize" "test_reinit_keys" {
		depends_on = [ibm_is_instance.testacc_instance]
		instance_id = ibm_is_instance.testacc_instance.id
		image = data.ibm_is_images.ubuntu_images.images[0].id
		keys = [ibm_is_ssh_key.testacc_sshkey.id, ibm_is_ssh_key.testacc_sshkey2.id]
		user_data = <<-EOF
			#!/bin/bash
			echo "Reinitialized instance" > /tmp/reinit.log
		EOF
	}
	`, vpcname, subnetname, acc.ISZoneName, acc.ISCIDR, sshname, publicKey, sshname2, publicKey, name, acc.ISZoneName)
}

// Test configuration for reinitialization with trusted profile
func testAccCheckIBMISInstanceReinitializeWithTrustedProfileConfig(vpcname, subnetname, sshname, publicKey, name, profileName string) string {
	return fmt.Sprintf(`
	data "ibm_is_images" "ubuntu_images" {
		visibility = "public"
		operating_system = "ubuntu-20-04-amd64"
	}
	
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
		public_key = "%s"
	}
	
	resource "ibm_iam_trusted_profile" "testacc_profile" {
		name = "%s"
	}
	
	resource "ibm_is_instance" "testacc_instance" {
		name    = "%s"
		image   = data.ibm_is_images.ubuntu_images.images[0].id
		profile = "bx2d-2x8"
		primary_network_interface {
			subnet = ibm_is_subnet.testacc_subnet.id
		}
		vpc  = ibm_is_vpc.testacc_vpc.id
		zone = "%s"
		keys = [ibm_is_ssh_key.testacc_sshkey.id]
	}
	
	resource "ibm_is_instance_reinitialize" "test_reinit_profile" {
		depends_on = [ibm_is_instance.testacc_instance, ibm_iam_trusted_profile.testacc_profile]
		instance_id = ibm_is_instance.testacc_instance.id
		image = data.ibm_is_images.ubuntu_images.images[0].id
		default_trusted_profile {
			auto_link = "true"
			target {
				id = ibm_iam_trusted_profile.testacc_profile.id
			}
		}
	}
	`, vpcname, subnetname, acc.ISZoneName, acc.ISCIDR, sshname, publicKey, profileName, name, acc.ISZoneName)
}

// Test configuration for reinitialization by snapshot (if snapshots are available)
func testAccCheckIBMISInstanceReinitializeBySnapshotConfig(vpcname, subnetname, sshname, publicKey, name, snapshotId string) string {
	return fmt.Sprintf(`
	data "ibm_is_images" "ubuntu_images" {
		visibility = "public"
		operating_system = "ubuntu-20-04-amd64"
	}
	
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
		public_key = "%s"
	}
	
	resource "ibm_is_instance" "testacc_instance" {
		name    = "%s"
		image   = data.ibm_is_images.ubuntu_images.images[0].id
		profile = "bx2d-2x8"
		primary_network_interface {
			subnet = ibm_is_subnet.testacc_subnet.id
		}
		vpc  = ibm_is_vpc.testacc_vpc.id
		zone = "%s"
		keys = [ibm_is_ssh_key.testacc_sshkey.id]
	}
	
	resource "ibm_is_instance_reinitialize" "test_reinit_snapshot" {
		depends_on = [ibm_is_instance.testacc_instance]
		instance_id = ibm_is_instance.testacc_instance.id
		boot_volume_attachment {
			volume {
				source_snapshot {
					id = "%s"
				}
			}
		}
	}
	`, vpcname, subnetname, acc.ISZoneName, acc.ISCIDR, sshname, publicKey, name, acc.ISZoneName, snapshotId)
}
