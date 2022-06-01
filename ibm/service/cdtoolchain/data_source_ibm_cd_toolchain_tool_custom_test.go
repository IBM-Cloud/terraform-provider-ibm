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

func TestAccIBMCdToolchainToolCustomDataSourceBasic(t *testing.T) {
	getIntegrationByIDResponseToolchainID := fmt.Sprintf("tf_toolchain_id_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCdToolchainToolCustomDataSourceConfigBasic(getIntegrationByIDResponseToolchainID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_custom.cd_toolchain_tool_custom", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_custom.cd_toolchain_tool_custom", "toolchain_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_custom.cd_toolchain_tool_custom", "integration_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_custom.cd_toolchain_tool_custom", "get_integration_by_id_response_id"),
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
	getIntegrationByIDResponseToolchainID := fmt.Sprintf("tf_toolchain_id_%d", acctest.RandIntRange(10, 100))
	getIntegrationByIDResponseName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCdToolchainToolCustomDataSourceConfig(getIntegrationByIDResponseToolchainID, getIntegrationByIDResponseName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_custom.cd_toolchain_tool_custom", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_custom.cd_toolchain_tool_custom", "toolchain_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_custom.cd_toolchain_tool_custom", "integration_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_custom.cd_toolchain_tool_custom", "get_integration_by_id_response_id"),
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

func testAccCheckIBMCdToolchainToolCustomDataSourceConfigBasic(getIntegrationByIDResponseToolchainID string) string {
	return fmt.Sprintf(`
		resource "ibm_cd_toolchain_tool_custom" "cd_toolchain_tool_custom" {
			toolchain_id = "%s"
		}

		data "ibm_cd_toolchain_tool_custom" "cd_toolchain_tool_custom" {
			toolchain_id = ibm_cd_toolchain_tool_custom.cd_toolchain_tool_custom.toolchain_id
			integration_id = "integration_id"
		}
	`, getIntegrationByIDResponseToolchainID)
}

func testAccCheckIBMCdToolchainToolCustomDataSourceConfig(getIntegrationByIDResponseToolchainID string, getIntegrationByIDResponseName string) string {
	return fmt.Sprintf(`
		resource "ibm_cd_toolchain_tool_custom" "cd_toolchain_tool_custom" {
			toolchain_id = "%s"
			name = "%s"
			parameters {
				type = "type"
				lifecyclePhase = "THINK"
				imageUrl = "imageUrl"
				documentationUrl = "documentationUrl"
				name = "name"
				dashboard_url = "dashboard_url"
				description = "description"
				additional-properties = "additional-properties"
			}
		}

		data "ibm_cd_toolchain_tool_custom" "cd_toolchain_tool_custom" {
			toolchain_id = ibm_cd_toolchain_tool_custom.cd_toolchain_tool_custom.toolchain_id
			integration_id = "integration_id"
		}
	`, getIntegrationByIDResponseToolchainID, getIntegrationByIDResponseName)
}
