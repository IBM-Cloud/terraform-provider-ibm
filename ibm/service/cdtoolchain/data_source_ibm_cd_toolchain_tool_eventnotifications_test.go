// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package cdtoolchain_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"
)

func TestAccIBMCdToolchainToolEventnotificationsDataSourceBasic(t *testing.T) {
	tcName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	rgName := acc.CdResourceGroupName
	enName := acc.CdEventNotificationsInstanceName

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCdToolchainToolEventnotificationsDataSourceConfigBasic(tcName, rgName, enName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_eventnotifications.cd_toolchain_tool_eventnotifications", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_eventnotifications.cd_toolchain_tool_eventnotifications", "toolchain_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_eventnotifications.cd_toolchain_tool_eventnotifications", "tool_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_eventnotifications.cd_toolchain_tool_eventnotifications", "resource_group_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_eventnotifications.cd_toolchain_tool_eventnotifications", "crn"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_eventnotifications.cd_toolchain_tool_eventnotifications", "toolchain_crn"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_eventnotifications.cd_toolchain_tool_eventnotifications", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_eventnotifications.cd_toolchain_tool_eventnotifications", "referent.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_eventnotifications.cd_toolchain_tool_eventnotifications", "updated_at"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_eventnotifications.cd_toolchain_tool_eventnotifications", "parameters.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_eventnotifications.cd_toolchain_tool_eventnotifications", "state"),
				),
			},
		},
	})
}

func TestAccIBMCdToolchainToolEventnotificationsDataSourceAllArgs(t *testing.T) {
	tcName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	rgName := acc.CdResourceGroupName
	enName := acc.CdEventNotificationsInstanceName
	toolName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCdToolchainToolEventnotificationsDataSourceConfig(tcName, rgName, enName, toolName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_eventnotifications.cd_toolchain_tool_eventnotifications", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_eventnotifications.cd_toolchain_tool_eventnotifications", "toolchain_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_eventnotifications.cd_toolchain_tool_eventnotifications", "tool_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_eventnotifications.cd_toolchain_tool_eventnotifications", "resource_group_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_eventnotifications.cd_toolchain_tool_eventnotifications", "crn"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_eventnotifications.cd_toolchain_tool_eventnotifications", "toolchain_crn"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_eventnotifications.cd_toolchain_tool_eventnotifications", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_eventnotifications.cd_toolchain_tool_eventnotifications", "referent.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_eventnotifications.cd_toolchain_tool_eventnotifications", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_eventnotifications.cd_toolchain_tool_eventnotifications", "updated_at"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_eventnotifications.cd_toolchain_tool_eventnotifications", "parameters.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_eventnotifications.cd_toolchain_tool_eventnotifications", "state"),
				),
			},
		},
	})
}

func testAccCheckIBMCdToolchainToolEventnotificationsDataSourceConfigBasic(tcName string, rgName string, enName string) string {
	return fmt.Sprintf(`
		data "ibm_resource_group" "resource_group" {
			name = "%s"
		}

		resource "ibm_cd_toolchain" "cd_toolchain" {
			name = "%s"
			resource_group_id = data.ibm_resource_group.resource_group.id
		}

		data "ibm_resource_instance" "en_resource_instance" {
			name = "%s"
		}

		resource "ibm_iam_authorization_policy" "s2sAuth1" {
			source_service_name         = "toolchain"
			source_resource_instance_id = ibm_cd_toolchain.cd_toolchain.id
			target_service_name         = "event-notifications"
			target_resource_instance_id = data.ibm_resource_instance.en_resource_instance.guid
			roles                       = ["Reader", "Event Source Manager"]
		}

		resource "ibm_cd_toolchain_tool_eventnotifications" "cd_toolchain_tool_eventnotifications" {
			toolchain_id = ibm_cd_toolchain.cd_toolchain.id
			parameters {
				name = "en_tool_01"
				instance_crn = data.ibm_resource_instance.en_resource_instance.crn
			}
			depends_on = [
				ibm_iam_authorization_policy.s2sAuth1
			]		  
		}

		data "ibm_cd_toolchain_tool_eventnotifications" "cd_toolchain_tool_eventnotifications" {
			toolchain_id = ibm_cd_toolchain_tool_eventnotifications.cd_toolchain_tool_eventnotifications.toolchain_id
			tool_id = ibm_cd_toolchain_tool_eventnotifications.cd_toolchain_tool_eventnotifications.tool_id
		}
	`, rgName, tcName, enName)
}

func testAccCheckIBMCdToolchainToolEventnotificationsDataSourceConfig(tcName string, rgName string, enName string, toolName string) string {
	return fmt.Sprintf(`
		data "ibm_resource_group" "resource_group" {
			name = "%s"
		}

		resource "ibm_cd_toolchain" "cd_toolchain" {
			name = "%s"
			resource_group_id = data.ibm_resource_group.resource_group.id
		}

		data "ibm_resource_instance" "en_resource_instance" {
			name              = "%s"
		}

		resource "ibm_iam_authorization_policy" "s2sAuth1" {
			source_service_name         = "toolchain"
			source_resource_instance_id = ibm_cd_toolchain.cd_toolchain.id
			target_service_name         = "event-notifications"
			target_resource_instance_id = data.ibm_resource_instance.en_resource_instance.guid
			roles                       = ["Reader", "Event Source Manager"]
		}

		resource "ibm_cd_toolchain_tool_eventnotifications" "cd_toolchain_tool_eventnotifications" {
			toolchain_id = ibm_cd_toolchain.cd_toolchain.id
			parameters {
				name = "en_tool_01"
				instance_crn = data.ibm_resource_instance.en_resource_instance.crn
			}
			name = "%s"
			depends_on = [
				ibm_iam_authorization_policy.s2sAuth1
			]		  
		}

		data "ibm_cd_toolchain_tool_eventnotifications" "cd_toolchain_tool_eventnotifications" {
			toolchain_id = ibm_cd_toolchain_tool_eventnotifications.cd_toolchain_tool_eventnotifications.toolchain_id
			tool_id = ibm_cd_toolchain_tool_eventnotifications.cd_toolchain_tool_eventnotifications.tool_id
		}
	`, rgName, tcName, enName, toolName)
}
