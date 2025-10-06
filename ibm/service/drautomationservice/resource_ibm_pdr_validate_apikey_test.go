// Copyright IBM Corp. 2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package drautomationservice_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.ibm.com/DRAutomation/dra-go-sdk/drautomationservicev1"
)

func TestAccIbmPdrValidateApikeyBasic(t *testing.T) {
	var conf drautomationservicev1.ValidationKeyResponse
	instanceID := "crn:v1:staging:public:power-dr-automation:global:a/a09202c1bfb04ceebfb4a9fd38c87721:050ebe3b-13f4-4db8-8ece-501a3c13be0965822::"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmPdrValidateApikeyDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmPdrValidateApikeyConfigBasic(instanceID),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmPdrValidateApikeyExists("ibm_pdr_validate_apikey.pdr_validate_apikey_instance", conf),
					resource.TestCheckResourceAttr("ibm_pdr_validate_apikey.pdr_validate_apikey_instance", "instance_id", instanceID),
				),
			},
		},
	})
}

func TestAccIbmPdrValidateApikeyAllArgs(t *testing.T) {
	var conf drautomationservicev1.ValidationKeyResponse
	instanceID := "crn:v1:staging:public:power-dr-automation:global:a/a09202c1bfb04ceebfb4a9fd38c87721:050ebe3b-13f4-4db8-8ece-501a3c13be0965822::"
	acceptLanguage := fmt.Sprintf("tf_accept_language_%d", acctest.RandIntRange(10, 100))
	ifNoneMatch := fmt.Sprintf("tf_if_none_match_%d", acctest.RandIntRange(10, 100))
	acceptLanguageUpdate := fmt.Sprintf("tf_accept_language_%d", acctest.RandIntRange(10, 100))
	ifNoneMatchUpdate := fmt.Sprintf("tf_if_none_match_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmPdrValidateApikeyDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmPdrValidateApikeyConfig(instanceID, acceptLanguage, ifNoneMatch),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmPdrValidateApikeyExists("ibm_pdr_validate_apikey.pdr_validate_apikey_instance", conf),
					resource.TestCheckResourceAttr("ibm_pdr_validate_apikey.pdr_validate_apikey_instance", "instance_id", instanceID),
					resource.TestCheckResourceAttr("ibm_pdr_validate_apikey.pdr_validate_apikey_instance", "accept_language", acceptLanguage),
					resource.TestCheckResourceAttr("ibm_pdr_validate_apikey.pdr_validate_apikey_instance", "if_none_match", ifNoneMatch),
				),
			},
			resource.TestStep{
				Config: testAccCheckIbmPdrValidateApikeyConfig(instanceID, acceptLanguageUpdate, ifNoneMatchUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_pdr_validate_apikey.pdr_validate_apikey_instance", "instance_id", instanceID),
					resource.TestCheckResourceAttr("ibm_pdr_validate_apikey.pdr_validate_apikey_instance", "accept_language", acceptLanguageUpdate),
					resource.TestCheckResourceAttr("ibm_pdr_validate_apikey.pdr_validate_apikey_instance", "if_none_match", ifNoneMatchUpdate),
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

func testAccCheckIbmPdrValidateApikeyConfigBasic(instanceID string) string {
	return fmt.Sprintf(`
		resource "ibm_pdr_validate_apikey" "pdr_validate_apikey_instance" {
			instance_id = "%s"
			api_key = "azGTysdgsvameQEAhya_1_fD"
		}
	`, instanceID)
}

func testAccCheckIbmPdrValidateApikeyConfig(instanceID string, acceptLanguage string, ifNoneMatch string) string {
	return fmt.Sprintf(`

		resource "ibm_pdr_validate_apikey" "pdr_validate_apikey_instance" {
			instance_id = "%s"
			api_key = "azGsdsdfTyameQEAhya_1_fD"
			accept_language = "%s"
			if_none_match = "%s"
		}
	`, instanceID, acceptLanguage, ifNoneMatch)
}

func testAccCheckIbmPdrValidateApikeyExists(n string, obj drautomationservicev1.ValidationKeyResponse) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		drAutomationServiceClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).DrAutomationServiceV1()
		if err != nil {
			return err
		}

		getServiceInstanceKeyV1Options := &drautomationservicev1.GetServiceInstanceKeyV1Options{}

		// parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		// if err != nil {
		// 	return err
		// }

		// getServiceInstanceKeyV1Options.SetInstanceID(parts[0])
		getServiceInstanceKeyV1Options.SetInstanceID(rs.Primary.ID)

		validationKeyResponse, _, err := drAutomationServiceClient.GetServiceInstanceKeyV1(getServiceInstanceKeyV1Options)
		if err != nil {
			return err
		}

		obj = *validationKeyResponse
		return nil
	}
}

func testAccCheckIbmPdrValidateApikeyDestroy(s *terraform.State) error {
	drAutomationServiceClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).DrAutomationServiceV1()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_pdr_validate_apikey" {
			continue
		}

		getServiceInstanceKeyV1Options := &drautomationservicev1.GetServiceInstanceKeyV1Options{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}
		getServiceInstanceKeyV1Options.SetInstanceID(fmt.Sprintf("%s/%s", parts[0], parts[1]))
		// Try to find the key
		_, response, err := drAutomationServiceClient.GetServiceInstanceKeyV1(getServiceInstanceKeyV1Options)
		fmt.Println(response)
		if err == nil {
			return fmt.Errorf("pdr_validate_apikey still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for pdr_validate_apikey (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
