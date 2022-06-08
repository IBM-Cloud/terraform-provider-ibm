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

func TestAccIBMCdToolchainToolSonarqubeDataSourceBasic(t *testing.T) {
	getToolByIDResponseToolchainID := fmt.Sprintf("tf_toolchain_id_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCdToolchainToolSonarqubeDataSourceConfigBasic(getToolByIDResponseToolchainID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_sonarqube.cd_toolchain_tool_sonarqube", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_sonarqube.cd_toolchain_tool_sonarqube", "toolchain_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_sonarqube.cd_toolchain_tool_sonarqube", "tool_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_sonarqube.cd_toolchain_tool_sonarqube", "get_tool_by_id_response_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_sonarqube.cd_toolchain_tool_sonarqube", "resource_group_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_sonarqube.cd_toolchain_tool_sonarqube", "crn"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_sonarqube.cd_toolchain_tool_sonarqube", "toolchain_crn"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_sonarqube.cd_toolchain_tool_sonarqube", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_sonarqube.cd_toolchain_tool_sonarqube", "referent.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_sonarqube.cd_toolchain_tool_sonarqube", "updated_at"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_sonarqube.cd_toolchain_tool_sonarqube", "parameters.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_sonarqube.cd_toolchain_tool_sonarqube", "state"),
				),
			},
		},
	})
}

func TestAccIBMCdToolchainToolSonarqubeDataSourceAllArgs(t *testing.T) {
	getToolByIDResponseToolchainID := fmt.Sprintf("tf_toolchain_id_%d", acctest.RandIntRange(10, 100))
	getToolByIDResponseName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCdToolchainToolSonarqubeDataSourceConfig(getToolByIDResponseToolchainID, getToolByIDResponseName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_sonarqube.cd_toolchain_tool_sonarqube", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_sonarqube.cd_toolchain_tool_sonarqube", "toolchain_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_sonarqube.cd_toolchain_tool_sonarqube", "tool_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_sonarqube.cd_toolchain_tool_sonarqube", "get_tool_by_id_response_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_sonarqube.cd_toolchain_tool_sonarqube", "resource_group_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_sonarqube.cd_toolchain_tool_sonarqube", "crn"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_sonarqube.cd_toolchain_tool_sonarqube", "toolchain_crn"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_sonarqube.cd_toolchain_tool_sonarqube", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_sonarqube.cd_toolchain_tool_sonarqube", "referent.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_sonarqube.cd_toolchain_tool_sonarqube", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_sonarqube.cd_toolchain_tool_sonarqube", "updated_at"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_sonarqube.cd_toolchain_tool_sonarqube", "parameters.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_sonarqube.cd_toolchain_tool_sonarqube", "state"),
				),
			},
		},
	})
}

func testAccCheckIBMCdToolchainToolSonarqubeDataSourceConfigBasic(getToolByIDResponseToolchainID string) string {
	return fmt.Sprintf(`
		resource "ibm_cd_toolchain_tool_sonarqube" "cd_toolchain_tool_sonarqube" {
			toolchain_id = "%s"
		}

		data "ibm_cd_toolchain_tool_sonarqube" "cd_toolchain_tool_sonarqube" {
			toolchain_id = ibm_cd_toolchain_tool_sonarqube.cd_toolchain_tool_sonarqube.toolchain_id
			tool_id = "tool_id"
		}
	`, getToolByIDResponseToolchainID)
}

func testAccCheckIBMCdToolchainToolSonarqubeDataSourceConfig(getToolByIDResponseToolchainID string, getToolByIDResponseName string) string {
	return fmt.Sprintf(`
		resource "ibm_cd_toolchain_tool_sonarqube" "cd_toolchain_tool_sonarqube" {
			toolchain_id = "%s"
			name = "%s"
			parameters {
				name = "name"
				dashboard_url = "dashboard_url"
				user_login = "user_login"
				user_password = "user_password"
				blind_connection = true
			}
		}

		data "ibm_cd_toolchain_tool_sonarqube" "cd_toolchain_tool_sonarqube" {
			toolchain_id = ibm_cd_toolchain_tool_sonarqube.cd_toolchain_tool_sonarqube.toolchain_id
			tool_id = "tool_id"
		}
	`, getToolByIDResponseToolchainID, getToolByIDResponseName)
}
