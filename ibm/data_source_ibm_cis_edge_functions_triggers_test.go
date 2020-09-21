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
					resource.TestCheckResourceAttrSet(node, "cis_edge_functions_triggers.0.pattern"),
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
