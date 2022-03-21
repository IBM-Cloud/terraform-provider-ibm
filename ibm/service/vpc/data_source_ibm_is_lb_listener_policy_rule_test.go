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

func TestAccIBMIsLbListenerPolicyRuleDataSourceBasic(t *testing.T) {
	vpcname := fmt.Sprintf("tflblisuat-vpc-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tflblisuat-subnet-%d", acctest.RandIntRange(10, 100))
	lbname := fmt.Sprintf("tflblisuat%d", acctest.RandIntRange(10, 100))
	lblistenerpolicyname := fmt.Sprintf("tflblisuat-listener-policy-%d", acctest.RandIntRange(10, 100))
	lblistenerpolicyRuleField1 := fmt.Sprintf("tflblipolicy-rule-field-%d", acctest.RandIntRange(10, 100))
	lblistenerpolicyRuleValue1 := fmt.Sprintf("tflblipolicy-rule-value-%d", acctest.RandIntRange(10, 100))

	priority := "1"
	protocol := "http"
	port := "8080"
	action := "forward"
	condition := "equals"
	typeh := "header"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsLbListenerPolicyRuleDataSourceConfigBasic(vpcname, subnetname, acc.ISZoneName, acc.ISCIDR, lbname, port, protocol, lblistenerpolicyname, action, priority, condition, typeh, lblistenerpolicyRuleField1, lblistenerpolicyRuleValue1),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_lb_listener_policy_rule.is_lb_listener_policy_rule", "rule"),
					resource.TestCheckResourceAttrSet("data.ibm_is_lb_listener_policy_rule.is_lb_listener_policy_rule", "lb"),
					resource.TestCheckResourceAttrSet("data.ibm_is_lb_listener_policy_rule.is_lb_listener_policy_rule", "listener"),
					resource.TestCheckResourceAttrSet("data.ibm_is_lb_listener_policy_rule.is_lb_listener_policy_rule", "policy"),
					resource.TestCheckResourceAttrSet("data.ibm_is_lb_listener_policy_rule.is_lb_listener_policy_rule", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_lb_listener_policy_rule.is_lb_listener_policy_rule", "condition"),
					resource.TestCheckResourceAttrSet("data.ibm_is_lb_listener_policy_rule.is_lb_listener_policy_rule", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_is_lb_listener_policy_rule.is_lb_listener_policy_rule", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_lb_listener_policy_rule.is_lb_listener_policy_rule", "provisioning_status"),
					resource.TestCheckResourceAttrSet("data.ibm_is_lb_listener_policy_rule.is_lb_listener_policy_rule", "type"),
					resource.TestCheckResourceAttrSet("data.ibm_is_lb_listener_policy_rule.is_lb_listener_policy_rule", "value"),
				),
			},
		},
	})
}

func testAccCheckIBMIsLbListenerPolicyRuleDataSourceConfigBasic(vpcname, subnetname, zone, cidr, lbname, port, protocol, lblistenerpolicyname, action, priority, condition, types, field, value string) string {
	return testAccCheckIBMISLBListenerPolicyRuleConfig(vpcname, subnetname, acc.ISZoneName, acc.ISCIDR, lbname, port, protocol, lblistenerpolicyname, action, priority, condition, types, field, value) + fmt.Sprintf(`
	data "ibm_is_lb_listener_policy_rule" "is_lb_listener_policy_rule" {
		lb = "${ibm_is_lb.testacc_LB.id}"
		listener = "${ibm_is_lb_listener.testacc_lb_listener.listener_id}"
		policy = "${ibm_is_lb_listener_policy.testacc_lb_listener_policy.policy_id}"
		rule = "${ibm_is_lb_listener_policy_rule.testacc_lb_listener_policy_rule.rule}"
	}
	`)
}
