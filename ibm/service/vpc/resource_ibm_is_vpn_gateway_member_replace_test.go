// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"fmt"
	"regexp"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"

	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

// TestAccIBMISVPNGatewayMemberReplace_basic moves one member of a regional
// gateway to a new subnet, then checks:
//   - the member really is in the new subnet (read from the API, not state)
//   - the gateway does not see the move as drift (no replacement planned)
//   - import works with the composite ID
func TestAccIBMISVPNGatewayMemberReplace_basic(t *testing.T) {
	var member *vpcv1.VPNGatewayMember
	vpcname := fmt.Sprintf("tf-vpc-%d", acctest.RandIntRange(10, 100))
	subnet1name := fmt.Sprintf("tf-subnet1-%d", acctest.RandIntRange(10, 100))
	subnet2name := fmt.Sprintf("tf-subnet2-%d", acctest.RandIntRange(10, 100))
	subnet3name := fmt.Sprintf("tf-subnet3-%d", acctest.RandIntRange(10, 100))
	vpnname := fmt.Sprintf("tf-vpngw-%d", acctest.RandIntRange(10, 100))
	res := "ibm_is_vpn_gateway_member_replace.example"
	config := testAccCheckIBMISVPNGatewayMemberReplaceConfig(vpcname, subnet1name, subnet2name, subnet3name, vpnname, `id = ibm_is_subnet.subnet3.id`)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISVPNGatewayMemberReplaceDestroy,
		Steps: []resource.TestStep{
			{
				Config: config,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISVPNGatewayMemberReplaceExists(res, &member),
					testAccCheckIBMISVPNGatewayMemberInSubnet(&member, "ibm_is_subnet.subnet3"),
					resource.TestCheckResourceAttrPair(res, "vpn_gateway_id", "ibm_is_vpn_gateway.example", "id"),
					resource.TestCheckResourceAttrPair(res, "subnet.0.id", "ibm_is_subnet.subnet3", "id"),
					resource.TestCheckResourceAttrPair(res, "subnet.0.crn", "ibm_is_subnet.subnet3", "crn"),
					resource.TestCheckResourceAttr(res, "lifecycle_state", "stable"),
					resource.TestCheckResourceAttrSet(res, "private_ip_address"),
					resource.TestCheckResourceAttrSet(res, "public_ip_address"),
					resource.TestCheckResourceAttrSet(res, "role"),
				),
			},
			{
				// After the move, neither the gateway nor this resource may plan a change.
				Config:   config,
				PlanOnly: true,
			},
			{
				ResourceName:      res,
				ImportState:       true,
				ImportStateVerify: true,
				// The member ID can change when the member is recreated by the replace.
				ImportStateVerifyIgnore: []string{"vpn_gateway_member_id"},
			},
			{
				ResourceName:  res,
				ImportState:   true,
				ImportStateId: "not-a-composite-id",
				ExpectError:   regexp.MustCompile(`does not contain /`),
			},
		},
	})
}

// TestAccIBMISVPNGatewayMemberReplace_subnetIdentity covers the crn and href
// branches of the subnet identity.
func TestAccIBMISVPNGatewayMemberReplace_subnetIdentity(t *testing.T) {
	var member *vpcv1.VPNGatewayMember
	vpcname := fmt.Sprintf("tf-vpc-%d", acctest.RandIntRange(10, 100))
	subnet1name := fmt.Sprintf("tf-subnet1-%d", acctest.RandIntRange(10, 100))
	subnet2name := fmt.Sprintf("tf-subnet2-%d", acctest.RandIntRange(10, 100))
	subnet3name := fmt.Sprintf("tf-subnet3-%d", acctest.RandIntRange(10, 100))
	vpnname := fmt.Sprintf("tf-vpngw-%d", acctest.RandIntRange(10, 100))
	res := "ibm_is_vpn_gateway_member_replace.example"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISVPNGatewayMemberReplaceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISVPNGatewayMemberReplaceConfig(vpcname, subnet1name, subnet2name, subnet3name, vpnname, `crn = ibm_is_subnet.subnet3.crn`),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISVPNGatewayMemberReplaceExists(res, &member),
					testAccCheckIBMISVPNGatewayMemberInSubnet(&member, "ibm_is_subnet.subnet3"),
					resource.TestCheckResourceAttrPair(res, "subnet.0.id", "ibm_is_subnet.subnet3", "id"),
				),
			},
			{
				Config: testAccCheckIBMISVPNGatewayMemberReplaceConfig(vpcname, subnet1name, subnet2name, subnet3name, vpnname, `href = ibm_is_subnet.subnet1.href`),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISVPNGatewayMemberReplaceExists(res, &member),
					testAccCheckIBMISVPNGatewayMemberInSubnet(&member, "ibm_is_subnet.subnet1"),
					resource.TestCheckResourceAttrPair(res, "subnet.0.id", "ibm_is_subnet.subnet1", "id"),
				),
			},
		},
	})
}

