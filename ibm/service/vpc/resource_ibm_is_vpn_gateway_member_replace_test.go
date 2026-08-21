// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"

	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccIBMISVPNGatewayMemberReplace_basic(t *testing.T) {
	var vpnGatewayMember *vpcv1.VPNGatewayMember

	vpcname := fmt.Sprintf("tf-vpc-%d", acctest.RandIntRange(10, 100))
	oldSubnetname := fmt.Sprintf("tf-old-subnet-%d", acctest.RandIntRange(10, 100))
	newSubnetname := fmt.Sprintf("tf-new-subnet-%d", acctest.RandIntRange(10, 100))
	vpnname := fmt.Sprintf("tf-vpngw-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISVPNGatewayMemberReplaceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISVPNGatewayMemberReplaceConfig(vpcname, oldSubnetname, newSubnetname, vpnname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISVPNGatewayMemberReplaceExists("ibm_is_vpn_gateway_member_replace.example", &vpnGatewayMember),
					resource.TestCheckResourceAttrSet("ibm_is_vpn_gateway_member_replace.example", "vpn_gateway_id"),
					resource.TestCheckResourceAttrSet("ibm_is_vpn_gateway_member_replace.example", "vpn_gateway_member_id"),
					resource.TestCheckResourceAttrSet("ibm_is_vpn_gateway_member_replace.example", "subnet.0.id"),
				),
			},
			{
				Config: testAccCheckIBMISVPNGatewayMemberReplaceConfigWithCRN(vpcname, oldSubnetname, newSubnetname, vpnname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISVPNGatewayMemberReplaceExists("ibm_is_vpn_gateway_member_replace.crn_example", &vpnGatewayMember),
					resource.TestCheckResourceAttrSet("ibm_is_vpn_gateway_member_replace.crn_example", "subnet.0.crn"),
				),
			},
			{
				Config: testAccCheckIBMISVPNGatewayMemberReplaceConfigWithHref(vpcname, oldSubnetname, newSubnetname, vpnname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISVPNGatewayMemberReplaceExists("ibm_is_vpn_gateway_member_replace.href_example", &vpnGatewayMember),
					resource.TestCheckResourceAttrSet("ibm_is_vpn_gateway_member_replace.href_example", "subnet.0.href"),
				),
			},
		},
	})
}

func TestAccIBMISVPNGatewayMemberReplace_import(t *testing.T) {
	vpcname := fmt.Sprintf("tf-vpc-%d", acctest.RandIntRange(10, 100))
	oldSubnetname := fmt.Sprintf("tf-old-subnet-%d", acctest.RandIntRange(10, 100))
	newSubnetname := fmt.Sprintf("tf-new-subnet-%d", acctest.RandIntRange(10, 100))
	vpnname := fmt.Sprintf("tf-vpngw-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISVPNGatewayMemberReplaceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISVPNGatewayMemberReplaceConfig(vpcname, oldSubnetname, newSubnetname, vpnname),
			},
			{
				ResourceName:            "ibm_is_vpn_gateway_member_replace.example",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"subnet"},
			},
		},
	})
}

func testAccCheckIBMISVPNGatewayMemberReplaceDestroy(s *terraform.State) error {
	sess, err := acc.TestAccProvider.Meta().(conns.ClientSession).VpcV1API()
	if err != nil {
		return err
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_is_vpn_gateway_member_replace" {
			continue
		}
		parts, err := flex.IdParts(rs.Primary.ID)
		if err != nil {
			return err
		}
		if len(parts) != 2 {
			return fmt.Errorf("Invalid ID format: %s", rs.Primary.ID)
		}
		gID := parts[0]
		memberID := parts[1]

		getVPNGatewayMemberOptions := &vpcv1.GetVPNGatewayMemberOptions{
			VPNGatewayID: &gID,
			ID:           &memberID,
		}
		member, response, err := sess.GetVPNGatewayMember(getVPNGatewayMemberOptions)
		if err == nil && member != nil {
			return fmt.Errorf("VPN Gateway Member still exists: %v", response)
		}
	}
	return nil
}

