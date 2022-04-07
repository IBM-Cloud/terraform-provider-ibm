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

func TestAccIBMISFloatingIP_basic(t *testing.T) {
	var ip string
	vpcname := fmt.Sprintf("tfip-vpc-%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tfip-%d", acctest.RandIntRange(10, 100))
	instancename := fmt.Sprintf("tfip-instance-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tfip-subnet-%d", acctest.RandIntRange(10, 100))
	publicKey := strings.TrimSpace(`
ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCKVmnMOlHKcZK8tpt3MP1lqOLAcqcJzhsvJcjscgVERRN7/9484SOBJ3HSKxxNG5JN8owAjy5f9yYwcUg+JaUVuytn5Pv3aeYROHGGg+5G346xaq3DAwX6Y5ykr2fvjObgncQBnuU5KHWCECO/4h8uWuwh/kfniXPVjFToc+gnkqA+3RKpAecZhFXwfalQ9mMuYGFxn+fwn8cYEApsJbsEmb0iJwPiZ5hjFC8wREuiTlhPHDgkBLOiycd20op2nXzDbHfCHInquEe/gYxEitALONxm0swBOwJZwlTDOB7C6y2dzlrtxr1L59m7pCkWI4EtTRLvleehBoj3u7jB4usR
`)
	sshname := fmt.Sprintf("tfip-sshname-%d", acctest.RandIntRange(10, 100))
	userData1 := "a"
	userData2 := "b"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISFloatingIPDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISFloatingIPConfig(vpcname, subnetname, sshname, publicKey, instancename, userData1, name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISFloatingIPExists("ibm_is_floating_ip.testacc_floatingip", ip),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "user_data", userData1),
					resource.TestCheckResourceAttr(
						"ibm_is_floating_ip.testacc_floatingip", "name", name),
					resource.TestCheckResourceAttr(
						"ibm_is_floating_ip.testacc_floatingip", "zone", acc.ISZoneName),
				),
			},
			{
				Config: testAccCheckIBMISFloatingIPConfig(vpcname, subnetname, sshname, publicKey, instancename, userData2, name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISFloatingIPExists("ibm_is_floating_ip.testacc_floatingip", ip),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "user_data", userData2),
					resource.TestCheckResourceAttr(
						"ibm_is_floating_ip.testacc_floatingip", "name", name),
					resource.TestCheckResourceAttr(
						"ibm_is_floating_ip.testacc_floatingip", "zone", acc.ISZoneName),
				),
			},
		},
	})
}

func TestAccIBMISFloatingIP_NoTarget(t *testing.T) {
	var ip string
	name := fmt.Sprintf("tfip-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISFloatingIPDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISFloatingIPNoTargetConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISFloatingIPExists("ibm_is_floating_ip.testacc_floatingip", ip),
					resource.TestCheckResourceAttr(
						"ibm_is_floating_ip.testacc_floatingip", "name", name),
					resource.TestCheckResourceAttr(
						"ibm_is_floating_ip.testacc_floatingip", "zone", acc.ISZoneName),
				),
			},
		},
	})
}

func testAccCheckIBMISFloatingIPDestroy(s *terraform.State) error {

	sess, _ := acc.TestAccProvider.Meta().(conns.ClientSession).VpcV1API()
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_is_floating_ip" {
			continue
		}

		getfipoptions := &vpcv1.GetFloatingIPOptions{
			ID: &rs.Primary.ID,
		}
		_, _, err := sess.GetFloatingIP(getfipoptions)
		if err == nil {
			return fmt.Errorf("Floating IP still exists: %s", rs.Primary.ID)
		}
	}

	return nil
}

func testAccCheckIBMISFloatingIPExists(n, ip string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return errors.New("No Record ID is set")
		}
		sess, _ := acc.TestAccProvider.Meta().(conns.ClientSession).VpcV1API()
		getfipoptions := &vpcv1.GetFloatingIPOptions{
			ID: &rs.Primary.ID,
		}
		foundip, _, err := sess.GetFloatingIP(getfipoptions)
		if err != nil {
			return err
		}
		ip = *foundip.ID
		return nil
	}
}

func testAccCheckIBMISFloatingIPConfig(vpcname, subnetname, sshname, publicKey, instancename, userData, name string) string {
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
		  subnet     = ibm_is_subnet.testacc_subnet.id
		}
		user_data = "%s"
		vpc  = ibm_is_vpc.testacc_vpc.id
		zone = "%s"
		keys = [ibm_is_ssh_key.testacc_sshkey.id]
	  }
	  
	  resource "ibm_is_floating_ip" "testacc_floatingip" {
		name   = "%s"
		target = ibm_is_instance.testacc_instance.primary_network_interface[0].id
	  }
`, vpcname, subnetname, acc.ISZoneName, acc.ISCIDR, sshname, publicKey, instancename, acc.IsImage, acc.InstanceProfileName, userData, acc.ISZoneName, name)
}

func testAccCheckIBMISFloatingIPNoTargetConfig(name string) string {
	return fmt.Sprintf(`
	  resource "ibm_is_floating_ip" "testacc_floatingip" {
		name   = "%s"
		zone   = "%s"
	  }
`, name, acc.ISZoneName)
}
