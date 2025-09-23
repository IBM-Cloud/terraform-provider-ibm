// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package classicinfrastructure_test

import (
	"fmt"
	"strconv"
	"testing"

	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/conns"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/softlayer/softlayer-go/services"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
)

func TestAccIBMSecurityGroupRule_basic(t *testing.T) {
	name1 := fmt.Sprintf("terraformsguat_create_step_name_%d", acctest.RandIntRange(10, 100))
	desc1 := fmt.Sprintf("terraformsguat_create_step_desc_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMSecurityGroupRuleDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMSecurityGroupRuleConfig(name1, desc1),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_security_group.testacc_sg", "name", name1),
					resource.TestCheckResourceAttr(
						"ibm_security_group.testacc_sg", "description", desc1),
					resource.TestCheckResourceAttr(
						"ibm_security_group_rule.testacc_sg_rule", "ether_type", "IPv4"),
					resource.TestCheckResourceAttr(
						"ibm_security_group_rule.testacc_sg_rule", "direction", "ingress"),
					resource.TestCheckResourceAttr(
						"ibm_security_group_rule.testacc_sg_rule", "protocol", "udp"),
					resource.TestCheckResourceAttr(
						"ibm_security_group_rule.testacc_sg_rule", "port_range_min", "80"),
					resource.TestCheckResourceAttr(
						"ibm_security_group_rule.testacc_sg_rule", "port_range_max", "8080"),
				),
			},

			{
				Config: testAccCheckIBMSecurityGroupRuleUpdateConfig(name1, desc1),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_security_group_rule.testacc_sg_rule", "ether_type", "IPv4"),
					resource.TestCheckResourceAttr(
						"ibm_security_group_rule.testacc_sg_rule", "direction", "ingress"),
					resource.TestCheckResourceAttr(
						"ibm_security_group_rule.testacc_sg_rule", "protocol", "tcp"),
					resource.TestCheckResourceAttr(
						"ibm_security_group_rule.testacc_sg_rule", "port_range_min", "70"),
					resource.TestCheckResourceAttr(
						"ibm_security_group_rule.testacc_sg_rule", "port_range_max", "8081"),
				),
			},
		},
	})
}

func TestAccIBMSecurityGroupRule_with_remote_group(t *testing.T) {
	name1 := fmt.Sprintf("terraformsguat_create_step_name_%d", acctest.RandIntRange(10, 100))
	desc1 := fmt.Sprintf("terraformsguat_create_step_desc_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMSecurityGroupRuleDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMSecurityGroupRuleConfigWithRemoteGroup(name1, desc1),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_security_group.testacc_sg", "name", name1),
					resource.TestCheckResourceAttr(
						"ibm_security_group.testacc_sg", "description", desc1),
					resource.TestCheckResourceAttr(
						"ibm_security_group_rule.testacc_sg_rule", "remote_group_id", "72101"),
					resource.TestCheckResourceAttr(
						"ibm_security_group_rule.testacc_sg_rule", "direction", "ingress"),
					resource.TestCheckResourceAttr(
						"ibm_security_group_rule.testacc_sg_rule", "protocol", "udp"),
					resource.TestCheckResourceAttr(
						"ibm_security_group_rule.testacc_sg_rule", "port_range_min", "80"),
					resource.TestCheckResourceAttr(
						"ibm_security_group_rule.testacc_sg_rule", "port_range_max", "8080"),
				),
			},
		},
	})
}

func TestAccIBMSecurityGroupRule_with_remote_ip(t *testing.T) {
	name1 := fmt.Sprintf("terraformsguat_create_step_name_%d", acctest.RandIntRange(10, 100))
	desc1 := fmt.Sprintf("terraformsguat_create_step_desc_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMSecurityGroupRuleDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMSecurityGroupRuleConfigWithRemoteIP(name1, desc1),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_security_group.testacc_sg", "name", name1),
					resource.TestCheckResourceAttr(
						"ibm_security_group.testacc_sg", "description", desc1),
					resource.TestCheckResourceAttr(
						"ibm_security_group_rule.testacc_sg_rule", "remote_ip", "10.0.0.2"),
					resource.TestCheckResourceAttr(
						"ibm_security_group_rule.testacc_sg_rule", "direction", "ingress"),
					resource.TestCheckResourceAttr(
						"ibm_security_group_rule.testacc_sg_rule", "protocol", "udp"),
					resource.TestCheckResourceAttr(
						"ibm_security_group_rule.testacc_sg_rule", "port_range_min", "80"),
					resource.TestCheckResourceAttr(
						"ibm_security_group_rule.testacc_sg_rule", "port_range_max", "8080"),
				),
			},
		},
	})
}

