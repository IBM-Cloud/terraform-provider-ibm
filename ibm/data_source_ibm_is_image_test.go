package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccIBMISImageDataSource_basic(t *testing.T) {
	resName := "data.ibm_is_image.test1"
	imageName := fmt.Sprintf("tfimage-name-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMISImageDataSourceConfig(imageName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resName, "name", imageName),
					resource.TestCheckResourceAttrSet(resName, "os"),
					resource.TestCheckResourceAttrSet(resName, "architecture"),
					resource.TestCheckResourceAttrSet(resName, "visibility"),
					resource.TestCheckResourceAttrSet(resName, "status"),
				),
			},
		},
	})
}

func TestAccIBMISImageDataSource_With_VisibiltyPublic(t *testing.T) {
	resName := "data.ibm_is_image.test1"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMISImageDataSourceWithVisibilityPublic("public"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resName, "name", IsImageName),
					resource.TestCheckResourceAttrSet(resName, "os"),
					resource.TestCheckResourceAttrSet(resName, "architecture"),
					resource.TestCheckResourceAttrSet(resName, "visibility"),
					resource.TestCheckResourceAttrSet(resName, "status"),
				),
			},
		},
	})
}

func TestAccIBMISImageDataSource_With_VisibiltyPrivate(t *testing.T) {
	resName := "data.ibm_is_image.test1"
	imageName := fmt.Sprintf("tfimage-name-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMISImageDataSourceWithVisibilityPrivate(imageName, "private"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resName, "name", imageName),
					resource.TestCheckResourceAttrSet(resName, "os"),
					resource.TestCheckResourceAttrSet(resName, "architecture"),
					resource.TestCheckResourceAttrSet(resName, "visibility"),
					resource.TestCheckResourceAttrSet(resName, "status"),
				),
			},
		},
	})
}

func testAccCheckIBMISImageDataSourceConfig(imageName string) string {
	return fmt.Sprintf(`
	resource "ibm_is_image" "isExampleImage" {
		href = "%s"
		name = "%s"
		operating_system = "%s"
	}
	data "ibm_is_image" "test1" {
		name = ibm_is_image.isExampleImage.name
	}`, image_cos_url, imageName, image_operating_system)
}

func testAccCheckIBMISImageDataSourceWithVisibilityPublic(visibility string) string {
	return fmt.Sprintf(`
	data "ibm_is_image" "test1" {
		name = "%s"
		visibility = "%s"
	}`, IsImageName, visibility)
}

func testAccCheckIBMISImageDataSourceWithVisibilityPrivate(imageName, visibility string) string {
	return fmt.Sprintf(`
	resource "ibm_is_image" "isExampleImage" {
		href = "%s"
		name = "%s"
		operating_system = "%s"
	}
	data "ibm_is_image" "test1" {
		name = ibm_is_image.isExampleImage.name
		visibility = "%s"
	}`, image_cos_url, imageName, image_operating_system, visibility)
}
