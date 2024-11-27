// Copyright IBM Corp. 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package cdtoolchain_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIBMCdToolchainToolGithubconsolidatedDataSourceBasic(t *testing.T) {
	tcName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	rgName := acc.CdResourceGroupName
	repoUrl := acc.CdGithubConsolidatedRepoUrl

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCdToolchainToolGithubconsolidatedDataSourceConfigBasic(tcName, rgName, repoUrl),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_githubconsolidated.cd_toolchain_tool_githubconsolidated", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_githubconsolidated.cd_toolchain_tool_githubconsolidated", "toolchain_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_githubconsolidated.cd_toolchain_tool_githubconsolidated", "tool_id"),
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

func TestAccIBMCdToolchainToolGithubconsolidatedDataSourceAllArgs(t *testing.T) {
	tcName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	rgName := acc.CdResourceGroupName
	repoUrl := acc.CdGithubConsolidatedRepoUrl
	toolName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCdToolchainToolGithubconsolidatedDataSourceConfig(tcName, rgName, repoUrl, toolName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_githubconsolidated.cd_toolchain_tool_githubconsolidated", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_githubconsolidated.cd_toolchain_tool_githubconsolidated", "toolchain_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_githubconsolidated.cd_toolchain_tool_githubconsolidated", "tool_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_githubconsolidated.cd_toolchain_tool_githubconsolidated", "resource_group_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_githubconsolidated.cd_toolchain_tool_githubconsolidated", "crn"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_githubconsolidated.cd_toolchain_tool_githubconsolidated", "toolchain_crn"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_githubconsolidated.cd_toolchain_tool_githubconsolidated", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_githubconsolidated.cd_toolchain_tool_githubconsolidated", "referent.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_githubconsolidated.cd_toolchain_tool_githubconsolidated", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_githubconsolidated.cd_toolchain_tool_githubconsolidated", "updated_at"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_githubconsolidated.cd_toolchain_tool_githubconsolidated", "parameters.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_githubconsolidated.cd_toolchain_tool_githubconsolidated", "state"),
				),
			},
		},
	})
}

func testAccCheckIBMCdToolchainToolGithubconsolidatedDataSourceConfigBasic(tcName string, rgName string, repoUrl string) string {
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
				repo_url = "%s"
				type = "link"
			}
		}

		data "ibm_cd_toolchain_tool_githubconsolidated" "cd_toolchain_tool_githubconsolidated" {
			toolchain_id = ibm_cd_toolchain_tool_githubconsolidated.cd_toolchain_tool_githubconsolidated.toolchain_id
			tool_id = ibm_cd_toolchain_tool_githubconsolidated.cd_toolchain_tool_githubconsolidated.tool_id
		}
	`, rgName, tcName, repoUrl)
}

func testAccCheckIBMCdToolchainToolGithubconsolidatedDataSourceConfig(tcName string, rgName string, repoUrl string, toolName string) string {
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
				repo_url = "%s"
				type = "link"
			}
			name = "%s"
		}

		data "ibm_cd_toolchain_tool_githubconsolidated" "cd_toolchain_tool_githubconsolidated" {
			toolchain_id = ibm_cd_toolchain_tool_githubconsolidated.cd_toolchain_tool_githubconsolidated.toolchain_id
			tool_id = ibm_cd_toolchain_tool_githubconsolidated.cd_toolchain_tool_githubconsolidated.tool_id
		}
	`, rgName, tcName, repoUrl, toolName)
}
