// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"errors"
	"fmt"
	"regexp"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/vpc"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/stretchr/testify/assert"
)

func TestAccIBMISVPNGateway_basic(t *testing.T) {
	var vpnGateway string
	vpcname := fmt.Sprintf("tfvpnuat-vpc-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tfvpnuat-subnet-%d", acctest.RandIntRange(10, 100))
	name1 := fmt.Sprintf("tfvpnuat-createname-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISVPNGatewayDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISVPNGatewayConfig(vpcname, subnetname, name1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISVPNGatewayExists("ibm_is_vpn_gateway.testacc_vpnGateway", vpnGateway),
					resource.TestCheckResourceAttr(
						"ibm_is_vpn_gateway.testacc_vpnGateway", "name", name1),
					resource.TestCheckResourceAttrSet("ibm_is_vpn_gateway.testacc_vpnGateway", "lifecycle_state"),
					resource.TestCheckResourceAttrSet("ibm_is_vpn_gateway.testacc_vpnGateway", "health_state"),
				),
			},
		},
	})
}

func TestAccIBMISVPNGateway_route(t *testing.T) {
	var vpnGateway string
	vpcname := fmt.Sprintf("tfvpnuat-vpc-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tfvpnuat-subnet-%d", acctest.RandIntRange(10, 100))
	name1 := fmt.Sprintf("tfvpnuat-createname-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISVPNGatewayDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISVPNGatewayRouteConfig(vpcname, subnetname, name1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISVPNGatewayExists("ibm_is_vpn_gateway.testacc_vpnGateway", vpnGateway),
					resource.TestCheckResourceAttr(
						"ibm_is_vpn_gateway.testacc_vpnGateway", "name", name1),
					resource.TestCheckResourceAttr(
						"ibm_is_vpn_gateway.testacc_vpnGateway", "mode", "route"),
					resource.TestCheckResourceAttr(
						"ibm_is_vpn_gateway.testacc_vpnGateway", "local_asn", "64520"),
				),
			},
			{
				Config: testAccCheckIBMISVPNGatewayRouteConfig(vpcname, subnetname, name1),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("ibm_is_vpn_gateway.testacc_vpnGateway", "vpc.#"),
					resource.TestCheckResourceAttrSet("ibm_is_vpn_gateway.testacc_vpnGateway", "vpc.0.name"),
					resource.TestCheckResourceAttrSet("ibm_is_vpn_gateway.testacc_vpnGateway", "vpc.0.crn"),
					resource.TestCheckResourceAttrSet("ibm_is_vpn_gateway.testacc_vpnGateway", "vpc.0.href"),
					resource.TestCheckResourceAttrSet("ibm_is_vpn_gateway.testacc_vpnGateway", "vpc.0.id"),
				),
			},
		},
	})
}

func testAccCheckIBMISVPNGatewayDestroy(s *terraform.State) error {

	sess, _ := acc.TestAccProvider.Meta().(conns.ClientSession).VpcV1API()
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_is_vpn_gateway" {
			continue
		}

		getvpngcptions := &vpcv1.GetVPNGatewayConnectionOptions{
			ID: &rs.Primary.ID,
		}
		_, _, err := sess.GetVPNGatewayConnection(getvpngcptions)

		if err == nil {
			return fmt.Errorf("vpnGateway still exists: %s", rs.Primary.ID)
		}
	}

	return nil
}

func testAccCheckIBMISVPNGatewayExists(n, vpnGatewayID string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return errors.New("No Record ID is set")
		}
		sess, _ := acc.TestAccProvider.Meta().(conns.ClientSession).VpcV1API()
		getvpngcptions := &vpcv1.GetVPNGatewayOptions{
			ID: &rs.Primary.ID,
		}
		foundvpnGatewayIntf, _, err := sess.GetVPNGateway(getvpngcptions)
		if err != nil {
			return err
		}
		foundvpnGateway := foundvpnGatewayIntf.(*vpcv1.VPNGateway)
		vpnGatewayID = *foundvpnGateway.ID
		return nil
	}
}

