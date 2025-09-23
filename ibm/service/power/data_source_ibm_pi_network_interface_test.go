// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package power_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/service/power"
)

func TestAccIBMPINetworkInterfaceDataSourceBasic(t *testing.T) {
	netIntData := "data.ibm_pi_network_interface.network_interface"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPINetworkInterfaceDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(netIntData, power.Attr_ID),
					resource.TestCheckResourceAttrSet(netIntData, power.Attr_IPAddress),
					resource.TestCheckResourceAttrSet(netIntData, power.Attr_MacAddress),
					resource.TestCheckResourceAttrSet(netIntData, power.Attr_Name),
					resource.TestCheckResourceAttrSet(netIntData, power.Attr_Status),
				),
			},
		},
	})
}

func testAccCheckIBMPINetworkInterfaceDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		data "ibm_pi_network_interface" "network_interface" {
			pi_cloud_instance_id = "%s"
			pi_network_id = "%s"
			pi_network_interface_id = "%s"
		}`, acc.Pi_cloud_instance_id, acc.Pi_network_id, acc.Pi_network_interface_id)
}
