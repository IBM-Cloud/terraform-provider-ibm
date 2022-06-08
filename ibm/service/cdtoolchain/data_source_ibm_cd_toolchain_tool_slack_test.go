// Copyright IBM Corp. 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package cdtoolchain_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIBMCdToolchainToolSlackDataSourceBasic(t *testing.T) {
	getToolByIDResponseToolchainID := fmt.Sprintf("tf_toolchain_id_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCdToolchainToolSlackDataSourceConfigBasic(getToolByIDResponseToolchainID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_slack.cd_toolchain_tool_slack", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_slack.cd_toolchain_tool_slack", "toolchain_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_slack.cd_toolchain_tool_slack", "tool_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_slack.cd_toolchain_tool_slack", "get_tool_by_id_response_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_slack.cd_toolchain_tool_slack", "resource_group_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_slack.cd_toolchain_tool_slack", "crn"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_slack.cd_toolchain_tool_slack", "toolchain_crn"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_slack.cd_toolchain_tool_slack", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_slack.cd_toolchain_tool_slack", "referent.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_slack.cd_toolchain_tool_slack", "updated_at"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_slack.cd_toolchain_tool_slack", "parameters.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_slack.cd_toolchain_tool_slack", "state"),
				),
			},
		},
	})
}

func TestAccIBMCdToolchainToolSlackDataSourceAllArgs(t *testing.T) {
	getToolByIDResponseToolchainID := fmt.Sprintf("tf_toolchain_id_%d", acctest.RandIntRange(10, 100))
	getToolByIDResponseName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCdToolchainToolSlackDataSourceConfig(getToolByIDResponseToolchainID, getToolByIDResponseName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_slack.cd_toolchain_tool_slack", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_slack.cd_toolchain_tool_slack", "toolchain_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_slack.cd_toolchain_tool_slack", "tool_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_slack.cd_toolchain_tool_slack", "get_tool_by_id_response_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_slack.cd_toolchain_tool_slack", "resource_group_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_slack.cd_toolchain_tool_slack", "crn"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_slack.cd_toolchain_tool_slack", "toolchain_crn"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_slack.cd_toolchain_tool_slack", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_slack.cd_toolchain_tool_slack", "referent.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_slack.cd_toolchain_tool_slack", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_slack.cd_toolchain_tool_slack", "updated_at"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_slack.cd_toolchain_tool_slack", "parameters.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_slack.cd_toolchain_tool_slack", "state"),
				),
			},
		},
	})
}

func testAccCheckIBMCdToolchainToolSlackDataSourceConfigBasic(getToolByIDResponseToolchainID string) string {
	return fmt.Sprintf(`
		resource "ibm_cd_toolchain_tool_slack" "cd_toolchain_tool_slack" {
			toolchain_id = "%s"
		}

		data "ibm_cd_toolchain_tool_slack" "cd_toolchain_tool_slack" {
			toolchain_id = ibm_cd_toolchain_tool_slack.cd_toolchain_tool_slack.toolchain_id
			tool_id = "tool_id"
		}
	`, getToolByIDResponseToolchainID)
}

func testAccCheckIBMCdToolchainToolSlackDataSourceConfig(getToolByIDResponseToolchainID string, getToolByIDResponseName string) string {
	return fmt.Sprintf(`
		resource "ibm_cd_toolchain_tool_slack" "cd_toolchain_tool_slack" {
			toolchain_id = "%s"
			name = "%s"
			parameters {
				api_token = "api_token"
				channel_name = "channel_name"
				team_url = "team_url"
				pipeline_start = true
				pipeline_success = true
				pipeline_fail = true
				toolchain_bind = true
				toolchain_unbind = true
			}
		}

		data "ibm_cd_toolchain_tool_slack" "cd_toolchain_tool_slack" {
			toolchain_id = ibm_cd_toolchain_tool_slack.cd_toolchain_tool_slack.toolchain_id
			tool_id = "tool_id"
		}
	`, getToolByIDResponseToolchainID, getToolByIDResponseName)
}
