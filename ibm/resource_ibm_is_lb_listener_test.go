// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"errors"
	"fmt"
	"testing"

	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccIBMISLBListener_basic(t *testing.T) {
	var lb string
	vpcname := fmt.Sprintf("tflblis-vpc-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tflblis-subnet-%d", acctest.RandIntRange(10, 100))
	lbname := fmt.Sprintf("tflblis%d", acctest.RandIntRange(10, 100))

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
			{
				Config: testAccCheckIBMISLBListenerConfig(vpcname, subnetname, ISZoneName, ISCIDR, lbname, port1, protocol1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISLBListenerExists("ibm_is_lb_listener.testacc_lb_listener", lb),
					resource.TestCheckResourceAttr(
						"ibm_is_lb.testacc_LB", "name", lbname),
					resource.TestCheckResourceAttr(
						"ibm_is_lb_listener.testacc_lb_listener", "port", port1),
					resource.TestCheckResourceAttr(
						"ibm_is_lb_listener.testacc_lb_listener", "protocol", protocol1),
					resource.TestCheckResourceAttr(
						"ibm_is_lb_listener.testacc_lb_listener", "accept_proxy_protocol", "true"),
				),
			},

			{
				Config: testAccCheckIBMISLBListenerConfigUpdate(vpcname, subnetname, ISZoneName, ISCIDR, lbname, port2, protocol2, connLimit),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISLBListenerExists("ibm_is_lb_listener.testacc_lb_listener", lb),
					resource.TestCheckResourceAttr(
						"ibm_is_lb_listener.testacc_lb_listener", "port", port2),
					resource.TestCheckResourceAttr(
						"ibm_is_lb_listener.testacc_lb_listener", "protocol", protocol2),
					resource.TestCheckResourceAttr(
						"ibm_is_lb_listener.testacc_lb_listener", "connection_limit", connLimit),
					resource.TestCheckResourceAttr(
						"ibm_is_lb_listener.testacc_lb_listener", "accept_proxy_protocol", "false"),
				),
			},
		},
	})
}
func TestAccIBMISNLBRouteModeListener_basic(t *testing.T) {
	var lb string
	vpcname := fmt.Sprintf("tflblis-vpc-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tflblis-subnet-%d", acctest.RandIntRange(10, 100))
	lbname := fmt.Sprintf("tflblis%d", acctest.RandIntRange(10, 100))

	protocol1 := "tcp"
	port1 := "1"
	port2 := "65535"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMISLBListenerDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISNLBRouteModeListenerConfig(vpcname, subnetname, ISZoneName, ISCIDR, lbname, port1, protocol1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISLBListenerExists("ibm_is_lb_listener.testacc_lb_listener", lb),
					resource.TestCheckResourceAttr(
						"ibm_is_lb.testacc_LB", "name", lbname),
					resource.TestCheckResourceAttr(
						"ibm_is_lb.testacc_LB", "type", "private"),
					resource.TestCheckResourceAttr(
						"ibm_is_lb.testacc_LB", "route_mode", "true"),
					resource.TestCheckResourceAttr(
						"ibm_is_lb_listener.testacc_lb_listener", "port", port1),
					resource.TestCheckResourceAttr(
						"ibm_is_lb_listener.testacc_lb_listener", "port_min", port1),
					resource.TestCheckResourceAttr(
						"ibm_is_lb_listener.testacc_lb_listener", "port_max", port2),
					resource.TestCheckResourceAttr(
						"ibm_is_lb_listener.testacc_lb_listener", "protocol", protocol1),
				),
			},
		},
	})
}

func TestAccIBMISLBListenerHttpRedirect_basic(t *testing.T) {
	var lb string
	vpcname := fmt.Sprintf("tflblis-vpc-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tflblis-subnet-%d", acctest.RandIntRange(10, 100))
	lbname := fmt.Sprintf("tflblis%d", acctest.RandIntRange(10, 100))

	protocol1 := "https"
	port1 := "9086"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMISLBListenerDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISLBListenerHttpsRedirectConfig(vpcname, subnetname, ISZoneName, ISCIDR, lbname, port1, protocol1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISLBListenerExists("ibm_is_lb_listener.lb_listener2", lb),
					resource.TestCheckResourceAttr(
						"ibm_is_lb.testacc_LB", "name", lbname),
					resource.TestCheckResourceAttr(
						"ibm_is_lb_listener.lb_listener2", "https_redirect_status_code", "301"),
					resource.TestCheckResourceAttr(
						"ibm_is_lb_listener.lb_listener2", "https_redirect_uri", "/example?doc=geta"),
				),
			},
			{
				Config: testAccCheckIBMISLBListenerHttpsRedirectConfigUpdate(vpcname, subnetname, ISZoneName, ISCIDR, lbname, port1, protocol1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISLBListenerExists("ibm_is_lb_listener.lb_listener2", lb),
					resource.TestCheckResourceAttr(
						"ibm_is_lb.testacc_LB", "name", lbname),
					resource.TestCheckResourceAttr(
						"ibm_is_lb_listener.lb_listener2", "https_redirect_status_code", "303"),
					resource.TestCheckResourceAttr(
						"ibm_is_lb_listener.lb_listener2", "https_redirect_uri", "/example?doc=updated"),
				),
			},
			{
				Config: testAccCheckIBMISLBListenerHttpsRedirectConfigRemove(vpcname, subnetname, ISZoneName, ISCIDR, lbname, port1, protocol1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISLBListenerExists("ibm_is_lb_listener.lb_listener2", lb),
					resource.TestCheckResourceAttr(
						"ibm_is_lb.testacc_LB", "name", lbname),
					resource.TestCheckResourceAttr(
						"ibm_is_lb_listener.lb_listener2", "https_redirect_uri", ""),
					resource.TestCheckResourceAttr(
						"ibm_is_lb_listener.lb_listener2", "https_redirect_listener", ""),
				),
			},
		},
	})
}

