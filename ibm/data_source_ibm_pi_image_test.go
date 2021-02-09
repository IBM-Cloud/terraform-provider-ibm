/**
 * © Copyright IBM Corporation 2020. All Rights Reserved.
 *
 * Licensed under the Mozilla Public License, version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at https://mozilla.org/MPL/2.0/
 */

package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccIBMPIImageDataSource_basic(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPIImageDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_pi_image.testacc_ds_image", "id"),
				),
			},
		},
	})
}

func testAccCheckIBMPIImageDataSourceConfig() string {
	return fmt.Sprintf(`
	resource "ibm_pi_image" "power_image" {
		pi_image_name       = "7200-04-01"
		pi_image_id         = "f31da27a-b634-45e5-913a-3f4d964e5a02"
		pi_cloud_instance_id = "%[1]s"
	  }
	data "ibm_pi_image" "testacc_ds_image" {
		pi_image_name = ibm_pi_image.power_image.image_id
		pi_cloud_instance_id = "%[1]s"
	}`, pi_cloud_instance_id)

}
