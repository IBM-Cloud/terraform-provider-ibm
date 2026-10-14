// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/vpc"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/stretchr/testify/assert"
)

func TestAccIBMIsDynamicRouteServerPeerGroupPeersDataSourceBasic(t *testing.T) {
	vpcname := fmt.Sprintf("tf-vpc-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tf-subnet-%d", acctest.RandIntRange(10, 100))
	vniname := fmt.Sprintf("tf-vni-%d", acctest.RandIntRange(10, 100))
	drsname := fmt.Sprintf("tf-drs-%d", acctest.RandIntRange(10, 100))
	pgname := fmt.Sprintf("tf-peer-group-%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf-peer-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsDynamicRouteServerPeerGroupPeersDataSourceConfigBasic(vpcname, subnetname, vniname, drsname, pgname, name),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_peer_group_peers.testacc_peers", "id"),
					resource.TestCheckResourceAttrPair("data.ibm_is_dynamic_route_server_peer_group_peers.testacc_peers", "dynamic_route_server_peer_group_id", isDynamicRouteServerPeerGroupResource, "dynamic_route_server_peer_group_id"),
					// One peer was created with the peer group and one with the peer resource.
					resource.TestCheckResourceAttr("data.ibm_is_dynamic_route_server_peer_group_peers.testacc_peers", "peers.#", "2"),
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_peer_group_peers.testacc_peers", "peers.0.id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_peer_group_peers.testacc_peers", "peers.0.name"),
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_peer_group_peers.testacc_peers", "peers.0.asn"),
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_peer_group_peers.testacc_peers", "peers.0.href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_peer_group_peers.testacc_peers", "peers.0.lifecycle_state"),
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_peer_group_peers.testacc_peers", "peers.0.resource_type"),
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_peer_group_peers.testacc_peers", "peers.0.endpoint.0.address"),
				),
			},
		},
	})
}

func testAccCheckIBMIsDynamicRouteServerPeerGroupPeersDataSourceConfigBasic(vpcname, subnetname, vniname, drsname, pgname, name string) string {
	return testAccCheckIBMIsDynamicRouteServerPeerGroupPeerConfig(vpcname, subnetname, vniname, drsname, pgname, name, 64514, 2, "enabled", "10.240.0.20") + `
	data "ibm_is_dynamic_route_server_peer_group_peers" "testacc_peers" {
		dynamic_route_server_id            = ibm_is_dynamic_route_server.testacc_drs.id
		dynamic_route_server_peer_group_id = ibm_is_dynamic_route_server_peer_group.testacc_peer_group.dynamic_route_server_peer_group_id
		depends_on                         = [ibm_is_dynamic_route_server_peer_group_peer.testacc_peer]
	}
	`
}

func TestDataSourceIBMIsDynamicRouteServerPeerGroupPeersSchemaShapes(t *testing.T) {
	ds := vpc.DataSourceIBMIsDynamicRouteServerPeerGroupPeers()
	assert.True(t, ds.Schema["dynamic_route_server_id"].Required)
	assert.True(t, ds.Schema["dynamic_route_server_peer_group_id"].Required)
	assert.True(t, ds.Schema["peers"].Computed)
	peer := ds.Schema["peers"].Elem.(*schema.Resource).Schema
	for _, key := range []string{"id", "asn", "bidirectional_forwarding_detection", "created_at", "creator", "endpoint", "href", "lifecycle_reasons", "lifecycle_state", "name", "priority", "resource_type", "sessions", "state", "status", "status_reasons"} {
		assert.True(t, peer[key].Computed, key)
	}
	endpoint := peer["endpoint"].Elem.(*schema.Resource).Schema
	for _, key := range []string{"address", "crn", "deleted", "gateway", "href", "id", "name", "resource_type"} {
		assert.True(t, endpoint[key].Computed, key)
	}
}

func TestDataSourceIBMIsDynamicRouteServerPeerGroupPeersDynamicRouteServerPeerGroupPeerToMap(t *testing.T) {
	t.Run("address subtype", func(t *testing.T) {
		result, err := vpc.DataSourceIBMIsDynamicRouteServerPeerGroupPeersDynamicRouteServerPeerGroupPeerToMap(&vpcv1.DynamicRouteServerPeerGroupPeerAddress{
			ID: core.StringPtr("p1"), Name: core.StringPtr("peer"), Asn: core.Int64Ptr(64513), Priority: core.Int64Ptr(0), State: core.StringPtr("enabled"),
			Endpoint: &vpcv1.DynamicRouteServerPeerGroupPeerAddressEndpointIPGateway{Address: core.StringPtr("10.240.0.10")},
		})
		assert.Nil(t, err)
		assert.Equal(t, "p1", result["id"])
		assert.Equal(t, 64513, result["asn"])
		assert.Equal(t, 0, result["priority"])
		assert.Equal(t, []map[string]interface{}{{"address": "10.240.0.10"}}, result["endpoint"])
		assert.Equal(t, []map[string]interface{}{}, result["sessions"])
	})
	t.Run("transit gateway subtype", func(t *testing.T) {
		result, err := vpc.DataSourceIBMIsDynamicRouteServerPeerGroupPeersDynamicRouteServerPeerGroupPeerToMap(&vpcv1.DynamicRouteServerPeerGroupPeerTransitGateway{
			ID: core.StringPtr("p2"), Endpoint: &vpcv1.DynamicRouteServerPeerGroupPeerTransitGatewayEndpointTransitGatewayReference{CRN: core.StringPtr("crn"), ID: core.StringPtr("tgw"), ResourceType: core.StringPtr("transit_gateway")},
		})
		assert.Nil(t, err)
		assert.Equal(t, "p2", result["id"])
		assert.Equal(t, []map[string]interface{}{{"crn": "crn", "id": "tgw", "resource_type": "transit_gateway"}}, result["endpoint"])
	})
	t.Run("base struct", func(t *testing.T) {
		result, err := vpc.DataSourceIBMIsDynamicRouteServerPeerGroupPeersDynamicRouteServerPeerGroupPeerToMap(&vpcv1.DynamicRouteServerPeerGroupPeer{ID: core.StringPtr("p3")})
		assert.Nil(t, err)
		assert.Equal(t, "p3", result["id"])
	})
	t.Run("nil is an error", func(t *testing.T) {
		_, err := vpc.DataSourceIBMIsDynamicRouteServerPeerGroupPeersDynamicRouteServerPeerGroupPeerToMap(nil)
		assert.NotNil(t, err)
	})
}
