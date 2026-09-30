// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package iamidentity_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

// ---------------------------------------------------------------------------
// consumable — list IDPs available for the account to bind
// ---------------------------------------------------------------------------

func TestAccIBMIamIdpAccountSettingsDataSourceConsumable(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckIAMIdpAccountSetting(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIamIdpAccountSettingsDataSourceConfigConsumable(acc.IAMIdpConsumerAccountId),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_iam_idp_account_settings.consumable", "id"),
					resource.TestCheckResourceAttr("data.ibm_iam_idp_account_settings.consumable", "account_id", acc.IAMIdpConsumerAccountId),
					resource.TestCheckResourceAttr("data.ibm_iam_idp_account_settings.consumable", "type", "consumable"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_idp_account_settings.consumable", "idps.#"),
				),
			},
		},
	})
}

// ---------------------------------------------------------------------------
// consumed — list IDPs already bound to the account
// ---------------------------------------------------------------------------

func TestAccIBMIamIdpAccountSettingsDataSourceConsumed(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckIAMIdpAccountSetting(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				// First bind the IdP so "consumed" list is non-empty.
				Config: testAccCheckIBMIamIdpAccountSettingsDataSourceConfigConsumed(
					acc.IAMIdpAccountId,
					acc.IAMIdpConsumerAccountId,
				),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ibm_iam_idp_account_settings.consumed", "type", "consumed"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_idp_account_settings.consumed", "idps.#"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_idp_account_settings.consumed", "idps.0.idp_id"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_idp_account_settings.consumed", "idps.0.idp_name"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_idp_account_settings.consumed", "idps.0.cloud_user_strategy"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_idp_account_settings.consumed", "idps.0.owner_account"),
				),
			},
		},
	})
}

// ---------------------------------------------------------------------------
// Helper configs
// ---------------------------------------------------------------------------

func testAccCheckIBMIamIdpAccountSettingsDataSourceConfigConsumable(consumerAccountID string) string {
	return fmt.Sprintf(`
data "ibm_iam_idp_account_settings" "consumable" {
  account_id = "%s"
  type       = "consumable"
}
`, consumerAccountID)
}

func testAccCheckIBMIamIdpAccountSettingsDataSourceConfigConsumed(ownerAccountID, consumerAccountID string) string {
	return fmt.Sprintf(`
resource "ibm_iam_idp" "iam_idp_for_consumed_ds" {
  account_id = "%s"
  name       = "tf-consumed-ds-idp"
  type       = "saml"
  active     = true

  properties {
    idp {
      entity_id           = "http://www.okta.com/exampleidp"
      redirect_binding_url = "https://example.okta.com/app/example/sso/saml"
    }
  }

  secrets {}

  share_scope {
    id   = "%s"
    type = "account"
  }
}

resource "ibm_iam_idp_account_setting" "iam_idp_as_for_consumed_ds" {
  account_id          = "%s"
  idp_id              = ibm_iam_idp.iam_idp_for_consumed_ds.idp_id
  cloud_user_strategy = "DYNAMIC"
  active              = true
  ui_default          = false
}

data "ibm_iam_idp_account_settings" "consumed" {
  account_id = "%s"
  type       = "consumed"
  depends_on = [ibm_iam_idp_account_setting.iam_idp_as_for_consumed_ds]
}
`, ownerAccountID, consumerAccountID, consumerAccountID, consumerAccountID)
}
