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

func TestNetworkACLRule_basicIcmpTcpUdp(t *testing.T) {
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

func TestNetworkACLRule_basicAny(t *testing.T) {
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
				Config: testAccCheckIBMISNetworkACLRuleConfig5(vpcName, ruleName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISNetworkACLRuleExists("ibm_is_network_acl_rule.testacc_nacl", nwACLRule),
					resource.TestCheckResourceAttr(
						"ibm_is_network_acl_rule.testacc_nacl", "name", ruleName),
				),
			},
			{
				Config: testAccCheckIBMISNetworkACLRuleConfig5Update(vpcName, updatedRuleName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISNetworkACLRuleExists("ibm_is_network_acl_rule.testacc_nacl", nwACLRule),
					resource.TestCheckResourceAttr(
						"ibm_is_network_acl_rule.testacc_nacl", "name", updatedRuleName),
				),
			},
		},
	})
}

func TestNetworkACLRule_basicIndividual(t *testing.T) {
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
				Config: testAccCheckIBMISNetworkACLRuleConfig6(vpcName, ruleName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISNetworkACLRuleExists("ibm_is_network_acl_rule.testacc_nacl", nwACLRule),
					resource.TestCheckResourceAttr(
						"ibm_is_network_acl_rule.testacc_nacl", "name", ruleName),
				),
			},
			{
				Config: testAccCheckIBMISNetworkACLRuleConfig6Update(vpcName, updatedRuleName),
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
		case "*vpcv1.NetworkACLRuleNetworkACLRuleProtocolAny":
			{
				rulex := foundNwACLRule.(*vpcv1.NetworkACLRuleNetworkACLRuleProtocolAny)
				nwACLRule = makeTerraformACLRuleID(nwACLID, *rulex.ID)

			}
		case "*vpcv1.NetworkACLRuleNetworkACLRuleProtocolIcmptcpudp":
			{
				rulex := foundNwACLRule.(*vpcv1.NetworkACLRuleNetworkACLRuleProtocolIcmptcpudp)
				nwACLRule = makeTerraformACLRuleID(nwACLID, *rulex.ID)

			}
		case "*vpcv1.NetworkACLRuleNetworkACLRuleProtocolIndividual":
			{
				rulex := foundNwACLRule.(*vpcv1.NetworkACLRuleNetworkACLRuleProtocolIndividual)
				nwACLRule = makeTerraformACLRuleID(nwACLID, *rulex.ID)
			}
		case "*vpcv1.NetworkACLRule":
			{
				rulex := foundNwACLRule.(*vpcv1.NetworkACLRule)
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

func testAccCheckIBMISNetworkACLRuleConfig5(vpcName, name string) string {
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
		protocol       = "any"
		}
	`, vpcName, name)
}

func testAccCheckIBMISNetworkACLRuleConfig5Update(vpcName, name string) string {
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

func testAccCheckIBMISNetworkACLRuleConfig6(vpcName, name string) string {
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
		protocol       = "number_99"
		}
	`, vpcName, name)
}

func testAccCheckIBMISNetworkACLRuleConfig6Update(vpcName, name string) string {
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

// TestNetworkACLRule_DeprecatedICMPToFlatMigration tests migration from
// deprecated icmp{} block to new flat struct (protocol + top-level type/code)
func TestNetworkACLRule_DeprecatedICMPToFlatMigration(t *testing.T) {
	var nwACLRule string
	vpcName := fmt.Sprintf("tf-nacl-vpc-migrate-%d", acctest.RandIntRange(10, 100))
	ruleName := fmt.Sprintf("tf-migrate-icmp-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: checkNetworkACLRuleDestroy,
		Steps: []resource.TestStep{
			{
				// Step 1: Create with deprecated icmp{} block
				Config: testAccCheckIBMISNetworkACLRuleDeprecatedICMP(vpcName, ruleName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISNetworkACLRuleExists("ibm_is_network_acl_rule.testacc_nacl", nwACLRule),
					resource.TestCheckResourceAttr(
						"ibm_is_network_acl_rule.testacc_nacl", "name", ruleName),
					resource.TestCheckResourceAttr(
						"ibm_is_network_acl_rule.testacc_nacl", "protocol", "icmp"),
				),
			},
			{
				// Step 2: Migrate to flat struct with protocol = "icmp" and top-level type/code
				Config: testAccCheckIBMISNetworkACLRuleFlatICMP(vpcName, ruleName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISNetworkACLRuleExists("ibm_is_network_acl_rule.testacc_nacl", nwACLRule),
					resource.TestCheckResourceAttr(
						"ibm_is_network_acl_rule.testacc_nacl", "name", ruleName),
					resource.TestCheckResourceAttr(
						"ibm_is_network_acl_rule.testacc_nacl", "protocol", "icmp"),
					resource.TestCheckResourceAttr(
						"ibm_is_network_acl_rule.testacc_nacl", "type", "8"),
					resource.TestCheckResourceAttr(
						"ibm_is_network_acl_rule.testacc_nacl", "code", "0"),
				),
			},
		},
	})
}

