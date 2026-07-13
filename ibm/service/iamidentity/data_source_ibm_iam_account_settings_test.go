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
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMIAMAccountSettingsDataSourceBasic(t *testing.T) {
	accountId := acc.IAMAccountId

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIamAccountSettingsDataSourceConfigBasic(accountId),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_iam_account_settings.iam_account_settings_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_account_settings.iam_account_settings_instance", "account_id"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_account_settings.iam_account_settings_instance", "restrict_create_service_id"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_account_settings.iam_account_settings_instance", "restrict_create_platform_apikey"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_account_settings.iam_account_settings_instance", "restrict_user_list_visibility"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_account_settings.iam_account_settings_instance", "restrict_user_domains.#"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_account_settings.iam_account_settings_instance", "entity_tag"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_account_settings.iam_account_settings_instance", "mfa"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_account_settings.iam_account_settings_instance", "user_mfa.#"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_account_settings.iam_account_settings_instance", "history.#"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_account_settings.iam_account_settings_instance", "session_expiration_in_seconds"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_account_settings.iam_account_settings_instance", "session_invalidation_in_seconds"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_account_settings.iam_account_settings_instance", "max_sessions_per_identity"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_account_settings.iam_account_settings_instance", "system_access_token_expiration_in_seconds"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_account_settings.iam_account_settings_instance", "system_refresh_token_expiration_in_seconds"),
				),
			},
		},
	})
}

func testAccCheckIBMIamAccountSettingsDataSourceConfigBasic(accountId string) string {
	return fmt.Sprintf(`
		data "ibm_iam_account_settings" "iam_account_settings_instance" {
			account_id = "%s"
		}
	`, accountId)
}
