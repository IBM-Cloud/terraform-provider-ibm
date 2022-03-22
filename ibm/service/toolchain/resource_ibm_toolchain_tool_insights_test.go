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

func TestAccIbmToolchainToolInsightsBasic(t *testing.T) {
	var conf toolchainv2.GetIntegrationByIdResponse
	toolchainID := fmt.Sprintf("tf_toolchain_id_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmToolchainToolInsightsDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmToolchainToolInsightsConfigBasic(toolchainID),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmToolchainToolInsightsExists("ibm_toolchain_tool_insights.toolchain_tool_insights", conf),
					resource.TestCheckResourceAttr("ibm_toolchain_tool_insights.toolchain_tool_insights", "toolchain_id", toolchainID),
				),
			},
		},
	})
}

func TestAccIbmToolchainToolInsightsAllArgs(t *testing.T) {
	var conf toolchainv2.GetIntegrationByIdResponse
	toolchainID := fmt.Sprintf("tf_toolchain_id_%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	nameUpdate := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmToolchainToolInsightsDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmToolchainToolInsightsConfig(toolchainID, name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmToolchainToolInsightsExists("ibm_toolchain_tool_insights.toolchain_tool_insights", conf),
					resource.TestCheckResourceAttr("ibm_toolchain_tool_insights.toolchain_tool_insights", "toolchain_id", toolchainID),
					resource.TestCheckResourceAttr("ibm_toolchain_tool_insights.toolchain_tool_insights", "name", name),
				),
			},
			resource.TestStep{
				Config: testAccCheckIbmToolchainToolInsightsConfig(toolchainID, nameUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_toolchain_tool_insights.toolchain_tool_insights", "toolchain_id", toolchainID),
					resource.TestCheckResourceAttr("ibm_toolchain_tool_insights.toolchain_tool_insights", "name", nameUpdate),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_toolchain_tool_insights.toolchain_tool_insights",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIbmToolchainToolInsightsConfigBasic(toolchainID string) string {
	return fmt.Sprintf(`

		resource "ibm_toolchain_tool_insights" "toolchain_tool_insights" {
			toolchain_id = "%s"
		}
	`, toolchainID)
}

func testAccCheckIbmToolchainToolInsightsConfig(toolchainID string, name string) string {
	return fmt.Sprintf(`

		resource "ibm_toolchain_tool_insights" "toolchain_tool_insights" {
			toolchain_id = "%s"
			name = "%s"
		}
	`, toolchainID, name)
}

func testAccCheckIbmToolchainToolInsightsExists(n string, obj toolchainv2.GetIntegrationByIdResponse) resource.TestCheckFunc {

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

func testAccCheckIbmToolchainToolInsightsDestroy(s *terraform.State) error {
	toolchainClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).ToolchainV2()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_toolchain_tool_insights" {
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
			return fmt.Errorf("toolchain_tool_insights still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for toolchain_tool_insights (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
