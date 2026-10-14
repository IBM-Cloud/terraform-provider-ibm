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
	vpcName := fmt.Sprintf("tf-drs-vpc-%d", acctest.RandIntRange(10, 100))
	subnetName := fmt.Sprintf("tf-drs-subnet-%d", acctest.RandIntRange(10, 100))
	vniName := fmt.Sprintf("tf-drs-vni-%d", acctest.RandIntRange(10, 100))
	memberVniName := fmt.Sprintf("tf-drs-member-vni-%d", acctest.RandIntRange(10, 100))
	memberName := fmt.Sprintf("tf-drs-member-%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf-drs-member2-%d", acctest.RandIntRange(10, 100))
	nameReplace := fmt.Sprintf("tf-drs-member3-%d", acctest.RandIntRange(10, 100))
	localAsn := fmt.Sprintf("%d", acctest.RandIntRange(64512, 65000))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIsDynamicRouteServerMemberDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsDynamicRouteServerMemberConfigBasic(vpcName, subnetName, vniName, memberVniName, memberName, localAsn, name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIsDynamicRouteServerMemberExists("ibm_is_dynamic_route_server_member.is_dynamic_route_server_member_instance", conf),
					resource.TestCheckResourceAttrPair("ibm_is_dynamic_route_server_member.is_dynamic_route_server_member_instance", "dynamic_route_server_id", "ibm_is_dynamic_route_server.is_dynamic_route_server_instance", "id"),
					resource.TestCheckResourceAttr("ibm_is_dynamic_route_server_member.is_dynamic_route_server_member_instance", "name", name),
					resource.TestCheckResourceAttr("ibm_is_dynamic_route_server_member.is_dynamic_route_server_member_instance", "virtual_network_interfaces.#", "1"),
					resource.TestCheckResourceAttrPair("ibm_is_dynamic_route_server_member.is_dynamic_route_server_member_instance", "virtual_network_interfaces.0.id", "ibm_is_virtual_network_interface.testacc_member_vni", "id"),
					resource.TestCheckResourceAttrSet("ibm_is_dynamic_route_server_member.is_dynamic_route_server_member_instance", "dynamic_route_server_member_id"),
					resource.TestCheckResourceAttrSet("ibm_is_dynamic_route_server_member.is_dynamic_route_server_member_instance", "href"),
					resource.TestCheckResourceAttr("ibm_is_dynamic_route_server_member.is_dynamic_route_server_member_instance", "lifecycle_state", "stable"),
					resource.TestCheckResourceAttr("ibm_is_dynamic_route_server_member.is_dynamic_route_server_member_instance", "resource_type", "dynamic_route_server_member"),
				),
			},
			resource.TestStep{
				// name has no patch counterpart, so changing it replaces the member.
				Config: testAccCheckIBMIsDynamicRouteServerMemberConfigBasic(vpcName, subnetName, vniName, memberVniName, memberName, localAsn, nameReplace),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIsDynamicRouteServerMemberExists("ibm_is_dynamic_route_server_member.is_dynamic_route_server_member_instance", conf),
					resource.TestCheckResourceAttr("ibm_is_dynamic_route_server_member.is_dynamic_route_server_member_instance", "name", nameReplace),
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

func TestAccIBMIsDynamicRouteServerMemberInlineVirtualNetworkInterface(t *testing.T) {
	var conf vpcv1.DynamicRouteServerMember
	vpcName := fmt.Sprintf("tf-drs-vpc-%d", acctest.RandIntRange(10, 100))
	subnetName := fmt.Sprintf("tf-drs-subnet-%d", acctest.RandIntRange(10, 100))
	vniName := fmt.Sprintf("tf-drs-vni-%d", acctest.RandIntRange(10, 100))
	inlineVniName := fmt.Sprintf("tf-drs-inline-vni-%d", acctest.RandIntRange(10, 100))
	memberName := fmt.Sprintf("tf-drs-member-%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf-drs-member2-%d", acctest.RandIntRange(10, 100))
	localAsn := fmt.Sprintf("%d", acctest.RandIntRange(64512, 65000))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIsDynamicRouteServerMemberDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsDynamicRouteServerMemberConfigInline(vpcName, subnetName, vniName, inlineVniName, memberName, localAsn, name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIsDynamicRouteServerMemberExists("ibm_is_dynamic_route_server_member.is_dynamic_route_server_member_instance", conf),
					resource.TestCheckResourceAttr("ibm_is_dynamic_route_server_member.is_dynamic_route_server_member_instance", "name", name),
					resource.TestCheckResourceAttr("ibm_is_dynamic_route_server_member.is_dynamic_route_server_member_instance", "virtual_network_interfaces.0.name", inlineVniName),
					resource.TestCheckResourceAttrPair("ibm_is_dynamic_route_server_member.is_dynamic_route_server_member_instance", "virtual_network_interfaces.0.subnet.0.id", "ibm_is_subnet.testacc_subnet", "id"),
					resource.TestCheckResourceAttrSet("ibm_is_dynamic_route_server_member.is_dynamic_route_server_member_instance", "virtual_network_interfaces.0.id"),
					resource.TestCheckResourceAttrSet("ibm_is_dynamic_route_server_member.is_dynamic_route_server_member_instance", "virtual_network_interfaces.0.primary_ip.0.address"),
					resource.TestCheckResourceAttr("ibm_is_dynamic_route_server_member.is_dynamic_route_server_member_instance", "virtual_network_interfaces.0.allow_ip_spoofing", "false"),
					resource.TestCheckResourceAttr("ibm_is_dynamic_route_server_member.is_dynamic_route_server_member_instance", "virtual_network_interfaces.0.auto_delete", "true"),
					resource.TestCheckResourceAttr("ibm_is_dynamic_route_server_member.is_dynamic_route_server_member_instance", "lifecycle_state", "stable"),
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

func testAccCheckIBMIsDynamicRouteServerMemberConfigBasic(vpcName string, subnetName string, vniName string, memberVniName string, memberName string, localAsn string, name string) string {
	return testAccCheckIBMIsDynamicRouteServerConfigBasic(vpcName, subnetName, vniName, memberName, localAsn) + fmt.Sprintf(`
		resource "ibm_is_virtual_network_interface" "testacc_member_vni" {
			name   = "%s"
			subnet = ibm_is_subnet.testacc_subnet.id
		}

		resource "ibm_is_dynamic_route_server_member" "is_dynamic_route_server_member_instance" {
			dynamic_route_server_id = ibm_is_dynamic_route_server.is_dynamic_route_server_instance.id
			name                    = "%s"
			virtual_network_interfaces {
				id = ibm_is_virtual_network_interface.testacc_member_vni.id
			}
		}
	`, memberVniName, name)
}

func testAccCheckIBMIsDynamicRouteServerMemberConfigInline(vpcName string, subnetName string, vniName string, inlineVniName string, memberName string, localAsn string, name string) string {
	return testAccCheckIBMIsDynamicRouteServerConfigBasic(vpcName, subnetName, vniName, memberName, localAsn) + fmt.Sprintf(`
		resource "ibm_is_dynamic_route_server_member" "is_dynamic_route_server_member_instance" {
			dynamic_route_server_id = ibm_is_dynamic_route_server.is_dynamic_route_server_instance.id
			name                    = "%s"
			virtual_network_interfaces {
				name = "%s"
				subnet {
					id = ibm_is_subnet.testacc_subnet.id
				}
			}
		}
	`, name, inlineVniName)
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

		_, response, err := vpcClient.GetDynamicRouteServerMember(getDynamicRouteServerMemberOptions)

		if err == nil {
			return fmt.Errorf("DynamicRouteServerMember still exists: %s", rs.Primary.ID)
		} else if response == nil || response.StatusCode != 404 {
			return fmt.Errorf("Error checking for DynamicRouteServerMember (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}

func TestResourceIBMIsDynamicRouteServerMemberVirtualNetworkInterfaceReferenceToMap(t *testing.T) {
	result, err := vpc.ResourceIBMIsDynamicRouteServerMemberVirtualNetworkInterfaceReferenceToMap(nil)
	assert.Nil(t, err)
	assert.Equal(t, map[string]interface{}{}, result)

	result, err = vpc.ResourceIBMIsDynamicRouteServerMemberVirtualNetworkInterfaceReferenceToMap(new(vpcv1.VirtualNetworkInterfaceReference))
	assert.Nil(t, err)
	assert.Equal(t, map[string]interface{}{}, result)

	// only an id: primary_ip and subnet are not emitted
	result, err = vpc.ResourceIBMIsDynamicRouteServerMemberVirtualNetworkInterfaceReferenceToMap(&vpcv1.VirtualNetworkInterfaceReference{ID: core.StringPtr("vni-id")})
	assert.Nil(t, err)
	assert.Equal(t, map[string]interface{}{"id": "vni-id"}, result)

	model := testDynamicRouteServerVirtualNetworkInterfaceReference()
	model.Deleted = &vpcv1.Deleted{MoreInfo: core.StringPtr("https://cloud.ibm.com/apidocs/vpc#deleted-resources")}
	expected := testDynamicRouteServerVirtualNetworkInterfaceReferenceMap()
	expected["deleted"] = []map[string]interface{}{{"more_info": "https://cloud.ibm.com/apidocs/vpc#deleted-resources"}}
	result, err = vpc.ResourceIBMIsDynamicRouteServerMemberVirtualNetworkInterfaceReferenceToMap(model)
	assert.Nil(t, err)
	assert.Equal(t, expected, result)
}

func TestResourceIBMIsDynamicRouteServerMemberDeletedToMap(t *testing.T) {
	result, err := vpc.ResourceIBMIsDynamicRouteServerMemberDeletedToMap(nil)
	assert.Nil(t, err)
	assert.Equal(t, map[string]interface{}{}, result)

	model := new(vpcv1.Deleted)
	model.MoreInfo = core.StringPtr("https://cloud.ibm.com/apidocs/vpc#deleted-resources")
	result, err = vpc.ResourceIBMIsDynamicRouteServerMemberDeletedToMap(model)
	assert.Nil(t, err)
	assert.Equal(t, map[string]interface{}{"more_info": "https://cloud.ibm.com/apidocs/vpc#deleted-resources"}, result)
}

func TestResourceIBMIsDynamicRouteServerMemberReservedIPReferenceToMap(t *testing.T) {
	result, err := vpc.ResourceIBMIsDynamicRouteServerMemberReservedIPReferenceToMap(nil)
	assert.Nil(t, err)
	assert.Equal(t, map[string]interface{}{}, result)

	model := testDynamicRouteServerReservedIPReference()
	model.Deleted = &vpcv1.Deleted{MoreInfo: core.StringPtr("https://cloud.ibm.com/apidocs/vpc#deleted-resources")}
	expected := testDynamicRouteServerReservedIPReferenceMap()
	expected["deleted"] = []map[string]interface{}{{"more_info": "https://cloud.ibm.com/apidocs/vpc#deleted-resources"}}
	result, err = vpc.ResourceIBMIsDynamicRouteServerMemberReservedIPReferenceToMap(model)
	assert.Nil(t, err)
	assert.Equal(t, expected, result)
}

func TestResourceIBMIsDynamicRouteServerMemberSubnetReferenceToMap(t *testing.T) {
	result, err := vpc.ResourceIBMIsDynamicRouteServerMemberSubnetReferenceToMap(nil)
	assert.Nil(t, err)
	assert.Equal(t, map[string]interface{}{}, result)

	result, err = vpc.ResourceIBMIsDynamicRouteServerMemberSubnetReferenceToMap(testDynamicRouteServerSubnetReference())
	assert.Nil(t, err)
	assert.Equal(t, testDynamicRouteServerSubnetReferenceMap(), result)
}

func TestResourceIBMIsDynamicRouteServerMemberDynamicRouteServerHealthReasonToMap(t *testing.T) {
	result, err := vpc.ResourceIBMIsDynamicRouteServerMemberDynamicRouteServerHealthReasonToMap(nil)
	assert.Nil(t, err)
	assert.Equal(t, map[string]interface{}{}, result)

	model := new(vpcv1.DynamicRouteServerHealthReason)
	model.Code = core.StringPtr("cannot_reserve_ip_address")
	model.Message = core.StringPtr("The IP address could not be reserved")
	result, err = vpc.ResourceIBMIsDynamicRouteServerMemberDynamicRouteServerHealthReasonToMap(model)
	assert.Nil(t, err)
	assert.Equal(t, map[string]interface{}{"code": "cannot_reserve_ip_address", "message": "The IP address could not be reserved"}, result)
}

func TestResourceIBMIsDynamicRouteServerMemberDynamicRouteServerLifecycleReasonToMap(t *testing.T) {
	result, err := vpc.ResourceIBMIsDynamicRouteServerMemberDynamicRouteServerLifecycleReasonToMap(nil)
	assert.Nil(t, err)
	assert.Equal(t, map[string]interface{}{}, result)

	model := new(vpcv1.DynamicRouteServerLifecycleReason)
	model.Code = core.StringPtr("resource_suspended_by_provider")
	model.Message = core.StringPtr("The resource has been suspended. Contact IBM support with the CRN for next steps.")
	model.MoreInfo = core.StringPtr("https://cloud.ibm.com/apidocs/vpc#resource-suspension")
	result, err = vpc.ResourceIBMIsDynamicRouteServerMemberDynamicRouteServerLifecycleReasonToMap(model)
	assert.Nil(t, err)
	assert.Equal(t, map[string]interface{}{
		"code":      "resource_suspended_by_provider",
		"message":   "The resource has been suspended. Contact IBM support with the CRN for next steps.",
		"more_info": "https://cloud.ibm.com/apidocs/vpc#resource-suspension",
	}, result)
}

func TestResourceIBMIsDynamicRouteServerMemberMapToDynamicRouteServerMemberPrototypeVirtualNetworkInterface(t *testing.T) {
	t.Run("nil map", func(t *testing.T) {
		_, err := vpc.ResourceIBMIsDynamicRouteServerMemberMapToDynamicRouteServerMemberPrototypeVirtualNetworkInterface(nil)
		assert.NotNil(t, err)
	})

	t.Run("empty map requires subnet", func(t *testing.T) {
		_, err := vpc.ResourceIBMIsDynamicRouteServerMemberMapToDynamicRouteServerMemberPrototypeVirtualNetworkInterface(map[string]interface{}{})
		assert.NotNil(t, err)
	})

	t.Run("identity by id", func(t *testing.T) {
		result, err := vpc.ResourceIBMIsDynamicRouteServerMemberMapToDynamicRouteServerMemberPrototypeVirtualNetworkInterface(map[string]interface{}{"id": "0717-54eb57ee-86f2-4796-90bb-d7874e0831ef", "href": "", "crn": ""})
		assert.Nil(t, err)
		assert.Equal(t, &vpcv1.DynamicRouteServerMemberPrototypeVirtualNetworkInterfaceVirtualNetworkInterfaceIdentityVirtualNetworkInterfaceIdentityByID{ID: core.StringPtr("0717-54eb57ee-86f2-4796-90bb-d7874e0831ef")}, result)
	})

	t.Run("identity by href", func(t *testing.T) {
		result, err := vpc.ResourceIBMIsDynamicRouteServerMemberMapToDynamicRouteServerMemberPrototypeVirtualNetworkInterface(map[string]interface{}{"href": "https://us-south.iaas.cloud.ibm.com/v1/virtual_network_interfaces/0717-54eb57ee-86f2-4796-90bb-d7874e0831ef"})
		assert.Nil(t, err)
		assert.Equal(t, &vpcv1.DynamicRouteServerMemberPrototypeVirtualNetworkInterfaceVirtualNetworkInterfaceIdentityVirtualNetworkInterfaceIdentityByHref{Href: core.StringPtr("https://us-south.iaas.cloud.ibm.com/v1/virtual_network_interfaces/0717-54eb57ee-86f2-4796-90bb-d7874e0831ef")}, result)
	})

	t.Run("identity by crn", func(t *testing.T) {
		result, err := vpc.ResourceIBMIsDynamicRouteServerMemberMapToDynamicRouteServerMemberPrototypeVirtualNetworkInterface(map[string]interface{}{"crn": "crn:v1:bluemix:public:is:us-south-1:a/aa2432b1fa4d4ace891e9b80fc104e34::virtual-network-interface:0717-54eb57ee-86f2-4796-90bb-d7874e0831ef"})
		assert.Nil(t, err)
		assert.Equal(t, &vpcv1.DynamicRouteServerMemberPrototypeVirtualNetworkInterfaceVirtualNetworkInterfaceIdentityVirtualNetworkInterfaceIdentityByCRN{CRN: core.StringPtr("crn:v1:bluemix:public:is:us-south-1:a/aa2432b1fa4d4ace891e9b80fc104e34::virtual-network-interface:0717-54eb57ee-86f2-4796-90bb-d7874e0831ef")}, result)
	})

	t.Run("identity with prototype fields ignores them", func(t *testing.T) {
		result, err := vpc.ResourceIBMIsDynamicRouteServerMemberMapToDynamicRouteServerMemberPrototypeVirtualNetworkInterface(map[string]interface{}{"id": "vni-id", "name": "ignored", "auto_delete": false})
		assert.Nil(t, err)
		assert.Equal(t, &vpcv1.DynamicRouteServerMemberPrototypeVirtualNetworkInterfaceVirtualNetworkInterfaceIdentityVirtualNetworkInterfaceIdentityByID{ID: core.StringPtr("vni-id")}, result)
	})

	t.Run("more than one identity key", func(t *testing.T) {
		_, err := vpc.ResourceIBMIsDynamicRouteServerMemberMapToDynamicRouteServerMemberPrototypeVirtualNetworkInterface(map[string]interface{}{"id": "a", "crn": "b"})
		assert.NotNil(t, err)
	})

	t.Run("prototype with all arguments", func(t *testing.T) {
		model := make(map[string]interface{})
		model["allow_ip_spoofing"] = true
		model["auto_delete"] = false
		model["enable_infrastructure_nat"] = true
		model["ips"] = []interface{}{map[string]interface{}{"address": "10.0.0.5", "auto_delete": false, "name": "my-reserved-ip"}}
		model["name"] = "my-virtual-network-interface"
		model["primary_ip"] = []interface{}{map[string]interface{}{"id": "0717-6d353a0f-aeb1-4ae1-832e-1110d10981bb"}}
		model["protocol_state_filtering_mode"] = "auto"
		model["resource_group"] = []interface{}{map[string]interface{}{"id": "fee82deba12e4c0fb69c3b09d1f12345"}}
		model["security_groups"] = []interface{}{map[string]interface{}{"id": "r006-be5df5ca-12a0-494b-907e-aa6ec2bfa271"}}
		model["subnet"] = []interface{}{map[string]interface{}{"id": "0717-7ec86020-1c6e-4889-b3f0-a15f2e50f87e"}}

		result, err := vpc.ResourceIBMIsDynamicRouteServerMemberMapToDynamicRouteServerMemberPrototypeVirtualNetworkInterface(model)
		assert.Nil(t, err)

		expected := new(vpcv1.DynamicRouteServerMemberPrototypeVirtualNetworkInterfaceVirtualNetworkInterfacePrototypeDynamicRouteServerContext)
		expected.AllowIPSpoofing = core.BoolPtr(true)
		expected.AutoDelete = core.BoolPtr(false)
		expected.EnableInfrastructureNat = core.BoolPtr(true)
		expected.Ips = []vpcv1.VirtualNetworkInterfaceIPPrototypeIntf{&vpcv1.VirtualNetworkInterfaceIPPrototypeReservedIPPrototypeVirtualNetworkInterfaceIPsContext{Address: core.StringPtr("10.0.0.5"), AutoDelete: core.BoolPtr(false), Name: core.StringPtr("my-reserved-ip")}}
		expected.Name = core.StringPtr("my-virtual-network-interface")
		expected.PrimaryIP = &vpcv1.VirtualNetworkInterfacePrimaryIPPrototypeReservedIPIdentityVirtualNetworkInterfacePrimaryIPContextByID{ID: core.StringPtr("0717-6d353a0f-aeb1-4ae1-832e-1110d10981bb")}
		expected.ProtocolStateFilteringMode = core.StringPtr("auto")
		expected.ResourceGroup = &vpcv1.ResourceGroupIdentityByID{ID: core.StringPtr("fee82deba12e4c0fb69c3b09d1f12345")}
		expected.SecurityGroups = []vpcv1.SecurityGroupIdentityIntf{&vpcv1.SecurityGroupIdentityByID{ID: core.StringPtr("r006-be5df5ca-12a0-494b-907e-aa6ec2bfa271")}}
		expected.Subnet = &vpcv1.SubnetIdentityByID{ID: core.StringPtr("0717-7ec86020-1c6e-4889-b3f0-a15f2e50f87e")}
		assert.Equal(t, expected, result)
	})

	t.Run("prototype with subnet only and empty lists", func(t *testing.T) {
		model := map[string]interface{}{
			"id": "", "href": "", "crn": "", "name": "", "protocol_state_filtering_mode": "",
			"ips": []interface{}{}, "security_groups": []interface{}{}, "resource_group": []interface{}{}, "primary_ip": []interface{}{},
			"subnet": []interface{}{map[string]interface{}{"crn": "crn:v1:bluemix:public:is:us-south-1:a/aa2432b1fa4d4ace891e9b80fc104e34::subnet:0717-7ec86020-1c6e-4889-b3f0-a15f2e50f87e"}},
		}
		result, err := vpc.ResourceIBMIsDynamicRouteServerMemberMapToDynamicRouteServerMemberPrototypeVirtualNetworkInterface(model)
		assert.Nil(t, err)
		expected := new(vpcv1.DynamicRouteServerMemberPrototypeVirtualNetworkInterfaceVirtualNetworkInterfacePrototypeDynamicRouteServerContext)
		expected.Subnet = &vpcv1.SubnetIdentityByCRN{CRN: core.StringPtr("crn:v1:bluemix:public:is:us-south-1:a/aa2432b1fa4d4ace891e9b80fc104e34::subnet:0717-7ec86020-1c6e-4889-b3f0-a15f2e50f87e")}
		assert.Equal(t, expected, result)
	})

	t.Run("prototype with invalid nested entries", func(t *testing.T) {
		base := func() map[string]interface{} {
			return map[string]interface{}{"subnet": []interface{}{map[string]interface{}{"id": "subnet-id"}}}
		}
		m := base()
		m["ips"] = []interface{}{map[string]interface{}{"id": "a", "href": "b"}}
		_, err := vpc.ResourceIBMIsDynamicRouteServerMemberMapToDynamicRouteServerMemberPrototypeVirtualNetworkInterface(m)
		assert.NotNil(t, err)

		m = base()
		m["security_groups"] = []interface{}{map[string]interface{}{}}
		_, err = vpc.ResourceIBMIsDynamicRouteServerMemberMapToDynamicRouteServerMemberPrototypeVirtualNetworkInterface(m)
		assert.NotNil(t, err)

		m = base()
		m["resource_group"] = []interface{}{map[string]interface{}{"id": ""}}
		_, err = vpc.ResourceIBMIsDynamicRouteServerMemberMapToDynamicRouteServerMemberPrototypeVirtualNetworkInterface(m)
		assert.NotNil(t, err)

		m = base()
		m["primary_ip"] = []interface{}{map[string]interface{}{"id": "a", "href": "b"}}
		_, err = vpc.ResourceIBMIsDynamicRouteServerMemberMapToDynamicRouteServerMemberPrototypeVirtualNetworkInterface(m)
		assert.NotNil(t, err)

		m = map[string]interface{}{"subnet": []interface{}{map[string]interface{}{"id": "a", "crn": "b"}}}
		_, err = vpc.ResourceIBMIsDynamicRouteServerMemberMapToDynamicRouteServerMemberPrototypeVirtualNetworkInterface(m)
		assert.NotNil(t, err)
	})
}

func TestResourceIBMIsDynamicRouteServerMemberMapToVirtualNetworkInterfaceIPPrototype(t *testing.T) {
	result, err := vpc.ResourceIBMIsDynamicRouteServerMemberMapToVirtualNetworkInterfaceIPPrototype(map[string]interface{}{"id": "0717-6d353a0f-aeb1-4ae1-832e-1110d10981bb"})
	assert.Nil(t, err)
	assert.Equal(t, &vpcv1.VirtualNetworkInterfaceIPPrototypeReservedIPIdentityVirtualNetworkInterfaceIPsContextByID{ID: core.StringPtr("0717-6d353a0f-aeb1-4ae1-832e-1110d10981bb")}, result)

	result, err = vpc.ResourceIBMIsDynamicRouteServerMemberMapToVirtualNetworkInterfaceIPPrototype(map[string]interface{}{"href": "https://us-south.iaas.cloud.ibm.com/v1/subnets/0717-bea6a632-5e13-42a4-b4b8-31dc877abfe4/reserved_ips/0717-6d353a0f-aeb1-4ae1-832e-1110d10981bb"})
	assert.Nil(t, err)
	assert.Equal(t, &vpcv1.VirtualNetworkInterfaceIPPrototypeReservedIPIdentityVirtualNetworkInterfaceIPsContextByHref{Href: core.StringPtr("https://us-south.iaas.cloud.ibm.com/v1/subnets/0717-bea6a632-5e13-42a4-b4b8-31dc877abfe4/reserved_ips/0717-6d353a0f-aeb1-4ae1-832e-1110d10981bb")}, result)

	// no identity: a reserved IP prototype, even when empty
	result, err = vpc.ResourceIBMIsDynamicRouteServerMemberMapToVirtualNetworkInterfaceIPPrototype(map[string]interface{}{})
	assert.Nil(t, err)
	assert.Equal(t, &vpcv1.VirtualNetworkInterfaceIPPrototypeReservedIPPrototypeVirtualNetworkInterfaceIPsContext{}, result)

	result, err = vpc.ResourceIBMIsDynamicRouteServerMemberMapToVirtualNetworkInterfaceIPPrototype(map[string]interface{}{"address": "10.0.0.5", "auto_delete": false, "name": "my-reserved-ip"})
	assert.Nil(t, err)
	assert.Equal(t, &vpcv1.VirtualNetworkInterfaceIPPrototypeReservedIPPrototypeVirtualNetworkInterfaceIPsContext{Address: core.StringPtr("10.0.0.5"), AutoDelete: core.BoolPtr(false), Name: core.StringPtr("my-reserved-ip")}, result)

	_, err = vpc.ResourceIBMIsDynamicRouteServerMemberMapToVirtualNetworkInterfaceIPPrototype(map[string]interface{}{"id": "a", "href": "b"})
	assert.NotNil(t, err)
}

func TestResourceIBMIsDynamicRouteServerMemberMapToVirtualNetworkInterfaceIPPrototypeReservedIPIdentityVirtualNetworkInterfaceIPsContext(t *testing.T) {
	result, err := vpc.ResourceIBMIsDynamicRouteServerMemberMapToVirtualNetworkInterfaceIPPrototypeReservedIPIdentityVirtualNetworkInterfaceIPsContext(map[string]interface{}{"id": "rip-id"})
	assert.Nil(t, err)
	assert.Equal(t, &vpcv1.VirtualNetworkInterfaceIPPrototypeReservedIPIdentityVirtualNetworkInterfaceIPsContextByID{ID: core.StringPtr("rip-id")}, result)

	result, err = vpc.ResourceIBMIsDynamicRouteServerMemberMapToVirtualNetworkInterfaceIPPrototypeReservedIPIdentityVirtualNetworkInterfaceIPsContext(map[string]interface{}{"href": "rip-href"})
	assert.Nil(t, err)
	assert.Equal(t, &vpcv1.VirtualNetworkInterfaceIPPrototypeReservedIPIdentityVirtualNetworkInterfaceIPsContextByHref{Href: core.StringPtr("rip-href")}, result)

	_, err = vpc.ResourceIBMIsDynamicRouteServerMemberMapToVirtualNetworkInterfaceIPPrototypeReservedIPIdentityVirtualNetworkInterfaceIPsContext(map[string]interface{}{})
	assert.NotNil(t, err)
}

func TestResourceIBMIsDynamicRouteServerMemberMapToVirtualNetworkInterfaceIPPrototypeReservedIPIdentityVirtualNetworkInterfaceIPsContextByID(t *testing.T) {
	result, err := vpc.ResourceIBMIsDynamicRouteServerMemberMapToVirtualNetworkInterfaceIPPrototypeReservedIPIdentityVirtualNetworkInterfaceIPsContextByID(map[string]interface{}{"id": "rip-id"})
	assert.Nil(t, err)
	assert.Equal(t, &vpcv1.VirtualNetworkInterfaceIPPrototypeReservedIPIdentityVirtualNetworkInterfaceIPsContextByID{ID: core.StringPtr("rip-id")}, result)

	_, err = vpc.ResourceIBMIsDynamicRouteServerMemberMapToVirtualNetworkInterfaceIPPrototypeReservedIPIdentityVirtualNetworkInterfaceIPsContextByID(map[string]interface{}{})
	assert.NotNil(t, err)
}

func TestResourceIBMIsDynamicRouteServerMemberMapToVirtualNetworkInterfaceIPPrototypeReservedIPIdentityVirtualNetworkInterfaceIPsContextByHref(t *testing.T) {
	result, err := vpc.ResourceIBMIsDynamicRouteServerMemberMapToVirtualNetworkInterfaceIPPrototypeReservedIPIdentityVirtualNetworkInterfaceIPsContextByHref(map[string]interface{}{"href": "rip-href"})
	assert.Nil(t, err)
	assert.Equal(t, &vpcv1.VirtualNetworkInterfaceIPPrototypeReservedIPIdentityVirtualNetworkInterfaceIPsContextByHref{Href: core.StringPtr("rip-href")}, result)

	_, err = vpc.ResourceIBMIsDynamicRouteServerMemberMapToVirtualNetworkInterfaceIPPrototypeReservedIPIdentityVirtualNetworkInterfaceIPsContextByHref(map[string]interface{}{"href": ""})
	assert.NotNil(t, err)
}

func TestResourceIBMIsDynamicRouteServerMemberMapToVirtualNetworkInterfaceIPPrototypeReservedIPPrototypeVirtualNetworkInterfaceIPsContext(t *testing.T) {
	result, err := vpc.ResourceIBMIsDynamicRouteServerMemberMapToVirtualNetworkInterfaceIPPrototypeReservedIPPrototypeVirtualNetworkInterfaceIPsContext(map[string]interface{}{})
	assert.Nil(t, err)
	assert.Equal(t, &vpcv1.VirtualNetworkInterfaceIPPrototypeReservedIPPrototypeVirtualNetworkInterfaceIPsContext{}, result)

	// auto_delete false is kept, it is not the same as unset
	result, err = vpc.ResourceIBMIsDynamicRouteServerMemberMapToVirtualNetworkInterfaceIPPrototypeReservedIPPrototypeVirtualNetworkInterfaceIPsContext(map[string]interface{}{"address": "10.0.0.5", "auto_delete": false, "name": ""})
	assert.Nil(t, err)
	assert.Equal(t, &vpcv1.VirtualNetworkInterfaceIPPrototypeReservedIPPrototypeVirtualNetworkInterfaceIPsContext{Address: core.StringPtr("10.0.0.5"), AutoDelete: core.BoolPtr(false)}, result)
}

func TestResourceIBMIsDynamicRouteServerMemberMapToVirtualNetworkInterfacePrimaryIPPrototype(t *testing.T) {
	result, err := vpc.ResourceIBMIsDynamicRouteServerMemberMapToVirtualNetworkInterfacePrimaryIPPrototype(map[string]interface{}{"id": "0717-6d353a0f-aeb1-4ae1-832e-1110d10981bb"})
	assert.Nil(t, err)
	assert.Equal(t, &vpcv1.VirtualNetworkInterfacePrimaryIPPrototypeReservedIPIdentityVirtualNetworkInterfacePrimaryIPContextByID{ID: core.StringPtr("0717-6d353a0f-aeb1-4ae1-832e-1110d10981bb")}, result)

	result, err = vpc.ResourceIBMIsDynamicRouteServerMemberMapToVirtualNetworkInterfacePrimaryIPPrototype(map[string]interface{}{"href": "rip-href"})
	assert.Nil(t, err)
	assert.Equal(t, &vpcv1.VirtualNetworkInterfacePrimaryIPPrototypeReservedIPIdentityVirtualNetworkInterfacePrimaryIPContextByHref{Href: core.StringPtr("rip-href")}, result)

	result, err = vpc.ResourceIBMIsDynamicRouteServerMemberMapToVirtualNetworkInterfacePrimaryIPPrototype(map[string]interface{}{"address": "10.0.0.5", "auto_delete": true, "name": "my-reserved-ip"})
	assert.Nil(t, err)
	assert.Equal(t, &vpcv1.VirtualNetworkInterfacePrimaryIPPrototypeReservedIPPrototypeVirtualNetworkInterfacePrimaryIPContext{Address: core.StringPtr("10.0.0.5"), AutoDelete: core.BoolPtr(true), Name: core.StringPtr("my-reserved-ip")}, result)

	result, err = vpc.ResourceIBMIsDynamicRouteServerMemberMapToVirtualNetworkInterfacePrimaryIPPrototype(map[string]interface{}{})
	assert.Nil(t, err)
	assert.Equal(t, &vpcv1.VirtualNetworkInterfacePrimaryIPPrototypeReservedIPPrototypeVirtualNetworkInterfacePrimaryIPContext{}, result)

	_, err = vpc.ResourceIBMIsDynamicRouteServerMemberMapToVirtualNetworkInterfacePrimaryIPPrototype(map[string]interface{}{"id": "a", "href": "b"})
	assert.NotNil(t, err)
}

func TestResourceIBMIsDynamicRouteServerMemberMapToVirtualNetworkInterfacePrimaryIPPrototypeReservedIPIdentityVirtualNetworkInterfacePrimaryIPContext(t *testing.T) {
	result, err := vpc.ResourceIBMIsDynamicRouteServerMemberMapToVirtualNetworkInterfacePrimaryIPPrototypeReservedIPIdentityVirtualNetworkInterfacePrimaryIPContext(map[string]interface{}{"id": "rip-id"})
	assert.Nil(t, err)
	assert.Equal(t, &vpcv1.VirtualNetworkInterfacePrimaryIPPrototypeReservedIPIdentityVirtualNetworkInterfacePrimaryIPContextByID{ID: core.StringPtr("rip-id")}, result)

	result, err = vpc.ResourceIBMIsDynamicRouteServerMemberMapToVirtualNetworkInterfacePrimaryIPPrototypeReservedIPIdentityVirtualNetworkInterfacePrimaryIPContext(map[string]interface{}{"href": "rip-href"})
	assert.Nil(t, err)
	assert.Equal(t, &vpcv1.VirtualNetworkInterfacePrimaryIPPrototypeReservedIPIdentityVirtualNetworkInterfacePrimaryIPContextByHref{Href: core.StringPtr("rip-href")}, result)

	_, err = vpc.ResourceIBMIsDynamicRouteServerMemberMapToVirtualNetworkInterfacePrimaryIPPrototypeReservedIPIdentityVirtualNetworkInterfacePrimaryIPContext(map[string]interface{}{"id": "", "href": ""})
	assert.NotNil(t, err)
}

func TestResourceIBMIsDynamicRouteServerMemberMapToVirtualNetworkInterfacePrimaryIPPrototypeReservedIPIdentityVirtualNetworkInterfacePrimaryIPContextByID(t *testing.T) {
	result, err := vpc.ResourceIBMIsDynamicRouteServerMemberMapToVirtualNetworkInterfacePrimaryIPPrototypeReservedIPIdentityVirtualNetworkInterfacePrimaryIPContextByID(map[string]interface{}{"id": "rip-id"})
	assert.Nil(t, err)
	assert.Equal(t, &vpcv1.VirtualNetworkInterfacePrimaryIPPrototypeReservedIPIdentityVirtualNetworkInterfacePrimaryIPContextByID{ID: core.StringPtr("rip-id")}, result)

	_, err = vpc.ResourceIBMIsDynamicRouteServerMemberMapToVirtualNetworkInterfacePrimaryIPPrototypeReservedIPIdentityVirtualNetworkInterfacePrimaryIPContextByID(map[string]interface{}{})
	assert.NotNil(t, err)
}

func TestResourceIBMIsDynamicRouteServerMemberMapToVirtualNetworkInterfacePrimaryIPPrototypeReservedIPIdentityVirtualNetworkInterfacePrimaryIPContextByHref(t *testing.T) {
	result, err := vpc.ResourceIBMIsDynamicRouteServerMemberMapToVirtualNetworkInterfacePrimaryIPPrototypeReservedIPIdentityVirtualNetworkInterfacePrimaryIPContextByHref(map[string]interface{}{"href": "rip-href"})
	assert.Nil(t, err)
	assert.Equal(t, &vpcv1.VirtualNetworkInterfacePrimaryIPPrototypeReservedIPIdentityVirtualNetworkInterfacePrimaryIPContextByHref{Href: core.StringPtr("rip-href")}, result)

	_, err = vpc.ResourceIBMIsDynamicRouteServerMemberMapToVirtualNetworkInterfacePrimaryIPPrototypeReservedIPIdentityVirtualNetworkInterfacePrimaryIPContextByHref(map[string]interface{}{})
	assert.NotNil(t, err)
}

func TestResourceIBMIsDynamicRouteServerMemberMapToVirtualNetworkInterfacePrimaryIPPrototypeReservedIPPrototypeVirtualNetworkInterfacePrimaryIPContext(t *testing.T) {
	result, err := vpc.ResourceIBMIsDynamicRouteServerMemberMapToVirtualNetworkInterfacePrimaryIPPrototypeReservedIPPrototypeVirtualNetworkInterfacePrimaryIPContext(map[string]interface{}{})
	assert.Nil(t, err)
	assert.Equal(t, &vpcv1.VirtualNetworkInterfacePrimaryIPPrototypeReservedIPPrototypeVirtualNetworkInterfacePrimaryIPContext{}, result)

	result, err = vpc.ResourceIBMIsDynamicRouteServerMemberMapToVirtualNetworkInterfacePrimaryIPPrototypeReservedIPPrototypeVirtualNetworkInterfacePrimaryIPContext(map[string]interface{}{"address": "10.0.0.5", "auto_delete": false, "name": "my-reserved-ip"})
	assert.Nil(t, err)
	assert.Equal(t, &vpcv1.VirtualNetworkInterfacePrimaryIPPrototypeReservedIPPrototypeVirtualNetworkInterfacePrimaryIPContext{Address: core.StringPtr("10.0.0.5"), AutoDelete: core.BoolPtr(false), Name: core.StringPtr("my-reserved-ip")}, result)
}

func TestResourceIBMIsDynamicRouteServerMemberMapToResourceGroupIdentity(t *testing.T) {
	result, err := vpc.ResourceIBMIsDynamicRouteServerMemberMapToResourceGroupIdentity(map[string]interface{}{"id": "fee82deba12e4c0fb69c3b09d1f12345"})
	assert.Nil(t, err)
	assert.Equal(t, &vpcv1.ResourceGroupIdentityByID{ID: core.StringPtr("fee82deba12e4c0fb69c3b09d1f12345")}, result)

	_, err = vpc.ResourceIBMIsDynamicRouteServerMemberMapToResourceGroupIdentity(map[string]interface{}{})
	assert.NotNil(t, err)
}

func TestResourceIBMIsDynamicRouteServerMemberMapToResourceGroupIdentityByID(t *testing.T) {
	result, err := vpc.ResourceIBMIsDynamicRouteServerMemberMapToResourceGroupIdentityByID(map[string]interface{}{"id": "fee82deba12e4c0fb69c3b09d1f12345"})
	assert.Nil(t, err)
	assert.Equal(t, &vpcv1.ResourceGroupIdentityByID{ID: core.StringPtr("fee82deba12e4c0fb69c3b09d1f12345")}, result)

	_, err = vpc.ResourceIBMIsDynamicRouteServerMemberMapToResourceGroupIdentityByID(map[string]interface{}{"id": ""})
	assert.NotNil(t, err)
}

func TestResourceIBMIsDynamicRouteServerMemberMapToSecurityGroupIdentity(t *testing.T) {
	result, err := vpc.ResourceIBMIsDynamicRouteServerMemberMapToSecurityGroupIdentity(map[string]interface{}{"id": "sg-id", "crn": "", "href": ""})
	assert.Nil(t, err)
	assert.Equal(t, &vpcv1.SecurityGroupIdentityByID{ID: core.StringPtr("sg-id")}, result)

	result, err = vpc.ResourceIBMIsDynamicRouteServerMemberMapToSecurityGroupIdentity(map[string]interface{}{"crn": "sg-crn"})
	assert.Nil(t, err)
	assert.Equal(t, &vpcv1.SecurityGroupIdentityByCRN{CRN: core.StringPtr("sg-crn")}, result)

	result, err = vpc.ResourceIBMIsDynamicRouteServerMemberMapToSecurityGroupIdentity(map[string]interface{}{"href": "sg-href"})
	assert.Nil(t, err)
	assert.Equal(t, &vpcv1.SecurityGroupIdentityByHref{Href: core.StringPtr("sg-href")}, result)

	_, err = vpc.ResourceIBMIsDynamicRouteServerMemberMapToSecurityGroupIdentity(map[string]interface{}{})
	assert.NotNil(t, err)

	_, err = vpc.ResourceIBMIsDynamicRouteServerMemberMapToSecurityGroupIdentity(map[string]interface{}{"id": "a", "href": "b"})
	assert.NotNil(t, err)
}

func TestResourceIBMIsDynamicRouteServerMemberMapToSecurityGroupIdentityByID(t *testing.T) {
	result, err := vpc.ResourceIBMIsDynamicRouteServerMemberMapToSecurityGroupIdentityByID(map[string]interface{}{"id": "sg-id"})
	assert.Nil(t, err)
	assert.Equal(t, &vpcv1.SecurityGroupIdentityByID{ID: core.StringPtr("sg-id")}, result)

	_, err = vpc.ResourceIBMIsDynamicRouteServerMemberMapToSecurityGroupIdentityByID(map[string]interface{}{})
	assert.NotNil(t, err)
}

func TestResourceIBMIsDynamicRouteServerMemberMapToSecurityGroupIdentityByCRN(t *testing.T) {
	result, err := vpc.ResourceIBMIsDynamicRouteServerMemberMapToSecurityGroupIdentityByCRN(map[string]interface{}{"crn": "sg-crn"})
	assert.Nil(t, err)
	assert.Equal(t, &vpcv1.SecurityGroupIdentityByCRN{CRN: core.StringPtr("sg-crn")}, result)

	_, err = vpc.ResourceIBMIsDynamicRouteServerMemberMapToSecurityGroupIdentityByCRN(map[string]interface{}{})
	assert.NotNil(t, err)
}

func TestResourceIBMIsDynamicRouteServerMemberMapToSecurityGroupIdentityByHref(t *testing.T) {
	result, err := vpc.ResourceIBMIsDynamicRouteServerMemberMapToSecurityGroupIdentityByHref(map[string]interface{}{"href": "sg-href"})
	assert.Nil(t, err)
	assert.Equal(t, &vpcv1.SecurityGroupIdentityByHref{Href: core.StringPtr("sg-href")}, result)

	_, err = vpc.ResourceIBMIsDynamicRouteServerMemberMapToSecurityGroupIdentityByHref(map[string]interface{}{})
	assert.NotNil(t, err)
}

func TestResourceIBMIsDynamicRouteServerMemberMapToSubnetIdentity(t *testing.T) {
	result, err := vpc.ResourceIBMIsDynamicRouteServerMemberMapToSubnetIdentity(map[string]interface{}{"id": "0717-7ec86020-1c6e-4889-b3f0-a15f2e50f87e"})
	assert.Nil(t, err)
	assert.Equal(t, &vpcv1.SubnetIdentityByID{ID: core.StringPtr("0717-7ec86020-1c6e-4889-b3f0-a15f2e50f87e")}, result)

	result, err = vpc.ResourceIBMIsDynamicRouteServerMemberMapToSubnetIdentity(map[string]interface{}{"crn": "subnet-crn"})
	assert.Nil(t, err)
	assert.Equal(t, &vpcv1.SubnetIdentityByCRN{CRN: core.StringPtr("subnet-crn")}, result)

	result, err = vpc.ResourceIBMIsDynamicRouteServerMemberMapToSubnetIdentity(map[string]interface{}{"href": "subnet-href"})
	assert.Nil(t, err)
	assert.Equal(t, &vpcv1.SubnetIdentityByHref{Href: core.StringPtr("subnet-href")}, result)

	_, err = vpc.ResourceIBMIsDynamicRouteServerMemberMapToSubnetIdentity(map[string]interface{}{"id": "", "crn": "", "href": ""})
	assert.NotNil(t, err)

	_, err = vpc.ResourceIBMIsDynamicRouteServerMemberMapToSubnetIdentity(map[string]interface{}{"id": "a", "crn": "b", "href": "c"})
	assert.NotNil(t, err)
}

func TestResourceIBMIsDynamicRouteServerMemberMapToSubnetIdentityByID(t *testing.T) {
	result, err := vpc.ResourceIBMIsDynamicRouteServerMemberMapToSubnetIdentityByID(map[string]interface{}{"id": "subnet-id"})
	assert.Nil(t, err)
	assert.Equal(t, &vpcv1.SubnetIdentityByID{ID: core.StringPtr("subnet-id")}, result)

	_, err = vpc.ResourceIBMIsDynamicRouteServerMemberMapToSubnetIdentityByID(map[string]interface{}{})
	assert.NotNil(t, err)
}

func TestResourceIBMIsDynamicRouteServerMemberMapToSubnetIdentityByCRN(t *testing.T) {
	result, err := vpc.ResourceIBMIsDynamicRouteServerMemberMapToSubnetIdentityByCRN(map[string]interface{}{"crn": "subnet-crn"})
	assert.Nil(t, err)
	assert.Equal(t, &vpcv1.SubnetIdentityByCRN{CRN: core.StringPtr("subnet-crn")}, result)

	_, err = vpc.ResourceIBMIsDynamicRouteServerMemberMapToSubnetIdentityByCRN(map[string]interface{}{})
	assert.NotNil(t, err)
}

func TestResourceIBMIsDynamicRouteServerMemberMapToSubnetIdentityByHref(t *testing.T) {
	result, err := vpc.ResourceIBMIsDynamicRouteServerMemberMapToSubnetIdentityByHref(map[string]interface{}{"href": "subnet-href"})
	assert.Nil(t, err)
	assert.Equal(t, &vpcv1.SubnetIdentityByHref{Href: core.StringPtr("subnet-href")}, result)

	_, err = vpc.ResourceIBMIsDynamicRouteServerMemberMapToSubnetIdentityByHref(map[string]interface{}{})
	assert.NotNil(t, err)
}

func TestResourceIBMIsDynamicRouteServerMemberMapToDynamicRouteServerMemberPrototypeVirtualNetworkInterfaceVirtualNetworkInterfacePrototypeDynamicRouteServerContext(t *testing.T) {
	_, err := vpc.ResourceIBMIsDynamicRouteServerMemberMapToDynamicRouteServerMemberPrototypeVirtualNetworkInterfaceVirtualNetworkInterfacePrototypeDynamicRouteServerContext(map[string]interface{}{})
	assert.NotNil(t, err)

	_, err = vpc.ResourceIBMIsDynamicRouteServerMemberMapToDynamicRouteServerMemberPrototypeVirtualNetworkInterfaceVirtualNetworkInterfacePrototypeDynamicRouteServerContext(map[string]interface{}{"subnet": []interface{}{}})
	assert.NotNil(t, err)

	model := map[string]interface{}{
		"subnet":                    []interface{}{map[string]interface{}{"href": "subnet-href"}},
		"allow_ip_spoofing":         false,
		"auto_delete":               true,
		"enable_infrastructure_nat": false,
		"ips": []interface{}{
			map[string]interface{}{"id": "rip-1"},
			map[string]interface{}{"address": "10.0.0.6"},
		},
		"primary_ip": []interface{}{map[string]interface{}{"address": "10.0.0.5", "auto_delete": true}},
		"security_groups": []interface{}{
			map[string]interface{}{"id": "sg-1"},
			map[string]interface{}{"crn": "sg-2-crn"},
		},
	}
	result, err := vpc.ResourceIBMIsDynamicRouteServerMemberMapToDynamicRouteServerMemberPrototypeVirtualNetworkInterfaceVirtualNetworkInterfacePrototypeDynamicRouteServerContext(model)
	assert.Nil(t, err)
	expected := new(vpcv1.DynamicRouteServerMemberPrototypeVirtualNetworkInterfaceVirtualNetworkInterfacePrototypeDynamicRouteServerContext)
	expected.Subnet = &vpcv1.SubnetIdentityByHref{Href: core.StringPtr("subnet-href")}
	expected.AllowIPSpoofing = core.BoolPtr(false)
	expected.AutoDelete = core.BoolPtr(true)
	expected.EnableInfrastructureNat = core.BoolPtr(false)
	expected.Ips = []vpcv1.VirtualNetworkInterfaceIPPrototypeIntf{
		&vpcv1.VirtualNetworkInterfaceIPPrototypeReservedIPIdentityVirtualNetworkInterfaceIPsContextByID{ID: core.StringPtr("rip-1")},
		&vpcv1.VirtualNetworkInterfaceIPPrototypeReservedIPPrototypeVirtualNetworkInterfaceIPsContext{Address: core.StringPtr("10.0.0.6")},
	}
	expected.PrimaryIP = &vpcv1.VirtualNetworkInterfacePrimaryIPPrototypeReservedIPPrototypeVirtualNetworkInterfacePrimaryIPContext{Address: core.StringPtr("10.0.0.5"), AutoDelete: core.BoolPtr(true)}
	expected.SecurityGroups = []vpcv1.SecurityGroupIdentityIntf{
		&vpcv1.SecurityGroupIdentityByID{ID: core.StringPtr("sg-1")},
		&vpcv1.SecurityGroupIdentityByCRN{CRN: core.StringPtr("sg-2-crn")},
	}
	assert.Equal(t, expected, result)
}

func TestResourceIBMIsDynamicRouteServerMemberMapToDynamicRouteServerMemberPrototypeVirtualNetworkInterfaceVirtualNetworkInterfaceIdentity(t *testing.T) {
	result, err := vpc.ResourceIBMIsDynamicRouteServerMemberMapToDynamicRouteServerMemberPrototypeVirtualNetworkInterfaceVirtualNetworkInterfaceIdentity(map[string]interface{}{"id": "vni-id"})
	assert.Nil(t, err)
	assert.Equal(t, &vpcv1.DynamicRouteServerMemberPrototypeVirtualNetworkInterfaceVirtualNetworkInterfaceIdentityVirtualNetworkInterfaceIdentityByID{ID: core.StringPtr("vni-id")}, result)

	result, err = vpc.ResourceIBMIsDynamicRouteServerMemberMapToDynamicRouteServerMemberPrototypeVirtualNetworkInterfaceVirtualNetworkInterfaceIdentity(map[string]interface{}{"href": "vni-href"})
	assert.Nil(t, err)
	assert.Equal(t, &vpcv1.DynamicRouteServerMemberPrototypeVirtualNetworkInterfaceVirtualNetworkInterfaceIdentityVirtualNetworkInterfaceIdentityByHref{Href: core.StringPtr("vni-href")}, result)

	result, err = vpc.ResourceIBMIsDynamicRouteServerMemberMapToDynamicRouteServerMemberPrototypeVirtualNetworkInterfaceVirtualNetworkInterfaceIdentity(map[string]interface{}{"crn": "vni-crn"})
	assert.Nil(t, err)
	assert.Equal(t, &vpcv1.DynamicRouteServerMemberPrototypeVirtualNetworkInterfaceVirtualNetworkInterfaceIdentityVirtualNetworkInterfaceIdentityByCRN{CRN: core.StringPtr("vni-crn")}, result)

	_, err = vpc.ResourceIBMIsDynamicRouteServerMemberMapToDynamicRouteServerMemberPrototypeVirtualNetworkInterfaceVirtualNetworkInterfaceIdentity(map[string]interface{}{})
	assert.NotNil(t, err)

	_, err = vpc.ResourceIBMIsDynamicRouteServerMemberMapToDynamicRouteServerMemberPrototypeVirtualNetworkInterfaceVirtualNetworkInterfaceIdentity(map[string]interface{}{"id": "a", "href": "b", "crn": "c"})
	assert.NotNil(t, err)
}

func TestResourceIBMIsDynamicRouteServerMemberMapToDynamicRouteServerMemberPrototypeVirtualNetworkInterfaceVirtualNetworkInterfaceIdentityVirtualNetworkInterfaceIdentityByID(t *testing.T) {
	result, err := vpc.ResourceIBMIsDynamicRouteServerMemberMapToDynamicRouteServerMemberPrototypeVirtualNetworkInterfaceVirtualNetworkInterfaceIdentityVirtualNetworkInterfaceIdentityByID(map[string]interface{}{"id": "vni-id"})
	assert.Nil(t, err)
	assert.Equal(t, &vpcv1.DynamicRouteServerMemberPrototypeVirtualNetworkInterfaceVirtualNetworkInterfaceIdentityVirtualNetworkInterfaceIdentityByID{ID: core.StringPtr("vni-id")}, result)

	_, err = vpc.ResourceIBMIsDynamicRouteServerMemberMapToDynamicRouteServerMemberPrototypeVirtualNetworkInterfaceVirtualNetworkInterfaceIdentityVirtualNetworkInterfaceIdentityByID(map[string]interface{}{})
	assert.NotNil(t, err)
}

func TestResourceIBMIsDynamicRouteServerMemberMapToDynamicRouteServerMemberPrototypeVirtualNetworkInterfaceVirtualNetworkInterfaceIdentityVirtualNetworkInterfaceIdentityByHref(t *testing.T) {
	result, err := vpc.ResourceIBMIsDynamicRouteServerMemberMapToDynamicRouteServerMemberPrototypeVirtualNetworkInterfaceVirtualNetworkInterfaceIdentityVirtualNetworkInterfaceIdentityByHref(map[string]interface{}{"href": "vni-href"})
	assert.Nil(t, err)
	assert.Equal(t, &vpcv1.DynamicRouteServerMemberPrototypeVirtualNetworkInterfaceVirtualNetworkInterfaceIdentityVirtualNetworkInterfaceIdentityByHref{Href: core.StringPtr("vni-href")}, result)

	_, err = vpc.ResourceIBMIsDynamicRouteServerMemberMapToDynamicRouteServerMemberPrototypeVirtualNetworkInterfaceVirtualNetworkInterfaceIdentityVirtualNetworkInterfaceIdentityByHref(map[string]interface{}{})
	assert.NotNil(t, err)
}

func TestResourceIBMIsDynamicRouteServerMemberMapToDynamicRouteServerMemberPrototypeVirtualNetworkInterfaceVirtualNetworkInterfaceIdentityVirtualNetworkInterfaceIdentityByCRN(t *testing.T) {
	result, err := vpc.ResourceIBMIsDynamicRouteServerMemberMapToDynamicRouteServerMemberPrototypeVirtualNetworkInterfaceVirtualNetworkInterfaceIdentityVirtualNetworkInterfaceIdentityByCRN(map[string]interface{}{"crn": "vni-crn"})
	assert.Nil(t, err)
	assert.Equal(t, &vpcv1.DynamicRouteServerMemberPrototypeVirtualNetworkInterfaceVirtualNetworkInterfaceIdentityVirtualNetworkInterfaceIdentityByCRN{CRN: core.StringPtr("vni-crn")}, result)

	_, err = vpc.ResourceIBMIsDynamicRouteServerMemberMapToDynamicRouteServerMemberPrototypeVirtualNetworkInterfaceVirtualNetworkInterfaceIdentityVirtualNetworkInterfaceIdentityByCRN(map[string]interface{}{})
	assert.NotNil(t, err)
}

func TestResourceIBMIsDynamicRouteServerMemberMergeVirtualNetworkInterfaces(t *testing.T) {
	t.Run("nil prior and nil current", func(t *testing.T) {
		assert.Equal(t, []map[string]interface{}{}, vpc.ResourceIBMIsDynamicRouteServerMemberMergeVirtualNetworkInterfaces(nil, nil))
	})

	t.Run("nil prior applies defaults", func(t *testing.T) {
		current := []map[string]interface{}{{"id": "vni-1", "primary_ip": []map[string]interface{}{{"id": "rip-1", "address": "10.0.0.5"}}}}
		result := vpc.ResourceIBMIsDynamicRouteServerMemberMergeVirtualNetworkInterfaces(nil, current)
		assert.Len(t, result, 1)
		assert.Equal(t, "vni-1", result[0]["id"])
		assert.Equal(t, false, result[0]["allow_ip_spoofing"])
		assert.Equal(t, true, result[0]["auto_delete"])
		assert.Equal(t, true, result[0]["enable_infrastructure_nat"])
		_, hasMode := result[0]["protocol_state_filtering_mode"]
		assert.False(t, hasMode)
		primaryIP := result[0]["primary_ip"].([]map[string]interface{})
		assert.Equal(t, true, primaryIP[0]["auto_delete"])
		assert.Equal(t, "10.0.0.5", primaryIP[0]["address"])
		// the input is not mutated
		_, mutated := current[0]["auto_delete"]
		assert.False(t, mutated)
	})

	t.Run("prior matched by id carries prototype fields", func(t *testing.T) {
		prior := []interface{}{
			map[string]interface{}{
				"id": "vni-2", "auto_delete": false, "allow_ip_spoofing": true, "enable_infrastructure_nat": false,
				"protocol_state_filtering_mode": "disabled",
				"ips":                           []interface{}{map[string]interface{}{"id": "rip-2"}},
				"security_groups":               []interface{}{map[string]interface{}{"id": "sg-1"}},
				"resource_group":                []interface{}{map[string]interface{}{"id": "rg-1"}},
				"primary_ip":                    []interface{}{map[string]interface{}{"auto_delete": false}},
			},
			map[string]interface{}{"id": "vni-1"},
		}
		current := []map[string]interface{}{
			{"id": "vni-1"},
			{"id": "vni-2", "primary_ip": []map[string]interface{}{{"id": "rip-9"}}},
		}
		result := vpc.ResourceIBMIsDynamicRouteServerMemberMergeVirtualNetworkInterfaces(prior, current)
		assert.Equal(t, true, result[0]["auto_delete"])
		assert.Equal(t, false, result[1]["auto_delete"])
		assert.Equal(t, true, result[1]["allow_ip_spoofing"])
		assert.Equal(t, false, result[1]["enable_infrastructure_nat"])
		assert.Equal(t, "disabled", result[1]["protocol_state_filtering_mode"])
		assert.Equal(t, []interface{}{map[string]interface{}{"id": "rip-2"}}, result[1]["ips"])
		assert.Equal(t, []interface{}{map[string]interface{}{"id": "sg-1"}}, result[1]["security_groups"])
		assert.Equal(t, []interface{}{map[string]interface{}{"id": "rg-1"}}, result[1]["resource_group"])
		primaryIP := result[1]["primary_ip"].([]map[string]interface{})
		assert.Equal(t, false, primaryIP[0]["auto_delete"])
		assert.Equal(t, "rip-9", primaryIP[0]["id"])
	})

	t.Run("prior inline prototype matched by name then by position", func(t *testing.T) {
		prior := []interface{}{
			map[string]interface{}{"id": "", "name": "inline-vni", "auto_delete": false},
			map[string]interface{}{"id": "", "name": "", "allow_ip_spoofing": true},
		}
		current := []map[string]interface{}{
			{"id": "vni-new-2", "name": "generated-name"},
			{"id": "vni-new-1", "name": "inline-vni"},
		}
		result := vpc.ResourceIBMIsDynamicRouteServerMemberMergeVirtualNetworkInterfaces(prior, current)
		assert.Equal(t, false, result[1]["auto_delete"])
		assert.Equal(t, false, result[0]["allow_ip_spoofing"])
		assert.Equal(t, true, result[0]["auto_delete"])
	})
}