func testAccCheckIBMISVPNGatewayConfig(vpc, subnet, name string) string {
	return fmt.Sprintf(`
	resource "ibm_is_vpc" "testacc_vpc" {
		name = "%s"
	}

	resource "ibm_is_subnet" "testacc_subnet" {
		name = "%s"
		vpc = "${ibm_is_vpc.testacc_vpc.id}"
		zone = "%s"
		ipv4_cidr_block = "%s"
	}
	resource "ibm_is_vpn_gateway" "testacc_vpnGateway" {
	name = "%s"
	subnet = "${ibm_is_subnet.testacc_subnet.id}"
	mode = "policy"
	}`, vpc, subnet, acc.ISZoneName, acc.ISCIDR, name)

}

func testAccCheckIBMISVPNGatewayRouteConfig(vpc, subnet, name string) string {
	return fmt.Sprintf(`
	resource "ibm_is_vpc" "testacc_vpc" {
		name = "%s"
	}
	resource "ibm_is_subnet" "testacc_subnet" {
		name = "%s"
		vpc = "${ibm_is_vpc.testacc_vpc.id}"
		zone = "%s"
		ipv4_cidr_block = "%s"
	}
	resource "ibm_is_vpn_gateway" "testacc_vpnGateway" {
		name = "%s"
		subnet = "${ibm_is_subnet.testacc_subnet.id}"
		mode = "route"
		local_asn = 64520
		lifecycle {
			ignore_changes = [
				advertised_cidrs
			]
		}
	}`, vpc, subnet, acc.ISZoneName, acc.ISCIDR, name)

}

func testAccCheckIBMISVPNGatewayTaintConfig(vpc, subnet, name string) string {
	return fmt.Sprintf(`
	resource "ibm_is_vpc" "testacc_vpc" {
		name = "%s"
	}

	resource "ibm_is_subnet" "testacc_subnet" {
		name = "%s"
		vpc = "${ibm_is_vpc.testacc_vpc.id}"
		zone = "%s"
		ipv4_cidr_block = "%s"
	}

	resource "ibm_is_vpn_gateway" "testacc_vpnGateway" {
		name 	= "%s"
		subnet 	= "${ibm_is_subnet.testacc_subnet.id}"
		mode 	= "policy"
		timeouts{
			create = "2m"
		}
		lifecycle {
			create_before_destroy = true
		}
	}`, vpc, subnet, acc.ISZoneName, acc.ISCIDR, name)

}
func testAccCheckIBMISVPNGatewayTaintConfig2(vpc, subnet, name string) string {
	return fmt.Sprintf(`
	resource "ibm_is_vpc" "testacc_vpc" {
		name = "%s"
	}

	resource "ibm_is_subnet" "testacc_subnet" {
		name = "%s"
		vpc = "${ibm_is_vpc.testacc_vpc.id}"
		zone = "%s"
		ipv4_cidr_block = "%s"
	}

	resource "ibm_is_vpn_gateway" "testacc_vpnGateway" {
		name 	= "%s"
		subnet 	= "${ibm_is_subnet.testacc_subnet.id}"
		mode 	= "policy"
		timeouts{
			create = "12m"
		}
		lifecycle {
			create_before_destroy = true
		}
	}`, vpc, subnet, acc.ISZoneName, acc.ISCIDR, name)

}

func TestAccIBMISVPNGateway_taint(t *testing.T) {
	var vpnGateway string
	vpcname := fmt.Sprintf("tfvpnuat-vpc-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tfvpnuat-subnet-%d", acctest.RandIntRange(10, 100))
	name1 := fmt.Sprintf("tfvpnuat-taintname-%d", acctest.RandIntRange(10, 100))
	name2 := fmt.Sprintf("tfvpnuat-createname-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISVPNGatewayDestroy,
		Steps: []resource.TestStep{
			{
				Config:      testAccCheckIBMISVPNGatewayTaintConfig(vpcname, subnetname, name1),
				ExpectError: regexp.MustCompile(fmt.Sprintf("timeout while waiting for state to become 'done,")),
			},
			{
				Config: testAccCheckIBMISVPNGatewayTaintConfig2(vpcname, subnetname, name2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISVPNGatewayExists("ibm_is_vpn_gateway.testacc_vpnGateway", vpnGateway),
					resource.TestCheckResourceAttr(
						"ibm_is_vpn_gateway.testacc_vpnGateway", "name", name2),
				),
			},
		},
	})
}

