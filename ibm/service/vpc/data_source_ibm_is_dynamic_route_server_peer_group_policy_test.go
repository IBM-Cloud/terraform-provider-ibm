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

const drsPolicyDataSourceName = "data.ibm_is_dynamic_route_server_peer_group_policy.is_dynamic_route_server_peer_group_policy_instance"

func TestAccIBMIsDynamicRouteServerPeerGroupPolicyDataSourceBasic(t *testing.T) {
	suffix := fmt.Sprintf("%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf-drs-policy-%s", suffix)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsDynamicRouteServerPeerGroupPolicyDataSourceConfig(suffix, name),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(drsPolicyDataSourceName, "id"),
					resource.TestCheckResourceAttrPair(drsPolicyDataSourceName, "dynamic_route_server_id", drsPolicyResourceName, "dynamic_route_server_id"),
					resource.TestCheckResourceAttrPair(drsPolicyDataSourceName, "dynamic_route_server_peer_group_id", drsPolicyResourceName, "dynamic_route_server_peer_group_id"),
					resource.TestCheckResourceAttrPair(drsPolicyDataSourceName, "dynamic_route_server_peer_group_policy_id", drsPolicyResourceName, "dynamic_route_server_peer_group_policy_id"),
					resource.TestCheckResourceAttr(drsPolicyDataSourceName, "name", name),
					resource.TestCheckResourceAttr(drsPolicyDataSourceName, "type", "vpc_routing_tables"),
					resource.TestCheckResourceAttr(drsPolicyDataSourceName, "state", "enabled"),
					resource.TestCheckResourceAttr(drsPolicyDataSourceName, "route_delete_delay", "15"),
					resource.TestCheckResourceAttr(drsPolicyDataSourceName, "lifecycle_state", "stable"),
					resource.TestCheckResourceAttr(drsPolicyDataSourceName, "next_hops.#", "1"),
					resource.TestCheckResourceAttrPair(drsPolicyDataSourceName, "next_hops.0.id", drsPolicyResourceName, "next_hops.0.id"),
					resource.TestCheckResourceAttr(drsPolicyDataSourceName, "routing_tables.#", "1"),
					resource.TestCheckResourceAttr(drsPolicyDataSourceName, "routing_tables.0.advertise", "true"),
					resource.TestCheckResourceAttrPair(drsPolicyDataSourceName, "routing_tables.0.vpc_routing_table.0.id", drsPolicyResourceName, "routing_tables.0.vpc_routing_table.0.id"),
					resource.TestCheckResourceAttrSet(drsPolicyDataSourceName, "created_at"),
					resource.TestCheckResourceAttrSet(drsPolicyDataSourceName, "href"),
					resource.TestCheckResourceAttrSet(drsPolicyDataSourceName, "resource_type"),
				),
			},
		},
	})
}

func testAccCheckIBMIsDynamicRouteServerPeerGroupPolicyDataSourceConfig(suffix string, name string) string {
	return testAccCheckIBMIsDynamicRouteServerPeerGroupPolicyConfigVPCRoutingTables(suffix, name, 15, true) + `
		data "ibm_is_dynamic_route_server_peer_group_policy" "is_dynamic_route_server_peer_group_policy_instance" {
			dynamic_route_server_id                      = ibm_is_dynamic_route_server_peer_group_policy.is_dynamic_route_server_peer_group_policy_instance.dynamic_route_server_id
			dynamic_route_server_peer_group_id           = ibm_is_dynamic_route_server_peer_group_policy.is_dynamic_route_server_peer_group_policy_instance.dynamic_route_server_peer_group_id
			dynamic_route_server_peer_group_policy_id = ibm_is_dynamic_route_server_peer_group_policy.is_dynamic_route_server_peer_group_policy_instance.dynamic_route_server_peer_group_policy_id
		}
	`
}

