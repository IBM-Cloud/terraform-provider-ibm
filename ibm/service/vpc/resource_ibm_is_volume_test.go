// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"errors"
	"fmt"
	"strings"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"

	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccIBMISVolume_basic(t *testing.T) {
	var vol string
	name := fmt.Sprintf("tf-vol-%d", acctest.RandIntRange(10, 100))
	name1 := fmt.Sprintf("tf-vol-upd-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISVolumeDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISVolumeConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISVolumeExists("ibm_is_volume.storage", vol),
					resource.TestCheckResourceAttr(
						"ibm_is_volume.storage", "name", name),
				),
			},

			{
				Config: testAccCheckIBMISVolumeConfig(name1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISVolumeExists("ibm_is_volume.storage", vol),
					resource.TestCheckResourceAttr(
						"ibm_is_volume.storage", "name", name1),
				),
			},
		},
	})
}

func TestAccIBMISVolumeUsertag_basic(t *testing.T) {
	var vol string
	name := fmt.Sprintf("tf-vol-%d", acctest.RandIntRange(10, 100))
	tagname := fmt.Sprintf("tfusertag%d", acctest.RandIntRange(10, 100))
	tagnameupdate := fmt.Sprintf("tfusertagupd%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISVolumeDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMISVolumeUsertagConfig(name, tagname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISVolumeExists("ibm_is_volume.storage", vol),
					resource.TestCheckResourceAttr(
						"ibm_is_volume.storage", "name", name),
					resource.TestCheckResourceAttrSet(
						"ibm_is_volume.storage", "tags.#"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_volume.storage", "tags.0"),
					resource.TestCheckResourceAttr(
						"ibm_is_volume.storage", "tags.0", tagname),
				),
			},

			resource.TestStep{
				Config: testAccCheckIBMISVolumeUsertagConfig(name, tagnameupdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISVolumeExists("ibm_is_volume.storage", vol),
					resource.TestCheckResourceAttrSet(
						"ibm_is_volume.storage", "tags.#"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_volume.storage", "tags.0"),
					resource.TestCheckResourceAttr(
						"ibm_is_volume.storage", "tags.0", tagnameupdate),
				),
			},
		},
	})
}

func TestAccIBMISVolumeUpdateCustom_basic(t *testing.T) {
	var vol string
	vpcname := fmt.Sprintf("tf-vpc-%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf-instnace-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tf-subnet-%d", acctest.RandIntRange(10, 100))
	publicKey := strings.TrimSpace(`
ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCKVmnMOlHKcZK8tpt3MP1lqOLAcqcJzhsvJcjscgVERRN7/9484SOBJ3HSKxxNG5JN8owAjy5f9yYwcUg+JaUVuytn5Pv3aeYROHGGg+5G346xaq3DAwX6Y5ykr2fvjObgncQBnuU5KHWCECO/4h8uWuwh/kfniXPVjFToc+gnkqA+3RKpAecZhFXwfalQ9mMuYGFxn+fwn8cYEApsJbsEmb0iJwPiZ5hjFC8wREuiTlhPHDgkBLOiycd20op2nXzDbHfCHInquEe/gYxEitALONxm0swBOwJZwlTDOB7C6y2dzlrtxr1L59m7pCkWI4EtTRLvleehBoj3u7jB4usR
`)
	sshname := fmt.Sprintf("tf-ssh-%d", acctest.RandIntRange(10, 100))
	volName := fmt.Sprintf("tf-vol-%d", acctest.RandIntRange(10, 100))
	iops1 := int64(600)
	iops2 := int64(900)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISVolumeDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISVolumeCustomConfig(vpcname, subnetname, sshname, publicKey, name, volName, iops1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISVolumeExists("ibm_is_volume.storage", vol),
					resource.TestCheckResourceAttr(
						"ibm_is_volume.storage", "name", volName),
					resource.TestCheckResourceAttr(
						"ibm_is_volume.storage", "iops", fmt.Sprintf("%d", iops1)),
				),
			},

			{
				Config: testAccCheckIBMISVolumeCustomConfig(vpcname, subnetname, sshname, publicKey, name, volName, iops2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISVolumeExists("ibm_is_volume.storage", vol),
					resource.TestCheckResourceAttr(
						"ibm_is_volume.storage", "name", volName),
					resource.TestCheckResourceAttr(
						"ibm_is_volume.storage", "iops", fmt.Sprintf("%d", iops2)),
				),
			},
		},
	})
}

