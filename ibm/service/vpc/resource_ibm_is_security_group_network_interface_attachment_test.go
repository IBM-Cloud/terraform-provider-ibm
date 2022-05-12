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
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"

	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
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
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
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
	sess, _ := acc.TestAccProvider.Meta().(conns.ClientSession).VpcV1API()
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_is_security_group_network_interface_attachment" {
			continue
		}

		parts, err := flex.IdParts(rs.Primary.ID)
		if err != nil {
			return err
		}

		sgID := parts[0]
		nicID := parts[1]
		getsgnicptions := &vpcv1.GetSecurityGroupTargetOptions{
			SecurityGroupID: &sgID,
			ID:              &nicID,
		}
		_, _, err1 := sess.GetSecurityGroupTarget(getsgnicptions)
		if err1 == nil {
			return fmt.Errorf("network interface still exists: %s", rs.Primary.ID)
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
		parts, err := flex.IdParts(rs.Primary.ID)
		if err != nil {
			return err
		}

		sgID := parts[0]
		nicID := parts[1]

		sess, _ := acc.TestAccProvider.Meta().(conns.ClientSession).VpcV1API()
		getsgnicptions := &vpcv1.GetSecurityGroupTargetOptions{
			SecurityGroupID: &sgID,
			ID:              &nicID,
		}
		found, _, err := sess.GetSecurityGroupTarget(getsgnicptions)
		if err != nil {
			return err
		}
		instance = *found.(*vpcv1.SecurityGroupTargetReference).ID
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
	  }`, vpcname, subnetname, acc.ISZoneName, acc.ISCIDR, sshname, publicKey, name, acc.IsImage, acc.InstanceProfileName, acc.ISZoneName, sgName)
}
