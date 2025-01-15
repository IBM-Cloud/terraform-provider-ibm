// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package schematics_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/IBM/schematics-go-sdk/schematicsv1"
)

func TestAccIBMSchematicsActionBasic(t *testing.T) {
	var conf schematicsv1.Action
	actionName := fmt.Sprintf("acc-test-schematics-actions_%s", acctest.RandString(10))
	description := fmt.Sprintf("acc-test-schematics-action_description_%s", acctest.RandString(10))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMSchematicsActionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMSchematicsActionConfigBasic(actionName, description),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMSchematicsActionExists("ibm_schematics_action.schematics_action", conf),
					resource.TestCheckResourceAttrSet("ibm_schematics_action.schematics_action", "id"),
					resource.TestCheckResourceAttr("ibm_schematics_action.schematics_action", "name", actionName),
					resource.TestCheckResourceAttr("ibm_schematics_action.schematics_action", "description", description),
					resource.TestCheckResourceAttrSet("ibm_schematics_action.schematics_action", "location"),
					resource.TestCheckResourceAttrSet("ibm_schematics_action.schematics_action", "resource_group"),
					resource.TestCheckResourceAttrSet("ibm_schematics_action.schematics_action", "user_state.#"),
					resource.TestCheckResourceAttrSet("ibm_schematics_action.schematics_action", "created_at"),
					resource.TestCheckResourceAttrSet("ibm_schematics_action.schematics_action", "created_by"),
					resource.TestCheckResourceAttrSet("ibm_schematics_action.schematics_action", "updated_at"),
					resource.TestCheckResourceAttrSet("ibm_schematics_action.schematics_action", "sys_lock.#"),
					resource.TestCheckResourceAttrSet("ibm_schematics_action.schematics_action", "state.#"),
				),
			},
		},
	})
}

func testAccCheckIBMSchematicsActionConfigBasic(actionName string, description string) string {

	return fmt.Sprintf(`

		resource "ibm_schematics_action" "schematics_action" {
			name = "%s"
			description = "%s"
			location = "us-east"
			resource_group = "Default"
		}
	`, actionName, description)
}

func testAccCheckIBMSchematicsActionExists(n string, obj schematicsv1.Action) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		schematicsClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).SchematicsV1()
		if err != nil {
			return err
		}

		getActionOptions := &schematicsv1.GetActionOptions{}

		getActionOptions.SetActionID(rs.Primary.ID)
		//getActionOptions.SetProfile("detailed")

		action, _, err := schematicsClient.GetAction(getActionOptions)
		if err != nil {
			return err
		}

		obj = *action
		return nil
	}
}

func testAccCheckIBMSchematicsActionDestroy(s *terraform.State) error {
	schematicsClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).SchematicsV1()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_schematics_action" {
			continue
		}

		getActionOptions := &schematicsv1.GetActionOptions{}

		getActionOptions.SetActionID(rs.Primary.ID)

		// Try to find the key
		_, response, err := schematicsClient.GetAction(getActionOptions)

		if err == nil {
			return fmt.Errorf("schematics_action still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("[ERROR] Error checking for schematics_action (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