func TestAccIBMISVPNGateway_regional_basic(t *testing.T) {
	var vpnGateway string
	vpcname := fmt.Sprintf("tfvpnuat-vpc-%d", acctest.RandIntRange(10, 100))
	subnet1name := fmt.Sprintf("tfvpnuat-subnet1-%d", acctest.RandIntRange(10, 100))
	subnet2name := fmt.Sprintf("tfvpnuat-subnet2-%d", acctest.RandIntRange(10, 100))
	name1 := fmt.Sprintf("tfvpnuat-regional-%d", acctest.RandIntRange(10, 100))
	name2 := fmt.Sprintf("tfvpnuat-regional-upd-%d", acctest.RandIntRange(10, 100))
	res := "ibm_is_vpn_gateway.testacc_vpnGateway"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISVPNGatewayDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISVPNGatewayRegionalConfig(vpcname, subnet1name, subnet2name, name1, acc.ISZoneName, acc.ISZoneName2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISVPNGatewayExists(res, vpnGateway),
					resource.TestCheckResourceAttr(res, "name", name1),
					resource.TestCheckResourceAttr(res, "availability_mode", "regional"),
					resource.TestCheckResourceAttr(res, "mode", "route"),
					resource.TestCheckResourceAttr(res, "local_asn", "64520"),
					resource.TestCheckResourceAttr(res, "subnet", ""),
					resource.TestCheckResourceAttr(res, "members.#", "2"),
					// The members must land in the subnets that were asked for (order is not guaranteed).
					resource.TestCheckTypeSetElemAttrPair(res, "members.*.private_ip.0.subnet.0.id", "ibm_is_subnet.testacc_subnet1", "id"),
					resource.TestCheckTypeSetElemAttrPair(res, "members.*.private_ip.0.subnet.0.id", "ibm_is_subnet.testacc_subnet2", "id"),
					resource.TestCheckResourceAttrSet(res, "members.0.id"),
					resource.TestCheckResourceAttrSet(res, "members.0.role"),
					resource.TestCheckResourceAttrSet(res, "members.0.lifecycle_state"),
					resource.TestCheckResourceAttrSet(res, "members.0.health_state"),
					resource.TestCheckResourceAttrSet(res, "members.0.private_ip.0.address"),
					resource.TestCheckResourceAttrSet(res, "members.1.private_ip.0.address"),
					resource.TestCheckResourceAttrSet(res, "members.0.public_ip.0.address"),
					// Legacy flat member attributes must still be filled in.
					resource.TestCheckResourceAttrSet(res, "members.0.address"),
					resource.TestCheckResourceAttrSet(res, "members.0.private_address"),
					resource.TestCheckResourceAttrSet(res, "public_ip_address"),
					resource.TestCheckResourceAttrSet(res, "public_ip_address2"),
					resource.TestCheckResourceAttrSet(res, "lifecycle_state"),
					resource.TestCheckResourceAttrSet(res, "health_state"),
					resource.TestCheckResourceAttrSet(res, "vpc.0.name"),
				),
			},
			{
				// Re-applying the same config must be a no-op (no drift from members).
				Config:   testAccCheckIBMISVPNGatewayRegionalConfig(vpcname, subnet1name, subnet2name, name1, acc.ISZoneName, acc.ISZoneName2),
				PlanOnly: true,
			},
			{
				// name is still updatable in place on a regional gateway.
				Config: testAccCheckIBMISVPNGatewayRegionalConfig(vpcname, subnet1name, subnet2name, name2, acc.ISZoneName, acc.ISZoneName2),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(res, "name", name2),
					resource.TestCheckResourceAttr(res, "availability_mode", "regional"),
				),
			},
			{
				ResourceName:      res,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"tags", "access_tags",
				},
			},
		},
	})
}

