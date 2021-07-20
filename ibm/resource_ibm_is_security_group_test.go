// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"errors"
	"fmt"
	"testing"

	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccIBMISSecurityGroup_basic(t *testing.T) {
	var securityGroup string

	vpcname := fmt.Sprintf("tfsg-vpc-%d", acctest.RandIntRange(10, 100))
	name1 := fmt.Sprintf("tfsg-createname-%d", acctest.RandIntRange(10, 100))
	//name2 := fmt.Sprintf("tfsg-updatename-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMISSecurityGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISsecurityGroupConfig(vpcname, name1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISSecurityGroupExists("ibm_is_security_group.testacc_security_group", securityGroup),
					resource.TestCheckResourceAttr(
						"ibm_is_security_group.testacc_security_group", "name", name1),
					resource.TestCheckResourceAttr(
						"ibm_is_security_group.testacc_security_group", "tags.#", "2"),
				),
			},
			{
				Config: testAccCheckIBMISsecurityGroupConfigUpdate(vpcname, name1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISSecurityGroupExists("ibm_is_security_group.testacc_security_group", securityGroup),
					resource.TestCheckResourceAttr(
						"ibm_is_security_group.testacc_security_group", "name", name1),
					resource.TestCheckResourceAttr(
						"ibm_is_security_group.testacc_security_group", "tags.#", "1"),
				),
			},
		},
	})
}

func testAccCheckIBMISSecurityGroupDestroy(s *terraform.State) error {
	sess, _ := testAccProvider.Meta().(ClientSession).VpcV1API()
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_is_security_group" {
			continue
		}

		getsgoptions := &vpcv1.GetSecurityGroupOptions{
			ID: &rs.Primary.ID,
		}
		_, _, err := sess.GetSecurityGroup(getsgoptions)

		if err == nil {
			return fmt.Errorf("securitygroup still exists: %s", rs.Primary.ID)
		}
	}

	return nil
}

func testAccCheckIBMISSecurityGroupExists(n, securityGroupID string) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return errors.New("No Record ID is set")
		}

		sess, _ := testAccProvider.Meta().(ClientSession).VpcV1API()
		getsgoptions := &vpcv1.GetSecurityGroupOptions{
			ID: &rs.Primary.ID,
		}
		foundsecurityGroup, _, err := sess.GetSecurityGroup(getsgoptions)
		if err != nil {
			return err
		}
		securityGroupID = *foundsecurityGroup.ID
		return nil
	}
}

func testAccCheckIBMISsecurityGroupConfig(vpcname, name string) string {
	return fmt.Sprintf(`
resource "ibm_is_vpc" "testacc_vpc" {
	name = "%s"
}

resource "ibm_is_security_group" "testacc_security_group" {
	name = "%s"
	vpc = "${ibm_is_vpc.testacc_vpc.id}"
	tags = ["Tag1", "tag2"]
}`, vpcname, name)

}

func testAccCheckIBMISsecurityGroupConfigUpdate(vpcname, name string) string {
	return fmt.Sprintf(`
resource "ibm_is_vpc" "testacc_vpc" {
	name = "%s"
}

resource "ibm_is_security_group" "testacc_security_group" {
	name = "%s"
	vpc = "${ibm_is_vpc.testacc_vpc.id}"
	tags = ["tag1"]
}`, vpcname, name)

}
