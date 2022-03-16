// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMIsSecurityGroupsDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsSecurityGroupsDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_security_groups.example", "security_groups.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_security_groups.example", "security_groups.0.created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_is_security_groups.example", "security_groups.0.crn"),
					resource.TestCheckResourceAttrSet("data.ibm_is_security_groups.example", "security_groups.0.href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_security_groups.example", "security_groups.0.name"),
					resource.TestCheckResourceAttrSet("data.ibm_is_security_groups.example", "security_groups.0.resource_group.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_security_groups.example", "security_groups.0.rules.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_security_groups.example", "security_groups.0.targets.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_security_groups.example", "security_groups.0.vpc.#"),
				),
			},
		},
	})
}

func testAccCheckIBMIsSecurityGroupsDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		data "ibm_is_security_groups" "example" {
		}
	`)
}
