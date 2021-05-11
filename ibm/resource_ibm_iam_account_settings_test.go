// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/IBM/platform-services-go-sdk/iamidentityv1"
)

var (
	restrict_create_service_id      = "NOT_SET"
	restrict_create_platform_apikey = "NOT_SET"
	entity_tag                      = "*"
	mfa_trait                       = "NONE"
	session_expiration_in_seconds   = "NOT_SET"
	session_invalidation_in_seconds = "NOT_SET"
	max_sessions_per_identity       = "NOT_SET"
)

func TestAccIBMIAMAccountSettingsBasic(t *testing.T) {
	var conf iamidentityv1.AccountSettingsResponse

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmIamAccountSettingsConfigBasic(),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmIamAccountSettingsExists("ibm_iam_account_settings.iam_account_settings", conf),
				),
			},
			resource.TestStep{
				Config: testAccCheckIbmIamAccountSettingsConfigBasic(),
				Check:  resource.ComposeAggregateTestCheckFunc(),
			},
		},
	})
}

func TestAccIBMIAMAccountSettingsAllArgs(t *testing.T) {
	var conf iamidentityv1.AccountSettingsResponse
	includeHistory := "false"
	includeHistoryUpdate := "true"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmIamAccountSettingsConfig(includeHistory),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmIamAccountSettingsExists("ibm_iam_account_settings.iam_account_settings", conf),
					resource.TestCheckResourceAttr("ibm_iam_account_settings.iam_account_settings", "include_history", includeHistory),
				),
			},
			resource.TestStep{
				Config: testAccCheckIbmIamAccountSettingsConfig(includeHistoryUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_iam_account_settings.iam_account_settings", "include_history", includeHistoryUpdate),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_iam_account_settings.iam_account_settings",
				ImportState:       true,
				ImportStateVerify: false,
			},
		},
	})
}

func TestAccIBMIAMAccountSettingsUpdate(t *testing.T) {
	var conf iamidentityv1.AccountSettingsResponse

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmIamAccountSettingsUpdateConfig(),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmIamAccountSettingsExists("ibm_iam_account_settings.iam_account_settings", conf),
					resource.TestCheckResourceAttr("ibm_iam_account_settings.iam_account_settings", "restrict_create_service_id", restrict_create_service_id),
					resource.TestCheckResourceAttr("ibm_iam_account_settings.iam_account_settings", "restrict_create_platform_apikey", restrict_create_platform_apikey),
					resource.TestCheckResourceAttr("ibm_iam_account_settings.iam_account_settings", "mfa", mfa_trait),
					resource.TestCheckResourceAttr("ibm_iam_account_settings.iam_account_settings", "session_expiration_in_seconds", session_expiration_in_seconds),
					resource.TestCheckResourceAttr("ibm_iam_account_settings.iam_account_settings", "session_invalidation_in_seconds", session_invalidation_in_seconds),
					resource.TestCheckResourceAttr("ibm_iam_account_settings.iam_account_settings", "max_sessions_per_identity", max_sessions_per_identity),
				),
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

func testAccCheckIbmIamAccountSettingsUpdateConfig() string {
	return fmt.Sprintf(`

		resource "ibm_iam_account_settings" "iam_account_settings" {
			restrict_create_service_id = "%s"
			restrict_create_platform_apikey = "%s"
			if_match = "%s"
			mfa = "%s"
			session_expiration_in_seconds = "%s"
			session_invalidation_in_seconds = "%s"
			max_sessions_per_identity = "%s"
		}
	`,
		restrict_create_service_id,
		restrict_create_platform_apikey,
		entity_tag,
		mfa_trait,
		session_expiration_in_seconds,
		session_invalidation_in_seconds,
		max_sessions_per_identity,
	)
}

func testAccCheckIbmIamAccountSettingsExists(n string, obj iamidentityv1.AccountSettingsResponse) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		iamIdentityClient, err := testAccProvider.Meta().(ClientSession).IAMIdentityV1API()
		if err != nil {
			return err
		}

		getAccountSettingsOptions := &iamidentityv1.GetAccountSettingsOptions{}
		getAccountSettingsOptions.SetAccountID(rs.Primary.ID)

		accountSettingsResponse, _, err := iamIdentityClient.GetAccountSettings(getAccountSettingsOptions)
		if err != nil {
			return err
		}

		entity_tag = *accountSettingsResponse.EntityTag

		obj = *accountSettingsResponse

		return nil
	}
}

func testAccCheckIbmIamAccountSettingsDestroy(s *terraform.State) error {
	// NOT SUPPORTED
	return nil
}
