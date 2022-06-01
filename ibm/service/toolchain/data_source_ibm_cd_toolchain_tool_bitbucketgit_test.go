// Copyright IBM Corp. 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package toolchain_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIBMCdToolchainToolBitbucketgitDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCdToolchainToolBitbucketgitDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_bitbucketgit.cd_toolchain_tool_bitbucketgit", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_bitbucketgit.cd_toolchain_tool_bitbucketgit", "toolchain_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_bitbucketgit.cd_toolchain_tool_bitbucketgit", "integration_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_bitbucketgit.cd_toolchain_tool_bitbucketgit", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_bitbucketgit.cd_toolchain_tool_bitbucketgit", "resource_group_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_bitbucketgit.cd_toolchain_tool_bitbucketgit", "crn"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_bitbucketgit.cd_toolchain_tool_bitbucketgit", "toolchain_crn"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_bitbucketgit.cd_toolchain_tool_bitbucketgit", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_bitbucketgit.cd_toolchain_tool_bitbucketgit", "referent.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_bitbucketgit.cd_toolchain_tool_bitbucketgit", "updated_at"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_bitbucketgit.cd_toolchain_tool_bitbucketgit", "parameters.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_bitbucketgit.cd_toolchain_tool_bitbucketgit", "state"),
				),
			},
		},
	})
}

func testAccCheckIBMCdToolchainToolBitbucketgitDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		data "ibm_cd_toolchain_tool_bitbucketgit" "cd_toolchain_tool_bitbucketgit" {
			toolchain_id = "toolchain_id"
			integration_id = "integration_id"
		}
	`)
}
