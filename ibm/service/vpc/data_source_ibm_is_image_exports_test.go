// Copyright IBM Corp. 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIBMIsImageExportsDataSourceBasic(t *testing.T) {
	imageExportJobImageID := fmt.Sprintf("tf_image_id_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIsImageExportsDataSourceConfigBasic(imageExportJobImageID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_image_exports.is_image_exports", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_image_exports.is_image_exports", "image"),
					resource.TestCheckResourceAttrSet("data.ibm_is_image_exports.is_image_exports", "export_jobs.#"),
				),
			},
		},
	})
}

func TestAccIBMIsImageExportsDataSourceAllArgs(t *testing.T) {
	imageExportJobImageID := fmt.Sprintf("tf_image_id_%d", acctest.RandIntRange(10, 100))
	imageExportJobFormat := "qcow2"
	imageExportJobName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsImageExportsDataSourceConfig(imageExportJobImageID, imageExportJobFormat, imageExportJobName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_image_exports.is_image_exports", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_image_exports.is_image_exports", "image"),
					resource.TestCheckResourceAttrSet("data.ibm_is_image_exports.is_image_exports", "export_jobs.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_image_exports.is_image_exports", "export_jobs.0.completed_at"),
					resource.TestCheckResourceAttrSet("data.ibm_is_image_exports.is_image_exports", "export_jobs.0.created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_is_image_exports.is_image_exports", "export_jobs.0.encrypted_data_key"),
					resource.TestCheckResourceAttr("data.ibm_is_image_exports.is_image_exports", "export_jobs.0.format", imageExportJobFormat),
					resource.TestCheckResourceAttrSet("data.ibm_is_image_exports.is_image_exports", "export_jobs.0.href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_image_exports.is_image_exports", "export_jobs.0.image_export_job"),
					resource.TestCheckResourceAttr("data.ibm_is_image_exports.is_image_exports", "export_jobs.0.name", imageExportJobName),
					resource.TestCheckResourceAttrSet("data.ibm_is_image_exports.is_image_exports", "export_jobs.0.resource_type"),
					resource.TestCheckResourceAttrSet("data.ibm_is_image_exports.is_image_exports", "export_jobs.0.started_at"),
					resource.TestCheckResourceAttrSet("data.ibm_is_image_exports.is_image_exports", "export_jobs.0.status"),
					resource.TestCheckResourceAttrSet("data.ibm_is_image_exports.is_image_exports", "export_jobs.0.storage_href"),
				),
			},
		},
	})
}

func testAccCheckIBMIsImageExportsDataSourceConfigBasic(imageExportJobImageID string) string {
	return fmt.Sprintf(`
		resource "ibm_is_image_export" "is_image_export" {
			image = "%s"
			storage_bucket = "%s"
		}

		data "ibm_is_image_exports" "is_image_exports" {
			image = ibm_is_image_export.is_image_export.image
		}
	`, imageExportJobImageID, acc.IsCosBucketName)
}

func testAccCheckIBMIsImageExportsDataSourceConfig(imageExportJobImageID string, imageExportJobFormat string, imageExportJobName string) string {
	return fmt.Sprintf(`
		resource "ibm_is_image_export" "is_image_export" {
			image = "%s"
			storage_bucket = "%s"
			format = "%s"
			name = "%s"
		}

		data "ibm_is_image_exports" "is_image_exports" {
			image = ibm_is_image_export.is_image_export.image
		}
	`, imageExportJobImageID, acc.IsCosBucketName, imageExportJobFormat, imageExportJobName)
}
