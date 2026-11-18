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
	"github.com/IBM/continuous-delivery-go-sdk/v2/cdtoolchainv2"
)

func TestAccIBMCdToolchainToolGithubconsolidatedBasic(t *testing.T) {
	var conf cdtoolchainv2.ToolchainTool
	tcName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	rgName := acc.CdResourceGroupName
	repoUrl := acc.CdGithubConsolidatedRepoUrl
	gitId := acc.CdGithubConsolidatedGitId

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCdToolchainToolGithubconsolidatedDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCdToolchainToolGithubconsolidatedConfigBasic(tcName, rgName, repoUrl, gitId),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCdToolchainToolGithubconsolidatedExists("ibm_cd_toolchain_tool_githubconsolidated.cd_toolchain_tool_githubconsolidated", conf),
					resource.TestCheckResourceAttrSet("ibm_cd_toolchain_tool_githubconsolidated.cd_toolchain_tool_githubconsolidated", "toolchain_id"),
				),
			},
		},
	})
}

func TestAccIBMCdToolchainToolGithubconsolidatedAllArgs(t *testing.T) {
	var conf cdtoolchainv2.ToolchainTool
	tcName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	rgName := acc.CdResourceGroupName
	repoUrl := acc.CdGithubConsolidatedRepoUrl
	gitId := acc.CdGithubConsolidatedGitId
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	nameUpdate := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCdToolchainToolGithubconsolidatedDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCdToolchainToolGithubconsolidatedConfig(tcName, rgName, repoUrl, gitId, name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCdToolchainToolGithubconsolidatedExists("ibm_cd_toolchain_tool_githubconsolidated.cd_toolchain_tool_githubconsolidated", conf),
					resource.TestCheckResourceAttrSet("ibm_cd_toolchain_tool_githubconsolidated.cd_toolchain_tool_githubconsolidated", "toolchain_id"),
					resource.TestCheckResourceAttr("ibm_cd_toolchain_tool_githubconsolidated.cd_toolchain_tool_githubconsolidated", "name", name),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMCdToolchainToolGithubconsolidatedConfig(tcName, rgName, repoUrl, gitId, nameUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("ibm_cd_toolchain_tool_githubconsolidated.cd_toolchain_tool_githubconsolidated", "toolchain_id"),
					resource.TestCheckResourceAttr("ibm_cd_toolchain_tool_githubconsolidated.cd_toolchain_tool_githubconsolidated", "name", nameUpdate),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_cd_toolchain_tool_githubconsolidated.cd_toolchain_tool_githubconsolidated",
				ImportState:       true,
				ImportStateVerify: false,
			},
		},
	})
}

func TestAccIBMCdToolchainToolGithubconsolidatedGithubAppAuth(t *testing.T) {
	var conf cdtoolchainv2.ToolchainTool
	tcName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	rgName := acc.CdResourceGroupName
	repoUrl := acc.CdGithubConsolidatedRepoUrl
	gitId := acc.CdGithubConsolidatedGitId
	resourceName := "ibm_cd_toolchain_tool_githubconsolidated.cd_toolchain_tool_githubconsolidated"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCdToolchainToolGithubconsolidatedDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCdToolchainToolGithubconsolidatedConfigGithubApp(tcName, rgName, repoUrl, gitId),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCdToolchainToolGithubconsolidatedExists(resourceName, conf),
					resource.TestCheckResourceAttrSet(resourceName, "toolchain_id"),
					resource.TestCheckResourceAttr(resourceName, "parameters.0.auth_type", "github_app"),
					resource.TestCheckResourceAttr(resourceName, "parameters.0.github_app_auth_method", "managed"),
				),
			},
		},
	})
}

func TestAccIBMCdToolchainToolGithubconsolidatedGithubAppAuthCustom(t *testing.T) {
	var conf cdtoolchainv2.ToolchainTool
	tcName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	rgName := acc.CdResourceGroupName
	repoUrl := acc.CdGithubConsolidatedRepoUrl
	gitId := acc.CdGithubConsolidatedGitId
	appId := acc.CdGithubConsolidatedAppId
	appKey := acc.CdGithubConsolidatedAppKey
	resourceName := "ibm_cd_toolchain_tool_githubconsolidated.cd_toolchain_tool_githubconsolidated"

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acc.TestAccPreCheck(t)
			testAccPreCheckIBMCdToolchainToolGithubconsolidatedCustomApp(t)
		},
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCdToolchainToolGithubconsolidatedDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCdToolchainToolGithubconsolidatedConfigGithubAppCustom(tcName, rgName, repoUrl, gitId, appId, appKey),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCdToolchainToolGithubconsolidatedExists(resourceName, conf),
					resource.TestCheckResourceAttrSet(resourceName, "toolchain_id"),
					resource.TestCheckResourceAttr(resourceName, "parameters.0.auth_type", "github_app"),
					resource.TestCheckResourceAttr(resourceName, "parameters.0.github_app_auth_method", "custom"),
					resource.TestCheckResourceAttr(resourceName, "parameters.0.github_app_credential_app_id", appId),
					resource.TestCheckResourceAttrSet(resourceName, "parameters.0.github_app_credential_private_key"),
				),
			},
		},
	})
}

// testAccPreCheckIBMCdToolchainToolGithubconsolidatedCustomApp skips the test when the custom
// GitHub App credentials are not provided through the environment.
func testAccPreCheckIBMCdToolchainToolGithubconsolidatedCustomApp(t *testing.T) {
	if acc.CdGithubConsolidatedAppId == "" || acc.CdGithubConsolidatedAppKey == "" {
		t.Skip("IBM_CD_GITHUB_CONSOLIDATED_APP_ID and IBM_CD_GITHUB_CONSOLIDATED_APP_PRIVATE_KEY must be set to test custom GitHub App auth")
	}
}

