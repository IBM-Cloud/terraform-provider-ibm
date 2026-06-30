// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"errors"
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"

	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestNetworkACLGen1(t *testing.T) {
	var nwACL string
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: checkNetworkACLDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISNetworkACLConfig(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISNetworkACLExists("ibm_is_network_acl.isExampleACL", nwACL),
					resource.TestCheckResourceAttr(
						"ibm_is_network_acl.isExampleACL", "name", "is-example-acl"),
					resource.TestCheckResourceAttr(
						"ibm_is_network_acl.isExampleACL", "rules.#", "2"),
				),
			},
		},
	})
}

func TestNetworkACLResourceGroupUpdate(t *testing.T) {
	var nwACL string
	setResourceGroup := false
	setResourceGroup1 := true
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: checkNetworkACLDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISNetworkACLResourceGroupConfig(setResourceGroup),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISNetworkACLExists("ibm_is_network_acl.isExampleACL", nwACL),
					resource.TestCheckResourceAttr(
						"ibm_is_vpc.testacc_vpc", "name", "tf-nwacl-vpc"),
					resource.TestCheckResourceAttr(
						"ibm_is_subnet.testacc_subnet", "name", "tf-nwacl-subnet"),
					resource.TestCheckResourceAttr(
						"ibm_is_network_acl.isExampleACL", "rules.#", "5"),
					resource.TestCheckResourceAttr(
						"ibm_is_network_acl.isExampleACL", "tags.#", "2"),
					resource.TestCheckResourceAttr(
						"ibm_is_network_acl.isExampleACL", "resource_group_name", "Default"),
				),
			},
			{
				Config: testAccCheckIBMISNetworkACLResourceGroupConfig(setResourceGroup1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISNetworkACLExists("ibm_is_network_acl.isExampleACL", nwACL),
					resource.TestCheckResourceAttr(
						"ibm_is_vpc.testacc_vpc", "name", "tf-nwacl-vpc"),
					resource.TestCheckResourceAttr(
						"ibm_is_subnet.testacc_subnet", "name", "tf-nwacl-subnet"),
					resource.TestCheckResourceAttr(
						"ibm_is_network_acl.isExampleACL", "rules.#", "5"),
					resource.TestCheckResourceAttr(
						"ibm_is_network_acl.isExampleACL", "tags.#", "2"),
					resource.TestCheckResourceAttrWith("ibm_is_network_acl.isExampleACL", "resource_group_name", func(v string) error {
						if v == "Default" {
							return fmt.Errorf("Attribute 'resource_group' is still Default")
						}
						return nil
					}),
				),
			},
		},
	})
}
func TestNetworkACLGen2(t *testing.T) {
	var nwACL string
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: checkNetworkACLDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISNetworkACLConfig1(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISNetworkACLExists("ibm_is_network_acl.isExampleACL", nwACL),
					resource.TestCheckResourceAttr(
						"ibm_is_network_acl.isExampleACL", "name", "is-example-acl"),
					resource.TestCheckResourceAttr(
						"ibm_is_network_acl.isExampleACL", "rules.#", "6"),
					resource.TestCheckResourceAttr(
						"ibm_is_network_acl.isExampleACL", "tags.#", "2"),
				),
			},
		},
	})
}

func checkNetworkACLDestroy(s *terraform.State) error {
	sess, _ := acc.TestAccProvider.Meta().(conns.ClientSession).VpcV1API()
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_is_network_acl" {
			continue
		}

		getnwacloptions := &vpcv1.GetNetworkACLOptions{
			ID: &rs.Primary.ID,
		}
		_, _, err := sess.GetNetworkACL(getnwacloptions)
		if err == nil {
			return fmt.Errorf("network acl still exists: %s", rs.Primary.ID)
		}
	}
	return nil
}

func testAccCheckIBMISNetworkACLExists(n, nwACL string) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return errors.New("No Record ID is set")
		}
		sess, _ := acc.TestAccProvider.Meta().(conns.ClientSession).VpcV1API()
		getnwacloptions := &vpcv1.GetNetworkACLOptions{
			ID: &rs.Primary.ID,
		}
		foundNwACL, _, err := sess.GetNetworkACL(getnwacloptions)
		if err != nil {
			return err
		}
		nwACL = *foundNwACL.ID
		return nil
	}
}

