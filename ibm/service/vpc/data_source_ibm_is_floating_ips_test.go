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

func TestAccIBMIsFloatingIpsDataSourceBasic(t *testing.T) {

	vpcname := fmt.Sprintf("tfip-vpc-%d", acctest.RandIntRange(10, 100))
	fipname := fmt.Sprintf("tfip-%d", acctest.RandIntRange(10, 100))
	instancename := fmt.Sprintf("tfip-instance-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tfip-subnet-%d", acctest.RandIntRange(10, 100))
	publicKey := strings.TrimSpace(`
ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCKVmnMOlHKcZK8tpt3MP1lqOLAcqcJzhsvJcjscgVERRN7/9484SOBJ3HSKxxNG5JN8owAjy5f9yYwcUg+JaUVuytn5Pv3aeYROHGGg+5G346xaq3DAwX6Y5ykr2fvjObgncQBnuU5KHWCECO/4h8uWuwh/kfniXPVjFToc+gnkqA+3RKpAecZhFXwfalQ9mMuYGFxn+fwn8cYEApsJbsEmb0iJwPiZ5hjFC8wREuiTlhPHDgkBLOiycd20op2nXzDbHfCHInquEe/gYxEitALONxm0swBOwJZwlTDOB7C6y2dzlrtxr1L59m7pCkWI4EtTRLvleehBoj3u7jB4usR
`)
	sshname := fmt.Sprintf("tfip-sshname-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIsFloatingIpsDataSourceConfigBasic(vpcname, subnetname, sshname, publicKey, instancename, fipname),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_floating_ips.is_floating_ips", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_floating_ips.is_floating_ips", "floating_ips.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_floating_ips.is_floating_ips", "floating_ips.0.target.0.primary_ip.0.address"),
					resource.TestCheckResourceAttrSet("data.ibm_is_floating_ips.is_floating_ips", "floating_ips.0.target.0.primary_ip.0.name"),
					resource.TestCheckResourceAttrSet("data.ibm_is_floating_ips.is_floating_ips", "floating_ips.0.target.0.primary_ip.0.href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_floating_ips.is_floating_ips", "floating_ips.0.target.0.primary_ip.0.reserved_ip"),
					resource.TestCheckResourceAttrSet("data.ibm_is_floating_ips.is_floating_ips", "floating_ips.0.target.0.primary_ip.0.resource_type"),
				),
			},
		},
	})
}

func testAccCheckIBMIsFloatingIpsDataSourceConfigBasic(vpcname, subnetname, sshname, publicKey, instancename, fipname string) string {
	// status filter defaults to empty
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
	  
	  resource "ibm_is_floating_ip" "testacc_floatingip" {
		name   = "%s"
		target = ibm_is_instance.testacc_instance.primary_network_interface[0].id
	  }

	  data "ibm_is_floating_ips" "is_floating_ips" {
		name   = ibm_is_floating_ip.testacc_floatingip.name
	  }
	  `, vpcname, subnetname, acc.ISZoneName, acc.ISCIDR, sshname, publicKey, instancename, acc.IsImage, acc.InstanceProfileName, acc.ISZoneName, fipname)
}
