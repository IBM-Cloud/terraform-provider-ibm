package ibm

import (
	"errors"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.ibm.com/Bluemix/riaas-go-client/clients/network"
	"github.ibm.com/Bluemix/riaas-go-client/riaas/models"
)

func TestAccIBMISVPCRoute_basic(t *testing.T) {
	var vpcRoute *models.Route
	name1 := fmt.Sprintf("terraformvpcuat-create-step-name-%d", acctest.RandInt())
	routeName := fmt.Sprintf("terraformvpcuat-create-prefix-name-%d", acctest.RandInt())
	routeName1 := fmt.Sprintf("terraformvpcuat-create-prefix-name-%d", acctest.RandInt())

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMISVPCRouteDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISVPCRouteConfig(name1, routeName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISVPCRouteExists("ibm_is_vpc_route.testacc_vpc_route", &vpcRoute),
					resource.TestCheckResourceAttr(
						"ibm_is_vpc_route.testacc_vpc_route", "name", routeName),
				),
			},
			{
				Config: testAccCheckIBMISVPCRouteConfig(name1, routeName1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISVPCRouteExists("ibm_is_vpc_route.testacc_vpc_route", &vpcRoute),
					resource.TestCheckResourceAttr(
						"ibm_is_vpc_route.testacc_vpc_route", "name", routeName1),
				),
			},
		},
	})
}

func testAccCheckIBMISVPCRouteDestroy(s *terraform.State) error {
	sess, _ := testAccProvider.Meta().(ClientSession).ISSession()

	vpcC := network.NewVPCClient(sess)

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
		_, err = vpcC.GetRoute(vpcID, routeID)

		if err == nil {
			return fmt.Errorf("vpc route still exists: %s", rs.Primary.ID)
		}
	}

	return nil
}

func testAccCheckIBMISVPCRouteExists(n string, vpcRoute **models.Route) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return errors.New("No Record ID is set")
		}

		sess, _ := testAccProvider.Meta().(ClientSession).ISSession()
		vpcC := network.NewVPCClient(sess)

		parts, err := idParts(rs.Primary.ID)
		if err != nil {
			return err
		}

		vpcID := parts[0]
		routeID := parts[1]
		route, err := vpcC.GetRoute(vpcID, routeID)

		if err != nil {
			return err
		}

		*vpcRoute = route
		return nil
	}
}

func testAccCheckIBMISVPCRouteConfig(name, routeName string) string {
	return fmt.Sprintf(`
resource "ibm_is_vpc" "testacc_vpc" {
    name = "%s"
}
resource "ibm_is_vpc_route" "testacc_vpc_route" {
    name = "%s"
    zone = "%s"
    vpc = "${ibm_is_vpc.testacc_vpc.id}"
	destination = "%s"
	next_hop = "%s"
}`, name, routeName, ISZoneName, ISRouteDestination, ISRouteNextHop)
}
