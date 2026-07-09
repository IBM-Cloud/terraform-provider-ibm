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
// Tests for surgical_rule_update flag
// ---------------------------------------------------------------------------

// TestNetworkACL_SurgicalUpdate tests the surgical update path (surgical_rule_update=true).
// Covers: add rule, remove rule, reorder rules, mutable field patch, protocol change.
func TestNetworkACL_SurgicalUpdate(t *testing.T) {
	var nwACL string
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: checkNetworkACLDestroy,
		Steps: []resource.TestStep{
			{
				// Step 1: Create with two rules, surgical mode enabled.
				Config: testAccNACLSurgical_Base(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISNetworkACLExists("ibm_is_network_acl.surgical_acl", nwACL),
					resource.TestCheckResourceAttr("ibm_is_network_acl.surgical_acl", "surgical_rule_update", "true"),
					resource.TestCheckResourceAttr("ibm_is_network_acl.surgical_acl", "rules.#", "2"),
					resource.TestCheckResourceAttr("ibm_is_network_acl.surgical_acl", "rules.0.name", "rule-a"),
					resource.TestCheckResourceAttr("ibm_is_network_acl.surgical_acl", "rules.1.name", "rule-b"),
				),
			},
			{
				// Step 2: Add a new rule at the end.
				Config: testAccNACLSurgical_AddRule(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISNetworkACLExists("ibm_is_network_acl.surgical_acl", nwACL),
					resource.TestCheckResourceAttr("ibm_is_network_acl.surgical_acl", "rules.#", "3"),
					resource.TestCheckResourceAttr("ibm_is_network_acl.surgical_acl", "rules.0.name", "rule-a"),
					resource.TestCheckResourceAttr("ibm_is_network_acl.surgical_acl", "rules.1.name", "rule-b"),
					resource.TestCheckResourceAttr("ibm_is_network_acl.surgical_acl", "rules.2.name", "rule-c"),
				),
			},
			{
				// Step 3: Remove rule-b (middle rule).
				Config: testAccNACLSurgical_RemoveRule(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISNetworkACLExists("ibm_is_network_acl.surgical_acl", nwACL),
					resource.TestCheckResourceAttr("ibm_is_network_acl.surgical_acl", "rules.#", "2"),
					resource.TestCheckResourceAttr("ibm_is_network_acl.surgical_acl", "rules.0.name", "rule-a"),
					resource.TestCheckResourceAttr("ibm_is_network_acl.surgical_acl", "rules.1.name", "rule-c"),
				),
			},
			{
				// Step 4: Reorder — swap rule-a and rule-c.
				Config: testAccNACLSurgical_Reorder(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISNetworkACLExists("ibm_is_network_acl.surgical_acl", nwACL),
					resource.TestCheckResourceAttr("ibm_is_network_acl.surgical_acl", "rules.#", "2"),
					resource.TestCheckResourceAttr("ibm_is_network_acl.surgical_acl", "rules.0.name", "rule-c"),
					resource.TestCheckResourceAttr("ibm_is_network_acl.surgical_acl", "rules.1.name", "rule-a"),
				),
			},
			{
				// Step 5: Patch mutable field (source CIDR) on rule-a only.
				Config: testAccNACLSurgical_PatchMutable(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISNetworkACLExists("ibm_is_network_acl.surgical_acl", nwACL),
					resource.TestCheckResourceAttr("ibm_is_network_acl.surgical_acl", "rules.#", "2"),
					resource.TestCheckResourceAttr("ibm_is_network_acl.surgical_acl", "rules.1.name", "rule-a"),
					resource.TestCheckResourceAttr("ibm_is_network_acl.surgical_acl", "rules.1.source", "10.0.0.0/8"),
				),
			},
			{
				// Step 6: Protocol change on rule-c (any → tcp) — delete+recreate at same position.
				Config: testAccNACLSurgical_ProtocolChange(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISNetworkACLExists("ibm_is_network_acl.surgical_acl", nwACL),
					resource.TestCheckResourceAttr("ibm_is_network_acl.surgical_acl", "rules.#", "2"),
					resource.TestCheckResourceAttr("ibm_is_network_acl.surgical_acl", "rules.0.name", "rule-c"),
					resource.TestCheckResourceAttr("ibm_is_network_acl.surgical_acl", "rules.0.protocol", "tcp"),
				),
			},
		},
	})
}

