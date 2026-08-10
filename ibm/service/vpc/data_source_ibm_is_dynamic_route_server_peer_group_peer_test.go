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

func TestAccIBMIsDynamicRouteServerPeerGroupPeerDataSourceBasic(t *testing.T) {
	dynamicRouteServerPeerGroupPeerDynamicRouteServerID := fmt.Sprintf("tf_dynamic_route_server_id_%d", acctest.RandIntRange(10, 100))
	dynamicRouteServerPeerGroupPeerDynamicRouteServerPeerGroupID := fmt.Sprintf("tf_dynamic_route_server_peer_group_id_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsDynamicRouteServerPeerGroupPeerDataSourceConfigBasic(dynamicRouteServerPeerGroupPeerDynamicRouteServerID, dynamicRouteServerPeerGroupPeerDynamicRouteServerPeerGroupID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_peer_group_peer.is_dynamic_route_server_peer_group_peer_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_peer_group_peer.is_dynamic_route_server_peer_group_peer_instance", "dynamic_route_server_id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_peer_group_peer.is_dynamic_route_server_peer_group_peer_instance", "dynamic_route_server_peer_group_id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_peer_group_peer.is_dynamic_route_server_peer_group_peer_instance", "is_dynamic_route_server_peer_group_peer_id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_peer_group_peer.is_dynamic_route_server_peer_group_peer_instance", "creator"),
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_peer_group_peer.is_dynamic_route_server_peer_group_peer_instance", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_peer_group_peer.is_dynamic_route_server_peer_group_peer_instance", "lifecycle_reasons.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_peer_group_peer.is_dynamic_route_server_peer_group_peer_instance", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_peer_group_peer.is_dynamic_route_server_peer_group_peer_instance", "resource_type"),
				),
			},
		},
	})
}

func TestAccIBMIsDynamicRouteServerPeerGroupPeerDataSourceAllArgs(t *testing.T) {
	dynamicRouteServerPeerGroupPeerDynamicRouteServerID := fmt.Sprintf("tf_dynamic_route_server_id_%d", acctest.RandIntRange(10, 100))
	dynamicRouteServerPeerGroupPeerDynamicRouteServerPeerGroupID := fmt.Sprintf("tf_dynamic_route_server_peer_group_id_%d", acctest.RandIntRange(10, 100))
	dynamicRouteServerPeerGroupPeerName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	dynamicRouteServerPeerGroupPeerAsn := fmt.Sprintf("%d", acctest.RandIntRange(10, 100))
	dynamicRouteServerPeerGroupPeerPriority := fmt.Sprintf("%d", acctest.RandIntRange(0, 4))
	dynamicRouteServerPeerGroupPeerState := "disabled"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsDynamicRouteServerPeerGroupPeerDataSourceConfig(dynamicRouteServerPeerGroupPeerDynamicRouteServerID, dynamicRouteServerPeerGroupPeerDynamicRouteServerPeerGroupID, dynamicRouteServerPeerGroupPeerName, dynamicRouteServerPeerGroupPeerAsn, dynamicRouteServerPeerGroupPeerPriority, dynamicRouteServerPeerGroupPeerState),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_peer_group_peer.is_dynamic_route_server_peer_group_peer_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_peer_group_peer.is_dynamic_route_server_peer_group_peer_instance", "dynamic_route_server_id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_peer_group_peer.is_dynamic_route_server_peer_group_peer_instance", "dynamic_route_server_peer_group_id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_peer_group_peer.is_dynamic_route_server_peer_group_peer_instance", "is_dynamic_route_server_peer_group_peer_id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_peer_group_peer.is_dynamic_route_server_peer_group_peer_instance", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_peer_group_peer.is_dynamic_route_server_peer_group_peer_instance", "creator"),
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_peer_group_peer.is_dynamic_route_server_peer_group_peer_instance", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_peer_group_peer.is_dynamic_route_server_peer_group_peer_instance", "lifecycle_reasons.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_peer_group_peer.is_dynamic_route_server_peer_group_peer_instance", "lifecycle_reasons.0.code"),
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_peer_group_peer.is_dynamic_route_server_peer_group_peer_instance", "lifecycle_reasons.0.message"),
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_peer_group_peer.is_dynamic_route_server_peer_group_peer_instance", "lifecycle_reasons.0.more_info"),
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_peer_group_peer.is_dynamic_route_server_peer_group_peer_instance", "lifecycle_state"),
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_peer_group_peer.is_dynamic_route_server_peer_group_peer_instance", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_peer_group_peer.is_dynamic_route_server_peer_group_peer_instance", "resource_type"),
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_peer_group_peer.is_dynamic_route_server_peer_group_peer_instance", "status"),
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_peer_group_peer.is_dynamic_route_server_peer_group_peer_instance", "status_reasons.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_peer_group_peer.is_dynamic_route_server_peer_group_peer_instance", "status_reasons.0.code"),
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_peer_group_peer.is_dynamic_route_server_peer_group_peer_instance", "status_reasons.0.message"),
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_peer_group_peer.is_dynamic_route_server_peer_group_peer_instance", "status_reasons.0.more_info"),
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_peer_group_peer.is_dynamic_route_server_peer_group_peer_instance", "asn"),
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_peer_group_peer.is_dynamic_route_server_peer_group_peer_instance", "bidirectional_forwarding_detection.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_peer_group_peer.is_dynamic_route_server_peer_group_peer_instance", "endpoint.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_peer_group_peer.is_dynamic_route_server_peer_group_peer_instance", "priority"),
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_peer_group_peer.is_dynamic_route_server_peer_group_peer_instance", "sessions.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_peer_group_peer.is_dynamic_route_server_peer_group_peer_instance", "sessions.0.established_at"),
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_peer_group_peer.is_dynamic_route_server_peer_group_peer_instance", "sessions.0.protocol_state"),
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_peer_group_peer.is_dynamic_route_server_peer_group_peer_instance", "state"),
				),
			},
		},
	})
}

