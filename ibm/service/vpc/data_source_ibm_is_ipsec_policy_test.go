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
	name := fmt.Sprintf("tfipsecc-name-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIsIpsecPolicyDataSourceConfigBasic(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_ipsec_policy.is_ipsec_policy", "authentication_algorithm"),
					resource.TestCheckResourceAttrSet("data.ibm_is_ipsec_policy.is_ipsec_policy", "connections.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_ipsec_policy.is_ipsec_policy", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_is_ipsec_policy.is_ipsec_policy", "encapsulation_mode"),
					resource.TestCheckResourceAttrSet("data.ibm_is_ipsec_policy.is_ipsec_policy", "encryption_algorithm"),
					resource.TestCheckResourceAttrSet("data.ibm_is_ipsec_policy.is_ipsec_policy", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_ipsec_policy.is_ipsec_policy", "key_lifetime"),
					resource.TestCheckResourceAttrSet("data.ibm_is_ipsec_policy.is_ipsec_policy", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_is_ipsec_policy.is_ipsec_policy", "pfs"),
					resource.TestCheckResourceAttrSet("data.ibm_is_ipsec_policy.is_ipsec_policy", "resource_group.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_ipsec_policy.is_ipsec_policy", "resource_type"),
					resource.TestCheckResourceAttrSet("data.ibm_is_ipsec_policy.is_ipsec_policy", "transform_protocol"),
				),
			},
			{
				Config: testAccCheckIBMIsIpsecPolicyDataSourceConfigBasic(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_ipsec_policy.is_ipsec_policy1", "authentication_algorithm"),
					resource.TestCheckResourceAttrSet("data.ibm_is_ipsec_policy.is_ipsec_policy1", "connections.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_ipsec_policy.is_ipsec_policy1", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_is_ipsec_policy.is_ipsec_policy1", "encapsulation_mode"),
					resource.TestCheckResourceAttrSet("data.ibm_is_ipsec_policy.is_ipsec_policy1", "encryption_algorithm"),
					resource.TestCheckResourceAttrSet("data.ibm_is_ipsec_policy.is_ipsec_policy1", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_ipsec_policy.is_ipsec_policy1", "key_lifetime"),
					resource.TestCheckResourceAttrSet("data.ibm_is_ipsec_policy.is_ipsec_policy1", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_is_ipsec_policy.is_ipsec_policy1", "pfs"),
					resource.TestCheckResourceAttrSet("data.ibm_is_ipsec_policy.is_ipsec_policy1", "resource_group.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_ipsec_policy.is_ipsec_policy1", "resource_type"),
					resource.TestCheckResourceAttrSet("data.ibm_is_ipsec_policy.is_ipsec_policy1", "transform_protocol"),
				),
			},
		},
	})
}

func testAccCheckIBMIsIpsecPolicyDataSourceConfigBasic(name string) string {
	return fmt.Sprintf(`
    	resource "ibm_is_ipsec_policy" "example" {
    		name = "%s"
    		authentication_algorithm = "sha1"
    		encryption_algorithm = "aes128"
    		pfs = "group_2"
    	}
		data "ibm_is_ipsec_policy" "is_ipsec_policy" {
			ipsec_policy = ibm_is_ipsec_policy.example.id
		}
		data "ibm_is_ipsec_policy" "is_ipsec_policy1" {
			name = ibm_is_ipsec_policy.example.name
		}
	`, name)
}
