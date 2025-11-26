// Copyright IBM Corp. 2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package power_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIBMPINetworkPeerDataSourceBasic(t *testing.T) {
	networkRes := "data.ibm_pi_network_peer.network_peer"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPINetworkPeerDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(networkRes, "id"),
				),
			},
		},
	})
}

func testAccCheckIBMPINetworkPeerDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		data "ibm_pi_network_peer" "network_peer" {
			pi_cloud_instance_id = "%[1]s"
			pi_network_peer_id   = "%[2]s"
		}`, acc.Pi_cloud_instance_id, acc.Pi_network_peer_id)
}
