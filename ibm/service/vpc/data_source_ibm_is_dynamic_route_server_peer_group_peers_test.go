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

func TestAccIBMIsDynamicRouteServerPeerGroupPeersDataSourceBasic(t *testing.T) {
	dynamicRouteServerPeerGroupPeerDynamicRouteServerID := fmt.Sprintf("tf_dynamic_route_server_id_%d", acctest.RandIntRange(10, 100))
	dynamicRouteServerPeerGroupPeerDynamicRouteServerPeerGroupID := fmt.Sprintf("tf_dynamic_route_server_peer_group_id_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsDynamicRouteServerPeerGroupPeersDataSourceConfigBasic(dynamicRouteServerPeerGroupPeerDynamicRouteServerID, dynamicRouteServerPeerGroupPeerDynamicRouteServerPeerGroupID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_peer_group_peers.is_dynamic_route_server_peer_group_peers_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_peer_group_peers.is_dynamic_route_server_peer_group_peers_instance", "dynamic_route_server_id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_peer_group_peers.is_dynamic_route_server_peer_group_peers_instance", "dynamic_route_server_peer_group_id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_peer_group_peers.is_dynamic_route_server_peer_group_peers_instance", "peers.#"),
				),
			},
		},
	})
}

func TestAccIBMIsDynamicRouteServerPeerGroupPeersDataSourceAllArgs(t *testing.T) {
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
				Config: testAccCheckIBMIsDynamicRouteServerPeerGroupPeersDataSourceConfig(dynamicRouteServerPeerGroupPeerDynamicRouteServerID, dynamicRouteServerPeerGroupPeerDynamicRouteServerPeerGroupID, dynamicRouteServerPeerGroupPeerName, dynamicRouteServerPeerGroupPeerAsn, dynamicRouteServerPeerGroupPeerPriority, dynamicRouteServerPeerGroupPeerState),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_peer_group_peers.is_dynamic_route_server_peer_group_peers_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_peer_group_peers.is_dynamic_route_server_peer_group_peers_instance", "dynamic_route_server_id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_peer_group_peers.is_dynamic_route_server_peer_group_peers_instance", "dynamic_route_server_peer_group_id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_peer_group_peers.is_dynamic_route_server_peer_group_peers_instance", "peers.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_peer_group_peers.is_dynamic_route_server_peer_group_peers_instance", "peers.0.created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_peer_group_peers.is_dynamic_route_server_peer_group_peers_instance", "peers.0.creator"),
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_peer_group_peers.is_dynamic_route_server_peer_group_peers_instance", "peers.0.href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_peer_group_peers.is_dynamic_route_server_peer_group_peers_instance", "peers.0.id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_peer_group_peers.is_dynamic_route_server_peer_group_peers_instance", "peers.0.lifecycle_state"),
					resource.TestCheckResourceAttr("data.ibm_is_dynamic_route_server_peer_group_peers.is_dynamic_route_server_peer_group_peers_instance", "peers.0.name", dynamicRouteServerPeerGroupPeerName),
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_peer_group_peers.is_dynamic_route_server_peer_group_peers_instance", "peers.0.resource_type"),
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_peer_group_peers.is_dynamic_route_server_peer_group_peers_instance", "peers.0.status"),
					resource.TestCheckResourceAttr("data.ibm_is_dynamic_route_server_peer_group_peers.is_dynamic_route_server_peer_group_peers_instance", "peers.0.asn", dynamicRouteServerPeerGroupPeerAsn),
					resource.TestCheckResourceAttr("data.ibm_is_dynamic_route_server_peer_group_peers.is_dynamic_route_server_peer_group_peers_instance", "peers.0.priority", dynamicRouteServerPeerGroupPeerPriority),
					resource.TestCheckResourceAttr("data.ibm_is_dynamic_route_server_peer_group_peers.is_dynamic_route_server_peer_group_peers_instance", "peers.0.state", dynamicRouteServerPeerGroupPeerState),
				),
			},
		},
	})
}

func testAccCheckIBMIsDynamicRouteServerPeerGroupPeersDataSourceConfigBasic(dynamicRouteServerPeerGroupPeerDynamicRouteServerID string, dynamicRouteServerPeerGroupPeerDynamicRouteServerPeerGroupID string) string {
	return fmt.Sprintf(`
		resource "ibm_is_dynamic_route_server_peer_group_peer" "is_dynamic_route_server_peer_group_peer_instance" {
			dynamic_route_server_id = "%s"
			dynamic_route_server_peer_group_id = "%s"
		}

		data "ibm_is_dynamic_route_server_peer_group_peers" "is_dynamic_route_server_peer_group_peers_instance" {
			dynamic_route_server_id = ibm_is_dynamic_route_server_peer_group_peer.is_dynamic_route_server_peer_group_peer_instance.dynamic_route_server_id
			dynamic_route_server_peer_group_id = ibm_is_dynamic_route_server_peer_group_peer.is_dynamic_route_server_peer_group_peer_instance.dynamic_route_server_peer_group_id
		}
	`, dynamicRouteServerPeerGroupPeerDynamicRouteServerID, dynamicRouteServerPeerGroupPeerDynamicRouteServerPeerGroupID)
}

