// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package power_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccIBMPIImageExport(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
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
			pi_image_id            = data.ibm_pi_image.power_image.id
			pi_cloud_instance_id   = "%[1]s"
			pi_image_bucket_name   = "%[2]s"
			pi_image_access_key    = "%[3]s"
			pi_image_secret_key    = "%[4]s"
			pi_image_bucket_region = "%[5]s"
		}
	`, acc.Pi_cloud_instance_id, acc.Pi_image_bucket_name, acc.Pi_image_bucket_access_key, acc.Pi_image_bucket_secret_key, acc.Pi_image_bucket_region, acc.Pi_image)
}
