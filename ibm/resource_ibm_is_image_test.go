package ibm

import (
	"errors"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.ibm.com/ibmcloud/vpc-go-sdk/vpcclassicv1"
	"github.ibm.com/ibmcloud/vpc-go-sdk/vpcv1"
)

func TestAccIBMISImage_basic(t *testing.T) {
	var image string
	name := fmt.Sprintf("tfimg-name-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheckImage(t) },
		Providers:    testAccProviders,
		CheckDestroy: checkImageDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMISImageConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISImageExists("ibm_is_image.isExampleImage", image),
					resource.TestCheckResourceAttr(
						"ibm_is_image.isExampleImage", "name", name),
				),
			},
		},
	})
}

func checkImageDestroy(s *terraform.State) error {
	userDetails, _ := testAccProvider.Meta().(ClientSession).BluemixUserDetails()

	if userDetails.generation == 1 {
		sess, _ := testAccProvider.Meta().(ClientSession).VpcClassicV1API()
		for _, rs := range s.RootModule().Resources {
			if rs.Type != "ibm_is_image" {
				continue
			}

			getimgoptions := &vpcclassicv1.GetImageOptions{
				ID: &rs.Primary.ID,
			}
			_, _, err := sess.GetImage(getimgoptions)
			if err == nil {
				return fmt.Errorf("Image still exists: %s", rs.Primary.ID)
			}
		}
	} else {
		sess, _ := testAccProvider.Meta().(ClientSession).VpcV1API()
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
	}
	return nil
}

func testAccCheckIBMISImageExists(n, image string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		fmt.Println("siv ", s.RootModule().Resources)
		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return errors.New("No Record ID is set")
		}
		userDetails, _ := testAccProvider.Meta().(ClientSession).BluemixUserDetails()

		if userDetails.generation == 1 {
			sess, _ := testAccProvider.Meta().(ClientSession).VpcClassicV1API()
			getimgoptions := &vpcclassicv1.GetImageOptions{
				ID: &rs.Primary.ID,
			}
			foundImage, _, err := sess.GetImage(getimgoptions)
			if err != nil {
				return err
			}
			image = *foundImage.ID
		} else {
			sess, _ := testAccProvider.Meta().(ClientSession).VpcV1API()
			getimgoptions := &vpcv1.GetImageOptions{
				ID: &rs.Primary.ID,
			}
			foundImage, _, err := sess.GetImage(getimgoptions)
			if err != nil {
				return err
			}
			image = *foundImage.ID
		}
		return nil
	}
}

func testAccCheckIBMISImageConfig(name string) string {
	return fmt.Sprintf(`
		resource "ibm_is_image" "isExampleImage" {
			href = "%s"
			name = "%s"
			operating_system = "%s"
		}
	`, image_cos_url, name, image_operating_system)
}