// TestNetworkACLRule_DeprecatedTCPToFlatMigration tests migration from
// deprecated tcp{} block to new flat struct (protocol + top-level ports)
func TestNetworkACLRule_DeprecatedTCPToFlatMigration(t *testing.T) {
	var nwACLRule string
	vpcName := fmt.Sprintf("tf-nacl-vpc-tcp-%d", acctest.RandIntRange(10, 100))
	ruleName := fmt.Sprintf("tf-migrate-tcp-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: checkNetworkACLRuleDestroy,
		Steps: []resource.TestStep{
			{
				// Step 1: Create with deprecated tcp{} block
				Config: testAccCheckIBMISNetworkACLRuleDeprecatedTCP(vpcName, ruleName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISNetworkACLRuleExists("ibm_is_network_acl_rule.testacc_nacl", nwACLRule),
					resource.TestCheckResourceAttr(
						"ibm_is_network_acl_rule.testacc_nacl", "name", ruleName),
					resource.TestCheckResourceAttr(
						"ibm_is_network_acl_rule.testacc_nacl", "protocol", "tcp"),
				),
			},
			{
				// Step 2: Migrate to flat struct with protocol = "tcp" and top-level ports
				Config: testAccCheckIBMISNetworkACLRuleFlatTCP(vpcName, ruleName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISNetworkACLRuleExists("ibm_is_network_acl_rule.testacc_nacl", nwACLRule),
					resource.TestCheckResourceAttr(
						"ibm_is_network_acl_rule.testacc_nacl", "name", ruleName),
					resource.TestCheckResourceAttr(
						"ibm_is_network_acl_rule.testacc_nacl", "protocol", "tcp"),
					resource.TestCheckResourceAttr(
						"ibm_is_network_acl_rule.testacc_nacl", "port_min", "443"),
					resource.TestCheckResourceAttr(
						"ibm_is_network_acl_rule.testacc_nacl", "port_max", "443"),
				),
			},
		},
	})
}

// TestNetworkACLRule_DeprecatedUDPToFlatMigration tests migration from
// deprecated udp{} block to new flat struct (protocol + top-level ports)
func TestNetworkACLRule_DeprecatedUDPToFlatMigration(t *testing.T) {
	var nwACLRule string
	vpcName := fmt.Sprintf("tf-nacl-vpc-udp-%d", acctest.RandIntRange(10, 100))
	ruleName := fmt.Sprintf("tf-migrate-udp-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: checkNetworkACLRuleDestroy,
		Steps: []resource.TestStep{
			{
				// Step 1: Create with deprecated udp{} block
				Config: testAccCheckIBMISNetworkACLRuleDeprecatedUDP(vpcName, ruleName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISNetworkACLRuleExists("ibm_is_network_acl_rule.testacc_nacl", nwACLRule),
					resource.TestCheckResourceAttr(
						"ibm_is_network_acl_rule.testacc_nacl", "name", ruleName),
					resource.TestCheckResourceAttr(
						"ibm_is_network_acl_rule.testacc_nacl", "protocol", "udp"),
				),
			},
			{
				// Step 2: Migrate to flat struct with protocol = "udp" and top-level ports
				Config: testAccCheckIBMISNetworkACLRuleFlatUDP(vpcName, ruleName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISNetworkACLRuleExists("ibm_is_network_acl_rule.testacc_nacl", nwACLRule),
					resource.TestCheckResourceAttr(
						"ibm_is_network_acl_rule.testacc_nacl", "name", ruleName),
					resource.TestCheckResourceAttr(
						"ibm_is_network_acl_rule.testacc_nacl", "protocol", "udp"),
					resource.TestCheckResourceAttr(
						"ibm_is_network_acl_rule.testacc_nacl", "port_min", "53"),
					resource.TestCheckResourceAttr(
						"ibm_is_network_acl_rule.testacc_nacl", "port_max", "53"),
				),
			},
		},
	})
}

// TestNetworkACLRule_ICMPZeroValues tests ICMP with type=0, code=0 (Echo Reply)
func TestNetworkACLRule_ICMPZeroValues(t *testing.T) {
	var nwACLRule string
	vpcName := fmt.Sprintf("tf-nacl-vpc-zero-%d", acctest.RandIntRange(10, 100))
	ruleName := fmt.Sprintf("tf-icmp-zero-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: checkNetworkACLRuleDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISNetworkACLRuleICMPZeroValues(vpcName, ruleName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISNetworkACLRuleExists("ibm_is_network_acl_rule.testacc_nacl", nwACLRule),
					resource.TestCheckResourceAttr(
						"ibm_is_network_acl_rule.testacc_nacl", "name", ruleName),
					resource.TestCheckResourceAttr(
						"ibm_is_network_acl_rule.testacc_nacl", "protocol", "icmp"),
					resource.TestCheckResourceAttr(
						"ibm_is_network_acl_rule.testacc_nacl", "type", "0"),
					resource.TestCheckResourceAttr(
						"ibm_is_network_acl_rule.testacc_nacl", "code", "0"),
				),
			},
		},
	})
}

