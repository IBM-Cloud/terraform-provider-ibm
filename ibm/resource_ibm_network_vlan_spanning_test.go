/**
 * © Copyright IBM Corporation 2020. All Rights Reserved.
 *
 * Licensed under the Mozilla Public License, version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at https://mozilla.org/MPL/2.0/
 */

package ibm

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccIBMNetworkVlanSpan_Basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMNetworkVlanSpanOnConfig_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_network_vlan_spanning.test_vlan", "vlan_spanning", "on"),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMNetworkVlanSpanOffConfig_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_network_vlan_spanning.test_vlan", "vlan_spanning", "off"),
				),
			},
		},
	})
}

const testAccCheckIBMNetworkVlanSpanOnConfig_basic = `
resource "ibm_network_vlan_spanning" "test_vlan" {
   "vlan_spanning" = "on"
}`
const testAccCheckIBMNetworkVlanSpanOffConfig_basic = `
resource "ibm_network_vlan_spanning" "test_vlan" {
   "vlan_spanning" = "off"
}`
