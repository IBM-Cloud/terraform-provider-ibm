// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package iamidentity_test

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM/platform-services-go-sdk/iamidentityv1"
)

func TestAccIBMAccountSettingsTemplateAssignmentBasic(t *testing.T) {
	var conf iamidentityv1.TemplateAssignmentResponse
	enterpriseAccountId := acc.IamIdentityEnterpriseAccountId
	targetId := acc.IamIdentityAssignmentTargetAccountId
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:                  func() { acc.TestAccPreCheckIamIdentityEnterpriseTemplates(t) },
		Providers:                 acc.TestAccProviders,
		CheckDestroy:              testAccCheckIBMAccountSettingsTemplateAssignmentResourceDestroy,
		PreventPostDestroyRefresh: true,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMAccountSettingsTemplateAssignmentConfigBasic(enterpriseAccountId, targetId, name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMAccountSettingsTemplateAssignmentExists("ibm_iam_account_settings_template_assignment.account_settings_template_assignment_instance", conf),
					resource.TestCheckResourceAttrSet("ibm_iam_account_settings_template_assignment.account_settings_template_assignment_instance", "id"),
					resource.TestCheckResourceAttrSet("ibm_iam_account_settings_template_assignment.account_settings_template_assignment_instance", "account_id"),
					resource.TestCheckResourceAttrSet("ibm_iam_account_settings_template_assignment.account_settings_template_assignment_instance", "template_id"),
					resource.TestCheckResourceAttrSet("ibm_iam_account_settings_template_assignment.account_settings_template_assignment_instance", "template_version"),
					resource.TestCheckResourceAttrSet("ibm_iam_account_settings_template_assignment.account_settings_template_assignment_instance", "target_type"),
					resource.TestCheckResourceAttrSet("ibm_iam_account_settings_template_assignment.account_settings_template_assignment_instance", "target"),
					resource.TestCheckResourceAttrSet("ibm_iam_account_settings_template_assignment.account_settings_template_assignment_instance", "resources.#"),
					resource.TestCheckResourceAttrSet("ibm_iam_account_settings_template_assignment.account_settings_template_assignment_instance", "status"),
					resource.TestCheckResourceAttrSet("ibm_iam_account_settings_template_assignment.account_settings_template_assignment_instance", "created_at"),
					resource.TestCheckResourceAttrSet("ibm_iam_account_settings_template_assignment.account_settings_template_assignment_instance", "created_by_id"),
					resource.TestCheckResourceAttrSet("ibm_iam_account_settings_template_assignment.account_settings_template_assignment_instance", "last_modified_at"),
					resource.TestCheckResourceAttrSet("ibm_iam_account_settings_template_assignment.account_settings_template_assignment_instance", "last_modified_by_id"),
					resource.TestCheckResourceAttrSet("ibm_iam_account_settings_template_assignment.account_settings_template_assignment_instance", "entity_tag"),
				),
			},
			{
				ExpectError: regexp.MustCompile("Template version '2' is not found."),
				Config:      testAccCheckIBMAccountSettingsTemplateAssignmentConfigBasicUpdate(enterpriseAccountId, targetId, name),
			},
		},
	})
}

func testAccCheckIBMAccountSettingsTemplateAssignmentConfigBasic(enterpriseAccountId string, targetId string, name string) string {
	return fmt.Sprintf(`
		resource "ibm_iam_account_settings_template" "account_settings_template" {
			account_id = "%s"
			name = "%s"
			account_settings {
				mfa = "LEVEL3"
			}
			committed = true
		}

		resource "ibm_iam_account_settings_template_assignment" "account_settings_template_assignment_instance" {
			template_id = split("/", ibm_iam_account_settings_template.account_settings_template.id)[0]
			template_version = ibm_iam_account_settings_template.account_settings_template.version
		  	target_type = "Account"
		  	target = "%s"

			depends_on = [
				ibm_iam_account_settings_template.account_settings_template
			]

		 	timeouts {
				create = "5m"
			}
		}
	`, enterpriseAccountId, name, targetId)
}

func testAccCheckIBMAccountSettingsTemplateAssignmentConfigBasicUpdate(enterpriseAccountId string, targetId string, name string) string {
	return fmt.Sprintf(`
		resource "ibm_iam_account_settings_template" "account_settings_template" {
			account_id = "%s"	
			name = "%s"
			account_settings {
				mfa = "LEVEL3"
			}
			committed = true
		}

		resource "ibm_iam_account_settings_template_assignment" "account_settings_template_assignment_instance" {
			template_id = split("/", ibm_iam_account_settings_template.account_settings_template.id)[0]
			template_version = 2
		  	target_type = "Account"
		  	target = "%s"

			depends_on = [
				ibm_iam_account_settings_template.account_settings_template
			]

		 	timeouts {
				update = "5m"
			}
		}
	`, enterpriseAccountId, name, targetId)
}

func testAccCheckIBMAccountSettingsTemplateAssignmentExists(n string, obj iamidentityv1.TemplateAssignmentResponse) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("not found: %s", n)
		}

		iamIdentityClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).IAMIdentityV1API()
		if err != nil {
			return err
		}

		getAccountSettingsAssignmentOptions := &iamidentityv1.GetAccountSettingsAssignmentOptions{}

		getAccountSettingsAssignmentOptions.SetAssignmentID(rs.Primary.ID)

		templateAssignmentResponse, _, err := iamIdentityClient.GetAccountSettingsAssignment(getAccountSettingsAssignmentOptions)
		if err != nil {
			return err
		}

		obj = *templateAssignmentResponse
		return nil
	}
}

func testAccCheckIBMAccountSettingsTemplateAssignmentResourceDestroy(s *terraform.State) error {
	iamIdentityClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).IAMIdentityV1API()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_iam_account_settings_template_assignment" {
			continue
		}

		getAccountSettingsAssignmentOptions := &iamidentityv1.GetAccountSettingsAssignmentOptions{}

		getAccountSettingsAssignmentOptions.SetAssignmentID(rs.Primary.ID)

		// Try to find the key
		_, response, err := iamIdentityClient.GetAccountSettingsAssignment(getAccountSettingsAssignmentOptions)

		if err == nil {
			return fmt.Errorf("account_settings_template_assignment still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("error checking for account_settings_template_assignment_instance (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
