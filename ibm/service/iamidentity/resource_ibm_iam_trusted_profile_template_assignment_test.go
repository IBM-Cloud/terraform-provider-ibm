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

func TestAccIBMTrustedProfileTemplateAssignmentBasic(t *testing.T) {
	var conf iamidentityv1.TemplateAssignmentResponse
	name := fmt.Sprintf("tf_tp_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acc.TestAccPreCheck(t)
			acc.TestAccPreCheckAssignmentTargetAccount(t)
		},
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMTrustedProfileTemplateAssignmentResourceDestroy,
		Steps: []resource.TestStep{
			{

				Config: testAccCheckIBMTrustedProfileTemplateAssignmentConfigBasic(name, 1),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMTrustedProfileTemplateAssignmentExists("ibm_iam_trusted_profile_template_assignment.trusted_profile_template_assignment_instance", conf),
					resource.TestCheckResourceAttrSet("ibm_iam_trusted_profile_template_assignment.trusted_profile_template_assignment_instance", "id"),
					resource.TestCheckResourceAttrSet("ibm_iam_trusted_profile_template_assignment.trusted_profile_template_assignment_instance", "account_id"),
					resource.TestCheckResourceAttrSet("ibm_iam_trusted_profile_template_assignment.trusted_profile_template_assignment_instance", "template_id"),
					resource.TestCheckResourceAttrSet("ibm_iam_trusted_profile_template_assignment.trusted_profile_template_assignment_instance", "template_version"),
					resource.TestCheckResourceAttrSet("ibm_iam_trusted_profile_template_assignment.trusted_profile_template_assignment_instance", "target_type"),
					resource.TestCheckResourceAttrSet("ibm_iam_trusted_profile_template_assignment.trusted_profile_template_assignment_instance", "target"),
					resource.TestCheckResourceAttrSet("ibm_iam_trusted_profile_template_assignment.trusted_profile_template_assignment_instance", "resources.#"),
					resource.TestCheckResourceAttrSet("ibm_iam_trusted_profile_template_assignment.trusted_profile_template_assignment_instance", "status"),
					resource.TestCheckResourceAttrSet("ibm_iam_trusted_profile_template_assignment.trusted_profile_template_assignment_instance", "created_at"),
					resource.TestCheckResourceAttrSet("ibm_iam_trusted_profile_template_assignment.trusted_profile_template_assignment_instance", "created_by_id"),
					resource.TestCheckResourceAttrSet("ibm_iam_trusted_profile_template_assignment.trusted_profile_template_assignment_instance", "last_modified_at"),
					resource.TestCheckResourceAttrSet("ibm_iam_trusted_profile_template_assignment.trusted_profile_template_assignment_instance", "last_modified_by_id"),
					resource.TestCheckResourceAttrSet("ibm_iam_trusted_profile_template_assignment.trusted_profile_template_assignment_instance", "entity_tag"),
				),
			},
			{
				ExpectError: regexp.MustCompile("Template version '3' is not found."),
				Config:      testAccCheckIBMTrustedProfileTemplateAssignmentConfigBasic(name, 3),
			},
			{
				Config: testAccCheckIBMTrustedProfileTemplateAssignmentConfigBasic(name, 2),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_iam_trusted_profile_template_assignment.trusted_profile_template_assignment_instance", "template_version", "2"),
				),
			},
		},
	})
}

func testAccCheckIBMTrustedProfileTemplateAssignmentConfigBasic(name string, version int) string {
	return fmt.Sprintf(`
		resource "ibm_iam_trusted_profile_template" "trusted_profile_template" {
			name = "%s"
			profile {
				name = "%s"
			}
			committed = "true"
		}

		resource "ibm_iam_trusted_profile_template" "trusted_profile_template_v2" {
			template_id = ibm_iam_trusted_profile_template.trusted_profile_template.id
			name = ibm_iam_trusted_profile_template.trusted_profile_template.name
			description = "v2 Description"
			profile {
				name = ibm_iam_trusted_profile_template.trusted_profile_template.name
			}
			committed = "true"
		}

		resource "ibm_iam_trusted_profile_template_assignment" "trusted_profile_template_assignment_instance" {
			template_id = split("/", ibm_iam_trusted_profile_template.trusted_profile_template.id)[0]
			template_version = %d
		  	target_type = "Account"
		  	target = "%s"
			depends_on = [
				ibm_iam_trusted_profile_template.trusted_profile_template,
				ibm_iam_trusted_profile_template.trusted_profile_template_v2
			]

		 	timeouts {
				create = "5m"
				update = "5m"
			}
		}
	`, name, name, version, acc.IamIdentityAssignmentTargetAccountId)
}

func testAccCheckIBMTrustedProfileTemplateAssignmentExists(n string, obj iamidentityv1.TemplateAssignmentResponse) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("not found: %s", n)
		}

		iamIdentityClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).IAMIdentityV1API()
		if err != nil {
			return err
		}

		getTrustedProfileAssignmentOptions := &iamidentityv1.GetTrustedProfileAssignmentOptions{}

		getTrustedProfileAssignmentOptions.SetAssignmentID(rs.Primary.ID)

		templateAssignmentResponse, _, err := iamIdentityClient.GetTrustedProfileAssignment(getTrustedProfileAssignmentOptions)
		if err != nil {
			return err
		}

		obj = *templateAssignmentResponse
		return nil
	}
}

func testAccCheckIBMTrustedProfileTemplateAssignmentResourceDestroy(s *terraform.State) error {
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
