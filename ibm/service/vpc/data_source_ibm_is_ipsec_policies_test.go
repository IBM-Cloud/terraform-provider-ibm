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

func TestAccIBMIsIpsecPoliciesDataSourceBasic(t *testing.T) {
	name1 := fmt.Sprintf("tfipsec-name1-%d", acctest.RandIntRange(10, 100))
	name2 := fmt.Sprintf("tfipsec-name2-%d", acctest.RandIntRange(10, 100))
	dataSourceName := "data.ibm_is_ipsec_policies.is_ipsec_policies"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIsIpsecPoliciesDataSourceConfig(name1, name2),
				Check: resource.ComposeTestCheckFunc(
					// Verify the list of policies
					resource.TestCheckResourceAttrSet(dataSourceName, "ipsec_policies.#"),

					// Verify all fields in the first policy entry
					resource.TestCheckResourceAttrSet(dataSourceName, "ipsec_policies.0.id"),
					resource.TestCheckResourceAttrSet(dataSourceName, "ipsec_policies.0.name"),
					resource.TestCheckResourceAttrSet(dataSourceName, "ipsec_policies.0.authentication_algorithm"),
					resource.TestCheckResourceAttrSet(dataSourceName, "ipsec_policies.0.encryption_algorithm"),
					resource.TestCheckResourceAttrSet(dataSourceName, "ipsec_policies.0.pfs"),
					resource.TestCheckResourceAttrSet(dataSourceName, "ipsec_policies.0.key_lifetime"),
					resource.TestCheckResourceAttrSet(dataSourceName, "ipsec_policies.0.encapsulation_mode"),
					resource.TestCheckResourceAttrSet(dataSourceName, "ipsec_policies.0.transform_protocol"),
					resource.TestCheckResourceAttrSet(dataSourceName, "ipsec_policies.0.href"),
					resource.TestCheckResourceAttrSet(dataSourceName, "ipsec_policies.0.created_at"),
					resource.TestCheckResourceAttrSet(dataSourceName, "ipsec_policies.0.resource_type"),

					// Verify connections array exists (may be empty)
					resource.TestCheckResourceAttrSet(dataSourceName, "ipsec_policies.0.connections.#"),

					// Verify resource_group nested structure
					resource.TestCheckResourceAttrSet(dataSourceName, "ipsec_policies.0.resource_group.#"),
					resource.TestCheckResourceAttrSet(dataSourceName, "ipsec_policies.0.resource_group.0.id"),
					resource.TestCheckResourceAttrSet(dataSourceName, "ipsec_policies.0.resource_group.0.name"),
					resource.TestCheckResourceAttrSet(dataSourceName, "ipsec_policies.0.resource_group.0.href"),

					// Verify specific expected values for fixed fields
					resource.TestCheckResourceAttr(dataSourceName, "ipsec_policies.0.encapsulation_mode", "tunnel"),
					resource.TestCheckResourceAttr(dataSourceName, "ipsec_policies.0.transform_protocol", "esp"),
				),
			},
		},
	})
}

func testAccCheckIBMIsIpsecPoliciesDataSourceConfig(name1, name2 string) string {
	return fmt.Sprintf(`
		// Create two policies with different configurations to ensure we test various values
		resource "ibm_is_ipsec_policy" "example1" {
			name = "%s"
			authentication_algorithm = "sha256"
			encryption_algorithm = "aes128"
			pfs = "group_14"
			key_lifetime = 3600
		}
		
		resource "ibm_is_ipsec_policy" "example2" {
			name = "%s"
			authentication_algorithm = "sha512"
			encryption_algorithm = "aes256"
			pfs = "group_19"
			key_lifetime = 7200
		}
		
		data "ibm_is_ipsec_policies" "is_ipsec_policies" {
			depends_on = [
				ibm_is_ipsec_policy.example1,
				ibm_is_ipsec_policy.example2
			]
		}
	`, name1, name2)
}