// TestNetworkACL_LegacyUpdate tests the legacy clear+recreate path (surgical_rule_update=false/absent).
// Covers: add rule, remove rule, mutable field change — all via full wipe+recreate.
func TestNetworkACL_LegacyUpdate(t *testing.T) {
	var nwACL string
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: checkNetworkACLDestroy,
		Steps: []resource.TestStep{
			{
				// Step 1: Create with two rules, legacy mode (no flag).
				Config: testAccNACLLegacy_Base(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISNetworkACLExists("ibm_is_network_acl.legacy_acl", nwACL),
					resource.TestCheckResourceAttr("ibm_is_network_acl.legacy_acl", "surgical_rule_update", "false"),
					resource.TestCheckResourceAttr("ibm_is_network_acl.legacy_acl", "rules.#", "2"),
					resource.TestCheckResourceAttr("ibm_is_network_acl.legacy_acl", "rules.0.name", "rule-x"),
					resource.TestCheckResourceAttr("ibm_is_network_acl.legacy_acl", "rules.1.name", "rule-y"),
				),
			},
			{
				// Step 2: Add a rule — triggers full clear+recreate.
				Config: testAccNACLLegacy_AddRule(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISNetworkACLExists("ibm_is_network_acl.legacy_acl", nwACL),
					resource.TestCheckResourceAttr("ibm_is_network_acl.legacy_acl", "rules.#", "3"),
					resource.TestCheckResourceAttr("ibm_is_network_acl.legacy_acl", "rules.0.name", "rule-x"),
					resource.TestCheckResourceAttr("ibm_is_network_acl.legacy_acl", "rules.1.name", "rule-y"),
					resource.TestCheckResourceAttr("ibm_is_network_acl.legacy_acl", "rules.2.name", "rule-z"),
				),
			},
			{
				// Step 3: Remove rule-y — triggers full clear+recreate.
				Config: testAccNACLLegacy_RemoveRule(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISNetworkACLExists("ibm_is_network_acl.legacy_acl", nwACL),
					resource.TestCheckResourceAttr("ibm_is_network_acl.legacy_acl", "rules.#", "2"),
					resource.TestCheckResourceAttr("ibm_is_network_acl.legacy_acl", "rules.0.name", "rule-x"),
					resource.TestCheckResourceAttr("ibm_is_network_acl.legacy_acl", "rules.1.name", "rule-z"),
				),
			},
			{
				// Step 4: Patch mutable field (destination) — triggers full clear+recreate.
				Config: testAccNACLLegacy_PatchMutable(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISNetworkACLExists("ibm_is_network_acl.legacy_acl", nwACL),
					resource.TestCheckResourceAttr("ibm_is_network_acl.legacy_acl", "rules.#", "2"),
					resource.TestCheckResourceAttr("ibm_is_network_acl.legacy_acl", "rules.0.name", "rule-x"),
					resource.TestCheckResourceAttr("ibm_is_network_acl.legacy_acl", "rules.0.destination", "10.0.0.0/8"),
				),
			},
		},
	})
}

// ---------------------------------------------------------------------------
// Surgical update configs
// ---------------------------------------------------------------------------

func testAccNACLSurgical_Base() string {
	return `
resource "ibm_is_vpc" "surgical_vpc" {
  name = "tf-surgical-nacl-vpc"
}
resource "ibm_is_network_acl" "surgical_acl" {
  name                 = "tf-surgical-nacl"
  vpc                  = ibm_is_vpc.surgical_vpc.id
  surgical_rule_update = true
  rules {
    name        = "rule-a"
    action      = "allow"
    source      = "0.0.0.0/0"
    destination = "0.0.0.0/0"
    direction   = "inbound"
    protocol    = "any"
  }
  rules {
    name        = "rule-b"
    action      = "allow"
    source      = "0.0.0.0/0"
    destination = "0.0.0.0/0"
    direction   = "outbound"
    protocol    = "any"
  }
}
`
}

