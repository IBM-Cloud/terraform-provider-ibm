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

func TestAccIBMIsLbListenersDataSourceBasic(t *testing.T) {
	vpcname := fmt.Sprintf("tflblis-vpc-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tflblis-subnet-%d", acctest.RandIntRange(10, 100))
	lbname := fmt.Sprintf("tflblis%d", acctest.RandIntRange(10, 100))
	protocol1 := "http"
	port1 := "8080"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsLbListenersDataSourceConfigBasic(vpcname, subnetname, acc.ISZoneName, acc.ISCIDR, lbname, port1, protocol1),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_lb_listeners.is_lb_listeners", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_lb_listeners.is_lb_listeners", "lb"),
					resource.TestCheckResourceAttrSet("data.ibm_is_lb_listeners.is_lb_listeners", "listeners.#"),
				),
			},
		},
	})
}

func testAccCheckIBMIsLbListenersDataSourceConfigBasic(vpcname, subnetname, zone, cidr, lbname, port, protocol string) string {
	return testAccCheckIBMISLBListenerConfig(vpcname, subnetname, acc.ISZoneName, acc.ISCIDR, lbname, port, protocol) + fmt.Sprintf(`

	data "ibm_is_lb_listeners" "is_lb_listeners" {
		lb = "${ibm_is_lb.testacc_LB.id}"
	}
	`)
}
