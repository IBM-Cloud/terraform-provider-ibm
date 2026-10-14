// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/vpc"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/stretchr/testify/assert"
)

const drsStaticRouteResourceName = "ibm_is_dynamic_route_server_static_route.is_dynamic_route_server_static_route_instance"

func TestAccIBMIsDynamicRouteServerStaticRouteBasic(t *testing.T) {
	var conf vpcv1.DynamicRouteServerStaticRoute
	var firstID string
	suffix := fmt.Sprintf("%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIsDynamicRouteServerStaticRouteDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsDynamicRouteServerStaticRouteConfigBasic(suffix, "192.168.3.0/24"),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIsDynamicRouteServerStaticRouteExists(drsStaticRouteResourceName, conf),
					testAccCheckIBMIsDynamicRouteServerPeerGroupPolicyRecordID(drsStaticRouteResourceName, &firstID),
					resource.TestCheckResourceAttr(drsStaticRouteResourceName, "destination", "192.168.3.0/24"),
					resource.TestCheckResourceAttr(drsStaticRouteResourceName, "route_delete_delay", "0"),
					resource.TestCheckResourceAttr(drsStaticRouteResourceName, "lifecycle_state", "stable"),
					resource.TestCheckResourceAttr(drsStaticRouteResourceName, "next_hops.#", "1"),
					resource.TestCheckResourceAttrPair(drsStaticRouteResourceName, "next_hops.0.id", "ibm_is_dynamic_route_server_peer_group.testacc_pg", "dynamic_route_server_peer_group_id"),
					resource.TestCheckResourceAttrSet(drsStaticRouteResourceName, "next_hops.0.name"),
					resource.TestCheckResourceAttr(drsStaticRouteResourceName, "routing_tables.#", "1"),
					resource.TestCheckResourceAttr(drsStaticRouteResourceName, "routing_tables.0.advertise", "false"),
					resource.TestCheckResourceAttrPair(drsStaticRouteResourceName, "routing_tables.0.vpc_routing_table.0.id", "ibm_is_vpc_routing_table.testacc_rt", "routing_table"),
					resource.TestCheckResourceAttrSet(drsStaticRouteResourceName, "name"),
					resource.TestCheckResourceAttrSet(drsStaticRouteResourceName, "href"),
					resource.TestCheckResourceAttrSet(drsStaticRouteResourceName, "dynamic_route_server_static_route_id"),
				),
			},
			resource.TestStep{
				// destination is ForceNew: changing it must re-create the static route.
				Config: testAccCheckIBMIsDynamicRouteServerStaticRouteConfigBasic(suffix, "192.168.4.0/24"),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIsDynamicRouteServerStaticRouteExists(drsStaticRouteResourceName, conf),
					testAccCheckIBMIsDynamicRouteServerPeerGroupPolicyIDChanged(drsStaticRouteResourceName, &firstID),
					resource.TestCheckResourceAttr(drsStaticRouteResourceName, "destination", "192.168.4.0/24"),
				),
			},
		},
	})
}

