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

func TestAccIBMIsLbListenerPolicyDataSourceBasic(t *testing.T) {
	vpcname := fmt.Sprintf("tflblisuat-vpc-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tflblisuat-subnet-%d", acctest.RandIntRange(10, 100))
	lbname := fmt.Sprintf("tflblisuatdata%d", acctest.RandIntRange(10, 100))
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
				Config: testAccCheckIBMIsLbListenerPolicyDataSourceConfigBasic(vpcname, subnetname, acc.ISZoneName, acc.ISCIDR, lbname, port, protocol, lblistenerpolicyname1, action, priority1),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_lb_listener_policy.is_lb_listener_policy", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_lb_listener_policy.is_lb_listener_policy", "lb"),
					resource.TestCheckResourceAttrSet("data.ibm_is_lb_listener_policy.is_lb_listener_policy", "listener"),
					resource.TestCheckResourceAttrSet("data.ibm_is_lb_listener_policy.is_lb_listener_policy", "policy_id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_lb_listener_policy.is_lb_listener_policy", "action"),
					resource.TestCheckResourceAttrSet("data.ibm_is_lb_listener_policy.is_lb_listener_policy", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_is_lb_listener_policy.is_lb_listener_policy", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_lb_listener_policy.is_lb_listener_policy", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_is_lb_listener_policy.is_lb_listener_policy", "priority"),
					resource.TestCheckResourceAttrSet("data.ibm_is_lb_listener_policy.is_lb_listener_policy", "provisioning_status"),
				),
			},
		},
	})
}

func testAccCheckIBMIsLbListenerPolicyDataSourceConfigBasic(vpcname, subnetname, zone, cidr, lbname, port, protocol, lblistenerpolicyname, action, priority string) string {
	return testAccCheckIBMISLBListenerPolicyConfig(vpcname, subnetname, acc.ISZoneName, acc.ISCIDR, lbname, port, protocol, lblistenerpolicyname, action, priority) + fmt.Sprintf(`

	data "ibm_is_lb_listener_policy" "is_lb_listener_policy" {
		lb = "${ibm_is_lb.testacc_LB.id}"
		listener = ibm_is_lb_listener.testacc_lb_listener.listener_id
		policy_id = ibm_is_lb_listener_policy.testacc_lb_listener_policy.policy_id
	}
	`)
}
