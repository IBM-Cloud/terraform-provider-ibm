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

func TestAccIBMIsDynamicRouteServerStaticRouteBasic(t *testing.T) {
	var conf vpcv1.DynamicRouteServerStaticRoute
	dynamicRouteServerID := fmt.Sprintf("tf_dynamic_route_server_id_%d", acctest.RandIntRange(10, 100))
	destination := fmt.Sprintf("tf_destination_%d", acctest.RandIntRange(10, 100))
	destinationUpdate := fmt.Sprintf("tf_destination_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIsDynamicRouteServerStaticRouteDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsDynamicRouteServerStaticRouteConfigBasic(dynamicRouteServerID, destination),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIsDynamicRouteServerStaticRouteExists("ibm_is_dynamic_route_server_static_route.is_dynamic_route_server_static_route_instance", conf),
					resource.TestCheckResourceAttr("ibm_is_dynamic_route_server_static_route.is_dynamic_route_server_static_route_instance", "dynamic_route_server_id", dynamicRouteServerID),
					resource.TestCheckResourceAttr("ibm_is_dynamic_route_server_static_route.is_dynamic_route_server_static_route_instance", "destination", destination),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMIsDynamicRouteServerStaticRouteConfigBasic(dynamicRouteServerID, destinationUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_is_dynamic_route_server_static_route.is_dynamic_route_server_static_route_instance", "dynamic_route_server_id", dynamicRouteServerID),
					resource.TestCheckResourceAttr("ibm_is_dynamic_route_server_static_route.is_dynamic_route_server_static_route_instance", "destination", destinationUpdate),
				),
			},
		},
	})
}

