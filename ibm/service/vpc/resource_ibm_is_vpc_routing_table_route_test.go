// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"errors"
	"fmt"
	"testing"

	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/flex"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccIBMISVPCRoutingTableRoute_basic(t *testing.T) {
	var vpcRouteTables string
	name1 := fmt.Sprintf("tfvpcuat-create-%d", acctest.RandIntRange(10, 100))
	subnetName := fmt.Sprintf("tfsubnet-%d", acctest.RandIntRange(10, 100))
	routeName := fmt.Sprintf("tfvpcuat-create-%d", acctest.RandIntRange(10, 100))
	routeName1 := fmt.Sprintf("tfvpcuat-create-%d", acctest.RandIntRange(10, 100))
	routeTableName := fmt.Sprintf("tfvpcrt-create-%d", acctest.RandIntRange(10, 100))
	routeTableName1 := fmt.Sprintf("tfvpcrt-up-create-%d", acctest.RandIntRange(10, 100))
	advertiseVal := fmt.Sprintf("tfpvpcuat-create-%d", acctest.RandIntRange(10, 50))
	advertiseValUpd := fmt.Sprintf("tfpvpcuat-update-%d", acctest.RandIntRange(60, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISVPCRouteTableRouteDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISVPCRouteTableRouteConfig(routeTableName, name1, subnetName, routeName, advertiseVal),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISVPCRouteTableRouteExists("ibm_is_vpc_routing_table_route.test_custom_route1", vpcRouteTables),
					resource.TestCheckResourceAttr(
						"ibm_is_vpc_routing_table_route.test_custom_route1", "name", routeName),
					resource.TestCheckResourceAttr(
						"ibm_is_vpc_routing_table_route.test_custom_route1", "advertise", advertiseVal),
				),
			},
			{
				Config: testAccCheckIBMISVPCRouteTableRouteConfig(routeTableName1, name1, subnetName, routeName1, advertiseValUpd),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISVPCRouteTableRouteExists("ibm_is_vpc_routing_table_route.test_custom_route1", vpcRouteTables),
					resource.TestCheckResourceAttr(
						"ibm_is_vpc_routing_table_route.test_custom_route1", "name", routeName1),
					resource.TestCheckResourceAttr(
						"ibm_is_vpc_routing_table_route.test_custom_route1", "advertise", advertiseValUpd),
				),
			},
		},
	})
}

func testAccCheckIBMISVPCRouteTableRouteDestroy(s *terraform.State) error {
	//userDetails, _ := acc.TestAccProvider.Meta().(conns.ClientSession).BluemixUserDetails()

	//sess, _ := acc.TestAccProvider.Meta().(conns.ClientSession).VpcV1API()
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_is_vpc_routing_table_route" {
			continue
		}
		parts, err := flex.IdParts(rs.Primary.ID)
		if err != nil {
			return err
		}
		sess, err := vpcClient(acc.TestAccProvider.Meta())
		if err != nil {
			return err
		}
		vpcID := parts[0]
		routeTableID := parts[1]
		routeID := parts[2]
		getVpcRoutingTableRouteOptions := sess.NewGetVPCRoutingTableRouteOptions(vpcID, routeTableID, routeID)
		_, _, err = sess.GetVPCRoutingTableRoute(getVpcRoutingTableRouteOptions)
		if err == nil {
			return fmt.Errorf("Routing table route still exists: %s", rs.Primary.ID)
		}
	}

	return nil
}

func testAccCheckIBMISVPCRouteTableRouteExists(n, vpcrouteTableID string) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		//sess, _ := acc.TestAccProvider.Meta().(conns.ClientSession).VpcV1API()

		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return errors.New("No Record ID is set")
		}

		parts, err := flex.IdParts(rs.Primary.ID)
		if err != nil {
			return err
		}

		vpcID := parts[0]
		routeTableID := parts[1]
		routeID := parts[2]
		sess, err := vpcClient(acc.TestAccProvider.Meta())
		getVpcRoutingTableRouteOptions := sess.NewGetVPCRoutingTableRouteOptions(vpcID, routeTableID, routeID)
		rtResponse, detail, err := sess.GetVPCRoutingTableRoute(getVpcRoutingTableRouteOptions)

		if err != nil {
			return fmt.Errorf("[ERROR] Error Getting Routing table route: %s\n%s", err, detail)
		}
		vpcrouteTableID = *rtResponse.ID
		return nil
	}
}

func testAccCheckIBMISVPCRouteTableRouteConfig(rtName, name, subnetName, routeName, advertise string) string {
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
  advertise = "%s"
  name = "%s"
  zone = "%s"
  next_hop = "%s"
  destination = ibm_is_subnet.test_cr_subnet1.ipv4_cidr_block
}
`, name, rtName, subnetName, acc.ISZoneName, acc.ISCIDR, advertise, routeName, acc.ISZoneName, acc.ISRouteNextHop)
}