func TestAccIBMISVolumeUpdateTier_basic(t *testing.T) {
	var vol string
	vpcname := fmt.Sprintf("tf-vpc-%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf-instnace-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tf-subnet-%d", acctest.RandIntRange(10, 100))
	publicKey := strings.TrimSpace(`
ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCKVmnMOlHKcZK8tpt3MP1lqOLAcqcJzhsvJcjscgVERRN7/9484SOBJ3HSKxxNG5JN8owAjy5f9yYwcUg+JaUVuytn5Pv3aeYROHGGg+5G346xaq3DAwX6Y5ykr2fvjObgncQBnuU5KHWCECO/4h8uWuwh/kfniXPVjFToc+gnkqA+3RKpAecZhFXwfalQ9mMuYGFxn+fwn8cYEApsJbsEmb0iJwPiZ5hjFC8wREuiTlhPHDgkBLOiycd20op2nXzDbHfCHInquEe/gYxEitALONxm0swBOwJZwlTDOB7C6y2dzlrtxr1L59m7pCkWI4EtTRLvleehBoj3u7jB4usR
`)
	sshname := fmt.Sprintf("tf-ssh-%d", acctest.RandIntRange(10, 100))
	volName := fmt.Sprintf("tf-vol-%d", acctest.RandIntRange(10, 100))
	profileName1 := "general-purpose"
	profileName2 := "5iops-tier"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISVolumeDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISVolumeTierConfig(vpcname, subnetname, sshname, publicKey, name, volName, profileName1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISVolumeExists("ibm_is_volume.storage", vol),
					resource.TestCheckResourceAttr(
						"ibm_is_volume.storage", "name", volName),
					resource.TestCheckResourceAttr(
						"ibm_is_volume.storage", "profile", profileName1),
				),
			},

			{
				Config: testAccCheckIBMISVolumeTierConfig(vpcname, subnetname, sshname, publicKey, name, volName, profileName2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISVolumeExists("ibm_is_volume.storage", vol),
					resource.TestCheckResourceAttr(
						"ibm_is_volume.storage", "name", volName),
					resource.TestCheckResourceAttr(
						"ibm_is_volume.storage", "profile", profileName2),
				),
			},
		},
	})
}

func TestAccIBMISVolumeUpdateCapacity_basic(t *testing.T) {
	var vol string
	vpcname := fmt.Sprintf("tf-vpc-%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf-instnace-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tf-subnet-%d", acctest.RandIntRange(10, 100))
	publicKey := strings.TrimSpace(`
ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCKVmnMOlHKcZK8tpt3MP1lqOLAcqcJzhsvJcjscgVERRN7/9484SOBJ3HSKxxNG5JN8owAjy5f9yYwcUg+JaUVuytn5Pv3aeYROHGGg+5G346xaq3DAwX6Y5ykr2fvjObgncQBnuU5KHWCECO/4h8uWuwh/kfniXPVjFToc+gnkqA+3RKpAecZhFXwfalQ9mMuYGFxn+fwn8cYEApsJbsEmb0iJwPiZ5hjFC8wREuiTlhPHDgkBLOiycd20op2nXzDbHfCHInquEe/gYxEitALONxm0swBOwJZwlTDOB7C6y2dzlrtxr1L59m7pCkWI4EtTRLvleehBoj3u7jB4usR
`)
	sshname := fmt.Sprintf("tf-ssh-%d", acctest.RandIntRange(10, 100))
	volName := fmt.Sprintf("tf-vol-%d", acctest.RandIntRange(10, 100))
	capacity1 := int64(100)
	capacity2 := int64(120)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISVolumeDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISVolumeCapacityConfig(vpcname, subnetname, sshname, publicKey, name, volName, capacity1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISVolumeExists("ibm_is_volume.storage", vol),
					resource.TestCheckResourceAttr(
						"ibm_is_volume.storage", "name", volName),
					resource.TestCheckResourceAttr(
						"ibm_is_volume.storage", "capacity", fmt.Sprintf("%d", capacity1)),
				),
			},

			{
				Config: testAccCheckIBMISVolumeCapacityConfig(vpcname, subnetname, sshname, publicKey, name, volName, capacity2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISVolumeExists("ibm_is_volume.storage", vol),
					resource.TestCheckResourceAttr(
						"ibm_is_volume.storage", "name", volName),
					resource.TestCheckResourceAttr(
						"ibm_is_volume.storage", "capacity", fmt.Sprintf("%d", capacity2)),
				),
			},
		},
	})
}

