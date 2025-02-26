// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0
package iampolicy_test

import (
	"fmt"
	// "regexp"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	// "github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	// "github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"

	// "github.com/IBM/platform-services-go-sdk/iampolicymanagementv1"
	// "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	// "github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccIBMIAMAccessManagementAccountSettings_User_Basic(t *testing.T) {
	// var conf iampolicymanagementv1.V2PolicyTemplateMetaData

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		// CheckDestroy: testAccCheckIBMIAMAccessManagementAccountSettingsDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIAMAccessManagementAccountSettingsUserBasic(acc.IAMAccountId),
				Check: resource.ComposeAggregateTestCheckFunc(
					// testAccCheckIBMIAMAccessManagementAccountSettingsExists("ibm_iam_access_management_account_settings.settings", conf),
					resource.TestCheckResourceAttr("ibm_iam_access_management_account_settings.settings", "external_account_identity_interaction.identity_types.user.external_allowed_accounts.#", "0"),
					resource.TestCheckResourceAttrSet("ibm_iam_access_management_account_settings.settings", "external_account_identity_interaction"),
				),
			},
			// {
			// 	Config: testAccCheckIBMIAMAccessManagementAccountSettingsUserUpdate(),
			// 	Check: resource.ComposeAggregateTestCheckFunc(
			// 		resource.TestCheckResourceAttr("ibm_iam_access_management_account_settings.settings", "tags.#", "2"),
			// 		resource.TestCheckResourceAttr("ibm_iam_access_management_account_settings.settings", "roles.#", "2"),
			// 	),
			// },
		},
	})
}

// func testAccCheckIBMIAMAccessManagementAccountSettingsDestroy(s *terraform.State) error {
// 	rsContClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).IAMPolicyManagementV1API()
// 	if err != nil {
// 		return err
// 	}
// 	for _, rs := range s.RootModule().Resources {
// 		if rs.Type != "ibm_iam_user_policy" {
// 			continue
// 		}
// 		policyID := rs.Primary.ID
// 		parts, err := flex.IdParts(policyID)
// 		if err != nil {
// 			return err
// 		}

// 		userPolicyID := parts[1]

// 		getPolicyOptions := rsContClient.NewGetV2PolicyOptions(
// 			userPolicyID,
// 		)

// 		// Try to find the key
// 		destroyedPolicy, response, err := rsContClient.GetV2Policy(getPolicyOptions)

// 		if err == nil && *destroyedPolicy.State != "deleted" {
// 			return fmt.Errorf("User policy still exists: %s\n", rs.Primary.ID)
// 		} else if response.StatusCode != 404 && destroyedPolicy.State != nil && *destroyedPolicy.State != "deleted" {
// 			return fmt.Errorf("[ERROR] Error waiting for user policy (%s) to be destroyed: %s", rs.Primary.ID, err)
// 		}
// 	}

// 	return nil
// }

func testAccCheckIBMIAMAccessManagementAccountSettingsUserBasic(accountID string) string {
	fmt.Printf("howdy accountID: %+v", accountID)
	return fmt.Sprintf(` 
		resource "ibm_iam_access_management_account_settings" "settings" {
			account_id = "%s"
			external_account_identity_interaction = <<EOF
			{
			  "identity_types":{
			    "user":{
			      "state":"monitor",
				    "external_allowed_accounts":[]
			    }
		    }
			}
			EOF
	  }
	`, accountID)
}

// func testAccCheckIBMIAMAccessManagementAccountSettingsUserUpdate(accountId string) string {
// 	return fmt.Sprintf(`
		
// 		resource "ibm_iam_access_management_account_settings" "settings" {
// 			external_account_identity_interaction {
// 				identity_types {
// 					user {
// 						state = "monitor"
// 						external_allowed_accounts = ["account1"]
// 					}
// 				}
// 			}
// 	  }
// 	`)
// }