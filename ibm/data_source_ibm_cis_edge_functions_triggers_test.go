/*
* IBM Confidential
* Object Code Only Source Materials
* 5747-SM3
* (c) Copyright IBM Corp. 2017,2021
*
* The source code for this program is not published or otherwise divested
* of its trade secrets, irrespective of what has been deposited with the
* U.S. Copyright Office.
 */

package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccIBMCisEdgeFunctionsTriggersDataSource_basic(t *testing.T) {
	node := "data.ibm_cis_edge_functions_triggers.test"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCisEdgeFunctionsTriggersDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(node, "cis_edge_functions_triggers.0.pattern_url"),
				),
			},
		},
	})
}

func testAccCheckIBMCisEdgeFunctionsTriggersDataSourceConfig() string {
	testName := "test"
	scriptName := "sample_script"
	pattern := fmt.Sprintf("example.%s/*", cisDomainStatic)
	return testAccCheckIBMCisEdgeFunctionsTriggerBasic(testName, pattern, scriptName) + fmt.Sprintf(`
	data "ibm_cis_edge_functions_triggers" "test" {
		cis_id    = ibm_cis_edge_functions_trigger.test.cis_id
		domain_id = ibm_cis_edge_functions_trigger.test.domain_id
	}`)
}
