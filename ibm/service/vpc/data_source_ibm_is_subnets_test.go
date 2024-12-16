// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccIBMISSubnetsDataSource_basic(t *testing.T) {
	var subnet string
	resName := "data.ibm_is_subnets.test1"
	vpcname := fmt.Sprintf("tfsubnet-vpc-%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tfsubnet-name-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISSubnetConfig(vpcname, name, acc.ISZoneName, acc.ISCIDR),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISSubnetExists("ibm_is_subnet.testacc_subnet", subnet),
					resource.TestCheckResourceAttr(
						"ibm_is_subnet.testacc_subnet", "name", name),
					resource.TestCheckResourceAttr(
						"ibm_is_subnet.testacc_subnet", "zone", acc.ISZoneName),
					resource.TestCheckResourceAttr(
						"ibm_is_subnet.testacc_subnet", "ipv4_cidr_block", acc.ISCIDR),
				),
			},
			{
				Config: testAccCheckIBMISSubnetsDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resName, "subnets.0.name"),
					resource.TestCheckResourceAttrSet(resName, "subnets.0.status"),
					resource.TestCheckResourceAttrSet(resName, "subnets.0.status"),
					resource.TestCheckResourceAttrSet(resName, "subnets.0.zone"),
					resource.TestCheckResourceAttrSet(resName, "subnets.0.crn"),
					resource.TestCheckResourceAttrSet(resName, "subnets.0.ipv4_cidr_block"),
					resource.TestCheckResourceAttrSet(resName, "subnets.0.available_ipv4_address_count"),
					resource.TestCheckResourceAttrSet(resName, "subnets.0.network_acl"),
					resource.TestCheckResourceAttrSet(resName, "subnets.0.total_ipv4_address_count"),
					resource.TestCheckResourceAttrSet(resName, "subnets.0.vpc"),
				),
			},
		},
	})
}

func TestAccIBMISSubnetsDataSource_basic_filterResourceGroup(t *testing.T) {
	var subnet string
	vpcname := fmt.Sprintf("tfsubnet-vpc-%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tfsubnet-name-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISSubnetsDataSourceFilterResourceGroupConfig(vpcname, name, acc.ISZoneName, acc.ISCIDR),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISSubnetExists("ibm_is_subnet.testacc_subnet", subnet),
					resource.TestCheckResourceAttr(
						"ibm_is_subnet.testacc_subnet", "name", name),
					resource.TestCheckResourceAttr(
						"ibm_is_subnet.testacc_subnet", "zone", acc.ISZoneName),
					resource.TestCheckResourceAttr(
						"ibm_is_subnet.testacc_subnet", "ipv4_cidr_block", acc.ISCIDR),
					resource.TestCheckResourceAttrSet("data.ibm_is_subnets.ds_subnets_resource_group", "subnets.0.name"),
					resource.TestCheckResourceAttrSet("data.ibm_is_subnets.ds_subnets_resource_group", "subnets.0.status"),
					resource.TestCheckResourceAttrSet("data.ibm_is_subnets.ds_subnets_resource_group", "subnets.0.status"),
					resource.TestCheckResourceAttrSet("data.ibm_is_subnets.ds_subnets_resource_group", "subnets.0.zone"),
					resource.TestCheckResourceAttrSet("data.ibm_is_subnets.ds_subnets_resource_group", "subnets.0.crn"),
					resource.TestCheckResourceAttrSet("data.ibm_is_subnets.ds_subnets_resource_group", "subnets.0.ipv4_cidr_block"),
					resource.TestCheckResourceAttrSet("data.ibm_is_subnets.ds_subnets_resource_group", "subnets.0.available_ipv4_address_count"),
					resource.TestCheckResourceAttrSet("data.ibm_is_subnets.ds_subnets_resource_group", "subnets.0.network_acl"),
					resource.TestCheckResourceAttrSet("data.ibm_is_subnets.ds_subnets_resource_group", "subnets.0.total_ipv4_address_count"),
					resource.TestCheckResourceAttrSet("data.ibm_is_subnets.ds_subnets_resource_group", "subnets.0.vpc"),
				),
			},
		},
	})
}

