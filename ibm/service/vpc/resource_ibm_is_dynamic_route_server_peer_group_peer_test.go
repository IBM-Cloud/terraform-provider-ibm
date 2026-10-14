// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/vpc"
	. "github.com/IBM-Cloud/terraform-provider-ibm/ibm/unittest"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/stretchr/testify/assert"
)

const isDynamicRouteServerPeerGroupPeerResource = "ibm_is_dynamic_route_server_peer_group_peer.testacc_peer"

func TestAccIBMIsDynamicRouteServerPeerGroupPeerBasic(t *testing.T) {
	var conf vpcv1.DynamicRouteServerPeerGroupPeerIntf
	vpcname := fmt.Sprintf("tf-vpc-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tf-subnet-%d", acctest.RandIntRange(10, 100))
	vniname := fmt.Sprintf("tf-vni-%d", acctest.RandIntRange(10, 100))
	drsname := fmt.Sprintf("tf-drs-%d", acctest.RandIntRange(10, 100))
	pgname := fmt.Sprintf("tf-peer-group-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIsDynamicRouteServerPeerGroupPeerDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsDynamicRouteServerPeerGroupPeerConfigBasic(vpcname, subnetname, vniname, drsname, pgname),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIsDynamicRouteServerPeerGroupPeerExists(isDynamicRouteServerPeerGroupPeerResource, &conf),
					resource.TestCheckResourceAttr(isDynamicRouteServerPeerGroupPeerResource, "asn", "64514"),
					resource.TestCheckResourceAttr(isDynamicRouteServerPeerGroupPeerResource, "endpoint.0.address", "10.240.0.20"),
					resource.TestCheckResourceAttrSet(isDynamicRouteServerPeerGroupPeerResource, "name"),
					resource.TestCheckResourceAttrSet(isDynamicRouteServerPeerGroupPeerResource, "dynamic_route_server_peer_group_peer_id"),
					resource.TestCheckResourceAttrSet(isDynamicRouteServerPeerGroupPeerResource, "href"),
					resource.TestCheckResourceAttrSet(isDynamicRouteServerPeerGroupPeerResource, "created_at"),
					resource.TestCheckResourceAttr(isDynamicRouteServerPeerGroupPeerResource, "lifecycle_state", "stable"),
					resource.TestCheckResourceAttr(isDynamicRouteServerPeerGroupPeerResource, "resource_type", "dynamic_route_server_peer_group_peer"),
					resource.TestCheckResourceAttrPair(isDynamicRouteServerPeerGroupPeerResource, "dynamic_route_server_peer_group_id", isDynamicRouteServerPeerGroupResource, "dynamic_route_server_peer_group_id"),
				),
			},
		},
	})
}

func TestAccIBMIsDynamicRouteServerPeerGroupPeerAllArgs(t *testing.T) {
	var conf vpcv1.DynamicRouteServerPeerGroupPeerIntf
	vpcname := fmt.Sprintf("tf-vpc-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tf-subnet-%d", acctest.RandIntRange(10, 100))
	vniname := fmt.Sprintf("tf-vni-%d", acctest.RandIntRange(10, 100))
	drsname := fmt.Sprintf("tf-drs-%d", acctest.RandIntRange(10, 100))
	pgname := fmt.Sprintf("tf-peer-group-%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf-peer-%d", acctest.RandIntRange(10, 100))
	nameUpdate := fmt.Sprintf("tf-peer-upd-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIsDynamicRouteServerPeerGroupPeerDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsDynamicRouteServerPeerGroupPeerConfig(vpcname, subnetname, vniname, drsname, pgname, name, 64514, 2, "enabled", "10.240.0.20"),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIsDynamicRouteServerPeerGroupPeerExists(isDynamicRouteServerPeerGroupPeerResource, &conf),
					resource.TestCheckResourceAttr(isDynamicRouteServerPeerGroupPeerResource, "name", name),
					resource.TestCheckResourceAttr(isDynamicRouteServerPeerGroupPeerResource, "asn", "64514"),
					resource.TestCheckResourceAttr(isDynamicRouteServerPeerGroupPeerResource, "priority", "2"),
					resource.TestCheckResourceAttr(isDynamicRouteServerPeerGroupPeerResource, "state", "enabled"),
					resource.TestCheckResourceAttr(isDynamicRouteServerPeerGroupPeerResource, "endpoint.0.address", "10.240.0.20"),
					resource.TestCheckResourceAttr(isDynamicRouteServerPeerGroupPeerResource, "endpoint.0.gateway.0.address", "10.240.0.1"),
				),
			},
			resource.TestStep{
				// Update every patchable field: asn, name, priority and state.
				Config: testAccCheckIBMIsDynamicRouteServerPeerGroupPeerConfig(vpcname, subnetname, vniname, drsname, pgname, nameUpdate, 64515, 3, "disabled", "10.240.0.20"),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIsDynamicRouteServerPeerGroupPeerExists(isDynamicRouteServerPeerGroupPeerResource, &conf),
					resource.TestCheckResourceAttr(isDynamicRouteServerPeerGroupPeerResource, "name", nameUpdate),
					resource.TestCheckResourceAttr(isDynamicRouteServerPeerGroupPeerResource, "asn", "64515"),
					resource.TestCheckResourceAttr(isDynamicRouteServerPeerGroupPeerResource, "priority", "3"),
					resource.TestCheckResourceAttr(isDynamicRouteServerPeerGroupPeerResource, "state", "disabled"),
				),
			},
			resource.TestStep{
				// Changing the endpoint forces a new peer.
				Config: testAccCheckIBMIsDynamicRouteServerPeerGroupPeerConfig(vpcname, subnetname, vniname, drsname, pgname, nameUpdate, 64515, 3, "disabled", "10.240.0.21"),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIsDynamicRouteServerPeerGroupPeerExists(isDynamicRouteServerPeerGroupPeerResource, &conf),
					resource.TestCheckResourceAttr(isDynamicRouteServerPeerGroupPeerResource, "endpoint.0.address", "10.240.0.21"),
				),
			},
			resource.TestStep{
				ResourceName:      isDynamicRouteServerPeerGroupPeerResource,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

// testAccCheckIBMIsDynamicRouteServerPeerGroupPeerParent creates the dependencies and the peer group a peer belongs to.
func testAccCheckIBMIsDynamicRouteServerPeerGroupPeerParent(vpcname, subnetname, vniname, drsname, pgname string) string {
	return testAccCheckIBMIsDynamicRouteServerPeerGroupConfigBasic(vpcname, subnetname, vniname, drsname, pgname)
}

func testAccCheckIBMIsDynamicRouteServerPeerGroupPeerConfigBasic(vpcname, subnetname, vniname, drsname, pgname string) string {
	return testAccCheckIBMIsDynamicRouteServerPeerGroupPeerParent(vpcname, subnetname, vniname, drsname, pgname) + `
	resource "ibm_is_dynamic_route_server_peer_group_peer" "testacc_peer" {
		dynamic_route_server_id            = ibm_is_dynamic_route_server.testacc_drs.id
		dynamic_route_server_peer_group_id = ibm_is_dynamic_route_server_peer_group.testacc_peer_group.dynamic_route_server_peer_group_id
		asn                                = 64514
		endpoint {
			address = "10.240.0.20"
		}
	}
	`
}

func testAccCheckIBMIsDynamicRouteServerPeerGroupPeerConfig(vpcname, subnetname, vniname, drsname, pgname, name string, asn, priority int, state, address string) string {
	return testAccCheckIBMIsDynamicRouteServerPeerGroupPeerParent(vpcname, subnetname, vniname, drsname, pgname) + fmt.Sprintf(`
	resource "ibm_is_dynamic_route_server_peer_group_peer" "testacc_peer" {
		dynamic_route_server_id            = ibm_is_dynamic_route_server.testacc_drs.id
		dynamic_route_server_peer_group_id = ibm_is_dynamic_route_server_peer_group.testacc_peer_group.dynamic_route_server_peer_group_id
		name                               = "%s"
		asn                                = %d
		priority                           = %d
		state                              = "%s"
		endpoint {
			address = "%s"
			gateway {
				address = "10.240.0.1"
			}
		}
	}
	`, name, asn, priority, state, address)
}

func testAccCheckIBMIsDynamicRouteServerPeerGroupPeerExists(n string, obj *vpcv1.DynamicRouteServerPeerGroupPeerIntf) resource.TestCheckFunc {
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
		if len(parts) != 3 {
			return fmt.Errorf("unexpected ID format %q", rs.Primary.ID)
		}

		getOptions := &vpcv1.GetDynamicRouteServerPeerGroupPeerOptions{}
		getOptions.SetDynamicRouteServerID(parts[0])
		getOptions.SetDynamicRouteServerPeerGroupID(parts[1])
		getOptions.SetID(parts[2])

		peer, _, err := vpcClient.GetDynamicRouteServerPeerGroupPeer(getOptions)
		if err != nil {
			return err
		}

		*obj = peer
		return nil
	}
}

