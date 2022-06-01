// Copyright IBM Corp. 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package toolchain_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIBMCdToolchainToolJiraDataSourceBasic(t *testing.T) {
	getIntegrationByIDResponseToolchainID := fmt.Sprintf("tf_toolchain_id_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCdToolchainToolJiraDataSourceConfigBasic(getIntegrationByIDResponseToolchainID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_jira.cd_toolchain_tool_jira", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_jira.cd_toolchain_tool_jira", "toolchain_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_jira.cd_toolchain_tool_jira", "integration_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_jira.cd_toolchain_tool_jira", "get_integration_by_id_response_id"),
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
	getIntegrationByIDResponseToolchainID := fmt.Sprintf("tf_toolchain_id_%d", acctest.RandIntRange(10, 100))
	getIntegrationByIDResponseName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCdToolchainToolJiraDataSourceConfig(getIntegrationByIDResponseToolchainID, getIntegrationByIDResponseName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_jira.cd_toolchain_tool_jira", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_jira.cd_toolchain_tool_jira", "toolchain_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_jira.cd_toolchain_tool_jira", "integration_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_jira.cd_toolchain_tool_jira", "get_integration_by_id_response_id"),
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

func testAccCheckIBMCdToolchainToolJiraDataSourceConfigBasic(getIntegrationByIDResponseToolchainID string) string {
	return fmt.Sprintf(`
		resource "ibm_cd_toolchain_tool_jira" "cd_toolchain_tool_jira" {
			toolchain_id = "%s"
		}

		data "ibm_cd_toolchain_tool_jira" "cd_toolchain_tool_jira" {
			toolchain_id = ibm_cd_toolchain_tool_jira.cd_toolchain_tool_jira.toolchain_id
			integration_id = "integration_id"
		}
	`, getIntegrationByIDResponseToolchainID)
}

func testAccCheckIBMCdToolchainToolJiraDataSourceConfig(getIntegrationByIDResponseToolchainID string, getIntegrationByIDResponseName string) string {
	return fmt.Sprintf(`
		resource "ibm_cd_toolchain_tool_jira" "cd_toolchain_tool_jira" {
			toolchain_id = "%s"
			name = "%s"
			parameters {
				type = "new"
				project_key = "project_key"
				project_name = "project_name"
				project_admin = "project_admin"
				api_url = "api_url"
				username = "username"
				password = "password"
				enable_traceability = true
			}
		}

		data "ibm_cd_toolchain_tool_jira" "cd_toolchain_tool_jira" {
			toolchain_id = ibm_cd_toolchain_tool_jira.cd_toolchain_tool_jira.toolchain_id
			integration_id = "integration_id"
		}
	`, getIntegrationByIDResponseToolchainID, getIntegrationByIDResponseName)
}
