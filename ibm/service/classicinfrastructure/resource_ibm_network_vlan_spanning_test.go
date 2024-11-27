// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package classicinfrastructure_test

import (
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccIBMNetworkVlanSpan_Basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
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
