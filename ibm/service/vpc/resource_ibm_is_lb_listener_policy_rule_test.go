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

func TestAccIBMISLBListenerPolicyRule_basic(t *testing.T) {
	var ruleID string
	vpcname := fmt.Sprintf("tflblisuat-vpc-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tflblisuat-subnet-%d", acctest.RandIntRange(10, 100))
	lbname := fmt.Sprintf("tflblisuat%d", acctest.RandIntRange(10, 100))
	lblistenerpolicyname := fmt.Sprintf("tflblisuat-listener-policy-%d", acctest.RandIntRange(10, 100))
	//lblistenerpolicyname2 := fmt.Sprintf("tflblisuat-listener-policy-%d", acctest.RandIntRange(10, 100))
	lblistenerpolicyRuleField1 := fmt.Sprintf("tflblipolicy-rule-field-%d", acctest.RandIntRange(10, 100))
	lblistenerpolicyRuleField2 := fmt.Sprintf("tflblipolicy-rule-field-%d", acctest.RandIntRange(10, 100))
	lblistenerpolicyRuleField3 := fmt.Sprintf("tflblipolicy-rule-field-%d", acctest.RandIntRange(10, 100))
	lblistenerpolicyRuleValue1 := fmt.Sprintf("tflblipolicy-rule-value-%d", acctest.RandIntRange(10, 100))
	lblistenerpolicyRuleValue2 := fmt.Sprintf("tflblipolicy-rule-value-%d", acctest.RandIntRange(10, 100))
	lblistenerpolicyRuleValue3 := fmt.Sprintf("tflblipolicy-rule-value-%d", acctest.RandIntRange(10, 100))

	priority := "1"
	protocol := "http"
	port := "8080"
	action := "forward"
	//priority2 := "2"
	condition := "equals"
	typeh := "header"
	typeb := "body"
	typeSni := "sni_hostname"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISLBListenerPolicyRuleDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISLBListenerPolicyRuleConfig(vpcname, subnetname, acc.ISZoneName, acc.ISCIDR, lbname, port, protocol, lblistenerpolicyname, action, priority, condition, typeh, lblistenerpolicyRuleField1, lblistenerpolicyRuleValue1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISLBListenerPolicyRuleExists("ibm_is_lb_listener_policy_rule.testacc_lb_listener_policy_rule", ruleID),
					resource.TestCheckResourceAttr(
						"ibm_is_lb.testacc_LB", "name", lbname),
					resource.TestCheckResourceAttr(
						"ibm_is_lb_listener_policy_rule.testacc_lb_listener_policy_rule", "field", lblistenerpolicyRuleField1),
					resource.TestCheckResourceAttr(
						"ibm_is_lb_listener_policy_rule.testacc_lb_listener_policy_rule", "value", lblistenerpolicyRuleValue1),
				),
			},

			{
				Config: testAccCheckIBMISLBListenerPolicySniRuleConfig(vpcname, subnetname, acc.ISZoneName, acc.ISCIDR, lbname, port, protocol, lblistenerpolicyname, action, priority, condition, typeSni, lblistenerpolicyRuleField1, lblistenerpolicyRuleValue1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISLBListenerPolicyRuleExists("ibm_is_lb_listener_policy_rule.testacc_lb_listener_policy_rule", ruleID),
					resource.TestCheckResourceAttr(
						"ibm_is_lb.testacc_LB", "name", lbname),
					resource.TestCheckResourceAttr(
						"ibm_is_lb_listener_policy_rule.testacc_lb_listener_policy_rule", "field", lblistenerpolicyRuleField3),
					resource.TestCheckResourceAttr(
						"ibm_is_lb_listener_policy_rule.testacc_lb_listener_policy_rule", "value", lblistenerpolicyRuleValue3),
				),
			},

			{
				Config: testAccCheckIBMISLBListenerPolicyRuleConfigUpdate(vpcname, subnetname, acc.ISZoneName, acc.ISCIDR, lbname, port, protocol, lblistenerpolicyname, priority, condition, typeb, lblistenerpolicyRuleField2, lblistenerpolicyRuleValue2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISLBListenerPolicyRuleExists("ibm_is_lb_listener_policy_rule.testacc_lb_listener_policy_rule", ruleID),
					resource.TestCheckResourceAttr(
						"ibm_is_lb_listener_policy_rule.testacc_lb_listener_policy_rule", "field", lblistenerpolicyRuleField2),
					resource.TestCheckResourceAttr(
						"ibm_is_lb_listener_policy_rule.testacc_lb_listener_policy_rule", "value", lblistenerpolicyRuleValue2),
				),
			},
		},
	})
}

