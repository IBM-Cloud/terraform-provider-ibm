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
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/vpc"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/stretchr/testify/assert"
)

func TestAccIBMIsDynamicRouteServerBasic(t *testing.T) {
	var conf vpcv1.DynamicRouteServer
	vpcName := fmt.Sprintf("tf-drs-vpc-%d", acctest.RandIntRange(10, 100))
	subnetName := fmt.Sprintf("tf-drs-subnet-%d", acctest.RandIntRange(10, 100))
	vniName := fmt.Sprintf("tf-drs-vni-%d", acctest.RandIntRange(10, 100))
	memberName := fmt.Sprintf("tf-drs-member-%d", acctest.RandIntRange(10, 100))
	localAsn := fmt.Sprintf("%d", acctest.RandIntRange(64512, 65000))
	localAsnUpdate := fmt.Sprintf("%d", acctest.RandIntRange(65001, 65534))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIsDynamicRouteServerDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsDynamicRouteServerConfigBasic(vpcName, subnetName, vniName, memberName, localAsn),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIsDynamicRouteServerExists("ibm_is_dynamic_route_server.is_dynamic_route_server_instance", conf),
					resource.TestCheckResourceAttr("ibm_is_dynamic_route_server.is_dynamic_route_server_instance", "local_asn", localAsn),
					resource.TestCheckResourceAttr("ibm_is_dynamic_route_server.is_dynamic_route_server_instance", "members.#", "1"),
					resource.TestCheckResourceAttr("ibm_is_dynamic_route_server.is_dynamic_route_server_instance", "members.0.name", memberName),
					resource.TestCheckResourceAttrPair("ibm_is_dynamic_route_server.is_dynamic_route_server_instance", "members.0.virtual_network_interfaces.0.id", "ibm_is_virtual_network_interface.testacc_vni", "id"),
					resource.TestCheckResourceAttrSet("ibm_is_dynamic_route_server.is_dynamic_route_server_instance", "members.0.id"),
					resource.TestCheckResourceAttrSet("ibm_is_dynamic_route_server.is_dynamic_route_server_instance", "crn"),
					resource.TestCheckResourceAttrSet("ibm_is_dynamic_route_server.is_dynamic_route_server_instance", "href"),
					resource.TestCheckResourceAttrSet("ibm_is_dynamic_route_server.is_dynamic_route_server_instance", "created_at"),
					resource.TestCheckResourceAttr("ibm_is_dynamic_route_server.is_dynamic_route_server_instance", "lifecycle_state", "stable"),
					resource.TestCheckResourceAttrSet("ibm_is_dynamic_route_server.is_dynamic_route_server_instance", "etag"),
					resource.TestCheckResourceAttrPair("ibm_is_dynamic_route_server.is_dynamic_route_server_instance", "vpc.0.id", "ibm_is_vpc.testacc_vpc", "id"),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMIsDynamicRouteServerConfigBasic(vpcName, subnetName, vniName, memberName, localAsnUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_is_dynamic_route_server.is_dynamic_route_server_instance", "local_asn", localAsnUpdate),
					resource.TestCheckResourceAttr("ibm_is_dynamic_route_server.is_dynamic_route_server_instance", "lifecycle_state", "stable"),
				),
			},
			resource.TestStep{
				ResourceName:            "ibm_is_dynamic_route_server.is_dynamic_route_server_instance",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"etag"},
			},
		},
	})
}

