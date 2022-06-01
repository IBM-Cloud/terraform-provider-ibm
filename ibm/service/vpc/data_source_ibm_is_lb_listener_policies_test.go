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

func TestAccIBMIsLbListenerPoliciesDataSourceBasic(t *testing.T) {
	vpcname := fmt.Sprintf("tflblisuat-vpc-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tflblisuat-subnet-%d", acctest.RandIntRange(10, 100))
	lbname := fmt.Sprintf("tflblisuat%d", acctest.RandIntRange(10, 100))
	lblistenerpolicyname1 := fmt.Sprintf("tflblisuat-listener-policy-%d", acctest.RandIntRange(10, 100))
	priority1 := "1"
	protocol := "http"
	port := "8080"
	action := "forward"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsLbListenerPoliciesDataSourceConfigBasic(vpcname, subnetname, acc.ISZoneName, acc.ISCIDR, lbname, port, protocol, lblistenerpolicyname1, action, priority1),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_lb_listener_policies.is_lb_listener_policies", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_lb_listener_policies.is_lb_listener_policies", "lb"),
					resource.TestCheckResourceAttrSet("data.ibm_is_lb_listener_policies.is_lb_listener_policies", "listener"),
					resource.TestCheckResourceAttrSet("data.ibm_is_lb_listener_policies.is_lb_listener_policies", "policies.#"),
				),
			},
		},
	})
}

func testAccCheckIBMIsLbListenerPoliciesDataSourceConfigBasic(vpcname, subnetname, zone, cidr, lbname, port, protocol, lblistenerpolicyname, action, priority string) string {
	return testAccCheckIBMISLBListenerPolicyConfig(vpcname, subnetname, acc.ISZoneName, acc.ISCIDR, lbname, port, protocol, lblistenerpolicyname, action, priority) + fmt.Sprintf(`
	data "ibm_is_lb_listener_policies" "is_lb_listener_policies" {
		lb = "${ibm_is_lb.testacc_LB.id}"
		listener = "${ibm_is_lb_listener.testacc_lb_listener.listener_id}"
	}
	`)
}
