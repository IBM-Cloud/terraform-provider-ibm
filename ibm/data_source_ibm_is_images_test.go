// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccIBMISImagesDataSource_basic(t *testing.T) {
	resName := "data.ibm_is_images.test1"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMISImagesDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resName, "images.0.name"),
					resource.TestCheckResourceAttrSet(resName, "images.0.status"),
					resource.TestCheckResourceAttrSet(resName, "images.0.architecture"),
				),
			},
		},
	})
}

func testAccCheckIBMISImagesDataSourceConfig() string {
	// status filter defaults to empty
	return fmt.Sprintf(`
      data "ibm_is_images" "test1" {
      }`)
}
