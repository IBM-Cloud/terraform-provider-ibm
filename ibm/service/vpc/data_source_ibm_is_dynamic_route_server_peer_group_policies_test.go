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

const drsPoliciesDataSourceName = "data.ibm_is_dynamic_route_server_peer_group_policies.is_dynamic_route_server_peer_group_policies_instance"

func TestAccIBMIsDynamicRouteServerPeerGroupPoliciesDataSourceBasic(t *testing.T) {
	suffix := fmt.Sprintf("%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf-drs-policy-%s", suffix)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsDynamicRouteServerPeerGroupPoliciesDataSourceConfig(suffix, name),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(drsPoliciesDataSourceName, "id"),
					resource.TestCheckResourceAttrPair(drsPoliciesDataSourceName, "dynamic_route_server_id", drsPolicyResourceName, "dynamic_route_server_id"),
					resource.TestCheckResourceAttrPair(drsPoliciesDataSourceName, "dynamic_route_server_peer_group_id", drsPolicyResourceName, "dynamic_route_server_peer_group_id"),
					resource.TestCheckResourceAttr(drsPoliciesDataSourceName, "sort", "name"),
					resource.TestCheckResourceAttr(drsPoliciesDataSourceName, "policies.#", "1"),
					resource.TestCheckResourceAttrPair(drsPoliciesDataSourceName, "policies.0.id", drsPolicyResourceName, "dynamic_route_server_peer_group_policy_id"),
					resource.TestCheckResourceAttr(drsPoliciesDataSourceName, "policies.0.name", name),
					resource.TestCheckResourceAttr(drsPoliciesDataSourceName, "policies.0.type", "custom_routes"),
					resource.TestCheckResourceAttr(drsPoliciesDataSourceName, "policies.0.state", "enabled"),
					resource.TestCheckResourceAttr(drsPoliciesDataSourceName, "policies.0.lifecycle_state", "stable"),
					resource.TestCheckResourceAttr(drsPoliciesDataSourceName, "policies.0.custom_routes.#", "1"),
					resource.TestCheckResourceAttr(drsPoliciesDataSourceName, "policies.0.custom_routes.0.destination", "192.168.3.0/24"),
					resource.TestCheckResourceAttrSet(drsPoliciesDataSourceName, "policies.0.created_at"),
					resource.TestCheckResourceAttrSet(drsPoliciesDataSourceName, "policies.0.href"),
					resource.TestCheckResourceAttrSet(drsPoliciesDataSourceName, "policies.0.resource_type"),
				),
			},
		},
	})
}

func testAccCheckIBMIsDynamicRouteServerPeerGroupPoliciesDataSourceConfig(suffix string, name string) string {
	return testAccCheckIBMIsDynamicRouteServerPeerGroupPolicyConfigCustomRoutes(suffix, name, "enabled", []string{"192.168.3.0/24"}) + `
		data "ibm_is_dynamic_route_server_peer_group_policies" "is_dynamic_route_server_peer_group_policies_instance" {
			dynamic_route_server_id            = ibm_is_dynamic_route_server_peer_group_policy.is_dynamic_route_server_peer_group_policy_instance.dynamic_route_server_id
			dynamic_route_server_peer_group_id = ibm_is_dynamic_route_server_peer_group_policy.is_dynamic_route_server_peer_group_policy_instance.dynamic_route_server_peer_group_id
			sort                               = "name"
		}
	`
}

