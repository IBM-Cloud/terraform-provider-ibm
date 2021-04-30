// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMSchematicsOutputDataSourceBasic(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMSchematicsOutputDataSourceConfigBasic(workspaceID, templateID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ibm_schematics_output.schematics_output", "workspace_id", workspaceID),
				),
			},
		},
	})
}

func testAccCheckIBMSchematicsOutputDataSourceConfigBasic(wID string, templateID string) string {
	return fmt.Sprintf(`
		  data "ibm_schematics_output" "schematics_output" {
			workspace_id = "%s"
			template_id = "%s"
		  }
	  `, workspaceID, templateID)
}
