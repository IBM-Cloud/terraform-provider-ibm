package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccCloudFunctionsTriggerDataSourceBasic(t *testing.T) {
	name := fmt.Sprintf("terraform_%d", acctest.RandInt())

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{

			resource.TestStep{
				Config: testAccCheckCloudFunctionsTriggerDataSource(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_cloud_functions_trigger.trigger", "name", name),
					resource.TestCheckResourceAttr("ibm_cloud_functions_trigger.trigger", "version", "0.0.1"),
					resource.TestCheckResourceAttr("ibm_cloud_functions_trigger.trigger", "publish", "false"),
					resource.TestCheckResourceAttr("data.ibm_cloud_functions_trigger.datatrigger", "name", name),
				),
			},
		},
	})
}

func testAccCheckCloudFunctionsTriggerDataSource(name string) string {
	return fmt.Sprintf(`
	
resource "ibm_cloud_functions_trigger" "trigger" {
	name = "%s"		  
}
data "ibm_cloud_functions_trigger" "datatrigger" {
	name = "${ibm_cloud_functions_trigger.trigger.name}"

}
`, name)

}
