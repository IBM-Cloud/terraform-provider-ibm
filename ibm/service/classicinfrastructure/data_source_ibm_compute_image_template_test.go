// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package classicinfrastructure_test

import (
	"regexp"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccIBMComputeImageTemplateDataSource_Basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
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