func testAccCheckIBMIsDynamicRouteServerPeerGroupPeerDataSourceConfigBasic(dynamicRouteServerPeerGroupPeerDynamicRouteServerID string, dynamicRouteServerPeerGroupPeerDynamicRouteServerPeerGroupID string) string {
	return fmt.Sprintf(`
		resource "ibm_is_dynamic_route_server_peer_group_peer" "is_dynamic_route_server_peer_group_peer_instance" {
			dynamic_route_server_id = "%s"
			dynamic_route_server_peer_group_id = "%s"
		}

		data "ibm_is_dynamic_route_server_peer_group_peer" "is_dynamic_route_server_peer_group_peer_instance" {
			dynamic_route_server_id = ibm_is_dynamic_route_server_peer_group_peer.is_dynamic_route_server_peer_group_peer_instance.dynamic_route_server_id
			dynamic_route_server_peer_group_id = ibm_is_dynamic_route_server_peer_group_peer.is_dynamic_route_server_peer_group_peer_instance.dynamic_route_server_peer_group_id
			is_dynamic_route_server_peer_group_peer_id = ibm_is_dynamic_route_server_peer_group_peer.is_dynamic_route_server_peer_group_peer_instance.is_dynamic_route_server_peer_group_peer_id
		}
	`, dynamicRouteServerPeerGroupPeerDynamicRouteServerID, dynamicRouteServerPeerGroupPeerDynamicRouteServerPeerGroupID)
}

