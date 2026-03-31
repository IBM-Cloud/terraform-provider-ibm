// Copyright IBM Corp. 2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package drautomationservice_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM/dra-go-sdk/drautomationservicev1"
)

func TestAccIBMPdrValidateApikeyBasic(t *testing.T) {
	var conf drautomationservicev1.ValidationKeyResponse
	instanceID := "3ad42074-e4f3-4b7b-a3d8-f88799e6da09"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMPdrValidateApikeyConfigBasic(instanceID),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMPdrValidateApikeyExists("ibm_pdr_validate_apikey.pdr_validate_apikey_instance", conf),
					resource.TestCheckResourceAttr("ibm_pdr_validate_apikey.pdr_validate_apikey_instance", "instance_id", instanceID),
				),
			},
		},
	})
}

func TestAccIBMPdrValidateApikeyAllArgs(t *testing.T) {
	var conf drautomationservicev1.ValidationKeyResponse
	instanceID := "3ad42074-e4f3-4b7b-a3d8-f88799e6da09"
	acceptLanguage := "it"
	acceptLanguageUpdate := "it"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMPdrValidateApikeyConfig(instanceID, acceptLanguage),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMPdrValidateApikeyExists("ibm_pdr_validate_apikey.pdr_validate_apikey_instance", conf),
					resource.TestCheckResourceAttr("ibm_pdr_validate_apikey.pdr_validate_apikey_instance", "instance_id", instanceID),
					resource.TestCheckResourceAttr("ibm_pdr_validate_apikey.pdr_validate_apikey_instance", "accept_language", acceptLanguage),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMPdrValidateApikeyConfig(instanceID, acceptLanguageUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_pdr_validate_apikey.pdr_validate_apikey_instance", "instance_id", instanceID),
					resource.TestCheckResourceAttr("ibm_pdr_validate_apikey.pdr_validate_apikey_instance", "accept_language", acceptLanguageUpdate),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_pdr_validate_apikey.pdr_validate_apikey_instance",
				ImportState:       true,
				ImportStateVerify: false,
			},
		},
	})
}

func testAccCheckIBMPdrValidateApikeyConfigBasic(instanceID string) string {
	apiKey := acc.DRApiKey
	return fmt.Sprintf(`
		resource "ibm_pdr_validate_apikey" "pdr_validate_apikey_instance" {
			instance_id = "%s"
			api_key = "%s"
		}
	`, instanceID, apiKey)
}

func testAccCheckIBMPdrValidateApikeyConfig(instanceID string, acceptLanguage string) string {
	apiKey := acc.DRApiKey
	return fmt.Sprintf(`
		resource "ibm_pdr_validate_apikey" "pdr_validate_apikey_instance" {
			instance_id = "%s"
			api_key = "%s"
			accept_language = "%s"
		}
	`, instanceID, apiKey, acceptLanguage)
}

func testAccCheckIBMPdrValidateApikeyExists(n string, obj drautomationservicev1.ValidationKeyResponse) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		drAutomationServiceClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).DrAutomationServiceV1()
		if err != nil {
			return err
		}

		getApikeyOptions := &drautomationservicev1.GetApikeyOptions{}

		getApikeyOptions.SetInstanceID(rs.Primary.ID)

		validationKeyResponse, _, err := drAutomationServiceClient.GetApikey(getApikeyOptions)
		if err != nil {
			return err
		}

		obj = *validationKeyResponse
		return nil
	}
}

func testAccCheckIBMPdrValidateApikeyDestroy(s *terraform.State) error {
	drAutomationServiceClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).DrAutomationServiceV1()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_pdr_validate_apikey" {
			continue
		}

		getApikeyOptions := &drautomationservicev1.GetApikeyOptions{}

		getApikeyOptions.SetInstanceID(rs.Primary.ID)

		// Try to find the key
		_, response, err := drAutomationServiceClient.GetApikey(getApikeyOptions)

		if err == nil {
			return fmt.Errorf("pdr_validate_apikey still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for pdr_validate_apikey (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
