// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package cdtoolchain_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/IBM/continuous-delivery-go-sdk/v2/cdtoolchainv2"
	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/conns"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/flex"
)

func TestAccIBMCdToolchainToolEventnotificationsBasic(t *testing.T) {
	var conf cdtoolchainv2.ToolchainTool
	rgName := acc.CdResourceGroupName
	tcName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	enName := acc.CdEventNotificationsInstanceName

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCdToolchainToolEventnotificationsDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCdToolchainToolEventnotificationsConfigBasic(tcName, rgName, enName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCdToolchainToolEventnotificationsExists("ibm_cd_toolchain_tool_eventnotifications.cd_toolchain_tool_eventnotifications", conf),
					resource.TestCheckResourceAttrSet("ibm_cd_toolchain_tool_eventnotifications.cd_toolchain_tool_eventnotifications", "toolchain_id"),
				),
			},
		},
	})
}

func TestAccIBMCdToolchainToolEventnotificationsAllArgs(t *testing.T) {
	var conf cdtoolchainv2.ToolchainTool
	rgName := acc.CdResourceGroupName
	tcName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	enName := acc.CdEventNotificationsInstanceName
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	nameUpdate := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCdToolchainToolEventnotificationsDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCdToolchainToolEventnotificationsConfig(tcName, rgName, enName, name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCdToolchainToolEventnotificationsExists("ibm_cd_toolchain_tool_eventnotifications.cd_toolchain_tool_eventnotifications", conf),
					resource.TestCheckResourceAttrSet("ibm_cd_toolchain_tool_eventnotifications.cd_toolchain_tool_eventnotifications", "toolchain_id"),
					resource.TestCheckResourceAttr("ibm_cd_toolchain_tool_eventnotifications.cd_toolchain_tool_eventnotifications", "name", name),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMCdToolchainToolEventnotificationsConfig(tcName, rgName, enName, nameUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("ibm_cd_toolchain_tool_eventnotifications.cd_toolchain_tool_eventnotifications", "toolchain_id"),
					resource.TestCheckResourceAttr("ibm_cd_toolchain_tool_eventnotifications.cd_toolchain_tool_eventnotifications", "name", nameUpdate),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_cd_toolchain_tool_eventnotifications.cd_toolchain_tool_eventnotifications",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIBMCdToolchainToolEventnotificationsConfigBasic(tcName string, rgName string, enName string) string {
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
	`, rgName, tcName, enName)
}

func testAccCheckIBMCdToolchainToolEventnotificationsConfig(tcName string, rgName string, enName string, name string) string {
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
			name = "%s"
			depends_on = [
				ibm_iam_authorization_policy.s2sAuth1
			]		  
		}
`, rgName, tcName, enName, name)
}

func testAccCheckIBMCdToolchainToolEventnotificationsExists(n string, obj cdtoolchainv2.ToolchainTool) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		cdToolchainClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).CdToolchainV2()
		if err != nil {
			return err
		}

		getToolByIDOptions := &cdtoolchainv2.GetToolByIDOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getToolByIDOptions.SetToolchainID(parts[0])
		getToolByIDOptions.SetToolID(parts[1])

		toolchainTool, _, err := cdToolchainClient.GetToolByID(getToolByIDOptions)
		if err != nil {
			return err
		}

		obj = *toolchainTool
		return nil
	}
}

func testAccCheckIBMCdToolchainToolEventnotificationsDestroy(s *terraform.State) error {
	cdToolchainClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).CdToolchainV2()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_cd_toolchain_tool_eventnotifications" {
			continue
		}

		getToolByIDOptions := &cdtoolchainv2.GetToolByIDOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getToolByIDOptions.SetToolchainID(parts[0])
		getToolByIDOptions.SetToolID(parts[1])

		// Try to find the key
		_, response, err := cdToolchainClient.GetToolByID(getToolByIDOptions)

		if err == nil {
			return fmt.Errorf("cd_toolchain_tool_eventnotifications still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for cd_toolchain_tool_eventnotifications (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