func TestAccIBMISVPNGatewayMemberReplace_errors(t *testing.T) {
	vpcname := fmt.Sprintf("tf-vpc-%d", acctest.RandIntRange(10, 100))
	subnet1name := fmt.Sprintf("tf-subnet1-%d", acctest.RandIntRange(10, 100))
	subnet2name := fmt.Sprintf("tf-subnet2-%d", acctest.RandIntRange(10, 100))
	subnet3name := fmt.Sprintf("tf-subnet3-%d", acctest.RandIntRange(10, 100))
	vpnname := fmt.Sprintf("tf-vpngw-%d", acctest.RandIntRange(10, 100))
	base := testAccCheckIBMISVPNGatewayMemberReplaceBaseConfig(vpcname, subnet1name, subnet2name, subnet3name, vpnname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				// Empty subnet block: rejected before any API call.
				Config: base + `
	resource "ibm_is_vpn_gateway_member_replace" "empty_subnet" {
		vpn_gateway_id        = ibm_is_vpn_gateway.example.id
		vpn_gateway_member_id = ibm_is_vpn_gateway.example.members[0].id
		subnet {}
	}`,
				ExpectError: regexp.MustCompile(`subnet must specify one of id, crn or href`),
			},
			{
				// Unknown member: the API error is surfaced.
				Config: base + `
	resource "ibm_is_vpn_gateway_member_replace" "bad_member" {
		vpn_gateway_id        = ibm_is_vpn_gateway.example.id
		vpn_gateway_member_id = "r006-00000000-0000-0000-0000-000000000000"
		subnet {
			id = ibm_is_subnet.subnet3.id
		}
	}`,
				ExpectError: regexp.MustCompile(`ReplaceVPNGatewayMemberWithContext failed`),
			},
		},
	})
}

func TestAccIBMISVPNGatewayMemberReplace_bothMembers(t *testing.T) {
	var member1, member2 *vpcv1.VPNGatewayMember
	vpcname := fmt.Sprintf("tf-vpc-%d", acctest.RandIntRange(10, 100))
	subnet1name := fmt.Sprintf("tf-subnet1-%d", acctest.RandIntRange(10, 100))
	subnet2name := fmt.Sprintf("tf-subnet2-%d", acctest.RandIntRange(10, 100))
	subnet3name := fmt.Sprintf("tf-subnet3-%d", acctest.RandIntRange(10, 100))
	subnet4name := fmt.Sprintf("tf-subnet4-%d", acctest.RandIntRange(10, 100))
	vpnname := fmt.Sprintf("tf-vpngw-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISVPNGatewayMemberReplaceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISVPNGatewayMemberReplaceBothConfig(vpcname, subnet1name, subnet2name, subnet3name, subnet4name, vpnname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISVPNGatewayMemberReplaceExists("ibm_is_vpn_gateway_member_replace.member1", &member1),
					testAccCheckIBMISVPNGatewayMemberReplaceExists("ibm_is_vpn_gateway_member_replace.member2", &member2),
					testAccCheckIBMISVPNGatewayMemberInSubnet(&member1, "ibm_is_subnet.subnet3"),
					testAccCheckIBMISVPNGatewayMemberInSubnet(&member2, "ibm_is_subnet.subnet4"),
				),
			},
		},
	})
}

// The member replace resource does not own the member, so after destroy the
// member is gone only because its gateway was destroyed in the same run.
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
		member, response, err := sess.GetVPNGatewayMember(&vpcv1.GetVPNGatewayMemberOptions{
			VPNGatewayID: &parts[0],
			ID:           &parts[1],
		})
		if err == nil && member != nil {
			return fmt.Errorf("VPN gateway member still exists: %v", response)
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
			return fmt.Errorf("No VPN gateway member replace ID is set")
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
		member, response, err := sess.GetVPNGatewayMember(&vpcv1.GetVPNGatewayMemberOptions{
			VPNGatewayID: &parts[0],
			ID:           &parts[1],
		})
		if err != nil {
			return fmt.Errorf("Error getting VPN gateway member: %s\n%s", err, response)
		}
		*vpnGatewayMember = member
		return nil
	}
}

