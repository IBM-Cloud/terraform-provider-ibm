// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

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
			{
				Config: testAccCheckIBMNetworkVlanSpanOnConfigBasic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_network_vlan_spanning.test_vlan", "vlan_spanning", "on"),
				),
			},
			{
				Config: testAccCheckIBMNetworkVlanSpanOffConfigBasic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_network_vlan_spanning.test_vlan", "vlan_spanning", "off"),
				),
			},
		},
	})
}

const testAccCheckIBMNetworkVlanSpanOnConfigBasic = `
resource "ibm_network_vlan_spanning" "test_vlan" {
   vlan_spanning = "on"
}`
const testAccCheckIBMNetworkVlanSpanOffConfigBasic = `
resource "ibm_network_vlan_spanning" "test_vlan" {
   vlan_spanning = "off"
}`
