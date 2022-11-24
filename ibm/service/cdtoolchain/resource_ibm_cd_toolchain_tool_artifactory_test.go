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
	"github.com/IBM/continuous-delivery-go-sdk/cdtoolchainv2"
)

func TestAccIBMCdToolchainToolArtifactoryBasic(t *testing.T) {
	var conf cdtoolchainv2.ToolchainTool
	rgName := acc.CdResourceGroupName
	tcName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCdToolchainToolArtifactoryDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCdToolchainToolArtifactoryConfigBasic(tcName, rgName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCdToolchainToolArtifactoryExists("ibm_cd_toolchain_tool_artifactory.cd_toolchain_tool_artifactory", conf),
					resource.TestCheckResourceAttrSet("ibm_cd_toolchain_tool_artifactory.cd_toolchain_tool_artifactory", "toolchain_id"),
				),
			},
		},
	})
}

func TestAccIBMCdToolchainToolArtifactoryAllArgs(t *testing.T) {
	var conf cdtoolchainv2.ToolchainTool
	rgName := acc.CdResourceGroupName
	tcName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	nameUpdate := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCdToolchainToolArtifactoryDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCdToolchainToolArtifactoryConfig(tcName, rgName, name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCdToolchainToolArtifactoryExists("ibm_cd_toolchain_tool_artifactory.cd_toolchain_tool_artifactory", conf),
					resource.TestCheckResourceAttrSet("ibm_cd_toolchain_tool_artifactory.cd_toolchain_tool_artifactory", "toolchain_id"),
					resource.TestCheckResourceAttr("ibm_cd_toolchain_tool_artifactory.cd_toolchain_tool_artifactory", "name", name),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMCdToolchainToolArtifactoryConfig(tcName, rgName, nameUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("ibm_cd_toolchain_tool_artifactory.cd_toolchain_tool_artifactory", "toolchain_id"),
					resource.TestCheckResourceAttr("ibm_cd_toolchain_tool_artifactory.cd_toolchain_tool_artifactory", "name", nameUpdate),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_cd_toolchain_tool_artifactory.cd_toolchain_tool_artifactory",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIBMCdToolchainToolArtifactoryConfigBasic(tcName string, rgName string) string {
	return fmt.Sprintf(`
		data "ibm_resource_group" "resource_group" {
			name = "%s"
		}

		resource "ibm_cd_toolchain" "cd_toolchain" {
			name = "%s"
			resource_group_id = data.ibm_resource_group.resource_group.id
		}

		resource "ibm_cd_toolchain_tool_artifactory" "cd_toolchain_tool_artifactory" {
			toolchain_id = ibm_cd_toolchain.cd_toolchain.id
			parameters {
				name = "artifactory-tool-01"
				dashboard_url = "https://mycompany.example.jfrog.io"
				type = "docker"
				user_id = "<user_id>"
				release_url = "release_url"
				mirror_url = "mirror_url"
				snapshot_url = "snapshot_url"
				repository_name = "default-docker-local"
				repository_url = "https://mycompany.example.jfrog.io/artifactory/default-docker-local"
				token = "<token>"
			}
		}
	`, rgName, tcName)
}

func testAccCheckIBMCdToolchainToolArtifactoryConfig(tcName string, rgName string, name string) string {
	return fmt.Sprintf(`
		data "ibm_resource_group" "resource_group" {
			name = "%s"
		}

		resource "ibm_cd_toolchain" "cd_toolchain" {
			name = "%s"
			resource_group_id = data.ibm_resource_group.resource_group.id
		}

		resource "ibm_cd_toolchain_tool_artifactory" "cd_toolchain_tool_artifactory" {
			toolchain_id = ibm_cd_toolchain.cd_toolchain.id
			parameters {
				name = "artifactory-tool-01"
				dashboard_url = "https://mycompany.example.jfrog.io"
				type = "docker"
				user_id = "<user_id>"
				release_url = "release_url"
				mirror_url = "mirror_url"
				snapshot_url = "snapshot_url"
				repository_name = "default-docker-local"
				repository_url = "https://mycompany.example.jfrog.io/artifactory/default-docker-local"
				token = "<token>"
			}
			name = "%s"
		}
	`, rgName, tcName, name)
}

func testAccCheckIBMCdToolchainToolArtifactoryExists(n string, obj cdtoolchainv2.ToolchainTool) resource.TestCheckFunc {

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

func testAccCheckIBMCdToolchainToolArtifactoryDestroy(s *terraform.State) error {
	cdToolchainClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).CdToolchainV2()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_cd_toolchain_tool_artifactory" {
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
			return fmt.Errorf("cd_toolchain_tool_artifactory still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for cd_toolchain_tool_artifactory (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
