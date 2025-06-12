// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"errors"
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccIBMISVPCRoutingTable_basic(t *testing.T) {
	var vpcRouteTables string
	name1 := fmt.Sprintf("tfvpc-create-%d", acctest.RandIntRange(10, 100))
	routeTableName := fmt.Sprintf("tfvpcrt-create-%d", acctest.RandIntRange(10, 100))
	routeTableName1 := fmt.Sprintf("tfvpcrt-up-create-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISVPCRouteTableDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISVPCRouteTableConfig(routeTableName, name1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISVPCRouteTableExists("ibm_is_vpc_routing_table.test_ibm_is_vpc_routing_table", vpcRouteTables),
					resource.TestCheckResourceAttr(
						"ibm_is_vpc_routing_table.test_ibm_is_vpc_routing_table", "name", routeTableName),
				),
			},
			{
				Config: testAccCheckIBMISVPCRouteTableConfig(routeTableName1, name1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISVPCRouteTableExists("ibm_is_vpc_routing_table.test_ibm_is_vpc_routing_table", vpcRouteTables),
					resource.TestCheckResourceAttr(
						"ibm_is_vpc_routing_table.test_ibm_is_vpc_routing_table", "name", routeTableName1),
				),
			},
		},
	})
}

func TestAccIBMISVPCRoutingTable_acceptRoutesFrom(t *testing.T) {
	var vpcRouteTables string
	name1 := fmt.Sprintf("tfvpc-create-%d", acctest.RandIntRange(10, 100))
	routeTableName := fmt.Sprintf("tfvpcrt-create-%d", acctest.RandIntRange(10, 100))
	routeTableName1 := fmt.Sprintf("tfvpcrt-up-create-%d", acctest.RandIntRange(10, 100))
	acceptRoutesFromVPNServer := "vpn_server"
	acceptRoutesFromVPNGateway := "vpn_gateway"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISVPCRouteTableDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISVPCRouteTableAcceptRoutesFromConfig(routeTableName, name1, acceptRoutesFromVPNServer),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISVPCRouteTableExists("ibm_is_vpc_routing_table.test_ibm_is_vpc_routing_table", vpcRouteTables),
					resource.TestCheckResourceAttr(
						"ibm_is_vpc_routing_table.test_ibm_is_vpc_routing_table", "name", routeTableName),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpc_routing_table.test_ibm_is_vpc_routing_table", "accept_routes_from_resource_type.#"),
				),
			},
			{
				Config: testAccCheckIBMISVPCRouteTableAcceptRoutesFromConfig(routeTableName1, name1, acceptRoutesFromVPNGateway),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISVPCRouteTableExists("ibm_is_vpc_routing_table.test_ibm_is_vpc_routing_table", vpcRouteTables),
					resource.TestCheckResourceAttr(
						"ibm_is_vpc_routing_table.test_ibm_is_vpc_routing_table", "name", routeTableName1),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpc_routing_table.test_ibm_is_vpc_routing_table", "accept_routes_from_resource_type.#"),
				),
			},
			{
				ResourceName:      "ibm_is_vpc_routing_table.test_ibm_is_vpc_routing_table",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

// advertise_routes_to
func TestAccIBMISVPCRoutingTable_advertiseRoutesTO(t *testing.T) {
	var vpcRouteTables string
	name1 := fmt.Sprintf("tfvpc-create-%d", acctest.RandIntRange(10, 100))
	routeTableName := fmt.Sprintf("tfvpcrt-create-%d", acctest.RandIntRange(10, 100))

	advertiseRoutesToDirectLink := "direct_link"
	advertiseRoutesToTransit_gateway := "transit_gateway"
	acceptRoutesFromVPNServer := "vpn_server"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISVPCRouteTableDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISVPCRouteTableAdvertiseRoutesToConfig(routeTableName, name1, acceptRoutesFromVPNServer, advertiseRoutesToDirectLink, advertiseRoutesToTransit_gateway),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISVPCRouteTableExists("ibm_is_vpc_routing_table.test_ibm_is_vpc_routing_table", vpcRouteTables),
					resource.TestCheckResourceAttr(
						"ibm_is_vpc_routing_table.test_ibm_is_vpc_routing_table", "advertise_routes_to.0", advertiseRoutesToDirectLink),
					resource.TestCheckResourceAttr(
						"ibm_is_vpc_routing_table.test_ibm_is_vpc_routing_table", "advertise_routes_to.1", advertiseRoutesToTransit_gateway),
					resource.TestCheckResourceAttr(
						"ibm_is_vpc_routing_table.test_ibm_is_vpc_routing_table", "accept_routes_from_resource_type.0", acceptRoutesFromVPNServer),
				),
			},
			{
				Config: testAccCheckIBMISVPCRouteTableAdvertiseRoutesToDLConfig(routeTableName, name1, acceptRoutesFromVPNServer, advertiseRoutesToDirectLink),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISVPCRouteTableExists("ibm_is_vpc_routing_table.test_ibm_is_vpc_routing_table", vpcRouteTables),
					resource.TestCheckResourceAttr(
						"ibm_is_vpc_routing_table.test_ibm_is_vpc_routing_table", "advertise_routes_to.0", advertiseRoutesToDirectLink),
					resource.TestCheckResourceAttr(
						"ibm_is_vpc_routing_table.test_ibm_is_vpc_routing_table", "accept_routes_from_resource_type.0", acceptRoutesFromVPNServer),
				),
			},
			{
				Config: testAccCheckIBMISVPCRouteTableAdvertiseRoutesToRemovalConfig(routeTableName, name1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISVPCRouteTableExists("ibm_is_vpc_routing_table.test_ibm_is_vpc_routing_table", vpcRouteTables),
					resource.TestCheckResourceAttr(
						"ibm_is_vpc_routing_table.test_ibm_is_vpc_routing_table", "advertise_routes_to.#", "0"),
					resource.TestCheckResourceAttr(
						"ibm_is_vpc_routing_table.test_ibm_is_vpc_routing_table", "accept_routes_from_resource_type.#", "0"),
				),
			},
		},
	})
}