func testAccCheckIBMIsDynamicRouteServerPeerGroupPeerDestroy(s *terraform.State) error {
	vpcClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).VpcV1API()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_is_dynamic_route_server_peer_group_peer" {
			continue
		}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}
		if len(parts) != 3 {
			return fmt.Errorf("unexpected ID format %q", rs.Primary.ID)
		}

		getOptions := &vpcv1.GetDynamicRouteServerPeerGroupPeerOptions{}
		getOptions.SetDynamicRouteServerID(parts[0])
		getOptions.SetDynamicRouteServerPeerGroupID(parts[1])
		getOptions.SetID(parts[2])

		_, response, err := vpcClient.GetDynamicRouteServerPeerGroupPeer(getOptions)
		if err == nil {
			return fmt.Errorf("dynamic route server peer group peer still exists: %s", rs.Primary.ID)
		} else if response == nil || response.StatusCode != 404 {
			return fmt.Errorf("Error checking for dynamic route server peer group peer (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}

func TestResourceIBMIsDynamicRouteServerPeerGroupPeerSchemaShapes(t *testing.T) {
	res := vpc.ResourceIBMIsDynamicRouteServerPeerGroupPeer()
	assert.NotNil(t, res.Importer)
	assert.NotNil(t, res.Timeouts)
	assert.True(t, res.Schema["dynamic_route_server_id"].ForceNew)
	assert.True(t, res.Schema["dynamic_route_server_peer_group_id"].ForceNew)
	assert.True(t, res.Schema["endpoint"].ForceNew)
	assert.Equal(t, 1, res.Schema["endpoint"].MaxItems)
	for _, key := range []string{"asn", "name", "priority", "state"} {
		assert.False(t, res.Schema[key].ForceNew, key)
	}
	for _, key := range []string{"created_at", "creator", "href", "lifecycle_state", "status", "resource_type", "sessions", "bidirectional_forwarding_detection", "dynamic_route_server_peer_group_peer_id", "etag"} {
		assert.True(t, res.Schema[key].Computed, key)
		assert.False(t, res.Schema[key].Optional, key)
	}
	endpoint := res.Schema["endpoint"].Elem.(*schema.Resource).Schema
	assert.Equal(t, 1, endpoint["gateway"].MaxItems)
	validator := vpc.ResourceIBMIsDynamicRouteServerPeerGroupPeerValidator()
	found := map[string]bool{}
	for _, entry := range validator.Schema {
		found[entry.Identifier] = true
	}
	for _, key := range []string{"dynamic_route_server_id", "dynamic_route_server_peer_group_id", "name", "priority", "state"} {
		assert.True(t, found[key], key)
	}
}

func TestResourceIBMIsDynamicRouteServerPeerGroupPeerIntfID(t *testing.T) {
	id, err := vpc.ResourceIBMIsDynamicRouteServerPeerGroupPeerIntfID(&vpcv1.DynamicRouteServerPeerGroupPeerAddress{ID: core.StringPtr("a")})
	assert.Nil(t, err)
	assert.Equal(t, "a", id)
	id, err = vpc.ResourceIBMIsDynamicRouteServerPeerGroupPeerIntfID(&vpcv1.DynamicRouteServerPeerGroupPeerTransitGateway{ID: core.StringPtr("b")})
	assert.Nil(t, err)
	assert.Equal(t, "b", id)
	id, err = vpc.ResourceIBMIsDynamicRouteServerPeerGroupPeerIntfID(&vpcv1.DynamicRouteServerPeerGroupPeer{ID: core.StringPtr("c")})
	assert.Nil(t, err)
	assert.Equal(t, "c", id)
	_, err = vpc.ResourceIBMIsDynamicRouteServerPeerGroupPeerIntfID(&vpcv1.DynamicRouteServerPeerGroupPeer{})
	assert.NotNil(t, err)
	_, err = vpc.ResourceIBMIsDynamicRouteServerPeerGroupPeerIntfID(nil)
	assert.NotNil(t, err)
}

func TestResourceIBMIsDynamicRouteServerPeerGroupPeerIntfLifecycleState(t *testing.T) {
	assert.Equal(t, "stable", vpc.ResourceIBMIsDynamicRouteServerPeerGroupPeerIntfLifecycleState(&vpcv1.DynamicRouteServerPeerGroupPeerAddress{LifecycleState: core.StringPtr("stable")}))
	assert.Equal(t, "pending", vpc.ResourceIBMIsDynamicRouteServerPeerGroupPeerIntfLifecycleState(&vpcv1.DynamicRouteServerPeerGroupPeerTransitGateway{LifecycleState: core.StringPtr("pending")}))
	assert.Equal(t, "failed", vpc.ResourceIBMIsDynamicRouteServerPeerGroupPeerIntfLifecycleState(&vpcv1.DynamicRouteServerPeerGroupPeer{LifecycleState: core.StringPtr("failed")}))
	assert.Equal(t, "", vpc.ResourceIBMIsDynamicRouteServerPeerGroupPeerIntfLifecycleState(&vpcv1.DynamicRouteServerPeerGroupPeer{}))
	assert.Equal(t, "", vpc.ResourceIBMIsDynamicRouteServerPeerGroupPeerIntfLifecycleState(nil))
}

func TestResourceIBMIsDynamicRouteServerPeerGroupPeerMapToDynamicRouteServerPeerGroupPeerPrototype(t *testing.T) {
	t.Run("all fields with IP gateway endpoint", func(t *testing.T) {
		result, err := vpc.ResourceIBMIsDynamicRouteServerPeerGroupPeerMapToDynamicRouteServerPeerGroupPeerPrototype(map[string]interface{}{
			"asn":      64513,
			"name":     "my-peer",
			"priority": 0,
			"state":    "enabled",
			"endpoint": []interface{}{map[string]interface{}{"address": "10.240.0.10", "gateway": []interface{}{map[string]interface{}{"address": "10.240.0.1"}}}},
		})
		assert.Nil(t, err)
		model := result.(*vpcv1.DynamicRouteServerPeerGroupPeerPrototypeDynamicRouteServerPeerGroupPeerAddressPrototype)
		assert.Equal(t, int64(64513), *model.Asn)
		assert.Equal(t, "my-peer", *model.Name)
		assert.Equal(t, int64(0), *model.Priority, "priority 0 is sent when present")
		assert.Equal(t, "enabled", *model.State)
		endpoint := model.Endpoint.(*vpcv1.DynamicRouteServerPeerGroupPeerAddressEndpointPrototypeDynamicRouteServerPeerGroupPeerAddressEndpointIPGatewayPrototype)
		assert.Equal(t, "10.240.0.10", *endpoint.Address)
		assert.Equal(t, "10.240.0.1", *endpoint.Gateway.Address)
	})
	t.Run("optional fields unset", func(t *testing.T) {
		result, err := vpc.ResourceIBMIsDynamicRouteServerPeerGroupPeerMapToDynamicRouteServerPeerGroupPeerPrototype(map[string]interface{}{
			"asn":      64513,
			"name":     "",
			"state":    "",
			"endpoint": []interface{}{map[string]interface{}{"id": "rip-id"}},
		})
		assert.Nil(t, err)
		model := result.(*vpcv1.DynamicRouteServerPeerGroupPeerPrototypeDynamicRouteServerPeerGroupPeerAddressPrototype)
		assert.Nil(t, model.Name)
		assert.Nil(t, model.Priority)
		assert.Nil(t, model.State)
		_, isByID := model.Endpoint.(*vpcv1.DynamicRouteServerPeerGroupPeerAddressEndpointPrototypeReservedIPIdentityByID)
		assert.True(t, isByID)
	})
	t.Run("missing asn is an error", func(t *testing.T) {
		_, err := vpc.ResourceIBMIsDynamicRouteServerPeerGroupPeerMapToDynamicRouteServerPeerGroupPeerPrototype(map[string]interface{}{"endpoint": []interface{}{map[string]interface{}{"id": "rip-id"}}})
		assert.NotNil(t, err)
	})
	t.Run("missing or empty endpoint is an error", func(t *testing.T) {
		_, err := vpc.ResourceIBMIsDynamicRouteServerPeerGroupPeerMapToDynamicRouteServerPeerGroupPeerPrototype(map[string]interface{}{"asn": 1})
		assert.NotNil(t, err)
		_, err = vpc.ResourceIBMIsDynamicRouteServerPeerGroupPeerMapToDynamicRouteServerPeerGroupPeerPrototype(map[string]interface{}{"asn": 1, "endpoint": []interface{}{}})
		assert.NotNil(t, err)
		_, err = vpc.ResourceIBMIsDynamicRouteServerPeerGroupPeerMapToDynamicRouteServerPeerGroupPeerPrototype(map[string]interface{}{"asn": 1, "endpoint": []interface{}{nil}})
		assert.NotNil(t, err)
		_, err = vpc.ResourceIBMIsDynamicRouteServerPeerGroupPeerMapToDynamicRouteServerPeerGroupPeerPrototype(map[string]interface{}{"asn": 1, "endpoint": []interface{}{map[string]interface{}{}}})
		assert.NotNil(t, err)
	})
	t.Run("nil map is an error", func(t *testing.T) {
		_, err := vpc.ResourceIBMIsDynamicRouteServerPeerGroupPeerMapToDynamicRouteServerPeerGroupPeerPrototype(nil)
		assert.NotNil(t, err)
	})
}

func TestResourceIBMIsDynamicRouteServerPeerGroupPeerMapToDynamicRouteServerPeerGroupPeerAddressEndpointPrototype(t *testing.T) {
	t.Run("by id", func(t *testing.T) {
		result, err := vpc.ResourceIBMIsDynamicRouteServerPeerGroupPeerMapToDynamicRouteServerPeerGroupPeerAddressEndpointPrototype(map[string]interface{}{"id": "rip-id", "href": "", "address": ""})
		assert.Nil(t, err)
		model := result.(*vpcv1.DynamicRouteServerPeerGroupPeerAddressEndpointPrototypeReservedIPIdentityByID)
		assert.Equal(t, "rip-id", *model.ID)
	})
	t.Run("by href", func(t *testing.T) {
		result, err := vpc.ResourceIBMIsDynamicRouteServerPeerGroupPeerMapToDynamicRouteServerPeerGroupPeerAddressEndpointPrototype(map[string]interface{}{"href": "https://example.com/rip"})
		assert.Nil(t, err)
		model := result.(*vpcv1.DynamicRouteServerPeerGroupPeerAddressEndpointPrototypeReservedIPIdentityByHref)
		assert.Equal(t, "https://example.com/rip", *model.Href)
	})
	t.Run("address without gateway", func(t *testing.T) {
		result, err := vpc.ResourceIBMIsDynamicRouteServerPeerGroupPeerMapToDynamicRouteServerPeerGroupPeerAddressEndpointPrototype(map[string]interface{}{"address": "10.240.0.10", "gateway": []interface{}{}})
		assert.Nil(t, err)
		model := result.(*vpcv1.DynamicRouteServerPeerGroupPeerAddressEndpointPrototypeDynamicRouteServerPeerGroupPeerAddressEndpointIPGatewayPrototype)
		assert.Equal(t, "10.240.0.10", *model.Address)
		assert.Nil(t, model.Gateway)
	})
	t.Run("address with gateway", func(t *testing.T) {
		result, err := vpc.ResourceIBMIsDynamicRouteServerPeerGroupPeerMapToDynamicRouteServerPeerGroupPeerAddressEndpointPrototype(map[string]interface{}{"address": "10.240.0.10", "gateway": []interface{}{map[string]interface{}{"address": "10.240.0.1"}}})
		assert.Nil(t, err)
		model := result.(*vpcv1.DynamicRouteServerPeerGroupPeerAddressEndpointPrototypeDynamicRouteServerPeerGroupPeerAddressEndpointIPGatewayPrototype)
		assert.Equal(t, "10.240.0.1", *model.Gateway.Address)
	})
	t.Run("id and address together is an error", func(t *testing.T) {
		_, err := vpc.ResourceIBMIsDynamicRouteServerPeerGroupPeerMapToDynamicRouteServerPeerGroupPeerAddressEndpointPrototype(map[string]interface{}{"id": "rip-id", "address": "10.240.0.10"})
		assert.NotNil(t, err)
	})
	t.Run("gateway without address is an error", func(t *testing.T) {
		_, err := vpc.ResourceIBMIsDynamicRouteServerPeerGroupPeerMapToDynamicRouteServerPeerGroupPeerAddressEndpointPrototype(map[string]interface{}{"address": "10.240.0.10", "gateway": []interface{}{map[string]interface{}{"address": ""}}})
		assert.NotNil(t, err)
	})
	t.Run("nothing set is an error", func(t *testing.T) {
		_, err := vpc.ResourceIBMIsDynamicRouteServerPeerGroupPeerMapToDynamicRouteServerPeerGroupPeerAddressEndpointPrototype(map[string]interface{}{})
		assert.NotNil(t, err)
		_, err = vpc.ResourceIBMIsDynamicRouteServerPeerGroupPeerMapToDynamicRouteServerPeerGroupPeerAddressEndpointPrototype(nil)
		assert.NotNil(t, err)
	})
}

func TestResourceIBMIsDynamicRouteServerPeerGroupPeerMapToDynamicRouteServerPeerGroupPeerAddressEndpointPrototypeReservedIPIdentityByID(t *testing.T) {
	result, err := vpc.ResourceIBMIsDynamicRouteServerPeerGroupPeerMapToDynamicRouteServerPeerGroupPeerAddressEndpointPrototypeReservedIPIdentityByID(map[string]interface{}{"id": "rip-id"})
	assert.Nil(t, err)
	assert.Equal(t, "rip-id", *result.ID)
	_, err = vpc.ResourceIBMIsDynamicRouteServerPeerGroupPeerMapToDynamicRouteServerPeerGroupPeerAddressEndpointPrototypeReservedIPIdentityByID(map[string]interface{}{"id": ""})
	assert.NotNil(t, err)
}

func TestResourceIBMIsDynamicRouteServerPeerGroupPeerMapToDynamicRouteServerPeerGroupPeerAddressEndpointPrototypeReservedIPIdentityByHref(t *testing.T) {
	result, err := vpc.ResourceIBMIsDynamicRouteServerPeerGroupPeerMapToDynamicRouteServerPeerGroupPeerAddressEndpointPrototypeReservedIPIdentityByHref(map[string]interface{}{"href": "https://example.com/rip"})
	assert.Nil(t, err)
	assert.Equal(t, "https://example.com/rip", *result.Href)
	_, err = vpc.ResourceIBMIsDynamicRouteServerPeerGroupPeerMapToDynamicRouteServerPeerGroupPeerAddressEndpointPrototypeReservedIPIdentityByHref(map[string]interface{}{})
	assert.NotNil(t, err)
}

func TestResourceIBMIsDynamicRouteServerPeerGroupPeerMapToDynamicRouteServerPeerGroupPeerAddressEndpointPrototypeDynamicRouteServerPeerGroupPeerAddressEndpointIPGatewayPrototype(t *testing.T) {
	result, err := vpc.ResourceIBMIsDynamicRouteServerPeerGroupPeerMapToDynamicRouteServerPeerGroupPeerAddressEndpointPrototypeDynamicRouteServerPeerGroupPeerAddressEndpointIPGatewayPrototype(map[string]interface{}{"address": "10.240.0.10"})
	assert.Nil(t, err)
	assert.Equal(t, "10.240.0.10", *result.Address)
	assert.Nil(t, result.Gateway)
	_, err = vpc.ResourceIBMIsDynamicRouteServerPeerGroupPeerMapToDynamicRouteServerPeerGroupPeerAddressEndpointPrototypeDynamicRouteServerPeerGroupPeerAddressEndpointIPGatewayPrototype(map[string]interface{}{"address": ""})
	assert.NotNil(t, err)
}

func TestResourceIBMIsDynamicRouteServerPeerGroupPeerMapToDynamicRouteServerPeerGroupPeerAddressEndpointIPGatewayGateway(t *testing.T) {
	result, err := vpc.ResourceIBMIsDynamicRouteServerPeerGroupPeerMapToDynamicRouteServerPeerGroupPeerAddressEndpointIPGatewayGateway(map[string]interface{}{"address": "10.240.0.1"})
	assert.Nil(t, err)
	assert.Equal(t, "10.240.0.1", *result.Address)
	_, err = vpc.ResourceIBMIsDynamicRouteServerPeerGroupPeerMapToDynamicRouteServerPeerGroupPeerAddressEndpointIPGatewayGateway(map[string]interface{}{})
	assert.NotNil(t, err)
}

func testPeerGroupPeerMemberReference() *vpcv1.DynamicRouteServerMemberReference {
	return &vpcv1.DynamicRouteServerMemberReference{
		Href:         core.StringPtr("https://example.com/member"),
		ID:           core.StringPtr("member-id"),
		Name:         core.StringPtr("member"),
		ResourceType: core.StringPtr("dynamic_route_server_member"),
		VirtualNetworkInterfaces: []vpcv1.VirtualNetworkInterfaceReference{{
			CRN:          core.StringPtr("crn:v1:bluemix:public:is:us-south-1:a/123::virtual-network-interface:vni"),
			Href:         core.StringPtr("https://example.com/vni"),
			ID:           core.StringPtr("vni-id"),
			Name:         core.StringPtr("vni"),
			ResourceType: core.StringPtr("virtual_network_interface"),
		}},
	}
}

func TestResourceIBMIsDynamicRouteServerPeerGroupPeerDynamicRouteServerPeerGroupPeerToMap(t *testing.T) {
	t.Run("address subtype with everything set", func(t *testing.T) {
		model := &vpcv1.DynamicRouteServerPeerGroupPeerAddress{
			CreatedAt:        CreateMockDateTime("2026-01-02T03:04:05.000Z"),
			Creator:          core.StringPtr("dynamic_route_server"),
			Href:             core.StringPtr("https://example.com/peer"),
			ID:               core.StringPtr("peer-id"),
			LifecycleReasons: []vpcv1.LifecycleReason{{Code: core.StringPtr("internal_error"), Message: core.StringPtr("msg")}},
			LifecycleState:   core.StringPtr("stable"),
			Name:             core.StringPtr("my-peer"),
			ResourceType:     core.StringPtr("dynamic_route_server_peer_group_peer"),
			Status:           core.StringPtr("up"),
			StatusReasons:    []vpcv1.DynamicRouteServerPeerGroupPeerStatusReason{{Code: core.StringPtr("c"), Message: core.StringPtr("m"), MoreInfo: core.StringPtr("i")}},
			Asn:              core.Int64Ptr(64513),
			Priority:         core.Int64Ptr(2),
			State:            core.StringPtr("enabled"),
			BidirectionalForwardingDetection: &vpcv1.DynamicRouteServerPeerGroupPeerAddressBidirectionalForwardingDetection{
				DetectMultiplier: core.Int64Ptr(3), Enabled: core.BoolPtr(false), Mode: core.StringPtr("asynchronous"), ReceiveInterval: core.Int64Ptr(300), Role: core.StringPtr("active"), TransmitInterval: core.Int64Ptr(300),
				Sessions: []vpcv1.DynamicRouteServerPeerGroupPeerAddressBidirectionalForwardingDetectionSession{{Local: testPeerGroupPeerMemberReference(), Remote: &vpcv1.DynamicRouteServerPeerGroupPeerReference{ID: core.StringPtr("remote-id")}, State: core.StringPtr("up")}},
			},
			Endpoint: &vpcv1.DynamicRouteServerPeerGroupPeerAddressEndpointByReservedIP{ID: core.StringPtr("rip-id"), Address: core.StringPtr("10.240.0.10"), Gateway: &vpcv1.IP{Address: core.StringPtr("10.240.0.1")}},
			Sessions: []vpcv1.DynamicRouteServerPeerGroupPeerAddressSession{
				{EstablishedAt: CreateMockDateTime("2026-01-02T03:04:05.000Z"), Local: testPeerGroupPeerMemberReference(), ProtocolState: core.StringPtr("established"), Remote: &vpcv1.DynamicRouteServerPeerGroupPeerName{Name: core.StringPtr("remote")}},
				{ProtocolState: core.StringPtr("idle")},
			},
		}
		result, err := vpc.ResourceIBMIsDynamicRouteServerPeerGroupPeerDynamicRouteServerPeerGroupPeerToMap(model)
		assert.Nil(t, err)
		assert.Equal(t, "peer-id", result["id"])
		assert.Equal(t, "dynamic_route_server", result["creator"])
		assert.Equal(t, "2026-01-02T03:04:05.000Z", result["created_at"])
		assert.Equal(t, 64513, result["asn"])
		assert.Equal(t, 2, result["priority"])
		assert.Equal(t, "enabled", result["state"])
		assert.Equal(t, "up", result["status"])
		assert.Equal(t, []map[string]interface{}{{"code": "internal_error", "message": "msg"}}, result["lifecycle_reasons"])
		assert.Equal(t, []map[string]interface{}{{"code": "c", "message": "m", "more_info": "i"}}, result["status_reasons"])
		bfd := result["bidirectional_forwarding_detection"].([]map[string]interface{})[0]
		assert.Equal(t, false, bfd["enabled"], "false is kept, not dropped")
		assert.Equal(t, 3, bfd["detect_multiplier"])
		bfdSession := bfd["sessions"].([]map[string]interface{})[0]
		assert.Equal(t, "up", bfdSession["state"])
		assert.Equal(t, "member-id", bfdSession["local"].([]map[string]interface{})[0]["id"])
		assert.Equal(t, "remote-id", bfdSession["remote"].([]map[string]interface{})[0]["id"])
		endpoint := result["endpoint"].([]map[string]interface{})[0]
		assert.Equal(t, "rip-id", endpoint["id"])
		assert.Equal(t, []map[string]interface{}{{"address": "10.240.0.1"}}, endpoint["gateway"])
		sessions := result["sessions"].([]map[string]interface{})
		assert.Len(t, sessions, 2)
		assert.Equal(t, "2026-01-02T03:04:05.000Z", sessions[0]["established_at"])
		assert.Equal(t, []map[string]interface{}{{"name": "remote"}}, sessions[0]["remote"])
		assert.Equal(t, map[string]interface{}{"protocol_state": "idle"}, sessions[1])
	})
	t.Run("transit gateway subtype", func(t *testing.T) {
		model := &vpcv1.DynamicRouteServerPeerGroupPeerTransitGateway{
			ID:       core.StringPtr("tgw-peer-id"),
			Creator:  core.StringPtr("transit_gateway"),
			Endpoint: &vpcv1.DynamicRouteServerPeerGroupPeerTransitGatewayEndpoint{CRN: core.StringPtr("crn:tgw"), ID: core.StringPtr("tgw-id"), ResourceType: core.StringPtr("transit_gateway")},
			Sessions: []vpcv1.DynamicRouteServerPeerGroupPeerTransitGatewaySession{{ProtocolState: core.StringPtr("established"), Remote: &vpcv1.DynamicRouteServerPeerGroupPeerTransitGatewaySessionRemote{Name: core.StringPtr("tgw")}}},
		}
		result, err := vpc.ResourceIBMIsDynamicRouteServerPeerGroupPeerDynamicRouteServerPeerGroupPeerToMap(model)
		assert.Nil(t, err)
		assert.Equal(t, "transit_gateway", result["creator"])
		_, hasAsn := result["asn"]
		assert.False(t, hasAsn)
		assert.Equal(t, []map[string]interface{}{{"crn": "crn:tgw", "id": "tgw-id", "resource_type": "transit_gateway"}}, result["endpoint"])
		assert.Equal(t, []map[string]interface{}{{"protocol_state": "established", "remote": []map[string]interface{}{{"name": "tgw"}}}}, result["sessions"])
	})
	t.Run("base struct with no sessions gives empty lists", func(t *testing.T) {
		result, err := vpc.ResourceIBMIsDynamicRouteServerPeerGroupPeerDynamicRouteServerPeerGroupPeerToMap(&vpcv1.DynamicRouteServerPeerGroupPeer{ID: core.StringPtr("id")})
		assert.Nil(t, err)
		assert.Equal(t, []map[string]interface{}{}, result["sessions"])
		assert.Equal(t, []map[string]interface{}{}, result["lifecycle_reasons"])
		assert.Equal(t, []map[string]interface{}{}, result["status_reasons"])
		_, hasEndpoint := result["endpoint"]
		assert.False(t, hasEndpoint)
	})
	t.Run("base struct with base endpoint", func(t *testing.T) {
		result, err := vpc.ResourceIBMIsDynamicRouteServerPeerGroupPeerDynamicRouteServerPeerGroupPeerToMap(&vpcv1.DynamicRouteServerPeerGroupPeer{
			ID:       core.StringPtr("id"),
			Endpoint: &vpcv1.DynamicRouteServerPeerGroupPeerAddressEndpoint{ID: core.StringPtr("tgw-id"), ResourceType: core.StringPtr("transit_gateway")},
		})
		assert.Nil(t, err)
		assert.Equal(t, []map[string]interface{}{{"id": "tgw-id", "resource_type": "transit_gateway"}}, result["endpoint"])
	})
	t.Run("nil interface and nil pointers are errors", func(t *testing.T) {
		_, err := vpc.ResourceIBMIsDynamicRouteServerPeerGroupPeerDynamicRouteServerPeerGroupPeerToMap(nil)
		assert.NotNil(t, err)
		var address *vpcv1.DynamicRouteServerPeerGroupPeerAddress
		_, err = vpc.ResourceIBMIsDynamicRouteServerPeerGroupPeerDynamicRouteServerPeerGroupPeerToMap(address)
		assert.NotNil(t, err)
		var tgw *vpcv1.DynamicRouteServerPeerGroupPeerTransitGateway
		_, err = vpc.ResourceIBMIsDynamicRouteServerPeerGroupPeerDynamicRouteServerPeerGroupPeerToMap(tgw)
		assert.NotNil(t, err)
		var base *vpcv1.DynamicRouteServerPeerGroupPeer
		_, err = vpc.ResourceIBMIsDynamicRouteServerPeerGroupPeerDynamicRouteServerPeerGroupPeerToMap(base)
		assert.NotNil(t, err)
	})
}

func TestResourceIBMIsDynamicRouteServerPeerGroupPeerDynamicRouteServerPeerGroupPeerAddressEndpointToMap(t *testing.T) {
	t.Run("IP gateway subtype without gateway", func(t *testing.T) {
		result, err := vpc.ResourceIBMIsDynamicRouteServerPeerGroupPeerDynamicRouteServerPeerGroupPeerAddressEndpointToMap(&vpcv1.DynamicRouteServerPeerGroupPeerAddressEndpointIPGateway{Address: core.StringPtr("10.240.0.10")})
		assert.Nil(t, err)
		assert.Equal(t, map[string]interface{}{"address": "10.240.0.10"}, result)
	})
	t.Run("reserved IP subtype with deleted", func(t *testing.T) {
		result, err := vpc.ResourceIBMIsDynamicRouteServerPeerGroupPeerDynamicRouteServerPeerGroupPeerAddressEndpointToMap(&vpcv1.DynamicRouteServerPeerGroupPeerAddressEndpointByReservedIP{
			ID: core.StringPtr("rip-id"), Href: core.StringPtr("h"), Name: core.StringPtr("n"), ResourceType: core.StringPtr("subnet_reserved_ip"), Address: core.StringPtr("10.240.0.10"),
			Deleted: &vpcv1.Deleted{MoreInfo: core.StringPtr("https://cloud.ibm.com/apidocs/vpc#deleted-resources")},
		})
		assert.Nil(t, err)
		assert.Equal(t, "rip-id", result["id"])
		assert.Equal(t, []map[string]interface{}{{"more_info": "https://cloud.ibm.com/apidocs/vpc#deleted-resources"}}, result["deleted"])
		_, hasGateway := result["gateway"]
		assert.False(t, hasGateway)
	})
	t.Run("base struct", func(t *testing.T) {
		result, err := vpc.ResourceIBMIsDynamicRouteServerPeerGroupPeerDynamicRouteServerPeerGroupPeerAddressEndpointToMap(&vpcv1.DynamicRouteServerPeerGroupPeerAddressEndpoint{
			Address: core.StringPtr("10.240.0.10"), Gateway: &vpcv1.DynamicRouteServerPeerGroupPeerAddressEndpointIPGatewayGateway{Address: core.StringPtr("10.240.0.1")}, Href: core.StringPtr("h"), ID: core.StringPtr("i"), Name: core.StringPtr("n"), ResourceType: core.StringPtr("r"),
			Deleted: &vpcv1.Deleted{MoreInfo: core.StringPtr("m")},
		})
		assert.Nil(t, err)
		assert.Equal(t, map[string]interface{}{
			"address": "10.240.0.10", "gateway": []map[string]interface{}{{"address": "10.240.0.1"}}, "href": "h", "id": "i", "name": "n", "resource_type": "r", "deleted": []map[string]interface{}{{"more_info": "m"}},
		}, result)
	})
	t.Run("errors", func(t *testing.T) {
		_, err := vpc.ResourceIBMIsDynamicRouteServerPeerGroupPeerDynamicRouteServerPeerGroupPeerAddressEndpointToMap(nil)
		assert.NotNil(t, err)
		var base *vpcv1.DynamicRouteServerPeerGroupPeerAddressEndpoint
		_, err = vpc.ResourceIBMIsDynamicRouteServerPeerGroupPeerDynamicRouteServerPeerGroupPeerAddressEndpointToMap(base)
		assert.NotNil(t, err)
		var ipGateway *vpcv1.DynamicRouteServerPeerGroupPeerAddressEndpointIPGateway
		_, err = vpc.ResourceIBMIsDynamicRouteServerPeerGroupPeerDynamicRouteServerPeerGroupPeerAddressEndpointToMap(ipGateway)
		assert.NotNil(t, err)
		var reservedIP *vpcv1.DynamicRouteServerPeerGroupPeerAddressEndpointByReservedIP
		_, err = vpc.ResourceIBMIsDynamicRouteServerPeerGroupPeerDynamicRouteServerPeerGroupPeerAddressEndpointToMap(reservedIP)
		assert.NotNil(t, err)
	})
}

func TestResourceIBMIsDynamicRouteServerPeerGroupPeerDynamicRouteServerPeerGroupPeerTransitGatewayEndpointToMap(t *testing.T) {
	result, err := vpc.ResourceIBMIsDynamicRouteServerPeerGroupPeerDynamicRouteServerPeerGroupPeerTransitGatewayEndpointToMap(&vpcv1.DynamicRouteServerPeerGroupPeerTransitGatewayEndpointTransitGatewayReference{CRN: core.StringPtr("crn"), ID: core.StringPtr("id"), ResourceType: core.StringPtr("transit_gateway")})
	assert.Nil(t, err)
	assert.Equal(t, map[string]interface{}{"crn": "crn", "id": "id", "resource_type": "transit_gateway"}, result)
	result, err = vpc.ResourceIBMIsDynamicRouteServerPeerGroupPeerDynamicRouteServerPeerGroupPeerTransitGatewayEndpointToMap(&vpcv1.DynamicRouteServerPeerGroupPeerTransitGatewayEndpoint{})
	assert.Nil(t, err)
	assert.Equal(t, map[string]interface{}{}, result)
	_, err = vpc.ResourceIBMIsDynamicRouteServerPeerGroupPeerDynamicRouteServerPeerGroupPeerTransitGatewayEndpointToMap(nil)
	assert.NotNil(t, err)
	var ref *vpcv1.DynamicRouteServerPeerGroupPeerTransitGatewayEndpointTransitGatewayReference
	_, err = vpc.ResourceIBMIsDynamicRouteServerPeerGroupPeerDynamicRouteServerPeerGroupPeerTransitGatewayEndpointToMap(ref)
	assert.NotNil(t, err)
}

func TestResourceIBMIsDynamicRouteServerPeerGroupPeerSmallToMaps(t *testing.T) {
	assert.Equal(t, map[string]interface{}{}, vpc.ResourceIBMIsDynamicRouteServerPeerGroupPeerDynamicRouteServerPeerGroupPeerAddressEndpointIPGatewayGatewayToMap(nil))
	assert.Equal(t, map[string]interface{}{"address": "a"}, vpc.ResourceIBMIsDynamicRouteServerPeerGroupPeerDynamicRouteServerPeerGroupPeerAddressEndpointIPGatewayGatewayToMap(&vpcv1.DynamicRouteServerPeerGroupPeerAddressEndpointIPGatewayGateway{Address: core.StringPtr("a")}))
	assert.Equal(t, map[string]interface{}{}, vpc.ResourceIBMIsDynamicRouteServerPeerGroupPeerIPToMap(nil))
	assert.Equal(t, map[string]interface{}{"address": "a"}, vpc.ResourceIBMIsDynamicRouteServerPeerGroupPeerIPToMap(&vpcv1.IP{Address: core.StringPtr("a")}))
	assert.Equal(t, map[string]interface{}{}, vpc.ResourceIBMIsDynamicRouteServerPeerGroupPeerDeletedToMap(nil))
	assert.Equal(t, map[string]interface{}{"more_info": "m"}, vpc.ResourceIBMIsDynamicRouteServerPeerGroupPeerDeletedToMap(&vpcv1.Deleted{MoreInfo: core.StringPtr("m")}))
	assert.Equal(t, map[string]interface{}{}, vpc.ResourceIBMIsDynamicRouteServerPeerGroupPeerLifecycleReasonToMap(nil))
	assert.Equal(t, map[string]interface{}{"code": "c", "message": "m", "more_info": "i"}, vpc.ResourceIBMIsDynamicRouteServerPeerGroupPeerLifecycleReasonToMap(&vpcv1.LifecycleReason{Code: core.StringPtr("c"), Message: core.StringPtr("m"), MoreInfo: core.StringPtr("i")}))
	assert.Equal(t, map[string]interface{}{}, vpc.ResourceIBMIsDynamicRouteServerPeerGroupPeerDynamicRouteServerPeerGroupPeerStatusReasonToMap(nil))
	assert.Equal(t, map[string]interface{}{"code": "c", "message": "m"}, vpc.ResourceIBMIsDynamicRouteServerPeerGroupPeerDynamicRouteServerPeerGroupPeerStatusReasonToMap(&vpcv1.DynamicRouteServerPeerGroupPeerStatusReason{Code: core.StringPtr("c"), Message: core.StringPtr("m")}))
	assert.Equal(t, map[string]interface{}{}, vpc.ResourceIBMIsDynamicRouteServerPeerGroupPeerDynamicRouteServerPeerGroupPeerNameToMap(nil))
	assert.Equal(t, map[string]interface{}{"name": "n"}, vpc.ResourceIBMIsDynamicRouteServerPeerGroupPeerDynamicRouteServerPeerGroupPeerNameToMap(&vpcv1.DynamicRouteServerPeerGroupPeerName{Name: core.StringPtr("n")}))
	assert.Equal(t, map[string]interface{}{}, vpc.ResourceIBMIsDynamicRouteServerPeerGroupPeerDynamicRouteServerPeerGroupPeerTransitGatewaySessionRemoteToMap(nil))
	assert.Equal(t, map[string]interface{}{"name": "n"}, vpc.ResourceIBMIsDynamicRouteServerPeerGroupPeerDynamicRouteServerPeerGroupPeerTransitGatewaySessionRemoteToMap(&vpcv1.DynamicRouteServerPeerGroupPeerTransitGatewaySessionRemote{Name: core.StringPtr("n")}))
	assert.Equal(t, map[string]interface{}{}, vpc.ResourceIBMIsDynamicRouteServerPeerGroupPeerDynamicRouteServerPeerGroupPeerReferenceToMap(nil))
	assert.Equal(t, map[string]interface{}{"id": "i", "deleted": []map[string]interface{}{{"more_info": "m"}}}, vpc.ResourceIBMIsDynamicRouteServerPeerGroupPeerDynamicRouteServerPeerGroupPeerReferenceToMap(&vpcv1.DynamicRouteServerPeerGroupPeerReference{ID: core.StringPtr("i"), Deleted: &vpcv1.Deleted{MoreInfo: core.StringPtr("m")}}))
	assert.Equal(t, map[string]interface{}{}, vpc.ResourceIBMIsDynamicRouteServerPeerGroupPeerVirtualNetworkInterfaceReferenceToMap(nil))
	assert.Equal(t, map[string]interface{}{"crn": "c", "id": "i"}, vpc.ResourceIBMIsDynamicRouteServerPeerGroupPeerVirtualNetworkInterfaceReferenceToMap(&vpcv1.VirtualNetworkInterfaceReference{CRN: core.StringPtr("c"), ID: core.StringPtr("i")}))
}

func TestResourceIBMIsDynamicRouteServerPeerGroupPeerDynamicRouteServerMemberReferenceToMap(t *testing.T) {
	assert.Equal(t, map[string]interface{}{}, vpc.ResourceIBMIsDynamicRouteServerPeerGroupPeerDynamicRouteServerMemberReferenceToMap(nil))
	empty := vpc.ResourceIBMIsDynamicRouteServerPeerGroupPeerDynamicRouteServerMemberReferenceToMap(&vpcv1.DynamicRouteServerMemberReference{})
	assert.Equal(t, []map[string]interface{}{}, empty["virtual_network_interfaces"])
	full := vpc.ResourceIBMIsDynamicRouteServerPeerGroupPeerDynamicRouteServerMemberReferenceToMap(testPeerGroupPeerMemberReference())
	assert.Equal(t, "member-id", full["id"])
	vnis := full["virtual_network_interfaces"].([]map[string]interface{})
	assert.Len(t, vnis, 1)
	assert.Equal(t, "vni-id", vnis[0]["id"])
}

func TestResourceIBMIsDynamicRouteServerPeerGroupPeerDynamicRouteServerPeerGroupPeerAddressBidirectionalForwardingDetectionToMap(t *testing.T) {
	_, err := vpc.ResourceIBMIsDynamicRouteServerPeerGroupPeerDynamicRouteServerPeerGroupPeerAddressBidirectionalForwardingDetectionToMap(nil)
	assert.NotNil(t, err)
	result, err := vpc.ResourceIBMIsDynamicRouteServerPeerGroupPeerDynamicRouteServerPeerGroupPeerAddressBidirectionalForwardingDetectionToMap(&vpcv1.DynamicRouteServerPeerGroupPeerAddressBidirectionalForwardingDetection{Enabled: core.BoolPtr(true), DetectMultiplier: core.Int64Ptr(0)})
	assert.Nil(t, err)
	assert.Equal(t, map[string]interface{}{"enabled": true, "detect_multiplier": 0, "sessions": []map[string]interface{}{}}, result)
	result, err = vpc.ResourceIBMIsDynamicRouteServerPeerGroupPeerDynamicRouteServerPeerGroupPeerAddressBidirectionalForwardingDetectionToMap(&vpcv1.DynamicRouteServerPeerGroupPeerAddressBidirectionalForwardingDetection{
		Sessions: []vpcv1.DynamicRouteServerPeerGroupPeerAddressBidirectionalForwardingDetectionSession{{State: core.StringPtr("up")}, {State: core.StringPtr("down")}},
	})
	assert.Nil(t, err)
	assert.Len(t, result["sessions"], 2)
}

func TestResourceIBMIsDynamicRouteServerPeerGroupPeerDynamicRouteServerPeerGroupPeerAddressBidirectionalForwardingDetectionSessionToMap(t *testing.T) {
	_, err := vpc.ResourceIBMIsDynamicRouteServerPeerGroupPeerDynamicRouteServerPeerGroupPeerAddressBidirectionalForwardingDetectionSessionToMap(nil)
	assert.NotNil(t, err)
	result, err := vpc.ResourceIBMIsDynamicRouteServerPeerGroupPeerDynamicRouteServerPeerGroupPeerAddressBidirectionalForwardingDetectionSessionToMap(&vpcv1.DynamicRouteServerPeerGroupPeerAddressBidirectionalForwardingDetectionSession{})
	assert.Nil(t, err)
	assert.Equal(t, map[string]interface{}{}, result)
	result, err = vpc.ResourceIBMIsDynamicRouteServerPeerGroupPeerDynamicRouteServerPeerGroupPeerAddressBidirectionalForwardingDetectionSessionToMap(&vpcv1.DynamicRouteServerPeerGroupPeerAddressBidirectionalForwardingDetectionSession{Local: testPeerGroupPeerMemberReference(), Remote: &vpcv1.DynamicRouteServerPeerGroupPeerReference{Name: core.StringPtr("r")}, State: core.StringPtr("up")})
	assert.Nil(t, err)
	assert.Equal(t, "up", result["state"])
	assert.Equal(t, "member-id", result["local"].([]map[string]interface{})[0]["id"])
	assert.Equal(t, "r", result["remote"].([]map[string]interface{})[0]["name"])
}

func TestResourceIBMIsDynamicRouteServerPeerGroupPeerDynamicRouteServerPeerGroupPeerAddressSessionToMap(t *testing.T) {
	_, err := vpc.ResourceIBMIsDynamicRouteServerPeerGroupPeerDynamicRouteServerPeerGroupPeerAddressSessionToMap(nil)
	assert.NotNil(t, err)
	result, err := vpc.ResourceIBMIsDynamicRouteServerPeerGroupPeerDynamicRouteServerPeerGroupPeerAddressSessionToMap(&vpcv1.DynamicRouteServerPeerGroupPeerAddressSession{})
	assert.Nil(t, err)
	assert.Equal(t, map[string]interface{}{}, result)
	result, err = vpc.ResourceIBMIsDynamicRouteServerPeerGroupPeerDynamicRouteServerPeerGroupPeerAddressSessionToMap(&vpcv1.DynamicRouteServerPeerGroupPeerAddressSession{EstablishedAt: CreateMockDateTime("2026-01-02T03:04:05.000Z"), Local: testPeerGroupPeerMemberReference(), ProtocolState: core.StringPtr("established"), Remote: &vpcv1.DynamicRouteServerPeerGroupPeerName{Name: core.StringPtr("remote")}})
	assert.Nil(t, err)
	assert.Equal(t, "2026-01-02T03:04:05.000Z", result["established_at"])
	assert.Equal(t, "established", result["protocol_state"])
	assert.Equal(t, []map[string]interface{}{{"name": "remote"}}, result["remote"])
}

func TestResourceIBMIsDynamicRouteServerPeerGroupPeerDynamicRouteServerPeerGroupPeerTransitGatewaySessionToMap(t *testing.T) {
	_, err := vpc.ResourceIBMIsDynamicRouteServerPeerGroupPeerDynamicRouteServerPeerGroupPeerTransitGatewaySessionToMap(nil)
	assert.NotNil(t, err)
	result, err := vpc.ResourceIBMIsDynamicRouteServerPeerGroupPeerDynamicRouteServerPeerGroupPeerTransitGatewaySessionToMap(&vpcv1.DynamicRouteServerPeerGroupPeerTransitGatewaySession{EstablishedAt: CreateMockDateTime("2026-01-02T03:04:05.000Z"), Local: testPeerGroupPeerMemberReference(), ProtocolState: core.StringPtr("established"), Remote: &vpcv1.DynamicRouteServerPeerGroupPeerTransitGatewaySessionRemote{Name: core.StringPtr("tgw")}})
	assert.Nil(t, err)
	assert.Equal(t, "2026-01-02T03:04:05.000Z", result["established_at"])
	assert.Equal(t, []map[string]interface{}{{"name": "tgw"}}, result["remote"])
	assert.Equal(t, "member-id", result["local"].([]map[string]interface{})[0]["id"])
}

func TestResourceIBMIsDynamicRouteServerPeerGroupPeerDynamicRouteServerPeerGroupPeerPatchAsPatch(t *testing.T) {
	assert.Equal(t, map[string]interface{}{}, vpc.ResourceIBMIsDynamicRouteServerPeerGroupPeerDynamicRouteServerPeerGroupPeerPatchAsPatch(nil, nil))
	assert.Equal(t, map[string]interface{}{}, vpc.ResourceIBMIsDynamicRouteServerPeerGroupPeerDynamicRouteServerPeerGroupPeerPatchAsPatch(&vpcv1.DynamicRouteServerPeerGroupPeerPatch{}, nil))
	patch := vpc.ResourceIBMIsDynamicRouteServerPeerGroupPeerDynamicRouteServerPeerGroupPeerPatchAsPatch(&vpcv1.DynamicRouteServerPeerGroupPeerPatch{Priority: core.Int64Ptr(0)}, nil)
	assert.Equal(t, map[string]interface{}{"priority": int64(0)}, patch, "priority 0 is a real value")
	patch = vpc.ResourceIBMIsDynamicRouteServerPeerGroupPeerDynamicRouteServerPeerGroupPeerPatchAsPatch(&vpcv1.DynamicRouteServerPeerGroupPeerPatch{Asn: core.Int64Ptr(64515), Name: core.StringPtr("n"), Priority: core.Int64Ptr(3), State: core.StringPtr("disabled")}, nil)
	assert.Equal(t, map[string]interface{}{"asn": int64(64515), "name": "n", "priority": int64(3), "state": "disabled"}, patch)
}
