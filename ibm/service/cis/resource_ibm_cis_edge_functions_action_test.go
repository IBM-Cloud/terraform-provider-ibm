// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package cis_test

import (
	"fmt"
	"testing"

	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/conns"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/flex"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccIBMCisEdgeFunctionsAction_Basic(t *testing.T) {
	var record string
	testName := "tf-acctest-basic"
	resourceName := "ibm_cis_edge_functions_action.tf-acctest-basic"
	actionName := "sample_script"
	content1 := "addEventListener('fetch', (event) => {\n\tevent.respondWith(handleRequest(event.request))\n})\n\n/**\n * Sample test function\n * Log a given request object\n * @param {Request} request\n */\nasync function handleRequest(request) {\n\tconsole.log('Got request', request)\n\tconst response = await fetch(request)\n\treturn response;\n}"
	content2 := "addEventListener('fetch', (event) => {\n\tevent.respondWith(handleRequest(event.request))\n})\n\n/**\n * Sample test function\n * @param {Request} request\n */\nasync function handleRequest(request) {\n\tconsole.log('Got request', request)\n\tconst response = await fetch(request)\n\treturn response;\n}"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheckCis(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCisEdgeFunctionsActionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCisEdgeFunctionsActionBasic(testName, actionName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMCisEdgeFunctionsActionExists(resourceName, &record),
					resource.TestCheckResourceAttr(
						resourceName, "action_name", actionName),
					resource.TestCheckResourceAttr(
						resourceName, "script", content1),
				),
			},
			{
				Config: testAccCheckIBMCisEdgeFunctionsActionUpdate(testName, actionName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMCisEdgeFunctionsActionExists(resourceName, &record),
					resource.TestCheckResourceAttr(
						resourceName, "action_name", actionName),
					resource.TestCheckResourceAttr(
						resourceName, "script", content2),
				),
			},
		},
	})
}

func TestAccIBMCisEdgeFunctionsAction_import(t *testing.T) {
	name := "ibm_cis_edge_functions_action.test"
	actionName := "sample_script"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCisEdgeFunctionsActionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCisEdgeFunctionsActionBasic("test", actionName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "action_name", actionName),
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

func TestAccIBMCisFunctionsAction_CreateAfterManualDestroy(t *testing.T) {
	var scriptOne, scriptTwo string
	name := "ibm_cis_edge_functions_action.test"
	scriptOne = "script_one"
	scriptTwo = "script_two"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCisEdgeFunctionsActionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCisEdgeFunctionsActionBasic("test", scriptOne),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMCisEdgeFunctionsActionExists(name, &scriptOne),
					testAccCheckIBMCisEdgeFunctionsActionDelete(&scriptOne),
				),
				ExpectNonEmptyPlan: true,
			},
			{
				Config: testAccCheckIBMCisEdgeFunctionsActionBasic("test", scriptTwo),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMCisEdgeFunctionsActionExists(name, &scriptTwo),
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

func testAccCheckIBMCisEdgeFunctionsActionDelete(tfActionID *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		cisClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).CisEdgeFunctionClientSession()
		if err != nil {
			return fmt.Errorf("[ERROR] Error in creating CIS object")
		}

		actionName, zoneID, cisID, _ := flex.ConvertTfToCisThreeVar(*tfActionID)
		cisClient.Crn = core.StringPtr(cisID)
		cisClient.ZoneIdentifier = core.StringPtr(zoneID)
		opt := cisClient.NewDeleteEdgeFunctionsActionOptions(actionName)
		_, response, err := cisClient.DeleteEdgeFunctionsAction(opt)
		if err != nil {
			return fmt.Errorf("Edge function action script deletion failed: %v", response)
		}
		return nil
	}
}

func testAccCheckIBMCisEdgeFunctionsActionDestroy(s *terraform.State) error {
	cisClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).CisEdgeFunctionClientSession()
	if err != nil {
		return fmt.Errorf("[ERROR] Error in creating CIS object")
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_cis_edge_functions_action" {
			continue
		}

		actionName, zoneID, cisID, _ := flex.ConvertTfToCisThreeVar(rs.Primary.ID)
		cisClient.Crn = core.StringPtr(cisID)
		cisClient.ZoneIdentifier = core.StringPtr(zoneID)
		opt := cisClient.NewGetEdgeFunctionsActionOptions(actionName)
		_, response, err := cisClient.GetEdgeFunctionsAction(opt)
		if err == nil {
			return fmt.Errorf("Edge function action script still exists: %v", response)
		}
	}

	return nil
}

func testAccCheckIBMCisEdgeFunctionsActionExists(n string, tfRecordID *string) resource.TestCheckFunc {
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
		actionName, zoneID, cisID, _ := flex.ConvertTfToCisThreeVar(rs.Primary.ID)
		cisClient.Crn = core.StringPtr(cisID)
		cisClient.ZoneIdentifier = core.StringPtr(zoneID)
		opt := cisClient.NewGetEdgeFunctionsActionOptions(actionName)
		_, resp, err := cisClient.GetEdgeFunctionsAction(opt)
		if err != nil {
			return fmt.Errorf("[ERROR] Error: %v", resp)
		}

		tfRecord := flex.ConvertCisToTfThreeVar(actionName, zoneID, cisID)
		*tfRecordID = tfRecord
		return nil
	}
}

func testAccCheckIBMCisEdgeFunctionsActionBasic(testName, actionName string) string {
	content := `addEventListener('fetch', (event) => {\n\tevent.respondWith(handleRequest(event.request))\n})\n\n/**\n * Sample test function\n * Log a given request object\n * @param {Request} request\n */\nasync function handleRequest(request) {\n\tconsole.log('Got request', request)\n\tconst response = await fetch(request)\n\treturn response;\n}`

	return testAccCheckIBMCisDomainDataSourceConfigBasic1() + fmt.Sprintf(`
	resource "ibm_cis_edge_functions_action" "%[1]s" {
		cis_id      = data.ibm_cis.cis.id
		domain_id   = data.ibm_cis_domain.cis_domain.domain_id
		action_name = "%[2]s"
		script      = "%[3]s"
	  }
	  `, testName, actionName, content)
}

func testAccCheckIBMCisEdgeFunctionsActionUpdate(testName, actionName string) string {
	content := `addEventListener('fetch', (event) => {\n\tevent.respondWith(handleRequest(event.request))\n})\n\n/**\n * Sample test function\n * @param {Request} request\n */\nasync function handleRequest(request) {\n\tconsole.log('Got request', request)\n\tconst response = await fetch(request)\n\treturn response;\n}`

	return testAccCheckIBMCisDomainDataSourceConfigBasic1() + fmt.Sprintf(`
	resource "ibm_cis_edge_functions_action" "%[1]s" {
		cis_id      = data.ibm_cis.cis.id
		domain_id   = data.ibm_cis_domain.cis_domain.domain_id
		action_name = "%[2]s"
		script      = "%[3]s"
	  }
	  `, testName, actionName, content)
}
