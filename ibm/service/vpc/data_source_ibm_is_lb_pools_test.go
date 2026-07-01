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
	clientCertCRN := "crn:v1:staging:public:secrets-manager:eu-gb:a/2d1bace7b46e4815a81e52c6ffeba5cf:2ca77a00-d2c6-41a2-93e4-6bfa23400b17:secret:7b8bea2d-124d-1264-98c9-678404ac947e"
	serverCACRN := "crn:v1:staging:public:secrets-manager:eu-gb:a/2d1bace7b46e4815a81e52c6ffeba5cf:2ca77a00-d2c6-41a2-93e4-6bfa23400b17:secret:6133d2b7-44b0-f6d1-87ff-67ae4f8f8a05"

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
