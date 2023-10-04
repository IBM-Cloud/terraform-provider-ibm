// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIBMIsVirtualNetworkInterfaceFloatingIPsDataSourceBasic(t *testing.T) {
	vpcname := fmt.Sprintf("tfp-vpc-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tfp-subnet-%d", acctest.RandIntRange(10, 100))
	vniname := fmt.Sprintf("tfp-createname-%d", acctest.RandIntRange(10, 100))
	floatingipname := fmt.Sprintf("tfp-reservedip-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsVirtualNetworkInterfaceFloatingIPsDataSourceConfigBasic(vpcname, subnetname, vniname, floatingipname),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_virtual_network_interface_floating_ips.is_floating_ips", "floating_ips.0.address"),
					resource.TestCheckResourceAttrSet("data.ibm_is_virtual_network_interface_floating_ips.is_floating_ips", "floating_ips.0.crn"),
					resource.TestCheckResourceAttrSet("data.ibm_is_virtual_network_interface_floating_ips.is_floating_ips", "floating_ips.0.href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_virtual_network_interface_floating_ips.is_floating_ips", "floating_ips.0.id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_virtual_network_interface_floating_ips.is_floating_ips", "floating_ips.0.name"),
				),
			},
		},
	})
}

func testAccCheckIBMIsVirtualNetworkInterfaceFloatingIPsDataSourceConfigBasic(vpcname, subnetname, vniname, floatingipname string) string {
	return testAccCheckIBMIsVirtualNetworkInterfaceFloatingIPConfigBasic(vpcname, subnetname, vniname, floatingipname) + fmt.Sprintf(`
		data "ibm_is_virtual_network_interface_floating_ips" "is_floating_ips" {
			virtual_network_interface = ibm_is_virtual_network_interface_floating_ip.testacc_vni_floatingip.virtual_network_interface
		}
	`)
}
