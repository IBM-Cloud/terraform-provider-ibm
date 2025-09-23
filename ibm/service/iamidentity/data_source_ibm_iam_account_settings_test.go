// Copyright IBM Corp. 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package iamidentity_test

import (
	"fmt"
	"testing"

	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMIAMAccountSettingsDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIbmIamAccountSettingsDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_iam_account_settings.iam_account_settings", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_account_settings.iam_account_settings", "account_id"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_account_settings.iam_account_settings", "restrict_create_service_id"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_account_settings.iam_account_settings", "restrict_create_platform_apikey"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_account_settings.iam_account_settings", "entity_tag"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_account_settings.iam_account_settings", "mfa"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_account_settings.iam_account_settings", "user_mfa.#"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_account_settings.iam_account_settings", "history.#"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_account_settings.iam_account_settings", "session_expiration_in_seconds"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_account_settings.iam_account_settings", "session_invalidation_in_seconds"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_account_settings.iam_account_settings", "max_sessions_per_identity"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_account_settings.iam_account_settings", "system_access_token_expiration_in_seconds"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_account_settings.iam_account_settings", "system_refresh_token_expiration_in_seconds"),
				),
			},
		},
	})
}

func testAccCheckIbmIamAccountSettingsDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		data "ibm_iam_account_settings" "iam_account_settings" {
		}
	`)
}
