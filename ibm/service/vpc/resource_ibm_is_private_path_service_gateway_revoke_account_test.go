// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/vpc-go-sdk/vpcv1"
)

func TestAccIBMIsPrivatePathServiceGatewayRevokeAccountBasic(t *testing.T) {
	var conf vpcv1.PrivatePathServiceGatewayRevokeAccount
	privatePathServiceGatewayID := fmt.Sprintf("tf_private_path_service_gateway_id_%d", acctest.RandIntRange(10, 100))
	accessPolicy := "deny"
	accessPolicyUpdate := "review"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIsPrivatePathServiceGatewayRevokeAccountDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsPrivatePathServiceGatewayRevokeAccountConfigBasic(privatePathServiceGatewayID, accessPolicy),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIsPrivatePathServiceGatewayRevokeAccountExists("ibm_is_private_path_service_gateway_account_policy.is_private_path_service_gateway_account_policy", conf),
					resource.TestCheckResourceAttr("ibm_is_private_path_service_gateway_account_policy.is_private_path_service_gateway_account_policy", "private_path_service_gateway_id", privatePathServiceGatewayID),
					resource.TestCheckResourceAttr("ibm_is_private_path_service_gateway_account_policy.is_private_path_service_gateway_account_policy", "access_policy", accessPolicy),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMIsPrivatePathServiceGatewayRevokeAccountConfigBasic(privatePathServiceGatewayID, accessPolicyUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_is_private_path_service_gateway_account_policy.is_private_path_service_gateway_account_policy", "private_path_service_gateway_id", privatePathServiceGatewayID),
					resource.TestCheckResourceAttr("ibm_is_private_path_service_gateway_account_policy.is_private_path_service_gateway_account_policy", "access_policy", accessPolicyUpdate),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_is_private_path_service_gateway_account_policy.is_private_path_service_gateway_account_policy",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIBMIsPrivatePathServiceGatewayRevokeAccountConfigBasic(privatePathServiceGatewayID string, accessPolicy string) string {
	return fmt.Sprintf(`

		resource "ibm_is_private_path_service_gateway_account_policy" "is_private_path_service_gateway_account_policy_instance" {
			private_path_service_gateway_id = "%s"
			access_policy = "%s"
			account {
				id = "fee82deba12e4c0fb69c3b09d1f12345"
			}
		}
	`, privatePathServiceGatewayID, accessPolicy)
}

func testAccCheckIBMIsPrivatePathServiceGatewayRevokeAccountExists(n string, obj vpcv1.PrivatePathServiceGatewayRevokeAccount) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		vpcClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).VpcV1API()
		if err != nil {
			return err
		}

		getPrivatePathServiceGatewayRevokeAccountOptions := &vpcv1.GetPrivatePathServiceGatewayRevokeAccountOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getPrivatePathServiceGatewayRevokeAccountOptions.SetPrivatePathServiceGatewayID(parts[0])
		getPrivatePathServiceGatewayRevokeAccountOptions.SetID(parts[1])

		privatePathServiceGatewayRevokeAccount, _, err := vpcClient.GetPrivatePathServiceGatewayRevokeAccount(getPrivatePathServiceGatewayRevokeAccountOptions)
		if err != nil {
			return err
		}

		obj = *privatePathServiceGatewayRevokeAccount
		return nil
	}
}

func testAccCheckIBMIsPrivatePathServiceGatewayRevokeAccountDestroy(s *terraform.State) error {
	vpcClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).VpcV1API()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_is_private_path_service_gateway_account_policy" {
			continue
		}

		getPrivatePathServiceGatewayRevokeAccountOptions := &vpcv1.GetPrivatePathServiceGatewayRevokeAccountOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getPrivatePathServiceGatewayRevokeAccountOptions.SetPrivatePathServiceGatewayID(parts[0])
		getPrivatePathServiceGatewayRevokeAccountOptions.SetID(parts[1])

		// Try to find the key
		_, response, err := vpcClient.GetPrivatePathServiceGatewayRevokeAccount(getPrivatePathServiceGatewayRevokeAccountOptions)

		if err == nil {
			return fmt.Errorf("PrivatePathServiceGatewayRevokeAccount still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for PrivatePathServiceGatewayRevokeAccount (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
