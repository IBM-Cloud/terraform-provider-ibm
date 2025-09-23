// Copyright IBM Corp. 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package cdtoolchain_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/IBM/continuous-delivery-go-sdk/v2/cdtoolchainv2"
	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/conns"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/flex"
)

func TestAccIBMCdToolchainToolJiraBasic(t *testing.T) {
	var conf cdtoolchainv2.ToolchainTool
	rgName := acc.CdResourceGroupName
	tcName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	projectKey := acc.CdJiraProjectKey
	apiUrl := acc.CdJiraApiUrl
	username := acc.CdJiraUsername
	apiToken := acc.CdJiraApiToken

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCdToolchainToolJiraDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCdToolchainToolJiraConfigBasic(tcName, rgName, projectKey, apiUrl, username, apiToken),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCdToolchainToolJiraExists("ibm_cd_toolchain_tool_jira.cd_toolchain_tool_jira", conf),
					resource.TestCheckResourceAttrSet("ibm_cd_toolchain_tool_jira.cd_toolchain_tool_jira", "toolchain_id"),
				),
			},
		},
	})
}

func TestAccIBMCdToolchainToolJiraAllArgs(t *testing.T) {
	var conf cdtoolchainv2.ToolchainTool
	rgName := acc.CdResourceGroupName
	tcName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	nameUpdate := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	projectKey := acc.CdJiraProjectKey
	apiUrl := acc.CdJiraApiUrl
	username := acc.CdJiraUsername
	apiToken := acc.CdJiraApiToken

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCdToolchainToolJiraDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCdToolchainToolJiraConfig(tcName, rgName, projectKey, apiUrl, username, apiToken, name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCdToolchainToolJiraExists("ibm_cd_toolchain_tool_jira.cd_toolchain_tool_jira", conf),
					resource.TestCheckResourceAttrSet("ibm_cd_toolchain_tool_jira.cd_toolchain_tool_jira", "toolchain_id"),
					resource.TestCheckResourceAttr("ibm_cd_toolchain_tool_jira.cd_toolchain_tool_jira", "name", name),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMCdToolchainToolJiraConfig(tcName, rgName, projectKey, apiUrl, username, apiToken, nameUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("ibm_cd_toolchain_tool_jira.cd_toolchain_tool_jira", "toolchain_id"),
					resource.TestCheckResourceAttr("ibm_cd_toolchain_tool_jira.cd_toolchain_tool_jira", "name", nameUpdate),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_cd_toolchain_tool_jira.cd_toolchain_tool_jira",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIBMCdToolchainToolJiraConfigBasic(tcName string, rgName string, projectKey string, apiUrl string, username string, apiToken string) string {
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
	`, rgName, tcName, projectKey, apiUrl, username, apiToken)
}

func testAccCheckIBMCdToolchainToolJiraConfig(tcName string, rgName string, projectKey string, apiUrl string, username string, apiToken string, name string) string {
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
	`, rgName, tcName, projectKey, apiUrl, username, apiToken, name)
}

func testAccCheckIBMCdToolchainToolJiraExists(n string, obj cdtoolchainv2.ToolchainTool) resource.TestCheckFunc {

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

func testAccCheckIBMCdToolchainToolJiraDestroy(s *terraform.State) error {
	cdToolchainClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).CdToolchainV2()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_cd_toolchain_tool_jira" {
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
			return fmt.Errorf("cd_toolchain_tool_jira still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for cd_toolchain_tool_jira (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
