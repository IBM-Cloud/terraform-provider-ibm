// Copyright IBM Corp. 2017, 2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"errors"
	"fmt"
	"strings"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccIBMISVPCDefaultNetworkACL_basic(t *testing.T) {
	var nwACLID string
	name := fmt.Sprintf("tfvpc-defaultacl-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISVPCDefaultNetworkACLDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISVPCDefaultNetworkACLConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISVPCDefaultNetworkACLExists("ibm_is_vpc_default_network_acl.test_default_acl", nwACLID),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpc_default_network_acl.test_default_acl", "name"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpc_default_network_acl.test_default_acl", "default_network_acl"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpc_default_network_acl.test_default_acl", "crn"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpc_default_network_acl.test_default_acl", "resource_group.#"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpc_default_network_acl.test_default_acl", "subnets.#"),
				),
			},
		},
	})
}

func TestAccIBMISVPCDefaultNetworkACL_name_update(t *testing.T) {
	var nwACLID string
	name := fmt.Sprintf("tfvpc-defaultacl-%d", acctest.RandIntRange(10, 100))
	updatedName := fmt.Sprintf("tfvpc-defaultacl-updated-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISVPCDefaultNetworkACLDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISVPCDefaultNetworkACLConfigWithName(name, name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISVPCDefaultNetworkACLExists("ibm_is_vpc_default_network_acl.test_default_acl", nwACLID),
					resource.TestCheckResourceAttr(
						"ibm_is_vpc_default_network_acl.test_default_acl", "name", name),
				),
			},
			{
				Config: testAccCheckIBMISVPCDefaultNetworkACLConfigWithName(name, updatedName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISVPCDefaultNetworkACLExists("ibm_is_vpc_default_network_acl.test_default_acl", nwACLID),
					resource.TestCheckResourceAttr(
						"ibm_is_vpc_default_network_acl.test_default_acl", "name", updatedName),
				),
			},
		},
	})
}

func TestAccIBMISVPCDefaultNetworkACL_tags(t *testing.T) {
	var nwACLID string
	name := fmt.Sprintf("tfvpc-defaultacl-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISVPCDefaultNetworkACLDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISVPCDefaultNetworkACLConfigWithTags(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISVPCDefaultNetworkACLExists("ibm_is_vpc_default_network_acl.test_default_acl", nwACLID),
					resource.TestCheckResourceAttr(
						"ibm_is_vpc_default_network_acl.test_default_acl", "tags.#", "2"),
					resource.TestCheckResourceAttr(
						"ibm_is_vpc_default_network_acl.test_default_acl", "tags.0", "env:test"),
					resource.TestCheckResourceAttr(
						"ibm_is_vpc_default_network_acl.test_default_acl", "tags.1", "team:dev"),
				),
			},
			{
				Config: testAccCheckIBMISVPCDefaultNetworkACLConfigWithUpdatedTags(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISVPCDefaultNetworkACLExists("ibm_is_vpc_default_network_acl.test_default_acl", nwACLID),
					resource.TestCheckResourceAttr(
						"ibm_is_vpc_default_network_acl.test_default_acl", "tags.#", "3"),
					resource.TestCheckResourceAttr(
						"ibm_is_vpc_default_network_acl.test_default_acl", "tags.0", "env:production"),
					resource.TestCheckResourceAttr(
						"ibm_is_vpc_default_network_acl.test_default_acl", "tags.1", "team:ops"),
					resource.TestCheckResourceAttr(
						"ibm_is_vpc_default_network_acl.test_default_acl", "tags.2", "project:web"),
				),
			},
		},
	})
}

