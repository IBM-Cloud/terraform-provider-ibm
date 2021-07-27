// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIbmIamApiKeyDataSourceBasic(t *testing.T) {
	apiKeyName := fmt.Sprintf("name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmIamApiKeyDataSourceConfigBasic(apiKeyName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_iam_api_key.iam_api_key", "apikey_id"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_api_key.iam_api_key", "entity_tag"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_api_key.iam_api_key", "crn"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_api_key.iam_api_key", "locked"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_api_key.iam_api_key", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_api_key.iam_api_key", "created_by"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_api_key.iam_api_key", "modified_at"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_api_key.iam_api_key", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_api_key.iam_api_key", "description"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_api_key.iam_api_key", "iam_id"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_api_key.iam_api_key", "account_id"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_api_key.iam_api_key", "apikey"),
				),
			},
		},
	})
}

func TestAccIbmIamApiKeyDataSourceAllArgs(t *testing.T) {
	apiKeyName := fmt.Sprintf("name_%d", acctest.RandIntRange(10, 100))
	apiKeyDescription := fmt.Sprintf("description_%d", acctest.RandIntRange(10, 100))
	apiKeyStoreValue := "false"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmIamApiKeyDataSourceConfig(apiKeyName, apiKeyDescription, apiKeyStoreValue),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_iam_api_key.iam_api_key", "apikey_id"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_api_key.iam_api_key", "entity_tag"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_api_key.iam_api_key", "crn"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_api_key.iam_api_key", "locked"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_api_key.iam_api_key", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_api_key.iam_api_key", "created_by"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_api_key.iam_api_key", "modified_at"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_api_key.iam_api_key", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_api_key.iam_api_key", "description"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_api_key.iam_api_key", "iam_id"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_api_key.iam_api_key", "account_id"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_api_key.iam_api_key", "apikey"),
				),
			},
		},
	})
}

func testAccCheckIbmIamApiKeyDataSourceConfigBasic(apiKeyName string) string {
	return fmt.Sprintf(`
		resource "ibm_iam_api_key" "iam_api_key" {
			name = "%s"
		}

		data "ibm_iam_api_key" "iam_api_key" {
			apikey_id = ibm_iam_api_key.iam_api_key.id
		}
	`, apiKeyName)
}

func testAccCheckIbmIamApiKeyDataSourceConfig(apiKeyName string, apiKeyDescription string, apiKeyStoreValue string) string {
	return fmt.Sprintf(`
		resource "ibm_iam_api_key" "iam_api_key" {
			name = "%s"
			description = "%s"
			store_value = %s
		}

		data "ibm_iam_api_key" "iam_api_key" {
			apikey_id = ibm_iam_api_key.iam_api_key.id
		}
	`, apiKeyName, apiKeyDescription, apiKeyStoreValue)
}
