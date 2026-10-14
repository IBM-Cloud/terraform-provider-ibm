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

func TestAccIBMIsDynamicRouteServerPeerGroupPolicyBasic(t *testing.T) {
	var conf vpcv1.DynamicRouteServerPeerGroupPolicy
	dynamicRouteServerID := fmt.Sprintf("tf_dynamic_route_server_id_%d", acctest.RandIntRange(10, 100))
	dynamicRouteServerPeerGroupID := fmt.Sprintf("tf_dynamic_route_server_peer_group_id_%d", acctest.RandIntRange(10, 100))
	typeVar := "custom_routes"
	typeVarUpdate := "vpc_routing_tables"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIsDynamicRouteServerPeerGroupPolicyDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsDynamicRouteServerPeerGroupPolicyConfigBasic(dynamicRouteServerID, dynamicRouteServerPeerGroupID, typeVar),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIsDynamicRouteServerPeerGroupPolicyExists("ibm_is_dynamic_route_server_peer_group_policy.is_dynamic_route_server_peer_group_policy_instance", conf),
					resource.TestCheckResourceAttr("ibm_is_dynamic_route_server_peer_group_policy.is_dynamic_route_server_peer_group_policy_instance", "dynamic_route_server_id", dynamicRouteServerID),
					resource.TestCheckResourceAttr("ibm_is_dynamic_route_server_peer_group_policy.is_dynamic_route_server_peer_group_policy_instance", "dynamic_route_server_peer_group_id", dynamicRouteServerPeerGroupID),
					resource.TestCheckResourceAttr("ibm_is_dynamic_route_server_peer_group_policy.is_dynamic_route_server_peer_group_policy_instance", "type", typeVar),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMIsDynamicRouteServerPeerGroupPolicyConfigBasic(dynamicRouteServerID, dynamicRouteServerPeerGroupID, typeVarUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_is_dynamic_route_server_peer_group_policy.is_dynamic_route_server_peer_group_policy_instance", "dynamic_route_server_id", dynamicRouteServerID),
					resource.TestCheckResourceAttr("ibm_is_dynamic_route_server_peer_group_policy.is_dynamic_route_server_peer_group_policy_instance", "dynamic_route_server_peer_group_id", dynamicRouteServerPeerGroupID),
					resource.TestCheckResourceAttr("ibm_is_dynamic_route_server_peer_group_policy.is_dynamic_route_server_peer_group_policy_instance", "type", typeVarUpdate),
				),
			},
		},
	})
}