func testAccCheckIBMISNetworkACLConfig() string {
	return fmt.Sprintf(`
	resource "ibm_is_vpc" "testacc_vpc" {
		name = "tf-nwacl-vpc"
	  }

	resource "ibm_is_network_acl" "isExampleACL" {
		name = "is-example-acl"
		vpc  = ibm_is_vpc.testacc_vpc.id
		rules {
		  name        = "outbound"
		  action      = "allow"
		  source      = "0.0.0.0/0"
		  destination = "0.0.0.0/0"
		  direction   = "outbound"
		  icmp {
			code = 8
			type = 1
		  }
		  # Optionals :
		  # port_max =
		  # port_min =
		}
		rules {
		  name        = "inbound"
		  action      = "allow"
		  source      = "0.0.0.0/0"
		  destination = "0.0.0.0/0"
		  direction   = "inbound"
		  icmp {
			code = 8
			type = 1
		  }
		  # Optionals :
		  # port_max =
		  # port_min =
		}
	  }
	`)
}

func testAccCheckIBMISNetworkACLConfig1() string {
	return fmt.Sprintf(`
	resource "ibm_is_vpc" "testacc_vpc" {
		name = "tf-nwacl-vpc"
	  }

	resource "ibm_is_network_acl" "isExampleACL" {
		name = "is-example-acl"
		tags = ["Tag1", "tag2"]
		vpc  = ibm_is_vpc.testacc_vpc.id
		rules {
		  name        = "outbound"
		  action      = "allow"
		  source      = "0.0.0.0/0"
		  destination = "0.0.0.0/0"
		  direction   = "outbound"
		  icmp {
			code = 8
			type = 1
		  }
		  # Optionals :
		  # port_max =
		  # port_min =
		}
		rules {
		  name        = "inbound"
		  action      = "allow"
		  source      = "0.0.0.0/0"
		  destination = "0.0.0.0/0"
		  direction   = "inbound"
		  icmp {
			code = 8
			type = 1
		  }
		  # Optionals :
		  # port_max =
		  # port_min =
		}
		rules {
		  name        = "icmnew"
		  action      = "allow"
		  source      = "0.0.0.0/0"
		  destination = "0.0.0.0/0"
		  direction   = "inbound"
		  protocol    = "icmp"
		  code 		  = 8
		  type 		  = 1
		}
		rules {
		  name        = "anyprotocol"
		  action      = "allow"
		  source      = "0.0.0.0/0"
		  destination = "0.0.0.0/0"
		  direction   = "inbound"
		  protocol    = "any"
		} 
		
		rules {
		  name        = "icmptcpudp"
		  action      = "allow"
		  source      = "0.0.0.0/0"
		  destination = "0.0.0.0/0"
		  direction   = "inbound"
		  protocol    = "icmp_tcp_udp"
		}
		rules {
		  name        = "individual"
		  action      = "allow"
		  source      = "0.0.0.0/0"
		  destination = "0.0.0.0/0"
		  direction   = "inbound"
		  protocol    = "number_99"
		}
	  }
	`)
}
func testAccCheckIBMISNetworkACLResourceGroupConfig(resourceGroupSelect bool) string {
	return fmt.Sprintf(`
	resource "ibm_is_vpc" "testacc_vpc" {
		name = "tf-nwacl-vpc"
	}

	resource "ibm_is_subnet" "testacc_subnet" {
		name           				 	= "tf-nwacl-subnet"
		vpc             				= ibm_is_vpc.testacc_vpc.id
		zone            				= "%s"
		total_ipv4_address_count 		= 16
		network_acl     				= ibm_is_network_acl.isExampleACL.id
	}

	resource "ibm_is_network_acl" "isExampleACL" {
		tags = ["Tag1", "tag2"]
		vpc  = ibm_is_vpc.testacc_vpc.id
		resource_group  				= %t ? "%s" : null
		rules {
		  name        = "outbound"
		  action      = "allow"
		  source      = "0.0.0.0/0"
		  destination = "0.0.0.0/0"
		  direction   = "outbound"
		  icmp {
			code = 8
			type = 1
		  }
		}
		rules {
		  name        = "inbound"
		  action      = "allow"
		  source      = "0.0.0.0/0"
		  destination = "0.0.0.0/0"
		  direction   = "inbound"
		  icmp {
			code = 8
			type = 1
		  }
		}

		lifecycle {
			create_before_destroy = true
		}
	  }
	`, acc.ISZoneName, resourceGroupSelect, acc.IsResourceGroupID)
}

