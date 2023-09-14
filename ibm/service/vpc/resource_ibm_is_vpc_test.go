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
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccIBMISVPC_basic(t *testing.T) {
	var vpc string
	name1 := fmt.Sprintf("terraformvpcuat-%d", acctest.RandIntRange(10, 100))
	name2 := fmt.Sprintf("terraformvpcuat-%d", acctest.RandIntRange(10, 100))
	apm := "manual"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISVPCDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISVPCConfig(name1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISVPCExists("ibm_is_vpc.testacc_vpc", vpc),
					resource.TestCheckResourceAttr(
						"ibm_is_vpc.testacc_vpc", "name", name1),
					resource.TestCheckResourceAttr(
						"ibm_is_vpc.testacc_vpc", "default_network_acl_name", "dnwacln"),
					resource.TestCheckResourceAttr(
						"ibm_is_vpc.testacc_vpc", "default_security_group_name", "dsgn"),
					resource.TestCheckResourceAttr(
						"ibm_is_vpc.testacc_vpc", "default_routing_table_name", "drtn"),
					resource.TestCheckResourceAttr(
						"ibm_is_vpc.testacc_vpc", "tags.#", "2"),
				),
			},
			{
				Config: testAccCheckIBMISVPCConfigUpdate(name1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISVPCExists("ibm_is_vpc.testacc_vpc", vpc),
					resource.TestCheckResourceAttr(
						"ibm_is_vpc.testacc_vpc", "name", name1),
					resource.TestCheckResourceAttr(
						"ibm_is_vpc.testacc_vpc", "tags.#", "1"),
				),
			},
			{
				Config: testAccCheckIBMISVPCConfig1(name2, apm),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISVPCExists("ibm_is_vpc.testacc_vpc1", vpc),
					resource.TestCheckResourceAttr(
						"ibm_is_vpc.testacc_vpc1", "name", name2),
					resource.TestCheckResourceAttr(
						"ibm_is_vpc.testacc_vpc1", "tags.#", "2"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpc.testacc_vpc1", "cse_source_addresses.#"),
				),
			},
		},
	})
}

func TestAccIBMISVPC_basic_apm(t *testing.T) {
	var vpc string
	name := fmt.Sprintf("terraformvpcuat-%d", acctest.RandIntRange(10, 100))
	apm1 := "auto"
	apm2 := "manual"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISVPCDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISVPCConfig2(name, apm1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISVPCExists("ibm_is_vpc.testacc_vpc1", vpc),
					resource.TestCheckResourceAttr(
						"ibm_is_vpc.testacc_vpc1", "name", name),
					resource.TestCheckResourceAttr(
						"ibm_is_vpc.testacc_vpc1", "address_prefix_management", apm1),
				),
			},
			{
				Config: testAccCheckIBMISVPCConfig2(name, apm2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISVPCExists("ibm_is_vpc.testacc_vpc1", vpc),
					resource.TestCheckResourceAttr(
						"ibm_is_vpc.testacc_vpc1", "name", name),
					resource.TestCheckResourceAttr(
						"ibm_is_vpc.testacc_vpc1", "address_prefix_management", apm2),
				),
			},
		},
	})
}

func TestAccIBMISVPC_securityGroups(t *testing.T) {
	var vpc string
	vpcname := fmt.Sprintf("terraformvpcuat-%d", acctest.RandIntRange(10, 100))
	sgname := fmt.Sprintf("terraformvpcsg-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISVPCDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISVPCSgConfig(vpcname, sgname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISVPCExists("ibm_is_vpc.testacc_vpc", vpc),
					resource.TestCheckResourceAttr(
						"ibm_is_vpc.testacc_vpc", "name", vpcname),
					resource.TestCheckResourceAttr(
						"ibm_is_security_group.testacc_security_group", "name", sgname),
					resource.TestCheckResourceAttrSet("ibm_is_vpc.testacc_vpc", "security_group.0.group_name"),
					resource.TestCheckResourceAttrSet("ibm_is_vpc.testacc_vpc", "security_group.0.group_id"),
				),
			},
		},
	})
}

