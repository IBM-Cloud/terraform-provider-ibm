/**
 * © Copyright IBM Corporation 2020. All Rights Reserved.
 *
 * Licensed under the Mozilla Public License, version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at https://mozilla.org/MPL/2.0/
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

func TestAccIBMISVPCRoutingTable_basic(t *testing.T) {
	var vpcRouteTables string
	name1 := fmt.Sprintf("tfvpc-create-%d", acctest.RandIntRange(10, 100))
	routeTableName := fmt.Sprintf("tfvpcrt-create-%d", acctest.RandIntRange(10, 100))
	routeTableName1 := fmt.Sprintf("tfvpcrt-up-create-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
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

func testAccCheckIBMISVPCRouteTableDestroy(s *terraform.State) error {
	//userDetails, _ := testAccProvider.Meta().(ClientSession).BluemixUserDetails()

	//sess, _ := testAccProvider.Meta().(ClientSession).VpcV1API()
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_is_vpc_routing_table" {
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
		sess, err := vpcClient(testAccProvider.Meta())
		getVpcRoutingTableOptions := sess.NewGetVPCRoutingTableOptions(vpcID, routeTableID)
		rtResponse, detail, err := sess.GetVPCRoutingTable(getVpcRoutingTableOptions)
		if err != nil {
			return fmt.Errorf("Error Getting Flow log: %s\n%s", err, detail)
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
