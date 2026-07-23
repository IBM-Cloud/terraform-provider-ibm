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

func TestAccIBMIsLbPoolDataSourceBasic(t *testing.T) {
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
				Config: testAccCheckIBMIsLbPoolDataSourceConfigBasic(vpcname, subnetname, acc.ISZoneName, acc.ISCIDR, name, poolName, alg1, protocol1, delay1, retries1, timeout1, healthType1),
				Check: resource.ComposeTestCheckFunc(
					// resource.TestCheckResourceAttrSet("data.ibm_is_lb_pool.is_lb_pool", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_lb_pool.is_lb_pool", "lb"),
					resource.TestCheckResourceAttrSet("data.ibm_is_lb_pool.is_lb_pool", "identifier"),
					resource.TestCheckResourceAttrSet("data.ibm_is_lb_pool.is_lb_pool", "algorithm"),
					resource.TestCheckResourceAttrSet("data.ibm_is_lb_pool.is_lb_pool", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_is_lb_pool.is_lb_pool", "health_monitor.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_lb_pool.is_lb_pool", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_lb_pool.is_lb_pool", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_is_lb_pool.is_lb_pool", "protocol"),
					resource.TestCheckResourceAttrSet("data.ibm_is_lb_pool.is_lb_pool", "provisioning_status"),
					resource.TestCheckResourceAttrSet("data.ibm_is_lb_pool.is_lb_pool", "proxy_protocol"),
				),
			},
		},
	})
}

func TestAccIBMIsLbPoolDataSource_mTLS(t *testing.T) {
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
				Config: testAccCheckIBMIsLbPoolDataSourceConfigmTLS(vpcname, subnetname, acc.ISZoneName, acc.ISCIDR, name, poolName, alg, protocol, delay, retries, timeout, healthType, clientCertCRN, serverCACRN),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_lb_pool.is_lb_pool_mtls", "lb"),
					resource.TestCheckResourceAttrSet("data.ibm_is_lb_pool.is_lb_pool_mtls", "identifier"),
					resource.TestCheckResourceAttr("data.ibm_is_lb_pool.is_lb_pool_mtls", "protocol", protocol),
					resource.TestCheckResourceAttrSet("data.ibm_is_lb_pool.is_lb_pool_mtls", "client_authentication.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_lb_pool.is_lb_pool_mtls", "client_authentication.0.certificate_instance.#"),
					resource.TestCheckResourceAttr("data.ibm_is_lb_pool.is_lb_pool_mtls", "client_authentication.0.certificate_instance.0.crn", clientCertCRN),
					resource.TestCheckResourceAttrSet("data.ibm_is_lb_pool.is_lb_pool_mtls", "server_authentication.#"),
					resource.TestCheckResourceAttr("data.ibm_is_lb_pool.is_lb_pool_mtls", "server_authentication.0.verify_certificate", "true"),
					resource.TestCheckResourceAttrSet("data.ibm_is_lb_pool.is_lb_pool_mtls", "server_authentication.0.certificate_authority.#"),
					resource.TestCheckResourceAttr("data.ibm_is_lb_pool.is_lb_pool_mtls", "server_authentication.0.certificate_authority.0.crn", serverCACRN),
				),
			},
		},
	})
}

func testAccCheckIBMIsLbPoolDataSourceConfigBasic(vpcname, subnetname, zone, cidr, name, poolName, algorithm, protocol, delay, retries, timeout, healthType string) string {
	return testAccCheckIBMISLBPoolConfig(vpcname, subnetname, zone, cidr, name, poolName, algorithm, protocol, delay, retries, timeout, healthType) + fmt.Sprintf(`
        data "ibm_is_lb_pool" "is_lb_pool" {
            lb = "${ibm_is_lb.testacc_LB.id}"
            identifier = "${element(split("/",ibm_is_lb_pool.testacc_lb_pool.id),1)}"
        }
    `)
}

func testAccCheckIBMIsLbPoolDataSourceConfigmTLS(vpcname, subnetname, zone, cidr, name, poolName, algorithm, protocol, delay, retries, timeout, healthType, clientCertCRN, serverCACRN string) string {
	return testAccCheckIBMISLBPoolmTLSConfig(vpcname, subnetname, zone, cidr, name, poolName, algorithm, protocol, delay, retries, timeout, healthType, clientCertCRN, serverCACRN) + fmt.Sprintf(`

	data "ibm_is_lb_pool" "is_lb_pool_mtls" {
		lb = ibm_is_lb.testacc_LB.id
		identifier = element(split("/", ibm_is_lb_pool.testacc_lb_pool_mtls.id), 1)
	}
	`)
}

func TestAccIBMIsLbPoolDataSourceHealthMonitor(t *testing.T) {
	vpcname := fmt.Sprintf("tflbp-vpc-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tflbpc-name-%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tfcreate%d", acctest.RandIntRange(10, 100))
	poolName := fmt.Sprintf("tflbpoolc%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIsLbPoolDataSourceHealthMonitorConfig(vpcname, subnetname, acc.ISZoneName, acc.ISCIDR, name, poolName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_lb_pool.is_lb_pool", "lb"),
					resource.TestCheckResourceAttrSet("data.ibm_is_lb_pool.is_lb_pool", "health_monitor.#"),
					resource.TestCheckResourceAttr("data.ibm_is_lb_pool.is_lb_pool", "health_monitor.0.request.#", "1"),
					resource.TestCheckResourceAttr("data.ibm_is_lb_pool.is_lb_pool", "health_monitor.0.request.0.method", "GET"),
					resource.TestCheckResourceAttr("data.ibm_is_lb_pool.is_lb_pool", "health_monitor.0.response.#", "1"),
					resource.TestCheckResourceAttr("data.ibm_is_lb_pool.is_lb_pool", "health_monitor.0.response.0.codes.#", "1"),
				),
			},
		},
	})
}