// TestNetworkACL_DeprecatedICMPToFlatMigration tests migration of inline rules from
// TestNetworkACL_DeprecatedToFlatMigration tests migration of inline rules from
// deprecated blocks (icmp{}, tcp{}, udp{}) to new flat struct (protocol + top-level attributes)
// and also tests switching between protocol blocks (icmp{} to tcp{})
func TestNetworkACL_DeprecatedToFlatMigration(t *testing.T) {
	var nwACL string
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: checkNetworkACLDestroy,
		Steps: []resource.TestStep{
			{
				// Step 1: Create with deprecated icmp{} block
				Config: testAccCheckIBMISNetworkACLDeprecatedICMP(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISNetworkACLExists("ibm_is_network_acl.test_acl", nwACL),
					resource.TestCheckResourceAttr(
						"ibm_is_network_acl.test_acl", "rules.#", "1"),
					resource.TestCheckResourceAttr(
						"ibm_is_network_acl.test_acl", "rules.0.protocol", "icmp"),
				),
			},
			{
				// Step 2: Migrate icmp{} to flat struct with protocol = "icmp"
				Config: testAccCheckIBMISNetworkACLFlatICMP(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISNetworkACLExists("ibm_is_network_acl.test_acl", nwACL),
					resource.TestCheckResourceAttr(
						"ibm_is_network_acl.test_acl", "rules.#", "1"),
					resource.TestCheckResourceAttr(
						"ibm_is_network_acl.test_acl", "rules.0.protocol", "icmp"),
				),
			},
			{
				// Step 3: Switch from icmp to tcp{} block (tests protocol switch)
				Config: testAccCheckIBMISNetworkACLDeprecatedTCP(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISNetworkACLExists("ibm_is_network_acl.test_acl", nwACL),
					resource.TestCheckResourceAttr(
						"ibm_is_network_acl.test_acl", "rules.#", "1"),
					resource.TestCheckResourceAttr(
						"ibm_is_network_acl.test_acl", "rules.0.protocol", "tcp"),
				),
			},
			{
				// Step 4: Migrate tcp{} to flat struct with protocol = "tcp"
				Config: testAccCheckIBMISNetworkACLFlatTCP(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISNetworkACLExists("ibm_is_network_acl.test_acl", nwACL),
					resource.TestCheckResourceAttr(
						"ibm_is_network_acl.test_acl", "rules.#", "1"),
					resource.TestCheckResourceAttr(
						"ibm_is_network_acl.test_acl", "rules.0.protocol", "tcp"),
				),
			},
			{
				// Step 5: Switch from tcp to udp{} block
				Config: testAccCheckIBMISNetworkACLDeprecatedUDP(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISNetworkACLExists("ibm_is_network_acl.test_acl", nwACL),
					resource.TestCheckResourceAttr(
						"ibm_is_network_acl.test_acl", "rules.#", "1"),
					resource.TestCheckResourceAttr(
						"ibm_is_network_acl.test_acl", "rules.0.protocol", "udp"),
				),
			},
			{
				// Step 6: Migrate udp{} to flat struct with protocol = "udp"
				Config: testAccCheckIBMISNetworkACLFlatUDP(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISNetworkACLExists("ibm_is_network_acl.test_acl", nwACL),
					resource.TestCheckResourceAttr(
						"ibm_is_network_acl.test_acl", "rules.#", "1"),
					resource.TestCheckResourceAttr(
						"ibm_is_network_acl.test_acl", "rules.0.protocol", "udp"),
				),
			},
		},
	})
}