// Both members in the same zone is a valid regional layout.
func TestAccIBMISVPNGateway_regional_same_zone(t *testing.T) {
	vpcname := fmt.Sprintf("tfvpnuat-vpc-%d", acctest.RandIntRange(10, 100))
	subnet1name := fmt.Sprintf("tfvpnuat-subnet1-%d", acctest.RandIntRange(10, 100))
	subnet2name := fmt.Sprintf("tfvpnuat-subnet2-%d", acctest.RandIntRange(10, 100))
	name1 := fmt.Sprintf("tfvpnuat-regional-sz-%d", acctest.RandIntRange(10, 100))
	res := "ibm_is_vpn_gateway.testacc_vpnGateway"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISVPNGatewayDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISVPNGatewayRegionalConfig(vpcname, subnet1name, subnet2name, name1, acc.ISZoneName, acc.ISZoneName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(res, "availability_mode", "regional"),
					resource.TestCheckResourceAttr(res, "members.#", "2"),
					resource.TestCheckTypeSetElemAttrPair(res, "members.*.private_ip.0.subnet.0.id", "ibm_is_subnet.testacc_subnet1", "id"),
					resource.TestCheckTypeSetElemAttrPair(res, "members.*.private_ip.0.subnet.0.id", "ibm_is_subnet.testacc_subnet2", "id"),
				),
			},
		},
	})
}

// A gateway created without availability_mode must read back as zonal and
// keep its subnet. This guards the existing behavior.
func TestAccIBMISVPNGateway_zonal_default(t *testing.T) {
	vpcname := fmt.Sprintf("tfvpnuat-vpc-%d", acctest.RandIntRange(10, 100))
	subnet1name := fmt.Sprintf("tfvpnuat-subnet1-%d", acctest.RandIntRange(10, 100))
	name1 := fmt.Sprintf("tfvpnuat-zonal-%d", acctest.RandIntRange(10, 100))
	res := "ibm_is_vpn_gateway.testacc_vpnGateway"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISVPNGatewayDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISVPNGatewayZonalConfig(vpcname, subnet1name, name1, ""),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(res, "availability_mode", "zonal"),
					resource.TestCheckResourceAttrPair(res, "subnet", "ibm_is_subnet.testacc_subnet1", "id"),
					resource.TestCheckResourceAttrSet(res, "members.0.id"),
					resource.TestCheckResourceAttrPair(res, "members.0.private_ip.0.subnet.0.id", "ibm_is_subnet.testacc_subnet1", "id"),
				),
			},
			{
				// Writing availability_mode = "zonal" explicitly must not cause a diff.
				Config:   testAccCheckIBMISVPNGatewayZonalConfig(vpcname, subnet1name, name1, "zonal"),
				PlanOnly: true,
			},
		},
	})
}

// availability_mode cannot be patched. Moving from zonal to regional in
// config must replace the gateway instead of calling update.
func TestAccIBMISVPNGateway_availability_mode_forces_new(t *testing.T) {
	var firstID string
	vpcname := fmt.Sprintf("tfvpnuat-vpc-%d", acctest.RandIntRange(10, 100))
	subnet1name := fmt.Sprintf("tfvpnuat-subnet1-%d", acctest.RandIntRange(10, 100))
	subnet2name := fmt.Sprintf("tfvpnuat-subnet2-%d", acctest.RandIntRange(10, 100))
	name1 := fmt.Sprintf("tfvpnuat-forcenew-%d", acctest.RandIntRange(10, 100))
	res := "ibm_is_vpn_gateway.testacc_vpnGateway"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISVPNGatewayDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISVPNGatewayZonalWithSecondSubnetConfig(vpcname, subnet1name, subnet2name, name1),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(res, "availability_mode", "zonal"),
					testAccCaptureResourceID(res, &firstID),
				),
			},
			{
				Config: testAccCheckIBMISVPNGatewayRegionalConfig(vpcname, subnet1name, subnet2name, name1, acc.ISZoneName, acc.ISZoneName2),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(res, "availability_mode", "regional"),
					resource.TestCheckResourceAttr(res, "members.#", "2"),
					testAccCheckResourceIDChanged(res, &firstID),
				),
			},
		},
	})
}