func TestDataSourceIBMIsDynamicRouteServerPeerGroupPoliciesDynamicRouteServerPeerGroupPolicyToMap(t *testing.T) {
	t.Run("nil interface", func(t *testing.T) {
		_, err := vpc.DataSourceIBMIsDynamicRouteServerPeerGroupPoliciesDynamicRouteServerPeerGroupPolicyToMap(nil)
		assert.NotNil(t, err)
	})
	t.Run("base struct", func(t *testing.T) {
		result, err := vpc.DataSourceIBMIsDynamicRouteServerPeerGroupPoliciesDynamicRouteServerPeerGroupPolicyToMap(&vpcv1.DynamicRouteServerPeerGroupPolicy{
			ID:    core.StringPtr("p1"),
			Name:  core.StringPtr("my-policy"),
			State: core.StringPtr("enabled"),
			Type:  core.StringPtr("custom_routes"),
			CustomRoutes: []vpcv1.DynamicRouteServerPeerGroupPolicyCustomRoute{
				{Destination: core.StringPtr("192.168.3.0/24")},
				{Destination: core.StringPtr("192.168.4.0/24")},
			},
		})
		assert.Nil(t, err)
		assert.Equal(t, "p1", result["id"])
		assert.Len(t, result["custom_routes"], 2)
		assert.Equal(t, []map[string]interface{}{}, result["peer_groups"])
	})
	t.Run("each subtype", func(t *testing.T) {
		for _, model := range []vpcv1.DynamicRouteServerPeerGroupPolicyIntf{
			&vpcv1.DynamicRouteServerPeerGroupPolicyCustomRoutes{ID: core.StringPtr("p1"), Type: core.StringPtr("custom_routes")},
			&vpcv1.DynamicRouteServerPeerGroupPolicyLearnedRoutes{ID: core.StringPtr("p2"), Type: core.StringPtr("learned_routes")},
			&vpcv1.DynamicRouteServerPeerGroupPolicyVPCAddressPrefixes{ID: core.StringPtr("p3"), Type: core.StringPtr("vpc_address_prefixes")},
			&vpcv1.DynamicRouteServerPeerGroupPolicyVPCRoutingTables{ID: core.StringPtr("p4"), Type: core.StringPtr("vpc_routing_tables"), RouteDeleteDelay: core.Int64Ptr(5)},
		} {
			result, err := vpc.DataSourceIBMIsDynamicRouteServerPeerGroupPoliciesDynamicRouteServerPeerGroupPolicyToMap(model)
			assert.Nil(t, err)
			assert.NotEmpty(t, result["id"])
			assert.NotEmpty(t, result["type"])
		}
	})
}

func TestDataSourceIBMIsDynamicRouteServerPeerGroupPoliciesDynamicRouteServerPeerGroupPolicyCustomRoutesToMap(t *testing.T) {
	_, err := vpc.DataSourceIBMIsDynamicRouteServerPeerGroupPoliciesDynamicRouteServerPeerGroupPolicyCustomRoutesToMap(nil)
	assert.NotNil(t, err)
	result, err := vpc.DataSourceIBMIsDynamicRouteServerPeerGroupPoliciesDynamicRouteServerPeerGroupPolicyCustomRoutesToMap(&vpcv1.DynamicRouteServerPeerGroupPolicyCustomRoutes{
		ID:           core.StringPtr("p1"),
		Type:         core.StringPtr("custom_routes"),
		CustomRoutes: []vpcv1.DynamicRouteServerPeerGroupPolicyCustomRoute{{Destination: core.StringPtr("192.168.3.0/24")}},
	})
	assert.Nil(t, err)
	assert.Equal(t, []map[string]interface{}{{"destination": "192.168.3.0/24"}}, result["custom_routes"])
}

func TestDataSourceIBMIsDynamicRouteServerPeerGroupPoliciesDynamicRouteServerPeerGroupPolicyLearnedRoutesToMap(t *testing.T) {
	_, err := vpc.DataSourceIBMIsDynamicRouteServerPeerGroupPoliciesDynamicRouteServerPeerGroupPolicyLearnedRoutesToMap(nil)
	assert.NotNil(t, err)
	result, err := vpc.DataSourceIBMIsDynamicRouteServerPeerGroupPoliciesDynamicRouteServerPeerGroupPolicyLearnedRoutesToMap(&vpcv1.DynamicRouteServerPeerGroupPolicyLearnedRoutes{
		ID:               core.StringPtr("p2"),
		Type:             core.StringPtr("learned_routes"),
		ExcludedPrefixes: []vpcv1.DynamicRouteServerPeerGroupPrefixWatcherExcludedPrefixPrototype{{Prefix: core.StringPtr("10.0.0.0/8")}},
		PeerGroups:       []vpcv1.DynamicRouteServerPeerGroupReference{{ID: core.StringPtr("pg1")}},
	})
	assert.Nil(t, err)
	assert.Equal(t, []map[string]interface{}{{"ge": 0, "le": 0, "prefix": "10.0.0.0/8"}}, result["excluded_prefixes"])
	assert.Equal(t, []map[string]interface{}{{"id": "pg1"}}, result["peer_groups"])
}

func TestDataSourceIBMIsDynamicRouteServerPeerGroupPoliciesDynamicRouteServerPeerGroupPolicyVPCAddressPrefixesToMap(t *testing.T) {
	_, err := vpc.DataSourceIBMIsDynamicRouteServerPeerGroupPoliciesDynamicRouteServerPeerGroupPolicyVPCAddressPrefixesToMap(nil)
	assert.NotNil(t, err)
	result, err := vpc.DataSourceIBMIsDynamicRouteServerPeerGroupPoliciesDynamicRouteServerPeerGroupPolicyVPCAddressPrefixesToMap(&vpcv1.DynamicRouteServerPeerGroupPolicyVPCAddressPrefixes{ID: core.StringPtr("p3"), Type: core.StringPtr("vpc_address_prefixes")})
	assert.Nil(t, err)
	assert.Equal(t, "vpc_address_prefixes", result["type"])
	assert.Equal(t, []map[string]interface{}{}, result["custom_routes"])
}

