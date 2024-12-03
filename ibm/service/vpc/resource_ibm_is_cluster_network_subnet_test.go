// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/vpc-go-sdk/vpcv1"
)

func TestAccIBMIsClusterNetworkSubnetBasic(t *testing.T) {
	var conf vpcv1.ClusterNetworkSubnet
	vpcname := fmt.Sprintf("tf-vpc-%d", acctest.RandIntRange(10, 100))
	totalipv4addresscount := int64(64)
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIsClusterNetworkSubnetDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsClusterNetworkSubnetConfigBasic(vpcname, totalipv4addresscount),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIsClusterNetworkSubnetExists("ibm_is_cluster_network_subnet.is_cluster_network_subnet_instance", conf),
					resource.TestCheckResourceAttr("ibm_is_vpc.is_vpc", "name", vpcname),
					resource.TestCheckResourceAttrSet("ibm_is_vpc.is_vpc", "id"),
					resource.TestCheckResourceAttrSet("ibm_is_cluster_network.is_cluster_network_instance", "id"),
					resource.TestCheckResourceAttrSet("ibm_is_cluster_network_subnet.is_cluster_network_subnet_instance", "available_ipv4_address_count"),
					resource.TestCheckResourceAttrSet("ibm_is_cluster_network_subnet.is_cluster_network_subnet_instance", "cluster_network_subnet_id"),
					resource.TestCheckResourceAttrSet("ibm_is_cluster_network_subnet.is_cluster_network_subnet_instance", "created_at"),
					resource.TestCheckResourceAttrSet("ibm_is_cluster_network_subnet.is_cluster_network_subnet_instance", "href"),
					resource.TestCheckResourceAttrSet("ibm_is_cluster_network_subnet.is_cluster_network_subnet_instance", "id"),
					resource.TestCheckResourceAttrSet("ibm_is_cluster_network_subnet.is_cluster_network_subnet_instance", "ip_version"),
					resource.TestCheckResourceAttrSet("ibm_is_cluster_network_subnet.is_cluster_network_subnet_instance", "ipv4_cidr_block"),
					resource.TestCheckResourceAttrSet("ibm_is_cluster_network_subnet.is_cluster_network_subnet_instance", "lifecycle_state"),
					resource.TestCheckResourceAttrSet("ibm_is_cluster_network_subnet.is_cluster_network_subnet_instance", "name"),
					resource.TestCheckResourceAttrSet("ibm_is_cluster_network_subnet.is_cluster_network_subnet_instance", "resource_type"),
					resource.TestCheckResourceAttrSet("ibm_is_cluster_network_subnet.is_cluster_network_subnet_instance", "total_ipv4_address_count"),
					resource.TestCheckResourceAttr("ibm_is_cluster_network_subnet.is_cluster_network_subnet_instance", "total_ipv4_address_count", "64"),
				),
			},
		},
	})
}