func TestAccIBMISVPNGateway_regional_with_advertised_cidrs(t *testing.T) {
	vpcname := fmt.Sprintf("tfvpnuat-vpc-%d", acctest.RandIntRange(10, 100))
	subnet1name := fmt.Sprintf("tfvpnuat-subnet1-%d", acctest.RandIntRange(10, 100))
	subnet2name := fmt.Sprintf("tfvpnuat-subnet2-%d", acctest.RandIntRange(10, 100))
	name1 := fmt.Sprintf("tfvpnuat-regional-cidrs-%d", acctest.RandIntRange(10, 100))
	res := "ibm_is_vpn_gateway.testacc_vpnGateway"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISVPNGatewayDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISVPNGatewayRegionalWithAdvertisedCIDRsConfig(vpcname, subnet1name, subnet2name, name1),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(res, "availability_mode", "regional"),
					resource.TestCheckResourceAttr("ibm_is_vpn_gateway_advertised_cidr.testacc_cidr", "cidr", "10.45.0.0/25"),
				),
			},
		},
	})
}

// Plan-time validation of the zonal and regional rules. These configs never
// reach the API, so fake subnet IDs are fine.
func TestAccIBMISVPNGateway_availability_validation(t *testing.T) {
	member := `
		members {
			private_ip {
				subnet {
					id = "%s"
				}
			}
		}`
	twoMembers := fmt.Sprintf(member, "subnet-a") + fmt.Sprintf(member, "subnet-b")
	cases := []struct {
		body string
		err  string
	}{
		{`availability_mode = "regional"
		mode = "policy"` + twoMembers, `supported only for route-based`},
		{`availability_mode = "regional"
		subnet = "subnet-a"` + twoMembers, `subnet must not be set`},
		{`availability_mode = "regional"` + fmt.Sprintf(member, "subnet-a"), `(minimum|At least 2|Too few|exactly 2 members)`},
		{`subnet = "subnet-a"` + twoMembers, `members can be set only when availability_mode is "regional"`},
		{`mode = "route"`, `subnet is required`},
		{`availability_mode = "global"
		subnet = "subnet-a"`, `availability_mode`},
	}
	steps := make([]resource.TestStep, 0, len(cases))
	for _, c := range cases {
		steps = append(steps, resource.TestStep{
			Config: fmt.Sprintf(`
	resource "ibm_is_vpn_gateway" "invalid" {
		name = "tfvpnuat-invalid"
		%s
	}`, c.body),
			PlanOnly:    true,
			ExpectError: regexp.MustCompile(c.err),
		})
	}
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps:     steps,
	})
}

func TestResourceIBMISVPNGatewayValidateAvailability(t *testing.T) {
	cases := []struct {
		name             string
		availabilityMode string
		mode             string
		subnetSet        bool
		members          int
		wantErr          string
	}{
		{"zonal default with subnet", "", "route", true, 0, ""},
		{"zonal explicit with subnet", "zonal", "policy", true, 0, ""},
		{"zonal without subnet", "zonal", "route", false, 0, "subnet is required"},
		{"zonal with members", "", "route", true, 2, "members can be set only"},
		{"regional route two members", "regional", "route", false, 2, ""},
		{"regional policy", "regional", "policy", false, 2, "route-based"},
		{"regional with subnet", "regional", "route", true, 2, "subnet must not be set"},
		{"regional one member", "regional", "route", false, 1, "exactly 2 members"},
		{"regional zero members", "regional", "route", false, 0, "exactly 2 members"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			err := vpc.ResourceIBMISVPNGatewayValidateAvailability(c.availabilityMode, c.mode, c.subnetSet, c.members)
			if c.wantErr == "" {
				assert.NoError(t, err)
				return
			}
			assert.ErrorContains(t, err, c.wantErr)
		})
	}
}

func TestResourceIBMISVPNGatewaySchemaAvailability(t *testing.T) {
	s := vpc.ResourceIBMISVPNGateway().Schema
	assert.True(t, s["availability_mode"].ForceNew, "availability_mode must not be patchable")
	assert.True(t, s["availability_mode"].Optional)
	assert.True(t, s["availability_mode"].Computed)
	assert.False(t, s["subnet"].Required, "subnet must be optional so regional gateways can omit it")
	assert.True(t, s["members"].Optional)
	assert.True(t, s["members"].Computed)
	assert.Equal(t, 2, s["members"].MaxItems)
	member := s["members"].Elem.(*schema.Resource).Schema
	assert.True(t, member["role"].Computed)
	assert.False(t, member["role"].Required, "role is not part of the member prototype")
	subnet := member["private_ip"].Elem.(*schema.Resource).Schema["subnet"].Elem.(*schema.Resource).Schema
	for _, k := range []string{"id", "crn", "href"} {
		assert.NotNil(t, subnet[k].DiffSuppressFunc, "members.private_ip.subnet.%s must ignore moves made after create", k)
	}
	assert.NoError(t, vpc.ResourceIBMISVPNGateway().InternalValidate(nil, true))
}

