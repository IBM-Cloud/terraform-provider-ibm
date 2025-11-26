// Copyright IBM Corp. 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package cdtoolchain_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/continuous-delivery-go-sdk/v2/cdtoolchainv2"
)

func TestAccIBMCdToolchainToolAppconfigBasic(t *testing.T) {
	var conf cdtoolchainv2.ToolchainTool
	rgName := acc.CdResourceGroupName
	tcName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	acName := acc.CdAppConfigInstanceName

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCdToolchainToolAppconfigDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCdToolchainToolAppconfigConfigBasic(tcName, rgName, acName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCdToolchainToolAppconfigExists("ibm_cd_toolchain_tool_appconfig.cd_toolchain_tool_appconfig", conf),
					resource.TestCheckResourceAttrSet("ibm_cd_toolchain_tool_appconfig.cd_toolchain_tool_appconfig", "toolchain_id"),
				),
			},
		},
	})
}

func TestAccIBMCdToolchainToolAppconfigAllArgs(t *testing.T) {
	var conf cdtoolchainv2.ToolchainTool
	rgName := acc.CdResourceGroupName
	tcName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	acName := acc.CdAppConfigInstanceName
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	nameUpdate := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCdToolchainToolAppconfigDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCdToolchainToolAppconfigConfig(tcName, rgName, acName, name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCdToolchainToolAppconfigExists("ibm_cd_toolchain_tool_appconfig.cd_toolchain_tool_appconfig", conf),
					resource.TestCheckResourceAttrSet("ibm_cd_toolchain_tool_appconfig.cd_toolchain_tool_appconfig", "toolchain_id"),
					resource.TestCheckResourceAttr("ibm_cd_toolchain_tool_appconfig.cd_toolchain_tool_appconfig", "name", name),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMCdToolchainToolAppconfigConfig(tcName, rgName, acName, nameUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("ibm_cd_toolchain_tool_appconfig.cd_toolchain_tool_appconfig", "toolchain_id"),
					resource.TestCheckResourceAttr("ibm_cd_toolchain_tool_appconfig.cd_toolchain_tool_appconfig", "name", nameUpdate),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_cd_toolchain_tool_appconfig.cd_toolchain_tool_appconfig",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIBMCdToolchainToolAppconfigConfigBasic(tcName string, rgName string, acName string) string {
	return fmt.Sprintf(`
		data "ibm_resource_group" "resource_group" {
			name = "%s"
		}
	
		resource "ibm_cd_toolchain" "cd_toolchain" {
			name = "%s"
			resource_group_id = data.ibm_resource_group.resource_group.id
		}

		data "ibm_resource_instance" "ac_resource_instance" {
			name              = "%s"
			location		  = "us-south"
			resource_group_id = data.ibm_resource_group.resource_group.id
			service           = "apprapp"
		}

		resource "ibm_cd_toolchain_tool_appconfig" "cd_toolchain_tool_appconfig" {
			toolchain_id = ibm_cd_toolchain.cd_toolchain.id
			parameters {
				name = "appconfig_tool_01"
				location = "us-south"
				resource_group_name = "%s"
				instance_id = data.ibm_resource_instance.ac_resource_instance.guid
				environment_id = "environment_01"
				collection_id = "collection_01"
			}
		}
	`, rgName, tcName, acName, rgName)
}

func testAccCheckIBMCdToolchainToolAppconfigConfig(tcName string, rgName string, acName string, name string) string {
	return fmt.Sprintf(`
		data "ibm_resource_group" "resource_group" {
			name = "%s"
		}
	
		resource "ibm_cd_toolchain" "cd_toolchain" {
			name = "%s"
			resource_group_id = data.ibm_resource_group.resource_group.id
		}

		data "ibm_resource_instance" "ac_resource_instance" {
			name              = "%s"
			location		  = "us-south"
			resource_group_id = data.ibm_resource_group.resource_group.id
			service           = "apprapp"
		}

		resource "ibm_cd_toolchain_tool_appconfig" "cd_toolchain_tool_appconfig" {
			toolchain_id = ibm_cd_toolchain.cd_toolchain.id
			parameters {
				name = "appconfig_tool_01"
				location = "us-south"
				resource_group_name = "%s"
				instance_id = data.ibm_resource_instance.ac_resource_instance.guid
				environment_id = "environment_01"
				collection_id = "collection_01"
			}
			name = "%s"
		}
	`, rgName, tcName, acName, rgName, name)
}

func testAccCheckIBMCdToolchainToolAppconfigExists(n string, obj cdtoolchainv2.ToolchainTool) resource.TestCheckFunc {

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

func testAccCheckIBMCdToolchainToolAppconfigDestroy(s *terraform.State) error {
	cdToolchainClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).CdToolchainV2()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_cd_toolchain_tool_appconfig" {
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
			return fmt.Errorf("cd_toolchain_tool_appconfig still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for cd_toolchain_tool_appconfig (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
