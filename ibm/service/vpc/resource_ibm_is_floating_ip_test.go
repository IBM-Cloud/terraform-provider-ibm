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
	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccIBMISFloatingIP_basic(t *testing.T) {
	var ip string
	vpcname := fmt.Sprintf("tfip-vpc-%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tfip-%d", acctest.RandIntRange(10, 100))
	updatedName := fmt.Sprintf("tfip-updated-%d", acctest.RandIntRange(10, 100))
	instancename := fmt.Sprintf("tfip-instance-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tfip-subnet-%d", acctest.RandIntRange(10, 100))
	publicKey := strings.TrimSpace(`
ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCKVmnMOlHKcZK8tpt3MP1lqOLAcqcJzhsvJcjscgVERRN7/9484SOBJ3HSKxxNG5JN8owAjy5f9yYwcUg+JaUVuytn5Pv3aeYROHGGg+5G346xaq3DAwX6Y5ykr2fvjObgncQBnuU5KHWCECO/4h8uWuwh/kfniXPVjFToc+gnkqA+3RKpAecZhFXwfalQ9mMuYGFxn+fwn8cYEApsJbsEmb0iJwPiZ5hjFC8wREuiTlhPHDgkBLOiycd20op2nXzDbHfCHInquEe/gYxEitALONxm0swBOwJZwlTDOB7C6y2dzlrtxr1L59m7pCkWI4EtTRLvleehBoj3u7jB4usR
`)
	sshname := fmt.Sprintf("tfip-sshname-%d", acctest.RandIntRange(10, 100))
	userData1 := "a"
	userData2 := "b"
	resourceKey := "ibm_is_floating_ip.testacc_floatingip"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISFloatingIPDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISFloatingIPConfig(vpcname, subnetname, sshname, publicKey, instancename, userData1, name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISFloatingIPExists(resourceKey, ip),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "user_data", userData1),
					resource.TestCheckResourceAttr(
						resourceKey, "name", name),
					resource.TestCheckResourceAttr(
						resourceKey, "zone", acc.ISZoneName),
					resource.TestCheckResourceAttrSet(resourceKey, "address"),
					resource.TestCheckResourceAttrSet(resourceKey, "status"),
					resource.TestCheckResourceAttrSet(resourceKey, "target"),
					resource.TestCheckResourceAttrSet(resourceKey, "target_list.0.primary_ip.0.address"),
					resource.TestCheckResourceAttrSet(resourceKey, "target_list.0.primary_ip.0.name"),
					resource.TestCheckResourceAttrSet(resourceKey, "target_list.0.primary_ip.0.href"),
					resource.TestCheckResourceAttrSet(resourceKey, "target_list.0.primary_ip.0.reserved_ip"),
					resource.TestCheckResourceAttrSet(resourceKey, "target_list.0.primary_ip.0.resource_type"),
					resource.TestCheckResourceAttrSet(resourceKey, "crn"),
					resource.TestCheckResourceAttrSet(resourceKey, "resource_controller_url"),
					resource.TestCheckResourceAttrSet(resourceKey, "resource_name"),
					resource.TestCheckResourceAttrSet(resourceKey, "resource_group_name"),
				),
			},
			{
				Config: testAccCheckIBMISFloatingIPConfig(vpcname, subnetname, sshname, publicKey, instancename, userData2, name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISFloatingIPExists(resourceKey, ip),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "user_data", userData2),
					resource.TestCheckResourceAttr(
						resourceKey, "name", name),
					resource.TestCheckResourceAttr(
						resourceKey, "zone", acc.ISZoneName),
					resource.TestCheckResourceAttrSet(resourceKey, "target_list.0.primary_ip.0.address"),
					resource.TestCheckResourceAttrSet(resourceKey, "target_list.0.primary_ip.0.name"),
					resource.TestCheckResourceAttrSet(resourceKey, "target_list.0.primary_ip.0.href"),
					resource.TestCheckResourceAttrSet(resourceKey, "target_list.0.primary_ip.0.reserved_ip"),
					resource.TestCheckResourceAttrSet(resourceKey, "target_list.0.primary_ip.0.resource_type"),
				),
			},
			// Test name update
			{
				Config: testAccCheckIBMISFloatingIPConfigNameUpdate(vpcname, subnetname, sshname, publicKey, instancename, userData2, updatedName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISFloatingIPExists(resourceKey, ip),
					resource.TestCheckResourceAttr(
						resourceKey, "name", updatedName),
					resource.TestCheckResourceAttrSet(resourceKey, "address"),
					resource.TestCheckResourceAttrSet(resourceKey, "status"),
					resource.TestCheckResourceAttrSet(resourceKey, "target"),
				),
			},
			// Test import
			{
				ResourceName:      resourceKey,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccIBMISFloatingIP_NoTarget(t *testing.T) {
	var ip string
	name := fmt.Sprintf("tfip-%d", acctest.RandIntRange(10, 100))
	updatedName := fmt.Sprintf("tfip-updated-%d", acctest.RandIntRange(10, 100))
	resourceKey := "ibm_is_floating_ip.testacc_floatingip"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISFloatingIPDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISFloatingIPNoTargetConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISFloatingIPExists(resourceKey, ip),
					resource.TestCheckResourceAttr(
						resourceKey, "name", name),
					resource.TestCheckResourceAttr(
						resourceKey, "zone", acc.ISZoneName),
					resource.TestCheckResourceAttrSet(resourceKey, "address"),
					resource.TestCheckResourceAttrSet(resourceKey, "status"),
					resource.TestCheckResourceAttr(resourceKey, "target", ""),
				),
			},
			// Test name update for zone-based FIP
			{
				Config: testAccCheckIBMISFloatingIPNoTargetConfig(updatedName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISFloatingIPExists(resourceKey, ip),
					resource.TestCheckResourceAttr(
						resourceKey, "name", updatedName),
					resource.TestCheckResourceAttr(
						resourceKey, "zone", acc.ISZoneName),
					resource.TestCheckResourceAttrSet(resourceKey, "address"),
				),
			},
			// Test import
			{
				ResourceName:      resourceKey,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccIBMISFloatingIP_ResourceGroup(t *testing.T) {
	var ip string
	name := fmt.Sprintf("tfip-rg-%d", acctest.RandIntRange(10, 100))
	resourceKey := "ibm_is_floating_ip.testacc_floatingip"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISFloatingIPDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISFloatingIPResourceGroupConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISFloatingIPExists(resourceKey, ip),
					resource.TestCheckResourceAttr(
						resourceKey, "name", name),
					resource.TestCheckResourceAttr(
						resourceKey, "zone", acc.ISZoneName),
					resource.TestCheckResourceAttrSet(resourceKey, "resource_group"),
					resource.TestCheckResourceAttrSet(resourceKey, "resource_group_name"),
					resource.TestCheckResourceAttrSet(resourceKey, "address"),
					resource.TestCheckResourceAttrSet(resourceKey, "status"),
				),
			},
		},
	})
}

