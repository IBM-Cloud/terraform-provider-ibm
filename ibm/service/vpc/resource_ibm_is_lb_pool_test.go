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

func TestAccIBMISLBPool_basic(t *testing.T) {
	var lb string
	vpcname := fmt.Sprintf("tflbp-vpc-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tflbpc-name-%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tfcreate%d", acctest.RandIntRange(10, 100))
	poolName := fmt.Sprintf("tflbpoolc%d", acctest.RandIntRange(10, 100))
	poolName1 := fmt.Sprintf("tflbpoolu%d", acctest.RandIntRange(10, 100))
	alg1 := "round_robin"
	protocol1 := "http"
	proxyProtocol1 := "disabled"
	delay1 := "45"
	retries1 := "5"
	timeout1 := "15"
	healthType1 := "http"

	alg2 := "least_connections"
	protocol2 := "tcp"
	proxyProtocol2 := "v2"
	delay2 := "60"
	retries2 := "3"
	timeout2 := "30"
	healthType2 := "tcp"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISLBPoolDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISLBPoolConfig(vpcname, subnetname, acc.ISZoneName, acc.ISCIDR, name, poolName, alg1, protocol1, delay1, retries1, timeout1, healthType1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISLBPoolExists("ibm_is_lb_pool.testacc_lb_pool", lb),
					resource.TestCheckResourceAttr(
						"ibm_is_lb.testacc_LB", "name", name),
					resource.TestCheckResourceAttr(
						"ibm_is_lb_pool.testacc_lb_pool", "name", poolName),
					resource.TestCheckResourceAttr(
						"ibm_is_lb_pool.testacc_lb_pool", "algorithm", alg1),
					resource.TestCheckResourceAttr(
						"ibm_is_lb_pool.testacc_lb_pool", "protocol", protocol1),
					resource.TestCheckResourceAttr(
						"ibm_is_lb_pool.testacc_lb_pool", "proxy_protocol", proxyProtocol1),
					resource.TestCheckResourceAttr(
						"ibm_is_lb_pool.testacc_lb_pool", "health_delay", delay1),
					resource.TestCheckResourceAttr(
						"ibm_is_lb_pool.testacc_lb_pool", "health_retries", retries1),
					resource.TestCheckResourceAttr(
						"ibm_is_lb_pool.testacc_lb_pool", "health_timeout", timeout1),
					resource.TestCheckResourceAttr(
						"ibm_is_lb_pool.testacc_lb_pool", "health_type", healthType1),
				),
			},

			{
				Config: testAccCheckIBMISLBPoolConfig(vpcname, subnetname, acc.ISZoneName, acc.ISCIDR, name, poolName1, alg2, protocol2, delay2, retries2, timeout2, healthType2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISLBPoolExists("ibm_is_lb_pool.testacc_lb_pool", lb),
					resource.TestCheckResourceAttr(
						"ibm_is_lb_pool.testacc_lb_pool", "name", poolName1),
					resource.TestCheckResourceAttr(
						"ibm_is_lb_pool.testacc_lb_pool", "algorithm", alg2),
					resource.TestCheckResourceAttr(
						"ibm_is_lb_pool.testacc_lb_pool", "protocol", protocol2),
					resource.TestCheckResourceAttr(
						"ibm_is_lb_pool.testacc_lb_pool", "proxy_protocol", proxyProtocol1),
					resource.TestCheckResourceAttr(
						"ibm_is_lb_pool.testacc_lb_pool", "health_delay", delay2),
					resource.TestCheckResourceAttr(
						"ibm_is_lb_pool.testacc_lb_pool", "health_retries", retries2),
					resource.TestCheckResourceAttr(
						"ibm_is_lb_pool.testacc_lb_pool", "health_timeout", timeout2),
					resource.TestCheckResourceAttr(
						"ibm_is_lb_pool.testacc_lb_pool", "health_type", healthType2),
				),
			},
			{
				Config: testAccCheckIBMISLBPoolConfigWithProxy(vpcname, subnetname, acc.ISZoneName, acc.ISCIDR, name, poolName, alg1, protocol1, proxyProtocol2, delay1, retries1, timeout1, healthType1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISLBPoolExists("ibm_is_lb_pool.testacc_lb_pool", lb),
					resource.TestCheckResourceAttr(
						"ibm_is_lb.testacc_LB", "name", name),
					resource.TestCheckResourceAttr(
						"ibm_is_lb_pool.testacc_lb_pool", "name", poolName),
					resource.TestCheckResourceAttr(
						"ibm_is_lb_pool.testacc_lb_pool", "algorithm", alg1),
					resource.TestCheckResourceAttr(
						"ibm_is_lb_pool.testacc_lb_pool", "protocol", protocol1),
					resource.TestCheckResourceAttr(
						"ibm_is_lb_pool.testacc_lb_pool", "proxy_protocol", proxyProtocol2),
					resource.TestCheckResourceAttr(
						"ibm_is_lb_pool.testacc_lb_pool", "health_delay", delay1),
					resource.TestCheckResourceAttr(
						"ibm_is_lb_pool.testacc_lb_pool", "health_retries", retries1),
					resource.TestCheckResourceAttr(
						"ibm_is_lb_pool.testacc_lb_pool", "health_timeout", timeout1),
					resource.TestCheckResourceAttr(
						"ibm_is_lb_pool.testacc_lb_pool", "health_type", healthType1),
				),
			},
		},
	})
}
func TestAccIBMISLBPool_basic_udp(t *testing.T) {
	var lb string
	vpcname := fmt.Sprintf("tflbp-vpc-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tflbpc-name-%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tfcreate%d", acctest.RandIntRange(10, 100))
	poolName := fmt.Sprintf("tflbpoolc%d", acctest.RandIntRange(10, 100))
	alg1 := "round_robin"
	protocol1 := "udp"
	delay1 := "5"
	retries1 := "2"
	timeout1 := "2"
	healthType1 := "http"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISLBPoolDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISLBPoolUdpConfig(vpcname, subnetname, acc.ISZoneName, acc.ISCIDR, name, poolName, alg1, protocol1, delay1, retries1, timeout1, healthType1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISLBPoolExists("ibm_is_lb_pool.testacc_lb_pool", lb),
					resource.TestCheckResourceAttr(
						"ibm_is_lb.testacc_LB", "name", name),
					resource.TestCheckResourceAttr(
						"ibm_is_lb_pool.testacc_lb_pool", "name", poolName),
					resource.TestCheckResourceAttr(
						"ibm_is_lb_pool.testacc_lb_pool", "algorithm", alg1),
					resource.TestCheckResourceAttr(
						"ibm_is_lb_pool.testacc_lb_pool", "protocol", protocol1),
					resource.TestCheckResourceAttr(
						"ibm_is_lb_pool.testacc_lb_pool", "health_delay", delay1),
					resource.TestCheckResourceAttr(
						"ibm_is_lb_pool.testacc_lb_pool", "health_retries", retries1),
					resource.TestCheckResourceAttr(
						"ibm_is_lb_pool.testacc_lb_pool", "health_timeout", timeout1),
					resource.TestCheckResourceAttr(
						"ibm_is_lb_pool.testacc_lb_pool", "health_type", healthType1),
				),
			},
		},
	})
}

