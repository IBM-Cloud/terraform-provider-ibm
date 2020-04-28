package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccFunctionActionDataSourceBasic(t *testing.T) {
	name := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{

			resource.TestStep{
				Config: testAccCheckFunctionActionDataSource(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_function_action.pythonzip", "name", name),
					resource.TestCheckResourceAttr("ibm_function_action.pythonzip", "exec.0.kind", "python"),
					resource.TestCheckResourceAttr("data.ibm_function_action.action", "name", name),
				),
			},
		},
	})
}

func testAccCheckFunctionActionDataSource(name string) string {
	return fmt.Sprintf(`
	
resource "ibm_function_action" "pythonzip" {
	name = "%s"		  
	exec = {
		kind = "python"
		code = "${base64encode("${file("test-fixtures/pythonaction.zip")}")}"
	}
}
data "ibm_function_action" "action" {
    name = "${ibm_function_action.pythonzip.name}"
}
`, name)

}
