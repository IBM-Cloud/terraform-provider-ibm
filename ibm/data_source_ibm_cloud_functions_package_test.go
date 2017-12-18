package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccCloudFunctionsPackageDataSourceBasic(t *testing.T) {
	name := fmt.Sprintf("terraform_%d", acctest.RandInt())

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{

			resource.TestStep{
				Config: testAccCheckCloudFunctionsPackageDataSource(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_cloud_functions_package.package", "name", name),
					resource.TestCheckResourceAttr("ibm_cloud_functions_package.package", "version", "0.0.1"),
					resource.TestCheckResourceAttr("ibm_cloud_functions_package.package", "publish", "false"),
					resource.TestCheckResourceAttr("ibm_cloud_functions_package.package", "parameters", "[]"),
					resource.TestCheckResourceAttr("data.ibm_cloud_functions_package.package", "name", name),
				),
			},
		},
	})
}

func testAccCheckCloudFunctionsPackageDataSource(name string) string {
	return fmt.Sprintf(`
	
resource "ibm_cloud_functions_package" "package" {
	   name = "%s"
}

data "ibm_cloud_functions_package" "package" {
    name = "${ibm_cloud_functions_package.package.name}"
}`, name)

}
