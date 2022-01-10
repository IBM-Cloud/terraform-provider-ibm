// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMPIImageEport(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPIImageExportConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("ibm_pi_image_export.power_image_export", "id"),
				),
			},
		},
	})
}

func testAccCheckIBMPIImageExportConfig() string {
	return fmt.Sprintf(`
	data "ibm_pi_image" "power_image" {
		pi_image_name        = "%[6]s"
		pi_cloud_instance_id = "%[1]s"
	  }
	resource "ibm_pi_image_export" "power_image_export" {
		pi_image_id         = data.ibm_pi_image.power_image.id
		pi_cloud_instance_id = "%[1]s"
		pi_image_bucket_name = "%[2]s"
		pi_image_access_key = "%[3]s"
		pi_image_secret_key = "%[4]s"
		pi_image_bucket_region = "%[5]s"
	  }
	`, pi_cloud_instance_id, pi_image_bucket_name, pi_image_bucket_access_key, pi_image_bucket_secret_key, pi_image_bucket_region, pi_image)
}
