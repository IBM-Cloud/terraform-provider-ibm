// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMISRegionDataSource_basic(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISRegionDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ibm_is_region.testacc_ds_region", "name", acc.RegionName),
					resource.TestCheckResourceAttr("data.ibm_is_region.testacc_default_region", "name", acc.RegionName),
				),
			},
		},
	})
}

func testAccCheckIBMISRegionDataSourceConfig() string {
	return fmt.Sprintf(`

data "ibm_is_region" "testacc_ds_region" {
	name = "%s"
}
data "ibm_is_region" "testacc_default_region" {
}
`, acc.RegionName)

}