func TestAccIBMISLBPool_port(t *testing.T) {
	var lb string
	vpcname := fmt.Sprintf("tflbp-vpc-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tflbp-subnet-%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tfcreate%d", acctest.RandIntRange(10, 100))
	poolName := fmt.Sprintf("tflbpoolc%d", acctest.RandIntRange(10, 100))
	alg1 := "round_robin"
	protocol1 := "http"
	proxyProtocol1 := "disabled"
	delay1 := "45"
	retries1 := "5"
	timeout1 := "15"
	healthType1 := "http"
	port := "2554"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISLBPoolDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISLBPoolPortConfig(vpcname, subnetname, acc.ISZoneName, acc.ISCIDR, name, poolName, alg1, protocol1, delay1, retries1, timeout1, healthType1, port),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISLBPoolExists("ibm_is_lb_pool.testacc_lb_pool", lb),
					resource.TestCheckResourceAttr(
						"ibm_is_lb.testacc_LB", "name", name),
					resource.TestCheckResourceAttr(
						"ibm_is_lb_pool.testacc_lb_pool", "name", poolName),
					resource.TestCheckResourceAttr(
						"ibm_is_lb_pool.testacc_lb_pool", "algorithm", alg1),
					resource.TestCheckResourceAttr(
						"ibm_is_lb_pool.testacc_lb_pool", "protocol", protocol1),
					resource.TestCheckResourceAttr(
						"ibm_is_lb_pool.testacc_lb_pool", "proxy_protocol", proxyProtocol1),
					resource.TestCheckResourceAttr(
						"ibm_is_lb_pool.testacc_lb_pool", "health_delay", delay1),
					resource.TestCheckResourceAttr(
						"ibm_is_lb_pool.testacc_lb_pool", "health_retries", retries1),
					resource.TestCheckResourceAttr(
						"ibm_is_lb_pool.testacc_lb_pool", "health_timeout", timeout1),
					resource.TestCheckResourceAttr(
						"ibm_is_lb_pool.testacc_lb_pool", "health_type", healthType1),
				),
			},
		},
	})
}