func TestAccIBMIsDynamicRouteServerPeerGroupPolicyAllArgs(t *testing.T) {
	var conf vpcv1.DynamicRouteServerPeerGroupPolicy
	dynamicRouteServerID := fmt.Sprintf("tf_dynamic_route_server_id_%d", acctest.RandIntRange(10, 100))
	dynamicRouteServerPeerGroupID := fmt.Sprintf("tf_dynamic_route_server_peer_group_id_%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	state := "disabled"
	typeVar := "custom_routes"
	routeDeleteDelay := fmt.Sprintf("%d", acctest.RandIntRange(0, 60))
	nameUpdate := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	stateUpdate := "enabled"
	typeVarUpdate := "vpc_routing_tables"
	routeDeleteDelayUpdate := fmt.Sprintf("%d", acctest.RandIntRange(0, 60))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIsDynamicRouteServerPeerGroupPolicyDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsDynamicRouteServerPeerGroupPolicyConfig(dynamicRouteServerID, dynamicRouteServerPeerGroupID, name, state, typeVar, routeDeleteDelay),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIsDynamicRouteServerPeerGroupPolicyExists("ibm_is_dynamic_route_server_peer_group_policy.is_dynamic_route_server_peer_group_policy_instance", conf),
					resource.TestCheckResourceAttr("ibm_is_dynamic_route_server_peer_group_policy.is_dynamic_route_server_peer_group_policy_instance", "dynamic_route_server_id", dynamicRouteServerID),
					resource.TestCheckResourceAttr("ibm_is_dynamic_route_server_peer_group_policy.is_dynamic_route_server_peer_group_policy_instance", "dynamic_route_server_peer_group_id", dynamicRouteServerPeerGroupID),
					resource.TestCheckResourceAttr("ibm_is_dynamic_route_server_peer_group_policy.is_dynamic_route_server_peer_group_policy_instance", "name", name),
					resource.TestCheckResourceAttr("ibm_is_dynamic_route_server_peer_group_policy.is_dynamic_route_server_peer_group_policy_instance", "state", state),
					resource.TestCheckResourceAttr("ibm_is_dynamic_route_server_peer_group_policy.is_dynamic_route_server_peer_group_policy_instance", "type", typeVar),
					resource.TestCheckResourceAttr("ibm_is_dynamic_route_server_peer_group_policy.is_dynamic_route_server_peer_group_policy_instance", "route_delete_delay", routeDeleteDelay),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMIsDynamicRouteServerPeerGroupPolicyConfig(dynamicRouteServerID, dynamicRouteServerPeerGroupID, nameUpdate, stateUpdate, typeVarUpdate, routeDeleteDelayUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_is_dynamic_route_server_peer_group_policy.is_dynamic_route_server_peer_group_policy_instance", "dynamic_route_server_id", dynamicRouteServerID),
					resource.TestCheckResourceAttr("ibm_is_dynamic_route_server_peer_group_policy.is_dynamic_route_server_peer_group_policy_instance", "dynamic_route_server_peer_group_id", dynamicRouteServerPeerGroupID),
					resource.TestCheckResourceAttr("ibm_is_dynamic_route_server_peer_group_policy.is_dynamic_route_server_peer_group_policy_instance", "name", nameUpdate),
					resource.TestCheckResourceAttr("ibm_is_dynamic_route_server_peer_group_policy.is_dynamic_route_server_peer_group_policy_instance", "state", stateUpdate),
					resource.TestCheckResourceAttr("ibm_is_dynamic_route_server_peer_group_policy.is_dynamic_route_server_peer_group_policy_instance", "type", typeVarUpdate),
					resource.TestCheckResourceAttr("ibm_is_dynamic_route_server_peer_group_policy.is_dynamic_route_server_peer_group_policy_instance", "route_delete_delay", routeDeleteDelayUpdate),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_is_dynamic_route_server_peer_group_policy.is_dynamic_route_server_peer_group_policy_instance",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIBMIsDynamicRouteServerPeerGroupPolicyConfigBasic(dynamicRouteServerID string, dynamicRouteServerPeerGroupID string, typeVar string) string {
	return fmt.Sprintf(`
		resource "ibm_is_dynamic_route_server_peer_group_policy" "is_dynamic_route_server_peer_group_policy_instance" {
			dynamic_route_server_id = "%s"
			dynamic_route_server_peer_group_id = "%s"
			type = "%s"
		}
	`, dynamicRouteServerID, dynamicRouteServerPeerGroupID, typeVar)
}

func testAccCheckIBMIsDynamicRouteServerPeerGroupPolicyConfig(dynamicRouteServerID string, dynamicRouteServerPeerGroupID string, name string, state string, typeVar string, routeDeleteDelay string) string {
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
	`, dynamicRouteServerID, dynamicRouteServerPeerGroupID, name, state, typeVar, routeDeleteDelay)
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

		getDynamicRouteServerPeerGroupPolicyOptions := &vpcv1.GetDynamicRouteServerPeerGroupPolicyOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getDynamicRouteServerPeerGroupPolicyOptions.SetDynamicRouteServerID(parts[0])
		getDynamicRouteServerPeerGroupPolicyOptions.SetDynamicRouteServerPeerGroupID(parts[1])
		getDynamicRouteServerPeerGroupPolicyOptions.SetID(parts[2])

		dynamicRouteServerPeerGroupPolicyIntf, _, err := vpcClient.GetDynamicRouteServerPeerGroupPolicy(getDynamicRouteServerPeerGroupPolicyOptions)
		if err != nil {
			return err
		}

		dynamicRouteServerPeerGroupPolicy := dynamicRouteServerPeerGroupPolicyIntf.(*vpcv1.DynamicRouteServerPeerGroupPolicy)
		obj = *dynamicRouteServerPeerGroupPolicy
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

		getDynamicRouteServerPeerGroupPolicyOptions := &vpcv1.GetDynamicRouteServerPeerGroupPolicyOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getDynamicRouteServerPeerGroupPolicyOptions.SetDynamicRouteServerID(parts[0])
		getDynamicRouteServerPeerGroupPolicyOptions.SetDynamicRouteServerPeerGroupID(parts[1])
		getDynamicRouteServerPeerGroupPolicyOptions.SetID(parts[2])

		// Try to find the key
		_, response, err := vpcClient.GetDynamicRouteServerPeerGroupPolicy(getDynamicRouteServerPeerGroupPolicyOptions)

		if err == nil {
			return fmt.Errorf("is_dynamic_route_server_peer_group_policy still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for is_dynamic_route_server_peer_group_policy (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}

func TestResourceIBMIsDynamicRouteServerPeerGroupPolicyDynamicRouteServerPeerGroupPolicyCustomRouteToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["destination"] = "192.168.3.0/24"

		assert.Equal(t, result, model)
	}

	model := new(vpcv1.DynamicRouteServerPeerGroupPolicyCustomRoute)
	model.Destination = core.StringPtr("192.168.3.0/24")

	result, err := vpc.ResourceIBMIsDynamicRouteServerPeerGroupPolicyDynamicRouteServerPeerGroupPolicyCustomRouteToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMIsDynamicRouteServerPeerGroupPolicyDynamicRouteServerPeerGroupPrefixWatcherExcludedPrefixPrototypeToMap(t *testing.T) {
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

	result, err := vpc.ResourceIBMIsDynamicRouteServerPeerGroupPolicyDynamicRouteServerPeerGroupPrefixWatcherExcludedPrefixPrototypeToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMIsDynamicRouteServerPeerGroupPolicyDynamicRouteServerPeerGroupReferenceToMap(t *testing.T) {
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

	result, err := vpc.ResourceIBMIsDynamicRouteServerPeerGroupPolicyDynamicRouteServerPeerGroupReferenceToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMIsDynamicRouteServerPeerGroupPolicyDeletedToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["more_info"] = "https://cloud.ibm.com/apidocs/vpc#deleted-resources"

		assert.Equal(t, result, model)
	}

	model := new(vpcv1.Deleted)
	model.MoreInfo = core.StringPtr("https://cloud.ibm.com/apidocs/vpc#deleted-resources")

	result, err := vpc.ResourceIBMIsDynamicRouteServerPeerGroupPolicyDeletedToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMIsDynamicRouteServerPeerGroupPolicyDynamicRouteServerPeerGroupPolicyNextHopToMap(t *testing.T) {
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

	result, err := vpc.ResourceIBMIsDynamicRouteServerPeerGroupPolicyDynamicRouteServerPeerGroupPolicyNextHopToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMIsDynamicRouteServerPeerGroupPolicyDynamicRouteServerPeerGroupPolicyNextHopDynamicRouteServerPeerGroupReferenceToMap(t *testing.T) {
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

	result, err := vpc.ResourceIBMIsDynamicRouteServerPeerGroupPolicyDynamicRouteServerPeerGroupPolicyNextHopDynamicRouteServerPeerGroupReferenceToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMIsDynamicRouteServerPeerGroupPolicyDynamicRouteServerPeerGroupPolicyVPCRoutingTablesRoutingTableToMap(t *testing.T) {
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

	result, err := vpc.ResourceIBMIsDynamicRouteServerPeerGroupPolicyDynamicRouteServerPeerGroupPolicyVPCRoutingTablesRoutingTableToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMIsDynamicRouteServerPeerGroupPolicyRoutingTableReferenceToMap(t *testing.T) {
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

	result, err := vpc.ResourceIBMIsDynamicRouteServerPeerGroupPolicyRoutingTableReferenceToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToDynamicRouteServerPeerGroupPolicyCustomRoutePrototype(t *testing.T) {
	checkResult := func(result *vpcv1.DynamicRouteServerPeerGroupPolicyCustomRoutePrototype) {
		model := new(vpcv1.DynamicRouteServerPeerGroupPolicyCustomRoutePrototype)
		model.Destination = core.StringPtr("192.168.3.0/24")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["destination"] = "192.168.3.0/24"

	result, err := vpc.ResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToDynamicRouteServerPeerGroupPolicyCustomRoutePrototype(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToDynamicRouteServerPeerGroupPrefixWatcherExcludedPrefixPrototype(t *testing.T) {
	checkResult := func(result *vpcv1.DynamicRouteServerPeerGroupPrefixWatcherExcludedPrefixPrototype) {
		model := new(vpcv1.DynamicRouteServerPeerGroupPrefixWatcherExcludedPrefixPrototype)
		model.Ge = core.Int64Ptr(int64(0))
		model.Le = core.Int64Ptr(int64(32))
		model.Prefix = core.StringPtr("testString")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["ge"] = int(0)
	model["le"] = int(32)
	model["prefix"] = "testString"

	result, err := vpc.ResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToDynamicRouteServerPeerGroupPrefixWatcherExcludedPrefixPrototype(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToDynamicRouteServerPeerGroupIdentity(t *testing.T) {
	checkResult := func(result vpcv1.DynamicRouteServerPeerGroupIdentityIntf) {
		model := new(vpcv1.DynamicRouteServerPeerGroupIdentity)
		model.ID = core.StringPtr("r006-8228af60-189d-11ed-861d-0242ac120002")
		model.Href = core.StringPtr("https://us-south.iaas.cloud.ibm.com/v1/dynamic_route_servers/r006-d7cc5196-9864-48c4-82d8-3f30da41fcc5/peer_groups/r006-8228af60-189d-11ed-861d-0242ac120002")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["id"] = "r006-8228af60-189d-11ed-861d-0242ac120002"
	model["href"] = "https://us-south.iaas.cloud.ibm.com/v1/dynamic_route_servers/r006-d7cc5196-9864-48c4-82d8-3f30da41fcc5/peer_groups/r006-8228af60-189d-11ed-861d-0242ac120002"

	result, err := vpc.ResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToDynamicRouteServerPeerGroupIdentity(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToDynamicRouteServerPeerGroupIdentityByID(t *testing.T) {
	checkResult := func(result *vpcv1.DynamicRouteServerPeerGroupIdentityByID) {
		model := new(vpcv1.DynamicRouteServerPeerGroupIdentityByID)
		model.ID = core.StringPtr("r006-8228af60-189d-11ed-861d-0242ac120002")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["id"] = "r006-8228af60-189d-11ed-861d-0242ac120002"

	result, err := vpc.ResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToDynamicRouteServerPeerGroupIdentityByID(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToDynamicRouteServerPeerGroupIdentityByHref(t *testing.T) {
	checkResult := func(result *vpcv1.DynamicRouteServerPeerGroupIdentityByHref) {
		model := new(vpcv1.DynamicRouteServerPeerGroupIdentityByHref)
		model.Href = core.StringPtr("https://us-south.iaas.cloud.ibm.com/v1/dynamic_route_servers/r006-d7cc5196-9864-48c4-82d8-3f30da41fcc5/peer_groups/r006-8228af60-189d-11ed-861d-0242ac120002")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["href"] = "https://us-south.iaas.cloud.ibm.com/v1/dynamic_route_servers/r006-d7cc5196-9864-48c4-82d8-3f30da41fcc5/peer_groups/r006-8228af60-189d-11ed-861d-0242ac120002"

	result, err := vpc.ResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToDynamicRouteServerPeerGroupIdentityByHref(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToDynamicRouteServerPeerGroupPolicyNextHopPrototype(t *testing.T) {
	checkResult := func(result vpcv1.DynamicRouteServerPeerGroupPolicyNextHopPrototypeIntf) {
		model := new(vpcv1.DynamicRouteServerPeerGroupPolicyNextHopPrototype)
		model.ID = core.StringPtr("r006-8228af60-189d-11ed-861d-0242ac120002")
		model.Href = core.StringPtr("https://us-south.iaas.cloud.ibm.com/v1/dynamic_route_servers/r006-d7cc5196-9864-48c4-82d8-3f30da41fcc5/peer_groups/r006-8228af60-189d-11ed-861d-0242ac120002")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["id"] = "r006-8228af60-189d-11ed-861d-0242ac120002"
	model["href"] = "https://us-south.iaas.cloud.ibm.com/v1/dynamic_route_servers/r006-d7cc5196-9864-48c4-82d8-3f30da41fcc5/peer_groups/r006-8228af60-189d-11ed-861d-0242ac120002"

	result, err := vpc.ResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToDynamicRouteServerPeerGroupPolicyNextHopPrototype(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToDynamicRouteServerPeerGroupPolicyNextHopPrototypeDynamicRouteServerPeerGroupIdentity(t *testing.T) {
	checkResult := func(result vpcv1.DynamicRouteServerPeerGroupPolicyNextHopPrototypeDynamicRouteServerPeerGroupIdentityIntf) {
		model := new(vpcv1.DynamicRouteServerPeerGroupPolicyNextHopPrototypeDynamicRouteServerPeerGroupIdentity)
		model.ID = core.StringPtr("r006-8228af60-189d-11ed-861d-0242ac120002")
		model.Href = core.StringPtr("https://us-south.iaas.cloud.ibm.com/v1/dynamic_route_servers/r006-d7cc5196-9864-48c4-82d8-3f30da41fcc5/peer_groups/r006-8228af60-189d-11ed-861d-0242ac120002")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["id"] = "r006-8228af60-189d-11ed-861d-0242ac120002"
	model["href"] = "https://us-south.iaas.cloud.ibm.com/v1/dynamic_route_servers/r006-d7cc5196-9864-48c4-82d8-3f30da41fcc5/peer_groups/r006-8228af60-189d-11ed-861d-0242ac120002"

	result, err := vpc.ResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToDynamicRouteServerPeerGroupPolicyNextHopPrototypeDynamicRouteServerPeerGroupIdentity(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToDynamicRouteServerPeerGroupPolicyNextHopPrototypeDynamicRouteServerPeerGroupIdentityDynamicRouteServerPeerGroupIdentityByID(t *testing.T) {
	checkResult := func(result *vpcv1.DynamicRouteServerPeerGroupPolicyNextHopPrototypeDynamicRouteServerPeerGroupIdentityDynamicRouteServerPeerGroupIdentityByID) {
		model := new(vpcv1.DynamicRouteServerPeerGroupPolicyNextHopPrototypeDynamicRouteServerPeerGroupIdentityDynamicRouteServerPeerGroupIdentityByID)
		model.ID = core.StringPtr("r006-8228af60-189d-11ed-861d-0242ac120002")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["id"] = "r006-8228af60-189d-11ed-861d-0242ac120002"

	result, err := vpc.ResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToDynamicRouteServerPeerGroupPolicyNextHopPrototypeDynamicRouteServerPeerGroupIdentityDynamicRouteServerPeerGroupIdentityByID(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToDynamicRouteServerPeerGroupPolicyNextHopPrototypeDynamicRouteServerPeerGroupIdentityDynamicRouteServerPeerGroupIdentityByHref(t *testing.T) {
	checkResult := func(result *vpcv1.DynamicRouteServerPeerGroupPolicyNextHopPrototypeDynamicRouteServerPeerGroupIdentityDynamicRouteServerPeerGroupIdentityByHref) {
		model := new(vpcv1.DynamicRouteServerPeerGroupPolicyNextHopPrototypeDynamicRouteServerPeerGroupIdentityDynamicRouteServerPeerGroupIdentityByHref)
		model.Href = core.StringPtr("https://us-south.iaas.cloud.ibm.com/v1/dynamic_route_servers/r006-d7cc5196-9864-48c4-82d8-3f30da41fcc5/peer_groups/r006-8228af60-189d-11ed-861d-0242ac120002")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["href"] = "https://us-south.iaas.cloud.ibm.com/v1/dynamic_route_servers/r006-d7cc5196-9864-48c4-82d8-3f30da41fcc5/peer_groups/r006-8228af60-189d-11ed-861d-0242ac120002"

	result, err := vpc.ResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToDynamicRouteServerPeerGroupPolicyNextHopPrototypeDynamicRouteServerPeerGroupIdentityDynamicRouteServerPeerGroupIdentityByHref(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToDynamicRouteServerPeerGroupPolicyVPCRoutingTablesRoutingTablePrototype(t *testing.T) {
	checkResult := func(result *vpcv1.DynamicRouteServerPeerGroupPolicyVPCRoutingTablesRoutingTablePrototype) {
		routingTableIdentityModel := new(vpcv1.RoutingTableIdentityByID)
		routingTableIdentityModel.ID = core.StringPtr("r006-6885e83f-03b2-4603-8a86-db2a0f55c840")

		model := new(vpcv1.DynamicRouteServerPeerGroupPolicyVPCRoutingTablesRoutingTablePrototype)
		model.Advertise = core.BoolPtr(false)
		model.VPCRoutingTable = routingTableIdentityModel

		assert.Equal(t, result, model)
	}

	routingTableIdentityModel := make(map[string]interface{})
	routingTableIdentityModel["id"] = "r006-6885e83f-03b2-4603-8a86-db2a0f55c840"

	model := make(map[string]interface{})
	model["advertise"] = false
	model["vpc_routing_table"] = []interface{}{routingTableIdentityModel}

	result, err := vpc.ResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToDynamicRouteServerPeerGroupPolicyVPCRoutingTablesRoutingTablePrototype(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToRoutingTableIdentity(t *testing.T) {
	checkResult := func(result vpcv1.RoutingTableIdentityIntf) {
		model := new(vpcv1.RoutingTableIdentity)
		model.ID = core.StringPtr("r006-6885e83f-03b2-4603-8a86-db2a0f55c840")
		model.CRN = core.StringPtr("crn:v1:bluemix:public:is:us-south:a/aa2432b1fa4d4ace891e9b80fc104e34::vpc-routing-table:r006-4727d842-f94f-4a2d-824a-9bc9b02c523b/r006-6885e83f-03b2-4603-8a86-db2a0f55c840")
		model.Href = core.StringPtr("https://us-south.iaas.cloud.ibm.com/v1/vpcs/r006-4727d842-f94f-4a2d-824a-9bc9b02c523b/routing_tables/r006-6885e83f-03b2-4603-8a86-db2a0f55c840")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["id"] = "r006-6885e83f-03b2-4603-8a86-db2a0f55c840"
	model["crn"] = "crn:v1:bluemix:public:is:us-south:a/aa2432b1fa4d4ace891e9b80fc104e34::vpc-routing-table:r006-4727d842-f94f-4a2d-824a-9bc9b02c523b/r006-6885e83f-03b2-4603-8a86-db2a0f55c840"
	model["href"] = "https://us-south.iaas.cloud.ibm.com/v1/vpcs/r006-4727d842-f94f-4a2d-824a-9bc9b02c523b/routing_tables/r006-6885e83f-03b2-4603-8a86-db2a0f55c840"

	result, err := vpc.ResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToRoutingTableIdentity(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToRoutingTableIdentityByID(t *testing.T) {
	checkResult := func(result *vpcv1.RoutingTableIdentityByID) {
		model := new(vpcv1.RoutingTableIdentityByID)
		model.ID = core.StringPtr("r006-6885e83f-03b2-4603-8a86-db2a0f55c840")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["id"] = "r006-6885e83f-03b2-4603-8a86-db2a0f55c840"

	result, err := vpc.ResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToRoutingTableIdentityByID(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToRoutingTableIdentityByCRN(t *testing.T) {
	checkResult := func(result *vpcv1.RoutingTableIdentityByCRN) {
		model := new(vpcv1.RoutingTableIdentityByCRN)
		model.CRN = core.StringPtr("crn:v1:bluemix:public:is:us-south:a/aa2432b1fa4d4ace891e9b80fc104e34::vpc-routing-table:r006-4727d842-f94f-4a2d-824a-9bc9b02c523b/r006-6885e83f-03b2-4603-8a86-db2a0f55c840")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["crn"] = "crn:v1:bluemix:public:is:us-south:a/aa2432b1fa4d4ace891e9b80fc104e34::vpc-routing-table:r006-4727d842-f94f-4a2d-824a-9bc9b02c523b/r006-6885e83f-03b2-4603-8a86-db2a0f55c840"

	result, err := vpc.ResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToRoutingTableIdentityByCRN(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToRoutingTableIdentityByHref(t *testing.T) {
	checkResult := func(result *vpcv1.RoutingTableIdentityByHref) {
		model := new(vpcv1.RoutingTableIdentityByHref)
		model.Href = core.StringPtr("https://us-south.iaas.cloud.ibm.com/v1/vpcs/r006-4727d842-f94f-4a2d-824a-9bc9b02c523b/routing_tables/r006-6885e83f-03b2-4603-8a86-db2a0f55c840")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["href"] = "https://us-south.iaas.cloud.ibm.com/v1/vpcs/r006-4727d842-f94f-4a2d-824a-9bc9b02c523b/routing_tables/r006-6885e83f-03b2-4603-8a86-db2a0f55c840"

	result, err := vpc.ResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToRoutingTableIdentityByHref(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToDynamicRouteServerPeerGroupPolicyVPCRoutingTablesRoutingTablePatch(t *testing.T) {
	checkResult := func(result *vpcv1.DynamicRouteServerPeerGroupPolicyVPCRoutingTablesRoutingTablePatch) {
		routingTableIdentityModel := new(vpcv1.RoutingTableIdentityByID)
		routingTableIdentityModel.ID = core.StringPtr("r006-6885e83f-03b2-4603-8a86-db2a0f55c840")

		model := new(vpcv1.DynamicRouteServerPeerGroupPolicyVPCRoutingTablesRoutingTablePatch)
		model.Advertise = core.BoolPtr(true)
		model.VPCRoutingTable = routingTableIdentityModel

		assert.Equal(t, result, model)
	}

	routingTableIdentityModel := make(map[string]interface{})
	routingTableIdentityModel["id"] = "r006-6885e83f-03b2-4603-8a86-db2a0f55c840"

	model := make(map[string]interface{})
	model["advertise"] = true
	model["vpc_routing_table"] = []interface{}{routingTableIdentityModel}

	result, err := vpc.ResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToDynamicRouteServerPeerGroupPolicyVPCRoutingTablesRoutingTablePatch(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToDynamicRouteServerPeerGroupPolicyPrototype(t *testing.T) {
	checkResult := func(result vpcv1.DynamicRouteServerPeerGroupPolicyPrototypeIntf) {
		dynamicRouteServerPeerGroupPolicyCustomRoutePrototypeModel := new(vpcv1.DynamicRouteServerPeerGroupPolicyCustomRoutePrototype)
		dynamicRouteServerPeerGroupPolicyCustomRoutePrototypeModel.Destination = core.StringPtr("192.168.3.0/24")

		model := new(vpcv1.DynamicRouteServerPeerGroupPolicyPrototype)
		model.Name = core.StringPtr("my-dynamic-route-server-peer-group-policy")
		model.State = core.StringPtr("enabled")
		model.Type = core.StringPtr("custom_routes")
		model.CustomRoutes = []vpcv1.DynamicRouteServerPeerGroupPolicyCustomRoutePrototype{*dynamicRouteServerPeerGroupPolicyCustomRoutePrototypeModel}
		// model.ExcludedPrefixes = []vpcv1.DynamicRouteServerPeerGroupPrefixWatcherExcludedPrefixPrototype{*dynamicRouteServerPeerGroupPrefixWatcherExcludedPrefixPrototypeModel}
		// model.PeerGroups = []vpcv1.DynamicRouteServerPeerGroupIdentityIntf{dynamicRouteServerPeerGroupIdentityModel}
		// model.NextHops = []vpcv1.DynamicRouteServerPeerGroupPolicyNextHopPrototypeIntf{dynamicRouteServerPeerGroupPolicyNextHopPrototypeModel}
		model.RouteDeleteDelay = core.Int64Ptr(int64(0))
		// model.RoutingTables = []vpcv1.DynamicRouteServerPeerGroupPolicyVPCRoutingTablesRoutingTablePrototype{*dynamicRouteServerPeerGroupPolicyVPCRoutingTablesRoutingTablePrototypeModel}

		assert.Equal(t, result, model)
	}

	dynamicRouteServerPeerGroupPolicyCustomRoutePrototypeModel := make(map[string]interface{})
	dynamicRouteServerPeerGroupPolicyCustomRoutePrototypeModel["destination"] = "192.168.3.0/24"

	model := make(map[string]interface{})
	model["name"] = "my-dynamic-route-server-peer-group-policy"
	model["state"] = "enabled"
	model["type"] = "custom_routes"
	model["custom_routes"] = []interface{}{dynamicRouteServerPeerGroupPolicyCustomRoutePrototypeModel}
	// model["excluded_prefixes"] = []interface{}{dynamicRouteServerPeerGroupPrefixWatcherExcludedPrefixPrototypeModel}
	// model["peer_groups"] = []interface{}{dynamicRouteServerPeerGroupIdentityModel}
	// model["next_hops"] = []interface{}{dynamicRouteServerPeerGroupPolicyNextHopPrototypeModel}
	model["route_delete_delay"] = int(0)
	// model["routing_tables"] = []interface{}{dynamicRouteServerPeerGroupPolicyVPCRoutingTablesRoutingTablePrototypeModel}

	result, err := vpc.ResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToDynamicRouteServerPeerGroupPolicyPrototype(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToDynamicRouteServerPeerGroupPolicyPrototypeDynamicRouteServerPeerGroupPolicyCustomRoutesPrototype(t *testing.T) {
	checkResult := func(result *vpcv1.DynamicRouteServerPeerGroupPolicyPrototypeDynamicRouteServerPeerGroupPolicyCustomRoutesPrototype) {
		dynamicRouteServerPeerGroupPolicyCustomRoutePrototypeModel := new(vpcv1.DynamicRouteServerPeerGroupPolicyCustomRoutePrototype)
		dynamicRouteServerPeerGroupPolicyCustomRoutePrototypeModel.Destination = core.StringPtr("192.168.3.0/24")

		model := new(vpcv1.DynamicRouteServerPeerGroupPolicyPrototypeDynamicRouteServerPeerGroupPolicyCustomRoutesPrototype)
		model.Name = core.StringPtr("my-dynamic-route-server-peer-group-policy")
		model.State = core.StringPtr("enabled")
		model.CustomRoutes = []vpcv1.DynamicRouteServerPeerGroupPolicyCustomRoutePrototype{*dynamicRouteServerPeerGroupPolicyCustomRoutePrototypeModel}
		model.Type = core.StringPtr("custom_routes")

		assert.Equal(t, result, model)
	}

	dynamicRouteServerPeerGroupPolicyCustomRoutePrototypeModel := make(map[string]interface{})
	dynamicRouteServerPeerGroupPolicyCustomRoutePrototypeModel["destination"] = "192.168.3.0/24"

	model := make(map[string]interface{})
	model["name"] = "my-dynamic-route-server-peer-group-policy"
	model["state"] = "enabled"
	model["custom_routes"] = []interface{}{dynamicRouteServerPeerGroupPolicyCustomRoutePrototypeModel}
	model["type"] = "custom_routes"

	result, err := vpc.ResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToDynamicRouteServerPeerGroupPolicyPrototypeDynamicRouteServerPeerGroupPolicyCustomRoutesPrototype(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToDynamicRouteServerPeerGroupPolicyPrototypeDynamicRouteServerPeerGroupPolicyLearnedRoutesPrototype(t *testing.T) {
	checkResult := func(result *vpcv1.DynamicRouteServerPeerGroupPolicyPrototypeDynamicRouteServerPeerGroupPolicyLearnedRoutesPrototype) {
		dynamicRouteServerPeerGroupPrefixWatcherExcludedPrefixPrototypeModel := new(vpcv1.DynamicRouteServerPeerGroupPrefixWatcherExcludedPrefixPrototype)
		dynamicRouteServerPeerGroupPrefixWatcherExcludedPrefixPrototypeModel.Ge = core.Int64Ptr(int64(28))
		dynamicRouteServerPeerGroupPrefixWatcherExcludedPrefixPrototypeModel.Le = core.Int64Ptr(int64(32))
		dynamicRouteServerPeerGroupPrefixWatcherExcludedPrefixPrototypeModel.Prefix = core.StringPtr("10.0.0.0/24")

		dynamicRouteServerPeerGroupIdentityModel := new(vpcv1.DynamicRouteServerPeerGroupIdentityByID)
		dynamicRouteServerPeerGroupIdentityModel.ID = core.StringPtr("r006-abcdef12-3456-7890-abcd-abcdef123456")

		model := new(vpcv1.DynamicRouteServerPeerGroupPolicyPrototypeDynamicRouteServerPeerGroupPolicyLearnedRoutesPrototype)
		model.Name = core.StringPtr("my-dynamic-route-server-peer-group-policy")
		model.State = core.StringPtr("enabled")
		model.ExcludedPrefixes = []vpcv1.DynamicRouteServerPeerGroupPrefixWatcherExcludedPrefixPrototype{*dynamicRouteServerPeerGroupPrefixWatcherExcludedPrefixPrototypeModel}
		model.PeerGroups = []vpcv1.DynamicRouteServerPeerGroupIdentityIntf{dynamicRouteServerPeerGroupIdentityModel}
		model.Type = core.StringPtr("learned_routes")

		assert.Equal(t, result, model)
	}

	dynamicRouteServerPeerGroupPrefixWatcherExcludedPrefixPrototypeModel := make(map[string]interface{})
	dynamicRouteServerPeerGroupPrefixWatcherExcludedPrefixPrototypeModel["ge"] = int(28)
	dynamicRouteServerPeerGroupPrefixWatcherExcludedPrefixPrototypeModel["le"] = int(32)
	dynamicRouteServerPeerGroupPrefixWatcherExcludedPrefixPrototypeModel["prefix"] = "10.0.0.0/24"

	dynamicRouteServerPeerGroupIdentityModel := make(map[string]interface{})
	dynamicRouteServerPeerGroupIdentityModel["id"] = "r006-abcdef12-3456-7890-abcd-abcdef123456"

	model := make(map[string]interface{})
	model["name"] = "my-dynamic-route-server-peer-group-policy"
	model["state"] = "enabled"
	model["excluded_prefixes"] = []interface{}{dynamicRouteServerPeerGroupPrefixWatcherExcludedPrefixPrototypeModel}
	model["peer_groups"] = []interface{}{dynamicRouteServerPeerGroupIdentityModel}
	model["type"] = "learned_routes"

	result, err := vpc.ResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToDynamicRouteServerPeerGroupPolicyPrototypeDynamicRouteServerPeerGroupPolicyLearnedRoutesPrototype(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToDynamicRouteServerPeerGroupPolicyPrototypeDynamicRouteServerPeerGroupPolicyVPCAddressPrefixesPrototype(t *testing.T) {
	checkResult := func(result *vpcv1.DynamicRouteServerPeerGroupPolicyPrototypeDynamicRouteServerPeerGroupPolicyVPCAddressPrefixesPrototype) {
		model := new(vpcv1.DynamicRouteServerPeerGroupPolicyPrototypeDynamicRouteServerPeerGroupPolicyVPCAddressPrefixesPrototype)
		model.Name = core.StringPtr("my-dynamic-route-server-peer-group-policy")
		model.State = core.StringPtr("enabled")
		model.Type = core.StringPtr("vpc_address_prefixes")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["name"] = "my-dynamic-route-server-peer-group-policy"
	model["state"] = "enabled"
	model["type"] = "vpc_address_prefixes"

	result, err := vpc.ResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToDynamicRouteServerPeerGroupPolicyPrototypeDynamicRouteServerPeerGroupPolicyVPCAddressPrefixesPrototype(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToDynamicRouteServerPeerGroupPolicyPrototypeDynamicRouteServerPeerGroupPolicyVPCRoutingTablesPrototype(t *testing.T) {
	checkResult := func(result *vpcv1.DynamicRouteServerPeerGroupPolicyPrototypeDynamicRouteServerPeerGroupPolicyVPCRoutingTablesPrototype) {
		dynamicRouteServerPeerGroupPrefixWatcherExcludedPrefixPrototypeModel := new(vpcv1.DynamicRouteServerPeerGroupPrefixWatcherExcludedPrefixPrototype)
		dynamicRouteServerPeerGroupPrefixWatcherExcludedPrefixPrototypeModel.Ge = core.Int64Ptr(int64(0))
		dynamicRouteServerPeerGroupPrefixWatcherExcludedPrefixPrototypeModel.Le = core.Int64Ptr(int64(32))
		dynamicRouteServerPeerGroupPrefixWatcherExcludedPrefixPrototypeModel.Prefix = core.StringPtr("testString")

		dynamicRouteServerPeerGroupPolicyNextHopPrototypeModel := new(vpcv1.DynamicRouteServerPeerGroupPolicyNextHopPrototypeDynamicRouteServerPeerGroupIdentityDynamicRouteServerPeerGroupIdentityByID)
		dynamicRouteServerPeerGroupPolicyNextHopPrototypeModel.ID = core.StringPtr("r006-8228af60-189d-11ed-861d-0242ac120002")

		routingTableIdentityModel := new(vpcv1.RoutingTableIdentityByID)
		routingTableIdentityModel.ID = core.StringPtr("r006-6885e83f-03b2-4603-8a86-db2a0f55c840")

		dynamicRouteServerPeerGroupPolicyVPCRoutingTablesRoutingTablePrototypeModel := new(vpcv1.DynamicRouteServerPeerGroupPolicyVPCRoutingTablesRoutingTablePrototype)
		dynamicRouteServerPeerGroupPolicyVPCRoutingTablesRoutingTablePrototypeModel.Advertise = core.BoolPtr(false)
		dynamicRouteServerPeerGroupPolicyVPCRoutingTablesRoutingTablePrototypeModel.VPCRoutingTable = routingTableIdentityModel

		model := new(vpcv1.DynamicRouteServerPeerGroupPolicyPrototypeDynamicRouteServerPeerGroupPolicyVPCRoutingTablesPrototype)
		model.Name = core.StringPtr("my-dynamic-route-server-peer-group-policy")
		model.State = core.StringPtr("enabled")
		model.ExcludedPrefixes = []vpcv1.DynamicRouteServerPeerGroupPrefixWatcherExcludedPrefixPrototype{*dynamicRouteServerPeerGroupPrefixWatcherExcludedPrefixPrototypeModel}
		model.NextHops = []vpcv1.DynamicRouteServerPeerGroupPolicyNextHopPrototypeIntf{dynamicRouteServerPeerGroupPolicyNextHopPrototypeModel}
		model.RouteDeleteDelay = core.Int64Ptr(int64(0))
		model.RoutingTables = []vpcv1.DynamicRouteServerPeerGroupPolicyVPCRoutingTablesRoutingTablePrototype{*dynamicRouteServerPeerGroupPolicyVPCRoutingTablesRoutingTablePrototypeModel}
		model.Type = core.StringPtr("vpc_routing_tables")

		assert.Equal(t, result, model)
	}

	dynamicRouteServerPeerGroupPrefixWatcherExcludedPrefixPrototypeModel := make(map[string]interface{})
	dynamicRouteServerPeerGroupPrefixWatcherExcludedPrefixPrototypeModel["ge"] = int(0)
	dynamicRouteServerPeerGroupPrefixWatcherExcludedPrefixPrototypeModel["le"] = int(32)
	dynamicRouteServerPeerGroupPrefixWatcherExcludedPrefixPrototypeModel["prefix"] = "testString"

	dynamicRouteServerPeerGroupPolicyNextHopPrototypeModel := make(map[string]interface{})
	dynamicRouteServerPeerGroupPolicyNextHopPrototypeModel["id"] = "r006-8228af60-189d-11ed-861d-0242ac120002"

	routingTableIdentityModel := make(map[string]interface{})
	routingTableIdentityModel["id"] = "r006-6885e83f-03b2-4603-8a86-db2a0f55c840"

	dynamicRouteServerPeerGroupPolicyVPCRoutingTablesRoutingTablePrototypeModel := make(map[string]interface{})
	dynamicRouteServerPeerGroupPolicyVPCRoutingTablesRoutingTablePrototypeModel["advertise"] = false
	dynamicRouteServerPeerGroupPolicyVPCRoutingTablesRoutingTablePrototypeModel["vpc_routing_table"] = []interface{}{routingTableIdentityModel}

	model := make(map[string]interface{})
	model["name"] = "my-dynamic-route-server-peer-group-policy"
	model["state"] = "enabled"
	model["excluded_prefixes"] = []interface{}{dynamicRouteServerPeerGroupPrefixWatcherExcludedPrefixPrototypeModel}
	model["next_hops"] = []interface{}{dynamicRouteServerPeerGroupPolicyNextHopPrototypeModel}
	model["route_delete_delay"] = int(0)
	model["routing_tables"] = []interface{}{dynamicRouteServerPeerGroupPolicyVPCRoutingTablesRoutingTablePrototypeModel}
	model["type"] = "vpc_routing_tables"

	result, err := vpc.ResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToDynamicRouteServerPeerGroupPolicyPrototypeDynamicRouteServerPeerGroupPolicyVPCRoutingTablesPrototype(model)
	assert.Nil(t, err)
	checkResult(result)
}
