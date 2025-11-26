// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM/vpc-go-sdk/vpcv1"
)

func TestAccIBMIsClusterNetworkBasic(t *testing.T) {
	var conf vpcv1.ClusterNetwork
	vpcname := fmt.Sprintf("tf-vpc-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIsClusterNetworkDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsClusterNetworkConfigBasic(vpcname),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIsClusterNetworkExists("ibm_is_cluster_network.is_cluster_network_instance", conf),
					resource.TestCheckResourceAttr("ibm_is_vpc.is_vpc", "name", vpcname),
					resource.TestCheckResourceAttrSet("ibm_is_vpc.is_vpc", "id"),
					resource.TestCheckResourceAttrSet("ibm_is_cluster_network.is_cluster_network_instance", "id"),
					resource.TestCheckResourceAttrSet("ibm_is_cluster_network.is_cluster_network_instance", "created_at"),
					resource.TestCheckResourceAttrSet("ibm_is_cluster_network.is_cluster_network_instance", "crn"),
					resource.TestCheckResourceAttrSet("ibm_is_cluster_network.is_cluster_network_instance", "href"),
					resource.TestCheckResourceAttrSet("ibm_is_cluster_network.is_cluster_network_instance", "lifecycle_state"),
					resource.TestCheckResourceAttrSet("ibm_is_cluster_network.is_cluster_network_instance", "subnet_prefixes.#"),
					resource.TestCheckResourceAttrSet("ibm_is_cluster_network.is_cluster_network_instance", "vpc.#"),
					resource.TestCheckResourceAttrSet("ibm_is_cluster_network.is_cluster_network_instance", "zone"),
					resource.TestCheckResourceAttr("ibm_is_cluster_network.is_cluster_network_instance", "vpc.0.name", vpcname),
					resource.TestCheckResourceAttr("ibm_is_cluster_network.is_cluster_network_instance", "profile", acc.ISClusterNetworkProfileName),
					resource.TestCheckResourceAttr("ibm_is_cluster_network.is_cluster_network_instance", "lifecycle_state", "stable"),
					resource.TestCheckResourceAttr("ibm_is_cluster_network.is_cluster_network_instance", "resource_type", "cluster_network"),
					resource.TestCheckResourceAttr("ibm_is_cluster_network.is_cluster_network_instance", "zone", acc.ISZoneName),
				),
			},
		},
	})
}

