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
	"github.ibm.com/org-ids/toolchain-go-sdk/cdtoolchainv2"
)

func TestAccIBMCdToolchainToolGithubpublicBasic(t *testing.T) {
	var conf cdtoolchainv2.GetIntegrationByIDResponse
	toolchainID := fmt.Sprintf("tf_toolchain_id_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCdToolchainToolGithubpublicDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCdToolchainToolGithubpublicConfigBasic(toolchainID),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCdToolchainToolGithubpublicExists("ibm_cd_toolchain_tool_githubpublic.cd_toolchain_tool_githubpublic", conf),
					resource.TestCheckResourceAttr("ibm_cd_toolchain_tool_githubpublic.cd_toolchain_tool_githubpublic", "toolchain_id", toolchainID),
				),
			},
		},
	})
}

func TestAccIBMCdToolchainToolGithubpublicAllArgs(t *testing.T) {
	var conf cdtoolchainv2.GetIntegrationByIDResponse
	toolchainID := fmt.Sprintf("tf_toolchain_id_%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	nameUpdate := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCdToolchainToolGithubpublicDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCdToolchainToolGithubpublicConfig(toolchainID, name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCdToolchainToolGithubpublicExists("ibm_cd_toolchain_tool_githubpublic.cd_toolchain_tool_githubpublic", conf),
					resource.TestCheckResourceAttr("ibm_cd_toolchain_tool_githubpublic.cd_toolchain_tool_githubpublic", "toolchain_id", toolchainID),
					resource.TestCheckResourceAttr("ibm_cd_toolchain_tool_githubpublic.cd_toolchain_tool_githubpublic", "name", name),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMCdToolchainToolGithubpublicConfig(toolchainID, nameUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_cd_toolchain_tool_githubpublic.cd_toolchain_tool_githubpublic", "toolchain_id", toolchainID),
					resource.TestCheckResourceAttr("ibm_cd_toolchain_tool_githubpublic.cd_toolchain_tool_githubpublic", "name", nameUpdate),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_cd_toolchain_tool_githubpublic.cd_toolchain_tool_githubpublic",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIBMCdToolchainToolGithubpublicConfigBasic(toolchainID string) string {
	return fmt.Sprintf(`

		resource "ibm_cd_toolchain_tool_githubpublic" "cd_toolchain_tool_githubpublic" {
			toolchain_id = "%s"
		}
	`, toolchainID)
}

func testAccCheckIBMCdToolchainToolGithubpublicConfig(toolchainID string, name string) string {
	return fmt.Sprintf(`

		resource "ibm_cd_toolchain_tool_githubpublic" "cd_toolchain_tool_githubpublic" {
			toolchain_id = "%s"
			name = "%s"
			parameters {
				legal = true
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
				auto_init = true
				enable_traceability = true
				authorized = "authorized"
				integration_owner = "integration_owner"
				auth_type = "oauth"
				api_token = "api_token"
			}
			initialization {
				legal = true
				repo_name = "repo_name"
				repo_url = "repo_url"
				source_repo_url = "source_repo_url"
				type = "new"
				private_repo = true
			}
		}
	`, toolchainID, name)
}

func testAccCheckIBMCdToolchainToolGithubpublicExists(n string, obj cdtoolchainv2.GetIntegrationByIDResponse) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		cdToolchainClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).CdToolchainV2()
		if err != nil {
			return err
		}

		getIntegrationByIDOptions := &cdtoolchainv2.GetIntegrationByIDOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getIntegrationByIDOptions.SetToolchainID(parts[0])
		getIntegrationByIDOptions.SetIntegrationID(parts[1])

		getIntegrationByIDResponse, _, err := cdToolchainClient.GetIntegrationByID(getIntegrationByIDOptions)
		if err != nil {
			return err
		}

		obj = *getIntegrationByIDResponse
		return nil
	}
}

func testAccCheckIBMCdToolchainToolGithubpublicDestroy(s *terraform.State) error {
	cdToolchainClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).CdToolchainV2()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_cd_toolchain_tool_githubpublic" {
			continue
		}

		getIntegrationByIDOptions := &cdtoolchainv2.GetIntegrationByIDOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getIntegrationByIDOptions.SetToolchainID(parts[0])
		getIntegrationByIDOptions.SetIntegrationID(parts[1])

		// Try to find the key
		_, response, err := cdToolchainClient.GetIntegrationByID(getIntegrationByIDOptions)

		if err == nil {
			return fmt.Errorf("cd_toolchain_tool_githubpublic still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for cd_toolchain_tool_githubpublic (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
