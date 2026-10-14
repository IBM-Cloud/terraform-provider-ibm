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
	. "github.com/IBM-Cloud/terraform-provider-ibm/ibm/unittest"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/stretchr/testify/assert"
)

func TestAccIBMIsDynamicRouteServerPeerGroupPeerBasic(t *testing.T) {
	var conf vpcv1.DynamicRouteServerPeerGroupPeer
	dynamicRouteServerID := fmt.Sprintf("tf_dynamic_route_server_id_%d", acctest.RandIntRange(10, 100))
	dynamicRouteServerPeerGroupID := fmt.Sprintf("tf_dynamic_route_server_peer_group_id_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIsDynamicRouteServerPeerGroupPeerDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsDynamicRouteServerPeerGroupPeerConfigBasic(dynamicRouteServerID, dynamicRouteServerPeerGroupID),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIsDynamicRouteServerPeerGroupPeerExists("ibm_is_dynamic_route_server_peer_group_peer.is_dynamic_route_server_peer_group_peer_instance", conf),
					resource.TestCheckResourceAttr("ibm_is_dynamic_route_server_peer_group_peer.is_dynamic_route_server_peer_group_peer_instance", "dynamic_route_server_id", dynamicRouteServerID),
					resource.TestCheckResourceAttr("ibm_is_dynamic_route_server_peer_group_peer.is_dynamic_route_server_peer_group_peer_instance", "dynamic_route_server_peer_group_id", dynamicRouteServerPeerGroupID),
				),
			},
		},
	})
}

