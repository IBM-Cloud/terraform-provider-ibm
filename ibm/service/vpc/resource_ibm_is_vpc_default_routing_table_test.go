// Copyright IBM Corp. 2017, 2025 All Rights Reserved.
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

func TestAccIBMISVPCDefaultRoutingTable_basic(t *testing.T) {
	var vpcRouteTableID string
	name1 := fmt.Sprintf("tfvpc-create-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISVPCDefaultRouteTableDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISVPCDefaultRouteTableConfig(name1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISVPCDefaultRouteTableExists("ibm_is_vpc_default_routing_table.test_ibm_is_vpc_default_routing_table", vpcRouteTableID),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpc_default_routing_table.test_ibm_is_vpc_default_routing_table", "name"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpc_default_routing_table.test_ibm_is_vpc_default_routing_table", "routing_table"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpc_default_routing_table.test_ibm_is_vpc_default_routing_table", "href"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpc_default_routing_table.test_ibm_is_vpc_default_routing_table", "lifecycle_state"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpc_default_routing_table.test_ibm_is_vpc_default_routing_table", "resource_type"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpc_default_routing_table.test_ibm_is_vpc_default_routing_table", "created_at"),
					resource.TestCheckResourceAttr(
						"ibm_is_vpc_default_routing_table.test_ibm_is_vpc_default_routing_table", "is_default", "true"),
					resource.TestCheckResourceAttr(
						"ibm_is_vpc_default_routing_table.test_ibm_is_vpc_default_routing_table", "route_direct_link_ingress", "false"),
					resource.TestCheckResourceAttr(
						"ibm_is_vpc_default_routing_table.test_ibm_is_vpc_default_routing_table", "route_internet_ingress", "false"),
					resource.TestCheckResourceAttr(
						"ibm_is_vpc_default_routing_table.test_ibm_is_vpc_default_routing_table", "route_transit_gateway_ingress", "false"),
					resource.TestCheckResourceAttr(
						"ibm_is_vpc_default_routing_table.test_ibm_is_vpc_default_routing_table", "route_vpc_zone_ingress", "false"),
				),
			},
		},
	})
}

