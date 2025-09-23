// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"fmt"
	"testing"

	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMIsNetworkACLDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIsNetworkACLDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_network_acl.is_network_acl", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_network_acl.is_network_acl", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_is_network_acl.is_network_acl", "crn"),
					resource.TestCheckResourceAttrSet("data.ibm_is_network_acl.is_network_acl", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_network_acl.is_network_acl", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_is_network_acl.is_network_acl", "resource_group.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_network_acl.is_network_acl", "rules.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_network_acl.is_network_acl", "subnets.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_network_acl.is_network_acl", "vpc.#"),
				),
			},
		},
	})
}

func testAccCheckIBMIsNetworkACLDataSourceConfigBasic() string {
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
	
      data "ibm_is_network_acl" "is_network_acl" {
    	vpc_name = ibm_is_vpc.testacc_vpc.name
		name = ibm_is_network_acl.isExampleACL.name
      }
	`)
}
