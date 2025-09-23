// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"fmt"
	"testing"

	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMISNetworkACLRuleDataSource_basic(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISNetworkACLRuleDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_network_acl_rule.testacc_nacl_rule", "name"),
				),
			},
		},
	})
}

func testAccCheckIBMISNetworkACLRuleDataSourceConfig() string {
	return fmt.Sprintf(`
		resource "ibm_is_vpc" "testacc_nacl_vpc" {
			name="tf-nacl-vpc"
		}
		resource "ibm_is_network_acl_rule" "testacc_nacl" {
			depends_on = [ibm_is_vpc.testacc_nacl_vpc]
			network_acl = ibm_is_vpc.testacc_nacl_vpc.default_network_acl
			name           = "tf-outbound-1"
			action         = "allow"
			source         = "0.0.0.0/0"
			destination    = "0.0.0.0/0"
			direction      = "outbound"
			icmp {
				code = 1
				type = 1
			}
		}
		data "ibm_is_network_acl_rule" "testacc_nacl_rule" {
			depends_on = [ibm_is_vpc.testacc_nacl_vpc, ibm_is_network_acl_rule.testacc_nacl]
			network_acl = ibm_is_vpc.testacc_nacl_vpc.default_network_acl
			name = "tf-outbound-1"
		}`)
}