func TestAccIBMISLBPool_SessionPersistence(t *testing.T) {
	var lb string
	vpcname := fmt.Sprintf("tflbp-vpc-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tflbp-subnet-%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tfcreate%d", acctest.RandIntRange(10, 100))
	poolName := fmt.Sprintf("tflbpoolc%d", acctest.RandIntRange(10, 100))
	alg1 := "round_robin"
	protocol1 := "http"
	proxyProtocol1 := "disabled"
	delay1 := "45"
	retries1 := "5"
	timeout1 := "15"
	healthType1 := "http"
	port := "2554"
	session_persistence_appcookie_type := "app_cookie"
	session_persistence_httpcookie_type := "http_cookie"
	app_cookie_name := "testacc_cookie"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISLBPoolDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISLBPoolSessionPersistenceConfig(vpcname, subnetname, acc.ISZoneName, acc.ISCIDR, name, poolName, alg1, protocol1, delay1, retries1, timeout1, healthType1, port, session_persistence_appcookie_type, app_cookie_name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISLBPoolExists("ibm_is_lb_pool.testacc_lb_pool", lb),
					resource.TestCheckResourceAttr(
						"ibm_is_lb.testacc_LB", "name", name),
					resource.TestCheckResourceAttr(
						"ibm_is_lb_pool.testacc_lb_pool", "name", poolName),
					resource.TestCheckResourceAttr(
						"ibm_is_lb_pool.testacc_lb_pool", "algorithm", alg1),
					resource.TestCheckResourceAttr(
						"ibm_is_lb_pool.testacc_lb_pool", "protocol", protocol1),
					resource.TestCheckResourceAttr(
						"ibm_is_lb_pool.testacc_lb_pool", "proxy_protocol", proxyProtocol1),
					resource.TestCheckResourceAttr(
						"ibm_is_lb_pool.testacc_lb_pool", "health_delay", delay1),
					resource.TestCheckResourceAttr(
						"ibm_is_lb_pool.testacc_lb_pool", "health_retries", retries1),
					resource.TestCheckResourceAttr(
						"ibm_is_lb_pool.testacc_lb_pool", "health_timeout", timeout1),
					resource.TestCheckResourceAttr(
						"ibm_is_lb_pool.testacc_lb_pool", "health_type", healthType1),
					resource.TestCheckResourceAttr(
						"ibm_is_lb_pool.testacc_lb_pool", "session_persistence_type", session_persistence_appcookie_type),
					resource.TestCheckResourceAttr(
						"ibm_is_lb_pool.testacc_lb_pool", "session_persistence_app_cookie_name", app_cookie_name),
					resource.TestCheckResourceAttr(
						"ibm_is_lb_pool.testacc_lb_pool", "session_persistence_http_cookie_name", ""),
				),
			},
			{
				Config: testAccCheckIBMISLBPoolSessionPersistenceConfigUpdate(vpcname, subnetname, acc.ISZoneName, acc.ISCIDR, name, poolName, alg1, protocol1, delay1, retries1, timeout1, healthType1, port, session_persistence_httpcookie_type),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISLBPoolExists("ibm_is_lb_pool.testacc_lb_pool", lb),
					resource.TestCheckResourceAttr(
						"ibm_is_lb.testacc_LB", "name", name),
					resource.TestCheckResourceAttr(
						"ibm_is_lb_pool.testacc_lb_pool", "name", poolName),
					resource.TestCheckResourceAttr(
						"ibm_is_lb_pool.testacc_lb_pool", "algorithm", alg1),
					resource.TestCheckResourceAttr(
						"ibm_is_lb_pool.testacc_lb_pool", "protocol", protocol1),
					resource.TestCheckResourceAttr(
						"ibm_is_lb_pool.testacc_lb_pool", "proxy_protocol", proxyProtocol1),
					resource.TestCheckResourceAttr(
						"ibm_is_lb_pool.testacc_lb_pool", "health_delay", delay1),
					resource.TestCheckResourceAttr(
						"ibm_is_lb_pool.testacc_lb_pool", "health_retries", retries1),
					resource.TestCheckResourceAttr(
						"ibm_is_lb_pool.testacc_lb_pool", "health_timeout", timeout1),
					resource.TestCheckResourceAttr(
						"ibm_is_lb_pool.testacc_lb_pool", "health_type", healthType1),
					resource.TestCheckResourceAttr(
						"ibm_is_lb_pool.testacc_lb_pool", "session_persistence_type", session_persistence_httpcookie_type),
					resource.TestCheckResourceAttr(
						"ibm_is_lb_pool.testacc_lb_pool", "session_persistence_app_cookie_name", ""),
					resource.TestCheckResourceAttrSet(
						"ibm_is_lb_pool.testacc_lb_pool", "session_persistence_http_cookie_name"),
				),
			},
			{
				Config: testAccCheckIBMISLBPoolSessionPersistenceConfigRemove(vpcname, subnetname, acc.ISZoneName, acc.ISCIDR, name, poolName, alg1, protocol1, delay1, retries1, timeout1, healthType1, port),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISLBPoolExists("ibm_is_lb_pool.testacc_lb_pool", lb),
					resource.TestCheckResourceAttr(
						"ibm_is_lb.testacc_LB", "name", name),
					resource.TestCheckResourceAttr(
						"ibm_is_lb_pool.testacc_lb_pool", "name", poolName),
					resource.TestCheckResourceAttr(
						"ibm_is_lb_pool.testacc_lb_pool", "algorithm", alg1),
					resource.TestCheckResourceAttr(
						"ibm_is_lb_pool.testacc_lb_pool", "protocol", protocol1),
					resource.TestCheckResourceAttr(
						"ibm_is_lb_pool.testacc_lb_pool", "proxy_protocol", proxyProtocol1),
					resource.TestCheckResourceAttr(
						"ibm_is_lb_pool.testacc_lb_pool", "health_delay", delay1),
					resource.TestCheckResourceAttr(
						"ibm_is_lb_pool.testacc_lb_pool", "health_retries", retries1),
					resource.TestCheckResourceAttr(
						"ibm_is_lb_pool.testacc_lb_pool", "health_timeout", timeout1),
					resource.TestCheckResourceAttr(
						"ibm_is_lb_pool.testacc_lb_pool", "health_type", healthType1),
					resource.TestCheckResourceAttr(
						"ibm_is_lb_pool.testacc_lb_pool", "session_persistence_type", ""),
					resource.TestCheckResourceAttr(
						"ibm_is_lb_pool.testacc_lb_pool", "session_persistence_app_cookie_name", ""),
					resource.TestCheckResourceAttr(
						"ibm_is_lb_pool.testacc_lb_pool", "session_persistence_http_cookie_name", ""),
				),
			},
		},
	})
}

