// Copyright IBM Corp. 2017, 2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"errors"
	"fmt"
	"strings"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccIBMISVPCDefaultSecurityGroup_basic(t *testing.T) {
	var securityGroupID string
	name := fmt.Sprintf("tfvpc-defaultsg-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISVPCDefaultSecurityGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISVPCDefaultSecurityGroupConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISVPCDefaultSecurityGroupExists("ibm_is_vpc_default_security_group.test_default_sg", securityGroupID),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpc_default_security_group.test_default_sg", "name"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpc_default_security_group.test_default_sg", "default_security_group"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpc_default_security_group.test_default_sg", "crn"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpc_default_security_group.test_default_sg", "resource_group.#"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpc_default_security_group.test_default_sg", "targets.#"),
				),
			},
		},
	})
}

func TestAccIBMISVPCDefaultSecurityGroup_name_update(t *testing.T) {
	var securityGroupID string
	name := fmt.Sprintf("tfvpc-defaultsg-%d", acctest.RandIntRange(10, 100))
	updatedName := fmt.Sprintf("tfvpc-defaultsg-updated-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISVPCDefaultSecurityGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISVPCDefaultSecurityGroupConfigWithName(name, name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISVPCDefaultSecurityGroupExists("ibm_is_vpc_default_security_group.test_default_sg", securityGroupID),
					resource.TestCheckResourceAttr(
						"ibm_is_vpc_default_security_group.test_default_sg", "name", name),
				),
			},
			{
				Config: testAccCheckIBMISVPCDefaultSecurityGroupConfigWithName(name, updatedName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISVPCDefaultSecurityGroupExists("ibm_is_vpc_default_security_group.test_default_sg", securityGroupID),
					resource.TestCheckResourceAttr(
						"ibm_is_vpc_default_security_group.test_default_sg", "name", updatedName),
				),
			},
		},
	})
}

func TestAccIBMISVPCDefaultSecurityGroup_tags(t *testing.T) {
	var securityGroupID string
	name := fmt.Sprintf("tfvpc-defaultsg-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISVPCDefaultSecurityGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISVPCDefaultSecurityGroupConfigWithTags(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISVPCDefaultSecurityGroupExists("ibm_is_vpc_default_security_group.test_default_sg", securityGroupID),
					resource.TestCheckResourceAttr(
						"ibm_is_vpc_default_security_group.test_default_sg", "tags.#", "2"),
					resource.TestCheckResourceAttr(
						"ibm_is_vpc_default_security_group.test_default_sg", "tags.0", "env:test"),
					resource.TestCheckResourceAttr(
						"ibm_is_vpc_default_security_group.test_default_sg", "tags.1", "team:dev"),
				),
			},
			{
				Config: testAccCheckIBMISVPCDefaultSecurityGroupConfigWithUpdatedTags(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISVPCDefaultSecurityGroupExists("ibm_is_vpc_default_security_group.test_default_sg", securityGroupID),
					resource.TestCheckResourceAttr(
						"ibm_is_vpc_default_security_group.test_default_sg", "tags.#", "3"),
					resource.TestCheckResourceAttr(
						"ibm_is_vpc_default_security_group.test_default_sg", "tags.0", "env:production"),
					resource.TestCheckResourceAttr(
						"ibm_is_vpc_default_security_group.test_default_sg", "tags.1", "team:ops"),
					resource.TestCheckResourceAttr(
						"ibm_is_vpc_default_security_group.test_default_sg", "tags.2", "project:web"),
				),
			},
		},
	})
}

