// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package classicinfrastructure_test

import (
	"fmt"
	"regexp"
	"testing"

	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMNetworkVlanDataSource_Basic(t *testing.T) {

	name := fmt.Sprintf("terraformuat_vlan_%s", acctest.RandString(2))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMNetworkVlanDataSourceConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMResources("data.ibm_network_vlan.tfacc_vlan", "number",
						"ibm_network_vlan.test_vlan_private", "vlan_number"),
					//resource.TestCheckResourceAttr("data.ibm_network_vlan.tfacc_vlan", "number", number),
					resource.TestCheckResourceAttr("data.ibm_network_vlan.tfacc_vlan", "name", name),
					resource.TestMatchResourceAttr("data.ibm_network_vlan.tfacc_vlan", "id", regexp.MustCompile("^[0-9]+$")),
					resource.TestCheckResourceAttr("data.ibm_network_vlan.tfacc_vlan", "subnets.#", "1"),
				),
			},
		},
	})
}

func testAccCheckIBMNetworkVlanDataSourceConfig(name string) string {
	return fmt.Sprintf(`
    resource "ibm_network_vlan" "test_vlan_private" {
    name            = "%s"
    datacenter      = "dal06"
    type            = "PRIVATE"
    
}
data "ibm_network_vlan" "tfacc_vlan" {
    number = "${ibm_network_vlan.test_vlan_private.vlan_number}"
    name = "${ibm_network_vlan.test_vlan_private.name}"
}`, name)
}
