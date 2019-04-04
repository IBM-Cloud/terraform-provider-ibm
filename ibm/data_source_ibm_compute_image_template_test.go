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
					resource.TestCheckNoResourceAttr(
						"data.ibm_compute_image_template.tfacc_img_tmpl",
						"most_recent",
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
					resource.TestCheckNoResourceAttr(
						"data.ibm_compute_image_template.tfacc_img_tmpl",
						"most_recent",
					),
					resource.TestMatchResourceAttr(
						"data.ibm_compute_image_template.tfacc_img_tmpl",
						"id",
						regexp.MustCompile("^[0-9]+$"),
					),
				),
			},
			// Tests looking up the most_recent of a public image
			{
				Config: testAccCheckIBMComputeImageTemplateDataSourceConfig_mostrecent,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"data.ibm_compute_image_template.tfacc_img_tmpl",
						"name",
						"25GB - Ubuntu / Ubuntu / 18.04-64 Minimal for VSI",
					),
					resource.TestCheckResourceAttr(
						"data.ibm_compute_image_template.tfacc_img_tmpl",
						"most_recent",
						"true",
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
				Config: testAccCheckIBMComputeImageTemplateDataSourceConfig_mostrecent2,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"data.ibm_compute_image_template.tfacc_img_tmpl",
						"name",
						"25GB - Ubuntu / Ubuntu / 18.04-64 Minimal for VSI",
					),
					resource.TestCheckResourceAttr(
						"data.ibm_compute_image_template.tfacc_img_tmpl",
						"most_recent",
						"false",
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

const testAccCheckIBMComputeImageTemplateDataSourceConfig_mostrecent = `
data "ibm_compute_image_template" "tfacc_img_tmpl" {
    name = "25GB - Ubuntu / Ubuntu / 18.04-64 Minimal for VSI"
    most_recent = true
}
`

const testAccCheckIBMComputeImageTemplateDataSourceConfig_mostrecent2 = `
data "ibm_compute_image_template" "tfacc_img_tmpl" {
    name = "25GB - Ubuntu / Ubuntu / 18.04-64 Minimal for VSI"
    most_recent = false
}
`