func testAccCheckIBMISLBPoolDestroy(s *terraform.State) error {

	sess, _ := acc.TestAccProvider.Meta().(conns.ClientSession).VpcV1API()
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_is_lb_pool" {
			continue
		}

		parts, err := flex.IdParts(rs.Primary.ID)
		if err != nil {
			return err
		}

		lbID := parts[0]
		lbPoolID := parts[1]
		getlbpptions := &vpcv1.GetLoadBalancerPoolOptions{
			LoadBalancerID: &lbID,
			ID:             &lbPoolID,
		}
		_, _, err1 := sess.GetLoadBalancerPool(getlbpptions)
		if err1 == nil {
			return fmt.Errorf("LB Pool still exists: %s", rs.Primary.ID)
		}
	}

	return nil
}

func testAccCheckIBMISLBPoolExists(n, lbPool string) resource.TestCheckFunc {
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
		lbPoolID := parts[1]

		sess, _ := acc.TestAccProvider.Meta().(conns.ClientSession).VpcV1API()
		getlbpptions := &vpcv1.GetLoadBalancerPoolOptions{
			LoadBalancerID: &lbID,
			ID:             &lbPoolID,
		}
		foundLBPool, _, err := sess.GetLoadBalancerPool(getlbpptions)
		if err != nil {
			return err
		}
		lbPool = *foundLBPool.ID

		return nil
	}
}

