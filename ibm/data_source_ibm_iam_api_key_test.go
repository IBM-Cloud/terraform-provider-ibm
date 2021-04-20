/**
 * (C) Copyright IBM Corp. 2021.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIbmIamApiKeyDataSourceBasic(t *testing.T) {
	apiKeyName := fmt.Sprintf("name_%d", acctest.RandIntRange(10, 100))
	apiKeyIamID := fmt.Sprintf("iam_id_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmIamApiKeyDataSourceConfigBasic(apiKeyName, apiKeyIamID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_iam_api_key.iam_api_key", "id"),
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
	apiKeyIamID := fmt.Sprintf("iam_id_%d", acctest.RandIntRange(10, 100))
	apiKeyDescription := fmt.Sprintf("description_%d", acctest.RandIntRange(10, 100))
	apiKeyAccountID := fmt.Sprintf("account_id_%d", acctest.RandIntRange(10, 100))
	apiKeyApikey := fmt.Sprintf("apikey_%d", acctest.RandIntRange(10, 100))
	apiKeyStoreValue := "false"
	apiKeyEntityLock := fmt.Sprintf("locked_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmIamApiKeyDataSourceConfig(apiKeyName, apiKeyIamID, apiKeyDescription, apiKeyAccountID, apiKeyApikey, apiKeyStoreValue, apiKeyEntityLock),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_iam_api_key.iam_api_key", "id"),
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

func testAccCheckIbmIamApiKeyDataSourceConfigBasic(apiKeyName string, apiKeyIamID string) string {
	return fmt.Sprintf(`
		resource "ibm_iam_api_key" "iam_api_key" {
			name = "%s"
			iam_id = "%s"
		}

		data "ibm_iam_api_key" "iam_api_key" {
			id = ibm_iam_api_key.iam_api_key.id
		}
	`, apiKeyName, apiKeyIamID)
}

func testAccCheckIbmIamApiKeyDataSourceConfig(apiKeyName string, apiKeyIamID string, apiKeyDescription string, apiKeyAccountID string, apiKeyApikey string, apiKeyStoreValue string, apiKeyEntityLock string) string {
	return fmt.Sprintf(`
		resource "ibm_iam_api_key" "iam_api_key" {
			name = "%s"
			iam_id = "%s"
			description = "%s"
			account_id = "%s"
			apikey = "%s"
			store_value = %s
			locked = "%s"
		}

		data "ibm_iam_api_key" "iam_api_key" {
			id = ibm_iam_api_key.iam_api_key.id
		}
	`, apiKeyName, apiKeyIamID, apiKeyDescription, apiKeyAccountID, apiKeyApikey, apiKeyStoreValue, apiKeyEntityLock)
}
