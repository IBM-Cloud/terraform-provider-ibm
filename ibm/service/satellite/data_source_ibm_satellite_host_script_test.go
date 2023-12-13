// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package satellite_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMSatelliteAttachHostScriptDataSourceBasic(t *testing.T) {
	locationName := fmt.Sprintf("tf-satellitelocation-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
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

// test coreos-enabled locations
func TestAccIBMSatelliteAttachHostScriptDataSourceBasicCoreos(t *testing.T) {
	locationName := fmt.Sprintf("tf-satellitelocation-coreos-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMSatelliteAttachHostScriptDataSourceConfigCoreos(locationName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ibm_satellite_attach_host_script.script", "host_provider", "ibm"),
					resource.TestCheckResourceAttr("data.ibm_satellite_attach_host_script.script", "host_link_agent_endpoint", "testendpoint"),
				),
			},
		},
	})
}

func testAccCheckIBMSatelliteAttachHostScriptDataSourceConfigCoreos(locationName string) string {
	return fmt.Sprintf(`
resource "ibm_satellite_location" "testacc_satellite" {
	location     = "%s"
	managed_from = "wdc04"
	coreos_enabled = true
	zones		 = ["us-east-1", "us-east-2", "us-east-3"]
}

data "ibm_satellite_attach_host_script" "script" {
	location       = ibm_satellite_location.testacc_satellite.id
	labels         = ["env:prod"]
	coreos_host	   = true
	host_provider  = "ibm"
	host_link_agent_endpoint = "testendpoint"
}`, locationName)
}
