// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"fmt"
	"strings"
	"testing"

	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"

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
	dataSourceName := "data.ibm_is_floating_ips.is_floating_ips"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIsFloatingIpsDataSourceConfigBasic(vpcname, subnetname, sshname, publicKey, instancename, fipname),
				Check: resource.ComposeTestCheckFunc(
					// Check basic data source attributes
					resource.TestCheckResourceAttrSet(dataSourceName, "id"),
					resource.TestCheckResourceAttrSet(dataSourceName, "floating_ips.#"),

					// Verify we can find our created floating IP in the list
					resource.TestCheckTypeSetElemNestedAttrs(dataSourceName, "floating_ips.*", map[string]string{
						"name": fipname,
					}),

					// Check detailed attributes of the first floating IP
					resource.TestCheckResourceAttrSet(dataSourceName, "floating_ips.0.id"),
					resource.TestCheckResourceAttrSet(dataSourceName, "floating_ips.0.name"),
					resource.TestCheckResourceAttrSet(dataSourceName, "floating_ips.0.address"),
					resource.TestCheckResourceAttrSet(dataSourceName, "floating_ips.0.status"),
					resource.TestCheckResourceAttrSet(dataSourceName, "floating_ips.0.crn"),
					resource.TestCheckResourceAttrSet(dataSourceName, "floating_ips.0.created_at"),
					resource.TestCheckResourceAttrSet(dataSourceName, "floating_ips.0.href"),

					// Check zone information
					resource.TestCheckResourceAttrSet(dataSourceName, "floating_ips.0.zone.#"),
					resource.TestCheckResourceAttrSet(dataSourceName, "floating_ips.0.zone.0.name"),
					resource.TestCheckResourceAttrSet(dataSourceName, "floating_ips.0.zone.0.href"),

					// Check resource group
					resource.TestCheckResourceAttrSet(dataSourceName, "floating_ips.0.resource_group.#"),
					resource.TestCheckResourceAttrSet(dataSourceName, "floating_ips.0.resource_group.0.id"),
					resource.TestCheckResourceAttrSet(dataSourceName, "floating_ips.0.resource_group.0.name"),
					resource.TestCheckResourceAttrSet(dataSourceName, "floating_ips.0.resource_group.0.href"),

					// Check target information if it exists
					resource.TestCheckResourceAttrSet(dataSourceName, "floating_ips.0.target.#"),
					resource.TestCheckResourceAttrSet(dataSourceName, "floating_ips.0.target.0.id"),
					resource.TestCheckResourceAttrSet(dataSourceName, "floating_ips.0.target.0.name"),
					resource.TestCheckResourceAttrSet(dataSourceName, "floating_ips.0.target.0.href"),
					resource.TestCheckResourceAttrSet(dataSourceName, "floating_ips.0.target.0.resource_type"),

					// Check primary_ip information
					resource.TestCheckResourceAttrSet(dataSourceName, "floating_ips.0.target.0.primary_ip.#"),
					resource.TestCheckResourceAttrSet(dataSourceName, "floating_ips.0.target.0.primary_ip.0.address"),
					resource.TestCheckResourceAttrSet(dataSourceName, "floating_ips.0.target.0.primary_ip.0.name"),
					resource.TestCheckResourceAttrSet(dataSourceName, "floating_ips.0.target.0.primary_ip.0.href"),
					resource.TestCheckResourceAttrSet(dataSourceName, "floating_ips.0.target.0.primary_ip.0.reserved_ip"),
					resource.TestCheckResourceAttrSet(dataSourceName, "floating_ips.0.target.0.primary_ip.0.resource_type"),
				),
			},
		},
	})
}

func testAccCheckIBMIsFloatingIpsDataSourceConfigBasic(vpcname, subnetname, sshname, publicKey, instancename, fipname string) string {
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
		vpc  = ibm_is_vpc.testacc_vpc.id
		zone = "%s"
		keys = [ibm_is_ssh_key.testacc_sshkey.id]
	}
	  
	resource "ibm_is_floating_ip" "testacc_floatingip" {
		name   = "%s"
		target = ibm_is_instance.testacc_instance.primary_network_interface[0].id
	}

	data "ibm_is_floating_ips" "is_floating_ips" {
		name = ibm_is_floating_ip.testacc_floatingip.name
		depends_on = [ibm_is_floating_ip.testacc_floatingip]
	}
	`, vpcname, subnetname, acc.ISZoneName, acc.ISCIDR, sshname, publicKey, instancename, acc.IsImage, acc.InstanceProfileName, acc.ISZoneName, fipname)
}

func TestAccIBMIsFloatingIpsDataSourceWithResourceGroup(t *testing.T) {
	vpcname := fmt.Sprintf("tfip-vpc-rg-%d", acctest.RandIntRange(10, 100))
	fipname := fmt.Sprintf("tfip-rg-%d", acctest.RandIntRange(10, 100))
	dataSourceName := "data.ibm_is_floating_ips.is_floating_ips_with_rg"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIsFloatingIpsDataSourceWithResourceGroup(vpcname, fipname),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(dataSourceName, "id"),
					resource.TestCheckResourceAttrSet(dataSourceName, "floating_ips.#"),
					resource.TestCheckResourceAttrSet(dataSourceName, "floating_ips.0.resource_group.0.id"),
					resource.TestCheckResourceAttrSet(dataSourceName, "floating_ips.0.resource_group.0.name"),
				),
			},
		},
	})
}

func testAccCheckIBMIsFloatingIpsDataSourceWithResourceGroup(vpcname, fipname string) string {
	return fmt.Sprintf(`
	data "ibm_resource_group" "default" {
		name = "Default"
	}

	resource "ibm_is_vpc" "testacc_vpc" {
		name = "%s"
	}

	resource "ibm_is_floating_ip" "testacc_floatingip_rg" {
		name = "%s"
		zone = "%s"
		resource_group = data.ibm_resource_group.default.id
	}

	data "ibm_is_floating_ips" "is_floating_ips_with_rg" {
		resource_group = data.ibm_resource_group.default.id
		depends_on = [ibm_is_floating_ip.testacc_floatingip_rg]
	}
	`, vpcname, fipname, acc.ISZoneName)
}
