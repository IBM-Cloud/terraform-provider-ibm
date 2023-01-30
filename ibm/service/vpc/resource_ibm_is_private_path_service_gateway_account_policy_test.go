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

func TestAccIBMIsPrivatePathServiceGatewayAccountPolicyBasic(t *testing.T) {
	var conf vpcv1.PrivatePathServiceGatewayAccountPolicy
	privatePathServiceGatewayID := fmt.Sprintf("tf_private_path_service_gateway_id_%d", acctest.RandIntRange(10, 100))
	accessPolicy := "deny"
	accessPolicyUpdate := "review"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIsPrivatePathServiceGatewayAccountPolicyDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsPrivatePathServiceGatewayAccountPolicyConfigBasic(privatePathServiceGatewayID, accessPolicy),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIsPrivatePathServiceGatewayAccountPolicyExists("ibm_is_private_path_service_gateway_account_policy.is_private_path_service_gateway_account_policy", conf),
					resource.TestCheckResourceAttr("ibm_is_private_path_service_gateway_account_policy.is_private_path_service_gateway_account_policy", "private_path_service_gateway_id", privatePathServiceGatewayID),
					resource.TestCheckResourceAttr("ibm_is_private_path_service_gateway_account_policy.is_private_path_service_gateway_account_policy", "access_policy", accessPolicy),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMIsPrivatePathServiceGatewayAccountPolicyConfigBasic(privatePathServiceGatewayID, accessPolicyUpdate),
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

func testAccCheckIBMIsPrivatePathServiceGatewayAccountPolicyConfigBasic(privatePathServiceGatewayID string, accessPolicy string) string {
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

func testAccCheckIBMIsPrivatePathServiceGatewayAccountPolicyExists(n string, obj vpcv1.PrivatePathServiceGatewayAccountPolicy) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		vpcClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).VpcV1()
		if err != nil {
			return err
		}

		getPrivatePathServiceGatewayAccountPolicyOptions := &vpcv1.GetPrivatePathServiceGatewayAccountPolicyOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getPrivatePathServiceGatewayAccountPolicyOptions.SetPrivatePathServiceGatewayID(parts[0])
		getPrivatePathServiceGatewayAccountPolicyOptions.SetID(parts[1])

		privatePathServiceGatewayAccountPolicy, _, err := vpcClient.GetPrivatePathServiceGatewayAccountPolicy(getPrivatePathServiceGatewayAccountPolicyOptions)
		if err != nil {
			return err
		}

		obj = *privatePathServiceGatewayAccountPolicy
		return nil
	}
}

func testAccCheckIBMIsPrivatePathServiceGatewayAccountPolicyDestroy(s *terraform.State) error {
	vpcClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).VpcV1()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_is_private_path_service_gateway_account_policy" {
			continue
		}

		getPrivatePathServiceGatewayAccountPolicyOptions := &vpcv1.GetPrivatePathServiceGatewayAccountPolicyOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getPrivatePathServiceGatewayAccountPolicyOptions.SetPrivatePathServiceGatewayID(parts[0])
		getPrivatePathServiceGatewayAccountPolicyOptions.SetID(parts[1])

		// Try to find the key
		_, response, err := vpcClient.GetPrivatePathServiceGatewayAccountPolicy(getPrivatePathServiceGatewayAccountPolicyOptions)

		if err == nil {
			return fmt.Errorf("PrivatePathServiceGatewayAccountPolicy still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for PrivatePathServiceGatewayAccountPolicy (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
