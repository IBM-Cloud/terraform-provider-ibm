// Copyright IBM Corp. 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package toolchain_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.ibm.com/org-ids/toolchain-go-sdk/toolchainv2"
)

func TestAccIbmToolchainToolKeyprotectBasic(t *testing.T) {
	var conf toolchainv2.GetIntegrationByIdResponse
	toolchainID := fmt.Sprintf("tf_toolchain_id_%d", acctest.RandIntRange(10, 100))
	serviceID := fmt.Sprintf("tf_service_id_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmToolchainToolKeyprotectDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmToolchainToolKeyprotectConfigBasic(toolchainID, serviceID),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmToolchainToolKeyprotectExists("ibm_toolchain_tool_keyprotect.toolchain_tool_keyprotect", conf),
					resource.TestCheckResourceAttr("ibm_toolchain_tool_keyprotect.toolchain_tool_keyprotect", "toolchain_id", toolchainID),
					resource.TestCheckResourceAttr("ibm_toolchain_tool_keyprotect.toolchain_tool_keyprotect", "service_id", serviceID),
				),
			},
		},
	})
}

func TestAccIbmToolchainToolKeyprotectAllArgs(t *testing.T) {
	var conf toolchainv2.GetIntegrationByIdResponse
	toolchainID := fmt.Sprintf("tf_toolchain_id_%d", acctest.RandIntRange(10, 100))
	serviceID := fmt.Sprintf("tf_service_id_%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	nameUpdate := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmToolchainToolKeyprotectDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmToolchainToolKeyprotectConfig(toolchainID, serviceID, name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmToolchainToolKeyprotectExists("ibm_toolchain_tool_keyprotect.toolchain_tool_keyprotect", conf),
					resource.TestCheckResourceAttr("ibm_toolchain_tool_keyprotect.toolchain_tool_keyprotect", "toolchain_id", toolchainID),
					resource.TestCheckResourceAttr("ibm_toolchain_tool_keyprotect.toolchain_tool_keyprotect", "service_id", serviceID),
					resource.TestCheckResourceAttr("ibm_toolchain_tool_keyprotect.toolchain_tool_keyprotect", "name", name),
				),
			},
			resource.TestStep{
				Config: testAccCheckIbmToolchainToolKeyprotectConfig(toolchainID, serviceID, nameUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_toolchain_tool_keyprotect.toolchain_tool_keyprotect", "toolchain_id", toolchainID),
					resource.TestCheckResourceAttr("ibm_toolchain_tool_keyprotect.toolchain_tool_keyprotect", "service_id", serviceID),
					resource.TestCheckResourceAttr("ibm_toolchain_tool_keyprotect.toolchain_tool_keyprotect", "name", nameUpdate),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_toolchain_tool_keyprotect.toolchain_tool_keyprotect",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIbmToolchainToolKeyprotectConfigBasic(toolchainID string, serviceID string) string {
	return fmt.Sprintf(`

		resource "ibm_toolchain_tool_keyprotect" "toolchain_tool_keyprotect" {
			toolchain_id = "%s"
			service_id = "%s"
		}
	`, toolchainID, serviceID)
}

func testAccCheckIbmToolchainToolKeyprotectConfig(toolchainID string, serviceID string, name string) string {
	return fmt.Sprintf(`

		resource "ibm_toolchain_tool_keyprotect" "toolchain_tool_keyprotect" {
			toolchain_id = "%s"
			service_id = "%s"
			name = "%s"
			parameters {
				name = "name"
				region = "region"
				resource_group = "resource_group"
				instance_name = "instance_name"
			}
			parameters_references = "FIXME"
		}
	`, toolchainID, serviceID, name)
}

func testAccCheckIbmToolchainToolKeyprotectExists(n string, obj toolchainv2.GetIntegrationByIdResponse) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		toolchainClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).ToolchainV2()
		if err != nil {
			return err
		}

		getIntegrationByIdOptions := &toolchainv2.GetIntegrationByIdOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getIntegrationByIdOptions.SetToolchainID(parts[0])
		getIntegrationByIdOptions.SetIntegrationID(parts[1])

		getIntegrationByIdResponse, _, err := toolchainClient.GetIntegrationByID(getIntegrationByIdOptions)
		if err != nil {
			return err
		}

		obj = *getIntegrationByIdResponse
		return nil
	}
}

func testAccCheckIbmToolchainToolKeyprotectDestroy(s *terraform.State) error {
	toolchainClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).ToolchainV2()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_toolchain_tool_keyprotect" {
			continue
		}

		getIntegrationByIdOptions := &toolchainv2.GetIntegrationByIdOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getIntegrationByIdOptions.SetToolchainID(parts[0])
		getIntegrationByIdOptions.SetIntegrationID(parts[1])

		// Try to find the key
		_, response, err := toolchainClient.GetIntegrationByID(getIntegrationByIdOptions)

		if err == nil {
			return fmt.Errorf("toolchain_tool_keyprotect still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for toolchain_tool_keyprotect (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
