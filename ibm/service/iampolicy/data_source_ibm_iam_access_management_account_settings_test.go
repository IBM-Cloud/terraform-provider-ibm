// Copyright IBM Corp. 2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package iampolicy_test

import (
	"fmt"

	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMIAMAccountSettingsExternalInteraction_Fetch(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIAMAccountSettingsExternalInteractionFetch(acc.IAMAccountId),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.ibm_iam_account_settings_external_interaction.initial_settings", "external_account_identity_interaction.0.identity_types.0.user.0.state", "enabled"),
					resource.TestCheckResourceAttr("data.ibm_iam_account_settings_external_interaction.initial_settings", "external_account_identity_interaction.0.identity_types.0.user.0.external_allowed_accounts.#", "0"),
					resource.TestCheckResourceAttr("data.ibm_iam_account_settings_external_interaction.initial_settings", "external_account_identity_interaction.0.identity_types.0.service_id.0.state", "enabled"),
					resource.TestCheckResourceAttr("data.ibm_iam_account_settings_external_interaction.initial_settings", "external_account_identity_interaction.0.identity_types.0.service_id.0.external_allowed_accounts.#", "0"),
					resource.TestCheckResourceAttr("data.ibm_iam_account_settings_external_interaction.initial_settings", "external_account_identity_interaction.0.identity_types.0.service.0.state", "enabled"),
					resource.TestCheckResourceAttr("data.ibm_iam_account_settings_external_interaction.initial_settings", "external_account_identity_interaction.0.identity_types.0.service.0.external_allowed_accounts.#", "0"),
				),
			},
		},
	})
}

func testAccCheckIBMIAMAccountSettingsExternalInteractionFetch(accountID string) string {
	return fmt.Sprintf(`
		data "ibm_iam_account_settings_external_interaction" "initial_settings" {
			account_id = "%s"
		}
	`, accountID)
}