func TestAccIBMIsDynamicRouteServerAllArgs(t *testing.T) {
	var conf vpcv1.DynamicRouteServer
	vpcName := fmt.Sprintf("tf-drs-vpc-%d", acctest.RandIntRange(10, 100))
	subnetName := fmt.Sprintf("tf-drs-subnet-%d", acctest.RandIntRange(10, 100))
	vniName := fmt.Sprintf("tf-drs-vni-%d", acctest.RandIntRange(10, 100))
	inlineVniName := fmt.Sprintf("tf-drs-inline-vni-%d", acctest.RandIntRange(10, 100))
	memberName := fmt.Sprintf("tf-drs-member-%d", acctest.RandIntRange(10, 100))
	memberNameReplace := fmt.Sprintf("tf-drs-member-new-%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf-drs-%d", acctest.RandIntRange(10, 100))
	nameUpdate := fmt.Sprintf("tf-drs-upd-%d", acctest.RandIntRange(10, 100))
	localAsn := fmt.Sprintf("%d", acctest.RandIntRange(64512, 65000))
	localAsnUpdate := fmt.Sprintf("%d", acctest.RandIntRange(65001, 65534))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIsDynamicRouteServerDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsDynamicRouteServerConfig(vpcName, subnetName, vniName, inlineVniName, memberName, localAsn, name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIsDynamicRouteServerExists("ibm_is_dynamic_route_server.is_dynamic_route_server_instance", conf),
					resource.TestCheckResourceAttr("ibm_is_dynamic_route_server.is_dynamic_route_server_instance", "local_asn", localAsn),
					resource.TestCheckResourceAttr("ibm_is_dynamic_route_server.is_dynamic_route_server_instance", "name", name),
					resource.TestCheckResourceAttr("ibm_is_dynamic_route_server.is_dynamic_route_server_instance", "members.#", "1"),
					resource.TestCheckResourceAttr("ibm_is_dynamic_route_server.is_dynamic_route_server_instance", "members.0.name", memberName),
					resource.TestCheckResourceAttr("ibm_is_dynamic_route_server.is_dynamic_route_server_instance", "members.0.virtual_network_interfaces.#", "2"),
					resource.TestCheckResourceAttrPair("ibm_is_dynamic_route_server.is_dynamic_route_server_instance", "members.0.virtual_network_interfaces.0.id", "ibm_is_virtual_network_interface.testacc_vni", "id"),
					resource.TestCheckResourceAttr("ibm_is_dynamic_route_server.is_dynamic_route_server_instance", "members.0.virtual_network_interfaces.1.name", inlineVniName),
					resource.TestCheckResourceAttrPair("ibm_is_dynamic_route_server.is_dynamic_route_server_instance", "members.0.virtual_network_interfaces.1.subnet.0.id", "ibm_is_subnet.testacc_subnet", "id"),
					resource.TestCheckResourceAttrSet("ibm_is_dynamic_route_server.is_dynamic_route_server_instance", "members.0.virtual_network_interfaces.1.id"),
					resource.TestCheckResourceAttrSet("ibm_is_dynamic_route_server.is_dynamic_route_server_instance", "members.0.virtual_network_interfaces.1.primary_ip.0.address"),
					resource.TestCheckResourceAttr("ibm_is_dynamic_route_server.is_dynamic_route_server_instance", "resource_group.0.id", acc.IsResourceGroupID),
					resource.TestCheckResourceAttr("ibm_is_dynamic_route_server.is_dynamic_route_server_instance", "lifecycle_state", "stable"),
				),
			},
			resource.TestStep{
				// local_asn and name are the only patchable fields; this step must update in place.
				Config: testAccCheckIBMIsDynamicRouteServerConfig(vpcName, subnetName, vniName, inlineVniName, memberName, localAsnUpdate, nameUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_is_dynamic_route_server.is_dynamic_route_server_instance", "local_asn", localAsnUpdate),
					resource.TestCheckResourceAttr("ibm_is_dynamic_route_server.is_dynamic_route_server_instance", "name", nameUpdate),
					resource.TestCheckResourceAttr("ibm_is_dynamic_route_server.is_dynamic_route_server_instance", "lifecycle_state", "stable"),
				),
			},
			resource.TestStep{
				// members is ForceNew; changing the member name replaces the dynamic route server.
				Config: testAccCheckIBMIsDynamicRouteServerConfig(vpcName, subnetName, vniName, inlineVniName, memberNameReplace, localAsnUpdate, nameUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIsDynamicRouteServerExists("ibm_is_dynamic_route_server.is_dynamic_route_server_instance", conf),
					resource.TestCheckResourceAttr("ibm_is_dynamic_route_server.is_dynamic_route_server_instance", "members.0.name", memberNameReplace),
				),
			},
			resource.TestStep{
				ResourceName:            "ibm_is_dynamic_route_server.is_dynamic_route_server_instance",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"etag"},
			},
		},
	})
}

// testAccCheckIBMIsDynamicRouteServerDependencies creates the VPC, subnet and virtual network interface
// that a dynamic route server member refers to.
func testAccCheckIBMIsDynamicRouteServerDependencies(vpcName string, subnetName string, vniName string) string {
	return fmt.Sprintf(`
		resource "ibm_is_vpc" "testacc_vpc" {
			name = "%s"
		}

		resource "ibm_is_subnet" "testacc_subnet" {
			name            = "%s"
			vpc             = ibm_is_vpc.testacc_vpc.id
			zone            = "%s"
			ipv4_cidr_block = "%s"
		}

		resource "ibm_is_virtual_network_interface" "testacc_vni" {
			name   = "%s"
			subnet = ibm_is_subnet.testacc_subnet.id
		}
	`, vpcName, subnetName, acc.ISZoneName, acc.ISCIDR, vniName)
}

func testAccCheckIBMIsDynamicRouteServerConfigBasic(vpcName string, subnetName string, vniName string, memberName string, localAsn string) string {
	return testAccCheckIBMIsDynamicRouteServerDependencies(vpcName, subnetName, vniName) + fmt.Sprintf(`
		resource "ibm_is_dynamic_route_server" "is_dynamic_route_server_instance" {
			local_asn = %s
			members {
				name = "%s"
				virtual_network_interfaces {
					id = ibm_is_virtual_network_interface.testacc_vni.id
				}
			}
		}
	`, localAsn, memberName)
}

