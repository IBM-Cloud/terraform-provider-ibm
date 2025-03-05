// Copyright IBM Corp. 2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package iampolicy_test

import (
	"fmt"

	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMIAMAccessManagementAccountSettings_Restore(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIAMAccessManagementAccountSettingsRestoreOriginal(acc.IAMAccountId),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.ibm_iam_access_management_account_settings.initial_settings", "external_account_identity_interaction.0.identity_types.0.user.0.state", "enabled"),
					resource.TestCheckResourceAttr("data.ibm_iam_access_management_account_settings.initial_settings", "external_account_identity_interaction.0.identity_types.0.user.0.external_allowed_accounts.#", "0"),
					resource.TestCheckResourceAttr("ibm_iam_access_management_account_settings.restored_settings", "external_account_identity_interaction.0.identity_types.0.user.0.state", "enabled"),
					resource.TestCheckResourceAttr("ibm_iam_access_management_account_settings.restored_settings", "external_account_identity_interaction.0.identity_types.0.user.0.external_allowed_accounts.#", "0"),
				),
			},
		},
	})
}

func testAccCheckIBMIAMAccessManagementAccountSettingsRestoreOriginal(accountID string) string {
	return fmt.Sprintf(`
		data "ibm_iam_access_management_account_settings" "initial_settings" {
			account_id = "%s"
		}

        resource "ibm_iam_access_management_account_settings" "restored_settings" {
			account_id = "%s"
		
			external_account_identity_interaction {
			  identity_types {
				user {
				  state                     = one(one(one(data.ibm_iam_access_management_account_settings.initial_settings.external_account_identity_interaction[*]).identity_types[*]).user[*]).state
				  external_allowed_accounts = one(one(one(data.ibm_iam_access_management_account_settings.initial_settings.external_account_identity_interaction[*]).identity_types[*]).user[*]).external_allowed_accounts
				}
		
				service_id {
			      state                     = one(one(one(data.ibm_iam_access_management_account_settings.initial_settings.external_account_identity_interaction[*]).identity_types[*]).service_id[*]).state
				  external_allowed_accounts = one(one(one(data.ibm_iam_access_management_account_settings.initial_settings.external_account_identity_interaction[*]).identity_types[*]).service_id[*]).external_allowed_accounts
			    }

                service {
			      state                     = one(one(one(data.ibm_iam_access_management_account_settings.initial_settings.external_account_identity_interaction[*]).identity_types[*]).service[*]).state
				  external_allowed_accounts = one(one(one(data.ibm_iam_access_management_account_settings.initial_settings.external_account_identity_interaction[*]).identity_types[*]).service[*]).external_allowed_accounts
			    }
			  }
			}
		}
	`, accountID, accountID)
}
