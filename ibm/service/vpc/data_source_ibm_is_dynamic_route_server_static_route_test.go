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

const drsStaticRouteDataSourceName = "data.ibm_is_dynamic_route_server_static_route.is_dynamic_route_server_static_route_instance"

func TestAccIBMIsDynamicRouteServerStaticRouteDataSourceBasic(t *testing.T) {
	suffix := fmt.Sprintf("%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf-drs-static-route-%s", suffix)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsDynamicRouteServerStaticRouteDataSourceConfig(suffix, name),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(drsStaticRouteDataSourceName, "id"),
					resource.TestCheckResourceAttrPair(drsStaticRouteDataSourceName, "dynamic_route_server_id", drsStaticRouteResourceName, "dynamic_route_server_id"),
					resource.TestCheckResourceAttrPair(drsStaticRouteDataSourceName, "dynamic_route_server_static_route_id", drsStaticRouteResourceName, "dynamic_route_server_static_route_id"),
					resource.TestCheckResourceAttr(drsStaticRouteDataSourceName, "name", name),
					resource.TestCheckResourceAttr(drsStaticRouteDataSourceName, "destination", "192.168.3.0/24"),
					resource.TestCheckResourceAttr(drsStaticRouteDataSourceName, "route_delete_delay", "20"),
					resource.TestCheckResourceAttr(drsStaticRouteDataSourceName, "lifecycle_state", "stable"),
					resource.TestCheckResourceAttr(drsStaticRouteDataSourceName, "next_hops.#", "1"),
					resource.TestCheckResourceAttrPair(drsStaticRouteDataSourceName, "next_hops.0.id", drsStaticRouteResourceName, "next_hops.0.id"),
					resource.TestCheckResourceAttr(drsStaticRouteDataSourceName, "routing_tables.#", "1"),
					resource.TestCheckResourceAttr(drsStaticRouteDataSourceName, "routing_tables.0.advertise", "true"),
					resource.TestCheckResourceAttrPair(drsStaticRouteDataSourceName, "routing_tables.0.vpc_routing_table.0.id", drsStaticRouteResourceName, "routing_tables.0.vpc_routing_table.0.id"),
					resource.TestCheckResourceAttrSet(drsStaticRouteDataSourceName, "href"),
					resource.TestCheckResourceAttrSet(drsStaticRouteDataSourceName, "resource_type"),
					resource.TestCheckResourceAttrSet(drsStaticRouteDataSourceName, "added_routes.#"),
				),
			},
		},
	})
}

func testAccCheckIBMIsDynamicRouteServerStaticRouteDataSourceConfig(suffix string, name string) string {
	return testAccCheckIBMIsDynamicRouteServerStaticRouteConfig(suffix, name, 20, true, "testacc_pg") + `
		data "ibm_is_dynamic_route_server_static_route" "is_dynamic_route_server_static_route_instance" {
			dynamic_route_server_id                 = ibm_is_dynamic_route_server_static_route.is_dynamic_route_server_static_route_instance.dynamic_route_server_id
			dynamic_route_server_static_route_id = ibm_is_dynamic_route_server_static_route.is_dynamic_route_server_static_route_instance.dynamic_route_server_static_route_id
		}
	`
}

func TestDataSourceIBMIsDynamicRouteServerStaticRouteDynamicRouteServerStaticRouteToMap(t *testing.T) {
	_, err := vpc.DataSourceIBMIsDynamicRouteServerStaticRouteDynamicRouteServerStaticRouteToMap(nil)
	assert.NotNil(t, err)
	result, err := vpc.DataSourceIBMIsDynamicRouteServerStaticRouteDynamicRouteServerStaticRouteToMap(&vpcv1.DynamicRouteServerStaticRoute{
		ID:               core.StringPtr("sr1"),
		Destination:      core.StringPtr("192.168.3.0/24"),
		RouteDeleteDelay: core.Int64Ptr(20),
		NextHops:         []vpcv1.DynamicRouteServerStaticRouteNextHopIntf{&vpcv1.DynamicRouteServerStaticRouteNextHopDynamicRouteServerPeerGroupReference{ID: core.StringPtr("pg1")}},
	})
	assert.Nil(t, err)
	assert.Equal(t, "sr1", result["id"])
	assert.Equal(t, 20, result["route_delete_delay"])
	assert.Equal(t, []map[string]interface{}{{"id": "pg1"}}, result["next_hops"])
	assert.Equal(t, []map[string]interface{}{}, result["routing_tables"])
}