func testAccCheckIBMIsDynamicRouteServerPeerGroupPeerDataSourceConfig(dynamicRouteServerPeerGroupPeerDynamicRouteServerID string, dynamicRouteServerPeerGroupPeerDynamicRouteServerPeerGroupID string, dynamicRouteServerPeerGroupPeerName string, dynamicRouteServerPeerGroupPeerAsn string, dynamicRouteServerPeerGroupPeerPriority string, dynamicRouteServerPeerGroupPeerState string) string {
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

		data "ibm_is_dynamic_route_server_peer_group_peer" "is_dynamic_route_server_peer_group_peer_instance" {
			dynamic_route_server_id = ibm_is_dynamic_route_server_peer_group_peer.is_dynamic_route_server_peer_group_peer_instance.dynamic_route_server_id
			dynamic_route_server_peer_group_id = ibm_is_dynamic_route_server_peer_group_peer.is_dynamic_route_server_peer_group_peer_instance.dynamic_route_server_peer_group_id
			is_dynamic_route_server_peer_group_peer_id = ibm_is_dynamic_route_server_peer_group_peer.is_dynamic_route_server_peer_group_peer_instance.is_dynamic_route_server_peer_group_peer_id
		}
	`, dynamicRouteServerPeerGroupPeerDynamicRouteServerID, dynamicRouteServerPeerGroupPeerDynamicRouteServerPeerGroupID, dynamicRouteServerPeerGroupPeerName, dynamicRouteServerPeerGroupPeerAsn, dynamicRouteServerPeerGroupPeerPriority, dynamicRouteServerPeerGroupPeerState)
}

func TestDataSourceIBMIsDynamicRouteServerPeerGroupPeerLifecycleReasonToMap(t *testing.T) {
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

	result, err := vpc.DataSourceIBMIsDynamicRouteServerPeerGroupPeerLifecycleReasonToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMIsDynamicRouteServerPeerGroupPeerDynamicRouteServerPeerGroupPeerStatusReasonToMap(t *testing.T) {
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

	result, err := vpc.DataSourceIBMIsDynamicRouteServerPeerGroupPeerDynamicRouteServerPeerGroupPeerStatusReasonToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMIsDynamicRouteServerPeerGroupPeerDynamicRouteServerPeerGroupPeerAddressBidirectionalForwardingDetectionToMap(t *testing.T) {
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

	result, err := vpc.DataSourceIBMIsDynamicRouteServerPeerGroupPeerDynamicRouteServerPeerGroupPeerAddressBidirectionalForwardingDetectionToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMIsDynamicRouteServerPeerGroupPeerDynamicRouteServerPeerGroupPeerAddressBidirectionalForwardingDetectionSessionToMap(t *testing.T) {
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

	result, err := vpc.DataSourceIBMIsDynamicRouteServerPeerGroupPeerDynamicRouteServerPeerGroupPeerAddressBidirectionalForwardingDetectionSessionToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMIsDynamicRouteServerPeerGroupPeerDynamicRouteServerMemberReferenceToMap(t *testing.T) {
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

	result, err := vpc.DataSourceIBMIsDynamicRouteServerPeerGroupPeerDynamicRouteServerMemberReferenceToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMIsDynamicRouteServerPeerGroupPeerDeletedToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["more_info"] = "https://cloud.ibm.com/apidocs/vpc#deleted-resources"

		assert.Equal(t, result, model)
	}

	model := new(vpcv1.Deleted)
	model.MoreInfo = core.StringPtr("https://cloud.ibm.com/apidocs/vpc#deleted-resources")

	result, err := vpc.DataSourceIBMIsDynamicRouteServerPeerGroupPeerDeletedToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMIsDynamicRouteServerPeerGroupPeerVirtualNetworkInterfaceReferenceToMap(t *testing.T) {
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

	result, err := vpc.DataSourceIBMIsDynamicRouteServerPeerGroupPeerVirtualNetworkInterfaceReferenceToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMIsDynamicRouteServerPeerGroupPeerReservedIPReferenceToMap(t *testing.T) {
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

	result, err := vpc.DataSourceIBMIsDynamicRouteServerPeerGroupPeerReservedIPReferenceToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMIsDynamicRouteServerPeerGroupPeerSubnetReferenceToMap(t *testing.T) {
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

	result, err := vpc.DataSourceIBMIsDynamicRouteServerPeerGroupPeerSubnetReferenceToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMIsDynamicRouteServerPeerGroupPeerDynamicRouteServerPeerGroupPeerReferenceToMap(t *testing.T) {
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

	result, err := vpc.DataSourceIBMIsDynamicRouteServerPeerGroupPeerDynamicRouteServerPeerGroupPeerReferenceToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMIsDynamicRouteServerPeerGroupPeerDynamicRouteServerPeerGroupPeerAddressEndpointToMap(t *testing.T) {
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

	result, err := vpc.DataSourceIBMIsDynamicRouteServerPeerGroupPeerDynamicRouteServerPeerGroupPeerAddressEndpointToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMIsDynamicRouteServerPeerGroupPeerDynamicRouteServerPeerGroupPeerAddressEndpointIPGatewayGatewayToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["address"] = "192.168.3.4"

		assert.Equal(t, result, model)
	}

	model := new(vpcv1.DynamicRouteServerPeerGroupPeerAddressEndpointIPGatewayGateway)
	model.Address = core.StringPtr("192.168.3.4")

	result, err := vpc.DataSourceIBMIsDynamicRouteServerPeerGroupPeerDynamicRouteServerPeerGroupPeerAddressEndpointIPGatewayGatewayToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMIsDynamicRouteServerPeerGroupPeerDynamicRouteServerPeerGroupPeerAddressEndpointIPGatewayToMap(t *testing.T) {
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

	result, err := vpc.DataSourceIBMIsDynamicRouteServerPeerGroupPeerDynamicRouteServerPeerGroupPeerAddressEndpointIPGatewayToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMIsDynamicRouteServerPeerGroupPeerDynamicRouteServerPeerGroupPeerAddressEndpointByReservedIPToMap(t *testing.T) {
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

	result, err := vpc.DataSourceIBMIsDynamicRouteServerPeerGroupPeerDynamicRouteServerPeerGroupPeerAddressEndpointByReservedIPToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMIsDynamicRouteServerPeerGroupPeerIPToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["address"] = "192.168.3.4"

		assert.Equal(t, result, model)
	}

	model := new(vpcv1.IP)
	model.Address = core.StringPtr("192.168.3.4")

	result, err := vpc.DataSourceIBMIsDynamicRouteServerPeerGroupPeerIPToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMIsDynamicRouteServerPeerGroupPeerDynamicRouteServerPeerGroupPeerAddressSessionToMap(t *testing.T) {
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

	result, err := vpc.DataSourceIBMIsDynamicRouteServerPeerGroupPeerDynamicRouteServerPeerGroupPeerAddressSessionToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMIsDynamicRouteServerPeerGroupPeerDynamicRouteServerPeerGroupPeerNameToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["name"] = "my-dynamic-route-server-peer-group-peer"

		assert.Equal(t, result, model)
	}

	model := new(vpcv1.DynamicRouteServerPeerGroupPeerName)
	model.Name = core.StringPtr("my-dynamic-route-server-peer-group-peer")

	result, err := vpc.DataSourceIBMIsDynamicRouteServerPeerGroupPeerDynamicRouteServerPeerGroupPeerNameToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}