// Config: Deprecated icmp{} block in inline rules
func testAccCheckIBMISNetworkACLDeprecatedICMP() string {
	return `
    resource "ibm_is_vpc" "testacc_vpc" {
        name = "tf-nwacl-migrate-vpc"
    }

    resource "ibm_is_network_acl" "test_acl" {
        name = "tf-nwacl-migrate"
        vpc  = ibm_is_vpc.testacc_vpc.id
        rules {
            name        = "test-rule"
            action      = "allow"
            source      = "0.0.0.0/0"
            destination = "0.0.0.0/0"
            direction   = "inbound"
            icmp {
                type = 8
                code = 0
            }
        }
    }
    `
}

// Config: Flat struct with protocol = "icmp" in inline rules
func testAccCheckIBMISNetworkACLFlatICMP() string {
	return `
    resource "ibm_is_vpc" "testacc_vpc" {
        name = "tf-nwacl-migrate-vpc"
    }

    resource "ibm_is_network_acl" "test_acl" {
        name = "tf-nwacl-migrate"
        vpc  = ibm_is_vpc.testacc_vpc.id
        rules {
            name        = "test-rule"
            action      = "allow"
            source      = "0.0.0.0/0"
            destination = "0.0.0.0/0"
            direction   = "inbound"
            protocol    = "icmp"
            type        = 8
            code        = 0
        }
    }
    `
}

// Config: Deprecated tcp{} block in inline rules
func testAccCheckIBMISNetworkACLDeprecatedTCP() string {
	return `
    resource "ibm_is_vpc" "testacc_vpc" {
        name = "tf-nwacl-migrate-vpc"
    }

    resource "ibm_is_network_acl" "test_acl" {
        name = "tf-nwacl-migrate"
        vpc  = ibm_is_vpc.testacc_vpc.id
        rules {
            name        = "test-rule"
            action      = "allow"
            source      = "0.0.0.0/0"
            destination = "0.0.0.0/0"
            direction   = "inbound"
            tcp {
                port_min = 80
                port_max = 80
            }
        }
    }
    `
}

// Config: Flat struct with protocol = "tcp" in inline rules
func testAccCheckIBMISNetworkACLFlatTCP() string {
	return `
    resource "ibm_is_vpc" "testacc_vpc" {
        name = "tf-nwacl-migrate-vpc"
    }

    resource "ibm_is_network_acl" "test_acl" {
        name = "tf-nwacl-migrate"
        vpc  = ibm_is_vpc.testacc_vpc.id
        rules {
            name        = "test-rule"
            action      = "allow"
            source      = "0.0.0.0/0"
            destination = "0.0.0.0/0"
            direction   = "inbound"
            protocol    = "tcp"
            port_min    = 443
            port_max    = 443
        }
    }
    `
}

// Config: Deprecated udp{} block in inline rules
func testAccCheckIBMISNetworkACLDeprecatedUDP() string {
	return `
    resource "ibm_is_vpc" "testacc_vpc" {
        name = "tf-nwacl-migrate-vpc"
    }

    resource "ibm_is_network_acl" "test_acl" {
        name = "tf-nwacl-migrate"
        vpc  = ibm_is_vpc.testacc_vpc.id
        rules {
            name        = "test-rule"
            action      = "allow"
            source      = "0.0.0.0/0"
            destination = "0.0.0.0/0"
            direction   = "inbound"
            udp {
                port_min = 53
                port_max = 53
            }
        }
    }
    `
}

// Config: Flat struct with protocol = "udp" in inline rules
func testAccCheckIBMISNetworkACLFlatUDP() string {
	return `
    resource "ibm_is_vpc" "testacc_vpc" {
        name = "tf-nwacl-migrate-vpc"
    }

    resource "ibm_is_network_acl" "test_acl" {
        name = "tf-nwacl-migrate"
        vpc  = ibm_is_vpc.testacc_vpc.id
        rules {
            name        = "test-rule"
            action      = "allow"
            source      = "0.0.0.0/0"
            destination = "0.0.0.0/0"
            direction   = "inbound"
            protocol    = "udp"
            port_min    = 53
            port_max    = 53
        }
    }
    `
}