func TestAccIBMIsDynamicRouteServerStaticRouteAllArgs(t *testing.T) {
	var conf vpcv1.DynamicRouteServerStaticRoute
	suffix := fmt.Sprintf("%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf-drs-static-route-%s", suffix)
	nameUpdate := fmt.Sprintf("tf-drs-static-route-upd-%s", suffix)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIsDynamicRouteServerStaticRouteDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsDynamicRouteServerStaticRouteConfig(suffix, name, 0, true, "testacc_pg"),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIsDynamicRouteServerStaticRouteExists(drsStaticRouteResourceName, conf),
					resource.TestCheckResourceAttr(drsStaticRouteResourceName, "name", name),
					resource.TestCheckResourceAttr(drsStaticRouteResourceName, "route_delete_delay", "0"),
					resource.TestCheckResourceAttr(drsStaticRouteResourceName, "routing_tables.0.advertise", "true"),
					resource.TestCheckResourceAttrPair(drsStaticRouteResourceName, "next_hops.0.id", "ibm_is_dynamic_route_server_peer_group.testacc_pg", "dynamic_route_server_peer_group_id"),
				),
			},
			resource.TestStep{
				// Update every patchable field: name, route_delete_delay, next_hops and routing_tables.
				Config: testAccCheckIBMIsDynamicRouteServerStaticRouteConfig(suffix, nameUpdate, 45, false, "testacc_pg2"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(drsStaticRouteResourceName, "name", nameUpdate),
					resource.TestCheckResourceAttr(drsStaticRouteResourceName, "route_delete_delay", "45"),
					resource.TestCheckResourceAttr(drsStaticRouteResourceName, "routing_tables.0.advertise", "false"),
					resource.TestCheckResourceAttrPair(drsStaticRouteResourceName, "next_hops.0.id", "ibm_is_dynamic_route_server_peer_group.testacc_pg2", "dynamic_route_server_peer_group_id"),
				),
			},
			resource.TestStep{
				ResourceName:      drsStaticRouteResourceName,
				ImportState:       true,
				ImportStateVerify: true,
				// etag is a response header and is not part of the imported state.
				ImportStateVerifyIgnore: []string{"etag"},
			},
		},
	})
}

func testAccCheckIBMIsDynamicRouteServerStaticRouteConfigBasic(suffix string, destination string) string {
	return testAccCheckIBMIsDynamicRouteServerPolicyDependencies(suffix) + fmt.Sprintf(`
		resource "ibm_is_dynamic_route_server_static_route" "is_dynamic_route_server_static_route_instance" {
			dynamic_route_server_id = ibm_is_dynamic_route_server.testacc_drs.id
			destination             = "%s"
			next_hops {
				id = ibm_is_dynamic_route_server_peer_group.testacc_pg.dynamic_route_server_peer_group_id
			}
			routing_tables {
				advertise = false
				vpc_routing_table {
					id = ibm_is_vpc_routing_table.testacc_rt.routing_table
				}
			}
		}
	`, destination)
}

func testAccCheckIBMIsDynamicRouteServerStaticRouteConfig(suffix string, name string, routeDeleteDelay int, advertise bool, nextHopPeerGroup string) string {
	return testAccCheckIBMIsDynamicRouteServerPolicyDependencies(suffix) + fmt.Sprintf(`
		resource "ibm_is_dynamic_route_server_static_route" "is_dynamic_route_server_static_route_instance" {
			dynamic_route_server_id = ibm_is_dynamic_route_server.testacc_drs.id
			destination             = "192.168.3.0/24"
			name                    = "%s"
			route_delete_delay      = %d
			next_hops {
				id = ibm_is_dynamic_route_server_peer_group.%s.dynamic_route_server_peer_group_id
			}
			routing_tables {
				advertise = %t
				vpc_routing_table {
					id = ibm_is_vpc_routing_table.testacc_rt.routing_table
				}
			}
		}
	`, name, routeDeleteDelay, nextHopPeerGroup, advertise)
}

func testAccCheckIBMIsDynamicRouteServerStaticRouteExists(n string, obj vpcv1.DynamicRouteServerStaticRoute) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		vpcClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).VpcV1API()
		if err != nil {
			return err
		}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}
		if len(parts) != 2 {
			return fmt.Errorf("unexpected ID format %q", rs.Primary.ID)
		}

		getOptions := &vpcv1.GetDynamicRouteServerStaticRouteOptions{}
		getOptions.SetDynamicRouteServerID(parts[0])
		getOptions.SetID(parts[1])

		staticRoute, _, err := vpcClient.GetDynamicRouteServerStaticRoute(getOptions)
		if err != nil {
			return err
		}
		obj = *staticRoute
		return nil
	}
}

