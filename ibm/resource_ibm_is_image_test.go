package ibm

import (
	"errors"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.ibm.com/Bluemix/riaas-go-client/clients/compute"
	"github.ibm.com/Bluemix/riaas-go-client/riaas/models"
)

func TestAccIBMISImage_basic(t *testing.T) {
	var image *models.Image

	href := fmt.Sprintf("cos://us-south/vnfimages/BIGIP-15.0.1-0.0.11.qcow2")
	operatingsystem := "centos-7-amd64"
	name := fmt.Sprintf("terraformimageuat-create-step-name-%d", acctest.RandInt())
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheckImage(t) },
		Providers:    testAccProviders,
		CheckDestroy: checkImageDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMISImageConfig(href, name, operatingsystem),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISImageExists("ibm_is_image.isExampleImage", &image),
					resource.TestCheckResourceAttr(
						"ibm_is_image.isExampleImage", "name", name),
				),
			},
		},
	})
}

func checkImageDestroy(s *terraform.State) error {
	sess, _ := testAccProvider.Meta().(ClientSession).ISSession()

	imageC := compute.NewImageClient(sess)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_is_image" {
			continue
		}

		_, err := imageC.Get(rs.Primary.ID)

		if err == nil {
			return fmt.Errorf("Image still exists: %s", rs.Primary.ID)
		}
	}

	return nil
}

func testAccCheckIBMISImageExists(n string, image **models.Image) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		fmt.Println("siv ", s.RootModule().Resources)
		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return errors.New("No Record ID is set")
		}

		sess, _ := testAccProvider.Meta().(ClientSession).ISSession()
		imageC := compute.NewImageClient(sess)
		foundImage, err := imageC.Get(rs.Primary.ID)

		if err != nil {
			return err
		}

		*image = foundImage
		return nil
	}
}

func testAccCheckIBMISImageConfig(href, name, operatingsystem string) string {
	return fmt.Sprintf(`
		resource "ibm_is_image" "isExampleImage" {
			href = "%s"
			name = "%s"
			operating_system = "%s"
		}
	`, image_cos_url, name, image_operating_system)
}
