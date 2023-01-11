// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIBMIsVirtualNetworkInterfaceDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsVirtualNetworkInterfaceDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_virtual_network_interface.is_virtual_network_interface", "virtual_network_interface"),
					resource.TestCheckResourceAttrSet("data.ibm_is_virtual_network_interface.is_virtual_network_interface", "auto_delete"),
					resource.TestCheckResourceAttrSet("data.ibm_is_virtual_network_interface.is_virtual_network_interface", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_is_virtual_network_interface.is_virtual_network_interface", "crn"),
					resource.TestCheckResourceAttrSet("data.ibm_is_virtual_network_interface.is_virtual_network_interface", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_virtual_network_interface.is_virtual_network_interface", "lifecycle_state"),
					resource.TestCheckResourceAttrSet("data.ibm_is_virtual_network_interface.is_virtual_network_interface", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_is_virtual_network_interface.is_virtual_network_interface", "resource_group.0.id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_virtual_network_interface.is_virtual_network_interface", "resource_type"),
					resource.TestCheckResourceAttrSet("data.ibm_is_virtual_network_interface.is_virtual_network_interface", "security_groups.0.id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_virtual_network_interface.is_virtual_network_interface", "subnet.0.id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_virtual_network_interface.is_virtual_network_interface", "vpc.0.id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_virtual_network_interface.is_virtual_network_interface", "zone.0.name"),
				),
			},
		},
	})
}

func testAccCheckIBMIsVirtualNetworkInterfaceDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		data "ibm_is_virtual_network_interface" "is_virtual_network_interface" {
			virtual_network_interface = "%s"
		}
	`, acc.VNIId)
}
