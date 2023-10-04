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
	subnetID := fmt.Sprintf("tf_subnet_id_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIsVirtualNetworkInterfaceIPDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsVirtualNetworkInterfaceIPConfigBasic(subnetID),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIsVirtualNetworkInterfaceIPExists("ibm_is_virtual_network_interface_ip.is_virtual_network_interface_ip", conf),
					resource.TestCheckResourceAttr("ibm_is_virtual_network_interface_ip.is_virtual_network_interface_ip", "subnet_id", subnetID),
				),
			},
		},
	})
}

func TestAccIBMIsVirtualNetworkInterfaceIPAllArgs(t *testing.T) {
	var conf vpcv1.ReservedIPReference
	subnetID := fmt.Sprintf("tf_subnet_id_%d", acctest.RandIntRange(10, 100))
	address := fmt.Sprintf("tf_address_%d", acctest.RandIntRange(10, 100))
	autoDelete := "false"
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	addressUpdate := fmt.Sprintf("tf_address_%d", acctest.RandIntRange(10, 100))
	autoDeleteUpdate := "true"
	nameUpdate := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIsVirtualNetworkInterfaceIPDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsVirtualNetworkInterfaceIPConfig(subnetID, address, autoDelete, name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIsVirtualNetworkInterfaceIPExists("ibm_is_virtual_network_interface_ip.is_virtual_network_interface_ip", conf),
					resource.TestCheckResourceAttr("ibm_is_virtual_network_interface_ip.is_virtual_network_interface_ip", "subnet_id", subnetID),
					resource.TestCheckResourceAttr("ibm_is_virtual_network_interface_ip.is_virtual_network_interface_ip", "address", address),
					resource.TestCheckResourceAttr("ibm_is_virtual_network_interface_ip.is_virtual_network_interface_ip", "auto_delete", autoDelete),
					resource.TestCheckResourceAttr("ibm_is_virtual_network_interface_ip.is_virtual_network_interface_ip", "name", name),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMIsVirtualNetworkInterfaceIPConfig(subnetID, addressUpdate, autoDeleteUpdate, nameUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_is_virtual_network_interface_ip.is_virtual_network_interface_ip", "subnet_id", subnetID),
					resource.TestCheckResourceAttr("ibm_is_virtual_network_interface_ip.is_virtual_network_interface_ip", "address", addressUpdate),
					resource.TestCheckResourceAttr("ibm_is_virtual_network_interface_ip.is_virtual_network_interface_ip", "auto_delete", autoDeleteUpdate),
					resource.TestCheckResourceAttr("ibm_is_virtual_network_interface_ip.is_virtual_network_interface_ip", "name", nameUpdate),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_is_virtual_network_interface_ip.is_virtual_network_interface_ip",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIBMIsVirtualNetworkInterfaceIPConfigBasic(subnetID string) string {
	return fmt.Sprintf(`
		resource "ibm_is_virtual_network_interface_ip" "is_virtual_network_interface_ip_instance" {
			subnet_id = "%s"
		}
	`, subnetID)
}

func testAccCheckIBMIsVirtualNetworkInterfaceIPConfig(subnetID string, address string, autoDelete string, name string) string {
	return fmt.Sprintf(`

		resource "ibm_is_virtual_network_interface_ip" "is_virtual_network_interface_ip_instance" {
			subnet_id = "%s"
			address = "%s"
			auto_delete = %s
			name = "%s"
			target {
				crn = "crn:v1:bluemix:public:is:us-south:a/123456::endpoint-gateway:r134-d7cc5196-9864-48c4-82d8-3f30da41fcc5"
				deleted {
					more_info = "https://cloud.ibm.com/apidocs/vpc#deleted-resources"
				}
				href = "https://us-south.iaas.cloud.ibm.com/v1/endpoint_gateways/r134-d7cc5196-9864-48c4-82d8-3f30da41fcc5"
				id = "r134-d7cc5196-9864-48c4-82d8-3f30da41fcc5"
				name = "my-endpoint-gateway"
				resource_type = "endpoint_gateway"
			}
		}
	`, subnetID, address, autoDelete, name)
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