// Config: Deprecated icmp{} block
func testAccCheckIBMISNetworkACLRuleDeprecatedICMP(vpcName, ruleName string) string {
	return fmt.Sprintf(`
    resource "ibm_is_vpc" "testacc_vpc" {
        name = "%s"
    }

    resource "ibm_is_network_acl_rule" "testacc_nacl" {
        network_acl = ibm_is_vpc.testacc_vpc.default_network_acl
        name        = "%s"
        action      = "allow"
        source      = "0.0.0.0/0"
        destination = "0.0.0.0/0"
        direction   = "inbound"
        icmp {
            type = 8
            code = 0
        }
    }
    `, vpcName, ruleName)
}

// Config: Flat struct with protocol = "icmp" and top-level type/code
func testAccCheckIBMISNetworkACLRuleFlatICMP(vpcName, ruleName string) string {
	return fmt.Sprintf(`
    resource "ibm_is_vpc" "testacc_vpc" {
        name = "%s"
    }

    resource "ibm_is_network_acl_rule" "testacc_nacl" {
        network_acl = ibm_is_vpc.testacc_vpc.default_network_acl
        name        = "%s"
        action      = "allow"
        source      = "0.0.0.0/0"
        destination = "0.0.0.0/0"
        direction   = "inbound"
        protocol    = "icmp"
        type        = 8
        code        = 0
    }
    `, vpcName, ruleName)
}

// Config: Deprecated tcp{} block
func testAccCheckIBMISNetworkACLRuleDeprecatedTCP(vpcName, ruleName string) string {
	return fmt.Sprintf(`
    resource "ibm_is_vpc" "testacc_vpc" {
        name = "%s"
    }

    resource "ibm_is_network_acl_rule" "testacc_nacl" {
        network_acl = ibm_is_vpc.testacc_vpc.default_network_acl
        name        = "%s"
        action      = "allow"
        source      = "0.0.0.0/0"
        destination = "0.0.0.0/0"
        direction   = "inbound"
        tcp {
            port_min = 80
            port_max = 80
        }
    }
    `, vpcName, ruleName)
}

// Config: Flat struct with protocol = "tcp" and top-level ports
func testAccCheckIBMISNetworkACLRuleFlatTCP(vpcName, ruleName string) string {
	return fmt.Sprintf(`
    resource "ibm_is_vpc" "testacc_vpc" {
        name = "%s"
    }

    resource "ibm_is_network_acl_rule" "testacc_nacl" {
        network_acl = ibm_is_vpc.testacc_vpc.default_network_acl
        name        = "%s"
        action      = "allow"
        source      = "0.0.0.0/0"
        destination = "0.0.0.0/0"
        direction   = "inbound"
        protocol    = "tcp"
        port_min    = 443
        port_max    = 443
    }
    `, vpcName, ruleName)
}

// Config: Deprecated udp{} block
func testAccCheckIBMISNetworkACLRuleDeprecatedUDP(vpcName, ruleName string) string {
	return fmt.Sprintf(`
    resource "ibm_is_vpc" "testacc_vpc" {
        name = "%s"
    }

    resource "ibm_is_network_acl_rule" "testacc_nacl" {
        network_acl = ibm_is_vpc.testacc_vpc.default_network_acl
        name        = "%s"
        action      = "allow"
        source      = "0.0.0.0/0"
        destination = "0.0.0.0/0"
        direction   = "inbound"
        udp {
            port_min = 53
            port_max = 53
        }
    }
    `, vpcName, ruleName)
}

// Config: Flat struct with protocol = "udp" and top-level ports
func testAccCheckIBMISNetworkACLRuleFlatUDP(vpcName, ruleName string) string {
	return fmt.Sprintf(`
    resource "ibm_is_vpc" "testacc_vpc" {
        name = "%s"
    }

    resource "ibm_is_network_acl_rule" "testacc_nacl" {
        network_acl = ibm_is_vpc.testacc_vpc.default_network_acl
        name        = "%s"
        action      = "allow"
        source      = "0.0.0.0/0"
        destination = "0.0.0.0/0"
        direction   = "inbound"
        protocol    = "udp"
        port_min    = 53
        port_max    = 53
    }
    `, vpcName, ruleName)
}

// Config: ICMP with zero values (type=0, code=0 - Echo Reply)
func testAccCheckIBMISNetworkACLRuleICMPZeroValues(vpcName, ruleName string) string {
	return fmt.Sprintf(`
    resource "ibm_is_vpc" "testacc_vpc" {
        name = "%s"
    }

    resource "ibm_is_network_acl_rule" "testacc_nacl" {
        network_acl = ibm_is_vpc.testacc_vpc.default_network_acl
        name        = "%s"
        action      = "allow"
        source      = "0.0.0.0/0"
        destination = "0.0.0.0/0"
        direction   = "inbound"
        protocol    = "icmp"
        type        = 0
        code        = 0
    }
    `, vpcName, ruleName)
}
