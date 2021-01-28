package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccIBMPIImagesDataSource_basic(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPIImagesDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_pi_images.testacc_ds_image", "id"),
				),
			},
		},
	})
}

func testAccCheckIBMPIImagesDataSourceConfig() string {
	return fmt.Sprintf(`
	resource "ibm_pi_image" "power_image" {
		pi_image_name       = "7100-05-04"
		pi_image_id         = "d469355f-effa-4c5d-9c85-33338d6c3789"
		pi_cloud_instance_id = "%[1]s"
	  }
	data "ibm_pi_images" "testacc_ds_image" {
		pi_cloud_instance_id = "%[1]s"
	}`, pi_cloud_instance_id)

}
