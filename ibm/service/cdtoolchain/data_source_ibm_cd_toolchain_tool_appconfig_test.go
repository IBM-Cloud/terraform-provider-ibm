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

func TestAccIBMCdToolchainToolAppconfigDataSourceBasic(t *testing.T) {
	tcName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	rgName := acc.CdResourceGroupName
	acName := acc.CdAppConfigInstanceName

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCdToolchainToolAppconfigDataSourceConfigBasic(tcName, rgName, acName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_appconfig.cd_toolchain_tool_appconfig", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_appconfig.cd_toolchain_tool_appconfig", "toolchain_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_appconfig.cd_toolchain_tool_appconfig", "tool_id"),
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
	tcName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	rgName := acc.CdResourceGroupName
	acName := acc.CdAppConfigInstanceName
	toolName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCdToolchainToolAppconfigDataSourceConfig(tcName, rgName, acName, toolName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_appconfig.cd_toolchain_tool_appconfig", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_appconfig.cd_toolchain_tool_appconfig", "toolchain_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_appconfig.cd_toolchain_tool_appconfig", "tool_id"),
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

func testAccCheckIBMCdToolchainToolAppconfigDataSourceConfigBasic(tcName string, rgName string, acName string) string {
	return fmt.Sprintf(`
		data "ibm_resource_group" "resource_group" {
			name = "%s"
		}

		resource "ibm_cd_toolchain" "cd_toolchain" {
			name = "%s"
			resource_group_id = data.ibm_resource_group.resource_group.id
		}

		data "ibm_resource_instance" "ac_resource_instance" {
			name              = "%s"
			location		  = "us-south"
			resource_group_id = data.ibm_resource_group.resource_group.id
			service           = "apprapp"
		}

		resource "ibm_cd_toolchain_tool_appconfig" "cd_toolchain_tool_appconfig" {
			toolchain_id = ibm_cd_toolchain.cd_toolchain.id
			parameters {
				name = "appconfig_tool_01"
				location = "us-south"
				resource_group_name = "%s"
				instance_id = data.ibm_resource_instance.ac_resource_instance.guid
				environment_id = "environment_01"
				collection_id = "collection_01"
			}
		}

		data "ibm_cd_toolchain_tool_appconfig" "cd_toolchain_tool_appconfig" {
			toolchain_id = ibm_cd_toolchain_tool_appconfig.cd_toolchain_tool_appconfig.toolchain_id
			tool_id = ibm_cd_toolchain_tool_appconfig.cd_toolchain_tool_appconfig.tool_id
		}
	`, rgName, tcName, acName, rgName)
}

func testAccCheckIBMCdToolchainToolAppconfigDataSourceConfig(tcName string, rgName string, acName string, toolName string) string {
	return fmt.Sprintf(`
		data "ibm_resource_group" "resource_group" {
			name = "%s"
		}

		resource "ibm_cd_toolchain" "cd_toolchain" {
			name = "%s"
			resource_group_id = data.ibm_resource_group.resource_group.id
		}

		data "ibm_resource_instance" "ac_resource_instance" {
			name              = "%s"
			location		  = "us-south"
			resource_group_id = data.ibm_resource_group.resource_group.id
			service           = "apprapp"
		}

		resource "ibm_cd_toolchain_tool_appconfig" "cd_toolchain_tool_appconfig" {
			toolchain_id = ibm_cd_toolchain.cd_toolchain.id
			parameters {
				name = "appconfig_tool_01"
				location = "us-south"
				resource_group_name = "%s"
				instance_id = data.ibm_resource_instance.ac_resource_instance.guid
				environment_id = "environment_01"
				collection_id = "collection_01"
			}
			name = "%s"
		}

		data "ibm_cd_toolchain_tool_appconfig" "cd_toolchain_tool_appconfig" {
			toolchain_id = ibm_cd_toolchain_tool_appconfig.cd_toolchain_tool_appconfig.toolchain_id
			tool_id = ibm_cd_toolchain_tool_appconfig.cd_toolchain_tool_appconfig.tool_id
		}
	`, rgName, tcName, acName, rgName, toolName)
}
