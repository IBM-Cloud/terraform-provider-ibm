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

func TestAccIBMCdToolchainToolHashicorpvaultDataSourceBasic(t *testing.T) {
	tcName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	rgName := acc.CdResourceGroupName

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCdToolchainToolHashicorpvaultDataSourceConfigBasic(tcName, rgName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_hashicorpvault.cd_toolchain_tool_hashicorpvault", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_hashicorpvault.cd_toolchain_tool_hashicorpvault", "toolchain_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_hashicorpvault.cd_toolchain_tool_hashicorpvault", "tool_id"),
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
	tcName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	rgName := acc.CdResourceGroupName
	toolName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCdToolchainToolHashicorpvaultDataSourceConfig(tcName, rgName, toolName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_hashicorpvault.cd_toolchain_tool_hashicorpvault", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_hashicorpvault.cd_toolchain_tool_hashicorpvault", "toolchain_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_hashicorpvault.cd_toolchain_tool_hashicorpvault", "tool_id"),
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

func testAccCheckIBMCdToolchainToolHashicorpvaultDataSourceConfigBasic(tcName string, rgName string) string {
	return fmt.Sprintf(`
		data "ibm_resource_group" "resource_group" {
			name = "%s"
		}

		resource "ibm_cd_toolchain" "cd_toolchain" {
			name = "%s"
			resource_group_id = data.ibm_resource_group.resource_group.id
		}

		resource "ibm_cd_toolchain_tool_hashicorpvault" "cd_toolchain_tool_hashicorpvault" {
			toolchain_id = ibm_cd_toolchain.cd_toolchain.id
			parameters {
				name = "hcv_tool_01"
				server_url = "https://hcv.mycompany.example.com:8200"
				authentication_method = "approle"
				token = "token"
				role_id = "<role_id>"
				secret_id = "<secret_id>"
				dashboard_url = "https://hcv.mycompany.example.com:8200/ui"
				path = "generic/project/test_project"
				secret_filter = "secret_filter"
				default_secret = "default_secret"
				username = "username"
				password = "password"
			}
		}

		data "ibm_cd_toolchain_tool_hashicorpvault" "cd_toolchain_tool_hashicorpvault" {
			toolchain_id = ibm_cd_toolchain_tool_hashicorpvault.cd_toolchain_tool_hashicorpvault.toolchain_id
			tool_id = ibm_cd_toolchain_tool_hashicorpvault.cd_toolchain_tool_hashicorpvault.tool_id
		}
	`, rgName, tcName)
}

func testAccCheckIBMCdToolchainToolHashicorpvaultDataSourceConfig(tcName string, rgName string, toolName string) string {
	return fmt.Sprintf(`
		data "ibm_resource_group" "resource_group" {
			name = "%s"
		}

		resource "ibm_cd_toolchain" "cd_toolchain" {
			name = "%s"
			resource_group_id = data.ibm_resource_group.resource_group.id
		}

		resource "ibm_cd_toolchain_tool_hashicorpvault" "cd_toolchain_tool_hashicorpvault" {
			toolchain_id = ibm_cd_toolchain.cd_toolchain.id
			parameters {
				name = "hcv_tool_01"
				server_url = "https://hcv.mycompany.example.com:8200"
				authentication_method = "approle"
				token = "token"
				role_id = "<role_id>"
				secret_id = "<secret_id>"
				dashboard_url = "https://hcv.mycompany.example.com:8200/ui"
				path = "generic/project/test_project"
				secret_filter = "secret_filter"
				default_secret = "default_secret"
				username = "username"
				password = "password"
			}
			name = "%s"
		}

		data "ibm_cd_toolchain_tool_hashicorpvault" "cd_toolchain_tool_hashicorpvault" {
			toolchain_id = ibm_cd_toolchain_tool_hashicorpvault.cd_toolchain_tool_hashicorpvault.toolchain_id
			tool_id = ibm_cd_toolchain_tool_hashicorpvault.cd_toolchain_tool_hashicorpvault.tool_id
		}
	`, rgName, tcName, toolName)
}
