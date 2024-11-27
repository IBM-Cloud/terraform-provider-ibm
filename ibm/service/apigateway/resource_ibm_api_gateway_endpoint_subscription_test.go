// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0
package apigateway_test

import (
	"fmt"
	"strings"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"

	apigatewaysdk "github.com/IBM/apigateway-go-sdk/apigatewaycontrollerapiv1"
	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccIBMAPIGatewaySubscription_Basic(t *testing.T) {
	var resultSubscription apigatewaysdk.V2Subscription
	name := fmt.Sprintf("tftest-%s", acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMAPIGatewaySubscriptionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMAPIGatewaySubscriptionBasic(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMAPIGatewaySubscriptionExists("ibm_api_gateway_endpoint_subscription.subscription", resultSubscription),
					resource.TestCheckResourceAttr("ibm_api_gateway_endpoint_subscription.subscription", "name", name),
				),
			},
		},
	})
}

func testAccCheckIBMAPIGatewaySubscriptionDestroy(s *terraform.State) error {
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_api_gateway_endpoint_subscription" {
			continue
		}
		apiClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).APIGateway()
		if err != nil {
			return err
		}
		sess, err := acc.TestAccProvider.Meta().(conns.ClientSession).BluemixSession()
		if err != nil {
			return err
		}
		parts := rs.Primary.ID
		partslist := strings.Split(parts, "//")
		ArtifactID := partslist[0]
		ClientID := partslist[1]
		oauthtoken := sess.Config.IAMAccessToken
		oauthtoken = strings.Replace(oauthtoken, "Bearer ", "", -1)

		payload := apigatewaysdk.GetSubscriptionOptions{
			ArtifactID:    &ArtifactID,
			ID:            &ClientID,
			Authorization: &oauthtoken,
		}
		_, _, err = apiClient.GetSubscription(&payload)
		if err != nil && !strings.Contains(err.Error(), "Not Found") {
			return fmt.Errorf("[ERROR] Error checking if instance (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}
	return nil
}
func TestAccIBMAPIGatewaySubscriptionImport(t *testing.T) {
	var resultSubscription apigatewaysdk.V2Subscription
	name := fmt.Sprintf("tftest-%s", acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMAPIGatewaySubscriptionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMAPIGatewaySubscriptionBasic(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMAPIGatewaySubscriptionExists("ibm_api_gateway_endpoint_subscription.subscription", resultSubscription),
					resource.TestCheckResourceAttr("ibm_api_gateway_endpoint_subscription.subscription", "name", name),
				),
			},
			{
				ResourceName:      "ibm_api_gateway_endpoint_subscription.subscription",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIBMAPIGatewaySubscriptionBasic(name string) string {
	return fmt.Sprintf(`
	resource "ibm_resource_instance" "apigateway"{
	name     = "testname"
	location = "global"
	service  = "api-gateway"
	plan     = "lite"
	}
	resource "ibm_api_gateway_endpoint" "endpoint"{
	service_instance_crn = ibm_resource_instance.apigateway.id
	name="test-endpoint"
	managed="true"
	open_api_doc_name = "../../test-fixtures/SDK-test.json"
	}
	resource "ibm_api_gateway_endpoint_subscription" "subscription"{
		artifact_id = ibm_api_gateway_endpoint.endpoint.endpoint_id
		client_id   = "test1234"
		name        = "%s"
		type        = "external"
	}
	  `, name)
}

func testAccCheckIBMAPIGatewaySubscriptionExists(n string, result apigatewaysdk.V2Subscription) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}
		apiClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).APIGateway()
		if err != nil {
			return err
		}
		sess, err := acc.TestAccProvider.Meta().(conns.ClientSession).BluemixSession()
		if err != nil {
			return err
		}
		parts := rs.Primary.ID
		partslist := strings.Split(parts, "//")
		ArtifactID := partslist[0]
		ClientID := partslist[1]
		oauthtoken := sess.Config.IAMAccessToken
		oauthtoken = strings.Replace(oauthtoken, "Bearer ", "", -1)

		payload := apigatewaysdk.GetSubscriptionOptions{
			ArtifactID:    &ArtifactID,
			ID:            &ClientID,
			Authorization: &oauthtoken,
		}
		endpoint, _, err := apiClient.GetSubscription(&payload)
		if err != nil {
			return err
		}

		result = *endpoint
		return nil
	}

}