func TestDataSourceIBMIsDynamicRouteServerStaticRouteDynamicRouteServerStaticRouteAddedRouteToMap(t *testing.T) {
	_, err := vpc.DataSourceIBMIsDynamicRouteServerStaticRouteDynamicRouteServerStaticRouteAddedRouteToMap(nil)
	assert.NotNil(t, err)
	result, err := vpc.DataSourceIBMIsDynamicRouteServerStaticRouteDynamicRouteServerStaticRouteAddedRouteToMap(&vpcv1.DynamicRouteServerStaticRouteAddedRoute{})
	assert.Nil(t, err)
	assert.Equal(t, map[string]interface{}{}, result)
	result, err = vpc.DataSourceIBMIsDynamicRouteServerStaticRouteDynamicRouteServerStaticRouteAddedRouteToMap(&vpcv1.DynamicRouteServerStaticRouteAddedRoute{VPCRoutingTableRoute: &vpcv1.RouteReference{ID: core.StringPtr("route1")}})
	assert.Nil(t, err)
	assert.Equal(t, map[string]interface{}{"vpc_routing_table_route": []map[string]interface{}{{"id": "route1"}}}, result)
}

func TestDataSourceIBMIsDynamicRouteServerStaticRouteRouteReferenceToMap(t *testing.T) {
	_, err := vpc.DataSourceIBMIsDynamicRouteServerStaticRouteRouteReferenceToMap(nil)
	assert.NotNil(t, err)
	result, err := vpc.DataSourceIBMIsDynamicRouteServerStaticRouteRouteReferenceToMap(&vpcv1.RouteReference{Href: core.StringPtr("href"), ID: core.StringPtr("route1"), Name: core.StringPtr("name"), Deleted: &vpcv1.Deleted{MoreInfo: core.StringPtr("info")}})
	assert.Nil(t, err)
	assert.Equal(t, map[string]interface{}{"href": "href", "id": "route1", "name": "name", "deleted": []map[string]interface{}{{"more_info": "info"}}}, result)
}

func TestDataSourceIBMIsDynamicRouteServerStaticRouteDeletedToMap(t *testing.T) {
	_, err := vpc.DataSourceIBMIsDynamicRouteServerStaticRouteDeletedToMap(nil)
	assert.NotNil(t, err)
	result, err := vpc.DataSourceIBMIsDynamicRouteServerStaticRouteDeletedToMap(&vpcv1.Deleted{MoreInfo: core.StringPtr("info")})
	assert.Nil(t, err)
	assert.Equal(t, map[string]interface{}{"more_info": "info"}, result)
}

func TestDataSourceIBMIsDynamicRouteServerStaticRouteLifecycleReasonToMap(t *testing.T) {
	_, err := vpc.DataSourceIBMIsDynamicRouteServerStaticRouteLifecycleReasonToMap(nil)
	assert.NotNil(t, err)
	result, err := vpc.DataSourceIBMIsDynamicRouteServerStaticRouteLifecycleReasonToMap(&vpcv1.LifecycleReason{Code: core.StringPtr("internal_error"), Message: core.StringPtr("msg"), MoreInfo: core.StringPtr("info")})
	assert.Nil(t, err)
	assert.Equal(t, map[string]interface{}{"code": "internal_error", "message": "msg", "more_info": "info"}, result)
}

