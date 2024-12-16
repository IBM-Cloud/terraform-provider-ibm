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

func TestAccIBMAPIGatewayEndpoint_Basic(t *testing.T) {
	var resultendpoint apigatewaysdk.V2Endpoint
	name := fmt.Sprintf("tftest-%s", acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMAPIGatewayEndpointDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMAPIGatewayEndpointBasic(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMAPIGatewayEndpointExists("ibm_api_gateway_endpoint.endpoint", resultendpoint),
					resource.TestCheckResourceAttr("ibm_api_gateway_endpoint.endpoint", "name", name),
				),
			},
		},
	})
}

func testAccCheckIBMAPIGatewayEndpointDestroy(s *terraform.State) error {
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_api_gateway_endpoint" {
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
		ServiceInstanceCrn := partslist[0]
		endpointID := partslist[1]
		oauthtoken := sess.Config.IAMAccessToken
		oauthtoken = strings.Replace(oauthtoken, "Bearer ", "", -1)

		payload := apigatewaysdk.GetEndpointOptions{
			ServiceInstanceCrn: &ServiceInstanceCrn,
			ID:                 &endpointID,
			Authorization:      &oauthtoken,
		}
		_, _, err = apiClient.GetEndpoint(&payload)
		if err != nil && !strings.Contains(err.Error(), "Not Found") {
			return fmt.Errorf("[ERROR] Error checking if instance (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}
	return nil
}
func TestAccIBMAPIGatewayEndpointImport(t *testing.T) {
	var resultendpoint apigatewaysdk.V2Endpoint
	name := fmt.Sprintf("tftest-%s", acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMAPIGatewayEndpointDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMAPIGatewayEndpointBasic(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMAPIGatewayEndpointExists("ibm_api_gateway_endpoint.endpoint", resultendpoint),
					resource.TestCheckResourceAttr("ibm_api_gateway_endpoint.endpoint", "name", name),
				),
			},
			{
				ResourceName:      "ibm_api_gateway_endpoint.endpoint",
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"open_api_doc_name", "type"},
			},
		},
	})
}

func testAccCheckIBMAPIGatewayEndpointBasic(name string) string {
	return fmt.Sprintf(`
	resource "ibm_resource_instance" "apigateway"{
	name     = "testname1"
	location = "global"
	service  = "api-gateway"
	plan     = "lite"
	}
	resource "ibm_api_gateway_endpoint" "endpoint"{
	service_instance_crn= ibm_resource_instance.apigateway.id
	name = "%s"
	managed="true"
    open_api_doc_name="../../test-fixtures/SDK-test.json"
	}
	  `, name)
}

func testAccCheckIBMAPIGatewayEndpointExists(n string, result apigatewaysdk.V2Endpoint) resource.TestCheckFunc {

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
		ServiceInstanceCrn := partslist[0]
		endpointID := partslist[1]
		oauthtoken := sess.Config.IAMAccessToken
		oauthtoken = strings.Replace(oauthtoken, "Bearer ", "", -1)

		payload := apigatewaysdk.GetEndpointOptions{
			ServiceInstanceCrn: &ServiceInstanceCrn,
			ID:                 &endpointID,
			Authorization:      &oauthtoken,
		}
		endpoint, _, err := apiClient.GetEndpoint(&payload)
		if err != nil {
			return err
		}

		result = *endpoint
		return nil
	}

}
