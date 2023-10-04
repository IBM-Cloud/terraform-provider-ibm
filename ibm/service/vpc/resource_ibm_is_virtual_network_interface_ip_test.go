// Copyright IBM Corp. 2023 All Rights VirtualNetworkInterface.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/vpc-go-sdk/vpcv1"
)

func TestAccIBMIsVirtualNetworkInterfaceIPBasic(t *testing.T) {
	var conf vpcv1.ReservedIPReference
	vpcname := fmt.Sprintf("tfp-vpc-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tfp-subnet-%d", acctest.RandIntRange(10, 100))
	vniname := fmt.Sprintf("tfp-createname-%d", acctest.RandIntRange(10, 100))
	reservedipname := fmt.Sprintf("tfp-reservedip-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIsVirtualNetworkInterfaceIPDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsVirtualNetworkInterfaceIPConfigBasic(vpcname, subnetname, vniname, reservedipname),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIsVirtualNetworkInterfaceIPExists("ibm_is_virtual_network_interface_ip.testacc_vni_reservedip", conf),
					resource.TestCheckResourceAttr("ibm_is_virtual_network_interface_ip.testacc_vni_reservedip", "resource_type", "subnet_reserved_ip"),
					resource.TestCheckResourceAttrSet("ibm_is_virtual_network_interface_ip.testacc_vni_reservedip", "address"),
					resource.TestCheckResourceAttrSet("ibm_is_virtual_network_interface_ip.testacc_vni_reservedip", "href"),
					resource.TestCheckResourceAttrSet("ibm_is_virtual_network_interface_ip.testacc_vni_reservedip", "id"),
					resource.TestCheckResourceAttrSet("ibm_is_virtual_network_interface_ip.testacc_vni_reservedip", "name"),
					resource.TestCheckResourceAttrSet("ibm_is_virtual_network_interface_ip.testacc_vni_reservedip", "reserved_ip"),
					resource.TestCheckResourceAttrSet("ibm_is_virtual_network_interface_ip.testacc_vni_reservedip", "virtual_network_interface"),
				),
			},
		},
	})
}

func testAccCheckIBMIsVirtualNetworkInterfaceIPConfigBasic(vpcname, subnetname, vniname, reservedipname string) string {
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

	resource "ibm_is_subnet_reserved_ip" "testacc_reservedip" {
		subnet = ibm_is_subnet.testacc_subnet.id
		name = "%s"
	}
	resource "ibm_is_virtual_network_interface_ip" "testacc_vni_reservedip" {
		virtual_network_interface = ibm_is_virtual_network_interface.testacc_vni.id
		reserved_ip = ibm_is_subnet_reserved_ip.testacc_reservedip.reserved_ip
	}

	`, vpcname, subnetname, acc.ISZoneName, vniname, reservedipname)
}

func testAccCheckIBMIsVirtualNetworkInterfaceIPExists(n string, obj vpcv1.ReservedIPReference) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		vpcClient, err := vpcClient(acc.TestAccProvider.Meta())
		if err != nil {
			return err
		}

		getSubnetVirtualNetworkInterfaceIPOptions := &vpcv1.GetVirtualNetworkInterfaceIPOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}
		getSubnetVirtualNetworkInterfaceIPOptions.SetVirtualNetworkInterfaceID(parts[0])
		getSubnetVirtualNetworkInterfaceIPOptions.SetID(parts[1])

		reservedIP, _, err := vpcClient.GetVirtualNetworkInterfaceIP(getSubnetVirtualNetworkInterfaceIPOptions)
		if err != nil {
			return err
		}
		obj = *reservedIP
		return nil
	}
}

func testAccCheckIBMIsVirtualNetworkInterfaceIPDestroy(s *terraform.State) error {
	vpcClient, err := vpcClient(acc.TestAccProvider.Meta())
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_is_virtual_network_interface_ip" {
			continue
		}
		getSubnetVirtualNetworkInterfaceIPOptions := &vpcv1.GetVirtualNetworkInterfaceIPOptions{}
		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}
		getSubnetVirtualNetworkInterfaceIPOptions.SetVirtualNetworkInterfaceID(parts[0])
		getSubnetVirtualNetworkInterfaceIPOptions.SetID(parts[1])

		// Try to find the key
		_, response, err := vpcClient.GetVirtualNetworkInterfaceIP(getSubnetVirtualNetworkInterfaceIPOptions)

		if err == nil {
			return fmt.Errorf("VirtualNetworkInterfaceIP still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for VirtualNetworkInterfaceIP (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}
	return nil
}