func TestDataSourceIBMIsDynamicRouteServerPeerGroupPoliciesDynamicRouteServerPeerGroupPolicyVPCRoutingTablesToMap(t *testing.T) {
	_, err := vpc.DataSourceIBMIsDynamicRouteServerPeerGroupPoliciesDynamicRouteServerPeerGroupPolicyVPCRoutingTablesToMap(nil)
	assert.NotNil(t, err)
	result, err := vpc.DataSourceIBMIsDynamicRouteServerPeerGroupPoliciesDynamicRouteServerPeerGroupPolicyVPCRoutingTablesToMap(&vpcv1.DynamicRouteServerPeerGroupPolicyVPCRoutingTables{
		ID:               core.StringPtr("p4"),
		Type:             core.StringPtr("vpc_routing_tables"),
		RouteDeleteDelay: core.Int64Ptr(0),
		NextHops:         []vpcv1.DynamicRouteServerPeerGroupPolicyNextHopIntf{&vpcv1.DynamicRouteServerPeerGroupPolicyNextHopDynamicRouteServerPeerGroupReference{ID: core.StringPtr("pg2")}},
		RoutingTables:    []vpcv1.DynamicRouteServerPeerGroupPolicyVPCRoutingTablesRoutingTable{{Advertise: core.BoolPtr(false), VPCRoutingTable: &vpcv1.RoutingTableReference{ID: core.StringPtr("rt1")}}},
	})
	assert.Nil(t, err)
	assert.Equal(t, 0, result["route_delete_delay"])
	assert.Equal(t, []map[string]interface{}{{"id": "pg2"}}, result["next_hops"])
	assert.Equal(t, []map[string]interface{}{{"advertise": false, "vpc_routing_table": []map[string]interface{}{{"id": "rt1"}}}}, result["routing_tables"])
}

func TestDataSourceIBMIsDynamicRouteServerPeerGroupPoliciesDynamicRouteServerPeerGroupPolicyCustomRouteToMap(t *testing.T) {
	_, err := vpc.DataSourceIBMIsDynamicRouteServerPeerGroupPoliciesDynamicRouteServerPeerGroupPolicyCustomRouteToMap(nil)
	assert.NotNil(t, err)
	result, err := vpc.DataSourceIBMIsDynamicRouteServerPeerGroupPoliciesDynamicRouteServerPeerGroupPolicyCustomRouteToMap(&vpcv1.DynamicRouteServerPeerGroupPolicyCustomRoute{Destination: core.StringPtr("192.168.3.0/24")})
	assert.Nil(t, err)
	assert.Equal(t, map[string]interface{}{"destination": "192.168.3.0/24"}, result)
}

func TestDataSourceIBMIsDynamicRouteServerPeerGroupPoliciesDynamicRouteServerPeerGroupPrefixWatcherExcludedPrefixPrototypeToMap(t *testing.T) {
	_, err := vpc.DataSourceIBMIsDynamicRouteServerPeerGroupPoliciesDynamicRouteServerPeerGroupPrefixWatcherExcludedPrefixPrototypeToMap(nil)
	assert.NotNil(t, err)
	result, err := vpc.DataSourceIBMIsDynamicRouteServerPeerGroupPoliciesDynamicRouteServerPeerGroupPrefixWatcherExcludedPrefixPrototypeToMap(&vpcv1.DynamicRouteServerPeerGroupPrefixWatcherExcludedPrefixPrototype{Ge: core.Int64Ptr(0), Le: core.Int64Ptr(32), Prefix: core.StringPtr("10.0.0.0/8")})
	assert.Nil(t, err)
	assert.Equal(t, map[string]interface{}{"ge": 0, "le": 32, "prefix": "10.0.0.0/8"}, result)
}

func TestDataSourceIBMIsDynamicRouteServerPeerGroupPoliciesDynamicRouteServerPeerGroupReferenceToMap(t *testing.T) {
	_, err := vpc.DataSourceIBMIsDynamicRouteServerPeerGroupPoliciesDynamicRouteServerPeerGroupReferenceToMap(nil)
	assert.NotNil(t, err)
	result, err := vpc.DataSourceIBMIsDynamicRouteServerPeerGroupPoliciesDynamicRouteServerPeerGroupReferenceToMap(&vpcv1.DynamicRouteServerPeerGroupReference{ID: core.StringPtr("pg1"), Href: core.StringPtr("href")})
	assert.Nil(t, err)
	assert.Equal(t, map[string]interface{}{"id": "pg1", "href": "href"}, result)
}

