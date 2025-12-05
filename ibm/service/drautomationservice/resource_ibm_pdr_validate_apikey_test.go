// Copyright IBM Corp. 2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package drautomationservice_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.ibm.com/DRAutomation/dra-go-sdk/drautomationservicev1"
)

func TestAccIBMPdrValidateApikeyBasic(t *testing.T) {
	var conf drautomationservicev1.ValidationKeyResponse
	instanceID := "123456d3-1122-3344-b67d-4389b44b7bf9"

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
	instanceID := "123456d3-1122-3344-b67d-4389b44b7bf9"
	acceptLanguage := fmt.Sprintf("tf_accept_language_%d", acctest.RandIntRange(10, 100))
	acceptLanguageUpdate := fmt.Sprintf("tf_accept_language_%d", acctest.RandIntRange(10, 100))

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
				ResourceName:      "ibm_pdr_validate_apikey.pdr_validate_apikey",
				ImportState:       true,
				ImportStateVerify: true,
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
			accept_language = "%s"
			api_key = "%s"
		}
	`, instanceID, acceptLanguage, apiKey)
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

		// parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		// if err != nil {
		// 	return err
		// }

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

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getApikeyOptions.SetInstanceID(parts[0])

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
