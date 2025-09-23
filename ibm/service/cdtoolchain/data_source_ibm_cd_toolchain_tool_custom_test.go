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

func TestAccIBMCdToolchainToolCustomDataSourceBasic(t *testing.T) {
	tcName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	rgName := acc.CdResourceGroupName

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCdToolchainToolCustomDataSourceConfigBasic(tcName, rgName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_custom.cd_toolchain_tool_custom", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_custom.cd_toolchain_tool_custom", "toolchain_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_custom.cd_toolchain_tool_custom", "tool_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_custom.cd_toolchain_tool_custom", "resource_group_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_custom.cd_toolchain_tool_custom", "crn"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_custom.cd_toolchain_tool_custom", "toolchain_crn"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_custom.cd_toolchain_tool_custom", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_custom.cd_toolchain_tool_custom", "referent.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_custom.cd_toolchain_tool_custom", "updated_at"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_custom.cd_toolchain_tool_custom", "parameters.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_custom.cd_toolchain_tool_custom", "state"),
				),
			},
		},
	})
}

func TestAccIBMCdToolchainToolCustomDataSourceAllArgs(t *testing.T) {
	tcName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	rgName := acc.CdResourceGroupName
	toolName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCdToolchainToolCustomDataSourceConfig(tcName, rgName, toolName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_custom.cd_toolchain_tool_custom", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_custom.cd_toolchain_tool_custom", "toolchain_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_custom.cd_toolchain_tool_custom", "tool_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_custom.cd_toolchain_tool_custom", "resource_group_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_custom.cd_toolchain_tool_custom", "crn"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_custom.cd_toolchain_tool_custom", "toolchain_crn"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_custom.cd_toolchain_tool_custom", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_custom.cd_toolchain_tool_custom", "referent.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_custom.cd_toolchain_tool_custom", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_custom.cd_toolchain_tool_custom", "updated_at"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_custom.cd_toolchain_tool_custom", "parameters.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_custom.cd_toolchain_tool_custom", "state"),
				),
			},
		},
	})
}

func testAccCheckIBMCdToolchainToolCustomDataSourceConfigBasic(tcName string, rgName string) string {
	return fmt.Sprintf(`
		data "ibm_resource_group" "resource_group" {
			name = "%s"
		}

		resource "ibm_cd_toolchain" "cd_toolchain" {
			name = "%s"
			resource_group_id = data.ibm_resource_group.resource_group.id
		}

		resource "ibm_cd_toolchain_tool_custom" "cd_toolchain_tool_custom" {
			toolchain_id = ibm_cd_toolchain.cd_toolchain.id
			parameters {
				type = "Delivery Pipeline"
				lifecycle_phase = "DELIVER"
				image_url = "image_url"
				documentation_url = "documentation_url"
				name = "My Build and Deploy Pipeline"
				dashboard_url = "https://cloud.ibm.com/devops/pipelines/tekton/ae47390c-9495-4b0b-a489-78464685acdd"
				description = "description"
				additional_properties = "additional_properties"
			}
		}

		data "ibm_cd_toolchain_tool_custom" "cd_toolchain_tool_custom" {
			toolchain_id = ibm_cd_toolchain_tool_custom.cd_toolchain_tool_custom.toolchain_id
			tool_id = ibm_cd_toolchain_tool_custom.cd_toolchain_tool_custom.tool_id
		}
	`, rgName, tcName)
}

func testAccCheckIBMCdToolchainToolCustomDataSourceConfig(tcName string, rgName string, toolName string) string {
	return fmt.Sprintf(`
		data "ibm_resource_group" "resource_group" {
			name = "%s"
		}

		resource "ibm_cd_toolchain" "cd_toolchain" {
			name = "%s"
			resource_group_id = data.ibm_resource_group.resource_group.id
		}

		resource "ibm_cd_toolchain_tool_custom" "cd_toolchain_tool_custom" {
			toolchain_id = ibm_cd_toolchain.cd_toolchain.id
			parameters {
				type = "Delivery Pipeline"
				lifecycle_phase = "DELIVER"
				image_url = "image_url"
				documentation_url = "documentation_url"
				name = "My Build and Deploy Pipeline"
				dashboard_url = "https://cloud.ibm.com/devops/pipelines/tekton/ae47390c-9495-4b0b-a489-78464685acdd"
				description = "description"
				additional_properties = "additional_properties"
			}
			name = "%s"
		}

		data "ibm_cd_toolchain_tool_custom" "cd_toolchain_tool_custom" {
			toolchain_id = ibm_cd_toolchain_tool_custom.cd_toolchain_tool_custom.toolchain_id
			tool_id = ibm_cd_toolchain_tool_custom.cd_toolchain_tool_custom.tool_id
		}
	`, rgName, tcName, toolName)
}