func testAccCheckIBMISVPCRouteTableDestroy(s *terraform.State) error {
	//userDetails, _ := acc.TestAccProvider.Meta().(conns.ClientSession).BluemixUserDetails()

	//sess, _ := acc.TestAccProvider.Meta().(conns.ClientSession).VpcV1API()
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_is_vpc_routing_table" {
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
		getVpcRoutingTableOptions := sess.NewGetVPCRoutingTableOptions(vpcID, routeTableID)
		_, _, err = sess.GetVPCRoutingTable(getVpcRoutingTableOptions)
		if err == nil {
			return fmt.Errorf("Routing table still exists: %s", rs.Primary.ID)
		}
	}

	return nil
}

func testAccCheckIBMISVPCRouteTableExists(n, vpcrouteTableID string) resource.TestCheckFunc {
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
		sess, err := vpcClient(acc.TestAccProvider.Meta())
		getVpcRoutingTableOptions := sess.NewGetVPCRoutingTableOptions(vpcID, routeTableID)
		rtResponse, detail, err := sess.GetVPCRoutingTable(getVpcRoutingTableOptions)
		if err != nil {
			return fmt.Errorf("[ERROR] Error Getting Flow log: %s\n%s", err, detail)
		}
		vpcrouteTableID = *rtResponse.ID
		return nil
	}
}

func testAccCheckIBMISVPCRouteTableConfig(rtName, name string) string {
	return fmt.Sprintf(`
resource "ibm_is_vpc" "testacc_vpc" {
	name = "%s"
}
resource "ibm_is_vpc_routing_table" "test_ibm_is_vpc_routing_table" {
	depends_on = [ibm_is_vpc.testacc_vpc]
	vpc = ibm_is_vpc.testacc_vpc.id
	name = "%s"
}`, name, rtName)
}

func testAccCheckIBMISVPCRouteTableAcceptRoutesFromConfig(rtName, name, acceptRoutesFromVPNServer string) string {
	return fmt.Sprintf(`
resource "ibm_is_vpc" "testacc_vpc" {
	name = "%s"
}
resource "ibm_is_vpc_routing_table" "test_ibm_is_vpc_routing_table" {
	depends_on = [ibm_is_vpc.testacc_vpc]
	vpc = ibm_is_vpc.testacc_vpc.id
	name = "%s"
	accept_routes_from_resource_type=["%s"]
}`, name, rtName, acceptRoutesFromVPNServer)
}

func testAccCheckIBMISVPCRouteTableAdvertiseRoutesToConfig(rtName, name, acceptRoutesFromVPNServer, advertiseRoutesTo1, advertiseRoutesTo2 string) string {
	return fmt.Sprintf(`
resource "ibm_is_vpc" "testacc_vpc" {
	name = "%s"
}
resource "ibm_is_vpc_routing_table" "test_ibm_is_vpc_routing_table" {
	depends_on = [ibm_is_vpc.testacc_vpc]
	route_direct_link_ingress = true
	route_transit_gateway_ingress = true
	vpc = ibm_is_vpc.testacc_vpc.id
	name = "%s"
	accept_routes_from_resource_type=["%s"]
	advertise_routes_to=["%s","%s"]
}`, name, rtName, acceptRoutesFromVPNServer, advertiseRoutesTo1, advertiseRoutesTo2)
}

func testAccCheckIBMISVPCRouteTableAdvertiseRoutesToDLConfig(rtName, name, acceptRoutesFromVPNServer, advertiseRoutesTo1 string) string {
	return fmt.Sprintf(`
resource "ibm_is_vpc" "testacc_vpc" {
	name = "%s"
}
resource "ibm_is_vpc_routing_table" "test_ibm_is_vpc_routing_table" {
	depends_on = [ibm_is_vpc.testacc_vpc]
	route_direct_link_ingress = true
	route_transit_gateway_ingress = true
	vpc = ibm_is_vpc.testacc_vpc.id
	name = "%s"
	accept_routes_from_resource_type=["%s"]
	advertise_routes_to=["%s"]
}`, name, rtName, acceptRoutesFromVPNServer, advertiseRoutesTo1)
}

func testAccCheckIBMISVPCRouteTableAdvertiseRoutesToRemovalConfig(rtName, name string) string {
	return fmt.Sprintf(`
resource "ibm_is_vpc" "testacc_vpc" {
	name = "%s"
}
resource "ibm_is_vpc_routing_table" "test_ibm_is_vpc_routing_table" {
	depends_on = [ibm_is_vpc.testacc_vpc]
	route_direct_link_ingress = true
	vpc = ibm_is_vpc.testacc_vpc.id
	name = "%s"
	accept_routes_from_resource_type=[]
	advertise_routes_to=[]
}`, name, rtName)
}
