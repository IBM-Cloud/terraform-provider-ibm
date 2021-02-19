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

func TestAccIBMTransitGatewaysLocationsDataSource_basic(t *testing.T) {
	resName := "data.ibm_tg_locations.test_tg_locations"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMTransitGatewaysLocationsDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resName, "locations.0.name"),
					resource.TestCheckResourceAttrSet(resName, "locations.0.type"),
					resource.TestCheckResourceAttrSet(resName, "locations.0.billing_location"),
				),
			},
		},
	})
}

func testAccCheckIBMTransitGatewaysLocationsDataSourceConfig() string {
	// status filter defaults to empty
	return fmt.Sprintf(`
	data "ibm_tg_locations" "test_tg_locations" {
		}   `)
}
