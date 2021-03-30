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

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/IBM/platform-services-go-sdk/iamidentityv1"
)

func TestAccIbmIamAccountSettingsBasic(t *testing.T) {
	var conf iamidentityv1.AccountSettingsResponse

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		//CheckDestroy: testAccCheckIbmIamAccountSettingsDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				ExpectNonEmptyPlan: true,
				Config:             testAccCheckIbmIamAccountSettingsConfigBasic(),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmIamAccountSettingsExists("ibm_iam_account_settings.iam_account_settings", conf),
				),
			},
			resource.TestStep{
				ExpectNonEmptyPlan: true,
				Config:             testAccCheckIbmIamAccountSettingsConfigBasic(),
				Check:              resource.ComposeAggregateTestCheckFunc(),
			},
		},
	})
}

func TestAccIbmIamAccountSettingsAllArgs(t *testing.T) {
	var conf iamidentityv1.AccountSettingsResponse
	includeHistory := "false"
	includeHistoryUpdate := "true"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		//CheckDestroy: testAccCheckIbmIamAccountSettingsDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				ExpectNonEmptyPlan: true,
				Config:             testAccCheckIbmIamAccountSettingsConfig(includeHistory),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmIamAccountSettingsExists("ibm_iam_account_settings.iam_account_settings", conf),
					resource.TestCheckResourceAttr("ibm_iam_account_settings.iam_account_settings", "include_history", includeHistory),
				),
			},
			resource.TestStep{
				ExpectNonEmptyPlan: true,
				Config:             testAccCheckIbmIamAccountSettingsConfig(includeHistoryUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_iam_account_settings.iam_account_settings", "include_history", includeHistoryUpdate),
				),
			},
			resource.TestStep{
				ExpectNonEmptyPlan: true,
				ResourceName:       "ibm_iam_account_settings.iam_account_settings",
				ImportState:        true,
				ImportStateVerify:  false,
			},
		},
	})
}

func testAccCheckIbmIamAccountSettingsConfigBasic() string {
	return fmt.Sprintf(`

		resource "ibm_iam_account_settings" "iam_account_settings" {
		}
	`)
}

func testAccCheckIbmIamAccountSettingsConfig(includeHistory string) string {
	return fmt.Sprintf(`

		resource "ibm_iam_account_settings" "iam_account_settings" {
			include_history = %s
		}
	`, includeHistory)
}

func testAccCheckIbmIamAccountSettingsExists(n string, obj iamidentityv1.AccountSettingsResponse) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		iamIdentityClient, err := testAccProvider.Meta().(ClientSession).IamIdentityV1()
		if err != nil {
			return err
		}

		getAccountSettingsOptions := &iamidentityv1.GetAccountSettingsOptions{}

		parts, err := idParts(rs.Primary.ID)
		if err != nil {
			return err
		}

		getAccountSettingsOptions.SetAccountID(parts[0])
		getAccountSettingsOptions.SetAccountID(parts[1])

		accountSettingsResponse, _, err := iamIdentityClient.GetAccountSettings(getAccountSettingsOptions)
		if err != nil {
			return err
		}

		obj = *accountSettingsResponse
		return nil
	}
}

func testAccCheckIbmIamAccountSettingsDestroy(s *terraform.State) error {
	// NOT SUPPORTED
	return nil
}
