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

func TestAccIBMIsDynamicRouteServerMemberBasic(t *testing.T) {
	var conf vpcv1.DynamicRouteServerMember
	dynamicRouteServerID := fmt.Sprintf("tf_dynamic_route_server_id_%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIsDynamicRouteServerMemberDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsDynamicRouteServerMemberConfigBasic(dynamicRouteServerID, name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIsDynamicRouteServerMemberExists("ibm_is_dynamic_route_server_member.is_dynamic_route_server_member_instance", conf),
					resource.TestCheckResourceAttr("ibm_is_dynamic_route_server_member.is_dynamic_route_server_member_instance", "dynamic_route_server_id", dynamicRouteServerID),
					resource.TestCheckResourceAttr("ibm_is_dynamic_route_server_member.is_dynamic_route_server_member_instance", "name", name),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_is_dynamic_route_server_member.is_dynamic_route_server_member_instance",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIBMIsDynamicRouteServerMemberConfigBasic(dynamicRouteServerID string, name string) string {
	return fmt.Sprintf(`
		resource "ibm_is_dynamic_route_server_member" "is_dynamic_route_server_member_instance" {
			dynamic_route_server_id = "%s"
			name = "%s"
			virtual_network_interfaces {
				crn = "crn:v1:bluemix:public:is:us-south-1:a/aa2432b1fa4d4ace891e9b80fc104e34::virtual-network-interface:0717-54eb57ee-86f2-4796-90bb-d7874e0831ef"
				deleted {
					more_info = "https://cloud.ibm.com/apidocs/vpc#deleted-resources"
				}
				href = "https://us-south.iaas.cloud.ibm.com/v1/virtual_network_interfaces/0717-54eb57ee-86f2-4796-90bb-d7874e0831ef"
				id = "0717-54eb57ee-86f2-4796-90bb-d7874e0831ef"
				name = "my-virtual-network-interface"
				primary_ip {
					address = "192.168.3.4"
					deleted {
						more_info = "https://cloud.ibm.com/apidocs/vpc#deleted-resources"
					}
					href = "https://us-south.iaas.cloud.ibm.com/v1/subnets/0717-bea6a632-5e13-42a4-b4b8-31dc877abfe4/reserved_ips/0717-6d353a0f-aeb1-4ae1-832e-1110d10981bb"
					id = "0717-6d353a0f-aeb1-4ae1-832e-1110d10981bb"
					name = "my-reserved-ip"
					resource_type = "subnet_reserved_ip"
				}
				resource_type = "virtual_network_interface"
				subnet {
					crn = "crn:v1:bluemix:public:is:us-south-1:a/aa2432b1fa4d4ace891e9b80fc104e34::subnet:0717-7ec86020-1c6e-4889-b3f0-a15f2e50f87e"
					deleted {
						more_info = "https://cloud.ibm.com/apidocs/vpc#deleted-resources"
					}
					href = "https://us-south.iaas.cloud.ibm.com/v1/subnets/0717-7ec86020-1c6e-4889-b3f0-a15f2e50f87e"
					id = "0717-7ec86020-1c6e-4889-b3f0-a15f2e50f87e"
					name = "my-subnet"
					resource_type = "subnet"
				}
			}
		}
	`, dynamicRouteServerID, name)
}

func testAccCheckIBMIsDynamicRouteServerMemberExists(n string, obj vpcv1.DynamicRouteServerMember) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		vpcClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).VpcV1API()
		if err != nil {
			return err
		}

		getDynamicRouteServerMemberOptions := &vpcv1.GetDynamicRouteServerMemberOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getDynamicRouteServerMemberOptions.SetDynamicRouteServerID(parts[0])
		getDynamicRouteServerMemberOptions.SetID(parts[1])

		dynamicRouteServerMember, _, err := vpcClient.GetDynamicRouteServerMember(getDynamicRouteServerMemberOptions)
		if err != nil {
			return err
		}

		obj = *dynamicRouteServerMember
		return nil
	}
}