func TestDataSourceIBMIsDynamicRouteServerStaticRouteDynamicRouteServerStaticRouteNextHopToMap(t *testing.T) {
	_, err := vpc.DataSourceIBMIsDynamicRouteServerStaticRouteDynamicRouteServerStaticRouteNextHopToMap(nil)
	assert.NotNil(t, err)
	result, err := vpc.DataSourceIBMIsDynamicRouteServerStaticRouteDynamicRouteServerStaticRouteNextHopToMap(&vpcv1.DynamicRouteServerStaticRouteNextHopDynamicRouteServerPeerGroupReference{ID: core.StringPtr("pg1")})
	assert.Nil(t, err)
	assert.Equal(t, map[string]interface{}{"id": "pg1"}, result)
	result, err = vpc.DataSourceIBMIsDynamicRouteServerStaticRouteDynamicRouteServerStaticRouteNextHopToMap(&vpcv1.DynamicRouteServerStaticRouteNextHop{Href: core.StringPtr("href")})
	assert.Nil(t, err)
	assert.Equal(t, map[string]interface{}{"href": "href"}, result)
}

func TestDataSourceIBMIsDynamicRouteServerStaticRouteDynamicRouteServerStaticRouteNextHopDynamicRouteServerPeerGroupReferenceToMap(t *testing.T) {
	_, err := vpc.DataSourceIBMIsDynamicRouteServerStaticRouteDynamicRouteServerStaticRouteNextHopDynamicRouteServerPeerGroupReferenceToMap(nil)
	assert.NotNil(t, err)
	result, err := vpc.DataSourceIBMIsDynamicRouteServerStaticRouteDynamicRouteServerStaticRouteNextHopDynamicRouteServerPeerGroupReferenceToMap(&vpcv1.DynamicRouteServerStaticRouteNextHopDynamicRouteServerPeerGroupReference{Name: core.StringPtr("name"), ResourceType: core.StringPtr("dynamic_route_server_peer_group")})
	assert.Nil(t, err)
	assert.Equal(t, map[string]interface{}{"name": "name", "resource_type": "dynamic_route_server_peer_group"}, result)
}

func TestDataSourceIBMIsDynamicRouteServerStaticRouteDynamicRouteServerStaticRouteRoutingTableToMap(t *testing.T) {
	_, err := vpc.DataSourceIBMIsDynamicRouteServerStaticRouteDynamicRouteServerStaticRouteRoutingTableToMap(nil)
	assert.NotNil(t, err)
	result, err := vpc.DataSourceIBMIsDynamicRouteServerStaticRouteDynamicRouteServerStaticRouteRoutingTableToMap(&vpcv1.DynamicRouteServerStaticRouteRoutingTable{})
	assert.Nil(t, err)
	assert.Equal(t, map[string]interface{}{"advertise": false}, result)
	result, err = vpc.DataSourceIBMIsDynamicRouteServerStaticRouteDynamicRouteServerStaticRouteRoutingTableToMap(&vpcv1.DynamicRouteServerStaticRouteRoutingTable{Advertise: core.BoolPtr(true), VPCRoutingTable: &vpcv1.RoutingTableReference{ID: core.StringPtr("rt1")}})
	assert.Nil(t, err)
	assert.Equal(t, map[string]interface{}{"advertise": true, "vpc_routing_table": []map[string]interface{}{{"id": "rt1"}}}, result)
}

func TestDataSourceIBMIsDynamicRouteServerStaticRouteRoutingTableReferenceToMap(t *testing.T) {
	_, err := vpc.DataSourceIBMIsDynamicRouteServerStaticRouteRoutingTableReferenceToMap(nil)
	assert.NotNil(t, err)
	result, err := vpc.DataSourceIBMIsDynamicRouteServerStaticRouteRoutingTableReferenceToMap(&vpcv1.RoutingTableReference{CRN: core.StringPtr("crn:rt"), Href: core.StringPtr("href"), ID: core.StringPtr("rt1"), Name: core.StringPtr("rt"), ResourceType: core.StringPtr("routing_table")})
	assert.Nil(t, err)
	assert.Equal(t, map[string]interface{}{"crn": "crn:rt", "href": "href", "id": "rt1", "name": "rt", "resource_type": "routing_table"}, result)
}
