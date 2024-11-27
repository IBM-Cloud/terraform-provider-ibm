// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"errors"
	"fmt"
	"strings"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"

	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccIBMISSecurityGroup_basic(t *testing.T) {
	var securityGroup string

	vpcname := fmt.Sprintf("tfsg-vpc-%d", acctest.RandIntRange(10, 100))
	name1 := fmt.Sprintf("tfsg-createname-%d", acctest.RandIntRange(10, 100))
	//name2 := fmt.Sprintf("tfsg-updatename-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
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
func TestAccIBMISSecurityGroup_wait(t *testing.T) {
	var securityGroup string

	vpcname := fmt.Sprintf("tfsg-vpc-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tfsg-subnet-%d", acctest.RandIntRange(10, 100))
	sshname := fmt.Sprintf("tfsg-ssh-%d", acctest.RandIntRange(10, 100))
	vsiname := fmt.Sprintf("tfsg-vsi-%d", acctest.RandIntRange(10, 100))
	bmname := fmt.Sprintf("tfsg-bm-%d", acctest.RandIntRange(10, 100))
	name1 := fmt.Sprintf("tfsg-createname-%d", acctest.RandIntRange(10, 100))
	publickey := strings.TrimSpace(`
	ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCKVmnMOlHKcZK8tpt3MP1lqOLAcqcJzhsvJcjscgVERRN7/9484SOBJ3HSKxxNG5JN8owAjy5f9yYwcUg+JaUVuytn5Pv3aeYROHGGg+5G346xaq3DAwX6Y5ykr2fvjObgncQBnuU5KHWCECO/4h8uWuwh/kfniXPVjFToc+gnkqA+3RKpAecZhFXwfalQ9mMuYGFxn+fwn8cYEApsJbsEmb0iJwPiZ5hjFC8wREuiTlhPHDgkBLOiycd20op2nXzDbHfCHInquEe/gYxEitALONxm0swBOwJZwlTDOB7C6y2dzlrtxr1L59m7pCkWI4EtTRLvleehBoj3u7jB4usR
	`)
	//name2 := fmt.Sprintf("tfsg-updatename-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISSecurityGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISsecurityGroupWaitConfig(name1, vpcname, subnetname, sshname, publickey, vsiname, bmname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISSecurityGroupExists("ibm_is_security_group.testacc_security_group", securityGroup),
					resource.TestCheckResourceAttr(
						"ibm_is_security_group.testacc_security_group", "name", name1),
					resource.TestCheckResourceAttr(
						"ibm_is_security_group.testacc_security_group", "tags.#", "2"),
					resource.TestCheckResourceAttr(
						"ibm_is_vpc.testacc_vpc", "name", vpcname),
					resource.TestCheckResourceAttr(
						"ibm_is_subnet.testacc_subnet", "name", subnetname),
					resource.TestCheckResourceAttr(
						"ibm_is_ssh_key.testacc_sshkey", "name", sshname),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "name", vsiname),
					resource.TestCheckResourceAttr(
						"ibm_is_bare_metal_server.testacc_bms", "name", bmname),
				),
			},
		},
	})
}

func testAccCheckIBMISSecurityGroupDestroy(s *terraform.State) error {
	sess, _ := acc.TestAccProvider.Meta().(conns.ClientSession).VpcV1API()
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

		sess, _ := acc.TestAccProvider.Meta().(conns.ClientSession).VpcV1API()
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
func testAccCheckIBMISsecurityGroupWaitConfig(name, vpcname, subnetname, sshname, publicKey, vsiname, bmname string) string {
	return fmt.Sprintf(`
resource "ibm_is_vpc" "testacc_vpc" {
	name = "%s"
}

resource ibm_is_subnet testacc_subnet {
	name = "%s"
	zone = "%s"
	vpc = "${ibm_is_vpc.testacc_vpc.id}"
	total_ipv4_address_count = 16
}

resource "ibm_is_security_group" "testacc_security_group" {
	name = "%s"
	vpc = "${ibm_is_vpc.testacc_vpc.id}"
	tags = ["tag1", "tag2"]
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
		subnet     = ibm_is_subnet.testacc_subnet.id
		security_groups = [ibm_is_security_group.testacc_security_group.id]
	}
	vpc  = ibm_is_vpc.testacc_vpc.id
	zone = "%s"
	keys = [ibm_is_ssh_key.testacc_sshkey.id]
	network_interfaces {
		subnet = ibm_is_subnet.testacc_subnet.id
		name   = "eth12"
		security_groups = [ibm_is_security_group.testacc_security_group.id]
	}
}

resource "ibm_is_bare_metal_server" "testacc_bms" {
	profile 			= "%s"
	name 				= "%s"
	image 				= "%s"
	zone 				= "%s"
	keys 				= [ibm_is_ssh_key.testacc_sshkey.id]
	primary_network_interface {
		subnet     		= ibm_is_subnet.testacc_subnet.id
		security_groups = [ibm_is_security_group.testacc_security_group.id]
	}
	vpc 				= ibm_is_vpc.testacc_vpc.id
}
`, vpcname, subnetname, acc.ISZoneName, name, sshname, publicKey, vsiname, acc.IsImage, acc.InstanceProfileName, acc.ISZoneName, acc.IsBareMetalServerProfileName, bmname, acc.IsBareMetalServerImage, acc.ISZoneName)

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
