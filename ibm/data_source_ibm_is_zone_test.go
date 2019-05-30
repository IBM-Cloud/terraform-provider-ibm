package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccIBMISZoneDataSource_basic(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMISZoneDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ibm_is_zone.testacc_ds_zone", "name", ISZoneName),
					resource.TestCheckResourceAttr("data.ibm_is_zone.testacc_ds_zone", "region", regionName),
				),
			},
		},
	})
}

func testAccCheckIBMISZoneDataSourceConfig() string {
	return fmt.Sprintf(`

data "ibm_is_zone" "testacc_ds_zone" {
	name = "%s",
	region = "%s",
}`, ISZoneName, regionName)
}
