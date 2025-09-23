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

func TestAccIBMCdToolchainToolJiraDataSourceBasic(t *testing.T) {
	tcName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	rgName := acc.CdResourceGroupName
	projectKey := acc.CdJiraProjectKey
	apiUrl := acc.CdJiraApiUrl
	username := acc.CdJiraUsername
	apiToken := acc.CdJiraApiToken

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCdToolchainToolJiraDataSourceConfigBasic(tcName, rgName, projectKey, apiUrl, username, apiToken),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_jira.cd_toolchain_tool_jira", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_jira.cd_toolchain_tool_jira", "toolchain_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_jira.cd_toolchain_tool_jira", "tool_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_jira.cd_toolchain_tool_jira", "resource_group_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_jira.cd_toolchain_tool_jira", "crn"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_jira.cd_toolchain_tool_jira", "toolchain_crn"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_jira.cd_toolchain_tool_jira", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_jira.cd_toolchain_tool_jira", "referent.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_jira.cd_toolchain_tool_jira", "updated_at"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_jira.cd_toolchain_tool_jira", "parameters.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_jira.cd_toolchain_tool_jira", "state"),
				),
			},
		},
	})
}

func TestAccIBMCdToolchainToolJiraDataSourceAllArgs(t *testing.T) {
	tcName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	rgName := acc.CdResourceGroupName
	toolName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	projectKey := acc.CdJiraProjectKey
	apiUrl := acc.CdJiraApiUrl
	username := acc.CdJiraUsername
	apiToken := acc.CdJiraApiToken

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCdToolchainToolJiraDataSourceConfig(tcName, rgName, projectKey, apiUrl, username, apiToken, toolName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_jira.cd_toolchain_tool_jira", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_jira.cd_toolchain_tool_jira", "toolchain_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_jira.cd_toolchain_tool_jira", "tool_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_jira.cd_toolchain_tool_jira", "resource_group_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_jira.cd_toolchain_tool_jira", "crn"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_jira.cd_toolchain_tool_jira", "toolchain_crn"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_jira.cd_toolchain_tool_jira", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_jira.cd_toolchain_tool_jira", "referent.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_jira.cd_toolchain_tool_jira", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_jira.cd_toolchain_tool_jira", "updated_at"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_jira.cd_toolchain_tool_jira", "parameters.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_jira.cd_toolchain_tool_jira", "state"),
				),
			},
		},
	})
}

func testAccCheckIBMCdToolchainToolJiraDataSourceConfigBasic(tcName string, rgName string, projectKey string, apiUrl string, username string, apiToken string) string {
	return fmt.Sprintf(`
		data "ibm_resource_group" "resource_group" {
			name = "%s"
		}

		resource "ibm_cd_toolchain" "cd_toolchain" {
			name = "%s"
			resource_group_id = data.ibm_resource_group.resource_group.id
		}

		resource "ibm_cd_toolchain_tool_jira" "cd_toolchain_tool_jira" {
			toolchain_id = ibm_cd_toolchain.cd_toolchain.id
			parameters {
				project_key = "%s"
				api_url = "%s"
				username = "%s"
				enable_traceability = true
				api_token = "%s"
			}
		}

		data "ibm_cd_toolchain_tool_jira" "cd_toolchain_tool_jira" {
			toolchain_id = ibm_cd_toolchain_tool_jira.cd_toolchain_tool_jira.toolchain_id
			tool_id = ibm_cd_toolchain_tool_jira.cd_toolchain_tool_jira.tool_id
		}
	`, rgName, tcName, projectKey, apiUrl, username, apiToken)
}

func testAccCheckIBMCdToolchainToolJiraDataSourceConfig(tcName string, rgName string, projectKey string, apiUrl string, username string, apiToken string, toolName string) string {
	return fmt.Sprintf(`
		data "ibm_resource_group" "resource_group" {
			name = "%s"
		}

		resource "ibm_cd_toolchain" "cd_toolchain" {
			name = "%s"
			resource_group_id = data.ibm_resource_group.resource_group.id
		}

		resource "ibm_cd_toolchain_tool_jira" "cd_toolchain_tool_jira" {
			toolchain_id = ibm_cd_toolchain.cd_toolchain.id
			parameters {
				project_key = "%s"
				api_url = "%s"
				username = "%s"
				enable_traceability = true
				api_token = "%s"
			}
			name = "%s"
		}

		data "ibm_cd_toolchain_tool_jira" "cd_toolchain_tool_jira" {
			toolchain_id = ibm_cd_toolchain_tool_jira.cd_toolchain_tool_jira.toolchain_id
			tool_id = ibm_cd_toolchain_tool_jira.cd_toolchain_tool_jira.tool_id
		}
	`, rgName, tcName, projectKey, apiUrl, username, apiToken, toolName)
}
