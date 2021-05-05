// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMSchematicsActionDataSourceBasic(t *testing.T) {

	actionName := fmt.Sprintf("acc-test-schematics-actions_%s", acctest.RandString(10))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMSchematicsActionDataSourceConfigBasic(actionName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "description"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "location"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "resource_group"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "user_state.#"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "created_by"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "updated_at"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "sys_lock.#"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "state.#"),
				),
			},
		},
	})
}

func testAccCheckIBMSchematicsActionDataSourceConfigBasic(actionName string) string {
	return fmt.Sprintf(`
		resource "ibm_schematics_action" "schematics_action" {
			name = "%s"
			description = "acc-test-schematics-actions"
			location = "us-east"
			resource_group = "default"
		}

		data "ibm_schematics_action" "schematics_action" {
			action_id = ibm_schematics_action.schematics_action.id
		}
	`, actionName)
}
