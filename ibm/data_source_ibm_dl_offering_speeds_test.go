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

func TestAccIBMDLOfferingSpeedsDataSource_basic(t *testing.T) {
	node1 := "data.ibm_dl_offering_speeds.test1"
	node2 := "data.ibm_dl_offering_speeds.test2"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMDLOfferingSpeedsDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(node1, "offering_speeds.0.link_speed"),
					resource.TestCheckResourceAttrSet(node2, "offering_speeds.0.link_speed"),
				),
			},
		},
	})
}

func testAccCheckIBMDLOfferingSpeedsDataSourceConfig() string {
	return fmt.Sprintf(`
	data "ibm_dl_offering_speeds" "test1" {
		offering_type = "dedicated"
	}

	data "ibm_dl_offering_speeds" "test2" {
		offering_type = "connect"
  	}
	`)
}
