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
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestNetworkACLRule_basicICMP(t *testing.T) {
	var nwACLRule string
	vpcName := fmt.Sprintf("tf-nacl-vpc-%d", acctest.RandIntRange(10, 100))
	ruleName := fmt.Sprintf("tf-outbound-icmp-%d", acctest.RandIntRange(10, 100))
	updatedRuleName := fmt.Sprintf("%s-update", ruleName)
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: checkNetworkACLRuleDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISNetworkACLRuleConfig1(vpcName, ruleName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISNetworkACLRuleExists("ibm_is_network_acl_rule.testacc_nacl", nwACLRule),
					resource.TestCheckResourceAttr(
						"ibm_is_network_acl_rule.testacc_nacl", "name", ruleName),
					resource.TestCheckResourceAttr(
						"ibm_is_network_acl_rule.testacc_nacl", "icmp.0.code", "1"),
					resource.TestCheckResourceAttr(
						"ibm_is_network_acl_rule.testacc_nacl", "icmp.0.type", "1"),
				),
			},
			{
				Config: testAccCheckIBMISNetworkACLRuleConfig1Update(vpcName, updatedRuleName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISNetworkACLRuleExists("ibm_is_network_acl_rule.testacc_nacl", nwACLRule),
					resource.TestCheckResourceAttr(
						"ibm_is_network_acl_rule.testacc_nacl", "name", updatedRuleName),
					resource.TestCheckResourceAttr(
						"ibm_is_network_acl_rule.testacc_nacl", "icmp.0.code", "2"),
					resource.TestCheckResourceAttr(
						"ibm_is_network_acl_rule.testacc_nacl", "icmp.0.type", "2"),
				),
			},
		},
	})
}

func TestNetworkACLRule_basicAll(t *testing.T) {
	var nwACLRule string
	vpcName := fmt.Sprintf("tf-nacl-vpc-%d", acctest.RandIntRange(10, 100))
	ruleName := fmt.Sprintf("tf-outbound-all-%d", acctest.RandIntRange(10, 100))
	updatedRuleName := fmt.Sprintf("%s-update", ruleName)
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: checkNetworkACLRuleDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISNetworkACLRuleConfig2(vpcName, ruleName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISNetworkACLRuleExists("ibm_is_network_acl_rule.testacc_nacl", nwACLRule),
					resource.TestCheckResourceAttr(
						"ibm_is_network_acl_rule.testacc_nacl", "name", ruleName),
				),
			},
			{
				Config: testAccCheckIBMISNetworkACLRuleConfig2Update(vpcName, updatedRuleName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISNetworkACLRuleExists("ibm_is_network_acl_rule.testacc_nacl", nwACLRule),
					resource.TestCheckResourceAttr(
						"ibm_is_network_acl_rule.testacc_nacl", "name", updatedRuleName),
				),
			},
		},
	})
}

func TestNetworkACLRule_basicTCP(t *testing.T) {
	var nwACLRule string
	vpcName := fmt.Sprintf("tf-nacl-vpc-%d", acctest.RandIntRange(10, 100))
	ruleName := fmt.Sprintf("tf-outbound-tcp-%d", acctest.RandIntRange(10, 100))
	updatedRuleName := fmt.Sprintf("%s-update", ruleName)
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: checkNetworkACLRuleDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISNetworkACLRuleConfig3(vpcName, ruleName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISNetworkACLRuleExists("ibm_is_network_acl_rule.testacc_nacl", nwACLRule),
					resource.TestCheckResourceAttr(
						"ibm_is_network_acl_rule.testacc_nacl", "name", ruleName),
					resource.TestCheckResourceAttr(
						"ibm_is_network_acl_rule.testacc_nacl", "tcp.0.source_port_min", "1000"),
					resource.TestCheckResourceAttr(
						"ibm_is_network_acl_rule.testacc_nacl", "tcp.0.source_port_max", "1101"),
				),
			},
			{
				Config: testAccCheckIBMISNetworkACLRuleConfig3Update(vpcName, updatedRuleName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISNetworkACLRuleExists("ibm_is_network_acl_rule.testacc_nacl", nwACLRule),
					resource.TestCheckResourceAttr(
						"ibm_is_network_acl_rule.testacc_nacl", "name", updatedRuleName),
					resource.TestCheckResourceAttr(
						"ibm_is_network_acl_rule.testacc_nacl", "tcp.0.source_port_min", "1002"),
					resource.TestCheckResourceAttr(
						"ibm_is_network_acl_rule.testacc_nacl", "tcp.0.source_port_max", "1101"),
				),
			},
		},
	})
}

