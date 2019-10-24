package ibm

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
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