func TestAccIBMISSubnetsDataSource_basic_filterRoutingTable(t *testing.T) {
	var subnet string
	vpcname := fmt.Sprintf("tfsubnet-vpc-%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tfsubnet-name-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISSubnetsDataSourceFilterRoutingTableConfig(vpcname, name, acc.ISZoneName, acc.ISCIDR),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISSubnetExists("ibm_is_subnet.testacc_subnet", subnet),
					resource.TestCheckResourceAttr(
						"ibm_is_subnet.testacc_subnet", "name", name),
					resource.TestCheckResourceAttr(
						"ibm_is_subnet.testacc_subnet", "zone", acc.ISZoneName),
					resource.TestCheckResourceAttr(
						"ibm_is_subnet.testacc_subnet", "ipv4_cidr_block", acc.ISCIDR),
					resource.TestCheckResourceAttrSet("data.ibm_is_subnets.ds_subnets_routing_table", "subnets.0.name"),
					resource.TestCheckResourceAttrSet("data.ibm_is_subnets.ds_subnets_routing_table", "subnets.0.status"),
					resource.TestCheckResourceAttrSet("data.ibm_is_subnets.ds_subnets_routing_table", "subnets.0.status"),
					resource.TestCheckResourceAttrSet("data.ibm_is_subnets.ds_subnets_routing_table", "subnets.0.zone"),
					resource.TestCheckResourceAttrSet("data.ibm_is_subnets.ds_subnets_routing_table", "subnets.0.crn"),
					resource.TestCheckResourceAttrSet("data.ibm_is_subnets.ds_subnets_routing_table", "subnets.0.ipv4_cidr_block"),
					resource.TestCheckResourceAttrSet("data.ibm_is_subnets.ds_subnets_routing_table", "subnets.0.available_ipv4_address_count"),
					resource.TestCheckResourceAttrSet("data.ibm_is_subnets.ds_subnets_routing_table", "subnets.0.network_acl"),
					resource.TestCheckResourceAttrSet("data.ibm_is_subnets.ds_subnets_routing_table", "subnets.0.total_ipv4_address_count"),
					resource.TestCheckResourceAttrSet("data.ibm_is_subnets.ds_subnets_routing_table", "subnets.0.vpc"),
				),
			},
		},
	})
}

func testAccCheckIBMISSubnetsDataSourceConfig() string {
	// status filter defaults to empty
	return fmt.Sprintf(`
	data "ibm_is_subnets" "test1" {
	}`)
}

func testAccCheckIBMISSubnetsDataSourceFilterResourceGroupConfig(vpcname, name, zone, cidr string) string {
	// status filter defaults to empty
	return fmt.Sprintf(`
	data "ibm_resource_group" "resourceGroup" {
		name = "Default"
	}
	  
	resource "ibm_is_vpc" "testacc_vpc" {
		name = "%s"
	}
	  
	resource "ibm_is_vpc_routing_table" "test_cr_route_table1" {
		name = "test-cr-route-table1"
		vpc  = ibm_is_vpc.testacc_vpc.id
	}
	  
	resource "ibm_is_subnet" "testacc_subnet" {
		name            = "%s"
		vpc             = ibm_is_vpc.testacc_vpc.id
		zone            = "%s"
		ipv4_cidr_block = "%s"
		routing_table   = ibm_is_vpc_routing_table.test_cr_route_table1.routing_table
		resource_group  = data.ibm_resource_group.resourceGroup.id
		tags            = ["Tag1", "tag2"]
	}
	  
	data "ibm_is_subnets" "ds_subnets_resource_group" {
		depends_on = [ibm_is_subnet.testacc_subnet]
		resource_group = data.ibm_resource_group.resourceGroup.id
	}
	`, vpcname, name, zone, cidr)
}

func testAccCheckIBMISSubnetsDataSourceFilterRoutingTableConfig(vpcname, name, zone, cidr string) string {
	// status filter defaults to empty
	return fmt.Sprintf(`
	data "ibm_resource_group" "resourceGroup" {
		name = "Default"
	}
	  
	resource "ibm_is_vpc" "testacc_vpc" {
		name = "%s"
	}
	  
	resource "ibm_is_vpc_routing_table" "test_cr_route_table1" {
		name = "test-cr-route-table1"
		vpc  = ibm_is_vpc.testacc_vpc.id
	}
	  
	resource "ibm_is_subnet" "testacc_subnet" {
		name            = "%s"
		vpc             = ibm_is_vpc.testacc_vpc.id
		zone            = "%s"
		ipv4_cidr_block = "%s"
		routing_table   = ibm_is_vpc_routing_table.test_cr_route_table1.routing_table
		resource_group  = data.ibm_resource_group.resourceGroup.id
		tags            = ["Tag1", "tag2"]
	}
	  
	data "ibm_is_subnets" "ds_subnets_routing_table" {
		depends_on = [ibm_is_subnet.testacc_subnet ]
		routing_table = ibm_is_vpc_routing_table.test_cr_route_table1.routing_table
	}
	`, vpcname, name, zone, cidr)
}
