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
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/stretchr/testify/assert"
)

func TestAccIBMIsDynamicRouteServerPeerGroupPolicyDataSourceBasic(t *testing.T) {
	dynamicRouteServerPeerGroupPolicyDynamicRouteServerID := fmt.Sprintf("tf_dynamic_route_server_id_%d", acctest.RandIntRange(10, 100))
	dynamicRouteServerPeerGroupPolicyDynamicRouteServerPeerGroupID := fmt.Sprintf("tf_dynamic_route_server_peer_group_id_%d", acctest.RandIntRange(10, 100))
	dynamicRouteServerPeerGroupPolicyType := "custom_routes"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsDynamicRouteServerPeerGroupPolicyDataSourceConfigBasic(dynamicRouteServerPeerGroupPolicyDynamicRouteServerID, dynamicRouteServerPeerGroupPolicyDynamicRouteServerPeerGroupID, dynamicRouteServerPeerGroupPolicyType),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_peer_group_policy.is_dynamic_route_server_peer_group_policy_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_peer_group_policy.is_dynamic_route_server_peer_group_policy_instance", "dynamic_route_server_id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_peer_group_policy.is_dynamic_route_server_peer_group_policy_instance", "dynamic_route_server_peer_group_id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_peer_group_policy.is_dynamic_route_server_peer_group_policy_instance", "is_dynamic_route_server_peer_group_policy_id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_peer_group_policy.is_dynamic_route_server_peer_group_policy_instance", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_peer_group_policy.is_dynamic_route_server_peer_group_policy_instance", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_peer_group_policy.is_dynamic_route_server_peer_group_policy_instance", "state"),
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_peer_group_policy.is_dynamic_route_server_peer_group_policy_instance", "type"),
				),
			},
		},
	})
}

func TestAccIBMIsDynamicRouteServerPeerGroupPolicyDataSourceAllArgs(t *testing.T) {
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
				Config: testAccCheckIBMIsDynamicRouteServerPeerGroupPolicyDataSourceConfig(dynamicRouteServerPeerGroupPolicyDynamicRouteServerID, dynamicRouteServerPeerGroupPolicyDynamicRouteServerPeerGroupID, dynamicRouteServerPeerGroupPolicyName, dynamicRouteServerPeerGroupPolicyState, dynamicRouteServerPeerGroupPolicyType, dynamicRouteServerPeerGroupPolicyRouteDeleteDelay),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_peer_group_policy.is_dynamic_route_server_peer_group_policy_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_peer_group_policy.is_dynamic_route_server_peer_group_policy_instance", "dynamic_route_server_id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_peer_group_policy.is_dynamic_route_server_peer_group_policy_instance", "dynamic_route_server_peer_group_id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_peer_group_policy.is_dynamic_route_server_peer_group_policy_instance", "is_dynamic_route_server_peer_group_policy_id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_peer_group_policy.is_dynamic_route_server_peer_group_policy_instance", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_peer_group_policy.is_dynamic_route_server_peer_group_policy_instance", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_peer_group_policy.is_dynamic_route_server_peer_group_policy_instance", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_peer_group_policy.is_dynamic_route_server_peer_group_policy_instance", "resource_type"),
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_peer_group_policy.is_dynamic_route_server_peer_group_policy_instance", "state"),
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_peer_group_policy.is_dynamic_route_server_peer_group_policy_instance", "type"),
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_peer_group_policy.is_dynamic_route_server_peer_group_policy_instance", "custom_routes.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_peer_group_policy.is_dynamic_route_server_peer_group_policy_instance", "custom_routes.0.destination"),
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_peer_group_policy.is_dynamic_route_server_peer_group_policy_instance", "excluded_prefixes.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_peer_group_policy.is_dynamic_route_server_peer_group_policy_instance", "excluded_prefixes.0.ge"),
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_peer_group_policy.is_dynamic_route_server_peer_group_policy_instance", "excluded_prefixes.0.le"),
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_peer_group_policy.is_dynamic_route_server_peer_group_policy_instance", "excluded_prefixes.0.prefix"),
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_peer_group_policy.is_dynamic_route_server_peer_group_policy_instance", "peer_groups.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_peer_group_policy.is_dynamic_route_server_peer_group_policy_instance", "peer_groups.0.href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_peer_group_policy.is_dynamic_route_server_peer_group_policy_instance", "peer_groups.0.id"),
					resource.TestCheckResourceAttr("data.ibm_is_dynamic_route_server_peer_group_policy.is_dynamic_route_server_peer_group_policy_instance", "peer_groups.0.name", dynamicRouteServerPeerGroupPolicyName),
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_peer_group_policy.is_dynamic_route_server_peer_group_policy_instance", "peer_groups.0.resource_type"),
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_peer_group_policy.is_dynamic_route_server_peer_group_policy_instance", "next_hops.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_peer_group_policy.is_dynamic_route_server_peer_group_policy_instance", "next_hops.0.href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_peer_group_policy.is_dynamic_route_server_peer_group_policy_instance", "next_hops.0.id"),
					resource.TestCheckResourceAttr("data.ibm_is_dynamic_route_server_peer_group_policy.is_dynamic_route_server_peer_group_policy_instance", "next_hops.0.name", dynamicRouteServerPeerGroupPolicyName),
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_peer_group_policy.is_dynamic_route_server_peer_group_policy_instance", "next_hops.0.resource_type"),
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_peer_group_policy.is_dynamic_route_server_peer_group_policy_instance", "route_delete_delay"),
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_peer_group_policy.is_dynamic_route_server_peer_group_policy_instance", "routing_tables.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_peer_group_policy.is_dynamic_route_server_peer_group_policy_instance", "routing_tables.0.advertise"),
				),
			},
		},
	})
}

