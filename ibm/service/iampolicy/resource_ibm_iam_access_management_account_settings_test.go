// Copyright IBM Corp. 2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package iampolicy_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMIAMAccessManagementAccountSettings_User(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIAMAccessManagementAccountSettingsRead(acc.IAMAccountId),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_iam_access_management_account_settings.settings", "external_account_identity_interaction.#", "1"),
					resource.TestCheckResourceAttr("ibm_iam_access_management_account_settings.settings", "external_account_identity_interaction.0.identity_types.#", "1"),
					resource.TestCheckResourceAttr("ibm_iam_access_management_account_settings.settings", "external_account_identity_interaction.0.identity_types.0.user.#", "1"),
					resource.TestCheckResourceAttr("ibm_iam_access_management_account_settings.settings", "external_account_identity_interaction.0.identity_types.0.user.0.state", "enabled"),
					resource.TestCheckResourceAttr("ibm_iam_access_management_account_settings.settings", "external_account_identity_interaction.0.identity_types.0.user.0.external_allowed_accounts.#", "0"),
				),
			},
			{
				Config: testAccCheckIBMIAMAccessManagementAccountSettingsUpdate(acc.IAMAccountId, "user", "service_id", "service"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_iam_access_management_account_settings.settings", "external_account_identity_interaction.#", "1"),
					resource.TestCheckResourceAttr("ibm_iam_access_management_account_settings.settings", "external_account_identity_interaction.0.identity_types.#", "1"),
					resource.TestCheckResourceAttr("ibm_iam_access_management_account_settings.settings", "external_account_identity_interaction.0.identity_types.0.user.#", "1"),
					resource.TestCheckResourceAttr("ibm_iam_access_management_account_settings.settings", "external_account_identity_interaction.0.identity_types.0.user.0.state", "monitor"),
					resource.TestCheckResourceAttr("ibm_iam_access_management_account_settings.settings", "external_account_identity_interaction.0.identity_types.0.user.0.external_allowed_accounts.#", "2"),
					resource.TestCheckResourceAttr("ibm_iam_access_management_account_settings.settings", "external_account_identity_interaction.0.identity_types.0.user.0.external_allowed_accounts.0", "account_id_01"),
					resource.TestCheckResourceAttr("ibm_iam_access_management_account_settings.settings", "external_account_identity_interaction.0.identity_types.0.user.0.external_allowed_accounts.1", "account_id_02"),
				),
			},
		},
	})
}

func TestAccIBMIAMAccessManagementAccountSettings_ServiceID(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIAMAccessManagementAccountSettingsRead(acc.IAMAccountId),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_iam_access_management_account_settings.settings", "external_account_identity_interaction.#", "1"),
					resource.TestCheckResourceAttr("ibm_iam_access_management_account_settings.settings", "external_account_identity_interaction.0.identity_types.#", "1"),
					resource.TestCheckResourceAttr("ibm_iam_access_management_account_settings.settings", "external_account_identity_interaction.0.identity_types.0.service_id.#", "1"),
					resource.TestCheckResourceAttr("ibm_iam_access_management_account_settings.settings", "external_account_identity_interaction.0.identity_types.0.service_id.0.state", "enabled"),
					resource.TestCheckResourceAttr("ibm_iam_access_management_account_settings.settings", "external_account_identity_interaction.0.identity_types.0.service_id.0.external_allowed_accounts.#", "0"),
				),
			},
			{
				Config: testAccCheckIBMIAMAccessManagementAccountSettingsUpdate(acc.IAMAccountId, "service_id", "user", "service"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_iam_access_management_account_settings.settings", "external_account_identity_interaction.#", "1"),
					resource.TestCheckResourceAttr("ibm_iam_access_management_account_settings.settings", "external_account_identity_interaction.0.identity_types.#", "1"),
					resource.TestCheckResourceAttr("ibm_iam_access_management_account_settings.settings", "external_account_identity_interaction.0.identity_types.0.service_id.#", "1"),
					resource.TestCheckResourceAttr("ibm_iam_access_management_account_settings.settings", "external_account_identity_interaction.0.identity_types.0.service_id.0.state", "monitor"),
					resource.TestCheckResourceAttr("ibm_iam_access_management_account_settings.settings", "external_account_identity_interaction.0.identity_types.0.service_id.0.external_allowed_accounts.#", "2"),
					resource.TestCheckResourceAttr("ibm_iam_access_management_account_settings.settings", "external_account_identity_interaction.0.identity_types.0.service_id.0.external_allowed_accounts.0", "account_id_01"),
					resource.TestCheckResourceAttr("ibm_iam_access_management_account_settings.settings", "external_account_identity_interaction.0.identity_types.0.service_id.0.external_allowed_accounts.1", "account_id_02"),
				),
			},
		},
	})
}