func TestDataSourceIBMIsDynamicRouteServerPeerGroupPoliciesDeletedToMap(t *testing.T) {
	_, err := vpc.DataSourceIBMIsDynamicRouteServerPeerGroupPoliciesDeletedToMap(nil)
	assert.NotNil(t, err)
	result, err := vpc.DataSourceIBMIsDynamicRouteServerPeerGroupPoliciesDeletedToMap(&vpcv1.Deleted{MoreInfo: core.StringPtr("info")})
	assert.Nil(t, err)
	assert.Equal(t, map[string]interface{}{"more_info": "info"}, result)
}

func TestDataSourceIBMIsDynamicRouteServerPeerGroupPoliciesDynamicRouteServerPeerGroupPolicyNextHopToMap(t *testing.T) {
	_, err := vpc.DataSourceIBMIsDynamicRouteServerPeerGroupPoliciesDynamicRouteServerPeerGroupPolicyNextHopToMap(nil)
	assert.NotNil(t, err)
	result, err := vpc.DataSourceIBMIsDynamicRouteServerPeerGroupPoliciesDynamicRouteServerPeerGroupPolicyNextHopToMap(&vpcv1.DynamicRouteServerPeerGroupPolicyNextHop{ID: core.StringPtr("pg1")})
	assert.Nil(t, err)
	assert.Equal(t, map[string]interface{}{"id": "pg1"}, result)
}

func TestDataSourceIBMIsDynamicRouteServerPeerGroupPoliciesDynamicRouteServerPeerGroupPolicyNextHopDynamicRouteServerPeerGroupReferenceToMap(t *testing.T) {
	_, err := vpc.DataSourceIBMIsDynamicRouteServerPeerGroupPoliciesDynamicRouteServerPeerGroupPolicyNextHopDynamicRouteServerPeerGroupReferenceToMap(nil)
	assert.NotNil(t, err)
	result, err := vpc.DataSourceIBMIsDynamicRouteServerPeerGroupPoliciesDynamicRouteServerPeerGroupPolicyNextHopDynamicRouteServerPeerGroupReferenceToMap(&vpcv1.DynamicRouteServerPeerGroupPolicyNextHopDynamicRouteServerPeerGroupReference{ID: core.StringPtr("pg1"), ResourceType: core.StringPtr("dynamic_route_server_peer_group")})
	assert.Nil(t, err)
	assert.Equal(t, map[string]interface{}{"id": "pg1", "resource_type": "dynamic_route_server_peer_group"}, result)
}

func TestDataSourceIBMIsDynamicRouteServerPeerGroupPoliciesDynamicRouteServerPeerGroupPolicyVPCRoutingTablesRoutingTableToMap(t *testing.T) {
	_, err := vpc.DataSourceIBMIsDynamicRouteServerPeerGroupPoliciesDynamicRouteServerPeerGroupPolicyVPCRoutingTablesRoutingTableToMap(nil)
	assert.NotNil(t, err)
	result, err := vpc.DataSourceIBMIsDynamicRouteServerPeerGroupPoliciesDynamicRouteServerPeerGroupPolicyVPCRoutingTablesRoutingTableToMap(&vpcv1.DynamicRouteServerPeerGroupPolicyVPCRoutingTablesRoutingTable{Advertise: core.BoolPtr(true)})
	assert.Nil(t, err)
	assert.Equal(t, map[string]interface{}{"advertise": true}, result)
}

func TestDataSourceIBMIsDynamicRouteServerPeerGroupPoliciesRoutingTableReferenceToMap(t *testing.T) {
	_, err := vpc.DataSourceIBMIsDynamicRouteServerPeerGroupPoliciesRoutingTableReferenceToMap(nil)
	assert.NotNil(t, err)
	result, err := vpc.DataSourceIBMIsDynamicRouteServerPeerGroupPoliciesRoutingTableReferenceToMap(&vpcv1.RoutingTableReference{ID: core.StringPtr("rt1"), Deleted: &vpcv1.Deleted{MoreInfo: core.StringPtr("info")}})
	assert.Nil(t, err)
	assert.Equal(t, map[string]interface{}{"id": "rt1", "deleted": []map[string]interface{}{{"more_info": "info"}}}, result)
}
