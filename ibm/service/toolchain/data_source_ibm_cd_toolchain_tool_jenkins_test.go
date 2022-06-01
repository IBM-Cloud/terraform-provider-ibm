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

func TestAccIBMCdToolchainToolJenkinsDataSourceBasic(t *testing.T) {
	getIntegrationByIDResponseToolchainID := fmt.Sprintf("tf_toolchain_id_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCdToolchainToolJenkinsDataSourceConfigBasic(getIntegrationByIDResponseToolchainID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_jenkins.cd_toolchain_tool_jenkins", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_jenkins.cd_toolchain_tool_jenkins", "toolchain_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_jenkins.cd_toolchain_tool_jenkins", "integration_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_jenkins.cd_toolchain_tool_jenkins", "get_integration_by_id_response_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_jenkins.cd_toolchain_tool_jenkins", "resource_group_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_jenkins.cd_toolchain_tool_jenkins", "crn"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_jenkins.cd_toolchain_tool_jenkins", "toolchain_crn"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_jenkins.cd_toolchain_tool_jenkins", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_jenkins.cd_toolchain_tool_jenkins", "referent.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_jenkins.cd_toolchain_tool_jenkins", "updated_at"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_jenkins.cd_toolchain_tool_jenkins", "parameters.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_jenkins.cd_toolchain_tool_jenkins", "state"),
				),
			},
		},
	})
}

func TestAccIBMCdToolchainToolJenkinsDataSourceAllArgs(t *testing.T) {
	getIntegrationByIDResponseToolchainID := fmt.Sprintf("tf_toolchain_id_%d", acctest.RandIntRange(10, 100))
	getIntegrationByIDResponseName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCdToolchainToolJenkinsDataSourceConfig(getIntegrationByIDResponseToolchainID, getIntegrationByIDResponseName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_jenkins.cd_toolchain_tool_jenkins", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_jenkins.cd_toolchain_tool_jenkins", "toolchain_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_jenkins.cd_toolchain_tool_jenkins", "integration_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_jenkins.cd_toolchain_tool_jenkins", "get_integration_by_id_response_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_jenkins.cd_toolchain_tool_jenkins", "resource_group_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_jenkins.cd_toolchain_tool_jenkins", "crn"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_jenkins.cd_toolchain_tool_jenkins", "toolchain_crn"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_jenkins.cd_toolchain_tool_jenkins", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_jenkins.cd_toolchain_tool_jenkins", "referent.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_jenkins.cd_toolchain_tool_jenkins", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_jenkins.cd_toolchain_tool_jenkins", "updated_at"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_jenkins.cd_toolchain_tool_jenkins", "parameters.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_jenkins.cd_toolchain_tool_jenkins", "state"),
				),
			},
		},
	})
}

func testAccCheckIBMCdToolchainToolJenkinsDataSourceConfigBasic(getIntegrationByIDResponseToolchainID string) string {
	return fmt.Sprintf(`
		resource "ibm_cd_toolchain_tool_jenkins" "cd_toolchain_tool_jenkins" {
			toolchain_id = "%s"
		}

		data "ibm_cd_toolchain_tool_jenkins" "cd_toolchain_tool_jenkins" {
			toolchain_id = ibm_cd_toolchain_tool_jenkins.cd_toolchain_tool_jenkins.toolchain_id
			integration_id = "integration_id"
		}
	`, getIntegrationByIDResponseToolchainID)
}

func testAccCheckIBMCdToolchainToolJenkinsDataSourceConfig(getIntegrationByIDResponseToolchainID string, getIntegrationByIDResponseName string) string {
	return fmt.Sprintf(`
		resource "ibm_cd_toolchain_tool_jenkins" "cd_toolchain_tool_jenkins" {
			toolchain_id = "%s"
			name = "%s"
			parameters {
				name = "name"
				dashboard_url = "dashboard_url"
				webhook_url = "webhook_url"
				api_user_name = "api_user_name"
				api_token = "api_token"
			}
		}

		data "ibm_cd_toolchain_tool_jenkins" "cd_toolchain_tool_jenkins" {
			toolchain_id = ibm_cd_toolchain_tool_jenkins.cd_toolchain_tool_jenkins.toolchain_id
			integration_id = "integration_id"
		}
	`, getIntegrationByIDResponseToolchainID, getIntegrationByIDResponseName)
}
