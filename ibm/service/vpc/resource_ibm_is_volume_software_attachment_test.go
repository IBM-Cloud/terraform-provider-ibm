// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/vpc-go-sdk/vpcv1"
)

func TestAccIBMIsVolumeSoftwareAttachmentBasic(t *testing.T) {
	var conf vpcv1.VolumeSoftwareAttachment
	vpcname := fmt.Sprintf("tf-vpc-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tf-subnet-%d", acctest.RandIntRange(10, 100))
	sshname := fmt.Sprintf("tf-ssh-%d", acctest.RandIntRange(10, 100))
	instanceName := fmt.Sprintf("tf-instance-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIsVolumeSoftwareAttachmentDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsVolumeSoftwareAttachmentConfigBasic(vpcname, subnetname, sshname, instanceName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIsVolumeSoftwareAttachmentExists("ibm_is_volume_software_attachment.is_volume_software_attachment_instance", conf),
					resource.TestCheckResourceAttrPair("ibm_is_volume_software_attachment.is_volume_software_attachment_instance", "volume_id", "ibm_is_instance.test_instance", "boot_volume.0.volume_id"),
					resource.TestCheckResourceAttrPair("ibm_is_volume_software_attachment.is_volume_software_attachment_instance", "volume_software_attachment_id", "data.ibm_is_volume.test_volume", "software_attachments.0.id"),
					resource.TestCheckResourceAttrSet("ibm_is_volume_software_attachment.is_volume_software_attachment_instance", "name"),
					resource.TestCheckResourceAttrSet("ibm_is_volume_software_attachment.is_volume_software_attachment_instance", "created_at"),
					resource.TestCheckResourceAttrSet("ibm_is_volume_software_attachment.is_volume_software_attachment_instance", "href"),
					resource.TestCheckResourceAttr("ibm_is_volume_software_attachment.is_volume_software_attachment_instance", "resource_type", "volume_software_attachment"),
				),
			},
		},
	})
}

