// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package iamidentity_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/platform-services-go-sdk/iamidentityv1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

// ---------------------------------------------------------------------------
// Acceptance tests
// ---------------------------------------------------------------------------

func TestAccIBMIamIdpAccountSettingBasic(t *testing.T) {
	var conf iamidentityv1.AccountIdpSettings
	name := fmt.Sprintf("tf_idp_binding_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheckIAMIdpAccountSetting(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIamIdpAccountSettingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIamIdpAccountSettingConfigBasic(acc.IAMIdpAccountId, name, acc.IAMIdpConsumerAccountId),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIamIdpAccountSettingExists("ibm_iam_idp_account_setting.iam_idp_account_setting_instance", conf),
					resource.TestCheckResourceAttrSet("ibm_iam_idp_account_setting.iam_idp_account_setting_instance", "id"),
					resource.TestCheckResourceAttr("ibm_iam_idp_account_setting.iam_idp_account_setting_instance", "account_id", acc.IAMIdpConsumerAccountId),
					resource.TestCheckResourceAttr("ibm_iam_idp_account_setting.iam_idp_account_setting_instance", "cloud_user_strategy", "DYNAMIC"),
					resource.TestCheckResourceAttr("ibm_iam_idp_account_setting.iam_idp_account_setting_instance", "active", "true"),
					resource.TestCheckResourceAttr("ibm_iam_idp_account_setting.iam_idp_account_setting_instance", "ui_default", "false"),
				),
			},
		},
	})
}

func TestAccIBMIamIdpAccountSettingUpdate(t *testing.T) {
	var conf iamidentityv1.AccountIdpSettings
	name := fmt.Sprintf("tf_idp_binding_upd_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheckIAMIdpAccountSetting(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIamIdpAccountSettingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIamIdpAccountSettingConfigBasic(acc.IAMIdpAccountId, name, acc.IAMIdpConsumerAccountId),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIamIdpAccountSettingExists("ibm_iam_idp_account_setting.iam_idp_account_setting_instance", conf),
					resource.TestCheckResourceAttr("ibm_iam_idp_account_setting.iam_idp_account_setting_instance", "cloud_user_strategy", "DYNAMIC"),
					resource.TestCheckResourceAttr("ibm_iam_idp_account_setting.iam_idp_account_setting_instance", "active", "true"),
				),
			},
			{
				Config: testAccCheckIBMIamIdpAccountSettingConfigUpdated(acc.IAMIdpAccountId, name, acc.IAMIdpConsumerAccountId),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_iam_idp_account_setting.iam_idp_account_setting_instance", "cloud_user_strategy", "STATIC"),
					resource.TestCheckResourceAttr("ibm_iam_idp_account_setting.iam_idp_account_setting_instance", "active", "false"),
					resource.TestCheckResourceAttr("ibm_iam_idp_account_setting.iam_idp_account_setting_instance", "ui_default", "false"),
				),
			},
			{
				ResourceName:      "ibm_iam_idp_account_setting.iam_idp_account_setting_instance",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

// ---------------------------------------------------------------------------
// Acceptance-test helper configs
// ---------------------------------------------------------------------------

// testAccCheckIBMIamIdpAccountSettingConfigBasic creates the owner IdP with the
// consumer in share_scope, then binds it to the consumer account.
func testAccCheckIBMIamIdpAccountSettingConfigBasic(ownerAccountID, idpName, consumerAccountID string) string {
	return fmt.Sprintf(`
resource "ibm_iam_idp" "iam_idp_for_binding" {
  account_id = "%s"
  name       = "%s"
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

resource "ibm_iam_idp_account_setting" "iam_idp_account_setting_instance" {
  account_id          = "%s"
  idp_id              = ibm_iam_idp.iam_idp_for_binding.idp_id
  cloud_user_strategy = "DYNAMIC"
  active              = true
  ui_default          = false
}
`, ownerAccountID, idpName, consumerAccountID, consumerAccountID)
}

func testAccCheckIBMIamIdpAccountSettingConfigUpdated(ownerAccountID, idpName, consumerAccountID string) string {
	return fmt.Sprintf(`
resource "ibm_iam_idp" "iam_idp_for_binding" {
  account_id = "%s"
  name       = "%s"
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

resource "ibm_iam_idp_account_setting" "iam_idp_account_setting_instance" {
  account_id          = "%s"
  idp_id              = ibm_iam_idp.iam_idp_for_binding.idp_id
  cloud_user_strategy = "STATIC"
  active              = false
  ui_default          = false
}
`, ownerAccountID, idpName, consumerAccountID, consumerAccountID)
}

// ---------------------------------------------------------------------------
// Acceptance-test check functions
// ---------------------------------------------------------------------------

func testAccCheckIBMIamIdpAccountSettingExists(n string, obj iamidentityv1.AccountIdpSettings) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		iamIdentityClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).IAMIdentityV1API()
		if err != nil {
			return err
		}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getOpts := iamIdentityClient.NewGetIDPSettingOptions(parts[0], parts[1])
		setting, _, err := iamIdentityClient.GetIDPSetting(getOpts)
		if err != nil {
			return err
		}

		obj = *setting
		return nil
	}
}

func testAccCheckIBMIamIdpAccountSettingDestroy(s *terraform.State) error {
	iamIdentityClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).IAMIdentityV1API()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_iam_idp_account_setting" {
			continue
		}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getOpts := iamIdentityClient.NewGetIDPSettingOptions(parts[0], parts[1])
		_, response, err := iamIdentityClient.GetIDPSetting(getOpts)
		if err == nil {
			return fmt.Errorf("ibm_iam_idp_account_setting still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking ibm_iam_idp_account_setting (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}
	return nil
}
