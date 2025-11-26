// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
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

func testAccCheckIBMIsLbPoolDataSourceConfigBasic(vpcname, subnetname, zone, cidr, name, poolName, algorithm, protocol, delay, retries, timeout, healthType string) string {
	return testAccCheckIBMISLBPoolConfig(vpcname, subnetname, zone, cidr, name, poolName, algorithm, protocol, delay, retries, timeout, healthType) + fmt.Sprintf(`
        data "ibm_is_lb_pool" "is_lb_pool" {
            lb = "${ibm_is_lb.testacc_LB.id}"
            identifier = "${element(split("/",ibm_is_lb_pool.testacc_lb_pool.id),1)}"
        }
    `)
}
