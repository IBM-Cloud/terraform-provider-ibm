// Copyright IBM Corp. 2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package cdtoolchain_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/continuous-delivery-go-sdk/v2/cdtoolchainv2"
)

func TestAccIBMCdToolchainToolCosBasic(t *testing.T) {
	var conf cdtoolchainv2.ToolchainTool
	toolchainID := acc.ToolchainID

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCdToolchainToolCosDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCdToolchainToolCosConfigBasic(toolchainID, acc.COSApiKey, acc.CosCRN, acc.BucketName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCdToolchainToolCosExists("ibm_cd_toolchain_tool_cos.cd_toolchain_tool_cos_instance", conf),
					resource.TestCheckResourceAttr("ibm_cd_toolchain_tool_cos.cd_toolchain_tool_cos_instance", "toolchain_id", toolchainID),
				),
			},
		},
	})
}

func TestAccIBMCdToolchainToolCosAllArgs(t *testing.T) {
	var conf cdtoolchainv2.ToolchainTool
	toolchainID := acc.ToolchainID
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	nameUpdate := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCdToolchainToolCosDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCdToolchainToolCosConfig(toolchainID, name, acc.COSApiKey, acc.CosCRN, acc.BucketName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCdToolchainToolCosExists("ibm_cd_toolchain_tool_cos.cd_toolchain_tool_cos_instance", conf),
					resource.TestCheckResourceAttr("ibm_cd_toolchain_tool_cos.cd_toolchain_tool_cos_instance", "toolchain_id", toolchainID),
					resource.TestCheckResourceAttr("ibm_cd_toolchain_tool_cos.cd_toolchain_tool_cos_instance", "name", name),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMCdToolchainToolCosConfig(toolchainID, nameUpdate, acc.COSApiKey, acc.CosCRN, acc.BucketName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_cd_toolchain_tool_cos.cd_toolchain_tool_cos_instance", "toolchain_id", toolchainID),
					resource.TestCheckResourceAttr("ibm_cd_toolchain_tool_cos.cd_toolchain_tool_cos_instance", "name", nameUpdate),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_cd_toolchain_tool_cos.cd_toolchain_tool_cos_instance",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIBMCdToolchainToolCosConfigBasic(toolchainID string, cosAPIKey string, cosCRN string, bucketName string) string {
	return fmt.Sprintf(`
		resource "ibm_cd_toolchain_tool_cos" "cd_toolchain_tool_cos_instance" {
			toolchain_id = "%s"
			parameters {
				name = "cos_tool_01"
				cos_api_key = "%s"
				instance_crn = "%s"
				bucket_name = "%s"
				endpoint = "s3.direct.us-south.cloud-object-storage.appdomain.cloud"
				hmac_access_key_id = "hmac_access_key_id"
				hmac_secret_access_key = "hmac_secret_access_key"
			}
		}
	`, toolchainID, cosAPIKey, cosCRN, bucketName)
}

func testAccCheckIBMCdToolchainToolCosConfig(toolchainID string, name string, cosAPIKey string, cosCRN string, bucketName string) string {
	return fmt.Sprintf(`

		resource "ibm_cd_toolchain_tool_cos" "cd_toolchain_tool_cos_instance" {
			toolchain_id = "%s"
			name = "%s"
			parameters {
				name = "cos_tool_01"
				cos_api_key = "%s"
				instance_crn = "%s"
				bucket_name = "%s"
				endpoint = "s3.direct.us-south.cloud-object-storage.appdomain.cloud"
				hmac_access_key_id = "hmac_access_key_id"
				hmac_secret_access_key = "hmac_secret_access_key"
			}
		}
	`, toolchainID, name, cosAPIKey, cosCRN, bucketName)
}

func testAccCheckIBMCdToolchainToolCosExists(n string, obj cdtoolchainv2.ToolchainTool) resource.TestCheckFunc {

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

func testAccCheckIBMCdToolchainToolCosDestroy(s *terraform.State) error {
	cdToolchainClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).CdToolchainV2()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_cd_toolchain_tool_cos" {
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
			return fmt.Errorf("cd_toolchain_tool_cos still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for cd_toolchain_tool_cos (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
