// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIBMIsVirtualNetworkInterfaceFloatingIPDataSourceBasic(t *testing.T) {
	vpcname := fmt.Sprintf("tfp-vpc-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tfp-subnet-%d", acctest.RandIntRange(10, 100))
	vniname := fmt.Sprintf("tfp-createname-%d", acctest.RandIntRange(10, 100))
	floatingipname := fmt.Sprintf("tfp-reservedip-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsVirtualNetworkInterfaceFloatingIPDataSourceConfigBasic(vpcname, subnetname, vniname, floatingipname),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_virtual_network_interface_floating_ip.is_floating_ip", "address"),
					resource.TestCheckResourceAttrSet("data.ibm_is_virtual_network_interface_floating_ip.is_floating_ip", "crn"),
					resource.TestCheckResourceAttrSet("data.ibm_is_virtual_network_interface_floating_ip.is_floating_ip", "floating_ip"),
					resource.TestCheckResourceAttrSet("data.ibm_is_virtual_network_interface_floating_ip.is_floating_ip", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_virtual_network_interface_floating_ip.is_floating_ip", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_virtual_network_interface_floating_ip.is_floating_ip", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_is_virtual_network_interface_floating_ip.is_floating_ip", "virtual_network_interface"),
				),
			},
		},
	})
}
func TestAccIBMIsVirtualNetworkInterfaceFloatingIPDataSource404(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config:      testAccCheckIBMIsVirtualNetworkInterfaceFloatingIPDataSourceConfigVNI404(),
				ExpectError: regexp.MustCompile("GetVirtualNetworkInterfaceFloatingIPWithContext failed"),
			},
			{
				Config:      testAccCheckIBMIsVirtualNetworkInterfaceFloatingIPDataSourceConfigFIP404(),
				ExpectError: regexp.MustCompile("GetVirtualNetworkInterfaceFloatingIPWithContext failed"),
			},
		},
	})
}

func testAccCheckIBMIsVirtualNetworkInterfaceFloatingIPDataSourceConfigBasic(vpcname, subnetname, vniname, floatingipname string) string {
	return testAccCheckIBMIsVirtualNetworkInterfaceFloatingIPConfigBasic(vpcname, subnetname, vniname, floatingipname) + fmt.Sprintf(`
		data "ibm_is_virtual_network_interface_floating_ip" "is_floating_ip" {
			virtual_network_interface = ibm_is_virtual_network_interface_floating_ip.testacc_vni_floatingip.virtual_network_interface
			floating_ip = ibm_is_virtual_network_interface_floating_ip.testacc_vni_floatingip.floating_ip
		}
	`)
}
func testAccCheckIBMIsVirtualNetworkInterfaceFloatingIPDataSourceConfigVNI404() string {
	return fmt.Sprintf(`
		data "ibm_is_virtual_network_interface_floating_ip" "is_floating_ip" {
			virtual_network_interface = "%s"
			floating_ip = "%s"
		}
	`, acc.VNIId, acc.VNIId)
}
func testAccCheckIBMIsVirtualNetworkInterfaceFloatingIPDataSourceConfigFIP404() string {
	return fmt.Sprintf(`
		data "ibm_is_virtual_network_interface_floating_ip" "is_floating_ip" {
			virtual_network_interface = "%s"
			floating_ip = "%s"
		}
	`, acc.FloatingIpID, acc.FloatingIpID)
}
