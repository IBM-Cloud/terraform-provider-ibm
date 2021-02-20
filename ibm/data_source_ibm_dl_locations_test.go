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
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"testing"
)

func TestAccIBMDLLocationsDataSource_basic(t *testing.T) {
	node := "data.ibm_dl_locations.test_dl_locations"
	offeringType := "dedicated"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMDLOfferingLocationsDataSourceConfig(offeringType),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(node, "locations.#"),
				),
			},
		},
	})
}

func testAccCheckIBMDLOfferingLocationsDataSourceConfig(offeringType string) string {
	return fmt.Sprintf(`
	   data "ibm_dl_locations" "test_dl_locations"{
		offering_type = "%s"
	 }

	  `, offeringType)
}
