// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package schematics_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccIBMSchematicsOutputDataSourceBasic(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMSchematicsOutputDataSourceConfigBasic(acc.WorkspaceID, acc.TemplateID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ibm_schematics_output.schematics_output", "workspace_id", acc.WorkspaceID),
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
	  `, acc.WorkspaceID, templateID)
}
