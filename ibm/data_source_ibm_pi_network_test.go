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

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccIBMPINetworkDataSource_basic(t *testing.T) {
	name := fmt.Sprintf("tf-pi-network-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPINetworkDataSourceConfig(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_pi_network.testacc_ds_network", "id"),
				),
			},
		},
	})
}

func testAccCheckIBMPINetworkDataSourceConfig(name string) string {
	return testAccCheckIBMPINetworkConfig(name) + fmt.Sprintf(`
data "ibm_pi_network" "testacc_ds_network" {
    pi_network_name = ibm_pi_network.power_networks.network_id
    pi_cloud_instance_id = "%s"
}`, pi_cloud_instance_id)

}