func testAccCheckIBMIsDynamicRouteServerPeerGroupPeersDataSourceConfig(dynamicRouteServerPeerGroupPeerDynamicRouteServerID string, dynamicRouteServerPeerGroupPeerDynamicRouteServerPeerGroupID string, dynamicRouteServerPeerGroupPeerName string, dynamicRouteServerPeerGroupPeerAsn string, dynamicRouteServerPeerGroupPeerPriority string, dynamicRouteServerPeerGroupPeerState string) string {
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

		data "ibm_is_dynamic_route_server_peer_group_peers" "is_dynamic_route_server_peer_group_peers_instance" {
			dynamic_route_server_id = ibm_is_dynamic_route_server_peer_group_peer.is_dynamic_route_server_peer_group_peer_instance.dynamic_route_server_id
			dynamic_route_server_peer_group_id = ibm_is_dynamic_route_server_peer_group_peer.is_dynamic_route_server_peer_group_peer_instance.dynamic_route_server_peer_group_id
		}
	`, dynamicRouteServerPeerGroupPeerDynamicRouteServerID, dynamicRouteServerPeerGroupPeerDynamicRouteServerPeerGroupID, dynamicRouteServerPeerGroupPeerName, dynamicRouteServerPeerGroupPeerAsn, dynamicRouteServerPeerGroupPeerPriority, dynamicRouteServerPeerGroupPeerState)
}

func TestDataSourceIBMIsDynamicRouteServerPeerGroupPeersDynamicRouteServerPeerGroupPeerToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		lifecycleReasonModel := make(map[string]interface{})
		lifecycleReasonModel["code"] = "resource_suspended_by_provider"
		lifecycleReasonModel["message"] = "The resource has been suspended. Contact IBM support with the CRN for next steps."
		lifecycleReasonModel["more_info"] = "https://cloud.ibm.com/apidocs/vpc#resource-suspension"

		dynamicRouteServerPeerGroupPeerStatusReasonModel := make(map[string]interface{})
		dynamicRouteServerPeerGroupPeerStatusReasonModel["code"] = "peer_not_responding"
		dynamicRouteServerPeerGroupPeerStatusReasonModel["message"] = "The connection is down because the peer is not responding."
		dynamicRouteServerPeerGroupPeerStatusReasonModel["more_info"] = "https://cloud.ibm.com/docs/drs?topic=vpc-drs-health#__TBD__"

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

		dynamicRouteServerPeerGroupPeerAddressBidirectionalForwardingDetectionModel := make(map[string]interface{})
		dynamicRouteServerPeerGroupPeerAddressBidirectionalForwardingDetectionModel["detect_multiplier"] = int(3)
		dynamicRouteServerPeerGroupPeerAddressBidirectionalForwardingDetectionModel["enabled"] = true
		dynamicRouteServerPeerGroupPeerAddressBidirectionalForwardingDetectionModel["mode"] = "asynchronous"
		dynamicRouteServerPeerGroupPeerAddressBidirectionalForwardingDetectionModel["receive_interval"] = int(300)
		dynamicRouteServerPeerGroupPeerAddressBidirectionalForwardingDetectionModel["role"] = "active"
		dynamicRouteServerPeerGroupPeerAddressBidirectionalForwardingDetectionModel["sessions"] = []map[string]interface{}{dynamicRouteServerPeerGroupPeerAddressBidirectionalForwardingDetectionSessionModel}
		dynamicRouteServerPeerGroupPeerAddressBidirectionalForwardingDetectionModel["transmit_interval"] = int(300)

		dynamicRouteServerPeerGroupPeerAddressEndpointIPGatewayGatewayModel := make(map[string]interface{})
		dynamicRouteServerPeerGroupPeerAddressEndpointIPGatewayGatewayModel["address"] = "192.168.3.4"

		dynamicRouteServerPeerGroupPeerAddressEndpointModel := make(map[string]interface{})
		dynamicRouteServerPeerGroupPeerAddressEndpointModel["address"] = "192.168.3.4"
		dynamicRouteServerPeerGroupPeerAddressEndpointModel["gateway"] = []map[string]interface{}{dynamicRouteServerPeerGroupPeerAddressEndpointIPGatewayGatewayModel}

		dynamicRouteServerPeerGroupPeerNameModel := make(map[string]interface{})
		dynamicRouteServerPeerGroupPeerNameModel["name"] = "my-dynamic-route-server-peer-group-peer"

		dynamicRouteServerPeerGroupPeerAddressSessionModel := make(map[string]interface{})
		dynamicRouteServerPeerGroupPeerAddressSessionModel["established_at"] = "2026-01-02T03:04:05.006Z"
		dynamicRouteServerPeerGroupPeerAddressSessionModel["local"] = []map[string]interface{}{dynamicRouteServerMemberReferenceModel}
		dynamicRouteServerPeerGroupPeerAddressSessionModel["protocol_state"] = "active"
		dynamicRouteServerPeerGroupPeerAddressSessionModel["remote"] = []map[string]interface{}{dynamicRouteServerPeerGroupPeerNameModel}

		model := make(map[string]interface{})
		model["created_at"] = "2026-01-02T03:04:05.006Z"
		model["creator"] = "dynamic_route_server"
		model["href"] = "https://us-south.iaas.cloud.ibm.com/v1/dynamic_route_servers/r006-d7cc5196-9864-48c4-82d8-3f30da41fcc5/peer_groups/r006-8228af60-189d-11ed-861d-0242ac120002/peers/r006-8228af60-189d-11ed-861d-0242ac120004"
		model["id"] = "r006-8228af60-189d-11ed-861d-0242ac120004"
		model["lifecycle_reasons"] = []map[string]interface{}{lifecycleReasonModel}
		model["lifecycle_state"] = "stable"
		model["name"] = "my-dynamic-route-server-peer-group-peer"
		model["resource_type"] = "dynamic_route_server_peer_group_peer"
		model["status"] = "up"
		model["status_reasons"] = []map[string]interface{}{dynamicRouteServerPeerGroupPeerStatusReasonModel}
		model["asn"] = int(64520)
		model["bidirectional_forwarding_detection"] = []map[string]interface{}{dynamicRouteServerPeerGroupPeerAddressBidirectionalForwardingDetectionModel}
		model["endpoint"] = []map[string]interface{}{dynamicRouteServerPeerGroupPeerAddressEndpointModel}
		model["priority"] = int(1)
		model["sessions"] = []map[string]interface{}{dynamicRouteServerPeerGroupPeerAddressSessionModel}
		model["state"] = "enabled"

		assert.Equal(t, result, model)
	}

	lifecycleReasonModel := new(vpcv1.LifecycleReason)
	lifecycleReasonModel.Code = core.StringPtr("resource_suspended_by_provider")
	lifecycleReasonModel.Message = core.StringPtr("The resource has been suspended. Contact IBM support with the CRN for next steps.")
	lifecycleReasonModel.MoreInfo = core.StringPtr("https://cloud.ibm.com/apidocs/vpc#resource-suspension")

	dynamicRouteServerPeerGroupPeerStatusReasonModel := new(vpcv1.DynamicRouteServerPeerGroupPeerStatusReason)
	dynamicRouteServerPeerGroupPeerStatusReasonModel.Code = core.StringPtr("peer_not_responding")
	dynamicRouteServerPeerGroupPeerStatusReasonModel.Message = core.StringPtr("The connection is down because the peer is not responding.")
	dynamicRouteServerPeerGroupPeerStatusReasonModel.MoreInfo = core.StringPtr("https://cloud.ibm.com/docs/drs?topic=vpc-drs-health#__TBD__")

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

	dynamicRouteServerPeerGroupPeerAddressBidirectionalForwardingDetectionModel := new(vpcv1.DynamicRouteServerPeerGroupPeerAddressBidirectionalForwardingDetection)
	dynamicRouteServerPeerGroupPeerAddressBidirectionalForwardingDetectionModel.DetectMultiplier = core.Int64Ptr(int64(3))
	dynamicRouteServerPeerGroupPeerAddressBidirectionalForwardingDetectionModel.Enabled = core.BoolPtr(true)
	dynamicRouteServerPeerGroupPeerAddressBidirectionalForwardingDetectionModel.Mode = core.StringPtr("asynchronous")
	dynamicRouteServerPeerGroupPeerAddressBidirectionalForwardingDetectionModel.ReceiveInterval = core.Int64Ptr(int64(300))
	dynamicRouteServerPeerGroupPeerAddressBidirectionalForwardingDetectionModel.Role = core.StringPtr("active")
	dynamicRouteServerPeerGroupPeerAddressBidirectionalForwardingDetectionModel.Sessions = []vpcv1.DynamicRouteServerPeerGroupPeerAddressBidirectionalForwardingDetectionSession{*dynamicRouteServerPeerGroupPeerAddressBidirectionalForwardingDetectionSessionModel}
	dynamicRouteServerPeerGroupPeerAddressBidirectionalForwardingDetectionModel.TransmitInterval = core.Int64Ptr(int64(300))

	dynamicRouteServerPeerGroupPeerAddressEndpointIPGatewayGatewayModel := new(vpcv1.DynamicRouteServerPeerGroupPeerAddressEndpointIPGatewayGateway)
	dynamicRouteServerPeerGroupPeerAddressEndpointIPGatewayGatewayModel.Address = core.StringPtr("192.168.3.4")

	dynamicRouteServerPeerGroupPeerAddressEndpointModel := new(vpcv1.DynamicRouteServerPeerGroupPeerAddressEndpointIPGateway)
	dynamicRouteServerPeerGroupPeerAddressEndpointModel.Address = core.StringPtr("192.168.3.4")
	dynamicRouteServerPeerGroupPeerAddressEndpointModel.Gateway = dynamicRouteServerPeerGroupPeerAddressEndpointIPGatewayGatewayModel

	dynamicRouteServerPeerGroupPeerNameModel := new(vpcv1.DynamicRouteServerPeerGroupPeerName)
	dynamicRouteServerPeerGroupPeerNameModel.Name = core.StringPtr("my-dynamic-route-server-peer-group-peer")

	dynamicRouteServerPeerGroupPeerAddressSessionModel := new(vpcv1.DynamicRouteServerPeerGroupPeerAddressSession)
	dynamicRouteServerPeerGroupPeerAddressSessionModel.EstablishedAt = CreateMockDateTime("2026-01-02T03:04:05.006Z")
	dynamicRouteServerPeerGroupPeerAddressSessionModel.Local = dynamicRouteServerMemberReferenceModel
	dynamicRouteServerPeerGroupPeerAddressSessionModel.ProtocolState = core.StringPtr("active")
	dynamicRouteServerPeerGroupPeerAddressSessionModel.Remote = dynamicRouteServerPeerGroupPeerNameModel

	model := new(vpcv1.DynamicRouteServerPeerGroupPeer)
	model.CreatedAt = CreateMockDateTime("2026-01-02T03:04:05.006Z")
	model.Creator = core.StringPtr("dynamic_route_server")
	model.Href = core.StringPtr("https://us-south.iaas.cloud.ibm.com/v1/dynamic_route_servers/r006-d7cc5196-9864-48c4-82d8-3f30da41fcc5/peer_groups/r006-8228af60-189d-11ed-861d-0242ac120002/peers/r006-8228af60-189d-11ed-861d-0242ac120004")
	model.ID = core.StringPtr("r006-8228af60-189d-11ed-861d-0242ac120004")
	model.LifecycleReasons = []vpcv1.LifecycleReason{*lifecycleReasonModel}
	model.LifecycleState = core.StringPtr("stable")
	model.Name = core.StringPtr("my-dynamic-route-server-peer-group-peer")
	model.ResourceType = core.StringPtr("dynamic_route_server_peer_group_peer")
	model.Status = core.StringPtr("up")
	model.StatusReasons = []vpcv1.DynamicRouteServerPeerGroupPeerStatusReason{*dynamicRouteServerPeerGroupPeerStatusReasonModel}
	model.Asn = core.Int64Ptr(int64(64520))
	model.BidirectionalForwardingDetection = dynamicRouteServerPeerGroupPeerAddressBidirectionalForwardingDetectionModel
	model.Endpoint = dynamicRouteServerPeerGroupPeerAddressEndpointModel
	model.Priority = core.Int64Ptr(int64(1))
	model.Sessions = []vpcv1.DynamicRouteServerPeerGroupPeerAddressSession{*dynamicRouteServerPeerGroupPeerAddressSessionModel}
	model.State = core.StringPtr("enabled")

	result, err := vpc.DataSourceIBMIsDynamicRouteServerPeerGroupPeersDynamicRouteServerPeerGroupPeerToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMIsDynamicRouteServerPeerGroupPeersLifecycleReasonToMap(t *testing.T) {
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

	result, err := vpc.DataSourceIBMIsDynamicRouteServerPeerGroupPeersLifecycleReasonToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMIsDynamicRouteServerPeerGroupPeersDynamicRouteServerPeerGroupPeerStatusReasonToMap(t *testing.T) {
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

	result, err := vpc.DataSourceIBMIsDynamicRouteServerPeerGroupPeersDynamicRouteServerPeerGroupPeerStatusReasonToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMIsDynamicRouteServerPeerGroupPeersDynamicRouteServerPeerGroupPeerAddressBidirectionalForwardingDetectionToMap(t *testing.T) {
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

	result, err := vpc.DataSourceIBMIsDynamicRouteServerPeerGroupPeersDynamicRouteServerPeerGroupPeerAddressBidirectionalForwardingDetectionToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMIsDynamicRouteServerPeerGroupPeersDynamicRouteServerPeerGroupPeerAddressBidirectionalForwardingDetectionSessionToMap(t *testing.T) {
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

	result, err := vpc.DataSourceIBMIsDynamicRouteServerPeerGroupPeersDynamicRouteServerPeerGroupPeerAddressBidirectionalForwardingDetectionSessionToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMIsDynamicRouteServerPeerGroupPeersDynamicRouteServerMemberReferenceToMap(t *testing.T) {
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

	result, err := vpc.DataSourceIBMIsDynamicRouteServerPeerGroupPeersDynamicRouteServerMemberReferenceToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMIsDynamicRouteServerPeerGroupPeersDeletedToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["more_info"] = "https://cloud.ibm.com/apidocs/vpc#deleted-resources"

		assert.Equal(t, result, model)
	}

	model := new(vpcv1.Deleted)
	model.MoreInfo = core.StringPtr("https://cloud.ibm.com/apidocs/vpc#deleted-resources")

	result, err := vpc.DataSourceIBMIsDynamicRouteServerPeerGroupPeersDeletedToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMIsDynamicRouteServerPeerGroupPeersVirtualNetworkInterfaceReferenceToMap(t *testing.T) {
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

	result, err := vpc.DataSourceIBMIsDynamicRouteServerPeerGroupPeersVirtualNetworkInterfaceReferenceToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMIsDynamicRouteServerPeerGroupPeersReservedIPReferenceToMap(t *testing.T) {
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

	result, err := vpc.DataSourceIBMIsDynamicRouteServerPeerGroupPeersReservedIPReferenceToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMIsDynamicRouteServerPeerGroupPeersSubnetReferenceToMap(t *testing.T) {
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

	result, err := vpc.DataSourceIBMIsDynamicRouteServerPeerGroupPeersSubnetReferenceToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMIsDynamicRouteServerPeerGroupPeersDynamicRouteServerPeerGroupPeerReferenceToMap(t *testing.T) {
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

	result, err := vpc.DataSourceIBMIsDynamicRouteServerPeerGroupPeersDynamicRouteServerPeerGroupPeerReferenceToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMIsDynamicRouteServerPeerGroupPeersDynamicRouteServerPeerGroupPeerAddressEndpointToMap(t *testing.T) {
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

	result, err := vpc.DataSourceIBMIsDynamicRouteServerPeerGroupPeersDynamicRouteServerPeerGroupPeerAddressEndpointToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMIsDynamicRouteServerPeerGroupPeersDynamicRouteServerPeerGroupPeerAddressEndpointIPGatewayGatewayToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["address"] = "192.168.3.4"

		assert.Equal(t, result, model)
	}

	model := new(vpcv1.DynamicRouteServerPeerGroupPeerAddressEndpointIPGatewayGateway)
	model.Address = core.StringPtr("192.168.3.4")

	result, err := vpc.DataSourceIBMIsDynamicRouteServerPeerGroupPeersDynamicRouteServerPeerGroupPeerAddressEndpointIPGatewayGatewayToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMIsDynamicRouteServerPeerGroupPeersDynamicRouteServerPeerGroupPeerAddressEndpointIPGatewayToMap(t *testing.T) {
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

	result, err := vpc.DataSourceIBMIsDynamicRouteServerPeerGroupPeersDynamicRouteServerPeerGroupPeerAddressEndpointIPGatewayToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMIsDynamicRouteServerPeerGroupPeersDynamicRouteServerPeerGroupPeerAddressEndpointByReservedIPToMap(t *testing.T) {
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

	result, err := vpc.DataSourceIBMIsDynamicRouteServerPeerGroupPeersDynamicRouteServerPeerGroupPeerAddressEndpointByReservedIPToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMIsDynamicRouteServerPeerGroupPeersIPToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["address"] = "192.168.3.4"

		assert.Equal(t, result, model)
	}

	model := new(vpcv1.IP)
	model.Address = core.StringPtr("192.168.3.4")

	result, err := vpc.DataSourceIBMIsDynamicRouteServerPeerGroupPeersIPToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMIsDynamicRouteServerPeerGroupPeersDynamicRouteServerPeerGroupPeerAddressSessionToMap(t *testing.T) {
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

	result, err := vpc.DataSourceIBMIsDynamicRouteServerPeerGroupPeersDynamicRouteServerPeerGroupPeerAddressSessionToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMIsDynamicRouteServerPeerGroupPeersDynamicRouteServerPeerGroupPeerNameToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["name"] = "my-dynamic-route-server-peer-group-peer"

		assert.Equal(t, result, model)
	}

	model := new(vpcv1.DynamicRouteServerPeerGroupPeerName)
	model.Name = core.StringPtr("my-dynamic-route-server-peer-group-peer")

	result, err := vpc.DataSourceIBMIsDynamicRouteServerPeerGroupPeersDynamicRouteServerPeerGroupPeerNameToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMIsDynamicRouteServerPeerGroupPeersDynamicRouteServerPeerGroupPeerAddressToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		lifecycleReasonModel := make(map[string]interface{})
		lifecycleReasonModel["code"] = "resource_suspended_by_provider"
		lifecycleReasonModel["message"] = "The resource has been suspended. Contact IBM support with the CRN for next steps."
		lifecycleReasonModel["more_info"] = "https://cloud.ibm.com/apidocs/vpc#resource-suspension"

		dynamicRouteServerPeerGroupPeerStatusReasonModel := make(map[string]interface{})
		dynamicRouteServerPeerGroupPeerStatusReasonModel["code"] = "peer_not_responding"
		dynamicRouteServerPeerGroupPeerStatusReasonModel["message"] = "The connection is down because the peer is not responding."
		dynamicRouteServerPeerGroupPeerStatusReasonModel["more_info"] = "https://cloud.ibm.com/docs/drs?topic=vpc-drs-health#__TBD__"

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

		dynamicRouteServerPeerGroupPeerAddressBidirectionalForwardingDetectionModel := make(map[string]interface{})
		dynamicRouteServerPeerGroupPeerAddressBidirectionalForwardingDetectionModel["detect_multiplier"] = int(3)
		dynamicRouteServerPeerGroupPeerAddressBidirectionalForwardingDetectionModel["enabled"] = true
		dynamicRouteServerPeerGroupPeerAddressBidirectionalForwardingDetectionModel["mode"] = "asynchronous"
		dynamicRouteServerPeerGroupPeerAddressBidirectionalForwardingDetectionModel["receive_interval"] = int(300)
		dynamicRouteServerPeerGroupPeerAddressBidirectionalForwardingDetectionModel["role"] = "active"
		dynamicRouteServerPeerGroupPeerAddressBidirectionalForwardingDetectionModel["sessions"] = []map[string]interface{}{dynamicRouteServerPeerGroupPeerAddressBidirectionalForwardingDetectionSessionModel}
		dynamicRouteServerPeerGroupPeerAddressBidirectionalForwardingDetectionModel["transmit_interval"] = int(300)

		dynamicRouteServerPeerGroupPeerAddressEndpointIPGatewayGatewayModel := make(map[string]interface{})
		dynamicRouteServerPeerGroupPeerAddressEndpointIPGatewayGatewayModel["address"] = "192.168.3.4"

		dynamicRouteServerPeerGroupPeerAddressEndpointModel := make(map[string]interface{})
		dynamicRouteServerPeerGroupPeerAddressEndpointModel["address"] = "192.168.3.4"
		dynamicRouteServerPeerGroupPeerAddressEndpointModel["gateway"] = []map[string]interface{}{dynamicRouteServerPeerGroupPeerAddressEndpointIPGatewayGatewayModel}

		dynamicRouteServerPeerGroupPeerNameModel := make(map[string]interface{})
		dynamicRouteServerPeerGroupPeerNameModel["name"] = "my-dynamic-route-server-peer-group-peer"

		dynamicRouteServerPeerGroupPeerAddressSessionModel := make(map[string]interface{})
		dynamicRouteServerPeerGroupPeerAddressSessionModel["established_at"] = "2026-01-02T03:04:05.006Z"
		dynamicRouteServerPeerGroupPeerAddressSessionModel["local"] = []map[string]interface{}{dynamicRouteServerMemberReferenceModel}
		dynamicRouteServerPeerGroupPeerAddressSessionModel["protocol_state"] = "active"
		dynamicRouteServerPeerGroupPeerAddressSessionModel["remote"] = []map[string]interface{}{dynamicRouteServerPeerGroupPeerNameModel}

		model := make(map[string]interface{})
		model["created_at"] = "2026-01-02T03:04:05.006Z"
		model["href"] = "https://us-south.iaas.cloud.ibm.com/v1/dynamic_route_servers/r006-d7cc5196-9864-48c4-82d8-3f30da41fcc5/peer_groups/r006-8228af60-189d-11ed-861d-0242ac120002/peers/r006-8228af60-189d-11ed-861d-0242ac120004"
		model["id"] = "r006-8228af60-189d-11ed-861d-0242ac120004"
		model["lifecycle_reasons"] = []map[string]interface{}{lifecycleReasonModel}
		model["lifecycle_state"] = "stable"
		model["name"] = "my-dynamic-route-server-peer-group-peer"
		model["resource_type"] = "dynamic_route_server_peer_group_peer"
		model["status"] = "up"
		model["status_reasons"] = []map[string]interface{}{dynamicRouteServerPeerGroupPeerStatusReasonModel}
		model["asn"] = int(64520)
		model["bidirectional_forwarding_detection"] = []map[string]interface{}{dynamicRouteServerPeerGroupPeerAddressBidirectionalForwardingDetectionModel}
		model["creator"] = "dynamic_route_server"
		model["endpoint"] = []map[string]interface{}{dynamicRouteServerPeerGroupPeerAddressEndpointModel}
		model["priority"] = int(1)
		model["sessions"] = []map[string]interface{}{dynamicRouteServerPeerGroupPeerAddressSessionModel}
		model["state"] = "enabled"

		assert.Equal(t, result, model)
	}

	lifecycleReasonModel := new(vpcv1.LifecycleReason)
	lifecycleReasonModel.Code = core.StringPtr("resource_suspended_by_provider")
	lifecycleReasonModel.Message = core.StringPtr("The resource has been suspended. Contact IBM support with the CRN for next steps.")
	lifecycleReasonModel.MoreInfo = core.StringPtr("https://cloud.ibm.com/apidocs/vpc#resource-suspension")

	dynamicRouteServerPeerGroupPeerStatusReasonModel := new(vpcv1.DynamicRouteServerPeerGroupPeerStatusReason)
	dynamicRouteServerPeerGroupPeerStatusReasonModel.Code = core.StringPtr("peer_not_responding")
	dynamicRouteServerPeerGroupPeerStatusReasonModel.Message = core.StringPtr("The connection is down because the peer is not responding.")
	dynamicRouteServerPeerGroupPeerStatusReasonModel.MoreInfo = core.StringPtr("https://cloud.ibm.com/docs/drs?topic=vpc-drs-health#__TBD__")

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

	dynamicRouteServerPeerGroupPeerAddressBidirectionalForwardingDetectionModel := new(vpcv1.DynamicRouteServerPeerGroupPeerAddressBidirectionalForwardingDetection)
	dynamicRouteServerPeerGroupPeerAddressBidirectionalForwardingDetectionModel.DetectMultiplier = core.Int64Ptr(int64(3))
	dynamicRouteServerPeerGroupPeerAddressBidirectionalForwardingDetectionModel.Enabled = core.BoolPtr(true)
	dynamicRouteServerPeerGroupPeerAddressBidirectionalForwardingDetectionModel.Mode = core.StringPtr("asynchronous")
	dynamicRouteServerPeerGroupPeerAddressBidirectionalForwardingDetectionModel.ReceiveInterval = core.Int64Ptr(int64(300))
	dynamicRouteServerPeerGroupPeerAddressBidirectionalForwardingDetectionModel.Role = core.StringPtr("active")
	dynamicRouteServerPeerGroupPeerAddressBidirectionalForwardingDetectionModel.Sessions = []vpcv1.DynamicRouteServerPeerGroupPeerAddressBidirectionalForwardingDetectionSession{*dynamicRouteServerPeerGroupPeerAddressBidirectionalForwardingDetectionSessionModel}
	dynamicRouteServerPeerGroupPeerAddressBidirectionalForwardingDetectionModel.TransmitInterval = core.Int64Ptr(int64(300))

	dynamicRouteServerPeerGroupPeerAddressEndpointIPGatewayGatewayModel := new(vpcv1.DynamicRouteServerPeerGroupPeerAddressEndpointIPGatewayGateway)
	dynamicRouteServerPeerGroupPeerAddressEndpointIPGatewayGatewayModel.Address = core.StringPtr("192.168.3.4")

	dynamicRouteServerPeerGroupPeerAddressEndpointModel := new(vpcv1.DynamicRouteServerPeerGroupPeerAddressEndpointIPGateway)
	dynamicRouteServerPeerGroupPeerAddressEndpointModel.Address = core.StringPtr("192.168.3.4")
	dynamicRouteServerPeerGroupPeerAddressEndpointModel.Gateway = dynamicRouteServerPeerGroupPeerAddressEndpointIPGatewayGatewayModel

	dynamicRouteServerPeerGroupPeerNameModel := new(vpcv1.DynamicRouteServerPeerGroupPeerName)
	dynamicRouteServerPeerGroupPeerNameModel.Name = core.StringPtr("my-dynamic-route-server-peer-group-peer")

	dynamicRouteServerPeerGroupPeerAddressSessionModel := new(vpcv1.DynamicRouteServerPeerGroupPeerAddressSession)
	dynamicRouteServerPeerGroupPeerAddressSessionModel.EstablishedAt = CreateMockDateTime("2026-01-02T03:04:05.006Z")
	dynamicRouteServerPeerGroupPeerAddressSessionModel.Local = dynamicRouteServerMemberReferenceModel
	dynamicRouteServerPeerGroupPeerAddressSessionModel.ProtocolState = core.StringPtr("active")
	dynamicRouteServerPeerGroupPeerAddressSessionModel.Remote = dynamicRouteServerPeerGroupPeerNameModel

	model := new(vpcv1.DynamicRouteServerPeerGroupPeerAddress)
	model.CreatedAt = CreateMockDateTime("2026-01-02T03:04:05.006Z")
	model.Href = core.StringPtr("https://us-south.iaas.cloud.ibm.com/v1/dynamic_route_servers/r006-d7cc5196-9864-48c4-82d8-3f30da41fcc5/peer_groups/r006-8228af60-189d-11ed-861d-0242ac120002/peers/r006-8228af60-189d-11ed-861d-0242ac120004")
	model.ID = core.StringPtr("r006-8228af60-189d-11ed-861d-0242ac120004")
	model.LifecycleReasons = []vpcv1.LifecycleReason{*lifecycleReasonModel}
	model.LifecycleState = core.StringPtr("stable")
	model.Name = core.StringPtr("my-dynamic-route-server-peer-group-peer")
	model.ResourceType = core.StringPtr("dynamic_route_server_peer_group_peer")
	model.Status = core.StringPtr("up")
	model.StatusReasons = []vpcv1.DynamicRouteServerPeerGroupPeerStatusReason{*dynamicRouteServerPeerGroupPeerStatusReasonModel}
	model.Asn = core.Int64Ptr(int64(64520))
	model.BidirectionalForwardingDetection = dynamicRouteServerPeerGroupPeerAddressBidirectionalForwardingDetectionModel
	model.Creator = core.StringPtr("dynamic_route_server")
	model.Endpoint = dynamicRouteServerPeerGroupPeerAddressEndpointModel
	model.Priority = core.Int64Ptr(int64(1))
	model.Sessions = []vpcv1.DynamicRouteServerPeerGroupPeerAddressSession{*dynamicRouteServerPeerGroupPeerAddressSessionModel}
	model.State = core.StringPtr("enabled")

	result, err := vpc.DataSourceIBMIsDynamicRouteServerPeerGroupPeersDynamicRouteServerPeerGroupPeerAddressToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMIsDynamicRouteServerPeerGroupPeersDynamicRouteServerPeerGroupPeerTransitGatewayToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		lifecycleReasonModel := make(map[string]interface{})
		lifecycleReasonModel["code"] = "resource_suspended_by_provider"
		lifecycleReasonModel["message"] = "The resource has been suspended. Contact IBM support with the CRN for next steps."
		lifecycleReasonModel["more_info"] = "https://cloud.ibm.com/apidocs/vpc#resource-suspension"

		dynamicRouteServerPeerGroupPeerStatusReasonModel := make(map[string]interface{})
		dynamicRouteServerPeerGroupPeerStatusReasonModel["code"] = "peer_not_responding"
		dynamicRouteServerPeerGroupPeerStatusReasonModel["message"] = "The connection is down because the peer is not responding."
		dynamicRouteServerPeerGroupPeerStatusReasonModel["more_info"] = "https://cloud.ibm.com/docs/drs?topic=vpc-drs-health#__TBD__"

		dynamicRouteServerPeerGroupPeerTransitGatewayEndpointModel := make(map[string]interface{})
		dynamicRouteServerPeerGroupPeerTransitGatewayEndpointModel["crn"] = "crn:v1:bluemix:public:transit:dal03:a/aa2432b1fa4d4ace891e9b80fc104e34::gateway:ef4dcb1a-fee4-41c7-9e11-9cd99e65c1f4"
		dynamicRouteServerPeerGroupPeerTransitGatewayEndpointModel["id"] = "ef4dcb1a-fee4-41c7-9e11-9cd99e65c1f4"
		dynamicRouteServerPeerGroupPeerTransitGatewayEndpointModel["resource_type"] = "transit_gateway"

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

		dynamicRouteServerPeerGroupPeerTransitGatewaySessionRemoteModel := make(map[string]interface{})
		dynamicRouteServerPeerGroupPeerTransitGatewaySessionRemoteModel["name"] = "my-transit-gateway-dynamic-route-server-connection-1"

		dynamicRouteServerPeerGroupPeerTransitGatewaySessionModel := make(map[string]interface{})
		dynamicRouteServerPeerGroupPeerTransitGatewaySessionModel["established_at"] = "2026-01-02T03:04:05.006Z"
		dynamicRouteServerPeerGroupPeerTransitGatewaySessionModel["local"] = []map[string]interface{}{dynamicRouteServerMemberReferenceModel}
		dynamicRouteServerPeerGroupPeerTransitGatewaySessionModel["protocol_state"] = "active"
		dynamicRouteServerPeerGroupPeerTransitGatewaySessionModel["remote"] = []map[string]interface{}{dynamicRouteServerPeerGroupPeerTransitGatewaySessionRemoteModel}

		model := make(map[string]interface{})
		model["created_at"] = "2026-01-02T03:04:05.006Z"
		model["href"] = "https://us-south.iaas.cloud.ibm.com/v1/dynamic_route_servers/r006-d7cc5196-9864-48c4-82d8-3f30da41fcc5/peer_groups/r006-8228af60-189d-11ed-861d-0242ac120002/peers/r006-8228af60-189d-11ed-861d-0242ac120004"
		model["id"] = "r006-8228af60-189d-11ed-861d-0242ac120004"
		model["lifecycle_reasons"] = []map[string]interface{}{lifecycleReasonModel}
		model["lifecycle_state"] = "stable"
		model["name"] = "my-dynamic-route-server-peer-group-peer"
		model["resource_type"] = "dynamic_route_server_peer_group_peer"
		model["status"] = "up"
		model["status_reasons"] = []map[string]interface{}{dynamicRouteServerPeerGroupPeerStatusReasonModel}
		model["creator"] = "transit_gateway"
		model["endpoint"] = []map[string]interface{}{dynamicRouteServerPeerGroupPeerTransitGatewayEndpointModel}
		model["sessions"] = []map[string]interface{}{dynamicRouteServerPeerGroupPeerTransitGatewaySessionModel}
		model["state"] = "enabled"

		assert.Equal(t, result, model)
	}

	lifecycleReasonModel := new(vpcv1.LifecycleReason)
	lifecycleReasonModel.Code = core.StringPtr("resource_suspended_by_provider")
	lifecycleReasonModel.Message = core.StringPtr("The resource has been suspended. Contact IBM support with the CRN for next steps.")
	lifecycleReasonModel.MoreInfo = core.StringPtr("https://cloud.ibm.com/apidocs/vpc#resource-suspension")

	dynamicRouteServerPeerGroupPeerStatusReasonModel := new(vpcv1.DynamicRouteServerPeerGroupPeerStatusReason)
	dynamicRouteServerPeerGroupPeerStatusReasonModel.Code = core.StringPtr("peer_not_responding")
	dynamicRouteServerPeerGroupPeerStatusReasonModel.Message = core.StringPtr("The connection is down because the peer is not responding.")
	dynamicRouteServerPeerGroupPeerStatusReasonModel.MoreInfo = core.StringPtr("https://cloud.ibm.com/docs/drs?topic=vpc-drs-health#__TBD__")

	dynamicRouteServerPeerGroupPeerTransitGatewayEndpointModel := new(vpcv1.DynamicRouteServerPeerGroupPeerTransitGatewayEndpointTransitGatewayReference)
	dynamicRouteServerPeerGroupPeerTransitGatewayEndpointModel.CRN = core.StringPtr("crn:v1:bluemix:public:transit:dal03:a/aa2432b1fa4d4ace891e9b80fc104e34::gateway:ef4dcb1a-fee4-41c7-9e11-9cd99e65c1f4")
	dynamicRouteServerPeerGroupPeerTransitGatewayEndpointModel.ID = core.StringPtr("ef4dcb1a-fee4-41c7-9e11-9cd99e65c1f4")
	dynamicRouteServerPeerGroupPeerTransitGatewayEndpointModel.ResourceType = core.StringPtr("transit_gateway")

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

	dynamicRouteServerPeerGroupPeerTransitGatewaySessionRemoteModel := new(vpcv1.DynamicRouteServerPeerGroupPeerTransitGatewaySessionRemote)
	dynamicRouteServerPeerGroupPeerTransitGatewaySessionRemoteModel.Name = core.StringPtr("my-transit-gateway-dynamic-route-server-connection-1")

	dynamicRouteServerPeerGroupPeerTransitGatewaySessionModel := new(vpcv1.DynamicRouteServerPeerGroupPeerTransitGatewaySession)
	dynamicRouteServerPeerGroupPeerTransitGatewaySessionModel.EstablishedAt = CreateMockDateTime("2026-01-02T03:04:05.006Z")
	dynamicRouteServerPeerGroupPeerTransitGatewaySessionModel.Local = dynamicRouteServerMemberReferenceModel
	dynamicRouteServerPeerGroupPeerTransitGatewaySessionModel.ProtocolState = core.StringPtr("active")
	dynamicRouteServerPeerGroupPeerTransitGatewaySessionModel.Remote = dynamicRouteServerPeerGroupPeerTransitGatewaySessionRemoteModel

	model := new(vpcv1.DynamicRouteServerPeerGroupPeerTransitGateway)
	model.CreatedAt = CreateMockDateTime("2026-01-02T03:04:05.006Z")
	model.Href = core.StringPtr("https://us-south.iaas.cloud.ibm.com/v1/dynamic_route_servers/r006-d7cc5196-9864-48c4-82d8-3f30da41fcc5/peer_groups/r006-8228af60-189d-11ed-861d-0242ac120002/peers/r006-8228af60-189d-11ed-861d-0242ac120004")
	model.ID = core.StringPtr("r006-8228af60-189d-11ed-861d-0242ac120004")
	model.LifecycleReasons = []vpcv1.LifecycleReason{*lifecycleReasonModel}
	model.LifecycleState = core.StringPtr("stable")
	model.Name = core.StringPtr("my-dynamic-route-server-peer-group-peer")
	model.ResourceType = core.StringPtr("dynamic_route_server_peer_group_peer")
	model.Status = core.StringPtr("up")
	model.StatusReasons = []vpcv1.DynamicRouteServerPeerGroupPeerStatusReason{*dynamicRouteServerPeerGroupPeerStatusReasonModel}
	model.Creator = core.StringPtr("transit_gateway")
	model.Endpoint = dynamicRouteServerPeerGroupPeerTransitGatewayEndpointModel
	model.Sessions = []vpcv1.DynamicRouteServerPeerGroupPeerTransitGatewaySession{*dynamicRouteServerPeerGroupPeerTransitGatewaySessionModel}
	model.State = core.StringPtr("enabled")

	result, err := vpc.DataSourceIBMIsDynamicRouteServerPeerGroupPeersDynamicRouteServerPeerGroupPeerTransitGatewayToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMIsDynamicRouteServerPeerGroupPeersDynamicRouteServerPeerGroupPeerTransitGatewayEndpointToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["crn"] = "crn:v1:bluemix:public:transit:dal03:a/aa2432b1fa4d4ace891e9b80fc104e34::gateway:ef4dcb1a-fee4-41c7-9e11-9cd99e65c1f4"
		model["id"] = "ef4dcb1a-fee4-41c7-9e11-9cd99e65c1f4"
		model["resource_type"] = "transit_gateway"

		assert.Equal(t, result, model)
	}

	model := new(vpcv1.DynamicRouteServerPeerGroupPeerTransitGatewayEndpoint)
	model.CRN = core.StringPtr("crn:v1:bluemix:public:transit:dal03:a/aa2432b1fa4d4ace891e9b80fc104e34::gateway:ef4dcb1a-fee4-41c7-9e11-9cd99e65c1f4")
	model.ID = core.StringPtr("ef4dcb1a-fee4-41c7-9e11-9cd99e65c1f4")
	model.ResourceType = core.StringPtr("transit_gateway")

	result, err := vpc.DataSourceIBMIsDynamicRouteServerPeerGroupPeersDynamicRouteServerPeerGroupPeerTransitGatewayEndpointToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMIsDynamicRouteServerPeerGroupPeersDynamicRouteServerPeerGroupPeerTransitGatewayEndpointTransitGatewayReferenceToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["crn"] = "crn:v1:bluemix:public:transit:dal03:a/aa2432b1fa4d4ace891e9b80fc104e34::gateway:ef4dcb1a-fee4-41c7-9e11-9cd99e65c1f4"
		model["id"] = "ef4dcb1a-fee4-41c7-9e11-9cd99e65c1f4"
		model["resource_type"] = "transit_gateway"

		assert.Equal(t, result, model)
	}

	model := new(vpcv1.DynamicRouteServerPeerGroupPeerTransitGatewayEndpointTransitGatewayReference)
	model.CRN = core.StringPtr("crn:v1:bluemix:public:transit:dal03:a/aa2432b1fa4d4ace891e9b80fc104e34::gateway:ef4dcb1a-fee4-41c7-9e11-9cd99e65c1f4")
	model.ID = core.StringPtr("ef4dcb1a-fee4-41c7-9e11-9cd99e65c1f4")
	model.ResourceType = core.StringPtr("transit_gateway")

	result, err := vpc.DataSourceIBMIsDynamicRouteServerPeerGroupPeersDynamicRouteServerPeerGroupPeerTransitGatewayEndpointTransitGatewayReferenceToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMIsDynamicRouteServerPeerGroupPeersDynamicRouteServerPeerGroupPeerTransitGatewaySessionToMap(t *testing.T) {
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

		dynamicRouteServerPeerGroupPeerTransitGatewaySessionRemoteModel := make(map[string]interface{})
		dynamicRouteServerPeerGroupPeerTransitGatewaySessionRemoteModel["name"] = "my-transit-gateway-dynamic-route-server-connection-1"

		model := make(map[string]interface{})
		model["established_at"] = "2026-01-02T03:04:05.006Z"
		model["local"] = []map[string]interface{}{dynamicRouteServerMemberReferenceModel}
		model["protocol_state"] = "active"
		model["remote"] = []map[string]interface{}{dynamicRouteServerPeerGroupPeerTransitGatewaySessionRemoteModel}

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

	dynamicRouteServerPeerGroupPeerTransitGatewaySessionRemoteModel := new(vpcv1.DynamicRouteServerPeerGroupPeerTransitGatewaySessionRemote)
	dynamicRouteServerPeerGroupPeerTransitGatewaySessionRemoteModel.Name = core.StringPtr("my-transit-gateway-dynamic-route-server-connection-1")

	model := new(vpcv1.DynamicRouteServerPeerGroupPeerTransitGatewaySession)
	model.EstablishedAt = CreateMockDateTime("2026-01-02T03:04:05.006Z")
	model.Local = dynamicRouteServerMemberReferenceModel
	model.ProtocolState = core.StringPtr("active")
	model.Remote = dynamicRouteServerPeerGroupPeerTransitGatewaySessionRemoteModel

	result, err := vpc.DataSourceIBMIsDynamicRouteServerPeerGroupPeersDynamicRouteServerPeerGroupPeerTransitGatewaySessionToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMIsDynamicRouteServerPeerGroupPeersDynamicRouteServerPeerGroupPeerTransitGatewaySessionRemoteToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["name"] = "my-transit-gateway-dynamic-route-server-connection-1"

		assert.Equal(t, result, model)
	}

	model := new(vpcv1.DynamicRouteServerPeerGroupPeerTransitGatewaySessionRemote)
	model.Name = core.StringPtr("my-transit-gateway-dynamic-route-server-connection-1")

	result, err := vpc.DataSourceIBMIsDynamicRouteServerPeerGroupPeersDynamicRouteServerPeerGroupPeerTransitGatewaySessionRemoteToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}
