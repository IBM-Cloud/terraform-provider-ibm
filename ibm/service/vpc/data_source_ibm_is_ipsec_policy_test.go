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

func TestAccIBMIsIpsecPolicyDataSourceBasic(t *testing.T) {
	name := fmt.Sprintf("tfipsec-data-%d", acctest.RandIntRange(10, 100))
	resourceKey := "ibm_is_ipsec_policy.example"
	dataSourceNameKey := "data.ibm_is_ipsec_policy.by_name"
	dataSourceIDKey := "data.ibm_is_ipsec_policy.by_id"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIsIpsecPolicyDataSourceConfig(name),
				Check: resource.ComposeTestCheckFunc(
					// Check when lookup is by name
					resource.TestCheckResourceAttrPair(dataSourceNameKey, "id", resourceKey, "id"),
					resource.TestCheckResourceAttrPair(dataSourceNameKey, "name", resourceKey, "name"),
					resource.TestCheckResourceAttrPair(dataSourceNameKey, "authentication_algorithm", resourceKey, "authentication_algorithm"),
					resource.TestCheckResourceAttrPair(dataSourceNameKey, "encryption_algorithm", resourceKey, "encryption_algorithm"),
					resource.TestCheckResourceAttrPair(dataSourceNameKey, "pfs", resourceKey, "pfs"),
					resource.TestCheckResourceAttrPair(dataSourceNameKey, "key_lifetime", resourceKey, "key_lifetime"),

					// Check computed fields for name lookup
					resource.TestCheckResourceAttrSet(dataSourceNameKey, "connections.#"),
					resource.TestCheckResourceAttrSet(dataSourceNameKey, "created_at"),
					resource.TestCheckResourceAttrSet(dataSourceNameKey, "href"),
					resource.TestCheckResourceAttrSet(dataSourceNameKey, "encapsulation_mode"),
					resource.TestCheckResourceAttrSet(dataSourceNameKey, "transform_protocol"),
					resource.TestCheckResourceAttrSet(dataSourceNameKey, "resource_group.#"),
					resource.TestCheckResourceAttrSet(dataSourceNameKey, "resource_type"),

					// Check specific values for computed fields
					resource.TestCheckResourceAttr(dataSourceNameKey, "encapsulation_mode", "tunnel"),
					resource.TestCheckResourceAttr(dataSourceNameKey, "transform_protocol", "esp"),

					// Check resource_group nested fields
					resource.TestCheckResourceAttrSet(dataSourceNameKey, "resource_group.0.id"),
					resource.TestCheckResourceAttrSet(dataSourceNameKey, "resource_group.0.name"),
					resource.TestCheckResourceAttrSet(dataSourceNameKey, "resource_group.0.href"),

					// Check when lookup is by ID
					resource.TestCheckResourceAttrPair(dataSourceIDKey, "id", resourceKey, "id"),
					resource.TestCheckResourceAttrPair(dataSourceIDKey, "name", resourceKey, "name"),
					resource.TestCheckResourceAttrPair(dataSourceIDKey, "authentication_algorithm", resourceKey, "authentication_algorithm"),
					resource.TestCheckResourceAttrPair(dataSourceIDKey, "encryption_algorithm", resourceKey, "encryption_algorithm"),
					resource.TestCheckResourceAttrPair(dataSourceIDKey, "pfs", resourceKey, "pfs"),
					resource.TestCheckResourceAttrPair(dataSourceIDKey, "key_lifetime", resourceKey, "key_lifetime"),

					// Check computed fields for ID lookup
					resource.TestCheckResourceAttrSet(dataSourceIDKey, "connections.#"),
					resource.TestCheckResourceAttrSet(dataSourceIDKey, "created_at"),
					resource.TestCheckResourceAttrSet(dataSourceIDKey, "href"),
					resource.TestCheckResourceAttrSet(dataSourceIDKey, "encapsulation_mode"),
					resource.TestCheckResourceAttrSet(dataSourceIDKey, "transform_protocol"),
					resource.TestCheckResourceAttrSet(dataSourceIDKey, "resource_group.#"),
					resource.TestCheckResourceAttrSet(dataSourceIDKey, "resource_type"),
				),
			},
		},
	})
}