func TestAccIBMIsClusterNetworkBasicAllArgs(t *testing.T) {
	var conf vpcv1.ClusterNetwork
	vpcname := fmt.Sprintf("tf-vpc-%d", acctest.RandIntRange(10, 100))
	subnetPrefixesCidr := acc.ISClusterNetworkSubnetPrefixesCidr
	name := fmt.Sprintf("tf-clusternetwork-%d", acctest.RandIntRange(10, 100))
	nameUpdate := fmt.Sprintf("tf-clusternetwork-updated-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIsClusterNetworkDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsClusterNetworkConfig(vpcname, name, subnetPrefixesCidr),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIsClusterNetworkExists("ibm_is_cluster_network.is_cluster_network_instance", conf),
					resource.TestCheckResourceAttr("ibm_is_cluster_network.is_cluster_network_instance", "name", name),
					resource.TestCheckResourceAttr("ibm_is_vpc.is_vpc", "name", vpcname),
					resource.TestCheckResourceAttrSet("ibm_is_vpc.is_vpc", "id"),
					resource.TestCheckResourceAttrSet("ibm_is_cluster_network.is_cluster_network_instance", "id"),
					resource.TestCheckResourceAttrSet("ibm_is_cluster_network.is_cluster_network_instance", "created_at"),
					resource.TestCheckResourceAttrSet("ibm_is_cluster_network.is_cluster_network_instance", "crn"),
					resource.TestCheckResourceAttrSet("ibm_is_cluster_network.is_cluster_network_instance", "href"),
					resource.TestCheckResourceAttrSet("ibm_is_cluster_network.is_cluster_network_instance", "lifecycle_state"),
					resource.TestCheckResourceAttrSet("ibm_is_cluster_network.is_cluster_network_instance", "subnet_prefixes.#"),
					resource.TestCheckResourceAttrSet("ibm_is_cluster_network.is_cluster_network_instance", "vpc.#"),
					resource.TestCheckResourceAttrSet("ibm_is_cluster_network.is_cluster_network_instance", "zone"),
					resource.TestCheckResourceAttr("ibm_is_cluster_network.is_cluster_network_instance", "vpc.0.name", vpcname),
					resource.TestCheckResourceAttr("ibm_is_cluster_network.is_cluster_network_instance", "profile", acc.ISClusterNetworkProfileName),
					resource.TestCheckResourceAttr("ibm_is_cluster_network.is_cluster_network_instance", "lifecycle_state", "stable"),
					resource.TestCheckResourceAttr("ibm_is_cluster_network.is_cluster_network_instance", "resource_type", "cluster_network"),
					resource.TestCheckResourceAttr("ibm_is_cluster_network.is_cluster_network_instance", "subnet_prefixes.0.cidr", acc.ISClusterNetworkSubnetPrefixesCidr),
					resource.TestCheckResourceAttr("ibm_is_cluster_network.is_cluster_network_instance", "zone", acc.ISZoneName),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMIsClusterNetworkConfig(vpcname, nameUpdate, subnetPrefixesCidr),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_is_cluster_network.is_cluster_network_instance", "name", nameUpdate),
					resource.TestCheckResourceAttr("ibm_is_vpc.is_vpc", "name", vpcname),
					resource.TestCheckResourceAttrSet("ibm_is_vpc.is_vpc", "id"),
					resource.TestCheckResourceAttrSet("ibm_is_cluster_network.is_cluster_network_instance", "id"),
					resource.TestCheckResourceAttrSet("ibm_is_cluster_network.is_cluster_network_instance", "created_at"),
					resource.TestCheckResourceAttrSet("ibm_is_cluster_network.is_cluster_network_instance", "crn"),
					resource.TestCheckResourceAttrSet("ibm_is_cluster_network.is_cluster_network_instance", "href"),
					resource.TestCheckResourceAttrSet("ibm_is_cluster_network.is_cluster_network_instance", "lifecycle_state"),
					resource.TestCheckResourceAttrSet("ibm_is_cluster_network.is_cluster_network_instance", "subnet_prefixes.#"),
					resource.TestCheckResourceAttrSet("ibm_is_cluster_network.is_cluster_network_instance", "vpc.#"),
					resource.TestCheckResourceAttrSet("ibm_is_cluster_network.is_cluster_network_instance", "zone"),
					resource.TestCheckResourceAttr("ibm_is_cluster_network.is_cluster_network_instance", "vpc.0.name", vpcname),
					resource.TestCheckResourceAttr("ibm_is_cluster_network.is_cluster_network_instance", "profile", acc.ISClusterNetworkProfileName),
					resource.TestCheckResourceAttr("ibm_is_cluster_network.is_cluster_network_instance", "lifecycle_state", "stable"),
					resource.TestCheckResourceAttr("ibm_is_cluster_network.is_cluster_network_instance", "subnet_prefixes.0.cidr", acc.ISClusterNetworkSubnetPrefixesCidr),
					resource.TestCheckResourceAttr("ibm_is_cluster_network.is_cluster_network_instance", "resource_type", "cluster_network"),
					resource.TestCheckResourceAttr("ibm_is_cluster_network.is_cluster_network_instance", "zone", acc.ISZoneName),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_is_cluster_network.is_cluster_network_instance",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIBMIsClusterNetworkConfigBasic(vpcname string) string {
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
	`, vpcname, acc.ISClusterNetworkProfileName, acc.ISZoneName)
}

func testAccCheckIBMIsClusterNetworkConfig(vpcname, clusterNetworkName, subnetPrefixesCidr string) string {
	return fmt.Sprintf(`
		resource "ibm_is_vpc" "is_vpc" {
  			name = "%s"
		}
		resource "ibm_is_cluster_network" "is_cluster_network_instance" {
			name = "%s"
			profile = "%s"
			subnet_prefixes {
				cidr = "%s"
			}
			vpc {
				id = ibm_is_vpc.is_vpc.id
			}
			zone  = "%s"
		}
	`, vpcname, clusterNetworkName, acc.ISClusterNetworkProfileName, subnetPrefixesCidr, acc.ISZoneName)
}

func testAccCheckIBMIsClusterNetworkExists(n string, obj vpcv1.ClusterNetwork) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		vpcClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).VpcV1API()
		if err != nil {
			return err
		}

		getClusterNetworkOptions := &vpcv1.GetClusterNetworkOptions{}

		getClusterNetworkOptions.SetID(rs.Primary.ID)

		clusterNetwork, _, err := vpcClient.GetClusterNetwork(getClusterNetworkOptions)
		if err != nil {
			return err
		}

		obj = *clusterNetwork
		return nil
	}
}

func testAccCheckIBMIsClusterNetworkDestroy(s *terraform.State) error {
	vpcClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).VpcV1API()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_is_cluster_network" {
			continue
		}

		getClusterNetworkOptions := &vpcv1.GetClusterNetworkOptions{}

		getClusterNetworkOptions.SetID(rs.Primary.ID)

		// Try to find the key
		_, response, err := vpcClient.GetClusterNetwork(getClusterNetworkOptions)

		if err == nil {
			return fmt.Errorf("ClusterNetwork still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for ClusterNetwork (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
