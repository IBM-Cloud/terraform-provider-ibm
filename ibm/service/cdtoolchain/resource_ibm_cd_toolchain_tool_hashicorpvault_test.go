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

func TestAccIBMCdToolchainToolHashicorpvaultBasic(t *testing.T) {
	var conf cdtoolchainv2.ToolchainTool
	rgName := acc.CdResourceGroupName
	tcName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCdToolchainToolHashicorpvaultDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCdToolchainToolHashicorpvaultConfigBasic(tcName, rgName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCdToolchainToolHashicorpvaultExists("ibm_cd_toolchain_tool_hashicorpvault.cd_toolchain_tool_hashicorpvault", conf),
					resource.TestCheckResourceAttrSet("ibm_cd_toolchain_tool_hashicorpvault.cd_toolchain_tool_hashicorpvault", "toolchain_id"),
				),
			},
		},
	})
}

func TestAccIBMCdToolchainToolHashicorpvaultAllArgs(t *testing.T) {
	var conf cdtoolchainv2.ToolchainTool
	rgName := acc.CdResourceGroupName
	tcName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	nameUpdate := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCdToolchainToolHashicorpvaultDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCdToolchainToolHashicorpvaultConfig(tcName, rgName, name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCdToolchainToolHashicorpvaultExists("ibm_cd_toolchain_tool_hashicorpvault.cd_toolchain_tool_hashicorpvault", conf),
					resource.TestCheckResourceAttrSet("ibm_cd_toolchain_tool_hashicorpvault.cd_toolchain_tool_hashicorpvault", "toolchain_id"),
					resource.TestCheckResourceAttr("ibm_cd_toolchain_tool_hashicorpvault.cd_toolchain_tool_hashicorpvault", "name", name),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMCdToolchainToolHashicorpvaultConfig(tcName, rgName, nameUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("ibm_cd_toolchain_tool_hashicorpvault.cd_toolchain_tool_hashicorpvault", "toolchain_id"),
					resource.TestCheckResourceAttr("ibm_cd_toolchain_tool_hashicorpvault.cd_toolchain_tool_hashicorpvault", "name", nameUpdate),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_cd_toolchain_tool_hashicorpvault.cd_toolchain_tool_hashicorpvault",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIBMCdToolchainToolHashicorpvaultConfigBasic(tcName string, rgName string) string {
	return fmt.Sprintf(`
		data "ibm_resource_group" "resource_group" {
			name = "%s"
		}

		resource "ibm_cd_toolchain" "cd_toolchain" {
			name = "%s"
			resource_group_id = data.ibm_resource_group.resource_group.id
		}

		resource "ibm_cd_toolchain_tool_hashicorpvault" "cd_toolchain_tool_hashicorpvault" {
			toolchain_id = ibm_cd_toolchain.cd_toolchain.id
			parameters {
				name = "hcv_tool_01"
				server_url = "https://hcv.mycompany.example.com:8200"
				authentication_method = "approle"
				token = "token"
				role_id = "<role_id>"
				secret_id = "<secret_id>"
				dashboard_url = "https://hcv.mycompany.example.com:8200/ui"
				path = "generic/project/test_project"
				secret_filter = "secret_filter"
				default_secret = "default_secret"
				username = "username"
				password = "password"
			}
		}
	`, rgName, tcName)
}

func testAccCheckIBMCdToolchainToolHashicorpvaultConfig(tcName string, rgName string, name string) string {
	return fmt.Sprintf(`
		data "ibm_resource_group" "resource_group" {
			name = "%s"
		}

		resource "ibm_cd_toolchain" "cd_toolchain" {
			name = "%s"
			resource_group_id = data.ibm_resource_group.resource_group.id
		}

		resource "ibm_cd_toolchain_tool_hashicorpvault" "cd_toolchain_tool_hashicorpvault" {
			toolchain_id = ibm_cd_toolchain.cd_toolchain.id
			parameters {
				name = "hcv_tool_01"
				server_url = "https://hcv.mycompany.example.com:8200"
				authentication_method = "approle"
				token = "token"
				role_id = "<role_id>"
				secret_id = "<secret_id>"
				dashboard_url = "https://hcv.mycompany.example.com:8200/ui"
				path = "generic/project/test_project"
				secret_filter = "secret_filter"
				default_secret = "default_secret"
				username = "username"
				password = "password"
			}
			name = "%s"
		}
	`, rgName, tcName, name)
}

func testAccCheckIBMCdToolchainToolHashicorpvaultExists(n string, obj cdtoolchainv2.ToolchainTool) resource.TestCheckFunc {

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

func testAccCheckIBMCdToolchainToolHashicorpvaultDestroy(s *terraform.State) error {
	cdToolchainClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).CdToolchainV2()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_cd_toolchain_tool_hashicorpvault" {
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
			return fmt.Errorf("cd_toolchain_tool_hashicorpvault still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for cd_toolchain_tool_hashicorpvault (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
