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
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/IBM/platform-services-go-sdk/iamidentityv1"
)

func TestAccIbmIamApiKeyBasic(t *testing.T) {
	var conf iamidentityv1.APIKey
	name := fmt.Sprintf("name_%d", acctest.RandIntRange(10, 100))
	iamID := fmt.Sprintf("iam_id_%d", acctest.RandIntRange(10, 100))
	nameUpdate := fmt.Sprintf("name_%d", acctest.RandIntRange(10, 100))
	iamIDUpdate := fmt.Sprintf("iam_id_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIbmIamApiKeyDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmIamApiKeyConfigBasic(name, iamID),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmIamApiKeyExists("ibm_iam_api_key.iam_api_key", conf),
					resource.TestCheckResourceAttr("ibm_iam_api_key.iam_api_key", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_api_key.iam_api_key", "iam_id", iamID),
				),
			},
			resource.TestStep{
				Config: testAccCheckIbmIamApiKeyConfigBasic(nameUpdate, iamIDUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_iam_api_key.iam_api_key", "name", nameUpdate),
					resource.TestCheckResourceAttr("ibm_iam_api_key.iam_api_key", "iam_id", iamIDUpdate),
				),
			},
		},
	})
}

func TestAccIbmIamApiKeyAllArgs(t *testing.T) {
	var conf iamidentityv1.APIKey
	name := fmt.Sprintf("name_%d", acctest.RandIntRange(10, 100))
	iamID := fmt.Sprintf("iam_id_%d", acctest.RandIntRange(10, 100))
	description := fmt.Sprintf("description_%d", acctest.RandIntRange(10, 100))
	accountID := fmt.Sprintf("account_id_%d", acctest.RandIntRange(10, 100))
	apikey := fmt.Sprintf("apikey_%d", acctest.RandIntRange(10, 100))
	storeValue := "false"
	entityLock := fmt.Sprintf("lock_%d", acctest.RandIntRange(10, 100))
	nameUpdate := fmt.Sprintf("name_%d", acctest.RandIntRange(10, 100))
	iamIDUpdate := fmt.Sprintf("iam_id_%d", acctest.RandIntRange(10, 100))
	descriptionUpdate := fmt.Sprintf("description_%d", acctest.RandIntRange(10, 100))
	accountIDUpdate := fmt.Sprintf("account_id_%d", acctest.RandIntRange(10, 100))
	apikeyUpdate := fmt.Sprintf("apikey_%d", acctest.RandIntRange(10, 100))
	storeValueUpdate := "true"
	entityLockUpdate := fmt.Sprintf("locked_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIbmIamApiKeyDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmIamApiKeyConfig(name, iamID, description, accountID, apikey, storeValue, entityLock),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmIamApiKeyExists("ibm_iam_api_key.iam_api_key", conf),
					resource.TestCheckResourceAttr("ibm_iam_api_key.iam_api_key", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_api_key.iam_api_key", "iam_id", iamID),
					resource.TestCheckResourceAttr("ibm_iam_api_key.iam_api_key", "description", description),
					resource.TestCheckResourceAttr("ibm_iam_api_key.iam_api_key", "account_id", accountID),
					resource.TestCheckResourceAttr("ibm_iam_api_key.iam_api_key", "apikey", apikey),
					resource.TestCheckResourceAttr("ibm_iam_api_key.iam_api_key", "store_value", storeValue),
					resource.TestCheckResourceAttr("ibm_iam_api_key.iam_api_key", "locked", entityLock),
				),
			},
			resource.TestStep{
				Config: testAccCheckIbmIamApiKeyConfig(nameUpdate, iamIDUpdate, descriptionUpdate, accountIDUpdate, apikeyUpdate, storeValueUpdate, entityLockUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_iam_api_key.iam_api_key", "name", nameUpdate),
					resource.TestCheckResourceAttr("ibm_iam_api_key.iam_api_key", "iam_id", iamIDUpdate),
					resource.TestCheckResourceAttr("ibm_iam_api_key.iam_api_key", "description", descriptionUpdate),
					resource.TestCheckResourceAttr("ibm_iam_api_key.iam_api_key", "account_id", accountIDUpdate),
					resource.TestCheckResourceAttr("ibm_iam_api_key.iam_api_key", "apikey", apikeyUpdate),
					resource.TestCheckResourceAttr("ibm_iam_api_key.iam_api_key", "store_value", storeValueUpdate),
					resource.TestCheckResourceAttr("ibm_iam_api_key.iam_api_key", "locked", entityLockUpdate),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_iam_api_key.iam_api_key",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIbmIamApiKeyConfigBasic(name string, iamID string) string {
	return fmt.Sprintf(`

		resource "ibm_iam_api_key" "iam_api_key" {
			name = "%s"
			iam_id = "%s"
		}
	`, name, iamID)
}

func testAccCheckIbmIamApiKeyConfig(name string, iamID string, description string, accountID string, apikey string, storeValue string, entityLock string) string {
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
	`, name, iamID, description, accountID, apikey, storeValue, entityLock)
}

func testAccCheckIbmIamApiKeyExists(n string, obj iamidentityv1.APIKey) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		iamIdentityClient, err := testAccProvider.Meta().(ClientSession).IAMIdentityV1API()
		if err != nil {
			return err
		}

		getApiKeyOptions := &iamidentityv1.GetAPIKeyOptions{}

		getApiKeyOptions.SetID(rs.Primary.ID)

		apiKey, _, err := iamIdentityClient.GetAPIKey(getApiKeyOptions)
		if err != nil {
			return err
		}

		obj = *apiKey
		return nil
	}
}

func testAccCheckIbmIamApiKeyDestroy(s *terraform.State) error {
	iamIdentityClient, err := testAccProvider.Meta().(ClientSession).IAMIdentityV1API()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_iam_api_key" {
			continue
		}

		getApiKeyOptions := &iamidentityv1.GetAPIKeyOptions{}

		getApiKeyOptions.SetID(rs.Primary.ID)

		// Try to find the key
		_, response, err := iamIdentityClient.GetAPIKey(getApiKeyOptions)

		if err == nil {
			return fmt.Errorf("iam_api_key still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for iam_api_key (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