func TestDataSourceIBMIsDynamicRouteServerPeerGroupPolicyDynamicRouteServerPeerGroupPolicyToMap(t *testing.T) {
	t.Run("nil interface", func(t *testing.T) {
		_, err := vpc.DataSourceIBMIsDynamicRouteServerPeerGroupPolicyDynamicRouteServerPeerGroupPolicyToMap(nil)
		assert.NotNil(t, err)
	})
	t.Run("base struct", func(t *testing.T) {
		result, err := vpc.DataSourceIBMIsDynamicRouteServerPeerGroupPolicyDynamicRouteServerPeerGroupPolicyToMap(&vpcv1.DynamicRouteServerPeerGroupPolicy{ID: core.StringPtr("p1"), Type: core.StringPtr("vpc_address_prefixes")})
		assert.Nil(t, err)
		assert.Equal(t, "p1", result["id"])
		assert.Equal(t, "vpc_address_prefixes", result["type"])
		assert.Equal(t, 0, result["route_delete_delay"])
	})
	t.Run("learned routes subtype", func(t *testing.T) {
		result, err := vpc.DataSourceIBMIsDynamicRouteServerPeerGroupPolicyDynamicRouteServerPeerGroupPolicyToMap(&vpcv1.DynamicRouteServerPeerGroupPolicyLearnedRoutes{
			ID:         core.StringPtr("p2"),
			Type:       core.StringPtr("learned_routes"),
			PeerGroups: []vpcv1.DynamicRouteServerPeerGroupReference{{ID: core.StringPtr("pg1")}, {ID: core.StringPtr("pg2")}},
		})
		assert.Nil(t, err)
		assert.Len(t, result["peer_groups"], 2)
	})
}

func TestDataSourceIBMIsDynamicRouteServerPeerGroupPolicyDynamicRouteServerPeerGroupPolicyCustomRouteToMap(t *testing.T) {
	_, err := vpc.DataSourceIBMIsDynamicRouteServerPeerGroupPolicyDynamicRouteServerPeerGroupPolicyCustomRouteToMap(nil)
	assert.NotNil(t, err)
	result, err := vpc.DataSourceIBMIsDynamicRouteServerPeerGroupPolicyDynamicRouteServerPeerGroupPolicyCustomRouteToMap(&vpcv1.DynamicRouteServerPeerGroupPolicyCustomRoute{Destination: core.StringPtr("192.168.3.0/24")})
	assert.Nil(t, err)
	assert.Equal(t, map[string]interface{}{"destination": "192.168.3.0/24"}, result)
}

func TestDataSourceIBMIsDynamicRouteServerPeerGroupPolicyDynamicRouteServerPeerGroupPrefixWatcherExcludedPrefixPrototypeToMap(t *testing.T) {
	_, err := vpc.DataSourceIBMIsDynamicRouteServerPeerGroupPolicyDynamicRouteServerPeerGroupPrefixWatcherExcludedPrefixPrototypeToMap(nil)
	assert.NotNil(t, err)
	result, err := vpc.DataSourceIBMIsDynamicRouteServerPeerGroupPolicyDynamicRouteServerPeerGroupPrefixWatcherExcludedPrefixPrototypeToMap(&vpcv1.DynamicRouteServerPeerGroupPrefixWatcherExcludedPrefixPrototype{Prefix: core.StringPtr("10.0.0.0/8")})
	assert.Nil(t, err)
	assert.Equal(t, map[string]interface{}{"ge": 0, "le": 0, "prefix": "10.0.0.0/8"}, result)
	result, err = vpc.DataSourceIBMIsDynamicRouteServerPeerGroupPolicyDynamicRouteServerPeerGroupPrefixWatcherExcludedPrefixPrototypeToMap(&vpcv1.DynamicRouteServerPeerGroupPrefixWatcherExcludedPrefixPrototype{Ge: core.Int64Ptr(16), Le: core.Int64Ptr(24), Prefix: core.StringPtr("10.0.0.0/8")})
	assert.Nil(t, err)
	assert.Equal(t, map[string]interface{}{"ge": 16, "le": 24, "prefix": "10.0.0.0/8"}, result)
}

func TestDataSourceIBMIsDynamicRouteServerPeerGroupPolicyDynamicRouteServerPeerGroupReferenceToMap(t *testing.T) {
	_, err := vpc.DataSourceIBMIsDynamicRouteServerPeerGroupPolicyDynamicRouteServerPeerGroupReferenceToMap(nil)
	assert.NotNil(t, err)
	result, err := vpc.DataSourceIBMIsDynamicRouteServerPeerGroupPolicyDynamicRouteServerPeerGroupReferenceToMap(&vpcv1.DynamicRouteServerPeerGroupReference{
		Deleted:      &vpcv1.Deleted{MoreInfo: core.StringPtr("info")},
		Href:         core.StringPtr("href"),
		ID:           core.StringPtr("pg1"),
		Name:         core.StringPtr("name"),
		ResourceType: core.StringPtr("dynamic_route_server_peer_group"),
	})
	assert.Nil(t, err)
	assert.Equal(t, map[string]interface{}{"deleted": []map[string]interface{}{{"more_info": "info"}}, "href": "href", "id": "pg1", "name": "name", "resource_type": "dynamic_route_server_peer_group"}, result)
}

