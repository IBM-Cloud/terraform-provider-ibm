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

func TestAccIBMISInstanceVolumeAttachment_basic(t *testing.T) {
	var instanceVolAtt string
	vpcname := fmt.Sprintf("tf-vpc-%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf-instnace-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tf-subnet-%d", acctest.RandIntRange(10, 100))
	publicKey := strings.TrimSpace(`
ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCKVmnMOlHKcZK8tpt3MP1lqOLAcqcJzhsvJcjscgVERRN7/9484SOBJ3HSKxxNG5JN8owAjy5f9yYwcUg+JaUVuytn5Pv3aeYROHGGg+5G346xaq3DAwX6Y5ykr2fvjObgncQBnuU5KHWCECO/4h8uWuwh/kfniXPVjFToc+gnkqA+3RKpAecZhFXwfalQ9mMuYGFxn+fwn8cYEApsJbsEmb0iJwPiZ5hjFC8wREuiTlhPHDgkBLOiycd20op2nXzDbHfCHInquEe/gYxEitALONxm0swBOwJZwlTDOB7C6y2dzlrtxr1L59m7pCkWI4EtTRLvleehBoj3u7jB4usR
`)
	sshname := fmt.Sprintf("tf-ssh-%d", acctest.RandIntRange(10, 100))
	attName := fmt.Sprintf("tf-volatt-%d", acctest.RandIntRange(10, 100))
	autoDelete := true
	volName := fmt.Sprintf("tf-vol-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMISInstanceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISInstanceVolumeAttachmentConfig(vpcname, subnetname, sshname, publicKey, name, attName, volName, autoDelete),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISInstanceVolumeAttachmentExists("ibm_is_instance_volume_attachment.testacc_att", instanceVolAtt),
					resource.TestCheckResourceAttr(
						"ibm_is_instance_volume_attachment.testacc_att", "name", attName),
					resource.TestCheckResourceAttr(
						"ibm_is_instance_volume_attachment.testacc_att", "delete_volume_on_instance_delete", fmt.Sprintf("%t", autoDelete)),
					resource.TestCheckResourceAttr(
						"ibm_is_instance_volume_attachment.testacc_att", "capacity", "20"),
				),
			},
		},
	})
}

func testAccCheckIBMISInstanceVolumeAttachmentDestroy(s *terraform.State) error {

	instanceC, _ := testAccProvider.Meta().(ClientSession).VpcV1API()
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_is_instance_volume_attachment" {
			continue
		}
		getinsvolAttOptions := &vpcv1.GetInstanceVolumeAttachmentOptions{
			ID: &rs.Primary.ID,
		}
		_, _, err := instanceC.GetInstanceVolumeAttachment(getinsvolAttOptions)

		if err == nil {
			return fmt.Errorf("instance volume attachment still exists: %s", rs.Primary.ID)
		}
	}

	return nil
}

func testAccCheckIBMISInstanceVolumeAttachmentExists(n string, instanceVolAtt string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return errors.New("No Record ID is set")
		}
		instanceId, id, err := parseVolAttTerraformID(rs.Primary.ID)
		instanceC, _ := testAccProvider.Meta().(ClientSession).VpcV1API()
		getinsVolAttOptions := &vpcv1.GetInstanceVolumeAttachmentOptions{
			ID:         &id,
			InstanceID: &instanceId,
		}
		foundins, _, err := instanceC.GetInstanceVolumeAttachment(getinsVolAttOptions)
		if err != nil {
			return err
		}
		instanceVolAtt = *foundins.ID
		return nil
	}
}

func testAccCheckIBMISInstanceVolumeAttachmentConfig(vpcname, subnetname, sshname, publicKey, name, attName, volName string, autoDelete bool) string {
	return fmt.Sprintf(`
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
	  resource "ibm_is_instance_volume_attachment" "testacc_att" {
		instance = ibm_is_instance.testacc_instance.id
	
		name 			= "%s"
		profile 		= "general-purpose"
		capacity	 	= "20"
	
		delete_volume_on_instance_delete = %t
		volume_name = "%s"
	  }
	 
	  `, vpcname, subnetname, ISZoneName, sshname, publicKey, name, isImage, instanceProfileName, ISZoneName, attName, autoDelete, volName)
}
