// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMIsLbPoolsDataSourceBasic(t *testing.T) {
	vpcname := fmt.Sprintf("tflbp-vpc-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tflbpc-name-%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tfcreate%d", acctest.RandIntRange(10, 100))
	poolName := fmt.Sprintf("tflbpoolc%d", acctest.RandIntRange(10, 100))
	alg1 := "round_robin"
	protocol1 := "http"
	delay1 := "45"
	retries1 := "5"
	timeout1 := "15"
	healthType1 := "http"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsLbPoolsDataSourceConfigBasic(vpcname, subnetname, acc.ISZoneName, acc.ISCIDR, name, poolName, alg1, protocol1, delay1, retries1, timeout1, healthType1),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_lb_pools.is_lb_pools", "lb"),
					resource.TestCheckResourceAttrSet("data.ibm_is_lb_pools.is_lb_pools", "pools.#"),
				),
			},
		},
	})
}

func TestAccIBMIsLbPoolsDataSource_mTLS(t *testing.T) {
	vpcname := fmt.Sprintf("tflbp-vpc-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tflbpc-name-%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tfcreate%d", acctest.RandIntRange(10, 100))
	poolName := fmt.Sprintf("tflbpoolc%d", acctest.RandIntRange(10, 100))
	alg := "round_robin"
	protocol := "https"
	delay := "45"
	retries := "5"
	timeout := "15"
	healthType := "https"

	// Example CRNs - replace with actual values from your test environment
	clientCertCRN := "crn:v1:bluemix:public:secrets-manager:us-south:a/7f75c7b025e54bc5635f754b2f888665:152af435-37ac-4b3e-83c3-828805bfc8e0:secret:1e5b9794-f576-de33-5e41-4a8c29d00132"
	serverCACRN := "crn:v1:bluemix:public:secrets-manager:us-south:a/7f75c7b025e54bc5635f754b2f888665:152af435-37ac-4b3e-83c3-828805bfc8e0:secret:4a1bc2d6-ccd3-ad25-e6b7-8a0c522038f6"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIsLbPoolsDataSourceConfigmTLS(vpcname, subnetname, acc.ISZoneName, acc.ISCIDR, name, poolName, alg, protocol, delay, retries, timeout, healthType, clientCertCRN, serverCACRN),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_lb_pools.is_lb_pools", "lb"),
					resource.TestCheckResourceAttrSet("data.ibm_is_lb_pools.is_lb_pools", "pools.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_lb_pools.is_lb_pools", "pools.0.client_authentication.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_lb_pools.is_lb_pools", "pools.0.client_authentication.0.certificate_instance.#"),
					resource.TestCheckResourceAttr("data.ibm_is_lb_pools.is_lb_pools", "pools.0.client_authentication.0.certificate_instance.0.crn", clientCertCRN),
					resource.TestCheckResourceAttrSet("data.ibm_is_lb_pools.is_lb_pools", "pools.0.server_authentication.#"),
					resource.TestCheckResourceAttr("data.ibm_is_lb_pools.is_lb_pools", "pools.0.server_authentication.0.verify_certificate", "true"),
					resource.TestCheckResourceAttrSet("data.ibm_is_lb_pools.is_lb_pools", "pools.0.server_authentication.0.certificate_authority.#"),
					resource.TestCheckResourceAttr("data.ibm_is_lb_pools.is_lb_pools", "pools.0.server_authentication.0.certificate_authority.0.crn", serverCACRN),
				),
			},
		},
	})
}

func testAccCheckIBMIsLbPoolsDataSourceConfigBasic(vpcname, subnetname, zone, cidr, name, poolName, algorithm, protocol, delay, retries, timeout, healthType string) string {
	return testAccCheckIBMISLBPoolConfig(vpcname, subnetname, zone, cidr, name, poolName, algorithm, protocol, delay, retries, timeout, healthType) + fmt.Sprintf(`
        data "ibm_is_lb_pools" "is_lb_pools" {
            lb = "${ibm_is_lb.testacc_LB.id}"
        }
    `)
}

func testAccCheckIBMIsLbPoolsDataSourceConfigmTLS(vpcname, subnetname, zone, cidr, name, poolName, algorithm, protocol, delay, retries, timeout, healthType, clientCertCRN, serverCACRN string) string {
	return testAccCheckIBMISLBPoolmTLSConfig(vpcname, subnetname, zone, cidr, name, poolName, algorithm, protocol, delay, retries, timeout, healthType, clientCertCRN, serverCACRN) + fmt.Sprintf(`

	data "ibm_is_lb_pools" "is_lb_pools" {
		lb = ibm_is_lb.testacc_LB.id
		depends_on = [ibm_is_lb_pool.testacc_lb_pool_mtls]
	}
	`)
}

// http bundle tests

// TestAccIBMIsLbPoolsDataSourceHealthMonitor validates that the pools list data
// source surfaces health_monitor.request and health_monitor.response when set.
func TestAccIBMIsLbPoolsDataSourceHealthMonitor(t *testing.T) {
	vpcname := fmt.Sprintf("tflbp-vpc-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tflbpc-name-%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tfcreate%d", acctest.RandIntRange(10, 100))
	poolName := fmt.Sprintf("tflbpoolc%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIsLbPoolsDataSourceHealthMonitorConfig(vpcname, subnetname, acc.ISZoneName, acc.ISCIDR, name, poolName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_lb_pools.is_lb_pools", "lb"),
					resource.TestCheckResourceAttrSet("data.ibm_is_lb_pools.is_lb_pools", "pools.#"),
					resource.TestCheckResourceAttr("data.ibm_is_lb_pools.is_lb_pools", "pools.0.health_monitor.0.request.#", "1"),
					resource.TestCheckResourceAttr("data.ibm_is_lb_pools.is_lb_pools", "pools.0.health_monitor.0.request.0.method", "GET"),
					resource.TestCheckResourceAttr("data.ibm_is_lb_pools.is_lb_pools", "pools.0.health_monitor.0.response.#", "1"),
					resource.TestCheckResourceAttr("data.ibm_is_lb_pools.is_lb_pools", "pools.0.health_monitor.0.response.0.codes.#", "1"),
				),
			},
		},
	})
}

func testAccCheckIBMIsLbPoolsDataSourceHealthMonitorConfig(vpcname, subnetname, zone, cidr, name, poolName string) string {
	return fmt.Sprintf(`
	resource "ibm_is_vpc" "testacc_vpc" {
		name = "%s"
	}
	resource "ibm_is_subnet" "testacc_subnet" {
		name            = "%s"
		vpc             = ibm_is_vpc.testacc_vpc.id
		zone            = "%s"
		ipv4_cidr_block = "%s"
	}
	resource "ibm_is_lb" "testacc_LB" {
		name    = "%s"
		subnets = [ibm_is_subnet.testacc_subnet.id]
	}
	resource "ibm_is_lb_pool" "testacc_lb_pool" {
		name           = "%s"
		lb             = ibm_is_lb.testacc_LB.id
		algorithm      = "round_robin"
		protocol       = "http"
		health_delay   = 45
		health_retries = 5
		health_timeout = 15
		health_type    = "http"
		health_monitor {
			request {
				method = "GET"
			}
			response {
				codes = ["200"]
			}
		}
	}
	data "ibm_is_lb_pools" "is_lb_pools" {
		lb = ibm_is_lb.testacc_LB.id
		depends_on = [ibm_is_lb_pool.testacc_lb_pool]
	}
	`, vpcname, subnetname, zone, cidr, name, poolName)
}
