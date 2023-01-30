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

func TestAccIBMIsPrivatePathServiceGatewayAccountPolicyDataSourceBasic(t *testing.T) {
	privatePathServiceGatewayAccountPolicyPrivatePathServiceGatewayID := fmt.Sprintf("tf_private_path_service_gateway_id_%d", acctest.RandIntRange(10, 100))
	privatePathServiceGatewayAccountPolicyAccessPolicy := "deny"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsPrivatePathServiceGatewayAccountPolicyDataSourceConfigBasic(privatePathServiceGatewayAccountPolicyPrivatePathServiceGatewayID, privatePathServiceGatewayAccountPolicyAccessPolicy),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_private_path_service_gateway_account_policy.is_private_path_service_gateway_account_policy", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_private_path_service_gateway_account_policy.is_private_path_service_gateway_account_policy", "private_path_service_gateway_id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_private_path_service_gateway_account_policy.is_private_path_service_gateway_account_policy", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_private_path_service_gateway_account_policy.is_private_path_service_gateway_account_policy", "access_policy"),
					resource.TestCheckResourceAttrSet("data.ibm_is_private_path_service_gateway_account_policy.is_private_path_service_gateway_account_policy", "account.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_private_path_service_gateway_account_policy.is_private_path_service_gateway_account_policy", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_is_private_path_service_gateway_account_policy.is_private_path_service_gateway_account_policy", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_private_path_service_gateway_account_policy.is_private_path_service_gateway_account_policy", "private_path_service_gateway_account_policy_id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_private_path_service_gateway_account_policy.is_private_path_service_gateway_account_policy", "resource_type"),
				),
			},
		},
	})
}

func testAccCheckIBMIsPrivatePathServiceGatewayAccountPolicyDataSourceConfigBasic(privatePathServiceGatewayAccountPolicyPrivatePathServiceGatewayID string, privatePathServiceGatewayAccountPolicyAccessPolicy string) string {
	return fmt.Sprintf(`
		resource "ibm_is_private_path_service_gateway_account_policy" "is_private_path_service_gateway_account_policy_instance" {
			private_path_service_gateway_id = "%s"
			access_policy = "%s"
			account {
				id = "fee82deba12e4c0fb69c3b09d1f12345"
			}
		}

		data "ibm_is_private_path_service_gateway_account_policy" "is_private_path_service_gateway_account_policy_instance" {
			private_path_service_gateway_id = ibm_is_private_path_service_gateway_account_policy.is_private_path_service_gateway_account_policy.private_path_service_gateway_id
			account_id = "id"
		}
	`, privatePathServiceGatewayAccountPolicyPrivatePathServiceGatewayID, privatePathServiceGatewayAccountPolicyAccessPolicy)
}
