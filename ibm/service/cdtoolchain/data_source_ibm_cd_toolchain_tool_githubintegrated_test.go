// Copyright IBM Corp. 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package cdtoolchain_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIBMCdToolchainToolGithubintegratedDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCdToolchainToolGithubintegratedDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_githubintegrated.cd_toolchain_tool_githubintegrated", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_githubintegrated.cd_toolchain_tool_githubintegrated", "toolchain_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_githubintegrated.cd_toolchain_tool_githubintegrated", "tool_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_githubintegrated.cd_toolchain_tool_githubintegrated", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_githubintegrated.cd_toolchain_tool_githubintegrated", "resource_group_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_githubintegrated.cd_toolchain_tool_githubintegrated", "crn"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_githubintegrated.cd_toolchain_tool_githubintegrated", "toolchain_crn"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_githubintegrated.cd_toolchain_tool_githubintegrated", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_githubintegrated.cd_toolchain_tool_githubintegrated", "referent.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_githubintegrated.cd_toolchain_tool_githubintegrated", "updated_at"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_githubintegrated.cd_toolchain_tool_githubintegrated", "parameters.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_githubintegrated.cd_toolchain_tool_githubintegrated", "state"),
				),
			},
		},
	})
}

func testAccCheckIBMCdToolchainToolGithubintegratedDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		data "ibm_cd_toolchain_tool_githubintegrated" "cd_toolchain_tool_githubintegrated" {
			toolchain_id = "toolchain_id"
			tool_id = "tool_id"
		}
	`)
}
