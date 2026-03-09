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
	"github.ibm.com/DRAutomation/dra-go-sdk/powerhaautomationservicev1"
)

func TestAccIBMPhaAPIKeyBasic(t *testing.T) {
	var conf powerhaautomationservicev1.APIKeyResponse
	phaInstanceID := "8ce2a099-a463-479a-9a1d-eedc19287a62"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMPhaAPIKeyConfigBasic(phaInstanceID),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMPhaAPIKeyExists("ibm_pha_api_key.pha_api_key_instance", conf),
					resource.TestCheckResourceAttr("ibm_pha_api_key.pha_api_key_instance", "pha_instance_id", phaInstanceID),
				),
			},
		},
	})
}

func TestAccIBMPhaAPIKeyAllArgs(t *testing.T) {
	var conf powerhaautomationservicev1.APIKeyResponse
	phaInstanceID := "8ce2a099-a463-479a-9a1d-eedc19287a62"
	acceptLanguage := "en"
	apiKey := ""

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMPhaAPIKeyConfig(phaInstanceID, acceptLanguage, apiKey),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMPhaAPIKeyExists("ibm_pha_api_key.pha_api_key_instance", conf),
					resource.TestCheckResourceAttr("ibm_pha_api_key.pha_api_key_instance", "pha_instance_id", phaInstanceID),
					resource.TestCheckResourceAttr("ibm_pha_api_key.pha_api_key_instance", "accept_language", acceptLanguage),
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

func testAccCheckIBMPhaAPIKeyConfigBasic(phaInstanceID string) string {
	return fmt.Sprintf(`
		resource "ibm_pha_api_key" "pha_api_key_instance" {
			pha_instance_id = "%s"
			api_key = ""
		}
	`, phaInstanceID)
}

func testAccCheckIBMPhaAPIKeyConfig(phaInstanceID string, acceptLanguage string, apiKey string) string {
	return fmt.Sprintf(`

		resource "ibm_pha_api_key" "pha_api_key_instance" {
			pha_instance_id = "%s"
			accept_language = "%s"
			api_key = "%s"
		}
	`, phaInstanceID, acceptLanguage, apiKey)
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
