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

	"github.com/IBM/platform-services-go-sdk/iampolicymanagementv1"
)

func TestAccIbmIamPolicyBasic(t *testing.T) {
	var conf iampolicymanagementv1.Policy
	typeVar := fmt.Sprintf("type_%d", acctest.RandIntRange(10, 100))
	typeVarUpdate := fmt.Sprintf("type_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIbmIamPolicyDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmIamPolicyConfigBasic(typeVar),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmIamPolicyExists("ibm_iam_policy.iam_policy", conf),
					resource.TestCheckResourceAttr("ibm_iam_policy.iam_policy", "type", typeVar),
				),
			},
			resource.TestStep{
				Config: testAccCheckIbmIamPolicyConfigBasic(typeVarUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_iam_policy.iam_policy", "type", typeVarUpdate),
				),
			},
		},
	})
}

func TestAccIbmIamPolicyAllArgs(t *testing.T) {
	var conf iampolicymanagementv1.Policy
	typeVar := fmt.Sprintf("type_%d", acctest.RandIntRange(10, 100))
	description := fmt.Sprintf("description_%d", acctest.RandIntRange(10, 100))
	acceptLanguage := fmt.Sprintf("Accept-Language_%d", acctest.RandIntRange(10, 100))
	typeVarUpdate := fmt.Sprintf("type_%d", acctest.RandIntRange(10, 100))
	descriptionUpdate := fmt.Sprintf("description_%d", acctest.RandIntRange(10, 100))
	acceptLanguageUpdate := fmt.Sprintf("Accept-Language_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIbmIamPolicyDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmIamPolicyConfig(typeVar, description, acceptLanguage),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmIamPolicyExists("ibm_iam_policy.iam_policy", conf),
					resource.TestCheckResourceAttr("ibm_iam_policy.iam_policy", "type", typeVar),
					resource.TestCheckResourceAttr("ibm_iam_policy.iam_policy", "description", description),
					resource.TestCheckResourceAttr("ibm_iam_policy.iam_policy", "Accept-Language", acceptLanguage),
				),
			},
			resource.TestStep{
				Config: testAccCheckIbmIamPolicyConfig(typeVarUpdate, descriptionUpdate, acceptLanguageUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_iam_policy.iam_policy", "type", typeVarUpdate),
					resource.TestCheckResourceAttr("ibm_iam_policy.iam_policy", "description", descriptionUpdate),
					resource.TestCheckResourceAttr("ibm_iam_policy.iam_policy", "Accept-Language", acceptLanguageUpdate),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_iam_policy.iam_policy",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIbmIamPolicyConfigBasic(typeVar string) string {
	return fmt.Sprintf(`

		resource "ibm_iam_policy" "iam_policy" {
			type = "%s"
			subjects = { example: "object" }
			roles {
				role_id = "role_id"
			}
			resources = { example: "object" }
		}
	`, typeVar)
}

func testAccCheckIbmIamPolicyConfig(typeVar string, description string, acceptLanguage string) string {
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
	`, typeVar, description, acceptLanguage)
}

func testAccCheckIbmIamPolicyExists(n string, obj iampolicymanagementv1.Policy) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		iamPolicyManagementClient, err := testAccProvider.Meta().(ClientSession).IamPolicyManagementV1()
		if err != nil {
			return err
		}

		getPolicyOptions := &iampolicymanagementv1.GetPolicyOptions{}

		getPolicyOptions.SetPolicyID(rs.Primary.ID)

		policy, _, err := iamPolicyManagementClient.GetPolicy(getPolicyOptions)
		if err != nil {
			return err
		}

		obj = *policy
		return nil
	}
}

func testAccCheckIbmIamPolicyDestroy(s *terraform.State) error {
	iamPolicyManagementClient, err := testAccProvider.Meta().(ClientSession).IamPolicyManagementV1()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_iam_policy" {
			continue
		}

		getPolicyOptions := &iampolicymanagementv1.GetPolicyOptions{}

		getPolicyOptions.SetPolicyID(rs.Primary.ID)

		// Try to find the key
		_, response, err := iamPolicyManagementClient.GetPolicy(getPolicyOptions)

		if err == nil {
			return fmt.Errorf("iam_policy still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for iam_policy (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
