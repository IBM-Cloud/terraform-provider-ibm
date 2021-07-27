// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMISOperatingSystemDataSource_basic(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISOperatingSystemDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_operating_system.testacc_ds_os", "name"),
				),
			},
		},
	})
}

func testAccCheckIBMISOperatingSystemDataSourceConfig() string {
	return fmt.Sprintf(`
		data "ibm_is_operating_systems" "testacc_ds_oslist" {
		}
		data "ibm_is_operating_system" "testacc_ds_os" {
			name = data.ibm_is_operating_systems.testacc_ds_oslist.operating_systems.0.name
		}`)
}
