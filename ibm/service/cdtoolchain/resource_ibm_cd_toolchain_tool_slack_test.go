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

func TestAccIBMCdToolchainToolSlackBasic(t *testing.T) {
	var conf cdtoolchainv2.ToolchainTool
	rgName := acc.CdResourceGroupName
	tcName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	channelName := acc.CdSlackChannelName
	teamName := acc.CdSlackTeamName
	webhook := acc.CdSlackWebhook

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCdToolchainToolSlackDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCdToolchainToolSlackConfigBasic(tcName, rgName, channelName, teamName, webhook),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCdToolchainToolSlackExists("ibm_cd_toolchain_tool_slack.cd_toolchain_tool_slack", conf),
					resource.TestCheckResourceAttrSet("ibm_cd_toolchain_tool_slack.cd_toolchain_tool_slack", "toolchain_id"),
				),
			},
		},
	})
}

func TestAccIBMCdToolchainToolSlackAllArgs(t *testing.T) {
	var conf cdtoolchainv2.ToolchainTool
	rgName := acc.CdResourceGroupName
	tcName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	channelName := acc.CdSlackChannelName
	teamName := acc.CdSlackTeamName
	webhook := acc.CdSlackWebhook
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	nameUpdate := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCdToolchainToolSlackDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCdToolchainToolSlackConfig(tcName, rgName, channelName, teamName, webhook, name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCdToolchainToolSlackExists("ibm_cd_toolchain_tool_slack.cd_toolchain_tool_slack", conf),
					resource.TestCheckResourceAttrSet("ibm_cd_toolchain_tool_slack.cd_toolchain_tool_slack", "toolchain_id"),
					resource.TestCheckResourceAttr("ibm_cd_toolchain_tool_slack.cd_toolchain_tool_slack", "name", name),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMCdToolchainToolSlackConfig(tcName, rgName, channelName, teamName, webhook, nameUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("ibm_cd_toolchain_tool_slack.cd_toolchain_tool_slack", "toolchain_id"),
					resource.TestCheckResourceAttr("ibm_cd_toolchain_tool_slack.cd_toolchain_tool_slack", "name", nameUpdate),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_cd_toolchain_tool_slack.cd_toolchain_tool_slack",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIBMCdToolchainToolSlackConfigBasic(tcName string, rgName string, channelName string, teamName string, webhook string) string {
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
	`, rgName, tcName, channelName, webhook, teamName)
}

func testAccCheckIBMCdToolchainToolSlackConfig(tcName string, rgName string, channelName string, teamName string, webhook string, name string) string {
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
	`, rgName, tcName, channelName, webhook, teamName, name)
}

func testAccCheckIBMCdToolchainToolSlackExists(n string, obj cdtoolchainv2.ToolchainTool) resource.TestCheckFunc {

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

func testAccCheckIBMCdToolchainToolSlackDestroy(s *terraform.State) error {
	cdToolchainClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).CdToolchainV2()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_cd_toolchain_tool_slack" {
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
			return fmt.Errorf("cd_toolchain_tool_slack still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for cd_toolchain_tool_slack (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
