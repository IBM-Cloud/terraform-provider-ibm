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

func TestAccIBMIsDynamicRouteServerStaticRouteDataSourceBasic(t *testing.T) {
	dynamicRouteServerStaticRouteDynamicRouteServerID := fmt.Sprintf("tf_dynamic_route_server_id_%d", acctest.RandIntRange(10, 100))
	dynamicRouteServerStaticRouteDestination := fmt.Sprintf("tf_destination_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsDynamicRouteServerStaticRouteDataSourceConfigBasic(dynamicRouteServerStaticRouteDynamicRouteServerID, dynamicRouteServerStaticRouteDestination),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_static_route.is_dynamic_route_server_static_route_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_static_route.is_dynamic_route_server_static_route_instance", "dynamic_route_server_id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_static_route.is_dynamic_route_server_static_route_instance", "is_dynamic_route_server_static_route_id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_static_route.is_dynamic_route_server_static_route_instance", "added_routes.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_static_route.is_dynamic_route_server_static_route_instance", "destination"),
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_static_route.is_dynamic_route_server_static_route_instance", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_static_route.is_dynamic_route_server_static_route_instance", "lifecycle_reasons.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_static_route.is_dynamic_route_server_static_route_instance", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_static_route.is_dynamic_route_server_static_route_instance", "next_hops.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_static_route.is_dynamic_route_server_static_route_instance", "resource_type"),
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_static_route.is_dynamic_route_server_static_route_instance", "route_delete_delay"),
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_static_route.is_dynamic_route_server_static_route_instance", "routing_tables.#"),
				),
			},
		},
	})
}

func TestAccIBMIsDynamicRouteServerStaticRouteDataSourceAllArgs(t *testing.T) {
	dynamicRouteServerStaticRouteDynamicRouteServerID := fmt.Sprintf("tf_dynamic_route_server_id_%d", acctest.RandIntRange(10, 100))
	dynamicRouteServerStaticRouteDestination := fmt.Sprintf("tf_destination_%d", acctest.RandIntRange(10, 100))
	dynamicRouteServerStaticRouteName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	dynamicRouteServerStaticRouteRouteDeleteDelay := fmt.Sprintf("%d", acctest.RandIntRange(0, 60))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsDynamicRouteServerStaticRouteDataSourceConfig(dynamicRouteServerStaticRouteDynamicRouteServerID, dynamicRouteServerStaticRouteDestination, dynamicRouteServerStaticRouteName, dynamicRouteServerStaticRouteRouteDeleteDelay),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_static_route.is_dynamic_route_server_static_route_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_static_route.is_dynamic_route_server_static_route_instance", "dynamic_route_server_id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_static_route.is_dynamic_route_server_static_route_instance", "is_dynamic_route_server_static_route_id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_static_route.is_dynamic_route_server_static_route_instance", "added_routes.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_static_route.is_dynamic_route_server_static_route_instance", "destination"),
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_static_route.is_dynamic_route_server_static_route_instance", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_static_route.is_dynamic_route_server_static_route_instance", "lifecycle_reasons.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_static_route.is_dynamic_route_server_static_route_instance", "lifecycle_reasons.0.code"),
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_static_route.is_dynamic_route_server_static_route_instance", "lifecycle_reasons.0.message"),
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_static_route.is_dynamic_route_server_static_route_instance", "lifecycle_reasons.0.more_info"),
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_static_route.is_dynamic_route_server_static_route_instance", "lifecycle_state"),
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_static_route.is_dynamic_route_server_static_route_instance", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_static_route.is_dynamic_route_server_static_route_instance", "next_hops.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_static_route.is_dynamic_route_server_static_route_instance", "next_hops.0.href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_static_route.is_dynamic_route_server_static_route_instance", "next_hops.0.id"),
					resource.TestCheckResourceAttr("data.ibm_is_dynamic_route_server_static_route.is_dynamic_route_server_static_route_instance", "next_hops.0.name", dynamicRouteServerStaticRouteName),
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_static_route.is_dynamic_route_server_static_route_instance", "next_hops.0.resource_type"),
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_static_route.is_dynamic_route_server_static_route_instance", "resource_type"),
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_static_route.is_dynamic_route_server_static_route_instance", "route_delete_delay"),
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_static_route.is_dynamic_route_server_static_route_instance", "routing_tables.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_static_route.is_dynamic_route_server_static_route_instance", "routing_tables.0.advertise"),
				),
			},
		},
	})
}