func testAccCheckIBMISLBListenerPolicyRuleDestroy(s *terraform.State) error {

	sess, _ := acc.TestAccProvider.Meta().(conns.ClientSession).VpcV1API()
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_is_lb_listener_policy_rule" {
			continue
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
		policyID := parts[2]
		ruleID := parts[3]

		getLbListenerPolicyRuleOptions := &vpcv1.GetLoadBalancerListenerPolicyRuleOptions{
			LoadBalancerID: &lbID,
			ListenerID:     &lbListenerID,
			PolicyID:       &policyID,
			ID:             &ruleID,
		}

		rule, _, err := sess.GetLoadBalancerListenerPolicyRule(getLbListenerPolicyRuleOptions)

		if err == nil {
			return fmt.Errorf("LBLIstenerPolicy still exists: %s %v", rs.Primary.ID, rule)
		}
	}

	return nil
}

func testAccCheckIBMISLBListenerPolicyRuleExists(n string, ruleID string) resource.TestCheckFunc {
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
		policyID := parts[2]
		ruleID := parts[3]

		sess, _ := acc.TestAccProvider.Meta().(conns.ClientSession).VpcV1API()
		getLbListenerPolicyRuleOptions := &vpcv1.GetLoadBalancerListenerPolicyRuleOptions{
			LoadBalancerID: &lbID,
			ListenerID:     &lbListenerID,
			PolicyID:       &policyID,
			ID:             &ruleID,
		}

		rule, _, err := sess.GetLoadBalancerListenerPolicyRule(getLbListenerPolicyRuleOptions)

		if err != nil {
			return err
		}

		ruleID = *rule.ID

		return nil
	}
}

func testAccCheckIBMISLBListenerPolicyRuleConfig(vpcname, subnetname, zone, cidr, lbname, port, protocol, lblistenerpolicyname, action, priority, condition, types, field, value string) string {
	return fmt.Sprintf(`
	resource "ibm_is_vpc" "testacc_vpc" {
		name = "%s"
	}

	resource "ibm_is_subnet" "testacc_subnet" {
		name = "%s"
		vpc = ibm_is_vpc.testacc_vpc.id
		zone = "%s"
		ipv4_cidr_block = "%s"
	}
	resource "ibm_is_lb" "testacc_LB" {
		name = "%s"
		subnets = [ibm_is_subnet.testacc_subnet.id]
	}
	resource "ibm_is_lb_listener" "testacc_lb_listener" {
		lb = ibm_is_lb.testacc_LB.id
		default_pool = ibm_is_lb_pool.testacc_pool.pool_id
		port = %s
		protocol = "%s"
	}

	resource "ibm_is_lb_pool" "testacc_pool" {
		name           = "test"
		lb             = ibm_is_lb.testacc_LB.id
		algorithm      = "round_robin"
		protocol       = "http"
		health_delay   = 60
		health_retries = 5
		health_timeout = 30
		health_type    = "http"
	}


	resource "ibm_is_lb_listener_policy" "testacc_lb_listener_policy" {
        lb = ibm_is_lb.testacc_LB.id
        listener = ibm_is_lb_listener.testacc_lb_listener.listener_id
        action = "%s"
		priority = %s
		name = "%s"
		target_id = ibm_is_lb_pool.testacc_pool.pool_id
	}


	resource "ibm_is_lb_listener_policy_rule" "testacc_lb_listener_policy_rule" {
		lb        = ibm_is_lb.testacc_LB.id
		listener  = ibm_is_lb_listener.testacc_lb_listener.listener_id
		policy    = ibm_is_lb_listener_policy.testacc_lb_listener_policy.policy_id
		condition = "%s"
		type      = "%s"
		field     = "%s"
		value     = "%s"
}`, vpcname, subnetname, zone, cidr, lbname, port, protocol, action, priority, lblistenerpolicyname, condition, types, field, value)

}