func TestNetworkACLRule_basicUDP(t *testing.T) {
	var nwACLRule string
	vpcName := fmt.Sprintf("tf-nacl-vpc-%d", acctest.RandIntRange(10, 100))
	ruleName := fmt.Sprintf("tf-outbound-udp-%d", acctest.RandIntRange(10, 100))
	updatedRuleName := fmt.Sprintf("%s-update", ruleName)
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: checkNetworkACLRuleDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISNetworkACLRuleConfig4(vpcName, ruleName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISNetworkACLRuleExists("ibm_is_network_acl_rule.testacc_nacl", nwACLRule),
					resource.TestCheckResourceAttr(
						"ibm_is_network_acl_rule.testacc_nacl", "name", ruleName),
					resource.TestCheckResourceAttr(
						"ibm_is_network_acl_rule.testacc_nacl", "udp.0.source_port_max", "101"),
					resource.TestCheckResourceAttr(
						"ibm_is_network_acl_rule.testacc_nacl", "udp.0.source_port_min", "1"),
				),
			},
			{
				Config: testAccCheckIBMISNetworkACLRuleConfig4Update(vpcName, updatedRuleName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISNetworkACLRuleExists("ibm_is_network_acl_rule.testacc_nacl", nwACLRule),
					resource.TestCheckResourceAttr(
						"ibm_is_network_acl_rule.testacc_nacl", "name", updatedRuleName),
					resource.TestCheckResourceAttr(
						"ibm_is_network_acl_rule.testacc_nacl", "udp.0.source_port_max", "101"),
					resource.TestCheckResourceAttr(
						"ibm_is_network_acl_rule.testacc_nacl", "udp.0.source_port_min", "2"),
				),
			},
		},
	})
}

func TestNetworkACLRule_basicBeforeRule(t *testing.T) {
	var nwACLRule string
	vpcName := fmt.Sprintf("tf-nacl-vpc-%d", acctest.RandIntRange(10, 100))
	ruleName := fmt.Sprintf("tf-outbound-udp-%d", acctest.RandIntRange(10, 100))
	ruleName1 := fmt.Sprintf("tf-outbound-udp1-%d", acctest.RandIntRange(10, 100))
	updatedRuleName := fmt.Sprintf("%s-update", ruleName)
	updatedRule1Name := fmt.Sprintf("%s-update", ruleName1)
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: checkNetworkACLRuleDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISNetworkACLRuleBeforeConfig(vpcName, ruleName, ruleName1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISNetworkACLRuleExists("ibm_is_network_acl_rule.testacc_nacl", nwACLRule),
					resource.TestCheckResourceAttr(
						"ibm_is_network_acl_rule.testacc_nacl", "name", ruleName),
					resource.TestCheckResourceAttr(
						"ibm_is_network_acl_rule.testacc_nacl", "udp.0.source_port_max", "101"),
					resource.TestCheckResourceAttr(
						"ibm_is_network_acl_rule.testacc_nacl", "udp.0.source_port_min", "1"),
				),
			},
			{
				Config: testAccCheckIBMISNetworkACLRuleBeforeUpdateConfig(vpcName, updatedRuleName, updatedRule1Name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISNetworkACLRuleExists("ibm_is_network_acl_rule.testacc_nacl", nwACLRule),
					resource.TestCheckResourceAttr(
						"ibm_is_network_acl_rule.testacc_nacl", "name", updatedRuleName),
					resource.TestCheckResourceAttr(
						"ibm_is_network_acl_rule.testacc_nacl", "udp.0.source_port_max", "101"),
					resource.TestCheckResourceAttr(
						"ibm_is_network_acl_rule.testacc_nacl", "udp.0.source_port_min", "1"),
					resource.TestCheckResourceAttr(
						"ibm_is_network_acl_rule.testacc_nacl", "before", "null"),
				),
			},
		},
	})
}

func checkNetworkACLRuleDestroy(s *terraform.State) error {
	sess, _ := acc.TestAccProvider.Meta().(conns.ClientSession).VpcV1API()
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_is_network_acl_rule" {
			continue
		}
		nwACLID, ruleID, err := parseNwACLTerraformID(rs.Primary.ID)
		getnwaclRuleoptions := &vpcv1.GetNetworkACLRuleOptions{
			ID:           &ruleID,
			NetworkACLID: &nwACLID,
		}
		_, _, err = sess.GetNetworkACLRule(getnwaclRuleoptions)
		if err == nil {
			return fmt.Errorf("network acl rule still exists: %s", rs.Primary.ID)
		}
	}
	return nil
}
func makeTerraformACLRuleID(id1, id2 string) string {
	// Include both network acl id and rule id to create a unique Terraform id.  As a bonus,
	// we can extract the network acl id as needed for API calls such as READ.
	return fmt.Sprintf("%s/%s", id1, id2)
}
func parseNwACLTerraformID(s string) (string, string, error) {
	segments := strings.Split(s, "/")
	if len(segments) != 2 {
		return "", "", fmt.Errorf("invalid terraform Id %s (incorrect number of segments)", s)
	}
	if segments[0] == "" || segments[1] == "" {
		return "", "", fmt.Errorf("invalid terraform Id %s (one or more empty segments)", s)
	}
	return segments[0], segments[1], nil
}

