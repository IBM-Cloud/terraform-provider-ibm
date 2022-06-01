// Copyright IBM Corp. 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package cdtoolchain_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIBMCdToolchainToolGithubconsolidatedDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCdToolchainToolGithubconsolidatedDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_githubconsolidated.cd_toolchain_tool_githubconsolidated", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_githubconsolidated.cd_toolchain_tool_githubconsolidated", "toolchain_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_githubconsolidated.cd_toolchain_tool_githubconsolidated", "integration_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_githubconsolidated.cd_toolchain_tool_githubconsolidated", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_githubconsolidated.cd_toolchain_tool_githubconsolidated", "resource_group_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_githubconsolidated.cd_toolchain_tool_githubconsolidated", "crn"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_githubconsolidated.cd_toolchain_tool_githubconsolidated", "toolchain_crn"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_githubconsolidated.cd_toolchain_tool_githubconsolidated", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_githubconsolidated.cd_toolchain_tool_githubconsolidated", "referent.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_githubconsolidated.cd_toolchain_tool_githubconsolidated", "updated_at"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_githubconsolidated.cd_toolchain_tool_githubconsolidated", "parameters.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_githubconsolidated.cd_toolchain_tool_githubconsolidated", "state"),
				),
			},
		},
	})
}

func testAccCheckIBMCdToolchainToolGithubconsolidatedDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		data "ibm_cd_toolchain_tool_githubconsolidated" "cd_toolchain_tool_githubconsolidated" {
			toolchain_id = "toolchain_id"
			integration_id = "integration_id"
		}
	`)
}
