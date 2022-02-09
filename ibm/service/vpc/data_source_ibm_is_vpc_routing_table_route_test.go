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

func TestAccIBMIBMIsVPCRoutingTableRouteDataSourceBasic(t *testing.T) {
	vpcname := fmt.Sprintf("tf-vpcname-%d", acctest.RandIntRange(100, 200))
	rtname := fmt.Sprintf("tf-rtname-%d", acctest.RandIntRange(100, 200))
	rname := fmt.Sprintf("tf-routename-%d", acctest.RandIntRange(100, 200))
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIBMIsVPCRoutingTableRouteDataSourceConfigBasic(vpcname, rtname, rname),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_vpc_routing_table_route.ibm_is_vpc_routing_table_route", "route_id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpc_routing_table_route.ibm_is_vpc_routing_table_route_name", "route_id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpc_routing_table_route.ibm_is_vpc_routing_table_route_name", "vpc"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpc_routing_table_route.ibm_is_vpc_routing_table_route", "vpc"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpc_routing_table_route.ibm_is_vpc_routing_table_route", "routing_table"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpc_routing_table_route.ibm_is_vpc_routing_table_route_name", "routing_table"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpc_routing_table_route.ibm_is_vpc_routing_table_route", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpc_routing_table_route.ibm_is_vpc_routing_table_route_name", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpc_routing_table_route.ibm_is_vpc_routing_table_route", "action"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpc_routing_table_route.ibm_is_vpc_routing_table_route_name", "action"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpc_routing_table_route.ibm_is_vpc_routing_table_route", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpc_routing_table_route.ibm_is_vpc_routing_table_route_name", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpc_routing_table_route.ibm_is_vpc_routing_table_route", "destination"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpc_routing_table_route.ibm_is_vpc_routing_table_route_name", "destination"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpc_routing_table_route.ibm_is_vpc_routing_table_route", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpc_routing_table_route.ibm_is_vpc_routing_table_route_name", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpc_routing_table_route.ibm_is_vpc_routing_table_route", "lifecycle_state"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpc_routing_table_route.ibm_is_vpc_routing_table_route_name", "lifecycle_state"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpc_routing_table_route.ibm_is_vpc_routing_table_route", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpc_routing_table_route.ibm_is_vpc_routing_table_route_name", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpc_routing_table_route.ibm_is_vpc_routing_table_route", "next_hop.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpc_routing_table_route.ibm_is_vpc_routing_table_route_name", "next_hop.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpc_routing_table_route.ibm_is_vpc_routing_table_route", "zone.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpc_routing_table_route.ibm_is_vpc_routing_table_route_name", "zone.#"),
				),
			},
		},
	})
}

func testAccCheckIBMIBMIsVPCRoutingTableRouteDataSourceConfigBasic(vpcname, rtname, rname string) string {
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
		resource "ibm_is_vpc_routing_table_route" "test_route_rt_route" {
			vpc 			= ibm_is_vpc.test_route_vpc.id
			routing_table 	= ibm_is_vpc_routing_table.test_route_rt.routing_table
			zone 			= "us-south-1"
			name 			= "%s"
			destination 	= "192.168.4.0/24"
			action 			= "deliver"
			next_hop 		= "0.0.0.0"
		  }
		data "ibm_is_vpc_routing_table_route" "ibm_is_vpc_routing_table_route" {
			vpc 			= ibm_is_vpc.test_route_vpc.id
			routing_table 	= ibm_is_vpc_routing_table.test_route_rt.routing_table
			route_id 		= ibm_is_vpc_routing_table_route.test_route_rt_route.route_id
		}
		data "ibm_is_vpc_routing_table_route" "ibm_is_vpc_routing_table_route_name" {
			vpc 			= ibm_is_vpc.test_route_vpc.id
			routing_table 	= ibm_is_vpc_routing_table.test_route_rt.routing_table
			name 			= ibm_is_vpc_routing_table_route.test_route_rt_route.name
		}
	`, vpcname, rtname, rname)
}
