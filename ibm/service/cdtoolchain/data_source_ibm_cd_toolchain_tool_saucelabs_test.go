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

func TestAccIBMCdToolchainToolSaucelabsDataSourceBasic(t *testing.T) {
	tcName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	rgName := acc.CdResourceGroupName
	accessKey := acc.CdSaucelabsAccessKey
	username := acc.CdSaucelabsUsername

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCdToolchainToolSaucelabsDataSourceConfigBasic(tcName, rgName, username, accessKey),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_saucelabs.cd_toolchain_tool_saucelabs", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_saucelabs.cd_toolchain_tool_saucelabs", "toolchain_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_saucelabs.cd_toolchain_tool_saucelabs", "tool_id"),
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
	tcName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	rgName := acc.CdResourceGroupName
	toolName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	accessKey := acc.CdSaucelabsAccessKey
	username := acc.CdSaucelabsUsername

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCdToolchainToolSaucelabsDataSourceConfig(tcName, rgName, username, accessKey, toolName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_saucelabs.cd_toolchain_tool_saucelabs", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_saucelabs.cd_toolchain_tool_saucelabs", "toolchain_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_saucelabs.cd_toolchain_tool_saucelabs", "tool_id"),
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

func testAccCheckIBMCdToolchainToolSaucelabsDataSourceConfigBasic(tcName string, rgName string, username string, accessKey string) string {
	return fmt.Sprintf(`
		data "ibm_resource_group" "resource_group" {
			name = "%s"
		}

		resource "ibm_cd_toolchain" "cd_toolchain" {
			name = "%s"
			resource_group_id = data.ibm_resource_group.resource_group.id
		}

		resource "ibm_cd_toolchain_tool_saucelabs" "cd_toolchain_tool_saucelabs" {
			toolchain_id = ibm_cd_toolchain.cd_toolchain.id
			parameters {
				username = "%s"
				access_key = "%s"
			}
		}

		data "ibm_cd_toolchain_tool_saucelabs" "cd_toolchain_tool_saucelabs" {
			toolchain_id = ibm_cd_toolchain_tool_saucelabs.cd_toolchain_tool_saucelabs.toolchain_id
			tool_id = ibm_cd_toolchain_tool_saucelabs.cd_toolchain_tool_saucelabs.tool_id
		}
	`, rgName, tcName, username, accessKey)
}

func testAccCheckIBMCdToolchainToolSaucelabsDataSourceConfig(tcName string, rgName string, username string, accessKey string, toolName string) string {
	return fmt.Sprintf(`
		data "ibm_resource_group" "resource_group" {
			name = "%s"
		}

		resource "ibm_cd_toolchain" "cd_toolchain" {
			name = "%s"
			resource_group_id = data.ibm_resource_group.resource_group.id
		}

		resource "ibm_cd_toolchain_tool_saucelabs" "cd_toolchain_tool_saucelabs" {
			toolchain_id = ibm_cd_toolchain.cd_toolchain.id
			parameters {
				username = "%s"
				access_key = "%s"
			}
			name = "%s"
		}

		data "ibm_cd_toolchain_tool_saucelabs" "cd_toolchain_tool_saucelabs" {
			toolchain_id = ibm_cd_toolchain_tool_saucelabs.cd_toolchain_tool_saucelabs.toolchain_id
			tool_id = ibm_cd_toolchain_tool_saucelabs.cd_toolchain_tool_saucelabs.tool_id
		}
	`, rgName, tcName, username, accessKey, toolName)
}
