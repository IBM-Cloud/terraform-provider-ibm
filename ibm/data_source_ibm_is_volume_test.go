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

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccIBMISVolumeDatasource_basic(t *testing.T) {
	name := fmt.Sprintf("tf-vol-%d", acctest.RandIntRange(10, 100))
	zone := "us-south-1"
	resName := "data.ibm_is_volume.testacc_dsvol"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISVolumeDataSourceConfig(name, zone),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						resName, "name", name),
					resource.TestCheckResourceAttr(
						resName, "zone", zone),
				),
			},
		},
	})
}

func testAccCheckIBMISVolumeDataSourceConfig(name, zone string) string {
	return fmt.Sprintf(`
	resource "ibm_is_volume" "testacc_volume"{
		name = "%s"
		profile = "10iops-tier"
		zone = "%s"
	}
	data "ibm_is_volume" "testacc_dsvol" {
		name = ibm_is_volume.testacc_volume.name
	}`, name, zone)
}
