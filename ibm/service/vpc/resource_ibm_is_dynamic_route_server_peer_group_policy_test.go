// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"fmt"
	"regexp"
	"strings"
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

const drsPolicyResourceName = "ibm_is_dynamic_route_server_peer_group_policy.is_dynamic_route_server_peer_group_policy_instance"

// testAccCheckIBMIsDynamicRouteServerPolicyDependencies builds the real dependency chain shared by the policy and
// static route tests: VPC, subnet, virtual network interface, dynamic route server, a peer group that acts as the
// policy target, a second peer group that acts as next hop / learned-routes source and a VPC routing table.
func testAccCheckIBMIsDynamicRouteServerPolicyDependencies(suffix string) string {
	return fmt.Sprintf(`
		resource "ibm_is_vpc" "testacc_vpc" {
			name = "tf-drs-vpc-%[1]s"
		}

		resource "ibm_is_subnet" "testacc_subnet" {
			name            = "tf-drs-subnet-%[1]s"
			vpc             = ibm_is_vpc.testacc_vpc.id
			zone            = "%[2]s"
			ipv4_cidr_block = "%[3]s"
		}

		resource "ibm_is_virtual_network_interface" "testacc_vni" {
			name   = "tf-drs-vni-%[1]s"
			subnet = ibm_is_subnet.testacc_subnet.id
		}

		resource "ibm_is_vpc_routing_table" "testacc_rt" {
			vpc  = ibm_is_vpc.testacc_vpc.id
			name = "tf-drs-rt-%[1]s"
		}

		resource "ibm_is_dynamic_route_server" "testacc_drs" {
			name      = "tf-drs-%[1]s"
			local_asn = 64512
			members {
				name = "tf-drs-member-%[1]s"
				virtual_network_interfaces {
					id = ibm_is_virtual_network_interface.testacc_vni.id
				}
			}
		}

		resource "ibm_is_dynamic_route_server_peer_group" "testacc_pg" {
			dynamic_route_server_id = ibm_is_dynamic_route_server.testacc_drs.id
			name                    = "tf-drs-pg-%[1]s"
			roles                   = ["next_hop"]
			health_monitor {
				mode  = "bfd"
				roles = ["next_hop"]
			}
			peers {
				name     = "tf-drs-peer-%[1]s"
				asn      = 64520
				priority = 100
				endpoint {
					address = "10.240.0.10"
				}
			}
		}

		resource "ibm_is_dynamic_route_server_peer_group" "testacc_pg2" {
			dynamic_route_server_id = ibm_is_dynamic_route_server.testacc_drs.id
			name                    = "tf-drs-pg2-%[1]s"
			roles                   = ["next_hop"]
			health_monitor {
				mode  = "bfd"
				roles = ["next_hop"]
			}
			peers {
				name     = "tf-drs-peer2-%[1]s"
				asn      = 64521
				priority = 110
				endpoint {
					address = "10.240.0.11"
				}
			}
		}
	`, suffix, acc.ISZoneName, acc.ISCIDR)
}

func TestAccIBMIsDynamicRouteServerPeerGroupPolicyBasic(t *testing.T) {
	var conf vpcv1.DynamicRouteServerPeerGroupPolicy
	var firstID string
	suffix := fmt.Sprintf("%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIsDynamicRouteServerPeerGroupPolicyDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsDynamicRouteServerPeerGroupPolicyConfigBasic(suffix, "vpc_address_prefixes"),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIsDynamicRouteServerPeerGroupPolicyExists(drsPolicyResourceName, conf),
					testAccCheckIBMIsDynamicRouteServerPeerGroupPolicyRecordID(drsPolicyResourceName, &firstID),
					resource.TestCheckResourceAttr(drsPolicyResourceName, "type", "vpc_address_prefixes"),
					resource.TestCheckResourceAttr(drsPolicyResourceName, "state", "enabled"),
					resource.TestCheckResourceAttr(drsPolicyResourceName, "route_delete_delay", "0"),
					resource.TestCheckResourceAttr(drsPolicyResourceName, "lifecycle_state", "stable"),
					resource.TestCheckResourceAttrSet(drsPolicyResourceName, "name"),
					resource.TestCheckResourceAttrSet(drsPolicyResourceName, "href"),
					resource.TestCheckResourceAttrSet(drsPolicyResourceName, "created_at"),
					resource.TestCheckResourceAttrSet(drsPolicyResourceName, "dynamic_route_server_peer_group_policy_id"),
				),
			},
			resource.TestStep{
				// type is ForceNew: changing it must re-create the policy.
				Config: testAccCheckIBMIsDynamicRouteServerPeerGroupPolicyConfigBasic(suffix, "custom_routes"),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIsDynamicRouteServerPeerGroupPolicyExists(drsPolicyResourceName, conf),
					testAccCheckIBMIsDynamicRouteServerPeerGroupPolicyIDChanged(drsPolicyResourceName, &firstID),
					resource.TestCheckResourceAttr(drsPolicyResourceName, "type", "custom_routes"),
				),
			},
		},
	})
}

// testAccCheckIBMIsDynamicRouteServerPeerGroupPolicyRecordID stores the current resource ID so a later step can
// verify that a ForceNew change produced a new resource.
func testAccCheckIBMIsDynamicRouteServerPeerGroupPolicyRecordID(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}
		*id = rs.Primary.ID
		return nil
	}
}

func testAccCheckIBMIsDynamicRouteServerPeerGroupPolicyIDChanged(n string, previous *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}
		if rs.Primary.ID == *previous {
			return fmt.Errorf("expected %s to be re-created, ID %s is unchanged", n, rs.Primary.ID)
		}
		return nil
	}
}

func TestAccIBMIsDynamicRouteServerPeerGroupPolicyCustomRoutes(t *testing.T) {
	var conf vpcv1.DynamicRouteServerPeerGroupPolicy
	suffix := fmt.Sprintf("%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf-drs-policy-%s", suffix)
	nameUpdate := fmt.Sprintf("tf-drs-policy-upd-%s", suffix)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIsDynamicRouteServerPeerGroupPolicyDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsDynamicRouteServerPeerGroupPolicyConfigCustomRoutes(suffix, name, "disabled", []string{"192.168.3.0/24"}),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIsDynamicRouteServerPeerGroupPolicyExists(drsPolicyResourceName, conf),
					resource.TestCheckResourceAttr(drsPolicyResourceName, "name", name),
					resource.TestCheckResourceAttr(drsPolicyResourceName, "state", "disabled"),
					resource.TestCheckResourceAttr(drsPolicyResourceName, "type", "custom_routes"),
					resource.TestCheckResourceAttr(drsPolicyResourceName, "custom_routes.#", "1"),
					resource.TestCheckResourceAttr(drsPolicyResourceName, "custom_routes.0.destination", "192.168.3.0/24"),
				),
			},
			resource.TestStep{
				// Update every patchable field of this subtype: name, state and custom_routes.
				Config: testAccCheckIBMIsDynamicRouteServerPeerGroupPolicyConfigCustomRoutes(suffix, nameUpdate, "enabled", []string{"192.168.3.0/24", "192.168.4.0/24"}),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(drsPolicyResourceName, "name", nameUpdate),
					resource.TestCheckResourceAttr(drsPolicyResourceName, "state", "enabled"),
					resource.TestCheckResourceAttr(drsPolicyResourceName, "custom_routes.#", "2"),
					resource.TestCheckResourceAttr(drsPolicyResourceName, "custom_routes.1.destination", "192.168.4.0/24"),
				),
			},
			resource.TestStep{
				ResourceName:      drsPolicyResourceName,
				ImportState:       true,
				ImportStateVerify: true,
				// etag is a response header and is not part of the imported state.
				ImportStateVerifyIgnore: []string{"etag"},
			},
		},
	})
}

