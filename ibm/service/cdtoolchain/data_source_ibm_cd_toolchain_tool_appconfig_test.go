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

func TestAccIBMCdToolchainToolAppconfigDataSourceBasic(t *testing.T) {
	getToolByIDResponseToolchainID := fmt.Sprintf("tf_toolchain_id_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCdToolchainToolAppconfigDataSourceConfigBasic(getToolByIDResponseToolchainID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_appconfig.cd_toolchain_tool_appconfig", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_appconfig.cd_toolchain_tool_appconfig", "toolchain_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_appconfig.cd_toolchain_tool_appconfig", "tool_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_appconfig.cd_toolchain_tool_appconfig", "get_tool_by_id_response_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_appconfig.cd_toolchain_tool_appconfig", "resource_group_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_appconfig.cd_toolchain_tool_appconfig", "crn"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_appconfig.cd_toolchain_tool_appconfig", "toolchain_crn"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_appconfig.cd_toolchain_tool_appconfig", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_appconfig.cd_toolchain_tool_appconfig", "referent.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_appconfig.cd_toolchain_tool_appconfig", "updated_at"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_appconfig.cd_toolchain_tool_appconfig", "parameters.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_appconfig.cd_toolchain_tool_appconfig", "state"),
				),
			},
		},
	})
}

func TestAccIBMCdToolchainToolAppconfigDataSourceAllArgs(t *testing.T) {
	getToolByIDResponseToolchainID := fmt.Sprintf("tf_toolchain_id_%d", acctest.RandIntRange(10, 100))
	getToolByIDResponseName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCdToolchainToolAppconfigDataSourceConfig(getToolByIDResponseToolchainID, getToolByIDResponseName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_appconfig.cd_toolchain_tool_appconfig", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_appconfig.cd_toolchain_tool_appconfig", "toolchain_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_appconfig.cd_toolchain_tool_appconfig", "tool_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_appconfig.cd_toolchain_tool_appconfig", "get_tool_by_id_response_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_appconfig.cd_toolchain_tool_appconfig", "resource_group_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_appconfig.cd_toolchain_tool_appconfig", "crn"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_appconfig.cd_toolchain_tool_appconfig", "toolchain_crn"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_appconfig.cd_toolchain_tool_appconfig", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_appconfig.cd_toolchain_tool_appconfig", "referent.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_appconfig.cd_toolchain_tool_appconfig", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_appconfig.cd_toolchain_tool_appconfig", "updated_at"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_appconfig.cd_toolchain_tool_appconfig", "parameters.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_appconfig.cd_toolchain_tool_appconfig", "state"),
				),
			},
		},
	})
}

func testAccCheckIBMCdToolchainToolAppconfigDataSourceConfigBasic(getToolByIDResponseToolchainID string) string {
	return fmt.Sprintf(`
		resource "ibm_cd_toolchain_tool_appconfig" "cd_toolchain_tool_appconfig" {
			toolchain_id = "%s"
		}

		data "ibm_cd_toolchain_tool_appconfig" "cd_toolchain_tool_appconfig" {
			toolchain_id = ibm_cd_toolchain_tool_appconfig.cd_toolchain_tool_appconfig.toolchain_id
			tool_id = "tool_id"
		}
	`, getToolByIDResponseToolchainID)
}

func testAccCheckIBMCdToolchainToolAppconfigDataSourceConfig(getToolByIDResponseToolchainID string, getToolByIDResponseName string) string {
	return fmt.Sprintf(`
		resource "ibm_cd_toolchain_tool_appconfig" "cd_toolchain_tool_appconfig" {
			toolchain_id = "%s"
			name = "%s"
			parameters {
				name = "name"
				region = "region"
				resource-group = "resource-group"
				instance-name = "instance-name"
				environment-name = "environment-name"
				collection-name = "collection-name"
				integration-status = "integration-status"
			}
		}

		data "ibm_cd_toolchain_tool_appconfig" "cd_toolchain_tool_appconfig" {
			toolchain_id = ibm_cd_toolchain_tool_appconfig.cd_toolchain_tool_appconfig.toolchain_id
			tool_id = "tool_id"
		}
	`, getToolByIDResponseToolchainID, getToolByIDResponseName)
}
