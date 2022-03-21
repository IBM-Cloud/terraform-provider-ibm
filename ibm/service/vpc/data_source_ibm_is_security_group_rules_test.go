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

func TestAccIBMIsSecurityGroupRulesDataSourceBasic(t *testing.T) {
	vpcname := fmt.Sprintf("tfsubnet-vpc-%d", acctest.RandIntRange(10, 100))
	sgname := fmt.Sprintf("tfsubnet-name-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsSecurityGroupRulesDataSourceConfigBasic(vpcname, sgname),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_security_group_rules.example", "security_group"),
					resource.TestCheckResourceAttrSet("data.ibm_is_security_group_rules.example", "rules.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_security_group_rules.example", "rules.0.direction"),
					resource.TestCheckResourceAttrSet("data.ibm_is_security_group_rules.example", "rules.0.href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_security_group_rules.example", "rules.0.ip_version"),
					resource.TestCheckResourceAttrSet("data.ibm_is_security_group_rules.example", "rules.0.protocol"),
					resource.TestCheckResourceAttrSet("data.ibm_is_security_group_rules.example", "rules.0.remote.#"),
				),
			},
		},
	})
}

func testAccCheckIBMIsSecurityGroupRulesDataSourceConfigBasic(vpcname, sgname string) string {
	return fmt.Sprintf(`

	resource "ibm_is_vpc" "example" {
		name = "%s"
	  }
	  
	  resource "ibm_is_security_group" "example" {
		name = "%s"
		vpc  = ibm_is_vpc.example.id
		depends_on = [
			ibm_is_vpc.example,
		]
	  }
	  
	  resource "ibm_is_security_group_rule" "example" {
		group     = ibm_is_security_group.example.id
		direction = "outbound"
		remote    = "127.0.0.1"
		tcp {
		  port_min = 8080
		  port_max = 8080
		}
		depends_on = [
			ibm_is_security_group.example,
		]
	  }
		data "ibm_is_security_group_rules" "example" {
			depends_on = [
				ibm_is_security_group_rule.example,
		]
			security_group = ibm_is_security_group.example.id
		}
	`, vpcname, sgname)
}
