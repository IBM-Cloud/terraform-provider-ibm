// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMIsVolumesDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsVolumesDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_volumes.is_volumes", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_volumes.is_volumes", "volumes.#"),
				),
			},
			{
				Config: testAccCheckIBMIsVolumesDataSourceConfigFilterByZone(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_volumes.is_volumes", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_volumes.is_volumes", "volumes.#"),
				),
			},
			{
				Config: testAccCheckIBMIsVolumesDataSourceConfigFilterByName(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_volumes.is_volumes", "id"),
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

func testAccCheckIBMIsVolumesDataSourceConfigFilterByZone() string {
	return fmt.Sprintf(`
		data "ibm_is_volumes" "is_volumes" {
			zone_name = "us-south-1"
		}
	`)
}

func testAccCheckIBMIsVolumesDataSourceConfigFilterByName() string {
	return fmt.Sprintf(`
		data "ibm_is_volumes" "is_volumes" {
			volume_name = "worrier-mailable-timpani-scowling"
		}
	`)
}