func TestAccIBMIsDynamicRouteServerPeerGroupPolicyLearnedRoutes(t *testing.T) {
	var conf vpcv1.DynamicRouteServerPeerGroupPolicy
	suffix := fmt.Sprintf("%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf-drs-policy-%s", suffix)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIsDynamicRouteServerPeerGroupPolicyDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsDynamicRouteServerPeerGroupPolicyConfigLearnedRoutes(suffix, name, "10.10.0.0/16", 0, 0),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIsDynamicRouteServerPeerGroupPolicyExists(drsPolicyResourceName, conf),
					resource.TestCheckResourceAttr(drsPolicyResourceName, "type", "learned_routes"),
					resource.TestCheckResourceAttr(drsPolicyResourceName, "peer_groups.#", "1"),
					resource.TestCheckResourceAttrPair(drsPolicyResourceName, "peer_groups.0.id", "ibm_is_dynamic_route_server_peer_group.testacc_pg2", "dynamic_route_server_peer_group_id"),
					resource.TestCheckResourceAttrSet(drsPolicyResourceName, "peer_groups.0.href"),
					resource.TestCheckResourceAttrSet(drsPolicyResourceName, "peer_groups.0.name"),
					resource.TestCheckResourceAttr(drsPolicyResourceName, "excluded_prefixes.#", "1"),
					resource.TestCheckResourceAttr(drsPolicyResourceName, "excluded_prefixes.0.prefix", "10.10.0.0/16"),
					resource.TestCheckResourceAttr(drsPolicyResourceName, "excluded_prefixes.0.ge", "0"),
					resource.TestCheckResourceAttr(drsPolicyResourceName, "excluded_prefixes.0.le", "0"),
				),
			},
			resource.TestStep{
				// Update excluded_prefixes (patchable) in place.
				Config: testAccCheckIBMIsDynamicRouteServerPeerGroupPolicyConfigLearnedRoutes(suffix, name, "10.20.0.0/16", 20, 24),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(drsPolicyResourceName, "excluded_prefixes.0.prefix", "10.20.0.0/16"),
					resource.TestCheckResourceAttr(drsPolicyResourceName, "excluded_prefixes.0.ge", "20"),
					resource.TestCheckResourceAttr(drsPolicyResourceName, "excluded_prefixes.0.le", "24"),
				),
			},
			resource.TestStep{
				ResourceName:            drsPolicyResourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"etag"},
			},
		},
	})
}

func TestAccIBMIsDynamicRouteServerPeerGroupPolicyVPCRoutingTables(t *testing.T) {
	var conf vpcv1.DynamicRouteServerPeerGroupPolicy
	suffix := fmt.Sprintf("%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf-drs-policy-%s", suffix)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIsDynamicRouteServerPeerGroupPolicyDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsDynamicRouteServerPeerGroupPolicyConfigVPCRoutingTables(suffix, name, 0, true),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIsDynamicRouteServerPeerGroupPolicyExists(drsPolicyResourceName, conf),
					resource.TestCheckResourceAttr(drsPolicyResourceName, "type", "vpc_routing_tables"),
					resource.TestCheckResourceAttr(drsPolicyResourceName, "route_delete_delay", "0"),
					resource.TestCheckResourceAttr(drsPolicyResourceName, "next_hops.#", "1"),
					resource.TestCheckResourceAttrPair(drsPolicyResourceName, "next_hops.0.id", "ibm_is_dynamic_route_server_peer_group.testacc_pg2", "dynamic_route_server_peer_group_id"),
					resource.TestCheckResourceAttrSet(drsPolicyResourceName, "next_hops.0.name"),
					resource.TestCheckResourceAttr(drsPolicyResourceName, "routing_tables.#", "1"),
					resource.TestCheckResourceAttr(drsPolicyResourceName, "routing_tables.0.advertise", "true"),
					resource.TestCheckResourceAttrPair(drsPolicyResourceName, "routing_tables.0.vpc_routing_table.0.id", "ibm_is_vpc_routing_table.testacc_rt", "routing_table"),
					resource.TestCheckResourceAttrSet(drsPolicyResourceName, "routing_tables.0.vpc_routing_table.0.name"),
				),
			},
			resource.TestStep{
				// Update route_delete_delay and routing_tables.advertise (patchable).
				Config: testAccCheckIBMIsDynamicRouteServerPeerGroupPolicyConfigVPCRoutingTables(suffix, name, 30, false),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(drsPolicyResourceName, "route_delete_delay", "30"),
					resource.TestCheckResourceAttr(drsPolicyResourceName, "routing_tables.0.advertise", "false"),
				),
			},
			resource.TestStep{
				ResourceName:            drsPolicyResourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"etag"},
			},
		},
	})
}

func TestAccIBMIsDynamicRouteServerPeerGroupPolicyInvalidCombination(t *testing.T) {
	suffix := fmt.Sprintf("%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIsDynamicRouteServerPeerGroupPolicyDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				// custom_routes is only valid for type custom_routes: the provider must reject this at plan time.
				Config: testAccCheckIBMIsDynamicRouteServerPolicyDependencies(suffix) + `
					resource "ibm_is_dynamic_route_server_peer_group_policy" "is_dynamic_route_server_peer_group_policy_instance" {
						dynamic_route_server_id            = ibm_is_dynamic_route_server.testacc_drs.id
						dynamic_route_server_peer_group_id = ibm_is_dynamic_route_server_peer_group.testacc_pg.dynamic_route_server_peer_group_id
						type                               = "vpc_address_prefixes"
						custom_routes {
							destination = "192.168.3.0/24"
						}
					}
				`,
				PlanOnly:    true,
				ExpectError: regexp.MustCompile(`"custom_routes" cannot be set when type is "vpc_address_prefixes"`),
			},
		},
	})
}

func testAccCheckIBMIsDynamicRouteServerPeerGroupPolicyConfigBasic(suffix string, typeVar string) string {
	return testAccCheckIBMIsDynamicRouteServerPolicyDependencies(suffix) + fmt.Sprintf(`
		resource "ibm_is_dynamic_route_server_peer_group_policy" "is_dynamic_route_server_peer_group_policy_instance" {
			dynamic_route_server_id            = ibm_is_dynamic_route_server.testacc_drs.id
			dynamic_route_server_peer_group_id = ibm_is_dynamic_route_server_peer_group.testacc_pg.dynamic_route_server_peer_group_id
			type                               = "%s"
		}
	`, typeVar)
}

func testAccCheckIBMIsDynamicRouteServerPeerGroupPolicyConfigCustomRoutes(suffix string, name string, state string, destinations []string) string {
	routes := ""
	for _, destination := range destinations {
		routes += fmt.Sprintf(`
			custom_routes {
				destination = "%s"
			}`, destination)
	}
	return testAccCheckIBMIsDynamicRouteServerPolicyDependencies(suffix) + fmt.Sprintf(`
		resource "ibm_is_dynamic_route_server_peer_group_policy" "is_dynamic_route_server_peer_group_policy_instance" {
			dynamic_route_server_id            = ibm_is_dynamic_route_server.testacc_drs.id
			dynamic_route_server_peer_group_id = ibm_is_dynamic_route_server_peer_group.testacc_pg.dynamic_route_server_peer_group_id
			name                               = "%s"
			state                              = "%s"
			type                               = "custom_routes"
			%s
		}
	`, name, state, routes)
}

func testAccCheckIBMIsDynamicRouteServerPeerGroupPolicyConfigLearnedRoutes(suffix string, name string, prefix string, ge int, le int) string {
	return testAccCheckIBMIsDynamicRouteServerPolicyDependencies(suffix) + fmt.Sprintf(`
		resource "ibm_is_dynamic_route_server_peer_group_policy" "is_dynamic_route_server_peer_group_policy_instance" {
			dynamic_route_server_id            = ibm_is_dynamic_route_server.testacc_drs.id
			dynamic_route_server_peer_group_id = ibm_is_dynamic_route_server_peer_group.testacc_pg.dynamic_route_server_peer_group_id
			name                               = "%s"
			type                               = "learned_routes"
			peer_groups {
				id = ibm_is_dynamic_route_server_peer_group.testacc_pg2.dynamic_route_server_peer_group_id
			}
			excluded_prefixes {
				prefix = "%s"
				ge     = %d
				le     = %d
			}
		}
	`, name, prefix, ge, le)
}

