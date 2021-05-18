// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMIAMAccountSettingsDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmIamAccountSettingsDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_iam_account_settings.iam_account_settings", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_account_settings.iam_account_settings", "account_id"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_account_settings.iam_account_settings", "restrict_create_service_id"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_account_settings.iam_account_settings", "restrict_create_platform_apikey"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_account_settings.iam_account_settings", "entity_tag"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_account_settings.iam_account_settings", "mfa"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_account_settings.iam_account_settings", "history.#"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_account_settings.iam_account_settings", "session_expiration_in_seconds"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_account_settings.iam_account_settings", "session_invalidation_in_seconds"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_account_settings.iam_account_settings", "max_sessions_per_identity"),
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
