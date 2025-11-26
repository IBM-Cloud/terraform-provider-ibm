// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"errors"
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"

	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccIBMISImageDeprecate_basic(t *testing.T) {
	var image string
	name := fmt.Sprintf("tfimg-name-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheckImage(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: checkImageDeprecateDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISImageDeprecateConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISImageDeprecateExists("ibm_is_image.isExampleImage", image),
					resource.TestCheckResourceAttr(
						"ibm_is_image.isExampleImage", "name", name),
					resource.TestCheckResourceAttr(
						"ibm_is_image_deprecate.deprecate", "status", "deprecated"),
				),
			},
		},
	})
}

func checkImageDeprecateDestroy(s *terraform.State) error {

	sess, _ := acc.TestAccProvider.Meta().(conns.ClientSession).VpcV1API()
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_is_image" {
			continue
		}

		getimgoptions := &vpcv1.GetImageOptions{
			ID: &rs.Primary.ID,
		}
		_, _, err := sess.GetImage(getimgoptions)
		if err == nil {
			return fmt.Errorf("Image still exists: %s", rs.Primary.ID)
		}
	}

	return nil
}

func testAccCheckIBMISImageDeprecateExists(n, image string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return errors.New("No Record ID is set")
		}

		sess, _ := acc.TestAccProvider.Meta().(conns.ClientSession).VpcV1API()
		getimgoptions := &vpcv1.GetImageOptions{
			ID: &rs.Primary.ID,
		}
		foundImage, _, err := sess.GetImage(getimgoptions)
		if err != nil {
			return err
		}
		image = *foundImage.ID

		return nil
	}
}

func testAccCheckIBMISImageDeprecateConfig(name string) string {
	return fmt.Sprintf(`
		resource "ibm_is_image" "isExampleImage" {
			href = "%s"
			name = "%s"
			operating_system = "%s"
		}
		resource ibm_is_image_deprecate deprecate {
			image = ibm_is_image.isExampleImage.id
		}

	`, acc.Image_cos_url, name, acc.Image_operating_system)
}
