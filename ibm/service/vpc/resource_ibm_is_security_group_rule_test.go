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
		case "*vpcv1.SecurityGroupRuleProtocolAny":
			{
				sgr := foundSecurityGroupRule.(*vpcv1.SecurityGroupRuleProtocolAny)
				securityGroupRuleID = *sgr.ID
			}
		case "*vpcv1.SecurityGroupRuleProtocolIcmptcpudp":
			{
				sgr := foundSecurityGroupRule.(*vpcv1.SecurityGroupRuleProtocolIcmptcpudp)
				securityGroupRuleID = *sgr.ID
			}
		case "*vpcv1.SecurityGroupRuleProtocolIndividual":
			{
				sgr := foundSecurityGroupRule.(*vpcv1.SecurityGroupRuleProtocolIndividual)
				securityGroupRuleID = *sgr.ID
			}
		case "*vpcv1.SecurityGroupRule":
			{
				sgr := foundSecurityGroupRule.(*vpcv1.SecurityGroupRule)
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

	  resource "ibm_is_security_group_rule" "testacc_security_group_rule" {
		group     = ibm_is_security_group.testacc_security_group.id
		direction = "inbound"
		remote    = "127.0.0.1"
		icmp {
		  code = 21
		  type = 31
		}
		name 	  = "test-name"
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

	  resource "ibm_is_security_group_rule" "testacc_security_group_rule_any" {
		depends_on = [ibm_is_security_group_rule.testacc_security_group_rule_tcp]
		group      = ibm_is_security_group.testacc_security_group.id
		direction  = "inbound"
		remote     = "127.0.0.1"
		protocol   = "any"
	  }
	  resource "ibm_is_security_group_rule" "testacc_security_group_rule_icmp_tcp_udp" {
		depends_on = [ibm_is_security_group_rule.testacc_security_group_rule_tcp]
		group      = ibm_is_security_group.testacc_security_group.id
		direction  = "inbound"
		remote     = "127.0.0.1"
		protocol   = "icmp_tcp_udp"
	  }
	  resource "ibm_is_security_group_rule" "testacc_security_group_rule_individual" {
		depends_on = [ibm_is_security_group_rule.testacc_security_group_rule_tcp]
		group      = ibm_is_security_group.testacc_security_group.id
		direction  = "inbound"
		remote     = "127.0.0.1"
		protocol   = "number_99"
	  }
	  resource "ibm_is_security_group_rule" "testacc_security_group_rule_icmp_new" {
		depends_on = [ibm_is_security_group_rule.testacc_security_group_rule_tcp]
		group      = ibm_is_security_group.testacc_security_group.id
		direction  = "inbound"
		remote     = "127.0.0.1"
		protocol   = "icmp"
		code = 20
		type = 30
	  }	
	  resource "ibm_is_security_group_rule" "testacc_security_group_rule_tcp_new" {
		depends_on = [ibm_is_security_group_rule.testacc_security_group_rule_tcp]
		group      = ibm_is_security_group.testacc_security_group.id
		direction  = "inbound"
		remote     = "127.0.0.1"
		protocol   = "tcp"
		port_min = 8080
		port_max = 8080
	  }		
	  resource "ibm_is_security_group_rule" "testacc_security_group_rule_udp_new" {
		depends_on = [ibm_is_security_group_rule.testacc_security_group_rule_tcp]
		group      = ibm_is_security_group.testacc_security_group.id
		direction  = "inbound"
		remote     = "127.0.0.1"
		protocol   = "udp"
		port_min = 8080
		port_max = 8080
	  }

 `, vpcname, name)

}

// TestAccIBMISSecurityGroupRule_DeprecatedToFlatMigration tests migration from
// deprecated blocks (icmp{}, tcp{}, udp{}) to new flat struct (protocol + top-level attributes)
func TestAccIBMISSecurityGroupRule_DeprecatedICMPToFlatMigration(t *testing.T) {
	var securityGroupRule string

	vpcname := fmt.Sprintf("tfsgrule-vpc-migrate-%d", acctest.RandIntRange(10, 100))
	sgname := fmt.Sprintf("tfsgrule-sg-migrate-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISSecurityGroupRuleDestroy,
		Steps: []resource.TestStep{
			{
				// Step 1: Create with deprecated icmp{} block
				Config: testAccCheckIBMISSecurityGroupRuleDeprecatedICMP(vpcname, sgname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISSecurityGroupRuleExists("ibm_is_security_group_rule.test_icmp", securityGroupRule),
					resource.TestCheckResourceAttr(
						"ibm_is_security_group_rule.test_icmp", "protocol", "icmp"),
				),
			},
			{
				// Step 2: Migrate to flat struct with protocol = "icmp" and top-level type/code
				Config: testAccCheckIBMISSecurityGroupRuleFlatICMP(vpcname, sgname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISSecurityGroupRuleExists("ibm_is_security_group_rule.test_icmp", securityGroupRule),
					resource.TestCheckResourceAttr(
						"ibm_is_security_group_rule.test_icmp", "protocol", "icmp"),
					resource.TestCheckResourceAttr(
						"ibm_is_security_group_rule.test_icmp", "type", "8"),
					resource.TestCheckResourceAttr(
						"ibm_is_security_group_rule.test_icmp", "code", "0"),
				),
			},
		},
	})
}

func TestAccIBMISSecurityGroupRule_DeprecatedTCPToFlatMigration(t *testing.T) {
	var securityGroupRule string

	vpcname := fmt.Sprintf("tfsgrule-vpc-tcp-%d", acctest.RandIntRange(10, 100))
	sgname := fmt.Sprintf("tfsgrule-sg-tcp-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISSecurityGroupRuleDestroy,
		Steps: []resource.TestStep{
			{
				// Step 1: Create with deprecated tcp{} block
				Config: testAccCheckIBMISSecurityGroupRuleDeprecatedTCP(vpcname, sgname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISSecurityGroupRuleExists("ibm_is_security_group_rule.test_tcp", securityGroupRule),
					resource.TestCheckResourceAttr(
						"ibm_is_security_group_rule.test_tcp", "protocol", "tcp"),
				),
			},
			{
				// Step 2: Migrate to flat struct with protocol = "tcp" and top-level ports
				Config: testAccCheckIBMISSecurityGroupRuleFlatTCP(vpcname, sgname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISSecurityGroupRuleExists("ibm_is_security_group_rule.test_tcp", securityGroupRule),
					resource.TestCheckResourceAttr(
						"ibm_is_security_group_rule.test_tcp", "protocol", "tcp"),
					resource.TestCheckResourceAttr(
						"ibm_is_security_group_rule.test_tcp", "port_min", "443"),
					resource.TestCheckResourceAttr(
						"ibm_is_security_group_rule.test_tcp", "port_max", "443"),
				),
			},
		},
	})
}

func TestAccIBMISSecurityGroupRule_DeprecatedUDPToFlatMigration(t *testing.T) {
	var securityGroupRule string

	vpcname := fmt.Sprintf("tfsgrule-vpc-udp-%d", acctest.RandIntRange(10, 100))
	sgname := fmt.Sprintf("tfsgrule-sg-udp-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISSecurityGroupRuleDestroy,
		Steps: []resource.TestStep{
			{
				// Step 1: Create with deprecated udp{} block
				Config: testAccCheckIBMISSecurityGroupRuleDeprecatedUDP(vpcname, sgname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISSecurityGroupRuleExists("ibm_is_security_group_rule.test_udp", securityGroupRule),
					resource.TestCheckResourceAttr(
						"ibm_is_security_group_rule.test_udp", "protocol", "udp"),
				),
			},
			{
				// Step 2: Migrate to flat struct with protocol = "udp" and top-level ports
				Config: testAccCheckIBMISSecurityGroupRuleFlatUDP(vpcname, sgname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISSecurityGroupRuleExists("ibm_is_security_group_rule.test_udp", securityGroupRule),
					resource.TestCheckResourceAttr(
						"ibm_is_security_group_rule.test_udp", "protocol", "udp"),
					resource.TestCheckResourceAttr(
						"ibm_is_security_group_rule.test_udp", "port_min", "53"),
					resource.TestCheckResourceAttr(
						"ibm_is_security_group_rule.test_udp", "port_max", "53"),
				),
			},
		},
	})
}

// Test ICMP with zero values (type=0, code=0 - Echo Reply)
func TestAccIBMISSecurityGroupRule_ICMPZeroValues(t *testing.T) {
	var securityGroupRule string

	vpcname := fmt.Sprintf("tfsgrule-vpc-zero-%d", acctest.RandIntRange(10, 100))
	sgname := fmt.Sprintf("tfsgrule-sg-zero-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISSecurityGroupRuleDestroy,
		Steps: []resource.TestStep{
			{
				// Create with ICMP type=0 (Echo Reply) and code=0
				Config: testAccCheckIBMISSecurityGroupRuleICMPZeroValues(vpcname, sgname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISSecurityGroupRuleExists("ibm_is_security_group_rule.test_icmp_zero", securityGroupRule),
					resource.TestCheckResourceAttr(
						"ibm_is_security_group_rule.test_icmp_zero", "protocol", "icmp"),
					resource.TestCheckResourceAttr(
						"ibm_is_security_group_rule.test_icmp_zero", "type", "0"),
					resource.TestCheckResourceAttr(
						"ibm_is_security_group_rule.test_icmp_zero", "code", "0"),
				),
			},
		},
	})
}

// Config: Deprecated icmp{} block
func testAccCheckIBMISSecurityGroupRuleDeprecatedICMP(vpcname, sgname string) string {
	return fmt.Sprintf(`
    resource "ibm_is_vpc" "testacc_vpc" {
        name = "%s"
    }

    resource "ibm_is_security_group" "testacc_security_group" {
        name = "%s"
        vpc  = ibm_is_vpc.testacc_vpc.id
    }

    resource "ibm_is_security_group_rule" "test_icmp" {
        group     = ibm_is_security_group.testacc_security_group.id
        direction = "inbound"
        remote    = "0.0.0.0/0"
        icmp {
            type = 8
            code = 0
        }
    }
    `, vpcname, sgname)
}

// Config: Flat struct with protocol = "icmp" and top-level type/code
func testAccCheckIBMISSecurityGroupRuleFlatICMP(vpcname, sgname string) string {
	return fmt.Sprintf(`
    resource "ibm_is_vpc" "testacc_vpc" {
        name = "%s"
    }

    resource "ibm_is_security_group" "testacc_security_group" {
        name = "%s"
        vpc  = ibm_is_vpc.testacc_vpc.id
    }

    resource "ibm_is_security_group_rule" "test_icmp" {
        group     = ibm_is_security_group.testacc_security_group.id
        direction = "inbound"
        remote    = "0.0.0.0/0"
        protocol  = "icmp"
        type      = 8
        code      = 0
    }
    `, vpcname, sgname)
}

// Config: Deprecated tcp{} block
func testAccCheckIBMISSecurityGroupRuleDeprecatedTCP(vpcname, sgname string) string {
	return fmt.Sprintf(`
    resource "ibm_is_vpc" "testacc_vpc" {
        name = "%s"
    }

    resource "ibm_is_security_group" "testacc_security_group" {
        name = "%s"
        vpc  = ibm_is_vpc.testacc_vpc.id
    }

    resource "ibm_is_security_group_rule" "test_tcp" {
        group     = ibm_is_security_group.testacc_security_group.id
        direction = "inbound"
        remote    = "0.0.0.0/0"
        tcp {
            port_min = 80
            port_max = 80
        }
    }
    `, vpcname, sgname)
}

// Config: Flat struct with protocol = "tcp" and top-level ports
func testAccCheckIBMISSecurityGroupRuleFlatTCP(vpcname, sgname string) string {
	return fmt.Sprintf(`
    resource "ibm_is_vpc" "testacc_vpc" {
        name = "%s"
    }

    resource "ibm_is_security_group" "testacc_security_group" {
        name = "%s"
        vpc  = ibm_is_vpc.testacc_vpc.id
    }

    resource "ibm_is_security_group_rule" "test_tcp" {
        group     = ibm_is_security_group.testacc_security_group.id
        direction = "inbound"
        remote    = "0.0.0.0/0"
        protocol  = "tcp"
        port_min  = 443
        port_max  = 443
    }
    `, vpcname, sgname)
}

// Config: Deprecated udp{} block
func testAccCheckIBMISSecurityGroupRuleDeprecatedUDP(vpcname, sgname string) string {
	return fmt.Sprintf(`
    resource "ibm_is_vpc" "testacc_vpc" {
        name = "%s"
    }

    resource "ibm_is_security_group" "testacc_security_group" {
        name = "%s"
        vpc  = ibm_is_vpc.testacc_vpc.id
    }

    resource "ibm_is_security_group_rule" "test_udp" {
        group     = ibm_is_security_group.testacc_security_group.id
        direction = "inbound"
        remote    = "0.0.0.0/0"
        udp {
            port_min = 53
            port_max = 53
        }
    }
    `, vpcname, sgname)
}

// Config: Flat struct with protocol = "udp" and top-level ports
func testAccCheckIBMISSecurityGroupRuleFlatUDP(vpcname, sgname string) string {
	return fmt.Sprintf(`
    resource "ibm_is_vpc" "testacc_vpc" {
        name = "%s"
    }

    resource "ibm_is_security_group" "testacc_security_group" {
        name = "%s"
        vpc  = ibm_is_vpc.testacc_vpc.id
    }

    resource "ibm_is_security_group_rule" "test_udp" {
        group     = ibm_is_security_group.testacc_security_group.id
        direction = "inbound"
        remote    = "0.0.0.0/0"
        protocol  = "udp"
        port_min  = 53
        port_max  = 53
    }
    `, vpcname, sgname)
}

// Config: ICMP with zero values (type=0, code=0 - Echo Reply)
func testAccCheckIBMISSecurityGroupRuleICMPZeroValues(vpcname, sgname string) string {
	return fmt.Sprintf(`
    resource "ibm_is_vpc" "testacc_vpc" {
        name = "%s"
    }

    resource "ibm_is_security_group" "testacc_security_group" {
        name = "%s"
        vpc  = ibm_is_vpc.testacc_vpc.id
    }

    resource "ibm_is_security_group_rule" "test_icmp_zero" {
        group     = ibm_is_security_group.testacc_security_group.id
        direction = "inbound"
        remote    = "0.0.0.0/0"
        protocol  = "icmp"
        type      = 0
        code      = 0
    }
    `, vpcname, sgname)
}