func TestAccIBMIsClusterNetworkSubnetBasicAllArgs(t *testing.T) {
	var conf vpcv1.ClusterNetworkSubnet
	vpcname := fmt.Sprintf("tf-vpcname-%d", acctest.RandIntRange(10, 100))
	name1 := fmt.Sprintf("tf-cluster-subnet1-%d", acctest.RandIntRange(10, 100))
	name1Updated := fmt.Sprintf("tf-cluster-subnet1-updated%d", acctest.RandIntRange(10, 100))
	name2 := fmt.Sprintf("tf-cluster-subnet2-%d", acctest.RandIntRange(10, 100))
	name2Updated := fmt.Sprintf("tf-cluster-subnet2-%d", acctest.RandIntRange(10, 100))
	totalIpv4AddressCount := int64(64)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIsClusterNetworkSubnetDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsClusterNetworkSubnetConfig(vpcname, name1, name2, totalIpv4AddressCount),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIsClusterNetworkSubnetExists("ibm_is_cluster_network_subnet.is_cluster_network_subnet_instance", conf),
					resource.TestCheckResourceAttr("ibm_is_vpc.is_vpc", "name", vpcname),
					resource.TestCheckResourceAttrSet("ibm_is_vpc.is_vpc", "id"),
					resource.TestCheckResourceAttrSet("ibm_is_cluster_network.is_cluster_network_instance", "id"),
					resource.TestCheckResourceAttr("ibm_is_cluster_network_subnet.is_cluster_network_subnet_instance", "name", name1),
					resource.TestCheckResourceAttr("ibm_is_cluster_network_subnet.is_cluster_network_subnet_instance", "total_ipv4_address_count", "64"),
					resource.TestCheckResourceAttr("ibm_is_cluster_network_subnet.is_cluster_network_subnet_instance_cidr", "name", name2),
					resource.TestCheckResourceAttr("ibm_is_cluster_network_subnet.is_cluster_network_subnet_instance_cidr", "ipv4_cidr_block", acc.ISCIDR),

					resource.TestCheckResourceAttrSet("ibm_is_cluster_network_subnet.is_cluster_network_subnet_instance", "available_ipv4_address_count"),
					resource.TestCheckResourceAttrSet("ibm_is_cluster_network_subnet.is_cluster_network_subnet_instance", "cluster_network_subnet_id"),
					resource.TestCheckResourceAttrSet("ibm_is_cluster_network_subnet.is_cluster_network_subnet_instance", "created_at"),
					resource.TestCheckResourceAttrSet("ibm_is_cluster_network_subnet.is_cluster_network_subnet_instance", "href"),
					resource.TestCheckResourceAttrSet("ibm_is_cluster_network_subnet.is_cluster_network_subnet_instance", "id"),
					resource.TestCheckResourceAttrSet("ibm_is_cluster_network_subnet.is_cluster_network_subnet_instance", "ip_version"),
					resource.TestCheckResourceAttrSet("ibm_is_cluster_network_subnet.is_cluster_network_subnet_instance", "ipv4_cidr_block"),
					resource.TestCheckResourceAttrSet("ibm_is_cluster_network_subnet.is_cluster_network_subnet_instance", "lifecycle_state"),
					resource.TestCheckResourceAttrSet("ibm_is_cluster_network_subnet.is_cluster_network_subnet_instance", "name"),
					resource.TestCheckResourceAttrSet("ibm_is_cluster_network_subnet.is_cluster_network_subnet_instance", "resource_type"),
					resource.TestCheckResourceAttrSet("ibm_is_cluster_network_subnet.is_cluster_network_subnet_instance", "total_ipv4_address_count"),
					resource.TestCheckResourceAttrSet("ibm_is_cluster_network_subnet.is_cluster_network_subnet_instance_cidr", "available_ipv4_address_count"),
					resource.TestCheckResourceAttrSet("ibm_is_cluster_network_subnet.is_cluster_network_subnet_instance_cidr", "cluster_network_subnet_id"),
					resource.TestCheckResourceAttrSet("ibm_is_cluster_network_subnet.is_cluster_network_subnet_instance_cidr", "created_at"),
					resource.TestCheckResourceAttrSet("ibm_is_cluster_network_subnet.is_cluster_network_subnet_instance_cidr", "href"),
					resource.TestCheckResourceAttrSet("ibm_is_cluster_network_subnet.is_cluster_network_subnet_instance_cidr", "id"),
					resource.TestCheckResourceAttrSet("ibm_is_cluster_network_subnet.is_cluster_network_subnet_instance_cidr", "ip_version"),
					resource.TestCheckResourceAttrSet("ibm_is_cluster_network_subnet.is_cluster_network_subnet_instance_cidr", "ipv4_cidr_block"),
					resource.TestCheckResourceAttrSet("ibm_is_cluster_network_subnet.is_cluster_network_subnet_instance_cidr", "lifecycle_state"),
					resource.TestCheckResourceAttrSet("ibm_is_cluster_network_subnet.is_cluster_network_subnet_instance_cidr", "name"),
					resource.TestCheckResourceAttrSet("ibm_is_cluster_network_subnet.is_cluster_network_subnet_instance_cidr", "resource_type"),
					resource.TestCheckResourceAttrSet("ibm_is_cluster_network_subnet.is_cluster_network_subnet_instance_cidr", "total_ipv4_address_count"),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMIsClusterNetworkSubnetConfig(vpcname, name1Updated, name2Updated, totalIpv4AddressCount),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_is_vpc.is_vpc", "name", vpcname),
					resource.TestCheckResourceAttrSet("ibm_is_vpc.is_vpc", "id"),
					resource.TestCheckResourceAttrSet("ibm_is_cluster_network.is_cluster_network_instance", "id"),
					resource.TestCheckResourceAttr("ibm_is_cluster_network_subnet.is_cluster_network_subnet_instance", "name", name1Updated),
					resource.TestCheckResourceAttr("ibm_is_cluster_network_subnet.is_cluster_network_subnet_instance", "total_ipv4_address_count", "64"),
					resource.TestCheckResourceAttr("ibm_is_cluster_network_subnet.is_cluster_network_subnet_instance_cidr", "name", name2Updated),
					resource.TestCheckResourceAttr("ibm_is_cluster_network_subnet.is_cluster_network_subnet_instance_cidr", "ipv4_cidr_block", acc.ISCIDR),

					resource.TestCheckResourceAttrSet("ibm_is_cluster_network_subnet.is_cluster_network_subnet_instance", "available_ipv4_address_count"),
					resource.TestCheckResourceAttrSet("ibm_is_cluster_network_subnet.is_cluster_network_subnet_instance", "cluster_network_subnet_id"),
					resource.TestCheckResourceAttrSet("ibm_is_cluster_network_subnet.is_cluster_network_subnet_instance", "created_at"),
					resource.TestCheckResourceAttrSet("ibm_is_cluster_network_subnet.is_cluster_network_subnet_instance", "href"),
					resource.TestCheckResourceAttrSet("ibm_is_cluster_network_subnet.is_cluster_network_subnet_instance", "id"),
					resource.TestCheckResourceAttrSet("ibm_is_cluster_network_subnet.is_cluster_network_subnet_instance", "ip_version"),
					resource.TestCheckResourceAttrSet("ibm_is_cluster_network_subnet.is_cluster_network_subnet_instance", "ipv4_cidr_block"),
					resource.TestCheckResourceAttrSet("ibm_is_cluster_network_subnet.is_cluster_network_subnet_instance", "lifecycle_state"),
					resource.TestCheckResourceAttrSet("ibm_is_cluster_network_subnet.is_cluster_network_subnet_instance", "name"),
					resource.TestCheckResourceAttrSet("ibm_is_cluster_network_subnet.is_cluster_network_subnet_instance", "resource_type"),
					resource.TestCheckResourceAttrSet("ibm_is_cluster_network_subnet.is_cluster_network_subnet_instance", "total_ipv4_address_count"),
					resource.TestCheckResourceAttrSet("ibm_is_cluster_network_subnet.is_cluster_network_subnet_instance_cidr", "available_ipv4_address_count"),
					resource.TestCheckResourceAttrSet("ibm_is_cluster_network_subnet.is_cluster_network_subnet_instance_cidr", "cluster_network_subnet_id"),
					resource.TestCheckResourceAttrSet("ibm_is_cluster_network_subnet.is_cluster_network_subnet_instance_cidr", "created_at"),
					resource.TestCheckResourceAttrSet("ibm_is_cluster_network_subnet.is_cluster_network_subnet_instance_cidr", "href"),
					resource.TestCheckResourceAttrSet("ibm_is_cluster_network_subnet.is_cluster_network_subnet_instance_cidr", "id"),
					resource.TestCheckResourceAttrSet("ibm_is_cluster_network_subnet.is_cluster_network_subnet_instance_cidr", "ip_version"),
					resource.TestCheckResourceAttrSet("ibm_is_cluster_network_subnet.is_cluster_network_subnet_instance_cidr", "ipv4_cidr_block"),
					resource.TestCheckResourceAttrSet("ibm_is_cluster_network_subnet.is_cluster_network_subnet_instance_cidr", "lifecycle_state"),
					resource.TestCheckResourceAttrSet("ibm_is_cluster_network_subnet.is_cluster_network_subnet_instance_cidr", "name"),
					resource.TestCheckResourceAttrSet("ibm_is_cluster_network_subnet.is_cluster_network_subnet_instance_cidr", "resource_type"),
					resource.TestCheckResourceAttrSet("ibm_is_cluster_network_subnet.is_cluster_network_subnet_instance_cidr", "total_ipv4_address_count"),
				),
			},
		},
	})
}

