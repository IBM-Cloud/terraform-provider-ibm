// Copyright IBM Corp. 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package catalogmanagement_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/IBM/platform-services-go-sdk/catalogmanagementv1"
	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"
)

func TestAccIBMCmValidationSimpleArgs(t *testing.T) {
	var conf catalogmanagementv1.Version
	versionLocator := "dba7e7dd-2bd7-4fcd-a846-4c370eab2672.98ba725b-86fa-4c6a-8430-70f38ec988da"

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
				name = "workspace name"
				description = "workspace description"
				region = "us-south"
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
