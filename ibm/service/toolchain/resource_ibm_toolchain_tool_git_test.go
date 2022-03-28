// Copyright IBM Corp. 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package toolchain_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.ibm.com/org-ids/toolchain-go-sdk/toolchainv2"
)

func TestAccIBMToolchainToolGitBasic(t *testing.T) {
	var conf toolchainv2.GetIntegrationByIDResponse
	toolchainID := fmt.Sprintf("tf_toolchain_id_%d", acctest.RandIntRange(10, 100))
	gitProvider := fmt.Sprintf("tf_git_provider_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMToolchainToolGitDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMToolchainToolGitConfigBasic(toolchainID, gitProvider),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMToolchainToolGitExists("ibm_toolchain_tool_git.toolchain_tool_git", conf),
					resource.TestCheckResourceAttr("ibm_toolchain_tool_git.toolchain_tool_git", "toolchain_id", toolchainID),
					resource.TestCheckResourceAttr("ibm_toolchain_tool_git.toolchain_tool_git", "git_provider", gitProvider),
				),
			},
		},
	})
}

func TestAccIBMToolchainToolGitAllArgs(t *testing.T) {
	var conf toolchainv2.GetIntegrationByIDResponse
	toolchainID := fmt.Sprintf("tf_toolchain_id_%d", acctest.RandIntRange(10, 100))
	gitProvider := fmt.Sprintf("tf_git_provider_%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	nameUpdate := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMToolchainToolGitDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMToolchainToolGitConfig(toolchainID, gitProvider, name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMToolchainToolGitExists("ibm_toolchain_tool_git.toolchain_tool_git", conf),
					resource.TestCheckResourceAttr("ibm_toolchain_tool_git.toolchain_tool_git", "toolchain_id", toolchainID),
					resource.TestCheckResourceAttr("ibm_toolchain_tool_git.toolchain_tool_git", "git_provider", gitProvider),
					resource.TestCheckResourceAttr("ibm_toolchain_tool_git.toolchain_tool_git", "name", name),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMToolchainToolGitConfig(toolchainID, gitProvider, nameUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_toolchain_tool_git.toolchain_tool_git", "toolchain_id", toolchainID),
					resource.TestCheckResourceAttr("ibm_toolchain_tool_git.toolchain_tool_git", "git_provider", gitProvider),
					resource.TestCheckResourceAttr("ibm_toolchain_tool_git.toolchain_tool_git", "name", nameUpdate),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_toolchain_tool_git.toolchain_tool_git",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIBMToolchainToolGitConfigBasic(toolchainID string, gitProvider string) string {
	return fmt.Sprintf(`

		resource "ibm_toolchain_tool_git" "toolchain_tool_git" {
			toolchain_id = "%s"
			git_provider = "%s"
		}
	`, toolchainID, gitProvider)
}

func testAccCheckIBMToolchainToolGitConfig(toolchainID string, gitProvider string, name string) string {
	return fmt.Sprintf(`

		resource "ibm_toolchain_tool_git" "toolchain_tool_git" {
			toolchain_id = "%s"
			git_provider = "%s"
			name = "%s"
			parameters {
				enable_traceability = true
				has_issues = true
				repo_name = "repo_name"
				repo_url = "repo_url"
				source_repo_url = "source_repo_url"
				type = "new"
				private_repo = true
				git_id = "git_id"
				owner_id = "owner_id"
			}
			initialization {
				repo_name = "repo_name"
				repo_url = "repo_url"
				source_repo_url = "source_repo_url"
				type = "new"
				private_repo = true
				git_id = "git_id"
				owner_id = "owner_id"
			}
		}
	`, toolchainID, gitProvider, name)
}

func testAccCheckIBMToolchainToolGitExists(n string, obj toolchainv2.GetIntegrationByIDResponse) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		toolchainClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).ToolchainV2()
		if err != nil {
			return err
		}

		getIntegrationByIDOptions := &toolchainv2.GetIntegrationByIDOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getIntegrationByIDOptions.SetToolchainID(parts[0])
		getIntegrationByIDOptions.SetIntegrationID(parts[1])

		getIntegrationByIDResponse, _, err := toolchainClient.GetIntegrationByID(getIntegrationByIDOptions)
		if err != nil {
			return err
		}

		obj = *getIntegrationByIDResponse
		return nil
	}
}

func testAccCheckIBMToolchainToolGitDestroy(s *terraform.State) error {
	toolchainClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).ToolchainV2()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_toolchain_tool_git" {
			continue
		}

		getIntegrationByIDOptions := &toolchainv2.GetIntegrationByIDOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getIntegrationByIDOptions.SetToolchainID(parts[0])
		getIntegrationByIDOptions.SetIntegrationID(parts[1])

		// Try to find the key
		_, response, err := toolchainClient.GetIntegrationByID(getIntegrationByIDOptions)

		if err == nil {
			return fmt.Errorf("toolchain_tool_git still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for toolchain_tool_git (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