func testAccCheckIBMISNetworkACLRuleExists(n, nwACLRule string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return errors.New("No Record ID is set")
		}

		sess, _ := acc.TestAccProvider.Meta().(conns.ClientSession).VpcV1API()
		nwACLID, ruleID, err := parseNwACLTerraformID(rs.Primary.ID)
		getnwaclRuleoptions := &vpcv1.GetNetworkACLRuleOptions{
			ID:           &ruleID,
			NetworkACLID: &nwACLID,
		}
		foundNwACLRule, _, err := sess.GetNetworkACLRule(getnwaclRuleoptions)
		if err != nil {
			return err
		}
		switch reflect.TypeOf(foundNwACLRule).String() {
		case "*vpcv1.NetworkACLRuleNetworkACLRuleProtocolIcmp":
			{
				rulex := foundNwACLRule.(*vpcv1.NetworkACLRuleNetworkACLRuleProtocolIcmp)
				nwACLRule = makeTerraformACLRuleID(nwACLID, *rulex.ID)
			}
		case "*vpcv1.NetworkACLRuleNetworkACLRuleProtocolTcpudp":
			{
				rulex := foundNwACLRule.(*vpcv1.NetworkACLRuleNetworkACLRuleProtocolTcpudp)
				nwACLRule = makeTerraformACLRuleID(nwACLID, *rulex.ID)

			}
		case "*vpcv1.NetworkACLRuleNetworkACLRuleProtocolAll":
			{
				rulex := foundNwACLRule.(*vpcv1.NetworkACLRuleNetworkACLRuleProtocolAll)
				nwACLRule = makeTerraformACLRuleID(nwACLID, *rulex.ID)

			}
		}
		return nil
	}
}

func testAccCheckIBMISNetworkACLRuleConfig1(vpcName, name string) string {
	return fmt.Sprintf(`
	  resource "ibm_is_vpc" "testacc_vpc" {
		name = "%s"
	  }
	  resource "ibm_is_network_acl_rule" "testacc_nacl" {
		network_acl = ibm_is_vpc.testacc_vpc.default_network_acl
		name           = "%s"
		action         = "allow"
		source         = "0.0.0.0/0"
		destination    = "0.0.0.0/0"
		direction      = "outbound"
		icmp {
			code = 1
			type = 1
			}
		}
	`, vpcName, name)
}

func testAccCheckIBMISNetworkACLRuleConfig1Update(vpcName, name string) string {
	return fmt.Sprintf(`
	  resource "ibm_is_vpc" "testacc_vpc" {
		name = "%s"
	  }
	  resource "ibm_is_network_acl_rule" "testacc_nacl" {
		network_acl = ibm_is_vpc.testacc_vpc.default_network_acl
		name           = "%s"
		action         = "allow"
		source         = "0.0.0.0/0"
		destination    = "0.0.0.0/0"
		direction      = "outbound"
		icmp {
			code = 2
			type = 2
			}
		}
	`, vpcName, name)
}
func testAccCheckIBMISNetworkACLRuleConfig2(vpcName, name string) string {
	return fmt.Sprintf(`
	  resource "ibm_is_vpc" "testacc_vpc" {
		name = "%s"
	  }
	  resource "ibm_is_network_acl_rule" "testacc_nacl" {
		network_acl = ibm_is_vpc.testacc_vpc.default_network_acl
		name           = "%s"
		action         = "allow"
		source         = "0.0.0.0/0"
		destination    = "0.0.0.0/0"
		direction      = "outbound"
		}
	`, vpcName, name)
}
func testAccCheckIBMISNetworkACLRuleConfig2Update(vpcName, name string) string {
	return fmt.Sprintf(`
	  resource "ibm_is_vpc" "testacc_vpc" {
		name = "%s"
	  }
	  resource "ibm_is_network_acl_rule" "testacc_nacl" {
		network_acl = ibm_is_vpc.testacc_vpc.default_network_acl
		name           = "%s"
		action         = "allow"
		source         = "0.0.0.0/0"
		destination    = "0.0.0.0/0"
		direction      = "outbound"
		}
	`, vpcName, name)
}