func TestAccIBMIsDynamicRouteServerStaticRouteAllArgs(t *testing.T) {
	var conf vpcv1.DynamicRouteServerStaticRoute
	dynamicRouteServerID := fmt.Sprintf("tf_dynamic_route_server_id_%d", acctest.RandIntRange(10, 100))
	destination := fmt.Sprintf("tf_destination_%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	routeDeleteDelay := fmt.Sprintf("%d", acctest.RandIntRange(0, 60))
	destinationUpdate := fmt.Sprintf("tf_destination_%d", acctest.RandIntRange(10, 100))
	nameUpdate := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	routeDeleteDelayUpdate := fmt.Sprintf("%d", acctest.RandIntRange(0, 60))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIsDynamicRouteServerStaticRouteDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsDynamicRouteServerStaticRouteConfig(dynamicRouteServerID, destination, name, routeDeleteDelay),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIsDynamicRouteServerStaticRouteExists("ibm_is_dynamic_route_server_static_route.is_dynamic_route_server_static_route_instance", conf),
					resource.TestCheckResourceAttr("ibm_is_dynamic_route_server_static_route.is_dynamic_route_server_static_route_instance", "dynamic_route_server_id", dynamicRouteServerID),
					resource.TestCheckResourceAttr("ibm_is_dynamic_route_server_static_route.is_dynamic_route_server_static_route_instance", "destination", destination),
					resource.TestCheckResourceAttr("ibm_is_dynamic_route_server_static_route.is_dynamic_route_server_static_route_instance", "name", name),
					resource.TestCheckResourceAttr("ibm_is_dynamic_route_server_static_route.is_dynamic_route_server_static_route_instance", "route_delete_delay", routeDeleteDelay),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMIsDynamicRouteServerStaticRouteConfig(dynamicRouteServerID, destinationUpdate, nameUpdate, routeDeleteDelayUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_is_dynamic_route_server_static_route.is_dynamic_route_server_static_route_instance", "dynamic_route_server_id", dynamicRouteServerID),
					resource.TestCheckResourceAttr("ibm_is_dynamic_route_server_static_route.is_dynamic_route_server_static_route_instance", "destination", destinationUpdate),
					resource.TestCheckResourceAttr("ibm_is_dynamic_route_server_static_route.is_dynamic_route_server_static_route_instance", "name", nameUpdate),
					resource.TestCheckResourceAttr("ibm_is_dynamic_route_server_static_route.is_dynamic_route_server_static_route_instance", "route_delete_delay", routeDeleteDelayUpdate),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_is_dynamic_route_server_static_route.is_dynamic_route_server_static_route_instance",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIBMIsDynamicRouteServerStaticRouteConfigBasic(dynamicRouteServerID string, destination string) string {
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
	`, dynamicRouteServerID, destination)
}

func testAccCheckIBMIsDynamicRouteServerStaticRouteConfig(dynamicRouteServerID string, destination string, name string, routeDeleteDelay string) string {
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
	`, dynamicRouteServerID, destination, name, routeDeleteDelay)
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

		getDynamicRouteServerStaticRouteOptions := &vpcv1.GetDynamicRouteServerStaticRouteOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getDynamicRouteServerStaticRouteOptions.SetDynamicRouteServerID(parts[0])
		getDynamicRouteServerStaticRouteOptions.SetID(parts[1])

		dynamicRouteServerStaticRoute, _, err := vpcClient.GetDynamicRouteServerStaticRoute(getDynamicRouteServerStaticRouteOptions)
		if err != nil {
			return err
		}

		obj = *dynamicRouteServerStaticRoute
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

		getDynamicRouteServerStaticRouteOptions := &vpcv1.GetDynamicRouteServerStaticRouteOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getDynamicRouteServerStaticRouteOptions.SetDynamicRouteServerID(parts[0])
		getDynamicRouteServerStaticRouteOptions.SetID(parts[1])

		// Try to find the key
		_, response, err := vpcClient.GetDynamicRouteServerStaticRoute(getDynamicRouteServerStaticRouteOptions)

		if err == nil {
			return fmt.Errorf("DynamicRouteServerStaticRoute still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for DynamicRouteServerStaticRoute (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}

func TestResourceIBMIsDynamicRouteServerStaticRouteDynamicRouteServerStaticRouteNextHopToMap(t *testing.T) {
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

	result, err := vpc.ResourceIBMIsDynamicRouteServerStaticRouteDynamicRouteServerStaticRouteNextHopToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMIsDynamicRouteServerStaticRouteDeletedToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["more_info"] = "https://cloud.ibm.com/apidocs/vpc#deleted-resources"

		assert.Equal(t, result, model)
	}

	model := new(vpcv1.Deleted)
	model.MoreInfo = core.StringPtr("https://cloud.ibm.com/apidocs/vpc#deleted-resources")

	result, err := vpc.ResourceIBMIsDynamicRouteServerStaticRouteDeletedToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMIsDynamicRouteServerStaticRouteDynamicRouteServerStaticRouteNextHopDynamicRouteServerPeerGroupReferenceToMap(t *testing.T) {
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

	result, err := vpc.ResourceIBMIsDynamicRouteServerStaticRouteDynamicRouteServerStaticRouteNextHopDynamicRouteServerPeerGroupReferenceToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMIsDynamicRouteServerStaticRouteDynamicRouteServerStaticRouteRoutingTableToMap(t *testing.T) {
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

	result, err := vpc.ResourceIBMIsDynamicRouteServerStaticRouteDynamicRouteServerStaticRouteRoutingTableToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMIsDynamicRouteServerStaticRouteRoutingTableReferenceToMap(t *testing.T) {
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

	result, err := vpc.ResourceIBMIsDynamicRouteServerStaticRouteRoutingTableReferenceToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMIsDynamicRouteServerStaticRouteDynamicRouteServerStaticRouteAddedRouteToMap(t *testing.T) {
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

	result, err := vpc.ResourceIBMIsDynamicRouteServerStaticRouteDynamicRouteServerStaticRouteAddedRouteToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMIsDynamicRouteServerStaticRouteRouteReferenceToMap(t *testing.T) {
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

	result, err := vpc.ResourceIBMIsDynamicRouteServerStaticRouteRouteReferenceToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMIsDynamicRouteServerStaticRouteLifecycleReasonToMap(t *testing.T) {
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

	result, err := vpc.ResourceIBMIsDynamicRouteServerStaticRouteLifecycleReasonToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMIsDynamicRouteServerStaticRouteMapToDynamicRouteServerStaticRouteNextHopPrototype(t *testing.T) {
	checkResult := func(result vpcv1.DynamicRouteServerStaticRouteNextHopPrototypeIntf) {
		model := new(vpcv1.DynamicRouteServerStaticRouteNextHopPrototype)
		model.ID = core.StringPtr("r006-8228af60-189d-11ed-861d-0242ac120002")
		model.Href = core.StringPtr("https://us-south.iaas.cloud.ibm.com/v1/dynamic_route_servers/r006-d7cc5196-9864-48c4-82d8-3f30da41fcc5/peer_groups/r006-8228af60-189d-11ed-861d-0242ac120002")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["id"] = "r006-8228af60-189d-11ed-861d-0242ac120002"
	model["href"] = "https://us-south.iaas.cloud.ibm.com/v1/dynamic_route_servers/r006-d7cc5196-9864-48c4-82d8-3f30da41fcc5/peer_groups/r006-8228af60-189d-11ed-861d-0242ac120002"

	result, err := vpc.ResourceIBMIsDynamicRouteServerStaticRouteMapToDynamicRouteServerStaticRouteNextHopPrototype(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMIsDynamicRouteServerStaticRouteMapToDynamicRouteServerStaticRouteNextHopPrototypeDynamicRouteServerPeerGroupIdentity(t *testing.T) {
	checkResult := func(result vpcv1.DynamicRouteServerStaticRouteNextHopPrototypeDynamicRouteServerPeerGroupIdentityIntf) {
		model := new(vpcv1.DynamicRouteServerStaticRouteNextHopPrototypeDynamicRouteServerPeerGroupIdentity)
		model.ID = core.StringPtr("r006-8228af60-189d-11ed-861d-0242ac120002")
		model.Href = core.StringPtr("https://us-south.iaas.cloud.ibm.com/v1/dynamic_route_servers/r006-d7cc5196-9864-48c4-82d8-3f30da41fcc5/peer_groups/r006-8228af60-189d-11ed-861d-0242ac120002")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["id"] = "r006-8228af60-189d-11ed-861d-0242ac120002"
	model["href"] = "https://us-south.iaas.cloud.ibm.com/v1/dynamic_route_servers/r006-d7cc5196-9864-48c4-82d8-3f30da41fcc5/peer_groups/r006-8228af60-189d-11ed-861d-0242ac120002"

	result, err := vpc.ResourceIBMIsDynamicRouteServerStaticRouteMapToDynamicRouteServerStaticRouteNextHopPrototypeDynamicRouteServerPeerGroupIdentity(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMIsDynamicRouteServerStaticRouteMapToDynamicRouteServerStaticRouteNextHopPrototypeDynamicRouteServerPeerGroupIdentityDynamicRouteServerPeerGroupIdentityByID(t *testing.T) {
	checkResult := func(result *vpcv1.DynamicRouteServerStaticRouteNextHopPrototypeDynamicRouteServerPeerGroupIdentityDynamicRouteServerPeerGroupIdentityByID) {
		model := new(vpcv1.DynamicRouteServerStaticRouteNextHopPrototypeDynamicRouteServerPeerGroupIdentityDynamicRouteServerPeerGroupIdentityByID)
		model.ID = core.StringPtr("r006-8228af60-189d-11ed-861d-0242ac120002")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["id"] = "r006-8228af60-189d-11ed-861d-0242ac120002"

	result, err := vpc.ResourceIBMIsDynamicRouteServerStaticRouteMapToDynamicRouteServerStaticRouteNextHopPrototypeDynamicRouteServerPeerGroupIdentityDynamicRouteServerPeerGroupIdentityByID(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMIsDynamicRouteServerStaticRouteMapToDynamicRouteServerStaticRouteNextHopPrototypeDynamicRouteServerPeerGroupIdentityDynamicRouteServerPeerGroupIdentityByHref(t *testing.T) {
	checkResult := func(result *vpcv1.DynamicRouteServerStaticRouteNextHopPrototypeDynamicRouteServerPeerGroupIdentityDynamicRouteServerPeerGroupIdentityByHref) {
		model := new(vpcv1.DynamicRouteServerStaticRouteNextHopPrototypeDynamicRouteServerPeerGroupIdentityDynamicRouteServerPeerGroupIdentityByHref)
		model.Href = core.StringPtr("https://us-south.iaas.cloud.ibm.com/v1/dynamic_route_servers/r006-d7cc5196-9864-48c4-82d8-3f30da41fcc5/peer_groups/r006-8228af60-189d-11ed-861d-0242ac120002")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["href"] = "https://us-south.iaas.cloud.ibm.com/v1/dynamic_route_servers/r006-d7cc5196-9864-48c4-82d8-3f30da41fcc5/peer_groups/r006-8228af60-189d-11ed-861d-0242ac120002"

	result, err := vpc.ResourceIBMIsDynamicRouteServerStaticRouteMapToDynamicRouteServerStaticRouteNextHopPrototypeDynamicRouteServerPeerGroupIdentityDynamicRouteServerPeerGroupIdentityByHref(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMIsDynamicRouteServerStaticRouteMapToDynamicRouteServerStaticRouteRoutingTablePrototype(t *testing.T) {
	checkResult := func(result *vpcv1.DynamicRouteServerStaticRouteRoutingTablePrototype) {
		routingTableIdentityModel := new(vpcv1.RoutingTableIdentityByID)
		routingTableIdentityModel.ID = core.StringPtr("r006-6885e83f-03b2-4603-8a86-db2a0f55c840")

		model := new(vpcv1.DynamicRouteServerStaticRouteRoutingTablePrototype)
		model.Advertise = core.BoolPtr(false)
		model.VPCRoutingTable = routingTableIdentityModel

		assert.Equal(t, result, model)
	}

	routingTableIdentityModel := make(map[string]interface{})
	routingTableIdentityModel["id"] = "r006-6885e83f-03b2-4603-8a86-db2a0f55c840"

	model := make(map[string]interface{})
	model["advertise"] = false
	model["vpc_routing_table"] = []interface{}{routingTableIdentityModel}

	result, err := vpc.ResourceIBMIsDynamicRouteServerStaticRouteMapToDynamicRouteServerStaticRouteRoutingTablePrototype(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMIsDynamicRouteServerStaticRouteMapToRoutingTableIdentity(t *testing.T) {
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

	result, err := vpc.ResourceIBMIsDynamicRouteServerStaticRouteMapToRoutingTableIdentity(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMIsDynamicRouteServerStaticRouteMapToRoutingTableIdentityByID(t *testing.T) {
	checkResult := func(result *vpcv1.RoutingTableIdentityByID) {
		model := new(vpcv1.RoutingTableIdentityByID)
		model.ID = core.StringPtr("r006-6885e83f-03b2-4603-8a86-db2a0f55c840")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["id"] = "r006-6885e83f-03b2-4603-8a86-db2a0f55c840"

	result, err := vpc.ResourceIBMIsDynamicRouteServerStaticRouteMapToRoutingTableIdentityByID(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMIsDynamicRouteServerStaticRouteMapToRoutingTableIdentityByCRN(t *testing.T) {
	checkResult := func(result *vpcv1.RoutingTableIdentityByCRN) {
		model := new(vpcv1.RoutingTableIdentityByCRN)
		model.CRN = core.StringPtr("crn:v1:bluemix:public:is:us-south:a/aa2432b1fa4d4ace891e9b80fc104e34::vpc-routing-table:r006-4727d842-f94f-4a2d-824a-9bc9b02c523b/r006-6885e83f-03b2-4603-8a86-db2a0f55c840")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["crn"] = "crn:v1:bluemix:public:is:us-south:a/aa2432b1fa4d4ace891e9b80fc104e34::vpc-routing-table:r006-4727d842-f94f-4a2d-824a-9bc9b02c523b/r006-6885e83f-03b2-4603-8a86-db2a0f55c840"

	result, err := vpc.ResourceIBMIsDynamicRouteServerStaticRouteMapToRoutingTableIdentityByCRN(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMIsDynamicRouteServerStaticRouteMapToRoutingTableIdentityByHref(t *testing.T) {
	checkResult := func(result *vpcv1.RoutingTableIdentityByHref) {
		model := new(vpcv1.RoutingTableIdentityByHref)
		model.Href = core.StringPtr("https://us-south.iaas.cloud.ibm.com/v1/vpcs/r006-4727d842-f94f-4a2d-824a-9bc9b02c523b/routing_tables/r006-6885e83f-03b2-4603-8a86-db2a0f55c840")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["href"] = "https://us-south.iaas.cloud.ibm.com/v1/vpcs/r006-4727d842-f94f-4a2d-824a-9bc9b02c523b/routing_tables/r006-6885e83f-03b2-4603-8a86-db2a0f55c840"

	result, err := vpc.ResourceIBMIsDynamicRouteServerStaticRouteMapToRoutingTableIdentityByHref(model)
	assert.Nil(t, err)
	checkResult(result)
}
