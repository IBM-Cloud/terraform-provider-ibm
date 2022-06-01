// Copyright IBM Corp. 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package toolchain_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIBMCdToolchainToolPipelineDataSourceBasic(t *testing.T) {
	getIntegrationByIDResponseToolchainID := fmt.Sprintf("tf_toolchain_id_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCdToolchainToolPipelineDataSourceConfigBasic(getIntegrationByIDResponseToolchainID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_pipeline.cd_toolchain_tool_pipeline", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_pipeline.cd_toolchain_tool_pipeline", "toolchain_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_pipeline.cd_toolchain_tool_pipeline", "integration_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_pipeline.cd_toolchain_tool_pipeline", "get_integration_by_id_response_id"),
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
	getIntegrationByIDResponseToolchainID := fmt.Sprintf("tf_toolchain_id_%d", acctest.RandIntRange(10, 100))
	getIntegrationByIDResponseName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCdToolchainToolPipelineDataSourceConfig(getIntegrationByIDResponseToolchainID, getIntegrationByIDResponseName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_pipeline.cd_toolchain_tool_pipeline", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_pipeline.cd_toolchain_tool_pipeline", "toolchain_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_pipeline.cd_toolchain_tool_pipeline", "integration_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_pipeline.cd_toolchain_tool_pipeline", "get_integration_by_id_response_id"),
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

func testAccCheckIBMCdToolchainToolPipelineDataSourceConfigBasic(getIntegrationByIDResponseToolchainID string) string {
	return fmt.Sprintf(`
		resource "ibm_cd_toolchain_tool_pipeline" "cd_toolchain_tool_pipeline" {
			toolchain_id = "%s"
		}

		data "ibm_cd_toolchain_tool_pipeline" "cd_toolchain_tool_pipeline" {
			toolchain_id = ibm_cd_toolchain_tool_pipeline.cd_toolchain_tool_pipeline.toolchain_id
			integration_id = "integration_id"
		}
	`, getIntegrationByIDResponseToolchainID)
}

func testAccCheckIBMCdToolchainToolPipelineDataSourceConfig(getIntegrationByIDResponseToolchainID string, getIntegrationByIDResponseName string) string {
	return fmt.Sprintf(`
		resource "ibm_cd_toolchain_tool_pipeline" "cd_toolchain_tool_pipeline" {
			toolchain_id = "%s"
			name = "%s"
			parameters {
				name = "name"
				type = "classic"
				ui_pipeline = true
			}
		}

		data "ibm_cd_toolchain_tool_pipeline" "cd_toolchain_tool_pipeline" {
			toolchain_id = ibm_cd_toolchain_tool_pipeline.cd_toolchain_tool_pipeline.toolchain_id
			integration_id = "integration_id"
		}
	`, getIntegrationByIDResponseToolchainID, getIntegrationByIDResponseName)
}
