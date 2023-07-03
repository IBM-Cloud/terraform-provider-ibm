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
	tcName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	rgName := acc.CdResourceGroupName

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCdToolchainToolArtifactoryDataSourceConfigBasic(tcName, rgName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_artifactory.cd_toolchain_tool_artifactory", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_artifactory.cd_toolchain_tool_artifactory", "toolchain_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_artifactory.cd_toolchain_tool_artifactory", "tool_id"),
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
	tcName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	rgName := acc.CdResourceGroupName
	toolName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCdToolchainToolArtifactoryDataSourceConfig(tcName, rgName, toolName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_artifactory.cd_toolchain_tool_artifactory", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_artifactory.cd_toolchain_tool_artifactory", "toolchain_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_artifactory.cd_toolchain_tool_artifactory", "tool_id"),
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

func testAccCheckIBMCdToolchainToolArtifactoryDataSourceConfigBasic(tcName string, rgName string) string {
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

		data "ibm_cd_toolchain_tool_artifactory" "cd_toolchain_tool_artifactory" {
			toolchain_id = ibm_cd_toolchain_tool_artifactory.cd_toolchain_tool_artifactory.toolchain_id
			tool_id = ibm_cd_toolchain_tool_artifactory.cd_toolchain_tool_artifactory.tool_id
		}
	`, rgName, tcName)
}

func testAccCheckIBMCdToolchainToolArtifactoryDataSourceConfig(tcName string, rgName string, toolName string) string {
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

		data "ibm_cd_toolchain_tool_artifactory" "cd_toolchain_tool_artifactory" {
			toolchain_id = ibm_cd_toolchain_tool_artifactory.cd_toolchain_tool_artifactory.toolchain_id
			tool_id = ibm_cd_toolchain_tool_artifactory.cd_toolchain_tool_artifactory.tool_id
		}
	`, rgName, tcName, toolName)
}