func TestAccIBMISVPC_noSGACLRules(t *testing.T) {
	var vpc string
	vpcname := fmt.Sprintf("terraformvpcuat-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISVPCDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISVPCNoSgAclRulesConfig(vpcname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISVPCExists("ibm_is_vpc.testacc_vpc", vpc),
					resource.TestCheckResourceAttr(
						"ibm_is_vpc.testacc_vpc", "name", vpcname),
					resource.TestCheckNoResourceAttr("ibm_is_vpc.testacc_vpc", "security_group.0.rules.#"),
					resource.TestCheckNoResourceAttr("ibm_is_vpc.testacc_vpc", "security_group.0.rules.#"),
				),
			},
		},
	})
}

func testAccCheckIBMISVPCDestroy(s *terraform.State) error {
	sess, _ := acc.TestAccProvider.Meta().(conns.ClientSession).VpcV1API()
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_is_vpc" {
			continue
		}

		getvpcoptions := &vpcv1.GetVPCOptions{
			ID: &rs.Primary.ID,
		}
		_, _, err := sess.GetVPC(getvpcoptions)

		if err == nil {
			return fmt.Errorf("vpc still exists: %s", rs.Primary.ID)
		}
	}
	return nil
}

func testAccCheckIBMISVPCExists(n, vpcID string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return errors.New("No Record ID is set")
		}
		sess, _ := acc.TestAccProvider.Meta().(conns.ClientSession).VpcV1API()
		getvpcoptions := &vpcv1.GetVPCOptions{
			ID: &rs.Primary.ID,
		}
		foundvpc, _, err := sess.GetVPC(getvpcoptions)
		if err != nil {
			return err
		}
		vpcID = *foundvpc.ID
		return nil
	}
}

func testAccCheckIBMISVPCConfig(name string) string {
	return fmt.Sprintf(`
	resource "ibm_is_vpc" "testacc_vpc" {
		name = "%s"
		default_network_acl_name = "dnwacln"
		default_security_group_name = "dsgn"
		default_routing_table_name = "drtn"
		tags = ["Tag1", "tag2"]
	}`, name)

}

func testAccCheckIBMISVPCConfigUpdate(name string) string {
	return fmt.Sprintf(`
	resource "ibm_is_vpc" "testacc_vpc" {
		name = "%s"
		tags = ["tag1"]
	}`, name)

}

func testAccCheckIBMISVPCConfig1(name string, apm string) string {
	return fmt.Sprintf(`
	resource "ibm_is_vpc" "testacc_vpc1" {
		name = "%s"
		address_prefix_management = "%s"
		tags = ["Tag1", "tag2"]
	}`, name, apm)

}
func testAccCheckIBMISVPCConfig2(name string, apm string) string {
	return fmt.Sprintf(`
	resource "ibm_is_vpc" "testacc_vpc1" {
		name = "%s"
		address_prefix_management = "%s"
	}`, name, apm)
}

func testAccCheckIBMISVPCSgConfig(vpcname string, sgname string) string {
	return fmt.Sprintf(`
	resource "ibm_is_vpc" "testacc_vpc" {
		name = "%s"
	  }
	  
	  resource "ibm_is_security_group" "testacc_security_group" {
		name = "%s"
		vpc  = ibm_is_vpc.testacc_vpc.id
	  }
	  
	  resource "ibm_is_security_group_rule" "testacc_security_group_rule_udp" {
		group      = ibm_is_security_group.testacc_security_group.id
		direction  = "inbound"
		remote     = "127.0.0.1"
		udp {
			port_min = 805
			port_max = 807
		}
	}  
`, vpcname, sgname)

}

func testAccCheckIBMISVPCNoSgAclRulesConfig(vpcname string) string {
	return fmt.Sprintf(`
	resource "ibm_is_vpc" "testacc_vpc" {
		name = "%s"
		no_sg_acl_rules = true
	  }
`, vpcname)

}
