// Copyright IBM Corp. 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package iamidentity_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/IBM/platform-services-go-sdk/iamidentityv1"
)

var (
	restrict_create_service_id                 = "NOT_SET"
	restrict_create_platform_apikey            = "NOT_SET"
	restrict_user_list_visibility              = "NOT_RESTRICTED"
	entity_tag                                 = "*"
	mfa_trait                                  = "NONE"
	session_expiration_in_seconds              = "NOT_SET"
	session_invalidation_in_seconds            = "NOT_SET"
	max_sessions_per_identity                  = "NOT_SET"
	system_access_token_expiration_in_seconds  = "3600"
	system_refresh_token_expiration_in_seconds = "259200"
)

func TestAccIBMIAMAccountSettingsBasic(t *testing.T) {
	var conf iamidentityv1.AccountSettingsResponse

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIbmIamAccountSettingsConfigBasic(),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmIamAccountSettingsExists("ibm_iam_account_settings.iam_account_settings", conf),
				),
			},
			{
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
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIbmIamAccountSettingsConfig(includeHistory, "*@companya.com"),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmIamAccountSettingsExists("ibm_iam_account_settings.iam_account_settings", conf),
					resource.TestCheckResourceAttr("ibm_iam_account_settings.iam_account_settings", "include_history", includeHistory),
				),
			},
			{
				Config: testAccCheckIbmIamAccountSettingsConfig(includeHistoryUpdate, "*@companyab.com"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_iam_account_settings.iam_account_settings", "include_history", includeHistoryUpdate),
				),
			},
			{
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
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIbmIamAccountSettingsUpdateConfig(),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmIamAccountSettingsExists("ibm_iam_account_settings.iam_account_settings", conf),
					resource.TestCheckResourceAttr("ibm_iam_account_settings.iam_account_settings", "restrict_create_service_id", restrict_create_service_id),
					resource.TestCheckResourceAttr("ibm_iam_account_settings.iam_account_settings", "restrict_create_platform_apikey", restrict_create_platform_apikey),
					resource.TestCheckResourceAttr("ibm_iam_account_settings.iam_account_settings", "restrict_user_list_visibility", restrict_user_list_visibility),
					resource.TestCheckResourceAttr("ibm_iam_account_settings.iam_account_settings", "mfa", mfa_trait),
					resource.TestCheckResourceAttr("ibm_iam_account_settings.iam_account_settings", "session_expiration_in_seconds", session_expiration_in_seconds),
					resource.TestCheckResourceAttr("ibm_iam_account_settings.iam_account_settings", "session_invalidation_in_seconds", session_invalidation_in_seconds),
					resource.TestCheckResourceAttr("ibm_iam_account_settings.iam_account_settings", "max_sessions_per_identity", max_sessions_per_identity),
					resource.TestCheckResourceAttr("ibm_iam_account_settings.iam_account_settings", "system_access_token_expiration_in_seconds", system_access_token_expiration_in_seconds),
					resource.TestCheckResourceAttr("ibm_iam_account_settings.iam_account_settings", "system_refresh_token_expiration_in_seconds", system_refresh_token_expiration_in_seconds),
				),
			},
		},
	})
}

func testAccCheckIbmIamAccountSettingsConfigBasic() string {
	return `

		resource "ibm_iam_account_settings" "iam_account_settings" {
		}
	`
}

func testAccCheckIbmIamAccountSettingsConfig(includeHistory string, emailPatterns string) string {
	return fmt.Sprintf(`

		resource "ibm_iam_account_settings" "iam_account_settings" {
			include_history = %s
			restrict_user_domains {
				realm_id                        = "IBMid"
				restrict_invitation             = false
				invitation_email_allow_patterns = ["%s"]
			}
		}
	`, includeHistory, emailPatterns)
}

func testAccCheckIbmIamAccountSettingsUpdateConfig() string {
	return fmt.Sprintf(`

		resource "ibm_iam_account_settings" "iam_account_settings" {
			restrict_create_service_id = "%s"
			restrict_create_platform_apikey = "%s"
			restrict_user_list_visibility = "%s"
			if_match = "%s"
			mfa = "%s"
			user_mfa {
				iam_id = "%s"
				mfa = "NONE"
			}
			restrict_user_domains {
				realm_id = "IBMid"
				invitation_email_allow_patterns = [
					"*@ibm.com",
					"**@corp.org"
					]
				restrict_invitation = false
			}
			session_expiration_in_seconds = "%s"
			session_invalidation_in_seconds = "%s"
			max_sessions_per_identity = "%s"
			system_access_token_expiration_in_seconds = "%s"
			system_refresh_token_expiration_in_seconds = "%s"
		}
	`,
		restrict_create_service_id,
		restrict_create_platform_apikey,
		restrict_user_list_visibility,
		entity_tag,
		mfa_trait,
		acc.Ibmid1,
		session_expiration_in_seconds,
		session_invalidation_in_seconds,
		max_sessions_per_identity,
		system_access_token_expiration_in_seconds,
		system_refresh_token_expiration_in_seconds,
	)
}

func testAccCheckIbmIamAccountSettingsUpdateConfigWithNoUserMfa() string {
	return fmt.Sprintf(`

		resource "ibm_iam_account_settings" "iam_account_settings" {
			restrict_create_service_id = "%s"
			restrict_create_platform_apikey = "%s"
			if_match = "%s"
			mfa = "%s"
			user_mfa {
			}
			session_expiration_in_seconds = "%s"
			session_invalidation_in_seconds = "%s"
			max_sessions_per_identity = "%s"
			system_access_token_expiration_in_seconds = "%s"
			system_refresh_token_expiration_in_seconds = "%s"
		}
	`,
		restrict_create_service_id,
		restrict_create_platform_apikey,
		entity_tag,
		mfa_trait,
		session_expiration_in_seconds,
		session_invalidation_in_seconds,
		max_sessions_per_identity,
		system_access_token_expiration_in_seconds,
		system_refresh_token_expiration_in_seconds,
	)
}

func testAccCheckIbmIamAccountSettingsUpdateConfigWithMultipleUserMfa() string {
	return fmt.Sprintf(`

		resource "ibm_iam_account_settings" "iam_account_settings" {
			restrict_create_service_id = "%s"
			restrict_create_platform_apikey = "%s"
			if_match = "%s"
			mfa = "%s"
			user_mfa {
				iam_id = "iam_id"
				mfa = "NONE"
			}
			user_mfa {
				iam_id = "iam_id"
				mfa = "NONE"
			}
			session_expiration_in_seconds = "%s"
			session_invalidation_in_seconds = "%s"
			max_sessions_per_identity = "%s"
			system_access_token_expiration_in_seconds = "%s"
			system_refresh_token_expiration_in_seconds = "%s"
		}
	`,
		restrict_create_service_id,
		restrict_create_platform_apikey,
		entity_tag,
		mfa_trait,
		session_expiration_in_seconds,
		session_invalidation_in_seconds,
		max_sessions_per_identity,
		system_access_token_expiration_in_seconds,
		system_refresh_token_expiration_in_seconds,
	)
}

func testAccCheckIbmIamAccountSettingsExists(n string, obj iamidentityv1.AccountSettingsResponse) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		iamIdentityClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).IAMIdentityV1API()
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