func TestAccIBMISFloatingIP_Tags(t *testing.T) {
	var ip string
	name := fmt.Sprintf("tfip-tags-%d", acctest.RandIntRange(10, 100))
	resourceKey := "ibm_is_floating_ip.testacc_floatingip"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISFloatingIPDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISFloatingIPTagsConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISFloatingIPExists(resourceKey, ip),
					resource.TestCheckResourceAttr(
						resourceKey, "name", name),
					resource.TestCheckResourceAttr(
						resourceKey, "zone", acc.ISZoneName),
					resource.TestCheckResourceAttr(resourceKey, "tags.#", "2"),
					resource.TestCheckResourceAttrSet(resourceKey, "address"),
					resource.TestCheckResourceAttrSet(resourceKey, "status"),
				),
			},
			// Test tag update
			{
				Config: testAccCheckIBMISFloatingIPTagsUpdateConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISFloatingIPExists(resourceKey, ip),
					resource.TestCheckResourceAttr(resourceKey, "tags.#", "3"),
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
			subnet = ibm_is_subnet.testacc_subnet.id
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

func testAccCheckIBMISFloatingIPConfigNameUpdate(vpcname, subnetname, sshname, publicKey, instancename, userData, name string) string {
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
			subnet = ibm_is_subnet.testacc_subnet.id
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
		name = "%s"
		zone = "%s"
	}
`, name, acc.ISZoneName)
}

func testAccCheckIBMISFloatingIPResourceGroupConfig(name string) string {
	return fmt.Sprintf(`
	data "ibm_resource_group" "test_group" {
		name = "Default"
	}
	
	resource "ibm_is_floating_ip" "testacc_floatingip" {
		name = "%s"
		zone = "%s"
		resource_group = data.ibm_resource_group.test_group.id
	}
`, name, acc.ISZoneName)
}

func testAccCheckIBMISFloatingIPTagsConfig(name string) string {
	return fmt.Sprintf(`
	resource "ibm_is_floating_ip" "testacc_floatingip" {
		name = "%s"
		zone = "%s"
		tags = ["tag1", "tag2"]
	}
`, name, acc.ISZoneName)
}

func testAccCheckIBMISFloatingIPTagsUpdateConfig(name string) string {
	return fmt.Sprintf(`
	resource "ibm_is_floating_ip" "testacc_floatingip" {
		name = "%s"
		zone = "%s"
		tags = ["tag1", "tag2", "tag3"]
	}
`, name, acc.ISZoneName)
}
