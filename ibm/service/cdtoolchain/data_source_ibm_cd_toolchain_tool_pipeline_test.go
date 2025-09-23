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

func TestAccIBMCdToolchainToolPipelineDataSourceBasic(t *testing.T) {
	tcName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	rgName := acc.CdResourceGroupName

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCdToolchainToolPipelineDataSourceConfigBasic(tcName, rgName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_pipeline.cd_toolchain_tool_pipeline", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_pipeline.cd_toolchain_tool_pipeline", "toolchain_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_pipeline.cd_toolchain_tool_pipeline", "tool_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_pipeline.cd_toolchain_tool_pipeline", "resource_group_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_pipeline.cd_toolchain_tool_pipeline", "crn"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_pipeline.cd_toolchain_tool_pipeline", "toolchain_crn"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_pipeline.cd_toolchain_tool_pipeline", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_pipeline.cd_toolchain_tool_pipeline", "referent.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_pipeline.cd_toolchain_tool_pipeline", "updated_at"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_pipeline.cd_toolchain_tool_pipeline", "parameters.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_pipeline.cd_toolchain_tool_pipeline", "state"),
				),
			},
		},
	})
}

func TestAccIBMCdToolchainToolPipelineDataSourceAllArgs(t *testing.T) {
	tcName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	rgName := acc.CdResourceGroupName
	toolName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCdToolchainToolPipelineDataSourceConfig(tcName, rgName, toolName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_pipeline.cd_toolchain_tool_pipeline", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_pipeline.cd_toolchain_tool_pipeline", "toolchain_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_pipeline.cd_toolchain_tool_pipeline", "tool_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_pipeline.cd_toolchain_tool_pipeline", "resource_group_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_pipeline.cd_toolchain_tool_pipeline", "crn"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_pipeline.cd_toolchain_tool_pipeline", "toolchain_crn"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_pipeline.cd_toolchain_tool_pipeline", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_pipeline.cd_toolchain_tool_pipeline", "referent.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_pipeline.cd_toolchain_tool_pipeline", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_pipeline.cd_toolchain_tool_pipeline", "updated_at"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_pipeline.cd_toolchain_tool_pipeline", "parameters.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_pipeline.cd_toolchain_tool_pipeline", "state"),
				),
			},
		},
	})
}

func testAccCheckIBMCdToolchainToolPipelineDataSourceConfigBasic(tcName string, rgName string) string {
	return fmt.Sprintf(`
		data "ibm_resource_group" "resource_group" {
			name = "%s"
		}

		resource "ibm_cd_toolchain" "cd_toolchain" {
			name = "%s"
			resource_group_id = data.ibm_resource_group.resource_group.id
		}

		resource "ibm_cd_toolchain_tool_pipeline" "cd_toolchain_tool_pipeline" {
			toolchain_id = ibm_cd_toolchain.cd_toolchain.id
			parameters {
				name = "pipeline-tool-01"
			}
		}

		data "ibm_cd_toolchain_tool_pipeline" "cd_toolchain_tool_pipeline" {
			toolchain_id = ibm_cd_toolchain_tool_pipeline.cd_toolchain_tool_pipeline.toolchain_id
			tool_id = ibm_cd_toolchain_tool_pipeline.cd_toolchain_tool_pipeline.tool_id
		}
	`, rgName, tcName)
}

func testAccCheckIBMCdToolchainToolPipelineDataSourceConfig(tcName string, rgName string, toolName string) string {
	return fmt.Sprintf(`
		data "ibm_resource_group" "resource_group" {
			name = "%s"
		}

		resource "ibm_cd_toolchain" "cd_toolchain" {
			name = "%s"
			resource_group_id = data.ibm_resource_group.resource_group.id
		}

		resource "ibm_cd_toolchain_tool_pipeline" "cd_toolchain_tool_pipeline" {
			toolchain_id = ibm_cd_toolchain.cd_toolchain.id
			parameters {
				name = "pipeline-tool-01"
			}
			name = "%s"
		}

		data "ibm_cd_toolchain_tool_pipeline" "cd_toolchain_tool_pipeline" {
			toolchain_id = ibm_cd_toolchain_tool_pipeline.cd_toolchain_tool_pipeline.toolchain_id
			tool_id = ibm_cd_toolchain_tool_pipeline.cd_toolchain_tool_pipeline.tool_id
		}
	`, rgName, tcName, toolName)
}
