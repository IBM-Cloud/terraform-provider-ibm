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
	localAsn := fmt.Sprintf("%d", acctest.RandIntRange(10, 100))
	localAsnUpdate := fmt.Sprintf("%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIsDynamicRouteServerDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsDynamicRouteServerConfigBasic(localAsn),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIsDynamicRouteServerExists("ibm_is_dynamic_route_server.is_dynamic_route_server_instance", conf),
					resource.TestCheckResourceAttr("ibm_is_dynamic_route_server.is_dynamic_route_server_instance", "local_asn", localAsn),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMIsDynamicRouteServerConfigBasic(localAsnUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_is_dynamic_route_server.is_dynamic_route_server_instance", "local_asn", localAsnUpdate),
				),
			},
		},
	})
}

func TestAccIBMIsDynamicRouteServerAllArgs(t *testing.T) {
	var conf vpcv1.DynamicRouteServer
	localAsn := fmt.Sprintf("%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	localAsnUpdate := fmt.Sprintf("%d", acctest.RandIntRange(10, 100))
	nameUpdate := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIsDynamicRouteServerDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsDynamicRouteServerConfig(localAsn, name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIsDynamicRouteServerExists("ibm_is_dynamic_route_server.is_dynamic_route_server_instance", conf),
					resource.TestCheckResourceAttr("ibm_is_dynamic_route_server.is_dynamic_route_server_instance", "local_asn", localAsn),
					resource.TestCheckResourceAttr("ibm_is_dynamic_route_server.is_dynamic_route_server_instance", "name", name),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMIsDynamicRouteServerConfig(localAsnUpdate, nameUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_is_dynamic_route_server.is_dynamic_route_server_instance", "local_asn", localAsnUpdate),
					resource.TestCheckResourceAttr("ibm_is_dynamic_route_server.is_dynamic_route_server_instance", "name", nameUpdate),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_is_dynamic_route_server.is_dynamic_route_server_instance",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIBMIsDynamicRouteServerConfigBasic(localAsn string) string {
	return fmt.Sprintf(`
		resource "ibm_is_dynamic_route_server" "is_dynamic_route_server_instance" {
			local_asn = %s
			members {
				deleted {
					more_info = "https://cloud.ibm.com/apidocs/vpc#deleted-resources"
				}
				href = "https://us-south.iaas.cloud.ibm.com/v1/dynamic_route_servers/r006-d7cc5196-9864-48c4-82d8-3f30da41fcc5/members/r006-d7cc5196-9924-48d4-82d8-3f30da41fdc"
				id = "0717-9dce0ab3-a719-4582-ad71-00abfaae43ed"
				name = "my-dynamic-route-server-member1"
				resource_type = "dynamic_route_server_member"
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
		}
	`, localAsn)
}

func testAccCheckIBMIsDynamicRouteServerConfig(localAsn string, name string) string {
	return fmt.Sprintf(`

		resource "ibm_is_dynamic_route_server" "is_dynamic_route_server_instance" {
			local_asn = %s
			members {
				deleted {
					more_info = "https://cloud.ibm.com/apidocs/vpc#deleted-resources"
				}
				href = "https://us-south.iaas.cloud.ibm.com/v1/dynamic_route_servers/r006-d7cc5196-9864-48c4-82d8-3f30da41fcc5/members/r006-d7cc5196-9924-48d4-82d8-3f30da41fdc"
				id = "0717-9dce0ab3-a719-4582-ad71-00abfaae43ed"
				name = "my-dynamic-route-server-member1"
				resource_type = "dynamic_route_server_member"
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
			name = "%s"
			resource_group {
				id = "fee82deba12e4c0fb69c3b09d1f12345"
			}
		}
	`, localAsn, name)
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

		// Try to find the key
		_, response, err := vpcClient.GetDynamicRouteServer(getDynamicRouteServerOptions)

		if err == nil {
			return fmt.Errorf("DynamicRouteServer still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for DynamicRouteServer (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}

func TestResourceIBMIsDynamicRouteServerDynamicRouteServerMemberReferenceToMap(t *testing.T) {
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

	result, err := vpc.ResourceIBMIsDynamicRouteServerDynamicRouteServerMemberReferenceToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMIsDynamicRouteServerDeletedToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["more_info"] = "https://cloud.ibm.com/apidocs/vpc#deleted-resources"

		assert.Equal(t, result, model)
	}

	model := new(vpcv1.Deleted)
	model.MoreInfo = core.StringPtr("https://cloud.ibm.com/apidocs/vpc#deleted-resources")

	result, err := vpc.ResourceIBMIsDynamicRouteServerDeletedToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMIsDynamicRouteServerVirtualNetworkInterfaceReferenceToMap(t *testing.T) {
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

	result, err := vpc.ResourceIBMIsDynamicRouteServerVirtualNetworkInterfaceReferenceToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMIsDynamicRouteServerReservedIPReferenceToMap(t *testing.T) {
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

	result, err := vpc.ResourceIBMIsDynamicRouteServerReservedIPReferenceToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMIsDynamicRouteServerSubnetReferenceToMap(t *testing.T) {
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

	result, err := vpc.ResourceIBMIsDynamicRouteServerSubnetReferenceToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMIsDynamicRouteServerResourceGroupReferenceToMap(t *testing.T) {
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

	result, err := vpc.ResourceIBMIsDynamicRouteServerResourceGroupReferenceToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMIsDynamicRouteServerDynamicRouteServerHealthReasonToMap(t *testing.T) {
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

	result, err := vpc.ResourceIBMIsDynamicRouteServerDynamicRouteServerHealthReasonToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMIsDynamicRouteServerDynamicRouteServerLifecycleReasonToMap(t *testing.T) {
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

	result, err := vpc.ResourceIBMIsDynamicRouteServerDynamicRouteServerLifecycleReasonToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMIsDynamicRouteServerVPCReferenceToMap(t *testing.T) {
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

	result, err := vpc.ResourceIBMIsDynamicRouteServerVPCReferenceToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMIsDynamicRouteServerMapToDynamicRouteServerMemberPrototype(t *testing.T) {
	checkResult := func(result *vpcv1.DynamicRouteServerMemberPrototype) {
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

		dynamicRouteServerMemberPrototypeVirtualNetworkInterfaceModel := new(vpcv1.DynamicRouteServerMemberPrototypeVirtualNetworkInterfaceVirtualNetworkInterfacePrototypeDynamicRouteServerContext)
		dynamicRouteServerMemberPrototypeVirtualNetworkInterfaceModel.AllowIPSpoofing = core.BoolPtr(true)
		dynamicRouteServerMemberPrototypeVirtualNetworkInterfaceModel.AutoDelete = core.BoolPtr(false)
		dynamicRouteServerMemberPrototypeVirtualNetworkInterfaceModel.EnableInfrastructureNat = core.BoolPtr(true)
		dynamicRouteServerMemberPrototypeVirtualNetworkInterfaceModel.Ips = []vpcv1.VirtualNetworkInterfaceIPPrototypeIntf{virtualNetworkInterfaceIPPrototypeModel}
		dynamicRouteServerMemberPrototypeVirtualNetworkInterfaceModel.Name = core.StringPtr("my-virtual-network-interface")
		dynamicRouteServerMemberPrototypeVirtualNetworkInterfaceModel.PrimaryIP = virtualNetworkInterfacePrimaryIPPrototypeModel
		dynamicRouteServerMemberPrototypeVirtualNetworkInterfaceModel.ProtocolStateFilteringMode = core.StringPtr("auto")
		dynamicRouteServerMemberPrototypeVirtualNetworkInterfaceModel.ResourceGroup = resourceGroupIdentityModel
		dynamicRouteServerMemberPrototypeVirtualNetworkInterfaceModel.SecurityGroups = []vpcv1.SecurityGroupIdentityIntf{securityGroupIdentityModel}
		dynamicRouteServerMemberPrototypeVirtualNetworkInterfaceModel.Subnet = subnetIdentityModel

		model := new(vpcv1.DynamicRouteServerMemberPrototype)
		model.Name = core.StringPtr("my-dynamic-route-server-member1")
		model.VirtualNetworkInterfaces = []vpcv1.DynamicRouteServerMemberPrototypeVirtualNetworkInterfaceIntf{dynamicRouteServerMemberPrototypeVirtualNetworkInterfaceModel}

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

	dynamicRouteServerMemberPrototypeVirtualNetworkInterfaceModel := make(map[string]interface{})
	dynamicRouteServerMemberPrototypeVirtualNetworkInterfaceModel["allow_ip_spoofing"] = true
	dynamicRouteServerMemberPrototypeVirtualNetworkInterfaceModel["auto_delete"] = false
	dynamicRouteServerMemberPrototypeVirtualNetworkInterfaceModel["enable_infrastructure_nat"] = true
	dynamicRouteServerMemberPrototypeVirtualNetworkInterfaceModel["ips"] = []interface{}{virtualNetworkInterfaceIPPrototypeModel}
	dynamicRouteServerMemberPrototypeVirtualNetworkInterfaceModel["name"] = "my-virtual-network-interface"
	dynamicRouteServerMemberPrototypeVirtualNetworkInterfaceModel["primary_ip"] = []interface{}{virtualNetworkInterfacePrimaryIPPrototypeModel}
	dynamicRouteServerMemberPrototypeVirtualNetworkInterfaceModel["protocol_state_filtering_mode"] = "auto"
	dynamicRouteServerMemberPrototypeVirtualNetworkInterfaceModel["resource_group"] = []interface{}{resourceGroupIdentityModel}
	dynamicRouteServerMemberPrototypeVirtualNetworkInterfaceModel["security_groups"] = []interface{}{securityGroupIdentityModel}
	dynamicRouteServerMemberPrototypeVirtualNetworkInterfaceModel["subnet"] = []interface{}{subnetIdentityModel}

	model := make(map[string]interface{})
	model["name"] = "my-dynamic-route-server-member1"
	model["virtual_network_interfaces"] = []interface{}{dynamicRouteServerMemberPrototypeVirtualNetworkInterfaceModel}

	result, err := vpc.ResourceIBMIsDynamicRouteServerMapToDynamicRouteServerMemberPrototype(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMIsDynamicRouteServerMapToDynamicRouteServerMemberPrototypeVirtualNetworkInterface(t *testing.T) {
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

	result, err := vpc.ResourceIBMIsDynamicRouteServerMapToDynamicRouteServerMemberPrototypeVirtualNetworkInterface(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMIsDynamicRouteServerMapToVirtualNetworkInterfaceIPPrototype(t *testing.T) {
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

	result, err := vpc.ResourceIBMIsDynamicRouteServerMapToVirtualNetworkInterfaceIPPrototype(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMIsDynamicRouteServerMapToVirtualNetworkInterfaceIPPrototypeReservedIPIdentityVirtualNetworkInterfaceIPsContext(t *testing.T) {
	checkResult := func(result vpcv1.VirtualNetworkInterfaceIPPrototypeReservedIPIdentityVirtualNetworkInterfaceIPsContextIntf) {
		model := new(vpcv1.VirtualNetworkInterfaceIPPrototypeReservedIPIdentityVirtualNetworkInterfaceIPsContext)
		model.ID = core.StringPtr("0717-6d353a0f-aeb1-4ae1-832e-1110d10981bb")
		model.Href = core.StringPtr("https://us-south.iaas.cloud.ibm.com/v1/subnets/0717-bea6a632-5e13-42a4-b4b8-31dc877abfe4/reserved_ips/0717-6d353a0f-aeb1-4ae1-832e-1110d10981bb")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["id"] = "0717-6d353a0f-aeb1-4ae1-832e-1110d10981bb"
	model["href"] = "https://us-south.iaas.cloud.ibm.com/v1/subnets/0717-bea6a632-5e13-42a4-b4b8-31dc877abfe4/reserved_ips/0717-6d353a0f-aeb1-4ae1-832e-1110d10981bb"

	result, err := vpc.ResourceIBMIsDynamicRouteServerMapToVirtualNetworkInterfaceIPPrototypeReservedIPIdentityVirtualNetworkInterfaceIPsContext(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMIsDynamicRouteServerMapToVirtualNetworkInterfaceIPPrototypeReservedIPIdentityVirtualNetworkInterfaceIPsContextByID(t *testing.T) {
	checkResult := func(result *vpcv1.VirtualNetworkInterfaceIPPrototypeReservedIPIdentityVirtualNetworkInterfaceIPsContextByID) {
		model := new(vpcv1.VirtualNetworkInterfaceIPPrototypeReservedIPIdentityVirtualNetworkInterfaceIPsContextByID)
		model.ID = core.StringPtr("0717-6d353a0f-aeb1-4ae1-832e-1110d10981bb")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["id"] = "0717-6d353a0f-aeb1-4ae1-832e-1110d10981bb"

	result, err := vpc.ResourceIBMIsDynamicRouteServerMapToVirtualNetworkInterfaceIPPrototypeReservedIPIdentityVirtualNetworkInterfaceIPsContextByID(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMIsDynamicRouteServerMapToVirtualNetworkInterfaceIPPrototypeReservedIPIdentityVirtualNetworkInterfaceIPsContextByHref(t *testing.T) {
	checkResult := func(result *vpcv1.VirtualNetworkInterfaceIPPrototypeReservedIPIdentityVirtualNetworkInterfaceIPsContextByHref) {
		model := new(vpcv1.VirtualNetworkInterfaceIPPrototypeReservedIPIdentityVirtualNetworkInterfaceIPsContextByHref)
		model.Href = core.StringPtr("https://us-south.iaas.cloud.ibm.com/v1/subnets/0717-bea6a632-5e13-42a4-b4b8-31dc877abfe4/reserved_ips/0717-6d353a0f-aeb1-4ae1-832e-1110d10981bb")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["href"] = "https://us-south.iaas.cloud.ibm.com/v1/subnets/0717-bea6a632-5e13-42a4-b4b8-31dc877abfe4/reserved_ips/0717-6d353a0f-aeb1-4ae1-832e-1110d10981bb"

	result, err := vpc.ResourceIBMIsDynamicRouteServerMapToVirtualNetworkInterfaceIPPrototypeReservedIPIdentityVirtualNetworkInterfaceIPsContextByHref(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMIsDynamicRouteServerMapToVirtualNetworkInterfaceIPPrototypeReservedIPPrototypeVirtualNetworkInterfaceIPsContext(t *testing.T) {
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

	result, err := vpc.ResourceIBMIsDynamicRouteServerMapToVirtualNetworkInterfaceIPPrototypeReservedIPPrototypeVirtualNetworkInterfaceIPsContext(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMIsDynamicRouteServerMapToVirtualNetworkInterfacePrimaryIPPrototype(t *testing.T) {
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

	result, err := vpc.ResourceIBMIsDynamicRouteServerMapToVirtualNetworkInterfacePrimaryIPPrototype(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMIsDynamicRouteServerMapToVirtualNetworkInterfacePrimaryIPPrototypeReservedIPIdentityVirtualNetworkInterfacePrimaryIPContext(t *testing.T) {
	checkResult := func(result vpcv1.VirtualNetworkInterfacePrimaryIPPrototypeReservedIPIdentityVirtualNetworkInterfacePrimaryIPContextIntf) {
		model := new(vpcv1.VirtualNetworkInterfacePrimaryIPPrototypeReservedIPIdentityVirtualNetworkInterfacePrimaryIPContext)
		model.ID = core.StringPtr("0717-6d353a0f-aeb1-4ae1-832e-1110d10981bb")
		model.Href = core.StringPtr("https://us-south.iaas.cloud.ibm.com/v1/subnets/0717-bea6a632-5e13-42a4-b4b8-31dc877abfe4/reserved_ips/0717-6d353a0f-aeb1-4ae1-832e-1110d10981bb")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["id"] = "0717-6d353a0f-aeb1-4ae1-832e-1110d10981bb"
	model["href"] = "https://us-south.iaas.cloud.ibm.com/v1/subnets/0717-bea6a632-5e13-42a4-b4b8-31dc877abfe4/reserved_ips/0717-6d353a0f-aeb1-4ae1-832e-1110d10981bb"

	result, err := vpc.ResourceIBMIsDynamicRouteServerMapToVirtualNetworkInterfacePrimaryIPPrototypeReservedIPIdentityVirtualNetworkInterfacePrimaryIPContext(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMIsDynamicRouteServerMapToVirtualNetworkInterfacePrimaryIPPrototypeReservedIPIdentityVirtualNetworkInterfacePrimaryIPContextByID(t *testing.T) {
	checkResult := func(result *vpcv1.VirtualNetworkInterfacePrimaryIPPrototypeReservedIPIdentityVirtualNetworkInterfacePrimaryIPContextByID) {
		model := new(vpcv1.VirtualNetworkInterfacePrimaryIPPrototypeReservedIPIdentityVirtualNetworkInterfacePrimaryIPContextByID)
		model.ID = core.StringPtr("0717-6d353a0f-aeb1-4ae1-832e-1110d10981bb")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["id"] = "0717-6d353a0f-aeb1-4ae1-832e-1110d10981bb"

	result, err := vpc.ResourceIBMIsDynamicRouteServerMapToVirtualNetworkInterfacePrimaryIPPrototypeReservedIPIdentityVirtualNetworkInterfacePrimaryIPContextByID(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMIsDynamicRouteServerMapToVirtualNetworkInterfacePrimaryIPPrototypeReservedIPIdentityVirtualNetworkInterfacePrimaryIPContextByHref(t *testing.T) {
	checkResult := func(result *vpcv1.VirtualNetworkInterfacePrimaryIPPrototypeReservedIPIdentityVirtualNetworkInterfacePrimaryIPContextByHref) {
		model := new(vpcv1.VirtualNetworkInterfacePrimaryIPPrototypeReservedIPIdentityVirtualNetworkInterfacePrimaryIPContextByHref)
		model.Href = core.StringPtr("https://us-south.iaas.cloud.ibm.com/v1/subnets/0717-bea6a632-5e13-42a4-b4b8-31dc877abfe4/reserved_ips/0717-6d353a0f-aeb1-4ae1-832e-1110d10981bb")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["href"] = "https://us-south.iaas.cloud.ibm.com/v1/subnets/0717-bea6a632-5e13-42a4-b4b8-31dc877abfe4/reserved_ips/0717-6d353a0f-aeb1-4ae1-832e-1110d10981bb"

	result, err := vpc.ResourceIBMIsDynamicRouteServerMapToVirtualNetworkInterfacePrimaryIPPrototypeReservedIPIdentityVirtualNetworkInterfacePrimaryIPContextByHref(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMIsDynamicRouteServerMapToVirtualNetworkInterfacePrimaryIPPrototypeReservedIPPrototypeVirtualNetworkInterfacePrimaryIPContext(t *testing.T) {
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

	result, err := vpc.ResourceIBMIsDynamicRouteServerMapToVirtualNetworkInterfacePrimaryIPPrototypeReservedIPPrototypeVirtualNetworkInterfacePrimaryIPContext(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMIsDynamicRouteServerMapToResourceGroupIdentity(t *testing.T) {
	checkResult := func(result vpcv1.ResourceGroupIdentityIntf) {
		model := new(vpcv1.ResourceGroupIdentity)
		model.ID = core.StringPtr("fee82deba12e4c0fb69c3b09d1f12345")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["id"] = "fee82deba12e4c0fb69c3b09d1f12345"

	result, err := vpc.ResourceIBMIsDynamicRouteServerMapToResourceGroupIdentity(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMIsDynamicRouteServerMapToResourceGroupIdentityByID(t *testing.T) {
	checkResult := func(result *vpcv1.ResourceGroupIdentityByID) {
		model := new(vpcv1.ResourceGroupIdentityByID)
		model.ID = core.StringPtr("fee82deba12e4c0fb69c3b09d1f12345")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["id"] = "fee82deba12e4c0fb69c3b09d1f12345"

	result, err := vpc.ResourceIBMIsDynamicRouteServerMapToResourceGroupIdentityByID(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMIsDynamicRouteServerMapToSecurityGroupIdentity(t *testing.T) {
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

	result, err := vpc.ResourceIBMIsDynamicRouteServerMapToSecurityGroupIdentity(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMIsDynamicRouteServerMapToSecurityGroupIdentityByID(t *testing.T) {
	checkResult := func(result *vpcv1.SecurityGroupIdentityByID) {
		model := new(vpcv1.SecurityGroupIdentityByID)
		model.ID = core.StringPtr("r006-be5df5ca-12a0-494b-907e-aa6ec2bfa271")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["id"] = "r006-be5df5ca-12a0-494b-907e-aa6ec2bfa271"

	result, err := vpc.ResourceIBMIsDynamicRouteServerMapToSecurityGroupIdentityByID(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMIsDynamicRouteServerMapToSecurityGroupIdentityByCRN(t *testing.T) {
	checkResult := func(result *vpcv1.SecurityGroupIdentityByCRN) {
		model := new(vpcv1.SecurityGroupIdentityByCRN)
		model.CRN = core.StringPtr("crn:v1:bluemix:public:is:us-south:a/aa2432b1fa4d4ace891e9b80fc104e34::security-group:r006-be5df5ca-12a0-494b-907e-aa6ec2bfa271")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["crn"] = "crn:v1:bluemix:public:is:us-south:a/aa2432b1fa4d4ace891e9b80fc104e34::security-group:r006-be5df5ca-12a0-494b-907e-aa6ec2bfa271"

	result, err := vpc.ResourceIBMIsDynamicRouteServerMapToSecurityGroupIdentityByCRN(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMIsDynamicRouteServerMapToSecurityGroupIdentityByHref(t *testing.T) {
	checkResult := func(result *vpcv1.SecurityGroupIdentityByHref) {
		model := new(vpcv1.SecurityGroupIdentityByHref)
		model.Href = core.StringPtr("https://us-south.iaas.cloud.ibm.com/v1/security_groups/r006-be5df5ca-12a0-494b-907e-aa6ec2bfa271")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["href"] = "https://us-south.iaas.cloud.ibm.com/v1/security_groups/r006-be5df5ca-12a0-494b-907e-aa6ec2bfa271"

	result, err := vpc.ResourceIBMIsDynamicRouteServerMapToSecurityGroupIdentityByHref(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMIsDynamicRouteServerMapToSubnetIdentity(t *testing.T) {
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

	result, err := vpc.ResourceIBMIsDynamicRouteServerMapToSubnetIdentity(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMIsDynamicRouteServerMapToSubnetIdentityByID(t *testing.T) {
	checkResult := func(result *vpcv1.SubnetIdentityByID) {
		model := new(vpcv1.SubnetIdentityByID)
		model.ID = core.StringPtr("0717-7ec86020-1c6e-4889-b3f0-a15f2e50f87e")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["id"] = "0717-7ec86020-1c6e-4889-b3f0-a15f2e50f87e"

	result, err := vpc.ResourceIBMIsDynamicRouteServerMapToSubnetIdentityByID(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMIsDynamicRouteServerMapToSubnetIdentityByCRN(t *testing.T) {
	checkResult := func(result *vpcv1.SubnetIdentityByCRN) {
		model := new(vpcv1.SubnetIdentityByCRN)
		model.CRN = core.StringPtr("crn:v1:bluemix:public:is:us-south-1:a/aa2432b1fa4d4ace891e9b80fc104e34::subnet:0717-7ec86020-1c6e-4889-b3f0-a15f2e50f87e")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["crn"] = "crn:v1:bluemix:public:is:us-south-1:a/aa2432b1fa4d4ace891e9b80fc104e34::subnet:0717-7ec86020-1c6e-4889-b3f0-a15f2e50f87e"

	result, err := vpc.ResourceIBMIsDynamicRouteServerMapToSubnetIdentityByCRN(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMIsDynamicRouteServerMapToSubnetIdentityByHref(t *testing.T) {
	checkResult := func(result *vpcv1.SubnetIdentityByHref) {
		model := new(vpcv1.SubnetIdentityByHref)
		model.Href = core.StringPtr("https://us-south.iaas.cloud.ibm.com/v1/subnets/0717-7ec86020-1c6e-4889-b3f0-a15f2e50f87e")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["href"] = "https://us-south.iaas.cloud.ibm.com/v1/subnets/0717-7ec86020-1c6e-4889-b3f0-a15f2e50f87e"

	result, err := vpc.ResourceIBMIsDynamicRouteServerMapToSubnetIdentityByHref(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMIsDynamicRouteServerMapToDynamicRouteServerMemberPrototypeVirtualNetworkInterfaceVirtualNetworkInterfacePrototypeDynamicRouteServerContext(t *testing.T) {
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

	result, err := vpc.ResourceIBMIsDynamicRouteServerMapToDynamicRouteServerMemberPrototypeVirtualNetworkInterfaceVirtualNetworkInterfacePrototypeDynamicRouteServerContext(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMIsDynamicRouteServerMapToDynamicRouteServerMemberPrototypeVirtualNetworkInterfaceVirtualNetworkInterfaceIdentity(t *testing.T) {
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

	result, err := vpc.ResourceIBMIsDynamicRouteServerMapToDynamicRouteServerMemberPrototypeVirtualNetworkInterfaceVirtualNetworkInterfaceIdentity(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMIsDynamicRouteServerMapToDynamicRouteServerMemberPrototypeVirtualNetworkInterfaceVirtualNetworkInterfaceIdentityVirtualNetworkInterfaceIdentityByID(t *testing.T) {
	checkResult := func(result *vpcv1.DynamicRouteServerMemberPrototypeVirtualNetworkInterfaceVirtualNetworkInterfaceIdentityVirtualNetworkInterfaceIdentityByID) {
		model := new(vpcv1.DynamicRouteServerMemberPrototypeVirtualNetworkInterfaceVirtualNetworkInterfaceIdentityVirtualNetworkInterfaceIdentityByID)
		model.ID = core.StringPtr("0717-54eb57ee-86f2-4796-90bb-d7874e0831ef")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["id"] = "0717-54eb57ee-86f2-4796-90bb-d7874e0831ef"

	result, err := vpc.ResourceIBMIsDynamicRouteServerMapToDynamicRouteServerMemberPrototypeVirtualNetworkInterfaceVirtualNetworkInterfaceIdentityVirtualNetworkInterfaceIdentityByID(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMIsDynamicRouteServerMapToDynamicRouteServerMemberPrototypeVirtualNetworkInterfaceVirtualNetworkInterfaceIdentityVirtualNetworkInterfaceIdentityByHref(t *testing.T) {
	checkResult := func(result *vpcv1.DynamicRouteServerMemberPrototypeVirtualNetworkInterfaceVirtualNetworkInterfaceIdentityVirtualNetworkInterfaceIdentityByHref) {
		model := new(vpcv1.DynamicRouteServerMemberPrototypeVirtualNetworkInterfaceVirtualNetworkInterfaceIdentityVirtualNetworkInterfaceIdentityByHref)
		model.Href = core.StringPtr("https://us-south.iaas.cloud.ibm.com/v1/virtual_network_interfaces/0717-54eb57ee-86f2-4796-90bb-d7874e0831ef")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["href"] = "https://us-south.iaas.cloud.ibm.com/v1/virtual_network_interfaces/0717-54eb57ee-86f2-4796-90bb-d7874e0831ef"

	result, err := vpc.ResourceIBMIsDynamicRouteServerMapToDynamicRouteServerMemberPrototypeVirtualNetworkInterfaceVirtualNetworkInterfaceIdentityVirtualNetworkInterfaceIdentityByHref(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMIsDynamicRouteServerMapToDynamicRouteServerMemberPrototypeVirtualNetworkInterfaceVirtualNetworkInterfaceIdentityVirtualNetworkInterfaceIdentityByCRN(t *testing.T) {
	checkResult := func(result *vpcv1.DynamicRouteServerMemberPrototypeVirtualNetworkInterfaceVirtualNetworkInterfaceIdentityVirtualNetworkInterfaceIdentityByCRN) {
		model := new(vpcv1.DynamicRouteServerMemberPrototypeVirtualNetworkInterfaceVirtualNetworkInterfaceIdentityVirtualNetworkInterfaceIdentityByCRN)
		model.CRN = core.StringPtr("crn:v1:bluemix:public:is:us-south-1:a/aa2432b1fa4d4ace891e9b80fc104e34::virtual-network-interface:0717-54eb57ee-86f2-4796-90bb-d7874e0831ef")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["crn"] = "crn:v1:bluemix:public:is:us-south-1:a/aa2432b1fa4d4ace891e9b80fc104e34::virtual-network-interface:0717-54eb57ee-86f2-4796-90bb-d7874e0831ef"

	result, err := vpc.ResourceIBMIsDynamicRouteServerMapToDynamicRouteServerMemberPrototypeVirtualNetworkInterfaceVirtualNetworkInterfaceIdentityVirtualNetworkInterfaceIdentityByCRN(model)
	assert.Nil(t, err)
	checkResult(result)
}
