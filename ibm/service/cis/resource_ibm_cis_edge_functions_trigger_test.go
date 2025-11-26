// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package cis_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccIBMCisEdgeFunctionsTrigger_Basic(t *testing.T) {
	var record string
	testName := "test"
	resourceName := fmt.Sprintf("ibm_cis_edge_functions_trigger.%s", testName)
	actionName := "sample_script"
	pattern1 := fmt.Sprintf("example.%s/*", acc.CisDomainStatic)
	pattern2 := fmt.Sprintf("example1.%s/*", acc.CisDomainStatic)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheckCis(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCisEdgeFunctionsActionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCisEdgeFunctionsTriggerBasic(testName, pattern1, actionName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMCisEdgeFunctionsTriggerExists(resourceName, &record),
					resource.TestCheckResourceAttr(
						resourceName, "action_name", actionName),
					resource.TestCheckResourceAttr(
						resourceName, "pattern_url", pattern1),
				),
			},
			{
				Config: testAccCheckIBMCisEdgeFunctionsTriggerUpdate(testName, pattern2, actionName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMCisEdgeFunctionsTriggerExists(resourceName, &record),
					resource.TestCheckResourceAttr(
						resourceName, "pattern_url", pattern2),
				),
			},
		},
	})
}

func TestAccIBMCisEdgeFunctionsTrigger_import(t *testing.T) {
	var record string

	name := "ibm_cis_edge_functions_trigger.test"
	actionName := "sample_script"
	pattern := fmt.Sprintf("example.%s/*", acc.CisDomainStatic)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCisEdgeFunctionsTriggerDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCisEdgeFunctionsTriggerBasic("test", pattern, actionName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMCisEdgeFunctionsTriggerExists(name, &record),
					resource.TestCheckResourceAttr(name, "action_name", actionName),
				),
			},
			{
				Config: testAccCheckIBMCisEdgeFunctionsTriggerUpdate("test", pattern, actionName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMCisEdgeFunctionsTriggerExists(name, &record),
					resource.TestCheckResourceAttr(
						name, "pattern_url", pattern),
				),
			},
			{
				ResourceName:      name,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccIBMCisFunctionsTrigger_CreateAfterManualDestroy(t *testing.T) {
	var scriptOne, scriptTwo string
	name := "ibm_cis_edge_functions_trigger.test"
	scriptOne = "script_one"
	scriptTwo = "script_two"
	pattern := fmt.Sprintf("example.%s/*", acc.CisDomainStatic)
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCisEdgeFunctionsTriggerDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCisEdgeFunctionsTriggerBasic("test", pattern, scriptOne),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMCisEdgeFunctionsTriggerExists(name, &scriptOne),
					testAccCheckIBMCisEdgeFunctionsTriggerDelete(&scriptOne),
				),
				ExpectNonEmptyPlan: true,
			},
			{
				Config: testAccCheckIBMCisEdgeFunctionsTriggerBasic("test", pattern, scriptTwo),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMCisEdgeFunctionsTriggerExists(name, &scriptTwo),
					func(state *terraform.State) error {
						if scriptOne == scriptTwo {
							return fmt.Errorf("Action script unchanged even after we thought we deleted it ( %s )",
								scriptTwo)
						}
						return nil
					},
				),
			},
		},
	})
}

func testAccCheckIBMCisEdgeFunctionsTriggerDelete(tfActionID *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		cisClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).CisEdgeFunctionClientSession()
		if err != nil {
			return fmt.Errorf("[ERROR] Error in creating CIS object")
		}

		actionName, zoneID, cisID, _ := flex.ConvertTfToCisThreeVar(*tfActionID)
		cisClient.Crn = core.StringPtr(cisID)
		cisClient.ZoneIdentifier = core.StringPtr(zoneID)
		opt := cisClient.NewDeleteEdgeFunctionsTriggerOptions(actionName)
		_, response, err := cisClient.DeleteEdgeFunctionsTrigger(opt)
		if err != nil {
			return fmt.Errorf("Edge function action script deletion failed: %v", response)
		}
		return nil
	}
}

func testAccCheckIBMCisEdgeFunctionsTriggerDestroy(s *terraform.State) error {
	cisClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).CisEdgeFunctionClientSession()
	if err != nil {
		return fmt.Errorf("[ERROR] Error in creating CIS object")
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_cis_edge_functions_trigger" {
			continue
		}

		triggerID, zoneID, cisID, _ := flex.ConvertTfToCisThreeVar(rs.Primary.ID)
		cisClient.Crn = core.StringPtr(cisID)
		cisClient.ZoneIdentifier = core.StringPtr(zoneID)
		opt := cisClient.NewGetEdgeFunctionsTriggerOptions(triggerID)
		_, response, err := cisClient.GetEdgeFunctionsTrigger(opt)
		if err == nil {
			return fmt.Errorf("Edge function action script trigger still exists: %v", response)
		}
	}

	return nil
}

func testAccCheckIBMCisEdgeFunctionsTriggerExists(n string, tfRecordID *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("[ERROR] Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("[ERROR] No Record ID is set")
		}

		// tfRecord := *tfRecordID
		cisClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).CisEdgeFunctionClientSession()
		if err != nil {
			return fmt.Errorf("[ERROR] Error in creating CIS object")
		}
		triggerID, zoneID, cisID, _ := flex.ConvertTfToCisThreeVar(rs.Primary.ID)
		cisClient.Crn = core.StringPtr(cisID)
		cisClient.ZoneIdentifier = core.StringPtr(zoneID)
		opt := cisClient.NewGetEdgeFunctionsTriggerOptions(triggerID)
		result, resp, err := cisClient.GetEdgeFunctionsTrigger(opt)
		if err != nil {
			return fmt.Errorf("[ERROR] Error: %v", resp)
		}
		if *result.Result.ID != triggerID {
			return fmt.Errorf("Trigger ID is not found")
		}

		tfRecord := flex.ConvertCisToTfThreeVar(*result.Result.ID, zoneID, cisID)
		*tfRecordID = tfRecord
		return nil
	}
}

func testAccCheckIBMCisEdgeFunctionsTriggerBasic(testName, pattern, actionName string) string {
	return testAccCheckIBMCisEdgeFunctionsActionBasic(testName, actionName) + fmt.Sprintf(`
	resource "ibm_cis_edge_functions_trigger" "%[1]s" {
		cis_id      = ibm_cis_edge_functions_action.test.cis_id
		domain_id   = ibm_cis_edge_functions_action.test.domain_id
		pattern_url = "%[2]s"
		action_name = "%[3]s"
	  }
	  `, testName, pattern, actionName)
}

func testAccCheckIBMCisEdgeFunctionsTriggerUpdate(testName, pattern, actionName string) string {
	return testAccCheckIBMCisEdgeFunctionsActionBasic(testName, actionName) + fmt.Sprintf(`
	resource "ibm_cis_edge_functions_trigger" "%[1]s" {
		cis_id      = ibm_cis_edge_functions_action.test.cis_id
		domain_id   = ibm_cis_edge_functions_action.test.domain_id
		pattern_url = "%[2]s"
	  }
	  `, testName, pattern)
}
