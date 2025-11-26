// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package schematics_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIbmSchematicsPoliciesDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmSchematicsPoliciesDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_schematics_policies.schematics_policies_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_policies.schematics_policies_instance", "offset"),
				),
			},
		},
	})
}

func testAccCheckIbmSchematicsPoliciesDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		data "ibm_schematics_policies" "schematics_policies_instance" {
		}
	`)
}
