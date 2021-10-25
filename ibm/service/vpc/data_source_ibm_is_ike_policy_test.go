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

func TestAccIBMIsIkePolicyDataSourceBasic(t *testing.T) {
	name := fmt.Sprintf("tfike-name-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIsIkePolicyDataSourceConfigBasic(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_ike_policy.is_ike_policy", "authentication_algorithm"),
					resource.TestCheckResourceAttrSet("data.ibm_is_ike_policy.is_ike_policy", "connections.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_ike_policy.is_ike_policy", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_is_ike_policy.is_ike_policy", "dh_group"),
					resource.TestCheckResourceAttrSet("data.ibm_is_ike_policy.is_ike_policy", "encryption_algorithm"),
					resource.TestCheckResourceAttrSet("data.ibm_is_ike_policy.is_ike_policy", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_ike_policy.is_ike_policy", "ike_version"),
					resource.TestCheckResourceAttrSet("data.ibm_is_ike_policy.is_ike_policy", "key_lifetime"),
					resource.TestCheckResourceAttrSet("data.ibm_is_ike_policy.is_ike_policy", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_is_ike_policy.is_ike_policy", "negotiation_mode"),
					resource.TestCheckResourceAttrSet("data.ibm_is_ike_policy.is_ike_policy", "resource_group.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_ike_policy.is_ike_policy", "resource_type"),
				),
			},
			{
				Config: testAccCheckIBMIsIkePolicyDataSourceConfigBasic(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_ike_policy.is_ike_policy1", "authentication_algorithm"),
					resource.TestCheckResourceAttrSet("data.ibm_is_ike_policy.is_ike_policy1", "connections.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_ike_policy.is_ike_policy1", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_is_ike_policy.is_ike_policy1", "dh_group"),
					resource.TestCheckResourceAttrSet("data.ibm_is_ike_policy.is_ike_policy1", "encryption_algorithm"),
					resource.TestCheckResourceAttrSet("data.ibm_is_ike_policy.is_ike_policy1", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_ike_policy.is_ike_policy1", "ike_version"),
					resource.TestCheckResourceAttrSet("data.ibm_is_ike_policy.is_ike_policy1", "key_lifetime"),
					resource.TestCheckResourceAttrSet("data.ibm_is_ike_policy.is_ike_policy1", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_is_ike_policy.is_ike_policy1", "negotiation_mode"),
					resource.TestCheckResourceAttrSet("data.ibm_is_ike_policy.is_ike_policy1", "resource_group.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_ike_policy.is_ike_policy1", "resource_type"),
				),
			},
		},
	})
}

func testAccCheckIBMIsIkePolicyDataSourceConfigBasic(name string) string {
	return fmt.Sprintf(`
    	resource "ibm_is_ike_policy" "example" {
    		name = "%s"
    		authentication_algorithm = "sha1"
    		encryption_algorithm = "aes128"
    		dh_group = 5
    		ike_version = 2
    		key_lifetime = 1800
    	}
		data "ibm_is_ike_policy" "is_ike_policy" {
			name = ibm_is_ike_policy.example.name
		}
		data "ibm_is_ike_policy" "is_ike_policy1" {
			ike_policy = ibm_is_ike_policy.example.id
		}
	`, name)
}