func testAccCheckIBMISVPNGatewayMemberReplaceExists(n string, vpnGatewayMember **vpcv1.VPNGatewayMember) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("No VPN Gateway Member Replace ID is set")
		}

		sess, err := acc.TestAccProvider.Meta().(conns.ClientSession).VpcV1API()
		if err != nil {
			return err
		}

		parts, err := flex.IdParts(rs.Primary.ID)
		if err != nil {
			return err
		}
		if len(parts) != 2 {
			return fmt.Errorf("Invalid ID format: %s", rs.Primary.ID)
		}
		gID := parts[0]
		memberID := parts[1]

		getVPNGatewayMemberOptions := &vpcv1.GetVPNGatewayMemberOptions{
			VPNGatewayID: &gID,
			ID:           &memberID,
		}

		member, response, err := sess.GetVPNGatewayMember(getVPNGatewayMemberOptions)
		if err != nil {
			if response != nil && response.StatusCode == 404 {
				return fmt.Errorf("VPN Gateway Member not found: %s", memberID)
			}
			return fmt.Errorf("Error getting VPN Gateway Member: %s\n%s", err, response)
		}
		*vpnGatewayMember = member
		return nil
	}
}

func testAccCheckIBMISVPNGatewayMemberReplaceConfig(vpcname, oldSubnetname, newSubnetname, vpnname string) string {
	return fmt.Sprintf(`
	  resource "ibm_is_vpc" "example" {
		name = "%s"
	  }
	  
	  resource "ibm_is_subnet" "old_subnet" {
		name            = "%s"
		vpc             = ibm_is_vpc.example.id
		zone            = "%s"
		ipv4_cidr_block = "10.240.0.0/24"
	  }
	  
	  resource "ibm_is_subnet" "new_subnet" {
		name            = "%s"
		vpc             = ibm_is_vpc.example.id
		zone            = "%s"
		ipv4_cidr_block = "10.240.1.0/24"
	  }
	  
	  resource "ibm_is_vpn_gateway" "example" {
		name   = "%s"
		subnet = ibm_is_subnet.old_subnet.id
		mode   = "route"
	  }
	  
	  resource "ibm_is_vpn_gateway_member_replace" "example" {
		vpn_gateway_id        = ibm_is_vpn_gateway.example.id
		vpn_gateway_member_id = ibm_is_vpn_gateway.example.members[0].id
		subnet {
		  id = ibm_is_subnet.new_subnet.id
		}
	  }`, vpcname, oldSubnetname, acc.ISZoneName, newSubnetname, acc.ISZoneName, vpnname)
}

func testAccCheckIBMISVPNGatewayMemberReplaceConfigWithCRN(vpcname, oldSubnetname, newSubnetname, vpnname string) string {
	return fmt.Sprintf(`
	  resource "ibm_is_vpc" "example" {
		name = "%s"
	  }
	  
	  resource "ibm_is_subnet" "old_subnet" {
		name            = "%s"
		vpc             = ibm_is_vpc.example.id
		zone            = "%s"
		ipv4_cidr_block = "10.240.2.0/24"
	  }
	  
	  resource "ibm_is_subnet" "new_subnet" {
		name            = "%s"
		vpc             = ibm_is_vpc.example.id
		zone            = "%s"
		ipv4_cidr_block = "10.240.3.0/24"
	  }
	  
	  resource "ibm_is_vpn_gateway" "example" {
		name   = "%s"
		subnet = ibm_is_subnet.old_subnet.id
		mode   = "route"
	  }
	  
	  resource "ibm_is_vpn_gateway_member_replace" "crn_example" {
		vpn_gateway_id        = ibm_is_vpn_gateway.example.id
		vpn_gateway_member_id = ibm_is_vpn_gateway.example.members[0].id
		subnet {
		  crn = ibm_is_subnet.new_subnet.crn
		}
	  }`, vpcname, oldSubnetname, acc.ISZoneName, newSubnetname, acc.ISZoneName, vpnname)
}

