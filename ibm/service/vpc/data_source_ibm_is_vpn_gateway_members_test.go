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
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
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
	model := new(vpcv1.PageLink)
	model.Href = core.StringPtr("testString")

	result, err := vpc.DataSourceIBMIsVPNGatewayMembersPageLinkToMap(model)
	assert.Nil(t, err)
	assert.Equal(t, map[string]interface{}{"href": "testString"}, result)

	// A missing page link must not panic.
	result, err = vpc.DataSourceIBMIsVPNGatewayMembersPageLinkToMap(nil)
	assert.Nil(t, err)
	assert.Empty(t, result)
}

// TestDataSourceIBMIsVPNGatewayMembersCollectionItemToMap checks that the
// collection item only carries keys that exist in the data source schema and
// that a pending member without IPs is handled.
func TestDataSourceIBMIsVPNGatewayMembersCollectionItemToMap(t *testing.T) {
	model := &vpcv1.VPNGatewayMember{
		HealthReasons:    []vpcv1.VPNGatewayMemberHealthReason{},
		HealthState:      core.StringPtr("inapplicable"),
		ID:               core.StringPtr("r006-member-1"),
		LifecycleReasons: []vpcv1.VPNGatewayMemberLifecycleReason{},
		LifecycleState:   core.StringPtr("pending"),
		Role:             core.StringPtr("standby"),
	}
	result, err := vpc.DataSourceIBMIsVPNGatewayMembersVPNGatewayMemberCollectionItemToMap(model)
	assert.Nil(t, err)
	assert.Equal(t, "r006-member-1", result["id"])
	assert.Equal(t, "pending", result["lifecycle_state"])
	assert.NotContains(t, result, "private_ip")
	assert.NotContains(t, result, "public_ip")
	assert.NotContains(t, result, "address")
	assert.NotContains(t, result, "private_address")

	schemaKeys := vpc.DataSourceIBMIsVPNGatewayMembers().Schema["members"].Elem.(*schema.Resource).Schema
	for k := range result {
		assert.Contains(t, schemaKeys, k)
	}
}
