/*
* Licensed Materials - Property of IBM
* (C) Copyright IBM Corp. 2017. All Rights Reserved.
* US Government Users Restricted Rights - Use, duplication or
* disclosure restricted by GSA ADP Schedule Contract with IBM Corp.
 */

package ibm

import (
	"github.com/hashicorp/terraform/helper/resource"
	"testing"
)

func TestAccIBMNetworkVlan_Basic(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMNetworkVlanConfig_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_network_vlan.test_vlan", "name", "test_vlan"),
					resource.TestCheckResourceAttr(
						"ibm_network_vlan.test_vlan", "datacenter", "lon02"),
					resource.TestCheckResourceAttr(
						"ibm_network_vlan.test_vlan", "type", "PUBLIC"),
					resource.TestCheckResourceAttr(
						"ibm_network_vlan.test_vlan", "softlayer_managed", "false"),
					resource.TestCheckResourceAttr(
						"ibm_network_vlan.test_vlan", "router_hostname", "fcr01a.lon02"),
					resource.TestCheckResourceAttr(
						"ibm_network_vlan.test_vlan", "subnet_size", "8"),
				),
			},

			resource.TestStep{
				Config: testAccCheckIBMNetworkVlanConfig_name_update,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_network_vlan.test_vlan", "name", "test_vlan_update"),
				),
			},
		},
	})
}

const testAccCheckIBMNetworkVlanConfig_basic = `
resource "ibm_network_vlan" "test_vlan" {
   name = "test_vlan"
   datacenter = "lon02"
   type = "PUBLIC"
   subnet_size = 8
   router_hostname = "fcr01a.lon02"
}`

const testAccCheckIBMNetworkVlanConfig_name_update = `
resource "ibm_network_vlan" "test_vlan" {
   name = "test_vlan_update"
   datacenter = "lon02"
   type = "PUBLIC"
   subnet_size = 8
   router_hostname = "fcr01a.lon02"
}`