func TestAccIBMIsIpsecPolicyDataSourceMultipleAlgorithms(t *testing.T) {
	name := fmt.Sprintf("tfipsec-data-multi-%d", acctest.RandIntRange(10, 100))
	resourceKey := "ibm_is_ipsec_policy.example"
	dataSourceNameKey := "data.ibm_is_ipsec_policy.by_name"
	dataSourceIDKey := "data.ibm_is_ipsec_policy.by_id"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIsIpsecPolicyDataSourceMultipleAlgorithmsConfig(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceNameKey, "id", resourceKey, "id"),
					resource.TestCheckResourceAttrPair(dataSourceNameKey, "name", resourceKey, "name"),
					resource.TestCheckResourceAttrPair(dataSourceNameKey, "authentication_algorithms.#", resourceKey, "authentication_algorithms.#"),
					resource.TestCheckResourceAttrPair(dataSourceNameKey, "authentication_algorithms.0", resourceKey, "authentication_algorithms.0"),
					resource.TestCheckResourceAttrPair(dataSourceNameKey, "authentication_algorithms.1", resourceKey, "authentication_algorithms.1"),
					resource.TestCheckResourceAttrPair(dataSourceNameKey, "encryption_algorithms.#", resourceKey, "encryption_algorithms.#"),
					resource.TestCheckResourceAttrPair(dataSourceNameKey, "encryption_algorithms.0", resourceKey, "encryption_algorithms.0"),
					resource.TestCheckResourceAttrPair(dataSourceNameKey, "encryption_algorithms.1", resourceKey, "encryption_algorithms.1"),
					resource.TestCheckResourceAttrPair(dataSourceNameKey, "pfs_groups.#", resourceKey, "pfs_groups.#"),
					resource.TestCheckResourceAttrPair(dataSourceNameKey, "pfs_groups.0", resourceKey, "pfs_groups.0"),
					resource.TestCheckResourceAttrPair(dataSourceNameKey, "pfs_groups.1", resourceKey, "pfs_groups.1"),
					resource.TestCheckResourceAttrPair(dataSourceIDKey, "authentication_algorithms.#", resourceKey, "authentication_algorithms.#"),
					resource.TestCheckResourceAttrPair(dataSourceIDKey, "authentication_algorithms.0", resourceKey, "authentication_algorithms.0"),
					resource.TestCheckResourceAttrPair(dataSourceIDKey, "authentication_algorithms.1", resourceKey, "authentication_algorithms.1"),
					resource.TestCheckResourceAttrPair(dataSourceIDKey, "encryption_algorithms.#", resourceKey, "encryption_algorithms.#"),
					resource.TestCheckResourceAttrPair(dataSourceIDKey, "encryption_algorithms.0", resourceKey, "encryption_algorithms.0"),
					resource.TestCheckResourceAttrPair(dataSourceIDKey, "encryption_algorithms.1", resourceKey, "encryption_algorithms.1"),
					resource.TestCheckResourceAttrPair(dataSourceIDKey, "pfs_groups.#", resourceKey, "pfs_groups.#"),
					resource.TestCheckResourceAttrPair(dataSourceIDKey, "pfs_groups.0", resourceKey, "pfs_groups.0"),
					resource.TestCheckResourceAttrPair(dataSourceIDKey, "pfs_groups.1", resourceKey, "pfs_groups.1"),
				),
			},
		},
	})
}

func testAccCheckIBMIsIpsecPolicyDataSourceConfig(name string) string {
	return fmt.Sprintf(`
		resource "ibm_is_ipsec_policy" "example" {
			name = "%s"
			authentication_algorithm = "sha256"
			encryption_algorithm = "aes128"
			pfs = "group_14"
			key_lifetime = 3600
		}
		
		data "ibm_is_ipsec_policy" "by_name" {
			name = ibm_is_ipsec_policy.example.name
			depends_on = [ibm_is_ipsec_policy.example]
		}
		
		data "ibm_is_ipsec_policy" "by_id" {
			ipsec_policy = ibm_is_ipsec_policy.example.id
			depends_on = [ibm_is_ipsec_policy.example]
		}
	`, name)
}

func testAccCheckIBMIsIpsecPolicyDataSourceMultipleAlgorithmsConfig(name string) string {
	return fmt.Sprintf(`
		resource "ibm_is_ipsec_policy" "example" {
			name = "%s"
			authentication_algorithms = ["sha512", "sha384"]
			encryption_algorithms = ["aes128", "aes192"]
			pfs_groups = ["group_14", "group_15"]
			key_lifetime = 3600
		}
		
		data "ibm_is_ipsec_policy" "by_name" {
			name = ibm_is_ipsec_policy.example.name
			depends_on = [ibm_is_ipsec_policy.example]
		}
		
		data "ibm_is_ipsec_policy" "by_id" {
			ipsec_policy = ibm_is_ipsec_policy.example.id
			depends_on = [ibm_is_ipsec_policy.example]
		}
	`, name)
}
