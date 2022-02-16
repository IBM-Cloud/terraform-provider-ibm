// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMISZoneDataSource_basic(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISZoneDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ibm_is_zone.testacc_ds_zone", "name", acc.ISZoneName),
					resource.TestCheckResourceAttr("data.ibm_is_zone.testacc_ds_zone", "region", acc.RegionName),
				),
			},
		},
	})
}

func testAccCheckIBMISZoneDataSourceConfig() string {
	return fmt.Sprintf(`

data "ibm_is_zone" "testacc_ds_zone" {
	name = "%s"
	region = "%s"
}`, acc.ISZoneName, acc.RegionName)
}
