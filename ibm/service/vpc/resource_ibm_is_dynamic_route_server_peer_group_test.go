// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/vpc"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/stretchr/testify/assert"
)

const isDynamicRouteServerPeerGroupResource = "ibm_is_dynamic_route_server_peer_group.testacc_peer_group"

func TestAccIBMIsDynamicRouteServerPeerGroupBasic(t *testing.T) {
	var conf vpcv1.DynamicRouteServerPeerGroup
	vpcname := fmt.Sprintf("tf-vpc-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tf-subnet-%d", acctest.RandIntRange(10, 100))
	vniname := fmt.Sprintf("tf-vni-%d", acctest.RandIntRange(10, 100))
	drsname := fmt.Sprintf("tf-drs-%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf-peer-group-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIsDynamicRouteServerPeerGroupDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsDynamicRouteServerPeerGroupConfigBasic(vpcname, subnetname, vniname, drsname, name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIsDynamicRouteServerPeerGroupExists(isDynamicRouteServerPeerGroupResource, &conf),
					resource.TestCheckResourceAttr(isDynamicRouteServerPeerGroupResource, "name", name),
					resource.TestCheckResourceAttr(isDynamicRouteServerPeerGroupResource, "roles.0", "next_hop"),
					resource.TestCheckResourceAttr(isDynamicRouteServerPeerGroupResource, "health_monitor.0.roles.0", "next_hop"),
					resource.TestCheckResourceAttr(isDynamicRouteServerPeerGroupResource, "peers.#", "1"),
					resource.TestCheckResourceAttr(isDynamicRouteServerPeerGroupResource, "peers.0.asn", "64513"),
					resource.TestCheckResourceAttr(isDynamicRouteServerPeerGroupResource, "peers.0.endpoint.0.address", "10.240.0.10"),
					resource.TestCheckResourceAttrSet(isDynamicRouteServerPeerGroupResource, "peers.0.id"),
					resource.TestCheckResourceAttrSet(isDynamicRouteServerPeerGroupResource, "dynamic_route_server_peer_group_id"),
					resource.TestCheckResourceAttrSet(isDynamicRouteServerPeerGroupResource, "href"),
					resource.TestCheckResourceAttrSet(isDynamicRouteServerPeerGroupResource, "created_at"),
					resource.TestCheckResourceAttr(isDynamicRouteServerPeerGroupResource, "lifecycle_state", "stable"),
					resource.TestCheckResourceAttr(isDynamicRouteServerPeerGroupResource, "resource_type", "dynamic_route_server_peer_group"),
				),
			},
		},
	})
}

