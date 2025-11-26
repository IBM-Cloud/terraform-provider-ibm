// Copyright IBM Corp. 2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package power_test

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/IBM-Cloud/power-go-client/clients/instance"
	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
)

func TestAccIBMPINetworkPeerRouteFilterBasic(t *testing.T) {
	direction := "import"
	index := fmt.Sprintf("%d", acctest.RandIntRange(10, 100))
	prefix := "192.168.91.0/24"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMPINetworkPeerRouteFilterDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPINetworkPeerRouteFilterConfigBasic(direction, index, prefix),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMPINetworkPeerRouteFilterExists("ibm_pi_network_peer_route_filter.network_peer_route_filter"),
					resource.TestCheckResourceAttr("ibm_pi_network_peer_route_filter.network_peer_route_filter", "pi_network_peer_id", acc.Pi_network_peer_id),
					resource.TestCheckResourceAttr("ibm_pi_network_peer_route_filter.network_peer_route_filter", "pi_direction", direction),
					resource.TestCheckResourceAttr("ibm_pi_network_peer_route_filter.network_peer_route_filter", "pi_index", index),
					resource.TestCheckResourceAttr("ibm_pi_network_peer_route_filter.network_peer_route_filter", "pi_prefix", prefix),
				),
			},
		},
	})
}

func TestAccIBMPINetworkPeerRouteFilterAllArgs(t *testing.T) {
	ge := fmt.Sprintf("%d", acctest.RandIntRange(25, 27))
	le := fmt.Sprintf("%d", acctest.RandIntRange(28, 32))
	action := "allow"
	direction := "import"
	index := fmt.Sprintf("%d", acctest.RandIntRange(10, 100))
	prefix := "192.168.92.0/24"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMPINetworkPeerRouteFilterDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPINetworkPeerRouteFilterConfig(ge, le, action, direction, index, prefix),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMPINetworkPeerRouteFilterExists("ibm_pi_network_peer_route_filter.network_peer_route_filter"),
					resource.TestCheckResourceAttr("ibm_pi_network_peer_route_filter.network_peer_route_filter", "pi_network_peer_id", acc.Pi_network_peer_id),
					resource.TestCheckResourceAttr("ibm_pi_network_peer_route_filter.network_peer_route_filter", "pi_ge", ge),
					resource.TestCheckResourceAttr("ibm_pi_network_peer_route_filter.network_peer_route_filter", "pi_le", le),
					resource.TestCheckResourceAttr("ibm_pi_network_peer_route_filter.network_peer_route_filter", "pi_action", action),
					resource.TestCheckResourceAttr("ibm_pi_network_peer_route_filter.network_peer_route_filter", "pi_direction", direction),
					resource.TestCheckResourceAttr("ibm_pi_network_peer_route_filter.network_peer_route_filter", "pi_index", index),
					resource.TestCheckResourceAttr("ibm_pi_network_peer_route_filter.network_peer_route_filter", "pi_prefix", prefix),
				),
			},
		},
	})
}

func testAccCheckIBMPINetworkPeerRouteFilterConfigBasic(direction, index, prefix string) string {
	return fmt.Sprintf(`
		resource "ibm_pi_network_peer_route_filter" "network_peer_route_filter" {
			pi_cloud_instance_id = "%[1]s"
			pi_network_peer_id   = "%[2]s"
			pi_direction         = "%[3]s"
			pi_index             = "%[4]s"
			pi_prefix            = "%[5]s"
		}
	`, acc.Pi_cloud_instance_id, acc.Pi_network_peer_id, direction, index, prefix)
}

func testAccCheckIBMPINetworkPeerRouteFilterConfig(ge, le, action, direction, index, prefix string) string {
	return fmt.Sprintf(`

		resource "ibm_pi_network_peer_route_filter" "network_peer_route_filter" {
			pi_cloud_instance_id = "%[1]s"
			pi_network_peer_id   = "%[2]s"
			pi_ge          		 = "%[3]s"
			pi_le 				 = "%[4]s"
			pi_action 			 = "%[5]s"
			pi_direction 		 = "%[6]s"
			pi_index 		 	 = "%[7]s"
			pi_prefix 			 = "%[8]s"
		}
	`, acc.Pi_cloud_instance_id, acc.Pi_network_peer_id, ge, le, action, direction, index, prefix)
}

func testAccCheckIBMPINetworkPeerRouteFilterExists(n string) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}
		if rs.Primary.ID == "" {
			return errors.New("No Record ID is set")
		}
		sess, err := acc.TestAccProvider.Meta().(conns.ClientSession).IBMPISession()
		if err != nil {
			return err
		}
		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}
		networkC := instance.NewIBMPINetworkPeerClient(context.Background(), sess, parts[0])

		_, err = networkC.GetNetworkPeersRouteFilter(parts[1], parts[2])
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccCheckIBMPINetworkPeerRouteFilterDestroy(s *terraform.State) error {
	sess, err := acc.TestAccProvider.Meta().(conns.ClientSession).IBMPISession()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_pi_network_peer_route_filter" {
			continue
		}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}
		networkC := instance.NewIBMPINetworkPeerClient(context.Background(), sess, parts[0])
		_, err = networkC.GetNetworkPeersRouteFilter(parts[1], parts[2])
		if err == nil {
			return fmt.Errorf("Network peer %s route filter still exists: %s", parts[1], parts[2])
		}
	}

	return nil
}
