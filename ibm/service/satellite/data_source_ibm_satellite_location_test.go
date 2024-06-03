// Copyright IBM Corp. 2017, 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package satellite_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccSatelliteLocationDataSourceBasic(t *testing.T) {
	name := fmt.Sprintf("tf-satellitelocation-%d", acctest.RandIntRange(10, 100))
	managed_from := "wdc04"
	physical_address := "test-road 10, 111 test-country"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{

			{
				Config: testAccCheckSatelliteLocationDataSource(name, managed_from, physical_address),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_satellite_location.location", "location", name),
					resource.TestCheckResourceAttr("ibm_satellite_location.location", "managed_from", managed_from),
					resource.TestCheckResourceAttr("ibm_satellite_location.location", "physical_address", physical_address),
				),
			},
		},
	})
}

func testAccCheckSatelliteLocationDataSource(name, managed_from string, physical_address string) string {
	return fmt.Sprintf(`

	resource "ibm_satellite_location" "location" {
		location      = "%s"
		managed_from  = "%s"
		physical_address = "%s"
		description	  = "satellite service"	
		zones		  = ["us-east-1", "us-east-2", "us-east-3"]
		tags		  = ["env:dev"]
	}

    data "ibm_satellite_location" "test_location" {
		location              = ibm_satellite_location.location.id	
}`, name, managed_from, physical_address)

}
