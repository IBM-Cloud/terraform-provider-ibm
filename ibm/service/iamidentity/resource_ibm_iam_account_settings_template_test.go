// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package iamidentity_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM/platform-services-go-sdk/iamidentityv1"
)

func TestAccIBMAccountSettingsTemplateBasic(t *testing.T) {
	var conf iamidentityv1.AccountSettingsTemplateResponse
	enterpriseAccountId := acc.IamIdentityEnterpriseAccountId
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	description := fmt.Sprintf("tf_desc_%d", acctest.RandIntRange(10, 100))
	descriptionUpdate := fmt.Sprintf("tf_desc_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheckIamIdentityEnterpriseTemplates(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMAccountSettingsTemplateDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMAccountSettingsTemplateConfigBasic(enterpriseAccountId, name, description),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMAccountSettingsTemplateExists("ibm_iam_account_settings_template.account_settings_template_instance", conf),
					resource.TestCheckResourceAttr("ibm_iam_account_settings_template.account_settings_template_instance", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_account_settings_template.account_settings_template_instance", "description", description),
				),
			},
			{
				Config: testAccCheckIBMAccountSettingsTemplateConfigBasic(enterpriseAccountId, name, descriptionUpdate),
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
	enterpriseAccountId := acc.IamIdentityEnterpriseAccountId
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheckIamIdentityEnterpriseTemplates(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMAccountSettingsTemplateDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMAccountSettingsTemplateVersionConfigBasic(enterpriseAccountId, name),
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
	enterpriseAccountId := acc.IamIdentityEnterpriseAccountId
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	description := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))
	mfa := "LEVEL1"
	descriptionUpdate := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))
	mfaUpdate := "LEVEL3"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheckIamIdentityEnterpriseTemplates(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMAccountSettingsTemplateDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMAccountSettingsTemplateConfig(enterpriseAccountId, name, description, "false", mfa),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMAccountSettingsTemplateExists("ibm_iam_account_settings_template.account_settings_template_instance", conf),
					resource.TestCheckResourceAttr("ibm_iam_account_settings_template.account_settings_template_instance", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_account_settings_template.account_settings_template_instance", "description", description),
					resource.TestCheckResourceAttr("ibm_iam_account_settings_template.account_settings_template_instance", "account_settings.0.mfa", mfa),
				),
			},
			{
				Config: testAccCheckIBMAccountSettingsTemplateConfig(enterpriseAccountId, name, descriptionUpdate, "false", mfaUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_iam_account_settings_template.account_settings_template_instance", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_account_settings_template.account_settings_template_instance", "description", descriptionUpdate),
					resource.TestCheckResourceAttr("ibm_iam_account_settings_template.account_settings_template_instance", "account_settings.0.mfa", mfaUpdate),
				),
			},
			{
				Config: testAccCheckIBMAccountSettingsTemplateConfig(enterpriseAccountId, name, descriptionUpdate, "true", mfaUpdate),
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

func testAccCheckIBMAccountSettingsTemplateConfigBasic(enterpriseAccountId string, name string, description string) string {
	return fmt.Sprintf(`
		resource "ibm_iam_account_settings_template" "account_settings_template_instance" {
			account_id = "%s"
			name = "%s"
			description = "%s"
			account_settings {
				mfa = "LEVEL3"
			}
		}
	`, enterpriseAccountId, name, description)
}

func testAccCheckIBMAccountSettingsTemplateVersionConfigBasic(enterpriseAccountId string, name string) string {
	return fmt.Sprintf(`
		resource "ibm_iam_account_settings_template" "account_settings_template_instance" {
			account_id = "%s"
			name = "%s"
			account_settings {
			
			}
		}
		resource "ibm_iam_account_settings_template" "account_settings_template_version" {
			template_id = ibm_iam_account_settings_template.account_settings_template_instance.id
			account_id = ibm_iam_account_settings_template.account_settings_template_instance.account_id
			name = "%s"
			description = "Description for version 2"
			account_settings {
			
			}
		}
	`, enterpriseAccountId, name, name)
}

func testAccCheckIBMAccountSettingsTemplateConfig(enterpriseAccountId string, name string, description string, committed string, mfa string) string {
	return fmt.Sprintf(`
		resource "ibm_iam_account_settings_template" "account_settings_template_instance" {
			account_id = "%s"
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
				restrict_user_list_visibility = "RESTRICTED"
				restrict_user_domains {
					account_sufficient = true
					restrictions {
						realm_id = "IBMid"
						invitation_email_allow_patterns = ["*.*@company.com"]
						restrict_invitation = true
					}
				}
			}
		}
	`, enterpriseAccountId, name, description, committed, mfa)
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
