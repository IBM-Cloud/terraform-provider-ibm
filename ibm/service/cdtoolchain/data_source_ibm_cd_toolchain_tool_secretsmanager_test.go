// Copyright IBM Corp. 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package cdtoolchain_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIBMCdToolchainToolSecretsmanagerDataSourceBasic(t *testing.T) {
	tcName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	rgName := acc.CdResourceGroupName
	smName := acc.CdSecretsManagerInstanceName

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCdToolchainToolSecretsmanagerDataSourceConfigBasic(tcName, rgName, smName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_secretsmanager.cd_toolchain_tool_secretsmanager", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_secretsmanager.cd_toolchain_tool_secretsmanager", "toolchain_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_secretsmanager.cd_toolchain_tool_secretsmanager", "tool_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_secretsmanager.cd_toolchain_tool_secretsmanager", "resource_group_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_secretsmanager.cd_toolchain_tool_secretsmanager", "crn"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_secretsmanager.cd_toolchain_tool_secretsmanager", "toolchain_crn"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_secretsmanager.cd_toolchain_tool_secretsmanager", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_secretsmanager.cd_toolchain_tool_secretsmanager", "referent.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_secretsmanager.cd_toolchain_tool_secretsmanager", "updated_at"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_secretsmanager.cd_toolchain_tool_secretsmanager", "parameters.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_secretsmanager.cd_toolchain_tool_secretsmanager", "state"),
				),
			},
		},
	})
}

func TestAccIBMCdToolchainToolSecretsmanagerDataSourceAllArgs(t *testing.T) {
	tcName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	rgName := acc.CdResourceGroupName
	smName := acc.CdSecretsManagerInstanceName
	toolName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCdToolchainToolSecretsmanagerDataSourceConfig(tcName, rgName, smName, toolName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_secretsmanager.cd_toolchain_tool_secretsmanager", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_secretsmanager.cd_toolchain_tool_secretsmanager", "toolchain_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_secretsmanager.cd_toolchain_tool_secretsmanager", "tool_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_secretsmanager.cd_toolchain_tool_secretsmanager", "resource_group_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_secretsmanager.cd_toolchain_tool_secretsmanager", "crn"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_secretsmanager.cd_toolchain_tool_secretsmanager", "toolchain_crn"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_secretsmanager.cd_toolchain_tool_secretsmanager", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_secretsmanager.cd_toolchain_tool_secretsmanager", "referent.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_secretsmanager.cd_toolchain_tool_secretsmanager", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_secretsmanager.cd_toolchain_tool_secretsmanager", "updated_at"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_secretsmanager.cd_toolchain_tool_secretsmanager", "parameters.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_secretsmanager.cd_toolchain_tool_secretsmanager", "state"),
				),
			},
		},
	})
}

func testAccCheckIBMCdToolchainToolSecretsmanagerDataSourceConfigBasic(tcName string, rgName string, smName string) string {
	return fmt.Sprintf(`
		data "ibm_resource_group" "resource_group" {
			name = "%s"
		}

		resource "ibm_cd_toolchain" "cd_toolchain" {
			name = "%s"
			resource_group_id = data.ibm_resource_group.resource_group.id
		}

		resource "ibm_cd_toolchain_tool_secretsmanager" "cd_toolchain_tool_secretsmanager" {
			toolchain_id = ibm_cd_toolchain.cd_toolchain.id
			parameters {
				name = "sm_tool_01"
				instance_name = "%s"
				location = "eu-gb"
				resource_group_name = "%s"
			}
		}

		data "ibm_cd_toolchain_tool_secretsmanager" "cd_toolchain_tool_secretsmanager" {
			toolchain_id = ibm_cd_toolchain_tool_secretsmanager.cd_toolchain_tool_secretsmanager.toolchain_id
			tool_id = ibm_cd_toolchain_tool_secretsmanager.cd_toolchain_tool_secretsmanager.tool_id
		}
	`, rgName, tcName, smName, rgName)
}

func testAccCheckIBMCdToolchainToolSecretsmanagerDataSourceConfig(tcName string, rgName string, smName string, toolName string) string {
	return fmt.Sprintf(`
		data "ibm_resource_group" "resource_group" {
			name = "%s"
		}

		resource "ibm_cd_toolchain" "cd_toolchain" {
			name = "%s"
			resource_group_id = data.ibm_resource_group.resource_group.id
		}

		resource "ibm_cd_toolchain_tool_secretsmanager" "cd_toolchain_tool_secretsmanager" {
			toolchain_id = ibm_cd_toolchain.cd_toolchain.id
			parameters {
				name = "sm_tool_01"
				instance_name = "%s"
				location = "eu-gb"
				resource_group_name = "%s"
			}
			name = "%s"
		}

		data "ibm_cd_toolchain_tool_secretsmanager" "cd_toolchain_tool_secretsmanager" {
			toolchain_id = ibm_cd_toolchain_tool_secretsmanager.cd_toolchain_tool_secretsmanager.toolchain_id
			tool_id = ibm_cd_toolchain_tool_secretsmanager.cd_toolchain_tool_secretsmanager.tool_id
		}
	`, rgName, tcName, smName, rgName, toolName)
}
