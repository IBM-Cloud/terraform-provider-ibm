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

func TestAccIBMCdToolchainToolKeyprotectDataSourceBasic(t *testing.T) {
	tcName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	rgName := acc.CdResourceGroupName
	kpName := acc.CdKeyProtectInstanceName

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCdToolchainToolKeyprotectDataSourceConfigBasic(tcName, rgName, kpName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_keyprotect.cd_toolchain_tool_keyprotect", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_keyprotect.cd_toolchain_tool_keyprotect", "toolchain_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_keyprotect.cd_toolchain_tool_keyprotect", "tool_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_keyprotect.cd_toolchain_tool_keyprotect", "resource_group_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_keyprotect.cd_toolchain_tool_keyprotect", "crn"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_keyprotect.cd_toolchain_tool_keyprotect", "toolchain_crn"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_keyprotect.cd_toolchain_tool_keyprotect", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_keyprotect.cd_toolchain_tool_keyprotect", "referent.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_keyprotect.cd_toolchain_tool_keyprotect", "updated_at"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_keyprotect.cd_toolchain_tool_keyprotect", "parameters.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_keyprotect.cd_toolchain_tool_keyprotect", "state"),
				),
			},
		},
	})
}

func TestAccIBMCdToolchainToolKeyprotectDataSourceAllArgs(t *testing.T) {
	tcName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	rgName := acc.CdResourceGroupName
	kpName := acc.CdKeyProtectInstanceName
	toolName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCdToolchainToolKeyprotectDataSourceConfig(tcName, rgName, kpName, toolName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_keyprotect.cd_toolchain_tool_keyprotect", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_keyprotect.cd_toolchain_tool_keyprotect", "toolchain_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_keyprotect.cd_toolchain_tool_keyprotect", "tool_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_keyprotect.cd_toolchain_tool_keyprotect", "resource_group_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_keyprotect.cd_toolchain_tool_keyprotect", "crn"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_keyprotect.cd_toolchain_tool_keyprotect", "toolchain_crn"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_keyprotect.cd_toolchain_tool_keyprotect", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_keyprotect.cd_toolchain_tool_keyprotect", "referent.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_keyprotect.cd_toolchain_tool_keyprotect", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_keyprotect.cd_toolchain_tool_keyprotect", "updated_at"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_keyprotect.cd_toolchain_tool_keyprotect", "parameters.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_keyprotect.cd_toolchain_tool_keyprotect", "state"),
				),
			},
		},
	})
}

func testAccCheckIBMCdToolchainToolKeyprotectDataSourceConfigBasic(tcName string, rgName string, kpName string) string {
	return fmt.Sprintf(`
		data "ibm_resource_group" "resource_group" {
			name = "%s"
		}

		resource "ibm_cd_toolchain" "cd_toolchain" {
			name = "%s"
			resource_group_id = data.ibm_resource_group.resource_group.id
		}

		resource "ibm_cd_toolchain_tool_keyprotect" "cd_toolchain_tool_keyprotect" {
			toolchain_id = ibm_cd_toolchain.cd_toolchain.id
			parameters {
				name = "kp_tool_01"
				instance_name = "%s"
				location = "us-south"
				resource_group_name = "%s"
			}
		}

		data "ibm_cd_toolchain_tool_keyprotect" "cd_toolchain_tool_keyprotect" {
			toolchain_id = ibm_cd_toolchain_tool_keyprotect.cd_toolchain_tool_keyprotect.toolchain_id
			tool_id = ibm_cd_toolchain_tool_keyprotect.cd_toolchain_tool_keyprotect.tool_id
		}
	`, rgName, tcName, kpName, rgName)
}

func testAccCheckIBMCdToolchainToolKeyprotectDataSourceConfig(tcName string, rgName string, kpName string, toolName string) string {
	return fmt.Sprintf(`
		data "ibm_resource_group" "resource_group" {
			name = "%s"
		}

		resource "ibm_cd_toolchain" "cd_toolchain" {
			name = "%s"
			resource_group_id = data.ibm_resource_group.resource_group.id
		}

		resource "ibm_cd_toolchain_tool_keyprotect" "cd_toolchain_tool_keyprotect" {
			toolchain_id = ibm_cd_toolchain.cd_toolchain.id
			parameters {
				name = "kp_tool_01"
				instance_name = "%s"
				location = "us-south"
				resource_group_name = "%s"
			}
			name = "%s"
		}

		data "ibm_cd_toolchain_tool_keyprotect" "cd_toolchain_tool_keyprotect" {
			toolchain_id = ibm_cd_toolchain_tool_keyprotect.cd_toolchain_tool_keyprotect.toolchain_id
			tool_id = ibm_cd_toolchain_tool_keyprotect.cd_toolchain_tool_keyprotect.tool_id
		}
	`, rgName, tcName, kpName, rgName, toolName)
}
