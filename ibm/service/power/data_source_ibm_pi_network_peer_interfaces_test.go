// Copyright IBM Corp. 2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package power_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIBMPINetworkPeerInterfacesDataSourceBasic(t *testing.T) {
	netPeerIntercafes := "data.ibm_pi_network_peer_interfaces.network_peer_interfaces"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPINetworkPeerInterfacesDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(netPeerIntercafes, "id"),
				),
			},
		},
	})
}

func testAccCheckIBMPINetworkPeerInterfacesDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		data "ibm_pi_network_peer_interfaces" "network_peer_interfaces" {
			pi_cloud_instance_id = "%s"
		}`, acc.Pi_cloud_instance_id)
}
