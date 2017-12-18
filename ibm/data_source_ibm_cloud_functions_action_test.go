package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccCloudFunctionsActionDataSourceBasic(t *testing.T) {
	name := fmt.Sprintf("terraform_%d", acctest.RandInt())

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{

			resource.TestStep{
				Config: testAccCheckCloudFunctionsActionDataSource(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_cloud_functions_action.pythonzip", "name", name),
					resource.TestCheckResourceAttr("ibm_cloud_functions_action.pythonzip", "exec.799663035.kind", "python"),
					resource.TestCheckResourceAttr("data.ibm_cloud_functions_action.action", "name", name),
				),
			},
		},
	})
}

func testAccCheckCloudFunctionsActionDataSource(name string) string {
	return fmt.Sprintf(`
	
resource "ibm_cloud_functions_action" "pythonzip" {
	name = "%s"		  
	exec = {
		kind = "python"
		code = "${base64encode("${file("test-fixtures/pythonaction.zip")}")}"
	}
}
data "ibm_cloud_functions_action" "action" {
    name = "${ibm_cloud_functions_action.pythonzip.name}"
}
`, name)

}
