// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"

	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccIBMISSecurityGroupRule_basic(t *testing.T) {
	var securityGroupRule string

	vpcname := fmt.Sprintf("tfsgrule-vpc-%d", acctest.RandIntRange(10, 100))
	name1 := fmt.Sprintf("tfsgrule-createname-%d", acctest.RandIntRange(10, 100))
	//name2 := fmt.Sprintf("tfsgrule-updatename-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISSecurityGroupRuleDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISsecurityGroupRuleConfig(vpcname, name1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISSecurityGroupRuleExists("ibm_is_security_group_rule.testacc_security_group_rule_all", securityGroupRule),
					resource.TestCheckResourceAttr(
						"ibm_is_security_group.testacc_security_group", "name", name1),
				),
			},
		},
	})
}
func parseISTerraformID(s string) (string, string, error) {
	segments := strings.Split(s, ".")
	if len(segments) != 2 {
		return "", "", fmt.Errorf("invalid terraform Id %s (incorrect number of segments)", s)
	}
	if segments[0] == "" || segments[1] == "" {
		return "", "", fmt.Errorf("invalid terraform Id %s (one or more empty segments)", s)
	}
	return segments[0], segments[1], nil
}
func testAccCheckIBMISSecurityGroupRuleDestroy(s *terraform.State) error {
	sess, _ := acc.TestAccProvider.Meta().(conns.ClientSession).VpcV1API()
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_is_security_group_rule" {
			continue
		}

		secgrpID, ruleID, err := parseISTerraformID(rs.Primary.ID)
		if err != nil {
			return err
		}
		getsgruleoptions := &vpcv1.GetSecurityGroupRuleOptions{
			SecurityGroupID: &secgrpID,
			ID:              &ruleID,
		}
		_, _, err1 := sess.GetSecurityGroupRule(getsgruleoptions)
		if err1 == nil {
			return fmt.Errorf("securitygrouprule still exists: %s", rs.Primary.ID)
		}
	}
	return nil
}

func testAccCheckIBMISSecurityGroupRuleExists(n, securityGroupRuleID string) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return errors.New("No Record ID is set")
		}
		secgrpID, ruleID, err := parseISTerraformID(rs.Primary.ID)
		if err != nil {
			return err
		}
		sess, _ := acc.TestAccProvider.Meta().(conns.ClientSession).VpcV1API()
		getsgruleoptions := &vpcv1.GetSecurityGroupRuleOptions{
			SecurityGroupID: &secgrpID,
			ID:              &ruleID,
		}
		foundSecurityGroupRule, _, err := sess.GetSecurityGroupRule(getsgruleoptions)
		if err != nil {
			return err
		}
		switch reflect.TypeOf(foundSecurityGroupRule).String() {
		case "*vpcv1.SecurityGroupRuleSecurityGroupRuleProtocolIcmp":
			{
				sgr := foundSecurityGroupRule.(*vpcv1.SecurityGroupRuleSecurityGroupRuleProtocolIcmp)
				securityGroupRuleID = *sgr.ID
			}
		case "*vpcv1.SecurityGroupRuleSecurityGroupRuleProtocolAll":
			{
				sgr := foundSecurityGroupRule.(*vpcv1.SecurityGroupRuleSecurityGroupRuleProtocolAll)
				securityGroupRuleID = *sgr.ID
			}
		case "*vpcv1.SecurityGroupRuleSecurityGroupRuleProtocolTcpudp":
			{
				sgr := foundSecurityGroupRule.(*vpcv1.SecurityGroupRuleSecurityGroupRuleProtocolTcpudp)
				securityGroupRuleID = *sgr.ID
			}
		}
		return nil
	}
}

func testAccCheckIBMISsecurityGroupRuleConfig(vpcname, name string) string {
	return fmt.Sprintf(`
	resource "ibm_is_vpc" "testacc_vpc" {
		name = "%s"
	  }
	  
	  resource "ibm_is_security_group" "testacc_security_group" {
		name = "%s"
		vpc  = ibm_is_vpc.testacc_vpc.id
	  }
	  
	  resource "ibm_is_security_group_rule" "testacc_security_group_rule_all" {
		group     = ibm_is_security_group.testacc_security_group.id
		direction = "inbound"
		remote    = "127.0.0.1"
	  }
	  
	  resource "ibm_is_security_group_rule" "testacc_security_group_rule_icmp" {
		depends_on = [ibm_is_security_group_rule.testacc_security_group_rule_all]
		group      = ibm_is_security_group.testacc_security_group.id
		direction  = "inbound"
		remote     = "127.0.0.1"
		icmp {
		  code = 20
		  type = 30
		}
	  }

	  resource "ibm_is_security_group_rule" "testacc_security_group_rule_icmp_code_any" {
		depends_on = [ibm_is_security_group_rule.testacc_security_group_rule_icmp]
		group      = ibm_is_security_group.testacc_security_group.id
		direction  = "inbound"
		remote     = "127.0.0.1"
		icmp {
		  type = 30
		}
	  }
	  
	  resource "ibm_is_security_group_rule" "testacc_security_group_rule_udp" {
		depends_on = [ibm_is_security_group_rule.testacc_security_group_rule_icmp]
		group      = ibm_is_security_group.testacc_security_group.id
		direction  = "inbound"
		remote     = "127.0.0.1"
		udp {
		  port_min = 805
		  port_max = 807
		}
	  }
	  
	  resource "ibm_is_security_group_rule" "testacc_security_group_rule_tcp" {
		depends_on = [ibm_is_security_group_rule.testacc_security_group_rule_udp]
		group      = ibm_is_security_group.testacc_security_group.id
		direction  = "inbound"
		remote     = "127.0.0.1"
		tcp {
		  port_min = 8080
		  port_max = 8080
		}
	  }

	  resource "ibm_is_security_group_rule" "testacc_security_group_rule_icmp_any" {
		depends_on = [ibm_is_security_group_rule.testacc_security_group_rule_tcp]
		group      = ibm_is_security_group.testacc_security_group.id
		direction  = "inbound"
		remote     = "127.0.0.1"
		icmp {
		}
	  }
	  
	  resource "ibm_is_security_group_rule" "testacc_security_group_rule_udp_any" {
		depends_on = [ibm_is_security_group_rule.testacc_security_group_rule_icmp_any]
		group      = ibm_is_security_group.testacc_security_group.id
		direction  = "inbound"
		remote     = "127.0.0.1"
		udp {
		}
	  }
	  
	  resource "ibm_is_security_group_rule" "testacc_security_group_rule_tcp_any" {
		depends_on = [ibm_is_security_group_rule.testacc_security_group_rule_udp_any]
		group      = ibm_is_security_group.testacc_security_group.id
		direction  = "inbound"
		remote     = "127.0.0.1"
		local 	   = "192.168.3.4"
		tcp {
		}
	  }
 `, vpcname, name)

}