func TestAccIBMISVolumeAttachmentDelete_basic(t *testing.T) {
	var vol string
	insname := fmt.Sprintf("tf-ins-%d", acctest.RandIntRange(10, 100))
	initialVolumeCapacityArray := fmt.Sprintf("[%d, %d]", 10, 20)
	finalVolumeCapacityArray := fmt.Sprintf("[%d]", 10)
	vpcname := fmt.Sprintf("tf-vpc-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tf-subnet-%d", acctest.RandIntRange(10, 100))
	publicKey := strings.TrimSpace(`
ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCKVmnMOlHKcZK8tpt3MP1lqOLAcqcJzhsvJcjscgVERRN7/9484SOBJ3HSKxxNG5JN8owAjy5f9yYwcUg+JaUVuytn5Pv3aeYROHGGg+5G346xaq3DAwX6Y5ykr2fvjObgncQBnuU5KHWCECO/4h8uWuwh/kfniXPVjFToc+gnkqA+3RKpAecZhFXwfalQ9mMuYGFxn+fwn8cYEApsJbsEmb0iJwPiZ5hjFC8wREuiTlhPHDgkBLOiycd20op2nXzDbHfCHInquEe/gYxEitALONxm0swBOwJZwlTDOB7C6y2dzlrtxr1L59m7pCkWI4EtTRLvleehBoj3u7jB4usR
`)
	sshname := fmt.Sprintf("tf-ssh-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISVolumeDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISVolumeAttachmentDeleteConfig(vpcname, subnetname, sshname, publicKey, insname, initialVolumeCapacityArray),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISVolumeExists("ibm_is_volume.storage.0", vol),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "name", insname),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "volume_attachments.#", "3"),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "volumes.#", "2"),
				),
			},

			{
				Config: testAccCheckIBMISVolumeAttachmentDeleteConfig(vpcname, subnetname, sshname, publicKey, insname, finalVolumeCapacityArray),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISVolumeExists("ibm_is_volume.storage.0", vol),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "name", insname),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "volume_attachments.#", "2"),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "volumes.#", "1"),
				),
			},
		},
	})
}

func testAccCheckIBMISVolumeDestroy(s *terraform.State) error {

	sess, _ := acc.TestAccProvider.Meta().(conns.ClientSession).VpcV1API()
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_is_vol" {
			continue
		}

		getvolumeoptions := &vpcv1.GetVolumeOptions{
			ID: &rs.Primary.ID,
		}
		_, _, err := sess.GetVolume(getvolumeoptions)

		if err == nil {
			return fmt.Errorf("[ERROR] Volume still exists: %s", rs.Primary.ID)
		}
	}

	return nil
}

func testAccCheckIBMISVolumeExists(n, volID string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return errors.New("No Record ID is set")
		}

		sess, _ := acc.TestAccProvider.Meta().(conns.ClientSession).VpcV1API()
		getvolumeoptions := &vpcv1.GetVolumeOptions{
			ID: &rs.Primary.ID,
		}
		foundvol, _, err := sess.GetVolume(getvolumeoptions)
		if err != nil {
			return err
		}
		volID = *foundvol.ID
		return nil
	}
}

func testAccCheckIBMISVolumeConfig(name string) string {
	return fmt.Sprintf(
		`
	resource "ibm_is_volume" "storage"{
		name = "%s"
		profile = "10iops-tier"
		zone = "us-south-1"
		# capacity= 200
	}
`, name)

}

func testAccCheckIBMISVolumeCustomConfig(vpcname, subnetname, sshname, publicKey, name, volName string, iops int64) string {
	return fmt.Sprintf(
		`
		resource "ibm_is_vpc" "testacc_vpc" {
			name = "%s"
		}
		  
		resource "ibm_is_subnet" "testacc_subnet" {
			name            			= "%s"
			vpc             			= ibm_is_vpc.testacc_vpc.id
			zone            			= "%s"
			total_ipv4_address_count 	= 16
		}
		  
		resource "ibm_is_ssh_key" "testacc_sshkey" {
			name       = "%s"
			public_key = "%s"
		}
		resource "ibm_is_volume" "storage"{
			name 	= "%s"
			profile = "custom"
			zone 	= "%s"
			iops 	= %d
		}		  
		resource "ibm_is_instance" "testacc_instance" {
			name    = "%s"
			image   = "%s"
			profile = "%s"
			volumes = [ibm_is_volume.storage.id]
			primary_network_interface {
				subnet     = ibm_is_subnet.testacc_subnet.id
			}
			vpc  = ibm_is_vpc.testacc_vpc.id
			zone = "%s"
			keys = [ibm_is_ssh_key.testacc_sshkey.id]
		}	

`, vpcname, subnetname, acc.ISZoneName, sshname, publicKey, volName, acc.ISZoneName, iops, name, acc.IsImage, acc.InstanceProfileName, acc.ISZoneName)

}

