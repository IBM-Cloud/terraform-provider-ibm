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

func TestAccIBMISImage_basic(t *testing.T) {
	var image string
	name := fmt.Sprintf("tfimg-name-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheckImage(t) },
		Providers:    testAccProviders,
		CheckDestroy: checkImageDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMISImageConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISImageExists("ibm_is_image.isExampleImage", image),
					resource.TestCheckResourceAttr(
						"ibm_is_image.isExampleImage", "name", name),
				),
			},
		},
	})
}

func TestAccIBMISImage_fromVolume(t *testing.T) {
	var image string
	vpcname := fmt.Sprintf("tf-vpc-%d", acctest.RandIntRange(10, 100))
	instanceName := fmt.Sprintf("tf-instnace-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tf-subnet-%d", acctest.RandIntRange(10, 100))
	publicKey := strings.TrimSpace(`
ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCKVmnMOlHKcZK8tpt3MP1lqOLAcqcJzhsvJcjscgVERRN7/9484SOBJ3HSKxxNG5JN8owAjy5f9yYwcUg+JaUVuytn5Pv3aeYROHGGg+5G346xaq3DAwX6Y5ykr2fvjObgncQBnuU5KHWCECO/4h8uWuwh/kfniXPVjFToc+gnkqA+3RKpAecZhFXwfalQ9mMuYGFxn+fwn8cYEApsJbsEmb0iJwPiZ5hjFC8wREuiTlhPHDgkBLOiycd20op2nXzDbHfCHInquEe/gYxEitALONxm0swBOwJZwlTDOB7C6y2dzlrtxr1L59m7pCkWI4EtTRLvleehBoj3u7jB4usR
`)
	sshname := fmt.Sprintf("tf-ssh-%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tfimg-name-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheckImage(t) },
		Providers:    testAccProviders,
		CheckDestroy: checkImageDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMISImageConfig1(vpcname, subnetname, sshname, publicKey, instanceName, name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISImageExists("ibm_is_image.isExampleImageFromVolume", image),
					resource.TestCheckResourceAttr(
						"ibm_is_image.isExampleImageFromVolume", "name", name),
				),
			},
		},
	})
}

func TestAccIBMISImage_encrypted(t *testing.T) {
	var image string
	name := fmt.Sprintf("tfimg-enc-name-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheckEncryptedImage(t) },
		Providers:    testAccProviders,
		CheckDestroy: checkImageDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMISImageEncryptedConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISImageExists("ibm_is_image.isExampleImageEncrypted", image),
					resource.TestCheckResourceAttr(
						"ibm_is_image.isExampleImageEncrypted", "name", name),
				),
			},
		},
	})
}
func checkImageDestroy(s *terraform.State) error {

	sess, _ := testAccProvider.Meta().(ClientSession).VpcV1API()
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_is_image" {
			continue
		}

		getimgoptions := &vpcv1.GetImageOptions{
			ID: &rs.Primary.ID,
		}
		_, _, err := sess.GetImage(getimgoptions)
		if err == nil {
			return fmt.Errorf("Image still exists: %s", rs.Primary.ID)
		}
	}

	return nil
}

func testAccCheckIBMISImageExists(n, image string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return errors.New("No Record ID is set")
		}

		sess, _ := testAccProvider.Meta().(ClientSession).VpcV1API()
		getimgoptions := &vpcv1.GetImageOptions{
			ID: &rs.Primary.ID,
		}
		foundImage, _, err := sess.GetImage(getimgoptions)
		if err != nil {
			return err
		}
		image = *foundImage.ID

		return nil
	}
}

func testAccCheckIBMISImageConfig(name string) string {
	return fmt.Sprintf(`
		resource "ibm_is_image" "isExampleImage" {
			href = "%s"
			name = "%s"
			operating_system = "%s"
		}
	`, image_cos_url, name, image_operating_system)
}
func testAccCheckIBMISImageConfig1(vpcname, subnetname, sshname, publicKey, instanceName, name string) string {
	return fmt.Sprintf(`
		  resource "ibm_is_vpc" "testacc_vpc" {
			name = "%s"
		  }
		  
		  resource "ibm_is_subnet" "testacc_subnet" {
			name            		 = "%s"
			vpc             		 = ibm_is_vpc.testacc_vpc.id
			zone            		 = "%s"
			total_ipv4_address_count = 32
		  }
		  
		  resource "ibm_is_ssh_key" "testacc_sshkey" {
			name       = "%s"
			public_key = "%s"
		  }
		  
		  resource "ibm_is_instance" "testacc_instance" {
			name    = "%s"
			image   = "%s"
			profile = "%s"
			primary_network_interface {
			  subnet     = ibm_is_subnet.testacc_subnet.id
			}
			vpc  = ibm_is_vpc.testacc_vpc.id
			zone = "%s"
			keys = [ibm_is_ssh_key.testacc_sshkey.id]
			network_interfaces {
			  subnet = ibm_is_subnet.testacc_subnet.id
			  name   = "eth1"
			}
		  }
		  
		  resource "ibm_is_image" "isExampleImageFromVolume" {
			name = "%s"
			source_volume = ibm_is_instance.testacc_instance.volume_attachments.0.volume_id
			timeouts {
				create = "45m"
			}
		  }`, vpcname, subnetname, ISZoneName, sshname, publicKey, instanceName, isImage, instanceProfileName, ISZoneName, name)
}

func testAccCheckIBMISImageEncryptedConfig(name string) string {
	return fmt.Sprintf(`
		resource "ibm_is_image" "isExampleImageEncrypted" {
			encrypted_data_key = "%s"
  			encryption_key = "%s"
			href = "%s"
			name = "%s"
			operating_system = "%s"
		}
		`, IsImageEncryptedDataKey, IsImageEncryptionKey, image_cos_url_encrypted, name, image_operating_system)
}
