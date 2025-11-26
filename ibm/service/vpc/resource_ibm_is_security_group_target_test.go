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
	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccIBMISSecurityGroupTarget_basic(t *testing.T) {
	var securityGroup string

	vpcname := fmt.Sprintf("tfsg-vpc-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tfsg-subnet-%d", acctest.RandIntRange(10, 100))
	lbname := fmt.Sprintf("tfsg-lb-%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tfsg-one-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISSecurityGroupTargetDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISsecurityGroupTargetConfig(vpcname, subnetname, acc.ISZoneName, acc.ISCIDR, lbname, name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISSecurityGroupTargetExists("ibm_is_security_group_target.testacc_security_group_target", &securityGroup),
					resource.TestCheckResourceAttr(
						"ibm_is_security_group_target.testacc_security_group_target", "name", lbname),
				),
			},
		},
	})
}

func testAccCheckIBMISSecurityGroupTargetDestroy(s *terraform.State) error {

	sess, err := acc.TestAccProvider.Meta().(conns.ClientSession).VpcV1API()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_is_security_group_target" {
			continue
		}
		parts, err := flex.IdParts(rs.Primary.ID)
		if err != nil {
			return err
		}
		securityGroupID := parts[0]
		targetID := parts[1]

		deleteSecurityGroupTargetBindingOptions := &vpcv1.DeleteSecurityGroupTargetBindingOptions{
			SecurityGroupID: &securityGroupID,
			ID:              &targetID,
		}

		response, err := sess.DeleteSecurityGroupTargetBinding(deleteSecurityGroupTargetBindingOptions)
		if err == nil {
			return fmt.Errorf("Security Group Targets still exists: %v", response)
		}
	}
	return nil
}

func testAccCheckIBMISSecurityGroupTargetExists(n string, securityGroupID *string) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("[ERROR] No Security Group Target ID is set")
		}

		sess, err := acc.TestAccProvider.Meta().(conns.ClientSession).VpcV1API()
		if err != nil {
			return err
		}

		parts, err := flex.IdParts(rs.Primary.ID)
		if err != nil {
			return err
		}
		securityGroupId := parts[0]
		targetID := parts[1]

		getSecurityGroupTargetOptions := &vpcv1.GetSecurityGroupTargetOptions{
			SecurityGroupID: &securityGroupId,
			ID:              &targetID,
		}

		_, response, err := sess.GetSecurityGroupTarget(getSecurityGroupTargetOptions)
		if err != nil {
			if response != nil && response.StatusCode == 404 {
				*securityGroupID = ""
				return nil
			}
			return fmt.Errorf("[ERROR] Error getting Security Group Target : %s\n%s", err, response)
		}

		*securityGroupID = fmt.Sprintf("%s/%s", securityGroupId, targetID)
		return nil
	}
}

func testAccCheckIBMISsecurityGroupTargetConfig(vpcname, subnetname, zoneName, cidr, lbname, name string) string {
	return fmt.Sprintf(`
resource "ibm_is_vpc" "testacc_vpc" {
    name = "%s"
}

resource "ibm_is_subnet" "testacc_subnet" {
    name = "%s"
    vpc = ibm_is_vpc.testacc_vpc.id
    zone = "%s"
    ipv4_cidr_block = "%s"
}

resource "ibm_is_security_group" "testacc_security_group_one" {
    name = "%s"
    vpc = "${ibm_is_vpc.testacc_vpc.id}"
}

resource "ibm_is_lb" "testacc_LB" {
    name = "%s"
    subnets = [ibm_is_subnet.testacc_subnet.id]
}

resource "ibm_is_security_group_target" "testacc_security_group_target" {
    security_group = ibm_is_security_group.testacc_security_group_one.id
    target = ibm_is_lb.testacc_LB.id
  }`, vpcname, subnetname, zoneName, cidr, name, lbname)

}

