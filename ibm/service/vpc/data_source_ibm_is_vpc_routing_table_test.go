// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMIBMIsVPCRoutingTableDataSourceBasic(t *testing.T) {
	vpcname := fmt.Sprintf("tf-vpcname-%d", acctest.RandIntRange(100, 200))
	rtname := fmt.Sprintf("tf-rtname-%d", acctest.RandIntRange(100, 200))
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIBMIsVPCRoutingTableDataSourceConfigBasic(vpcname, rtname),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_vpc_routing_table.ibm_is_vpc_routing_table", "routing_table"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpc_routing_table.ibm_is_vpc_routing_table", "vpc"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpc_routing_table.ibm_is_vpc_routing_table", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpc_routing_table.ibm_is_vpc_routing_table", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpc_routing_table.ibm_is_vpc_routing_table", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpc_routing_table.ibm_is_vpc_routing_table", "is_default"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpc_routing_table.ibm_is_vpc_routing_table", "lifecycle_state"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpc_routing_table.ibm_is_vpc_routing_table", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpc_routing_table.ibm_is_vpc_routing_table", "resource_type"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpc_routing_table.ibm_is_vpc_routing_table", "route_direct_link_ingress"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpc_routing_table.ibm_is_vpc_routing_table", "route_transit_gateway_ingress"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpc_routing_table.ibm_is_vpc_routing_table", "route_vpc_zone_ingress"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpc_routing_table.ibm_is_vpc_routing_table_name", "routing_table"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpc_routing_table.ibm_is_vpc_routing_table_name", "vpc"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpc_routing_table.ibm_is_vpc_routing_table_name", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpc_routing_table.ibm_is_vpc_routing_table_name", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpc_routing_table.ibm_is_vpc_routing_table_name", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpc_routing_table.ibm_is_vpc_routing_table_name", "is_default"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpc_routing_table.ibm_is_vpc_routing_table_name", "lifecycle_state"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpc_routing_table.ibm_is_vpc_routing_table_name", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpc_routing_table.ibm_is_vpc_routing_table_name", "resource_type"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpc_routing_table.ibm_is_vpc_routing_table_name", "route_direct_link_ingress"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpc_routing_table.ibm_is_vpc_routing_table_name", "route_transit_gateway_ingress"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpc_routing_table.ibm_is_vpc_routing_table_name", "route_vpc_zone_ingress"),
				),
			},
		},
	})
}

func testAccCheckIBMIBMIsVPCRoutingTableDataSourceConfigBasic(vpcname, rtname string) string {
	return fmt.Sprintf(`
	resource "ibm_is_vpc" "test_route_vpc" {
		name 	= "%s"
	}
	resource "ibm_is_vpc_routing_table" "test_route_rt" {
		vpc 							= ibm_is_vpc.test_route_vpc.id
		name 							= "%s"
		route_direct_link_ingress 		= true
		route_transit_gateway_ingress 	= false
		route_vpc_zone_ingress 			= false
	}
	data "ibm_is_vpc_routing_table" "ibm_is_vpc_routing_table" {
		vpc 			= ibm_is_vpc.test_route_vpc.id
		routing_table 	= ibm_is_vpc_routing_table.test_route_rt.routing_table
	}
	data "ibm_is_vpc_routing_table" "ibm_is_vpc_routing_table_name" {
		vpc 			= ibm_is_vpc.test_route_vpc.id
		name 			= ibm_is_vpc_routing_table.test_route_rt.name
	}
	`, vpcname, rtname)
}
