// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMSatelliteAttachHostScriptDataSourceBasic(t *testing.T) {
	locationName := fmt.Sprintf("tf-satellitelocation-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMSatelliteAttachHostScriptDataSourceConfig(locationName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ibm_satellite_attach_host_script.script", "host_provider", "ibm"),
				),
			},
		},
	})
}

func testAccCheckIBMSatelliteAttachHostScriptDataSourceConfig(locationName string) string {
	return fmt.Sprintf(`
resource "ibm_satellite_location" "testacc_satellite" {
	location     = "%s"
	managed_from = "wdc04"
	zones		 = ["us-east-1", "us-east-2", "us-east-3"]
}

data "ibm_satellite_attach_host_script" "script" {
	location       = ibm_satellite_location.testacc_satellite.id
	labels         = ["env:prod"]
	host_provider  = "ibm"
}`, locationName)
}
