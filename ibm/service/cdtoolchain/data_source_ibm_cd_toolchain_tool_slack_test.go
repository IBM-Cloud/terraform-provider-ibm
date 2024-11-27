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

func TestAccIBMCdToolchainToolSlackDataSourceBasic(t *testing.T) {
	tcName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	rgName := acc.CdResourceGroupName
	channelName := acc.CdSlackChannelName
	teamName := acc.CdSlackTeamName
	webhook := acc.CdSlackWebhook

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCdToolchainToolSlackDataSourceConfigBasic(tcName, rgName, channelName, teamName, webhook),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_slack.cd_toolchain_tool_slack", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_slack.cd_toolchain_tool_slack", "toolchain_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_slack.cd_toolchain_tool_slack", "tool_id"),
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
	tcName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	rgName := acc.CdResourceGroupName
	channelName := acc.CdSlackChannelName
	teamName := acc.CdSlackTeamName
	webhook := acc.CdSlackWebhook
	toolName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCdToolchainToolSlackDataSourceConfig(tcName, rgName, channelName, teamName, webhook, toolName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_slack.cd_toolchain_tool_slack", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_slack.cd_toolchain_tool_slack", "toolchain_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_slack.cd_toolchain_tool_slack", "tool_id"),
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

func testAccCheckIBMCdToolchainToolSlackDataSourceConfigBasic(tcName string, rgName string, channelName string, teamName string, webhook string) string {
	return fmt.Sprintf(`
		data "ibm_resource_group" "resource_group" {
			name = "%s"
		}

		resource "ibm_cd_toolchain" "cd_toolchain" {
			name = "%s"
			resource_group_id = data.ibm_resource_group.resource_group.id
		}

		resource "ibm_cd_toolchain_tool_slack" "cd_toolchain_tool_slack" {
			toolchain_id = ibm_cd_toolchain.cd_toolchain.id
			parameters {
				channel_name = "%s"
				pipeline_start = true
				pipeline_success = true
				pipeline_fail = true
				toolchain_bind = true
				toolchain_unbind = true
				webhook = "%s"
				team_name = "%s"
			}
		}

		data "ibm_cd_toolchain_tool_slack" "cd_toolchain_tool_slack" {
			toolchain_id = ibm_cd_toolchain_tool_slack.cd_toolchain_tool_slack.toolchain_id
			tool_id = ibm_cd_toolchain_tool_slack.cd_toolchain_tool_slack.tool_id
		}
	`, rgName, tcName, channelName, webhook, teamName)
}

func testAccCheckIBMCdToolchainToolSlackDataSourceConfig(tcName string, rgName string, channelName string, teamName string, webhook string, toolName string) string {
	return fmt.Sprintf(`
		data "ibm_resource_group" "resource_group" {
			name = "%s"
		}

		resource "ibm_cd_toolchain" "cd_toolchain" {
			name = "%s"
			resource_group_id = data.ibm_resource_group.resource_group.id
		}

		resource "ibm_cd_toolchain_tool_slack" "cd_toolchain_tool_slack" {
			toolchain_id = ibm_cd_toolchain.cd_toolchain.id
			parameters {
				channel_name = "%s"
				pipeline_start = true
				pipeline_success = true
				pipeline_fail = true
				toolchain_bind = true
				toolchain_unbind = true
				webhook = "%s"
				team_name = "%s"
			}
			name = "%s"
		}

		data "ibm_cd_toolchain_tool_slack" "cd_toolchain_tool_slack" {
			toolchain_id = ibm_cd_toolchain_tool_slack.cd_toolchain_tool_slack.toolchain_id
			tool_id = ibm_cd_toolchain_tool_slack.cd_toolchain_tool_slack.tool_id
		}
	`, rgName, tcName, channelName, webhook, teamName, toolName)
}
