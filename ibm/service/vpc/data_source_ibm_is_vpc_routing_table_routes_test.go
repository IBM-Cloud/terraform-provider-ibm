// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"fmt"
	"testing"

	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMISVPCRoutingTableRoutesDataSource_basic(t *testing.T) {
	node := "data.ibm_is_vpc_routing_table_routes.routes_test"
	name1 := fmt.Sprintf("tfvpcuat-create-data-%d", acctest.RandIntRange(10, 100))
	subnetName := fmt.Sprintf("tfsubnet-create-data-%d", acctest.RandIntRange(10, 100))
	routeName := fmt.Sprintf("tfvpcuat-create-data-%d", acctest.RandIntRange(10, 100))
	routeTableName := fmt.Sprintf("tfvpcrt-create-data-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISVPCRoutingTableRoutesDataSourceConfig(routeTableName, name1, subnetName, routeName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(node, "routes.#"),
				),
			},
		},
	})
}

func testAccCheckIBMISVPCRoutingTableRoutesDataSourceConfig(rtName, name, subnetName, routeName string) string {
	return fmt.Sprintf(`
resource "ibm_is_vpc" "testacc_vpc" {
    name = "%s"
}
resource "ibm_is_vpc_routing_table" "test_ibm_is_vpc_routing_table" {
	depends_on = [ibm_is_vpc.testacc_vpc]
	vpc = ibm_is_vpc.testacc_vpc.id
	name = "%s"
}
resource "ibm_is_subnet" "test_cr_subnet1" {
	depends_on = [ibm_is_vpc_routing_table.test_ibm_is_vpc_routing_table]
	name = "%s"
	vpc = ibm_is_vpc.testacc_vpc.id
	zone = "%s"
	ipv4_cidr_block = "%s"
	routing_table = ibm_is_vpc_routing_table.test_ibm_is_vpc_routing_table.routing_table
}
//custom route for source
resource "ibm_is_vpc_routing_table_route" "test_custom_route1" {
  depends_on = [ibm_is_vpc_routing_table.test_ibm_is_vpc_routing_table, ibm_is_subnet.test_cr_subnet1]
  vpc = ibm_is_vpc.testacc_vpc.id
  routing_table = ibm_is_vpc_routing_table.test_ibm_is_vpc_routing_table.routing_table
  name = "%s"
  zone = "%s"
  next_hop = "%s"
  destination = ibm_is_subnet.test_cr_subnet1.ipv4_cidr_block
}

data "ibm_is_vpc_routing_table_routes" "routes_test" {
	vpc = ibm_is_vpc.testacc_vpc.id
	routing_table = ibm_is_vpc_routing_table.test_ibm_is_vpc_routing_table.routing_table
  }
  
`, name, rtName, subnetName, acc.ISZoneName, acc.ISCIDR, routeName, acc.ISZoneName, acc.ISRouteNextHop)
}