// ---------------------------------------------------------------------------
// TestNetworkACL_InlineRuleUpdate* — cover the clear-all + recreate-all path
// that fires on any inline rule change (field edit, add, remove, reorder,
// protocol change).
// ---------------------------------------------------------------------------

// TestNetworkACL_InlineRuleFieldEdit verifies that editing a single field
// (action, source, destination) on an existing rule triggers clear+recreate
// and the new values are reflected in state.
func TestNetworkACL_InlineRuleFieldEdit(t *testing.T) {
	var nwACL string
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: checkNetworkACLDestroy,
		Steps: []resource.TestStep{
			{
				// Step 1: create with action=allow, source=0.0.0.0/0
				Config: testAccCheckIBMISNetworkACLInlineUpdateBase(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISNetworkACLExists("ibm_is_network_acl.test_update_acl", nwACL),
					resource.TestCheckResourceAttr(
						"ibm_is_network_acl.test_update_acl", "rules.#", "1"),
					resource.TestCheckResourceAttr(
						"ibm_is_network_acl.test_update_acl", "rules.0.action", "allow"),
					resource.TestCheckResourceAttr(
						"ibm_is_network_acl.test_update_acl", "rules.0.source", "0.0.0.0/0"),
				),
			},
			{
				// Step 2: change action to deny and narrow source — triggers clear+recreate
				Config: testAccCheckIBMISNetworkACLInlineUpdateFieldEdit(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISNetworkACLExists("ibm_is_network_acl.test_update_acl", nwACL),
					resource.TestCheckResourceAttr(
						"ibm_is_network_acl.test_update_acl", "rules.#", "1"),
					resource.TestCheckResourceAttr(
						"ibm_is_network_acl.test_update_acl", "rules.0.action", "deny"),
					resource.TestCheckResourceAttr(
						"ibm_is_network_acl.test_update_acl", "rules.0.source", "10.0.0.0/8"),
				),
			},
		},
	})
}

// TestNetworkACL_InlineRuleAdd verifies that adding a new rule triggers
// clear-all + recreate-all and both rules appear in state.
func TestNetworkACL_InlineRuleAdd(t *testing.T) {
	var nwACL string
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: checkNetworkACLDestroy,
		Steps: []resource.TestStep{
			{
				// Step 1: one rule
				Config: testAccCheckIBMISNetworkACLInlineUpdateBase(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISNetworkACLExists("ibm_is_network_acl.test_update_acl", nwACL),
					resource.TestCheckResourceAttr(
						"ibm_is_network_acl.test_update_acl", "rules.#", "1"),
				),
			},
			{
				// Step 2: add a second rule — triggers clear+recreate
				Config: testAccCheckIBMISNetworkACLInlineUpdateAddRule(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISNetworkACLExists("ibm_is_network_acl.test_update_acl", nwACL),
					resource.TestCheckResourceAttr(
						"ibm_is_network_acl.test_update_acl", "rules.#", "2"),
					resource.TestCheckResourceAttr(
						"ibm_is_network_acl.test_update_acl", "rules.0.name", "rule-one"),
					resource.TestCheckResourceAttr(
						"ibm_is_network_acl.test_update_acl", "rules.1.name", "rule-two"),
				),
			},
		},
	})
}

// TestNetworkACL_InlineRuleRemove verifies that removing a rule triggers
// clear-all + recreate-all and only the remaining rule is in state.
func TestNetworkACL_InlineRuleRemove(t *testing.T) {
	var nwACL string
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: checkNetworkACLDestroy,
		Steps: []resource.TestStep{
			{
				// Step 1: two rules
				Config: testAccCheckIBMISNetworkACLInlineUpdateAddRule(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISNetworkACLExists("ibm_is_network_acl.test_update_acl", nwACL),
					resource.TestCheckResourceAttr(
						"ibm_is_network_acl.test_update_acl", "rules.#", "2"),
				),
			},
			{
				// Step 2: remove rule-two — triggers clear+recreate, only rule-one remains
				Config: testAccCheckIBMISNetworkACLInlineUpdateBase(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISNetworkACLExists("ibm_is_network_acl.test_update_acl", nwACL),
					resource.TestCheckResourceAttr(
						"ibm_is_network_acl.test_update_acl", "rules.#", "1"),
					resource.TestCheckResourceAttr(
						"ibm_is_network_acl.test_update_acl", "rules.0.name", "rule-one"),
				),
			},
		},
	})
}

