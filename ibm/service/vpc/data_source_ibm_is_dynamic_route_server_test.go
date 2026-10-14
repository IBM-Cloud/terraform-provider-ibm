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

func TestAccIBMIsDynamicRouteServerDataSourceBasic(t *testing.T) {
	vpcName := fmt.Sprintf("tf-drs-vpc-%d", acctest.RandIntRange(10, 100))
	subnetName := fmt.Sprintf("tf-drs-subnet-%d", acctest.RandIntRange(10, 100))
	vniName := fmt.Sprintf("tf-drs-vni-%d", acctest.RandIntRange(10, 100))
	memberName := fmt.Sprintf("tf-drs-member-%d", acctest.RandIntRange(10, 100))
	localAsn := fmt.Sprintf("%d", acctest.RandIntRange(64512, 65000))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsDynamicRouteServerDataSourceConfigBasic(vpcName, subnetName, vniName, memberName, localAsn),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("data.ibm_is_dynamic_route_server.is_dynamic_route_server_instance", "id", "ibm_is_dynamic_route_server.is_dynamic_route_server_instance", "id"),
					resource.TestCheckResourceAttrPair("data.ibm_is_dynamic_route_server.is_dynamic_route_server_instance", "identifier", "ibm_is_dynamic_route_server.is_dynamic_route_server_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server.is_dynamic_route_server_instance", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server.is_dynamic_route_server_instance", "crn"),
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server.is_dynamic_route_server_instance", "health_reasons.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server.is_dynamic_route_server_instance", "health_state"),
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server.is_dynamic_route_server_instance", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server.is_dynamic_route_server_instance", "lifecycle_reasons.#"),
					resource.TestCheckResourceAttr("data.ibm_is_dynamic_route_server.is_dynamic_route_server_instance", "lifecycle_state", "stable"),
					resource.TestCheckResourceAttr("data.ibm_is_dynamic_route_server.is_dynamic_route_server_instance", "local_asn", localAsn),
					resource.TestCheckResourceAttr("data.ibm_is_dynamic_route_server.is_dynamic_route_server_instance", "members.#", "1"),
					resource.TestCheckResourceAttr("data.ibm_is_dynamic_route_server.is_dynamic_route_server_instance", "members.0.name", memberName),
					resource.TestCheckResourceAttrPair("data.ibm_is_dynamic_route_server.is_dynamic_route_server_instance", "members.0.virtual_network_interfaces.0.id", "ibm_is_virtual_network_interface.testacc_vni", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server.is_dynamic_route_server_instance", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server.is_dynamic_route_server_instance", "resource_group.0.id"),
					resource.TestCheckResourceAttr("data.ibm_is_dynamic_route_server.is_dynamic_route_server_instance", "resource_type", "dynamic_route_server"),
					resource.TestCheckResourceAttrPair("data.ibm_is_dynamic_route_server.is_dynamic_route_server_instance", "vpc.0.id", "ibm_is_vpc.testacc_vpc", "id"),
				),
			},
		},
	})
}

