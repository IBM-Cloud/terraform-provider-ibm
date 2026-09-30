// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.113.1-d76630af-20260320-135953
 */

package iamidentity_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMIamAPIKeyDataSourceBasic(t *testing.T) {
	apiKeyName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIamAPIKeyDataSourceConfigBasic(apiKeyName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_iam_api_key.iam_api_key_instance_data", "apikey_id"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_api_key.iam_api_key_instance_data", "entity_tag"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_api_key.iam_api_key_instance_data", "crn"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_api_key.iam_api_key_instance_data", "locked"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_api_key.iam_api_key_instance_data", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_api_key.iam_api_key_instance_data", "created_by"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_api_key.iam_api_key_instance_data", "modified_at"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_api_key.iam_api_key_instance_data", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_api_key.iam_api_key_instance_data", "iam_id"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_api_key.iam_api_key_instance_data", "account_id"),
				),
			},
		},
	})
}

func TestAccIBMIamAPIKeyDataSourceAllArgs(t *testing.T) {
	apiKeyName := fmt.Sprintf("name_%d", acctest.RandIntRange(10, 100))
	apiKeyDescription := fmt.Sprintf("description_%d", acctest.RandIntRange(10, 100))
	apiKeyStoreValue := "false"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIamAPIKeyDataSourceConfig(apiKeyName, apiKeyDescription, apiKeyStoreValue),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_iam_api_key.iam_api_key_instance_data", "apikey_id"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_api_key.iam_api_key_instance_data", "entity_tag"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_api_key.iam_api_key_instance_data", "crn"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_api_key.iam_api_key_instance_data", "locked"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_api_key.iam_api_key_instance_data", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_api_key.iam_api_key_instance_data", "created_by"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_api_key.iam_api_key_instance_data", "modified_at"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_api_key.iam_api_key_instance_data", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_api_key.iam_api_key_instance_data", "iam_id"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_api_key.iam_api_key_instance_data", "account_id"),
				),
			},
		},
	})
}

func testAccCheckIBMIamAPIKeyDataSourceConfigBasic(apiKeyName string) string {
	return fmt.Sprintf(`
		resource "ibm_iam_api_key" "iam_api_key_instance" {
			name = "%s"
		}

		data "ibm_iam_api_key" "iam_api_key_instance_data" {
			apikey_id = ibm_iam_api_key.iam_api_key_instance.id
		}
	`, apiKeyName)
}

func testAccCheckIBMIamAPIKeyDataSourceConfig(apiKeyName string, apiKeyDescription string, apiKeyStoreValue string) string {
	return fmt.Sprintf(`
		resource "ibm_iam_api_key" "iam_api_key_instance" {
			name = "%s"
			description = "%s"
			store_value = %s
		}

		data "ibm_iam_api_key" "iam_api_key_instance_data" {
			apikey_id = ibm_iam_api_key.iam_api_key_instance.id
		}
	`, apiKeyName, apiKeyDescription, apiKeyStoreValue)
}
