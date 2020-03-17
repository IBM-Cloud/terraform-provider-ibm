package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccIBMISSubnetsDataSource_basic(t *testing.T) {
	resName := "data.ibm_is_subnets.test1"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMISSubnetsDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resName, "subnets.0.name"),
					resource.TestCheckResourceAttrSet(resName, "subnets.0.status"),
				),
			},
		},
	})
}

func testAccCheckIBMISSubnetsDataSourceConfig() string {
	// status filter defaults to empty
	return fmt.Sprintf(`
      data "ibm_is_subnets" "test1" {
      }`)
}