func testAccCheckIBMISNetworkACLRuleConfig3(vpcName, name string) string {
	return fmt.Sprintf(`
	   resource "ibm_is_vpc" "testacc_vpc" {
			name = "%s"
		}
	  resource "ibm_is_network_acl_rule" "testacc_nacl" {
		network_acl = ibm_is_vpc.testacc_vpc.default_network_acl
		name           = "%s"
		action         = "allow"
		source         = "0.0.0.0/0"
		destination    = "0.0.0.0/0"
		direction      = "outbound"
		tcp {
			source_port_max = 1101
			source_port_min = 1000
			port_min = 2020
			port_max = 2200
	   		}
		}
	`, vpcName, name)
}
func testAccCheckIBMISNetworkACLRuleConfig3Update(vpcName, name string) string {
	return fmt.Sprintf(`
	   resource "ibm_is_vpc" "testacc_vpc" {
			name = "%s"
		}
	  resource "ibm_is_network_acl_rule" "testacc_nacl" {
		network_acl = ibm_is_vpc.testacc_vpc.default_network_acl
		name           = "%s"
		action         = "allow"
		source         = "0.0.0.0/0"
		destination    = "0.0.0.0/0"
		direction      = "outbound"
		tcp {
			source_port_max = 1101
			source_port_min = 1002
			port_min = 2020
			port_max = 2200
	   		}
		}
	`, vpcName, name)
}

func testAccCheckIBMISNetworkACLRuleConfig4(vpcName, name string) string {
	return fmt.Sprintf(`
	resource "ibm_is_vpc" "testacc_vpc" {
		name = "%s"
	}
	resource "ibm_is_network_acl_rule" "testacc_nacl" {
		network_acl = ibm_is_vpc.testacc_vpc.default_network_acl
		name           = "%s"
		action         = "allow"
		source         = "0.0.0.0/0"
		destination    = "0.0.0.0/0"
		direction      = "outbound"
		udp {
			source_port_max = 101
			source_port_min = 1
			port_min = 202
			port_max = 220
	   		}
		}
	`, vpcName, name)
}
func testAccCheckIBMISNetworkACLRuleConfig4Update(vpcName, name string) string {
	return fmt.Sprintf(`
	resource "ibm_is_vpc" "testacc_vpc" {
		name = "%s"
	}
	resource "ibm_is_network_acl_rule" "testacc_nacl" {
		network_acl = ibm_is_vpc.testacc_vpc.default_network_acl
		name           = "%s"
		action         = "allow"
		source         = "0.0.0.0/0"
		destination    = "0.0.0.0/0"
		direction      = "outbound"
		udp {
			source_port_max = 101
			source_port_min = 2
			port_min = 202
			port_max = 220
	   		}
		}
	`, vpcName, name)
}

func testAccCheckIBMISNetworkACLRuleBeforeConfig(vpcName, name, name1 string) string {
	return fmt.Sprintf(`
	resource "ibm_is_vpc" "testacc_vpc" {
		name = "%s"
	}
	resource "ibm_is_network_acl_rule" "testacc_nacl" {
		network_acl = ibm_is_vpc.testacc_vpc.default_network_acl
		name           = "%s"
		action         = "allow"
		source         = "0.0.0.0/0"
		destination    = "0.0.0.0/0"
		direction      = "outbound"
		before = ibm_is_network_acl_rule.testacc_nacl_1.rule_id
		udp {
			source_port_max = 101
			source_port_min = 1
			port_min = 202
			port_max = 220
			}
	}
	resource "ibm_is_network_acl_rule" "testacc_nacl_1" {
		network_acl = ibm_is_vpc.testacc_vpc.default_network_acl
		name           = "%s"
		action         = "allow"
		source         = "0.0.0.0/0"
		destination    = "0.0.0.0/0"
		direction      = "outbound"
		udp {
			source_port_max = 101
			source_port_min = 1
			port_min = 202
			port_max = 220
			}
	}
	`, vpcName, name, name1)
}
func testAccCheckIBMISNetworkACLRuleBeforeUpdateConfig(vpcName, name, name1 string) string {
	return fmt.Sprintf(`
	resource "ibm_is_vpc" "testacc_vpc" {
		name = "%s"
	}
	resource "ibm_is_network_acl_rule" "testacc_nacl" {
		network_acl = ibm_is_vpc.testacc_vpc.default_network_acl
		name           = "%s"
		action         = "allow"
		source         = "0.0.0.0/0"
		destination    = "0.0.0.0/0"
		direction      = "outbound"
		before = "null"
		udp {
			source_port_max = 101
			source_port_min = 1
			port_min = 202
			port_max = 220
			}
	}
	resource "ibm_is_network_acl_rule" "testacc_nacl_1" {
		network_acl = ibm_is_vpc.testacc_vpc.default_network_acl
		name           = "%s"
		action         = "allow"
		source         = "0.0.0.0/0"
		destination    = "0.0.0.0/0"
		direction      = "outbound"
		udp {
			source_port_max = 101
			source_port_min = 1
			port_min = 202
			port_max = 220
			}
	}
	`, vpcName, name, name1)
}
