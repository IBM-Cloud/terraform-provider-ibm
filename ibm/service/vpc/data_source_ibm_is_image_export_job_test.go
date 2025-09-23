// Copyright IBM Corp. 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"
)

func TestAccIBMIsImageExportDataSourceAllArgs(t *testing.T) {
	imageExportJobFormat := "qcow2"
	imageExportJobName := fmt.Sprintf("tf-exportjob-name%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsImageExportDataSourceConfig(imageExportJobFormat, imageExportJobName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_image_export_job.is_image_export_job", "image_export_job"),
					resource.TestCheckResourceAttrSet("data.ibm_is_image_export_job.is_image_export_job", "image"),
					resource.TestCheckResourceAttrSet("data.ibm_is_image_export_job.is_image_export_job", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_is_image_export_job.is_image_export_job", "format"),
					resource.TestCheckResourceAttrSet("data.ibm_is_image_export_job.is_image_export_job", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_image_export_job.is_image_export_job", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_is_image_export_job.is_image_export_job", "resource_type"),
					resource.TestCheckResourceAttrSet("data.ibm_is_image_export_job.is_image_export_job", "status"),
					resource.TestCheckResourceAttrSet("data.ibm_is_image_export_job.is_image_export_job", "storage_bucket.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_image_export_job.is_image_export_job", "storage_href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_image_export_job.is_image_export_job", "storage_object.#"),
				),
			},
		},
	})
}

func testAccCheckIBMIsImageExportDataSourceConfig(imageExportJobFormat, imageExportJobName string) string {
	return fmt.Sprintf(`
		resource "ibm_is_image_export_job" "is_image_export_job" {
			image = "%s"
			storage_bucket {
				name = "%s"
			} 
			format = "%s"
			name = "%s"
		}

		data "ibm_is_image_export_job" "is_image_export_job" {
			image = ibm_is_image_export_job.is_image_export_job.image
			image_export_job = ibm_is_image_export_job.is_image_export_job.image_export_job
		}
	`, acc.IsImage, acc.IsCosBucketName, imageExportJobFormat, imageExportJobName)
}