func testAccCheckIBMISVPNGatewayMemberReplaceConfigWithHref(vpcname, oldSubnetname, newSubnetname, vpnname string) string {
	return fmt.Sprintf(`
	  resource "ibm_is_vpc" "example" {
		name = "%s"
	  }
	  
	  resource "ibm_is_subnet" "old_subnet" {
		name            = "%s"
		vpc             = ibm_is_vpc.example.id
		zone            = "%s"
		ipv4_cidr_block = "10.240.4.0/24"
	  }
	  
	  resource "ibm_is_subnet" "new_subnet" {
		name            = "%s"
		vpc             = ibm_is_vpc.example.id
		zone            = "%s"
		ipv4_cidr_block = "10.240.5.0/24"
	  }
	  
	  resource "ibm_is_vpn_gateway" "example" {
		name   = "%s"
		subnet = ibm_is_subnet.old_subnet.id
		mode   = "route"
	  }
	  
	  resource "ibm_is_vpn_gateway_member_replace" "href_example" {
		vpn_gateway_id        = ibm_is_vpn_gateway.example.id
		vpn_gateway_member_id = ibm_is_vpn_gateway.example.members[0].id
		subnet {
		  href = ibm_is_subnet.new_subnet.href
		}
	  }`, vpcname, oldSubnetname, acc.ISZoneName, newSubnetname, acc.ISZoneName, vpnname)
}

func TestAccIBMISVPNGatewayMemberReplace_regional(t *testing.T) {
	var vpnGatewayMember *vpcv1.VPNGatewayMember

	vpcname := fmt.Sprintf("tf-vpc-%d", acctest.RandIntRange(10, 100))
	subnet1name := fmt.Sprintf("tf-subnet1-%d", acctest.RandIntRange(10, 100))
	subnet2name := fmt.Sprintf("tf-subnet2-%d", acctest.RandIntRange(10, 100))
	subnet3name := fmt.Sprintf("tf-subnet3-%d", acctest.RandIntRange(10, 100))
	vpnname := fmt.Sprintf("tf-vpngw-regional-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISVPNGatewayMemberReplaceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISVPNGatewayMemberReplaceRegionalConfig(vpcname, subnet1name, subnet2name, subnet3name, vpnname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISVPNGatewayMemberReplaceExists("ibm_is_vpn_gateway_member_replace.example", &vpnGatewayMember),
					resource.TestCheckResourceAttrSet("ibm_is_vpn_gateway_member_replace.example", "vpn_gateway_id"),
					resource.TestCheckResourceAttrSet("ibm_is_vpn_gateway_member_replace.example", "vpn_gateway_member_id"),
					resource.TestCheckResourceAttrSet("ibm_is_vpn_gateway_member_replace.example", "subnet.0.id"),
					resource.TestCheckResourceAttrSet("ibm_is_vpn_gateway_member_replace.example", "private_ip.0.address"),
					resource.TestCheckResourceAttrSet("ibm_is_vpn_gateway_member_replace.example", "public_ip.0.address"),
				),
			},
		},
	})
}

func TestAccIBMISVPNGatewayMemberReplace_regional_multiple(t *testing.T) {
	var vpnGatewayMember1 *vpcv1.VPNGatewayMember
	var vpnGatewayMember2 *vpcv1.VPNGatewayMember

	vpcname := fmt.Sprintf("tf-vpc-%d", acctest.RandIntRange(10, 100))
	subnet1name := fmt.Sprintf("tf-subnet1-%d", acctest.RandIntRange(10, 100))
	subnet2name := fmt.Sprintf("tf-subnet2-%d", acctest.RandIntRange(10, 100))
	subnet3name := fmt.Sprintf("tf-subnet3-%d", acctest.RandIntRange(10, 100))
	subnet4name := fmt.Sprintf("tf-subnet4-%d", acctest.RandIntRange(10, 100))
	vpnname := fmt.Sprintf("tf-vpngw-regional-multi-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISVPNGatewayMemberReplaceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISVPNGatewayMemberReplaceRegionalMultipleConfig(vpcname, subnet1name, subnet2name, subnet3name, subnet4name, vpnname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISVPNGatewayMemberReplaceExists("ibm_is_vpn_gateway_member_replace.member1", &vpnGatewayMember1),
					testAccCheckIBMISVPNGatewayMemberReplaceExists("ibm_is_vpn_gateway_member_replace.member2", &vpnGatewayMember2),
					resource.TestCheckResourceAttrSet("ibm_is_vpn_gateway_member_replace.member1", "vpn_gateway_id"),
					resource.TestCheckResourceAttrSet("ibm_is_vpn_gateway_member_replace.member1", "vpn_gateway_member_id"),
					resource.TestCheckResourceAttrSet("ibm_is_vpn_gateway_member_replace.member2", "vpn_gateway_id"),
					resource.TestCheckResourceAttrSet("ibm_is_vpn_gateway_member_replace.member2", "vpn_gateway_member_id"),
				),
			},
		},
	})
}