func testAccCheckIBMIsClusterNetworkSubnetConfigBasic(vpcname string, totalipv4addresscount int64) string {
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
			total_ipv4_address_count = %d
		}
	`, vpcname, acc.ISClusterNetworkProfileName, acc.ISZoneName, totalipv4addresscount)
}

func testAccCheckIBMIsClusterNetworkSubnetConfig(vpcname, subnetname1, subnetname2 string, totalIpv4AddressCount int64) string {
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
			total_ipv4_address_count = %d
		}
		resource "ibm_is_cluster_network_subnet" "is_cluster_network_subnet_instance_cidr" {
			cluster_network_id = ibm_is_cluster_network.is_cluster_network_instance.id
			name = "%s"
			ipv4_cidr_block = "%s"
		}
	`, vpcname, acc.ISClusterNetworkProfileName, acc.ISZoneName, subnetname1, totalIpv4AddressCount, subnetname2, acc.ISCIDR)
}

func testAccCheckIBMIsClusterNetworkSubnetExists(n string, obj vpcv1.ClusterNetworkSubnet) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		vpcClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).VpcV1API()
		if err != nil {
			return err
		}

		getClusterNetworkSubnetOptions := &vpcv1.GetClusterNetworkSubnetOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getClusterNetworkSubnetOptions.SetClusterNetworkID(parts[0])
		getClusterNetworkSubnetOptions.SetID(parts[1])

		clusterNetworkSubnet, _, err := vpcClient.GetClusterNetworkSubnet(getClusterNetworkSubnetOptions)
		if err != nil {
			return err
		}

		obj = *clusterNetworkSubnet
		return nil
	}
}

func testAccCheckIBMIsClusterNetworkSubnetDestroy(s *terraform.State) error {
	vpcClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).VpcV1API()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_is_cluster_network_subnet" {
			continue
		}

		getClusterNetworkSubnetOptions := &vpcv1.GetClusterNetworkSubnetOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getClusterNetworkSubnetOptions.SetClusterNetworkID(parts[0])
		getClusterNetworkSubnetOptions.SetID(parts[1])

		// Try to find the key
		_, response, err := vpcClient.GetClusterNetworkSubnet(getClusterNetworkSubnetOptions)

		if err == nil {
			return fmt.Errorf("ClusterNetworkSubnet still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for ClusterNetworkSubnet (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
