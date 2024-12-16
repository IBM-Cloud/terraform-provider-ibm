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

func TestAccIBMISSecurityGroupTargetDataSource_vni(t *testing.T) {
	var securityGroup string
	terraformTag := "data.ibm_is_security_group_target.testacc_security_group_target"
	vpcname := fmt.Sprintf("tfsg-vpc-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tfsg-subnet-%d", acctest.RandIntRange(10, 100))
	lbname := fmt.Sprintf("tfsg-lb-%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tfsg-one-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISSecurityGroupDataSourceTargetDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISsecurityGroupTargetDataSourceVniConfig(vpcname, subnetname, acc.ISZoneName, acc.ISCIDR, lbname, name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISSecurityGroupDataSourceTargetExists("ibm_is_security_group_target.testacc_security_group_target", &securityGroup),
					resource.TestCheckResourceAttrSet(
						terraformTag, "crn"),
					resource.TestCheckResourceAttr(
						terraformTag, "resource_type", "virtual_network_interface"),
					resource.TestCheckResourceAttrSet(
						terraformTag, "target"),
				),
			},
		},
	})
}

func testAccCheckIBMISSecurityGroupDataSourceTargetDestroy(s *terraform.State) error {

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

func testAccCheckIBMISSecurityGroupDataSourceTargetExists(n string, securityGroupID *string) resource.TestCheckFunc {

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
func testAccCheckIBMISsecurityGroupTargetDataSourceVniConfig(vpcname, subnetname, zoneName, cidr, vniname, name string) string {
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

	resource "ibm_is_virtual_network_interface" "testacc_vni"{
		name 						= "%s"
		allow_ip_spoofing 			= false
		enable_infrastructure_nat 	= true
		primary_ip {
			auto_delete 	= false
			address 		= cidrhost(cidrsubnet(ibm_is_subnet.testacc_subnet.ipv4_cidr_block, 4, 6), 0)
		}
		subnet = ibm_is_subnet.testacc_subnet.id
	}

	resource "ibm_is_security_group_target" "testacc_security_group_target" {
	    security_group = ibm_is_security_group.testacc_security_group_one.id
	    target = ibm_is_virtual_network_interface.testacc_vni.id
	  }

	data "ibm_is_security_group_target" "testacc_security_group_target" {
		security_group = ibm_is_security_group_target.testacc_security_group_target.security_group
		name = ibm_is_security_group_target.testacc_security_group_target.name
	}

	`, vpcname, subnetname, zoneName, cidr, name, vniname)
}
