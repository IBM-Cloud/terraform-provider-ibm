package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccIBMCisEdgeFunctionsActionsDataSource_basic(t *testing.T) {
	node := "data.ibm_cis_edge_functions_actions.test"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCisEdgeFunctionsActionsDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(node, "cis_edge_functions_actions.0.etag"),
				),
			},
		},
	})
}

func testAccCheckIBMCisEdgeFunctionsActionsDataSourceConfig() string {
	testName := "tf-acctest-basic"
	scriptName := "sample_script"

	return testAccCheckIBMCisEdgeFunctionsActionBasic(testName, scriptName) + fmt.Sprintf(`
	data "ibm_cis_edge_functions_actions" "test" {
		cis_id    = ibm_cis_edge_functions_action.tf-acctest-basic.cis_id
		domain_id = ibm_cis_edge_functions_action.tf-acctest-basic.domain_id
	  }`)
}
