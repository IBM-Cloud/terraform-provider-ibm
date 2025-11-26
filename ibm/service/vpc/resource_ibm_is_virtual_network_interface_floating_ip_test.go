// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/vpc"
	"github.com/IBM/vpc-go-sdk/vpcv1"
)

func TestAccIBMIsVirtualNetworkInterfaceFloatingIPBasic(t *testing.T) {
	var conf vpcv1.FloatingIPReference
	vpcname := fmt.Sprintf("tfp-vpc-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tfp-subnet-%d", acctest.RandIntRange(10, 100))
	vniname := fmt.Sprintf("tfp-createname-%d", acctest.RandIntRange(10, 100))
	floatingipname := fmt.Sprintf("tfp-reservedip-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIsVirtualNetworkInterfaceFloatingIPDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsVirtualNetworkInterfaceFloatingIPConfigBasic(vpcname, subnetname, vniname, floatingipname),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIsVirtualNetworkInterfaceFloatingIPExists("ibm_is_virtual_network_interface_floating_ip.testacc_vni_floatingip", conf),
					resource.TestCheckResourceAttrSet("ibm_is_virtual_network_interface_floating_ip.testacc_vni_floatingip", "address"),
					resource.TestCheckResourceAttrSet("ibm_is_virtual_network_interface_floating_ip.testacc_vni_floatingip", "crn"),
					resource.TestCheckResourceAttrSet("ibm_is_virtual_network_interface_floating_ip.testacc_vni_floatingip", "href"),
					resource.TestCheckResourceAttrSet("ibm_is_virtual_network_interface_floating_ip.testacc_vni_floatingip", "id"),
					resource.TestCheckResourceAttrSet("ibm_is_virtual_network_interface_floating_ip.testacc_vni_floatingip", "name"),
					resource.TestCheckResourceAttrSet("ibm_is_virtual_network_interface_floating_ip.testacc_vni_floatingip", "virtual_network_interface"),
					resource.TestCheckResourceAttrSet("ibm_is_virtual_network_interface_floating_ip.testacc_vni_floatingip", "floating_ip"),
				),
			},
		},
	})
}

func testAccCheckIBMIsVirtualNetworkInterfaceFloatingIPConfigBasic(vpcname, subnetname, vniname, floatingipname string) string {
	return fmt.Sprintf(`

	resource "ibm_is_vpc" "testacc_vpc" {
		name = "%s"
	}
	
	resource "ibm_is_subnet" "testacc_subnet" {
		name = "%s"
		vpc = ibm_is_vpc.testacc_vpc.id
		zone = "%s"
		total_ipv4_address_count = 16
	
	}
	
	resource "ibm_is_virtual_network_interface" "testacc_vni"{
		name = "%s"
		subnet = ibm_is_subnet.testacc_subnet.id
	}
	
	resource "ibm_is_floating_ip" "testacc_floatingip" {
		name = "%s"
		zone = ibm_is_subnet.testacc_subnet.zone
	}
	resource "ibm_is_virtual_network_interface_floating_ip" "testacc_vni_floatingip" {
		virtual_network_interface = ibm_is_virtual_network_interface.testacc_vni.id
		floating_ip = ibm_is_floating_ip.testacc_floatingip.id
	}
	`, vpcname, subnetname, acc.ISZoneName, vniname, floatingipname)
}
func testAccCheckIBMIsVirtualNetworkInterfaceFloatingIPExists(n string, obj vpcv1.FloatingIPReference) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		sess, err := vpcClient(acc.TestAccProvider.Meta())
		if err != nil {
			return err
		}
		vniid, fipid, err := vpc.ParseVNIFloatingIpTerraformID(rs.Primary.ID)
		if err != nil {
			return err
		}
		getVirtualNetworkInterfaceFloatingIPOptions := &vpcv1.GetNetworkInterfaceFloatingIPOptions{}

		getVirtualNetworkInterfaceFloatingIPOptions.SetVirtualNetworkInterfaceID(vniid)
		getVirtualNetworkInterfaceFloatingIPOptions.SetID(fipid)

		floatingIP, _, err := sess.GetNetworkInterfaceFloatingIP(getVirtualNetworkInterfaceFloatingIPOptions)
		if err != nil {
			return err
		}

		obj = *floatingIP
		return nil
	}
}

func testAccCheckIBMIsVirtualNetworkInterfaceFloatingIPDestroy(s *terraform.State) error {
	sess, err := vpcClient(acc.TestAccProvider.Meta())
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_is_virtual_network_interface_floating_ip" {
			continue
		}
		vniid, fipid, err := vpc.ParseVNIFloatingIpTerraformID(rs.Primary.ID)
		if err != nil {
			return err
		}
		getVirtualNetworkInterfaceFloatingIPOptions := &vpcv1.GetNetworkInterfaceFloatingIPOptions{}

		getVirtualNetworkInterfaceFloatingIPOptions.SetVirtualNetworkInterfaceID(vniid)
		getVirtualNetworkInterfaceFloatingIPOptions.SetID(fipid)

		// Try to find the key
		_, response, err := sess.GetNetworkInterfaceFloatingIP(getVirtualNetworkInterfaceFloatingIPOptions)

		if err == nil {
			return fmt.Errorf("VirtualNetworkInterfaceFloatingIP still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for VirtualNetworkInterfaceFloatingIP (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
