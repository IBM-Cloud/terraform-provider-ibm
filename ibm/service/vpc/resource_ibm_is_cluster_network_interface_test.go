// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/IBM/vpc-go-sdk/vpcv1"
	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/conns"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/flex"
)

func TestAccIBMIsClusterNetworkInterfaceBasic(t *testing.T) {
	var conf vpcv1.ClusterNetworkInterface
	vpcname := fmt.Sprintf("tf-vpc-%d", acctest.RandIntRange(10, 100))
	clustersubnetname := fmt.Sprintf("tf-clustersubnet-%d", acctest.RandIntRange(10, 100))
	clustersubnetreservedipname := fmt.Sprintf("tf-clustersubnet-reservedip-%d", acctest.RandIntRange(10, 100))
	clusterinterfacename := fmt.Sprintf("tf-clusterinterface-%d", acctest.RandIntRange(10, 100))
	clusterinterfacenameupdated := fmt.Sprintf("tf-clusterinterfaceupdated-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIsClusterNetworkInterfaceDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsClusterNetworkInterfaceConfigBasic(vpcname, clustersubnetname, clustersubnetreservedipname, clusterinterfacename),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIsClusterNetworkInterfaceExists("ibm_is_cluster_network_interface.is_cluster_network_interface_instance", conf),
					resource.TestCheckResourceAttr("ibm_is_vpc.is_vpc", "name", vpcname),
					resource.TestCheckResourceAttrSet("ibm_is_vpc.is_vpc", "id"),
					resource.TestCheckResourceAttrSet("ibm_is_cluster_network.is_cluster_network_instance", "id"),
					resource.TestCheckResourceAttrSet("ibm_is_cluster_network_subnet.is_cluster_network_subnet_instance", "id"),
					resource.TestCheckResourceAttrSet("ibm_is_cluster_network_subnet_reserved_ip.is_cluster_network_subnet_reserved_ip_instance", "id"),
					resource.TestCheckResourceAttrSet("ibm_is_cluster_network_interface.is_cluster_network_interface_instance", "id"),
					resource.TestCheckResourceAttrSet("ibm_is_cluster_network_interface.is_cluster_network_interface_instance", "allow_ip_spoofing"),
					resource.TestCheckResourceAttrSet("ibm_is_cluster_network_interface.is_cluster_network_interface_instance", "auto_delete"),
					resource.TestCheckResourceAttrSet("ibm_is_cluster_network_interface.is_cluster_network_interface_instance", "cluster_network_id"),
					resource.TestCheckResourceAttrSet("ibm_is_cluster_network_interface.is_cluster_network_interface_instance", "cluster_network_interface_id"),
					resource.TestCheckResourceAttrSet("ibm_is_cluster_network_interface.is_cluster_network_interface_instance", "created_at"),
					resource.TestCheckResourceAttrSet("ibm_is_cluster_network_interface.is_cluster_network_interface_instance", "enable_infrastructure_nat"),
					resource.TestCheckResourceAttrSet("ibm_is_cluster_network_interface.is_cluster_network_interface_instance", "href"),
					resource.TestCheckResourceAttrSet("ibm_is_cluster_network_interface.is_cluster_network_interface_instance", "lifecycle_state"),
					resource.TestCheckResourceAttr("ibm_is_cluster_network_interface.is_cluster_network_interface_instance", "lifecycle_state", "stable"),
					resource.TestCheckResourceAttrSet("ibm_is_cluster_network_interface.is_cluster_network_interface_instance", "mac_address"),
					resource.TestCheckResourceAttrSet("ibm_is_cluster_network_interface.is_cluster_network_interface_instance", "resource_type"),
					resource.TestCheckResourceAttr("ibm_is_cluster_network_interface.is_cluster_network_interface_instance", "resource_type", "cluster_network_interface"),
					resource.TestCheckResourceAttrSet("ibm_is_cluster_network_interface.is_cluster_network_interface_instance", "vpc.#"),
					resource.TestCheckResourceAttrSet("ibm_is_cluster_network_interface.is_cluster_network_interface_instance", "primary_ip.#"),
					resource.TestCheckResourceAttrSet("ibm_is_cluster_network_interface.is_cluster_network_interface_instance", "zone.#"),
					resource.TestCheckResourceAttrSet("ibm_is_cluster_network_interface.is_cluster_network_interface_instance", "subnet.#"),
					resource.TestCheckResourceAttrSet("ibm_is_cluster_network_interface.is_cluster_network_interface_instance", "name"),
					resource.TestCheckResourceAttr("ibm_is_cluster_network_interface.is_cluster_network_interface_instance", "name", clusterinterfacename),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMIsClusterNetworkInterfaceConfigBasic(vpcname, clustersubnetname, clustersubnetreservedipname, clusterinterfacenameupdated),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIsClusterNetworkInterfaceExists("ibm_is_cluster_network_interface.is_cluster_network_interface_instance", conf),
					resource.TestCheckResourceAttr("ibm_is_vpc.is_vpc", "name", vpcname),
					resource.TestCheckResourceAttrSet("ibm_is_vpc.is_vpc", "id"),
					resource.TestCheckResourceAttrSet("ibm_is_cluster_network.is_cluster_network_instance", "id"),
					resource.TestCheckResourceAttrSet("ibm_is_cluster_network_subnet.is_cluster_network_subnet_instance", "id"),
					resource.TestCheckResourceAttrSet("ibm_is_cluster_network_subnet_reserved_ip.is_cluster_network_subnet_reserved_ip_instance", "id"),
					resource.TestCheckResourceAttrSet("ibm_is_cluster_network_interface.is_cluster_network_interface_instance", "id"),
					resource.TestCheckResourceAttrSet("ibm_is_cluster_network_interface.is_cluster_network_interface_instance", "allow_ip_spoofing"),
					resource.TestCheckResourceAttrSet("ibm_is_cluster_network_interface.is_cluster_network_interface_instance", "auto_delete"),
					resource.TestCheckResourceAttrSet("ibm_is_cluster_network_interface.is_cluster_network_interface_instance", "cluster_network_id"),
					resource.TestCheckResourceAttrSet("ibm_is_cluster_network_interface.is_cluster_network_interface_instance", "cluster_network_interface_id"),
					resource.TestCheckResourceAttrSet("ibm_is_cluster_network_interface.is_cluster_network_interface_instance", "created_at"),
					resource.TestCheckResourceAttrSet("ibm_is_cluster_network_interface.is_cluster_network_interface_instance", "enable_infrastructure_nat"),
					resource.TestCheckResourceAttrSet("ibm_is_cluster_network_interface.is_cluster_network_interface_instance", "href"),
					resource.TestCheckResourceAttrSet("ibm_is_cluster_network_interface.is_cluster_network_interface_instance", "lifecycle_state"),
					resource.TestCheckResourceAttr("ibm_is_cluster_network_interface.is_cluster_network_interface_instance", "lifecycle_state", "stable"),
					resource.TestCheckResourceAttrSet("ibm_is_cluster_network_interface.is_cluster_network_interface_instance", "mac_address"),
					resource.TestCheckResourceAttrSet("ibm_is_cluster_network_interface.is_cluster_network_interface_instance", "resource_type"),
					resource.TestCheckResourceAttr("ibm_is_cluster_network_interface.is_cluster_network_interface_instance", "resource_type", "cluster_network_interface"),
					resource.TestCheckResourceAttrSet("ibm_is_cluster_network_interface.is_cluster_network_interface_instance", "vpc.#"),
					resource.TestCheckResourceAttrSet("ibm_is_cluster_network_interface.is_cluster_network_interface_instance", "primary_ip.#"),
					resource.TestCheckResourceAttrSet("ibm_is_cluster_network_interface.is_cluster_network_interface_instance", "zone.#"),
					resource.TestCheckResourceAttrSet("ibm_is_cluster_network_interface.is_cluster_network_interface_instance", "subnet.#"),
					resource.TestCheckResourceAttrSet("ibm_is_cluster_network_interface.is_cluster_network_interface_instance", "name"),
					resource.TestCheckResourceAttr("ibm_is_cluster_network_interface.is_cluster_network_interface_instance", "name", clusterinterfacenameupdated),
				),
			},
		},
	})
}

