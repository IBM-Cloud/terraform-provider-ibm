package ibm

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccIBMComputeImageTemplateDataSource_Basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			// Tests looking up private or shared images
			{
				Config: testAccCheckIBMComputeImageTemplateDataSourceConfig_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"data.ibm_compute_image_template.tfacc_img_tmpl",
						"name",
						"jumpbox",
					),
					resource.TestMatchResourceAttr(
						"data.ibm_compute_image_template.tfacc_img_tmpl",
						"id",
						regexp.MustCompile("^[0-9]+$"),
					),
				),
			},
			// Tests looking up a public image
			{
				Config: testAccCheckIBMComputeImageTemplateDataSourceConfig_basic2,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"data.ibm_compute_image_template.tfacc_img_tmpl",
						"name",
						"RightImage_Ubuntu_12.04_amd64_v13.5",
					),
					resource.TestMatchResourceAttr(
						"data.ibm_compute_image_template.tfacc_img_tmpl",
						"id",
						regexp.MustCompile("^[0-9]+$"),
					),
				),
			},
			// Tests looking up the latest of a public image
			{
				Config: testAccCheckIBMComputeImageTemplateDataSourceConfig_latest,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"data.ibm_compute_image_template.tfacc_img_tmpl",
						"name",
						"25GB - Ubuntu / Ubuntu / 18.04-64 Minimal for VSI",
					),
					resource.TestMatchResourceAttr(
						"data.ibm_compute_image_template.tfacc_img_tmpl",
						"id",
						regexp.MustCompile("^[0-9]+$"),
					),
				),
			},
			// Tests looking up the first returned of a public image
			{
				Config: testAccCheckIBMComputeImageTemplateDataSourceConfig_notlatest,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"data.ibm_compute_image_template.tfacc_img_tmpl",
						"name",
						"25GB - Ubuntu / Ubuntu / 18.04-64 Minimal for VSI",
					),
					resource.TestMatchResourceAttr(
						"data.ibm_compute_image_template.tfacc_img_tmpl",
						"id",
						regexp.MustCompile("^[0-9]+$"),
					),
				),
			},
		},
	})
}

const testAccCheckIBMComputeImageTemplateDataSourceConfig_basic = `
data "ibm_compute_image_template" "tfacc_img_tmpl" {
    name = "jumpbox"
}
`

const testAccCheckIBMComputeImageTemplateDataSourceConfig_basic2 = `
data "ibm_compute_image_template" "tfacc_img_tmpl" {
    name = "RightImage_Ubuntu_12.04_amd64_v13.5"
}
`

const testAccCheckIBMComputeImageTemplateDataSourceConfig_latest = `
data "ibm_compute_image_template" "tfacc_img_tmpl" {
    name = "25GB - Ubuntu / Ubuntu / 18.04-64 Minimal for VSI"
    latest = true
}
`

const testAccCheckIBMComputeImageTemplateDataSourceConfig_notlatest = `
data "ibm_compute_image_template" "tfacc_img_tmpl" {
    name = "25GB - Ubuntu / Ubuntu / 18.04-64 Minimal for VSI"
    latest = false
}
`