func TestAccIBMIsDynamicRouteServerDataSourceByName(t *testing.T) {
	vpcName := fmt.Sprintf("tf-drs-vpc-%d", acctest.RandIntRange(10, 100))
	subnetName := fmt.Sprintf("tf-drs-subnet-%d", acctest.RandIntRange(10, 100))
	vniName := fmt.Sprintf("tf-drs-vni-%d", acctest.RandIntRange(10, 100))
	inlineVniName := fmt.Sprintf("tf-drs-inline-vni-%d", acctest.RandIntRange(10, 100))
	memberName := fmt.Sprintf("tf-drs-member-%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf-drs-%d", acctest.RandIntRange(10, 100))
	localAsn := fmt.Sprintf("%d", acctest.RandIntRange(64512, 65000))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsDynamicRouteServerDataSourceConfig(vpcName, subnetName, vniName, inlineVniName, memberName, localAsn, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("data.ibm_is_dynamic_route_server.is_dynamic_route_server_instance", "id", "ibm_is_dynamic_route_server.is_dynamic_route_server_instance", "id"),
					resource.TestCheckResourceAttr("data.ibm_is_dynamic_route_server.is_dynamic_route_server_instance", "name", name),
					resource.TestCheckResourceAttr("data.ibm_is_dynamic_route_server.is_dynamic_route_server_instance", "local_asn", localAsn),
					resource.TestCheckResourceAttr("data.ibm_is_dynamic_route_server.is_dynamic_route_server_instance", "members.0.virtual_network_interfaces.#", "2"),
					resource.TestCheckResourceAttr("data.ibm_is_dynamic_route_server.is_dynamic_route_server_instance", "members.0.virtual_network_interfaces.1.name", inlineVniName),
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server.is_dynamic_route_server_instance", "members.0.virtual_network_interfaces.1.primary_ip.0.address"),
					resource.TestCheckResourceAttrPair("data.ibm_is_dynamic_route_server.is_dynamic_route_server_instance", "members.0.virtual_network_interfaces.1.subnet.0.id", "ibm_is_subnet.testacc_subnet", "id"),
					resource.TestCheckResourceAttr("data.ibm_is_dynamic_route_server.is_dynamic_route_server_instance", "resource_group.0.id", acc.IsResourceGroupID),
				),
			},
		},
	})
}

func testAccCheckIBMIsDynamicRouteServerDataSourceConfigBasic(vpcName string, subnetName string, vniName string, memberName string, localAsn string) string {
	return testAccCheckIBMIsDynamicRouteServerConfigBasic(vpcName, subnetName, vniName, memberName, localAsn) + `
		data "ibm_is_dynamic_route_server" "is_dynamic_route_server_instance" {
			identifier = ibm_is_dynamic_route_server.is_dynamic_route_server_instance.id
		}
	`
}

func testAccCheckIBMIsDynamicRouteServerDataSourceConfig(vpcName string, subnetName string, vniName string, inlineVniName string, memberName string, localAsn string, name string) string {
	return testAccCheckIBMIsDynamicRouteServerConfig(vpcName, subnetName, vniName, inlineVniName, memberName, localAsn, name) + `
		data "ibm_is_dynamic_route_server" "is_dynamic_route_server_instance" {
			name = ibm_is_dynamic_route_server.is_dynamic_route_server_instance.name
		}
	`
}

