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

func TestAccIBMCdToolchainToolPagerdutyDataSourceBasic(t *testing.T) {
	getIntegrationByIDResponseToolchainID := fmt.Sprintf("tf_toolchain_id_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCdToolchainToolPagerdutyDataSourceConfigBasic(getIntegrationByIDResponseToolchainID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_pagerduty.cd_toolchain_tool_pagerduty", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_pagerduty.cd_toolchain_tool_pagerduty", "toolchain_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_pagerduty.cd_toolchain_tool_pagerduty", "integration_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_pagerduty.cd_toolchain_tool_pagerduty", "get_integration_by_id_response_id"),
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
	getIntegrationByIDResponseToolchainID := fmt.Sprintf("tf_toolchain_id_%d", acctest.RandIntRange(10, 100))
	getIntegrationByIDResponseName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCdToolchainToolPagerdutyDataSourceConfig(getIntegrationByIDResponseToolchainID, getIntegrationByIDResponseName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_pagerduty.cd_toolchain_tool_pagerduty", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_pagerduty.cd_toolchain_tool_pagerduty", "toolchain_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_pagerduty.cd_toolchain_tool_pagerduty", "integration_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_pagerduty.cd_toolchain_tool_pagerduty", "get_integration_by_id_response_id"),
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

func testAccCheckIBMCdToolchainToolPagerdutyDataSourceConfigBasic(getIntegrationByIDResponseToolchainID string) string {
	return fmt.Sprintf(`
		resource "ibm_cd_toolchain_tool_pagerduty" "cd_toolchain_tool_pagerduty" {
			toolchain_id = "%s"
		}

		data "ibm_cd_toolchain_tool_pagerduty" "cd_toolchain_tool_pagerduty" {
			toolchain_id = ibm_cd_toolchain_tool_pagerduty.cd_toolchain_tool_pagerduty.toolchain_id
			integration_id = "integration_id"
		}
	`, getIntegrationByIDResponseToolchainID)
}

func testAccCheckIBMCdToolchainToolPagerdutyDataSourceConfig(getIntegrationByIDResponseToolchainID string, getIntegrationByIDResponseName string) string {
	return fmt.Sprintf(`
		resource "ibm_cd_toolchain_tool_pagerduty" "cd_toolchain_tool_pagerduty" {
			toolchain_id = "%s"
			name = "%s"
			parameters {
				key_type = "api"
				api_key = "api_key"
				service_name = "service_name"
				user_email = "user_email"
				user_phone = "user_phone"
				service_url = "service_url"
				service_key = "service_key"
				service_id = "service_id"
			}
		}

		data "ibm_cd_toolchain_tool_pagerduty" "cd_toolchain_tool_pagerduty" {
			toolchain_id = ibm_cd_toolchain_tool_pagerduty.cd_toolchain_tool_pagerduty.toolchain_id
			integration_id = "integration_id"
		}
	`, getIntegrationByIDResponseToolchainID, getIntegrationByIDResponseName)
}
