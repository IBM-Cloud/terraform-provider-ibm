// Copyright IBM Corp. 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package toolchain_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIBMCdToolchainToolGitlabDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCdToolchainToolGitlabDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_gitlab.cd_toolchain_tool_gitlab", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_gitlab.cd_toolchain_tool_gitlab", "toolchain_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_gitlab.cd_toolchain_tool_gitlab", "integration_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_gitlab.cd_toolchain_tool_gitlab", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_gitlab.cd_toolchain_tool_gitlab", "resource_group_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_gitlab.cd_toolchain_tool_gitlab", "crn"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_gitlab.cd_toolchain_tool_gitlab", "toolchain_crn"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_gitlab.cd_toolchain_tool_gitlab", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_gitlab.cd_toolchain_tool_gitlab", "referent.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_gitlab.cd_toolchain_tool_gitlab", "updated_at"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_gitlab.cd_toolchain_tool_gitlab", "parameters.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_gitlab.cd_toolchain_tool_gitlab", "state"),
				),
			},
		},
	})
}

func testAccCheckIBMCdToolchainToolGitlabDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		data "ibm_cd_toolchain_tool_gitlab" "cd_toolchain_tool_gitlab" {
			toolchain_id = "toolchain_id"
			integration_id = "integration_id"
		}
	`)
}