func testAccCheckIBMISLBListenerPolicyRuleConfigUpdate(vpcname, subnetname, zone, cidr, lbname, port, protocol, lblistenerpolicyname, priority, condition, types, field, value string) string {
	return fmt.Sprintf(`
	resource "ibm_is_vpc" "testacc_vpc" {
		name = "%s"
	}

	resource "ibm_is_subnet" "testacc_subnet" {
		name = "%s"
		vpc = ibm_is_vpc.testacc_vpc.id
		zone = "%s"
		ipv4_cidr_block = "%s"
	}
	resource "ibm_is_lb" "testacc_LB" {
		name = "%s"
		subnets = [ibm_is_subnet.testacc_subnet.id]
	}
	resource "ibm_is_lb_listener" "testacc_lb_listener" {
		lb = ibm_is_lb.testacc_LB.id
		port = %s
		protocol = "%s"
		default_pool = ibm_is_lb_pool.testacc_pool.pool_id
	}

	resource "ibm_is_lb_pool" "testacc_pool" {
		name           = "test"
		lb             = ibm_is_lb.testacc_LB.id
		algorithm      = "round_robin"
		protocol       = "http"
		health_delay   = 60
		health_retries = 5
		health_timeout = 30
		health_type    = "http"
	}

	resource "ibm_is_lb_listener_policy" "testacc_lb_listener_policy" {
        lb = ibm_is_lb.testacc_LB.id
        listener = ibm_is_lb_listener.testacc_lb_listener.listener_id
        action = "forward"
		priority = %s
		name = "%s"
		target_id = ibm_is_lb_pool.testacc_pool.pool_id
	}

	resource "ibm_is_lb_listener_policy_rule" "testacc_lb_listener_policy_rule" {
		lb        = ibm_is_lb.testacc_LB.id
		listener  = ibm_is_lb_listener.testacc_lb_listener.listener_id
		policy    = ibm_is_lb_listener_policy.testacc_lb_listener_policy.policy_id
		condition = "%s"
		type      = "%s"
		field     = "%s"
		value     = "%s"
}`, vpcname, subnetname, zone, cidr, lbname, port, protocol, priority, lblistenerpolicyname, condition, types, field, value)

}

func testAccCheckIBMISLBListenerPolicySniRuleConfig(vpcname, subnetname, zone, cidr, lbname, port, protocol, lblistenerpolicyname, action, priority, condition, types, field, value string) string {
	return fmt.Sprintf(`
	resource "ibm_is_vpc" "testacc_vpc" {
		name = "%s"
	}

	resource "ibm_is_subnet" "testacc_subnet" {
		name = "%s"
		vpc = ibm_is_vpc.testacc_vpc.id
		zone = "%s"
		ipv4_cidr_block = "%s"
	}
	resource "ibm_is_lb" "testacc_LB" {
		name = "%s"
		subnets = [ibm_is_subnet.testacc_subnet.id]
	}
	resource "ibm_is_lb_listener" "testacc_lb_listener" {
		lb = ibm_is_lb.testacc_LB.id
		default_pool = ibm_is_lb_pool.testacc_pool.pool_id
		port = %s
		protocol = "%s"
	}

	resource "ibm_is_lb_pool" "testacc_pool" {
		name           = "test"
		lb             = ibm_is_lb.testacc_LB.id
		algorithm      = "round_robin"
		protocol       = "http"
		health_delay   = 60
		health_retries = 5
		health_timeout = 30
		health_type    = "http"
	}


	resource "ibm_is_lb_listener_policy" "testacc_lb_listener_policy" {
        lb = ibm_is_lb.testacc_LB.id
        listener = ibm_is_lb_listener.testacc_lb_listener.listener_id
        action = "%s"
		priority = %s
		name = "%s"
		target_id = ibm_is_lb_pool.testacc_pool.pool_id
	}


	resource "ibm_is_lb_listener_policy_rule" "testacc_lb_listener_policy_rule" {
		lb        = ibm_is_lb.testacc_LB.id
		listener  = ibm_is_lb_listener.testacc_lb_listener.listener_id
		policy    = ibm_is_lb_listener_policy.testacc_lb_listener_policy.policy_id
		condition = "%s"
		type      = "%s"
		field     = "%s"
		value     = "%s"
}`, vpcname, subnetname, zone, cidr, lbname, port, protocol, action, priority, lblistenerpolicyname, condition, types, field, value)
}