func testAccCheckIBMIsDynamicRouteServerStaticRouteDataSourceConfigBasic(dynamicRouteServerStaticRouteDynamicRouteServerID string, dynamicRouteServerStaticRouteDestination string) string {
	return fmt.Sprintf(`
		resource "ibm_is_dynamic_route_server_static_route" "is_dynamic_route_server_static_route_instance" {
			dynamic_route_server_id = "%s"
			destination = "%s"
			next_hops {
				deleted {
					more_info = "https://cloud.ibm.com/apidocs/vpc#deleted-resources"
				}
				href = "https://us-south.iaas.cloud.ibm.com/v1/dynamic_route_servers/r006-d7cc5196-9864-48c4-82d8-3f30da41fcc5/peer_groups/r006-8228af60-189d-11ed-861d-0242ac120002"
				id = "r006-8228af60-189d-11ed-861d-0242ac120004"
				name = "my-dynamic-route-server-peer-group"
				resource_type = "dynamic_route_server_peer_group"
			}
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

		data "ibm_is_dynamic_route_server_static_route" "is_dynamic_route_server_static_route_instance" {
			dynamic_route_server_id = ibm_is_dynamic_route_server_static_route.is_dynamic_route_server_static_route_instance.dynamic_route_server_id
			is_dynamic_route_server_static_route_id = ibm_is_dynamic_route_server_static_route.is_dynamic_route_server_static_route_instance.is_dynamic_route_server_static_route_id
		}
	`, dynamicRouteServerStaticRouteDynamicRouteServerID, dynamicRouteServerStaticRouteDestination)
}

func testAccCheckIBMIsDynamicRouteServerStaticRouteDataSourceConfig(dynamicRouteServerStaticRouteDynamicRouteServerID string, dynamicRouteServerStaticRouteDestination string, dynamicRouteServerStaticRouteName string, dynamicRouteServerStaticRouteRouteDeleteDelay string) string {
	return fmt.Sprintf(`
		resource "ibm_is_dynamic_route_server_static_route" "is_dynamic_route_server_static_route_instance" {
			dynamic_route_server_id = "%s"
			destination = "%s"
			name = "%s"
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

		data "ibm_is_dynamic_route_server_static_route" "is_dynamic_route_server_static_route_instance" {
			dynamic_route_server_id = ibm_is_dynamic_route_server_static_route.is_dynamic_route_server_static_route_instance.dynamic_route_server_id
			is_dynamic_route_server_static_route_id = ibm_is_dynamic_route_server_static_route.is_dynamic_route_server_static_route_instance.is_dynamic_route_server_static_route_id
		}
	`, dynamicRouteServerStaticRouteDynamicRouteServerID, dynamicRouteServerStaticRouteDestination, dynamicRouteServerStaticRouteName, dynamicRouteServerStaticRouteRouteDeleteDelay)
}