// testAccCheckIBMISVPNGatewayMemberInSubnet checks the live member (from the
// API) against the subnet in state, so the test proves the move happened.
func testAccCheckIBMISVPNGatewayMemberInSubnet(member **vpcv1.VPNGatewayMember, subnetResource string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[subnetResource]
		if !ok {
			return fmt.Errorf("Not found: %s", subnetResource)
		}
		m := *member
		if m == nil || m.PrivateIP == nil || m.PrivateIP.Subnet == nil || m.PrivateIP.Subnet.ID == nil {
			return fmt.Errorf("VPN gateway member has no private_ip.subnet")
		}
		if *m.PrivateIP.Subnet.ID != rs.Primary.ID {
			return fmt.Errorf("VPN gateway member is in subnet %s, expected %s", *m.PrivateIP.Subnet.ID, rs.Primary.ID)
		}
		return nil
	}
}

func testAccCheckIBMISVPNGatewayMemberReplaceBaseConfig(vpcname, subnet1name, subnet2name, subnet3name, vpnname string) string {
	return fmt.Sprintf(`
	resource "ibm_is_vpc" "example" {
		name = "%s"
	}

	resource "ibm_is_subnet" "subnet1" {
		name                     = "%s"
		vpc                      = ibm_is_vpc.example.id
		zone                     = "%s"
		total_ipv4_address_count = 16
	}

	resource "ibm_is_subnet" "subnet2" {
		name                     = "%s"
		vpc                      = ibm_is_vpc.example.id
		zone                     = "%s"
		total_ipv4_address_count = 16
	}

	resource "ibm_is_subnet" "subnet3" {
		name                     = "%s"
		vpc                      = ibm_is_vpc.example.id
		zone                     = "%s"
		total_ipv4_address_count = 16
	}

	resource "ibm_is_vpn_gateway" "example" {
		name              = "%s"
		availability_mode = "regional"
		mode              = "route"
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
	`, vpcname, subnet1name, acc.ISZoneName, subnet2name, acc.ISZoneName2, subnet3name, acc.ISZoneName, vpnname)
}

func testAccCheckIBMISVPNGatewayMemberReplaceConfig(vpcname, subnet1name, subnet2name, subnet3name, vpnname, subnetIdentity string) string {
	return testAccCheckIBMISVPNGatewayMemberReplaceBaseConfig(vpcname, subnet1name, subnet2name, subnet3name, vpnname) + fmt.Sprintf(`
	resource "ibm_is_vpn_gateway_member_replace" "example" {
		vpn_gateway_id        = ibm_is_vpn_gateway.example.id
		vpn_gateway_member_id = ibm_is_vpn_gateway.example.members[0].id
		subnet {
			%s
		}
	}`, subnetIdentity)
}

func testAccCheckIBMISVPNGatewayMemberReplaceBothConfig(vpcname, subnet1name, subnet2name, subnet3name, subnet4name, vpnname string) string {
	return testAccCheckIBMISVPNGatewayMemberReplaceBaseConfig(vpcname, subnet1name, subnet2name, subnet3name, vpnname) + fmt.Sprintf(`
	resource "ibm_is_subnet" "subnet4" {
		name                     = "%s"
		vpc                      = ibm_is_vpc.example.id
		zone                     = "%s"
		total_ipv4_address_count = 16
	}

	resource "ibm_is_vpn_gateway_member_replace" "member1" {
		vpn_gateway_id        = ibm_is_vpn_gateway.example.id
		vpn_gateway_member_id = ibm_is_vpn_gateway.example.members[0].id
		subnet {
			id = ibm_is_subnet.subnet3.id
		}
	}

	# One member at a time, so the gateway always keeps one working member.
	resource "ibm_is_vpn_gateway_member_replace" "member2" {
		depends_on            = [ibm_is_vpn_gateway_member_replace.member1]
		vpn_gateway_id        = ibm_is_vpn_gateway.example.id
		vpn_gateway_member_id = ibm_is_vpn_gateway.example.members[1].id
		subnet {
			id = ibm_is_subnet.subnet4.id
		}
	}`, subnet4name, acc.ISZoneName2)
}