func TestAccIBMISVPCDefaultRoutingTable_acceptRoutesFrom(t *testing.T) {
	var vpcRouteTableID string
	name1 := fmt.Sprintf("tfvpc-create-%d", acctest.RandIntRange(10, 100))
	acceptRoutesFromVPNServer := "vpn_server"
	acceptRoutesFromVPNGateway := "vpn_gateway"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISVPCDefaultRouteTableDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISVPCDefaultRouteTableAcceptRoutesFromConfig(name1, acceptRoutesFromVPNServer),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISVPCDefaultRouteTableExists("ibm_is_vpc_default_routing_table.test_ibm_is_vpc_default_routing_table", vpcRouteTableID),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpc_default_routing_table.test_ibm_is_vpc_default_routing_table", "name"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpc_default_routing_table.test_ibm_is_vpc_default_routing_table", "accept_routes_from_resource_type.#"),
					resource.TestCheckResourceAttr(
						"ibm_is_vpc_default_routing_table.test_ibm_is_vpc_default_routing_table", "accept_routes_from_resource_type.0", acceptRoutesFromVPNServer),
				),
			},
			{
				Config: testAccCheckIBMISVPCDefaultRouteTableAcceptRoutesFromConfig(name1, acceptRoutesFromVPNGateway),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISVPCDefaultRouteTableExists("ibm_is_vpc_default_routing_table.test_ibm_is_vpc_default_routing_table", vpcRouteTableID),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpc_default_routing_table.test_ibm_is_vpc_default_routing_table", "name"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpc_default_routing_table.test_ibm_is_vpc_default_routing_table", "accept_routes_from_resource_type.#"),
					resource.TestCheckResourceAttr(
						"ibm_is_vpc_default_routing_table.test_ibm_is_vpc_default_routing_table", "accept_routes_from_resource_type.0", acceptRoutesFromVPNGateway),
				),
			},
			{
				ResourceName:      "ibm_is_vpc_default_routing_table.test_ibm_is_vpc_default_routing_table",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccIBMISVPCDefaultRoutingTable_advertiseRoutesTo(t *testing.T) {
	var vpcRouteTableID string
	name1 := fmt.Sprintf("tfvpc-create-%d", acctest.RandIntRange(10, 100))
	advertiseRoutesToDirectLink := "direct_link"
	advertiseRoutesToTransitGateway := "transit_gateway"
	acceptRoutesFromVPNServer := "vpn_server"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISVPCDefaultRouteTableDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISVPCDefaultRouteTableAdvertiseRoutesToConfig(name1, acceptRoutesFromVPNServer, advertiseRoutesToDirectLink, advertiseRoutesToTransitGateway),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISVPCDefaultRouteTableExists("ibm_is_vpc_default_routing_table.test_ibm_is_vpc_default_routing_table", vpcRouteTableID),
					resource.TestCheckResourceAttr(
						"ibm_is_vpc_default_routing_table.test_ibm_is_vpc_default_routing_table", "advertise_routes_to.0", advertiseRoutesToDirectLink),
					resource.TestCheckResourceAttr(
						"ibm_is_vpc_default_routing_table.test_ibm_is_vpc_default_routing_table", "advertise_routes_to.1", advertiseRoutesToTransitGateway),
					resource.TestCheckResourceAttr(
						"ibm_is_vpc_default_routing_table.test_ibm_is_vpc_default_routing_table", "accept_routes_from_resource_type.0", acceptRoutesFromVPNServer),
					resource.TestCheckResourceAttr(
						"ibm_is_vpc_default_routing_table.test_ibm_is_vpc_default_routing_table", "route_direct_link_ingress", "true"),
					resource.TestCheckResourceAttr(
						"ibm_is_vpc_default_routing_table.test_ibm_is_vpc_default_routing_table", "route_transit_gateway_ingress", "true"),
				),
			},
			{
				Config: testAccCheckIBMISVPCDefaultRouteTableAdvertiseRoutesToDLConfig(name1, acceptRoutesFromVPNServer, advertiseRoutesToDirectLink),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISVPCDefaultRouteTableExists("ibm_is_vpc_default_routing_table.test_ibm_is_vpc_default_routing_table", vpcRouteTableID),
					resource.TestCheckResourceAttr(
						"ibm_is_vpc_default_routing_table.test_ibm_is_vpc_default_routing_table", "advertise_routes_to.0", advertiseRoutesToDirectLink),
					resource.TestCheckResourceAttr(
						"ibm_is_vpc_default_routing_table.test_ibm_is_vpc_default_routing_table", "accept_routes_from_resource_type.0", acceptRoutesFromVPNServer),
					resource.TestCheckResourceAttr(
						"ibm_is_vpc_default_routing_table.test_ibm_is_vpc_default_routing_table", "route_direct_link_ingress", "true"),
				),
			},
			{
				Config: testAccCheckIBMISVPCDefaultRouteTableAdvertiseRoutesToRemovalConfig(name1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISVPCDefaultRouteTableExists("ibm_is_vpc_default_routing_table.test_ibm_is_vpc_default_routing_table", vpcRouteTableID),
					resource.TestCheckResourceAttr(
						"ibm_is_vpc_default_routing_table.test_ibm_is_vpc_default_routing_table", "advertise_routes_to.#", "0"),
					resource.TestCheckResourceAttr(
						"ibm_is_vpc_default_routing_table.test_ibm_is_vpc_default_routing_table", "accept_routes_from_resource_type.#", "0"),
					resource.TestCheckResourceAttr(
						"ibm_is_vpc_default_routing_table.test_ibm_is_vpc_default_routing_table", "route_direct_link_ingress", "true"),
				),
			},
		},
	})
}