func testAccCheckIBMIsDynamicRouteServerPeerGroupPolicyConfigVPCRoutingTables(suffix string, name string, routeDeleteDelay int, advertise bool) string {
	return testAccCheckIBMIsDynamicRouteServerPolicyDependencies(suffix) + fmt.Sprintf(`
		resource "ibm_is_dynamic_route_server_peer_group_policy" "is_dynamic_route_server_peer_group_policy_instance" {
			dynamic_route_server_id            = ibm_is_dynamic_route_server.testacc_drs.id
			dynamic_route_server_peer_group_id = ibm_is_dynamic_route_server_peer_group.testacc_pg.dynamic_route_server_peer_group_id
			name                               = "%s"
			type                               = "vpc_routing_tables"
			route_delete_delay                 = %d
			next_hops {
				id = ibm_is_dynamic_route_server_peer_group.testacc_pg2.dynamic_route_server_peer_group_id
			}
			routing_tables {
				advertise = %t
				vpc_routing_table {
					id = ibm_is_vpc_routing_table.testacc_rt.routing_table
				}
			}
		}
	`, name, routeDeleteDelay, advertise)
}

func testAccCheckIBMIsDynamicRouteServerPeerGroupPolicyExists(n string, obj vpcv1.DynamicRouteServerPeerGroupPolicy) resource.TestCheckFunc {
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

		getOptions := &vpcv1.GetDynamicRouteServerPeerGroupPolicyOptions{}
		getOptions.SetDynamicRouteServerID(parts[0])
		getOptions.SetDynamicRouteServerPeerGroupID(parts[1])
		getOptions.SetID(parts[2])

		policyIntf, _, err := vpcClient.GetDynamicRouteServerPeerGroupPolicy(getOptions)
		if err != nil {
			return err
		}
		policy, err := vpc.ResourceIBMIsDynamicRouteServerPeerGroupPolicyToBase(policyIntf)
		if err != nil {
			return err
		}
		obj = *policy
		return nil
	}
}

