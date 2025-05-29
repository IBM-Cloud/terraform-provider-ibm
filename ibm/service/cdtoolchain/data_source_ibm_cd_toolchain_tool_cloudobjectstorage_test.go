// Copyright IBM Corp. 2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.96.0-d6dec9d7-20241008-212902
 */

package cdtoolchain_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIBMCdToolchainToolCloudobjectstorageDataSourceBasic(t *testing.T) {
	toolchainToolToolchainID := acc.ToolchainID

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCdToolchainToolCloudobjectstorageDataSourceConfigBasic(toolchainToolToolchainID, acc.COSApiKey, acc.CosCRN, acc.BucketName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_cloudobjectstorage.cd_toolchain_tool_cloudobjectstorage_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_cloudobjectstorage.cd_toolchain_tool_cloudobjectstorage_instance", "toolchain_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_cloudobjectstorage.cd_toolchain_tool_cloudobjectstorage_instance", "tool_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_cloudobjectstorage.cd_toolchain_tool_cloudobjectstorage_instance", "resource_group_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_cloudobjectstorage.cd_toolchain_tool_cloudobjectstorage_instance", "crn"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_cloudobjectstorage.cd_toolchain_tool_cloudobjectstorage_instance", "toolchain_crn"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_cloudobjectstorage.cd_toolchain_tool_cloudobjectstorage_instance", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_cloudobjectstorage.cd_toolchain_tool_cloudobjectstorage_instance", "referent.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_cloudobjectstorage.cd_toolchain_tool_cloudobjectstorage_instance", "updated_at"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_cloudobjectstorage.cd_toolchain_tool_cloudobjectstorage_instance", "parameters.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_cloudobjectstorage.cd_toolchain_tool_cloudobjectstorage_instance", "state"),
				),
			},
		},
	})
}

func TestAccIBMCdToolchainToolCloudobjectstorageDataSourceAllArgs(t *testing.T) {
	toolchainToolToolchainID := acc.ToolchainID
	toolchainToolName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCdToolchainToolCloudobjectstorageDataSourceConfig(toolchainToolToolchainID, toolchainToolName, acc.COSApiKey, acc.CosCRN, acc.BucketName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_cloudobjectstorage.cd_toolchain_tool_cloudobjectstorage_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_cloudobjectstorage.cd_toolchain_tool_cloudobjectstorage_instance", "toolchain_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_cloudobjectstorage.cd_toolchain_tool_cloudobjectstorage_instance", "tool_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_cloudobjectstorage.cd_toolchain_tool_cloudobjectstorage_instance", "resource_group_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_cloudobjectstorage.cd_toolchain_tool_cloudobjectstorage_instance", "crn"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_cloudobjectstorage.cd_toolchain_tool_cloudobjectstorage_instance", "toolchain_crn"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_cloudobjectstorage.cd_toolchain_tool_cloudobjectstorage_instance", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_cloudobjectstorage.cd_toolchain_tool_cloudobjectstorage_instance", "referent.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_cloudobjectstorage.cd_toolchain_tool_cloudobjectstorage_instance", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_cloudobjectstorage.cd_toolchain_tool_cloudobjectstorage_instance", "updated_at"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_cloudobjectstorage.cd_toolchain_tool_cloudobjectstorage_instance", "parameters.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_cloudobjectstorage.cd_toolchain_tool_cloudobjectstorage_instance", "state"),
				),
			},
		},
	})
}

func testAccCheckIBMCdToolchainToolCloudobjectstorageDataSourceConfigBasic(toolchainToolToolchainID string, cosAPIKey string, cosCRN string, bucketName string) string {
	return fmt.Sprintf(`
		resource "ibm_cd_toolchain_tool_cloudobjectstorage" "cd_toolchain_tool_cloudobjectstorage_instance" {
			toolchain_id = "%s"
			parameters {
				name = "cos_tool_01"
				auth_type = "apikey"
				cos_api_key = "%s"
				instance_crn = "%s"
				bucket_name = "%s"
				endpoint = "s3.direct.us-south.cloud-object-storage.appdomain.cloud"
				hmac_access_key_id = "hmac_access_key_id"
				hmac_secret_access_key = "hmac_secret_access_key"
			}
		}

		data "ibm_cd_toolchain_tool_cloudobjectstorage" "cd_toolchain_tool_cloudobjectstorage_instance" {
			toolchain_id = ibm_cd_toolchain_tool_cloudobjectstorage.cd_toolchain_tool_cloudobjectstorage_instance.toolchain_id
			tool_id = ibm_cd_toolchain_tool_cloudobjectstorage.cd_toolchain_tool_cloudobjectstorage_instance.tool_id
		}
	`, toolchainToolToolchainID, cosAPIKey, cosCRN, bucketName)
}

func testAccCheckIBMCdToolchainToolCloudobjectstorageDataSourceConfig(toolchainToolToolchainID string, toolchainToolName string, cosAPIKey string, cosCRN string, bucketName string) string {
	return fmt.Sprintf(`
		resource "ibm_cd_toolchain_tool_cloudobjectstorage" "cd_toolchain_tool_cloudobjectstorage_instance" {
			toolchain_id = "%s"
			name = "%s"
			parameters {
				name = "cos_tool_01"
				auth_type = "apikey"
				cos_api_key = "%s"
				instance_crn = "%s"
				bucket_name = "%s"
				endpoint = "s3.direct.us-south.cloud-object-storage.appdomain.cloud"
				hmac_access_key_id = "hmac_access_key_id"
				hmac_secret_access_key = "hmac_secret_access_key"
			}
		}

		data "ibm_cd_toolchain_tool_cloudobjectstorage" "cd_toolchain_tool_cloudobjectstorage_instance" {
			toolchain_id = ibm_cd_toolchain_tool_cloudobjectstorage.cd_toolchain_tool_cloudobjectstorage_instance.toolchain_id
			tool_id = ibm_cd_toolchain_tool_cloudobjectstorage.cd_toolchain_tool_cloudobjectstorage_instance.tool_id
		}
	`, toolchainToolToolchainID, toolchainToolName, cosAPIKey, cosCRN, bucketName)
}
