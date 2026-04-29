// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package powerhaautomationservice_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM/dra-go-sdk/powerhaautomationservicev1"
)

func TestAccIBMPhaAPIKeyBasic(t *testing.T) {
	var conf powerhaautomationservicev1.APIKeyResponse
	instanceID := "8eefautr-4c02-0009-0086-8bd4d8cf61b6"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMPhaAPIKeyConfigBasic(instanceID),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMPhaAPIKeyExists("ibm_pha_api_key.pha_api_key_instance", conf),
					resource.TestCheckResourceAttr("ibm_pha_api_key.pha_api_key_instance", "instance_id", instanceID),
				),
			},
		},
	})
}

func TestAccIBMPhaAPIKeyAllArgs(t *testing.T) {
	var conf powerhaautomationservicev1.APIKeyResponse
	instanceID := "8eefautr-4c02-0009-0086-8bd4d8cf61b6"
	acceptLanguage := "en"
	ifNoneMatch := ""
	apiKey := ""

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMPhaAPIKeyConfig(instanceID, acceptLanguage, ifNoneMatch, apiKey),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMPhaAPIKeyExists("ibm_pha_api_key.pha_api_key_instance", conf),
					resource.TestCheckResourceAttr("ibm_pha_api_key.pha_api_key_instance", "instance_id", instanceID),
					resource.TestCheckResourceAttr("ibm_pha_api_key.pha_api_key_instance", "accept_language", acceptLanguage),
					resource.TestCheckResourceAttr("ibm_pha_api_key.pha_api_key_instance", "if_none_match", ifNoneMatch),
					resource.TestCheckResourceAttr("ibm_pha_api_key.pha_api_key_instance", "api_key", apiKey),
				),
			},
			// resource.TestStep{
			// 	ResourceName:      "ibm_pha_api_key.pha_api_key_instance",
			// 	ImportState:       true,
			// 	ImportStateVerify: true,
			// },
		},
	})
}

func testAccCheckIBMPhaAPIKeyConfigBasic(instanceID string) string {
	return fmt.Sprintf(`
		resource "ibm_pha_api_key" "pha_api_key_instance" {
			instance_id = "%s"
		}
	`, instanceID)
}

func testAccCheckIBMPhaAPIKeyConfig(instanceID string, acceptLanguage string, ifNoneMatch string, apiKey string) string {
	return fmt.Sprintf(`

		resource "ibm_pha_api_key" "pha_api_key_instance" {
			instance_id = "%s"
			accept_language = "%s"
			if_none_match = "%s"
			api_key = "%s"
		}
	`, instanceID, acceptLanguage, ifNoneMatch, apiKey)
}

func testAccCheckIBMPhaAPIKeyExists(n string, obj powerhaautomationservicev1.APIKeyResponse) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		powerhaAutomationServiceClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).PowerhaAutomationServiceV1()
		if err != nil {
			return err
		}

		getAPIKeyOptions := &powerhaautomationservicev1.GetAPIKeyOptions{}

		// parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		// if err != nil {
		// 	return err
		// }

		// getAPIKeyOptions.SetPhaInstanceID(parts[0])
		getAPIKeyOptions.SetPhaInstanceID(rs.Primary.ID)

		apiKeyResponse, _, err := powerhaAutomationServiceClient.GetAPIKey(getAPIKeyOptions)
		if err != nil {
			return err
		}

		obj = *apiKeyResponse
		return nil
	}
}

func testAccCheckIBMPhaAPIKeyDestroy(s *terraform.State) error {
	powerhaAutomationServiceClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).PowerhaAutomationServiceV1()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_pha_api_key" {
			continue
		}

		getAPIKeyOptions := &powerhaautomationservicev1.GetAPIKeyOptions{}

		// parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		// if err != nil {
		// 	return err
		// }

		// getAPIKeyOptions.SetPhaInstanceID(parts[0])
		getAPIKeyOptions.SetPhaInstanceID(rs.Primary.ID)

		// Try to find the key
		_, response, err := powerhaAutomationServiceClient.GetAPIKey(getAPIKeyOptions)

		if err == nil {
			return fmt.Errorf("pha_api_key still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for pha_api_key (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
