// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMIsIkePoliciesDataSourceBasic(t *testing.T) {
	name := fmt.Sprintf("tfike-name-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIsIkePoliciesDataSourceConfigBasic(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_ike_policies.is_ike_policies", "ike_policies.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_ike_policies.is_ike_policies", "ike_policies.0.authentication_algorithm"),
					resource.TestCheckResourceAttrSet("data.ibm_is_ike_policies.is_ike_policies", "ike_policies.0.dh_group"),
					resource.TestCheckResourceAttrSet("data.ibm_is_ike_policies.is_ike_policies", "ike_policies.0.ike_version"),
					resource.TestCheckResourceAttrSet("data.ibm_is_ike_policies.is_ike_policies", "ike_policies.0.key_lifetime"),
					resource.TestCheckResourceAttrSet("data.ibm_is_ike_policies.is_ike_policies", "ike_policies.0.name"),
					resource.TestCheckResourceAttrSet("data.ibm_is_ike_policies.is_ike_policies", "ike_policies.0.negotiation_mode"),
				),
			},
		},
	})
}

func testAccCheckIBMIsIkePoliciesDataSourceConfigBasic(name string) string {
	return fmt.Sprintf(`
		resource "ibm_is_ike_policy" "example" {
			name = "%s"
			authentication_algorithm = "sha1"
			encryption_algorithm = "aes128"
			dh_group = 5
			ike_version = 2
			key_lifetime = 1800
		}
		data "ibm_is_ike_policies" "is_ike_policies" {
			depends_on = [ibm_is_ike_policy.example]
		}
	`, name)
}