func TestAccIBMIsDynamicRouteServerPeerGroupPeerAllArgs(t *testing.T) {
	var conf vpcv1.DynamicRouteServerPeerGroupPeer
	dynamicRouteServerID := fmt.Sprintf("tf_dynamic_route_server_id_%d", acctest.RandIntRange(10, 100))
	dynamicRouteServerPeerGroupID := fmt.Sprintf("tf_dynamic_route_server_peer_group_id_%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	asn := fmt.Sprintf("%d", acctest.RandIntRange(10, 100))
	priority := fmt.Sprintf("%d", acctest.RandIntRange(0, 4))
	state := "disabled"
	nameUpdate := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	asnUpdate := fmt.Sprintf("%d", acctest.RandIntRange(10, 100))
	priorityUpdate := fmt.Sprintf("%d", acctest.RandIntRange(0, 4))
	stateUpdate := "enabled"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIsDynamicRouteServerPeerGroupPeerDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsDynamicRouteServerPeerGroupPeerConfig(dynamicRouteServerID, dynamicRouteServerPeerGroupID, name, asn, priority, state),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIsDynamicRouteServerPeerGroupPeerExists("ibm_is_dynamic_route_server_peer_group_peer.is_dynamic_route_server_peer_group_peer_instance", conf),
					resource.TestCheckResourceAttr("ibm_is_dynamic_route_server_peer_group_peer.is_dynamic_route_server_peer_group_peer_instance", "dynamic_route_server_id", dynamicRouteServerID),
					resource.TestCheckResourceAttr("ibm_is_dynamic_route_server_peer_group_peer.is_dynamic_route_server_peer_group_peer_instance", "dynamic_route_server_peer_group_id", dynamicRouteServerPeerGroupID),
					resource.TestCheckResourceAttr("ibm_is_dynamic_route_server_peer_group_peer.is_dynamic_route_server_peer_group_peer_instance", "name", name),
					resource.TestCheckResourceAttr("ibm_is_dynamic_route_server_peer_group_peer.is_dynamic_route_server_peer_group_peer_instance", "asn", asn),
					resource.TestCheckResourceAttr("ibm_is_dynamic_route_server_peer_group_peer.is_dynamic_route_server_peer_group_peer_instance", "priority", priority),
					resource.TestCheckResourceAttr("ibm_is_dynamic_route_server_peer_group_peer.is_dynamic_route_server_peer_group_peer_instance", "state", state),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMIsDynamicRouteServerPeerGroupPeerConfig(dynamicRouteServerID, dynamicRouteServerPeerGroupID, nameUpdate, asnUpdate, priorityUpdate, stateUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_is_dynamic_route_server_peer_group_peer.is_dynamic_route_server_peer_group_peer_instance", "dynamic_route_server_id", dynamicRouteServerID),
					resource.TestCheckResourceAttr("ibm_is_dynamic_route_server_peer_group_peer.is_dynamic_route_server_peer_group_peer_instance", "dynamic_route_server_peer_group_id", dynamicRouteServerPeerGroupID),
					resource.TestCheckResourceAttr("ibm_is_dynamic_route_server_peer_group_peer.is_dynamic_route_server_peer_group_peer_instance", "name", nameUpdate),
					resource.TestCheckResourceAttr("ibm_is_dynamic_route_server_peer_group_peer.is_dynamic_route_server_peer_group_peer_instance", "asn", asnUpdate),
					resource.TestCheckResourceAttr("ibm_is_dynamic_route_server_peer_group_peer.is_dynamic_route_server_peer_group_peer_instance", "priority", priorityUpdate),
					resource.TestCheckResourceAttr("ibm_is_dynamic_route_server_peer_group_peer.is_dynamic_route_server_peer_group_peer_instance", "state", stateUpdate),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_is_dynamic_route_server_peer_group_peer.is_dynamic_route_server_peer_group_peer_instance",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIBMIsDynamicRouteServerPeerGroupPeerConfigBasic(dynamicRouteServerID string, dynamicRouteServerPeerGroupID string) string {
	return fmt.Sprintf(`
		resource "ibm_is_dynamic_route_server_peer_group_peer" "is_dynamic_route_server_peer_group_peer_instance" {
			dynamic_route_server_id = "%s"
			dynamic_route_server_peer_group_id = "%s"
		}
	`, dynamicRouteServerID, dynamicRouteServerPeerGroupID)
}

func testAccCheckIBMIsDynamicRouteServerPeerGroupPeerConfig(dynamicRouteServerID string, dynamicRouteServerPeerGroupID string, name string, asn string, priority string, state string) string {
	return fmt.Sprintf(`

		resource "ibm_is_dynamic_route_server_peer_group_peer" "is_dynamic_route_server_peer_group_peer_instance" {
			dynamic_route_server_id = "%s"
			dynamic_route_server_peer_group_id = "%s"
			name = "%s"
			asn = %s
			endpoint {
				address = "192.168.3.4"
				gateway {
					address = "192.168.3.4"
				}
			}
			priority = %s
			state = "%s"
		}
	`, dynamicRouteServerID, dynamicRouteServerPeerGroupID, name, asn, priority, state)
}

func testAccCheckIBMIsDynamicRouteServerPeerGroupPeerExists(n string, obj vpcv1.DynamicRouteServerPeerGroupPeer) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		vpcClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).VpcV1API()
		if err != nil {
			return err
		}

		getDynamicRouteServerPeerGroupPeerOptions := &vpcv1.GetDynamicRouteServerPeerGroupPeerOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getDynamicRouteServerPeerGroupPeerOptions.SetDynamicRouteServerID(parts[0])
		getDynamicRouteServerPeerGroupPeerOptions.SetDynamicRouteServerPeerGroupID(parts[1])
		getDynamicRouteServerPeerGroupPeerOptions.SetID(parts[2])

		dynamicRouteServerPeerGroupPeerIntf, _, err := vpcClient.GetDynamicRouteServerPeerGroupPeer(getDynamicRouteServerPeerGroupPeerOptions)
		if err != nil {
			return err
		}

		dynamicRouteServerPeerGroupPeer := dynamicRouteServerPeerGroupPeerIntf.(*vpcv1.DynamicRouteServerPeerGroupPeer)
		obj = *dynamicRouteServerPeerGroupPeer
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

		getDynamicRouteServerPeerGroupPeerOptions := &vpcv1.GetDynamicRouteServerPeerGroupPeerOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getDynamicRouteServerPeerGroupPeerOptions.SetDynamicRouteServerID(parts[0])
		getDynamicRouteServerPeerGroupPeerOptions.SetDynamicRouteServerPeerGroupID(parts[1])
		getDynamicRouteServerPeerGroupPeerOptions.SetID(parts[2])

		// Try to find the key
		_, response, err := vpcClient.GetDynamicRouteServerPeerGroupPeer(getDynamicRouteServerPeerGroupPeerOptions)

		if err == nil {
			return fmt.Errorf("is_dynamic_route_server_peer_group_peer still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for is_dynamic_route_server_peer_group_peer (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}

func TestResourceIBMIsDynamicRouteServerPeerGroupPeerDynamicRouteServerPeerGroupPeerAddressEndpointToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		dynamicRouteServerPeerGroupPeerAddressEndpointIPGatewayGatewayModel := make(map[string]interface{})
		dynamicRouteServerPeerGroupPeerAddressEndpointIPGatewayGatewayModel["address"] = "192.168.3.4"

		model := make(map[string]interface{})
		model["address"] = "192.168.3.4"
		model["gateway"] = []map[string]interface{}{dynamicRouteServerPeerGroupPeerAddressEndpointIPGatewayGatewayModel}
		// model["deleted"] = []map[string]interface{}{deletedModel}
		model["href"] = "https://us-south.iaas.cloud.ibm.com/v1/subnets/0717-bea6a632-5e13-42a4-b4b8-31dc877abfe4/reserved_ips/0717-6d353a0f-aeb1-4ae1-832e-1110d10981bb"
		model["id"] = "0717-6d353a0f-aeb1-4ae1-832e-1110d10981bb"
		model["name"] = "my-reserved-ip"
		model["resource_type"] = "subnet_reserved_ip"

		assert.Equal(t, result, model)
	}

	dynamicRouteServerPeerGroupPeerAddressEndpointIPGatewayGatewayModel := new(vpcv1.DynamicRouteServerPeerGroupPeerAddressEndpointIPGatewayGateway)
	dynamicRouteServerPeerGroupPeerAddressEndpointIPGatewayGatewayModel.Address = core.StringPtr("192.168.3.4")

	model := new(vpcv1.DynamicRouteServerPeerGroupPeerAddressEndpoint)
	model.Address = core.StringPtr("192.168.3.4")
	model.Gateway = dynamicRouteServerPeerGroupPeerAddressEndpointIPGatewayGatewayModel
	// model.Deleted = deletedModel
	model.Href = core.StringPtr("https://us-south.iaas.cloud.ibm.com/v1/subnets/0717-bea6a632-5e13-42a4-b4b8-31dc877abfe4/reserved_ips/0717-6d353a0f-aeb1-4ae1-832e-1110d10981bb")
	model.ID = core.StringPtr("0717-6d353a0f-aeb1-4ae1-832e-1110d10981bb")
	model.Name = core.StringPtr("my-reserved-ip")
	model.ResourceType = core.StringPtr("subnet_reserved_ip")

	result, err := vpc.ResourceIBMIsDynamicRouteServerPeerGroupPeerDynamicRouteServerPeerGroupPeerAddressEndpointToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMIsDynamicRouteServerPeerGroupPeerDynamicRouteServerPeerGroupPeerAddressEndpointIPGatewayGatewayToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["address"] = "192.168.3.4"

		assert.Equal(t, result, model)
	}

	model := new(vpcv1.DynamicRouteServerPeerGroupPeerAddressEndpointIPGatewayGateway)
	model.Address = core.StringPtr("192.168.3.4")

	result, err := vpc.ResourceIBMIsDynamicRouteServerPeerGroupPeerDynamicRouteServerPeerGroupPeerAddressEndpointIPGatewayGatewayToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMIsDynamicRouteServerPeerGroupPeerDeletedToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["more_info"] = "https://cloud.ibm.com/apidocs/vpc#deleted-resources"

		assert.Equal(t, result, model)
	}

	model := new(vpcv1.Deleted)
	model.MoreInfo = core.StringPtr("https://cloud.ibm.com/apidocs/vpc#deleted-resources")

	result, err := vpc.ResourceIBMIsDynamicRouteServerPeerGroupPeerDeletedToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMIsDynamicRouteServerPeerGroupPeerDynamicRouteServerPeerGroupPeerAddressEndpointIPGatewayToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		dynamicRouteServerPeerGroupPeerAddressEndpointIPGatewayGatewayModel := make(map[string]interface{})
		dynamicRouteServerPeerGroupPeerAddressEndpointIPGatewayGatewayModel["address"] = "192.168.3.4"

		model := make(map[string]interface{})
		model["address"] = "192.168.3.4"
		model["gateway"] = []map[string]interface{}{dynamicRouteServerPeerGroupPeerAddressEndpointIPGatewayGatewayModel}

		assert.Equal(t, result, model)
	}

	dynamicRouteServerPeerGroupPeerAddressEndpointIPGatewayGatewayModel := new(vpcv1.DynamicRouteServerPeerGroupPeerAddressEndpointIPGatewayGateway)
	dynamicRouteServerPeerGroupPeerAddressEndpointIPGatewayGatewayModel.Address = core.StringPtr("192.168.3.4")

	model := new(vpcv1.DynamicRouteServerPeerGroupPeerAddressEndpointIPGateway)
	model.Address = core.StringPtr("192.168.3.4")
	model.Gateway = dynamicRouteServerPeerGroupPeerAddressEndpointIPGatewayGatewayModel

	result, err := vpc.ResourceIBMIsDynamicRouteServerPeerGroupPeerDynamicRouteServerPeerGroupPeerAddressEndpointIPGatewayToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMIsDynamicRouteServerPeerGroupPeerDynamicRouteServerPeerGroupPeerAddressEndpointByReservedIPToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		deletedModel := make(map[string]interface{})
		deletedModel["more_info"] = "https://cloud.ibm.com/apidocs/vpc#deleted-resources"

		ipModel := make(map[string]interface{})
		ipModel["address"] = "192.168.3.4"

		model := make(map[string]interface{})
		model["address"] = "192.168.3.4"
		model["deleted"] = []map[string]interface{}{deletedModel}
		model["gateway"] = []map[string]interface{}{ipModel}
		model["href"] = "https://us-south.iaas.cloud.ibm.com/v1/subnets/0717-bea6a632-5e13-42a4-b4b8-31dc877abfe4/reserved_ips/0717-6d353a0f-aeb1-4ae1-832e-1110d10981bb"
		model["id"] = "0717-6d353a0f-aeb1-4ae1-832e-1110d10981bb"
		model["name"] = "my-reserved-ip"
		model["resource_type"] = "subnet_reserved_ip"

		assert.Equal(t, result, model)
	}

	deletedModel := new(vpcv1.Deleted)
	deletedModel.MoreInfo = core.StringPtr("https://cloud.ibm.com/apidocs/vpc#deleted-resources")

	ipModel := new(vpcv1.IP)
	ipModel.Address = core.StringPtr("192.168.3.4")

	model := new(vpcv1.DynamicRouteServerPeerGroupPeerAddressEndpointByReservedIP)
	model.Address = core.StringPtr("192.168.3.4")
	model.Deleted = deletedModel
	model.Gateway = ipModel
	model.Href = core.StringPtr("https://us-south.iaas.cloud.ibm.com/v1/subnets/0717-bea6a632-5e13-42a4-b4b8-31dc877abfe4/reserved_ips/0717-6d353a0f-aeb1-4ae1-832e-1110d10981bb")
	model.ID = core.StringPtr("0717-6d353a0f-aeb1-4ae1-832e-1110d10981bb")
	model.Name = core.StringPtr("my-reserved-ip")
	model.ResourceType = core.StringPtr("subnet_reserved_ip")

	result, err := vpc.ResourceIBMIsDynamicRouteServerPeerGroupPeerDynamicRouteServerPeerGroupPeerAddressEndpointByReservedIPToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMIsDynamicRouteServerPeerGroupPeerIPToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["address"] = "192.168.3.4"

		assert.Equal(t, result, model)
	}

	model := new(vpcv1.IP)
	model.Address = core.StringPtr("192.168.3.4")

	result, err := vpc.ResourceIBMIsDynamicRouteServerPeerGroupPeerIPToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMIsDynamicRouteServerPeerGroupPeerLifecycleReasonToMap(t *testing.T) {
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

	result, err := vpc.ResourceIBMIsDynamicRouteServerPeerGroupPeerLifecycleReasonToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMIsDynamicRouteServerPeerGroupPeerDynamicRouteServerPeerGroupPeerStatusReasonToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["code"] = "peer_not_responding"
		model["message"] = "The connection is down because the peer is not responding."
		model["more_info"] = "https://cloud.ibm.com/docs/drs?topic=vpc-drs-health#__TBD__"

		assert.Equal(t, result, model)
	}

	model := new(vpcv1.DynamicRouteServerPeerGroupPeerStatusReason)
	model.Code = core.StringPtr("peer_not_responding")
	model.Message = core.StringPtr("The connection is down because the peer is not responding.")
	model.MoreInfo = core.StringPtr("https://cloud.ibm.com/docs/drs?topic=vpc-drs-health#__TBD__")

	result, err := vpc.ResourceIBMIsDynamicRouteServerPeerGroupPeerDynamicRouteServerPeerGroupPeerStatusReasonToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMIsDynamicRouteServerPeerGroupPeerDynamicRouteServerPeerGroupPeerAddressBidirectionalForwardingDetectionToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		deletedModel := make(map[string]interface{})
		deletedModel["more_info"] = "https://cloud.ibm.com/apidocs/vpc#deleted-resources"

		reservedIPReferenceModel := make(map[string]interface{})
		reservedIPReferenceModel["address"] = "192.168.3.4"
		reservedIPReferenceModel["deleted"] = []map[string]interface{}{deletedModel}
		reservedIPReferenceModel["href"] = "https://us-south.iaas.cloud.ibm.com/v1/subnets/0717-bea6a632-5e13-42a4-b4b8-31dc877abfe4/reserved_ips/0717-6d353a0f-aeb1-4ae1-832e-1110d10981bb"
		reservedIPReferenceModel["id"] = "0717-6d353a0f-aeb1-4ae1-832e-1110d10981bb"
		reservedIPReferenceModel["name"] = "my-reserved-ip"
		reservedIPReferenceModel["resource_type"] = "subnet_reserved_ip"

		subnetReferenceModel := make(map[string]interface{})
		subnetReferenceModel["crn"] = "crn:v1:bluemix:public:is:us-south-1:a/aa2432b1fa4d4ace891e9b80fc104e34::subnet:0717-7ec86020-1c6e-4889-b3f0-a15f2e50f87e"
		subnetReferenceModel["deleted"] = []map[string]interface{}{deletedModel}
		subnetReferenceModel["href"] = "https://us-south.iaas.cloud.ibm.com/v1/subnets/0717-7ec86020-1c6e-4889-b3f0-a15f2e50f87e"
		subnetReferenceModel["id"] = "0717-7ec86020-1c6e-4889-b3f0-a15f2e50f87e"
		subnetReferenceModel["name"] = "my-subnet"
		subnetReferenceModel["resource_type"] = "subnet"

		virtualNetworkInterfaceReferenceModel := make(map[string]interface{})
		virtualNetworkInterfaceReferenceModel["crn"] = "crn:v1:bluemix:public:is:us-south-1:a/aa2432b1fa4d4ace891e9b80fc104e34::virtual-network-interface:0717-54eb57ee-86f2-4796-90bb-d7874e0831ef"
		virtualNetworkInterfaceReferenceModel["deleted"] = []map[string]interface{}{deletedModel}
		virtualNetworkInterfaceReferenceModel["href"] = "https://us-south.iaas.cloud.ibm.com/v1/virtual_network_interfaces/0717-54eb57ee-86f2-4796-90bb-d7874e0831ef"
		virtualNetworkInterfaceReferenceModel["id"] = "0717-54eb57ee-86f2-4796-90bb-d7874e0831ef"
		virtualNetworkInterfaceReferenceModel["name"] = "my-virtual-network-interface"
		virtualNetworkInterfaceReferenceModel["primary_ip"] = []map[string]interface{}{reservedIPReferenceModel}
		virtualNetworkInterfaceReferenceModel["resource_type"] = "virtual_network_interface"
		virtualNetworkInterfaceReferenceModel["subnet"] = []map[string]interface{}{subnetReferenceModel}

		dynamicRouteServerMemberReferenceModel := make(map[string]interface{})
		dynamicRouteServerMemberReferenceModel["deleted"] = []map[string]interface{}{deletedModel}
		dynamicRouteServerMemberReferenceModel["href"] = "https://us-south.iaas.cloud.ibm.com/v1/dynamic_route_servers/r006-d7cc5196-9864-48c4-82d8-3f30da41fcc5/members/r006-d7cc5196-9924-48d4-82d8-3f30da41fdc"
		dynamicRouteServerMemberReferenceModel["id"] = "0717-9dce0ab3-a719-4582-ad71-00abfaae43ed"
		dynamicRouteServerMemberReferenceModel["name"] = "my-dynamic-route-server-member1"
		dynamicRouteServerMemberReferenceModel["resource_type"] = "dynamic_route_server_member"
		dynamicRouteServerMemberReferenceModel["virtual_network_interfaces"] = []map[string]interface{}{virtualNetworkInterfaceReferenceModel}

		dynamicRouteServerPeerGroupPeerReferenceModel := make(map[string]interface{})
		dynamicRouteServerPeerGroupPeerReferenceModel["deleted"] = []map[string]interface{}{deletedModel}
		dynamicRouteServerPeerGroupPeerReferenceModel["href"] = "https://us-south.iaas.cloud.ibm.com/v1/dynamic_route_servers/r006-d7cc5196-9864-48c4-82d8-3f30da41fcc5/peer_groups/r006-8228af60-189d-11ed-861d-0242ac120002/peers/r006-8228af60-189d-11ed-861d-0242ac120004"
		dynamicRouteServerPeerGroupPeerReferenceModel["id"] = "r006-8228af60-189d-11ed-861d-0242ac120004"
		dynamicRouteServerPeerGroupPeerReferenceModel["name"] = "my-dynamic-route-server-peer-group-peer"
		dynamicRouteServerPeerGroupPeerReferenceModel["resource_type"] = "dynamic_route_server_peer_group_peer"

		dynamicRouteServerPeerGroupPeerAddressBidirectionalForwardingDetectionSessionModel := make(map[string]interface{})
		dynamicRouteServerPeerGroupPeerAddressBidirectionalForwardingDetectionSessionModel["local"] = []map[string]interface{}{dynamicRouteServerMemberReferenceModel}
		dynamicRouteServerPeerGroupPeerAddressBidirectionalForwardingDetectionSessionModel["remote"] = []map[string]interface{}{dynamicRouteServerPeerGroupPeerReferenceModel}
		dynamicRouteServerPeerGroupPeerAddressBidirectionalForwardingDetectionSessionModel["state"] = "admin_down"

		model := make(map[string]interface{})
		model["detect_multiplier"] = int(3)
		model["enabled"] = true
		model["mode"] = "asynchronous"
		model["receive_interval"] = int(300)
		model["role"] = "active"
		model["sessions"] = []map[string]interface{}{dynamicRouteServerPeerGroupPeerAddressBidirectionalForwardingDetectionSessionModel}
		model["transmit_interval"] = int(300)

		assert.Equal(t, result, model)
	}

	deletedModel := new(vpcv1.Deleted)
	deletedModel.MoreInfo = core.StringPtr("https://cloud.ibm.com/apidocs/vpc#deleted-resources")

	reservedIPReferenceModel := new(vpcv1.ReservedIPReference)
	reservedIPReferenceModel.Address = core.StringPtr("192.168.3.4")
	reservedIPReferenceModel.Deleted = deletedModel
	reservedIPReferenceModel.Href = core.StringPtr("https://us-south.iaas.cloud.ibm.com/v1/subnets/0717-bea6a632-5e13-42a4-b4b8-31dc877abfe4/reserved_ips/0717-6d353a0f-aeb1-4ae1-832e-1110d10981bb")
	reservedIPReferenceModel.ID = core.StringPtr("0717-6d353a0f-aeb1-4ae1-832e-1110d10981bb")
	reservedIPReferenceModel.Name = core.StringPtr("my-reserved-ip")
	reservedIPReferenceModel.ResourceType = core.StringPtr("subnet_reserved_ip")

	subnetReferenceModel := new(vpcv1.SubnetReference)
	subnetReferenceModel.CRN = core.StringPtr("crn:v1:bluemix:public:is:us-south-1:a/aa2432b1fa4d4ace891e9b80fc104e34::subnet:0717-7ec86020-1c6e-4889-b3f0-a15f2e50f87e")
	subnetReferenceModel.Deleted = deletedModel
	subnetReferenceModel.Href = core.StringPtr("https://us-south.iaas.cloud.ibm.com/v1/subnets/0717-7ec86020-1c6e-4889-b3f0-a15f2e50f87e")
	subnetReferenceModel.ID = core.StringPtr("0717-7ec86020-1c6e-4889-b3f0-a15f2e50f87e")
	subnetReferenceModel.Name = core.StringPtr("my-subnet")
	subnetReferenceModel.ResourceType = core.StringPtr("subnet")

	virtualNetworkInterfaceReferenceModel := new(vpcv1.VirtualNetworkInterfaceReference)
	virtualNetworkInterfaceReferenceModel.CRN = core.StringPtr("crn:v1:bluemix:public:is:us-south-1:a/aa2432b1fa4d4ace891e9b80fc104e34::virtual-network-interface:0717-54eb57ee-86f2-4796-90bb-d7874e0831ef")
	virtualNetworkInterfaceReferenceModel.Deleted = deletedModel
	virtualNetworkInterfaceReferenceModel.Href = core.StringPtr("https://us-south.iaas.cloud.ibm.com/v1/virtual_network_interfaces/0717-54eb57ee-86f2-4796-90bb-d7874e0831ef")
	virtualNetworkInterfaceReferenceModel.ID = core.StringPtr("0717-54eb57ee-86f2-4796-90bb-d7874e0831ef")
	virtualNetworkInterfaceReferenceModel.Name = core.StringPtr("my-virtual-network-interface")
	virtualNetworkInterfaceReferenceModel.PrimaryIP = reservedIPReferenceModel
	virtualNetworkInterfaceReferenceModel.ResourceType = core.StringPtr("virtual_network_interface")
	virtualNetworkInterfaceReferenceModel.Subnet = subnetReferenceModel

	dynamicRouteServerMemberReferenceModel := new(vpcv1.DynamicRouteServerMemberReference)
	dynamicRouteServerMemberReferenceModel.Deleted = deletedModel
	dynamicRouteServerMemberReferenceModel.Href = core.StringPtr("https://us-south.iaas.cloud.ibm.com/v1/dynamic_route_servers/r006-d7cc5196-9864-48c4-82d8-3f30da41fcc5/members/r006-d7cc5196-9924-48d4-82d8-3f30da41fdc")
	dynamicRouteServerMemberReferenceModel.ID = core.StringPtr("0717-9dce0ab3-a719-4582-ad71-00abfaae43ed")
	dynamicRouteServerMemberReferenceModel.Name = core.StringPtr("my-dynamic-route-server-member1")
	dynamicRouteServerMemberReferenceModel.ResourceType = core.StringPtr("dynamic_route_server_member")
	dynamicRouteServerMemberReferenceModel.VirtualNetworkInterfaces = []vpcv1.VirtualNetworkInterfaceReference{*virtualNetworkInterfaceReferenceModel}

	dynamicRouteServerPeerGroupPeerReferenceModel := new(vpcv1.DynamicRouteServerPeerGroupPeerReference)
	dynamicRouteServerPeerGroupPeerReferenceModel.Deleted = deletedModel
	dynamicRouteServerPeerGroupPeerReferenceModel.Href = core.StringPtr("https://us-south.iaas.cloud.ibm.com/v1/dynamic_route_servers/r006-d7cc5196-9864-48c4-82d8-3f30da41fcc5/peer_groups/r006-8228af60-189d-11ed-861d-0242ac120002/peers/r006-8228af60-189d-11ed-861d-0242ac120004")
	dynamicRouteServerPeerGroupPeerReferenceModel.ID = core.StringPtr("r006-8228af60-189d-11ed-861d-0242ac120004")
	dynamicRouteServerPeerGroupPeerReferenceModel.Name = core.StringPtr("my-dynamic-route-server-peer-group-peer")
	dynamicRouteServerPeerGroupPeerReferenceModel.ResourceType = core.StringPtr("dynamic_route_server_peer_group_peer")

	dynamicRouteServerPeerGroupPeerAddressBidirectionalForwardingDetectionSessionModel := new(vpcv1.DynamicRouteServerPeerGroupPeerAddressBidirectionalForwardingDetectionSession)
	dynamicRouteServerPeerGroupPeerAddressBidirectionalForwardingDetectionSessionModel.Local = dynamicRouteServerMemberReferenceModel
	dynamicRouteServerPeerGroupPeerAddressBidirectionalForwardingDetectionSessionModel.Remote = dynamicRouteServerPeerGroupPeerReferenceModel
	dynamicRouteServerPeerGroupPeerAddressBidirectionalForwardingDetectionSessionModel.State = core.StringPtr("admin_down")

	model := new(vpcv1.DynamicRouteServerPeerGroupPeerAddressBidirectionalForwardingDetection)
	model.DetectMultiplier = core.Int64Ptr(int64(3))
	model.Enabled = core.BoolPtr(true)
	model.Mode = core.StringPtr("asynchronous")
	model.ReceiveInterval = core.Int64Ptr(int64(300))
	model.Role = core.StringPtr("active")
	model.Sessions = []vpcv1.DynamicRouteServerPeerGroupPeerAddressBidirectionalForwardingDetectionSession{*dynamicRouteServerPeerGroupPeerAddressBidirectionalForwardingDetectionSessionModel}
	model.TransmitInterval = core.Int64Ptr(int64(300))

	result, err := vpc.ResourceIBMIsDynamicRouteServerPeerGroupPeerDynamicRouteServerPeerGroupPeerAddressBidirectionalForwardingDetectionToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMIsDynamicRouteServerPeerGroupPeerDynamicRouteServerPeerGroupPeerAddressBidirectionalForwardingDetectionSessionToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		deletedModel := make(map[string]interface{})
		deletedModel["more_info"] = "https://cloud.ibm.com/apidocs/vpc#deleted-resources"

		reservedIPReferenceModel := make(map[string]interface{})
		reservedIPReferenceModel["address"] = "192.168.3.4"
		reservedIPReferenceModel["deleted"] = []map[string]interface{}{deletedModel}
		reservedIPReferenceModel["href"] = "https://us-south.iaas.cloud.ibm.com/v1/subnets/0717-bea6a632-5e13-42a4-b4b8-31dc877abfe4/reserved_ips/0717-6d353a0f-aeb1-4ae1-832e-1110d10981bb"
		reservedIPReferenceModel["id"] = "0717-6d353a0f-aeb1-4ae1-832e-1110d10981bb"
		reservedIPReferenceModel["name"] = "my-reserved-ip"
		reservedIPReferenceModel["resource_type"] = "subnet_reserved_ip"

		subnetReferenceModel := make(map[string]interface{})
		subnetReferenceModel["crn"] = "crn:v1:bluemix:public:is:us-south-1:a/aa2432b1fa4d4ace891e9b80fc104e34::subnet:0717-7ec86020-1c6e-4889-b3f0-a15f2e50f87e"
		subnetReferenceModel["deleted"] = []map[string]interface{}{deletedModel}
		subnetReferenceModel["href"] = "https://us-south.iaas.cloud.ibm.com/v1/subnets/0717-7ec86020-1c6e-4889-b3f0-a15f2e50f87e"
		subnetReferenceModel["id"] = "0717-7ec86020-1c6e-4889-b3f0-a15f2e50f87e"
		subnetReferenceModel["name"] = "my-subnet"
		subnetReferenceModel["resource_type"] = "subnet"

		virtualNetworkInterfaceReferenceModel := make(map[string]interface{})
		virtualNetworkInterfaceReferenceModel["crn"] = "crn:v1:bluemix:public:is:us-south-1:a/aa2432b1fa4d4ace891e9b80fc104e34::virtual-network-interface:0717-54eb57ee-86f2-4796-90bb-d7874e0831ef"
		virtualNetworkInterfaceReferenceModel["deleted"] = []map[string]interface{}{deletedModel}
		virtualNetworkInterfaceReferenceModel["href"] = "https://us-south.iaas.cloud.ibm.com/v1/virtual_network_interfaces/0717-54eb57ee-86f2-4796-90bb-d7874e0831ef"
		virtualNetworkInterfaceReferenceModel["id"] = "0717-54eb57ee-86f2-4796-90bb-d7874e0831ef"
		virtualNetworkInterfaceReferenceModel["name"] = "my-virtual-network-interface"
		virtualNetworkInterfaceReferenceModel["primary_ip"] = []map[string]interface{}{reservedIPReferenceModel}
		virtualNetworkInterfaceReferenceModel["resource_type"] = "virtual_network_interface"
		virtualNetworkInterfaceReferenceModel["subnet"] = []map[string]interface{}{subnetReferenceModel}

		dynamicRouteServerMemberReferenceModel := make(map[string]interface{})
		dynamicRouteServerMemberReferenceModel["deleted"] = []map[string]interface{}{deletedModel}
		dynamicRouteServerMemberReferenceModel["href"] = "https://us-south.iaas.cloud.ibm.com/v1/dynamic_route_servers/r006-d7cc5196-9864-48c4-82d8-3f30da41fcc5/members/r006-d7cc5196-9924-48d4-82d8-3f30da41fdc"
		dynamicRouteServerMemberReferenceModel["id"] = "0717-9dce0ab3-a719-4582-ad71-00abfaae43ed"
		dynamicRouteServerMemberReferenceModel["name"] = "my-dynamic-route-server-member1"
		dynamicRouteServerMemberReferenceModel["resource_type"] = "dynamic_route_server_member"
		dynamicRouteServerMemberReferenceModel["virtual_network_interfaces"] = []map[string]interface{}{virtualNetworkInterfaceReferenceModel}

		dynamicRouteServerPeerGroupPeerReferenceModel := make(map[string]interface{})
		dynamicRouteServerPeerGroupPeerReferenceModel["deleted"] = []map[string]interface{}{deletedModel}
		dynamicRouteServerPeerGroupPeerReferenceModel["href"] = "https://us-south.iaas.cloud.ibm.com/v1/dynamic_route_servers/r006-d7cc5196-9864-48c4-82d8-3f30da41fcc5/peer_groups/r006-8228af60-189d-11ed-861d-0242ac120002/peers/r006-8228af60-189d-11ed-861d-0242ac120004"
		dynamicRouteServerPeerGroupPeerReferenceModel["id"] = "r006-8228af60-189d-11ed-861d-0242ac120004"
		dynamicRouteServerPeerGroupPeerReferenceModel["name"] = "my-dynamic-route-server-peer-group-peer"
		dynamicRouteServerPeerGroupPeerReferenceModel["resource_type"] = "dynamic_route_server_peer_group_peer"

		model := make(map[string]interface{})
		model["local"] = []map[string]interface{}{dynamicRouteServerMemberReferenceModel}
		model["remote"] = []map[string]interface{}{dynamicRouteServerPeerGroupPeerReferenceModel}
		model["state"] = "admin_down"

		assert.Equal(t, result, model)
	}

	deletedModel := new(vpcv1.Deleted)
	deletedModel.MoreInfo = core.StringPtr("https://cloud.ibm.com/apidocs/vpc#deleted-resources")

	reservedIPReferenceModel := new(vpcv1.ReservedIPReference)
	reservedIPReferenceModel.Address = core.StringPtr("192.168.3.4")
	reservedIPReferenceModel.Deleted = deletedModel
	reservedIPReferenceModel.Href = core.StringPtr("https://us-south.iaas.cloud.ibm.com/v1/subnets/0717-bea6a632-5e13-42a4-b4b8-31dc877abfe4/reserved_ips/0717-6d353a0f-aeb1-4ae1-832e-1110d10981bb")
	reservedIPReferenceModel.ID = core.StringPtr("0717-6d353a0f-aeb1-4ae1-832e-1110d10981bb")
	reservedIPReferenceModel.Name = core.StringPtr("my-reserved-ip")
	reservedIPReferenceModel.ResourceType = core.StringPtr("subnet_reserved_ip")

	subnetReferenceModel := new(vpcv1.SubnetReference)
	subnetReferenceModel.CRN = core.StringPtr("crn:v1:bluemix:public:is:us-south-1:a/aa2432b1fa4d4ace891e9b80fc104e34::subnet:0717-7ec86020-1c6e-4889-b3f0-a15f2e50f87e")
	subnetReferenceModel.Deleted = deletedModel
	subnetReferenceModel.Href = core.StringPtr("https://us-south.iaas.cloud.ibm.com/v1/subnets/0717-7ec86020-1c6e-4889-b3f0-a15f2e50f87e")
	subnetReferenceModel.ID = core.StringPtr("0717-7ec86020-1c6e-4889-b3f0-a15f2e50f87e")
	subnetReferenceModel.Name = core.StringPtr("my-subnet")
	subnetReferenceModel.ResourceType = core.StringPtr("subnet")

	virtualNetworkInterfaceReferenceModel := new(vpcv1.VirtualNetworkInterfaceReference)
	virtualNetworkInterfaceReferenceModel.CRN = core.StringPtr("crn:v1:bluemix:public:is:us-south-1:a/aa2432b1fa4d4ace891e9b80fc104e34::virtual-network-interface:0717-54eb57ee-86f2-4796-90bb-d7874e0831ef")
	virtualNetworkInterfaceReferenceModel.Deleted = deletedModel
	virtualNetworkInterfaceReferenceModel.Href = core.StringPtr("https://us-south.iaas.cloud.ibm.com/v1/virtual_network_interfaces/0717-54eb57ee-86f2-4796-90bb-d7874e0831ef")
	virtualNetworkInterfaceReferenceModel.ID = core.StringPtr("0717-54eb57ee-86f2-4796-90bb-d7874e0831ef")
	virtualNetworkInterfaceReferenceModel.Name = core.StringPtr("my-virtual-network-interface")
	virtualNetworkInterfaceReferenceModel.PrimaryIP = reservedIPReferenceModel
	virtualNetworkInterfaceReferenceModel.ResourceType = core.StringPtr("virtual_network_interface")
	virtualNetworkInterfaceReferenceModel.Subnet = subnetReferenceModel

	dynamicRouteServerMemberReferenceModel := new(vpcv1.DynamicRouteServerMemberReference)
	dynamicRouteServerMemberReferenceModel.Deleted = deletedModel
	dynamicRouteServerMemberReferenceModel.Href = core.StringPtr("https://us-south.iaas.cloud.ibm.com/v1/dynamic_route_servers/r006-d7cc5196-9864-48c4-82d8-3f30da41fcc5/members/r006-d7cc5196-9924-48d4-82d8-3f30da41fdc")
	dynamicRouteServerMemberReferenceModel.ID = core.StringPtr("0717-9dce0ab3-a719-4582-ad71-00abfaae43ed")
	dynamicRouteServerMemberReferenceModel.Name = core.StringPtr("my-dynamic-route-server-member1")
	dynamicRouteServerMemberReferenceModel.ResourceType = core.StringPtr("dynamic_route_server_member")
	dynamicRouteServerMemberReferenceModel.VirtualNetworkInterfaces = []vpcv1.VirtualNetworkInterfaceReference{*virtualNetworkInterfaceReferenceModel}

	dynamicRouteServerPeerGroupPeerReferenceModel := new(vpcv1.DynamicRouteServerPeerGroupPeerReference)
	dynamicRouteServerPeerGroupPeerReferenceModel.Deleted = deletedModel
	dynamicRouteServerPeerGroupPeerReferenceModel.Href = core.StringPtr("https://us-south.iaas.cloud.ibm.com/v1/dynamic_route_servers/r006-d7cc5196-9864-48c4-82d8-3f30da41fcc5/peer_groups/r006-8228af60-189d-11ed-861d-0242ac120002/peers/r006-8228af60-189d-11ed-861d-0242ac120004")
	dynamicRouteServerPeerGroupPeerReferenceModel.ID = core.StringPtr("r006-8228af60-189d-11ed-861d-0242ac120004")
	dynamicRouteServerPeerGroupPeerReferenceModel.Name = core.StringPtr("my-dynamic-route-server-peer-group-peer")
	dynamicRouteServerPeerGroupPeerReferenceModel.ResourceType = core.StringPtr("dynamic_route_server_peer_group_peer")

	model := new(vpcv1.DynamicRouteServerPeerGroupPeerAddressBidirectionalForwardingDetectionSession)
	model.Local = dynamicRouteServerMemberReferenceModel
	model.Remote = dynamicRouteServerPeerGroupPeerReferenceModel
	model.State = core.StringPtr("admin_down")

	result, err := vpc.ResourceIBMIsDynamicRouteServerPeerGroupPeerDynamicRouteServerPeerGroupPeerAddressBidirectionalForwardingDetectionSessionToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMIsDynamicRouteServerPeerGroupPeerDynamicRouteServerMemberReferenceToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		deletedModel := make(map[string]interface{})
		deletedModel["more_info"] = "https://cloud.ibm.com/apidocs/vpc#deleted-resources"

		reservedIPReferenceModel := make(map[string]interface{})
		reservedIPReferenceModel["address"] = "192.168.3.4"
		reservedIPReferenceModel["deleted"] = []map[string]interface{}{deletedModel}
		reservedIPReferenceModel["href"] = "https://us-south.iaas.cloud.ibm.com/v1/subnets/0717-bea6a632-5e13-42a4-b4b8-31dc877abfe4/reserved_ips/0717-6d353a0f-aeb1-4ae1-832e-1110d10981bb"
		reservedIPReferenceModel["id"] = "0717-6d353a0f-aeb1-4ae1-832e-1110d10981bb"
		reservedIPReferenceModel["name"] = "my-reserved-ip"
		reservedIPReferenceModel["resource_type"] = "subnet_reserved_ip"

		subnetReferenceModel := make(map[string]interface{})
		subnetReferenceModel["crn"] = "crn:v1:bluemix:public:is:us-south-1:a/aa2432b1fa4d4ace891e9b80fc104e34::subnet:0717-7ec86020-1c6e-4889-b3f0-a15f2e50f87e"
		subnetReferenceModel["deleted"] = []map[string]interface{}{deletedModel}
		subnetReferenceModel["href"] = "https://us-south.iaas.cloud.ibm.com/v1/subnets/0717-7ec86020-1c6e-4889-b3f0-a15f2e50f87e"
		subnetReferenceModel["id"] = "0717-7ec86020-1c6e-4889-b3f0-a15f2e50f87e"
		subnetReferenceModel["name"] = "my-subnet"
		subnetReferenceModel["resource_type"] = "subnet"

		virtualNetworkInterfaceReferenceModel := make(map[string]interface{})
		virtualNetworkInterfaceReferenceModel["crn"] = "crn:v1:bluemix:public:is:us-south-1:a/aa2432b1fa4d4ace891e9b80fc104e34::virtual-network-interface:0717-54eb57ee-86f2-4796-90bb-d7874e0831ef"
		virtualNetworkInterfaceReferenceModel["deleted"] = []map[string]interface{}{deletedModel}
		virtualNetworkInterfaceReferenceModel["href"] = "https://us-south.iaas.cloud.ibm.com/v1/virtual_network_interfaces/0717-54eb57ee-86f2-4796-90bb-d7874e0831ef"
		virtualNetworkInterfaceReferenceModel["id"] = "0717-54eb57ee-86f2-4796-90bb-d7874e0831ef"
		virtualNetworkInterfaceReferenceModel["name"] = "my-virtual-network-interface"
		virtualNetworkInterfaceReferenceModel["primary_ip"] = []map[string]interface{}{reservedIPReferenceModel}
		virtualNetworkInterfaceReferenceModel["resource_type"] = "virtual_network_interface"
		virtualNetworkInterfaceReferenceModel["subnet"] = []map[string]interface{}{subnetReferenceModel}

		model := make(map[string]interface{})
		model["deleted"] = []map[string]interface{}{deletedModel}
		model["href"] = "https://us-south.iaas.cloud.ibm.com/v1/dynamic_route_servers/r006-d7cc5196-9864-48c4-82d8-3f30da41fcc5/members/r006-d7cc5196-9924-48d4-82d8-3f30da41fdc"
		model["id"] = "0717-9dce0ab3-a719-4582-ad71-00abfaae43ed"
		model["name"] = "my-dynamic-route-server-member1"
		model["resource_type"] = "dynamic_route_server_member"
		model["virtual_network_interfaces"] = []map[string]interface{}{virtualNetworkInterfaceReferenceModel}

		assert.Equal(t, result, model)
	}

	deletedModel := new(vpcv1.Deleted)
	deletedModel.MoreInfo = core.StringPtr("https://cloud.ibm.com/apidocs/vpc#deleted-resources")

	reservedIPReferenceModel := new(vpcv1.ReservedIPReference)
	reservedIPReferenceModel.Address = core.StringPtr("192.168.3.4")
	reservedIPReferenceModel.Deleted = deletedModel
	reservedIPReferenceModel.Href = core.StringPtr("https://us-south.iaas.cloud.ibm.com/v1/subnets/0717-bea6a632-5e13-42a4-b4b8-31dc877abfe4/reserved_ips/0717-6d353a0f-aeb1-4ae1-832e-1110d10981bb")
	reservedIPReferenceModel.ID = core.StringPtr("0717-6d353a0f-aeb1-4ae1-832e-1110d10981bb")
	reservedIPReferenceModel.Name = core.StringPtr("my-reserved-ip")
	reservedIPReferenceModel.ResourceType = core.StringPtr("subnet_reserved_ip")

	subnetReferenceModel := new(vpcv1.SubnetReference)
	subnetReferenceModel.CRN = core.StringPtr("crn:v1:bluemix:public:is:us-south-1:a/aa2432b1fa4d4ace891e9b80fc104e34::subnet:0717-7ec86020-1c6e-4889-b3f0-a15f2e50f87e")
	subnetReferenceModel.Deleted = deletedModel
	subnetReferenceModel.Href = core.StringPtr("https://us-south.iaas.cloud.ibm.com/v1/subnets/0717-7ec86020-1c6e-4889-b3f0-a15f2e50f87e")
	subnetReferenceModel.ID = core.StringPtr("0717-7ec86020-1c6e-4889-b3f0-a15f2e50f87e")
	subnetReferenceModel.Name = core.StringPtr("my-subnet")
	subnetReferenceModel.ResourceType = core.StringPtr("subnet")

	virtualNetworkInterfaceReferenceModel := new(vpcv1.VirtualNetworkInterfaceReference)
	virtualNetworkInterfaceReferenceModel.CRN = core.StringPtr("crn:v1:bluemix:public:is:us-south-1:a/aa2432b1fa4d4ace891e9b80fc104e34::virtual-network-interface:0717-54eb57ee-86f2-4796-90bb-d7874e0831ef")
	virtualNetworkInterfaceReferenceModel.Deleted = deletedModel
	virtualNetworkInterfaceReferenceModel.Href = core.StringPtr("https://us-south.iaas.cloud.ibm.com/v1/virtual_network_interfaces/0717-54eb57ee-86f2-4796-90bb-d7874e0831ef")
	virtualNetworkInterfaceReferenceModel.ID = core.StringPtr("0717-54eb57ee-86f2-4796-90bb-d7874e0831ef")
	virtualNetworkInterfaceReferenceModel.Name = core.StringPtr("my-virtual-network-interface")
	virtualNetworkInterfaceReferenceModel.PrimaryIP = reservedIPReferenceModel
	virtualNetworkInterfaceReferenceModel.ResourceType = core.StringPtr("virtual_network_interface")
	virtualNetworkInterfaceReferenceModel.Subnet = subnetReferenceModel

	model := new(vpcv1.DynamicRouteServerMemberReference)
	model.Deleted = deletedModel
	model.Href = core.StringPtr("https://us-south.iaas.cloud.ibm.com/v1/dynamic_route_servers/r006-d7cc5196-9864-48c4-82d8-3f30da41fcc5/members/r006-d7cc5196-9924-48d4-82d8-3f30da41fdc")
	model.ID = core.StringPtr("0717-9dce0ab3-a719-4582-ad71-00abfaae43ed")
	model.Name = core.StringPtr("my-dynamic-route-server-member1")
	model.ResourceType = core.StringPtr("dynamic_route_server_member")
	model.VirtualNetworkInterfaces = []vpcv1.VirtualNetworkInterfaceReference{*virtualNetworkInterfaceReferenceModel}

	result, err := vpc.ResourceIBMIsDynamicRouteServerPeerGroupPeerDynamicRouteServerMemberReferenceToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMIsDynamicRouteServerPeerGroupPeerVirtualNetworkInterfaceReferenceToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		deletedModel := make(map[string]interface{})
		deletedModel["more_info"] = "https://cloud.ibm.com/apidocs/vpc#deleted-resources"

		reservedIPReferenceModel := make(map[string]interface{})
		reservedIPReferenceModel["address"] = "192.168.3.4"
		reservedIPReferenceModel["deleted"] = []map[string]interface{}{deletedModel}
		reservedIPReferenceModel["href"] = "https://us-south.iaas.cloud.ibm.com/v1/subnets/0717-bea6a632-5e13-42a4-b4b8-31dc877abfe4/reserved_ips/0717-6d353a0f-aeb1-4ae1-832e-1110d10981bb"
		reservedIPReferenceModel["id"] = "0717-6d353a0f-aeb1-4ae1-832e-1110d10981bb"
		reservedIPReferenceModel["name"] = "my-reserved-ip"
		reservedIPReferenceModel["resource_type"] = "subnet_reserved_ip"

		subnetReferenceModel := make(map[string]interface{})
		subnetReferenceModel["crn"] = "crn:v1:bluemix:public:is:us-south-1:a/aa2432b1fa4d4ace891e9b80fc104e34::subnet:0717-7ec86020-1c6e-4889-b3f0-a15f2e50f87e"
		subnetReferenceModel["deleted"] = []map[string]interface{}{deletedModel}
		subnetReferenceModel["href"] = "https://us-south.iaas.cloud.ibm.com/v1/subnets/0717-7ec86020-1c6e-4889-b3f0-a15f2e50f87e"
		subnetReferenceModel["id"] = "0717-7ec86020-1c6e-4889-b3f0-a15f2e50f87e"
		subnetReferenceModel["name"] = "my-subnet"
		subnetReferenceModel["resource_type"] = "subnet"

		model := make(map[string]interface{})
		model["crn"] = "crn:v1:bluemix:public:is:us-south-1:a/aa2432b1fa4d4ace891e9b80fc104e34::virtual-network-interface:0717-54eb57ee-86f2-4796-90bb-d7874e0831ef"
		model["deleted"] = []map[string]interface{}{deletedModel}
		model["href"] = "https://us-south.iaas.cloud.ibm.com/v1/virtual_network_interfaces/0717-54eb57ee-86f2-4796-90bb-d7874e0831ef"
		model["id"] = "0717-54eb57ee-86f2-4796-90bb-d7874e0831ef"
		model["name"] = "my-virtual-network-interface"
		model["primary_ip"] = []map[string]interface{}{reservedIPReferenceModel}
		model["resource_type"] = "virtual_network_interface"
		model["subnet"] = []map[string]interface{}{subnetReferenceModel}

		assert.Equal(t, result, model)
	}

	deletedModel := new(vpcv1.Deleted)
	deletedModel.MoreInfo = core.StringPtr("https://cloud.ibm.com/apidocs/vpc#deleted-resources")

	reservedIPReferenceModel := new(vpcv1.ReservedIPReference)
	reservedIPReferenceModel.Address = core.StringPtr("192.168.3.4")
	reservedIPReferenceModel.Deleted = deletedModel
	reservedIPReferenceModel.Href = core.StringPtr("https://us-south.iaas.cloud.ibm.com/v1/subnets/0717-bea6a632-5e13-42a4-b4b8-31dc877abfe4/reserved_ips/0717-6d353a0f-aeb1-4ae1-832e-1110d10981bb")
	reservedIPReferenceModel.ID = core.StringPtr("0717-6d353a0f-aeb1-4ae1-832e-1110d10981bb")
	reservedIPReferenceModel.Name = core.StringPtr("my-reserved-ip")
	reservedIPReferenceModel.ResourceType = core.StringPtr("subnet_reserved_ip")

	subnetReferenceModel := new(vpcv1.SubnetReference)
	subnetReferenceModel.CRN = core.StringPtr("crn:v1:bluemix:public:is:us-south-1:a/aa2432b1fa4d4ace891e9b80fc104e34::subnet:0717-7ec86020-1c6e-4889-b3f0-a15f2e50f87e")
	subnetReferenceModel.Deleted = deletedModel
	subnetReferenceModel.Href = core.StringPtr("https://us-south.iaas.cloud.ibm.com/v1/subnets/0717-7ec86020-1c6e-4889-b3f0-a15f2e50f87e")
	subnetReferenceModel.ID = core.StringPtr("0717-7ec86020-1c6e-4889-b3f0-a15f2e50f87e")
	subnetReferenceModel.Name = core.StringPtr("my-subnet")
	subnetReferenceModel.ResourceType = core.StringPtr("subnet")

	model := new(vpcv1.VirtualNetworkInterfaceReference)
	model.CRN = core.StringPtr("crn:v1:bluemix:public:is:us-south-1:a/aa2432b1fa4d4ace891e9b80fc104e34::virtual-network-interface:0717-54eb57ee-86f2-4796-90bb-d7874e0831ef")
	model.Deleted = deletedModel
	model.Href = core.StringPtr("https://us-south.iaas.cloud.ibm.com/v1/virtual_network_interfaces/0717-54eb57ee-86f2-4796-90bb-d7874e0831ef")
	model.ID = core.StringPtr("0717-54eb57ee-86f2-4796-90bb-d7874e0831ef")
	model.Name = core.StringPtr("my-virtual-network-interface")
	model.PrimaryIP = reservedIPReferenceModel
	model.ResourceType = core.StringPtr("virtual_network_interface")
	model.Subnet = subnetReferenceModel

	result, err := vpc.ResourceIBMIsDynamicRouteServerPeerGroupPeerVirtualNetworkInterfaceReferenceToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMIsDynamicRouteServerPeerGroupPeerReservedIPReferenceToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		deletedModel := make(map[string]interface{})
		deletedModel["more_info"] = "https://cloud.ibm.com/apidocs/vpc#deleted-resources"

		model := make(map[string]interface{})
		model["address"] = "192.168.3.4"
		model["deleted"] = []map[string]interface{}{deletedModel}
		model["href"] = "https://us-south.iaas.cloud.ibm.com/v1/subnets/0717-bea6a632-5e13-42a4-b4b8-31dc877abfe4/reserved_ips/0717-6d353a0f-aeb1-4ae1-832e-1110d10981bb"
		model["id"] = "0717-6d353a0f-aeb1-4ae1-832e-1110d10981bb"
		model["name"] = "my-reserved-ip"
		model["resource_type"] = "subnet_reserved_ip"

		assert.Equal(t, result, model)
	}

	deletedModel := new(vpcv1.Deleted)
	deletedModel.MoreInfo = core.StringPtr("https://cloud.ibm.com/apidocs/vpc#deleted-resources")

	model := new(vpcv1.ReservedIPReference)
	model.Address = core.StringPtr("192.168.3.4")
	model.Deleted = deletedModel
	model.Href = core.StringPtr("https://us-south.iaas.cloud.ibm.com/v1/subnets/0717-bea6a632-5e13-42a4-b4b8-31dc877abfe4/reserved_ips/0717-6d353a0f-aeb1-4ae1-832e-1110d10981bb")
	model.ID = core.StringPtr("0717-6d353a0f-aeb1-4ae1-832e-1110d10981bb")
	model.Name = core.StringPtr("my-reserved-ip")
	model.ResourceType = core.StringPtr("subnet_reserved_ip")

	result, err := vpc.ResourceIBMIsDynamicRouteServerPeerGroupPeerReservedIPReferenceToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMIsDynamicRouteServerPeerGroupPeerSubnetReferenceToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		deletedModel := make(map[string]interface{})
		deletedModel["more_info"] = "https://cloud.ibm.com/apidocs/vpc#deleted-resources"

		model := make(map[string]interface{})
		model["crn"] = "crn:v1:bluemix:public:is:us-south-1:a/aa2432b1fa4d4ace891e9b80fc104e34::subnet:0717-7ec86020-1c6e-4889-b3f0-a15f2e50f87e"
		model["deleted"] = []map[string]interface{}{deletedModel}
		model["href"] = "https://us-south.iaas.cloud.ibm.com/v1/subnets/0717-7ec86020-1c6e-4889-b3f0-a15f2e50f87e"
		model["id"] = "0717-7ec86020-1c6e-4889-b3f0-a15f2e50f87e"
		model["name"] = "my-subnet"
		model["resource_type"] = "subnet"

		assert.Equal(t, result, model)
	}

	deletedModel := new(vpcv1.Deleted)
	deletedModel.MoreInfo = core.StringPtr("https://cloud.ibm.com/apidocs/vpc#deleted-resources")

	model := new(vpcv1.SubnetReference)
	model.CRN = core.StringPtr("crn:v1:bluemix:public:is:us-south-1:a/aa2432b1fa4d4ace891e9b80fc104e34::subnet:0717-7ec86020-1c6e-4889-b3f0-a15f2e50f87e")
	model.Deleted = deletedModel
	model.Href = core.StringPtr("https://us-south.iaas.cloud.ibm.com/v1/subnets/0717-7ec86020-1c6e-4889-b3f0-a15f2e50f87e")
	model.ID = core.StringPtr("0717-7ec86020-1c6e-4889-b3f0-a15f2e50f87e")
	model.Name = core.StringPtr("my-subnet")
	model.ResourceType = core.StringPtr("subnet")

	result, err := vpc.ResourceIBMIsDynamicRouteServerPeerGroupPeerSubnetReferenceToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMIsDynamicRouteServerPeerGroupPeerDynamicRouteServerPeerGroupPeerReferenceToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		deletedModel := make(map[string]interface{})
		deletedModel["more_info"] = "https://cloud.ibm.com/apidocs/vpc#deleted-resources"

		model := make(map[string]interface{})
		model["deleted"] = []map[string]interface{}{deletedModel}
		model["href"] = "https://us-south.iaas.cloud.ibm.com/v1/dynamic_route_servers/r006-d7cc5196-9864-48c4-82d8-3f30da41fcc5/peer_groups/r006-8228af60-189d-11ed-861d-0242ac120002/peers/r006-8228af60-189d-11ed-861d-0242ac120004"
		model["id"] = "r006-8228af60-189d-11ed-861d-0242ac120004"
		model["name"] = "my-dynamic-route-server-peer-group-peer"
		model["resource_type"] = "dynamic_route_server_peer_group_peer"

		assert.Equal(t, result, model)
	}

	deletedModel := new(vpcv1.Deleted)
	deletedModel.MoreInfo = core.StringPtr("https://cloud.ibm.com/apidocs/vpc#deleted-resources")

	model := new(vpcv1.DynamicRouteServerPeerGroupPeerReference)
	model.Deleted = deletedModel
	model.Href = core.StringPtr("https://us-south.iaas.cloud.ibm.com/v1/dynamic_route_servers/r006-d7cc5196-9864-48c4-82d8-3f30da41fcc5/peer_groups/r006-8228af60-189d-11ed-861d-0242ac120002/peers/r006-8228af60-189d-11ed-861d-0242ac120004")
	model.ID = core.StringPtr("r006-8228af60-189d-11ed-861d-0242ac120004")
	model.Name = core.StringPtr("my-dynamic-route-server-peer-group-peer")
	model.ResourceType = core.StringPtr("dynamic_route_server_peer_group_peer")

	result, err := vpc.ResourceIBMIsDynamicRouteServerPeerGroupPeerDynamicRouteServerPeerGroupPeerReferenceToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMIsDynamicRouteServerPeerGroupPeerDynamicRouteServerPeerGroupPeerAddressSessionToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		deletedModel := make(map[string]interface{})
		deletedModel["more_info"] = "https://cloud.ibm.com/apidocs/vpc#deleted-resources"

		reservedIPReferenceModel := make(map[string]interface{})
		reservedIPReferenceModel["address"] = "192.168.3.4"
		reservedIPReferenceModel["deleted"] = []map[string]interface{}{deletedModel}
		reservedIPReferenceModel["href"] = "https://us-south.iaas.cloud.ibm.com/v1/subnets/0717-bea6a632-5e13-42a4-b4b8-31dc877abfe4/reserved_ips/0717-6d353a0f-aeb1-4ae1-832e-1110d10981bb"
		reservedIPReferenceModel["id"] = "0717-6d353a0f-aeb1-4ae1-832e-1110d10981bb"
		reservedIPReferenceModel["name"] = "my-reserved-ip"
		reservedIPReferenceModel["resource_type"] = "subnet_reserved_ip"

		subnetReferenceModel := make(map[string]interface{})
		subnetReferenceModel["crn"] = "crn:v1:bluemix:public:is:us-south-1:a/aa2432b1fa4d4ace891e9b80fc104e34::subnet:0717-7ec86020-1c6e-4889-b3f0-a15f2e50f87e"
		subnetReferenceModel["deleted"] = []map[string]interface{}{deletedModel}
		subnetReferenceModel["href"] = "https://us-south.iaas.cloud.ibm.com/v1/subnets/0717-7ec86020-1c6e-4889-b3f0-a15f2e50f87e"
		subnetReferenceModel["id"] = "0717-7ec86020-1c6e-4889-b3f0-a15f2e50f87e"
		subnetReferenceModel["name"] = "my-subnet"
		subnetReferenceModel["resource_type"] = "subnet"

		virtualNetworkInterfaceReferenceModel := make(map[string]interface{})
		virtualNetworkInterfaceReferenceModel["crn"] = "crn:v1:bluemix:public:is:us-south-1:a/aa2432b1fa4d4ace891e9b80fc104e34::virtual-network-interface:0717-54eb57ee-86f2-4796-90bb-d7874e0831ef"
		virtualNetworkInterfaceReferenceModel["deleted"] = []map[string]interface{}{deletedModel}
		virtualNetworkInterfaceReferenceModel["href"] = "https://us-south.iaas.cloud.ibm.com/v1/virtual_network_interfaces/0717-54eb57ee-86f2-4796-90bb-d7874e0831ef"
		virtualNetworkInterfaceReferenceModel["id"] = "0717-54eb57ee-86f2-4796-90bb-d7874e0831ef"
		virtualNetworkInterfaceReferenceModel["name"] = "my-virtual-network-interface"
		virtualNetworkInterfaceReferenceModel["primary_ip"] = []map[string]interface{}{reservedIPReferenceModel}
		virtualNetworkInterfaceReferenceModel["resource_type"] = "virtual_network_interface"
		virtualNetworkInterfaceReferenceModel["subnet"] = []map[string]interface{}{subnetReferenceModel}

		dynamicRouteServerMemberReferenceModel := make(map[string]interface{})
		dynamicRouteServerMemberReferenceModel["deleted"] = []map[string]interface{}{deletedModel}
		dynamicRouteServerMemberReferenceModel["href"] = "https://us-south.iaas.cloud.ibm.com/v1/dynamic_route_servers/r006-d7cc5196-9864-48c4-82d8-3f30da41fcc5/members/r006-d7cc5196-9924-48d4-82d8-3f30da41fdc"
		dynamicRouteServerMemberReferenceModel["id"] = "0717-9dce0ab3-a719-4582-ad71-00abfaae43ed"
		dynamicRouteServerMemberReferenceModel["name"] = "my-dynamic-route-server-member1"
		dynamicRouteServerMemberReferenceModel["resource_type"] = "dynamic_route_server_member"
		dynamicRouteServerMemberReferenceModel["virtual_network_interfaces"] = []map[string]interface{}{virtualNetworkInterfaceReferenceModel}

		dynamicRouteServerPeerGroupPeerNameModel := make(map[string]interface{})
		dynamicRouteServerPeerGroupPeerNameModel["name"] = "my-dynamic-route-server-peer-group-peer"

		model := make(map[string]interface{})
		model["established_at"] = "2026-01-02T03:04:05.006Z"
		model["local"] = []map[string]interface{}{dynamicRouteServerMemberReferenceModel}
		model["protocol_state"] = "active"
		model["remote"] = []map[string]interface{}{dynamicRouteServerPeerGroupPeerNameModel}

		assert.Equal(t, result, model)
	}

	deletedModel := new(vpcv1.Deleted)
	deletedModel.MoreInfo = core.StringPtr("https://cloud.ibm.com/apidocs/vpc#deleted-resources")

	reservedIPReferenceModel := new(vpcv1.ReservedIPReference)
	reservedIPReferenceModel.Address = core.StringPtr("192.168.3.4")
	reservedIPReferenceModel.Deleted = deletedModel
	reservedIPReferenceModel.Href = core.StringPtr("https://us-south.iaas.cloud.ibm.com/v1/subnets/0717-bea6a632-5e13-42a4-b4b8-31dc877abfe4/reserved_ips/0717-6d353a0f-aeb1-4ae1-832e-1110d10981bb")
	reservedIPReferenceModel.ID = core.StringPtr("0717-6d353a0f-aeb1-4ae1-832e-1110d10981bb")
	reservedIPReferenceModel.Name = core.StringPtr("my-reserved-ip")
	reservedIPReferenceModel.ResourceType = core.StringPtr("subnet_reserved_ip")

	subnetReferenceModel := new(vpcv1.SubnetReference)
	subnetReferenceModel.CRN = core.StringPtr("crn:v1:bluemix:public:is:us-south-1:a/aa2432b1fa4d4ace891e9b80fc104e34::subnet:0717-7ec86020-1c6e-4889-b3f0-a15f2e50f87e")
	subnetReferenceModel.Deleted = deletedModel
	subnetReferenceModel.Href = core.StringPtr("https://us-south.iaas.cloud.ibm.com/v1/subnets/0717-7ec86020-1c6e-4889-b3f0-a15f2e50f87e")
	subnetReferenceModel.ID = core.StringPtr("0717-7ec86020-1c6e-4889-b3f0-a15f2e50f87e")
	subnetReferenceModel.Name = core.StringPtr("my-subnet")
	subnetReferenceModel.ResourceType = core.StringPtr("subnet")

	virtualNetworkInterfaceReferenceModel := new(vpcv1.VirtualNetworkInterfaceReference)
	virtualNetworkInterfaceReferenceModel.CRN = core.StringPtr("crn:v1:bluemix:public:is:us-south-1:a/aa2432b1fa4d4ace891e9b80fc104e34::virtual-network-interface:0717-54eb57ee-86f2-4796-90bb-d7874e0831ef")
	virtualNetworkInterfaceReferenceModel.Deleted = deletedModel
	virtualNetworkInterfaceReferenceModel.Href = core.StringPtr("https://us-south.iaas.cloud.ibm.com/v1/virtual_network_interfaces/0717-54eb57ee-86f2-4796-90bb-d7874e0831ef")
	virtualNetworkInterfaceReferenceModel.ID = core.StringPtr("0717-54eb57ee-86f2-4796-90bb-d7874e0831ef")
	virtualNetworkInterfaceReferenceModel.Name = core.StringPtr("my-virtual-network-interface")
	virtualNetworkInterfaceReferenceModel.PrimaryIP = reservedIPReferenceModel
	virtualNetworkInterfaceReferenceModel.ResourceType = core.StringPtr("virtual_network_interface")
	virtualNetworkInterfaceReferenceModel.Subnet = subnetReferenceModel

	dynamicRouteServerMemberReferenceModel := new(vpcv1.DynamicRouteServerMemberReference)
	dynamicRouteServerMemberReferenceModel.Deleted = deletedModel
	dynamicRouteServerMemberReferenceModel.Href = core.StringPtr("https://us-south.iaas.cloud.ibm.com/v1/dynamic_route_servers/r006-d7cc5196-9864-48c4-82d8-3f30da41fcc5/members/r006-d7cc5196-9924-48d4-82d8-3f30da41fdc")
	dynamicRouteServerMemberReferenceModel.ID = core.StringPtr("0717-9dce0ab3-a719-4582-ad71-00abfaae43ed")
	dynamicRouteServerMemberReferenceModel.Name = core.StringPtr("my-dynamic-route-server-member1")
	dynamicRouteServerMemberReferenceModel.ResourceType = core.StringPtr("dynamic_route_server_member")
	dynamicRouteServerMemberReferenceModel.VirtualNetworkInterfaces = []vpcv1.VirtualNetworkInterfaceReference{*virtualNetworkInterfaceReferenceModel}

	dynamicRouteServerPeerGroupPeerNameModel := new(vpcv1.DynamicRouteServerPeerGroupPeerName)
	dynamicRouteServerPeerGroupPeerNameModel.Name = core.StringPtr("my-dynamic-route-server-peer-group-peer")

	model := new(vpcv1.DynamicRouteServerPeerGroupPeerAddressSession)
	model.EstablishedAt = CreateMockDateTime("2026-01-02T03:04:05.006Z")
	model.Local = dynamicRouteServerMemberReferenceModel
	model.ProtocolState = core.StringPtr("active")
	model.Remote = dynamicRouteServerPeerGroupPeerNameModel

	result, err := vpc.ResourceIBMIsDynamicRouteServerPeerGroupPeerDynamicRouteServerPeerGroupPeerAddressSessionToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMIsDynamicRouteServerPeerGroupPeerDynamicRouteServerPeerGroupPeerNameToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["name"] = "my-dynamic-route-server-peer-group-peer"

		assert.Equal(t, result, model)
	}

	model := new(vpcv1.DynamicRouteServerPeerGroupPeerName)
	model.Name = core.StringPtr("my-dynamic-route-server-peer-group-peer")

	result, err := vpc.ResourceIBMIsDynamicRouteServerPeerGroupPeerDynamicRouteServerPeerGroupPeerNameToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMIsDynamicRouteServerPeerGroupPeerMapToDynamicRouteServerPeerGroupPeerAddressEndpointPrototype(t *testing.T) {
	checkResult := func(result vpcv1.DynamicRouteServerPeerGroupPeerAddressEndpointPrototypeIntf) {
		model := new(vpcv1.DynamicRouteServerPeerGroupPeerAddressEndpointPrototype)
		model.ID = core.StringPtr("0717-6d353a0f-aeb1-4ae1-832e-1110d10981bb")
		model.Href = core.StringPtr("https://us-south.iaas.cloud.ibm.com/v1/subnets/0717-bea6a632-5e13-42a4-b4b8-31dc877abfe4/reserved_ips/0717-6d353a0f-aeb1-4ae1-832e-1110d10981bb")
		model.Address = core.StringPtr("192.168.3.4")
		// model.Gateway = dynamicRouteServerPeerGroupPeerAddressEndpointIPGatewayGatewayModel

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["id"] = "0717-6d353a0f-aeb1-4ae1-832e-1110d10981bb"
	model["href"] = "https://us-south.iaas.cloud.ibm.com/v1/subnets/0717-bea6a632-5e13-42a4-b4b8-31dc877abfe4/reserved_ips/0717-6d353a0f-aeb1-4ae1-832e-1110d10981bb"
	model["address"] = "192.168.3.4"
	// model["gateway"] = []interface{}{dynamicRouteServerPeerGroupPeerAddressEndpointIPGatewayGatewayModel}

	result, err := vpc.ResourceIBMIsDynamicRouteServerPeerGroupPeerMapToDynamicRouteServerPeerGroupPeerAddressEndpointPrototype(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMIsDynamicRouteServerPeerGroupPeerMapToDynamicRouteServerPeerGroupPeerAddressEndpointIPGatewayGateway(t *testing.T) {
	checkResult := func(result *vpcv1.DynamicRouteServerPeerGroupPeerAddressEndpointIPGatewayGateway) {
		model := new(vpcv1.DynamicRouteServerPeerGroupPeerAddressEndpointIPGatewayGateway)
		model.Address = core.StringPtr("192.168.3.4")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["address"] = "192.168.3.4"

	result, err := vpc.ResourceIBMIsDynamicRouteServerPeerGroupPeerMapToDynamicRouteServerPeerGroupPeerAddressEndpointIPGatewayGateway(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMIsDynamicRouteServerPeerGroupPeerMapToDynamicRouteServerPeerGroupPeerAddressEndpointPrototypeReservedIPIdentity(t *testing.T) {
	checkResult := func(result vpcv1.DynamicRouteServerPeerGroupPeerAddressEndpointPrototypeReservedIPIdentityIntf) {
		model := new(vpcv1.DynamicRouteServerPeerGroupPeerAddressEndpointPrototypeReservedIPIdentity)
		model.ID = core.StringPtr("0717-6d353a0f-aeb1-4ae1-832e-1110d10981bb")
		model.Href = core.StringPtr("https://us-south.iaas.cloud.ibm.com/v1/subnets/0717-bea6a632-5e13-42a4-b4b8-31dc877abfe4/reserved_ips/0717-6d353a0f-aeb1-4ae1-832e-1110d10981bb")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["id"] = "0717-6d353a0f-aeb1-4ae1-832e-1110d10981bb"
	model["href"] = "https://us-south.iaas.cloud.ibm.com/v1/subnets/0717-bea6a632-5e13-42a4-b4b8-31dc877abfe4/reserved_ips/0717-6d353a0f-aeb1-4ae1-832e-1110d10981bb"

	result, err := vpc.ResourceIBMIsDynamicRouteServerPeerGroupPeerMapToDynamicRouteServerPeerGroupPeerAddressEndpointPrototypeReservedIPIdentity(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMIsDynamicRouteServerPeerGroupPeerMapToDynamicRouteServerPeerGroupPeerAddressEndpointPrototypeReservedIPIdentityByID(t *testing.T) {
	checkResult := func(result *vpcv1.DynamicRouteServerPeerGroupPeerAddressEndpointPrototypeReservedIPIdentityByID) {
		model := new(vpcv1.DynamicRouteServerPeerGroupPeerAddressEndpointPrototypeReservedIPIdentityByID)
		model.ID = core.StringPtr("0717-6d353a0f-aeb1-4ae1-832e-1110d10981bb")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["id"] = "0717-6d353a0f-aeb1-4ae1-832e-1110d10981bb"

	result, err := vpc.ResourceIBMIsDynamicRouteServerPeerGroupPeerMapToDynamicRouteServerPeerGroupPeerAddressEndpointPrototypeReservedIPIdentityByID(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMIsDynamicRouteServerPeerGroupPeerMapToDynamicRouteServerPeerGroupPeerAddressEndpointPrototypeReservedIPIdentityByHref(t *testing.T) {
	checkResult := func(result *vpcv1.DynamicRouteServerPeerGroupPeerAddressEndpointPrototypeReservedIPIdentityByHref) {
		model := new(vpcv1.DynamicRouteServerPeerGroupPeerAddressEndpointPrototypeReservedIPIdentityByHref)
		model.Href = core.StringPtr("https://us-south.iaas.cloud.ibm.com/v1/subnets/0717-bea6a632-5e13-42a4-b4b8-31dc877abfe4/reserved_ips/0717-6d353a0f-aeb1-4ae1-832e-1110d10981bb")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["href"] = "https://us-south.iaas.cloud.ibm.com/v1/subnets/0717-bea6a632-5e13-42a4-b4b8-31dc877abfe4/reserved_ips/0717-6d353a0f-aeb1-4ae1-832e-1110d10981bb"

	result, err := vpc.ResourceIBMIsDynamicRouteServerPeerGroupPeerMapToDynamicRouteServerPeerGroupPeerAddressEndpointPrototypeReservedIPIdentityByHref(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMIsDynamicRouteServerPeerGroupPeerMapToDynamicRouteServerPeerGroupPeerAddressEndpointPrototypeDynamicRouteServerPeerGroupPeerAddressEndpointIPGatewayPrototype(t *testing.T) {
	checkResult := func(result *vpcv1.DynamicRouteServerPeerGroupPeerAddressEndpointPrototypeDynamicRouteServerPeerGroupPeerAddressEndpointIPGatewayPrototype) {
		dynamicRouteServerPeerGroupPeerAddressEndpointIPGatewayGatewayModel := new(vpcv1.DynamicRouteServerPeerGroupPeerAddressEndpointIPGatewayGateway)
		dynamicRouteServerPeerGroupPeerAddressEndpointIPGatewayGatewayModel.Address = core.StringPtr("192.168.3.4")

		model := new(vpcv1.DynamicRouteServerPeerGroupPeerAddressEndpointPrototypeDynamicRouteServerPeerGroupPeerAddressEndpointIPGatewayPrototype)
		model.Address = core.StringPtr("192.168.3.4")
		model.Gateway = dynamicRouteServerPeerGroupPeerAddressEndpointIPGatewayGatewayModel

		assert.Equal(t, result, model)
	}

	dynamicRouteServerPeerGroupPeerAddressEndpointIPGatewayGatewayModel := make(map[string]interface{})
	dynamicRouteServerPeerGroupPeerAddressEndpointIPGatewayGatewayModel["address"] = "192.168.3.4"

	model := make(map[string]interface{})
	model["address"] = "192.168.3.4"
	model["gateway"] = []interface{}{dynamicRouteServerPeerGroupPeerAddressEndpointIPGatewayGatewayModel}

	result, err := vpc.ResourceIBMIsDynamicRouteServerPeerGroupPeerMapToDynamicRouteServerPeerGroupPeerAddressEndpointPrototypeDynamicRouteServerPeerGroupPeerAddressEndpointIPGatewayPrototype(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMIsDynamicRouteServerPeerGroupPeerMapToDynamicRouteServerPeerGroupPeerPrototype(t *testing.T) {
	checkResult := func(result vpcv1.DynamicRouteServerPeerGroupPeerPrototypeIntf) {
		dynamicRouteServerPeerGroupPeerAddressEndpointPrototypeModel := new(vpcv1.DynamicRouteServerPeerGroupPeerAddressEndpointPrototypeReservedIPIdentityByID)
		dynamicRouteServerPeerGroupPeerAddressEndpointPrototypeModel.ID = core.StringPtr("0717-6d353a0f-aeb1-4ae1-832e-1110d10981bb")

		model := new(vpcv1.DynamicRouteServerPeerGroupPeerPrototype)
		model.Asn = core.Int64Ptr(int64(64520))
		model.Endpoint = dynamicRouteServerPeerGroupPeerAddressEndpointPrototypeModel
		model.Name = core.StringPtr("my-dynamic-route-server-peer-group-peer")
		model.Priority = core.Int64Ptr(int64(1))
		model.State = core.StringPtr("enabled")

		assert.Equal(t, result, model)
	}

	dynamicRouteServerPeerGroupPeerAddressEndpointPrototypeModel := make(map[string]interface{})
	dynamicRouteServerPeerGroupPeerAddressEndpointPrototypeModel["id"] = "0717-6d353a0f-aeb1-4ae1-832e-1110d10981bb"

	model := make(map[string]interface{})
	model["asn"] = int(64520)
	model["endpoint"] = []interface{}{dynamicRouteServerPeerGroupPeerAddressEndpointPrototypeModel}
	model["name"] = "my-dynamic-route-server-peer-group-peer"
	model["priority"] = int(1)
	model["state"] = "enabled"

	result, err := vpc.ResourceIBMIsDynamicRouteServerPeerGroupPeerMapToDynamicRouteServerPeerGroupPeerPrototype(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMIsDynamicRouteServerPeerGroupPeerMapToDynamicRouteServerPeerGroupPeerPrototypeDynamicRouteServerPeerGroupPeerAddressPrototype(t *testing.T) {
	checkResult := func(result *vpcv1.DynamicRouteServerPeerGroupPeerPrototypeDynamicRouteServerPeerGroupPeerAddressPrototype) {
		dynamicRouteServerPeerGroupPeerAddressEndpointPrototypeModel := new(vpcv1.DynamicRouteServerPeerGroupPeerAddressEndpointPrototypeReservedIPIdentityByID)
		dynamicRouteServerPeerGroupPeerAddressEndpointPrototypeModel.ID = core.StringPtr("0717-6d353a0f-aeb1-4ae1-832e-1110d10981bb")

		model := new(vpcv1.DynamicRouteServerPeerGroupPeerPrototypeDynamicRouteServerPeerGroupPeerAddressPrototype)
		model.Asn = core.Int64Ptr(int64(64520))
		model.Endpoint = dynamicRouteServerPeerGroupPeerAddressEndpointPrototypeModel
		model.Name = core.StringPtr("my-dynamic-route-server-peer-group-peer")
		model.Priority = core.Int64Ptr(int64(1))
		model.State = core.StringPtr("enabled")

		assert.Equal(t, result, model)
	}

	dynamicRouteServerPeerGroupPeerAddressEndpointPrototypeModel := make(map[string]interface{})
	dynamicRouteServerPeerGroupPeerAddressEndpointPrototypeModel["id"] = "0717-6d353a0f-aeb1-4ae1-832e-1110d10981bb"

	model := make(map[string]interface{})
	model["asn"] = int(64520)
	model["endpoint"] = []interface{}{dynamicRouteServerPeerGroupPeerAddressEndpointPrototypeModel}
	model["name"] = "my-dynamic-route-server-peer-group-peer"
	model["priority"] = int(1)
	model["state"] = "enabled"

	result, err := vpc.ResourceIBMIsDynamicRouteServerPeerGroupPeerMapToDynamicRouteServerPeerGroupPeerPrototypeDynamicRouteServerPeerGroupPeerAddressPrototype(model)
	assert.Nil(t, err)
	checkResult(result)
}