func testAccCheckIBMIsDynamicRouteServerPeerGroupPolicyDataSourceConfigBasic(dynamicRouteServerPeerGroupPolicyDynamicRouteServerID string, dynamicRouteServerPeerGroupPolicyDynamicRouteServerPeerGroupID string, dynamicRouteServerPeerGroupPolicyType string) string {
	return fmt.Sprintf(`
		resource "ibm_is_dynamic_route_server_peer_group_policy" "is_dynamic_route_server_peer_group_policy_instance" {
			dynamic_route_server_id = "%s"
			dynamic_route_server_peer_group_id = "%s"
			type = "%s"
		}

		data "ibm_is_dynamic_route_server_peer_group_policy" "is_dynamic_route_server_peer_group_policy_instance" {
			dynamic_route_server_id = ibm_is_dynamic_route_server_peer_group_policy.is_dynamic_route_server_peer_group_policy_instance.dynamic_route_server_id
			dynamic_route_server_peer_group_id = ibm_is_dynamic_route_server_peer_group_policy.is_dynamic_route_server_peer_group_policy_instance.dynamic_route_server_peer_group_id
			is_dynamic_route_server_peer_group_policy_id = ibm_is_dynamic_route_server_peer_group_policy.is_dynamic_route_server_peer_group_policy_instance.is_dynamic_route_server_peer_group_policy_id
		}
	`, dynamicRouteServerPeerGroupPolicyDynamicRouteServerID, dynamicRouteServerPeerGroupPolicyDynamicRouteServerPeerGroupID, dynamicRouteServerPeerGroupPolicyType)
}

