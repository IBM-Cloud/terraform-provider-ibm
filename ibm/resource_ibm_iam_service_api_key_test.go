// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"testing"

	"github.com/IBM/platform-services-go-sdk/iamidentityv1"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccIBMIAMServiceAPIKey_Basic(t *testing.T) {
	var apiKey string
	serviceName := fmt.Sprintf("terraform_iam_ser_%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("terraform_iam_%d", acctest.RandIntRange(10, 100))
	updateName := fmt.Sprintf("terraform_iam_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMIAMServiceAPIKeyDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIAMServiceAPIKeyBasic(serviceName, name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIAMServiceAPIKeyExists("ibm_iam_service_api_key.testacc_apiKey", apiKey),
					resource.TestCheckResourceAttr("ibm_iam_service_api_key.testacc_apiKey", "name", name),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMIAMServiceAPIKeyUpdateWithSameName(serviceName, name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIAMServiceAPIKeyExists("ibm_iam_service_api_key.testacc_apiKey", apiKey),
					resource.TestCheckResourceAttr("ibm_iam_service_api_key.testacc_apiKey", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_service_api_key.testacc_apiKey", "description", "Service API Key for test scenario1"),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMIAMServiceAPIKeyUpdate(serviceName, updateName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_iam_service_api_key.testacc_apiKey", "name", updateName),
					resource.TestCheckResourceAttr("ibm_iam_service_api_key.testacc_apiKey", "description", "Service API Key for test scenario2"),
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

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMIAMServiceAPIKeyDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIAMServiceAPIKeyImport(serviceName, name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIAMServiceAPIKeyExists(resourceName, apiKey),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "Service API Key for test scenario2"),
				),
			},
			resource.TestStep{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIBMIAMServiceAPIKeyDestroy(s *terraform.State) error {
	rsContClient, err := testAccProvider.Meta().(ClientSession).IAMIdentityV1API()
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

func testAccCheckIBMIAMServiceAPIKeyExists(n string, apiKey string) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		rsContClient, err := testAccProvider.Meta().(ClientSession).IAMIdentityV1API()
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

		apiKey = *foundAPIKey.ID
		return nil
	}
}

func testAccCheckIBMIAMServiceAPIKeyBasic(serviceName, name string) string {
	return fmt.Sprintf(`
		
		resource "ibm_iam_service_id" "serviceID" {
			name = "%s"
			tags = ["tag1", "tag2"]
		  }
		  resource "ibm_iam_service_api_key" "testacc_apiKey" {
			name = "%s"
			iam_service_id = ibm_iam_service_id.serviceID.iam_id
	  	}
	`, serviceName, name)
}

func testAccCheckIBMIAMServiceAPIKeyUpdateWithSameName(serviceName, name string) string {
	return fmt.Sprintf(`
		
		resource "ibm_iam_service_id" "serviceID" {
			name        = "%s"
			tags        = ["tag1", "tag2", "db"]
		  }
		  resource "ibm_iam_service_api_key" "testacc_apiKey" {
			name = "%s"
			description = "Service API Key for test scenario1"
			iam_service_id = ibm_iam_service_id.serviceID.iam_id
	  	}
	`, serviceName, name)
}

func testAccCheckIBMIAMServiceAPIKeyUpdate(serviceName, updateName string) string {
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
	  	}
	`, serviceName, updateName)
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