func TestAccIBMISVPCDefaultRoutingTable_routeIngress(t *testing.T) {
	var vpcRouteTableID string
	name1 := fmt.Sprintf("tfvpc-create-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISVPCDefaultRouteTableDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISVPCDefaultRouteTableRouteIngressConfig(name1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISVPCDefaultRouteTableExists("ibm_is_vpc_default_routing_table.test_ibm_is_vpc_default_routing_table", vpcRouteTableID),
					resource.TestCheckResourceAttr(
						"ibm_is_vpc_default_routing_table.test_ibm_is_vpc_default_routing_table", "route_direct_link_ingress", "true"),
					resource.TestCheckResourceAttr(
						"ibm_is_vpc_default_routing_table.test_ibm_is_vpc_default_routing_table", "route_transit_gateway_ingress", "true"),
					resource.TestCheckResourceAttr(
						"ibm_is_vpc_default_routing_table.test_ibm_is_vpc_default_routing_table", "route_vpc_zone_ingress", "true"),
				),
			},
			{
				Config: testAccCheckIBMISVPCDefaultRouteTableRouteIngressUpdateConfig(name1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISVPCDefaultRouteTableExists("ibm_is_vpc_default_routing_table.test_ibm_is_vpc_default_routing_table", vpcRouteTableID),
					resource.TestCheckResourceAttr(
						"ibm_is_vpc_default_routing_table.test_ibm_is_vpc_default_routing_table", "route_direct_link_ingress", "false"),
					resource.TestCheckResourceAttr(
						"ibm_is_vpc_default_routing_table.test_ibm_is_vpc_default_routing_table", "route_transit_gateway_ingress", "false"),
					resource.TestCheckResourceAttr(
						"ibm_is_vpc_default_routing_table.test_ibm_is_vpc_default_routing_table", "route_vpc_zone_ingress", "false"),
				),
			},
		},
	})
}

func testAccCheckIBMISVPCDefaultRouteTableDestroy(s *terraform.State) error {
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_is_vpc_default_routing_table" {
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
		routingTable, _, err := sess.GetVPCRoutingTable(getVpcRoutingTableOptions)
		if err == nil {
			// Default routing table cannot be deleted, so we check if it exists and is still default
			if routingTable != nil && *routingTable.IsDefault {
				// This is expected - default routing table should still exist
				continue
			}
			return fmt.Errorf("Default routing table still exists but is not marked as default: %s", rs.Primary.ID)
		}
	}
	return nil
}

func testAccCheckIBMISVPCDefaultRouteTableExists(n, vpcrouteTableID string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
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
		if err != nil {
			return err
		}
		getVpcRoutingTableOptions := sess.NewGetVPCRoutingTableOptions(vpcID, routeTableID)
		rtResponse, detail, err := sess.GetVPCRoutingTable(getVpcRoutingTableOptions)
		if err != nil {
			return fmt.Errorf("[ERROR] Error Getting VPC Default Routing Table: %s\n%s", err, detail)
		}

		// Verify it's actually the default routing table
		if !*rtResponse.IsDefault {
			return fmt.Errorf("[ERROR] Retrieved routing table is not the default routing table")
		}

		vpcrouteTableID = *rtResponse.ID
		return nil
	}
}

func testAccCheckIBMISVPCDefaultRouteTableConfig(name string) string {
	return fmt.Sprintf(`
resource "ibm_is_vpc" "testacc_vpc" {
	name = "%s"
}

resource "ibm_is_vpc_default_routing_table" "test_ibm_is_vpc_default_routing_table" {
	depends_on = [ibm_is_vpc.testacc_vpc]
	vpc = ibm_is_vpc.testacc_vpc.id
}`, name)
}

