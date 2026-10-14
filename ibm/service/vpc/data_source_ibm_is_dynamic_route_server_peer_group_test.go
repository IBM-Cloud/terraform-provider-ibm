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
	. "github.com/IBM-Cloud/terraform-provider-ibm/ibm/unittest"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/stretchr/testify/assert"
)

func TestAccIBMIsDynamicRouteServerPeerGroupDataSourceBasic(t *testing.T) {
	vpcname := fmt.Sprintf("tf-vpc-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tf-subnet-%d", acctest.RandIntRange(10, 100))
	vniname := fmt.Sprintf("tf-vni-%d", acctest.RandIntRange(10, 100))
	drsname := fmt.Sprintf("tf-drs-%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf-peer-group-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsDynamicRouteServerPeerGroupDataSourceConfigBasic(vpcname, subnetname, vniname, drsname, name),
				Check: resource.ComposeAggregateTestCheckFunc(
					// Lookup by identifier.
					resource.TestCheckResourceAttrPair("data.ibm_is_dynamic_route_server_peer_group.by_id", "dynamic_route_server_peer_group_id", isDynamicRouteServerPeerGroupResource, "dynamic_route_server_peer_group_id"),
					resource.TestCheckResourceAttr("data.ibm_is_dynamic_route_server_peer_group.by_id", "name", name),
					resource.TestCheckResourceAttr("data.ibm_is_dynamic_route_server_peer_group.by_id", "roles.0", "next_hop"),
					resource.TestCheckResourceAttr("data.ibm_is_dynamic_route_server_peer_group.by_id", "health_monitor.0.roles.0", "next_hop"),
					resource.TestCheckResourceAttr("data.ibm_is_dynamic_route_server_peer_group.by_id", "peers.#", "1"),
					resource.TestCheckResourceAttr("data.ibm_is_dynamic_route_server_peer_group.by_id", "peers.0.asn", "64513"),
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_peer_group.by_id", "peers.0.id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_peer_group.by_id", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_peer_group.by_id", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_peer_group.by_id", "lifecycle_state"),
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_peer_group.by_id", "resource_type"),
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_peer_group.by_id", "state"),
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_peer_group.by_id", "status"),
					// Lookup by name.
					resource.TestCheckResourceAttrPair("data.ibm_is_dynamic_route_server_peer_group.by_name", "dynamic_route_server_peer_group_id", isDynamicRouteServerPeerGroupResource, "dynamic_route_server_peer_group_id"),
					resource.TestCheckResourceAttr("data.ibm_is_dynamic_route_server_peer_group.by_name", "name", name),
				),
			},
		},
	})
}

func testAccCheckIBMIsDynamicRouteServerPeerGroupDataSourceConfigBasic(vpcname, subnetname, vniname, drsname, name string) string {
	return testAccCheckIBMIsDynamicRouteServerPeerGroupConfigBasic(vpcname, subnetname, vniname, drsname, name) + `
	data "ibm_is_dynamic_route_server_peer_group" "by_id" {
		dynamic_route_server_id            = ibm_is_dynamic_route_server.testacc_drs.id
		dynamic_route_server_peer_group_id = ibm_is_dynamic_route_server_peer_group.testacc_peer_group.dynamic_route_server_peer_group_id
	}

	data "ibm_is_dynamic_route_server_peer_group" "by_name" {
		dynamic_route_server_id = ibm_is_dynamic_route_server.testacc_drs.id
		name                    = ibm_is_dynamic_route_server_peer_group.testacc_peer_group.name
	}
	`
}

func TestDataSourceIBMIsDynamicRouteServerPeerGroupSchemaShapes(t *testing.T) {
	ds := vpc.DataSourceIBMIsDynamicRouteServerPeerGroup()
	assert.True(t, ds.Schema["dynamic_route_server_id"].Required)
	assert.True(t, ds.Schema["dynamic_route_server_peer_group_id"].Optional)
	assert.True(t, ds.Schema["name"].Optional)
	assert.Equal(t, []string{"dynamic_route_server_peer_group_id", "name"}, ds.Schema["name"].ExactlyOneOf)
	for _, key := range []string{"created_at", "creator", "health_monitor", "href", "lifecycle_reasons", "lifecycle_state", "peers", "prefix_watcher", "resource_type", "roles", "state", "status", "status_reasons"} {
		assert.True(t, ds.Schema[key].Computed, key)
		assert.False(t, ds.Schema[key].Optional, key)
	}
}

func TestDataSourceIBMIsDynamicRouteServerPeerGroupFindByName(t *testing.T) {
	groups := []vpcv1.DynamicRouteServerPeerGroup{
		{ID: core.StringPtr("a"), Name: core.StringPtr("alpha")},
		{ID: core.StringPtr("b"), Name: core.StringPtr("beta")},
		{ID: core.StringPtr("c")},
	}
	found, err := vpc.DataSourceIBMIsDynamicRouteServerPeerGroupFindByName(groups, "beta")
	assert.Nil(t, err)
	assert.Equal(t, "b", *found.ID)

	_, err = vpc.DataSourceIBMIsDynamicRouteServerPeerGroupFindByName(groups, "gamma")
	assert.NotNil(t, err)
	_, err = vpc.DataSourceIBMIsDynamicRouteServerPeerGroupFindByName(groups, "")
	assert.NotNil(t, err)
	_, err = vpc.DataSourceIBMIsDynamicRouteServerPeerGroupFindByName(nil, "alpha")
	assert.NotNil(t, err)
	_, err = vpc.DataSourceIBMIsDynamicRouteServerPeerGroupFindByName(append(groups, vpcv1.DynamicRouteServerPeerGroup{ID: core.StringPtr("d"), Name: core.StringPtr("alpha")}), "alpha")
	assert.NotNil(t, err, "duplicate names are an error")
}

