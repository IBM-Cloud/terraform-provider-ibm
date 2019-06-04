package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccIBMISImageDataSource_basic(t *testing.T) {
	resName := "data.ibm_is_image.test1"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMISImageDataSourceConfig("ubuntu-18.04-amd64"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resName, "name", "ubuntu-18.04-amd64"),
					resource.TestCheckResourceAttrSet(resName, "os"),
					resource.TestCheckResourceAttrSet(resName, "architecture"),
					resource.TestCheckResourceAttrSet(resName, "visibility"),
					resource.TestCheckResourceAttrSet(resName, "status"),
				),
			},
		},
	})
}

func TestAccIBMISImageDataSource_With_Visibilty(t *testing.T) {
	resName := "data.ibm_is_image.test1"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMISImageDataSourceWithVisibility("centos-7.x-amd64", "public"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resName, "name", "centos-7.x-amd64"),
					resource.TestCheckResourceAttrSet(resName, "os"),
					resource.TestCheckResourceAttrSet(resName, "architecture"),
					resource.TestCheckResourceAttrSet(resName, "visibility"),
					resource.TestCheckResourceAttrSet(resName, "status"),
				),
			},
		},
	})
}

func testAccCheckIBMISImageDataSourceConfig(image string) string {
	return fmt.Sprintf(`

data "ibm_is_image" "test1" {
	name = "%s",
}`, image)
}

func testAccCheckIBMISImageDataSourceWithVisibility(image, visibility string) string {
	return fmt.Sprintf(`

data "ibm_is_image" "test1" {
	name = "%s",
	visibility = "%s"
}`, image, visibility)
}
