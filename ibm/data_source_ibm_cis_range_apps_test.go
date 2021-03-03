// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccIBMCisRangeAppsDataSource_basic(t *testing.T) {
	node := "data.ibm_cis_range_apps.test"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCisRangeAppsDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(node, "range_apps.0.id"),
					resource.TestCheckResourceAttrSet(node, "range_apps.0.origin_direct.0"),
				),
			},
		},
	})
}

func testAccCheckIBMCisRangeAppsDataSourceConfig() string {
	return testAccCheckCisRangeAppConfigBasic() + fmt.Sprintf(`
	data "ibm_cis_range_apps" "test" {
		cis_id     = ibm_cis_range_app.app.cis_id
		domain_id  = ibm_cis_range_app.app.domain_id
	}`)
}
