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

func TestAccIBMCdToolchainToolSonarqubeDataSourceBasic(t *testing.T) {
	tcName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	rgName := acc.CdResourceGroupName

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCdToolchainToolSonarqubeDataSourceConfigBasic(tcName, rgName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_sonarqube.cd_toolchain_tool_sonarqube", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_sonarqube.cd_toolchain_tool_sonarqube", "toolchain_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_sonarqube.cd_toolchain_tool_sonarqube", "tool_id"),
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
	tcName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	rgName := acc.CdResourceGroupName
	toolName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCdToolchainToolSonarqubeDataSourceConfig(tcName, rgName, toolName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_sonarqube.cd_toolchain_tool_sonarqube", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_sonarqube.cd_toolchain_tool_sonarqube", "toolchain_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_sonarqube.cd_toolchain_tool_sonarqube", "tool_id"),
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

func testAccCheckIBMCdToolchainToolSonarqubeDataSourceConfigBasic(tcName string, rgName string) string {
	return fmt.Sprintf(`
		data "ibm_resource_group" "resource_group" {
			name = "%s"
		}

		resource "ibm_cd_toolchain" "cd_toolchain" {
			name = "%s"
			resource_group_id = data.ibm_resource_group.resource_group.id
		}

		resource "ibm_cd_toolchain_tool_sonarqube" "cd_toolchain_tool_sonarqube" {
			toolchain_id = ibm_cd_toolchain.cd_toolchain.id
			parameters {
				name = "my-sonarqube"
				user_login = "<user_login>"
				user_password = "<user_password>"
				blind_connection = true
				server_url = "https://my.sonarqube.server.com/"
			}
		}

		data "ibm_cd_toolchain_tool_sonarqube" "cd_toolchain_tool_sonarqube" {
			toolchain_id = ibm_cd_toolchain_tool_sonarqube.cd_toolchain_tool_sonarqube.toolchain_id
			tool_id = ibm_cd_toolchain_tool_sonarqube.cd_toolchain_tool_sonarqube.tool_id
		}
	`, rgName, tcName)
}

func testAccCheckIBMCdToolchainToolSonarqubeDataSourceConfig(tcName string, rgName string, toolName string) string {
	return fmt.Sprintf(`
		data "ibm_resource_group" "resource_group" {
			name = "%s"
		}

		resource "ibm_cd_toolchain" "cd_toolchain" {
			name = "%s"
			resource_group_id = data.ibm_resource_group.resource_group.id
		}

		resource "ibm_cd_toolchain_tool_sonarqube" "cd_toolchain_tool_sonarqube" {
			toolchain_id = ibm_cd_toolchain.cd_toolchain.id
			parameters {
				name = "my-sonarqube"
				user_login = "<user_login>"
				user_password = "<user_password>"
				blind_connection = true
				server_url = "https://my.sonarqube.server.com/"
			}
			name = "%s"
		}

		data "ibm_cd_toolchain_tool_sonarqube" "cd_toolchain_tool_sonarqube" {
			toolchain_id = ibm_cd_toolchain_tool_sonarqube.cd_toolchain_tool_sonarqube.toolchain_id
			tool_id = ibm_cd_toolchain_tool_sonarqube.cd_toolchain_tool_sonarqube.tool_id
		}
	`, rgName, tcName, toolName)
}