func TestAccIBMISVPCDefaultSecurityGroup_access_tags(t *testing.T) {
	var securityGroupID string
	name := fmt.Sprintf("tfvpc-defaultsg-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISVPCDefaultSecurityGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISVPCDefaultSecurityGroupConfigWithAccessTags(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISVPCDefaultSecurityGroupExists("ibm_is_vpc_default_security_group.test_default_sg", securityGroupID),
					resource.TestCheckResourceAttr(
						"ibm_is_vpc_default_security_group.test_default_sg", "access_tags.#", "1"),
				),
			},
			{
				ResourceName:      "ibm_is_vpc_default_security_group.test_default_sg",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccIBMISVPCDefaultSecurityGroup_with_targets(t *testing.T) {
	var securityGroupID string
	vpcname := fmt.Sprintf("tfvpc-defaultsg-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tfsubnet-defaultsg-%d", acctest.RandIntRange(10, 100))
	sshname := fmt.Sprintf("tfssh-defaultsg-%d", acctest.RandIntRange(10, 100))
	vsiname := fmt.Sprintf("tfvsi-defaultsg-%d", acctest.RandIntRange(10, 100))
	publickey := strings.TrimSpace(`
	ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCKVmnMOlHKcZK8tpt3MP1lqOLAcqcJzhsvJcjscgVERRN7/9484SOBJ3HSKxxNG5JN8owAjy5f9yYwcUg+JaUVuytn5Pv3aeYROHGGg+5G346xaq3DAwX6Y5ykr2fvjObgncQBnuU5KHWCECO/4h8uWuwh/kfniXPVjFToc+gnkqA+3RKpAecZhFXwfalQ9mMuYGFxn+fwn8cYEApsJbsEmb0iJwPiZ5hjFC8wREuiTlhPHDgkBLOiycd20op2nXzDbHfCHInquEe/gYxEitALONxm0swBOwJZwlTDOB7C6y2dzlrtxr1L59m7pCkWI4EtTRLvleehBoj3u7jB4usR
	`)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISVPCDefaultSecurityGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISVPCDefaultSecurityGroupConfigWithTargets(vpcname, subnetname, sshname, publickey, vsiname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISVPCDefaultSecurityGroupExists("ibm_is_vpc_default_security_group.test_default_sg", securityGroupID),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpc_default_security_group.test_default_sg", "targets.#"),
					resource.TestCheckResourceAttr(
						"ibm_is_vpc.testacc_vpc", "name", vpcname),
					resource.TestCheckResourceAttr(
						"ibm_is_subnet.testacc_subnet", "name", subnetname),
					resource.TestCheckResourceAttr(
						"ibm_is_ssh_key.testacc_sshkey", "name", sshname),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "name", vsiname),
				),
			},
		},
	})
}

func testAccCheckIBMISVPCDefaultSecurityGroupDestroy(s *terraform.State) error {
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_is_vpc_default_security_group" {
			continue
		}

		parts := strings.Split(rs.Primary.ID, "/")
		if len(parts) != 2 {
			return fmt.Errorf("Invalid ID format for default security group")
		}

		sess, err := vpcClient(acc.TestAccProvider.Meta())
		if err != nil {
			return err
		}

		securityGroupID := parts[1]
		getSecurityGroupOptions := sess.NewGetSecurityGroupOptions(securityGroupID)
		foundSecurityGroup, _, err := sess.GetSecurityGroup(getSecurityGroupOptions)
		if err == nil {
			// Default security group should still exist, but we check if it's properly configured
			if foundSecurityGroup != nil {
				// This is expected - default security group should still exist
				continue
			}
			return fmt.Errorf("Default security group still exists but not found properly: %s", rs.Primary.ID)
		}
	}
	return nil
}

func testAccCheckIBMISVPCDefaultSecurityGroupExists(n, securityGroupID string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return errors.New("No Record ID is set")
		}

		parts, err := flex.IdParts(rs.Primary.ID)
		if err != nil {
			return err
		}

		if len(parts) != 2 {
			return fmt.Errorf("Invalid ID format: %s", rs.Primary.ID)
		}

		securityGroupID = parts[1]
		sess, err := vpcClient(acc.TestAccProvider.Meta())
		if err != nil {
			return err
		}

		getSecurityGroupOptions := sess.NewGetSecurityGroupOptions(securityGroupID)
		foundSecurityGroup, response, err := sess.GetSecurityGroup(getSecurityGroupOptions)
		if err != nil {
			return fmt.Errorf("[ERROR] Error Getting VPC Default Security Group: %s\n%s", err, response)
		}

		securityGroupID = *foundSecurityGroup.ID
		return nil
	}
}