func TestDataSourceIBMIsDynamicRouteServerStaticRouteDynamicRouteServerStaticRouteAddedRouteToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		deletedModel := make(map[string]interface{})
		deletedModel["more_info"] = "https://cloud.ibm.com/apidocs/vpc#deleted-resources"

		routeReferenceModel := make(map[string]interface{})
		routeReferenceModel["deleted"] = []map[string]interface{}{deletedModel}
		routeReferenceModel["href"] = "https://us-south.iaas.cloud.ibm.com/v1/vpcs/r006-4727d842-f94f-4a2d-824a-9bc9b02c523b/routing_tables/r006-6885e83f-03b2-4603-8a86-db2a0f55c840/routes/r006-67b7b783-9a0f-41c1-a7f7-eccff87fb8f1"
		routeReferenceModel["id"] = "r006-67b7b783-9a0f-41c1-a7f7-eccff87fb8f1"
		routeReferenceModel["name"] = "my-vpc-routing-table-route"

		model := make(map[string]interface{})
		model["vpc_routing_table_route"] = []map[string]interface{}{routeReferenceModel}

		assert.Equal(t, result, model)
	}

	deletedModel := new(vpcv1.Deleted)
	deletedModel.MoreInfo = core.StringPtr("https://cloud.ibm.com/apidocs/vpc#deleted-resources")

	routeReferenceModel := new(vpcv1.RouteReference)
	routeReferenceModel.Deleted = deletedModel
	routeReferenceModel.Href = core.StringPtr("https://us-south.iaas.cloud.ibm.com/v1/vpcs/r006-4727d842-f94f-4a2d-824a-9bc9b02c523b/routing_tables/r006-6885e83f-03b2-4603-8a86-db2a0f55c840/routes/r006-67b7b783-9a0f-41c1-a7f7-eccff87fb8f1")
	routeReferenceModel.ID = core.StringPtr("r006-67b7b783-9a0f-41c1-a7f7-eccff87fb8f1")
	routeReferenceModel.Name = core.StringPtr("my-vpc-routing-table-route")

	model := new(vpcv1.DynamicRouteServerStaticRouteAddedRoute)
	model.VPCRoutingTableRoute = routeReferenceModel

	result, err := vpc.DataSourceIBMIsDynamicRouteServerStaticRouteDynamicRouteServerStaticRouteAddedRouteToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMIsDynamicRouteServerStaticRouteRouteReferenceToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		deletedModel := make(map[string]interface{})
		deletedModel["more_info"] = "https://cloud.ibm.com/apidocs/vpc#deleted-resources"

		model := make(map[string]interface{})
		model["deleted"] = []map[string]interface{}{deletedModel}
		model["href"] = "https://us-south.iaas.cloud.ibm.com/v1/vpcs/r006-4727d842-f94f-4a2d-824a-9bc9b02c523b/routing_tables/r006-6885e83f-03b2-4603-8a86-db2a0f55c840/routes/r006-67b7b783-9a0f-41c1-a7f7-eccff87fb8f1"
		model["id"] = "r006-67b7b783-9a0f-41c1-a7f7-eccff87fb8f1"
		model["name"] = "my-vpc-routing-table-route"

		assert.Equal(t, result, model)
	}

	deletedModel := new(vpcv1.Deleted)
	deletedModel.MoreInfo = core.StringPtr("https://cloud.ibm.com/apidocs/vpc#deleted-resources")

	model := new(vpcv1.RouteReference)
	model.Deleted = deletedModel
	model.Href = core.StringPtr("https://us-south.iaas.cloud.ibm.com/v1/vpcs/r006-4727d842-f94f-4a2d-824a-9bc9b02c523b/routing_tables/r006-6885e83f-03b2-4603-8a86-db2a0f55c840/routes/r006-67b7b783-9a0f-41c1-a7f7-eccff87fb8f1")
	model.ID = core.StringPtr("r006-67b7b783-9a0f-41c1-a7f7-eccff87fb8f1")
	model.Name = core.StringPtr("my-vpc-routing-table-route")

	result, err := vpc.DataSourceIBMIsDynamicRouteServerStaticRouteRouteReferenceToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMIsDynamicRouteServerStaticRouteDeletedToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["more_info"] = "https://cloud.ibm.com/apidocs/vpc#deleted-resources"

		assert.Equal(t, result, model)
	}

	model := new(vpcv1.Deleted)
	model.MoreInfo = core.StringPtr("https://cloud.ibm.com/apidocs/vpc#deleted-resources")

	result, err := vpc.DataSourceIBMIsDynamicRouteServerStaticRouteDeletedToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMIsDynamicRouteServerStaticRouteLifecycleReasonToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["code"] = "resource_suspended_by_provider"
		model["message"] = "The resource has been suspended. Contact IBM support with the CRN for next steps."
		model["more_info"] = "https://cloud.ibm.com/apidocs/vpc#resource-suspension"

		assert.Equal(t, result, model)
	}

	model := new(vpcv1.LifecycleReason)
	model.Code = core.StringPtr("resource_suspended_by_provider")
	model.Message = core.StringPtr("The resource has been suspended. Contact IBM support with the CRN for next steps.")
	model.MoreInfo = core.StringPtr("https://cloud.ibm.com/apidocs/vpc#resource-suspension")

	result, err := vpc.DataSourceIBMIsDynamicRouteServerStaticRouteLifecycleReasonToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMIsDynamicRouteServerStaticRouteDynamicRouteServerStaticRouteNextHopToMap(t *testing.T) {
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

	model := new(vpcv1.DynamicRouteServerStaticRouteNextHop)
	model.Deleted = deletedModel
	model.Href = core.StringPtr("https://us-south.iaas.cloud.ibm.com/v1/dynamic_route_servers/r006-d7cc5196-9864-48c4-82d8-3f30da41fcc5/peer_groups/r006-8228af60-189d-11ed-861d-0242ac120002")
	model.ID = core.StringPtr("r006-8228af60-189d-11ed-861d-0242ac120004")
	model.Name = core.StringPtr("my-dynamic-route-server-peer-group")
	model.ResourceType = core.StringPtr("dynamic_route_server_peer_group")

	result, err := vpc.DataSourceIBMIsDynamicRouteServerStaticRouteDynamicRouteServerStaticRouteNextHopToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMIsDynamicRouteServerStaticRouteDynamicRouteServerStaticRouteNextHopDynamicRouteServerPeerGroupReferenceToMap(t *testing.T) {
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

	model := new(vpcv1.DynamicRouteServerStaticRouteNextHopDynamicRouteServerPeerGroupReference)
	model.Deleted = deletedModel
	model.Href = core.StringPtr("https://us-south.iaas.cloud.ibm.com/v1/dynamic_route_servers/r006-d7cc5196-9864-48c4-82d8-3f30da41fcc5/peer_groups/r006-8228af60-189d-11ed-861d-0242ac120002")
	model.ID = core.StringPtr("r006-8228af60-189d-11ed-861d-0242ac120004")
	model.Name = core.StringPtr("my-dynamic-route-server-peer-group")
	model.ResourceType = core.StringPtr("dynamic_route_server_peer_group")

	result, err := vpc.DataSourceIBMIsDynamicRouteServerStaticRouteDynamicRouteServerStaticRouteNextHopDynamicRouteServerPeerGroupReferenceToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMIsDynamicRouteServerStaticRouteDynamicRouteServerStaticRouteRoutingTableToMap(t *testing.T) {
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

	model := new(vpcv1.DynamicRouteServerStaticRouteRoutingTable)
	model.Advertise = core.BoolPtr(true)
	model.VPCRoutingTable = routingTableReferenceModel

	result, err := vpc.DataSourceIBMIsDynamicRouteServerStaticRouteDynamicRouteServerStaticRouteRoutingTableToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMIsDynamicRouteServerStaticRouteRoutingTableReferenceToMap(t *testing.T) {
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

	result, err := vpc.DataSourceIBMIsDynamicRouteServerStaticRouteRoutingTableReferenceToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}
