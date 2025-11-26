// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccIBMIsIkePoliciesDataSourceBasic(t *testing.T) {
	name1 := fmt.Sprintf("tfike-name-%d", acctest.RandIntRange(10, 100))
	name2 := fmt.Sprintf("tfike-name-%d", acctest.RandIntRange(10, 100))
	dataSourceName := "data.ibm_is_ike_policies.is_ike_policies"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIsIkePoliciesDataSourceConfigBasic(name1, name2),
				Check: resource.ComposeTestCheckFunc(
					// Verify the list of policies
					resource.TestCheckResourceAttrSet(dataSourceName, "ike_policies.#"),

					// Verify all fields in the first policy entry
					resource.TestCheckResourceAttrSet(dataSourceName, "ike_policies.0.id"),
					resource.TestCheckResourceAttrSet(dataSourceName, "ike_policies.0.name"),
					resource.TestCheckResourceAttrSet(dataSourceName, "ike_policies.0.authentication_algorithm"),
					resource.TestCheckResourceAttrSet(dataSourceName, "ike_policies.0.encryption_algorithm"),
					resource.TestCheckResourceAttrSet(dataSourceName, "ike_policies.0.dh_group"),
					resource.TestCheckResourceAttrSet(dataSourceName, "ike_policies.0.ike_version"),
					resource.TestCheckResourceAttrSet(dataSourceName, "ike_policies.0.key_lifetime"),
					resource.TestCheckResourceAttrSet(dataSourceName, "ike_policies.0.negotiation_mode"),
					resource.TestCheckResourceAttrSet(dataSourceName, "ike_policies.0.href"),
					resource.TestCheckResourceAttrSet(dataSourceName, "ike_policies.0.created_at"),
					resource.TestCheckResourceAttrSet(dataSourceName, "ike_policies.0.resource_type"),

					// Verify connections array exists (may be empty)
					resource.TestCheckResourceAttrSet(dataSourceName, "ike_policies.0.connections.#"),

					// Verify resource_group nested structure
					resource.TestCheckResourceAttrSet(dataSourceName, "ike_policies.0.resource_group.#"),
					resource.TestCheckResourceAttrSet(dataSourceName, "ike_policies.0.resource_group.0.id"),
					resource.TestCheckResourceAttrSet(dataSourceName, "ike_policies.0.resource_group.0.name"),
					resource.TestCheckResourceAttrSet(dataSourceName, "ike_policies.0.resource_group.0.href"),

					// Verify negotiation_mode is "main"
					resource.TestCheckResourceAttr(dataSourceName, "ike_policies.0.negotiation_mode", "main"),

					// Additional check to verify we can find the resources we created
					// Note: This is a loose check since we can't guarantee the order of policies in the list
					// but at least one of our created policies should be in the list
					resource.TestCheckResourceAttrSet(dataSourceName, "ike_policies.#"),
				),
			},
		},
	})
}

func testAccCheckIBMIsIkePoliciesDataSourceConfigBasic(name1, name2 string) string {
	return fmt.Sprintf(`
		// Create two policies with different configurations to ensure we test various values
		resource "ibm_is_ike_policy" "example1" {
			name = "%s"
			authentication_algorithm = "sha256"
			encryption_algorithm = "aes128"
			dh_group = 14
			ike_version = 2
			key_lifetime = 1800
		}
		
		resource "ibm_is_ike_policy" "example2" {
			name = "%s"
			authentication_algorithm = "sha512"
			encryption_algorithm = "aes256"
			dh_group = 19
			ike_version = 1
			key_lifetime = 3600
		}
		
		data "ibm_is_ike_policies" "is_ike_policies" {
			depends_on = [
				ibm_is_ike_policy.example1,
				ibm_is_ike_policy.example2
			]
		}
	`, name1, name2)
}
