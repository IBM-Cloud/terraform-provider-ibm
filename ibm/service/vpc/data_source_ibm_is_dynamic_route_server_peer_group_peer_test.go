// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/vpc"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/stretchr/testify/assert"
)

func TestAccIBMIsDynamicRouteServerPeerGroupPeerDataSourceBasic(t *testing.T) {
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
				Config: testAccCheckIBMIsDynamicRouteServerPeerGroupPeerDataSourceConfigBasic(vpcname, subnetname, vniname, drsname, pgname, name),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrPair("data.ibm_is_dynamic_route_server_peer_group_peer.by_id", "dynamic_route_server_peer_group_peer_id", isDynamicRouteServerPeerGroupPeerResource, "dynamic_route_server_peer_group_peer_id"),
					resource.TestCheckResourceAttr("data.ibm_is_dynamic_route_server_peer_group_peer.by_id", "name", name),
					resource.TestCheckResourceAttr("data.ibm_is_dynamic_route_server_peer_group_peer.by_id", "asn", "64514"),
					resource.TestCheckResourceAttr("data.ibm_is_dynamic_route_server_peer_group_peer.by_id", "priority", "2"),
					resource.TestCheckResourceAttr("data.ibm_is_dynamic_route_server_peer_group_peer.by_id", "state", "enabled"),
					resource.TestCheckResourceAttr("data.ibm_is_dynamic_route_server_peer_group_peer.by_id", "endpoint.0.address", "10.240.0.20"),
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_peer_group_peer.by_id", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_peer_group_peer.by_id", "creator"),
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_peer_group_peer.by_id", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_peer_group_peer.by_id", "lifecycle_state"),
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_peer_group_peer.by_id", "resource_type"),
					resource.TestCheckResourceAttrPair("data.ibm_is_dynamic_route_server_peer_group_peer.by_name", "dynamic_route_server_peer_group_peer_id", isDynamicRouteServerPeerGroupPeerResource, "dynamic_route_server_peer_group_peer_id"),
					resource.TestCheckResourceAttr("data.ibm_is_dynamic_route_server_peer_group_peer.by_name", "asn", "64514"),
				),
			},
		},
	})
}

func testAccCheckIBMIsDynamicRouteServerPeerGroupPeerDataSourceConfigBasic(vpcname, subnetname, vniname, drsname, pgname, name string) string {
	return testAccCheckIBMIsDynamicRouteServerPeerGroupPeerConfig(vpcname, subnetname, vniname, drsname, pgname, name, 64514, 2, "enabled", "10.240.0.20") + `
	data "ibm_is_dynamic_route_server_peer_group_peer" "by_id" {
		dynamic_route_server_id                 = ibm_is_dynamic_route_server.testacc_drs.id
		dynamic_route_server_peer_group_id      = ibm_is_dynamic_route_server_peer_group.testacc_peer_group.dynamic_route_server_peer_group_id
		dynamic_route_server_peer_group_peer_id = ibm_is_dynamic_route_server_peer_group_peer.testacc_peer.dynamic_route_server_peer_group_peer_id
	}

	data "ibm_is_dynamic_route_server_peer_group_peer" "by_name" {
		dynamic_route_server_id            = ibm_is_dynamic_route_server.testacc_drs.id
		dynamic_route_server_peer_group_id = ibm_is_dynamic_route_server_peer_group.testacc_peer_group.dynamic_route_server_peer_group_id
		name                               = ibm_is_dynamic_route_server_peer_group_peer.testacc_peer.name
	}
	`
}

func TestDataSourceIBMIsDynamicRouteServerPeerGroupPeerSchemaShapes(t *testing.T) {
	ds := vpc.DataSourceIBMIsDynamicRouteServerPeerGroupPeer()
	assert.True(t, ds.Schema["dynamic_route_server_id"].Required)
	assert.True(t, ds.Schema["dynamic_route_server_peer_group_id"].Required)
	assert.True(t, ds.Schema["dynamic_route_server_peer_group_peer_id"].Optional)
	assert.True(t, ds.Schema["name"].Optional)
	assert.Equal(t, []string{"dynamic_route_server_peer_group_peer_id", "name"}, ds.Schema["dynamic_route_server_peer_group_peer_id"].ExactlyOneOf)
	_, hasID := ds.Schema["id"]
	assert.False(t, hasID)
	for _, key := range []string{"asn", "bidirectional_forwarding_detection", "created_at", "creator", "endpoint", "href", "lifecycle_reasons", "lifecycle_state", "priority", "resource_type", "sessions", "state", "status", "status_reasons"} {
		assert.True(t, ds.Schema[key].Computed, key)
		assert.False(t, ds.Schema[key].Optional, key)
	}
}

func TestDataSourceIBMIsDynamicRouteServerPeerGroupPeerFindByName(t *testing.T) {
	peers := []vpcv1.DynamicRouteServerPeerGroupPeerIntf{
		&vpcv1.DynamicRouteServerPeerGroupPeerAddress{ID: core.StringPtr("a"), Name: core.StringPtr("alpha")},
		&vpcv1.DynamicRouteServerPeerGroupPeerTransitGateway{ID: core.StringPtr("b"), Name: core.StringPtr("beta")},
		&vpcv1.DynamicRouteServerPeerGroupPeer{ID: core.StringPtr("c"), Name: core.StringPtr("gamma")},
		&vpcv1.DynamicRouteServerPeerGroupPeer{ID: core.StringPtr("d")},
	}
	for name, id := range map[string]string{"alpha": "a", "beta": "b", "gamma": "c"} {
		found, err := vpc.DataSourceIBMIsDynamicRouteServerPeerGroupPeerFindByName(peers, name)
		assert.Nil(t, err)
		foundID, err := vpc.ResourceIBMIsDynamicRouteServerPeerGroupPeerIntfID(found)
		assert.Nil(t, err)
		assert.Equal(t, id, foundID)
	}
	_, err := vpc.DataSourceIBMIsDynamicRouteServerPeerGroupPeerFindByName(peers, "delta")
	assert.NotNil(t, err)
	_, err = vpc.DataSourceIBMIsDynamicRouteServerPeerGroupPeerFindByName(peers, "")
	assert.NotNil(t, err)
	_, err = vpc.DataSourceIBMIsDynamicRouteServerPeerGroupPeerFindByName(nil, "alpha")
	assert.NotNil(t, err)
	_, err = vpc.DataSourceIBMIsDynamicRouteServerPeerGroupPeerFindByName(append(peers, &vpcv1.DynamicRouteServerPeerGroupPeer{ID: core.StringPtr("e"), Name: core.StringPtr("alpha")}), "alpha")
	assert.NotNil(t, err, "duplicate names are an error")
	_, err = vpc.DataSourceIBMIsDynamicRouteServerPeerGroupPeerFindByName([]vpcv1.DynamicRouteServerPeerGroupPeerIntf{nil}, "alpha")
	assert.NotNil(t, err, "unknown subtype is an error")
}
