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

func TestAccIBMCdToolchainToolRationalteamconcertDataSourceBasic(t *testing.T) {
	getToolByIDResponseToolchainID := fmt.Sprintf("tf_toolchain_id_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCdToolchainToolRationalteamconcertDataSourceConfigBasic(getToolByIDResponseToolchainID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_rationalteamconcert.cd_toolchain_tool_rationalteamconcert", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_rationalteamconcert.cd_toolchain_tool_rationalteamconcert", "toolchain_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_rationalteamconcert.cd_toolchain_tool_rationalteamconcert", "tool_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_rationalteamconcert.cd_toolchain_tool_rationalteamconcert", "get_tool_by_id_response_id"),
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
	getToolByIDResponseToolchainID := fmt.Sprintf("tf_toolchain_id_%d", acctest.RandIntRange(10, 100))
	getToolByIDResponseName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCdToolchainToolRationalteamconcertDataSourceConfig(getToolByIDResponseToolchainID, getToolByIDResponseName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_rationalteamconcert.cd_toolchain_tool_rationalteamconcert", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_rationalteamconcert.cd_toolchain_tool_rationalteamconcert", "toolchain_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_rationalteamconcert.cd_toolchain_tool_rationalteamconcert", "tool_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_rationalteamconcert.cd_toolchain_tool_rationalteamconcert", "get_tool_by_id_response_id"),
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

func testAccCheckIBMCdToolchainToolRationalteamconcertDataSourceConfigBasic(getToolByIDResponseToolchainID string) string {
	return fmt.Sprintf(`
		resource "ibm_cd_toolchain_tool_rationalteamconcert" "cd_toolchain_tool_rationalteamconcert" {
			toolchain_id = "%s"
		}

		data "ibm_cd_toolchain_tool_rationalteamconcert" "cd_toolchain_tool_rationalteamconcert" {
			toolchain_id = ibm_cd_toolchain_tool_rationalteamconcert.cd_toolchain_tool_rationalteamconcert.toolchain_id
			tool_id = "tool_id"
		}
	`, getToolByIDResponseToolchainID)
}

func testAccCheckIBMCdToolchainToolRationalteamconcertDataSourceConfig(getToolByIDResponseToolchainID string, getToolByIDResponseName string) string {
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
			tool_id = "tool_id"
		}
	`, getToolByIDResponseToolchainID, getToolByIDResponseName)
}
