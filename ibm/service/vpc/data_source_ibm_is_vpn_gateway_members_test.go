// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.114.0-a902401e-20260427-192904
 */

package vpc_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/stretchr/testify/assert"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/vpc"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/vpc-go-sdk/vpcv1"
)

func TestAccIBMIsVPNGatewayMembersDataSourceBasic(t *testing.T) {
	vpcname := fmt.Sprintf("tfvpnuat-vpc-%d", acctest.RandIntRange(10, 100))
	subnet1name := fmt.Sprintf("tfvpnuat-subnet1-%d", acctest.RandIntRange(10, 100))
	subnet2name := fmt.Sprintf("tfvpnuat-subnet2-%d", acctest.RandIntRange(10, 100))
	vpngwname := fmt.Sprintf("tfvpnuat-vpngw-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIsVPNGatewayMembersDataSourceConfigBasic(vpcname, subnet1name, subnet2name, vpngwname),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_vpn_gateway_members.is_vpn_gateway_members_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpn_gateway_members.is_vpn_gateway_members_instance", "vpn_gateway_id"),
					resource.TestCheckResourceAttr("data.ibm_is_vpn_gateway_members.is_vpn_gateway_members_instance", "members.#", "2"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpn_gateway_members.is_vpn_gateway_members_instance", "members.0.id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpn_gateway_members.is_vpn_gateway_members_instance", "members.0.health_state"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpn_gateway_members.is_vpn_gateway_members_instance", "members.0.lifecycle_state"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpn_gateway_members.is_vpn_gateway_members_instance", "members.0.private_ip.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpn_gateway_members.is_vpn_gateway_members_instance", "members.0.public_ip.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpn_gateway_members.is_vpn_gateway_members_instance", "members.0.role"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpn_gateway_members.is_vpn_gateway_members_instance", "members.1.id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpn_gateway_members.is_vpn_gateway_members_instance", "total_count"),
				),
			},
		},
	})
}

func testAccCheckIBMIsVPNGatewayMembersDataSourceConfigBasic(vpc, subnet1, subnet2, vpngwname string) string {
	return fmt.Sprintf(`
		resource "ibm_is_vpc" "example" {
			name = "%s"
		}
		
		resource "ibm_is_subnet" "example1" {
			name = "%s"
			vpc = ibm_is_vpc.example.id
			zone = "%s"
			ipv4_cidr_block = "10.240.40.0/24"
		}
		
		resource "ibm_is_subnet" "example2" {
			name = "%s"
			vpc = ibm_is_vpc.example.id
			zone = "%s"
			ipv4_cidr_block = "10.240.41.0/24"
		}
		
		resource "ibm_is_vpn_gateway" "example" {
			name = "%s"
			availability_mode = "regional"
			mode = "route"
			members {
				private_ip {
					subnet {
						id = ibm_is_subnet.example1.id
					}
				}
			}
			members {
				private_ip {
					subnet {
						id = ibm_is_subnet.example2.id
					}
				}
			}
		}
		
		data "ibm_is_vpn_gateway_members" "is_vpn_gateway_members_instance" {
			vpn_gateway_id = ibm_is_vpn_gateway.example.id
		}
	`, vpc, subnet1, acc.ISZoneName, subnet2, acc.ISZoneName2, vpngwname)
}

func TestDataSourceIBMIsVPNGatewayMembersPageLinkToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["href"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(vpcv1.PageLink)
	model.Href = core.StringPtr("testString")

	result, err := vpc.DataSourceIBMIsVPNGatewayMembersPageLinkToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMIsVPNGatewayMembersVPNGatewayMemberHealthReasonToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["code"] = "cannot_reserve_ip_address"
		model["message"] = "IP address exhaustion (release addresses on the VPN's subnet)."
		model["more_info"] = "https://cloud.ibm.com/docs/vpc?topic=vpc-vpn-health"

		assert.Equal(t, result, model)
	}

	model := new(vpcv1.VPNGatewayMemberHealthReason)
	model.Code = core.StringPtr("cannot_reserve_ip_address")
	model.Message = core.StringPtr("IP address exhaustion (release addresses on the VPN's subnet).")
	model.MoreInfo = core.StringPtr("https://cloud.ibm.com/docs/vpc?topic=vpc-vpn-health")

	result, err := vpc.DataSourceIBMIsVPNGatewayMembersVPNGatewayMemberHealthReasonToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMIsVPNGatewayMembersVPNGatewayMemberLifecycleReasonToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["code"] = "resource_suspended_by_provider"
		model["message"] = "The resource has been suspended. Contact IBM support with the CRN for next steps."
		model["more_info"] = "https://cloud.ibm.com/apidocs/vpc#resource-suspension"

		assert.Equal(t, result, model)
	}

	model := new(vpcv1.VPNGatewayMemberLifecycleReason)
	model.Code = core.StringPtr("resource_suspended_by_provider")
	model.Message = core.StringPtr("The resource has been suspended. Contact IBM support with the CRN for next steps.")
	model.MoreInfo = core.StringPtr("https://cloud.ibm.com/apidocs/vpc#resource-suspension")

	result, err := vpc.DataSourceIBMIsVPNGatewayMembersVPNGatewayMemberLifecycleReasonToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMIsVPNGatewayMembersDeletedToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["more_info"] = "https://cloud.ibm.com/apidocs/vpc#deleted-resources"

		assert.Equal(t, result, model)
	}

	model := new(vpcv1.Deleted)
	model.MoreInfo = core.StringPtr("https://cloud.ibm.com/apidocs/vpc#deleted-resources")

	result, err := vpc.DataSourceIBMIsVPNGatewayMembersDeletedToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMIsVPNGatewayMembersSubnetReferenceToMap(t *testing.T) {
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

	result, err := vpc.DataSourceIBMIsVPNGatewayMembersSubnetReferenceToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMIsVPNGatewayMembersIPToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["address"] = "192.168.3.4"

		assert.Equal(t, result, model)
	}

	model := new(vpcv1.IP)
	model.Address = core.StringPtr("192.168.3.4")

	result, err := vpc.DataSourceIBMIsVPNGatewayMembersIPToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}
