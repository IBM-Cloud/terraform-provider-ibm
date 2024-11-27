// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIBMIsVirtualNetworkInterfacesDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsVirtualNetworkInterfacesDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_virtual_network_interfaces.is_virtual_network_interfaces", "virtual_network_interfaces.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_virtual_network_interfaces.is_virtual_network_interfaces", "virtual_network_interfaces.0.auto_delete"),
					resource.TestCheckResourceAttrSet("data.ibm_is_virtual_network_interfaces.is_virtual_network_interfaces", "virtual_network_interfaces.0.created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_is_virtual_network_interfaces.is_virtual_network_interfaces", "virtual_network_interfaces.0.crn"),
					resource.TestCheckResourceAttrSet("data.ibm_is_virtual_network_interfaces.is_virtual_network_interfaces", "virtual_network_interfaces.0.href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_virtual_network_interfaces.is_virtual_network_interfaces", "virtual_network_interfaces.0.id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_virtual_network_interfaces.is_virtual_network_interfaces", "virtual_network_interfaces.0.lifecycle_state"),
					resource.TestCheckResourceAttrSet("data.ibm_is_virtual_network_interfaces.is_virtual_network_interfaces", "virtual_network_interfaces.0.name"),
					resource.TestCheckResourceAttrSet("data.ibm_is_virtual_network_interfaces.is_virtual_network_interfaces", "virtual_network_interfaces.0.resource_group.0.id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_virtual_network_interfaces.is_virtual_network_interfaces", "virtual_network_interfaces.0.resource_group.0.href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_virtual_network_interfaces.is_virtual_network_interfaces", "virtual_network_interfaces.0.resource_group.0.name"),
					resource.TestCheckResourceAttrSet("data.ibm_is_virtual_network_interfaces.is_virtual_network_interfaces", "virtual_network_interfaces.0.resource_type"),
					resource.TestCheckResourceAttrSet("data.ibm_is_virtual_network_interfaces.is_virtual_network_interfaces", "virtual_network_interfaces.0.security_groups.0.id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_virtual_network_interfaces.is_virtual_network_interfaces", "virtual_network_interfaces.0.subnet.0.id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_virtual_network_interfaces.is_virtual_network_interfaces", "virtual_network_interfaces.0.vpc.0.id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_virtual_network_interfaces.is_virtual_network_interfaces", "virtual_network_interfaces.0.zone.0.name"),
				),
			},
		},
	})
}
func TestAccIBMIsVirtualNetworkInterfacesDataSourceVniBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsVirtualNetworkInterfacesDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_virtual_network_interfaces.is_virtual_network_interfaces", "virtual_network_interfaces.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_virtual_network_interfaces.is_virtual_network_interfaces", "virtual_network_interfaces.0.auto_delete"),
					resource.TestCheckResourceAttrSet("data.ibm_is_virtual_network_interfaces.is_virtual_network_interfaces", "virtual_network_interfaces.0.created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_is_virtual_network_interfaces.is_virtual_network_interfaces", "virtual_network_interfaces.0.crn"),
					resource.TestCheckResourceAttrSet("data.ibm_is_virtual_network_interfaces.is_virtual_network_interfaces", "virtual_network_interfaces.0.href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_virtual_network_interfaces.is_virtual_network_interfaces", "virtual_network_interfaces.0.id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_virtual_network_interfaces.is_virtual_network_interfaces", "virtual_network_interfaces.0.lifecycle_state"),
					resource.TestCheckResourceAttrSet("data.ibm_is_virtual_network_interfaces.is_virtual_network_interfaces", "virtual_network_interfaces.0.name"),
					resource.TestCheckResourceAttrSet("data.ibm_is_virtual_network_interfaces.is_virtual_network_interfaces", "virtual_network_interfaces.0.resource_group.0.id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_virtual_network_interfaces.is_virtual_network_interfaces", "virtual_network_interfaces.0.resource_group.0.href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_virtual_network_interfaces.is_virtual_network_interfaces", "virtual_network_interfaces.0.resource_group.0.name"),
					resource.TestCheckResourceAttrSet("data.ibm_is_virtual_network_interfaces.is_virtual_network_interfaces", "virtual_network_interfaces.0.resource_type"),
					resource.TestCheckResourceAttrSet("data.ibm_is_virtual_network_interfaces.is_virtual_network_interfaces", "virtual_network_interfaces.0.security_groups.0.id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_virtual_network_interfaces.is_virtual_network_interfaces", "virtual_network_interfaces.0.subnet.0.id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_virtual_network_interfaces.is_virtual_network_interfaces", "virtual_network_interfaces.0.vpc.0.id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_virtual_network_interfaces.is_virtual_network_interfaces", "virtual_network_interfaces.0.zone.0.name"),
					resource.TestCheckResourceAttrSet("data.ibm_is_virtual_network_interfaces.is_virtual_network_interfaces", "virtual_network_interfaces.0.allow_ip_spoofing"),
					resource.TestCheckResourceAttrSet("data.ibm_is_virtual_network_interfaces.is_virtual_network_interfaces", "virtual_network_interfaces.0.enable_infrastructure_nat"),
				),
			},
		},
	})
}

func testAccCheckIBMIsVirtualNetworkInterfacesDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		data "ibm_is_virtual_network_interfaces" "is_virtual_network_interfaces" {
		}
	`)
}
