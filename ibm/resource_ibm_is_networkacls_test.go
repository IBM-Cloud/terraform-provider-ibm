// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"errors"
	"fmt"
	"testing"

	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestNetworkACLGen1(t *testing.T) {
	var nwACL string
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: checkNetworkACLDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
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

func TestNetworkACLGen2(t *testing.T) {
	var nwACL string
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: checkNetworkACLDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMISNetworkACLConfig1(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISNetworkACLExists("ibm_is_network_acl.isExampleACL", nwACL),
					resource.TestCheckResourceAttr(
						"ibm_is_network_acl.isExampleACL", "name", "is-example-acl"),
					resource.TestCheckResourceAttr(
						"ibm_is_network_acl.isExampleACL", "rules.#", "2"),
					resource.TestCheckResourceAttr(
						"ibm_is_network_acl.isExampleACL", "tags.#", "2"),
				),
			},
		},
	})
}

func checkNetworkACLDestroy(s *terraform.State) error {
	sess, _ := testAccProvider.Meta().(ClientSession).VpcV1API()
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
		sess, _ := testAccProvider.Meta().(ClientSession).VpcV1API()
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
	  }
	`)
}
