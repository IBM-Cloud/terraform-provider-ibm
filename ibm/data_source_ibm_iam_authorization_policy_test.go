// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMIAMAuthorizationPolicyDataSource_Basic(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIAMAuthorizationPolicyDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ibm_iam_authorization_policy.testacc_ds_authorization_policy", "policies.#", "1"),
				),
			},
		},
	})
}

func TestAccIBMIAMAuthorizationPolicyDataSource_Multiple_Policies(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIAMAuthorizationPolicyDataSourceMultiplePolicies(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_iam_authorization_policy.testacc_ds_authorization_policy", "policies.#"),
				),
			},
		},
	})
}

func testAccCheckIBMIAMAuthorizationPolicyDataSourceConfig() string {
	return `
		resource "ibm_iam_authorization_policy" "policy" {
			source_service_name         = "databases-for-redis"
			target_service_name         = "kms"
			roles                       = ["Reader", "Authorization Delegator"]
		}

		data "ibm_iam_authorization_policy" "testacc_ds_authorization_policy" {
			account_id = ""
			id = ibm_iam_authorization_policy.policy.id
		}`
}

func testAccCheckIBMIAMAuthorizationPolicyDataSourceMultiplePolicies() string {
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

		data "ibm_iam_authorization_policy" "testacc_ds_authorization_policy" {
			sort = "-id"
		}`
}
