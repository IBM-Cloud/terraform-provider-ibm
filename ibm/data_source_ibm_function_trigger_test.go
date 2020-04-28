package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccFunctionTriggerDataSourceBasic(t *testing.T) {
	name := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{

			resource.TestStep{
				Config: testAccCheckFunctionTriggerDataSource(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_function_trigger.trigger", "name", name),
					resource.TestCheckResourceAttr("ibm_function_trigger.trigger", "version", "0.0.1"),
					resource.TestCheckResourceAttr("ibm_function_trigger.trigger", "publish", "false"),
					resource.TestCheckResourceAttr("data.ibm_function_trigger.datatrigger", "name", name),
				),
			},
		},
	})
}

func testAccCheckFunctionTriggerDataSource(name string) string {
	return fmt.Sprintf(`
	
resource "ibm_function_trigger" "trigger" {
	name = "%s"		  
}
data "ibm_function_trigger" "datatrigger" {
	name = "${ibm_function_trigger.trigger.name}"

}
`, name)

}