func testAccCheckIBMISVolumeTierConfig(vpcname, subnetname, sshname, publicKey, name, volName, profileName string) string {
	return fmt.Sprintf(
		`
		resource "ibm_is_vpc" "testacc_vpc" {
			name = "%s"
		}
		  
		resource "ibm_is_subnet" "testacc_subnet" {
			name            			= "%s"
			vpc             			= ibm_is_vpc.testacc_vpc.id
			zone            			= "%s"
			total_ipv4_address_count 	= 16
		}
		  
		resource "ibm_is_ssh_key" "testacc_sshkey" {
			name       = "%s"
			public_key = "%s"
		}
		resource "ibm_is_volume" "storage"{
			name 	= "%s"
			profile = "%s"
			zone 	= "%s"
		}		  
		resource "ibm_is_instance" "testacc_instance" {
			name    = "%s"
			image   = "%s"
			profile = "%s"
			volumes = [ibm_is_volume.storage.id]
			primary_network_interface {
				subnet     = ibm_is_subnet.testacc_subnet.id
			}
			vpc  = ibm_is_vpc.testacc_vpc.id
			zone = "%s"
			keys = [ibm_is_ssh_key.testacc_sshkey.id]
		}	

`, vpcname, subnetname, acc.ISZoneName, sshname, publicKey, volName, profileName, acc.ISZoneName, name, acc.IsImage, acc.InstanceProfileName, acc.ISZoneName)

}

func testAccCheckIBMISVolumeAttachmentDeleteConfig(vpcname, subnetname, sshname, publicKey, insname, capacityArray string) string {
	return fmt.Sprintf(
		`
		variable "vsi_vol_size" {
			description = "capacity array"
			default     =  %s
		}

		resource "ibm_is_volume" "storage"{
			name 	 = "tf-vol-att-${count.index}"
			count 	 = length(var.vsi_vol_size)
			profile  = "general-purpose"
			zone 	 = "%s"
			capacity = var.vsi_vol_size[count.index]
		}

		resource "ibm_is_vpc" "testacc_vpc" {
			name = "%s"
		}
		  
		resource "ibm_is_subnet" "testacc_subnet" {
			name            			= "%s"
			vpc             			= ibm_is_vpc.testacc_vpc.id
			zone            			= "%s"
			total_ipv4_address_count 	= 16
		}
		  
		resource "ibm_is_ssh_key" "testacc_sshkey" {
			name       					= "%s"
			public_key 					= "%s"
		}
		  
		resource "ibm_is_instance" "testacc_instance" {
			name    		= "%s"
			image   		= "%s"
			profile 		= "%s"
			volumes = ibm_is_volume.storage[*].id
			primary_network_interface {
				subnet     = ibm_is_subnet.testacc_subnet.id
			}
			vpc  = ibm_is_vpc.testacc_vpc.id
			zone = "%s"
			keys = [ibm_is_ssh_key.testacc_sshkey.id]
		}
`, capacityArray, acc.ISZoneName, vpcname, subnetname, acc.ISZoneName, sshname, publicKey, insname, acc.IsImage, acc.InstanceProfileName, acc.ISZoneName)

}

func testAccCheckIBMISVolumeCapacityConfig(vpcname, subnetname, sshname, publicKey, name, volName string, capacity int64) string {
	return fmt.Sprintf(
		`
		resource "ibm_is_vpc" "testacc_vpc" {
			name = "%s"
		}
		  
		resource "ibm_is_subnet" "testacc_subnet" {
			name            			= "%s"
			vpc             			= ibm_is_vpc.testacc_vpc.id
			zone            			= "%s"
			total_ipv4_address_count 	= 16
		}
		  
		resource "ibm_is_ssh_key" "testacc_sshkey" {
			name       = "%s"
			public_key = "%s"
		}
		resource "ibm_is_volume" "storage"{
			name 		= "%s"
			profile 	= "10iops-tier"
			zone 		= "%s"
			capacity 	= %d
		}		  
		resource "ibm_is_instance" "testacc_instance" {
			name    = "%s"
			image   = "%s"
			profile = "%s"
			volumes = [ibm_is_volume.storage.id]
			primary_network_interface {
				subnet     = ibm_is_subnet.testacc_subnet.id
			}
			vpc  = ibm_is_vpc.testacc_vpc.id
			zone = "%s"
			keys = [ibm_is_ssh_key.testacc_sshkey.id]
		}	

`, vpcname, subnetname, acc.ISZoneName, sshname, publicKey, volName, acc.ISZoneName, capacity, name, acc.IsImage, acc.InstanceProfileName, acc.ISZoneName)

}

func testAccCheckIBMISVolumeUsertagConfig(name, usertag string) string {
	return fmt.Sprintf(
		`
    resource "ibm_is_volume" "storage"{
        name = "%s"
        profile = "10iops-tier"
        zone = "us-south-1"
        # capacity= 200
        tags = ["%s"]
    }
`, name, usertag)

}