func testAccNACLSurgical_AddRule() string {
	return `
resource "ibm_is_vpc" "surgical_vpc" {
  name = "tf-surgical-nacl-vpc"
}
resource "ibm_is_network_acl" "surgical_acl" {
  name                 = "tf-surgical-nacl"
  vpc                  = ibm_is_vpc.surgical_vpc.id
  surgical_rule_update = true
  rules {
    name        = "rule-a"
    action      = "allow"
    source      = "0.0.0.0/0"
    destination = "0.0.0.0/0"
    direction   = "inbound"
    protocol    = "any"
  }
  rules {
    name        = "rule-b"
    action      = "allow"
    source      = "0.0.0.0/0"
    destination = "0.0.0.0/0"
    direction   = "outbound"
    protocol    = "any"
  }
  rules {
    name        = "rule-c"
    action      = "deny"
    source      = "0.0.0.0/0"
    destination = "0.0.0.0/0"
    direction   = "inbound"
    protocol    = "any"
  }
}
`
}

func testAccNACLSurgical_RemoveRule() string {
	return `
resource "ibm_is_vpc" "surgical_vpc" {
  name = "tf-surgical-nacl-vpc"
}
resource "ibm_is_network_acl" "surgical_acl" {
  name                 = "tf-surgical-nacl"
  vpc                  = ibm_is_vpc.surgical_vpc.id
  surgical_rule_update = true
  rules {
    name        = "rule-a"
    action      = "allow"
    source      = "0.0.0.0/0"
    destination = "0.0.0.0/0"
    direction   = "inbound"
    protocol    = "any"
  }
  rules {
    name        = "rule-c"
    action      = "deny"
    source      = "0.0.0.0/0"
    destination = "0.0.0.0/0"
    direction   = "inbound"
    protocol    = "any"
  }
}
`
}

func testAccNACLSurgical_Reorder() string {
	return `
resource "ibm_is_vpc" "surgical_vpc" {
  name = "tf-surgical-nacl-vpc"
}
resource "ibm_is_network_acl" "surgical_acl" {
  name                 = "tf-surgical-nacl"
  vpc                  = ibm_is_vpc.surgical_vpc.id
  surgical_rule_update = true
  rules {
    name        = "rule-c"
    action      = "deny"
    source      = "0.0.0.0/0"
    destination = "0.0.0.0/0"
    direction   = "inbound"
    protocol    = "any"
  }
  rules {
    name        = "rule-a"
    action      = "allow"
    source      = "0.0.0.0/0"
    destination = "0.0.0.0/0"
    direction   = "inbound"
    protocol    = "any"
  }
}
`
}

func testAccNACLSurgical_PatchMutable() string {
	return `
resource "ibm_is_vpc" "surgical_vpc" {
  name = "tf-surgical-nacl-vpc"
}
resource "ibm_is_network_acl" "surgical_acl" {
  name                 = "tf-surgical-nacl"
  vpc                  = ibm_is_vpc.surgical_vpc.id
  surgical_rule_update = true
  rules {
    name        = "rule-c"
    action      = "deny"
    source      = "0.0.0.0/0"
    destination = "0.0.0.0/0"
    direction   = "inbound"
    protocol    = "any"
  }
  rules {
    name        = "rule-a"
    action      = "allow"
    source      = "10.0.0.0/8"
    destination = "0.0.0.0/0"
    direction   = "inbound"
    protocol    = "any"
  }
}
`
}

