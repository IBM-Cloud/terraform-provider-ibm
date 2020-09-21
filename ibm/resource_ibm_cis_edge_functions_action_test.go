package ibm

import (
	"fmt"
	"testing"

	"github.com/IBM/go-sdk-core/v3/core"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

const cisEdgeFunctionsActionScriptTmpFile = "/tmp/script.js"

func TestAccIBMCisEdgeFunctionsAction_Basic(t *testing.T) {
	var record string
	testName := "tf-acctest-basic"
	resourceName := "ibm_cis_edge_functions_action.tf-acctest-basic"
	scriptName := "sample_script"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheckCis(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMCisEdgeFunctionsActionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCisEdgeFunctionsActionBasic(testName, scriptName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMCisEdgeFunctionsActionExists(resourceName, &record),
					resource.TestCheckResourceAttr(
						resourceName, "script_name", scriptName),
				),
			},
		},
	})
}

func TestAccIBMCisEdgeFunctionsAction_import(t *testing.T) {
	name := "ibm_cis_edge_functions_action.test"
	scriptName := "sample_script"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMCisEdgeFunctionsActionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCisEdgeFunctionsActionBasic("test", scriptName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "script_name", scriptName),
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

func TestAccIBMCisFunctionsAction_CreateAfterManualDestroy(t *testing.T) {
	var scriptOne, scriptTwo string
	name := "ibm_cis_edge_functions_action.test"
	scriptOne = "script_one"
	scriptTwo = "script_two"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
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
		cisClient, err := testAccProvider.Meta().(ClientSession).CisEdgeFunctionClientSession()
		if err != nil {
			return fmt.Errorf("Error in creating CIS object")
		}

		scriptName, zoneID, cisID, err := convertTfToCisThreeVar(*tfActionID)
		cisClient.Crn = core.StringPtr(cisID)
		cisClient.ZoneIdentifier = core.StringPtr(zoneID)
		opt := cisClient.NewDeleteEdgeFunctionsActionOptions(scriptName)
		_, response, err := cisClient.DeleteEdgeFunctionsAction(opt)
		if err != nil {
			return fmt.Errorf("Edge function action script deletion failed: %v", response)
		}
		return nil
	}
}

func testAccCheckIBMCisEdgeFunctionsActionDestroy(s *terraform.State) error {
	cisClient, err := testAccProvider.Meta().(ClientSession).CisEdgeFunctionClientSession()
	if err != nil {
		return fmt.Errorf("Error in creating CIS object")
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_cis_edge_functions_action" {
			continue
		}

		scriptName, zoneID, cisID, err := convertTfToCisThreeVar(rs.Primary.ID)
		cisClient.Crn = core.StringPtr(cisID)
		cisClient.ZoneIdentifier = core.StringPtr(zoneID)
		opt := cisClient.NewGetEdgeFunctionsActionOptions(scriptName)
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
		scriptName, zoneID, cisID, err := convertTfToCisThreeVar(rs.Primary.ID)
		cisClient.Crn = core.StringPtr(cisID)
		cisClient.ZoneIdentifier = core.StringPtr(zoneID)
		opt := cisClient.NewGetEdgeFunctionsActionOptions(scriptName)
		_, resp, err := cisClient.GetEdgeFunctionsAction(opt)
		if err != nil {
			return fmt.Errorf("Error: %v", resp)
		}

		tfRecord = convertCisToTfThreeVar(scriptName, zoneID, cisID)
		*tfRecordID = tfRecord
		return nil
	}
}

func testAccCheckIBMCisEdgeFunctionsActionBasic(testName, scriptName string) string {
	content := `addEventListener('fetch', (event) => {\n\tevent.respondWith(handleRequest(event.request))\n})\n\n/**\n * Sample test function\n * Log a given request object\n * @param {Request} request\n */\nasync function handleRequest(request) {\n\tconsole.log('Got request', request)\n\tconst response = await fetch(request)\n\treturn response;\n}`

	return testAccCheckIBMCisDomainDataSourceConfigBasic1() + fmt.Sprintf(`
	resource "ibm_cis_edge_functions_action" "%[1]s" {
		cis_id      = data.ibm_cis.cis.id
		domain_id   = data.ibm_cis_domain.cis_domain.domain_id
		script_name = "%[2]s"
		script = "%[3]s"
	  }
	  `, testName, scriptName, content)
}