func TestAccIBMISSecurityGroupTarget_UpdatePendingRetry(t *testing.T) {
	var securityGroup1, securityGroup2 string

	vpcname := fmt.Sprintf("tfsg-vpc-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tfsg-subnet-%d", acctest.RandIntRange(10, 100))
	lbname := fmt.Sprintf("tfsg-lb-%d", acctest.RandIntRange(10, 100))
	sg1name := fmt.Sprintf("tfsg-one-%d", acctest.RandIntRange(10, 100))
	sg2name := fmt.Sprintf("tfsg-two-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISSecurityGroupTargetDestroy,
		Steps: []resource.TestStep{
			// Create first security group target binding
			{
				Config: testAccCheckIBMISsecurityGroupTargetConfig(vpcname, subnetname, acc.ISZoneName, acc.ISCIDR, lbname, sg1name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISSecurityGroupTargetExists("ibm_is_security_group_target.testacc_security_group_target", &securityGroup1),
					resource.TestCheckResourceAttr(
						"ibm_is_security_group_target.testacc_security_group_target", "name", lbname),
					resource.TestCheckResourceAttrSet(
						"ibm_is_security_group_target.testacc_security_group_target", "security_group"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_security_group_target.testacc_security_group_target", "target"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_security_group_target.testacc_security_group_target", "crn"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_security_group_target.testacc_security_group_target", "resource_type"),
				),
			},
			// Create second security group target binding for the same LB - should trigger UPDATE_PENDING condition
			{
				Config: testAccCheckIBMISsecurityGroupTargetDualConfig(vpcname, subnetname, acc.ISZoneName, acc.ISCIDR, lbname, sg1name, sg2name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISSecurityGroupTargetExists("ibm_is_security_group_target.testacc_security_group_target", &securityGroup1),
					testAccCheckIBMISSecurityGroupTargetExists("ibm_is_security_group_target.testacc_security_group_target2", &securityGroup2),
					resource.TestCheckResourceAttr(
						"ibm_is_security_group_target.testacc_security_group_target", "name", lbname),
					resource.TestCheckResourceAttr(
						"ibm_is_security_group_target.testacc_security_group_target2", "name", lbname),
					resource.TestCheckResourceAttrSet(
						"ibm_is_security_group_target.testacc_security_group_target2", "security_group"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_security_group_target.testacc_security_group_target2", "target"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_security_group_target.testacc_security_group_target2", "crn"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_security_group_target.testacc_security_group_target2", "resource_type"),
				),
			},
		},
	})
}

// Config for dual security groups attached to the same load balancer
func testAccCheckIBMISsecurityGroupTargetDualConfig(vpcname, subnetname, zoneName, cidr, lbname, sg1name, sg2name string) string {
	return fmt.Sprintf(`
		resource "ibm_is_vpc" "testacc_vpc" {
			name = "%s"
		}

		resource "ibm_is_subnet" "testacc_subnet" {
			name = "%s"
			vpc = ibm_is_vpc.testacc_vpc.id
			zone = "%s"
			ipv4_cidr_block = "%s"
		}

		resource "ibm_is_security_group" "testacc_security_group_one" {
			name = "%s"
			vpc = ibm_is_vpc.testacc_vpc.id
		}

		resource "ibm_is_security_group" "testacc_security_group_two" {
			name = "%s"
			vpc = ibm_is_vpc.testacc_vpc.id
		}

		resource "ibm_is_lb" "testacc_LB" {
			name = "%s"
			subnets = [ibm_is_subnet.testacc_subnet.id]
		}

		resource "ibm_is_security_group_target" "testacc_security_group_target" {
			security_group = ibm_is_security_group.testacc_security_group_one.id
			target = ibm_is_lb.testacc_LB.id
		}

		resource "ibm_is_security_group_target" "testacc_security_group_target2" {
			security_group = ibm_is_security_group.testacc_security_group_two.id
			target = ibm_is_lb.testacc_LB.id
		}`,
		vpcname, subnetname, zoneName, cidr, sg1name, sg2name, lbname)
}
