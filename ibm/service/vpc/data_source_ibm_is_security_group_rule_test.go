// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMIsSecurityGroupRuleDataSourceBasic(t *testing.T) {
	vpcname := fmt.Sprintf("tfsecgrprl-vpc-%d", acctest.RandIntRange(10, 100))
	sgname := fmt.Sprintf("tfsecgrprl-name-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsSecurityGroupRuleDataSourceConfigBasic(vpcname, sgname),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_security_group_rule.example", "security_group"),
					resource.TestCheckResourceAttrSet("data.ibm_is_security_group_rule.example", "direction"),
					resource.TestCheckResourceAttrSet("data.ibm_is_security_group_rule.example", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_security_group_rule.example", "ip_version"),
					resource.TestCheckResourceAttrSet("data.ibm_is_security_group_rule.example", "protocol"),
					resource.TestCheckResourceAttrSet("data.ibm_is_security_group_rule.example", "remote.#"),
				),
			},
		},
	})
}

func testAccCheckIBMIsSecurityGroupRuleDataSourceConfigBasic(vpcname, sgname string) string {
	return fmt.Sprintf(`

	resource "ibm_is_vpc" "example" {
        name = "%s"
      }
      
      resource "ibm_is_security_group" "example" {
        name = "%s"
        vpc  = ibm_is_vpc.example.id
      }
      
      resource "ibm_is_security_group_rule" "example" {
        depends_on = [
            ibm_is_security_group.example,
        ]
        group     = ibm_is_security_group.example.id
        direction = "inbound"
        remote    = "127.0.0.1"
        udp {
          port_min = 805
          port_max = 807
        }
      }

      data "ibm_is_security_group_rule" "example" {
        depends_on = [
            ibm_is_security_group_rule.example,
        ]
          security_group_rule = ibm_is_security_group_rule.example.rule_id
          security_group = ibm_is_security_group.example.id
      }
	`, vpcname, sgname)
}