func testAccCheckIBMISLBListenerDestroy(s *terraform.State) error {

	sess, _ := testAccProvider.Meta().(ClientSession).VpcV1API()
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
		getlblptions := &vpcv1.GetLoadBalancerListenerOptions{
			LoadBalancerID: &lbID,
			ID:             &lbListenerID,
		}
		_, _, err1 := sess.GetLoadBalancerListener(getlblptions)
		if err1 == nil {
			return fmt.Errorf("LB Listener still exists: %s", rs.Primary.ID)
		}
	}

	return nil
}

func testAccCheckIBMISLBListenerExists(n, LBListener string) resource.TestCheckFunc {
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

		lbID := parts[0]
		lbListenerID := parts[1]

		sess, _ := testAccProvider.Meta().(ClientSession).VpcV1API()
		getlblptions := &vpcv1.GetLoadBalancerListenerOptions{
			LoadBalancerID: &lbID,
			ID:             &lbListenerID,
		}
		foundLBListener, _, err := sess.GetLoadBalancerListener(getlblptions)
		if err != nil {
			return err
		}
		LBListener = *foundLBListener.ID

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
		accept_proxy_protocol = true
    }`, vpcname, subnetname, zone, cidr, lbname, port, protocol)

}

func testAccCheckIBMISLBListenerHttpsRedirectConfig(vpcname, subnetname, zone, cidr, lbname, port, protocol string) string {
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
	resource "ibm_is_lb_listener" "lb_listener1"{
		lb       = ibm_is_lb.testacc_LB.id
		port     = "9086"
		protocol = "https"
		certificate_instance="%s"
	  }
	  
	  resource "ibm_is_lb_listener" "lb_listener2"{
		lb       = ibm_is_lb.testacc_LB.id
		port     = "9087"
		protocol = "http"
		https_redirect_listener = ibm_is_lb_listener.lb_listener1.listener_id
		https_redirect_status_code = 301
		https_redirect_uri = "/example?doc=geta" 
	  }`, vpcname, subnetname, zone, cidr, lbname, lbListerenerCertificateInstance)

}

func testAccCheckIBMISLBListenerHttpsRedirectConfigUpdate(vpcname, subnetname, zone, cidr, lbname, port, protocol string) string {
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
	resource "ibm_is_lb_listener" "lb_listener1"{
		lb       = ibm_is_lb.testacc_LB.id
		port     = "9086"
		protocol = "https"
		certificate_instance="%s"
	  }
	  
	  resource "ibm_is_lb_listener" "lb_listener2"{
		lb       = ibm_is_lb.testacc_LB.id
		port     = "9087"
		protocol = "http"
		https_redirect_listener = ibm_is_lb_listener.lb_listener1.listener_id
		https_redirect_status_code = 303
		https_redirect_uri = "/example?doc=updated" 
	  }`, vpcname, subnetname, zone, cidr, lbname, lbListerenerCertificateInstance)

}

func testAccCheckIBMISLBListenerHttpsRedirectConfigRemove(vpcname, subnetname, zone, cidr, lbname, port, protocol string) string {
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
	resource "ibm_is_lb_listener" "lb_listener1"{
		lb       = ibm_is_lb.testacc_LB.id
		port     = "9086"
		protocol = "https"
		certificate_instance="%s"
	  }
	  
	  resource "ibm_is_lb_listener" "lb_listener2"{
		lb       = ibm_is_lb.testacc_LB.id
		port     = "9087"
		protocol = "http"
	  }`, vpcname, subnetname, zone, cidr, lbname, lbListerenerCertificateInstance)

}
func testAccCheckIBMISNLBRouteModeListenerConfig(vpcname, subnetname, zone, cidr, lbname, port, protocol string) string {
	return fmt.Sprintf(`
	resource "ibm_is_vpc" "testacc_vpc" {
		name = "%s"
	}

	resource "ibm_is_subnet" "testacc_subnet" {
		name 			= "%s"
		vpc 			= "${ibm_is_vpc.testacc_vpc.id}"
		zone 			= "%s"
		ipv4_cidr_block = "%s"
	}
	resource "ibm_is_lb" "testacc_LB" {
		name			= "%s"
		subnets 		= ["${ibm_is_subnet.testacc_subnet.id}"]
		profile 		= "network-fixed"
		route_mode 		= true
		type 			= "private"
	}
	resource "ibm_is_lb_listener" "testacc_lb_listener" {
		lb 			= "${ibm_is_lb.testacc_LB.id}"
		port 		= %s
		protocol 	= "%s"
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
		accept_proxy_protocol = false
}`, vpcname, subnetname, zone, cidr, lbname, port, protocol, connLimit)

}
