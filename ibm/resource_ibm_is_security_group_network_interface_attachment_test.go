// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"errors"
	"fmt"
	"strings"
	"testing"

	"github.com/IBM/vpc-go-sdk/vpcclassicv1"
	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccIBMISSecurityGroupNwInterfaceAttachment_basic(t *testing.T) {
	var instanceNic string
	vpcname := fmt.Sprintf("tf-vpc-%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tf-subnet-%d", acctest.RandIntRange(10, 100))
	sgName := fmt.Sprintf("tfsg-%d", acctest.RandIntRange(10, 100))
	publicKey := strings.TrimSpace(`
ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCKVmnMOlHKcZK8tpt3MP1lqOLAcqcJzhsvJcjscgVERRN7/9484SOBJ3HSKxxNG5JN8owAjy5f9yYwcUg+JaUVuytn5Pv3aeYROHGGg+5G346xaq3DAwX6Y5ykr2fvjObgncQBnuU5KHWCECO/4h8uWuwh/kfniXPVjFToc+gnkqA+3RKpAecZhFXwfalQ9mMuYGFxn+fwn8cYEApsJbsEmb0iJwPiZ5hjFC8wREuiTlhPHDgkBLOiycd20op2nXzDbHfCHInquEe/gYxEitALONxm0swBOwJZwlTDOB7C6y2dzlrtxr1L59m7pCkWI4EtTRLvleehBoj3u7jB4usR
`)
	sshname := fmt.Sprintf("tfssh-name-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMISSecurityGroupNwInterfaceAttachmentDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISSecurityGroupNwInterfaceAttachmentConfig(vpcname, subnetname, sshname, publicKey, name, sgName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISSecurityGroupNwInterfaceAttachmentExists("ibm_is_security_group_network_interface_attachment.sgnic", instanceNic),
					resource.TestCheckResourceAttrSet(
						"ibm_is_security_group_network_interface_attachment.sgnic", "security_group"),
				),
			},
		},
	})
}

func testAccCheckIBMISSecurityGroupNwInterfaceAttachmentDestroy(s *terraform.State) error {
	userDetails, _ := testAccProvider.Meta().(ClientSession).BluemixUserDetails()

	if userDetails.generation == 1 {
		sess, _ := testAccProvider.Meta().(ClientSession).VpcClassicV1API()
		for _, rs := range s.RootModule().Resources {
			if rs.Type != "ibm_is_security_group_network_interface_attachment" {
				continue
			}

			parts, err := idParts(rs.Primary.ID)
			if err != nil {
				return err
			}

			sgID := parts[0]
			nicID := parts[1]
			getsgnicptions := &vpcclassicv1.GetSecurityGroupNetworkInterfaceOptions{
				SecurityGroupID: &sgID,
				ID:              &nicID,
			}
			_, _, err1 := sess.GetSecurityGroupNetworkInterface(getsgnicptions)
			if err1 == nil {
				return fmt.Errorf("network interface still exists: %s", rs.Primary.ID)
			}
		}
	} else {
		sess, _ := testAccProvider.Meta().(ClientSession).VpcV1API()
		for _, rs := range s.RootModule().Resources {
			if rs.Type != "ibm_is_security_group_network_interface_attachment" {
				continue
			}

			parts, err := idParts(rs.Primary.ID)
			if err != nil {
				return err
			}

			sgID := parts[0]
			nicID := parts[1]
			getsgnicptions := &vpcv1.GetSecurityGroupNetworkInterfaceOptions{
				SecurityGroupID: &sgID,
				ID:              &nicID,
			}
			_, _, err1 := sess.GetSecurityGroupNetworkInterface(getsgnicptions)
			if err1 == nil {
				return fmt.Errorf("network interface still exists: %s", rs.Primary.ID)
			}
		}
	}

	return nil
}

func testAccCheckIBMISSecurityGroupNwInterfaceAttachmentExists(n, instance string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return errors.New("No Record ID is set")
		}
		parts, err := idParts(rs.Primary.ID)
		if err != nil {
			return err
		}

		sgID := parts[0]
		nicID := parts[1]

		userDetails, _ := testAccProvider.Meta().(ClientSession).BluemixUserDetails()
		if userDetails.generation == 1 {
			sess, _ := testAccProvider.Meta().(ClientSession).VpcClassicV1API()
			getsgnicptions := &vpcclassicv1.GetSecurityGroupNetworkInterfaceOptions{
				SecurityGroupID: &sgID,
				ID:              &nicID,
			}
			found, _, err := sess.GetSecurityGroupNetworkInterface(getsgnicptions)
			if err != nil {
				return err
			}
			instance = *found.ID
		} else {
			sess, _ := testAccProvider.Meta().(ClientSession).VpcV1API()
			getsgnicptions := &vpcv1.GetSecurityGroupNetworkInterfaceOptions{
				SecurityGroupID: &sgID,
				ID:              &nicID,
			}
			found, _, err := sess.GetSecurityGroupNetworkInterface(getsgnicptions)
			if err != nil {
				return err
			}
			instance = *found.ID
		}
		return nil
	}
}

func testAccCheckIBMISSecurityGroupNwInterfaceAttachmentConfig(vpcname, subnetname, sshname, publicKey, name, sgName string) string {
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
		public_key = "%s"
	  }
	  
	  resource "ibm_is_instance" "testacc_instance" {
		name    = "%s"
		image   = "%s"
		profile = "%s"
	  
		primary_network_interface {
		  port_speed = "100"
		  subnet     = ibm_is_subnet.testacc_subnet.id
		}
	  
		vpc  = ibm_is_vpc.testacc_vpc.id
		zone = "%s"
		keys = [ibm_is_ssh_key.testacc_sshkey.id]
	  }
	  
	  resource "ibm_is_security_group" "testacc_security_group" {
		name = "%s"
		vpc  = ibm_is_vpc.testacc_vpc.id
	  }
	  
	  resource "ibm_is_security_group_network_interface_attachment" "sgnic" {
		security_group    = ibm_is_security_group.testacc_security_group.id
		network_interface = ibm_is_instance.testacc_instance.primary_network_interface[0].id
	  }`, vpcname, subnetname, ISZoneName, ISCIDR, sshname, publicKey, name, isImage, instanceProfileName, ISZoneName, sgName)
}
