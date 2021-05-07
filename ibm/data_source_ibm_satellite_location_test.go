// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccSatelliteLocationDataSourceBasic(t *testing.T) {
	name := fmt.Sprintf("tf-satellitelocation-%d", acctest.RandIntRange(10, 100))
	managed_from := "wdc04"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{

			resource.TestStep{
				Config: testAccCheckSatelliteLocationDataSource(name, managed_from),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_satellite_location.location", "location", name),
					resource.TestCheckResourceAttr("ibm_satellite_location.location", "managed_from", managed_from),
				),
			},
		},
	})
}

func testAccCheckSatelliteLocationDataSource(name, managed_from string) string {
	return fmt.Sprintf(`

	resource "ibm_satellite_location" "location" {
		location      = "%s"
		managed_from  = "%s"
		description	  = "satellite service"	
		zones		  = ["us-east-1", "us-east-2", "us-east-3"]
		tags		  = ["env:dev"]
	}

    data "ibm_satellite_location" "test_location" {
		location              = ibm_satellite_location.location.id	
}`, name, managed_from)

}
