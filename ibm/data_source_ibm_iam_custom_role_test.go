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

func TestAccIbmIamCustomRoleDataSourceBasic(t *testing.T) {
	customRoleDisplayName := fmt.Sprintf("display_name_%d", acctest.RandIntRange(10, 100))
	customRoleName := fmt.Sprintf("name_%d", acctest.RandIntRange(10, 100))
	customRoleAccountID := fmt.Sprintf("account_id_%d", acctest.RandIntRange(10, 100))
	customRoleServiceName := fmt.Sprintf("service_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmIamCustomRoleDataSourceConfigBasic(customRoleDisplayName, customRoleName, customRoleAccountID, customRoleServiceName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_iam_custom_role.iam_custom_role", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_custom_role.iam_custom_role", "role_id"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_custom_role.iam_custom_role", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_custom_role.iam_custom_role", "display_name"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_custom_role.iam_custom_role", "description"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_custom_role.iam_custom_role", "actions.#"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_custom_role.iam_custom_role", "crn"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_custom_role.iam_custom_role", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_custom_role.iam_custom_role", "account_id"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_custom_role.iam_custom_role", "service_name"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_custom_role.iam_custom_role", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_custom_role.iam_custom_role", "created_by_id"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_custom_role.iam_custom_role", "last_modified_at"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_custom_role.iam_custom_role", "last_modified_by_id"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_custom_role.iam_custom_role", "href"),
				),
			},
		},
	})
}

func TestAccIbmIamCustomRoleDataSourceAllArgs(t *testing.T) {
	customRoleDisplayName := fmt.Sprintf("display_name_%d", acctest.RandIntRange(10, 100))
	customRoleName := fmt.Sprintf("name_%d", acctest.RandIntRange(10, 100))
	customRoleAccountID := fmt.Sprintf("account_id_%d", acctest.RandIntRange(10, 100))
	customRoleServiceName := fmt.Sprintf("service_name_%d", acctest.RandIntRange(10, 100))
	customRoleDescription := fmt.Sprintf("description_%d", acctest.RandIntRange(10, 100))
	customRoleAcceptLanguage := fmt.Sprintf("accept_language_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmIamCustomRoleDataSourceConfig(customRoleDisplayName, customRoleName, customRoleAccountID, customRoleServiceName, customRoleDescription, customRoleAcceptLanguage),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_iam_custom_role.iam_custom_role", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_custom_role.iam_custom_role", "role_id"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_custom_role.iam_custom_role", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_custom_role.iam_custom_role", "display_name"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_custom_role.iam_custom_role", "description"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_custom_role.iam_custom_role", "actions.#"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_custom_role.iam_custom_role", "crn"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_custom_role.iam_custom_role", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_custom_role.iam_custom_role", "account_id"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_custom_role.iam_custom_role", "service_name"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_custom_role.iam_custom_role", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_custom_role.iam_custom_role", "created_by_id"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_custom_role.iam_custom_role", "last_modified_at"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_custom_role.iam_custom_role", "last_modified_by_id"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_custom_role.iam_custom_role", "href"),
				),
			},
		},
	})
}

func testAccCheckIbmIamCustomRoleDataSourceConfigBasic(customRoleDisplayName string, customRoleName string, customRoleAccountID string, customRoleServiceName string) string {
	return fmt.Sprintf(`
		resource "ibm_iam_custom_role" "iam_custom_role" {
			display_name = "%s"
			actions = "FIXME"
			name = "%s"
			account_id = "%s"
			service_name = "%s"
		}

		data "ibm_iam_custom_role" "iam_custom_role" {
			role_id = "role_id"
		}
	`, customRoleDisplayName, customRoleName, customRoleAccountID, customRoleServiceName)
}

func testAccCheckIbmIamCustomRoleDataSourceConfig(customRoleDisplayName string, customRoleName string, customRoleAccountID string, customRoleServiceName string, customRoleDescription string, customRoleAcceptLanguage string) string {
	return fmt.Sprintf(`
		resource "ibm_iam_custom_role" "iam_custom_role" {
			display_name = "%s"
			actions = "FIXME"
			name = "%s"
			account_id = "%s"
			service_name = "%s"
			description = "%s"
			Accept-Language = "%s"
		}

		data "ibm_iam_custom_role" "iam_custom_role" {
			role_id = "role_id"
		}
	`, customRoleDisplayName, customRoleName, customRoleAccountID, customRoleServiceName, customRoleDescription, customRoleAcceptLanguage)
}
