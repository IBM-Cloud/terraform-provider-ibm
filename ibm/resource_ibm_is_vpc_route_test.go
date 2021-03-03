// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"errors"
	"fmt"
	"testing"

	"github.com/IBM/vpc-go-sdk/vpcclassicv1"
	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccIBMISVPCRoute_basic(t *testing.T) {
	var vpcRoute string
	name1 := fmt.Sprintf("tfvpcuat-create-%d", acctest.RandIntRange(10, 100))
	subnetName := fmt.Sprintf("tfsubnet-%d", acctest.RandIntRange(10, 100))
	routeName := fmt.Sprintf("tfvpcuat-create-%d", acctest.RandIntRange(10, 100))
	routeName1 := fmt.Sprintf("tfvpcuat-create-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMISVPCRouteDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISVPCRouteConfig(name1, subnetName, routeName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISVPCRouteExists("ibm_is_vpc_route.testacc_vpc_route", vpcRoute),
					resource.TestCheckResourceAttr(
						"ibm_is_vpc_route.testacc_vpc_route", "name", routeName),
				),
			},
			{
				Config: testAccCheckIBMISVPCRouteConfig(name1, subnetName, routeName1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISVPCRouteExists("ibm_is_vpc_route.testacc_vpc_route", vpcRoute),
					resource.TestCheckResourceAttr(
						"ibm_is_vpc_route.testacc_vpc_route", "name", routeName1),
				),
			},
		},
	})
}

func testAccCheckIBMISVPCRouteDestroy(s *terraform.State) error {
	userDetails, _ := testAccProvider.Meta().(ClientSession).BluemixUserDetails()

	if userDetails.generation == 1 {
		sess, _ := testAccProvider.Meta().(ClientSession).VpcClassicV1API()
		for _, rs := range s.RootModule().Resources {
			if rs.Type != "ibm_is_vpc_route" {
				continue
			}
			parts, err := idParts(rs.Primary.ID)
			if err != nil {
				return err
			}

			vpcID := parts[0]
			routeID := parts[1]
			getVpcRouteOptions := &vpcclassicv1.GetVPCRouteOptions{
				VPCID: &vpcID,
				ID:    &routeID,
			}
			_, _, err1 := sess.GetVPCRoute(getVpcRouteOptions)

			if err1 == nil {
				return fmt.Errorf("vpc route still exists: %s", rs.Primary.ID)
			}
		}
	} else {
		sess, _ := testAccProvider.Meta().(ClientSession).VpcV1API()
		for _, rs := range s.RootModule().Resources {
			if rs.Type != "ibm_is_vpc_route" {
				continue
			}
			parts, err := idParts(rs.Primary.ID)
			if err != nil {
				return err
			}

			vpcID := parts[0]
			routeID := parts[1]
			getVpcRouteOptions := &vpcv1.GetVPCRouteOptions{
				VPCID: &vpcID,
				ID:    &routeID,
			}
			_, _, err1 := sess.GetVPCRoute(getVpcRouteOptions)

			if err1 == nil {
				return fmt.Errorf("vpc route still exists: %s", rs.Primary.ID)
			}
		}
	}
	return nil
}

func testAccCheckIBMISVPCRouteExists(n, vpcrouteID string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
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
		routeID := parts[1]
		userDetails, _ := testAccProvider.Meta().(ClientSession).BluemixUserDetails()

		if userDetails.generation == 1 {
			sess, _ := testAccProvider.Meta().(ClientSession).VpcClassicV1API()
			getVpcRouteOptions := &vpcclassicv1.GetVPCRouteOptions{
				VPCID: &vpcID,
				ID:    &routeID,
			}
			foundroute, _, err := sess.GetVPCRoute(getVpcRouteOptions)
			if err != nil {
				return err
			}
			vpcrouteID = *foundroute.ID
		} else {
			sess, _ := testAccProvider.Meta().(ClientSession).VpcV1API()
			getVpcRouteOptions := &vpcv1.GetVPCRouteOptions{
				VPCID: &vpcID,
				ID:    &routeID,
			}
			foundroute, _, err := sess.GetVPCRoute(getVpcRouteOptions)
			if err != nil {
				return err
			}
			vpcrouteID = *foundroute.ID
		}
		return nil
	}
}

func testAccCheckIBMISVPCRouteConfig(name, subnetName, routeName string) string {
	return fmt.Sprintf(`
resource "ibm_is_vpc" "testacc_vpc" {
    name = "%s"
}

resource "ibm_is_subnet" "testacc_subnet" {
	name = "%s"
	vpc = ibm_is_vpc.testacc_vpc.id
	zone = "%s"
	ipv4_cidr_block = "%s"
}

resource "ibm_is_vpc_route" "testacc_vpc_route" {
    name = "%s"
    zone = "%s"
    vpc = ibm_is_vpc.testacc_vpc.id
	destination = "%s"
	next_hop = "%s"
	depends_on  = [ibm_is_subnet.testacc_subnet]
}`, name, subnetName, ISZoneName, ISCIDR, routeName, ISZoneName, ISRouteDestination, ISRouteNextHop)
}