func TestAccIBMIsDynamicRouteServerPeerGroupAllArgs(t *testing.T) {
	var conf vpcv1.DynamicRouteServerPeerGroup
	vpcname := fmt.Sprintf("tf-vpc-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tf-subnet-%d", acctest.RandIntRange(10, 100))
	vniname := fmt.Sprintf("tf-vni-%d", acctest.RandIntRange(10, 100))
	drsname := fmt.Sprintf("tf-drs-%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf-peer-group-%d", acctest.RandIntRange(10, 100))
	nameUpdate := fmt.Sprintf("tf-peer-group-upd-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIsDynamicRouteServerPeerGroupDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsDynamicRouteServerPeerGroupConfig(vpcname, subnetname, vniname, drsname, name, "bgp", "enabled", "192.168.0.0/16", 16, 24, "next_hop"),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIsDynamicRouteServerPeerGroupExists(isDynamicRouteServerPeerGroupResource, &conf),
					resource.TestCheckResourceAttr(isDynamicRouteServerPeerGroupResource, "name", name),
					resource.TestCheckResourceAttr(isDynamicRouteServerPeerGroupResource, "health_monitor.0.mode", "bgp"),
					resource.TestCheckResourceAttr(isDynamicRouteServerPeerGroupResource, "state", "enabled"),
					resource.TestCheckResourceAttr(isDynamicRouteServerPeerGroupResource, "roles.0", "next_hop"),
					resource.TestCheckResourceAttr(isDynamicRouteServerPeerGroupResource, "prefix_watcher.0.monitored_prefixes.0.prefix", "192.168.0.0/16"),
					resource.TestCheckResourceAttr(isDynamicRouteServerPeerGroupResource, "prefix_watcher.0.monitored_prefixes.0.ge", "16"),
					resource.TestCheckResourceAttr(isDynamicRouteServerPeerGroupResource, "prefix_watcher.0.monitored_prefixes.0.le", "24"),
					resource.TestCheckResourceAttr(isDynamicRouteServerPeerGroupResource, "prefix_watcher.0.excluded_prefixes.0.prefix", "192.168.100.0/24"),
					resource.TestCheckResourceAttr(isDynamicRouteServerPeerGroupResource, "peers.0.name", name+"-peer"),
					resource.TestCheckResourceAttr(isDynamicRouteServerPeerGroupResource, "peers.0.priority", "1"),
					resource.TestCheckResourceAttr(isDynamicRouteServerPeerGroupResource, "peers.0.state", "enabled"),
					resource.TestCheckResourceAttr(isDynamicRouteServerPeerGroupResource, "peers.0.endpoint.0.gateway.0.address", "10.240.0.1"),
				),
			},
			resource.TestStep{
				// Update every patchable field: name, health_monitor.mode, prefix_watcher and state.
				Config: testAccCheckIBMIsDynamicRouteServerPeerGroupConfig(vpcname, subnetname, vniname, drsname, nameUpdate, "bfd", "disabled", "172.16.0.0/12", 12, 20, "next_hop"),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIsDynamicRouteServerPeerGroupExists(isDynamicRouteServerPeerGroupResource, &conf),
					resource.TestCheckResourceAttr(isDynamicRouteServerPeerGroupResource, "name", nameUpdate),
					resource.TestCheckResourceAttr(isDynamicRouteServerPeerGroupResource, "health_monitor.0.mode", "bfd"),
					resource.TestCheckResourceAttr(isDynamicRouteServerPeerGroupResource, "state", "disabled"),
					resource.TestCheckResourceAttr(isDynamicRouteServerPeerGroupResource, "prefix_watcher.0.monitored_prefixes.0.prefix", "172.16.0.0/12"),
					resource.TestCheckResourceAttr(isDynamicRouteServerPeerGroupResource, "prefix_watcher.0.monitored_prefixes.0.ge", "12"),
					resource.TestCheckResourceAttr(isDynamicRouteServerPeerGroupResource, "prefix_watcher.0.monitored_prefixes.0.le", "20"),
				),
			},
			resource.TestStep{
				// Changing roles forces a new peer group.
				Config: testAccCheckIBMIsDynamicRouteServerPeerGroupConfig(vpcname, subnetname, vniname, drsname, nameUpdate, "bfd", "disabled", "172.16.0.0/12", 12, 20, "router"),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIsDynamicRouteServerPeerGroupExists(isDynamicRouteServerPeerGroupResource, &conf),
					resource.TestCheckResourceAttr(isDynamicRouteServerPeerGroupResource, "roles.0", "router"),
					resource.TestCheckResourceAttr(isDynamicRouteServerPeerGroupResource, "health_monitor.0.roles.0", "router"),
				),
			},
			resource.TestStep{
				ResourceName:      isDynamicRouteServerPeerGroupResource,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

// testAccCheckIBMIsDynamicRouteServerPeerGroupDependencies creates the VPC, subnet, virtual network interface and
// dynamic route server that a peer group needs.
func testAccCheckIBMIsDynamicRouteServerPeerGroupDependencies(vpcname, subnetname, vniname, drsname string) string {
	return fmt.Sprintf(`
	resource "ibm_is_vpc" "testacc_vpc" {
		name = "%s"
	}

	resource "ibm_is_subnet" "testacc_subnet" {
		name            = "%s"
		vpc             = ibm_is_vpc.testacc_vpc.id
		zone            = "%s"
		ipv4_cidr_block = "%s"
	}

	resource "ibm_is_virtual_network_interface" "testacc_vni" {
		name   = "%s"
		subnet = ibm_is_subnet.testacc_subnet.id
	}

	resource "ibm_is_dynamic_route_server" "testacc_drs" {
		name      = "%s"
		local_asn = 64512
		members {
			name = "%s-member"
			virtual_network_interfaces {
				id = ibm_is_virtual_network_interface.testacc_vni.id
			}
		}
	}
	`, vpcname, subnetname, acc.ISZoneName, acc.ISCIDR, vniname, drsname, drsname)
}

func testAccCheckIBMIsDynamicRouteServerPeerGroupConfigBasic(vpcname, subnetname, vniname, drsname, name string) string {
	return testAccCheckIBMIsDynamicRouteServerPeerGroupDependencies(vpcname, subnetname, vniname, drsname) + fmt.Sprintf(`
	resource "ibm_is_dynamic_route_server_peer_group" "testacc_peer_group" {
		dynamic_route_server_id = ibm_is_dynamic_route_server.testacc_drs.id
		name                    = "%s"
		roles                   = ["next_hop"]
		health_monitor {
			roles = ["next_hop"]
		}
		peers {
			asn = 64513
			endpoint {
				address = "10.240.0.10"
			}
		}
	}
	`, name)
}

func testAccCheckIBMIsDynamicRouteServerPeerGroupConfig(vpcname, subnetname, vniname, drsname, name, mode, state, prefix string, ge, le int, role string) string {
	return testAccCheckIBMIsDynamicRouteServerPeerGroupDependencies(vpcname, subnetname, vniname, drsname) + fmt.Sprintf(`
	resource "ibm_is_dynamic_route_server_peer_group" "testacc_peer_group" {
		dynamic_route_server_id = ibm_is_dynamic_route_server.testacc_drs.id
		name                    = "%s"
		roles                   = ["%s"]
		state                   = "%s"
		health_monitor {
			mode  = "%s"
			roles = ["%s"]
		}
		peers {
			asn      = 64513
			name     = "%s-peer"
			priority = 1
			state    = "enabled"
			endpoint {
				address = "10.240.0.10"
				gateway {
					address = "10.240.0.1"
				}
			}
		}
		prefix_watcher {
			monitored_prefixes {
				prefix = "%s"
				ge     = %d
				le     = %d
			}
			excluded_prefixes {
				prefix = "192.168.100.0/24"
			}
		}
	}
	`, name, role, state, mode, role, name, prefix, ge, le)
}

func testAccCheckIBMIsDynamicRouteServerPeerGroupExists(n string, obj *vpcv1.DynamicRouteServerPeerGroup) resource.TestCheckFunc {
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

		getOptions := &vpcv1.GetDynamicRouteServerPeerGroupOptions{}
		getOptions.SetDynamicRouteServerID(parts[0])
		getOptions.SetID(parts[1])

		peerGroup, _, err := vpcClient.GetDynamicRouteServerPeerGroup(getOptions)
		if err != nil {
			return err
		}

		*obj = *peerGroup
		return nil
	}
}

