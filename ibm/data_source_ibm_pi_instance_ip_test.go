/**
 * © Copyright IBM Corporation 2020. All Rights Reserved.
 *
 * Licensed under the Mozilla Public License, version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at https://mozilla.org/MPL/2.0/
 */

package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccIBMPIInstanceIPDataSource_basic(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMPIInstanceIPDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ibm_pi_instance_ip.testacc_ds_instance_ip", "pi_network_name", pi_network_name),
					resource.TestCheckResourceAttr("data.ibm_pi_instance_ip.testacc_ds_instance_ip", "pi_instance_name", pi_instance_name),
					resource.TestCheckResourceAttr("data.ibm_pi_instance_ip.testacc_ds_instance_ip", "pi_cloud_instance_id", pi_cloud_instance_id),
				),
			},
		},
	})
}

func testAccCheckIBMPIInstanceIPDataSourceConfig() string {
	return fmt.Sprintf(`
	resource "ibm_pi_network" "power_networks" {
		pi_cloud_instance_id = "%[3]s"
		pi_network_name      = "%[1]s"
		pi_network_type      = "pub-vlan"
	}
data "ibm_pi_instance_ip" "testacc_ds_instance_ip" {
    pi_network_name = "%[1]s"
    pi_instance_name = "%[2]s"
    pi_cloud_instance_id = "%[3]s"
}`, pi_network_name, pi_instance_name, pi_cloud_instance_id)

}
