/*
* IBM Confidential
* Object Code Only Source Materials
* 5747-SM3
* (c) Copyright IBM Corp. 2017,2021
*
* The source code for this program is not published or otherwise divested
* of its trade secrets, irrespective of what has been deposited with the
* U.S. Copyright Office.
 */

package ibm

import (
	"errors"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"testing"
)

func TestAccIBMISVPCRoutingTableRoute_basic(t *testing.T) {
	var vpcRouteTables string
	name1 := fmt.Sprintf("tfvpcuat-create-%d", acctest.RandIntRange(10, 100))
	subnetName := fmt.Sprintf("tfsubnet-%d", acctest.RandIntRange(10, 100))
	routeName := fmt.Sprintf("tfvpcuat-create-%d", acctest.RandIntRange(10, 100))
	routeName1 := fmt.Sprintf("tfvpcuat-create-%d", acctest.RandIntRange(10, 100))
	routeTableName := fmt.Sprintf("tfvpcrt-create-%d", acctest.RandIntRange(10, 100))
	routeTableName1 := fmt.Sprintf("tfvpcrt-up-create-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMISVPCRouteTableRouteDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISVPCRouteTableRouteConfig(routeTableName, name1, subnetName, routeName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISVPCRouteTableRouteExists("ibm_is_vpc_routing_table_route.test_custom_route1", vpcRouteTables),
					resource.TestCheckResourceAttr(
						"ibm_is_vpc_routing_table_route.test_custom_route1", "name", routeName),
				),
			},
			{
				Config: testAccCheckIBMISVPCRouteTableRouteConfig(routeTableName1, name1, subnetName, routeName1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISVPCRouteTableRouteExists("ibm_is_vpc_routing_table_route.test_custom_route1", vpcRouteTables),
					resource.TestCheckResourceAttr(
						"ibm_is_vpc_routing_table_route.test_custom_route1", "name", routeName1),
				),
			},
		},
	})
}

func testAccCheckIBMISVPCRouteTableRouteDestroy(s *terraform.State) error {
	//userDetails, _ := testAccProvider.Meta().(ClientSession).BluemixUserDetails()

	//sess, _ := testAccProvider.Meta().(ClientSession).VpcV1API()
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_is_vpc_routing_table_route" {
			continue
		}
		parts, err := idParts(rs.Primary.ID)
		if err != nil {
			return err
		}
		sess, err := vpcClient(testAccProvider.Meta())
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

		//sess, _ := testAccProvider.Meta().(ClientSession).VpcV1API()

		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return errors.New("No Record ID is set")
		}

		parts, err := idParts(rs.Primary.ID)
		if err != nil {
			return err
		}

		vpcID := parts[0]
		routeTableID := parts[1]
		routeID := parts[2]
		sess, err := vpcClient(testAccProvider.Meta())
		getVpcRoutingTableRouteOptions := sess.NewGetVPCRoutingTableRouteOptions(vpcID, routeTableID, routeID)
		rtResponse, detail, err := sess.GetVPCRoutingTableRoute(getVpcRoutingTableRouteOptions)

		if err != nil {
			return fmt.Errorf("Error Getting Routing table route: %s\n%s", err, detail)
		}
		vpcrouteTableID = *rtResponse.ID
		return nil
	}
}

func testAccCheckIBMISVPCRouteTableRouteConfig(rtName, name, subnetName, routeName string) string {
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
`, name, rtName, subnetName, ISZoneName, ISCIDR, routeName, ISZoneName, ISRouteNextHop)
}