// testAccCheckIBMCdToolchainToolGithubconsolidatedGitIdAttr returns the git_id attribute line for the
// initialization block, or an empty string when no git_id is set so the provider default applies.
func testAccCheckIBMCdToolchainToolGithubconsolidatedGitIdAttr(gitId string) string {
	if gitId == "" {
		return ""
	}
	return fmt.Sprintf(`git_id = "%s"`, gitId)
}

func testAccCheckIBMCdToolchainToolGithubconsolidatedConfigGithubApp(tcName string, rgName string, repoUrl string, gitId string) string {
	gitIdAttr := testAccCheckIBMCdToolchainToolGithubconsolidatedGitIdAttr(gitId)
	return fmt.Sprintf(`
		data "ibm_resource_group" "resource_group" {
			name = "%s"
		}

		resource "ibm_cd_toolchain" "cd_toolchain" {
			name = "%s"
			resource_group_id = data.ibm_resource_group.resource_group.id
		}

		resource "ibm_cd_toolchain_tool_githubconsolidated" "cd_toolchain_tool_githubconsolidated" {
			toolchain_id = ibm_cd_toolchain.cd_toolchain.id
			parameters {
				auth_type = "github_app"
				github_app_auth_method = "managed"
			}
			initialization {
				%s
				repo_url = "%s"
				type = "link"
			}
		}
	`, rgName, tcName, gitIdAttr, repoUrl)
}

func testAccCheckIBMCdToolchainToolGithubconsolidatedConfigGithubAppCustom(tcName string, rgName string, repoUrl string, gitId string, appId string, appKey string) string {
	gitIdAttr := testAccCheckIBMCdToolchainToolGithubconsolidatedGitIdAttr(gitId)
	return fmt.Sprintf(`
		data "ibm_resource_group" "resource_group" {
			name = "%s"
		}

		resource "ibm_cd_toolchain" "cd_toolchain" {
			name = "%s"
			resource_group_id = data.ibm_resource_group.resource_group.id
		}

		resource "ibm_cd_toolchain_tool_githubconsolidated" "cd_toolchain_tool_githubconsolidated" {
			toolchain_id = ibm_cd_toolchain.cd_toolchain.id
			parameters {
				auth_type = "github_app"
				github_app_auth_method = "custom"
				github_app_credential_app_id = "%s"
				github_app_credential_private_key = %q
			}
			initialization {
				%s
				repo_url = "%s"
				type = "link"
			}
		}
	`, rgName, tcName, appId, appKey, gitIdAttr, repoUrl)
}

func testAccCheckIBMCdToolchainToolGithubconsolidatedConfigBasic(tcName string, rgName string, repoUrl string, gitId string) string {
	gitIdAttr := testAccCheckIBMCdToolchainToolGithubconsolidatedGitIdAttr(gitId)
	return fmt.Sprintf(`
		data "ibm_resource_group" "resource_group" {
			name = "%s"
		}

		resource "ibm_cd_toolchain" "cd_toolchain" {
			name = "%s"
			resource_group_id = data.ibm_resource_group.resource_group.id
		}

		resource "ibm_cd_toolchain_tool_githubconsolidated" "cd_toolchain_tool_githubconsolidated" {
			toolchain_id = ibm_cd_toolchain.cd_toolchain.id
			parameters {
				toolchain_issues_enabled = true
				enable_traceability = true
			}
			initialization {
				%s
				repo_url = "%s"
				type = "link"
			}
		}
	`, rgName, tcName, gitIdAttr, repoUrl)
}

func testAccCheckIBMCdToolchainToolGithubconsolidatedConfig(tcName string, rgName string, repoUrl string, gitId string, name string) string {
	gitIdAttr := testAccCheckIBMCdToolchainToolGithubconsolidatedGitIdAttr(gitId)
	return fmt.Sprintf(`
		data "ibm_resource_group" "resource_group" {
			name = "%s"
		}

		resource "ibm_cd_toolchain" "cd_toolchain" {
			name = "%s"
			resource_group_id = data.ibm_resource_group.resource_group.id
		}

		resource "ibm_cd_toolchain_tool_githubconsolidated" "cd_toolchain_tool_githubconsolidated" {
			toolchain_id = ibm_cd_toolchain.cd_toolchain.id
			parameters {
				toolchain_issues_enabled = true
				enable_traceability = true
			}
			initialization {
				%s
				repo_url = "%s"
				type = "link"
			}
			name = "%s"
		}
	`, rgName, tcName, gitIdAttr, repoUrl, name)
}

func testAccCheckIBMCdToolchainToolGithubconsolidatedExists(n string, obj cdtoolchainv2.ToolchainTool) resource.TestCheckFunc {

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

		toolchainTool, _, err := cdToolchainClient.GetToolByID(getToolByIDOptions)
		if err != nil {
			return err
		}

		obj = *toolchainTool
		return nil
	}
}

func testAccCheckIBMCdToolchainToolGithubconsolidatedDestroy(s *terraform.State) error {
	cdToolchainClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).CdToolchainV2()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_cd_toolchain_tool_githubconsolidated" {
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
			return fmt.Errorf("cd_toolchain_tool_githubconsolidated still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for cd_toolchain_tool_githubconsolidated (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
