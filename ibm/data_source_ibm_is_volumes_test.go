// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMIsVolumesDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsVolumesDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_volumes.is_volumes", "id"),
					// resource.TestCheckResourceAttrSet("data.ibm_is_volumes.is_volumes", "first.#"),
					// resource.TestCheckResourceAttrSet("data.ibm_is_volumes.is_volumes", "limit"),
					resource.TestCheckResourceAttrSet("data.ibm_is_volumes.is_volumes", "volumes.#"),
				),
			},
		},
	})
}

func testAccCheckIBMIsVolumesDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		data "ibm_is_volumes" "is_volumes" {
		}
	`)
}
