package ibm

import (
	"errors"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.ibm.com/Bluemix/riaas-go-client/clients/lbaas"
	"github.ibm.com/riaas/rias-api/riaas/models"
)

func TestAccIBMISLBListener_basic(t *testing.T) {
	var lb *models.Listener
	vpcname := fmt.Sprintf("terraformLBLisuat_vpc_%d", acctest.RandInt())
	subnetname := fmt.Sprintf("terraformLBLisuat_subnet_%d", acctest.RandInt())
	lbname := fmt.Sprintf("terraformLBLisuat_lb_%d", acctest.RandInt())

	protocol1 := "http"
	port1 := "8080"

	protocol2 := "tcp"
	port2 := "9080"

	connLimit := "100"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMISLBListenerDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMISLBListenerConfig(vpcname, subnetname, ISZoneName, ISCIDR, lbname, port1, protocol1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISLBListenerExists("ibm_is_lb_listener.testacc_lb_listener", &lb),
					resource.TestCheckResourceAttr(
						"ibm_is_lb.testacc_LB", "name", lbname),
					resource.TestCheckResourceAttr(
						"ibm_is_lb_listener.testacc_lb_listener", "port", port1),
					resource.TestCheckResourceAttr(
						"ibm_is_lb_listener.testacc_lb_listener", "protocol", protocol1),
				),
			},

			resource.TestStep{
				Config: testAccCheckIBMISLBListenerConfigUpdate(vpcname, subnetname, ISZoneName, ISCIDR, lbname, port2, protocol2, connLimit),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISLBListenerExists("ibm_is_lb_listener.testacc_lb_listener", &lb),
					resource.TestCheckResourceAttr(
						"ibm_is_lb_listener.testacc_lb_listener", "port", port2),
					resource.TestCheckResourceAttr(
						"ibm_is_lb_listener.testacc_lb_listener", "protocol", protocol2),
					resource.TestCheckResourceAttr(
						"ibm_is_lb_listener.testacc_lb_listener", "connection_limit", connLimit),
				),
			},
		},
	})
}

func testAccCheckIBMISLBListenerDestroy(s *terraform.State) error {
	sess, _ := testAccProvider.Meta().(ClientSession).ISSession()

	lbc := lbaas.NewLoadBalancerClient(sess)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_is_lb_listener" {
			continue
		}
		parts, err := idParts(rs.Primary.ID)
		if err != nil {
			return err
		}

		lbID := parts[0]
		lbListenerID := parts[1]
		_, err = lbc.GetListener(lbID, lbListenerID)

		if err == nil {
			return fmt.Errorf("LB Listener still exists: %s", rs.Primary.ID)
		}
	}

	return nil
}

func testAccCheckIBMISLBListenerExists(n string, LBListener **models.Listener) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return errors.New("No Record ID is set")
		}

		sess, _ := testAccProvider.Meta().(ClientSession).ISSession()
		client := lbaas.NewLoadBalancerClient(sess)
		parts, err := idParts(rs.Primary.ID)
		if err != nil {
			return err
		}

		lbID := parts[0]
		lbListenerID := parts[1]
		foundLBListener, err := client.GetListener(lbID, lbListenerID)

		if err != nil {
			return err
		}

		*LBListener = foundLBListener
		return nil
	}
}

func testAccCheckIBMISLBListenerConfig(vpcname, subnetname, zone, cidr, lbname, port, protocol string) string {
	return fmt.Sprintf(`

	resource "ibm_is_vpc" "testacc_vpc" {
		name = "%s"
	}

	resource "ibm_is_subnet" "testacc_subnet" {
		name = "%s"
		vpc = "${ibm_is_vpc.testacc_vpc.id}"
		zone = "%s"
		ipv4_cidr_block = "%s"
	}
	resource "ibm_is_lb" "testacc_LB" {
		name = "%s"
		subnets = ["${ibm_is_subnet.testacc_subnet.id}"]
	}
	resource "ibm_is_lb_listener" "testacc_lb_listener" {
		lb = "${ibm_is_lb.testacc_LB.id}"
		port = %s
		protocol = "%s"
}`, vpcname, subnetname, zone, cidr, lbname, port, protocol)

}

func testAccCheckIBMISLBListenerConfigUpdate(vpcname, subnetname, zone, cidr, lbname, port, protocol, connLimit string) string {
	return fmt.Sprintf(`

	resource "ibm_is_vpc" "testacc_vpc" {
		name = "%s"
	}

	resource "ibm_is_subnet" "testacc_subnet" {
		name = "%s"
		vpc = "${ibm_is_vpc.testacc_vpc.id}"
		zone = "%s"
		ipv4_cidr_block = "%s"
	}
	resource "ibm_is_lb" "testacc_LB" {
		name = "%s"
		subnets = ["${ibm_is_subnet.testacc_subnet.id}"]
	}
	resource "ibm_is_lb_listener" "testacc_lb_listener" {
		lb = "${ibm_is_lb.testacc_LB.id}"
		port = %s
		protocol = "%s"
		connection_limit = %s
}`, vpcname, subnetname, zone, cidr, lbname, port, protocol, connLimit)

}
