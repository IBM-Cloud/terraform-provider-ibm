package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccIBMISRegionDataSource_basic(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMISRegionDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ibm_is_region.testacc_ds_region", "name", regionName),
				),
			},
		},
	})
}

func testAccCheckIBMISRegionDataSourceConfig() string {
	return fmt.Sprintf(`

data "ibm_is_region" "testacc_ds_region" {
	name = "%s"
}`, regionName)

}
