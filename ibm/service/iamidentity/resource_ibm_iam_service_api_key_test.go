// Copyright IBM Corp. 2017, 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package iamidentity_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"

	"github.com/IBM/platform-services-go-sdk/iamidentityv1"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccIBMIAMServiceAPIKey_Basic(t *testing.T) {
	var apiKey string
	serviceName := fmt.Sprintf("terraform_iam_ser_%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("terraform_iam_%d", acctest.RandIntRange(10, 100))
	updateName := fmt.Sprintf("terraform_iam_%d", acctest.RandIntRange(10, 100))
	storeValue := true
	expiresAt := "2040-01-28T15:00+0000"
	expiresAtUpdate := "2035-01-28T15:00+0000"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIAMServiceAPIKeyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIAMServiceAPIKeyBasic(serviceName, name, storeValue, expiresAt),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIAMServiceAPIKeyExistsWithValidation("ibm_iam_service_api_key.testacc_apiKey", apiKey, storeValue),
					resource.TestCheckResourceAttr("ibm_iam_service_api_key.testacc_apiKey", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_service_api_key.testacc_apiKey", "expires_at", expiresAt),
					resource.TestCheckResourceAttrSet("ibm_iam_service_api_key.testacc_apiKey", "apikey"),
				),
			},
			{
				Config: testAccCheckIBMIAMServiceAPIKeyUpdateWithSameName(serviceName, name, expiresAtUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIAMServiceAPIKeyExistsWithValidation("ibm_iam_service_api_key.testacc_apiKey", apiKey, storeValue),
					resource.TestCheckResourceAttr("ibm_iam_service_api_key.testacc_apiKey", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_service_api_key.testacc_apiKey", "expires_at", expiresAtUpdate),
					resource.TestCheckResourceAttr("ibm_iam_service_api_key.testacc_apiKey", "description", "Service API Key for test scenario1"),
				),
			},
			{
				Config: testAccCheckIBMIAMServiceAPIKeyUpdate(serviceName, updateName, expiresAt),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_iam_service_api_key.testacc_apiKey", "name", updateName),
					resource.TestCheckResourceAttr("ibm_iam_service_api_key.testacc_apiKey", "expires_at", expiresAt),
					resource.TestCheckResourceAttr("ibm_iam_service_api_key.testacc_apiKey", "description", "Service API Key for test scenario2"),
				),
			},
		},
	})
}

func TestAccIBMIAMServiceAPIKey_doNotStoreApikeyValue(t *testing.T) {
	var apiKey string
	serviceName := fmt.Sprintf("terraform_iam_ser_%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("terraform_iam_%d", acctest.RandIntRange(10, 100))
	storeValue := false
	expiresAt := "2040-01-28T15:00+0000"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIAMServiceAPIKeyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIAMServiceAPIKeyBasic(serviceName, name, storeValue, expiresAt),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIAMServiceAPIKeyExistsWithValidation("ibm_iam_service_api_key.testacc_apiKey", apiKey, storeValue),
					resource.TestCheckResourceAttr("ibm_iam_service_api_key.testacc_apiKey", "name", name),
				),
			},
		},
	})
}

func TestAccIBMIAMServiceAPIKey_import(t *testing.T) {
	var apiKey string
	serviceName := fmt.Sprintf("terraform_iam_ser_%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("terraform_iam_%d", acctest.RandIntRange(10, 100))
	resourceName := "ibm_iam_service_api_key.testacc_apiKey"
	storeValue := true

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIAMServiceAPIKeyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIAMServiceAPIKeyImport(serviceName, name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIAMServiceAPIKeyExistsWithValidation(resourceName, apiKey, storeValue),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "Service API Key for test scenario2"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"store_value",
				},
			},
		},
	})
}

func testAccCheckIBMIAMServiceAPIKeyDestroy(s *terraform.State) error {
	rsContClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).IAMIdentityV1API()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_iam_service_api_key" {
			continue
		}

		getAPIKeyOptions := &iamidentityv1.GetAPIKeyOptions{
			ID: &rs.Primary.ID,
		}

		_, _, err := rsContClient.GetAPIKey(getAPIKeyOptions)
		if err == nil {
			return fmt.Errorf("Service API Key Still Exists: %s", rs.Primary.ID)
		}
	}

	return nil
}

func testAccCheckIBMIAMServiceAPIKeyExistsWithValidation(n string, apiKey string, apikeyValueExpected bool) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		rsContClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).IAMIdentityV1API()
		if err != nil {
			return err
		}

		getAPIKeyOptions := &iamidentityv1.GetAPIKeyOptions{
			ID: &rs.Primary.ID,
		}

		foundAPIKey, _, err := rsContClient.GetAPIKey(getAPIKeyOptions)
		if err != nil {
			return err
		}

		if apikeyValueExpected {
			if foundAPIKey.Apikey == nil {
				return fmt.Errorf("apikey value should be present")
			}
		} else {
			if foundAPIKey.Apikey != nil {
				return fmt.Errorf("apikey value should not be present")
			}
		}

		apiKey = *foundAPIKey.ID
		return nil
	}
}

func testAccCheckIBMIAMServiceAPIKeyBasic(serviceName, name string, storeValue bool, expiresAt string) string {
	return fmt.Sprintf(`
		
		resource "ibm_iam_service_id" "serviceID" {
			name = "%s"
			tags = ["tag1", "tag2"]
		  }
		  resource "ibm_iam_service_api_key" "testacc_apiKey" {
			name = "%s"
			iam_service_id = ibm_iam_service_id.serviceID.iam_id
			store_value = "%t"
			expires_at = "%s"
	  	}
	`, serviceName, name, storeValue, expiresAt)
}

func testAccCheckIBMIAMServiceAPIKeyUpdateWithSameName(serviceName, name string, expiresAt string) string {
	return fmt.Sprintf(`
		
		resource "ibm_iam_service_id" "serviceID" {
			name        = "%s"
			tags        = ["tag1", "tag2", "db"]
		  }
		  resource "ibm_iam_service_api_key" "testacc_apiKey" {
			name = "%s"
			description = "Service API Key for test scenario1"
			iam_service_id = ibm_iam_service_id.serviceID.iam_id
			expires_at = "%s"
	  	}
	`, serviceName, name, expiresAt)
}

func testAccCheckIBMIAMServiceAPIKeyUpdate(serviceName, updateName string, expiresAt string) string {
	return fmt.Sprintf(`

		resource "ibm_iam_service_id" "serviceID" {
			name              = "%s"		
			description       = "ServiceID for test scenario2"
			tags              = ["tag1"]
		}
		resource "ibm_iam_service_api_key" "testacc_apiKey" {
			name = "%s"
			description = "Service API Key for test scenario2"
			iam_service_id = ibm_iam_service_id.serviceID.iam_id
			expires_at = "%s"
	  	}
	`, serviceName, updateName, expiresAt)
}

func testAccCheckIBMIAMServiceAPIKeyImport(serviceName, name string) string {
	return fmt.Sprintf(`

		resource "ibm_iam_service_id" "serviceID" {
			name              = "%s"		
			description       = "ServiceID for test scenario2"
		}
		resource "ibm_iam_service_api_key" "testacc_apiKey" {
			name = "%s"
			description = "Service API Key for test scenario2"
			iam_service_id = ibm_iam_service_id.serviceID.iam_id
	  	}
	`, serviceName, name)
}
