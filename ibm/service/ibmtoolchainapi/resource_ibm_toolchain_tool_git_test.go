// Copyright IBM Corp. 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibmtoolchainapi_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.ibm.com/org-ids/toolchain-go-sdk/ibmtoolchainapiv2"
)

func TestAccIbmToolchainToolGitBasic(t *testing.T) {
	var conf ibmtoolchainapiv2.ServiceResponse
	toolchainID := fmt.Sprintf("tf_toolchain_id_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmToolchainToolGitDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmToolchainToolGitConfigBasic(toolchainID),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmToolchainToolGitExists("ibm_toolchain_tool_git.toolchain_tool_git", conf),
					resource.TestCheckResourceAttr("ibm_toolchain_tool_git.toolchain_tool_git", "toolchain_id", toolchainID),
				),
			},
		},
	})
}

func TestAccIbmToolchainToolGitAllArgs(t *testing.T) {
	var conf ibmtoolchainapiv2.ServiceResponse
	toolchainID := fmt.Sprintf("tf_toolchain_id_%d", acctest.RandIntRange(10, 100))
	provider := fmt.Sprintf("tf_provider_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmToolchainToolGitDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmToolchainToolGitConfig(toolchainID, provider),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmToolchainToolGitExists("ibm_toolchain_tool_git.toolchain_tool_git", conf),
					resource.TestCheckResourceAttr("ibm_toolchain_tool_git.toolchain_tool_git", "toolchain_id", toolchainID),
					resource.TestCheckResourceAttr("ibm_toolchain_tool_git.toolchain_tool_git", "provider", provider),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_toolchain_tool_git.toolchain_tool_git",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIbmToolchainToolGitConfigBasic(toolchainID string) string {
	return fmt.Sprintf(`

		resource "ibm_toolchain_tool_git" "toolchain_tool_git" {
			toolchain_id = "%s"
		}
	`, toolchainID)
}

func testAccCheckIbmToolchainToolGitConfig(toolchainID string, provider string) string {
	return fmt.Sprintf(`

		resource "ibm_toolchain_tool_git" "toolchain_tool_git" {
			toolchain_id = "%s"
			provider = "%s"
			parameters {
				repo_url = "repo_url"
				action = "clone"
				legal = true
				enable_traceability = true
				has_issues = true
			}
			parameters_references = "FIXME"
			container {
				guid = "d02d29f1-e7bb-4977-8a6f-26d7b7bb893e"
				type = "organization_guid"
			}
		}
	`, toolchainID, provider)
}

func testAccCheckIbmToolchainToolGitExists(n string, obj ibmtoolchainapiv2.ServiceResponse) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		ibmToolchainApiClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).IbmToolchainApiV2()
		if err != nil {
			return err
		}

		getServiceInstanceOptions := &ibmtoolchainapiv2.GetServiceInstanceOptions{}

		getServiceInstanceOptions.SetServiceInstanceID(rs.Primary.ID)

		serviceResponse, _, err := ibmToolchainApiClient.GetServiceInstance(getServiceInstanceOptions)
		if err != nil {
			return err
		}

		obj = *serviceResponse
		return nil
	}
}

func testAccCheckIbmToolchainToolGitDestroy(s *terraform.State) error {
	ibmToolchainApiClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).IbmToolchainApiV2()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_toolchain_tool_git" {
			continue
		}

		getServiceInstanceOptions := &ibmtoolchainapiv2.GetServiceInstanceOptions{}

		getServiceInstanceOptions.SetServiceInstanceID(rs.Primary.ID)

		// Try to find the key
		_, response, err := ibmToolchainApiClient.GetServiceInstance(getServiceInstanceOptions)

		if err == nil {
			return fmt.Errorf("toolchain_tool_git still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for toolchain_tool_git (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