func testAccCheckIBMIsClusterNetworkInterfaceConfigBasic(vpcname, clustersubnetname, clustersubnetreservedipname, clusternetworkinterfacename string) string {
	return fmt.Sprintf(`
		resource "ibm_is_vpc" "is_vpc" {
  			name = "%s"
		}
		resource "ibm_is_cluster_network" "is_cluster_network_instance" {
			profile = "%s"
			vpc {
				id = ibm_is_vpc.is_vpc.id
			}
			zone  = "%s"
		}
		resource "ibm_is_cluster_network_subnet" "is_cluster_network_subnet_instance" {
			cluster_network_id = ibm_is_cluster_network.is_cluster_network_instance.id
			name = "%s"
			total_ipv4_address_count = 64
		}
		resource "ibm_is_cluster_network_subnet_reserved_ip" "is_cluster_network_subnet_reserved_ip_instance" {
			cluster_network_id 			= ibm_is_cluster_network.is_cluster_network_instance.id
			cluster_network_subnet_id 	= ibm_is_cluster_network_subnet.is_cluster_network_subnet_instance.cluster_network_subnet_id
			address 					= "${replace(ibm_is_cluster_network_subnet.is_cluster_network_subnet_instance.ipv4_cidr_block, "0/26", "11")}"
  			name						= "%s"
		}
		resource "ibm_is_cluster_network_interface" "is_cluster_network_interface_instance" {
			cluster_network_id = ibm_is_cluster_network.is_cluster_network_instance.id
			name = "%s"
			primary_ip {
				id = ibm_is_cluster_network_subnet_reserved_ip.is_cluster_network_subnet_reserved_ip_instance.cluster_network_subnet_reserved_ip_id
			}
			subnet {
				id = ibm_is_cluster_network_subnet.is_cluster_network_subnet_instance.cluster_network_subnet_id
			}
		}
	`, vpcname, acc.ISClusterNetworkProfileName, acc.ISZoneName, clustersubnetname, clustersubnetreservedipname, clusternetworkinterfacename)
}

