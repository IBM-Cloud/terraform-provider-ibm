// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.114.0-a902401e-20260427-192904
 */

package vpc_test

import (
	"fmt"
	"regexp"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMIsVPNGatewayMemberDataSourceBasic(t *testing.T) {
	vpcname := fmt.Sprintf("tfvpnuat-vpc-%d", acctest.RandIntRange(10, 100))
	subnet1name := fmt.Sprintf("tfvpnuat-subnet1-%d", acctest.RandIntRange(10, 100))
	subnet2name := fmt.Sprintf("tfvpnuat-subnet2-%d", acctest.RandIntRange(10, 100))
	vpngwname := fmt.Sprintf("tfvpnuat-vpngw-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIsVPNGatewayMemberDataSourceConfigBasic(vpcname, subnet1name, subnet2name, vpngwname),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_vpn_gateway_member.is_vpn_gateway_member_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpn_gateway_member.is_vpn_gateway_member_instance", "vpn_gateway_id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpn_gateway_member.is_vpn_gateway_member_instance", "vpn_gateway_member_id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpn_gateway_member.is_vpn_gateway_member_instance", "health_state"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpn_gateway_member.is_vpn_gateway_member_instance", "lifecycle_state"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpn_gateway_member.is_vpn_gateway_member_instance", "private_ip.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpn_gateway_member.is_vpn_gateway_member_instance", "private_ip.0.address"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpn_gateway_member.is_vpn_gateway_member_instance", "public_ip.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpn_gateway_member.is_vpn_gateway_member_instance", "public_ip.0.address"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpn_gateway_member.is_vpn_gateway_member_instance", "role"),
				),
			},
		},
	})
}

func testAccCheckIBMIsVPNGatewayMemberDataSourceConfigBasic(vpc, subnet1, subnet2, vpngwname string) string {
	return fmt.Sprintf(`
		resource "ibm_is_vpc" "example" {
			name = "%s"
		}
		
		resource "ibm_is_subnet" "example1" {
			name = "%s"
			vpc = ibm_is_vpc.example.id
			zone = "%s"
			ipv4_cidr_block = "10.240.30.0/24"
		}
		
		resource "ibm_is_subnet" "example2" {
			name = "%s"
			vpc = ibm_is_vpc.example.id
			zone = "%s"
			ipv4_cidr_block = "10.240.31.0/24"
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
		
		data "ibm_is_vpn_gateway_member" "is_vpn_gateway_member_instance" {
			vpn_gateway_id = ibm_is_vpn_gateway.example.id
			vpn_gateway_member_id = ibm_is_vpn_gateway.example.members[0].id
		}
	`, vpc, subnet1, acc.ISZoneName, subnet2, acc.ISZoneName2, vpngwname)
}

// TestAccIBMIsVPNGatewayMemberDataSourceZonal reads a member of a zonal
// gateway, to make sure the new member data source works for gateways that
// were created without availability_mode.
func TestAccIBMIsVPNGatewayMemberDataSourceZonal(t *testing.T) {
	vpcname := fmt.Sprintf("tfvpnuat-vpc-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tfvpnuat-subnet-%d", acctest.RandIntRange(10, 100))
	vpngwname := fmt.Sprintf("tfvpnuat-vpngw-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIsVPNGatewayMemberDataSourceConfigZonal(vpcname, subnetname, vpngwname),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("data.ibm_is_vpn_gateway_member.zonal", "id", "ibm_is_vpn_gateway.example", "members.0.id"),
					resource.TestCheckResourceAttrPair("data.ibm_is_vpn_gateway_member.zonal", "private_ip.0.subnet.0.id", "ibm_is_subnet.example", "id"),
					resource.TestCheckResourceAttrPair("data.ibm_is_vpn_gateway_member.zonal", "role", "ibm_is_vpn_gateway.example", "members.0.role"),
					resource.TestCheckResourceAttr("data.ibm_is_vpn_gateway_member.zonal", "health_reasons.#", "0"),
					resource.TestCheckResourceAttr("data.ibm_is_vpn_gateway_member.zonal", "lifecycle_state", "stable"),
				),
			},
		},
	})
}

func testAccCheckIBMIsVPNGatewayMemberDataSourceConfigZonal(vpc, subnet, vpngwname string) string {
	return fmt.Sprintf(`
		resource "ibm_is_vpc" "example" {
			name = "%s"
		}
		resource "ibm_is_subnet" "example" {
			name            = "%s"
			vpc             = ibm_is_vpc.example.id
			zone            = "%s"
			ipv4_cidr_block = "10.240.32.0/24"
		}
		resource "ibm_is_vpn_gateway" "example" {
			name   = "%s"
			subnet = ibm_is_subnet.example.id
			mode   = "route"
		}
		data "ibm_is_vpn_gateway_member" "zonal" {
			vpn_gateway_id        = ibm_is_vpn_gateway.example.id
			vpn_gateway_member_id = ibm_is_vpn_gateway.example.members[0].id
		}
	`, vpc, subnet, acc.ISZoneName, vpngwname)
}

func TestAccIBMIsVPNGatewayMemberDataSourceNotFound(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: `
					data "ibm_is_vpn_gateway_member" "missing" {
						vpn_gateway_id        = "r006-00000000-0000-0000-0000-000000000000"
						vpn_gateway_member_id = "r006-00000000-0000-0000-0000-000000000001"
					}
				`,
				ExpectError: regexp.MustCompile("GetVPNGatewayMemberWithContext failed"),
			},
		},
	})
}
