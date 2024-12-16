// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package cis_test

import (
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccIBMCisRangeAppsDataSource_basic(t *testing.T) {
	node := "data.ibm_cis_range_apps.test"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
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
	return testAccCheckCisRangeAppConfigBasic() + `
	data "ibm_cis_range_apps" "test" {
		cis_id     = ibm_cis_range_app.app.cis_id
		domain_id  = ibm_cis_range_app.app.domain_id
	}`
}
