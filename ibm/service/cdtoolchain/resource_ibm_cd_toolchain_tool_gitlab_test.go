// Copyright IBM Corp. 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package cdtoolchain_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/continuous-delivery-go-sdk/cdtoolchainv2"
)

func TestAccIBMCdToolchainToolGitlabBasic(t *testing.T) {
	var conf cdtoolchainv2.GetToolByIDResponse
	toolchainID := fmt.Sprintf("tf_toolchain_id_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCdToolchainToolGitlabDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCdToolchainToolGitlabConfigBasic(toolchainID),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCdToolchainToolGitlabExists("ibm_cd_toolchain_tool_gitlab.cd_toolchain_tool_gitlab", conf),
					resource.TestCheckResourceAttr("ibm_cd_toolchain_tool_gitlab.cd_toolchain_tool_gitlab", "toolchain_id", toolchainID),
				),
			},
		},
	})
}

func TestAccIBMCdToolchainToolGitlabAllArgs(t *testing.T) {
	var conf cdtoolchainv2.GetToolByIDResponse
	toolchainID := fmt.Sprintf("tf_toolchain_id_%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	nameUpdate := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCdToolchainToolGitlabDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCdToolchainToolGitlabConfig(toolchainID, name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCdToolchainToolGitlabExists("ibm_cd_toolchain_tool_gitlab.cd_toolchain_tool_gitlab", conf),
					resource.TestCheckResourceAttr("ibm_cd_toolchain_tool_gitlab.cd_toolchain_tool_gitlab", "toolchain_id", toolchainID),
					resource.TestCheckResourceAttr("ibm_cd_toolchain_tool_gitlab.cd_toolchain_tool_gitlab", "name", name),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMCdToolchainToolGitlabConfig(toolchainID, nameUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_cd_toolchain_tool_gitlab.cd_toolchain_tool_gitlab", "toolchain_id", toolchainID),
					resource.TestCheckResourceAttr("ibm_cd_toolchain_tool_gitlab.cd_toolchain_tool_gitlab", "name", nameUpdate),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_cd_toolchain_tool_gitlab.cd_toolchain_tool_gitlab",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIBMCdToolchainToolGitlabConfigBasic(toolchainID string) string {
	return fmt.Sprintf(`

		resource "ibm_cd_toolchain_tool_gitlab" "cd_toolchain_tool_gitlab" {
			toolchain_id = "%s"
		}
	`, toolchainID)
}

func testAccCheckIBMCdToolchainToolGitlabConfig(toolchainID string, name string) string {
	return fmt.Sprintf(`

		resource "ibm_cd_toolchain_tool_gitlab" "cd_toolchain_tool_gitlab" {
			toolchain_id = "%s"
			name = "%s"
			parameters {
				git_id = "git_id"
				title = "title"
				api_root_url = "api_root_url"
				default_branch = "default_branch"
				root_url = "root_url"
				access_token = "access_token"
				blind_connection = true
				owner_id = "owner_id"
				repo_name = "repo_name"
				repo_url = "repo_url"
				source_repo_url = "source_repo_url"
				token_url = "token_url"
				type = "new"
				private_repo = true
				has_issues = true
				enable_traceability = true
				authorized = "authorized"
				integration_owner = "integration_owner"
				auth_type = "oauth"
				api_token = "api_token"
			}
			initialization {
				repo_name = "repo_name"
				repo_url = "repo_url"
				source_repo_url = "source_repo_url"
				type = "new"
				private_repo = true
			}
		}
	`, toolchainID, name)
}

func testAccCheckIBMCdToolchainToolGitlabExists(n string, obj cdtoolchainv2.GetToolByIDResponse) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		cdToolchainClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).CdToolchainV2()
		if err != nil {
			return err
		}

		getToolByIDOptions := &cdtoolchainv2.GetToolByIDOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getToolByIDOptions.SetToolchainID(parts[0])
		getToolByIDOptions.SetToolID(parts[1])

		getToolByIDResponse, _, err := cdToolchainClient.GetToolByID(getToolByIDOptions)
		if err != nil {
			return err
		}

		obj = *getToolByIDResponse
		return nil
	}
}

func testAccCheckIBMCdToolchainToolGitlabDestroy(s *terraform.State) error {
	cdToolchainClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).CdToolchainV2()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_cd_toolchain_tool_gitlab" {
			continue
		}

		getToolByIDOptions := &cdtoolchainv2.GetToolByIDOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getToolByIDOptions.SetToolchainID(parts[0])
		getToolByIDOptions.SetToolID(parts[1])

		// Try to find the key
		_, response, err := cdToolchainClient.GetToolByID(getToolByIDOptions)

		if err == nil {
			return fmt.Errorf("cd_toolchain_tool_gitlab still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for cd_toolchain_tool_gitlab (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