func testAccCheckIBMIsDynamicRouteServerPeerGroupPolicyDestroy(s *terraform.State) error {
	vpcClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).VpcV1API()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_is_dynamic_route_server_peer_group_policy" {
			continue
		}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}
		if len(parts) != 3 {
			return fmt.Errorf("unexpected ID format %q", rs.Primary.ID)
		}

		getOptions := &vpcv1.GetDynamicRouteServerPeerGroupPolicyOptions{}
		getOptions.SetDynamicRouteServerID(parts[0])
		getOptions.SetDynamicRouteServerPeerGroupID(parts[1])
		getOptions.SetID(parts[2])

		_, response, err := vpcClient.GetDynamicRouteServerPeerGroupPolicy(getOptions)
		if err == nil {
			return fmt.Errorf("ibm_is_dynamic_route_server_peer_group_policy still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for ibm_is_dynamic_route_server_peer_group_policy (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}

// ---- unit tests ----

func drsPolicyTestPeerGroupReference(id string) *vpcv1.DynamicRouteServerPeerGroupReference {
	return &vpcv1.DynamicRouteServerPeerGroupReference{
		Href:         core.StringPtr("https://us-south.iaas.cloud.ibm.com/v1/dynamic_route_servers/r006-d7cc/peer_groups/" + id),
		ID:           core.StringPtr(id),
		Name:         core.StringPtr("my-peer-group-" + id),
		ResourceType: core.StringPtr("dynamic_route_server_peer_group"),
	}
}

func drsPolicyTestRoutingTableReference() *vpcv1.RoutingTableReference {
	return &vpcv1.RoutingTableReference{
		CRN:          core.StringPtr("crn:v1:bluemix:public:is:us-south:a/123456::vpc-routing-table:r006-4727/r006-6885"),
		Href:         core.StringPtr("https://us-south.iaas.cloud.ibm.com/v1/vpcs/r006-4727/routing_tables/r006-6885"),
		ID:           core.StringPtr("r006-6885"),
		Name:         core.StringPtr("my-routing-table-1"),
		ResourceType: core.StringPtr("routing_table"),
	}
}

func TestResourceIBMIsDynamicRouteServerPeerGroupPolicyToBase(t *testing.T) {
	t.Run("base struct is returned as is", func(t *testing.T) {
		base := &vpcv1.DynamicRouteServerPeerGroupPolicy{ID: core.StringPtr("p1"), Type: core.StringPtr("custom_routes")}
		result, err := vpc.ResourceIBMIsDynamicRouteServerPeerGroupPolicyToBase(base)
		assert.Nil(t, err)
		assert.Equal(t, base, result)
	})
	t.Run("custom routes subtype", func(t *testing.T) {
		model := &vpcv1.DynamicRouteServerPeerGroupPolicyCustomRoutes{
			ID:             core.StringPtr("p1"),
			Type:           core.StringPtr("custom_routes"),
			LifecycleState: core.StringPtr("stable"),
			CustomRoutes:   []vpcv1.DynamicRouteServerPeerGroupPolicyCustomRoute{{Destination: core.StringPtr("192.168.3.0/24")}},
		}
		result, err := vpc.ResourceIBMIsDynamicRouteServerPeerGroupPolicyToBase(model)
		assert.Nil(t, err)
		assert.Equal(t, "p1", *result.ID)
		assert.Equal(t, "stable", *result.LifecycleState)
		assert.Len(t, result.CustomRoutes, 1)
	})
	t.Run("learned routes subtype", func(t *testing.T) {
		model := &vpcv1.DynamicRouteServerPeerGroupPolicyLearnedRoutes{
			ID:         core.StringPtr("p2"),
			Type:       core.StringPtr("learned_routes"),
			PeerGroups: []vpcv1.DynamicRouteServerPeerGroupReference{*drsPolicyTestPeerGroupReference("pg1")},
		}
		result, err := vpc.ResourceIBMIsDynamicRouteServerPeerGroupPolicyToBase(model)
		assert.Nil(t, err)
		assert.Len(t, result.PeerGroups, 1)
		assert.Nil(t, result.CustomRoutes)
	})
	t.Run("vpc address prefixes subtype", func(t *testing.T) {
		model := &vpcv1.DynamicRouteServerPeerGroupPolicyVPCAddressPrefixes{ID: core.StringPtr("p3"), Type: core.StringPtr("vpc_address_prefixes")}
		result, err := vpc.ResourceIBMIsDynamicRouteServerPeerGroupPolicyToBase(model)
		assert.Nil(t, err)
		assert.Equal(t, "vpc_address_prefixes", *result.Type)
	})
	t.Run("vpc routing tables subtype", func(t *testing.T) {
		model := &vpcv1.DynamicRouteServerPeerGroupPolicyVPCRoutingTables{
			ID:               core.StringPtr("p4"),
			Type:             core.StringPtr("vpc_routing_tables"),
			RouteDeleteDelay: core.Int64Ptr(30),
			NextHops:         []vpcv1.DynamicRouteServerPeerGroupPolicyNextHopIntf{&vpcv1.DynamicRouteServerPeerGroupPolicyNextHopDynamicRouteServerPeerGroupReference{ID: core.StringPtr("pg2")}},
			RoutingTables:    []vpcv1.DynamicRouteServerPeerGroupPolicyVPCRoutingTablesRoutingTable{{Advertise: core.BoolPtr(true), VPCRoutingTable: drsPolicyTestRoutingTableReference()}},
		}
		result, err := vpc.ResourceIBMIsDynamicRouteServerPeerGroupPolicyToBase(model)
		assert.Nil(t, err)
		assert.Equal(t, int64(30), *result.RouteDeleteDelay)
		assert.Len(t, result.NextHops, 1)
		assert.Len(t, result.RoutingTables, 1)
	})
	t.Run("nil subtype pointer is an error", func(t *testing.T) {
		var model *vpcv1.DynamicRouteServerPeerGroupPolicyCustomRoutes
		_, err := vpc.ResourceIBMIsDynamicRouteServerPeerGroupPolicyToBase(model)
		assert.NotNil(t, err)
	})
	t.Run("nil interface is an error", func(t *testing.T) {
		_, err := vpc.ResourceIBMIsDynamicRouteServerPeerGroupPolicyToBase(nil)
		assert.NotNil(t, err)
	})
}

func TestResourceIBMIsDynamicRouteServerPeerGroupPolicyDynamicRouteServerPeerGroupPolicyToMap(t *testing.T) {
	t.Run("nil model", func(t *testing.T) {
		_, err := vpc.ResourceIBMIsDynamicRouteServerPeerGroupPolicyDynamicRouteServerPeerGroupPolicyToMap(nil)
		assert.NotNil(t, err)
	})
	t.Run("minimal policy has empty lists and zero delay", func(t *testing.T) {
		model := &vpcv1.DynamicRouteServerPeerGroupPolicy{
			ID:    core.StringPtr("p1"),
			Name:  core.StringPtr("my-policy"),
			State: core.StringPtr("enabled"),
			Type:  core.StringPtr("vpc_address_prefixes"),
		}
		result, err := vpc.ResourceIBMIsDynamicRouteServerPeerGroupPolicyDynamicRouteServerPeerGroupPolicyToMap(model)
		assert.Nil(t, err)
		assert.Equal(t, "p1", result["id"])
		assert.Equal(t, 0, result["route_delete_delay"])
		assert.Equal(t, []map[string]interface{}{}, result["custom_routes"])
		assert.Equal(t, []map[string]interface{}{}, result["next_hops"])
		assert.Equal(t, []map[string]interface{}{}, result["lifecycle_reasons"])
		_, hasCreatedAt := result["created_at"]
		assert.False(t, hasCreatedAt)
	})
	t.Run("full vpc routing tables policy", func(t *testing.T) {
		model := &vpcv1.DynamicRouteServerPeerGroupPolicy{
			ID:               core.StringPtr("p4"),
			Href:             core.StringPtr("https://us-south.iaas.cloud.ibm.com/v1/dynamic_route_servers/r006/peer_groups/pg/policies/p4"),
			Name:             core.StringPtr("my-policy"),
			State:            core.StringPtr("enabled"),
			Type:             core.StringPtr("vpc_routing_tables"),
			LifecycleState:   core.StringPtr("stable"),
			LifecycleReasons: []vpcv1.DynamicRouteServerPeerGroupPolicyLifecycleReason{{Code: core.StringPtr("internal_error"), Message: core.StringPtr("internal error")}},
			ResourceType:     core.StringPtr("dynamic_route_server_peer_group_policy"),
			RouteDeleteDelay: core.Int64Ptr(45),
			ExcludedPrefixes: []vpcv1.DynamicRouteServerPeerGroupPrefixWatcherExcludedPrefixPrototype{{Prefix: core.StringPtr("10.0.0.0/8")}},
			NextHops: []vpcv1.DynamicRouteServerPeerGroupPolicyNextHopIntf{
				&vpcv1.DynamicRouteServerPeerGroupPolicyNextHopDynamicRouteServerPeerGroupReference{ID: core.StringPtr("pg2"), Name: core.StringPtr("pg2")},
				&vpcv1.DynamicRouteServerPeerGroupPolicyNextHop{ID: core.StringPtr("pg3")},
			},
			RoutingTables: []vpcv1.DynamicRouteServerPeerGroupPolicyVPCRoutingTablesRoutingTable{
				{Advertise: core.BoolPtr(false), VPCRoutingTable: drsPolicyTestRoutingTableReference()},
			},
		}
		result, err := vpc.ResourceIBMIsDynamicRouteServerPeerGroupPolicyDynamicRouteServerPeerGroupPolicyToMap(model)
		assert.Nil(t, err)
		assert.Equal(t, 45, result["route_delete_delay"])
		assert.Equal(t, "stable", result["lifecycle_state"])
		assert.Len(t, result["lifecycle_reasons"], 1)
		assert.Len(t, result["next_hops"], 2)
		nextHops := result["next_hops"].([]map[string]interface{})
		assert.Equal(t, "pg2", nextHops[0]["id"])
		assert.Equal(t, "pg3", nextHops[1]["id"])
		routingTables := result["routing_tables"].([]map[string]interface{})
		assert.Equal(t, false, routingTables[0]["advertise"])
		assert.Equal(t, "r006-6885", routingTables[0]["vpc_routing_table"].([]map[string]interface{})[0]["id"])
		excluded := result["excluded_prefixes"].([]map[string]interface{})
		assert.Equal(t, 0, excluded[0]["ge"])
		assert.Equal(t, 0, excluded[0]["le"])
	})
	t.Run("unrecognized next hop subtype propagates error", func(t *testing.T) {
		model := &vpcv1.DynamicRouteServerPeerGroupPolicy{
			NextHops: []vpcv1.DynamicRouteServerPeerGroupPolicyNextHopIntf{nil},
		}
		_, err := vpc.ResourceIBMIsDynamicRouteServerPeerGroupPolicyDynamicRouteServerPeerGroupPolicyToMap(model)
		assert.NotNil(t, err)
	})
}

func TestResourceIBMIsDynamicRouteServerPeerGroupPolicyDynamicRouteServerPeerGroupPolicyCustomRouteToMap(t *testing.T) {
	t.Run("nil", func(t *testing.T) {
		_, err := vpc.ResourceIBMIsDynamicRouteServerPeerGroupPolicyDynamicRouteServerPeerGroupPolicyCustomRouteToMap(nil)
		assert.NotNil(t, err)
	})
	t.Run("empty destination", func(t *testing.T) {
		result, err := vpc.ResourceIBMIsDynamicRouteServerPeerGroupPolicyDynamicRouteServerPeerGroupPolicyCustomRouteToMap(&vpcv1.DynamicRouteServerPeerGroupPolicyCustomRoute{})
		assert.Nil(t, err)
		assert.Equal(t, map[string]interface{}{}, result)
	})
	t.Run("destination", func(t *testing.T) {
		result, err := vpc.ResourceIBMIsDynamicRouteServerPeerGroupPolicyDynamicRouteServerPeerGroupPolicyCustomRouteToMap(&vpcv1.DynamicRouteServerPeerGroupPolicyCustomRoute{Destination: core.StringPtr("192.168.3.0/24")})
		assert.Nil(t, err)
		assert.Equal(t, map[string]interface{}{"destination": "192.168.3.0/24"}, result)
	})
}

func TestResourceIBMIsDynamicRouteServerPeerGroupPolicyDynamicRouteServerPeerGroupPrefixWatcherExcludedPrefixPrototypeToMap(t *testing.T) {
	t.Run("nil", func(t *testing.T) {
		_, err := vpc.ResourceIBMIsDynamicRouteServerPeerGroupPolicyDynamicRouteServerPeerGroupPrefixWatcherExcludedPrefixPrototypeToMap(nil)
		assert.NotNil(t, err)
	})
	t.Run("unset ge/le become zero", func(t *testing.T) {
		result, err := vpc.ResourceIBMIsDynamicRouteServerPeerGroupPolicyDynamicRouteServerPeerGroupPrefixWatcherExcludedPrefixPrototypeToMap(&vpcv1.DynamicRouteServerPeerGroupPrefixWatcherExcludedPrefixPrototype{Prefix: core.StringPtr("10.0.0.0/8")})
		assert.Nil(t, err)
		assert.Equal(t, map[string]interface{}{"ge": 0, "le": 0, "prefix": "10.0.0.0/8"}, result)
	})
	t.Run("explicit zero and non-zero values", func(t *testing.T) {
		result, err := vpc.ResourceIBMIsDynamicRouteServerPeerGroupPolicyDynamicRouteServerPeerGroupPrefixWatcherExcludedPrefixPrototypeToMap(&vpcv1.DynamicRouteServerPeerGroupPrefixWatcherExcludedPrefixPrototype{Ge: core.Int64Ptr(0), Le: core.Int64Ptr(24), Prefix: core.StringPtr("10.0.0.0/8")})
		assert.Nil(t, err)
		assert.Equal(t, map[string]interface{}{"ge": 0, "le": 24, "prefix": "10.0.0.0/8"}, result)
	})
}

func TestResourceIBMIsDynamicRouteServerPeerGroupPolicyDynamicRouteServerPeerGroupReferenceToMap(t *testing.T) {
	t.Run("nil", func(t *testing.T) {
		_, err := vpc.ResourceIBMIsDynamicRouteServerPeerGroupPolicyDynamicRouteServerPeerGroupReferenceToMap(nil)
		assert.NotNil(t, err)
	})
	t.Run("reference without deleted", func(t *testing.T) {
		result, err := vpc.ResourceIBMIsDynamicRouteServerPeerGroupPolicyDynamicRouteServerPeerGroupReferenceToMap(drsPolicyTestPeerGroupReference("pg1"))
		assert.Nil(t, err)
		assert.Equal(t, "pg1", result["id"])
		assert.Equal(t, "my-peer-group-pg1", result["name"])
		_, hasDeleted := result["deleted"]
		assert.False(t, hasDeleted)
	})
	t.Run("reference with deleted", func(t *testing.T) {
		model := drsPolicyTestPeerGroupReference("pg1")
		model.Deleted = &vpcv1.Deleted{MoreInfo: core.StringPtr("https://cloud.ibm.com/apidocs/vpc#deleted-resources")}
		result, err := vpc.ResourceIBMIsDynamicRouteServerPeerGroupPolicyDynamicRouteServerPeerGroupReferenceToMap(model)
		assert.Nil(t, err)
		assert.Equal(t, []map[string]interface{}{{"more_info": "https://cloud.ibm.com/apidocs/vpc#deleted-resources"}}, result["deleted"])
	})
}

func TestResourceIBMIsDynamicRouteServerPeerGroupPolicyDeletedToMap(t *testing.T) {
	_, err := vpc.ResourceIBMIsDynamicRouteServerPeerGroupPolicyDeletedToMap(nil)
	assert.NotNil(t, err)

	result, err := vpc.ResourceIBMIsDynamicRouteServerPeerGroupPolicyDeletedToMap(&vpcv1.Deleted{})
	assert.Nil(t, err)
	assert.Equal(t, map[string]interface{}{}, result)

	result, err = vpc.ResourceIBMIsDynamicRouteServerPeerGroupPolicyDeletedToMap(&vpcv1.Deleted{MoreInfo: core.StringPtr("info")})
	assert.Nil(t, err)
	assert.Equal(t, map[string]interface{}{"more_info": "info"}, result)
}

func TestResourceIBMIsDynamicRouteServerPeerGroupPolicyDynamicRouteServerPeerGroupPolicyNextHopToMap(t *testing.T) {
	t.Run("peer group reference subtype", func(t *testing.T) {
		model := &vpcv1.DynamicRouteServerPeerGroupPolicyNextHopDynamicRouteServerPeerGroupReference{
			Href:         core.StringPtr("href"),
			ID:           core.StringPtr("pg1"),
			Name:         core.StringPtr("name"),
			ResourceType: core.StringPtr("dynamic_route_server_peer_group"),
		}
		result, err := vpc.ResourceIBMIsDynamicRouteServerPeerGroupPolicyDynamicRouteServerPeerGroupPolicyNextHopToMap(model)
		assert.Nil(t, err)
		assert.Equal(t, map[string]interface{}{"href": "href", "id": "pg1", "name": "name", "resource_type": "dynamic_route_server_peer_group"}, result)
	})
	t.Run("base struct", func(t *testing.T) {
		model := &vpcv1.DynamicRouteServerPeerGroupPolicyNextHop{ID: core.StringPtr("pg1"), Deleted: &vpcv1.Deleted{MoreInfo: core.StringPtr("info")}}
		result, err := vpc.ResourceIBMIsDynamicRouteServerPeerGroupPolicyDynamicRouteServerPeerGroupPolicyNextHopToMap(model)
		assert.Nil(t, err)
		assert.Equal(t, "pg1", result["id"])
		assert.Equal(t, []map[string]interface{}{{"more_info": "info"}}, result["deleted"])
	})
	t.Run("nil pointers", func(t *testing.T) {
		var reference *vpcv1.DynamicRouteServerPeerGroupPolicyNextHopDynamicRouteServerPeerGroupReference
		_, err := vpc.ResourceIBMIsDynamicRouteServerPeerGroupPolicyDynamicRouteServerPeerGroupPolicyNextHopToMap(reference)
		assert.NotNil(t, err)
		var base *vpcv1.DynamicRouteServerPeerGroupPolicyNextHop
		_, err = vpc.ResourceIBMIsDynamicRouteServerPeerGroupPolicyDynamicRouteServerPeerGroupPolicyNextHopToMap(base)
		assert.NotNil(t, err)
	})
	t.Run("unrecognized subtype", func(t *testing.T) {
		_, err := vpc.ResourceIBMIsDynamicRouteServerPeerGroupPolicyDynamicRouteServerPeerGroupPolicyNextHopToMap(nil)
		assert.NotNil(t, err)
		assert.True(t, strings.Contains(err.Error(), "Unrecognized"))
	})
}

func TestResourceIBMIsDynamicRouteServerPeerGroupPolicyDynamicRouteServerPeerGroupPolicyNextHopDynamicRouteServerPeerGroupReferenceToMap(t *testing.T) {
	_, err := vpc.ResourceIBMIsDynamicRouteServerPeerGroupPolicyDynamicRouteServerPeerGroupPolicyNextHopDynamicRouteServerPeerGroupReferenceToMap(nil)
	assert.NotNil(t, err)

	result, err := vpc.ResourceIBMIsDynamicRouteServerPeerGroupPolicyDynamicRouteServerPeerGroupPolicyNextHopDynamicRouteServerPeerGroupReferenceToMap(&vpcv1.DynamicRouteServerPeerGroupPolicyNextHopDynamicRouteServerPeerGroupReference{ID: core.StringPtr("pg1")})
	assert.Nil(t, err)
	assert.Equal(t, map[string]interface{}{"id": "pg1"}, result)
}

func TestResourceIBMIsDynamicRouteServerPeerGroupPolicyDynamicRouteServerPeerGroupPolicyVPCRoutingTablesRoutingTableToMap(t *testing.T) {
	t.Run("nil", func(t *testing.T) {
		_, err := vpc.ResourceIBMIsDynamicRouteServerPeerGroupPolicyDynamicRouteServerPeerGroupPolicyVPCRoutingTablesRoutingTableToMap(nil)
		assert.NotNil(t, err)
	})
	t.Run("advertise unset defaults to false and no routing table", func(t *testing.T) {
		result, err := vpc.ResourceIBMIsDynamicRouteServerPeerGroupPolicyDynamicRouteServerPeerGroupPolicyVPCRoutingTablesRoutingTableToMap(&vpcv1.DynamicRouteServerPeerGroupPolicyVPCRoutingTablesRoutingTable{})
		assert.Nil(t, err)
		assert.Equal(t, map[string]interface{}{"advertise": false}, result)
	})
	t.Run("advertise true with routing table", func(t *testing.T) {
		result, err := vpc.ResourceIBMIsDynamicRouteServerPeerGroupPolicyDynamicRouteServerPeerGroupPolicyVPCRoutingTablesRoutingTableToMap(&vpcv1.DynamicRouteServerPeerGroupPolicyVPCRoutingTablesRoutingTable{Advertise: core.BoolPtr(true), VPCRoutingTable: drsPolicyTestRoutingTableReference()})
		assert.Nil(t, err)
		assert.Equal(t, true, result["advertise"])
		vpcRoutingTable := result["vpc_routing_table"].([]map[string]interface{})
		assert.Len(t, vpcRoutingTable, 1)
		assert.Equal(t, "my-routing-table-1", vpcRoutingTable[0]["name"])
	})
}

func TestResourceIBMIsDynamicRouteServerPeerGroupPolicyRoutingTableReferenceToMap(t *testing.T) {
	_, err := vpc.ResourceIBMIsDynamicRouteServerPeerGroupPolicyRoutingTableReferenceToMap(nil)
	assert.NotNil(t, err)

	model := drsPolicyTestRoutingTableReference()
	model.Deleted = &vpcv1.Deleted{MoreInfo: core.StringPtr("info")}
	result, err := vpc.ResourceIBMIsDynamicRouteServerPeerGroupPolicyRoutingTableReferenceToMap(model)
	assert.Nil(t, err)
	assert.Equal(t, *model.CRN, result["crn"])
	assert.Equal(t, "r006-6885", result["id"])
	assert.Equal(t, "routing_table", result["resource_type"])
	assert.Equal(t, []map[string]interface{}{{"more_info": "info"}}, result["deleted"])
}

func TestResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToDynamicRouteServerPeerGroupPolicyCustomRoutePrototype(t *testing.T) {
	_, err := vpc.ResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToDynamicRouteServerPeerGroupPolicyCustomRoutePrototype(map[string]interface{}{})
	assert.NotNil(t, err)

	result, err := vpc.ResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToDynamicRouteServerPeerGroupPolicyCustomRoutePrototype(map[string]interface{}{"destination": "192.168.3.0/24"})
	assert.Nil(t, err)
	assert.Equal(t, "192.168.3.0/24", *result.Destination)
}

func TestResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToDynamicRouteServerPeerGroupPrefixWatcherExcludedPrefixPrototype(t *testing.T) {
	_, err := vpc.ResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToDynamicRouteServerPeerGroupPrefixWatcherExcludedPrefixPrototype(map[string]interface{}{"ge": 0})
	assert.NotNil(t, err)

	result, err := vpc.ResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToDynamicRouteServerPeerGroupPrefixWatcherExcludedPrefixPrototype(map[string]interface{}{"prefix": "10.0.0.0/8"})
	assert.Nil(t, err)
	assert.Nil(t, result.Ge)
	assert.Nil(t, result.Le)

	result, err = vpc.ResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToDynamicRouteServerPeerGroupPrefixWatcherExcludedPrefixPrototype(map[string]interface{}{"ge": 0, "le": 24, "prefix": "10.0.0.0/8"})
	assert.Nil(t, err)
	assert.Equal(t, int64(0), *result.Ge)
	assert.Equal(t, int64(24), *result.Le)
	assert.Equal(t, "10.0.0.0/8", *result.Prefix)
}

func TestResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToDynamicRouteServerPeerGroupIdentity(t *testing.T) {
	t.Run("by id", func(t *testing.T) {
		result, err := vpc.ResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToDynamicRouteServerPeerGroupIdentity(map[string]interface{}{"id": "pg1", "href": ""})
		assert.Nil(t, err)
		assert.Equal(t, &vpcv1.DynamicRouteServerPeerGroupIdentityByID{ID: core.StringPtr("pg1")}, result)
	})
	t.Run("by href", func(t *testing.T) {
		result, err := vpc.ResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToDynamicRouteServerPeerGroupIdentity(map[string]interface{}{"href": "https://pg"})
		assert.Nil(t, err)
		assert.Equal(t, &vpcv1.DynamicRouteServerPeerGroupIdentityByHref{Href: core.StringPtr("https://pg")}, result)
	})
	t.Run("id wins over href", func(t *testing.T) {
		result, err := vpc.ResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToDynamicRouteServerPeerGroupIdentity(map[string]interface{}{"id": "pg1", "href": "https://pg"})
		assert.Nil(t, err)
		_, isByID := result.(*vpcv1.DynamicRouteServerPeerGroupIdentityByID)
		assert.True(t, isByID)
	})
	t.Run("missing both", func(t *testing.T) {
		_, err := vpc.ResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToDynamicRouteServerPeerGroupIdentity(map[string]interface{}{"name": "x"})
		assert.NotNil(t, err)
	})
	t.Run("nil map", func(t *testing.T) {
		_, err := vpc.ResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToDynamicRouteServerPeerGroupIdentity(nil)
		assert.NotNil(t, err)
	})
}

func TestResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToDynamicRouteServerPeerGroupIdentityByID(t *testing.T) {
	_, err := vpc.ResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToDynamicRouteServerPeerGroupIdentityByID(map[string]interface{}{})
	assert.NotNil(t, err)
	result, err := vpc.ResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToDynamicRouteServerPeerGroupIdentityByID(map[string]interface{}{"id": "pg1"})
	assert.Nil(t, err)
	assert.Equal(t, "pg1", *result.ID)
}

func TestResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToDynamicRouteServerPeerGroupIdentityByHref(t *testing.T) {
	_, err := vpc.ResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToDynamicRouteServerPeerGroupIdentityByHref(map[string]interface{}{})
	assert.NotNil(t, err)
	result, err := vpc.ResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToDynamicRouteServerPeerGroupIdentityByHref(map[string]interface{}{"href": "https://pg"})
	assert.Nil(t, err)
	assert.Equal(t, "https://pg", *result.Href)
}

func TestResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToDynamicRouteServerPeerGroupPolicyNextHopPrototype(t *testing.T) {
	t.Run("by id", func(t *testing.T) {
		result, err := vpc.ResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToDynamicRouteServerPeerGroupPolicyNextHopPrototype(map[string]interface{}{"id": "pg1"})
		assert.Nil(t, err)
		assert.Equal(t, &vpcv1.DynamicRouteServerPeerGroupPolicyNextHopPrototypeDynamicRouteServerPeerGroupIdentityDynamicRouteServerPeerGroupIdentityByID{ID: core.StringPtr("pg1")}, result)
	})
	t.Run("by href", func(t *testing.T) {
		result, err := vpc.ResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToDynamicRouteServerPeerGroupPolicyNextHopPrototype(map[string]interface{}{"id": "", "href": "https://pg"})
		assert.Nil(t, err)
		assert.Equal(t, &vpcv1.DynamicRouteServerPeerGroupPolicyNextHopPrototypeDynamicRouteServerPeerGroupIdentityDynamicRouteServerPeerGroupIdentityByHref{Href: core.StringPtr("https://pg")}, result)
	})
	t.Run("missing", func(t *testing.T) {
		_, err := vpc.ResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToDynamicRouteServerPeerGroupPolicyNextHopPrototype(map[string]interface{}{})
		assert.NotNil(t, err)
	})
}

func TestResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToDynamicRouteServerPeerGroupPolicyNextHopPrototypeDynamicRouteServerPeerGroupIdentity(t *testing.T) {
	result, err := vpc.ResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToDynamicRouteServerPeerGroupPolicyNextHopPrototypeDynamicRouteServerPeerGroupIdentity(map[string]interface{}{"id": "pg1"})
	assert.Nil(t, err)
	assert.NotNil(t, result)
	_, err = vpc.ResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToDynamicRouteServerPeerGroupPolicyNextHopPrototypeDynamicRouteServerPeerGroupIdentity(nil)
	assert.NotNil(t, err)
}

func TestResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToDynamicRouteServerPeerGroupPolicyNextHopPrototypeDynamicRouteServerPeerGroupIdentityDynamicRouteServerPeerGroupIdentityByID(t *testing.T) {
	_, err := vpc.ResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToDynamicRouteServerPeerGroupPolicyNextHopPrototypeDynamicRouteServerPeerGroupIdentityDynamicRouteServerPeerGroupIdentityByID(map[string]interface{}{"id": ""})
	assert.NotNil(t, err)
	result, err := vpc.ResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToDynamicRouteServerPeerGroupPolicyNextHopPrototypeDynamicRouteServerPeerGroupIdentityDynamicRouteServerPeerGroupIdentityByID(map[string]interface{}{"id": "pg1"})
	assert.Nil(t, err)
	assert.Equal(t, "pg1", *result.ID)
}

func TestResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToDynamicRouteServerPeerGroupPolicyNextHopPrototypeDynamicRouteServerPeerGroupIdentityDynamicRouteServerPeerGroupIdentityByHref(t *testing.T) {
	_, err := vpc.ResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToDynamicRouteServerPeerGroupPolicyNextHopPrototypeDynamicRouteServerPeerGroupIdentityDynamicRouteServerPeerGroupIdentityByHref(map[string]interface{}{})
	assert.NotNil(t, err)
	result, err := vpc.ResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToDynamicRouteServerPeerGroupPolicyNextHopPrototypeDynamicRouteServerPeerGroupIdentityDynamicRouteServerPeerGroupIdentityByHref(map[string]interface{}{"href": "https://pg"})
	assert.Nil(t, err)
	assert.Equal(t, "https://pg", *result.Href)
}

func TestResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToDynamicRouteServerPeerGroupPolicyVPCRoutingTablesRoutingTablePrototype(t *testing.T) {
	t.Run("advertise false with id", func(t *testing.T) {
		result, err := vpc.ResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToDynamicRouteServerPeerGroupPolicyVPCRoutingTablesRoutingTablePrototype(map[string]interface{}{
			"advertise":         false,
			"vpc_routing_table": []interface{}{map[string]interface{}{"id": "rt1"}},
		})
		assert.Nil(t, err)
		assert.Equal(t, false, *result.Advertise)
		assert.Equal(t, &vpcv1.RoutingTableIdentityByID{ID: core.StringPtr("rt1")}, result.VPCRoutingTable)
	})
	t.Run("advertise true with crn", func(t *testing.T) {
		result, err := vpc.ResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToDynamicRouteServerPeerGroupPolicyVPCRoutingTablesRoutingTablePrototype(map[string]interface{}{
			"advertise":         true,
			"vpc_routing_table": []interface{}{map[string]interface{}{"crn": "crn:rt"}},
		})
		assert.Nil(t, err)
		assert.Equal(t, true, *result.Advertise)
		assert.Equal(t, &vpcv1.RoutingTableIdentityByCRN{CRN: core.StringPtr("crn:rt")}, result.VPCRoutingTable)
	})
	t.Run("missing vpc_routing_table", func(t *testing.T) {
		_, err := vpc.ResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToDynamicRouteServerPeerGroupPolicyVPCRoutingTablesRoutingTablePrototype(map[string]interface{}{"advertise": true})
		assert.NotNil(t, err)
		_, err = vpc.ResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToDynamicRouteServerPeerGroupPolicyVPCRoutingTablesRoutingTablePrototype(map[string]interface{}{"advertise": true, "vpc_routing_table": []interface{}{nil}})
		assert.NotNil(t, err)
	})
	t.Run("empty identity", func(t *testing.T) {
		_, err := vpc.ResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToDynamicRouteServerPeerGroupPolicyVPCRoutingTablesRoutingTablePrototype(map[string]interface{}{"advertise": true, "vpc_routing_table": []interface{}{map[string]interface{}{"name": "x"}}})
		assert.NotNil(t, err)
	})
}

func TestResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToRoutingTableIdentity(t *testing.T) {
	result, err := vpc.ResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToRoutingTableIdentity(map[string]interface{}{"id": "rt1"})
	assert.Nil(t, err)
	assert.Equal(t, &vpcv1.RoutingTableIdentityByID{ID: core.StringPtr("rt1")}, result)

	result, err = vpc.ResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToRoutingTableIdentity(map[string]interface{}{"crn": "crn:rt"})
	assert.Nil(t, err)
	assert.Equal(t, &vpcv1.RoutingTableIdentityByCRN{CRN: core.StringPtr("crn:rt")}, result)

	result, err = vpc.ResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToRoutingTableIdentity(map[string]interface{}{"href": "https://rt"})
	assert.Nil(t, err)
	assert.Equal(t, &vpcv1.RoutingTableIdentityByHref{Href: core.StringPtr("https://rt")}, result)

	_, err = vpc.ResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToRoutingTableIdentity(map[string]interface{}{"id": "", "crn": "", "href": ""})
	assert.NotNil(t, err)
}

func TestResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToRoutingTableIdentityByID(t *testing.T) {
	_, err := vpc.ResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToRoutingTableIdentityByID(map[string]interface{}{})
	assert.NotNil(t, err)
	result, err := vpc.ResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToRoutingTableIdentityByID(map[string]interface{}{"id": "rt1"})
	assert.Nil(t, err)
	assert.Equal(t, "rt1", *result.ID)
}

func TestResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToRoutingTableIdentityByCRN(t *testing.T) {
	_, err := vpc.ResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToRoutingTableIdentityByCRN(map[string]interface{}{})
	assert.NotNil(t, err)
	result, err := vpc.ResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToRoutingTableIdentityByCRN(map[string]interface{}{"crn": "crn:rt"})
	assert.Nil(t, err)
	assert.Equal(t, "crn:rt", *result.CRN)
}

func TestResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToRoutingTableIdentityByHref(t *testing.T) {
	_, err := vpc.ResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToRoutingTableIdentityByHref(map[string]interface{}{})
	assert.NotNil(t, err)
	result, err := vpc.ResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToRoutingTableIdentityByHref(map[string]interface{}{"href": "https://rt"})
	assert.Nil(t, err)
	assert.Equal(t, "https://rt", *result.Href)
}

func TestResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToDynamicRouteServerPeerGroupPolicyVPCRoutingTablesRoutingTablePatch(t *testing.T) {
	result, err := vpc.ResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToDynamicRouteServerPeerGroupPolicyVPCRoutingTablesRoutingTablePatch(map[string]interface{}{
		"advertise":         false,
		"vpc_routing_table": []interface{}{map[string]interface{}{"href": "https://rt"}},
	})
	assert.Nil(t, err)
	assert.Equal(t, false, *result.Advertise)
	assert.Equal(t, &vpcv1.RoutingTableIdentityByHref{Href: core.StringPtr("https://rt")}, result.VPCRoutingTable)

	_, err = vpc.ResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToDynamicRouteServerPeerGroupPolicyVPCRoutingTablesRoutingTablePatch(map[string]interface{}{"advertise": false})
	assert.NotNil(t, err)
}

func TestResourceIBMIsDynamicRouteServerPeerGroupPolicyValidateTypeFields(t *testing.T) {
	getter := func(m map[string]interface{}) policyGetter { return policyGetter(m) }
	assert.Nil(t, vpc.ResourceIBMIsDynamicRouteServerPeerGroupPolicyValidateTypeFields("custom_routes", getter(map[string]interface{}{"custom_routes": []interface{}{map[string]interface{}{}}})))
	assert.Nil(t, vpc.ResourceIBMIsDynamicRouteServerPeerGroupPolicyValidateTypeFields("vpc_address_prefixes", getter(map[string]interface{}{"custom_routes": []interface{}{}, "route_delete_delay": 0})))
	assert.Nil(t, vpc.ResourceIBMIsDynamicRouteServerPeerGroupPolicyValidateTypeFields("learned_routes", getter(map[string]interface{}{"peer_groups": []interface{}{map[string]interface{}{}}, "excluded_prefixes": []interface{}{map[string]interface{}{}}})))
	assert.Nil(t, vpc.ResourceIBMIsDynamicRouteServerPeerGroupPolicyValidateTypeFields("vpc_routing_tables", getter(map[string]interface{}{"next_hops": []interface{}{map[string]interface{}{}}, "routing_tables": []interface{}{map[string]interface{}{}}, "route_delete_delay": 10})))

	assert.NotNil(t, vpc.ResourceIBMIsDynamicRouteServerPeerGroupPolicyValidateTypeFields("custom_routes", getter(map[string]interface{}{"next_hops": []interface{}{map[string]interface{}{}}})))
	assert.NotNil(t, vpc.ResourceIBMIsDynamicRouteServerPeerGroupPolicyValidateTypeFields("custom_routes", getter(map[string]interface{}{"route_delete_delay": 5})))
	assert.NotNil(t, vpc.ResourceIBMIsDynamicRouteServerPeerGroupPolicyValidateTypeFields("learned_routes", getter(map[string]interface{}{"routing_tables": []interface{}{map[string]interface{}{}}})))
	assert.NotNil(t, vpc.ResourceIBMIsDynamicRouteServerPeerGroupPolicyValidateTypeFields("vpc_address_prefixes", getter(map[string]interface{}{"peer_groups": []interface{}{map[string]interface{}{}}})))
	assert.NotNil(t, vpc.ResourceIBMIsDynamicRouteServerPeerGroupPolicyValidateTypeFields("vpc_routing_tables", getter(map[string]interface{}{"custom_routes": []interface{}{map[string]interface{}{}}})))
	assert.NotNil(t, vpc.ResourceIBMIsDynamicRouteServerPeerGroupPolicyValidateTypeFields("bogus", getter(map[string]interface{}{})))
}

// policyGetter adapts a map to the Get interface used by the type validation.
type policyGetter map[string]interface{}

func (m policyGetter) Get(key string) interface{} { return m[key] }

func TestResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToDynamicRouteServerPeerGroupPolicyPrototype(t *testing.T) {
	t.Run("missing type", func(t *testing.T) {
		_, err := vpc.ResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToDynamicRouteServerPeerGroupPolicyPrototype(map[string]interface{}{})
		assert.NotNil(t, err)
	})
	t.Run("unknown type", func(t *testing.T) {
		_, err := vpc.ResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToDynamicRouteServerPeerGroupPolicyPrototype(map[string]interface{}{"type": "bogus"})
		assert.NotNil(t, err)
	})
	t.Run("custom routes", func(t *testing.T) {
		result, err := vpc.ResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToDynamicRouteServerPeerGroupPolicyPrototype(map[string]interface{}{
			"type":          "custom_routes",
			"name":          "my-policy",
			"state":         "enabled",
			"custom_routes": []interface{}{map[string]interface{}{"destination": "192.168.3.0/24"}},
		})
		assert.Nil(t, err)
		model, ok := result.(*vpcv1.DynamicRouteServerPeerGroupPolicyPrototypeDynamicRouteServerPeerGroupPolicyCustomRoutesPrototype)
		assert.True(t, ok)
		assert.Equal(t, "custom_routes", *model.Type)
		assert.Equal(t, "my-policy", *model.Name)
		assert.Len(t, model.CustomRoutes, 1)
	})
	t.Run("learned routes", func(t *testing.T) {
		result, err := vpc.ResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToDynamicRouteServerPeerGroupPolicyPrototype(map[string]interface{}{
			"type":        "learned_routes",
			"peer_groups": []interface{}{map[string]interface{}{"id": "pg1"}, map[string]interface{}{"href": "https://pg2"}},
		})
		assert.Nil(t, err)
		model, ok := result.(*vpcv1.DynamicRouteServerPeerGroupPolicyPrototypeDynamicRouteServerPeerGroupPolicyLearnedRoutesPrototype)
		assert.True(t, ok)
		assert.Len(t, model.PeerGroups, 2)
		assert.Nil(t, model.ExcludedPrefixes)
		assert.Nil(t, model.Name)
	})
	t.Run("vpc address prefixes", func(t *testing.T) {
		result, err := vpc.ResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToDynamicRouteServerPeerGroupPolicyPrototype(map[string]interface{}{"type": "vpc_address_prefixes", "state": "disabled"})
		assert.Nil(t, err)
		model, ok := result.(*vpcv1.DynamicRouteServerPeerGroupPolicyPrototypeDynamicRouteServerPeerGroupPolicyVPCAddressPrefixesPrototype)
		assert.True(t, ok)
		assert.Equal(t, "disabled", *model.State)
	})
	t.Run("vpc routing tables", func(t *testing.T) {
		result, err := vpc.ResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToDynamicRouteServerPeerGroupPolicyPrototype(map[string]interface{}{
			"type":               "vpc_routing_tables",
			"route_delete_delay": 0,
			"next_hops":          []interface{}{map[string]interface{}{"id": "pg2"}},
			"routing_tables":     []interface{}{map[string]interface{}{"advertise": false, "vpc_routing_table": []interface{}{map[string]interface{}{"id": "rt1"}}}},
		})
		assert.Nil(t, err)
		model, ok := result.(*vpcv1.DynamicRouteServerPeerGroupPolicyPrototypeDynamicRouteServerPeerGroupPolicyVPCRoutingTablesPrototype)
		assert.True(t, ok)
		assert.Equal(t, int64(0), *model.RouteDeleteDelay)
		assert.Len(t, model.NextHops, 1)
		assert.Len(t, model.RoutingTables, 1)
	})
	t.Run("field not allowed for type", func(t *testing.T) {
		_, err := vpc.ResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToDynamicRouteServerPeerGroupPolicyPrototype(map[string]interface{}{
			"type":          "vpc_address_prefixes",
			"custom_routes": []interface{}{map[string]interface{}{"destination": "192.168.3.0/24"}},
		})
		assert.NotNil(t, err)
	})
	t.Run("invalid nested item", func(t *testing.T) {
		_, err := vpc.ResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToDynamicRouteServerPeerGroupPolicyPrototype(map[string]interface{}{
			"type":      "vpc_routing_tables",
			"next_hops": []interface{}{map[string]interface{}{}},
		})
		assert.NotNil(t, err)
	})
}

func TestResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToDynamicRouteServerPeerGroupPolicyPrototypeDynamicRouteServerPeerGroupPolicyCustomRoutesPrototype(t *testing.T) {
	result, err := vpc.ResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToDynamicRouteServerPeerGroupPolicyPrototypeDynamicRouteServerPeerGroupPolicyCustomRoutesPrototype(map[string]interface{}{"type": "custom_routes"})
	assert.Nil(t, err)
	assert.Equal(t, []vpcv1.DynamicRouteServerPeerGroupPolicyCustomRoutePrototype{}, result.CustomRoutes)

	_, err = vpc.ResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToDynamicRouteServerPeerGroupPolicyPrototypeDynamicRouteServerPeerGroupPolicyCustomRoutesPrototype(map[string]interface{}{"custom_routes": []interface{}{map[string]interface{}{}}})
	assert.NotNil(t, err)
}

func TestResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToDynamicRouteServerPeerGroupPolicyPrototypeDynamicRouteServerPeerGroupPolicyLearnedRoutesPrototype(t *testing.T) {
	result, err := vpc.ResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToDynamicRouteServerPeerGroupPolicyPrototypeDynamicRouteServerPeerGroupPolicyLearnedRoutesPrototype(map[string]interface{}{
		"excluded_prefixes": []interface{}{map[string]interface{}{"ge": 0, "le": 0, "prefix": "10.0.0.0/8"}},
		"peer_groups":       []interface{}{map[string]interface{}{"id": "pg1"}},
	})
	assert.Nil(t, err)
	assert.Equal(t, "learned_routes", *result.Type)
	assert.Len(t, result.ExcludedPrefixes, 1)
	assert.Len(t, result.PeerGroups, 1)

	_, err = vpc.ResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToDynamicRouteServerPeerGroupPolicyPrototypeDynamicRouteServerPeerGroupPolicyLearnedRoutesPrototype(map[string]interface{}{
		"excluded_prefixes": []interface{}{map[string]interface{}{"ge": 0}},
	})
	assert.NotNil(t, err)
}

func TestResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToDynamicRouteServerPeerGroupPolicyPrototypeDynamicRouteServerPeerGroupPolicyVPCAddressPrefixesPrototype(t *testing.T) {
	result, err := vpc.ResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToDynamicRouteServerPeerGroupPolicyPrototypeDynamicRouteServerPeerGroupPolicyVPCAddressPrefixesPrototype(map[string]interface{}{"name": "", "state": ""})
	assert.Nil(t, err)
	assert.Equal(t, "vpc_address_prefixes", *result.Type)
	assert.Nil(t, result.Name)
	assert.Nil(t, result.State)
}

func TestResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToDynamicRouteServerPeerGroupPolicyPrototypeDynamicRouteServerPeerGroupPolicyVPCRoutingTablesPrototype(t *testing.T) {
	result, err := vpc.ResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToDynamicRouteServerPeerGroupPolicyPrototypeDynamicRouteServerPeerGroupPolicyVPCRoutingTablesPrototype(map[string]interface{}{
		"name":               "my-policy",
		"route_delete_delay": 30,
		"next_hops":          []interface{}{map[string]interface{}{"id": "pg2"}, map[string]interface{}{"href": "https://pg3"}},
		"routing_tables":     []interface{}{map[string]interface{}{"advertise": true, "vpc_routing_table": []interface{}{map[string]interface{}{"crn": "crn:rt"}}}},
	})
	assert.Nil(t, err)
	assert.Equal(t, "vpc_routing_tables", *result.Type)
	assert.Equal(t, int64(30), *result.RouteDeleteDelay)
	assert.Len(t, result.NextHops, 2)
	assert.Len(t, result.RoutingTables, 1)
	assert.Nil(t, result.ExcludedPrefixes)

	_, err = vpc.ResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToDynamicRouteServerPeerGroupPolicyPrototypeDynamicRouteServerPeerGroupPolicyVPCRoutingTablesPrototype(map[string]interface{}{
		"routing_tables": []interface{}{map[string]interface{}{"advertise": true}},
	})
	assert.NotNil(t, err)
}
