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

func TestAccIBMCdToolchainToolRationalteamconcertDataSourceBasic(t *testing.T) {
	getIntegrationByIDResponseToolchainID := fmt.Sprintf("tf_toolchain_id_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCdToolchainToolRationalteamconcertDataSourceConfigBasic(getIntegrationByIDResponseToolchainID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_rationalteamconcert.cd_toolchain_tool_rationalteamconcert", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_rationalteamconcert.cd_toolchain_tool_rationalteamconcert", "toolchain_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_rationalteamconcert.cd_toolchain_tool_rationalteamconcert", "integration_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_rationalteamconcert.cd_toolchain_tool_rationalteamconcert", "get_integration_by_id_response_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_rationalteamconcert.cd_toolchain_tool_rationalteamconcert", "resource_group_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_rationalteamconcert.cd_toolchain_tool_rationalteamconcert", "crn"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_rationalteamconcert.cd_toolchain_tool_rationalteamconcert", "toolchain_crn"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_rationalteamconcert.cd_toolchain_tool_rationalteamconcert", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_rationalteamconcert.cd_toolchain_tool_rationalteamconcert", "referent.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_rationalteamconcert.cd_toolchain_tool_rationalteamconcert", "updated_at"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_rationalteamconcert.cd_toolchain_tool_rationalteamconcert", "parameters.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_rationalteamconcert.cd_toolchain_tool_rationalteamconcert", "state"),
				),
			},
		},
	})
}

func TestAccIBMCdToolchainToolRationalteamconcertDataSourceAllArgs(t *testing.T) {
	getIntegrationByIDResponseToolchainID := fmt.Sprintf("tf_toolchain_id_%d", acctest.RandIntRange(10, 100))
	getIntegrationByIDResponseName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCdToolchainToolRationalteamconcertDataSourceConfig(getIntegrationByIDResponseToolchainID, getIntegrationByIDResponseName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_rationalteamconcert.cd_toolchain_tool_rationalteamconcert", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_rationalteamconcert.cd_toolchain_tool_rationalteamconcert", "toolchain_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_rationalteamconcert.cd_toolchain_tool_rationalteamconcert", "integration_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_rationalteamconcert.cd_toolchain_tool_rationalteamconcert", "get_integration_by_id_response_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_rationalteamconcert.cd_toolchain_tool_rationalteamconcert", "resource_group_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_rationalteamconcert.cd_toolchain_tool_rationalteamconcert", "crn"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_rationalteamconcert.cd_toolchain_tool_rationalteamconcert", "toolchain_crn"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_rationalteamconcert.cd_toolchain_tool_rationalteamconcert", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_rationalteamconcert.cd_toolchain_tool_rationalteamconcert", "referent.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_rationalteamconcert.cd_toolchain_tool_rationalteamconcert", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_rationalteamconcert.cd_toolchain_tool_rationalteamconcert", "updated_at"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_rationalteamconcert.cd_toolchain_tool_rationalteamconcert", "parameters.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_rationalteamconcert.cd_toolchain_tool_rationalteamconcert", "state"),
				),
			},
		},
	})
}

func testAccCheckIBMCdToolchainToolRationalteamconcertDataSourceConfigBasic(getIntegrationByIDResponseToolchainID string) string {
	return fmt.Sprintf(`
		resource "ibm_cd_toolchain_tool_rationalteamconcert" "cd_toolchain_tool_rationalteamconcert" {
			toolchain_id = "%s"
		}

		data "ibm_cd_toolchain_tool_rationalteamconcert" "cd_toolchain_tool_rationalteamconcert" {
			toolchain_id = ibm_cd_toolchain_tool_rationalteamconcert.cd_toolchain_tool_rationalteamconcert.toolchain_id
			integration_id = "integration_id"
		}
	`, getIntegrationByIDResponseToolchainID)
}

func testAccCheckIBMCdToolchainToolRationalteamconcertDataSourceConfig(getIntegrationByIDResponseToolchainID string, getIntegrationByIDResponseName string) string {
	return fmt.Sprintf(`
		resource "ibm_cd_toolchain_tool_rationalteamconcert" "cd_toolchain_tool_rationalteamconcert" {
			toolchain_id = "%s"
			name = "%s"
			parameters {
				server_url = "server_url"
				user_id = "user_id"
				password = "password"
				type = "new"
				project_area = "project_area"
				process_template = "process_template"
				enable_traceability = true
			}
		}

		data "ibm_cd_toolchain_tool_rationalteamconcert" "cd_toolchain_tool_rationalteamconcert" {
			toolchain_id = ibm_cd_toolchain_tool_rationalteamconcert.cd_toolchain_tool_rationalteamconcert.toolchain_id
			integration_id = "integration_id"
		}
	`, getIntegrationByIDResponseToolchainID, getIntegrationByIDResponseName)
}
