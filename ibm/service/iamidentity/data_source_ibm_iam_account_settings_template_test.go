// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package iamidentity_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIBMAccountSettingsTemplateDataSourceBasic(t *testing.T) {
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	description := fmt.Sprintf("tf_desc_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMAccountSettingsTemplateDataSourceConfigBasic(name, description),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_iam_account_settings_template.account_settings_template", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_account_settings_template.account_settings_template", "template_id"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_account_settings_template.account_settings_template", "version"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_account_settings_template.account_settings_template", "account_id"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_account_settings_template.account_settings_template", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_account_settings_template.account_settings_template", "committed"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_account_settings_template.account_settings_template", "account_settings.#"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_account_settings_template.account_settings_template", "entity_tag"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_account_settings_template.account_settings_template", "crn"),
				),
			},
		},
	})
}

func TestAccIBMAccountSettingsTemplateDataSourceAllArgs(t *testing.T) {
	accountSettingsTemplateResponseName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	accountSettingsTemplateResponseDescription := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMAccountSettingsTemplateDataSourceConfig(accountSettingsTemplateResponseName, accountSettingsTemplateResponseDescription),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_iam_account_settings_template.account_settings_template", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_account_settings_template.account_settings_template", "version"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_account_settings_template.account_settings_template", "include_history"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_account_settings_template.account_settings_template", "account_id"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_account_settings_template.account_settings_template", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_account_settings_template.account_settings_template", "description"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_account_settings_template.account_settings_template", "committed"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_account_settings_template.account_settings_template", "account_settings.#"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_account_settings_template.account_settings_template", "account_settings.0.restrict_create_service_id"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_account_settings_template.account_settings_template", "account_settings.0.restrict_create_platform_apikey"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_account_settings_template.account_settings_template", "account_settings.0.allowed_ip_addresses"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_account_settings_template.account_settings_template", "account_settings.0.mfa"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_account_settings_template.account_settings_template", "account_settings.0.user_mfa.#"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_account_settings_template.account_settings_template", "account_settings.0.session_expiration_in_seconds"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_account_settings_template.account_settings_template", "account_settings.0.session_invalidation_in_seconds"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_account_settings_template.account_settings_template", "account_settings.0.max_sessions_per_identity"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_account_settings_template.account_settings_template", "account_settings.0.system_access_token_expiration_in_seconds"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_account_settings_template.account_settings_template", "account_settings.0.system_refresh_token_expiration_in_seconds"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_account_settings_template.account_settings_template", "entity_tag"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_account_settings_template.account_settings_template", "crn"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_account_settings_template.account_settings_template", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_account_settings_template.account_settings_template", "created_by_id"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_account_settings_template.account_settings_template", "last_modified_at"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_account_settings_template.account_settings_template", "last_modified_by_id"),
				),
			},
		},
	})
}

func testAccCheckIBMAccountSettingsTemplateDataSourceConfigBasic(name string, description string) string {
	return fmt.Sprintf(`
		resource "ibm_iam_account_settings_template" "account_settings_template" {
			name = "%s"
			description = "%s"
			account_settings {
				mfa = "LEVEL3"
			}
		}

		data "ibm_iam_account_settings_template" "account_settings_template" {
			template_id = ibm_iam_account_settings_template.account_settings_template.id
		}
	`, name, description)
}

func testAccCheckIBMAccountSettingsTemplateDataSourceConfig(accountSettingsTemplateResponseName string, accountSettingsTemplateResponseDescription string) string {
	return fmt.Sprintf(`
		resource "ibm_iam_account_settings_template" "account_settings_template" {
			name = "%s"
			description = "%s"
			account_settings {
				restrict_create_service_id = "RESTRICTED"
				restrict_create_platform_apikey = "RESTRICTED"
				allowed_ip_addresses = "127.0.0.1"
				mfa = "LEVEL1"
				user_mfa {
					iam_id = "IBMid-123456789"
					mfa = "NONE"
				}
				session_expiration_in_seconds = "1800"
				session_invalidation_in_seconds = "900"
				max_sessions_per_identity = "5"
				system_access_token_expiration_in_seconds = "NOT_SET"
				system_refresh_token_expiration_in_seconds = "NOT_SET"
			}
		}

		data "ibm_iam_account_settings_template" "account_settings_template" {
			template_id = ibm_iam_account_settings_template.account_settings_template.id
		}
	`, accountSettingsTemplateResponseName, accountSettingsTemplateResponseDescription)
}
