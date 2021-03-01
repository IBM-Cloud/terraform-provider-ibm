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

func TestAccIbmIamPolicyDataSourceBasic(t *testing.T) {
	policyType := fmt.Sprintf("type_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmIamPolicyDataSourceConfigBasic(policyType),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_iam_policy.iam_policy", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_policy.iam_policy", "policy_id"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_policy.iam_policy", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_policy.iam_policy", "type"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_policy.iam_policy", "description"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_policy.iam_policy", "subjects.#"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_policy.iam_policy", "roles.#"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_policy.iam_policy", "resources.#"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_policy.iam_policy", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_policy.iam_policy", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_policy.iam_policy", "created_by_id"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_policy.iam_policy", "last_modified_at"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_policy.iam_policy", "last_modified_by_id"),
				),
			},
		},
	})
}

func TestAccIbmIamPolicyDataSourceAllArgs(t *testing.T) {
	policyType := fmt.Sprintf("type_%d", acctest.RandIntRange(10, 100))
	policyDescription := fmt.Sprintf("description_%d", acctest.RandIntRange(10, 100))
	policyAcceptLanguage := fmt.Sprintf("accept_language_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmIamPolicyDataSourceConfig(policyType, policyDescription, policyAcceptLanguage),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_iam_policy.iam_policy", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_policy.iam_policy", "policy_id"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_policy.iam_policy", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_policy.iam_policy", "type"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_policy.iam_policy", "description"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_policy.iam_policy", "subjects.#"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_policy.iam_policy", "roles.#"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_policy.iam_policy", "roles.0.role_id"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_policy.iam_policy", "roles.0.display_name"),
					resource.TestCheckResourceAttr("data.ibm_iam_policy.iam_policy", "roles.0.description", policyDescription),
					resource.TestCheckResourceAttrSet("data.ibm_iam_policy.iam_policy", "resources.#"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_policy.iam_policy", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_policy.iam_policy", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_policy.iam_policy", "created_by_id"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_policy.iam_policy", "last_modified_at"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_policy.iam_policy", "last_modified_by_id"),
				),
			},
		},
	})
}

func testAccCheckIbmIamPolicyDataSourceConfigBasic(policyType string) string {
	return fmt.Sprintf(`
		resource "ibm_iam_policy" "iam_policy" {
			type = "%s"
			subjects = { example: "object" }
			roles {
				role_id = "role_id"
			}
			resources = { example: "object" }
		}

		data "ibm_iam_policy" "iam_policy" {
			policy_id = "policy_id"
		}
	`, policyType)
}

func testAccCheckIbmIamPolicyDataSourceConfig(policyType string, policyDescription string, policyAcceptLanguage string) string {
	return fmt.Sprintf(`
		resource "ibm_iam_policy" "iam_policy" {
			type = "%s"
			subjects = { example: "object" }
			roles {
				role_id = "role_id"
			}
			resources = { example: "object" }
			description = "%s"
			Accept-Language = "%s"
		}

		data "ibm_iam_policy" "iam_policy" {
			policy_id = "policy_id"
		}
	`, policyType, policyDescription, policyAcceptLanguage)
}
