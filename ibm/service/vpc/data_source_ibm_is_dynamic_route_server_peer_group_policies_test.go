// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.114.0-a902401e-20260427-192904
 */

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

func TestAccIBMIsDynamicRouteServerPeerGroupPoliciesDataSourceBasic(t *testing.T) {
	dynamicRouteServerPeerGroupPolicyDynamicRouteServerID := fmt.Sprintf("tf_dynamic_route_server_id_%d", acctest.RandIntRange(10, 100))
	dynamicRouteServerPeerGroupPolicyDynamicRouteServerPeerGroupID := fmt.Sprintf("tf_dynamic_route_server_peer_group_id_%d", acctest.RandIntRange(10, 100))
	dynamicRouteServerPeerGroupPolicyType := "custom_routes"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsDynamicRouteServerPeerGroupPoliciesDataSourceConfigBasic(dynamicRouteServerPeerGroupPolicyDynamicRouteServerID, dynamicRouteServerPeerGroupPolicyDynamicRouteServerPeerGroupID, dynamicRouteServerPeerGroupPolicyType),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_peer_group_policies.is_dynamic_route_server_peer_group_policies_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_peer_group_policies.is_dynamic_route_server_peer_group_policies_instance", "dynamic_route_server_id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_peer_group_policies.is_dynamic_route_server_peer_group_policies_instance", "dynamic_route_server_peer_group_id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_peer_group_policies.is_dynamic_route_server_peer_group_policies_instance", "policies.#"),
					resource.TestCheckResourceAttr("data.ibm_is_dynamic_route_server_peer_group_policies.is_dynamic_route_server_peer_group_policies_instance", "policies.0.type", dynamicRouteServerPeerGroupPolicyType),
				),
			},
		},
	})
}

func TestAccIBMIsDynamicRouteServerPeerGroupPoliciesDataSourceAllArgs(t *testing.T) {
	dynamicRouteServerPeerGroupPolicyDynamicRouteServerID := fmt.Sprintf("tf_dynamic_route_server_id_%d", acctest.RandIntRange(10, 100))
	dynamicRouteServerPeerGroupPolicyDynamicRouteServerPeerGroupID := fmt.Sprintf("tf_dynamic_route_server_peer_group_id_%d", acctest.RandIntRange(10, 100))
	dynamicRouteServerPeerGroupPolicyName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	dynamicRouteServerPeerGroupPolicyState := "disabled"
	dynamicRouteServerPeerGroupPolicyType := "custom_routes"
	dynamicRouteServerPeerGroupPolicyRouteDeleteDelay := fmt.Sprintf("%d", acctest.RandIntRange(0, 60))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsDynamicRouteServerPeerGroupPoliciesDataSourceConfig(dynamicRouteServerPeerGroupPolicyDynamicRouteServerID, dynamicRouteServerPeerGroupPolicyDynamicRouteServerPeerGroupID, dynamicRouteServerPeerGroupPolicyName, dynamicRouteServerPeerGroupPolicyState, dynamicRouteServerPeerGroupPolicyType, dynamicRouteServerPeerGroupPolicyRouteDeleteDelay),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_peer_group_policies.is_dynamic_route_server_peer_group_policies_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_peer_group_policies.is_dynamic_route_server_peer_group_policies_instance", "dynamic_route_server_id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_peer_group_policies.is_dynamic_route_server_peer_group_policies_instance", "dynamic_route_server_peer_group_id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_peer_group_policies.is_dynamic_route_server_peer_group_policies_instance", "sort"),
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_peer_group_policies.is_dynamic_route_server_peer_group_policies_instance", "policies.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_peer_group_policies.is_dynamic_route_server_peer_group_policies_instance", "policies.0.created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_peer_group_policies.is_dynamic_route_server_peer_group_policies_instance", "policies.0.href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_peer_group_policies.is_dynamic_route_server_peer_group_policies_instance", "policies.0.id"),
					resource.TestCheckResourceAttr("data.ibm_is_dynamic_route_server_peer_group_policies.is_dynamic_route_server_peer_group_policies_instance", "policies.0.name", dynamicRouteServerPeerGroupPolicyName),
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_peer_group_policies.is_dynamic_route_server_peer_group_policies_instance", "policies.0.resource_type"),
					resource.TestCheckResourceAttr("data.ibm_is_dynamic_route_server_peer_group_policies.is_dynamic_route_server_peer_group_policies_instance", "policies.0.state", dynamicRouteServerPeerGroupPolicyState),
					resource.TestCheckResourceAttr("data.ibm_is_dynamic_route_server_peer_group_policies.is_dynamic_route_server_peer_group_policies_instance", "policies.0.type", dynamicRouteServerPeerGroupPolicyType),
					resource.TestCheckResourceAttr("data.ibm_is_dynamic_route_server_peer_group_policies.is_dynamic_route_server_peer_group_policies_instance", "policies.0.route_delete_delay", dynamicRouteServerPeerGroupPolicyRouteDeleteDelay),
				),
			},
		},
	})
}

func testAccCheckIBMIsDynamicRouteServerPeerGroupPoliciesDataSourceConfigBasic(dynamicRouteServerPeerGroupPolicyDynamicRouteServerID string, dynamicRouteServerPeerGroupPolicyDynamicRouteServerPeerGroupID string, dynamicRouteServerPeerGroupPolicyType string) string {
	return fmt.Sprintf(`
		resource "ibm_is_dynamic_route_server_peer_group_policy" "is_dynamic_route_server_peer_group_policy_instance" {
			dynamic_route_server_id = "%s"
			dynamic_route_server_peer_group_id = "%s"
			type = "%s"
		}

		data "ibm_is_dynamic_route_server_peer_group_policies" "is_dynamic_route_server_peer_group_policies_instance" {
			dynamic_route_server_id = ibm_is_dynamic_route_server_peer_group_policy.is_dynamic_route_server_peer_group_policy_instance.dynamic_route_server_id
			dynamic_route_server_peer_group_id = ibm_is_dynamic_route_server_peer_group_policy.is_dynamic_route_server_peer_group_policy_instance.dynamic_route_server_peer_group_id
			sort = "name"
		}
	`, dynamicRouteServerPeerGroupPolicyDynamicRouteServerID, dynamicRouteServerPeerGroupPolicyDynamicRouteServerPeerGroupID, dynamicRouteServerPeerGroupPolicyType)
}

