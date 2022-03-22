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

func TestAccIbmToolchainToolSonarqubeBasic(t *testing.T) {
	var conf toolchainv2.GetIntegrationByIdResponse
	toolchainID := fmt.Sprintf("tf_toolchain_id_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmToolchainToolSonarqubeDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmToolchainToolSonarqubeConfigBasic(toolchainID),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmToolchainToolSonarqubeExists("ibm_toolchain_tool_sonarqube.toolchain_tool_sonarqube", conf),
					resource.TestCheckResourceAttr("ibm_toolchain_tool_sonarqube.toolchain_tool_sonarqube", "toolchain_id", toolchainID),
				),
			},
		},
	})
}

func TestAccIbmToolchainToolSonarqubeAllArgs(t *testing.T) {
	var conf toolchainv2.GetIntegrationByIdResponse
	toolchainID := fmt.Sprintf("tf_toolchain_id_%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	nameUpdate := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmToolchainToolSonarqubeDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmToolchainToolSonarqubeConfig(toolchainID, name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmToolchainToolSonarqubeExists("ibm_toolchain_tool_sonarqube.toolchain_tool_sonarqube", conf),
					resource.TestCheckResourceAttr("ibm_toolchain_tool_sonarqube.toolchain_tool_sonarqube", "toolchain_id", toolchainID),
					resource.TestCheckResourceAttr("ibm_toolchain_tool_sonarqube.toolchain_tool_sonarqube", "name", name),
				),
			},
			resource.TestStep{
				Config: testAccCheckIbmToolchainToolSonarqubeConfig(toolchainID, nameUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_toolchain_tool_sonarqube.toolchain_tool_sonarqube", "toolchain_id", toolchainID),
					resource.TestCheckResourceAttr("ibm_toolchain_tool_sonarqube.toolchain_tool_sonarqube", "name", nameUpdate),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_toolchain_tool_sonarqube.toolchain_tool_sonarqube",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIbmToolchainToolSonarqubeConfigBasic(toolchainID string) string {
	return fmt.Sprintf(`

		resource "ibm_toolchain_tool_sonarqube" "toolchain_tool_sonarqube" {
			toolchain_id = "%s"
		}
	`, toolchainID)
}

func testAccCheckIbmToolchainToolSonarqubeConfig(toolchainID string, name string) string {
	return fmt.Sprintf(`

		resource "ibm_toolchain_tool_sonarqube" "toolchain_tool_sonarqube" {
			toolchain_id = "%s"
			name = "%s"
			parameters {
				name = "name"
				dashboard_url = "dashboard_url"
				user_login = "user_login"
				user_password = "user_password"
				blind_connection = true
			}
			parameters_references = "FIXME"
		}
	`, toolchainID, name)
}

func testAccCheckIbmToolchainToolSonarqubeExists(n string, obj toolchainv2.GetIntegrationByIdResponse) resource.TestCheckFunc {

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

func testAccCheckIbmToolchainToolSonarqubeDestroy(s *terraform.State) error {
	toolchainClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).ToolchainV2()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_toolchain_tool_sonarqube" {
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
			return fmt.Errorf("toolchain_tool_sonarqube still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for toolchain_tool_sonarqube (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
