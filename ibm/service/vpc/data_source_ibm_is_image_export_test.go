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

func TestAccIBMIsImageExportDataSourceBasic(t *testing.T) {
	imageExportJobImageID := fmt.Sprintf("tf_image_id_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIsImageExportDataSourceConfigBasic(imageExportJobImageID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_image_export.is_image_export", "image_export_job"),
					resource.TestCheckResourceAttrSet("data.ibm_is_image_export.is_image_export", "image"),
					resource.TestCheckResourceAttrSet("data.ibm_is_image_export.is_image_export", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_is_image_export.is_image_export", "format"),
					resource.TestCheckResourceAttrSet("data.ibm_is_image_export.is_image_export", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_image_export.is_image_export", "image_export_job"),
					resource.TestCheckResourceAttrSet("data.ibm_is_image_export.is_image_export", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_is_image_export.is_image_export", "resource_type"),
					resource.TestCheckResourceAttrSet("data.ibm_is_image_export.is_image_export", "status"),
					resource.TestCheckResourceAttrSet("data.ibm_is_image_export.is_image_export", "status_reasons.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_image_export.is_image_export", "storage_bucket.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_image_export.is_image_export", "storage_href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_image_export.is_image_export", "storage_object.#"),
				),
			},
		},
	})
}

func TestAccIBMIsImageExportDataSourceAllArgs(t *testing.T) {
	imageExportJobImageID := fmt.Sprintf("tf_image_id_%d", acctest.RandIntRange(10, 100))
	imageExportJobFormat := "qcow2"
	imageExportJobName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsImageExportDataSourceConfig(imageExportJobImageID, imageExportJobFormat, imageExportJobName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_image_export.is_image_export", "image_export_job"),
					resource.TestCheckResourceAttrSet("data.ibm_is_image_export.is_image_export", "image"),
					resource.TestCheckResourceAttrSet("data.ibm_is_image_export.is_image_export", "completed_at"),
					resource.TestCheckResourceAttrSet("data.ibm_is_image_export.is_image_export", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_is_image_export.is_image_export", "encrypted_data_key"),
					resource.TestCheckResourceAttrSet("data.ibm_is_image_export.is_image_export", "format"),
					resource.TestCheckResourceAttrSet("data.ibm_is_image_export.is_image_export", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_image_export.is_image_export", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_is_image_export.is_image_export", "resource_type"),
					resource.TestCheckResourceAttrSet("data.ibm_is_image_export.is_image_export", "started_at"),
					resource.TestCheckResourceAttrSet("data.ibm_is_image_export.is_image_export", "status"),
					resource.TestCheckResourceAttrSet("data.ibm_is_image_export.is_image_export", "status_reasons.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_image_export.is_image_export", "status_reasons.0.code"),
					resource.TestCheckResourceAttrSet("data.ibm_is_image_export.is_image_export", "status_reasons.0.message"),
					resource.TestCheckResourceAttrSet("data.ibm_is_image_export.is_image_export", "status_reasons.0.more_info"),
					resource.TestCheckResourceAttrSet("data.ibm_is_image_export.is_image_export", "storage_bucket.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_image_export.is_image_export", "storage_href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_image_export.is_image_export", "storage_object.#"),
				),
			},
		},
	})
}

func testAccCheckIBMIsImageExportDataSourceConfigBasic(imageExportJobImageID string) string {
	return fmt.Sprintf(`
		resource "ibm_is_image_export" "is_image_export" {
			image = "%s"
			storage_bucket = "%s"
		}

		data "ibm_is_image_export" "is_image_export" {
			image = ibm_is_image_export.is_image_export.image
			image_export_job = ibm_is_image_export.is_image_export.image_export_job
		}
	`, imageExportJobImageID, acc.IsCosBucketName)
}

func testAccCheckIBMIsImageExportDataSourceConfig(imageExportJobImageID string, imageExportJobFormat string, imageExportJobName string) string {
	return fmt.Sprintf(`
		resource "ibm_is_image_export" "is_image_export" {
			image = "%s"
			storage_bucket_name = "%s"
			format = "%s"
			name = "%s"
		}

		data "ibm_is_image_export" "is_image_export" {
			image = ibm_is_image_export.is_image_export.image
			image_export_job = ibm_is_image_export.is_image_export.image_export_job
		}
	`, imageExportJobImageID, acc.IsCosBucketName, imageExportJobFormat, imageExportJobName)
}
