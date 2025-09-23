// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package iamidentity_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/IBM/platform-services-go-sdk/iamidentityv1"
	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/conns"
)

func TestAccIBMAccountSettingsTemplateBasic(t *testing.T) {
	var conf iamidentityv1.AccountSettingsTemplateResponse
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	description := fmt.Sprintf("tf_desc_%d", acctest.RandIntRange(10, 100))
	descriptionUpdate := fmt.Sprintf("tf_desc_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMAccountSettingsTemplateDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMAccountSettingsTemplateConfigBasic(name, description),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMAccountSettingsTemplateExists("ibm_iam_account_settings_template.account_settings_template_instance", conf),
					resource.TestCheckResourceAttr("ibm_iam_account_settings_template.account_settings_template_instance", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_account_settings_template.account_settings_template_instance", "description", description),
				),
			},
			{
				Config: testAccCheckIBMAccountSettingsTemplateConfigBasic(name, descriptionUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMAccountSettingsTemplateExists("ibm_iam_account_settings_template.account_settings_template_instance", conf),
					resource.TestCheckResourceAttr("ibm_iam_account_settings_template.account_settings_template_instance", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_account_settings_template.account_settings_template_instance", "description", descriptionUpdate),
				),
			},
		},
	})
}

func TestAccIBMAccountSettingsTemplateVersionBasic(t *testing.T) {
	var conf iamidentityv1.AccountSettingsTemplateResponse
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMAccountSettingsTemplateDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMAccountSettingsTemplateVersionConfigBasic(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMAccountSettingsTemplateExists("ibm_iam_account_settings_template.account_settings_template_instance", conf),
					testAccCheckIBMAccountSettingsTemplateExists("ibm_iam_account_settings_template.account_settings_template_version", conf),
					resource.TestCheckResourceAttr("ibm_iam_account_settings_template.account_settings_template_instance", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_account_settings_template.account_settings_template_version", "name", name),
				),
			},
		},
	})
}

func TestAccIBMAccountSettingsTemplateAllArgs(t *testing.T) {
	var conf iamidentityv1.AccountSettingsTemplateResponse
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	description := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))
	mfa := "LEVEL1"
	descriptionUpdate := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))
	mfaUpdate := "LEVEL3"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMAccountSettingsTemplateDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMAccountSettingsTemplateConfig(name, description, "false", mfa),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMAccountSettingsTemplateExists("ibm_iam_account_settings_template.account_settings_template_instance", conf),
					resource.TestCheckResourceAttr("ibm_iam_account_settings_template.account_settings_template_instance", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_account_settings_template.account_settings_template_instance", "description", description),
					resource.TestCheckResourceAttr("ibm_iam_account_settings_template.account_settings_template_instance", "account_settings.0.mfa", mfa),
				),
			},
			{
				Config: testAccCheckIBMAccountSettingsTemplateConfig(name, descriptionUpdate, "false", mfaUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_iam_account_settings_template.account_settings_template_instance", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_account_settings_template.account_settings_template_instance", "description", descriptionUpdate),
					resource.TestCheckResourceAttr("ibm_iam_account_settings_template.account_settings_template_instance", "account_settings.0.mfa", mfaUpdate),
				),
			},
			{
				Config: testAccCheckIBMAccountSettingsTemplateConfig(name, descriptionUpdate, "true", mfaUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_iam_account_settings_template.account_settings_template_instance", "committed", "true"),
				),
			},
			{
				ResourceName:      "ibm_iam_account_settings_template.account_settings_template_instance",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIBMAccountSettingsTemplateConfigBasic(name string, description string) string {
	return fmt.Sprintf(`
		resource "ibm_iam_account_settings_template" "account_settings_template_instance" {
			name = "%s"
			description = "%s"
			account_settings {
				mfa = "LEVEL3"
			}
		}
	`, name, description)
}

func testAccCheckIBMAccountSettingsTemplateVersionConfigBasic(name string) string {
	return fmt.Sprintf(`
		resource "ibm_iam_account_settings_template" "account_settings_template_instance" {
			name = "%s"
			account_settings {
			
			}
		}
		resource "ibm_iam_account_settings_template" "account_settings_template_version" {
			template_id = ibm_iam_account_settings_template.account_settings_template_instance.id
			name = "%s"
			description = "Description for version 2"
			account_settings {
			
			}
		}
	`, name, name)
}

func testAccCheckIBMAccountSettingsTemplateConfig(name string, description string, committed string, mfa string) string {
	return fmt.Sprintf(`
		resource "ibm_iam_account_settings_template" "account_settings_template_instance" {
			name = "%s"
			description = "%s"
			committed = "%s"
			account_settings {
				restrict_create_service_id = "RESTRICTED"
				restrict_create_platform_apikey = "RESTRICTED"
				allowed_ip_addresses = ""
				mfa = "%s"
				user_mfa {
					iam_id = "IBMid-123456789"
					mfa = "NONE"
				}
				session_expiration_in_seconds = "1800"
				session_invalidation_in_seconds = "1800"
				max_sessions_per_identity = "5"
				system_access_token_expiration_in_seconds = "NOT_SET"
				system_refresh_token_expiration_in_seconds = "NOT_SET"
			}
		}
	`, name, description, committed, mfa)
}

func testAccCheckIBMAccountSettingsTemplateExists(n string, obj iamidentityv1.AccountSettingsTemplateResponse) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		iamIdentityClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).IAMIdentityV1API()
		if err != nil {
			return err
		}

		getAccountSettingsTemplateVersionOptions := &iamidentityv1.GetAccountSettingsTemplateVersionOptions{}

		id, version := parseAccountSettingsResourceId(rs.Primary.ID)
		getAccountSettingsTemplateVersionOptions.SetTemplateID(id)
		getAccountSettingsTemplateVersionOptions.SetVersion(version)

		accountSettingsTemplateResponse, _, err := iamIdentityClient.GetAccountSettingsTemplateVersion(getAccountSettingsTemplateVersionOptions)
		if err != nil {
			return err
		}

		obj = *accountSettingsTemplateResponse
		return nil
	}
}

func testAccCheckIBMAccountSettingsTemplateDestroy(s *terraform.State) error {
	iamIdentityClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).IAMIdentityV1API()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_iam_account_settings_template" {
			continue
		}

		getAccountSettingsTemplateVersionOptions := &iamidentityv1.GetAccountSettingsTemplateVersionOptions{}

		id, version := parseAccountSettingsResourceId(rs.Primary.ID)
		getAccountSettingsTemplateVersionOptions.SetTemplateID(id)
		getAccountSettingsTemplateVersionOptions.SetVersion(version)

		// Try to find the key
		_, response, err := iamIdentityClient.GetAccountSettingsTemplateVersion(getAccountSettingsTemplateVersionOptions)

		if err == nil {
			return fmt.Errorf("account_settings_template still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for account_settings_template (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}

func parseAccountSettingsResourceId(ID string) (templateId, templateVersion string) {
	resourceIdParts := strings.Split(ID, "/")

	if len(resourceIdParts) == 1 {
		return resourceIdParts[0], ""
	}

	return resourceIdParts[0], resourceIdParts[1]
}