func TestDataSourceIBMIsDynamicRouteServerPeerGroupDynamicRouteServerPeerGroupToMap(t *testing.T) {
	t.Run("nil is an error", func(t *testing.T) {
		_, err := vpc.DataSourceIBMIsDynamicRouteServerPeerGroupDynamicRouteServerPeerGroupToMap(nil)
		assert.NotNil(t, err)
	})
	t.Run("empty struct gives empty lists", func(t *testing.T) {
		result, err := vpc.DataSourceIBMIsDynamicRouteServerPeerGroupDynamicRouteServerPeerGroupToMap(&vpcv1.DynamicRouteServerPeerGroup{})
		assert.Nil(t, err)
		assert.Equal(t, []string{}, result["creator"])
		assert.Equal(t, []string{}, result["roles"])
		assert.Equal(t, []map[string]interface{}{}, result["peers"])
		assert.Equal(t, []map[string]interface{}{}, result["lifecycle_reasons"])
		assert.Equal(t, []map[string]interface{}{}, result["status_reasons"])
		_, hasHealthMonitor := result["health_monitor"]
		assert.False(t, hasHealthMonitor)
		_, hasPrefixWatcher := result["prefix_watcher"]
		assert.False(t, hasPrefixWatcher)
	})
	t.Run("full struct with both peer subtypes", func(t *testing.T) {
		model := &vpcv1.DynamicRouteServerPeerGroup{
			ID:               core.StringPtr("pg-id"),
			CreatedAt:        CreateMockDateTime("2026-01-02T03:04:05.000Z"),
			Creator:          []string{"dynamic_route_server"},
			HealthMonitor:    &vpcv1.DynamicRouteServerPeerGroupHealthMonitorNextHop{Mode: core.StringPtr("bfd"), Roles: []string{"next_hop"}},
			Href:             core.StringPtr("https://example.com/pg"),
			LifecycleReasons: []vpcv1.LifecycleReason{{Code: core.StringPtr("c"), Message: core.StringPtr("m")}},
			LifecycleState:   core.StringPtr("stable"),
			Name:             core.StringPtr("pg"),
			Peers: []vpcv1.DynamicRouteServerPeerGroupPeerIntf{
				&vpcv1.DynamicRouteServerPeerGroupPeerAddress{ID: core.StringPtr("p1"), Asn: core.Int64Ptr(1)},
				&vpcv1.DynamicRouteServerPeerGroupPeerTransitGateway{ID: core.StringPtr("p2")},
			},
			PrefixWatcher: &vpcv1.DynamicRouteServerPeerGroupPrefixWatcher{MonitoredPrefixes: []vpcv1.DynamicRouteServerPeerGroupPrefixWatcherMonitoredPrefix{{Prefix: core.StringPtr("10.0.0.0/8")}}},
			ResourceType:  core.StringPtr("dynamic_route_server_peer_group"),
			Roles:         []string{"next_hop", "router"},
			State:         core.StringPtr("enabled"),
			Status:        core.StringPtr("up"),
			StatusReasons: []vpcv1.DynamicRouteServerPeerGroupStatusReason{{Code: core.StringPtr("c"), Message: core.StringPtr("m")}},
		}
		result, err := vpc.DataSourceIBMIsDynamicRouteServerPeerGroupDynamicRouteServerPeerGroupToMap(model)
		assert.Nil(t, err)
		assert.Equal(t, "pg-id", result["id"])
		assert.Equal(t, "2026-01-02T03:04:05.000Z", result["created_at"])
		assert.Equal(t, []map[string]interface{}{{"mode": "bfd", "roles": []string{"next_hop"}}}, result["health_monitor"])
		peers := result["peers"].([]map[string]interface{})
		assert.Len(t, peers, 2)
		assert.Equal(t, "p1", peers[0]["id"])
		assert.Equal(t, 1, peers[0]["asn"])
		assert.Equal(t, "p2", peers[1]["id"])
		prefixWatcher := result["prefix_watcher"].([]map[string]interface{})[0]
		assert.Equal(t, []map[string]interface{}{{"prefix": "10.0.0.0/8"}}, prefixWatcher["monitored_prefixes"])
		assert.Equal(t, []string{"next_hop", "router"}, result["roles"])
		assert.Equal(t, "up", result["status"])
	})
	t.Run("unknown peer subtype is an error", func(t *testing.T) {
		_, err := vpc.DataSourceIBMIsDynamicRouteServerPeerGroupDynamicRouteServerPeerGroupToMap(&vpcv1.DynamicRouteServerPeerGroup{Peers: []vpcv1.DynamicRouteServerPeerGroupPeerIntf{nil}})
		assert.NotNil(t, err)
	})
}