func testAccCheckIBMISVPCDefaultRouteTableAcceptRoutesFromConfig(name, acceptRoutesFromVPNServer string) string {
	return fmt.Sprintf(`
resource "ibm_is_vpc" "testacc_vpc" {
	name = "%s"
}

resource "ibm_is_vpc_default_routing_table" "test_ibm_is_vpc_default_routing_table" {
	depends_on = [ibm_is_vpc.testacc_vpc]
	vpc = ibm_is_vpc.testacc_vpc.id
	accept_routes_from_resource_type = ["%s"]
}`, name, acceptRoutesFromVPNServer)
}

func testAccCheckIBMISVPCDefaultRouteTableAdvertiseRoutesToConfig(name, acceptRoutesFromVPNServer, advertiseRoutesTo1, advertiseRoutesTo2 string) string {
	return fmt.Sprintf(`
resource "ibm_is_vpc" "testacc_vpc" {
	name = "%s"
}

resource "ibm_is_vpc_default_routing_table" "test_ibm_is_vpc_default_routing_table" {
	depends_on = [ibm_is_vpc.testacc_vpc]
	vpc = ibm_is_vpc.testacc_vpc.id
	route_direct_link_ingress = true
	route_transit_gateway_ingress = true
	accept_routes_from_resource_type = ["%s"]
	advertise_routes_to = ["%s", "%s"]
}`, name, acceptRoutesFromVPNServer, advertiseRoutesTo1, advertiseRoutesTo2)
}

func testAccCheckIBMISVPCDefaultRouteTableAdvertiseRoutesToDLConfig(name, acceptRoutesFromVPNServer, advertiseRoutesTo1 string) string {
	return fmt.Sprintf(`
resource "ibm_is_vpc" "testacc_vpc" {
	name = "%s"
}

resource "ibm_is_vpc_default_routing_table" "test_ibm_is_vpc_default_routing_table" {
	depends_on = [ibm_is_vpc.testacc_vpc]
	vpc = ibm_is_vpc.testacc_vpc.id
	route_direct_link_ingress = true
	route_transit_gateway_ingress = true
	accept_routes_from_resource_type = ["%s"]
	advertise_routes_to = ["%s"]
}`, name, acceptRoutesFromVPNServer, advertiseRoutesTo1)
}

func testAccCheckIBMISVPCDefaultRouteTableAdvertiseRoutesToRemovalConfig(name string) string {
	return fmt.Sprintf(`
resource "ibm_is_vpc" "testacc_vpc" {
	name = "%s"
}

resource "ibm_is_vpc_default_routing_table" "test_ibm_is_vpc_default_routing_table" {
	depends_on = [ibm_is_vpc.testacc_vpc]
	vpc = ibm_is_vpc.testacc_vpc.id
	route_direct_link_ingress = true
	accept_routes_from_resource_type = []
	advertise_routes_to = []
}`, name)
}

func testAccCheckIBMISVPCDefaultRouteTableRouteIngressConfig(name string) string {
	return fmt.Sprintf(`
resource "ibm_is_vpc" "testacc_vpc" {
	name = "%s"
}

resource "ibm_is_vpc_default_routing_table" "test_ibm_is_vpc_default_routing_table" {
	depends_on = [ibm_is_vpc.testacc_vpc]
	vpc = ibm_is_vpc.testacc_vpc.id
	route_direct_link_ingress = true
	route_transit_gateway_ingress = true
	route_vpc_zone_ingress = true
}`, name)
}

func testAccCheckIBMISVPCDefaultRouteTableRouteIngressUpdateConfig(name string) string {
	return fmt.Sprintf(`
resource "ibm_is_vpc" "testacc_vpc" {
	name = "%s"
}

resource "ibm_is_vpc_default_routing_table" "test_ibm_is_vpc_default_routing_table" {
	depends_on = [ibm_is_vpc.testacc_vpc]
	vpc = ibm_is_vpc.testacc_vpc.id
	route_direct_link_ingress = false
	route_transit_gateway_ingress = false
	route_vpc_zone_ingress = false
}`, name)
}
