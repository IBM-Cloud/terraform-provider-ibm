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

func TestAccIBMCdToolchainToolArtifactoryDataSourceBasic(t *testing.T) {
	toolchainToolToolchainID := fmt.Sprintf("tf_toolchain_id_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCdToolchainToolArtifactoryDataSourceConfigBasic(toolchainToolToolchainID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_artifactory.cd_toolchain_tool_artifactory", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_artifactory.cd_toolchain_tool_artifactory", "toolchain_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_artifactory.cd_toolchain_tool_artifactory", "tool_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_artifactory.cd_toolchain_tool_artifactory", "toolchain_tool_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_artifactory.cd_toolchain_tool_artifactory", "resource_group_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_artifactory.cd_toolchain_tool_artifactory", "crn"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_artifactory.cd_toolchain_tool_artifactory", "toolchain_crn"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_artifactory.cd_toolchain_tool_artifactory", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_artifactory.cd_toolchain_tool_artifactory", "referent.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_artifactory.cd_toolchain_tool_artifactory", "updated_at"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_artifactory.cd_toolchain_tool_artifactory", "parameters.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_artifactory.cd_toolchain_tool_artifactory", "state"),
				),
			},
		},
	})
}

func TestAccIBMCdToolchainToolArtifactoryDataSourceAllArgs(t *testing.T) {
	toolchainToolToolchainID := fmt.Sprintf("tf_toolchain_id_%d", acctest.RandIntRange(10, 100))
	toolchainToolName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCdToolchainToolArtifactoryDataSourceConfig(toolchainToolToolchainID, toolchainToolName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_artifactory.cd_toolchain_tool_artifactory", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_artifactory.cd_toolchain_tool_artifactory", "toolchain_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_artifactory.cd_toolchain_tool_artifactory", "tool_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_artifactory.cd_toolchain_tool_artifactory", "toolchain_tool_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_artifactory.cd_toolchain_tool_artifactory", "resource_group_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_artifactory.cd_toolchain_tool_artifactory", "crn"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_artifactory.cd_toolchain_tool_artifactory", "toolchain_crn"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_artifactory.cd_toolchain_tool_artifactory", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_artifactory.cd_toolchain_tool_artifactory", "referent.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_artifactory.cd_toolchain_tool_artifactory", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_artifactory.cd_toolchain_tool_artifactory", "updated_at"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_artifactory.cd_toolchain_tool_artifactory", "parameters.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_artifactory.cd_toolchain_tool_artifactory", "state"),
				),
			},
		},
	})
}

func testAccCheckIBMCdToolchainToolArtifactoryDataSourceConfigBasic(toolchainToolToolchainID string) string {
	return fmt.Sprintf(`
		resource "ibm_cd_toolchain_tool_artifactory" "cd_toolchain_tool_artifactory" {
			toolchain_id = "%s"
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
				api_key = "<api_key>"
			}
		}

		data "ibm_cd_toolchain_tool_artifactory" "cd_toolchain_tool_artifactory" {
			toolchain_id = ibm_cd_toolchain_tool_artifactory.cd_toolchain_tool_artifactory.toolchain_id
			tool_id = "tool_id"
		}
	`, toolchainToolToolchainID)
}

func testAccCheckIBMCdToolchainToolArtifactoryDataSourceConfig(toolchainToolToolchainID string, toolchainToolName string) string {
	return fmt.Sprintf(`
		resource "ibm_cd_toolchain_tool_artifactory" "cd_toolchain_tool_artifactory" {
			toolchain_id = "%s"
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
				api_key = "<api_key>"
			}
			name = "%s"
		}

		data "ibm_cd_toolchain_tool_artifactory" "cd_toolchain_tool_artifactory" {
			toolchain_id = ibm_cd_toolchain_tool_artifactory.cd_toolchain_tool_artifactory.toolchain_id
			tool_id = "tool_id"
		}
	`, toolchainToolToolchainID, toolchainToolName)
}