func TestAccIBMIAMAccessManagementAccountSettings_Service(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIAMAccessManagementAccountSettingsRead(acc.IAMAccountId),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_iam_access_management_account_settings.settings", "external_account_identity_interaction.#", "1"),
					resource.TestCheckResourceAttr("ibm_iam_access_management_account_settings.settings", "external_account_identity_interaction.0.identity_types.#", "1"),
					resource.TestCheckResourceAttr("ibm_iam_access_management_account_settings.settings", "external_account_identity_interaction.0.identity_types.0.service.#", "1"),
					resource.TestCheckResourceAttr("ibm_iam_access_management_account_settings.settings", "external_account_identity_interaction.0.identity_types.0.service.0.state", "enabled"),
					resource.TestCheckResourceAttr("ibm_iam_access_management_account_settings.settings", "external_account_identity_interaction.0.identity_types.0.service.0.external_allowed_accounts.#", "0"),
				),
			},
			{
				Config: testAccCheckIBMIAMAccessManagementAccountSettingsUpdate(acc.IAMAccountId, "service", "user", "service_id"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_iam_access_management_account_settings.settings", "external_account_identity_interaction.#", "1"),
					resource.TestCheckResourceAttr("ibm_iam_access_management_account_settings.settings", "external_account_identity_interaction.0.identity_types.#", "1"),
					resource.TestCheckResourceAttr("ibm_iam_access_management_account_settings.settings", "external_account_identity_interaction.0.identity_types.0.service.#", "1"),
					resource.TestCheckResourceAttr("ibm_iam_access_management_account_settings.settings", "external_account_identity_interaction.0.identity_types.0.service.0.state", "monitor"),
					resource.TestCheckResourceAttr("ibm_iam_access_management_account_settings.settings", "external_account_identity_interaction.0.identity_types.0.service.0.external_allowed_accounts.#", "2"),
					resource.TestCheckResourceAttr("ibm_iam_access_management_account_settings.settings", "external_account_identity_interaction.0.identity_types.0.service.0.external_allowed_accounts.0", "account_id_01"),
					resource.TestCheckResourceAttr("ibm_iam_access_management_account_settings.settings", "external_account_identity_interaction.0.identity_types.0.service.0.external_allowed_accounts.1", "account_id_02"),
				),
			},
		},
	})
}

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

func testAccCheckIBMIAMAccessManagementAccountSettingsRead(accountID string) string {
	return fmt.Sprintf(` 
		resource "ibm_iam_access_management_account_settings" "settings" {
			account_id = "%s"

			external_account_identity_interaction {
			  identity_types {
			    user {
			      state                     = "enabled"
                  external_allowed_accounts = []
			    }

                service_id {
			      state                     = "enabled"
                  external_allowed_accounts = []
			    }

                service {
			      state                     = "enabled"
                  external_allowed_accounts = []
			    }
              }
            }
	  }
	`, accountID)
}

func testAccCheckIBMIAMAccessManagementAccountSettingsUpdate(accountID string, type1 string, type2 string, type3 string) string {
	return fmt.Sprintf(` 
		resource "ibm_iam_access_management_account_settings" "settings" {
			account_id = "%s"

			external_account_identity_interaction {
			  identity_types {
			    %s {
			      state                     = "monitor"
                  external_allowed_accounts = ["account_id_01", "account_id_02"]
			    }

                %s {
			      state                     = "enabled"
                  external_allowed_accounts = []
			    }

                %s {
			      state                     = "enabled"
                  external_allowed_accounts = []
			    }
              }
            }
	  }
	`, accountID, type1, type2, type3)
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