func testAccCheckIBMIsDynamicRouteServerConfig(vpcName string, subnetName string, vniName string, inlineVniName string, memberName string, localAsn string, name string) string {
	return testAccCheckIBMIsDynamicRouteServerDependencies(vpcName, subnetName, vniName) + fmt.Sprintf(`
		resource "ibm_is_dynamic_route_server" "is_dynamic_route_server_instance" {
			local_asn = %s
			name      = "%s"
			resource_group {
				id = "%s"
			}
			members {
				name = "%s"
				virtual_network_interfaces {
					id = ibm_is_virtual_network_interface.testacc_vni.id
				}
				virtual_network_interfaces {
					name = "%s"
					subnet {
						id = ibm_is_subnet.testacc_subnet.id
					}
				}
			}
		}
	`, localAsn, name, acc.IsResourceGroupID, memberName, inlineVniName)
}

func testAccCheckIBMIsDynamicRouteServerExists(n string, obj vpcv1.DynamicRouteServer) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		vpcClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).VpcV1API()
		if err != nil {
			return err
		}

		getDynamicRouteServerOptions := &vpcv1.GetDynamicRouteServerOptions{}

		getDynamicRouteServerOptions.SetID(rs.Primary.ID)

		dynamicRouteServer, _, err := vpcClient.GetDynamicRouteServer(getDynamicRouteServerOptions)
		if err != nil {
			return err
		}

		obj = *dynamicRouteServer
		return nil
	}
}