func testAccCheckIBMIsDynamicRouteServerPeerGroupPoliciesDataSourceConfig(dynamicRouteServerPeerGroupPolicyDynamicRouteServerID string, dynamicRouteServerPeerGroupPolicyDynamicRouteServerPeerGroupID string, dynamicRouteServerPeerGroupPolicyName string, dynamicRouteServerPeerGroupPolicyState string, dynamicRouteServerPeerGroupPolicyType string, dynamicRouteServerPeerGroupPolicyRouteDeleteDelay string) string {
	return fmt.Sprintf(`
		resource "ibm_is_dynamic_route_server_peer_group_policy" "is_dynamic_route_server_peer_group_policy_instance" {
			dynamic_route_server_id = "%s"
			dynamic_route_server_peer_group_id = "%s"
			name = "%s"
			state = "%s"
			type = "%s"
			custom_routes {
				destination = "192.168.3.0/24"
			}
			excluded_prefixes {
				ge = 0
				le = 32
				prefix = "prefix"
			}
			peer_groups {
				deleted {
					more_info = "https://cloud.ibm.com/apidocs/vpc#deleted-resources"
				}
				href = "https://us-south.iaas.cloud.ibm.com/v1/dynamic_route_servers/r006-d7cc5196-9864-48c4-82d8-3f30da41fcc5/peer_groups/r006-8228af60-189d-11ed-861d-0242ac120002"
				id = "r006-8228af60-189d-11ed-861d-0242ac120004"
				name = "my-dynamic-route-server-peer-group"
				resource_type = "dynamic_route_server_peer_group"
			}
			next_hops {
				deleted {
					more_info = "https://cloud.ibm.com/apidocs/vpc#deleted-resources"
				}
				href = "https://us-south.iaas.cloud.ibm.com/v1/dynamic_route_servers/r006-d7cc5196-9864-48c4-82d8-3f30da41fcc5/peer_groups/r006-8228af60-189d-11ed-861d-0242ac120002"
				id = "r006-8228af60-189d-11ed-861d-0242ac120004"
				name = "my-dynamic-route-server-peer-group"
				resource_type = "dynamic_route_server_peer_group"
			}
			route_delete_delay = %s
			routing_tables {
				advertise = true
				vpc_routing_table {
					crn = "crn:v1:bluemix:public:is:us-south:a/aa2432b1fa4d4ace891e9b80fc104e34::vpc-routing-table:r006-4727d842-f94f-4a2d-824a-9bc9b02c523b/r006-6885e83f-03b2-4603-8a86-db2a0f55c840"
					deleted {
						more_info = "https://cloud.ibm.com/apidocs/vpc#deleted-resources"
					}
					href = "https://us-south.iaas.cloud.ibm.com/v1/vpcs/r006-4727d842-f94f-4a2d-824a-9bc9b02c523b/routing_tables/r006-6885e83f-03b2-4603-8a86-db2a0f55c840"
					id = "r006-6885e83f-03b2-4603-8a86-db2a0f55c840"
					name = "my-routing-table-1"
					resource_type = "routing_table"
				}
			}
		}

		data "ibm_is_dynamic_route_server_peer_group_policies" "is_dynamic_route_server_peer_group_policies_instance" {
			dynamic_route_server_id = ibm_is_dynamic_route_server_peer_group_policy.is_dynamic_route_server_peer_group_policy_instance.dynamic_route_server_id
			dynamic_route_server_peer_group_id = ibm_is_dynamic_route_server_peer_group_policy.is_dynamic_route_server_peer_group_policy_instance.dynamic_route_server_peer_group_id
			sort = "name"
		}
	`, dynamicRouteServerPeerGroupPolicyDynamicRouteServerID, dynamicRouteServerPeerGroupPolicyDynamicRouteServerPeerGroupID, dynamicRouteServerPeerGroupPolicyName, dynamicRouteServerPeerGroupPolicyState, dynamicRouteServerPeerGroupPolicyType, dynamicRouteServerPeerGroupPolicyRouteDeleteDelay)
}

