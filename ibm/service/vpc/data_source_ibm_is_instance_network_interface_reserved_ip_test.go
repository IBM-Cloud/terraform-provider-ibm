// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
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

func TestAccIBMISInstanceNICReservedIP_basic(t *testing.T) {
	vpcname := fmt.Sprintf("tf-vpc-%d", acctest.RandIntRange(10, 100))
	insname := fmt.Sprintf("tf-instance-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tf-subnet-%d", acctest.RandIntRange(10, 100))
	ripname := fmt.Sprintf("tf-rip-%d", acctest.RandIntRange(10, 100))
	publicKey := strings.TrimSpace(`
    ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCKVmnMOlHKcZK8tpt3MP1lqOLAcqcJzhsvJcjscgVERRN7/9484SOBJ3HSKxxNG5JN8owAjy5f9yYwcUg+JaUVuytn5Pv3aeYROHGGg+5G346xaq3DAwX6Y5ykr2fvjObgncQBnuU5KHWCECO/4h8uWuwh/kfniXPVjFToc+gnkqA+3RKpAecZhFXwfalQ9mMuYGFxn+fwn8cYEApsJbsEmb0iJwPiZ5hjFC8wREuiTlhPHDgkBLOiycd20op2nXzDbHfCHInquEe/gYxEitALONxm0swBOwJZwlTDOB7C6y2dzlrtxr1L59m7pCkWI4EtTRLvleehBoj3u7jB4usR
    `)
	sshname := fmt.Sprintf("tf-ssh-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccIBMISInstanceNICReservedIPdataSoruceConfig(vpcname, subnetname, ripname, sshname, publicKey, insname),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_is_vpc.testacc_vpc", "name", vpcname),
					resource.TestCheckResourceAttr("ibm_is_subnet.testacc_subnet", "name", subnetname),
					resource.TestCheckResourceAttr("ibm_is_ssh_key.testacc_sshkey", "name", sshname),
					resource.TestCheckResourceAttr("ibm_is_instance.testacc_instance", "name", insname),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_network_interface_reserved_ip.instance_network_interface_reserved_ip", "address"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_network_interface_reserved_ip.instance_network_interface_reserved_ip", "auto_delete"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_network_interface_reserved_ip.instance_network_interface_reserved_ip", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_network_interface_reserved_ip.instance_network_interface_reserved_ip", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_network_interface_reserved_ip.instance_network_interface_reserved_ip", "id"),
					resource.TestCheckResourceAttr("data.ibm_is_instance_network_interface_reserved_ip.instance_network_interface_reserved_ip", "name", ripname),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_network_interface_reserved_ip.instance_network_interface_reserved_ip", "resource_type"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_network_interface_reserved_ip.instance_network_interface_reserved_ip", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_network_interface_reserved_ip.instance_network_interface_reserved_ip", "target"),
				),
			},
		},
	})
}

func testAccIBMISInstanceNICReservedIPdataSoruceConfig(vpcname, subnetname, ripname, sshname, publicKey, insname string) string {
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
	resource "ibm_is_subnet_reserved_ip" "testacc_rip" {
		subnet = ibm_is_subnet.testacc_subnet.id
		name = "%s"
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
			primary_ip {
				reserved_ip = ibm_is_subnet_reserved_ip.testacc_rip.reserved_ip
			}
		}
		vpc  = ibm_is_vpc.testacc_vpc.id
		zone = "%s"
		keys = [ibm_is_ssh_key.testacc_sshkey.id]
	}

	data "ibm_is_instance_network_interface_reserved_ip" "instance_network_interface_reserved_ip" {
		instance = ibm_is_instance.testacc_instance.id
		network_interface = ibm_is_instance.testacc_instance.primary_network_interface.0.id
		reserved_ip = ibm_is_instance.testacc_instance.primary_network_interface.0.primary_ip.0.reserved_ip
	}
      `, vpcname, subnetname, acc.ISZoneName, acc.ISCIDR, ripname, sshname, publicKey, insname, acc.IsImage, acc.InstanceProfileName, acc.ISZoneName)
}
