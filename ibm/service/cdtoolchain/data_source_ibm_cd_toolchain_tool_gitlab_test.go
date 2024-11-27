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

func TestAccIBMCdToolchainToolGitlabDataSourceBasic(t *testing.T) {
	tcName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	rgName := acc.CdResourceGroupName
	repoUrl := acc.CdGitlabRepoUrl

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCdToolchainToolGitlabDataSourceConfigBasic(tcName, rgName, repoUrl),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_gitlab.cd_toolchain_tool_gitlab", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_gitlab.cd_toolchain_tool_gitlab", "toolchain_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_gitlab.cd_toolchain_tool_gitlab", "tool_id"),
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

func TestAccIBMCdToolchainToolGitlabDataSourceAllArgs(t *testing.T) {
	tcName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	rgName := acc.CdResourceGroupName
	repoUrl := acc.CdGitlabRepoUrl
	toolName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCdToolchainToolGitlabDataSourceConfig(tcName, rgName, repoUrl, toolName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_gitlab.cd_toolchain_tool_gitlab", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_gitlab.cd_toolchain_tool_gitlab", "toolchain_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_gitlab.cd_toolchain_tool_gitlab", "tool_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_gitlab.cd_toolchain_tool_gitlab", "resource_group_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_gitlab.cd_toolchain_tool_gitlab", "crn"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_gitlab.cd_toolchain_tool_gitlab", "toolchain_crn"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_gitlab.cd_toolchain_tool_gitlab", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_gitlab.cd_toolchain_tool_gitlab", "referent.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_gitlab.cd_toolchain_tool_gitlab", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_gitlab.cd_toolchain_tool_gitlab", "updated_at"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_gitlab.cd_toolchain_tool_gitlab", "parameters.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_gitlab.cd_toolchain_tool_gitlab", "state"),
				),
			},
		},
	})
}

func testAccCheckIBMCdToolchainToolGitlabDataSourceConfigBasic(tcName string, rgName string, repoUrl string) string {
	return fmt.Sprintf(`
		data "ibm_resource_group" "resource_group" {
			name = "%s"
		}

		resource "ibm_cd_toolchain" "cd_toolchain" {
			name = "%s"
			resource_group_id = data.ibm_resource_group.resource_group.id
		}

		resource "ibm_cd_toolchain_tool_gitlab" "cd_toolchain_tool_gitlab" {
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

		data "ibm_cd_toolchain_tool_gitlab" "cd_toolchain_tool_gitlab" {
			toolchain_id = ibm_cd_toolchain_tool_gitlab.cd_toolchain_tool_gitlab.toolchain_id
			tool_id = ibm_cd_toolchain_tool_gitlab.cd_toolchain_tool_gitlab.tool_id
		}
	`, rgName, tcName, repoUrl)
}

func testAccCheckIBMCdToolchainToolGitlabDataSourceConfig(tcName string, rgName string, repoUrl string, toolName string) string {
	return fmt.Sprintf(`
		data "ibm_resource_group" "resource_group" {
			name = "%s"
		}

		resource "ibm_cd_toolchain" "cd_toolchain" {
			name = "%s"
			resource_group_id = data.ibm_resource_group.resource_group.id
		}

		resource "ibm_cd_toolchain_tool_gitlab" "cd_toolchain_tool_gitlab" {
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

		data "ibm_cd_toolchain_tool_gitlab" "cd_toolchain_tool_gitlab" {
			toolchain_id = ibm_cd_toolchain_tool_gitlab.cd_toolchain_tool_gitlab.toolchain_id
			tool_id = ibm_cd_toolchain_tool_gitlab.cd_toolchain_tool_gitlab.tool_id
		}
	`, rgName, tcName, repoUrl, toolName)
}
