// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package iamidentity_test

import (
	"fmt"
	"testing"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM/platform-services-go-sdk/iamidentityv1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIBMTrustedProfileTemplateAssignmentDataSourceBasic(t *testing.T) {
	enterpriseAccountId := acc.IamIdentityEnterpriseAccountId
	targetId := acc.IamIdentityAssignmentTargetAccountId
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:                  func() { acc.TestAccPreCheckIamIdentityEnterpriseTemplates(t) },
		Providers:                 acc.TestAccProviders,
		CheckDestroy:              testAccCheckIBMTrustedProfileTemplateAssignmentDataSourceDestroy,
		PreventPostDestroyRefresh: true,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMTrustedProfileTemplateAssignmentDataSourceConfigBasic(enterpriseAccountId, targetId, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profile_template_assignment.trusted_profile_template_assignment_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profile_template_assignment.trusted_profile_template_assignment_instance", "account_id"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profile_template_assignment.trusted_profile_template_assignment_instance", "template_id"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profile_template_assignment.trusted_profile_template_assignment_instance", "template_version"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profile_template_assignment.trusted_profile_template_assignment_instance", "target_type"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profile_template_assignment.trusted_profile_template_assignment_instance", "target"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profile_template_assignment.trusted_profile_template_assignment_instance", "resources.#"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profile_template_assignment.trusted_profile_template_assignment_instance", "status"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profile_template_assignment.trusted_profile_template_assignment_instance", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profile_template_assignment.trusted_profile_template_assignment_instance", "created_by_id"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profile_template_assignment.trusted_profile_template_assignment_instance", "last_modified_at"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profile_template_assignment.trusted_profile_template_assignment_instance", "last_modified_by_id"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profile_template_assignment.trusted_profile_template_assignment_instance", "entity_tag"),
				),
			},
		},
	})
}

func testAccCheckIBMTrustedProfileTemplateAssignmentDataSourceConfigBasic(enterpriseAccountId string, targetId string, name string) string {
	return fmt.Sprintf(`
		resource "ibm_iam_trusted_profile_template" "trusted_profile_template" {
			account_id = "%s"
			name = "%s"
			profile {
				name = "%s"
			}
			committed = "true"
		}

		resource "ibm_iam_trusted_profile_template_assignment" "trusted_profile_template_assignment_instance" {
			template_id = split("/", ibm_iam_trusted_profile_template.trusted_profile_template.id)[0]
			template_version = ibm_iam_trusted_profile_template.trusted_profile_template.version
		  	target_type = "Account"
		  	target = "%s"
			depends_on = [
				ibm_iam_trusted_profile_template.trusted_profile_template
			]

		 	timeouts {
				create = "5m"
			}
		}

		data "ibm_iam_trusted_profile_template_assignment" "trusted_profile_template_assignment_instance" {
			assignment_id = ibm_iam_trusted_profile_template_assignment.trusted_profile_template_assignment_instance.id
		}
	`, enterpriseAccountId, name, name, targetId)
}

func testAccCheckIBMTrustedProfileTemplateAssignmentDataSourceDestroy(s *terraform.State) error {
	iamIdentityClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).IAMIdentityV1API()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_iam_trusted_profile_template_assignment" {
			continue
		}

		getTrustedProfileAssignmentOptions := &iamidentityv1.GetTrustedProfileAssignmentOptions{}

		getTrustedProfileAssignmentOptions.SetAssignmentID(rs.Primary.ID)

		// Try to find the key
		_, response, err := iamIdentityClient.GetTrustedProfileAssignment(getTrustedProfileAssignmentOptions)

		if err == nil {
			return fmt.Errorf("trusted_profile_template_assignment still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("error checking for trusted_profile_template_assignment (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