func testAccCheckIBMIsDynamicRouteServerDestroy(s *terraform.State) error {
	vpcClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).VpcV1API()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_is_dynamic_route_server" {
			continue
		}

		getDynamicRouteServerOptions := &vpcv1.GetDynamicRouteServerOptions{}

		getDynamicRouteServerOptions.SetID(rs.Primary.ID)

		_, response, err := vpcClient.GetDynamicRouteServer(getDynamicRouteServerOptions)

		if err == nil {
			return fmt.Errorf("DynamicRouteServer still exists: %s", rs.Primary.ID)
		} else if response == nil || response.StatusCode != 404 {
			return fmt.Errorf("Error checking for DynamicRouteServer (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}

func testDynamicRouteServerReservedIPReference() *vpcv1.ReservedIPReference {
	model := new(vpcv1.ReservedIPReference)
	model.Address = core.StringPtr("192.168.3.4")
	model.Href = core.StringPtr("https://us-south.iaas.cloud.ibm.com/v1/subnets/0717-bea6a632-5e13-42a4-b4b8-31dc877abfe4/reserved_ips/0717-6d353a0f-aeb1-4ae1-832e-1110d10981bb")
	model.ID = core.StringPtr("0717-6d353a0f-aeb1-4ae1-832e-1110d10981bb")
	model.Name = core.StringPtr("my-reserved-ip")
	model.ResourceType = core.StringPtr("subnet_reserved_ip")
	return model
}

func testDynamicRouteServerSubnetReference() *vpcv1.SubnetReference {
	model := new(vpcv1.SubnetReference)
	model.CRN = core.StringPtr("crn:v1:bluemix:public:is:us-south-1:a/aa2432b1fa4d4ace891e9b80fc104e34::subnet:0717-7ec86020-1c6e-4889-b3f0-a15f2e50f87e")
	model.Href = core.StringPtr("https://us-south.iaas.cloud.ibm.com/v1/subnets/0717-7ec86020-1c6e-4889-b3f0-a15f2e50f87e")
	model.ID = core.StringPtr("0717-7ec86020-1c6e-4889-b3f0-a15f2e50f87e")
	model.Name = core.StringPtr("my-subnet")
	model.ResourceType = core.StringPtr("subnet")
	return model
}

func testDynamicRouteServerVirtualNetworkInterfaceReference() *vpcv1.VirtualNetworkInterfaceReference {
	model := new(vpcv1.VirtualNetworkInterfaceReference)
	model.CRN = core.StringPtr("crn:v1:bluemix:public:is:us-south-1:a/aa2432b1fa4d4ace891e9b80fc104e34::virtual-network-interface:0717-54eb57ee-86f2-4796-90bb-d7874e0831ef")
	model.Href = core.StringPtr("https://us-south.iaas.cloud.ibm.com/v1/virtual_network_interfaces/0717-54eb57ee-86f2-4796-90bb-d7874e0831ef")
	model.ID = core.StringPtr("0717-54eb57ee-86f2-4796-90bb-d7874e0831ef")
	model.Name = core.StringPtr("my-virtual-network-interface")
	model.PrimaryIP = testDynamicRouteServerReservedIPReference()
	model.ResourceType = core.StringPtr("virtual_network_interface")
	model.Subnet = testDynamicRouteServerSubnetReference()
	return model
}

func testDynamicRouteServerReservedIPReferenceMap() map[string]interface{} {
	m := make(map[string]interface{})
	m["address"] = "192.168.3.4"
	m["href"] = "https://us-south.iaas.cloud.ibm.com/v1/subnets/0717-bea6a632-5e13-42a4-b4b8-31dc877abfe4/reserved_ips/0717-6d353a0f-aeb1-4ae1-832e-1110d10981bb"
	m["id"] = "0717-6d353a0f-aeb1-4ae1-832e-1110d10981bb"
	m["name"] = "my-reserved-ip"
	m["resource_type"] = "subnet_reserved_ip"
	return m
}

func testDynamicRouteServerSubnetReferenceMap() map[string]interface{} {
	m := make(map[string]interface{})
	m["crn"] = "crn:v1:bluemix:public:is:us-south-1:a/aa2432b1fa4d4ace891e9b80fc104e34::subnet:0717-7ec86020-1c6e-4889-b3f0-a15f2e50f87e"
	m["href"] = "https://us-south.iaas.cloud.ibm.com/v1/subnets/0717-7ec86020-1c6e-4889-b3f0-a15f2e50f87e"
	m["id"] = "0717-7ec86020-1c6e-4889-b3f0-a15f2e50f87e"
	m["name"] = "my-subnet"
	m["resource_type"] = "subnet"
	return m
}

func testDynamicRouteServerVirtualNetworkInterfaceReferenceMap() map[string]interface{} {
	m := make(map[string]interface{})
	m["crn"] = "crn:v1:bluemix:public:is:us-south-1:a/aa2432b1fa4d4ace891e9b80fc104e34::virtual-network-interface:0717-54eb57ee-86f2-4796-90bb-d7874e0831ef"
	m["href"] = "https://us-south.iaas.cloud.ibm.com/v1/virtual_network_interfaces/0717-54eb57ee-86f2-4796-90bb-d7874e0831ef"
	m["id"] = "0717-54eb57ee-86f2-4796-90bb-d7874e0831ef"
	m["name"] = "my-virtual-network-interface"
	m["primary_ip"] = []map[string]interface{}{testDynamicRouteServerReservedIPReferenceMap()}
	m["resource_type"] = "virtual_network_interface"
	m["subnet"] = []map[string]interface{}{testDynamicRouteServerSubnetReferenceMap()}
	return m
}

func TestResourceIBMIsDynamicRouteServerDynamicRouteServerMemberReferenceToMap(t *testing.T) {
	t.Run("nil model", func(t *testing.T) {
		result, err := vpc.ResourceIBMIsDynamicRouteServerDynamicRouteServerMemberReferenceToMap(nil)
		assert.Nil(t, err)
		assert.Equal(t, map[string]interface{}{}, result)
	})

	t.Run("empty model", func(t *testing.T) {
		result, err := vpc.ResourceIBMIsDynamicRouteServerDynamicRouteServerMemberReferenceToMap(new(vpcv1.DynamicRouteServerMemberReference))
		assert.Nil(t, err)
		assert.Equal(t, map[string]interface{}{"virtual_network_interfaces": []map[string]interface{}{}}, result)
	})

	t.Run("single virtual network interface with deleted", func(t *testing.T) {
		model := new(vpcv1.DynamicRouteServerMemberReference)
		model.Deleted = &vpcv1.Deleted{MoreInfo: core.StringPtr("https://cloud.ibm.com/apidocs/vpc#deleted-resources")}
		model.Href = core.StringPtr("https://us-south.iaas.cloud.ibm.com/v1/dynamic_route_servers/r006-d7cc5196-9864-48c4-82d8-3f30da41fcc5/members/r006-d7cc5196-9924-48d4-82d8-3f30da41fdc")
		model.ID = core.StringPtr("0717-9dce0ab3-a719-4582-ad71-00abfaae43ed")
		model.Name = core.StringPtr("my-dynamic-route-server-member1")
		model.ResourceType = core.StringPtr("dynamic_route_server_member")
		model.VirtualNetworkInterfaces = []vpcv1.VirtualNetworkInterfaceReference{*testDynamicRouteServerVirtualNetworkInterfaceReference()}

		expected := make(map[string]interface{})
		expected["deleted"] = []map[string]interface{}{{"more_info": "https://cloud.ibm.com/apidocs/vpc#deleted-resources"}}
		expected["href"] = "https://us-south.iaas.cloud.ibm.com/v1/dynamic_route_servers/r006-d7cc5196-9864-48c4-82d8-3f30da41fcc5/members/r006-d7cc5196-9924-48d4-82d8-3f30da41fdc"
		expected["id"] = "0717-9dce0ab3-a719-4582-ad71-00abfaae43ed"
		expected["name"] = "my-dynamic-route-server-member1"
		expected["resource_type"] = "dynamic_route_server_member"
		expected["virtual_network_interfaces"] = []map[string]interface{}{testDynamicRouteServerVirtualNetworkInterfaceReferenceMap()}

		result, err := vpc.ResourceIBMIsDynamicRouteServerDynamicRouteServerMemberReferenceToMap(model)
		assert.Nil(t, err)
		assert.Equal(t, expected, result)
	})

	t.Run("multiple virtual network interfaces", func(t *testing.T) {
		model := new(vpcv1.DynamicRouteServerMemberReference)
		model.ID = core.StringPtr("member-id")
		model.VirtualNetworkInterfaces = []vpcv1.VirtualNetworkInterfaceReference{*testDynamicRouteServerVirtualNetworkInterfaceReference(), *testDynamicRouteServerVirtualNetworkInterfaceReference()}

		result, err := vpc.ResourceIBMIsDynamicRouteServerDynamicRouteServerMemberReferenceToMap(model)
		assert.Nil(t, err)
		assert.Equal(t, "member-id", result["id"])
		assert.Len(t, result["virtual_network_interfaces"], 2)
		_, hasDeleted := result["deleted"]
		assert.False(t, hasDeleted)
	})
}

func TestResourceIBMIsDynamicRouteServerDeletedToMap(t *testing.T) {
	result, err := vpc.ResourceIBMIsDynamicRouteServerDeletedToMap(nil)
	assert.Nil(t, err)
	assert.Equal(t, map[string]interface{}{}, result)

	result, err = vpc.ResourceIBMIsDynamicRouteServerDeletedToMap(new(vpcv1.Deleted))
	assert.Nil(t, err)
	assert.Equal(t, map[string]interface{}{}, result)

	model := new(vpcv1.Deleted)
	model.MoreInfo = core.StringPtr("https://cloud.ibm.com/apidocs/vpc#deleted-resources")
	result, err = vpc.ResourceIBMIsDynamicRouteServerDeletedToMap(model)
	assert.Nil(t, err)
	assert.Equal(t, map[string]interface{}{"more_info": "https://cloud.ibm.com/apidocs/vpc#deleted-resources"}, result)
}

func TestResourceIBMIsDynamicRouteServerVirtualNetworkInterfaceReferenceToMap(t *testing.T) {
	result, err := vpc.ResourceIBMIsDynamicRouteServerVirtualNetworkInterfaceReferenceToMap(nil)
	assert.Nil(t, err)
	assert.Equal(t, map[string]interface{}{}, result)

	result, err = vpc.ResourceIBMIsDynamicRouteServerVirtualNetworkInterfaceReferenceToMap(new(vpcv1.VirtualNetworkInterfaceReference))
	assert.Nil(t, err)
	assert.Equal(t, map[string]interface{}{}, result)

	model := testDynamicRouteServerVirtualNetworkInterfaceReference()
	model.Deleted = &vpcv1.Deleted{MoreInfo: core.StringPtr("https://cloud.ibm.com/apidocs/vpc#deleted-resources")}
	expected := testDynamicRouteServerVirtualNetworkInterfaceReferenceMap()
	expected["deleted"] = []map[string]interface{}{{"more_info": "https://cloud.ibm.com/apidocs/vpc#deleted-resources"}}
	result, err = vpc.ResourceIBMIsDynamicRouteServerVirtualNetworkInterfaceReferenceToMap(model)
	assert.Nil(t, err)
	assert.Equal(t, expected, result)
}

func TestResourceIBMIsDynamicRouteServerReservedIPReferenceToMap(t *testing.T) {
	result, err := vpc.ResourceIBMIsDynamicRouteServerReservedIPReferenceToMap(nil)
	assert.Nil(t, err)
	assert.Equal(t, map[string]interface{}{}, result)

	result, err = vpc.ResourceIBMIsDynamicRouteServerReservedIPReferenceToMap(testDynamicRouteServerReservedIPReference())
	assert.Nil(t, err)
	assert.Equal(t, testDynamicRouteServerReservedIPReferenceMap(), result)
}

func TestResourceIBMIsDynamicRouteServerSubnetReferenceToMap(t *testing.T) {
	result, err := vpc.ResourceIBMIsDynamicRouteServerSubnetReferenceToMap(nil)
	assert.Nil(t, err)
	assert.Equal(t, map[string]interface{}{}, result)

	result, err = vpc.ResourceIBMIsDynamicRouteServerSubnetReferenceToMap(testDynamicRouteServerSubnetReference())
	assert.Nil(t, err)
	assert.Equal(t, testDynamicRouteServerSubnetReferenceMap(), result)
}

func TestResourceIBMIsDynamicRouteServerResourceGroupReferenceToMap(t *testing.T) {
	result, err := vpc.ResourceIBMIsDynamicRouteServerResourceGroupReferenceToMap(nil)
	assert.Nil(t, err)
	assert.Equal(t, map[string]interface{}{}, result)

	model := new(vpcv1.ResourceGroupReference)
	model.Href = core.StringPtr("https://resource-controller.cloud.ibm.com/v2/resource_groups/fee82deba12e4c0fb69c3b09d1f12345")
	model.ID = core.StringPtr("fee82deba12e4c0fb69c3b09d1f12345")
	model.Name = core.StringPtr("my-resource-group")
	result, err = vpc.ResourceIBMIsDynamicRouteServerResourceGroupReferenceToMap(model)
	assert.Nil(t, err)
	assert.Equal(t, map[string]interface{}{
		"href": "https://resource-controller.cloud.ibm.com/v2/resource_groups/fee82deba12e4c0fb69c3b09d1f12345",
		"id":   "fee82deba12e4c0fb69c3b09d1f12345",
		"name": "my-resource-group",
	}, result)
}

func TestResourceIBMIsDynamicRouteServerDynamicRouteServerHealthReasonToMap(t *testing.T) {
	result, err := vpc.ResourceIBMIsDynamicRouteServerDynamicRouteServerHealthReasonToMap(nil)
	assert.Nil(t, err)
	assert.Equal(t, map[string]interface{}{}, result)

	model := new(vpcv1.DynamicRouteServerHealthReason)
	model.Code = core.StringPtr("cannot_reserve_ip_address")
	model.Message = core.StringPtr("The IP address could not be reserved")
	result, err = vpc.ResourceIBMIsDynamicRouteServerDynamicRouteServerHealthReasonToMap(model)
	assert.Nil(t, err)
	assert.Equal(t, map[string]interface{}{"code": "cannot_reserve_ip_address", "message": "The IP address could not be reserved"}, result)

	model.MoreInfo = core.StringPtr("https://cloud.ibm.com/docs/vpc?topic=vpc-drs")
	result, err = vpc.ResourceIBMIsDynamicRouteServerDynamicRouteServerHealthReasonToMap(model)
	assert.Nil(t, err)
	assert.Equal(t, "https://cloud.ibm.com/docs/vpc?topic=vpc-drs", result["more_info"])
}

func TestResourceIBMIsDynamicRouteServerDynamicRouteServerLifecycleReasonToMap(t *testing.T) {
	result, err := vpc.ResourceIBMIsDynamicRouteServerDynamicRouteServerLifecycleReasonToMap(nil)
	assert.Nil(t, err)
	assert.Equal(t, map[string]interface{}{}, result)

	model := new(vpcv1.DynamicRouteServerLifecycleReason)
	model.Code = core.StringPtr("resource_suspended_by_provider")
	model.Message = core.StringPtr("The resource has been suspended. Contact IBM support with the CRN for next steps.")
	model.MoreInfo = core.StringPtr("https://cloud.ibm.com/apidocs/vpc#resource-suspension")
	result, err = vpc.ResourceIBMIsDynamicRouteServerDynamicRouteServerLifecycleReasonToMap(model)
	assert.Nil(t, err)
	assert.Equal(t, map[string]interface{}{
		"code":      "resource_suspended_by_provider",
		"message":   "The resource has been suspended. Contact IBM support with the CRN for next steps.",
		"more_info": "https://cloud.ibm.com/apidocs/vpc#resource-suspension",
	}, result)
}

func TestResourceIBMIsDynamicRouteServerVPCReferenceToMap(t *testing.T) {
	result, err := vpc.ResourceIBMIsDynamicRouteServerVPCReferenceToMap(nil)
	assert.Nil(t, err)
	assert.Equal(t, map[string]interface{}{}, result)

	model := new(vpcv1.VPCReference)
	model.CRN = core.StringPtr("crn:v1:bluemix:public:is:us-south:a/aa2432b1fa4d4ace891e9b80fc104e34::vpc:r006-4727d842-f94f-4a2d-824a-9bc9b02c523b")
	model.Deleted = &vpcv1.Deleted{MoreInfo: core.StringPtr("https://cloud.ibm.com/apidocs/vpc#deleted-resources")}
	model.Href = core.StringPtr("https://us-south.iaas.cloud.ibm.com/v1/vpcs/r006-4727d842-f94f-4a2d-824a-9bc9b02c523b")
	model.ID = core.StringPtr("r006-4727d842-f94f-4a2d-824a-9bc9b02c523b")
	model.Name = core.StringPtr("my-vpc")
	model.ResourceType = core.StringPtr("vpc")
	result, err = vpc.ResourceIBMIsDynamicRouteServerVPCReferenceToMap(model)
	assert.Nil(t, err)
	assert.Equal(t, map[string]interface{}{
		"crn":           "crn:v1:bluemix:public:is:us-south:a/aa2432b1fa4d4ace891e9b80fc104e34::vpc:r006-4727d842-f94f-4a2d-824a-9bc9b02c523b",
		"deleted":       []map[string]interface{}{{"more_info": "https://cloud.ibm.com/apidocs/vpc#deleted-resources"}},
		"href":          "https://us-south.iaas.cloud.ibm.com/v1/vpcs/r006-4727d842-f94f-4a2d-824a-9bc9b02c523b",
		"id":            "r006-4727d842-f94f-4a2d-824a-9bc9b02c523b",
		"name":          "my-vpc",
		"resource_type": "vpc",
	}, result)
}

func TestResourceIBMIsDynamicRouteServerMapToDynamicRouteServerMemberPrototype(t *testing.T) {
	t.Run("nil map", func(t *testing.T) {
		_, err := vpc.ResourceIBMIsDynamicRouteServerMapToDynamicRouteServerMemberPrototype(nil)
		assert.NotNil(t, err)
	})

	t.Run("missing name", func(t *testing.T) {
		model := map[string]interface{}{"virtual_network_interfaces": []interface{}{map[string]interface{}{"id": "vni-id"}}}
		_, err := vpc.ResourceIBMIsDynamicRouteServerMapToDynamicRouteServerMemberPrototype(model)
		assert.NotNil(t, err)
	})

	t.Run("empty virtual network interfaces", func(t *testing.T) {
		model := map[string]interface{}{"name": "member", "virtual_network_interfaces": []interface{}{}}
		_, err := vpc.ResourceIBMIsDynamicRouteServerMapToDynamicRouteServerMemberPrototype(model)
		assert.NotNil(t, err)

		model = map[string]interface{}{"name": "member"}
		_, err = vpc.ResourceIBMIsDynamicRouteServerMapToDynamicRouteServerMemberPrototype(model)
		assert.NotNil(t, err)
	})

	t.Run("invalid virtual network interface propagates error", func(t *testing.T) {
		model := map[string]interface{}{"name": "member", "virtual_network_interfaces": []interface{}{map[string]interface{}{"id": "a", "crn": "b"}}}
		_, err := vpc.ResourceIBMIsDynamicRouteServerMapToDynamicRouteServerMemberPrototype(model)
		assert.NotNil(t, err)
	})

	t.Run("identity and prototype members", func(t *testing.T) {
		model := make(map[string]interface{})
		model["name"] = "my-dynamic-route-server-member1"
		model["virtual_network_interfaces"] = []interface{}{
			map[string]interface{}{"id": "0717-54eb57ee-86f2-4796-90bb-d7874e0831ef"},
			map[string]interface{}{
				"name":   "my-virtual-network-interface",
				"subnet": []interface{}{map[string]interface{}{"id": "0717-7ec86020-1c6e-4889-b3f0-a15f2e50f87e"}},
			},
		}

		result, err := vpc.ResourceIBMIsDynamicRouteServerMapToDynamicRouteServerMemberPrototype(model)
		assert.Nil(t, err)

		expected := new(vpcv1.DynamicRouteServerMemberPrototype)
		expected.Name = core.StringPtr("my-dynamic-route-server-member1")
		expected.VirtualNetworkInterfaces = []vpcv1.DynamicRouteServerMemberPrototypeVirtualNetworkInterfaceIntf{
			&vpcv1.DynamicRouteServerMemberPrototypeVirtualNetworkInterfaceVirtualNetworkInterfaceIdentityVirtualNetworkInterfaceIdentityByID{ID: core.StringPtr("0717-54eb57ee-86f2-4796-90bb-d7874e0831ef")},
			&vpcv1.DynamicRouteServerMemberPrototypeVirtualNetworkInterfaceVirtualNetworkInterfacePrototypeDynamicRouteServerContext{
				Name:   core.StringPtr("my-virtual-network-interface"),
				Subnet: &vpcv1.SubnetIdentityByID{ID: core.StringPtr("0717-7ec86020-1c6e-4889-b3f0-a15f2e50f87e")},
			},
		}
		assert.Equal(t, expected, result)
	})
}

func TestResourceIBMIsDynamicRouteServerMapToResourceGroupIdentity(t *testing.T) {
	_, err := vpc.ResourceIBMIsDynamicRouteServerMapToResourceGroupIdentity(map[string]interface{}{})
	assert.NotNil(t, err)

	_, err = vpc.ResourceIBMIsDynamicRouteServerMapToResourceGroupIdentity(map[string]interface{}{"id": ""})
	assert.NotNil(t, err)

	result, err := vpc.ResourceIBMIsDynamicRouteServerMapToResourceGroupIdentity(map[string]interface{}{"id": "fee82deba12e4c0fb69c3b09d1f12345"})
	assert.Nil(t, err)
	assert.Equal(t, &vpcv1.ResourceGroupIdentityByID{ID: core.StringPtr("fee82deba12e4c0fb69c3b09d1f12345")}, result)
}

func TestResourceIBMIsDynamicRouteServerMapToResourceGroupIdentityByID(t *testing.T) {
	_, err := vpc.ResourceIBMIsDynamicRouteServerMapToResourceGroupIdentityByID(map[string]interface{}{})
	assert.NotNil(t, err)

	result, err := vpc.ResourceIBMIsDynamicRouteServerMapToResourceGroupIdentityByID(map[string]interface{}{"id": "fee82deba12e4c0fb69c3b09d1f12345"})
	assert.Nil(t, err)
	assert.Equal(t, &vpcv1.ResourceGroupIdentityByID{ID: core.StringPtr("fee82deba12e4c0fb69c3b09d1f12345")}, result)
}

func TestResourceIBMIsDynamicRouteServerMergeMembers(t *testing.T) {
	current := []map[string]interface{}{
		{
			"name": "member-a",
			"id":   "member-a-id",
			"virtual_network_interfaces": []map[string]interface{}{
				{"id": "vni-1", "name": "vni-one"},
			},
		},
		{
			"name": "member-b",
			"id":   "member-b-id",
			"virtual_network_interfaces": []map[string]interface{}{
				{"id": "vni-2", "name": "vni-two"},
			},
		},
	}

	t.Run("nil prior uses defaults", func(t *testing.T) {
		result := vpc.ResourceIBMIsDynamicRouteServerMergeMembers(nil, current)
		assert.Len(t, result, 2)
		vnis := result[0]["virtual_network_interfaces"].([]map[string]interface{})
		assert.Equal(t, false, vnis[0]["allow_ip_spoofing"])
		assert.Equal(t, true, vnis[0]["auto_delete"])
		assert.Equal(t, true, vnis[0]["enable_infrastructure_nat"])
	})

	t.Run("empty current", func(t *testing.T) {
		result := vpc.ResourceIBMIsDynamicRouteServerMergeMembers([]interface{}{map[string]interface{}{"name": "x"}}, []map[string]interface{}{})
		assert.Equal(t, []map[string]interface{}{}, result)
	})

	t.Run("prior matched by name regardless of order", func(t *testing.T) {
		prior := []interface{}{
			map[string]interface{}{
				"name": "member-b",
				"virtual_network_interfaces": []interface{}{
					map[string]interface{}{"id": "vni-2", "auto_delete": false, "allow_ip_spoofing": true},
				},
			},
			map[string]interface{}{
				"name": "member-a",
				"virtual_network_interfaces": []interface{}{
					map[string]interface{}{"name": "vni-one", "auto_delete": true, "enable_infrastructure_nat": false, "protocol_state_filtering_mode": "enabled"},
				},
			},
		}
		result := vpc.ResourceIBMIsDynamicRouteServerMergeMembers(prior, current)
		vnisA := result[0]["virtual_network_interfaces"].([]map[string]interface{})
		assert.Equal(t, false, vnisA[0]["enable_infrastructure_nat"])
		assert.Equal(t, "enabled", vnisA[0]["protocol_state_filtering_mode"])
		vnisB := result[1]["virtual_network_interfaces"].([]map[string]interface{})
		assert.Equal(t, false, vnisB[0]["auto_delete"])
		assert.Equal(t, true, vnisB[0]["allow_ip_spoofing"])
		// the API values are untouched
		assert.Equal(t, "member-a-id", result[0]["id"])
	})
}

func TestResourceIBMIsDynamicRouteServerDynamicRouteServerPatchAsPatch(t *testing.T) {
	assert.Equal(t, map[string]interface{}{}, vpc.ResourceIBMIsDynamicRouteServerDynamicRouteServerPatchAsPatch(nil, nil))
	assert.Equal(t, map[string]interface{}{}, vpc.ResourceIBMIsDynamicRouteServerDynamicRouteServerPatchAsPatch(&vpcv1.DynamicRouteServerPatch{}, nil))

	patch := vpc.ResourceIBMIsDynamicRouteServerDynamicRouteServerPatchAsPatch(&vpcv1.DynamicRouteServerPatch{LocalAsn: core.Int64Ptr(64512)}, nil)
	assert.Equal(t, map[string]interface{}{"local_asn": int64(64512)}, patch)

	patch = vpc.ResourceIBMIsDynamicRouteServerDynamicRouteServerPatchAsPatch(&vpcv1.DynamicRouteServerPatch{Name: core.StringPtr("new-name")}, nil)
	assert.Equal(t, map[string]interface{}{"name": "new-name"}, patch)

	// an empty name is never sent, the service keeps the generated name
	patch = vpc.ResourceIBMIsDynamicRouteServerDynamicRouteServerPatchAsPatch(&vpcv1.DynamicRouteServerPatch{Name: core.StringPtr("")}, nil)
	assert.Equal(t, map[string]interface{}{}, patch)
}
