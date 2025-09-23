// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package schematics_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"
)

func TestAccIbmSchematicsAgentsDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmSchematicsAgentsDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_schematics_agents.schematics_agents_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_agents.schematics_agents_instance", "offset"),
				),
			},
		},
	})
}

func testAccCheckIbmSchematicsAgentsDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		data "ibm_schematics_agents" "schematics_agents_instance" {
		}
	`)
}
