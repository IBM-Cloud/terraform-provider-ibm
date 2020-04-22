package ibm

import (
	"errors"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.ibm.com/ibmcloud/vpc-go-sdk/vpcclassicv1"
	"github.ibm.com/ibmcloud/vpc-go-sdk/vpcv1"
)

func TestAccIBMISLBListenerPolicyRule_basic(t *testing.T) {
	var ruleID string
	vpcname := fmt.Sprintf("terraformLBLisuat-vpc-%d", acctest.RandInt())
	subnetname := fmt.Sprintf("terraformLBLisuat-subnet-%d", acctest.RandInt())
	lbname := fmt.Sprintf("tflblisuat%d", acctest.RandInt())
	lblistenerpolicyname1 := fmt.Sprintf("tflblisuat-listener-policy-%d", acctest.RandInt())
	lblistenerpolicyname2 := fmt.Sprintf("tflblisuat-listener-policy-%d", acctest.RandInt())
	lblistenerpolicyRuleField1 := fmt.Sprintf("tflblipolicy-rule-field-%d", acctest.RandInt())
	lblistenerpolicyRuleField2 := fmt.Sprintf("tflblipolicy-rule-field-%d", acctest.RandInt())
	lblistenerpolicyRuleValue1 := fmt.Sprintf("tflblipolicy-rule-value-%d", acctest.RandInt())
	lblistenerpolicyRuleValue2 := fmt.Sprintf("tflblipolicy-rule-value-%d", acctest.RandInt())

	priority1 := "1"
	protocol := "http"
	port := "8080"
	action := "forward"
	priority2 := "2"
	condition := "equals"
	types := "header"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMISLBListenerPolicyRuleDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMISLBListenerPolicyRuleConfig(vpcname, subnetname, ISZoneName, ISCIDR, lbname, port, protocol, lblistenerpolicyname1, action, priority1, condition, types, lblistenerpolicyRuleField1, lblistenerpolicyRuleValue1),
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

			resource.TestStep{
				Config: testAccCheckIBMISLBListenerPolicyRuleConfigUpdate(vpcname, subnetname, ISZoneName, ISCIDR, lbname, port, protocol, lblistenerpolicyname2, priority2, condition, types, lblistenerpolicyRuleField2, lblistenerpolicyRuleValue2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISLBListenerPolicyRuleExists("ibm_is_lb_listener_policy_rule.testacc_lb_listener_policy_rule", ruleID),
					resource.TestCheckResourceAttr(
						"ibm_is_lb_listener_policy.testacc_lb_listener_policy", "name", lblistenerpolicyname2),
					resource.TestCheckResourceAttr(
						"ibm_is_lb_listener_policy.testacc_lb_listener_policy", "proprity", priority2),
				),
			},
		},
	})
}

func testAccCheckIBMISLBListenerPolicyRuleDestroy(s *terraform.State) error {
	userDetails, _ := testAccProvider.Meta().(ClientSession).BluemixUserDetails()

	if userDetails.generation == 1 {
		sess, _ := testAccProvider.Meta().(ClientSession).VpcClassicV1API()
		for _, rs := range s.RootModule().Resources {
			if rs.Type != "ibm_is_lb_listener_policy_rule" {
				continue
			}

			if rs.Primary.ID == "" {
				return errors.New("No Record ID is set")
			}

			parts, err := idParts(rs.Primary.ID)
			if err != nil {
				return err
			}

			lbID := parts[0]
			lbListenerID := parts[1]
			policyID := parts[2]
			ruleID := parts[3]

			getLbListenerPolicyRuleOptions := &vpcclassicv1.GetLoadBalancerListenerPolicyRuleOptions{
				LoadBalancerID: &lbID,
				ListenerID:     &lbListenerID,
				PolicyID:       &policyID,
				ID:             &ruleID,
			}

			rule, _, err := sess.GetLoadBalancerListenerPolicyRule(getLbListenerPolicyRuleOptions)

			if err == nil {
				return fmt.Errorf("LBLIstenerPolicyRule still exists: %s %v", rs.Primary.ID, rule)
			}
		}
	} else {
		sess, _ := testAccProvider.Meta().(ClientSession).VpcV1API()
		for _, rs := range s.RootModule().Resources {
			if rs.Type != "ibm_is_lb_listener_policy_rule" {
				continue
			}

			if rs.Primary.ID == "" {
				return errors.New("No Record ID is set")
			}

			parts, err := idParts(rs.Primary.ID)
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

		parts, err := idParts(rs.Primary.ID)
		if err != nil {
			return err
		}

		lbID := parts[0]
		lbListenerID := parts[1]
		policyID := parts[2]
		ruleID := parts[3]

		userDetails, _ := testAccProvider.Meta().(ClientSession).BluemixUserDetails()

		if userDetails.generation == 1 {
			sess, _ := testAccProvider.Meta().(ClientSession).VpcClassicV1API()

			getLbListenerPolicyRuleOptions := &vpcclassicv1.GetLoadBalancerListenerPolicyRuleOptions{
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

		} else {
			sess, _ := testAccProvider.Meta().(ClientSession).VpcV1API()
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
		}
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
		subnets = ["ibm_is_subnet.testacc_subnet.id"]
	}
	resource "ibm_is_lb_listener" "testacc_lb_listener" {
		lb = ibm_is_lb.testacc_LB.id
		port = %s
		protocol = "%s"
	}

	resource "ibm_is_lb_listener_policy" "testacc_lb_listener_policy" {
        lb = ibm_is_lb.testacc_LB.id
        listener = ibm_is_lb_listener.testacc_lb_listener.listener_id
        action = "%s"
		priority = %s
		name = "%s"
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
		vpc = "${ibm_is_vpc.testacc_vpc.id}"
		zone = "%s"
		ipv4_cidr_block = "%s"
	}
	resource "ibm_is_lb" "testacc_LB" {
		name = "%s"
		subnets = ["${ibm_is_subnet.testacc_subnet.id}"]
	}
	resource "ibm_is_lb_listener" "testacc_lb_listener" {
		lb = "${ibm_is_lb.testacc_LB.id}"
		port = %s
		protocol = "%s"
	}

	resource "ibm_is_lb_listener_policy" "testacc_lb_listener_policy" {
        lb = ibm_is_lb.testacc_LB.id
        listener = ibm_is_lb_listener.testacc_lb_listener.listener_id
        action = "forward"
		priority = %s
		name = "%s"
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