func testAccCheckIBMISVPNGatewayMemberReplaceRegionalConfig(vpcname, subnet1name, subnet2name, subnet3name, vpnname string) string {
	return fmt.Sprintf(`
	  resource "ibm_is_vpc" "example" {
		name = "%s"
	  }
	  
	  resource "ibm_is_subnet" "subnet1" {
		name            = "%s"
		vpc             = ibm_is_vpc.example.id
		zone            = "%s"
		ipv4_cidr_block = "10.240.50.0/24"
	  }
	  
	  resource "ibm_is_subnet" "subnet2" {
		name            = "%s"
		vpc             = ibm_is_vpc.example.id
		zone            = "%s"
		ipv4_cidr_block = "10.240.51.0/24"
	  }
	  
	  resource "ibm_is_subnet" "subnet3" {
		name            = "%s"
		vpc             = ibm_is_vpc.example.id
		zone            = "%s"
		ipv4_cidr_block = "10.240.52.0/24"
	  }
	  
	  resource "ibm_is_vpn_gateway" "example" {
		name   = "%s"
		availability_mode = "regional"
		mode   = "route"
		members {
			private_ip {
				subnet {
					id = ibm_is_subnet.subnet1.id
				}
			}
		}
		members {
			private_ip {
				subnet {
					id = ibm_is_subnet.subnet2.id
				}
			}
		}
	  }
	  
	  resource "ibm_is_vpn_gateway_member_replace" "example" {
		vpn_gateway_id        = ibm_is_vpn_gateway.example.id
		vpn_gateway_member_id = ibm_is_vpn_gateway.example.members[0].id
		subnet {
		  id = ibm_is_subnet.subnet3.id
		}
	  }`, vpcname, subnet1name, acc.ISZoneName, subnet2name, acc.ISZoneName2, subnet3name, acc.ISZoneName, vpnname)
}

func testAccCheckIBMISVPNGatewayMemberReplaceRegionalMultipleConfig(vpcname, subnet1name, subnet2name, subnet3name, subnet4name, vpnname string) string {
	return fmt.Sprintf(`
	  resource "ibm_is_vpc" "example" {
		name = "%s"
	  }
	  
	  resource "ibm_is_subnet" "subnet1" {
		name            = "%s"
		vpc             = ibm_is_vpc.example.id
		zone            = "%s"
		ipv4_cidr_block = "10.240.60.0/24"
	  }
	  
	  resource "ibm_is_subnet" "subnet2" {
		name            = "%s"
		vpc             = ibm_is_vpc.example.id
		zone            = "%s"
		ipv4_cidr_block = "10.240.61.0/24"
	  }
	  
	  resource "ibm_is_subnet" "subnet3" {
		name            = "%s"
		vpc             = ibm_is_vpc.example.id
		zone            = "%s"
		ipv4_cidr_block = "10.240.62.0/24"
	  }
	  
	  resource "ibm_is_subnet" "subnet4" {
		name            = "%s"
		vpc             = ibm_is_vpc.example.id
		zone            = "%s"
		ipv4_cidr_block = "10.240.63.0/24"
	  }
	  
	  resource "ibm_is_vpn_gateway" "example" {
		name   = "%s"
		availability_mode = "regional"
		mode   = "route"
		members {
			private_ip {
				subnet {
					id = ibm_is_subnet.subnet1.id
				}
			}
		}
		members {
			private_ip {
				subnet {
					id = ibm_is_subnet.subnet2.id
				}
			}
		}
	  }
	  
	  resource "ibm_is_vpn_gateway_member_replace" "member1" {
		vpn_gateway_id        = ibm_is_vpn_gateway.example.id
		vpn_gateway_member_id = ibm_is_vpn_gateway.example.members[0].id
		subnet {
		  id = ibm_is_subnet.subnet3.id
		}
	  }
	  
	  resource "ibm_is_vpn_gateway_member_replace" "member2" {
		depends_on = [ibm_is_vpn_gateway_member_replace.member1]
		vpn_gateway_id        = ibm_is_vpn_gateway.example.id
		vpn_gateway_member_id = ibm_is_vpn_gateway.example.members[1].id
		subnet {
		  id = ibm_is_subnet.subnet4.id
		}
	  }`, vpcname, subnet1name, acc.ISZoneName, subnet2name, acc.ISZoneName2, subnet3name, acc.ISZoneName, subnet4name, acc.ISZoneName2, vpnname)
}
