// Copyright IBM Corp. 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package cdtoolchain_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"
)

func TestAccIBMCdToolchainToolNexusDataSourceBasic(t *testing.T) {
	tcName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	rgName := acc.CdResourceGroupName

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCdToolchainToolNexusDataSourceConfigBasic(tcName, rgName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_nexus.cd_toolchain_tool_nexus", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_nexus.cd_toolchain_tool_nexus", "toolchain_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_nexus.cd_toolchain_tool_nexus", "tool_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_nexus.cd_toolchain_tool_nexus", "resource_group_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_nexus.cd_toolchain_tool_nexus", "crn"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_nexus.cd_toolchain_tool_nexus", "toolchain_crn"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_nexus.cd_toolchain_tool_nexus", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_nexus.cd_toolchain_tool_nexus", "referent.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_nexus.cd_toolchain_tool_nexus", "updated_at"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_nexus.cd_toolchain_tool_nexus", "parameters.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_nexus.cd_toolchain_tool_nexus", "state"),
				),
			},
		},
	})
}

func TestAccIBMCdToolchainToolNexusDataSourceAllArgs(t *testing.T) {
	tcName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	rgName := acc.CdResourceGroupName
	toolName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCdToolchainToolNexusDataSourceConfig(tcName, rgName, toolName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_nexus.cd_toolchain_tool_nexus", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_nexus.cd_toolchain_tool_nexus", "toolchain_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_nexus.cd_toolchain_tool_nexus", "tool_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_nexus.cd_toolchain_tool_nexus", "resource_group_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_nexus.cd_toolchain_tool_nexus", "crn"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_nexus.cd_toolchain_tool_nexus", "toolchain_crn"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_nexus.cd_toolchain_tool_nexus", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_nexus.cd_toolchain_tool_nexus", "referent.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_nexus.cd_toolchain_tool_nexus", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_nexus.cd_toolchain_tool_nexus", "updated_at"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_nexus.cd_toolchain_tool_nexus", "parameters.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_nexus.cd_toolchain_tool_nexus", "state"),
				),
			},
		},
	})
}

func testAccCheckIBMCdToolchainToolNexusDataSourceConfigBasic(tcName string, rgName string) string {
	return fmt.Sprintf(`
		data "ibm_resource_group" "resource_group" {
			name = "%s"
		}

		resource "ibm_cd_toolchain" "cd_toolchain" {
			name = "%s"
			resource_group_id = data.ibm_resource_group.resource_group.id
		}

		resource "ibm_cd_toolchain_tool_nexus" "cd_toolchain_tool_nexus" {
			toolchain_id = ibm_cd_toolchain.cd_toolchain.id
			parameters {
				name = "my-nexus"
				type = "npm"
				user_id = "<user_id>"
				token = "<token>"
				release_url = "release_url"
				mirror_url = "mirror_url"
				snapshot_url = "snapshot_url"
				server_url = "https://my.nexus.server.com/"
			}
		}

		data "ibm_cd_toolchain_tool_nexus" "cd_toolchain_tool_nexus" {
			toolchain_id = ibm_cd_toolchain_tool_nexus.cd_toolchain_tool_nexus.toolchain_id
			tool_id = ibm_cd_toolchain_tool_nexus.cd_toolchain_tool_nexus.tool_id
		}
	`, rgName, tcName)
}

func testAccCheckIBMCdToolchainToolNexusDataSourceConfig(tcName string, rgName string, toolName string) string {
	return fmt.Sprintf(`
		data "ibm_resource_group" "resource_group" {
			name = "%s"
		}

		resource "ibm_cd_toolchain" "cd_toolchain" {
			name = "%s"
			resource_group_id = data.ibm_resource_group.resource_group.id
		}

		resource "ibm_cd_toolchain_tool_nexus" "cd_toolchain_tool_nexus" {
			toolchain_id = ibm_cd_toolchain.cd_toolchain.id
			parameters {
				name = "my-nexus"
				type = "npm"
				user_id = "<user_id>"
				token = "<token>"
				release_url = "release_url"
				mirror_url = "mirror_url"
				snapshot_url = "snapshot_url"
				server_url = "https://my.nexus.server.com/"
			}
			name = "%s"
		}

		data "ibm_cd_toolchain_tool_nexus" "cd_toolchain_tool_nexus" {
			toolchain_id = ibm_cd_toolchain_tool_nexus.cd_toolchain_tool_nexus.toolchain_id
			tool_id = ibm_cd_toolchain_tool_nexus.cd_toolchain_tool_nexus.tool_id
		}
	`, rgName, tcName, toolName)
}
