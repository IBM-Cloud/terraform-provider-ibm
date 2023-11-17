// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package iamidentity_test

import (
	"fmt"
	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM/platform-services-go-sdk/iamidentityv1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"testing"
)

func TestAccIBMAccountSettingsTemplateAssignmentDataSourceBasic(t *testing.T) {
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acc.TestAccPreCheckEnterprise(t)
			acc.TestAccPreCheckAssignmentTargetAccount(t)
		},
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMAccountSettingsTemplateAssignmentDataSourceDestroy,
		Steps: []resource.TestStep{
			{
				Config:             testAccCheckIBMAccountSettingsTemplateAssignmentDataSourceConfigBasic(name),
				ExpectNonEmptyPlan: true,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_iam_account_settings_template_assignment.account_settings_template_assignment_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_account_settings_template_assignment.account_settings_template_assignment_instance", "account_id"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_account_settings_template_assignment.account_settings_template_assignment_instance", "template_id"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_account_settings_template_assignment.account_settings_template_assignment_instance", "template_version"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_account_settings_template_assignment.account_settings_template_assignment_instance", "target_type"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_account_settings_template_assignment.account_settings_template_assignment_instance", "target"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_account_settings_template_assignment.account_settings_template_assignment_instance", "resources.#"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_account_settings_template_assignment.account_settings_template_assignment_instance", "status"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_account_settings_template_assignment.account_settings_template_assignment_instance", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_account_settings_template_assignment.account_settings_template_assignment_instance", "created_by_id"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_account_settings_template_assignment.account_settings_template_assignment_instance", "last_modified_at"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_account_settings_template_assignment.account_settings_template_assignment_instance", "last_modified_by_id"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_account_settings_template_assignment.account_settings_template_assignment_instance", "entity_tag"),
				),
			},
		},
	})
}

func testAccCheckIBMAccountSettingsTemplateAssignmentDataSourceConfigBasic(name string) string {
	return fmt.Sprintf(`
		resource "ibm_iam_account_settings_template" "account_settings_template" {
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

		 	timeouts {
				create = "5m"
			}
		}

		data "ibm_iam_account_settings_template_assignment" "account_settings_template_assignment_instance" {
			assignment_id = ibm_iam_account_settings_template_assignment.account_settings_template_assignment_instance.id
		}
	`, name, acc.IamIdentityAssignmentTargetAccountId)
}

func testAccCheckIBMAccountSettingsTemplateAssignmentDataSourceDestroy(s *terraform.State) error {
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