func testAccCheckIBMIsClusterNetworkInterfaceExists(n string, obj vpcv1.ClusterNetworkInterface) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		vpcClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).VpcV1API()
		if err != nil {
			return err
		}

		getClusterNetworkInterfaceOptions := &vpcv1.GetClusterNetworkInterfaceOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getClusterNetworkInterfaceOptions.SetClusterNetworkID(parts[0])
		getClusterNetworkInterfaceOptions.SetID(parts[1])

		clusterNetworkInterface, _, err := vpcClient.GetClusterNetworkInterface(getClusterNetworkInterfaceOptions)
		if err != nil {
			return err
		}

		obj = *clusterNetworkInterface
		return nil
	}
}

func testAccCheckIBMIsClusterNetworkInterfaceDestroy(s *terraform.State) error {
	vpcClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).VpcV1API()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_is_cluster_network_interface" {
			continue
		}

		getClusterNetworkInterfaceOptions := &vpcv1.GetClusterNetworkInterfaceOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getClusterNetworkInterfaceOptions.SetClusterNetworkID(parts[0])
		getClusterNetworkInterfaceOptions.SetID(parts[1])

		// Try to find the key
		_, response, err := vpcClient.GetClusterNetworkInterface(getClusterNetworkInterfaceOptions)

		if err == nil {
			return fmt.Errorf("ClusterNetworkInterface still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for ClusterNetworkInterface (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
