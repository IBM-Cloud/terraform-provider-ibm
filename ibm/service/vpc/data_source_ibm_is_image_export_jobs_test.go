// Copyright IBM Corp. 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIBMIsImageExportsDataSourceAllArgs(t *testing.T) {

	imageExportJobFormat := "qcow2"
	imageExportJobName := fmt.Sprintf("tf-exportjob-name%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsImageExportsDataSourceConfig(imageExportJobFormat, imageExportJobName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_image_export_jobs.is_image_exports", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_image_export_jobs.is_image_exports", "image"),
					resource.TestCheckResourceAttrSet("data.ibm_is_image_export_jobs.is_image_exports", "export_jobs.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_image_export_jobs.is_image_exports", "export_jobs.0.created_at"),
					resource.TestCheckResourceAttr("data.ibm_is_image_export_jobs.is_image_exports", "export_jobs.0.format", imageExportJobFormat),
					resource.TestCheckResourceAttrSet("data.ibm_is_image_export_jobs.is_image_exports", "export_jobs.0.href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_image_export_jobs.is_image_exports", "export_jobs.0.image_export_job"),
					resource.TestCheckResourceAttr("data.ibm_is_image_export_jobs.is_image_exports", "export_jobs.0.name", imageExportJobName),
					resource.TestCheckResourceAttrSet("data.ibm_is_image_export_jobs.is_image_exports", "export_jobs.0.resource_type"),
					resource.TestCheckResourceAttrSet("data.ibm_is_image_export_jobs.is_image_exports", "export_jobs.0.status"),
					resource.TestCheckResourceAttrSet("data.ibm_is_image_export_jobs.is_image_exports", "export_jobs.0.storage_bucket.0.name"),
					resource.TestCheckResourceAttrSet("data.ibm_is_image_export_jobs.is_image_exports", "export_jobs.0.storage_bucket.0.crn"),
					resource.TestCheckResourceAttrSet("data.ibm_is_image_export_jobs.is_image_exports", "export_jobs.0.storage_href"),
				),
			},
		},
	})
}

func testAccCheckIBMIsImageExportsDataSourceConfig(imageExportJobFormat string, imageExportJobName string) string {
	return fmt.Sprintf(`
		resource "ibm_is_image_export_job" "is_image_export" {
			image = "%s"
			storage_bucket {
				name = "%s"
			}
			format = "%s"
			name = "%s"
		}

		data "ibm_is_image_export_jobs" "is_image_exports" {
			image = ibm_is_image_export_job.is_image_export.image
		}
	`, acc.IsImage, acc.IsCosBucketName, imageExportJobFormat, imageExportJobName)
}
