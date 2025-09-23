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

func TestAccIBMCdToolchainToolDevopsinsightsDataSourceBasic(t *testing.T) {
	tcName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	rgName := acc.CdResourceGroupName

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCdToolchainToolDevopsinsightsDataSourceConfigBasic(tcName, rgName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_devopsinsights.cd_toolchain_tool_devopsinsights", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_devopsinsights.cd_toolchain_tool_devopsinsights", "toolchain_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_devopsinsights.cd_toolchain_tool_devopsinsights", "tool_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_devopsinsights.cd_toolchain_tool_devopsinsights", "resource_group_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_devopsinsights.cd_toolchain_tool_devopsinsights", "crn"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_devopsinsights.cd_toolchain_tool_devopsinsights", "toolchain_crn"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_devopsinsights.cd_toolchain_tool_devopsinsights", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_devopsinsights.cd_toolchain_tool_devopsinsights", "referent.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_devopsinsights.cd_toolchain_tool_devopsinsights", "updated_at"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_devopsinsights.cd_toolchain_tool_devopsinsights", "state"),
				),
			},
		},
	})
}

func TestAccIBMCdToolchainToolDevopsinsightsDataSourceAllArgs(t *testing.T) {
	tcName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	rgName := acc.CdResourceGroupName
	toolName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCdToolchainToolDevopsinsightsDataSourceConfig(tcName, rgName, toolName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_devopsinsights.cd_toolchain_tool_devopsinsights", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_devopsinsights.cd_toolchain_tool_devopsinsights", "toolchain_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_devopsinsights.cd_toolchain_tool_devopsinsights", "tool_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_devopsinsights.cd_toolchain_tool_devopsinsights", "resource_group_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_devopsinsights.cd_toolchain_tool_devopsinsights", "crn"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_devopsinsights.cd_toolchain_tool_devopsinsights", "toolchain_crn"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_devopsinsights.cd_toolchain_tool_devopsinsights", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_devopsinsights.cd_toolchain_tool_devopsinsights", "referent.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_devopsinsights.cd_toolchain_tool_devopsinsights", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_devopsinsights.cd_toolchain_tool_devopsinsights", "updated_at"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_devopsinsights.cd_toolchain_tool_devopsinsights", "state"),
				),
			},
		},
	})
}

func testAccCheckIBMCdToolchainToolDevopsinsightsDataSourceConfigBasic(tcName string, rgName string) string {
	return fmt.Sprintf(`
		data "ibm_resource_group" "resource_group" {
			name = "%s"
		}

		resource "ibm_cd_toolchain" "cd_toolchain" {
			name = "%s"
			resource_group_id = data.ibm_resource_group.resource_group.id
		}

		resource "ibm_cd_toolchain_tool_devopsinsights" "cd_toolchain_tool_devopsinsights" {
			toolchain_id = ibm_cd_toolchain.cd_toolchain.id
		}

		data "ibm_cd_toolchain_tool_devopsinsights" "cd_toolchain_tool_devopsinsights" {
			toolchain_id = ibm_cd_toolchain_tool_devopsinsights.cd_toolchain_tool_devopsinsights.toolchain_id
			tool_id = ibm_cd_toolchain_tool_devopsinsights.cd_toolchain_tool_devopsinsights.tool_id
		}
	`, rgName, tcName)
}

func testAccCheckIBMCdToolchainToolDevopsinsightsDataSourceConfig(tcName string, rgName string, toolName string) string {
	return fmt.Sprintf(`
		data "ibm_resource_group" "resource_group" {
			name = "%s"
		}

		resource "ibm_cd_toolchain" "cd_toolchain" {
			name = "%s"
			resource_group_id = data.ibm_resource_group.resource_group.id
		}

		resource "ibm_cd_toolchain_tool_devopsinsights" "cd_toolchain_tool_devopsinsights" {
			toolchain_id = ibm_cd_toolchain.cd_toolchain.id
			name = "%s"
		}

		data "ibm_cd_toolchain_tool_devopsinsights" "cd_toolchain_tool_devopsinsights" {
			toolchain_id = ibm_cd_toolchain_tool_devopsinsights.cd_toolchain_tool_devopsinsights.toolchain_id
			tool_id = ibm_cd_toolchain_tool_devopsinsights.cd_toolchain_tool_devopsinsights.tool_id
		}
	`, rgName, tcName, toolName)
}
