// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMIAMAuthorizationPoliciesDataSource_Basic(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIAMAuthorizationPoliciesDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_iam_authorization_policies.testacc_ds_authorization_policies", "id"),
				),
			},
		},
	})
}

func TestAccIBMIAMAuthorizationPoliciesDataSource_Multiple_Policies(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIAMAuthorizationPoliciesDataSourceMultiplePolicies(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("ibm_iam_authorization_policy.policy", "id"),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMIAMAuthorizationPoliciesDataSourceConfigSort(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_iam_authorization_policies.testacc_ds_authorization_policies", "policies.#"),
				),
			},
		},
	})
}

func testAccCheckIBMIAMAuthorizationPoliciesDataSourceConfig() string {
	return `
		data "ibm_iam_authorization_policies" "testacc_ds_authorization_policies" {
		}`
}

func testAccCheckIBMIAMAuthorizationPoliciesDataSourceConfigSort() string {
	return `
	data "ibm_iam_authorization_policies" "testacc_ds_authorization_policies" {
		sort = "-id"
	}`
}

func testAccCheckIBMIAMAuthorizationPoliciesDataSourceMultiplePolicies() string {
	return `
		resource "ibm_iam_authorization_policy" "policy" {
			source_service_name         = "databases-for-redis"
			target_service_name         = "kms"
			roles                       = ["Reader", "Authorization Delegator"]
		}

		resource "ibm_iam_authorization_policy" "policy1" {
			source_service_name  = "is"
			source_resource_type = "load-balancer"
			target_service_name  = "cloudcerts"
			roles                = ["Reader"]
		}
		`
}
