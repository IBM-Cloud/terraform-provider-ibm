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

func TestAccIBMCdToolchainToolHashicorpvaultDataSourceBasic(t *testing.T) {
	getToolByIDResponseToolchainID := fmt.Sprintf("tf_toolchain_id_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCdToolchainToolHashicorpvaultDataSourceConfigBasic(getToolByIDResponseToolchainID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_hashicorpvault.cd_toolchain_tool_hashicorpvault", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_hashicorpvault.cd_toolchain_tool_hashicorpvault", "toolchain_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_hashicorpvault.cd_toolchain_tool_hashicorpvault", "tool_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_hashicorpvault.cd_toolchain_tool_hashicorpvault", "get_tool_by_id_response_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_hashicorpvault.cd_toolchain_tool_hashicorpvault", "resource_group_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_hashicorpvault.cd_toolchain_tool_hashicorpvault", "crn"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_hashicorpvault.cd_toolchain_tool_hashicorpvault", "toolchain_crn"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_hashicorpvault.cd_toolchain_tool_hashicorpvault", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_hashicorpvault.cd_toolchain_tool_hashicorpvault", "referent.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_hashicorpvault.cd_toolchain_tool_hashicorpvault", "updated_at"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_hashicorpvault.cd_toolchain_tool_hashicorpvault", "parameters.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_hashicorpvault.cd_toolchain_tool_hashicorpvault", "state"),
				),
			},
		},
	})
}

func TestAccIBMCdToolchainToolHashicorpvaultDataSourceAllArgs(t *testing.T) {
	getToolByIDResponseToolchainID := fmt.Sprintf("tf_toolchain_id_%d", acctest.RandIntRange(10, 100))
	getToolByIDResponseName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCdToolchainToolHashicorpvaultDataSourceConfig(getToolByIDResponseToolchainID, getToolByIDResponseName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_hashicorpvault.cd_toolchain_tool_hashicorpvault", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_hashicorpvault.cd_toolchain_tool_hashicorpvault", "toolchain_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_hashicorpvault.cd_toolchain_tool_hashicorpvault", "tool_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_hashicorpvault.cd_toolchain_tool_hashicorpvault", "get_tool_by_id_response_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_hashicorpvault.cd_toolchain_tool_hashicorpvault", "resource_group_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_hashicorpvault.cd_toolchain_tool_hashicorpvault", "crn"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_hashicorpvault.cd_toolchain_tool_hashicorpvault", "toolchain_crn"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_hashicorpvault.cd_toolchain_tool_hashicorpvault", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_hashicorpvault.cd_toolchain_tool_hashicorpvault", "referent.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_hashicorpvault.cd_toolchain_tool_hashicorpvault", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_hashicorpvault.cd_toolchain_tool_hashicorpvault", "updated_at"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_hashicorpvault.cd_toolchain_tool_hashicorpvault", "parameters.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_hashicorpvault.cd_toolchain_tool_hashicorpvault", "state"),
				),
			},
		},
	})
}

func testAccCheckIBMCdToolchainToolHashicorpvaultDataSourceConfigBasic(getToolByIDResponseToolchainID string) string {
	return fmt.Sprintf(`
		resource "ibm_cd_toolchain_tool_hashicorpvault" "cd_toolchain_tool_hashicorpvault" {
			toolchain_id = "%s"
		}

		data "ibm_cd_toolchain_tool_hashicorpvault" "cd_toolchain_tool_hashicorpvault" {
			toolchain_id = ibm_cd_toolchain_tool_hashicorpvault.cd_toolchain_tool_hashicorpvault.toolchain_id
			tool_id = "tool_id"
		}
	`, getToolByIDResponseToolchainID)
}

func testAccCheckIBMCdToolchainToolHashicorpvaultDataSourceConfig(getToolByIDResponseToolchainID string, getToolByIDResponseName string) string {
	return fmt.Sprintf(`
		resource "ibm_cd_toolchain_tool_hashicorpvault" "cd_toolchain_tool_hashicorpvault" {
			toolchain_id = "%s"
			name = "%s"
			parameters {
				name = "name"
				server_url = "server_url"
				authentication_method = "token"
				token = "token"
				role_id = "role_id"
				secret_id = "secret_id"
				dashboard_url = "dashboard_url"
				path = "path"
				secret_filter = "secret_filter"
				default_secret = "default_secret"
				username = "username"
				password = "password"
			}
		}

		data "ibm_cd_toolchain_tool_hashicorpvault" "cd_toolchain_tool_hashicorpvault" {
			toolchain_id = ibm_cd_toolchain_tool_hashicorpvault.cd_toolchain_tool_hashicorpvault.toolchain_id
			tool_id = "tool_id"
		}
	`, getToolByIDResponseToolchainID, getToolByIDResponseName)
}
