// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIBMIsPrivatePathServiceGatewayAccountPoliciesDataSourceBasic(t *testing.T) {
	accessPolicy1 := "deny"
	accessPolicy := "deny"
	vpcname := fmt.Sprintf("tflb-vpc-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tflb-subnet-name-%d", acctest.RandIntRange(10, 100))
	lbname := fmt.Sprintf("tf-test-lb%dd", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf-test-ppsg%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsPrivatePathServiceGatewayAccountPoliciesDataSourceConfigBasic(vpcname, subnetname, acc.ISZoneName, acc.ISCIDR, lbname, accessPolicy, name, accessPolicy1),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_private_path_service_gateway_account_policies.is_private_path_service_gateway_account_policies", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_private_path_service_gateway_account_policies.is_private_path_service_gateway_account_policies", "account_policies.#"),
					resource.TestCheckResourceAttr("data.ibm_is_private_path_service_gateway_account_policies.is_private_path_service_gateway_account_policies", "account_policies.0.access_policy", accessPolicy1),
					resource.TestCheckResourceAttrSet("data.ibm_is_private_path_service_gateway_account_policies.is_private_path_service_gateway_account_policies", "account_policies.0.account.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_private_path_service_gateway_account_policies.is_private_path_service_gateway_account_policies", "account_policies.0.account.0.id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_private_path_service_gateway_account_policies.is_private_path_service_gateway_account_policies", "account_policies.0.account.0.resource_type"),
					resource.TestCheckResourceAttrSet("data.ibm_is_private_path_service_gateway_account_policies.is_private_path_service_gateway_account_policies", "account_policies.0.created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_is_private_path_service_gateway_account_policies.is_private_path_service_gateway_account_policies", "account_policies.0.id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_private_path_service_gateway_account_policies.is_private_path_service_gateway_account_policies", "account_policies.0.href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_private_path_service_gateway_account_policies.is_private_path_service_gateway_account_policies", "account_policies.0.resource_type"),
				),
			},
		},
	})
}

func testAccCheckIBMIsPrivatePathServiceGatewayAccountPoliciesDataSourceConfigBasic(vpcname, subnetname, zone, cidr, lbname, accessPolicy, name, accessPolicy1 string) string {
	return testAccCheckIBMIsPrivatePathServiceGatewayConfigBasic(vpcname, subnetname, acc.ISZoneName, acc.ISCIDR, lbname, accessPolicy, name) + fmt.Sprintf(`
		resource "ibm_is_private_path_service_gateway_account_policy" "is_private_path_service_gateway_account_policy" {
			private_path_service_gateway = ibm_is_private_path_service_gateway.is_private_path_service_gateway.id
			access_policy = "%s"
			account = "%s"
		}

		data "ibm_is_private_path_service_gateway_account_policies" "is_private_path_service_gateway_account_policies" {
			private_path_service_gateway = ibm_is_private_path_service_gateway.is_private_path_service_gateway.id
		}
	`, accessPolicy1, acc.AccountId)
}
