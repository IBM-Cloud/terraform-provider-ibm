// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"errors"
	"fmt"
	"strings"
	"testing"

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
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMISVolumeDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMISVolumeConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISVolumeExists("ibm_is_volume.storage", vol),
					resource.TestCheckResourceAttr(
						"ibm_is_volume.storage", "name", name),
				),
			},

			resource.TestStep{
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
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMISVolumeDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
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

			resource.TestStep{
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

	sess, _ := testAccProvider.Meta().(ClientSession).VpcV1API()
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_is_vol" {
			continue
		}

		getvolumeoptions := &vpcv1.GetVolumeOptions{
			ID: &rs.Primary.ID,
		}
		_, _, err := sess.GetVolume(getvolumeoptions)

		if err == nil {
			return fmt.Errorf("Volume still exists: %s", rs.Primary.ID)
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

		sess, _ := testAccProvider.Meta().(ClientSession).VpcV1API()
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
		zone = "us-south-3"
		# capacity= 200
	}
`, name)

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
`, capacityArray, ISZoneName, vpcname, subnetname, ISZoneName, sshname, publicKey, insname, isImage, instanceProfileName, ISZoneName)

}
