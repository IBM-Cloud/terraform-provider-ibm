// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"fmt"
	"testing"

	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMISOperatingSystemsDataSource_basic(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISOperatingSystemsDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_operating_systems.testacc_ds_oslist", "operating_systems.0.name"),
					resource.TestCheckResourceAttrSet("data.ibm_is_operating_systems.testacc_ds_oslist", "operating_systems.0.allow_user_image_creation"),
					resource.TestCheckResourceAttrSet("data.ibm_is_operating_systems.testacc_ds_oslist", "operating_systems.0.user_data_format"),
				),
			},
		},
	})
}

func testAccCheckIBMISOperatingSystemsDataSourceConfig() string {
	return fmt.Sprintf(`
		data "ibm_is_operating_systems" "testacc_ds_oslist" {
			}`)
}