func testAccCheckIBMIsDynamicRouteServerPeerGroupPolicyDataSourceConfig(dynamicRouteServerPeerGroupPolicyDynamicRouteServerID string, dynamicRouteServerPeerGroupPolicyDynamicRouteServerPeerGroupID string, dynamicRouteServerPeerGroupPolicyName string, dynamicRouteServerPeerGroupPolicyState string, dynamicRouteServerPeerGroupPolicyType string, dynamicRouteServerPeerGroupPolicyRouteDeleteDelay string) string {
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

		data "ibm_is_dynamic_route_server_peer_group_policy" "is_dynamic_route_server_peer_group_policy_instance" {
			dynamic_route_server_id = ibm_is_dynamic_route_server_peer_group_policy.is_dynamic_route_server_peer_group_policy_instance.dynamic_route_server_id
			dynamic_route_server_peer_group_id = ibm_is_dynamic_route_server_peer_group_policy.is_dynamic_route_server_peer_group_policy_instance.dynamic_route_server_peer_group_id
			is_dynamic_route_server_peer_group_policy_id = ibm_is_dynamic_route_server_peer_group_policy.is_dynamic_route_server_peer_group_policy_instance.is_dynamic_route_server_peer_group_policy_id
		}
	`, dynamicRouteServerPeerGroupPolicyDynamicRouteServerID, dynamicRouteServerPeerGroupPolicyDynamicRouteServerPeerGroupID, dynamicRouteServerPeerGroupPolicyName, dynamicRouteServerPeerGroupPolicyState, dynamicRouteServerPeerGroupPolicyType, dynamicRouteServerPeerGroupPolicyRouteDeleteDelay)
}

func TestDataSourceIBMIsDynamicRouteServerPeerGroupPolicyDynamicRouteServerPeerGroupPolicyCustomRouteToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["destination"] = "192.168.3.0/24"

		assert.Equal(t, result, model)
	}

	model := new(vpcv1.DynamicRouteServerPeerGroupPolicyCustomRoute)
	model.Destination = core.StringPtr("192.168.3.0/24")

	result, err := vpc.DataSourceIBMIsDynamicRouteServerPeerGroupPolicyDynamicRouteServerPeerGroupPolicyCustomRouteToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMIsDynamicRouteServerPeerGroupPolicyDynamicRouteServerPeerGroupPrefixWatcherExcludedPrefixPrototypeToMap(t *testing.T) {
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

	result, err := vpc.DataSourceIBMIsDynamicRouteServerPeerGroupPolicyDynamicRouteServerPeerGroupPrefixWatcherExcludedPrefixPrototypeToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMIsDynamicRouteServerPeerGroupPolicyDynamicRouteServerPeerGroupReferenceToMap(t *testing.T) {
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

	result, err := vpc.DataSourceIBMIsDynamicRouteServerPeerGroupPolicyDynamicRouteServerPeerGroupReferenceToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMIsDynamicRouteServerPeerGroupPolicyDeletedToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["more_info"] = "https://cloud.ibm.com/apidocs/vpc#deleted-resources"

		assert.Equal(t, result, model)
	}

	model := new(vpcv1.Deleted)
	model.MoreInfo = core.StringPtr("https://cloud.ibm.com/apidocs/vpc#deleted-resources")

	result, err := vpc.DataSourceIBMIsDynamicRouteServerPeerGroupPolicyDeletedToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMIsDynamicRouteServerPeerGroupPolicyDynamicRouteServerPeerGroupPolicyNextHopToMap(t *testing.T) {
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

	result, err := vpc.DataSourceIBMIsDynamicRouteServerPeerGroupPolicyDynamicRouteServerPeerGroupPolicyNextHopToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMIsDynamicRouteServerPeerGroupPolicyDynamicRouteServerPeerGroupPolicyNextHopDynamicRouteServerPeerGroupReferenceToMap(t *testing.T) {
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

	result, err := vpc.DataSourceIBMIsDynamicRouteServerPeerGroupPolicyDynamicRouteServerPeerGroupPolicyNextHopDynamicRouteServerPeerGroupReferenceToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMIsDynamicRouteServerPeerGroupPolicyDynamicRouteServerPeerGroupPolicyVPCRoutingTablesRoutingTableToMap(t *testing.T) {
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

	result, err := vpc.DataSourceIBMIsDynamicRouteServerPeerGroupPolicyDynamicRouteServerPeerGroupPolicyVPCRoutingTablesRoutingTableToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMIsDynamicRouteServerPeerGroupPolicyRoutingTableReferenceToMap(t *testing.T) {
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

	result, err := vpc.DataSourceIBMIsDynamicRouteServerPeerGroupPolicyRoutingTableReferenceToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}