func testAccCheckIBMISLBPoolConfig(vpcname, subnetname, zone, cidr, name, poolName, algorithm, protocol, delay, retries, timeout, healthType string) string {
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
	resource "ibm_is_lb_pool" "testacc_lb_pool" {
		name = "%s"
		lb = "${ibm_is_lb.testacc_LB.id}"
		algorithm = "%s"
		protocol = "%s"
		health_delay= %s
		health_retries = %s
		health_timeout = %s
		health_type = "%s"
}`, vpcname, subnetname, zone, cidr, name, poolName, algorithm, protocol, delay, retries, timeout, healthType)

}
func testAccCheckIBMISLBPoolUdpConfig(vpcname, subnetname, zone, cidr, name, poolName, algorithm, protocol, delay, retries, timeout, healthType string) string {
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
		name 	= "%s"
		subnets = ["${ibm_is_subnet.testacc_subnet.id}"]
		profile = "network-fixed"
		type 	= "public"
	}
	resource "ibm_is_lb_pool" "testacc_lb_pool" {
		name 				= "%s"
		lb 					= "${ibm_is_lb.testacc_LB.id}"
		algorithm          	= "%s"
		protocol           	= "%s"
		health_delay       	= %s
		health_retries     	= %s
		health_timeout     	= %s
		health_type        	= "%s"
		health_monitor_url 	= "/"
}`, vpcname, subnetname, zone, cidr, name, poolName, algorithm, protocol, delay, retries, timeout, healthType)

}

func testAccCheckIBMISLBPoolPortConfig(vpcname, subnetname, zone, cidr, name, poolName, algorithm, protocol, delay, retries, timeout, healthType, port string) string {
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
	resource "ibm_is_lb_pool" "testacc_lb_pool" {
		name = "%s"
		lb = "${ibm_is_lb.testacc_LB.id}"
		algorithm = "%s"
		protocol = "%s"
		health_delay= %s
		health_retries = %s
		health_timeout = %s
		health_type = "%s"
		health_monitor_port = %s
}`, vpcname, subnetname, zone, cidr, name, poolName, algorithm, protocol, delay, retries, timeout, healthType, port)

}

func testAccCheckIBMISLBPoolSessionPersistenceConfig(vpcname, subnetname, zone, cidr, name, poolName, algorithm, protocol, delay, retries, timeout, healthType, port, session_persistence_type, app_cookie_name string) string {
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
	resource "ibm_is_lb_pool" "testacc_lb_pool" {
		name = "%s"
		lb = "${ibm_is_lb.testacc_LB.id}"
		algorithm = "%s"
		protocol = "%s"
		health_delay= %s
		health_retries = %s
		health_timeout = %s
		health_type = "%s"
		health_monitor_port = %s
		session_persistence_type = "%s"
		session_persistence_app_cookie_name = "%s"
}`, vpcname, subnetname, zone, cidr, name, poolName, algorithm, protocol, delay, retries, timeout, healthType, port, session_persistence_type, app_cookie_name)

}

func testAccCheckIBMISLBPoolSessionPersistenceConfigUpdate(vpcname, subnetname, zone, cidr, name, poolName, algorithm, protocol, delay, retries, timeout, healthType, port, session_persistence_type string) string {
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
	resource "ibm_is_lb_pool" "testacc_lb_pool" {
		name = "%s"
		lb = "${ibm_is_lb.testacc_LB.id}"
		algorithm = "%s"
		protocol = "%s"
		health_delay= %s
		health_retries = %s
		health_timeout = %s
		health_type = "%s"
		health_monitor_port = %s
		session_persistence_type = "%s"
}`, vpcname, subnetname, zone, cidr, name, poolName, algorithm, protocol, delay, retries, timeout, healthType, port, session_persistence_type)

}

func testAccCheckIBMISLBPoolSessionPersistenceConfigRemove(vpcname, subnetname, zone, cidr, name, poolName, algorithm, protocol, delay, retries, timeout, healthType, port string) string {
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
	resource "ibm_is_lb_pool" "testacc_lb_pool" {
		name = "%s"
		lb = "${ibm_is_lb.testacc_LB.id}"
		algorithm = "%s"
		protocol = "%s"
		health_delay= %s
		health_retries = %s
		health_timeout = %s
		health_type = "%s"
		health_monitor_port = %s
}`, vpcname, subnetname, zone, cidr, name, poolName, algorithm, protocol, delay, retries, timeout, healthType, port)

}

func testAccCheckIBMISLBPoolConfigWithProxy(vpcname, subnetname, zone, cidr, name, poolName, algorithm, protocol, proxyProtocol, delay, retries, timeout, healthType string) string {
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
	resource "ibm_is_lb_pool" "testacc_lb_pool" {
		name = "%s"
		lb = "${ibm_is_lb.testacc_LB.id}"
		algorithm = "%s"
		protocol = "%s"
		proxy_protocol = "%s"
		health_delay= %s
		health_retries = %s
		health_timeout = %s
		health_type = "%s"
}`, vpcname, subnetname, zone, cidr, name, poolName, algorithm, protocol, proxyProtocol, delay, retries, timeout, healthType)

}
