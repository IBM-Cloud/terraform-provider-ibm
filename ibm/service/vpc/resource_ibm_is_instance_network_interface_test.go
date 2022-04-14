// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"fmt"
	"strings"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/IBM/vpc-go-sdk/vpcv1"
)

func TestAccIBMIsInstanceNetworkInterfaceAllArgs(t *testing.T) {
	var conf vpcv1.NetworkInterface
	allowIPSpoofing := "false"
	name := fmt.Sprintf("tf-net-int%d", acctest.RandIntRange(10, 100))
	secGrpName := fmt.Sprintf("tf-sec-grp%d", acctest.RandIntRange(10, 100))
	primaryIpv4Address := "10.240.0.6"
	allowIPSpoofingUpdate := "true"
	nameUpdate := fmt.Sprintf("tf-net-int%d", acctest.RandIntRange(10, 100))
	vpcname := fmt.Sprintf("tf-vpc-%d", acctest.RandIntRange(10, 100))
	insname := fmt.Sprintf("tf-instance-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tf-subnet-%d", acctest.RandIntRange(10, 100))
	publicKey := strings.TrimSpace(`
    ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCKVmnMOlHKcZK8tpt3MP1lqOLAcqcJzhsvJcjscgVERRN7/9484SOBJ3HSKxxNG5JN8owAjy5f9yYwcUg+JaUVuytn5Pv3aeYROHGGg+5G346xaq3DAwX6Y5ykr2fvjObgncQBnuU5KHWCECO/4h8uWuwh/kfniXPVjFToc+gnkqA+3RKpAecZhFXwfalQ9mMuYGFxn+fwn8cYEApsJbsEmb0iJwPiZ5hjFC8wREuiTlhPHDgkBLOiycd20op2nXzDbHfCHInquEe/gYxEitALONxm0swBOwJZwlTDOB7C6y2dzlrtxr1L59m7pCkWI4EtTRLvleehBoj3u7jB4usR
    `)
	sshname := fmt.Sprintf("tf-ssh-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIsInstanceNetworkInterfaceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIsInstanceNetworkInterfaceConfig(vpcname, subnetname, sshname, publicKey, insname, allowIPSpoofing, name, primaryIpv4Address, secGrpName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIsInstanceNetworkInterfaceExists("ibm_is_instance_network_interface.is_instance_network_interface", conf),
					resource.TestCheckResourceAttr("ibm_is_instance_network_interface.is_instance_network_interface", "allow_ip_spoofing", allowIPSpoofing),
					resource.TestCheckResourceAttr("ibm_is_instance_network_interface.is_instance_network_interface", "name", name),
					resource.TestCheckResourceAttr("ibm_is_instance_network_interface.is_instance_network_interface", "primary_ipv4_address", primaryIpv4Address),
				),
			},
			{
				Config: testAccCheckIBMIsInstanceNetworkInterfaceConfig(vpcname, subnetname, sshname, publicKey, insname, allowIPSpoofingUpdate, nameUpdate, primaryIpv4Address, secGrpName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_is_instance_network_interface.is_instance_network_interface", "allow_ip_spoofing", allowIPSpoofingUpdate),
					resource.TestCheckResourceAttr("ibm_is_instance_network_interface.is_instance_network_interface", "name", nameUpdate),
					resource.TestCheckResourceAttr("ibm_is_instance_network_interface.is_instance_network_interface", "primary_ipv4_address", primaryIpv4Address),
				),
			},
			{
				ResourceName:      "ibm_is_instance_network_interface.is_instance_network_interface",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}
