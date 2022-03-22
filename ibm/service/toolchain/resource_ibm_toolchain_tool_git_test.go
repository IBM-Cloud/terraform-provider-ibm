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

func TestAccIbmToolchainToolGitBasic(t *testing.T) {
	var conf toolchainv2.GetIntegrationByIdResponse
	toolchainID := fmt.Sprintf("tf_toolchain_id_%d", acctest.RandIntRange(10, 100))
	gitProvider := fmt.Sprintf("tf_git_provider_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmToolchainToolGitDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmToolchainToolGitConfigBasic(toolchainID, gitProvider),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmToolchainToolGitExists("ibm_toolchain_tool_git.toolchain_tool_git", conf),
					resource.TestCheckResourceAttr("ibm_toolchain_tool_git.toolchain_tool_git", "toolchain_id", toolchainID),
					resource.TestCheckResourceAttr("ibm_toolchain_tool_git.toolchain_tool_git", "git_provider", gitProvider),
				),
			},
		},
	})
}

func TestAccIbmToolchainToolGitAllArgs(t *testing.T) {
	var conf toolchainv2.GetIntegrationByIdResponse
	toolchainID := fmt.Sprintf("tf_toolchain_id_%d", acctest.RandIntRange(10, 100))
	gitProvider := fmt.Sprintf("tf_git_provider_%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	nameUpdate := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmToolchainToolGitDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmToolchainToolGitConfig(toolchainID, gitProvider, name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmToolchainToolGitExists("ibm_toolchain_tool_git.toolchain_tool_git", conf),
					resource.TestCheckResourceAttr("ibm_toolchain_tool_git.toolchain_tool_git", "toolchain_id", toolchainID),
					resource.TestCheckResourceAttr("ibm_toolchain_tool_git.toolchain_tool_git", "git_provider", gitProvider),
					resource.TestCheckResourceAttr("ibm_toolchain_tool_git.toolchain_tool_git", "name", name),
				),
			},
			resource.TestStep{
				Config: testAccCheckIbmToolchainToolGitConfig(toolchainID, gitProvider, nameUpdate),
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

func testAccCheckIbmToolchainToolGitConfigBasic(toolchainID string, gitProvider string) string {
	return fmt.Sprintf(`

		resource "ibm_toolchain_tool_git" "toolchain_tool_git" {
			toolchain_id = "%s"
			git_provider = "%s"
		}
	`, toolchainID, gitProvider)
}

func testAccCheckIbmToolchainToolGitConfig(toolchainID string, gitProvider string, name string) string {
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
			}
			initialization {
				repo_name = "repo_name"
				repo_url = "repo_url"
				source_repo_url = "source_repo_url"
				type = "new"
				private_repo = true
			}
			parameters_references = "FIXME"
		}
	`, toolchainID, gitProvider, name)
}

func testAccCheckIbmToolchainToolGitExists(n string, obj toolchainv2.GetIntegrationByIdResponse) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		toolchainClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).ToolchainV2()
		if err != nil {
			return err
		}

		getIntegrationByIdOptions := &toolchainv2.GetIntegrationByIdOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getIntegrationByIdOptions.SetToolchainID(parts[0])
		getIntegrationByIdOptions.SetIntegrationID(parts[1])

		getIntegrationByIdResponse, _, err := toolchainClient.GetIntegrationByID(getIntegrationByIdOptions)
		if err != nil {
			return err
		}

		obj = *getIntegrationByIdResponse
		return nil
	}
}

func testAccCheckIbmToolchainToolGitDestroy(s *terraform.State) error {
	toolchainClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).ToolchainV2()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_toolchain_tool_git" {
			continue
		}

		getIntegrationByIdOptions := &toolchainv2.GetIntegrationByIdOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getIntegrationByIdOptions.SetToolchainID(parts[0])
		getIntegrationByIdOptions.SetIntegrationID(parts[1])

		// Try to find the key
		_, response, err := toolchainClient.GetIntegrationByID(getIntegrationByIdOptions)

		if err == nil {
			return fmt.Errorf("toolchain_tool_git still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for toolchain_tool_git (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
