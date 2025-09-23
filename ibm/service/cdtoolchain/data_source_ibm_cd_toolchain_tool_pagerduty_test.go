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

func TestAccIBMCdToolchainToolPagerdutyDataSourceBasic(t *testing.T) {
	tcName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	rgName := acc.CdResourceGroupName

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCdToolchainToolPagerdutyDataSourceConfigBasic(tcName, rgName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_pagerduty.cd_toolchain_tool_pagerduty", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_pagerduty.cd_toolchain_tool_pagerduty", "toolchain_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_pagerduty.cd_toolchain_tool_pagerduty", "tool_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_pagerduty.cd_toolchain_tool_pagerduty", "resource_group_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_pagerduty.cd_toolchain_tool_pagerduty", "crn"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_pagerduty.cd_toolchain_tool_pagerduty", "toolchain_crn"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_pagerduty.cd_toolchain_tool_pagerduty", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_pagerduty.cd_toolchain_tool_pagerduty", "referent.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_pagerduty.cd_toolchain_tool_pagerduty", "updated_at"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_pagerduty.cd_toolchain_tool_pagerduty", "parameters.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_pagerduty.cd_toolchain_tool_pagerduty", "state"),
				),
			},
		},
	})
}

func TestAccIBMCdToolchainToolPagerdutyDataSourceAllArgs(t *testing.T) {
	tcName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	rgName := acc.CdResourceGroupName
	toolName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCdToolchainToolPagerdutyDataSourceConfig(tcName, rgName, toolName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_pagerduty.cd_toolchain_tool_pagerduty", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_pagerduty.cd_toolchain_tool_pagerduty", "toolchain_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_pagerduty.cd_toolchain_tool_pagerduty", "tool_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_pagerduty.cd_toolchain_tool_pagerduty", "resource_group_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_pagerduty.cd_toolchain_tool_pagerduty", "crn"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_pagerduty.cd_toolchain_tool_pagerduty", "toolchain_crn"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_pagerduty.cd_toolchain_tool_pagerduty", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_pagerduty.cd_toolchain_tool_pagerduty", "referent.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_pagerduty.cd_toolchain_tool_pagerduty", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_pagerduty.cd_toolchain_tool_pagerduty", "updated_at"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_pagerduty.cd_toolchain_tool_pagerduty", "parameters.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_pagerduty.cd_toolchain_tool_pagerduty", "state"),
				),
			},
		},
	})
}

func testAccCheckIBMCdToolchainToolPagerdutyDataSourceConfigBasic(tcName string, rgName string) string {
	return fmt.Sprintf(`
		data "ibm_resource_group" "resource_group" {
			name = "%s"
		}

		resource "ibm_cd_toolchain" "cd_toolchain" {
			name = "%s"
			resource_group_id = data.ibm_resource_group.resource_group.id
		}

		resource "ibm_cd_toolchain_tool_pagerduty" "cd_toolchain_tool_pagerduty" {
			toolchain_id = ibm_cd_toolchain.cd_toolchain.id
			parameters {
				service_url = "https://mycompany.example.pagerduty.com/services/AS34FR4"
				service_key = "12345678901234567890123456789012"
			}
		}

		data "ibm_cd_toolchain_tool_pagerduty" "cd_toolchain_tool_pagerduty" {
			toolchain_id = ibm_cd_toolchain_tool_pagerduty.cd_toolchain_tool_pagerduty.toolchain_id
			tool_id = ibm_cd_toolchain_tool_pagerduty.cd_toolchain_tool_pagerduty.tool_id
		}
	`, rgName, tcName)
}

func testAccCheckIBMCdToolchainToolPagerdutyDataSourceConfig(tcName string, rgName string, toolName string) string {
	return fmt.Sprintf(`
		data "ibm_resource_group" "resource_group" {
			name = "%s"
		}

		resource "ibm_cd_toolchain" "cd_toolchain" {
			name = "%s"
			resource_group_id = data.ibm_resource_group.resource_group.id
		}

		resource "ibm_cd_toolchain_tool_pagerduty" "cd_toolchain_tool_pagerduty" {
			toolchain_id = ibm_cd_toolchain.cd_toolchain.id
			parameters {
				service_url = "https://mycompany.example.pagerduty.com/services/AS34FR4"
				service_key = "12345678901234567890123456789012"
			}
			name = "%s"
		}

		data "ibm_cd_toolchain_tool_pagerduty" "cd_toolchain_tool_pagerduty" {
			toolchain_id = ibm_cd_toolchain_tool_pagerduty.cd_toolchain_tool_pagerduty.toolchain_id
			tool_id = ibm_cd_toolchain_tool_pagerduty.cd_toolchain_tool_pagerduty.tool_id
		}
	`, rgName, tcName, toolName)
}