func TestDataSourceIBMIsDynamicRouteServerPeerGroupPoliciesDynamicRouteServerPeerGroupPolicyToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		dynamicRouteServerPeerGroupPolicyCustomRouteModel := make(map[string]interface{})
		dynamicRouteServerPeerGroupPolicyCustomRouteModel["destination"] = "192.168.3.0/24"

		model := make(map[string]interface{})
		model["created_at"] = "2026-01-02T03:04:05.006Z"
		model["href"] = "https://us-south.iaas.cloud.ibm.com/v1/dynamic_route_servers/r006-d7cc5196-9864-48c4-82d8-3f30da41fcc5/peer_groups/r006-d7cc5196-9864-48c4-82d8-3f30da41fcc5/policies/r006-d7cc5196-9864-48c4-82d8-3f30da41fcc5"
		model["id"] = "r006-d7cc5196-9864-48c4-82d8-3f30da41fcc5"
		model["name"] = "my-dynamic-route-server-peer-group-policy"
		model["resource_type"] = "dynamic_route_server_peer_group_policy"
		model["state"] = "disabled"
		model["type"] = "custom_routes"
		model["custom_routes"] = []map[string]interface{}{dynamicRouteServerPeerGroupPolicyCustomRouteModel}
		// model["excluded_prefixes"] = []map[string]interface{}{dynamicRouteServerPeerGroupPrefixWatcherExcludedPrefixPrototypeModel}
		// model["peer_groups"] = []map[string]interface{}{dynamicRouteServerPeerGroupReferenceModel}
		// model["next_hops"] = []map[string]interface{}{dynamicRouteServerPeerGroupPolicyNextHopModel}
		model["route_delete_delay"] = int(0)
		// model["routing_tables"] = []map[string]interface{}{dynamicRouteServerPeerGroupPolicyVPCRoutingTablesRoutingTableModel}

		assert.Equal(t, result, model)
	}

	dynamicRouteServerPeerGroupPolicyCustomRouteModel := new(vpcv1.DynamicRouteServerPeerGroupPolicyCustomRoute)
	dynamicRouteServerPeerGroupPolicyCustomRouteModel.Destination = core.StringPtr("192.168.3.0/24")

	model := new(vpcv1.DynamicRouteServerPeerGroupPolicy)
	model.CreatedAt = CreateMockDateTime("2026-01-02T03:04:05.006Z")
	model.Href = core.StringPtr("https://us-south.iaas.cloud.ibm.com/v1/dynamic_route_servers/r006-d7cc5196-9864-48c4-82d8-3f30da41fcc5/peer_groups/r006-d7cc5196-9864-48c4-82d8-3f30da41fcc5/policies/r006-d7cc5196-9864-48c4-82d8-3f30da41fcc5")
	model.ID = core.StringPtr("r006-d7cc5196-9864-48c4-82d8-3f30da41fcc5")
	model.Name = core.StringPtr("my-dynamic-route-server-peer-group-policy")
	model.ResourceType = core.StringPtr("dynamic_route_server_peer_group_policy")
	model.State = core.StringPtr("disabled")
	model.Type = core.StringPtr("custom_routes")
	model.CustomRoutes = []vpcv1.DynamicRouteServerPeerGroupPolicyCustomRoute{*dynamicRouteServerPeerGroupPolicyCustomRouteModel}
	// model.ExcludedPrefixes = []vpcv1.DynamicRouteServerPeerGroupPrefixWatcherExcludedPrefixPrototype{*dynamicRouteServerPeerGroupPrefixWatcherExcludedPrefixPrototypeModel}
	// model.PeerGroups = []vpcv1.DynamicRouteServerPeerGroupReference{*dynamicRouteServerPeerGroupReferenceModel}
	// model.NextHops = []vpcv1.DynamicRouteServerPeerGroupPolicyNextHopIntf{dynamicRouteServerPeerGroupPolicyNextHopModel}
	model.RouteDeleteDelay = core.Int64Ptr(int64(0))
	// model.RoutingTables = []vpcv1.DynamicRouteServerPeerGroupPolicyVPCRoutingTablesRoutingTable{*dynamicRouteServerPeerGroupPolicyVPCRoutingTablesRoutingTableModel}

	result, err := vpc.DataSourceIBMIsDynamicRouteServerPeerGroupPoliciesDynamicRouteServerPeerGroupPolicyToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMIsDynamicRouteServerPeerGroupPoliciesDynamicRouteServerPeerGroupPolicyCustomRouteToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["destination"] = "192.168.3.0/24"

		assert.Equal(t, result, model)
	}

	model := new(vpcv1.DynamicRouteServerPeerGroupPolicyCustomRoute)
	model.Destination = core.StringPtr("192.168.3.0/24")

	result, err := vpc.DataSourceIBMIsDynamicRouteServerPeerGroupPoliciesDynamicRouteServerPeerGroupPolicyCustomRouteToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMIsDynamicRouteServerPeerGroupPoliciesDynamicRouteServerPeerGroupPrefixWatcherExcludedPrefixPrototypeToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["ge"] = int(0)
		model["le"] = int(32)
		model["prefix"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(vpcv1.DynamicRouteServerPeerGroupPrefixWatcherExcludedPrefixPrototype)
	model.Ge = core.Int64Ptr(int64(0))
	model.Le = core.Int64Ptr(int64(32))
	model.Prefix = core.StringPtr("testString")

	result, err := vpc.DataSourceIBMIsDynamicRouteServerPeerGroupPoliciesDynamicRouteServerPeerGroupPrefixWatcherExcludedPrefixPrototypeToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMIsDynamicRouteServerPeerGroupPoliciesDynamicRouteServerPeerGroupReferenceToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		deletedModel := make(map[string]interface{})
		deletedModel["more_info"] = "https://cloud.ibm.com/apidocs/vpc#deleted-resources"

		model := make(map[string]interface{})
		model["deleted"] = []map[string]interface{}{deletedModel}
		model["href"] = "https://us-south.iaas.cloud.ibm.com/v1/dynamic_route_servers/r006-d7cc5196-9864-48c4-82d8-3f30da41fcc5/peer_groups/r006-8228af60-189d-11ed-861d-0242ac120002"
		model["id"] = "r006-8228af60-189d-11ed-861d-0242ac120004"
		model["name"] = "my-dynamic-route-server-peer-group"
		model["resource_type"] = "dynamic_route_server_peer_group"

		assert.Equal(t, result, model)
	}

	deletedModel := new(vpcv1.Deleted)
	deletedModel.MoreInfo = core.StringPtr("https://cloud.ibm.com/apidocs/vpc#deleted-resources")

	model := new(vpcv1.DynamicRouteServerPeerGroupReference)
	model.Deleted = deletedModel
	model.Href = core.StringPtr("https://us-south.iaas.cloud.ibm.com/v1/dynamic_route_servers/r006-d7cc5196-9864-48c4-82d8-3f30da41fcc5/peer_groups/r006-8228af60-189d-11ed-861d-0242ac120002")
	model.ID = core.StringPtr("r006-8228af60-189d-11ed-861d-0242ac120004")
	model.Name = core.StringPtr("my-dynamic-route-server-peer-group")
	model.ResourceType = core.StringPtr("dynamic_route_server_peer_group")

	result, err := vpc.DataSourceIBMIsDynamicRouteServerPeerGroupPoliciesDynamicRouteServerPeerGroupReferenceToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMIsDynamicRouteServerPeerGroupPoliciesDeletedToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["more_info"] = "https://cloud.ibm.com/apidocs/vpc#deleted-resources"

		assert.Equal(t, result, model)
	}

	model := new(vpcv1.Deleted)
	model.MoreInfo = core.StringPtr("https://cloud.ibm.com/apidocs/vpc#deleted-resources")

	result, err := vpc.DataSourceIBMIsDynamicRouteServerPeerGroupPoliciesDeletedToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMIsDynamicRouteServerPeerGroupPoliciesDynamicRouteServerPeerGroupPolicyNextHopToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		deletedModel := make(map[string]interface{})
		deletedModel["more_info"] = "https://cloud.ibm.com/apidocs/vpc#deleted-resources"

		model := make(map[string]interface{})
		model["deleted"] = []map[string]interface{}{deletedModel}
		model["href"] = "https://us-south.iaas.cloud.ibm.com/v1/dynamic_route_servers/r006-d7cc5196-9864-48c4-82d8-3f30da41fcc5/peer_groups/r006-8228af60-189d-11ed-861d-0242ac120002"
		model["id"] = "r006-8228af60-189d-11ed-861d-0242ac120004"
		model["name"] = "my-dynamic-route-server-peer-group"
		model["resource_type"] = "dynamic_route_server_peer_group"

		assert.Equal(t, result, model)
	}

	deletedModel := new(vpcv1.Deleted)
	deletedModel.MoreInfo = core.StringPtr("https://cloud.ibm.com/apidocs/vpc#deleted-resources")

	model := new(vpcv1.DynamicRouteServerPeerGroupPolicyNextHop)
	model.Deleted = deletedModel
	model.Href = core.StringPtr("https://us-south.iaas.cloud.ibm.com/v1/dynamic_route_servers/r006-d7cc5196-9864-48c4-82d8-3f30da41fcc5/peer_groups/r006-8228af60-189d-11ed-861d-0242ac120002")
	model.ID = core.StringPtr("r006-8228af60-189d-11ed-861d-0242ac120004")
	model.Name = core.StringPtr("my-dynamic-route-server-peer-group")
	model.ResourceType = core.StringPtr("dynamic_route_server_peer_group")

	result, err := vpc.DataSourceIBMIsDynamicRouteServerPeerGroupPoliciesDynamicRouteServerPeerGroupPolicyNextHopToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMIsDynamicRouteServerPeerGroupPoliciesDynamicRouteServerPeerGroupPolicyNextHopDynamicRouteServerPeerGroupReferenceToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		deletedModel := make(map[string]interface{})
		deletedModel["more_info"] = "https://cloud.ibm.com/apidocs/vpc#deleted-resources"

		model := make(map[string]interface{})
		model["deleted"] = []map[string]interface{}{deletedModel}
		model["href"] = "https://us-south.iaas.cloud.ibm.com/v1/dynamic_route_servers/r006-d7cc5196-9864-48c4-82d8-3f30da41fcc5/peer_groups/r006-8228af60-189d-11ed-861d-0242ac120002"
		model["id"] = "r006-8228af60-189d-11ed-861d-0242ac120004"
		model["name"] = "my-dynamic-route-server-peer-group"
		model["resource_type"] = "dynamic_route_server_peer_group"

		assert.Equal(t, result, model)
	}

	deletedModel := new(vpcv1.Deleted)
	deletedModel.MoreInfo = core.StringPtr("https://cloud.ibm.com/apidocs/vpc#deleted-resources")

	model := new(vpcv1.DynamicRouteServerPeerGroupPolicyNextHopDynamicRouteServerPeerGroupReference)
	model.Deleted = deletedModel
	model.Href = core.StringPtr("https://us-south.iaas.cloud.ibm.com/v1/dynamic_route_servers/r006-d7cc5196-9864-48c4-82d8-3f30da41fcc5/peer_groups/r006-8228af60-189d-11ed-861d-0242ac120002")
	model.ID = core.StringPtr("r006-8228af60-189d-11ed-861d-0242ac120004")
	model.Name = core.StringPtr("my-dynamic-route-server-peer-group")
	model.ResourceType = core.StringPtr("dynamic_route_server_peer_group")

	result, err := vpc.DataSourceIBMIsDynamicRouteServerPeerGroupPoliciesDynamicRouteServerPeerGroupPolicyNextHopDynamicRouteServerPeerGroupReferenceToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMIsDynamicRouteServerPeerGroupPoliciesDynamicRouteServerPeerGroupPolicyVPCRoutingTablesRoutingTableToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		deletedModel := make(map[string]interface{})
		deletedModel["more_info"] = "https://cloud.ibm.com/apidocs/vpc#deleted-resources"

		routingTableReferenceModel := make(map[string]interface{})
		routingTableReferenceModel["crn"] = "crn:v1:bluemix:public:is:us-south:a/aa2432b1fa4d4ace891e9b80fc104e34::vpc-routing-table:r006-4727d842-f94f-4a2d-824a-9bc9b02c523b/r006-6885e83f-03b2-4603-8a86-db2a0f55c840"
		routingTableReferenceModel["deleted"] = []map[string]interface{}{deletedModel}
		routingTableReferenceModel["href"] = "https://us-south.iaas.cloud.ibm.com/v1/vpcs/r006-4727d842-f94f-4a2d-824a-9bc9b02c523b/routing_tables/r006-6885e83f-03b2-4603-8a86-db2a0f55c840"
		routingTableReferenceModel["id"] = "r006-6885e83f-03b2-4603-8a86-db2a0f55c840"
		routingTableReferenceModel["name"] = "my-routing-table-1"
		routingTableReferenceModel["resource_type"] = "routing_table"

		model := make(map[string]interface{})
		model["advertise"] = true
		model["vpc_routing_table"] = []map[string]interface{}{routingTableReferenceModel}

		assert.Equal(t, result, model)
	}

	deletedModel := new(vpcv1.Deleted)
	deletedModel.MoreInfo = core.StringPtr("https://cloud.ibm.com/apidocs/vpc#deleted-resources")

	routingTableReferenceModel := new(vpcv1.RoutingTableReference)
	routingTableReferenceModel.CRN = core.StringPtr("crn:v1:bluemix:public:is:us-south:a/aa2432b1fa4d4ace891e9b80fc104e34::vpc-routing-table:r006-4727d842-f94f-4a2d-824a-9bc9b02c523b/r006-6885e83f-03b2-4603-8a86-db2a0f55c840")
	routingTableReferenceModel.Deleted = deletedModel
	routingTableReferenceModel.Href = core.StringPtr("https://us-south.iaas.cloud.ibm.com/v1/vpcs/r006-4727d842-f94f-4a2d-824a-9bc9b02c523b/routing_tables/r006-6885e83f-03b2-4603-8a86-db2a0f55c840")
	routingTableReferenceModel.ID = core.StringPtr("r006-6885e83f-03b2-4603-8a86-db2a0f55c840")
	routingTableReferenceModel.Name = core.StringPtr("my-routing-table-1")
	routingTableReferenceModel.ResourceType = core.StringPtr("routing_table")

	model := new(vpcv1.DynamicRouteServerPeerGroupPolicyVPCRoutingTablesRoutingTable)
	model.Advertise = core.BoolPtr(true)
	model.VPCRoutingTable = routingTableReferenceModel

	result, err := vpc.DataSourceIBMIsDynamicRouteServerPeerGroupPoliciesDynamicRouteServerPeerGroupPolicyVPCRoutingTablesRoutingTableToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMIsDynamicRouteServerPeerGroupPoliciesRoutingTableReferenceToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		deletedModel := make(map[string]interface{})
		deletedModel["more_info"] = "https://cloud.ibm.com/apidocs/vpc#deleted-resources"

		model := make(map[string]interface{})
		model["crn"] = "crn:v1:bluemix:public:is:us-south:a/aa2432b1fa4d4ace891e9b80fc104e34::vpc-routing-table:r006-4727d842-f94f-4a2d-824a-9bc9b02c523b/r006-6885e83f-03b2-4603-8a86-db2a0f55c840"
		model["deleted"] = []map[string]interface{}{deletedModel}
		model["href"] = "https://us-south.iaas.cloud.ibm.com/v1/vpcs/r006-4727d842-f94f-4a2d-824a-9bc9b02c523b/routing_tables/r006-6885e83f-03b2-4603-8a86-db2a0f55c840"
		model["id"] = "r006-6885e83f-03b2-4603-8a86-db2a0f55c840"
		model["name"] = "my-routing-table-1"
		model["resource_type"] = "routing_table"

		assert.Equal(t, result, model)
	}

	deletedModel := new(vpcv1.Deleted)
	deletedModel.MoreInfo = core.StringPtr("https://cloud.ibm.com/apidocs/vpc#deleted-resources")

	model := new(vpcv1.RoutingTableReference)
	model.CRN = core.StringPtr("crn:v1:bluemix:public:is:us-south:a/aa2432b1fa4d4ace891e9b80fc104e34::vpc-routing-table:r006-4727d842-f94f-4a2d-824a-9bc9b02c523b/r006-6885e83f-03b2-4603-8a86-db2a0f55c840")
	model.Deleted = deletedModel
	model.Href = core.StringPtr("https://us-south.iaas.cloud.ibm.com/v1/vpcs/r006-4727d842-f94f-4a2d-824a-9bc9b02c523b/routing_tables/r006-6885e83f-03b2-4603-8a86-db2a0f55c840")
	model.ID = core.StringPtr("r006-6885e83f-03b2-4603-8a86-db2a0f55c840")
	model.Name = core.StringPtr("my-routing-table-1")
	model.ResourceType = core.StringPtr("routing_table")

	result, err := vpc.DataSourceIBMIsDynamicRouteServerPeerGroupPoliciesRoutingTableReferenceToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMIsDynamicRouteServerPeerGroupPoliciesDynamicRouteServerPeerGroupPolicyCustomRoutesToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		dynamicRouteServerPeerGroupPolicyCustomRouteModel := make(map[string]interface{})
		dynamicRouteServerPeerGroupPolicyCustomRouteModel["destination"] = "192.168.3.0/24"

		model := make(map[string]interface{})
		model["created_at"] = "2026-01-02T03:04:05.006Z"
		model["href"] = "https://us-south.iaas.cloud.ibm.com/v1/dynamic_route_servers/r006-d7cc5196-9864-48c4-82d8-3f30da41fcc5/peer_groups/r006-d7cc5196-9864-48c4-82d8-3f30da41fcc5/policies/r006-d7cc5196-9864-48c4-82d8-3f30da41fcc5"
		model["id"] = "r006-d7cc5196-9864-48c4-82d8-3f30da41fcc5"
		model["name"] = "my-dynamic-route-server-peer-group-policy"
		model["resource_type"] = "dynamic_route_server_peer_group_policy"
		model["state"] = "disabled"
		model["custom_routes"] = []map[string]interface{}{dynamicRouteServerPeerGroupPolicyCustomRouteModel}
		model["type"] = "custom_routes"

		assert.Equal(t, result, model)
	}

	dynamicRouteServerPeerGroupPolicyCustomRouteModel := new(vpcv1.DynamicRouteServerPeerGroupPolicyCustomRoute)
	dynamicRouteServerPeerGroupPolicyCustomRouteModel.Destination = core.StringPtr("192.168.3.0/24")

	model := new(vpcv1.DynamicRouteServerPeerGroupPolicyCustomRoutes)
	model.CreatedAt = CreateMockDateTime("2026-01-02T03:04:05.006Z")
	model.Href = core.StringPtr("https://us-south.iaas.cloud.ibm.com/v1/dynamic_route_servers/r006-d7cc5196-9864-48c4-82d8-3f30da41fcc5/peer_groups/r006-d7cc5196-9864-48c4-82d8-3f30da41fcc5/policies/r006-d7cc5196-9864-48c4-82d8-3f30da41fcc5")
	model.ID = core.StringPtr("r006-d7cc5196-9864-48c4-82d8-3f30da41fcc5")
	model.Name = core.StringPtr("my-dynamic-route-server-peer-group-policy")
	model.ResourceType = core.StringPtr("dynamic_route_server_peer_group_policy")
	model.State = core.StringPtr("disabled")
	model.CustomRoutes = []vpcv1.DynamicRouteServerPeerGroupPolicyCustomRoute{*dynamicRouteServerPeerGroupPolicyCustomRouteModel}
	model.Type = core.StringPtr("custom_routes")

	result, err := vpc.DataSourceIBMIsDynamicRouteServerPeerGroupPoliciesDynamicRouteServerPeerGroupPolicyCustomRoutesToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMIsDynamicRouteServerPeerGroupPoliciesDynamicRouteServerPeerGroupPolicyLearnedRoutesToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		dynamicRouteServerPeerGroupPrefixWatcherExcludedPrefixPrototypeModel := make(map[string]interface{})
		dynamicRouteServerPeerGroupPrefixWatcherExcludedPrefixPrototypeModel["ge"] = int(0)
		dynamicRouteServerPeerGroupPrefixWatcherExcludedPrefixPrototypeModel["le"] = int(32)
		dynamicRouteServerPeerGroupPrefixWatcherExcludedPrefixPrototypeModel["prefix"] = "testString"

		deletedModel := make(map[string]interface{})
		deletedModel["more_info"] = "https://cloud.ibm.com/apidocs/vpc#deleted-resources"

		dynamicRouteServerPeerGroupReferenceModel := make(map[string]interface{})
		dynamicRouteServerPeerGroupReferenceModel["deleted"] = []map[string]interface{}{deletedModel}
		dynamicRouteServerPeerGroupReferenceModel["href"] = "https://us-south.iaas.cloud.ibm.com/v1/dynamic_route_servers/r006-d7cc5196-9864-48c4-82d8-3f30da41fcc5/peer_groups/r006-8228af60-189d-11ed-861d-0242ac120002"
		dynamicRouteServerPeerGroupReferenceModel["id"] = "r006-8228af60-189d-11ed-861d-0242ac120004"
		dynamicRouteServerPeerGroupReferenceModel["name"] = "my-dynamic-route-server-peer-group"
		dynamicRouteServerPeerGroupReferenceModel["resource_type"] = "dynamic_route_server_peer_group"

		model := make(map[string]interface{})
		model["created_at"] = "2026-01-02T03:04:05.006Z"
		model["href"] = "https://us-south.iaas.cloud.ibm.com/v1/dynamic_route_servers/r006-d7cc5196-9864-48c4-82d8-3f30da41fcc5/peer_groups/r006-d7cc5196-9864-48c4-82d8-3f30da41fcc5/policies/r006-d7cc5196-9864-48c4-82d8-3f30da41fcc5"
		model["id"] = "r006-d7cc5196-9864-48c4-82d8-3f30da41fcc5"
		model["name"] = "my-dynamic-route-server-peer-group-policy"
		model["resource_type"] = "dynamic_route_server_peer_group_policy"
		model["state"] = "disabled"
		model["excluded_prefixes"] = []map[string]interface{}{dynamicRouteServerPeerGroupPrefixWatcherExcludedPrefixPrototypeModel}
		model["peer_groups"] = []map[string]interface{}{dynamicRouteServerPeerGroupReferenceModel}
		model["type"] = "learned_routes"

		assert.Equal(t, result, model)
	}

	dynamicRouteServerPeerGroupPrefixWatcherExcludedPrefixPrototypeModel := new(vpcv1.DynamicRouteServerPeerGroupPrefixWatcherExcludedPrefixPrototype)
	dynamicRouteServerPeerGroupPrefixWatcherExcludedPrefixPrototypeModel.Ge = core.Int64Ptr(int64(0))
	dynamicRouteServerPeerGroupPrefixWatcherExcludedPrefixPrototypeModel.Le = core.Int64Ptr(int64(32))
	dynamicRouteServerPeerGroupPrefixWatcherExcludedPrefixPrototypeModel.Prefix = core.StringPtr("testString")

	deletedModel := new(vpcv1.Deleted)
	deletedModel.MoreInfo = core.StringPtr("https://cloud.ibm.com/apidocs/vpc#deleted-resources")

	dynamicRouteServerPeerGroupReferenceModel := new(vpcv1.DynamicRouteServerPeerGroupReference)
	dynamicRouteServerPeerGroupReferenceModel.Deleted = deletedModel
	dynamicRouteServerPeerGroupReferenceModel.Href = core.StringPtr("https://us-south.iaas.cloud.ibm.com/v1/dynamic_route_servers/r006-d7cc5196-9864-48c4-82d8-3f30da41fcc5/peer_groups/r006-8228af60-189d-11ed-861d-0242ac120002")
	dynamicRouteServerPeerGroupReferenceModel.ID = core.StringPtr("r006-8228af60-189d-11ed-861d-0242ac120004")
	dynamicRouteServerPeerGroupReferenceModel.Name = core.StringPtr("my-dynamic-route-server-peer-group")
	dynamicRouteServerPeerGroupReferenceModel.ResourceType = core.StringPtr("dynamic_route_server_peer_group")

	model := new(vpcv1.DynamicRouteServerPeerGroupPolicyLearnedRoutes)
	model.CreatedAt = CreateMockDateTime("2026-01-02T03:04:05.006Z")
	model.Href = core.StringPtr("https://us-south.iaas.cloud.ibm.com/v1/dynamic_route_servers/r006-d7cc5196-9864-48c4-82d8-3f30da41fcc5/peer_groups/r006-d7cc5196-9864-48c4-82d8-3f30da41fcc5/policies/r006-d7cc5196-9864-48c4-82d8-3f30da41fcc5")
	model.ID = core.StringPtr("r006-d7cc5196-9864-48c4-82d8-3f30da41fcc5")
	model.Name = core.StringPtr("my-dynamic-route-server-peer-group-policy")
	model.ResourceType = core.StringPtr("dynamic_route_server_peer_group_policy")
	model.State = core.StringPtr("disabled")
	model.ExcludedPrefixes = []vpcv1.DynamicRouteServerPeerGroupPrefixWatcherExcludedPrefixPrototype{*dynamicRouteServerPeerGroupPrefixWatcherExcludedPrefixPrototypeModel}
	model.PeerGroups = []vpcv1.DynamicRouteServerPeerGroupReference{*dynamicRouteServerPeerGroupReferenceModel}
	model.Type = core.StringPtr("learned_routes")

	result, err := vpc.DataSourceIBMIsDynamicRouteServerPeerGroupPoliciesDynamicRouteServerPeerGroupPolicyLearnedRoutesToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMIsDynamicRouteServerPeerGroupPoliciesDynamicRouteServerPeerGroupPolicyVPCAddressPrefixesToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["created_at"] = "2026-01-02T03:04:05.006Z"
		model["href"] = "https://us-south.iaas.cloud.ibm.com/v1/dynamic_route_servers/r006-d7cc5196-9864-48c4-82d8-3f30da41fcc5/peer_groups/r006-d7cc5196-9864-48c4-82d8-3f30da41fcc5/policies/r006-d7cc5196-9864-48c4-82d8-3f30da41fcc5"
		model["id"] = "r006-d7cc5196-9864-48c4-82d8-3f30da41fcc5"
		model["name"] = "my-dynamic-route-server-peer-group-policy"
		model["resource_type"] = "dynamic_route_server_peer_group_policy"
		model["state"] = "disabled"
		model["type"] = "vpc_address_prefixes"

		assert.Equal(t, result, model)
	}

	model := new(vpcv1.DynamicRouteServerPeerGroupPolicyVPCAddressPrefixes)
	model.CreatedAt = CreateMockDateTime("2026-01-02T03:04:05.006Z")
	model.Href = core.StringPtr("https://us-south.iaas.cloud.ibm.com/v1/dynamic_route_servers/r006-d7cc5196-9864-48c4-82d8-3f30da41fcc5/peer_groups/r006-d7cc5196-9864-48c4-82d8-3f30da41fcc5/policies/r006-d7cc5196-9864-48c4-82d8-3f30da41fcc5")
	model.ID = core.StringPtr("r006-d7cc5196-9864-48c4-82d8-3f30da41fcc5")
	model.Name = core.StringPtr("my-dynamic-route-server-peer-group-policy")
	model.ResourceType = core.StringPtr("dynamic_route_server_peer_group_policy")
	model.State = core.StringPtr("disabled")
	model.Type = core.StringPtr("vpc_address_prefixes")

	result, err := vpc.DataSourceIBMIsDynamicRouteServerPeerGroupPoliciesDynamicRouteServerPeerGroupPolicyVPCAddressPrefixesToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMIsDynamicRouteServerPeerGroupPoliciesDynamicRouteServerPeerGroupPolicyVPCRoutingTablesToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		dynamicRouteServerPeerGroupPrefixWatcherExcludedPrefixPrototypeModel := make(map[string]interface{})
		dynamicRouteServerPeerGroupPrefixWatcherExcludedPrefixPrototypeModel["ge"] = int(0)
		dynamicRouteServerPeerGroupPrefixWatcherExcludedPrefixPrototypeModel["le"] = int(32)
		dynamicRouteServerPeerGroupPrefixWatcherExcludedPrefixPrototypeModel["prefix"] = "testString"

		deletedModel := make(map[string]interface{})
		deletedModel["more_info"] = "https://cloud.ibm.com/apidocs/vpc#deleted-resources"

		dynamicRouteServerPeerGroupPolicyNextHopModel := make(map[string]interface{})
		dynamicRouteServerPeerGroupPolicyNextHopModel["deleted"] = []map[string]interface{}{deletedModel}
		dynamicRouteServerPeerGroupPolicyNextHopModel["href"] = "https://us-south.iaas.cloud.ibm.com/v1/dynamic_route_servers/r006-d7cc5196-9864-48c4-82d8-3f30da41fcc5/peer_groups/r006-8228af60-189d-11ed-861d-0242ac120002"
		dynamicRouteServerPeerGroupPolicyNextHopModel["id"] = "r006-8228af60-189d-11ed-861d-0242ac120004"
		dynamicRouteServerPeerGroupPolicyNextHopModel["name"] = "my-dynamic-route-server-peer-group"
		dynamicRouteServerPeerGroupPolicyNextHopModel["resource_type"] = "dynamic_route_server_peer_group"

		routingTableReferenceModel := make(map[string]interface{})
		routingTableReferenceModel["crn"] = "crn:v1:bluemix:public:is:us-south:a/aa2432b1fa4d4ace891e9b80fc104e34::vpc-routing-table:r006-4727d842-f94f-4a2d-824a-9bc9b02c523b/r006-6885e83f-03b2-4603-8a86-db2a0f55c840"
		routingTableReferenceModel["deleted"] = []map[string]interface{}{deletedModel}
		routingTableReferenceModel["href"] = "https://us-south.iaas.cloud.ibm.com/v1/vpcs/r006-4727d842-f94f-4a2d-824a-9bc9b02c523b/routing_tables/r006-6885e83f-03b2-4603-8a86-db2a0f55c840"
		routingTableReferenceModel["id"] = "r006-6885e83f-03b2-4603-8a86-db2a0f55c840"
		routingTableReferenceModel["name"] = "my-routing-table-1"
		routingTableReferenceModel["resource_type"] = "routing_table"

		dynamicRouteServerPeerGroupPolicyVPCRoutingTablesRoutingTableModel := make(map[string]interface{})
		dynamicRouteServerPeerGroupPolicyVPCRoutingTablesRoutingTableModel["advertise"] = true
		dynamicRouteServerPeerGroupPolicyVPCRoutingTablesRoutingTableModel["vpc_routing_table"] = []map[string]interface{}{routingTableReferenceModel}

		model := make(map[string]interface{})
		model["created_at"] = "2026-01-02T03:04:05.006Z"
		model["href"] = "https://us-south.iaas.cloud.ibm.com/v1/dynamic_route_servers/r006-d7cc5196-9864-48c4-82d8-3f30da41fcc5/peer_groups/r006-d7cc5196-9864-48c4-82d8-3f30da41fcc5/policies/r006-d7cc5196-9864-48c4-82d8-3f30da41fcc5"
		model["id"] = "r006-d7cc5196-9864-48c4-82d8-3f30da41fcc5"
		model["name"] = "my-dynamic-route-server-peer-group-policy"
		model["resource_type"] = "dynamic_route_server_peer_group_policy"
		model["state"] = "disabled"
		model["excluded_prefixes"] = []map[string]interface{}{dynamicRouteServerPeerGroupPrefixWatcherExcludedPrefixPrototypeModel}
		model["next_hops"] = []map[string]interface{}{dynamicRouteServerPeerGroupPolicyNextHopModel}
		model["route_delete_delay"] = int(0)
		model["routing_tables"] = []map[string]interface{}{dynamicRouteServerPeerGroupPolicyVPCRoutingTablesRoutingTableModel}
		model["type"] = "vpc_routing_tables"

		assert.Equal(t, result, model)
	}

	dynamicRouteServerPeerGroupPrefixWatcherExcludedPrefixPrototypeModel := new(vpcv1.DynamicRouteServerPeerGroupPrefixWatcherExcludedPrefixPrototype)
	dynamicRouteServerPeerGroupPrefixWatcherExcludedPrefixPrototypeModel.Ge = core.Int64Ptr(int64(0))
	dynamicRouteServerPeerGroupPrefixWatcherExcludedPrefixPrototypeModel.Le = core.Int64Ptr(int64(32))
	dynamicRouteServerPeerGroupPrefixWatcherExcludedPrefixPrototypeModel.Prefix = core.StringPtr("testString")

	deletedModel := new(vpcv1.Deleted)
	deletedModel.MoreInfo = core.StringPtr("https://cloud.ibm.com/apidocs/vpc#deleted-resources")

	dynamicRouteServerPeerGroupPolicyNextHopModel := new(vpcv1.DynamicRouteServerPeerGroupPolicyNextHopDynamicRouteServerPeerGroupReference)
	dynamicRouteServerPeerGroupPolicyNextHopModel.Deleted = deletedModel
	dynamicRouteServerPeerGroupPolicyNextHopModel.Href = core.StringPtr("https://us-south.iaas.cloud.ibm.com/v1/dynamic_route_servers/r006-d7cc5196-9864-48c4-82d8-3f30da41fcc5/peer_groups/r006-8228af60-189d-11ed-861d-0242ac120002")
	dynamicRouteServerPeerGroupPolicyNextHopModel.ID = core.StringPtr("r006-8228af60-189d-11ed-861d-0242ac120004")
	dynamicRouteServerPeerGroupPolicyNextHopModel.Name = core.StringPtr("my-dynamic-route-server-peer-group")
	dynamicRouteServerPeerGroupPolicyNextHopModel.ResourceType = core.StringPtr("dynamic_route_server_peer_group")

	routingTableReferenceModel := new(vpcv1.RoutingTableReference)
	routingTableReferenceModel.CRN = core.StringPtr("crn:v1:bluemix:public:is:us-south:a/aa2432b1fa4d4ace891e9b80fc104e34::vpc-routing-table:r006-4727d842-f94f-4a2d-824a-9bc9b02c523b/r006-6885e83f-03b2-4603-8a86-db2a0f55c840")
	routingTableReferenceModel.Deleted = deletedModel
	routingTableReferenceModel.Href = core.StringPtr("https://us-south.iaas.cloud.ibm.com/v1/vpcs/r006-4727d842-f94f-4a2d-824a-9bc9b02c523b/routing_tables/r006-6885e83f-03b2-4603-8a86-db2a0f55c840")
	routingTableReferenceModel.ID = core.StringPtr("r006-6885e83f-03b2-4603-8a86-db2a0f55c840")
	routingTableReferenceModel.Name = core.StringPtr("my-routing-table-1")
	routingTableReferenceModel.ResourceType = core.StringPtr("routing_table")

	dynamicRouteServerPeerGroupPolicyVPCRoutingTablesRoutingTableModel := new(vpcv1.DynamicRouteServerPeerGroupPolicyVPCRoutingTablesRoutingTable)
	dynamicRouteServerPeerGroupPolicyVPCRoutingTablesRoutingTableModel.Advertise = core.BoolPtr(true)
	dynamicRouteServerPeerGroupPolicyVPCRoutingTablesRoutingTableModel.VPCRoutingTable = routingTableReferenceModel

	model := new(vpcv1.DynamicRouteServerPeerGroupPolicyVPCRoutingTables)
	model.CreatedAt = CreateMockDateTime("2026-01-02T03:04:05.006Z")
	model.Href = core.StringPtr("https://us-south.iaas.cloud.ibm.com/v1/dynamic_route_servers/r006-d7cc5196-9864-48c4-82d8-3f30da41fcc5/peer_groups/r006-d7cc5196-9864-48c4-82d8-3f30da41fcc5/policies/r006-d7cc5196-9864-48c4-82d8-3f30da41fcc5")
	model.ID = core.StringPtr("r006-d7cc5196-9864-48c4-82d8-3f30da41fcc5")
	model.Name = core.StringPtr("my-dynamic-route-server-peer-group-policy")
	model.ResourceType = core.StringPtr("dynamic_route_server_peer_group_policy")
	model.State = core.StringPtr("disabled")
	model.ExcludedPrefixes = []vpcv1.DynamicRouteServerPeerGroupPrefixWatcherExcludedPrefixPrototype{*dynamicRouteServerPeerGroupPrefixWatcherExcludedPrefixPrototypeModel}
	model.NextHops = []vpcv1.DynamicRouteServerPeerGroupPolicyNextHopIntf{dynamicRouteServerPeerGroupPolicyNextHopModel}
	model.RouteDeleteDelay = core.Int64Ptr(int64(0))
	model.RoutingTables = []vpcv1.DynamicRouteServerPeerGroupPolicyVPCRoutingTablesRoutingTable{*dynamicRouteServerPeerGroupPolicyVPCRoutingTablesRoutingTableModel}
	model.Type = core.StringPtr("vpc_routing_tables")

	result, err := vpc.DataSourceIBMIsDynamicRouteServerPeerGroupPoliciesDynamicRouteServerPeerGroupPolicyVPCRoutingTablesToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}
