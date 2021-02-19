/*
* IBM Confidential
* Object Code Only Source Materials
* 5747-SM3
* (c) Copyright IBM Corp. 2017,2021
*
* The source code for this program is not published or otherwise divested
* of its trade secrets, irrespective of what has been deposited with the
* U.S. Copyright Office.
 */

package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccIBMPICatalogImagesDataSource_basic(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPICatalogImagesDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_pi_catalog_images.power_catalog_images", "id"),
				),
			},
		},
	})
}

func testAccCheckIBMPICatalogImagesDataSourceConfig() string {
	return fmt.Sprintf(`
	
data "ibm_pi_catalog_images" "power_catalog_images" {
    pi_cloud_instance_id = "%s"
}`, pi_cloud_instance_id)

}
