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

func TestAccIBMCdToolchainToolNexusDataSourceBasic(t *testing.T) {
	getIntegrationByIDResponseToolchainID := fmt.Sprintf("tf_toolchain_id_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCdToolchainToolNexusDataSourceConfigBasic(getIntegrationByIDResponseToolchainID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_nexus.cd_toolchain_tool_nexus", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_nexus.cd_toolchain_tool_nexus", "toolchain_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_nexus.cd_toolchain_tool_nexus", "integration_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_nexus.cd_toolchain_tool_nexus", "get_integration_by_id_response_id"),
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
	getIntegrationByIDResponseToolchainID := fmt.Sprintf("tf_toolchain_id_%d", acctest.RandIntRange(10, 100))
	getIntegrationByIDResponseName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCdToolchainToolNexusDataSourceConfig(getIntegrationByIDResponseToolchainID, getIntegrationByIDResponseName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_nexus.cd_toolchain_tool_nexus", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_nexus.cd_toolchain_tool_nexus", "toolchain_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_nexus.cd_toolchain_tool_nexus", "integration_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_nexus.cd_toolchain_tool_nexus", "get_integration_by_id_response_id"),
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

func testAccCheckIBMCdToolchainToolNexusDataSourceConfigBasic(getIntegrationByIDResponseToolchainID string) string {
	return fmt.Sprintf(`
		resource "ibm_cd_toolchain_tool_nexus" "cd_toolchain_tool_nexus" {
			toolchain_id = "%s"
		}

		data "ibm_cd_toolchain_tool_nexus" "cd_toolchain_tool_nexus" {
			toolchain_id = ibm_cd_toolchain_tool_nexus.cd_toolchain_tool_nexus.toolchain_id
			integration_id = "integration_id"
		}
	`, getIntegrationByIDResponseToolchainID)
}

func testAccCheckIBMCdToolchainToolNexusDataSourceConfig(getIntegrationByIDResponseToolchainID string, getIntegrationByIDResponseName string) string {
	return fmt.Sprintf(`
		resource "ibm_cd_toolchain_tool_nexus" "cd_toolchain_tool_nexus" {
			toolchain_id = "%s"
			name = "%s"
			parameters {
				name = "name"
				dashboard_url = "dashboard_url"
				type = "npm"
				user_id = "user_id"
				token = "token"
				release_url = "release_url"
				mirror_url = "mirror_url"
				snapshot_url = "snapshot_url"
			}
		}

		data "ibm_cd_toolchain_tool_nexus" "cd_toolchain_tool_nexus" {
			toolchain_id = ibm_cd_toolchain_tool_nexus.cd_toolchain_tool_nexus.toolchain_id
			integration_id = "integration_id"
		}
	`, getIntegrationByIDResponseToolchainID, getIntegrationByIDResponseName)
}