func testAccCheckIBMIsLbPoolDataSourceHealthMonitorConfig(vpcname, subnetname, zone, cidr, name, poolName string) string {
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
	data "ibm_is_lb_pool" "is_lb_pool" {
		lb         = ibm_is_lb.testacc_LB.id
		identifier = element(split("/", ibm_is_lb_pool.testacc_lb_pool.id), 1)
	}
	`, vpcname, subnetname, zone, cidr, name, poolName)
}

// TestAccIBMIsLbPoolDataSourceHealthMonitorHttps validates that the data source
// correctly reads advanced health_monitor fields (request, response) for HTTPS pools.
func TestAccIBMIsLbPoolDataSourceHealthMonitorHttps(t *testing.T) {
	vpcname := fmt.Sprintf("tflbp-vpc-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tflbpc-name-%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tfcreate%d", acctest.RandIntRange(10, 100))
	poolName := fmt.Sprintf("tflbpoolc%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIsLbPoolDataSourceHealthMonitorHttpsConfig(vpcname, subnetname, acc.ISZoneName, acc.ISCIDR, name, poolName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_lb_pool.is_lb_pool", "lb"),
					resource.TestCheckResourceAttr("data.ibm_is_lb_pool.is_lb_pool", "health_monitor.0.type", "https"),
					resource.TestCheckResourceAttr("data.ibm_is_lb_pool.is_lb_pool", "health_monitor.0.request.#", "1"),
					resource.TestCheckResourceAttr("data.ibm_is_lb_pool.is_lb_pool", "health_monitor.0.request.0.method", "POST"),
					resource.TestCheckResourceAttr("data.ibm_is_lb_pool.is_lb_pool", "health_monitor.0.request.0.headers.#", "1"),
					resource.TestCheckResourceAttr("data.ibm_is_lb_pool.is_lb_pool", "health_monitor.0.request.0.headers.0.field", "Content-Type"),
					resource.TestCheckResourceAttr("data.ibm_is_lb_pool.is_lb_pool", "health_monitor.0.response.#", "1"),
					resource.TestCheckResourceAttr("data.ibm_is_lb_pool.is_lb_pool", "health_monitor.0.response.0.codes.#", "1"),
					resource.TestCheckResourceAttr("data.ibm_is_lb_pool.is_lb_pool", "health_monitor.0.response.0.body_regex", ".*healthy.*"),
				),
			},
		},
	})
}

// TestAccIBMIsLbPoolDataSourceNoHealthMonitorRequest validates that a pool
// without request/response in health_monitor returns empty lists in the data source.
func TestAccIBMIsLbPoolDataSourceNoHealthMonitorRequest(t *testing.T) {
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
			{
				Config: testAccCheckIBMIsLbPoolDataSourceConfigBasic(vpcname, subnetname, acc.ISZoneName, acc.ISCIDR, name, poolName, alg1, protocol1, delay1, retries1, timeout1, healthType1),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_lb_pool.is_lb_pool", "lb"),
					resource.TestCheckResourceAttrSet("data.ibm_is_lb_pool.is_lb_pool", "health_monitor.#"),
					resource.TestCheckResourceAttr("data.ibm_is_lb_pool.is_lb_pool", "health_monitor.0.type", "http"),
					resource.TestCheckResourceAttr("data.ibm_is_lb_pool.is_lb_pool", "health_monitor.0.delay", "45"),
					resource.TestCheckResourceAttr("data.ibm_is_lb_pool.is_lb_pool", "health_monitor.0.max_retries", "5"),
					resource.TestCheckResourceAttr("data.ibm_is_lb_pool.is_lb_pool", "health_monitor.0.timeout", "15"),
					// No advanced request/response for a plain pool without health_monitor request
					resource.TestCheckResourceAttr("data.ibm_is_lb_pool.is_lb_pool", "health_monitor.0.request.#", "0"),
					resource.TestCheckResourceAttr("data.ibm_is_lb_pool.is_lb_pool", "health_monitor.0.response.#", "0"),
				),
			},
		},
	})
}

func testAccCheckIBMIsLbPoolDataSourceHealthMonitorHttpsConfig(vpcname, subnetname, zone, cidr, name, poolName string) string {
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
		name               = "%s"
		lb                 = ibm_is_lb.testacc_LB.id
		algorithm          = "round_robin"
		protocol           = "https"
		health_delay       = 45
		health_retries     = 5
		health_timeout     = 15
		health_type        = "https"
		health_monitor_url = "/healthz"
		health_monitor {
			request {
				method = "POST"
				headers {
					field = "Content-Type"
					value = "application/json"
				}
			}
			response {
				codes      = ["200"]
				body_regex = ".*healthy.*"
			}
		}
	}
	data "ibm_is_lb_pool" "is_lb_pool" {
		lb         = ibm_is_lb.testacc_LB.id
		identifier = element(split("/", ibm_is_lb_pool.testacc_lb_pool.id), 1)
	}
	`, vpcname, subnetname, zone, cidr, name, poolName)
}
