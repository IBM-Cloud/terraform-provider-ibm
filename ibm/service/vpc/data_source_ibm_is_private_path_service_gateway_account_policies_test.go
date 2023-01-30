// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIBMIsPrivatePathServiceGatewayAccountPoliciesDataSourceBasic(t *testing.T) {
	ppsgId := fmt.Sprintf("tf_private_path_service_gateway_id_%d", acctest.RandIntRange(10, 100))
	accessPolicy := "deny"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsPrivatePathServiceGatewayAccountPoliciesDataSourceConfigBasic(ppsgId, accessPolicy),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_private_path_service_gateway_account_policies.is_private_path_service_gateway_account_policies", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_private_path_service_gateway_account_policies.is_private_path_service_gateway_account_policies", "account_policies.#"),
					resource.TestCheckResourceAttr("data.ibm_is_private_path_service_gateway_account_policies.is_private_path_service_gateway_account_policies", "account_policies.0.access_policy", accessPolicy),
					resource.TestCheckResourceAttrSet("data.ibm_is_private_path_service_gateway_account_policies.is_private_path_service_gateway_account_policies", "account_policies.0.account.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_private_path_service_gateway_account_policies.is_private_path_service_gateway_account_policies", "account_policies.0.account.id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_private_path_service_gateway_account_policies.is_private_path_service_gateway_account_policies", "account_policies.0.account.resource_type"),
					resource.TestCheckResourceAttrSet("data.ibm_is_private_path_service_gateway_account_policies.is_private_path_service_gateway_account_policies", "account_policies.0.created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_is_private_path_service_gateway_account_policies.is_private_path_service_gateway_account_policies", "account_policies.0.href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_private_path_service_gateway_account_policies.is_private_path_service_gateway_account_policies", "account_policies.0.resource_type"),
				),
			},
		},
	})
}

func testAccCheckIBMIsPrivatePathServiceGatewayAccountPoliciesDataSourceConfigBasic(ppsgId, accessPolicy string) string {
	return fmt.Sprintf(`
		resource "ibm_is_private_path_service_gateway_account_policy" "is_private_path_service_gateway_account_policy_instance" {
			private_path_service_gateway = "%s"
			access_policy = "%s"
			account = "%s"
		}

		data "ibm_is_private_path_service_gateway_account_policies" "is_private_path_service_gateway_account_policies_instance" {
			private_path_service_gateway_id = ibm_is_private_path_service_gateway_account_policy.is_private_path_service_gateway_account_policy.private_path_service_gateway_id
		}
	`, ppsgId, accessPolicy, acc.AccountId)
}
