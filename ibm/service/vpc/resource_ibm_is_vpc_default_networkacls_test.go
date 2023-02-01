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

func TestVPCDefaultNetworkACL(t *testing.T) {
	var nwACL string
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: checkVPCDefaultNetworkACLDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISVPCDefaultNetworkACLConfig1(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISVPCDefaultNetworkACLExists("ibm_is_vpc_default_network_acl.isExampleACL", nwACL),
					resource.TestCheckResourceAttr(
						"ibm_is_vpc_default_network_acl.isExampleACL", "name", "is-example-acl"),
					resource.TestCheckResourceAttr(
						"ibm_is_vpc_default_network_acl.isExampleACL", "rules.#", "2"),
					resource.TestCheckResourceAttr(
						"ibm_is_vpc_default_network_acl.isExampleACL", "tags.#", "2"),
				),
			},
		},
	})
}

func checkVPCDefaultNetworkACLDestroy(s *terraform.State) error {
	sess, _ := acc.TestAccProvider.Meta().(conns.ClientSession).VpcV1API()
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_is_vpc_default_network_acl" {
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

func testAccCheckIBMISVPCDefaultNetworkACLExists(n, nwACL string) resource.TestCheckFunc {
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

func testAccCheckIBMISVPCDefaultNetworkACLConfig1() string {
	return fmt.Sprintf(`
	resource "ibm_is_vpc" "testacc_vpc" {
		name = "tf-nwacl-vpc"
	  }

	resource "ibm_is_vpc_default_network_acl" "isExampleACL" {
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
