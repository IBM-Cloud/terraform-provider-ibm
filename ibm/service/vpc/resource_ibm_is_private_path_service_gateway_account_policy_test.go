// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/vpc-go-sdk/vpcv1"
)

func TestAccIBMIsPrivatePathServiceGatewayAccountPolicyBasic(t *testing.T) {
	var conf vpcv1.PrivatePathServiceGatewayAccountPolicy
	accessPolicy := "deny"
	accessPolicyUpdate := "review"
	accessPolicy1 := "review"
	vpcname := fmt.Sprintf("tflb-vpc-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tflb-subnet-name-%d", acctest.RandIntRange(10, 100))
	lbname := fmt.Sprintf("tf-test-lb%dd", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf-test-ppsg%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIsPrivatePathServiceGatewayAccountPolicyDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsPrivatePathServiceGatewayAccountPolicyConfigBasic(vpcname, subnetname, acc.ISZoneName, acc.ISCIDR, lbname, accessPolicy, name, accessPolicy1),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIsPrivatePathServiceGatewayAccountPolicyExists("ibm_is_private_path_service_gateway_account_policy.is_private_path_service_gateway_account_policy", conf),
					resource.TestCheckResourceAttr("ibm_is_private_path_service_gateway_account_policy.is_private_path_service_gateway_account_policy", "access_policy", accessPolicy1),
					resource.TestCheckResourceAttrSet("data.ibm_is_private_path_service_gateway_account_policy.is_private_path_service_gateway_account_policy", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_private_path_service_gateway_account_policy.is_private_path_service_gateway_account_policy", "private_path_service_gateway"),
					resource.TestCheckResourceAttrSet("data.ibm_is_private_path_service_gateway_account_policy.is_private_path_service_gateway_account_policy", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_private_path_service_gateway_account_policy.is_private_path_service_gateway_account_policy", "access_policy"),
					resource.TestCheckResourceAttrSet("data.ibm_is_private_path_service_gateway_account_policy.is_private_path_service_gateway_account_policy", "account.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_private_path_service_gateway_account_policy.is_private_path_service_gateway_account_policy", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_is_private_path_service_gateway_account_policy.is_private_path_service_gateway_account_policy", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_private_path_service_gateway_account_policy.is_private_path_service_gateway_account_policy", "private_path_service_gateway_account_policy"),
					resource.TestCheckResourceAttrSet("data.ibm_is_private_path_service_gateway_account_policy.is_private_path_service_gateway_account_policy", "resource_type"),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMIsPrivatePathServiceGatewayAccountPolicyConfigBasic(vpcname, subnetname, acc.ISZoneName, acc.ISCIDR, lbname, accessPolicy, name, accessPolicyUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_is_private_path_service_gateway_account_policy.is_private_path_service_gateway_account_policy", "access_policy", accessPolicyUpdate),
					resource.TestCheckResourceAttrSet("data.ibm_is_private_path_service_gateway_account_policy.is_private_path_service_gateway_account_policy", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_private_path_service_gateway_account_policy.is_private_path_service_gateway_account_policy", "private_path_service_gateway"),
					resource.TestCheckResourceAttrSet("data.ibm_is_private_path_service_gateway_account_policy.is_private_path_service_gateway_account_policy", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_private_path_service_gateway_account_policy.is_private_path_service_gateway_account_policy", "access_policy"),
					resource.TestCheckResourceAttrSet("data.ibm_is_private_path_service_gateway_account_policy.is_private_path_service_gateway_account_policy", "account.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_private_path_service_gateway_account_policy.is_private_path_service_gateway_account_policy", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_is_private_path_service_gateway_account_policy.is_private_path_service_gateway_account_policy", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_private_path_service_gateway_account_policy.is_private_path_service_gateway_account_policy", "private_path_service_gateway_account_policy"),
					resource.TestCheckResourceAttrSet("data.ibm_is_private_path_service_gateway_account_policy.is_private_path_service_gateway_account_policy", "resource_type"),
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

func testAccCheckIBMIsPrivatePathServiceGatewayAccountPolicyConfigBasic(vpcname, subnetname, zone, cidr, lbname, accessPolicy, name, accessPolicy1 string) string {
	return testAccCheckIBMIsPrivatePathServiceGatewayConfigBasic(vpcname, subnetname, acc.ISZoneName, acc.ISCIDR, lbname, accessPolicy, name) + fmt.Sprintf(`

		resource "ibm_is_private_path_service_gateway_account_policy" "is_private_path_service_gateway_account_policy" {
			private_path_service_gateway = ibm_is_private_path_service_gateway.is_private_path_service_gateway.id
			access_policy = "%s"
			account = "%s"
			
		}
	`, accessPolicy1, acc.AccountId)
}

func testAccCheckIBMIsPrivatePathServiceGatewayAccountPolicyExists(n string, obj vpcv1.PrivatePathServiceGatewayAccountPolicy) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		vpcClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).VpcV1API()
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
	vpcClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).VpcV1API()
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
