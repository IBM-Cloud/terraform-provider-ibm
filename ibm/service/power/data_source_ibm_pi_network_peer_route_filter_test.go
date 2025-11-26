// Copyright IBM Corp. 2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package power_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIBMPINetworkPeerRouteFilterDataSourceBasic(t *testing.T) {
	networkPeerRouteFilter := "data.ibm_pi_network_peer_route_filter.network_peer_route_filter"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPINetworkPeerRouteFilterDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(networkPeerRouteFilter, "id"),
				),
			},
		},
	})
}

func testAccCheckIBMPINetworkPeerRouteFilterDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		data "ibm_pi_network_peer_route_filter" "network_peer_route_filter" {
			pi_cloud_instance_id = "%s"
			pi_network_peer_id   = "%s"
			pi_route_filter_id   = "%s"
		}`, acc.Pi_cloud_instance_id, acc.Pi_network_peer_id, acc.Pi_route_filter_id)
}