func TestAccIBMISVPCDefaultNetworkACL_access_tags(t *testing.T) {
	var nwACLID string
	name := fmt.Sprintf("tfvpc-defaultacl-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISVPCDefaultNetworkACLDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISVPCDefaultNetworkACLConfigWithAccessTags(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISVPCDefaultNetworkACLExists("ibm_is_vpc_default_network_acl.test_default_acl", nwACLID),
					resource.TestCheckResourceAttr(
						"ibm_is_vpc_default_network_acl.test_default_acl", "access_tags.#", "1"),
				),
			},
			{
				ResourceName:      "ibm_is_vpc_default_network_acl.test_default_acl",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIBMISVPCDefaultNetworkACLDestroy(s *terraform.State) error {
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_is_vpc_default_network_acl" {
			continue
		}

		parts := strings.Split(rs.Primary.ID, "/")
		if len(parts) != 2 {
			return fmt.Errorf("Invalid ID format for default network ACL")
		}

		sess, err := vpcClient(acc.TestAccProvider.Meta())
		if err != nil {
			return err
		}

		networkACLID := parts[1]
		getNetworkACLOptions := sess.NewGetNetworkACLOptions(networkACLID)
		foundNetworkACL, _, err := sess.GetNetworkACL(getNetworkACLOptions)
		if err == nil {
			// Default network ACL should still exist, but we check if it's properly configured
			if foundNetworkACL != nil {
				// This is expected - default network ACL should still exist
				continue
			}
			return fmt.Errorf("Default network ACL still exists but not found properly: %s", rs.Primary.ID)
		}
	}
	return nil
}

func testAccCheckIBMISVPCDefaultNetworkACLExists(n, nwACLID string) resource.TestCheckFunc {
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

		if len(parts) != 2 {
			return fmt.Errorf("Invalid ID format: %s", rs.Primary.ID)
		}

		networkACLID := parts[1]
		sess, err := vpcClient(acc.TestAccProvider.Meta())
		if err != nil {
			return err
		}

		getNetworkACLOptions := sess.NewGetNetworkACLOptions(networkACLID)
		foundNetworkACL, response, err := sess.GetNetworkACL(getNetworkACLOptions)
		if err != nil {
			return fmt.Errorf("[ERROR] Error Getting VPC Default Network ACL: %s\n%s", err, response)
		}

		nwACLID = *foundNetworkACL.ID
		return nil
	}
}

func testAccCheckIBMISVPCDefaultNetworkACLConfig(name string) string {
	return fmt.Sprintf(`
resource "ibm_is_vpc" "testacc_vpc" {
	name = "%s"
}

resource "ibm_is_vpc_default_network_acl" "test_default_acl" {
	vpc = ibm_is_vpc.testacc_vpc.id
}`, name)
}

func testAccCheckIBMISVPCDefaultNetworkACLConfigWithName(vpcName, aclName string) string {
	return fmt.Sprintf(`
resource "ibm_is_vpc" "testacc_vpc" {
	name = "%s"
}

resource "ibm_is_vpc_default_network_acl" "test_default_acl" {
	vpc  = ibm_is_vpc.testacc_vpc.id
	name = "%s"
}`, vpcName, aclName)
}

func testAccCheckIBMISVPCDefaultNetworkACLConfigWithTags(name string) string {
	return fmt.Sprintf(`
resource "ibm_is_vpc" "testacc_vpc" {
	name = "%s"
}

resource "ibm_is_vpc_default_network_acl" "test_default_acl" {
	vpc  = ibm_is_vpc.testacc_vpc.id
	name = "%s-default-acl"
	tags = ["env:test", "team:dev"]
}`, name, name)
}

func testAccCheckIBMISVPCDefaultNetworkACLConfigWithUpdatedTags(name string) string {
	return fmt.Sprintf(`
resource "ibm_is_vpc" "testacc_vpc" {
	name = "%s"
}

resource "ibm_is_vpc_default_network_acl" "test_default_acl" {
	vpc  = ibm_is_vpc.testacc_vpc.id
	name = "%s-default-acl"
	tags = ["env:production", "team:ops", "project:web"]
}`, name, name)
}

func testAccCheckIBMISVPCDefaultNetworkACLConfigWithAccessTags(name string) string {
	return fmt.Sprintf(`
resource "ibm_is_vpc" "testacc_vpc" {
	name = "%s"
}

resource "ibm_is_vpc_default_network_acl" "test_default_acl" {
	vpc         = ibm_is_vpc.testacc_vpc.id
	name        = "%s-default-acl"
	access_tags = ["project:test"]
}`, name, name)
}
