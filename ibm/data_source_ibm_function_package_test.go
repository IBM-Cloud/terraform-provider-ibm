package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccFunctionPackageDataSourceBasic(t *testing.T) {
	name := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{

			resource.TestStep{
				Config: testAccCheckFunctionPackageDataSource(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_function_package.package", "name", name),
					resource.TestCheckResourceAttr("ibm_function_package.package", "version", "0.0.1"),
					resource.TestCheckResourceAttr("ibm_function_package.package", "publish", "false"),
					resource.TestCheckResourceAttr("ibm_function_package.package", "parameters", "[]"),
					resource.TestCheckResourceAttr("data.ibm_function_package.package", "name", name),
				),
			},
		},
	})
}

func testAccCheckFunctionPackageDataSource(name string) string {
	return fmt.Sprintf(`
	
resource "ibm_function_package" "package" {
	   name = "%s"
}

data "ibm_function_package" "package" {
    name = ibm_function_package.package.name
}`, name)

}