func TestAccIBMIsVolumeSoftwareAttachmentAllArgs(t *testing.T) {
	var conf vpcv1.VolumeSoftwareAttachment
	vpcname := fmt.Sprintf("tf-vpc-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tf-subnet-%d", acctest.RandIntRange(10, 100))
	sshname := fmt.Sprintf("tf-ssh-%d", acctest.RandIntRange(10, 100))
	instanceName := fmt.Sprintf("tf-instance-%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf-name-%d", acctest.RandIntRange(10, 100))
	nameUpdate := fmt.Sprintf("tf-name-upd-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIsVolumeSoftwareAttachmentDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsVolumeSoftwareAttachmentConfig(vpcname, subnetname, sshname, instanceName, name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIsVolumeSoftwareAttachmentExists("ibm_is_volume_software_attachment.is_volume_software_attachment_instance", conf),
					resource.TestCheckResourceAttrPair("ibm_is_volume_software_attachment.is_volume_software_attachment_instance", "volume_id", "ibm_is_instance.test_instance", "boot_volume.0.volume_id"),
					resource.TestCheckResourceAttr("ibm_is_volume_software_attachment.is_volume_software_attachment_instance", "name", name),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMIsVolumeSoftwareAttachmentConfig(vpcname, subnetname, sshname, instanceName, nameUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_is_volume_software_attachment.is_volume_software_attachment_instance", "name", nameUpdate),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_is_volume_software_attachment.is_volume_software_attachment_instance",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIBMIsVolumeSoftwareAttachmentConfigBasic(vpcname, subnetname, sshname, instanceName string) string {
	publicKey := strings.TrimSpace(`
ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCKVmnMOlHKcZK8tpt3MP1lqOLAcqcJzhsvJcjscgVERRN7/9484SOBJ3HSKxxNG5JN8owAjy5f9yYwcUg+JaUVuytn5Pv3aeYROHGGg+5G346xaq3DAwX6Y5ykr2fvjObgncQBnuU5KHWCECO/4h8uWuwh/kfniXPVjFToc+gnkqA+3RKpAecZhFXwfalQ9mMuYGFxn+fwn8cYEApsJbsEmb0iJwPiZ5hjFC8wREuiTlhPHDgkBLOiycd20op2nXzDbHfCHInquEe/gYxEitALONxm0swBOwJZwlTDOB7C6y2dzlrtxr1L59m7pCkWI4EtTRLvleehBoj3u7jB4usR
`)
	return fmt.Sprintf(`
		resource "ibm_is_vpc" "test_vpc" {
			name = "%s"
		}

		resource "ibm_is_subnet" "test_subnet" {
			name                     = "%s"
			vpc                      = ibm_is_vpc.test_vpc.id
			zone                     = "%s"
			total_ipv4_address_count = 64
		}

		resource "ibm_is_ssh_key" "test_key" {
			name       = "%s"
			public_key = "%s"
		}

		resource "ibm_is_instance" "test_instance" {
			name    = "%s"
			profile = "%s"
			catalog_offering {
				version_crn = "%s"
				plan_crn    = "%s"
			}
			vpc  = ibm_is_vpc.test_vpc.id
			zone = ibm_is_subnet.test_subnet.zone
			keys = [ibm_is_ssh_key.test_key.id]

			primary_network_attachment {
				virtual_network_interface {
					subnet = ibm_is_subnet.test_subnet.id
				}
			}
		}

		data "ibm_is_volume" "test_volume" {
			identifier = ibm_is_instance.test_instance.boot_volume[0].volume_id
		}

		resource "ibm_is_volume_software_attachment" "is_volume_software_attachment_instance" {
			volume_id                     = ibm_is_instance.test_instance.boot_volume[0].volume_id
			volume_software_attachment_id = data.ibm_is_volume.test_volume.software_attachments.0.id
		}
	`, vpcname, subnetname, acc.ISZoneName, sshname, publicKey, instanceName, acc.InstanceProfileName, acc.ISCatalogImageOfferingCRN, acc.ISCatalogImagePlanCRN)
}

func testAccCheckIBMIsVolumeSoftwareAttachmentConfig(vpcname, subnetname, sshname, instanceName, name string) string {
	publicKey := strings.TrimSpace(`
ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCKVmnMOlHKcZK8tpt3MP1lqOLAcqcJzhsvJcjscgVERRN7/9484SOBJ3HSKxxNG5JN8owAjy5f9yYwcUg+JaUVuytn5Pv3aeYROHGGg+5G346xaq3DAwX6Y5ykr2fvjObgncQBnuU5KHWCECO/4h8uWuwh/kfniXPVjFToc+gnkqA+3RKpAecZhFXwfalQ9mMuYGFxn+fwn8cYEApsJbsEmb0iJwPiZ5hjFC8wREuiTlhPHDgkBLOiycd20op2nXzDbHfCHInquEe/gYxEitALONxm0swBOwJZwlTDOB7C6y2dzlrtxr1L59m7pCkWI4EtTRLvleehBoj3u7jB4usR
`)
	return fmt.Sprintf(`
		resource "ibm_is_vpc" "test_vpc" {
			name = "%s"
		}

		resource "ibm_is_subnet" "test_subnet" {
			name                     = "%s"
			vpc                      = ibm_is_vpc.test_vpc.id
			zone                     = "%s"
			total_ipv4_address_count = 64
		}

		resource "ibm_is_ssh_key" "test_key" {
			name       = "%s"
			public_key = "%s"
		}

		resource "ibm_is_instance" "test_instance" {
			name    = "%s"
			profile = "%s"
			catalog_offering {
				version_crn = "%s"
				plan_crn    = "%s"
			}
			vpc  = ibm_is_vpc.test_vpc.id
			zone = ibm_is_subnet.test_subnet.zone
			keys = [ibm_is_ssh_key.test_key.id]

			primary_network_attachment {
				virtual_network_interface {
					subnet = ibm_is_subnet.test_subnet.id
				}
			}
		}

		data "ibm_is_volume" "test_volume" {
			identifier = ibm_is_instance.test_instance.boot_volume[0].volume_id
		}

		resource "ibm_is_volume_software_attachment" "is_volume_software_attachment_instance" {
			volume_id                     = ibm_is_instance.test_instance.boot_volume[0].volume_id
			volume_software_attachment_id = data.ibm_is_volume.test_volume.software_attachments.0.id
			name                          = "%s"
		}
	`, vpcname, subnetname, acc.ISZoneName, sshname, publicKey, instanceName, acc.InstanceProfileName, acc.ISCatalogImageOfferingCRN, acc.ISCatalogImagePlanCRN, name)
}

func testAccCheckIBMIsVolumeSoftwareAttachmentExists(n string, obj vpcv1.VolumeSoftwareAttachment) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		vpcClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).VpcV1API()
		if err != nil {
			return err
		}

		getVolumeSoftwareAttachmentOptions := &vpcv1.GetVolumeSoftwareAttachmentOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getVolumeSoftwareAttachmentOptions.SetVolumeID(parts[0])
		getVolumeSoftwareAttachmentOptions.SetID(parts[1])

		volumeSoftwareAttachment, _, err := vpcClient.GetVolumeSoftwareAttachment(getVolumeSoftwareAttachmentOptions)
		if err != nil {
			return err
		}

		obj = *volumeSoftwareAttachment
		return nil
	}
}

func testAccCheckIBMIsVolumeSoftwareAttachmentDestroy(s *terraform.State) error {
	vpcClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).VpcV1API()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_is_volume_software_attachment" {
			continue
		}

		getVolumeSoftwareAttachmentOptions := &vpcv1.GetVolumeSoftwareAttachmentOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getVolumeSoftwareAttachmentOptions.SetVolumeID(parts[0])
		getVolumeSoftwareAttachmentOptions.SetID(parts[1])

		// The attachment is removed together with its volume, which is destroyed in the same test.
		_, response, err := vpcClient.GetVolumeSoftwareAttachment(getVolumeSoftwareAttachmentOptions)

		if err == nil {
			return fmt.Errorf("VolumeSoftwareAttachment still exists: %s", rs.Primary.ID)
		} else if response == nil || response.StatusCode != 404 {
			return fmt.Errorf("Error checking for VolumeSoftwareAttachment (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