// TestNetworkACL_InlineRuleReorder verifies that reordering rules triggers
// clear-all + recreate-all and the new order is reflected in state.
func TestNetworkACL_InlineRuleReorder(t *testing.T) {
	var nwACL string
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: checkNetworkACLDestroy,
		Steps: []resource.TestStep{
			{
				// Step 1: rule-one first, rule-two second
				Config: testAccCheckIBMISNetworkACLInlineUpdateAddRule(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISNetworkACLExists("ibm_is_network_acl.test_update_acl", nwACL),
					resource.TestCheckResourceAttr(
						"ibm_is_network_acl.test_update_acl", "rules.0.name", "rule-one"),
					resource.TestCheckResourceAttr(
						"ibm_is_network_acl.test_update_acl", "rules.1.name", "rule-two"),
				),
			},
			{
				// Step 2: swap order — triggers clear+recreate
				Config: testAccCheckIBMISNetworkACLInlineUpdateReorder(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISNetworkACLExists("ibm_is_network_acl.test_update_acl", nwACL),
					resource.TestCheckResourceAttr(
						"ibm_is_network_acl.test_update_acl", "rules.0.name", "rule-two"),
					resource.TestCheckResourceAttr(
						"ibm_is_network_acl.test_update_acl", "rules.1.name", "rule-one"),
				),
			},
		},
	})
}

// TestNetworkACL_InlineRuleProtocolChange verifies that changing the protocol
// of a rule (tcp → udp) triggers clear-all + recreate-all.
func TestNetworkACL_InlineRuleProtocolChange(t *testing.T) {
	var nwACL string
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: checkNetworkACLDestroy,
		Steps: []resource.TestStep{
			{
				// Step 1: tcp rule
				Config: testAccCheckIBMISNetworkACLInlineUpdateTCP(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISNetworkACLExists("ibm_is_network_acl.test_update_acl", nwACL),
					resource.TestCheckResourceAttr(
						"ibm_is_network_acl.test_update_acl", "rules.#", "1"),
					resource.TestCheckResourceAttr(
						"ibm_is_network_acl.test_update_acl", "rules.0.protocol", "tcp"),
				),
			},
			{
				// Step 2: change protocol to udp — triggers clear+recreate
				Config: testAccCheckIBMISNetworkACLInlineUpdateUDP(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISNetworkACLExists("ibm_is_network_acl.test_update_acl", nwACL),
					resource.TestCheckResourceAttr(
						"ibm_is_network_acl.test_update_acl", "rules.#", "1"),
					resource.TestCheckResourceAttr(
						"ibm_is_network_acl.test_update_acl", "rules.0.protocol", "udp"),
				),
			},
			{
				// Step 3: change protocol to icmp — triggers clear+recreate
				Config: testAccCheckIBMISNetworkACLInlineUpdateICMP(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISNetworkACLExists("ibm_is_network_acl.test_update_acl", nwACL),
					resource.TestCheckResourceAttr(
						"ibm_is_network_acl.test_update_acl", "rules.#", "1"),
					resource.TestCheckResourceAttr(
						"ibm_is_network_acl.test_update_acl", "rules.0.protocol", "icmp"),
				),
			},
		},
	})
}

// ---------------------------------------------------------------------------
// Config helpers for inline-rule update tests
// ---------------------------------------------------------------------------

func testAccCheckIBMISNetworkACLInlineUpdateBase() string {
	return `
    resource "ibm_is_vpc" "testacc_vpc" {
        name = "tf-nwacl-update-vpc"
    }
    resource "ibm_is_network_acl" "test_update_acl" {
        name = "tf-nwacl-update"
        vpc  = ibm_is_vpc.testacc_vpc.id
        rules {
            name        = "rule-one"
            action      = "allow"
            source      = "0.0.0.0/0"
            destination = "0.0.0.0/0"
            direction   = "inbound"
            protocol    = "any"
        }
    }
    `
}