func testAccCheckIBMIsDynamicRouteServerPeerGroupDestroy(s *terraform.State) error {
	vpcClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).VpcV1API()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_is_dynamic_route_server_peer_group" {
			continue
		}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}
		if len(parts) != 2 {
			return fmt.Errorf("unexpected ID format %q", rs.Primary.ID)
		}

		getOptions := &vpcv1.GetDynamicRouteServerPeerGroupOptions{}
		getOptions.SetDynamicRouteServerID(parts[0])
		getOptions.SetID(parts[1])

		_, response, err := vpcClient.GetDynamicRouteServerPeerGroup(getOptions)
		if err == nil {
			return fmt.Errorf("dynamic route server peer group still exists: %s", rs.Primary.ID)
		} else if response == nil || response.StatusCode != 404 {
			return fmt.Errorf("Error checking for dynamic route server peer group (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}

func TestResourceIBMIsDynamicRouteServerPeerGroupDynamicRouteServerPeerGroupHealthMonitorToMap(t *testing.T) {
	t.Run("next hop subtype", func(t *testing.T) {
		model := &vpcv1.DynamicRouteServerPeerGroupHealthMonitorNextHop{Mode: core.StringPtr("bfd"), Roles: []string{"next_hop"}}
		result, err := vpc.ResourceIBMIsDynamicRouteServerPeerGroupDynamicRouteServerPeerGroupHealthMonitorToMap(model)
		assert.Nil(t, err)
		assert.Equal(t, map[string]interface{}{"mode": "bfd", "roles": []string{"next_hop"}}, result)
	})
	t.Run("router subtype", func(t *testing.T) {
		model := &vpcv1.DynamicRouteServerPeerGroupHealthMonitorRouter{Mode: core.StringPtr("bgp"), Roles: []string{"router"}}
		result, err := vpc.ResourceIBMIsDynamicRouteServerPeerGroupDynamicRouteServerPeerGroupHealthMonitorToMap(model)
		assert.Nil(t, err)
		assert.Equal(t, map[string]interface{}{"mode": "bgp", "roles": []string{"router"}}, result)
	})
	t.Run("base struct with multiple roles", func(t *testing.T) {
		model := &vpcv1.DynamicRouteServerPeerGroupHealthMonitor{Mode: core.StringPtr("none"), Roles: []string{"next_hop", "router"}}
		result, err := vpc.ResourceIBMIsDynamicRouteServerPeerGroupDynamicRouteServerPeerGroupHealthMonitorToMap(model)
		assert.Nil(t, err)
		assert.Equal(t, map[string]interface{}{"mode": "none", "roles": []string{"next_hop", "router"}}, result)
	})
	t.Run("base struct with nothing set returns empty roles", func(t *testing.T) {
		result, err := vpc.ResourceIBMIsDynamicRouteServerPeerGroupDynamicRouteServerPeerGroupHealthMonitorToMap(&vpcv1.DynamicRouteServerPeerGroupHealthMonitor{})
		assert.Nil(t, err)
		assert.Equal(t, map[string]interface{}{"roles": []string{}}, result)
	})
	t.Run("nil interface is an error", func(t *testing.T) {
		result, err := vpc.ResourceIBMIsDynamicRouteServerPeerGroupDynamicRouteServerPeerGroupHealthMonitorToMap(nil)
		assert.NotNil(t, err)
		assert.Nil(t, result)
	})
	t.Run("nil typed pointer is an error", func(t *testing.T) {
		var model *vpcv1.DynamicRouteServerPeerGroupHealthMonitorNextHop
		result, err := vpc.ResourceIBMIsDynamicRouteServerPeerGroupDynamicRouteServerPeerGroupHealthMonitorToMap(model)
		assert.NotNil(t, err)
		assert.Nil(t, result)
	})
}

func TestResourceIBMIsDynamicRouteServerPeerGroupDynamicRouteServerPeerGroupPeerToMap(t *testing.T) {
	t.Run("address subtype with IP gateway endpoint", func(t *testing.T) {
		model := &vpcv1.DynamicRouteServerPeerGroupPeerAddress{
			ID:       core.StringPtr("peer-id"),
			Name:     core.StringPtr("my-peer"),
			Asn:      core.Int64Ptr(64513),
			Priority: core.Int64Ptr(0),
			State:    core.StringPtr("enabled"),
			Href:     core.StringPtr("https://example.com/peer"),
			Endpoint: &vpcv1.DynamicRouteServerPeerGroupPeerAddressEndpointIPGateway{
				Address: core.StringPtr("10.240.0.10"),
				Gateway: &vpcv1.DynamicRouteServerPeerGroupPeerAddressEndpointIPGatewayGateway{Address: core.StringPtr("10.240.0.1")},
			},
		}
		result, err := vpc.ResourceIBMIsDynamicRouteServerPeerGroupDynamicRouteServerPeerGroupPeerToMap(model)
		assert.Nil(t, err)
		assert.Equal(t, "peer-id", result["id"])
		assert.Equal(t, "my-peer", result["name"])
		assert.Equal(t, 64513, result["asn"])
		assert.Equal(t, 0, result["priority"])
		assert.Equal(t, "enabled", result["state"])
		_, hasHref := result["href"]
		assert.False(t, hasHref, "response-only fields are not part of the peers input block")
		endpoints := result["endpoint"].([]map[string]interface{})
		assert.Len(t, endpoints, 1)
		assert.Equal(t, "10.240.0.10", endpoints[0]["address"])
		assert.Equal(t, []map[string]interface{}{{"address": "10.240.0.1"}}, endpoints[0]["gateway"])
		_, hasID := endpoints[0]["id"]
		assert.False(t, hasID)
	})
	t.Run("address subtype with reserved IP endpoint", func(t *testing.T) {
		model := &vpcv1.DynamicRouteServerPeerGroupPeerAddress{
			ID:  core.StringPtr("peer-id"),
			Asn: core.Int64Ptr(64513),
			Endpoint: &vpcv1.DynamicRouteServerPeerGroupPeerAddressEndpointByReservedIP{
				ID:           core.StringPtr("rip-id"),
				Href:         core.StringPtr("https://example.com/rip"),
				Address:      core.StringPtr("10.240.0.11"),
				Gateway:      &vpcv1.IP{Address: core.StringPtr("10.240.0.1")},
				Name:         core.StringPtr("my-rip"),
				ResourceType: core.StringPtr("subnet_reserved_ip"),
			},
		}
		result, err := vpc.ResourceIBMIsDynamicRouteServerPeerGroupDynamicRouteServerPeerGroupPeerToMap(model)
		assert.Nil(t, err)
		endpoints := result["endpoint"].([]map[string]interface{})
		assert.Equal(t, "rip-id", endpoints[0]["id"])
		assert.Equal(t, "https://example.com/rip", endpoints[0]["href"])
		assert.Equal(t, "10.240.0.11", endpoints[0]["address"])
		assert.Equal(t, []map[string]interface{}{{"address": "10.240.0.1"}}, endpoints[0]["gateway"])
		_, hasName := endpoints[0]["name"]
		assert.False(t, hasName)
	})
	t.Run("transit gateway subtype", func(t *testing.T) {
		model := &vpcv1.DynamicRouteServerPeerGroupPeerTransitGateway{
			ID:    core.StringPtr("tgw-peer-id"),
			Name:  core.StringPtr("tgw-peer"),
			State: core.StringPtr("enabled"),
			Endpoint: &vpcv1.DynamicRouteServerPeerGroupPeerTransitGatewayEndpointTransitGatewayReference{
				CRN: core.StringPtr("crn:v1:bluemix:public:transit:us-south:a/123::gateway:abc"), ID: core.StringPtr("abc"), ResourceType: core.StringPtr("transit_gateway"),
			},
		}
		result, err := vpc.ResourceIBMIsDynamicRouteServerPeerGroupDynamicRouteServerPeerGroupPeerToMap(model)
		assert.Nil(t, err)
		assert.Equal(t, "tgw-peer-id", result["id"])
		_, hasAsn := result["asn"]
		assert.False(t, hasAsn)
		endpoints := result["endpoint"].([]map[string]interface{})
		assert.Equal(t, map[string]interface{}{"id": "abc"}, endpoints[0])
	})
	t.Run("base struct without endpoint", func(t *testing.T) {
		model := &vpcv1.DynamicRouteServerPeerGroupPeer{ID: core.StringPtr("base-id"), Asn: core.Int64Ptr(1)}
		result, err := vpc.ResourceIBMIsDynamicRouteServerPeerGroupDynamicRouteServerPeerGroupPeerToMap(model)
		assert.Nil(t, err)
		assert.Equal(t, map[string]interface{}{"id": "base-id", "asn": 1}, result)
	})
	t.Run("nil interface is an error", func(t *testing.T) {
		result, err := vpc.ResourceIBMIsDynamicRouteServerPeerGroupDynamicRouteServerPeerGroupPeerToMap(nil)
		assert.NotNil(t, err)
		assert.Nil(t, result)
	})
}

func TestResourceIBMIsDynamicRouteServerPeerGroupDynamicRouteServerPeerGroupPrefixWatcherToMap(t *testing.T) {
	t.Run("nil returns empty lists", func(t *testing.T) {
		result := vpc.ResourceIBMIsDynamicRouteServerPeerGroupDynamicRouteServerPeerGroupPrefixWatcherToMap(nil)
		assert.Equal(t, map[string]interface{}{"excluded_prefixes": []map[string]interface{}{}, "monitored_prefixes": []map[string]interface{}{}}, result)
	})
	t.Run("empty struct returns empty lists", func(t *testing.T) {
		result := vpc.ResourceIBMIsDynamicRouteServerPeerGroupDynamicRouteServerPeerGroupPrefixWatcherToMap(&vpcv1.DynamicRouteServerPeerGroupPrefixWatcher{})
		assert.Equal(t, []map[string]interface{}{}, result["excluded_prefixes"])
		assert.Equal(t, []map[string]interface{}{}, result["monitored_prefixes"])
	})
	t.Run("single and multiple prefixes", func(t *testing.T) {
		model := &vpcv1.DynamicRouteServerPeerGroupPrefixWatcher{
			ExcludedPrefixes: []vpcv1.DynamicRouteServerPeerGroupPrefixWatcherExcludedPrefix{{Prefix: core.StringPtr("10.0.0.0/8")}},
			MonitoredPrefixes: []vpcv1.DynamicRouteServerPeerGroupPrefixWatcherMonitoredPrefix{
				{Prefix: core.StringPtr("192.168.0.0/16"), Ge: core.Int64Ptr(16), Le: core.Int64Ptr(24)},
				{Prefix: core.StringPtr("172.16.0.0/12"), Ge: core.Int64Ptr(0)},
			},
		}
		result := vpc.ResourceIBMIsDynamicRouteServerPeerGroupDynamicRouteServerPeerGroupPrefixWatcherToMap(model)
		assert.Equal(t, []map[string]interface{}{{"prefix": "10.0.0.0/8"}}, result["excluded_prefixes"])
		assert.Equal(t, []map[string]interface{}{
			{"prefix": "192.168.0.0/16", "ge": 16, "le": 24},
			{"prefix": "172.16.0.0/12", "ge": 0},
		}, result["monitored_prefixes"])
	})
}

func TestResourceIBMIsDynamicRouteServerPeerGroupDynamicRouteServerPeerGroupPrefixWatcherExcludedPrefixToMap(t *testing.T) {
	assert.Equal(t, map[string]interface{}{}, vpc.ResourceIBMIsDynamicRouteServerPeerGroupDynamicRouteServerPeerGroupPrefixWatcherExcludedPrefixToMap(nil))
	model := &vpcv1.DynamicRouteServerPeerGroupPrefixWatcherExcludedPrefix{Prefix: core.StringPtr("10.0.0.0/8"), Ge: core.Int64Ptr(8), Le: core.Int64Ptr(16)}
	assert.Equal(t, map[string]interface{}{"prefix": "10.0.0.0/8", "ge": 8, "le": 16}, vpc.ResourceIBMIsDynamicRouteServerPeerGroupDynamicRouteServerPeerGroupPrefixWatcherExcludedPrefixToMap(model))
}

func TestResourceIBMIsDynamicRouteServerPeerGroupDynamicRouteServerPeerGroupPrefixWatcherMonitoredPrefixToMap(t *testing.T) {
	assert.Equal(t, map[string]interface{}{}, vpc.ResourceIBMIsDynamicRouteServerPeerGroupDynamicRouteServerPeerGroupPrefixWatcherMonitoredPrefixToMap(nil))
	model := &vpcv1.DynamicRouteServerPeerGroupPrefixWatcherMonitoredPrefix{Prefix: core.StringPtr("10.0.0.0/8")}
	assert.Equal(t, map[string]interface{}{"prefix": "10.0.0.0/8"}, vpc.ResourceIBMIsDynamicRouteServerPeerGroupDynamicRouteServerPeerGroupPrefixWatcherMonitoredPrefixToMap(model))
}

func TestResourceIBMIsDynamicRouteServerPeerGroupDynamicRouteServerPeerGroupStatusReasonToMap(t *testing.T) {
	assert.Equal(t, map[string]interface{}{}, vpc.ResourceIBMIsDynamicRouteServerPeerGroupDynamicRouteServerPeerGroupStatusReasonToMap(nil))
	model := &vpcv1.DynamicRouteServerPeerGroupStatusReason{Code: core.StringPtr("internal_error"), Message: core.StringPtr("Internal error")}
	assert.Equal(t, map[string]interface{}{"code": "internal_error", "message": "Internal error"}, vpc.ResourceIBMIsDynamicRouteServerPeerGroupDynamicRouteServerPeerGroupStatusReasonToMap(model))
	model.MoreInfo = core.StringPtr("https://cloud.ibm.com/apidocs/vpc")
	assert.Equal(t, "https://cloud.ibm.com/apidocs/vpc", vpc.ResourceIBMIsDynamicRouteServerPeerGroupDynamicRouteServerPeerGroupStatusReasonToMap(model)["more_info"])
}

func TestResourceIBMIsDynamicRouteServerPeerGroupMapToDynamicRouteServerPeerGroupHealthMonitorPrototype(t *testing.T) {
	t.Run("mode and single role", func(t *testing.T) {
		result, err := vpc.ResourceIBMIsDynamicRouteServerPeerGroupMapToDynamicRouteServerPeerGroupHealthMonitorPrototype(map[string]interface{}{"mode": "bfd", "roles": []interface{}{"next_hop"}})
		assert.Nil(t, err)
		model := result.(*vpcv1.DynamicRouteServerPeerGroupHealthMonitorPrototype)
		assert.Equal(t, "bfd", *model.Mode)
		assert.Equal(t, []string{"next_hop"}, model.Roles)
	})
	t.Run("empty mode is omitted and multiple roles kept", func(t *testing.T) {
		result, err := vpc.ResourceIBMIsDynamicRouteServerPeerGroupMapToDynamicRouteServerPeerGroupHealthMonitorPrototype(map[string]interface{}{"mode": "", "roles": []interface{}{"next_hop", "router"}})
		assert.Nil(t, err)
		model := result.(*vpcv1.DynamicRouteServerPeerGroupHealthMonitorPrototype)
		assert.Nil(t, model.Mode)
		assert.Equal(t, []string{"next_hop", "router"}, model.Roles)
	})
	t.Run("missing roles is an error", func(t *testing.T) {
		_, err := vpc.ResourceIBMIsDynamicRouteServerPeerGroupMapToDynamicRouteServerPeerGroupHealthMonitorPrototype(map[string]interface{}{"mode": "bfd"})
		assert.NotNil(t, err)
		_, err = vpc.ResourceIBMIsDynamicRouteServerPeerGroupMapToDynamicRouteServerPeerGroupHealthMonitorPrototype(map[string]interface{}{"roles": []interface{}{}})
		assert.NotNil(t, err)
	})
	t.Run("empty role string is an error", func(t *testing.T) {
		_, err := vpc.ResourceIBMIsDynamicRouteServerPeerGroupMapToDynamicRouteServerPeerGroupHealthMonitorPrototype(map[string]interface{}{"roles": []interface{}{""}})
		assert.NotNil(t, err)
	})
	t.Run("nil map is an error", func(t *testing.T) {
		_, err := vpc.ResourceIBMIsDynamicRouteServerPeerGroupMapToDynamicRouteServerPeerGroupHealthMonitorPrototype(nil)
		assert.NotNil(t, err)
	})
}

func TestResourceIBMIsDynamicRouteServerPeerGroupMapToDynamicRouteServerPeerGroupHealthMonitorPatch(t *testing.T) {
	result, err := vpc.ResourceIBMIsDynamicRouteServerPeerGroupMapToDynamicRouteServerPeerGroupHealthMonitorPatch(map[string]interface{}{"mode": "bgp", "roles": []interface{}{"next_hop"}})
	assert.Nil(t, err)
	assert.Equal(t, "bgp", *result.Mode)

	result, err = vpc.ResourceIBMIsDynamicRouteServerPeerGroupMapToDynamicRouteServerPeerGroupHealthMonitorPatch(map[string]interface{}{"mode": ""})
	assert.Nil(t, err)
	assert.Nil(t, result.Mode)

	result, err = vpc.ResourceIBMIsDynamicRouteServerPeerGroupMapToDynamicRouteServerPeerGroupHealthMonitorPatch(nil)
	assert.Nil(t, err)
	assert.Nil(t, result.Mode)
}

func TestResourceIBMIsDynamicRouteServerPeerGroupMapToDynamicRouteServerPeerGroupPrefixWatcherPrototype(t *testing.T) {
	t.Run("nil map gives empty lists", func(t *testing.T) {
		result, err := vpc.ResourceIBMIsDynamicRouteServerPeerGroupMapToDynamicRouteServerPeerGroupPrefixWatcherPrototype(nil)
		assert.Nil(t, err)
		assert.Equal(t, []vpcv1.DynamicRouteServerPeerGroupPrefixWatcherExcludedPrefixPrototype{}, result.ExcludedPrefixes)
		assert.Equal(t, []vpcv1.DynamicRouteServerPeerGroupPrefixWatcherMonitoredPrefixPrototype{}, result.MonitoredPrefixes)
	})
	t.Run("empty lists stay empty", func(t *testing.T) {
		result, err := vpc.ResourceIBMIsDynamicRouteServerPeerGroupMapToDynamicRouteServerPeerGroupPrefixWatcherPrototype(map[string]interface{}{"excluded_prefixes": []interface{}{}, "monitored_prefixes": []interface{}{}})
		assert.Nil(t, err)
		assert.Len(t, result.ExcludedPrefixes, 0)
		assert.Len(t, result.MonitoredPrefixes, 0)
	})
	t.Run("single and multiple prefixes", func(t *testing.T) {
		result, err := vpc.ResourceIBMIsDynamicRouteServerPeerGroupMapToDynamicRouteServerPeerGroupPrefixWatcherPrototype(map[string]interface{}{
			"excluded_prefixes": []interface{}{map[string]interface{}{"prefix": "10.0.0.0/8", "ge": 8, "le": 16}},
			"monitored_prefixes": []interface{}{
				map[string]interface{}{"prefix": "192.168.0.0/16", "ge": 16, "le": 24},
				map[string]interface{}{"prefix": "172.16.0.0/12", "ge": 0, "le": 0},
			},
		})
		assert.Nil(t, err)
		assert.Len(t, result.ExcludedPrefixes, 1)
		assert.Equal(t, int64(8), *result.ExcludedPrefixes[0].Ge)
		assert.Len(t, result.MonitoredPrefixes, 2)
		assert.Equal(t, "172.16.0.0/12", *result.MonitoredPrefixes[1].Prefix)
		assert.Nil(t, result.MonitoredPrefixes[1].Ge, "zero ge is treated as unset")
		assert.Nil(t, result.MonitoredPrefixes[1].Le, "zero le is treated as unset")
	})
	t.Run("missing prefix is an error", func(t *testing.T) {
		_, err := vpc.ResourceIBMIsDynamicRouteServerPeerGroupMapToDynamicRouteServerPeerGroupPrefixWatcherPrototype(map[string]interface{}{"monitored_prefixes": []interface{}{map[string]interface{}{"ge": 8}}})
		assert.NotNil(t, err)
		_, err = vpc.ResourceIBMIsDynamicRouteServerPeerGroupMapToDynamicRouteServerPeerGroupPrefixWatcherPrototype(map[string]interface{}{"excluded_prefixes": []interface{}{"not-a-block"}})
		assert.NotNil(t, err)
	})
}

func TestResourceIBMIsDynamicRouteServerPeerGroupMapToDynamicRouteServerPeerGroupPrefixWatcherPatch(t *testing.T) {
	result, err := vpc.ResourceIBMIsDynamicRouteServerPeerGroupMapToDynamicRouteServerPeerGroupPrefixWatcherPatch(map[string]interface{}{
		"monitored_prefixes": []interface{}{map[string]interface{}{"prefix": "192.168.0.0/16", "ge": 16, "le": 24}},
	})
	assert.Nil(t, err)
	assert.Equal(t, []vpcv1.DynamicRouteServerPeerGroupPrefixWatcherExcludedPrefixPrototype{}, result.ExcludedPrefixes)
	assert.Len(t, result.MonitoredPrefixes, 1)
	assert.Equal(t, int64(24), *result.MonitoredPrefixes[0].Le)

	_, err = vpc.ResourceIBMIsDynamicRouteServerPeerGroupMapToDynamicRouteServerPeerGroupPrefixWatcherPatch(map[string]interface{}{"excluded_prefixes": []interface{}{map[string]interface{}{"prefix": ""}}})
	assert.NotNil(t, err)
}

func TestResourceIBMIsDynamicRouteServerPeerGroupMapToDynamicRouteServerPeerGroupPrefixWatcherExcludedPrefixPrototype(t *testing.T) {
	result, err := vpc.ResourceIBMIsDynamicRouteServerPeerGroupMapToDynamicRouteServerPeerGroupPrefixWatcherExcludedPrefixPrototype(map[string]interface{}{"prefix": "10.0.0.0/8", "ge": 8, "le": 16})
	assert.Nil(t, err)
	assert.Equal(t, "10.0.0.0/8", *result.Prefix)
	assert.Equal(t, int64(8), *result.Ge)
	assert.Equal(t, int64(16), *result.Le)

	result, err = vpc.ResourceIBMIsDynamicRouteServerPeerGroupMapToDynamicRouteServerPeerGroupPrefixWatcherExcludedPrefixPrototype(map[string]interface{}{"prefix": "10.0.0.0/8"})
	assert.Nil(t, err)
	assert.Nil(t, result.Ge)
	assert.Nil(t, result.Le)

	_, err = vpc.ResourceIBMIsDynamicRouteServerPeerGroupMapToDynamicRouteServerPeerGroupPrefixWatcherExcludedPrefixPrototype(map[string]interface{}{})
	assert.NotNil(t, err)
	_, err = vpc.ResourceIBMIsDynamicRouteServerPeerGroupMapToDynamicRouteServerPeerGroupPrefixWatcherExcludedPrefixPrototype(nil)
	assert.NotNil(t, err)
}

func TestResourceIBMIsDynamicRouteServerPeerGroupMapToDynamicRouteServerPeerGroupPrefixWatcherMonitoredPrefixPrototype(t *testing.T) {
	result, err := vpc.ResourceIBMIsDynamicRouteServerPeerGroupMapToDynamicRouteServerPeerGroupPrefixWatcherMonitoredPrefixPrototype(map[string]interface{}{"prefix": "192.168.0.0/16", "ge": 16, "le": 24})
	assert.Nil(t, err)
	assert.Equal(t, "192.168.0.0/16", *result.Prefix)
	assert.Equal(t, int64(16), *result.Ge)
	assert.Equal(t, int64(24), *result.Le)

	_, err = vpc.ResourceIBMIsDynamicRouteServerPeerGroupMapToDynamicRouteServerPeerGroupPrefixWatcherMonitoredPrefixPrototype(map[string]interface{}{"ge": 16})
	assert.NotNil(t, err)
}

func TestResourceIBMIsDynamicRouteServerPeerGroupDynamicRouteServerPeerGroupPatchAsPatch(t *testing.T) {
	t.Run("nil patch gives empty map", func(t *testing.T) {
		assert.Equal(t, map[string]interface{}{}, vpc.ResourceIBMIsDynamicRouteServerPeerGroupDynamicRouteServerPeerGroupPatchAsPatch(nil, nil))
	})
	t.Run("only changed fields are sent", func(t *testing.T) {
		patch := vpc.ResourceIBMIsDynamicRouteServerPeerGroupDynamicRouteServerPeerGroupPatchAsPatch(&vpcv1.DynamicRouteServerPeerGroupPatch{Name: core.StringPtr("new-name")}, nil)
		assert.Equal(t, map[string]interface{}{"name": "new-name"}, patch)
	})
	t.Run("every patchable field", func(t *testing.T) {
		patch := vpc.ResourceIBMIsDynamicRouteServerPeerGroupDynamicRouteServerPeerGroupPatchAsPatch(&vpcv1.DynamicRouteServerPeerGroupPatch{
			HealthMonitor: &vpcv1.DynamicRouteServerPeerGroupHealthMonitorPatch{Mode: core.StringPtr("bfd")},
			Name:          core.StringPtr("new-name"),
			State:         core.StringPtr("disabled"),
			PrefixWatcher: &vpcv1.DynamicRouteServerPeerGroupPrefixWatcherPatch{
				ExcludedPrefixes:  []vpcv1.DynamicRouteServerPeerGroupPrefixWatcherExcludedPrefixPrototype{},
				MonitoredPrefixes: []vpcv1.DynamicRouteServerPeerGroupPrefixWatcherMonitoredPrefixPrototype{{Prefix: core.StringPtr("10.0.0.0/8"), Ge: core.Int64Ptr(8)}},
			},
		}, nil)
		assert.Equal(t, map[string]interface{}{"mode": "bfd"}, patch["health_monitor"])
		assert.Equal(t, "new-name", patch["name"])
		assert.Equal(t, "disabled", patch["state"])
		prefixWatcher := patch["prefix_watcher"].(map[string]interface{})
		assert.Equal(t, []map[string]interface{}{}, prefixWatcher["excluded_prefixes"], "an emptied list is sent as empty, not omitted")
		assert.Equal(t, []map[string]interface{}{{"prefix": "10.0.0.0/8", "ge": int64(8)}}, prefixWatcher["monitored_prefixes"])
	})
}

func TestResourceIBMIsDynamicRouteServerPeerGroupSchemaShapes(t *testing.T) {
	res := vpc.ResourceIBMIsDynamicRouteServerPeerGroup()
	assert.NotNil(t, res.Importer)
	assert.NotNil(t, res.Timeouts)
	assert.True(t, res.Schema["dynamic_route_server_id"].ForceNew)
	assert.True(t, res.Schema["roles"].ForceNew)
	assert.True(t, res.Schema["peers"].ForceNew)
	assert.Equal(t, 1, res.Schema["health_monitor"].MaxItems)
	assert.Equal(t, 1, res.Schema["prefix_watcher"].MaxItems)
	healthMonitor := res.Schema["health_monitor"].Elem.(*schema.Resource).Schema
	assert.True(t, healthMonitor["roles"].ForceNew)
	assert.False(t, healthMonitor["mode"].ForceNew)
	peer := res.Schema["peers"].Elem.(*schema.Resource).Schema
	assert.Equal(t, 1, peer["endpoint"].MaxItems)
	for _, key := range []string{"created_at", "href", "lifecycle_state", "status", "resource_type", "dynamic_route_server_peer_group_id", "etag"} {
		assert.True(t, res.Schema[key].Computed, key)
		assert.False(t, res.Schema[key].Optional, key)
		assert.False(t, res.Schema[key].Required, key)
	}
	validator := vpc.ResourceIBMIsDynamicRouteServerPeerGroupValidator()
	identifiers := []string{}
	for _, entry := range validator.Schema {
		identifiers = append(identifiers, entry.Identifier)
	}
	for _, key := range []string{"dynamic_route_server_id", "name", "state", "health_monitor_mode"} {
		assert.Contains(t, strings.Join(identifiers, ","), key)
	}
}