func TestResourceIBMIsVPNGatewayVPNGatewayMemberToMap(t *testing.T) {
	t.Run("nil member", func(t *testing.T) {
		result, err := vpc.ResourceIBMIsVPNGatewayVPNGatewayMemberToMap(nil)
		assert.NoError(t, err)
		assert.Empty(t, result)
	})

	t.Run("pending member without ips", func(t *testing.T) {
		model := &vpcv1.VPNGatewayMember{
			ID:             core.StringPtr("r006-member"),
			LifecycleState: core.StringPtr("pending"),
		}
		result, err := vpc.ResourceIBMIsVPNGatewayVPNGatewayMemberToMap(model)
		assert.NoError(t, err)
		assert.Equal(t, "r006-member", result["id"])
		assert.Equal(t, "pending", result["lifecycle_state"])
		assert.NotContains(t, result, "health_state")
		assert.NotContains(t, result, "role")
		assert.NotContains(t, result, "private_ip")
		assert.NotContains(t, result, "public_ip")
		assert.Equal(t, []map[string]interface{}{}, result["health_reasons"])
		assert.Equal(t, []map[string]interface{}{}, result["lifecycle_reasons"])
	})

	t.Run("full member", func(t *testing.T) {
		model := &vpcv1.VPNGatewayMember{
			HealthReasons: []vpcv1.VPNGatewayMemberHealthReason{{
				Code:    core.StringPtr("cannot_reserve_ip_address"),
				Message: core.StringPtr("IP address exhaustion"),
			}},
			HealthState: core.StringPtr("degraded"),
			ID:          core.StringPtr("r006-member"),
			LifecycleReasons: []vpcv1.VPNGatewayMemberLifecycleReason{{
				Code:     core.StringPtr("internal_error"),
				Message:  core.StringPtr("internal error"),
				MoreInfo: core.StringPtr("https://cloud.ibm.com/docs"),
			}},
			LifecycleState: core.StringPtr("stable"),
			PrivateIP: &vpcv1.ReservedIPReferenceVPNGatewayMemberContext{
				Address:      core.StringPtr("10.240.0.5"),
				Href:         core.StringPtr("https://example/reserved_ips/r006-rip"),
				ID:           core.StringPtr("r006-rip"),
				Name:         core.StringPtr("rip"),
				ResourceType: core.StringPtr("subnet_reserved_ip"),
				Subnet: &vpcv1.SubnetReference{
					CRN:          core.StringPtr("crn:v1:subnet"),
					Href:         core.StringPtr("https://example/subnets/s1"),
					ID:           core.StringPtr("s1"),
					Name:         core.StringPtr("subnet-1"),
					ResourceType: core.StringPtr("subnet"),
					Deleted:      &vpcv1.Deleted{},
				},
			},
			PublicIP: &vpcv1.IP{Address: core.StringPtr("169.0.0.1")},
			Role:     core.StringPtr("active"),
		}
		result, err := vpc.ResourceIBMIsVPNGatewayVPNGatewayMemberToMap(model)
		assert.NoError(t, err)
		assert.Equal(t, "degraded", result["health_state"])
		assert.Equal(t, "active", result["role"])
		assert.Equal(t, "169.0.0.1", result["address"])
		assert.Equal(t, "169.0.0.1", result["public_ip_address"])
		assert.Equal(t, "10.240.0.5", result["private_address"])
		assert.Equal(t, "10.240.0.5", result["private_ip_address"])
		privateIP := result["private_ip"].([]map[string]interface{})[0]
		subnet := privateIP["subnet"].([]map[string]interface{})[0]
		assert.Equal(t, "s1", subnet["id"])
		// Deleted without more_info must not panic and gives an empty map.
		assert.Equal(t, []map[string]interface{}{{}}, subnet["deleted"])
		healthReason := result["health_reasons"].([]map[string]interface{})[0]
		assert.NotContains(t, healthReason, "more_info")
		lifecycleReason := result["lifecycle_reasons"].([]map[string]interface{})[0]
		assert.Equal(t, "https://cloud.ibm.com/docs", lifecycleReason["more_info"])
	})
}