func TestDataSourceIBMIsDynamicRouteServerPeerGroupPolicyDeletedToMap(t *testing.T) {
	_, err := vpc.DataSourceIBMIsDynamicRouteServerPeerGroupPolicyDeletedToMap(nil)
	assert.NotNil(t, err)
	result, err := vpc.DataSourceIBMIsDynamicRouteServerPeerGroupPolicyDeletedToMap(&vpcv1.Deleted{MoreInfo: core.StringPtr("info")})
	assert.Nil(t, err)
	assert.Equal(t, map[string]interface{}{"more_info": "info"}, result)
}

func TestDataSourceIBMIsDynamicRouteServerPeerGroupPolicyDynamicRouteServerPeerGroupPolicyNextHopToMap(t *testing.T) {
	result, err := vpc.DataSourceIBMIsDynamicRouteServerPeerGroupPolicyDynamicRouteServerPeerGroupPolicyNextHopToMap(&vpcv1.DynamicRouteServerPeerGroupPolicyNextHopDynamicRouteServerPeerGroupReference{ID: core.StringPtr("pg1")})
	assert.Nil(t, err)
	assert.Equal(t, map[string]interface{}{"id": "pg1"}, result)
	result, err = vpc.DataSourceIBMIsDynamicRouteServerPeerGroupPolicyDynamicRouteServerPeerGroupPolicyNextHopToMap(&vpcv1.DynamicRouteServerPeerGroupPolicyNextHop{ID: core.StringPtr("pg2")})
	assert.Nil(t, err)
	assert.Equal(t, map[string]interface{}{"id": "pg2"}, result)
	_, err = vpc.DataSourceIBMIsDynamicRouteServerPeerGroupPolicyDynamicRouteServerPeerGroupPolicyNextHopToMap(nil)
	assert.NotNil(t, err)
}

func TestDataSourceIBMIsDynamicRouteServerPeerGroupPolicyDynamicRouteServerPeerGroupPolicyNextHopDynamicRouteServerPeerGroupReferenceToMap(t *testing.T) {
	_, err := vpc.DataSourceIBMIsDynamicRouteServerPeerGroupPolicyDynamicRouteServerPeerGroupPolicyNextHopDynamicRouteServerPeerGroupReferenceToMap(nil)
	assert.NotNil(t, err)
	result, err := vpc.DataSourceIBMIsDynamicRouteServerPeerGroupPolicyDynamicRouteServerPeerGroupPolicyNextHopDynamicRouteServerPeerGroupReferenceToMap(&vpcv1.DynamicRouteServerPeerGroupPolicyNextHopDynamicRouteServerPeerGroupReference{Name: core.StringPtr("name")})
	assert.Nil(t, err)
	assert.Equal(t, map[string]interface{}{"name": "name"}, result)
}

func TestDataSourceIBMIsDynamicRouteServerPeerGroupPolicyDynamicRouteServerPeerGroupPolicyVPCRoutingTablesRoutingTableToMap(t *testing.T) {
	_, err := vpc.DataSourceIBMIsDynamicRouteServerPeerGroupPolicyDynamicRouteServerPeerGroupPolicyVPCRoutingTablesRoutingTableToMap(nil)
	assert.NotNil(t, err)
	result, err := vpc.DataSourceIBMIsDynamicRouteServerPeerGroupPolicyDynamicRouteServerPeerGroupPolicyVPCRoutingTablesRoutingTableToMap(&vpcv1.DynamicRouteServerPeerGroupPolicyVPCRoutingTablesRoutingTable{
		Advertise:       core.BoolPtr(false),
		VPCRoutingTable: &vpcv1.RoutingTableReference{ID: core.StringPtr("rt1"), Name: core.StringPtr("rt")},
	})
	assert.Nil(t, err)
	assert.Equal(t, false, result["advertise"])
	assert.Equal(t, []map[string]interface{}{{"id": "rt1", "name": "rt"}}, result["vpc_routing_table"])
}

func TestDataSourceIBMIsDynamicRouteServerPeerGroupPolicyRoutingTableReferenceToMap(t *testing.T) {
	_, err := vpc.DataSourceIBMIsDynamicRouteServerPeerGroupPolicyRoutingTableReferenceToMap(nil)
	assert.NotNil(t, err)
	result, err := vpc.DataSourceIBMIsDynamicRouteServerPeerGroupPolicyRoutingTableReferenceToMap(&vpcv1.RoutingTableReference{CRN: core.StringPtr("crn:rt"), ID: core.StringPtr("rt1")})
	assert.Nil(t, err)
	assert.Equal(t, map[string]interface{}{"crn": "crn:rt", "id": "rt1"}, result)
}
