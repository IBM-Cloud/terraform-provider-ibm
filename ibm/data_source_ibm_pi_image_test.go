/* IBM Confidential
*  Object Code Only Source Materials
*  5747-SM3
*  (c) Copyright IBM Corp. 2017,2021
*
*  The source code for this program is not published or otherwise divested
*  of its trade secrets, irrespective of what has been deposited with the
*  U.S. Copyright Office. */

package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccIBMPIImageDataSource_basic(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPIImageDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_pi_image.testacc_ds_image", "id"),
				),
			},
		},
	})
}

func testAccCheckIBMPIImageDataSourceConfig() string {
	return fmt.Sprintf(`
	resource "ibm_pi_image" "power_image" {
		pi_image_name       = "7200-04-01"
		pi_image_id         = "f31da27a-b634-45e5-913a-3f4d964e5a02"
		pi_cloud_instance_id = "%[1]s"
	  }
	data "ibm_pi_image" "testacc_ds_image" {
		pi_image_name = ibm_pi_image.power_image.image_id
		pi_cloud_instance_id = "%[1]s"
	}`, pi_cloud_instance_id)

}