func testAccCheckIBMISVPCDefaultSecurityGroupConfig(name string) string {
	return fmt.Sprintf(`
resource "ibm_is_vpc" "testacc_vpc" {
	name = "%s"
}

resource "ibm_is_vpc_default_security_group" "test_default_sg" {
	vpc = ibm_is_vpc.testacc_vpc.id
}`, name)
}

func testAccCheckIBMISVPCDefaultSecurityGroupConfigWithName(vpcName, sgName string) string {
	return fmt.Sprintf(`
resource "ibm_is_vpc" "testacc_vpc" {
	name = "%s"
}

resource "ibm_is_vpc_default_security_group" "test_default_sg" {
	vpc  = ibm_is_vpc.testacc_vpc.id
	name = "%s"
}`, vpcName, sgName)
}

func testAccCheckIBMISVPCDefaultSecurityGroupConfigWithTags(name string) string {
	return fmt.Sprintf(`
resource "ibm_is_vpc" "testacc_vpc" {
	name = "%s"
}

resource "ibm_is_vpc_default_security_group" "test_default_sg" {
	vpc  = ibm_is_vpc.testacc_vpc.id
	name = "%s-default-sg"
	tags = ["env:test", "team:dev"]
}`, name, name)
}

func testAccCheckIBMISVPCDefaultSecurityGroupConfigWithUpdatedTags(name string) string {
	return fmt.Sprintf(`
resource "ibm_is_vpc" "testacc_vpc" {
	name = "%s"
}

resource "ibm_is_vpc_default_security_group" "test_default_sg" {
	vpc  = ibm_is_vpc.testacc_vpc.id
	name = "%s-default-sg"
	tags = ["env:production", "team:ops", "project:web"]
}`, name, name)
}

func testAccCheckIBMISVPCDefaultSecurityGroupConfigWithAccessTags(name string) string {
	return fmt.Sprintf(`
resource "ibm_is_vpc" "testacc_vpc" {
	name = "%s"
}

resource "ibm_is_vpc_default_security_group" "test_default_sg" {
	vpc         = ibm_is_vpc.testacc_vpc.id
	name        = "%s-default-sg"
	access_tags = ["project:test"]
}`, name, name)
}

func testAccCheckIBMISVPCDefaultSecurityGroupConfigWithTargets(vpcname, subnetname, sshname, publicKey, vsiname string) string {
	return fmt.Sprintf(`
resource "ibm_is_vpc" "testacc_vpc" {
	name = "%s"
}

resource "ibm_is_subnet" "testacc_subnet" {
	name = "%s"
	zone = "%s"
	vpc = ibm_is_vpc.testacc_vpc.id
	total_ipv4_address_count = 16
}

resource "ibm_is_vpc_default_security_group" "test_default_sg" {
	vpc  = ibm_is_vpc.testacc_vpc.id
	name = "%s-default-sg"
	tags = ["test", "targets"]
}

resource "ibm_is_ssh_key" "testacc_sshkey" {
	name       = "%s"
	public_key = "%s"
}
  
resource "ibm_is_instance" "testacc_instance" {
	name    = "%s"
	image   = "%s"
	profile = "%s"
	primary_network_interface {
		subnet = ibm_is_subnet.testacc_subnet.id
		security_groups = [ibm_is_vpc_default_security_group.test_default_sg.default_security_group]
	}
	vpc  = ibm_is_vpc.testacc_vpc.id
	zone = "%s"
	keys = [ibm_is_ssh_key.testacc_sshkey.id]
}`, vpcname, subnetname, acc.ISZoneName, vpcname, sshname, publicKey, vsiname, acc.IsImage, acc.InstanceProfileName, acc.ISZoneName)
}