func testAccCaptureResourceID(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}
		*id = rs.Primary.ID
		return nil
	}
}

func testAccCheckResourceIDChanged(n string, oldID *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}
		if rs.Primary.ID == *oldID {
			return fmt.Errorf("expected %s to be replaced, but ID is still %s", n, *oldID)
		}
		return nil
	}
}

func testAccCheckIBMISVPNGatewayRegionalConfig(vpc, subnet1, subnet2, name, zone1, zone2 string) string {
	return fmt.Sprintf(`
	resource "ibm_is_vpc" "testacc_vpc" {
		name = "%s"
	}

	resource "ibm_is_subnet" "testacc_subnet1" {
		name                     = "%s"
		vpc                      = ibm_is_vpc.testacc_vpc.id
		zone                     = "%s"
		total_ipv4_address_count = 16
	}

	resource "ibm_is_subnet" "testacc_subnet2" {
		name                     = "%s"
		vpc                      = ibm_is_vpc.testacc_vpc.id
		zone                     = "%s"
		total_ipv4_address_count = 16
	}

	resource "ibm_is_vpn_gateway" "testacc_vpnGateway" {
		name              = "%s"
		availability_mode = "regional"
		mode              = "route"
		local_asn         = 64520
		members {
			private_ip {
				subnet {
					id = ibm_is_subnet.testacc_subnet1.id
				}
			}
		}
		members {
			private_ip {
				subnet {
					id = ibm_is_subnet.testacc_subnet2.id
				}
			}
		}
	}`, vpc, subnet1, zone1, subnet2, zone2, name)
}

func testAccCheckIBMISVPNGatewayZonalConfig(vpc, subnet1, name, availabilityMode string) string {
	availabilityModeLine := ""
	if availabilityMode != "" {
		availabilityModeLine = fmt.Sprintf("availability_mode = %q", availabilityMode)
	}
	return fmt.Sprintf(`
	resource "ibm_is_vpc" "testacc_vpc" {
		name = "%s"
	}

	resource "ibm_is_subnet" "testacc_subnet1" {
		name                     = "%s"
		vpc                      = ibm_is_vpc.testacc_vpc.id
		zone                     = "%s"
		total_ipv4_address_count = 16
	}

	resource "ibm_is_vpn_gateway" "testacc_vpnGateway" {
		name   = "%s"
		subnet = ibm_is_subnet.testacc_subnet1.id
		mode   = "route"
		%s
	}`, vpc, subnet1, acc.ISZoneName, name, availabilityModeLine)
}

func testAccCheckIBMISVPNGatewayZonalWithSecondSubnetConfig(vpc, subnet1, subnet2, name string) string {
	return fmt.Sprintf(`
	resource "ibm_is_vpc" "testacc_vpc" {
		name = "%s"
	}

	resource "ibm_is_subnet" "testacc_subnet1" {
		name                     = "%s"
		vpc                      = ibm_is_vpc.testacc_vpc.id
		zone                     = "%s"
		total_ipv4_address_count = 16
	}

	resource "ibm_is_subnet" "testacc_subnet2" {
		name                     = "%s"
		vpc                      = ibm_is_vpc.testacc_vpc.id
		zone                     = "%s"
		total_ipv4_address_count = 16
	}

	resource "ibm_is_vpn_gateway" "testacc_vpnGateway" {
		name   = "%s"
		subnet = ibm_is_subnet.testacc_subnet1.id
		mode   = "route"
	}`, vpc, subnet1, acc.ISZoneName, subnet2, acc.ISZoneName2, name)
}

func testAccCheckIBMISVPNGatewayRegionalWithAdvertisedCIDRsConfig(vpc, subnet1, subnet2, name string) string {
	return testAccCheckIBMISVPNGatewayRegionalConfig(vpc, subnet1, subnet2, name, acc.ISZoneName, acc.ISZoneName2) + `
	resource "ibm_is_vpn_gateway_advertised_cidr" "testacc_cidr" {
		vpn_gateway = ibm_is_vpn_gateway.testacc_vpnGateway.id
		cidr        = "10.45.0.0/25"
	}`
}
