// Copyright IBM Corp. 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package cdtoolchain_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"
)

func TestAccIBMCdToolchainToolHostedgitDataSourceBasic(t *testing.T) {
	tcName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	rgName := acc.CdResourceGroupName
	repoUrl := acc.CdHostedGitRepoUrl

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCdToolchainToolHostedgitDataSourceConfigBasic(tcName, rgName, repoUrl),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_hostedgit.cd_toolchain_tool_hostedgit", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_hostedgit.cd_toolchain_tool_hostedgit", "toolchain_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_hostedgit.cd_toolchain_tool_hostedgit", "tool_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_hostedgit.cd_toolchain_tool_hostedgit", "resource_group_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_hostedgit.cd_toolchain_tool_hostedgit", "crn"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_hostedgit.cd_toolchain_tool_hostedgit", "toolchain_crn"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_hostedgit.cd_toolchain_tool_hostedgit", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_hostedgit.cd_toolchain_tool_hostedgit", "referent.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_hostedgit.cd_toolchain_tool_hostedgit", "updated_at"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_hostedgit.cd_toolchain_tool_hostedgit", "parameters.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_hostedgit.cd_toolchain_tool_hostedgit", "state"),
				),
			},
		},
	})
}

func TestAccIBMCdToolchainToolHostedgitDataSourceAllArgs(t *testing.T) {
	tcName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	rgName := acc.CdResourceGroupName
	repoUrl := acc.CdHostedGitRepoUrl
	toolName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCdToolchainToolHostedgitDataSourceConfig(tcName, rgName, repoUrl, toolName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_hostedgit.cd_toolchain_tool_hostedgit", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_hostedgit.cd_toolchain_tool_hostedgit", "toolchain_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_hostedgit.cd_toolchain_tool_hostedgit", "tool_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_hostedgit.cd_toolchain_tool_hostedgit", "resource_group_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_hostedgit.cd_toolchain_tool_hostedgit", "crn"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_hostedgit.cd_toolchain_tool_hostedgit", "toolchain_crn"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_hostedgit.cd_toolchain_tool_hostedgit", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_hostedgit.cd_toolchain_tool_hostedgit", "referent.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_hostedgit.cd_toolchain_tool_hostedgit", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_hostedgit.cd_toolchain_tool_hostedgit", "updated_at"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_hostedgit.cd_toolchain_tool_hostedgit", "parameters.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_hostedgit.cd_toolchain_tool_hostedgit", "state"),
				),
			},
		},
	})
}

func testAccCheckIBMCdToolchainToolHostedgitDataSourceConfigBasic(tcName string, rgName string, repoUrl string) string {
	return fmt.Sprintf(`
		data "ibm_resource_group" "resource_group" {
			name = "%s"
		}

		resource "ibm_cd_toolchain" "cd_toolchain" {
			name = "%s"
			resource_group_id = data.ibm_resource_group.resource_group.id
		}

		resource "ibm_cd_toolchain_tool_hostedgit" "cd_toolchain_tool_hostedgit" {
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

		data "ibm_cd_toolchain_tool_hostedgit" "cd_toolchain_tool_hostedgit" {
			toolchain_id = ibm_cd_toolchain_tool_hostedgit.cd_toolchain_tool_hostedgit.toolchain_id
			tool_id = ibm_cd_toolchain_tool_hostedgit.cd_toolchain_tool_hostedgit.tool_id
		}
	`, rgName, tcName, repoUrl)
}

func testAccCheckIBMCdToolchainToolHostedgitDataSourceConfig(tcName string, rgName string, repoUrl string, toolName string) string {
	return fmt.Sprintf(`
		data "ibm_resource_group" "resource_group" {
			name = "%s"
		}

		resource "ibm_cd_toolchain" "cd_toolchain" {
			name = "%s"
			resource_group_id = data.ibm_resource_group.resource_group.id
		}

		resource "ibm_cd_toolchain_tool_hostedgit" "cd_toolchain_tool_hostedgit" {
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

		data "ibm_cd_toolchain_tool_hostedgit" "cd_toolchain_tool_hostedgit" {
			toolchain_id = ibm_cd_toolchain_tool_hostedgit.cd_toolchain_tool_hostedgit.toolchain_id
			tool_id = ibm_cd_toolchain_tool_hostedgit.cd_toolchain_tool_hostedgit.tool_id
		}
	`, rgName, tcName, repoUrl, toolName)
}
