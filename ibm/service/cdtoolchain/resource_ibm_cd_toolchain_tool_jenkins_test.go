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

func TestAccIBMCdToolchainToolJenkinsBasic(t *testing.T) {
	var conf cdtoolchainv2.ToolchainTool
	rgName := acc.CdResourceGroupName
	tcName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCdToolchainToolJenkinsDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCdToolchainToolJenkinsConfigBasic(tcName, rgName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCdToolchainToolJenkinsExists("ibm_cd_toolchain_tool_jenkins.cd_toolchain_tool_jenkins", conf),
					resource.TestCheckResourceAttrSet("ibm_cd_toolchain_tool_jenkins.cd_toolchain_tool_jenkins", "toolchain_id"),
				),
			},
		},
	})
}

func TestAccIBMCdToolchainToolJenkinsAllArgs(t *testing.T) {
	var conf cdtoolchainv2.ToolchainTool
	rgName := acc.CdResourceGroupName
	tcName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	nameUpdate := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCdToolchainToolJenkinsDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCdToolchainToolJenkinsConfig(tcName, rgName, name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCdToolchainToolJenkinsExists("ibm_cd_toolchain_tool_jenkins.cd_toolchain_tool_jenkins", conf),
					resource.TestCheckResourceAttrSet("ibm_cd_toolchain_tool_jenkins.cd_toolchain_tool_jenkins", "toolchain_id"),
					resource.TestCheckResourceAttr("ibm_cd_toolchain_tool_jenkins.cd_toolchain_tool_jenkins", "name", name),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMCdToolchainToolJenkinsConfig(tcName, rgName, nameUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("ibm_cd_toolchain_tool_jenkins.cd_toolchain_tool_jenkins", "toolchain_id"),
					resource.TestCheckResourceAttr("ibm_cd_toolchain_tool_jenkins.cd_toolchain_tool_jenkins", "name", nameUpdate),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_cd_toolchain_tool_jenkins.cd_toolchain_tool_jenkins",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIBMCdToolchainToolJenkinsConfigBasic(tcName string, rgName string) string {
	return fmt.Sprintf(`
		data "ibm_resource_group" "resource_group" {
			name = "%s"
		}

		resource "ibm_cd_toolchain" "cd_toolchain" {
			name = "%s"
			resource_group_id = data.ibm_resource_group.resource_group.id
		}

		resource "ibm_cd_toolchain_tool_jenkins" "cd_toolchain_tool_jenkins" {
			toolchain_id = ibm_cd_toolchain.cd_toolchain.id
			parameters {
				name = "jenkins-tool-01"
				dashboard_url = "https://jenkins.mycompany.example.com"
				api_user_name = "<api_user_name>"
				api_token = "<api_token>"
			}
		}
	`, rgName, tcName)
}

func testAccCheckIBMCdToolchainToolJenkinsConfig(tcName string, rgName string, name string) string {
	return fmt.Sprintf(`
		data "ibm_resource_group" "resource_group" {
			name = "%s"
		}

		resource "ibm_cd_toolchain" "cd_toolchain" {
			name = "%s"
			resource_group_id = data.ibm_resource_group.resource_group.id
		}

		resource "ibm_cd_toolchain_tool_jenkins" "cd_toolchain_tool_jenkins" {
			toolchain_id = ibm_cd_toolchain.cd_toolchain.id
			parameters {
				name = "jenkins-tool-01"
				dashboard_url = "https://jenkins.mycompany.example.com"
				api_user_name = "<api_user_name>"
				api_token = "<api_token>"
			}
			name = "%s"
		}
	`, rgName, tcName, name)
}

func testAccCheckIBMCdToolchainToolJenkinsExists(n string, obj cdtoolchainv2.ToolchainTool) resource.TestCheckFunc {

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

func testAccCheckIBMCdToolchainToolJenkinsDestroy(s *terraform.State) error {
	cdToolchainClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).CdToolchainV2()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_cd_toolchain_tool_jenkins" {
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
			return fmt.Errorf("cd_toolchain_tool_jenkins still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for cd_toolchain_tool_jenkins (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
