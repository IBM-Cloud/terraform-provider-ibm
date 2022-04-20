// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"errors"
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"

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
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISLBListenerDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISLBListenerConfig(vpcname, subnetname, acc.ISZoneName, acc.ISCIDR, lbname, port1, protocol1),
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
				Config: testAccCheckIBMISLBListenerConfigUpdate(vpcname, subnetname, acc.ISZoneName, acc.ISCIDR, lbname, port2, protocol2, connLimit),
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
func TestAccIBMISLBListener_basic_udp(t *testing.T) {
	var lb string
	vpcname := fmt.Sprintf("tflblis-vpc-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tflblis-subnet-%d", acctest.RandIntRange(10, 100))
	lbname := fmt.Sprintf("tflblis%d", acctest.RandIntRange(10, 100))

	protocol1 := "udp"
	port1 := "8080"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISLBListenerDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISLBUdpListenerConfig(vpcname, subnetname, acc.ISZoneName, acc.ISCIDR, lbname, port1, protocol1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISLBListenerExists("ibm_is_lb_listener.testacc_lb_listener", lb),
					resource.TestCheckResourceAttr(
						"ibm_is_lb.testacc_LB", "name", lbname),
					resource.TestCheckResourceAttr(
						"ibm_is_lb_listener.testacc_lb_listener", "port", port1),
					resource.TestCheckResourceAttr(
						"ibm_is_lb_listener.testacc_lb_listener", "protocol", protocol1),
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
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISLBListenerDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISNLBRouteModeListenerConfig(vpcname, subnetname, acc.ISZoneName, acc.ISCIDR, lbname, port1, protocol1),
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
func TestAccIBMISNLBPortRangeListener_basic(t *testing.T) {
	var lb string
	vpcname := fmt.Sprintf("tflblis-vpc-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tflblis-subnet-%d", acctest.RandIntRange(10, 100))
	lbname := fmt.Sprintf("tflblis%d", acctest.RandIntRange(10, 100))

	protocol1 := "tcp"
	portMin := "20"
	portMax := "40"
	portMin1 := "20"
	portMax2 := "40"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISLBListenerDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISNLBPortRangeListenerConfig(vpcname, subnetname, acc.ISZoneName, acc.ISCIDR, lbname, portMin, portMax, protocol1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISLBListenerExists("ibm_is_lb_listener.testacc_lb_listener", lb),
					resource.TestCheckResourceAttr(
						"ibm_is_lb.testacc_LB", "name", lbname),
					resource.TestCheckResourceAttr(
						"ibm_is_lb.testacc_LB", "type", "public"),
					resource.TestCheckResourceAttr(
						"ibm_is_lb.testacc_LB", "route_mode", "false"),
					resource.TestCheckResourceAttr(
						"ibm_is_lb_listener.testacc_lb_listener", "port", portMin),
					resource.TestCheckResourceAttr(
						"ibm_is_lb_listener.testacc_lb_listener", "port_min", portMin),
					resource.TestCheckResourceAttr(
						"ibm_is_lb_listener.testacc_lb_listener", "port_max", portMax),
					resource.TestCheckResourceAttr(
						"ibm_is_lb_listener.testacc_lb_listener", "protocol", protocol1),
				),
			},
			{
				Config: testAccCheckIBMISNLBPortRangeListenerConfig(vpcname, subnetname, acc.ISZoneName, acc.ISCIDR, lbname, portMin1, portMax2, protocol1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISLBListenerExists("ibm_is_lb_listener.testacc_lb_listener", lb),
					resource.TestCheckResourceAttr(
						"ibm_is_lb.testacc_LB", "name", lbname),
					resource.TestCheckResourceAttr(
						"ibm_is_lb.testacc_LB", "type", "public"),
					resource.TestCheckResourceAttr(
						"ibm_is_lb.testacc_LB", "route_mode", "false"),
					resource.TestCheckResourceAttr(
						"ibm_is_lb_listener.testacc_lb_listener", "port", portMin1),
					resource.TestCheckResourceAttr(
						"ibm_is_lb_listener.testacc_lb_listener", "port_min", portMin1),
					resource.TestCheckResourceAttr(
						"ibm_is_lb_listener.testacc_lb_listener", "port_max", portMax2),
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
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISLBListenerDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISLBListenerHttpsRedirectConfig(vpcname, subnetname, acc.ISZoneName, acc.ISCIDR, lbname, port1, protocol1),
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
				Config: testAccCheckIBMISLBListenerHttpsRedirectConfigUpdate(vpcname, subnetname, acc.ISZoneName, acc.ISCIDR, lbname, port1, protocol1),
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
				Config: testAccCheckIBMISLBListenerHttpsRedirectConfigRemove(vpcname, subnetname, acc.ISZoneName, acc.ISCIDR, lbname, port1, protocol1),
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

	sess, _ := acc.TestAccProvider.Meta().(conns.ClientSession).VpcV1API()
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_is_lb_listener" {
			continue
		}

		parts, err := flex.IdParts(rs.Primary.ID)
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

		parts, err := flex.IdParts(rs.Primary.ID)
		if err != nil {
			return err
		}

		lbID := parts[0]
		lbListenerID := parts[1]

		sess, _ := acc.TestAccProvider.Meta().(conns.ClientSession).VpcV1API()
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
func testAccCheckIBMISLBUdpListenerConfig(vpcname, subnetname, zone, cidr, lbname, port, protocol string) string {
	return fmt.Sprintf(`
	resource "ibm_is_vpc" "testacc_vpc" {
		name = "%s"
	}

	resource "ibm_is_subnet" "testacc_subnet" {
		name 		= "%s"
		vpc 		= "${ibm_is_vpc.testacc_vpc.id}"
		zone 		= "%s"
		ipv4_cidr_block = "%s"
	}
	resource "ibm_is_lb" "testacc_LB" {
		name 	= "%s"
		subnets = ["${ibm_is_subnet.testacc_subnet.id}"]
		profile = "network-fixed"
		type 	= "public"
	}
	resource "ibm_is_lb_listener" "testacc_lb_listener" {
		lb 			= "${ibm_is_lb.testacc_LB.id}"
		port 		= %s
		protocol 	= "%s"
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
	  }`, vpcname, subnetname, zone, cidr, lbname, acc.LbListerenerCertificateInstance)

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
	  }`, vpcname, subnetname, zone, cidr, lbname, acc.LbListerenerCertificateInstance)

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
	  }`, vpcname, subnetname, zone, cidr, lbname, acc.LbListerenerCertificateInstance)

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
func testAccCheckIBMISNLBPortRangeListenerConfig(vpcname, subnetname, zone, cidr, lbname, portMin, portMax, protocol string) string {
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
		type 			= "public"
	}
	resource "ibm_is_lb_listener" "testacc_lb_listener" {
		lb 			= "${ibm_is_lb.testacc_LB.id}"
		port_min 	= %s
		port_max 	= %s
		protocol 	= "%s"
}`, vpcname, subnetname, zone, cidr, lbname, portMin, portMax, protocol)

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
