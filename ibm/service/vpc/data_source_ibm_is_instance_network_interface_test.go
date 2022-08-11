// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"fmt"
	"strings"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMIsInstanceNetworkInterfaceDataSourceBasic(t *testing.T) {

	networkInterfaceName := fmt.Sprintf("tf-net-int%d", acctest.RandIntRange(10, 100))
	vpcname := fmt.Sprintf("tf-vpc-%d", acctest.RandIntRange(10, 100))
	insname := fmt.Sprintf("tf-instance-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tf-subnet-%d", acctest.RandIntRange(10, 100))
	publicKey := strings.TrimSpace(`
    ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCKVmnMOlHKcZK8tpt3MP1lqOLAcqcJzhsvJcjscgVERRN7/9484SOBJ3HSKxxNG5JN8owAjy5f9yYwcUg+JaUVuytn5Pv3aeYROHGGg+5G346xaq3DAwX6Y5ykr2fvjObgncQBnuU5KHWCECO/4h8uWuwh/kfniXPVjFToc+gnkqA+3RKpAecZhFXwfalQ9mMuYGFxn+fwn8cYEApsJbsEmb0iJwPiZ5hjFC8wREuiTlhPHDgkBLOiycd20op2nXzDbHfCHInquEe/gYxEitALONxm0swBOwJZwlTDOB7C6y2dzlrtxr1L59m7pCkWI4EtTRLvleehBoj3u7jB4usR
    `)
	sshname := fmt.Sprintf("tf-ssh-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIsInstanceNetworkInterfaceDataSourceConfigBasic(vpcname, subnetname, sshname, publicKey, insname, networkInterfaceName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_network_interface.is_instance_network_interface", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_network_interface.is_instance_network_interface", "allow_ip_spoofing"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_network_interface.is_instance_network_interface", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_network_interface.is_instance_network_interface", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_network_interface.is_instance_network_interface", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_network_interface.is_instance_network_interface", "port_speed"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_network_interface.is_instance_network_interface", "primary_ipv4_address"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_network_interface.is_instance_network_interface", "resource_type"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_network_interface.is_instance_network_interface", "security_groups.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_network_interface.is_instance_network_interface", "status"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_network_interface.is_instance_network_interface", "subnet.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_network_interface.is_instance_network_interface", "type"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_network_interface.is_instance_network_interface", "primary_ip.0.address"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_network_interface.is_instance_network_interface", "primary_ip.0.name"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_network_interface.is_instance_network_interface", "primary_ip.0.href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_network_interface.is_instance_network_interface", "primary_ip.0.reserved_ip"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_network_interface.is_instance_network_interface", "primary_ip.0.resource_type"),
				),
			},
		},
	})
}

func TestAccIBMIsInstanceNetworkInterfaceDataSourceAllArgs(t *testing.T) {
	allowIPSpoofing := "false"
	networkInterfaceName := fmt.Sprintf("tf-net-int%d", acctest.RandIntRange(10, 100))
	primaryIpv4Address := "10.240.0.6"
	vpcname := fmt.Sprintf("tf-vpc-%d", acctest.RandIntRange(10, 100))
	insname := fmt.Sprintf("tf-instance-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tf-subnet-%d", acctest.RandIntRange(10, 100))
	publicKey := strings.TrimSpace(`
    ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCKVmnMOlHKcZK8tpt3MP1lqOLAcqcJzhsvJcjscgVERRN7/9484SOBJ3HSKxxNG5JN8owAjy5f9yYwcUg+JaUVuytn5Pv3aeYROHGGg+5G346xaq3DAwX6Y5ykr2fvjObgncQBnuU5KHWCECO/4h8uWuwh/kfniXPVjFToc+gnkqA+3RKpAecZhFXwfalQ9mMuYGFxn+fwn8cYEApsJbsEmb0iJwPiZ5hjFC8wREuiTlhPHDgkBLOiycd20op2nXzDbHfCHInquEe/gYxEitALONxm0swBOwJZwlTDOB7C6y2dzlrtxr1L59m7pCkWI4EtTRLvleehBoj3u7jB4usR
    `)
	sshname := fmt.Sprintf("tf-ssh-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIsInstanceNetworkInterfaceDataSourceConfig(vpcname, subnetname, sshname, publicKey, insname, allowIPSpoofing, networkInterfaceName, primaryIpv4Address),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_network_interface.is_instance_network_interface", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_network_interface.is_instance_network_interface", "allow_ip_spoofing"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_network_interface.is_instance_network_interface", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_network_interface.is_instance_network_interface", "href"),
					resource.TestCheckResourceAttr("data.ibm_is_instance_network_interface.is_instance_network_interface", "name", networkInterfaceName),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_network_interface.is_instance_network_interface", "port_speed"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_network_interface.is_instance_network_interface", "primary_ipv4_address"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_network_interface.is_instance_network_interface", "resource_type"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_network_interface.is_instance_network_interface", "security_groups.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_network_interface.is_instance_network_interface", "security_groups.0.crn"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_network_interface.is_instance_network_interface", "security_groups.0.href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_network_interface.is_instance_network_interface", "security_groups.0.id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_network_interface.is_instance_network_interface", "security_groups.0.name"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_network_interface.is_instance_network_interface", "status"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_network_interface.is_instance_network_interface", "subnet.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_network_interface.is_instance_network_interface", "type"),
				),
			},
		},
	})
}

func testAccCheckIBMIsInstanceNetworkInterfaceDataSourceConfigBasic(vpcname, subnetname, sshname, publicKey, insname, name string) string {

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
   		resource "ibm_is_instance_network_interface" "is_instance_network_interface" {
   			instance = ibm_is_instance.testacc_instance.id
   			subnet = ibm_is_subnet.testacc_subnet.id
   			name = "%s"
   		}
   		data "ibm_is_instance_network_interface" "is_instance_network_interface" {
   			instance_name = ibm_is_instance.testacc_instance.name
   			network_interface_name = ibm_is_instance_network_interface.is_instance_network_interface.name
   		}
	`, vpcname, subnetname, acc.ISZoneName, acc.ISCIDR, sshname, publicKey, insname, acc.IsImage, acc.InstanceProfileName, acc.ISZoneName, name)
}

func testAccCheckIBMIsInstanceNetworkInterfaceDataSourceConfig(vpcname, subnetname, sshname, publicKey, insname, allowIPSpoofing, name, primaryIpv4Address string) string {
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
		resource "ibm_is_instance_network_interface" "is_instance_network_interface" {
			instance = ibm_is_instance.testacc_instance.id
			subnet = ibm_is_subnet.testacc_subnet.id
			allow_ip_spoofing = %s
			name = "%s"
			primary_ipv4_address = "%s"
		}
		data "ibm_is_instance_network_interface" "is_instance_network_interface" {
			instance_name = ibm_is_instance.testacc_instance.name
			network_interface_name = ibm_is_instance_network_interface.is_instance_network_interface.name
		}
	`, vpcname, subnetname, acc.ISZoneName, acc.ISCIDR, sshname, publicKey, insname, acc.IsImage, acc.InstanceProfileName, acc.ISZoneName, allowIPSpoofing, name, primaryIpv4Address)
}
