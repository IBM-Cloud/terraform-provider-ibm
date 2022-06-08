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

func TestAccIBMCdToolchainToolSaucelabsDataSourceBasic(t *testing.T) {
	getToolByIDResponseToolchainID := fmt.Sprintf("tf_toolchain_id_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCdToolchainToolSaucelabsDataSourceConfigBasic(getToolByIDResponseToolchainID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_saucelabs.cd_toolchain_tool_saucelabs", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_saucelabs.cd_toolchain_tool_saucelabs", "toolchain_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_saucelabs.cd_toolchain_tool_saucelabs", "tool_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_saucelabs.cd_toolchain_tool_saucelabs", "get_tool_by_id_response_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_saucelabs.cd_toolchain_tool_saucelabs", "resource_group_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_saucelabs.cd_toolchain_tool_saucelabs", "crn"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_saucelabs.cd_toolchain_tool_saucelabs", "toolchain_crn"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_saucelabs.cd_toolchain_tool_saucelabs", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_saucelabs.cd_toolchain_tool_saucelabs", "referent.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_saucelabs.cd_toolchain_tool_saucelabs", "updated_at"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_saucelabs.cd_toolchain_tool_saucelabs", "parameters.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_saucelabs.cd_toolchain_tool_saucelabs", "state"),
				),
			},
		},
	})
}

func TestAccIBMCdToolchainToolSaucelabsDataSourceAllArgs(t *testing.T) {
	getToolByIDResponseToolchainID := fmt.Sprintf("tf_toolchain_id_%d", acctest.RandIntRange(10, 100))
	getToolByIDResponseName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCdToolchainToolSaucelabsDataSourceConfig(getToolByIDResponseToolchainID, getToolByIDResponseName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_saucelabs.cd_toolchain_tool_saucelabs", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_saucelabs.cd_toolchain_tool_saucelabs", "toolchain_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_saucelabs.cd_toolchain_tool_saucelabs", "tool_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_saucelabs.cd_toolchain_tool_saucelabs", "get_tool_by_id_response_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_saucelabs.cd_toolchain_tool_saucelabs", "resource_group_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_saucelabs.cd_toolchain_tool_saucelabs", "crn"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_saucelabs.cd_toolchain_tool_saucelabs", "toolchain_crn"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_saucelabs.cd_toolchain_tool_saucelabs", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_saucelabs.cd_toolchain_tool_saucelabs", "referent.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_saucelabs.cd_toolchain_tool_saucelabs", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_saucelabs.cd_toolchain_tool_saucelabs", "updated_at"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_saucelabs.cd_toolchain_tool_saucelabs", "parameters.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_saucelabs.cd_toolchain_tool_saucelabs", "state"),
				),
			},
		},
	})
}

func testAccCheckIBMCdToolchainToolSaucelabsDataSourceConfigBasic(getToolByIDResponseToolchainID string) string {
	return fmt.Sprintf(`
		resource "ibm_cd_toolchain_tool_saucelabs" "cd_toolchain_tool_saucelabs" {
			toolchain_id = "%s"
		}

		data "ibm_cd_toolchain_tool_saucelabs" "cd_toolchain_tool_saucelabs" {
			toolchain_id = ibm_cd_toolchain_tool_saucelabs.cd_toolchain_tool_saucelabs.toolchain_id
			tool_id = "tool_id"
		}
	`, getToolByIDResponseToolchainID)
}

func testAccCheckIBMCdToolchainToolSaucelabsDataSourceConfig(getToolByIDResponseToolchainID string, getToolByIDResponseName string) string {
	return fmt.Sprintf(`
		resource "ibm_cd_toolchain_tool_saucelabs" "cd_toolchain_tool_saucelabs" {
			toolchain_id = "%s"
			name = "%s"
			parameters {
				username = "username"
				key = "key"
			}
		}

		data "ibm_cd_toolchain_tool_saucelabs" "cd_toolchain_tool_saucelabs" {
			toolchain_id = ibm_cd_toolchain_tool_saucelabs.cd_toolchain_tool_saucelabs.toolchain_id
			tool_id = "tool_id"
		}
	`, getToolByIDResponseToolchainID, getToolByIDResponseName)
}