func TestAccIBMIsInstanceNetworkInterface_rip(t *testing.T) {
	var conf vpcv1.NetworkInterface
	allowIPSpoofing := "false"
	name := fmt.Sprintf("tf-net-int%d", acctest.RandIntRange(10, 100))
	secGrpName := fmt.Sprintf("tf-sec-grp%d", acctest.RandIntRange(10, 100))
	primaryIpv4Address := "10.240.0.6"
	vpcname := fmt.Sprintf("tf-vpc-%d", acctest.RandIntRange(10, 100))
	insname := fmt.Sprintf("tf-instance-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tf-subnet-%d", acctest.RandIntRange(10, 100))
	publicKey := strings.TrimSpace(`
    ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCKVmnMOlHKcZK8tpt3MP1lqOLAcqcJzhsvJcjscgVERRN7/9484SOBJ3HSKxxNG5JN8owAjy5f9yYwcUg+JaUVuytn5Pv3aeYROHGGg+5G346xaq3DAwX6Y5ykr2fvjObgncQBnuU5KHWCECO/4h8uWuwh/kfniXPVjFToc+gnkqA+3RKpAecZhFXwfalQ9mMuYGFxn+fwn8cYEApsJbsEmb0iJwPiZ5hjFC8wREuiTlhPHDgkBLOiycd20op2nXzDbHfCHInquEe/gYxEitALONxm0swBOwJZwlTDOB7C6y2dzlrtxr1L59m7pCkWI4EtTRLvleehBoj3u7jB4usR
    `)
	sshname := fmt.Sprintf("tf-ssh-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIsInstanceNetworkInterfaceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIsInstanceNetworkInterfaceRipConfig(vpcname, subnetname, sshname, publicKey, insname, allowIPSpoofing, name, primaryIpv4Address, secGrpName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIsInstanceNetworkInterfaceExists("ibm_is_instance_network_interface.is_instance_network_interface", conf),
					resource.TestCheckResourceAttr("ibm_is_instance_network_interface.is_instance_network_interface", "allow_ip_spoofing", allowIPSpoofing),
					resource.TestCheckResourceAttr("ibm_is_instance_network_interface.is_instance_network_interface", "name", name),
					resource.TestCheckResourceAttr("ibm_is_instance_network_interface.is_instance_network_interface", "primary_ipv4_address", primaryIpv4Address),
				),
			},
			{
				ResourceName:      "ibm_is_instance_network_interface.is_instance_network_interface",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIBMIsInstanceNetworkInterfaceConfig(vpcname, subnetname, sshname, publicKey, insname, allowIPSpoofing, name, primaryIpv4Address, secGrpName string) string {
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
		vpc  = ibm_is_vpc.testacc_vpc.id
		zone = "%s"
		keys = [ibm_is_ssh_key.testacc_sshkey.id]
	}

	resource "ibm_is_security_group" "testacc_security_group" {
		name = "%s"
		vpc  = ibm_is_vpc.testacc_vpc.id
	}
	resource "ibm_is_instance_network_interface" "is_instance_network_interface" {
		instance = ibm_is_instance.testacc_instance.id
		subnet = ibm_is_subnet.testacc_subnet.id
		allow_ip_spoofing = %s
		name = "%s"
		primary_ipv4_address = "%s"
	}
	`, vpcname, subnetname, acc.ISZoneName, acc.ISCIDR, sshname, publicKey, insname, acc.IsImage, acc.InstanceProfileName, acc.ISZoneName, secGrpName, allowIPSpoofing, name, primaryIpv4Address)
}
func testAccCheckIBMIsInstanceNetworkInterfaceRipConfig(vpcname, subnetname, sshname, publicKey, insname, allowIPSpoofing, name, primaryIpv4Address, secGrpName string) string {
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
		vpc  = ibm_is_vpc.testacc_vpc.id
		zone = "%s"
		keys = [ibm_is_ssh_key.testacc_sshkey.id]
	}

	resource "ibm_is_security_group" "testacc_security_group" {
		name = "%s"
		vpc  = ibm_is_vpc.testacc_vpc.id
	}
	resource "ibm_is_instance_network_interface" "is_instance_network_interface" {
		instance = ibm_is_instance.testacc_instance.id
		subnet = ibm_is_subnet.testacc_subnet.id
		allow_ip_spoofing = %s
		name = "%s"
		primary_ip {
			address = "%s"
		}
	}
	`, vpcname, subnetname, acc.ISZoneName, acc.ISCIDR, sshname, publicKey, insname, acc.IsImage, acc.InstanceProfileName, acc.ISZoneName, secGrpName, allowIPSpoofing, name, primaryIpv4Address)
}

func testAccCheckIBMIsInstanceNetworkInterfaceExists(n string, obj vpcv1.NetworkInterface) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		vpcClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).VpcV1API()
		if err != nil {
			return err
		}

		getInstanceNetworkInterfaceOptions := &vpcv1.GetInstanceNetworkInterfaceOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getInstanceNetworkInterfaceOptions.SetInstanceID(parts[0])
		getInstanceNetworkInterfaceOptions.SetID(parts[1])

		networkInterface, _, err := vpcClient.GetInstanceNetworkInterface(getInstanceNetworkInterfaceOptions)
		if err != nil {
			return err
		}

		obj = *networkInterface
		return nil
	}
}

func testAccCheckIBMIsInstanceNetworkInterfaceDestroy(s *terraform.State) error {
	vpcClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).VpcV1API()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_is_instance_network_interface" {
			continue
		}

		getInstanceNetworkInterfaceOptions := &vpcv1.GetInstanceNetworkInterfaceOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getInstanceNetworkInterfaceOptions.SetInstanceID(parts[0])
		getInstanceNetworkInterfaceOptions.SetID(parts[1])

		// Try to find the key
		_, response, err := vpcClient.GetInstanceNetworkInterface(getInstanceNetworkInterfaceOptions)

		if err == nil {
			return fmt.Errorf("NetworkInterface still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("[ERROR] Error checking for NetworkInterface (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