func TestDataSourceIBMIsDynamicRouteServerDynamicRouteServerHealthReasonToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["code"] = "cannot_start_capacity"
		model["message"] = "Cannot start one or more members because resource capacity is unavailable."
		model["more_info"] = "https://cloud.ibm.com/docs/__TBD__"

		assert.Equal(t, result, model)
	}

	model := new(vpcv1.DynamicRouteServerHealthReason)
	model.Code = core.StringPtr("cannot_start_capacity")
	model.Message = core.StringPtr("Cannot start one or more members because resource capacity is unavailable.")
	model.MoreInfo = core.StringPtr("https://cloud.ibm.com/docs/__TBD__")

	result, err := vpc.DataSourceIBMIsDynamicRouteServerDynamicRouteServerHealthReasonToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMIsDynamicRouteServerDynamicRouteServerLifecycleReasonToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["code"] = "resource_suspended_by_provider"
		model["message"] = "The resource has been suspended. Contact IBM support with the CRN for next steps."
		model["more_info"] = "https://cloud.ibm.com/apidocs/vpc#resource-suspension"

		assert.Equal(t, result, model)
	}

	model := new(vpcv1.DynamicRouteServerLifecycleReason)
	model.Code = core.StringPtr("resource_suspended_by_provider")
	model.Message = core.StringPtr("The resource has been suspended. Contact IBM support with the CRN for next steps.")
	model.MoreInfo = core.StringPtr("https://cloud.ibm.com/apidocs/vpc#resource-suspension")

	result, err := vpc.DataSourceIBMIsDynamicRouteServerDynamicRouteServerLifecycleReasonToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMIsDynamicRouteServerDynamicRouteServerMemberReferenceToMap(t *testing.T) {
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

	result, err := vpc.DataSourceIBMIsDynamicRouteServerDynamicRouteServerMemberReferenceToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMIsDynamicRouteServerDeletedToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["more_info"] = "https://cloud.ibm.com/apidocs/vpc#deleted-resources"

		assert.Equal(t, result, model)
	}

	model := new(vpcv1.Deleted)
	model.MoreInfo = core.StringPtr("https://cloud.ibm.com/apidocs/vpc#deleted-resources")

	result, err := vpc.DataSourceIBMIsDynamicRouteServerDeletedToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMIsDynamicRouteServerVirtualNetworkInterfaceReferenceToMap(t *testing.T) {
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

	result, err := vpc.DataSourceIBMIsDynamicRouteServerVirtualNetworkInterfaceReferenceToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMIsDynamicRouteServerReservedIPReferenceToMap(t *testing.T) {
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

	result, err := vpc.DataSourceIBMIsDynamicRouteServerReservedIPReferenceToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMIsDynamicRouteServerSubnetReferenceToMap(t *testing.T) {
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

	result, err := vpc.DataSourceIBMIsDynamicRouteServerSubnetReferenceToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMIsDynamicRouteServerResourceGroupReferenceToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["href"] = "https://resource-controller.cloud.ibm.com/v2/resource_groups/fee82deba12e4c0fb69c3b09d1f12345"
		model["id"] = "fee82deba12e4c0fb69c3b09d1f12345"
		model["name"] = "my-resource-group"

		assert.Equal(t, result, model)
	}

	model := new(vpcv1.ResourceGroupReference)
	model.Href = core.StringPtr("https://resource-controller.cloud.ibm.com/v2/resource_groups/fee82deba12e4c0fb69c3b09d1f12345")
	model.ID = core.StringPtr("fee82deba12e4c0fb69c3b09d1f12345")
	model.Name = core.StringPtr("my-resource-group")

	result, err := vpc.DataSourceIBMIsDynamicRouteServerResourceGroupReferenceToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMIsDynamicRouteServerVPCReferenceToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		deletedModel := make(map[string]interface{})
		deletedModel["more_info"] = "https://cloud.ibm.com/apidocs/vpc#deleted-resources"

		model := make(map[string]interface{})
		model["crn"] = "crn:v1:bluemix:public:is:us-south:a/aa2432b1fa4d4ace891e9b80fc104e34::vpc:r006-4727d842-f94f-4a2d-824a-9bc9b02c523b"
		model["deleted"] = []map[string]interface{}{deletedModel}
		model["href"] = "https://us-south.iaas.cloud.ibm.com/v1/vpcs/r006-4727d842-f94f-4a2d-824a-9bc9b02c523b"
		model["id"] = "r006-4727d842-f94f-4a2d-824a-9bc9b02c523b"
		model["name"] = "my-vpc"
		model["resource_type"] = "vpc"

		assert.Equal(t, result, model)
	}

	deletedModel := new(vpcv1.Deleted)
	deletedModel.MoreInfo = core.StringPtr("https://cloud.ibm.com/apidocs/vpc#deleted-resources")

	model := new(vpcv1.VPCReference)
	model.CRN = core.StringPtr("crn:v1:bluemix:public:is:us-south:a/aa2432b1fa4d4ace891e9b80fc104e34::vpc:r006-4727d842-f94f-4a2d-824a-9bc9b02c523b")
	model.Deleted = deletedModel
	model.Href = core.StringPtr("https://us-south.iaas.cloud.ibm.com/v1/vpcs/r006-4727d842-f94f-4a2d-824a-9bc9b02c523b")
	model.ID = core.StringPtr("r006-4727d842-f94f-4a2d-824a-9bc9b02c523b")
	model.Name = core.StringPtr("my-vpc")
	model.ResourceType = core.StringPtr("vpc")

	result, err := vpc.DataSourceIBMIsDynamicRouteServerVPCReferenceToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMIsDynamicRouteServerToMapNilAndEmpty(t *testing.T) {
	empty := map[string]interface{}{}

	result, err := vpc.DataSourceIBMIsDynamicRouteServerDynamicRouteServerHealthReasonToMap(nil)
	assert.Nil(t, err)
	assert.Equal(t, empty, result)

	result, err = vpc.DataSourceIBMIsDynamicRouteServerDynamicRouteServerLifecycleReasonToMap(nil)
	assert.Nil(t, err)
	assert.Equal(t, empty, result)

	result, err = vpc.DataSourceIBMIsDynamicRouteServerDynamicRouteServerMemberReferenceToMap(nil)
	assert.Nil(t, err)
	assert.Equal(t, empty, result)

	result, err = vpc.DataSourceIBMIsDynamicRouteServerDynamicRouteServerMemberReferenceToMap(new(vpcv1.DynamicRouteServerMemberReference))
	assert.Nil(t, err)
	assert.Equal(t, map[string]interface{}{"virtual_network_interfaces": []map[string]interface{}{}}, result)

	result, err = vpc.DataSourceIBMIsDynamicRouteServerDeletedToMap(nil)
	assert.Nil(t, err)
	assert.Equal(t, empty, result)

	result, err = vpc.DataSourceIBMIsDynamicRouteServerVirtualNetworkInterfaceReferenceToMap(nil)
	assert.Nil(t, err)
	assert.Equal(t, empty, result)

	// no primary IP and no subnet: those keys are absent instead of panicking
	result, err = vpc.DataSourceIBMIsDynamicRouteServerVirtualNetworkInterfaceReferenceToMap(&vpcv1.VirtualNetworkInterfaceReference{ID: core.StringPtr("vni-id")})
	assert.Nil(t, err)
	assert.Equal(t, map[string]interface{}{"id": "vni-id"}, result)

	result, err = vpc.DataSourceIBMIsDynamicRouteServerReservedIPReferenceToMap(nil)
	assert.Nil(t, err)
	assert.Equal(t, empty, result)

	result, err = vpc.DataSourceIBMIsDynamicRouteServerSubnetReferenceToMap(nil)
	assert.Nil(t, err)
	assert.Equal(t, empty, result)

	result, err = vpc.DataSourceIBMIsDynamicRouteServerResourceGroupReferenceToMap(nil)
	assert.Nil(t, err)
	assert.Equal(t, empty, result)

	result, err = vpc.DataSourceIBMIsDynamicRouteServerVPCReferenceToMap(nil)
	assert.Nil(t, err)
	assert.Equal(t, empty, result)

	result, err = vpc.DataSourceIBMIsDynamicRouteServerVPCReferenceToMap(new(vpcv1.VPCReference))
	assert.Nil(t, err)
	assert.Equal(t, empty, result)
}

func TestDataSourceIBMIsDynamicRouteServerDynamicRouteServerMemberReferenceToMapMultiple(t *testing.T) {
	model := new(vpcv1.DynamicRouteServerMemberReference)
	model.ID = core.StringPtr("member-id")
	model.VirtualNetworkInterfaces = []vpcv1.VirtualNetworkInterfaceReference{
		{ID: core.StringPtr("vni-1")},
		{ID: core.StringPtr("vni-2")},
	}
	result, err := vpc.DataSourceIBMIsDynamicRouteServerDynamicRouteServerMemberReferenceToMap(model)
	assert.Nil(t, err)
	assert.Equal(t, []map[string]interface{}{{"id": "vni-1"}, {"id": "vni-2"}}, result["virtual_network_interfaces"])
}
