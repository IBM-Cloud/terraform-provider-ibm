// Copyright IBM Corp. 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/IBM/vpc-go-sdk/vpcv1"
	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/conns"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/flex"
)

func TestAccIBMIsImageExportBasic(t *testing.T) {
	var conf vpcv1.ImageExportJob

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIsImageExportDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsImageExportConfigBasic(),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIsImageExportExists("ibm_is_image_export_job.is_image_export", conf),
					resource.TestCheckResourceAttr("ibm_is_image_export_job.is_image_export", "image", acc.IsImage),
					resource.TestCheckResourceAttrSet("ibm_is_image_export_job.is_image_export", "status"),
					resource.TestCheckResourceAttrSet("ibm_is_image_export_job.is_image_export", "storage_bucket.0.name"),
					resource.TestCheckResourceAttrSet("ibm_is_image_export_job.is_image_export", "storage_bucket.0.crn"),
				),
			},
		},
	})
}

func TestAccIBMIsImageExportAllArgs(t *testing.T) {
	var conf vpcv1.ImageExportJob
	format := "qcow2"
	name := fmt.Sprintf("tf-imageexport-job%d", acctest.RandIntRange(10, 100))
	nameUpdate := fmt.Sprintf("tf-name%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIsImageExportDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsImageExportConfig(format, name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIsImageExportExists("ibm_is_image_export_job.is_image_export", conf),
					resource.TestCheckResourceAttr("ibm_is_image_export_job.is_image_export", "format", format),
					resource.TestCheckResourceAttr("ibm_is_image_export_job.is_image_export", "name", name),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMIsImageExportConfig(format, nameUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_is_image_export_job.is_image_export", "name", nameUpdate),
				),
			},
		},
	})
}

func testAccCheckIBMIsImageExportConfigBasic() string {
	return fmt.Sprintf(`

		resource "ibm_is_image_export_job" "is_image_export" {
			image = "%s"
			storage_bucket {
				name = "%s"
			}
		}
	`, acc.IsImage, acc.IsCosBucketName)
}

func testAccCheckIBMIsImageExportConfig(format string, name string) string {
	return fmt.Sprintf(`

		resource "ibm_is_image_export_job" "is_image_export" {
			image = "%s"
			storage_bucket {
				crn = "%s"
			}
			format = "%s"
			name = "%s"
		}
	`, acc.IsImage, acc.IsCosBucketCRN, format, name)
}

func testAccCheckIBMIsImageExportExists(n string, obj vpcv1.ImageExportJob) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		vpcClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).VpcV1API()
		if err != nil {
			return err
		}

		getImageExportJobOptions := &vpcv1.GetImageExportJobOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getImageExportJobOptions.SetImageID(parts[0])
		getImageExportJobOptions.SetID(parts[1])

		imageExportJob, _, err := vpcClient.GetImageExportJob(getImageExportJobOptions)
		if err != nil {
			return err
		}

		obj = *imageExportJob
		return nil
	}
}

func testAccCheckIBMIsImageExportDestroy(s *terraform.State) error {
	vpcClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).VpcV1API()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_is_image_export_job" {
			continue
		}

		getImageExportJobOptions := &vpcv1.GetImageExportJobOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getImageExportJobOptions.SetImageID(parts[0])
		getImageExportJobOptions.SetID(parts[1])

		// Try to find the key
		_, response, err := vpcClient.GetImageExportJob(getImageExportJobOptions)

		if err == nil {
			return fmt.Errorf("ImageExportJob still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for ImageExportJob (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