func testAccCheckIBMIsDynamicRouteServerMemberDestroy(s *terraform.State) error {
	vpcClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).VpcV1API()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_is_dynamic_route_server_member" {
			continue
		}

		getDynamicRouteServerMemberOptions := &vpcv1.GetDynamicRouteServerMemberOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getDynamicRouteServerMemberOptions.SetDynamicRouteServerID(parts[0])
		getDynamicRouteServerMemberOptions.SetID(parts[1])

		// Try to find the key
		_, response, err := vpcClient.GetDynamicRouteServerMember(getDynamicRouteServerMemberOptions)

		if err == nil {
			return fmt.Errorf("DynamicRouteServerMember still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for DynamicRouteServerMember (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}

func TestResourceIBMIsDynamicRouteServerMemberVirtualNetworkInterfaceReferenceToMap(t *testing.T) {
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

	result, err := vpc.ResourceIBMIsDynamicRouteServerMemberVirtualNetworkInterfaceReferenceToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMIsDynamicRouteServerMemberDeletedToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["more_info"] = "https://cloud.ibm.com/apidocs/vpc#deleted-resources"

		assert.Equal(t, result, model)
	}

	model := new(vpcv1.Deleted)
	model.MoreInfo = core.StringPtr("https://cloud.ibm.com/apidocs/vpc#deleted-resources")

	result, err := vpc.ResourceIBMIsDynamicRouteServerMemberDeletedToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMIsDynamicRouteServerMemberReservedIPReferenceToMap(t *testing.T) {
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

	result, err := vpc.ResourceIBMIsDynamicRouteServerMemberReservedIPReferenceToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMIsDynamicRouteServerMemberSubnetReferenceToMap(t *testing.T) {
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

	result, err := vpc.ResourceIBMIsDynamicRouteServerMemberSubnetReferenceToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMIsDynamicRouteServerMemberDynamicRouteServerHealthReasonToMap(t *testing.T) {
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

	result, err := vpc.ResourceIBMIsDynamicRouteServerMemberDynamicRouteServerHealthReasonToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMIsDynamicRouteServerMemberDynamicRouteServerLifecycleReasonToMap(t *testing.T) {
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

	result, err := vpc.ResourceIBMIsDynamicRouteServerMemberDynamicRouteServerLifecycleReasonToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMIsDynamicRouteServerMemberMapToDynamicRouteServerMemberPrototypeVirtualNetworkInterface(t *testing.T) {
	checkResult := func(result vpcv1.DynamicRouteServerMemberPrototypeVirtualNetworkInterfaceIntf) {
		virtualNetworkInterfaceIPPrototypeModel := new(vpcv1.VirtualNetworkInterfaceIPPrototypeReservedIPPrototypeVirtualNetworkInterfaceIPsContext)
		virtualNetworkInterfaceIPPrototypeModel.Address = core.StringPtr("10.0.0.5")
		virtualNetworkInterfaceIPPrototypeModel.AutoDelete = core.BoolPtr(false)
		virtualNetworkInterfaceIPPrototypeModel.Name = core.StringPtr("my-reserved-ip")

		virtualNetworkInterfacePrimaryIPPrototypeModel := new(vpcv1.VirtualNetworkInterfacePrimaryIPPrototypeReservedIPIdentityVirtualNetworkInterfacePrimaryIPContextByID)
		virtualNetworkInterfacePrimaryIPPrototypeModel.ID = core.StringPtr("0717-6d353a0f-aeb1-4ae1-832e-1110d10981bb")

		resourceGroupIdentityModel := new(vpcv1.ResourceGroupIdentityByID)
		resourceGroupIdentityModel.ID = core.StringPtr("fee82deba12e4c0fb69c3b09d1f12345")

		securityGroupIdentityModel := new(vpcv1.SecurityGroupIdentityByID)
		securityGroupIdentityModel.ID = core.StringPtr("r006-be5df5ca-12a0-494b-907e-aa6ec2bfa271")

		subnetIdentityModel := new(vpcv1.SubnetIdentityByID)
		subnetIdentityModel.ID = core.StringPtr("0717-7ec86020-1c6e-4889-b3f0-a15f2e50f87e")

		model := new(vpcv1.DynamicRouteServerMemberPrototypeVirtualNetworkInterface)
		model.AllowIPSpoofing = core.BoolPtr(true)
		model.AutoDelete = core.BoolPtr(false)
		model.EnableInfrastructureNat = core.BoolPtr(true)
		model.Ips = []vpcv1.VirtualNetworkInterfaceIPPrototypeIntf{virtualNetworkInterfaceIPPrototypeModel}
		model.Name = core.StringPtr("my-virtual-network-interface")
		model.PrimaryIP = virtualNetworkInterfacePrimaryIPPrototypeModel
		model.ProtocolStateFilteringMode = core.StringPtr("auto")
		model.ResourceGroup = resourceGroupIdentityModel
		model.SecurityGroups = []vpcv1.SecurityGroupIdentityIntf{securityGroupIdentityModel}
		model.Subnet = subnetIdentityModel
		model.ID = core.StringPtr("0717-54eb57ee-86f2-4796-90bb-d7874e0831ef")
		model.Href = core.StringPtr("https://us-south.iaas.cloud.ibm.com/v1/virtual_network_interfaces/0717-54eb57ee-86f2-4796-90bb-d7874e0831ef")
		model.CRN = core.StringPtr("crn:v1:bluemix:public:is:us-south-1:a/aa2432b1fa4d4ace891e9b80fc104e34::virtual-network-interface:0717-54eb57ee-86f2-4796-90bb-d7874e0831ef")

		assert.Equal(t, result, model)
	}

	virtualNetworkInterfaceIPPrototypeModel := make(map[string]interface{})
	virtualNetworkInterfaceIPPrototypeModel["address"] = "10.0.0.5"
	virtualNetworkInterfaceIPPrototypeModel["auto_delete"] = false
	virtualNetworkInterfaceIPPrototypeModel["name"] = "my-reserved-ip"

	virtualNetworkInterfacePrimaryIPPrototypeModel := make(map[string]interface{})
	virtualNetworkInterfacePrimaryIPPrototypeModel["id"] = "0717-6d353a0f-aeb1-4ae1-832e-1110d10981bb"

	resourceGroupIdentityModel := make(map[string]interface{})
	resourceGroupIdentityModel["id"] = "fee82deba12e4c0fb69c3b09d1f12345"

	securityGroupIdentityModel := make(map[string]interface{})
	securityGroupIdentityModel["id"] = "r006-be5df5ca-12a0-494b-907e-aa6ec2bfa271"

	subnetIdentityModel := make(map[string]interface{})
	subnetIdentityModel["id"] = "0717-7ec86020-1c6e-4889-b3f0-a15f2e50f87e"

	model := make(map[string]interface{})
	model["allow_ip_spoofing"] = true
	model["auto_delete"] = false
	model["enable_infrastructure_nat"] = true
	model["ips"] = []interface{}{virtualNetworkInterfaceIPPrototypeModel}
	model["name"] = "my-virtual-network-interface"
	model["primary_ip"] = []interface{}{virtualNetworkInterfacePrimaryIPPrototypeModel}
	model["protocol_state_filtering_mode"] = "auto"
	model["resource_group"] = []interface{}{resourceGroupIdentityModel}
	model["security_groups"] = []interface{}{securityGroupIdentityModel}
	model["subnet"] = []interface{}{subnetIdentityModel}
	model["id"] = "0717-54eb57ee-86f2-4796-90bb-d7874e0831ef"
	model["href"] = "https://us-south.iaas.cloud.ibm.com/v1/virtual_network_interfaces/0717-54eb57ee-86f2-4796-90bb-d7874e0831ef"
	model["crn"] = "crn:v1:bluemix:public:is:us-south-1:a/aa2432b1fa4d4ace891e9b80fc104e34::virtual-network-interface:0717-54eb57ee-86f2-4796-90bb-d7874e0831ef"

	result, err := vpc.ResourceIBMIsDynamicRouteServerMemberMapToDynamicRouteServerMemberPrototypeVirtualNetworkInterface(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMIsDynamicRouteServerMemberMapToVirtualNetworkInterfaceIPPrototype(t *testing.T) {
	checkResult := func(result vpcv1.VirtualNetworkInterfaceIPPrototypeIntf) {
		model := new(vpcv1.VirtualNetworkInterfaceIPPrototype)
		model.ID = core.StringPtr("0717-6d353a0f-aeb1-4ae1-832e-1110d10981bb")
		model.Href = core.StringPtr("https://us-south.iaas.cloud.ibm.com/v1/subnets/0717-bea6a632-5e13-42a4-b4b8-31dc877abfe4/reserved_ips/0717-6d353a0f-aeb1-4ae1-832e-1110d10981bb")
		model.Address = core.StringPtr("192.168.3.4")
		model.AutoDelete = core.BoolPtr(false)
		model.Name = core.StringPtr("my-reserved-ip")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["id"] = "0717-6d353a0f-aeb1-4ae1-832e-1110d10981bb"
	model["href"] = "https://us-south.iaas.cloud.ibm.com/v1/subnets/0717-bea6a632-5e13-42a4-b4b8-31dc877abfe4/reserved_ips/0717-6d353a0f-aeb1-4ae1-832e-1110d10981bb"
	model["address"] = "192.168.3.4"
	model["auto_delete"] = false
	model["name"] = "my-reserved-ip"

	result, err := vpc.ResourceIBMIsDynamicRouteServerMemberMapToVirtualNetworkInterfaceIPPrototype(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMIsDynamicRouteServerMemberMapToVirtualNetworkInterfaceIPPrototypeReservedIPIdentityVirtualNetworkInterfaceIPsContext(t *testing.T) {
	checkResult := func(result vpcv1.VirtualNetworkInterfaceIPPrototypeReservedIPIdentityVirtualNetworkInterfaceIPsContextIntf) {
		model := new(vpcv1.VirtualNetworkInterfaceIPPrototypeReservedIPIdentityVirtualNetworkInterfaceIPsContext)
		model.ID = core.StringPtr("0717-6d353a0f-aeb1-4ae1-832e-1110d10981bb")
		model.Href = core.StringPtr("https://us-south.iaas.cloud.ibm.com/v1/subnets/0717-bea6a632-5e13-42a4-b4b8-31dc877abfe4/reserved_ips/0717-6d353a0f-aeb1-4ae1-832e-1110d10981bb")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["id"] = "0717-6d353a0f-aeb1-4ae1-832e-1110d10981bb"
	model["href"] = "https://us-south.iaas.cloud.ibm.com/v1/subnets/0717-bea6a632-5e13-42a4-b4b8-31dc877abfe4/reserved_ips/0717-6d353a0f-aeb1-4ae1-832e-1110d10981bb"

	result, err := vpc.ResourceIBMIsDynamicRouteServerMemberMapToVirtualNetworkInterfaceIPPrototypeReservedIPIdentityVirtualNetworkInterfaceIPsContext(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMIsDynamicRouteServerMemberMapToVirtualNetworkInterfaceIPPrototypeReservedIPIdentityVirtualNetworkInterfaceIPsContextByID(t *testing.T) {
	checkResult := func(result *vpcv1.VirtualNetworkInterfaceIPPrototypeReservedIPIdentityVirtualNetworkInterfaceIPsContextByID) {
		model := new(vpcv1.VirtualNetworkInterfaceIPPrototypeReservedIPIdentityVirtualNetworkInterfaceIPsContextByID)
		model.ID = core.StringPtr("0717-6d353a0f-aeb1-4ae1-832e-1110d10981bb")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["id"] = "0717-6d353a0f-aeb1-4ae1-832e-1110d10981bb"

	result, err := vpc.ResourceIBMIsDynamicRouteServerMemberMapToVirtualNetworkInterfaceIPPrototypeReservedIPIdentityVirtualNetworkInterfaceIPsContextByID(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMIsDynamicRouteServerMemberMapToVirtualNetworkInterfaceIPPrototypeReservedIPIdentityVirtualNetworkInterfaceIPsContextByHref(t *testing.T) {
	checkResult := func(result *vpcv1.VirtualNetworkInterfaceIPPrototypeReservedIPIdentityVirtualNetworkInterfaceIPsContextByHref) {
		model := new(vpcv1.VirtualNetworkInterfaceIPPrototypeReservedIPIdentityVirtualNetworkInterfaceIPsContextByHref)
		model.Href = core.StringPtr("https://us-south.iaas.cloud.ibm.com/v1/subnets/0717-bea6a632-5e13-42a4-b4b8-31dc877abfe4/reserved_ips/0717-6d353a0f-aeb1-4ae1-832e-1110d10981bb")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["href"] = "https://us-south.iaas.cloud.ibm.com/v1/subnets/0717-bea6a632-5e13-42a4-b4b8-31dc877abfe4/reserved_ips/0717-6d353a0f-aeb1-4ae1-832e-1110d10981bb"

	result, err := vpc.ResourceIBMIsDynamicRouteServerMemberMapToVirtualNetworkInterfaceIPPrototypeReservedIPIdentityVirtualNetworkInterfaceIPsContextByHref(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMIsDynamicRouteServerMemberMapToVirtualNetworkInterfaceIPPrototypeReservedIPPrototypeVirtualNetworkInterfaceIPsContext(t *testing.T) {
	checkResult := func(result *vpcv1.VirtualNetworkInterfaceIPPrototypeReservedIPPrototypeVirtualNetworkInterfaceIPsContext) {
		model := new(vpcv1.VirtualNetworkInterfaceIPPrototypeReservedIPPrototypeVirtualNetworkInterfaceIPsContext)
		model.Address = core.StringPtr("192.168.3.4")
		model.AutoDelete = core.BoolPtr(false)
		model.Name = core.StringPtr("my-reserved-ip")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["address"] = "192.168.3.4"
	model["auto_delete"] = false
	model["name"] = "my-reserved-ip"

	result, err := vpc.ResourceIBMIsDynamicRouteServerMemberMapToVirtualNetworkInterfaceIPPrototypeReservedIPPrototypeVirtualNetworkInterfaceIPsContext(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMIsDynamicRouteServerMemberMapToVirtualNetworkInterfacePrimaryIPPrototype(t *testing.T) {
	checkResult := func(result vpcv1.VirtualNetworkInterfacePrimaryIPPrototypeIntf) {
		model := new(vpcv1.VirtualNetworkInterfacePrimaryIPPrototype)
		model.ID = core.StringPtr("0717-6d353a0f-aeb1-4ae1-832e-1110d10981bb")
		model.Href = core.StringPtr("https://us-south.iaas.cloud.ibm.com/v1/subnets/0717-bea6a632-5e13-42a4-b4b8-31dc877abfe4/reserved_ips/0717-6d353a0f-aeb1-4ae1-832e-1110d10981bb")
		model.Address = core.StringPtr("192.168.3.4")
		model.AutoDelete = core.BoolPtr(false)
		model.Name = core.StringPtr("my-reserved-ip")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["id"] = "0717-6d353a0f-aeb1-4ae1-832e-1110d10981bb"
	model["href"] = "https://us-south.iaas.cloud.ibm.com/v1/subnets/0717-bea6a632-5e13-42a4-b4b8-31dc877abfe4/reserved_ips/0717-6d353a0f-aeb1-4ae1-832e-1110d10981bb"
	model["address"] = "192.168.3.4"
	model["auto_delete"] = false
	model["name"] = "my-reserved-ip"

	result, err := vpc.ResourceIBMIsDynamicRouteServerMemberMapToVirtualNetworkInterfacePrimaryIPPrototype(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMIsDynamicRouteServerMemberMapToVirtualNetworkInterfacePrimaryIPPrototypeReservedIPIdentityVirtualNetworkInterfacePrimaryIPContext(t *testing.T) {
	checkResult := func(result vpcv1.VirtualNetworkInterfacePrimaryIPPrototypeReservedIPIdentityVirtualNetworkInterfacePrimaryIPContextIntf) {
		model := new(vpcv1.VirtualNetworkInterfacePrimaryIPPrototypeReservedIPIdentityVirtualNetworkInterfacePrimaryIPContext)
		model.ID = core.StringPtr("0717-6d353a0f-aeb1-4ae1-832e-1110d10981bb")
		model.Href = core.StringPtr("https://us-south.iaas.cloud.ibm.com/v1/subnets/0717-bea6a632-5e13-42a4-b4b8-31dc877abfe4/reserved_ips/0717-6d353a0f-aeb1-4ae1-832e-1110d10981bb")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["id"] = "0717-6d353a0f-aeb1-4ae1-832e-1110d10981bb"
	model["href"] = "https://us-south.iaas.cloud.ibm.com/v1/subnets/0717-bea6a632-5e13-42a4-b4b8-31dc877abfe4/reserved_ips/0717-6d353a0f-aeb1-4ae1-832e-1110d10981bb"

	result, err := vpc.ResourceIBMIsDynamicRouteServerMemberMapToVirtualNetworkInterfacePrimaryIPPrototypeReservedIPIdentityVirtualNetworkInterfacePrimaryIPContext(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMIsDynamicRouteServerMemberMapToVirtualNetworkInterfacePrimaryIPPrototypeReservedIPIdentityVirtualNetworkInterfacePrimaryIPContextByID(t *testing.T) {
	checkResult := func(result *vpcv1.VirtualNetworkInterfacePrimaryIPPrototypeReservedIPIdentityVirtualNetworkInterfacePrimaryIPContextByID) {
		model := new(vpcv1.VirtualNetworkInterfacePrimaryIPPrototypeReservedIPIdentityVirtualNetworkInterfacePrimaryIPContextByID)
		model.ID = core.StringPtr("0717-6d353a0f-aeb1-4ae1-832e-1110d10981bb")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["id"] = "0717-6d353a0f-aeb1-4ae1-832e-1110d10981bb"

	result, err := vpc.ResourceIBMIsDynamicRouteServerMemberMapToVirtualNetworkInterfacePrimaryIPPrototypeReservedIPIdentityVirtualNetworkInterfacePrimaryIPContextByID(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMIsDynamicRouteServerMemberMapToVirtualNetworkInterfacePrimaryIPPrototypeReservedIPIdentityVirtualNetworkInterfacePrimaryIPContextByHref(t *testing.T) {
	checkResult := func(result *vpcv1.VirtualNetworkInterfacePrimaryIPPrototypeReservedIPIdentityVirtualNetworkInterfacePrimaryIPContextByHref) {
		model := new(vpcv1.VirtualNetworkInterfacePrimaryIPPrototypeReservedIPIdentityVirtualNetworkInterfacePrimaryIPContextByHref)
		model.Href = core.StringPtr("https://us-south.iaas.cloud.ibm.com/v1/subnets/0717-bea6a632-5e13-42a4-b4b8-31dc877abfe4/reserved_ips/0717-6d353a0f-aeb1-4ae1-832e-1110d10981bb")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["href"] = "https://us-south.iaas.cloud.ibm.com/v1/subnets/0717-bea6a632-5e13-42a4-b4b8-31dc877abfe4/reserved_ips/0717-6d353a0f-aeb1-4ae1-832e-1110d10981bb"

	result, err := vpc.ResourceIBMIsDynamicRouteServerMemberMapToVirtualNetworkInterfacePrimaryIPPrototypeReservedIPIdentityVirtualNetworkInterfacePrimaryIPContextByHref(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMIsDynamicRouteServerMemberMapToVirtualNetworkInterfacePrimaryIPPrototypeReservedIPPrototypeVirtualNetworkInterfacePrimaryIPContext(t *testing.T) {
	checkResult := func(result *vpcv1.VirtualNetworkInterfacePrimaryIPPrototypeReservedIPPrototypeVirtualNetworkInterfacePrimaryIPContext) {
		model := new(vpcv1.VirtualNetworkInterfacePrimaryIPPrototypeReservedIPPrototypeVirtualNetworkInterfacePrimaryIPContext)
		model.Address = core.StringPtr("192.168.3.4")
		model.AutoDelete = core.BoolPtr(false)
		model.Name = core.StringPtr("my-reserved-ip")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["address"] = "192.168.3.4"
	model["auto_delete"] = false
	model["name"] = "my-reserved-ip"

	result, err := vpc.ResourceIBMIsDynamicRouteServerMemberMapToVirtualNetworkInterfacePrimaryIPPrototypeReservedIPPrototypeVirtualNetworkInterfacePrimaryIPContext(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMIsDynamicRouteServerMemberMapToResourceGroupIdentity(t *testing.T) {
	checkResult := func(result vpcv1.ResourceGroupIdentityIntf) {
		model := new(vpcv1.ResourceGroupIdentity)
		model.ID = core.StringPtr("fee82deba12e4c0fb69c3b09d1f12345")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["id"] = "fee82deba12e4c0fb69c3b09d1f12345"

	result, err := vpc.ResourceIBMIsDynamicRouteServerMemberMapToResourceGroupIdentity(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMIsDynamicRouteServerMemberMapToResourceGroupIdentityByID(t *testing.T) {
	checkResult := func(result *vpcv1.ResourceGroupIdentityByID) {
		model := new(vpcv1.ResourceGroupIdentityByID)
		model.ID = core.StringPtr("fee82deba12e4c0fb69c3b09d1f12345")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["id"] = "fee82deba12e4c0fb69c3b09d1f12345"

	result, err := vpc.ResourceIBMIsDynamicRouteServerMemberMapToResourceGroupIdentityByID(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMIsDynamicRouteServerMemberMapToSecurityGroupIdentity(t *testing.T) {
	checkResult := func(result vpcv1.SecurityGroupIdentityIntf) {
		model := new(vpcv1.SecurityGroupIdentity)
		model.ID = core.StringPtr("r006-be5df5ca-12a0-494b-907e-aa6ec2bfa271")
		model.CRN = core.StringPtr("crn:v1:bluemix:public:is:us-south:a/aa2432b1fa4d4ace891e9b80fc104e34::security-group:r006-be5df5ca-12a0-494b-907e-aa6ec2bfa271")
		model.Href = core.StringPtr("https://us-south.iaas.cloud.ibm.com/v1/security_groups/r006-be5df5ca-12a0-494b-907e-aa6ec2bfa271")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["id"] = "r006-be5df5ca-12a0-494b-907e-aa6ec2bfa271"
	model["crn"] = "crn:v1:bluemix:public:is:us-south:a/aa2432b1fa4d4ace891e9b80fc104e34::security-group:r006-be5df5ca-12a0-494b-907e-aa6ec2bfa271"
	model["href"] = "https://us-south.iaas.cloud.ibm.com/v1/security_groups/r006-be5df5ca-12a0-494b-907e-aa6ec2bfa271"

	result, err := vpc.ResourceIBMIsDynamicRouteServerMemberMapToSecurityGroupIdentity(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMIsDynamicRouteServerMemberMapToSecurityGroupIdentityByID(t *testing.T) {
	checkResult := func(result *vpcv1.SecurityGroupIdentityByID) {
		model := new(vpcv1.SecurityGroupIdentityByID)
		model.ID = core.StringPtr("r006-be5df5ca-12a0-494b-907e-aa6ec2bfa271")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["id"] = "r006-be5df5ca-12a0-494b-907e-aa6ec2bfa271"

	result, err := vpc.ResourceIBMIsDynamicRouteServerMemberMapToSecurityGroupIdentityByID(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMIsDynamicRouteServerMemberMapToSecurityGroupIdentityByCRN(t *testing.T) {
	checkResult := func(result *vpcv1.SecurityGroupIdentityByCRN) {
		model := new(vpcv1.SecurityGroupIdentityByCRN)
		model.CRN = core.StringPtr("crn:v1:bluemix:public:is:us-south:a/aa2432b1fa4d4ace891e9b80fc104e34::security-group:r006-be5df5ca-12a0-494b-907e-aa6ec2bfa271")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["crn"] = "crn:v1:bluemix:public:is:us-south:a/aa2432b1fa4d4ace891e9b80fc104e34::security-group:r006-be5df5ca-12a0-494b-907e-aa6ec2bfa271"

	result, err := vpc.ResourceIBMIsDynamicRouteServerMemberMapToSecurityGroupIdentityByCRN(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMIsDynamicRouteServerMemberMapToSecurityGroupIdentityByHref(t *testing.T) {
	checkResult := func(result *vpcv1.SecurityGroupIdentityByHref) {
		model := new(vpcv1.SecurityGroupIdentityByHref)
		model.Href = core.StringPtr("https://us-south.iaas.cloud.ibm.com/v1/security_groups/r006-be5df5ca-12a0-494b-907e-aa6ec2bfa271")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["href"] = "https://us-south.iaas.cloud.ibm.com/v1/security_groups/r006-be5df5ca-12a0-494b-907e-aa6ec2bfa271"

	result, err := vpc.ResourceIBMIsDynamicRouteServerMemberMapToSecurityGroupIdentityByHref(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMIsDynamicRouteServerMemberMapToSubnetIdentity(t *testing.T) {
	checkResult := func(result vpcv1.SubnetIdentityIntf) {
		model := new(vpcv1.SubnetIdentity)
		model.ID = core.StringPtr("0717-7ec86020-1c6e-4889-b3f0-a15f2e50f87e")
		model.CRN = core.StringPtr("crn:v1:bluemix:public:is:us-south-1:a/aa2432b1fa4d4ace891e9b80fc104e34::subnet:0717-7ec86020-1c6e-4889-b3f0-a15f2e50f87e")
		model.Href = core.StringPtr("https://us-south.iaas.cloud.ibm.com/v1/subnets/0717-7ec86020-1c6e-4889-b3f0-a15f2e50f87e")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["id"] = "0717-7ec86020-1c6e-4889-b3f0-a15f2e50f87e"
	model["crn"] = "crn:v1:bluemix:public:is:us-south-1:a/aa2432b1fa4d4ace891e9b80fc104e34::subnet:0717-7ec86020-1c6e-4889-b3f0-a15f2e50f87e"
	model["href"] = "https://us-south.iaas.cloud.ibm.com/v1/subnets/0717-7ec86020-1c6e-4889-b3f0-a15f2e50f87e"

	result, err := vpc.ResourceIBMIsDynamicRouteServerMemberMapToSubnetIdentity(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMIsDynamicRouteServerMemberMapToSubnetIdentityByID(t *testing.T) {
	checkResult := func(result *vpcv1.SubnetIdentityByID) {
		model := new(vpcv1.SubnetIdentityByID)
		model.ID = core.StringPtr("0717-7ec86020-1c6e-4889-b3f0-a15f2e50f87e")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["id"] = "0717-7ec86020-1c6e-4889-b3f0-a15f2e50f87e"

	result, err := vpc.ResourceIBMIsDynamicRouteServerMemberMapToSubnetIdentityByID(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMIsDynamicRouteServerMemberMapToSubnetIdentityByCRN(t *testing.T) {
	checkResult := func(result *vpcv1.SubnetIdentityByCRN) {
		model := new(vpcv1.SubnetIdentityByCRN)
		model.CRN = core.StringPtr("crn:v1:bluemix:public:is:us-south-1:a/aa2432b1fa4d4ace891e9b80fc104e34::subnet:0717-7ec86020-1c6e-4889-b3f0-a15f2e50f87e")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["crn"] = "crn:v1:bluemix:public:is:us-south-1:a/aa2432b1fa4d4ace891e9b80fc104e34::subnet:0717-7ec86020-1c6e-4889-b3f0-a15f2e50f87e"

	result, err := vpc.ResourceIBMIsDynamicRouteServerMemberMapToSubnetIdentityByCRN(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMIsDynamicRouteServerMemberMapToSubnetIdentityByHref(t *testing.T) {
	checkResult := func(result *vpcv1.SubnetIdentityByHref) {
		model := new(vpcv1.SubnetIdentityByHref)
		model.Href = core.StringPtr("https://us-south.iaas.cloud.ibm.com/v1/subnets/0717-7ec86020-1c6e-4889-b3f0-a15f2e50f87e")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["href"] = "https://us-south.iaas.cloud.ibm.com/v1/subnets/0717-7ec86020-1c6e-4889-b3f0-a15f2e50f87e"

	result, err := vpc.ResourceIBMIsDynamicRouteServerMemberMapToSubnetIdentityByHref(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMIsDynamicRouteServerMemberMapToDynamicRouteServerMemberPrototypeVirtualNetworkInterfaceVirtualNetworkInterfacePrototypeDynamicRouteServerContext(t *testing.T) {
	checkResult := func(result *vpcv1.DynamicRouteServerMemberPrototypeVirtualNetworkInterfaceVirtualNetworkInterfacePrototypeDynamicRouteServerContext) {
		virtualNetworkInterfaceIPPrototypeModel := new(vpcv1.VirtualNetworkInterfaceIPPrototypeReservedIPPrototypeVirtualNetworkInterfaceIPsContext)
		virtualNetworkInterfaceIPPrototypeModel.Address = core.StringPtr("10.0.0.5")
		virtualNetworkInterfaceIPPrototypeModel.AutoDelete = core.BoolPtr(false)
		virtualNetworkInterfaceIPPrototypeModel.Name = core.StringPtr("my-reserved-ip")

		virtualNetworkInterfacePrimaryIPPrototypeModel := new(vpcv1.VirtualNetworkInterfacePrimaryIPPrototypeReservedIPIdentityVirtualNetworkInterfacePrimaryIPContextByID)
		virtualNetworkInterfacePrimaryIPPrototypeModel.ID = core.StringPtr("0717-6d353a0f-aeb1-4ae1-832e-1110d10981bb")

		resourceGroupIdentityModel := new(vpcv1.ResourceGroupIdentityByID)
		resourceGroupIdentityModel.ID = core.StringPtr("fee82deba12e4c0fb69c3b09d1f12345")

		securityGroupIdentityModel := new(vpcv1.SecurityGroupIdentityByID)
		securityGroupIdentityModel.ID = core.StringPtr("r006-be5df5ca-12a0-494b-907e-aa6ec2bfa271")

		subnetIdentityModel := new(vpcv1.SubnetIdentityByID)
		subnetIdentityModel.ID = core.StringPtr("0717-7ec86020-1c6e-4889-b3f0-a15f2e50f87e")

		model := new(vpcv1.DynamicRouteServerMemberPrototypeVirtualNetworkInterfaceVirtualNetworkInterfacePrototypeDynamicRouteServerContext)
		model.AllowIPSpoofing = core.BoolPtr(true)
		model.AutoDelete = core.BoolPtr(false)
		model.EnableInfrastructureNat = core.BoolPtr(true)
		model.Ips = []vpcv1.VirtualNetworkInterfaceIPPrototypeIntf{virtualNetworkInterfaceIPPrototypeModel}
		model.Name = core.StringPtr("my-virtual-network-interface")
		model.PrimaryIP = virtualNetworkInterfacePrimaryIPPrototypeModel
		model.ProtocolStateFilteringMode = core.StringPtr("auto")
		model.ResourceGroup = resourceGroupIdentityModel
		model.SecurityGroups = []vpcv1.SecurityGroupIdentityIntf{securityGroupIdentityModel}
		model.Subnet = subnetIdentityModel

		assert.Equal(t, result, model)
	}

	virtualNetworkInterfaceIPPrototypeModel := make(map[string]interface{})
	virtualNetworkInterfaceIPPrototypeModel["address"] = "10.0.0.5"
	virtualNetworkInterfaceIPPrototypeModel["auto_delete"] = false
	virtualNetworkInterfaceIPPrototypeModel["name"] = "my-reserved-ip"

	virtualNetworkInterfacePrimaryIPPrototypeModel := make(map[string]interface{})
	virtualNetworkInterfacePrimaryIPPrototypeModel["id"] = "0717-6d353a0f-aeb1-4ae1-832e-1110d10981bb"

	resourceGroupIdentityModel := make(map[string]interface{})
	resourceGroupIdentityModel["id"] = "fee82deba12e4c0fb69c3b09d1f12345"

	securityGroupIdentityModel := make(map[string]interface{})
	securityGroupIdentityModel["id"] = "r006-be5df5ca-12a0-494b-907e-aa6ec2bfa271"

	subnetIdentityModel := make(map[string]interface{})
	subnetIdentityModel["id"] = "0717-7ec86020-1c6e-4889-b3f0-a15f2e50f87e"

	model := make(map[string]interface{})
	model["allow_ip_spoofing"] = true
	model["auto_delete"] = false
	model["enable_infrastructure_nat"] = true
	model["ips"] = []interface{}{virtualNetworkInterfaceIPPrototypeModel}
	model["name"] = "my-virtual-network-interface"
	model["primary_ip"] = []interface{}{virtualNetworkInterfacePrimaryIPPrototypeModel}
	model["protocol_state_filtering_mode"] = "auto"
	model["resource_group"] = []interface{}{resourceGroupIdentityModel}
	model["security_groups"] = []interface{}{securityGroupIdentityModel}
	model["subnet"] = []interface{}{subnetIdentityModel}

	result, err := vpc.ResourceIBMIsDynamicRouteServerMemberMapToDynamicRouteServerMemberPrototypeVirtualNetworkInterfaceVirtualNetworkInterfacePrototypeDynamicRouteServerContext(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMIsDynamicRouteServerMemberMapToDynamicRouteServerMemberPrototypeVirtualNetworkInterfaceVirtualNetworkInterfaceIdentity(t *testing.T) {
	checkResult := func(result vpcv1.DynamicRouteServerMemberPrototypeVirtualNetworkInterfaceVirtualNetworkInterfaceIdentityIntf) {
		model := new(vpcv1.DynamicRouteServerMemberPrototypeVirtualNetworkInterfaceVirtualNetworkInterfaceIdentity)
		model.ID = core.StringPtr("0717-54eb57ee-86f2-4796-90bb-d7874e0831ef")
		model.Href = core.StringPtr("https://us-south.iaas.cloud.ibm.com/v1/virtual_network_interfaces/0717-54eb57ee-86f2-4796-90bb-d7874e0831ef")
		model.CRN = core.StringPtr("crn:v1:bluemix:public:is:us-south-1:a/aa2432b1fa4d4ace891e9b80fc104e34::virtual-network-interface:0717-54eb57ee-86f2-4796-90bb-d7874e0831ef")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["id"] = "0717-54eb57ee-86f2-4796-90bb-d7874e0831ef"
	model["href"] = "https://us-south.iaas.cloud.ibm.com/v1/virtual_network_interfaces/0717-54eb57ee-86f2-4796-90bb-d7874e0831ef"
	model["crn"] = "crn:v1:bluemix:public:is:us-south-1:a/aa2432b1fa4d4ace891e9b80fc104e34::virtual-network-interface:0717-54eb57ee-86f2-4796-90bb-d7874e0831ef"

	result, err := vpc.ResourceIBMIsDynamicRouteServerMemberMapToDynamicRouteServerMemberPrototypeVirtualNetworkInterfaceVirtualNetworkInterfaceIdentity(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMIsDynamicRouteServerMemberMapToDynamicRouteServerMemberPrototypeVirtualNetworkInterfaceVirtualNetworkInterfaceIdentityVirtualNetworkInterfaceIdentityByID(t *testing.T) {
	checkResult := func(result *vpcv1.DynamicRouteServerMemberPrototypeVirtualNetworkInterfaceVirtualNetworkInterfaceIdentityVirtualNetworkInterfaceIdentityByID) {
		model := new(vpcv1.DynamicRouteServerMemberPrototypeVirtualNetworkInterfaceVirtualNetworkInterfaceIdentityVirtualNetworkInterfaceIdentityByID)
		model.ID = core.StringPtr("0717-54eb57ee-86f2-4796-90bb-d7874e0831ef")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["id"] = "0717-54eb57ee-86f2-4796-90bb-d7874e0831ef"

	result, err := vpc.ResourceIBMIsDynamicRouteServerMemberMapToDynamicRouteServerMemberPrototypeVirtualNetworkInterfaceVirtualNetworkInterfaceIdentityVirtualNetworkInterfaceIdentityByID(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMIsDynamicRouteServerMemberMapToDynamicRouteServerMemberPrototypeVirtualNetworkInterfaceVirtualNetworkInterfaceIdentityVirtualNetworkInterfaceIdentityByHref(t *testing.T) {
	checkResult := func(result *vpcv1.DynamicRouteServerMemberPrototypeVirtualNetworkInterfaceVirtualNetworkInterfaceIdentityVirtualNetworkInterfaceIdentityByHref) {
		model := new(vpcv1.DynamicRouteServerMemberPrototypeVirtualNetworkInterfaceVirtualNetworkInterfaceIdentityVirtualNetworkInterfaceIdentityByHref)
		model.Href = core.StringPtr("https://us-south.iaas.cloud.ibm.com/v1/virtual_network_interfaces/0717-54eb57ee-86f2-4796-90bb-d7874e0831ef")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["href"] = "https://us-south.iaas.cloud.ibm.com/v1/virtual_network_interfaces/0717-54eb57ee-86f2-4796-90bb-d7874e0831ef"

	result, err := vpc.ResourceIBMIsDynamicRouteServerMemberMapToDynamicRouteServerMemberPrototypeVirtualNetworkInterfaceVirtualNetworkInterfaceIdentityVirtualNetworkInterfaceIdentityByHref(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMIsDynamicRouteServerMemberMapToDynamicRouteServerMemberPrototypeVirtualNetworkInterfaceVirtualNetworkInterfaceIdentityVirtualNetworkInterfaceIdentityByCRN(t *testing.T) {
	checkResult := func(result *vpcv1.DynamicRouteServerMemberPrototypeVirtualNetworkInterfaceVirtualNetworkInterfaceIdentityVirtualNetworkInterfaceIdentityByCRN) {
		model := new(vpcv1.DynamicRouteServerMemberPrototypeVirtualNetworkInterfaceVirtualNetworkInterfaceIdentityVirtualNetworkInterfaceIdentityByCRN)
		model.CRN = core.StringPtr("crn:v1:bluemix:public:is:us-south-1:a/aa2432b1fa4d4ace891e9b80fc104e34::virtual-network-interface:0717-54eb57ee-86f2-4796-90bb-d7874e0831ef")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["crn"] = "crn:v1:bluemix:public:is:us-south-1:a/aa2432b1fa4d4ace891e9b80fc104e34::virtual-network-interface:0717-54eb57ee-86f2-4796-90bb-d7874e0831ef"

	result, err := vpc.ResourceIBMIsDynamicRouteServerMemberMapToDynamicRouteServerMemberPrototypeVirtualNetworkInterfaceVirtualNetworkInterfaceIdentityVirtualNetworkInterfaceIdentityByCRN(model)
	assert.Nil(t, err)
	checkResult(result)
}
