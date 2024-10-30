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
	accessPolicy := "deny"
	accessPolicy1 := "review"
	vpcname := fmt.Sprintf("tflb-vpc-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tflb-subnet-name-%d", acctest.RandIntRange(10, 100))
	lbname := fmt.Sprintf("tf-test-lb%dd", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf-test-ppsg%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsPrivatePathServiceGatewayAccountPolicyDataSourceConfigBasic(vpcname, subnetname, acc.ISZoneName, acc.ISCIDR, lbname, accessPolicy, name, accessPolicy1),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_private_path_service_gateway_account_policy.is_private_path_service_gateway_account_policy", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_private_path_service_gateway_account_policy.is_private_path_service_gateway_account_policy", "private_path_service_gateway"),
					resource.TestCheckResourceAttrSet("data.ibm_is_private_path_service_gateway_account_policy.is_private_path_service_gateway_account_policy", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_private_path_service_gateway_account_policy.is_private_path_service_gateway_account_policy", "access_policy"),
					resource.TestCheckResourceAttrSet("data.ibm_is_private_path_service_gateway_account_policy.is_private_path_service_gateway_account_policy", "account.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_private_path_service_gateway_account_policy.is_private_path_service_gateway_account_policy", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_is_private_path_service_gateway_account_policy.is_private_path_service_gateway_account_policy", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_private_path_service_gateway_account_policy.is_private_path_service_gateway_account_policy", "resource_type"),
				),
			},
		},
	})
}

func testAccCheckIBMIsPrivatePathServiceGatewayAccountPolicyDataSourceConfigBasic(vpcname, subnetname, zone, cidr, lbname, accessPolicy, name, accessPolicy1 string) string {
	return testAccCheckIBMIsPrivatePathServiceGatewayConfigBasic(vpcname, subnetname, acc.ISZoneName, acc.ISCIDR, lbname, accessPolicy, name) + fmt.Sprintf(`
		resource "ibm_is_private_path_service_gateway_account_policy" "is_private_path_service_gateway_account_policy" {
			private_path_service_gateway = ibm_is_private_path_service_gateway.is_private_path_service_gateway.id
			access_policy = "%s"
			account = "%s"
		}
		data "ibm_is_private_path_service_gateway_account_policy" "is_private_path_service_gateway_account_policy" {
			private_path_service_gateway = ibm_is_private_path_service_gateway.is_private_path_service_gateway.id
			account_policy = ibm_is_private_path_service_gateway_account_policy.is_private_path_service_gateway_account_policy.id
		}
	`, accessPolicy1, acc.AccountId)
}