func testAccCheckIBMIsDynamicRouteServerStaticRouteDestroy(s *terraform.State) error {
	vpcClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).VpcV1API()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_is_dynamic_route_server_static_route" {
			continue
		}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}
		if len(parts) != 2 {
			return fmt.Errorf("unexpected ID format %q", rs.Primary.ID)
		}

		getOptions := &vpcv1.GetDynamicRouteServerStaticRouteOptions{}
		getOptions.SetDynamicRouteServerID(parts[0])
		getOptions.SetID(parts[1])

		_, response, err := vpcClient.GetDynamicRouteServerStaticRoute(getOptions)
		if err == nil {
			return fmt.Errorf("ibm_is_dynamic_route_server_static_route still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for ibm_is_dynamic_route_server_static_route (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}

// ---- unit tests ----

func drsStaticRouteTestRoutingTableReference() *vpcv1.RoutingTableReference {
	return &vpcv1.RoutingTableReference{
		CRN:          core.StringPtr("crn:v1:bluemix:public:is:us-south:a/123456::vpc-routing-table:r006-4727/r006-6885"),
		Href:         core.StringPtr("https://us-south.iaas.cloud.ibm.com/v1/vpcs/r006-4727/routing_tables/r006-6885"),
		ID:           core.StringPtr("r006-6885"),
		Name:         core.StringPtr("my-routing-table-1"),
		ResourceType: core.StringPtr("routing_table"),
	}
}

func TestResourceIBMIsDynamicRouteServerStaticRouteDynamicRouteServerStaticRouteToMap(t *testing.T) {
	t.Run("nil", func(t *testing.T) {
		_, err := vpc.ResourceIBMIsDynamicRouteServerStaticRouteDynamicRouteServerStaticRouteToMap(nil)
		assert.NotNil(t, err)
	})
	t.Run("minimal", func(t *testing.T) {
		result, err := vpc.ResourceIBMIsDynamicRouteServerStaticRouteDynamicRouteServerStaticRouteToMap(&vpcv1.DynamicRouteServerStaticRoute{
			ID:          core.StringPtr("sr1"),
			Destination: core.StringPtr("192.168.3.0/24"),
		})
		assert.Nil(t, err)
		assert.Equal(t, "sr1", result["id"])
		assert.Equal(t, "192.168.3.0/24", result["destination"])
		assert.Equal(t, 0, result["route_delete_delay"])
		assert.Equal(t, []map[string]interface{}{}, result["added_routes"])
		assert.Equal(t, []map[string]interface{}{}, result["next_hops"])
		assert.Equal(t, []map[string]interface{}{}, result["routing_tables"])
		assert.Equal(t, []map[string]interface{}{}, result["lifecycle_reasons"])
		_, hasState := result["lifecycle_state"]
		assert.False(t, hasState)
	})
	t.Run("full", func(t *testing.T) {
		result, err := vpc.ResourceIBMIsDynamicRouteServerStaticRouteDynamicRouteServerStaticRouteToMap(&vpcv1.DynamicRouteServerStaticRoute{
			ID:               core.StringPtr("sr1"),
			Href:             core.StringPtr("https://us-south.iaas.cloud.ibm.com/v1/dynamic_route_servers/r006/static_routes/sr1"),
			Name:             core.StringPtr("my-static-route"),
			Destination:      core.StringPtr("192.168.3.0/24"),
			LifecycleState:   core.StringPtr("stable"),
			LifecycleReasons: []vpcv1.LifecycleReason{{Code: core.StringPtr("internal_error"), Message: core.StringPtr("internal error")}},
			ResourceType:     core.StringPtr("dynamic_route_server_static_route"),
			RouteDeleteDelay: core.Int64Ptr(0),
			AddedRoutes: []vpcv1.DynamicRouteServerStaticRouteAddedRoute{
				{VPCRoutingTableRoute: &vpcv1.RouteReference{ID: core.StringPtr("route1"), Name: core.StringPtr("route1")}},
				{},
			},
			NextHops: []vpcv1.DynamicRouteServerStaticRouteNextHopIntf{
				&vpcv1.DynamicRouteServerStaticRouteNextHopDynamicRouteServerPeerGroupReference{ID: core.StringPtr("pg1")},
				&vpcv1.DynamicRouteServerStaticRouteNextHop{ID: core.StringPtr("pg2")},
			},
			RoutingTables: []vpcv1.DynamicRouteServerStaticRouteRoutingTable{
				{Advertise: core.BoolPtr(true), VPCRoutingTable: drsStaticRouteTestRoutingTableReference()},
				{Advertise: core.BoolPtr(false), VPCRoutingTable: drsStaticRouteTestRoutingTableReference()},
			},
		})
		assert.Nil(t, err)
		assert.Equal(t, 0, result["route_delete_delay"])
		assert.Equal(t, "stable", result["lifecycle_state"])
		assert.Len(t, result["lifecycle_reasons"], 1)
		addedRoutes := result["added_routes"].([]map[string]interface{})
		assert.Len(t, addedRoutes, 2)
		assert.Equal(t, "route1", addedRoutes[0]["vpc_routing_table_route"].([]map[string]interface{})[0]["id"])
		assert.Equal(t, map[string]interface{}{}, addedRoutes[1])
		nextHops := result["next_hops"].([]map[string]interface{})
		assert.Equal(t, "pg1", nextHops[0]["id"])
		assert.Equal(t, "pg2", nextHops[1]["id"])
		routingTables := result["routing_tables"].([]map[string]interface{})
		assert.Equal(t, true, routingTables[0]["advertise"])
		assert.Equal(t, false, routingTables[1]["advertise"])
	})
	t.Run("unrecognized next hop subtype propagates error", func(t *testing.T) {
		_, err := vpc.ResourceIBMIsDynamicRouteServerStaticRouteDynamicRouteServerStaticRouteToMap(&vpcv1.DynamicRouteServerStaticRoute{
			NextHops: []vpcv1.DynamicRouteServerStaticRouteNextHopIntf{nil},
		})
		assert.NotNil(t, err)
	})
}

func TestResourceIBMIsDynamicRouteServerStaticRouteDynamicRouteServerStaticRouteNextHopToMap(t *testing.T) {
	t.Run("peer group reference subtype", func(t *testing.T) {
		result, err := vpc.ResourceIBMIsDynamicRouteServerStaticRouteDynamicRouteServerStaticRouteNextHopToMap(&vpcv1.DynamicRouteServerStaticRouteNextHopDynamicRouteServerPeerGroupReference{
			Href:         core.StringPtr("href"),
			ID:           core.StringPtr("pg1"),
			Name:         core.StringPtr("name"),
			ResourceType: core.StringPtr("dynamic_route_server_peer_group"),
			Deleted:      &vpcv1.Deleted{MoreInfo: core.StringPtr("info")},
		})
		assert.Nil(t, err)
		assert.Equal(t, map[string]interface{}{"href": "href", "id": "pg1", "name": "name", "resource_type": "dynamic_route_server_peer_group", "deleted": []map[string]interface{}{{"more_info": "info"}}}, result)
	})
	t.Run("base struct", func(t *testing.T) {
		result, err := vpc.ResourceIBMIsDynamicRouteServerStaticRouteDynamicRouteServerStaticRouteNextHopToMap(&vpcv1.DynamicRouteServerStaticRouteNextHop{ID: core.StringPtr("pg1")})
		assert.Nil(t, err)
		assert.Equal(t, map[string]interface{}{"id": "pg1"}, result)
	})
	t.Run("nil pointers and unrecognized", func(t *testing.T) {
		var reference *vpcv1.DynamicRouteServerStaticRouteNextHopDynamicRouteServerPeerGroupReference
		_, err := vpc.ResourceIBMIsDynamicRouteServerStaticRouteDynamicRouteServerStaticRouteNextHopToMap(reference)
		assert.NotNil(t, err)
		var base *vpcv1.DynamicRouteServerStaticRouteNextHop
		_, err = vpc.ResourceIBMIsDynamicRouteServerStaticRouteDynamicRouteServerStaticRouteNextHopToMap(base)
		assert.NotNil(t, err)
		_, err = vpc.ResourceIBMIsDynamicRouteServerStaticRouteDynamicRouteServerStaticRouteNextHopToMap(nil)
		assert.NotNil(t, err)
	})
}

func TestResourceIBMIsDynamicRouteServerStaticRouteDeletedToMap(t *testing.T) {
	_, err := vpc.ResourceIBMIsDynamicRouteServerStaticRouteDeletedToMap(nil)
	assert.NotNil(t, err)
	result, err := vpc.ResourceIBMIsDynamicRouteServerStaticRouteDeletedToMap(&vpcv1.Deleted{MoreInfo: core.StringPtr("info")})
	assert.Nil(t, err)
	assert.Equal(t, map[string]interface{}{"more_info": "info"}, result)
}

func TestResourceIBMIsDynamicRouteServerStaticRouteDynamicRouteServerStaticRouteNextHopDynamicRouteServerPeerGroupReferenceToMap(t *testing.T) {
	_, err := vpc.ResourceIBMIsDynamicRouteServerStaticRouteDynamicRouteServerStaticRouteNextHopDynamicRouteServerPeerGroupReferenceToMap(nil)
	assert.NotNil(t, err)
	result, err := vpc.ResourceIBMIsDynamicRouteServerStaticRouteDynamicRouteServerStaticRouteNextHopDynamicRouteServerPeerGroupReferenceToMap(&vpcv1.DynamicRouteServerStaticRouteNextHopDynamicRouteServerPeerGroupReference{})
	assert.Nil(t, err)
	assert.Equal(t, map[string]interface{}{}, result)
}

func TestResourceIBMIsDynamicRouteServerStaticRouteDynamicRouteServerStaticRouteRoutingTableToMap(t *testing.T) {
	_, err := vpc.ResourceIBMIsDynamicRouteServerStaticRouteDynamicRouteServerStaticRouteRoutingTableToMap(nil)
	assert.NotNil(t, err)

	result, err := vpc.ResourceIBMIsDynamicRouteServerStaticRouteDynamicRouteServerStaticRouteRoutingTableToMap(&vpcv1.DynamicRouteServerStaticRouteRoutingTable{})
	assert.Nil(t, err)
	assert.Equal(t, map[string]interface{}{"advertise": false}, result)

	result, err = vpc.ResourceIBMIsDynamicRouteServerStaticRouteDynamicRouteServerStaticRouteRoutingTableToMap(&vpcv1.DynamicRouteServerStaticRouteRoutingTable{Advertise: core.BoolPtr(true), VPCRoutingTable: drsStaticRouteTestRoutingTableReference()})
	assert.Nil(t, err)
	assert.Equal(t, true, result["advertise"])
	assert.Equal(t, "r006-6885", result["vpc_routing_table"].([]map[string]interface{})[0]["id"])
}

func TestResourceIBMIsDynamicRouteServerStaticRouteRoutingTableReferenceToMap(t *testing.T) {
	_, err := vpc.ResourceIBMIsDynamicRouteServerStaticRouteRoutingTableReferenceToMap(nil)
	assert.NotNil(t, err)
	model := drsStaticRouteTestRoutingTableReference()
	model.Deleted = &vpcv1.Deleted{MoreInfo: core.StringPtr("info")}
	result, err := vpc.ResourceIBMIsDynamicRouteServerStaticRouteRoutingTableReferenceToMap(model)
	assert.Nil(t, err)
	assert.Equal(t, "my-routing-table-1", result["name"])
	assert.Equal(t, "routing_table", result["resource_type"])
	assert.Equal(t, []map[string]interface{}{{"more_info": "info"}}, result["deleted"])
}

func TestResourceIBMIsDynamicRouteServerStaticRouteDynamicRouteServerStaticRouteAddedRouteToMap(t *testing.T) {
	_, err := vpc.ResourceIBMIsDynamicRouteServerStaticRouteDynamicRouteServerStaticRouteAddedRouteToMap(nil)
	assert.NotNil(t, err)

	result, err := vpc.ResourceIBMIsDynamicRouteServerStaticRouteDynamicRouteServerStaticRouteAddedRouteToMap(&vpcv1.DynamicRouteServerStaticRouteAddedRoute{})
	assert.Nil(t, err)
	assert.Equal(t, map[string]interface{}{}, result)

	result, err = vpc.ResourceIBMIsDynamicRouteServerStaticRouteDynamicRouteServerStaticRouteAddedRouteToMap(&vpcv1.DynamicRouteServerStaticRouteAddedRoute{
		VPCRoutingTableRoute: &vpcv1.RouteReference{Href: core.StringPtr("href"), ID: core.StringPtr("route1"), Name: core.StringPtr("route1"), Deleted: &vpcv1.Deleted{MoreInfo: core.StringPtr("info")}},
	})
	assert.Nil(t, err)
	route := result["vpc_routing_table_route"].([]map[string]interface{})[0]
	assert.Equal(t, "route1", route["id"])
	assert.Equal(t, []map[string]interface{}{{"more_info": "info"}}, route["deleted"])
}

func TestResourceIBMIsDynamicRouteServerStaticRouteRouteReferenceToMap(t *testing.T) {
	_, err := vpc.ResourceIBMIsDynamicRouteServerStaticRouteRouteReferenceToMap(nil)
	assert.NotNil(t, err)
	result, err := vpc.ResourceIBMIsDynamicRouteServerStaticRouteRouteReferenceToMap(&vpcv1.RouteReference{ID: core.StringPtr("route1")})
	assert.Nil(t, err)
	assert.Equal(t, map[string]interface{}{"id": "route1"}, result)
}

func TestResourceIBMIsDynamicRouteServerStaticRouteLifecycleReasonToMap(t *testing.T) {
	_, err := vpc.ResourceIBMIsDynamicRouteServerStaticRouteLifecycleReasonToMap(nil)
	assert.NotNil(t, err)
	result, err := vpc.ResourceIBMIsDynamicRouteServerStaticRouteLifecycleReasonToMap(&vpcv1.LifecycleReason{Code: core.StringPtr("internal_error"), Message: core.StringPtr("msg")})
	assert.Nil(t, err)
	assert.Equal(t, map[string]interface{}{"code": "internal_error", "message": "msg"}, result)
	result, err = vpc.ResourceIBMIsDynamicRouteServerStaticRouteLifecycleReasonToMap(&vpcv1.LifecycleReason{Code: core.StringPtr("internal_error"), Message: core.StringPtr("msg"), MoreInfo: core.StringPtr("info")})
	assert.Nil(t, err)
	assert.Equal(t, "info", result["more_info"])
}

func TestResourceIBMIsDynamicRouteServerStaticRouteMapToDynamicRouteServerStaticRouteNextHopPrototype(t *testing.T) {
	result, err := vpc.ResourceIBMIsDynamicRouteServerStaticRouteMapToDynamicRouteServerStaticRouteNextHopPrototype(map[string]interface{}{"id": "pg1"})
	assert.Nil(t, err)
	assert.Equal(t, &vpcv1.DynamicRouteServerStaticRouteNextHopPrototypeDynamicRouteServerPeerGroupIdentityDynamicRouteServerPeerGroupIdentityByID{ID: core.StringPtr("pg1")}, result)

	result, err = vpc.ResourceIBMIsDynamicRouteServerStaticRouteMapToDynamicRouteServerStaticRouteNextHopPrototype(map[string]interface{}{"id": "", "href": "https://pg"})
	assert.Nil(t, err)
	assert.Equal(t, &vpcv1.DynamicRouteServerStaticRouteNextHopPrototypeDynamicRouteServerPeerGroupIdentityDynamicRouteServerPeerGroupIdentityByHref{Href: core.StringPtr("https://pg")}, result)

	_, err = vpc.ResourceIBMIsDynamicRouteServerStaticRouteMapToDynamicRouteServerStaticRouteNextHopPrototype(map[string]interface{}{"name": "x"})
	assert.NotNil(t, err)
	_, err = vpc.ResourceIBMIsDynamicRouteServerStaticRouteMapToDynamicRouteServerStaticRouteNextHopPrototype(nil)
	assert.NotNil(t, err)
}

func TestResourceIBMIsDynamicRouteServerStaticRouteMapToDynamicRouteServerStaticRouteNextHopPrototypeDynamicRouteServerPeerGroupIdentity(t *testing.T) {
	result, err := vpc.ResourceIBMIsDynamicRouteServerStaticRouteMapToDynamicRouteServerStaticRouteNextHopPrototypeDynamicRouteServerPeerGroupIdentity(map[string]interface{}{"id": "pg1", "href": "https://pg"})
	assert.Nil(t, err)
	_, isByID := result.(*vpcv1.DynamicRouteServerStaticRouteNextHopPrototypeDynamicRouteServerPeerGroupIdentityDynamicRouteServerPeerGroupIdentityByID)
	assert.True(t, isByID)
	_, err = vpc.ResourceIBMIsDynamicRouteServerStaticRouteMapToDynamicRouteServerStaticRouteNextHopPrototypeDynamicRouteServerPeerGroupIdentity(map[string]interface{}{})
	assert.NotNil(t, err)
}

func TestResourceIBMIsDynamicRouteServerStaticRouteMapToDynamicRouteServerStaticRouteNextHopPrototypeDynamicRouteServerPeerGroupIdentityDynamicRouteServerPeerGroupIdentityByID(t *testing.T) {
	_, err := vpc.ResourceIBMIsDynamicRouteServerStaticRouteMapToDynamicRouteServerStaticRouteNextHopPrototypeDynamicRouteServerPeerGroupIdentityDynamicRouteServerPeerGroupIdentityByID(map[string]interface{}{})
	assert.NotNil(t, err)
	result, err := vpc.ResourceIBMIsDynamicRouteServerStaticRouteMapToDynamicRouteServerStaticRouteNextHopPrototypeDynamicRouteServerPeerGroupIdentityDynamicRouteServerPeerGroupIdentityByID(map[string]interface{}{"id": "pg1"})
	assert.Nil(t, err)
	assert.Equal(t, "pg1", *result.ID)
}

func TestResourceIBMIsDynamicRouteServerStaticRouteMapToDynamicRouteServerStaticRouteNextHopPrototypeDynamicRouteServerPeerGroupIdentityDynamicRouteServerPeerGroupIdentityByHref(t *testing.T) {
	_, err := vpc.ResourceIBMIsDynamicRouteServerStaticRouteMapToDynamicRouteServerStaticRouteNextHopPrototypeDynamicRouteServerPeerGroupIdentityDynamicRouteServerPeerGroupIdentityByHref(map[string]interface{}{"href": ""})
	assert.NotNil(t, err)
	result, err := vpc.ResourceIBMIsDynamicRouteServerStaticRouteMapToDynamicRouteServerStaticRouteNextHopPrototypeDynamicRouteServerPeerGroupIdentityDynamicRouteServerPeerGroupIdentityByHref(map[string]interface{}{"href": "https://pg"})
	assert.Nil(t, err)
	assert.Equal(t, "https://pg", *result.Href)
}

func TestResourceIBMIsDynamicRouteServerStaticRouteMapToDynamicRouteServerStaticRouteRoutingTablePrototype(t *testing.T) {
	t.Run("advertise false with id", func(t *testing.T) {
		result, err := vpc.ResourceIBMIsDynamicRouteServerStaticRouteMapToDynamicRouteServerStaticRouteRoutingTablePrototype(map[string]interface{}{
			"advertise":         false,
			"vpc_routing_table": []interface{}{map[string]interface{}{"id": "rt1"}},
		})
		assert.Nil(t, err)
		assert.Equal(t, false, *result.Advertise)
		assert.Equal(t, &vpcv1.RoutingTableIdentityByID{ID: core.StringPtr("rt1")}, result.VPCRoutingTable)
	})
	t.Run("advertise true with href", func(t *testing.T) {
		result, err := vpc.ResourceIBMIsDynamicRouteServerStaticRouteMapToDynamicRouteServerStaticRouteRoutingTablePrototype(map[string]interface{}{
			"advertise":         true,
			"vpc_routing_table": []interface{}{map[string]interface{}{"href": "https://rt"}},
		})
		assert.Nil(t, err)
		assert.Equal(t, true, *result.Advertise)
		assert.Equal(t, &vpcv1.RoutingTableIdentityByHref{Href: core.StringPtr("https://rt")}, result.VPCRoutingTable)
	})
	t.Run("missing or empty vpc_routing_table", func(t *testing.T) {
		_, err := vpc.ResourceIBMIsDynamicRouteServerStaticRouteMapToDynamicRouteServerStaticRouteRoutingTablePrototype(map[string]interface{}{"advertise": true})
		assert.NotNil(t, err)
		_, err = vpc.ResourceIBMIsDynamicRouteServerStaticRouteMapToDynamicRouteServerStaticRouteRoutingTablePrototype(map[string]interface{}{"advertise": true, "vpc_routing_table": []interface{}{}})
		assert.NotNil(t, err)
		_, err = vpc.ResourceIBMIsDynamicRouteServerStaticRouteMapToDynamicRouteServerStaticRouteRoutingTablePrototype(map[string]interface{}{"advertise": true, "vpc_routing_table": []interface{}{map[string]interface{}{}}})
		assert.NotNil(t, err)
	})
}

func TestResourceIBMIsDynamicRouteServerStaticRouteMapToRoutingTableIdentity(t *testing.T) {
	result, err := vpc.ResourceIBMIsDynamicRouteServerStaticRouteMapToRoutingTableIdentity(map[string]interface{}{"id": "rt1"})
	assert.Nil(t, err)
	assert.Equal(t, &vpcv1.RoutingTableIdentityByID{ID: core.StringPtr("rt1")}, result)

	result, err = vpc.ResourceIBMIsDynamicRouteServerStaticRouteMapToRoutingTableIdentity(map[string]interface{}{"crn": "crn:rt"})
	assert.Nil(t, err)
	assert.Equal(t, &vpcv1.RoutingTableIdentityByCRN{CRN: core.StringPtr("crn:rt")}, result)

	result, err = vpc.ResourceIBMIsDynamicRouteServerStaticRouteMapToRoutingTableIdentity(map[string]interface{}{"href": "https://rt"})
	assert.Nil(t, err)
	assert.Equal(t, &vpcv1.RoutingTableIdentityByHref{Href: core.StringPtr("https://rt")}, result)

	_, err = vpc.ResourceIBMIsDynamicRouteServerStaticRouteMapToRoutingTableIdentity(map[string]interface{}{})
	assert.NotNil(t, err)
}

func TestResourceIBMIsDynamicRouteServerStaticRouteMapToRoutingTableIdentityByID(t *testing.T) {
	_, err := vpc.ResourceIBMIsDynamicRouteServerStaticRouteMapToRoutingTableIdentityByID(map[string]interface{}{})
	assert.NotNil(t, err)
	result, err := vpc.ResourceIBMIsDynamicRouteServerStaticRouteMapToRoutingTableIdentityByID(map[string]interface{}{"id": "rt1"})
	assert.Nil(t, err)
	assert.Equal(t, "rt1", *result.ID)
}

func TestResourceIBMIsDynamicRouteServerStaticRouteMapToRoutingTableIdentityByCRN(t *testing.T) {
	_, err := vpc.ResourceIBMIsDynamicRouteServerStaticRouteMapToRoutingTableIdentityByCRN(map[string]interface{}{})
	assert.NotNil(t, err)
	result, err := vpc.ResourceIBMIsDynamicRouteServerStaticRouteMapToRoutingTableIdentityByCRN(map[string]interface{}{"crn": "crn:rt"})
	assert.Nil(t, err)
	assert.Equal(t, "crn:rt", *result.CRN)
}

func TestResourceIBMIsDynamicRouteServerStaticRouteMapToRoutingTableIdentityByHref(t *testing.T) {
	_, err := vpc.ResourceIBMIsDynamicRouteServerStaticRouteMapToRoutingTableIdentityByHref(map[string]interface{}{})
	assert.NotNil(t, err)
	result, err := vpc.ResourceIBMIsDynamicRouteServerStaticRouteMapToRoutingTableIdentityByHref(map[string]interface{}{"href": "https://rt"})
	assert.Nil(t, err)
	assert.Equal(t, "https://rt", *result.Href)
}