func TestAccIBMSecurityGroupRule_with_cross_refernce_another_security_group(t *testing.T) {
	name1 := fmt.Sprintf("terraformsguat_create_step_name_%d", acctest.RandIntRange(10, 100))
	desc1 := fmt.Sprintf("terraformsguat_create_step_desc_%d", acctest.RandIntRange(10, 100))
	name2 := fmt.Sprintf("terraformsguat_create_step_name_%d", acctest.RandIntRange(10, 100))
	desc2 := fmt.Sprintf("terraformsguat_create_step_desc_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMSecurityGroupRuleDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMSecurityGroupRuleConfigWithSecurityGroupReference(name1, desc1, name2, desc2),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_security_group.testacc_sg1", "name", name1),
					resource.TestCheckResourceAttr(
						"ibm_security_group.testacc_sg1", "description", desc1),
					resource.TestCheckResourceAttr(
						"ibm_security_group.testacc_sg2", "name", name2),
					resource.TestCheckResourceAttr(
						"ibm_security_group.testacc_sg2", "description", desc2),
					resource.TestCheckResourceAttr(
						"ibm_security_group_rule.testacc_sg_rule", "protocol", "udp"),
				),
			},
		},
	})
}

func testAccCheckIBMSecurityGroupRuleDestroy(s *terraform.State) error {
	service := services.GetNetworkSecurityGroupService(acc.TestAccProvider.Meta().(conns.ClientSession).SoftLayerSession())

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_security_group_rule" {
			continue
		}

		sgID, _ := strconv.Atoi(rs.Primary.ID)

		// Try to find the key
		_, err := service.Id(sgID).GetObject()

		if err == nil {
			return fmt.Errorf("Security Group Rule %d still exists", sgID)
		}
	}

	return nil
}

func testAccCheckIBMSecurityGroupRuleConfig(name, description string) string {
	return fmt.Sprintf(`
resource "ibm_security_group" "testacc_sg" {
    name = "%s"
	description = "%s"
}
resource "ibm_security_group_rule" "testacc_sg_rule" {
    direction = "ingress"
	port_range_min = 80
	port_range_max = 8080
	protocol = "udp"
	security_group_id = "${ibm_security_group.testacc_sg.id}"
}
`, name, description)

}

func testAccCheckIBMSecurityGroupRuleUpdateConfig(name, description string) string {
	return fmt.Sprintf(`
resource "ibm_security_group" "testacc_sg" {
    name = "%s"
	description = "%s"
}
resource "ibm_security_group_rule" "testacc_sg_rule" {
    direction = "ingress"
	port_range_min = 70
	port_range_max = 8081
	protocol = "tcp"
	security_group_id = "${ibm_security_group.testacc_sg.id}"
}
`, name, description)

}

func testAccCheckIBMSecurityGroupRuleConfigWithRemoteGroup(name, description string) string {
	return fmt.Sprintf(`
resource "ibm_security_group" "testacc_sg" {
    name = "%s"
	description = "%s"
}
resource "ibm_security_group_rule" "testacc_sg_rule" {
    direction = "ingress"
	port_range_min = 80
	port_range_max = 8080
	protocol = "udp"
	remote_group_id = 72101
	security_group_id = "${ibm_security_group.testacc_sg.id}"
}
`, name, description)

}

func testAccCheckIBMSecurityGroupRuleConfigWithRemoteIP(name, description string) string {
	return fmt.Sprintf(`
resource "ibm_security_group" "testacc_sg" {
    name = "%s"
	description = "%s"
}
resource "ibm_security_group_rule" "testacc_sg_rule" {
    direction = "ingress"
	port_range_min = 80
	port_range_max = 8080
	protocol = "udp"
	remote_ip = "10.0.0.2"
	security_group_id = "${ibm_security_group.testacc_sg.id}"
}
`, name, description)

}

func testAccCheckIBMSecurityGroupRuleConfigWithSecurityGroupReference(name1, description1, name2, description2 string) string {
	return fmt.Sprintf(`
resource "ibm_security_group" "testacc_sg1" {
    name = "%s"
	description = "%s"
}
resource "ibm_security_group" "testacc_sg2" {
    name = "%s"
	description = "%s"
}

resource "ibm_security_group_rule" "testacc_sg_rule" {
    direction = "ingress"
	port_range_min = 80
	port_range_max = 8080
	protocol = "udp"
	remote_group_id = "${ibm_security_group.testacc_sg2.id}"
	security_group_id = "${ibm_security_group.testacc_sg1.id}"
}
`, name1, description1, name2, description2)

}
