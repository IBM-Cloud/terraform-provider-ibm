package ibm

import (
	"fmt"
	"testing"

	"github.com/IBM/go-sdk-core/v3/core"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccIBMCisEdgeFunctionsTrigger_Basic(t *testing.T) {
	var record string
	testName := "test"
	resourceName := fmt.Sprintf("ibm_cis_edge_functions_trigger.%s", testName)
	scriptName := "sample_script"
	pattern := fmt.Sprintf("example.%s/*", cisDomainStatic)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheckCis(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMCisEdgeFunctionsActionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCisEdgeFunctionsTriggerBasic(testName, pattern, scriptName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMCisEdgeFunctionsTriggerExists(resourceName, &record),
					resource.TestCheckResourceAttr(
						resourceName, "script", scriptName),
				),
			},
		},
	})
}

func TestAccIBMCisEdgeFunctionsTrigger_import(t *testing.T) {
	name := "ibm_cis_edge_functions_trigger.test"
	scriptName := "sample_script"
	pattern := fmt.Sprintf("example.%s/*", cisDomainStatic)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMCisEdgeFunctionsTriggerDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCisEdgeFunctionsTriggerBasic("test", pattern, scriptName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "script", scriptName),
				),
			},
			{
				ResourceName:      name,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"wait_time_minutes"},
			},
		},
	})
}

func TestAccIBMCisFunctionsTrigger_CreateAfterManualDestroy(t *testing.T) {
	var scriptOne, scriptTwo string
	name := "ibm_cis_edge_functions_trigger.test"
	scriptOne = "script_one"
	scriptTwo = "script_two"
	pattern := fmt.Sprintf("example.%s/*", cisDomainStatic)
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
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
		cisClient, err := testAccProvider.Meta().(ClientSession).CisEdgeFunctionClientSession()
		if err != nil {
			return fmt.Errorf("Error in creating CIS object")
		}

		scriptName, zoneID, cisID, _ := convertTfToCisThreeVar(*tfActionID)
		cisClient.Crn = core.StringPtr(cisID)
		cisClient.ZoneIdentifier = core.StringPtr(zoneID)
		opt := cisClient.NewDeleteEdgeFunctionsTriggerOptions(scriptName)
		_, response, err := cisClient.DeleteEdgeFunctionsTrigger(opt)
		if err != nil {
			return fmt.Errorf("Edge function action script deletion failed: %v", response)
		}
		return nil
	}
}

func testAccCheckIBMCisEdgeFunctionsTriggerDestroy(s *terraform.State) error {
	cisClient, err := testAccProvider.Meta().(ClientSession).CisEdgeFunctionClientSession()
	if err != nil {
		return fmt.Errorf("Error in creating CIS object")
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_cis_edge_functions_trigger" {
			continue
		}

		triggerID, zoneID, cisID, _ := convertTfToCisThreeVar(rs.Primary.ID)
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
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Record ID is set")
		}

		tfRecord := *tfRecordID
		cisClient, err := testAccProvider.Meta().(ClientSession).CisEdgeFunctionClientSession()
		if err != nil {
			return fmt.Errorf("Error in creating CIS object")
		}
		triggerID, zoneID, cisID, _ := convertTfToCisThreeVar(rs.Primary.ID)
		cisClient.Crn = core.StringPtr(cisID)
		cisClient.ZoneIdentifier = core.StringPtr(zoneID)
		opt := cisClient.NewGetEdgeFunctionsTriggerOptions(triggerID)
		result, resp, err := cisClient.GetEdgeFunctionsTrigger(opt)
		if err != nil {
			return fmt.Errorf("Error: %v", resp)
		}

		tfRecord = convertCisToTfThreeVar(*result.Result.ID, zoneID, cisID)
		*tfRecordID = tfRecord
		return nil
	}
}

func testAccCheckIBMCisEdgeFunctionsTriggerBasic(testName, pattern, scriptName string) string {
	return testAccCheckIBMCisEdgeFunctionsActionBasic(testName, scriptName) + fmt.Sprintf(`
	resource "ibm_cis_edge_functions_trigger" "%[1]s" {
		cis_id    = ibm_cis_edge_functions_action.test.cis_id
		domain_id = ibm_cis_edge_functions_action.test.domain_id
		pattern   = "%[2]s"
		script    = "%[3]s"
	  }
	  `, testName, pattern, scriptName)
}
