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

func TestAccIBMIsClusterNetworkSubnetReservedIPBasic(t *testing.T) {
	var conf vpcv1.ClusterNetworkSubnetReservedIP
	vpcname := fmt.Sprintf("tf-vpc-%d", acctest.RandIntRange(10, 100))
	clustersubnetname := fmt.Sprintf("tf-clustersubnet-%d", acctest.RandIntRange(10, 100))
	clustersubnetreservedipname := fmt.Sprintf("tf-clustersubnet-reservedip-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIsClusterNetworkSubnetReservedIPDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsClusterNetworkSubnetReservedIPConfigBasic(vpcname, clustersubnetname, clustersubnetreservedipname),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIsClusterNetworkSubnetReservedIPExists("ibm_is_cluster_network_subnet_reserved_ip.is_cluster_network_subnet_reserved_ip_instance", conf),
					resource.TestCheckResourceAttr("ibm_is_vpc.is_vpc", "name", vpcname),
					resource.TestCheckResourceAttrSet("ibm_is_vpc.is_vpc", "id"),
					resource.TestCheckResourceAttrSet("ibm_is_cluster_network.is_cluster_network_instance", "id"),
					resource.TestCheckResourceAttrSet("ibm_is_cluster_network_subnet.is_cluster_network_subnet_instance", "id"),
					resource.TestCheckResourceAttr("ibm_is_cluster_network_subnet.is_cluster_network_subnet_instance", "name", clustersubnetname),
					resource.TestCheckResourceAttrSet("ibm_is_cluster_network_subnet_reserved_ip.is_cluster_network_subnet_reserved_ip_instance", "id"),
					resource.TestCheckResourceAttr("ibm_is_cluster_network_subnet_reserved_ip.is_cluster_network_subnet_reserved_ip_instance", "name", clustersubnetreservedipname),
					resource.TestCheckResourceAttrSet("ibm_is_cluster_network_subnet_reserved_ip.is_cluster_network_subnet_reserved_ip_instance", "address"),
					resource.TestCheckResourceAttrSet("ibm_is_cluster_network_subnet_reserved_ip.is_cluster_network_subnet_reserved_ip_instance", "auto_delete"),
					resource.TestCheckResourceAttrSet("ibm_is_cluster_network_subnet_reserved_ip.is_cluster_network_subnet_reserved_ip_instance", "cluster_network_id"),
					resource.TestCheckResourceAttrSet("ibm_is_cluster_network_subnet_reserved_ip.is_cluster_network_subnet_reserved_ip_instance", "cluster_network_subnet_id"),
					resource.TestCheckResourceAttrSet("ibm_is_cluster_network_subnet_reserved_ip.is_cluster_network_subnet_reserved_ip_instance", "cluster_network_subnet_reserved_ip_id"),
					resource.TestCheckResourceAttrSet("ibm_is_cluster_network_subnet_reserved_ip.is_cluster_network_subnet_reserved_ip_instance", "created_at"),
					resource.TestCheckResourceAttrSet("ibm_is_cluster_network_subnet_reserved_ip.is_cluster_network_subnet_reserved_ip_instance", "href"),
					resource.TestCheckResourceAttrSet("ibm_is_cluster_network_subnet_reserved_ip.is_cluster_network_subnet_reserved_ip_instance", "lifecycle_state"),
					resource.TestCheckResourceAttr("ibm_is_cluster_network_subnet_reserved_ip.is_cluster_network_subnet_reserved_ip_instance", "lifecycle_state", "stable"),
					resource.TestCheckResourceAttrSet("ibm_is_cluster_network_subnet_reserved_ip.is_cluster_network_subnet_reserved_ip_instance", "resource_type"),
					resource.TestCheckResourceAttr("ibm_is_cluster_network_subnet_reserved_ip.is_cluster_network_subnet_reserved_ip_instance", "resource_type", "cluster_network_subnet_reserved_ip"),
					resource.TestCheckResourceAttrSet("ibm_is_cluster_network_subnet_reserved_ip.is_cluster_network_subnet_reserved_ip_instance", "owner"),
				),
			},
		},
	})
}

func testAccCheckIBMIsClusterNetworkSubnetReservedIPConfigBasic(vpcname, clustersubnetname, clustersubnetreservedipname string) string {
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
	`, vpcname, acc.ISClusterNetworkProfileName, acc.ISZoneName, clustersubnetname, clustersubnetreservedipname)
}

func testAccCheckIBMIsClusterNetworkSubnetReservedIPExists(n string, obj vpcv1.ClusterNetworkSubnetReservedIP) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		vpcClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).VpcV1API()
		if err != nil {
			return err
		}

		getClusterNetworkSubnetReservedIPOptions := &vpcv1.GetClusterNetworkSubnetReservedIPOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getClusterNetworkSubnetReservedIPOptions.SetClusterNetworkID(parts[0])
		getClusterNetworkSubnetReservedIPOptions.SetClusterNetworkSubnetID(parts[1])
		getClusterNetworkSubnetReservedIPOptions.SetID(parts[2])

		clusterNetworkSubnetReservedIP, _, err := vpcClient.GetClusterNetworkSubnetReservedIP(getClusterNetworkSubnetReservedIPOptions)
		if err != nil {
			return err
		}

		obj = *clusterNetworkSubnetReservedIP
		return nil
	}
}

func testAccCheckIBMIsClusterNetworkSubnetReservedIPDestroy(s *terraform.State) error {
	vpcClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).VpcV1API()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_is_cluster_network_subnet_reserved_ip" {
			continue
		}

		getClusterNetworkSubnetReservedIPOptions := &vpcv1.GetClusterNetworkSubnetReservedIPOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getClusterNetworkSubnetReservedIPOptions.SetClusterNetworkID(parts[0])
		getClusterNetworkSubnetReservedIPOptions.SetClusterNetworkSubnetID(parts[1])
		getClusterNetworkSubnetReservedIPOptions.SetID(parts[2])

		// Try to find the key
		_, response, err := vpcClient.GetClusterNetworkSubnetReservedIP(getClusterNetworkSubnetReservedIPOptions)

		if err == nil {
			return fmt.Errorf("ClusterNetworkSubnetReservedIP still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for ClusterNetworkSubnetReservedIP (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