func testAccNACLSurgical_ProtocolChange() string {
	return `
resource "ibm_is_vpc" "surgical_vpc" {
  name = "tf-surgical-nacl-vpc"
}
resource "ibm_is_network_acl" "surgical_acl" {
  name                 = "tf-surgical-nacl"
  vpc                  = ibm_is_vpc.surgical_vpc.id
  surgical_rule_update = true
  rules {
    name        = "rule-c"
    action      = "allow"
    source      = "0.0.0.0/0"
    destination = "0.0.0.0/0"
    direction   = "inbound"
    protocol    = "tcp"
    port_min    = 80
    port_max    = 80
  }
  rules {
    name        = "rule-a"
    action      = "allow"
    source      = "10.0.0.0/8"
    destination = "0.0.0.0/0"
    direction   = "inbound"
    protocol    = "any"
  }
}
`
}

// ---------------------------------------------------------------------------
// Legacy update configs
// ---------------------------------------------------------------------------

func testAccNACLLegacy_Base() string {
	return `
resource "ibm_is_vpc" "legacy_vpc" {
  name = "tf-legacy-nacl-vpc"
}
resource "ibm_is_network_acl" "legacy_acl" {
  name                 = "tf-legacy-nacl"
  vpc                  = ibm_is_vpc.legacy_vpc.id
  surgical_rule_update = false
  rules {
    name        = "rule-x"
    action      = "allow"
    source      = "0.0.0.0/0"
    destination = "0.0.0.0/0"
    direction   = "inbound"
    protocol    = "any"
  }
  rules {
    name        = "rule-y"
    action      = "allow"
    source      = "0.0.0.0/0"
    destination = "0.0.0.0/0"
    direction   = "outbound"
    protocol    = "any"
  }
}
`
}

func testAccNACLLegacy_AddRule() string {
	return `
resource "ibm_is_vpc" "legacy_vpc" {
  name = "tf-legacy-nacl-vpc"
}
resource "ibm_is_network_acl" "legacy_acl" {
  name                 = "tf-legacy-nacl"
  vpc                  = ibm_is_vpc.legacy_vpc.id
  surgical_rule_update = false
  rules {
    name        = "rule-x"
    action      = "allow"
    source      = "0.0.0.0/0"
    destination = "0.0.0.0/0"
    direction   = "inbound"
    protocol    = "any"
  }
  rules {
    name        = "rule-y"
    action      = "allow"
    source      = "0.0.0.0/0"
    destination = "0.0.0.0/0"
    direction   = "outbound"
    protocol    = "any"
  }
  rules {
    name        = "rule-z"
    action      = "deny"
    source      = "0.0.0.0/0"
    destination = "0.0.0.0/0"
    direction   = "inbound"
    protocol    = "any"
  }
}
`
}

func testAccNACLLegacy_RemoveRule() string {
	return `
resource "ibm_is_vpc" "legacy_vpc" {
  name = "tf-legacy-nacl-vpc"
}
resource "ibm_is_network_acl" "legacy_acl" {
  name                 = "tf-legacy-nacl"
  vpc                  = ibm_is_vpc.legacy_vpc.id
  surgical_rule_update = false
  rules {
    name        = "rule-x"
    action      = "allow"
    source      = "0.0.0.0/0"
    destination = "0.0.0.0/0"
    direction   = "inbound"
    protocol    = "any"
  }
  rules {
    name        = "rule-z"
    action      = "deny"
    source      = "0.0.0.0/0"
    destination = "0.0.0.0/0"
    direction   = "inbound"
    protocol    = "any"
  }
}
`
}

func testAccNACLLegacy_PatchMutable() string {
	return `
resource "ibm_is_vpc" "legacy_vpc" {
  name = "tf-legacy-nacl-vpc"
}
resource "ibm_is_network_acl" "legacy_acl" {
  name                 = "tf-legacy-nacl"
  vpc                  = ibm_is_vpc.legacy_vpc.id
  surgical_rule_update = false
  rules {
    name        = "rule-x"
    action      = "allow"
    source      = "0.0.0.0/0"
    destination = "10.0.0.0/8"
    direction   = "inbound"
    protocol    = "any"
  }
  rules {
    name        = "rule-z"
    action      = "deny"
    source      = "0.0.0.0/0"
    destination = "0.0.0.0/0"
    direction   = "inbound"
    protocol    = "any"
  }
}
`
}