func testAccCheckIBMISNetworkACLInlineUpdateFieldEdit() string {
	return `
    resource "ibm_is_vpc" "testacc_vpc" {
        name = "tf-nwacl-update-vpc"
    }
    resource "ibm_is_network_acl" "test_update_acl" {
        name = "tf-nwacl-update"
        vpc  = ibm_is_vpc.testacc_vpc.id
        rules {
            name        = "rule-one"
            action      = "deny"
            source      = "10.0.0.0/8"
            destination = "0.0.0.0/0"
            direction   = "inbound"
            protocol    = "any"
        }
    }
    `
}

func testAccCheckIBMISNetworkACLInlineUpdateAddRule() string {
	return `
    resource "ibm_is_vpc" "testacc_vpc" {
        name = "tf-nwacl-update-vpc"
    }
    resource "ibm_is_network_acl" "test_update_acl" {
        name = "tf-nwacl-update"
        vpc  = ibm_is_vpc.testacc_vpc.id
        rules {
            name        = "rule-one"
            action      = "allow"
            source      = "0.0.0.0/0"
            destination = "0.0.0.0/0"
            direction   = "inbound"
            protocol    = "any"
        }
        rules {
            name        = "rule-two"
            action      = "allow"
            source      = "0.0.0.0/0"
            destination = "0.0.0.0/0"
            direction   = "outbound"
            protocol    = "any"
        }
    }
    `
}

func testAccCheckIBMISNetworkACLInlineUpdateReorder() string {
	return `
    resource "ibm_is_vpc" "testacc_vpc" {
        name = "tf-nwacl-update-vpc"
    }
    resource "ibm_is_network_acl" "test_update_acl" {
        name = "tf-nwacl-update"
        vpc  = ibm_is_vpc.testacc_vpc.id
        rules {
            name        = "rule-two"
            action      = "allow"
            source      = "0.0.0.0/0"
            destination = "0.0.0.0/0"
            direction   = "outbound"
            protocol    = "any"
        }
        rules {
            name        = "rule-one"
            action      = "allow"
            source      = "0.0.0.0/0"
            destination = "0.0.0.0/0"
            direction   = "inbound"
            protocol    = "any"
        }
    }
    `
}

func testAccCheckIBMISNetworkACLInlineUpdateTCP() string {
	return `
    resource "ibm_is_vpc" "testacc_vpc" {
        name = "tf-nwacl-update-vpc"
    }
    resource "ibm_is_network_acl" "test_update_acl" {
        name = "tf-nwacl-update"
        vpc  = ibm_is_vpc.testacc_vpc.id
        rules {
            name        = "rule-one"
            action      = "allow"
            source      = "0.0.0.0/0"
            destination = "0.0.0.0/0"
            direction   = "inbound"
            protocol    = "tcp"
            port_min    = 80
            port_max    = 80
        }
    }
    `
}

func testAccCheckIBMISNetworkACLInlineUpdateUDP() string {
	return `
    resource "ibm_is_vpc" "testacc_vpc" {
        name = "tf-nwacl-update-vpc"
    }
    resource "ibm_is_network_acl" "test_update_acl" {
        name = "tf-nwacl-update"
        vpc  = ibm_is_vpc.testacc_vpc.id
        rules {
            name        = "rule-one"
            action      = "allow"
            source      = "0.0.0.0/0"
            destination = "0.0.0.0/0"
            direction   = "inbound"
            protocol    = "udp"
            port_min    = 53
            port_max    = 53
        }
    }
    `
}

func testAccCheckIBMISNetworkACLInlineUpdateICMP() string {
	return `
    resource "ibm_is_vpc" "testacc_vpc" {
        name = "tf-nwacl-update-vpc"
    }
    resource "ibm_is_network_acl" "test_update_acl" {
        name = "tf-nwacl-update"
        vpc  = ibm_is_vpc.testacc_vpc.id
        rules {
            name        = "rule-one"
            action      = "allow"
            source      = "0.0.0.0/0"
            destination = "0.0.0.0/0"
            direction   = "inbound"
            protocol    = "icmp"
            type        = 8
            code        = 0
        }
    }
    `
}
