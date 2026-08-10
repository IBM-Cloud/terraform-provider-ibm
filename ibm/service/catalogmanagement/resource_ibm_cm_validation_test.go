// Copyright IBM Corp. 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package catalogmanagement_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM/platform-services-go-sdk/catalogmanagementv1"
)

func TestAccIBMCmValidationSimpleArgs(t *testing.T) {
	var conf catalogmanagementv1.Version
	// this needs to be a real VL for a version that can be validated (i.e. the version is not published/ready).  A real validation
	// will be done, so there will be a workspace created.
	versionLocator := "f418eaf7-602a-472d-9e9f-ab75a7c193ab.46f1f0e2-3205-41d8-99e1-14dad4e1dfb8"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCmValidationDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCmValidationSimpleConfig(versionLocator),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCmValidationExists("ibm_cm_validation.cm_validation", conf),
					resource.TestCheckResourceAttr("ibm_cm_validation.cm_validation", "version_locator", versionLocator),
				),
			},
			{ // step 2 will wait for step 1 ^^ to finish.  'target' is not set until step 1 finishes.
				Config: testAccCheckIBMCmValidationSimpleConfig(versionLocator),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_cm_validation.cm_validation", "state", "valid"),         // state = valid
					resource.TestCheckResourceAttrSet("ibm_cm_validation.cm_validation", "target.workspace_id"), // Check workspace_id key exists and has a value
				),
			},
		},
	})
}

func testAccCheckIBMCmValidationSimpleConfig(versionLocator string) string {
	return fmt.Sprintf(`
		resource "ibm_cm_validation" "cm_validation" {
			version_locator = "%s"
			revalidate_if_validated = true
			mark_version_consumable = false

			override_values = {
				name = "My TF"
			}

			environment_variables {
				name = "test"
				value = "test"
				secure = true
			}

			schematics {
				name = "acceptance-test-workspace"
				description = "workspace description"
				region = "eu-de"
			}
		}
	`, versionLocator)
}

func testAccCheckIBMCmValidationExists(n string, obj catalogmanagementv1.Version) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		return nil
	}
}

func testAccCheckIBMCmValidationDestroy(s *terraform.State) error {
	return nil
}
