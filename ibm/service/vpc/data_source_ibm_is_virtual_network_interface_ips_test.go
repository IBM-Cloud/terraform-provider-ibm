// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIBMIsVirtualNetworkInterfaceIpsDataSourceBasic(t *testing.T) {
	vpcname := fmt.Sprintf("tfp-vpc-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tfp-subnet-%d", acctest.RandIntRange(10, 100))
	vniname := fmt.Sprintf("tfp-createname-%d", acctest.RandIntRange(10, 100))
	reservedipname := fmt.Sprintf("tfp-reservedip-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsVirtualNetworkInterfaceIpsDataSourceConfigBasic(vpcname, subnetname, vniname, reservedipname),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_virtual_network_interface_ips.is_reserved_ips", "reserved_ips.0.id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_virtual_network_interface_ips.is_reserved_ips", "reserved_ips.0.name"),
					resource.TestCheckResourceAttrSet("data.ibm_is_virtual_network_interface_ips.is_reserved_ips", "reserved_ips.0.href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_virtual_network_interface_ips.is_reserved_ips", "reserved_ips.0.address"),
					resource.TestCheckResourceAttrSet("data.ibm_is_virtual_network_interface_ips.is_reserved_ips", "reserved_ips.0.resource_type"),
					resource.TestCheckResourceAttrSet("data.ibm_is_virtual_network_interface_ips.is_reserved_ips", "reserved_ips.0.reserved_ip"),
				),
			},
		},
	})
}

func testAccCheckIBMIsVirtualNetworkInterfaceIpsDataSourceConfigBasic(vpcname, subnetname, vniname, reservedipname string) string {
	return testAccCheckIBMIsVirtualNetworkInterfaceIPConfigBasic(vpcname, subnetname, vniname, reservedipname) + fmt.Sprintf(`
		data "ibm_is_virtual_network_interface_ips" "is_reserved_ips" {
			virtual_network_interface = ibm_is_virtual_network_interface_ip.testacc_vni_reservedip.virtual_network_interface
		}
	`)
}
