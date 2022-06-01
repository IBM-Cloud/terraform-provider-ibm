// Copyright IBM Corp. 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package cdtoolchain_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.ibm.com/org-ids/toolchain-go-sdk/cdtoolchainv2"
)

func TestAccIBMCdToolchainToolPagerdutyBasic(t *testing.T) {
	var conf cdtoolchainv2.GetIntegrationByIDResponse
	toolchainID := fmt.Sprintf("tf_toolchain_id_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCdToolchainToolPagerdutyDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCdToolchainToolPagerdutyConfigBasic(toolchainID),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCdToolchainToolPagerdutyExists("ibm_cd_toolchain_tool_pagerduty.cd_toolchain_tool_pagerduty", conf),
					resource.TestCheckResourceAttr("ibm_cd_toolchain_tool_pagerduty.cd_toolchain_tool_pagerduty", "toolchain_id", toolchainID),
				),
			},
		},
	})
}

func TestAccIBMCdToolchainToolPagerdutyAllArgs(t *testing.T) {
	var conf cdtoolchainv2.GetIntegrationByIDResponse
	toolchainID := fmt.Sprintf("tf_toolchain_id_%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	nameUpdate := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCdToolchainToolPagerdutyDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCdToolchainToolPagerdutyConfig(toolchainID, name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCdToolchainToolPagerdutyExists("ibm_cd_toolchain_tool_pagerduty.cd_toolchain_tool_pagerduty", conf),
					resource.TestCheckResourceAttr("ibm_cd_toolchain_tool_pagerduty.cd_toolchain_tool_pagerduty", "toolchain_id", toolchainID),
					resource.TestCheckResourceAttr("ibm_cd_toolchain_tool_pagerduty.cd_toolchain_tool_pagerduty", "name", name),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMCdToolchainToolPagerdutyConfig(toolchainID, nameUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_cd_toolchain_tool_pagerduty.cd_toolchain_tool_pagerduty", "toolchain_id", toolchainID),
					resource.TestCheckResourceAttr("ibm_cd_toolchain_tool_pagerduty.cd_toolchain_tool_pagerduty", "name", nameUpdate),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_cd_toolchain_tool_pagerduty.cd_toolchain_tool_pagerduty",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIBMCdToolchainToolPagerdutyConfigBasic(toolchainID string) string {
	return fmt.Sprintf(`

		resource "ibm_cd_toolchain_tool_pagerduty" "cd_toolchain_tool_pagerduty" {
			toolchain_id = "%s"
		}
	`, toolchainID)
}

func testAccCheckIBMCdToolchainToolPagerdutyConfig(toolchainID string, name string) string {
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
	`, toolchainID, name)
}

func testAccCheckIBMCdToolchainToolPagerdutyExists(n string, obj cdtoolchainv2.GetIntegrationByIDResponse) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		cdToolchainClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).CdToolchainV2()
		if err != nil {
			return err
		}

		getIntegrationByIDOptions := &cdtoolchainv2.GetIntegrationByIDOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getIntegrationByIDOptions.SetToolchainID(parts[0])
		getIntegrationByIDOptions.SetIntegrationID(parts[1])

		getIntegrationByIDResponse, _, err := cdToolchainClient.GetIntegrationByID(getIntegrationByIDOptions)
		if err != nil {
			return err
		}

		obj = *getIntegrationByIDResponse
		return nil
	}
}

func testAccCheckIBMCdToolchainToolPagerdutyDestroy(s *terraform.State) error {
	cdToolchainClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).CdToolchainV2()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_cd_toolchain_tool_pagerduty" {
			continue
		}

		getIntegrationByIDOptions := &cdtoolchainv2.GetIntegrationByIDOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getIntegrationByIDOptions.SetToolchainID(parts[0])
		getIntegrationByIDOptions.SetIntegrationID(parts[1])

		// Try to find the key
		_, response, err := cdToolchainClient.GetIntegrationByID(getIntegrationByIDOptions)

		if err == nil {
			return fmt.Errorf("cd_toolchain_tool_pagerduty still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for cd_toolchain_tool_pagerduty (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
