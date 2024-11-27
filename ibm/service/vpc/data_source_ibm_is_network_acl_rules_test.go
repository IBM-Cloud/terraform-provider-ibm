// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccIBMISNetworkACLRulesDataSource_basic(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISNetworkACLRulesDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_network_acl_rules.testacc_ds_ruleslist", "rules.0.name"),
				),
			},
		},
	})
}

func testAccCheckIBMISNetworkACLRulesDataSourceConfig() string {
	return fmt.Sprintf(`
		resource "ibm_is_vpc" "testacc_nacl_vpc" {
			name="tf-nacl-vpc"
		}
		resource "ibm_is_network_acl_rule" "testacc_nacl" {
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
		data "ibm_is_network_acl_rules" "testacc_ds_ruleslist" {
			network_acl = ibm_is_vpc.testacc_nacl_vpc.default_network_acl
		}`)
}
