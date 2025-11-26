// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.98.0-8be2046a-20241205-162752
 */

package cloudshell_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIBMCloudShellAccountSettingsDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckCloudShell(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCloudShellAccountSettingsDataSourceConfigBasic(acc.CloudShellAccountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_cloud_shell_account_settings.cloud_shell_account_settings_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_cloud_shell_account_settings.cloud_shell_account_settings_instance", "account_id"),
				),
			},
		},
	})
}

func TestAccIBMCloudShellAccountSettingsDataSourceAllArgs(t *testing.T) {
	accountSettingsDefaultEnableNewFeatures := "false"
	accountSettingsDefaultEnableNewRegions := "true"
	accountSettingsEnabled := "true"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckCloudShell(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCloudShellAccountSettingsDataSourceConfig(acc.CloudShellAccountID, accountSettingsDefaultEnableNewFeatures, accountSettingsDefaultEnableNewRegions, accountSettingsEnabled),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ibm_cloud_shell_account_settings.account_settings_after_update", "id", acc.CloudShellAccountID),
					resource.TestCheckResourceAttr("data.ibm_cloud_shell_account_settings.account_settings_after_update", "account_id", acc.CloudShellAccountID),
					resource.TestCheckResourceAttrSet("data.ibm_cloud_shell_account_settings.account_settings_after_update", "rev"),
					resource.TestCheckResourceAttrSet("data.ibm_cloud_shell_account_settings.account_settings_after_update", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_cloud_shell_account_settings.account_settings_after_update", "created_by"),
					resource.TestCheckResourceAttr("data.ibm_cloud_shell_account_settings.account_settings_after_update", "default_enable_new_features", accountSettingsDefaultEnableNewFeatures),
					resource.TestCheckResourceAttr("data.ibm_cloud_shell_account_settings.account_settings_after_update", "default_enable_new_regions", accountSettingsDefaultEnableNewRegions),
					resource.TestCheckResourceAttr("data.ibm_cloud_shell_account_settings.account_settings_after_update", "enabled", accountSettingsEnabled),
					resource.TestCheckResourceAttrSet("data.ibm_cloud_shell_account_settings.account_settings_after_update", "features.#"),
					resource.TestCheckResourceAttr("data.ibm_cloud_shell_account_settings.account_settings_after_update", "features.0.enabled", "true"),
					resource.TestCheckResourceAttr("data.ibm_cloud_shell_account_settings.account_settings_after_update", "features.0.key", "server.file_manager"),
					resource.TestCheckResourceAttr("data.ibm_cloud_shell_account_settings.account_settings_after_update", "features.1.enabled", "false"),
					resource.TestCheckResourceAttr("data.ibm_cloud_shell_account_settings.account_settings_after_update", "features.1.key", "server.web_preview"),
					resource.TestCheckResourceAttrSet("data.ibm_cloud_shell_account_settings.account_settings_after_update", "regions.#"),
					resource.TestCheckResourceAttr("data.ibm_cloud_shell_account_settings.account_settings_after_update", "regions.0.enabled", "true"),
					resource.TestCheckResourceAttr("data.ibm_cloud_shell_account_settings.account_settings_after_update", "regions.0.key", "eu-de"),
					resource.TestCheckResourceAttr("data.ibm_cloud_shell_account_settings.account_settings_after_update", "regions.1.enabled", "false"),
					resource.TestCheckResourceAttr("data.ibm_cloud_shell_account_settings.account_settings_after_update", "regions.1.key", "us-south"),
					resource.TestCheckResourceAttrSet("data.ibm_cloud_shell_account_settings.account_settings_after_update", "type"),
				),
			},
		},
	})
}

func testAccCheckIBMCloudShellAccountSettingsDataSourceConfigBasic(accountSettingsAccountID string) string {
	return fmt.Sprintf(`
		 data "ibm_cloud_shell_account_settings" "cloud_shell_account_settings_instance" {
			 account_id = "%s"
		 }
	 `, accountSettingsAccountID)
}

func testAccCheckIBMCloudShellAccountSettingsDataSourceConfig(accountSettingsAccountID string, accountSettingsDefaultEnableNewFeatures string, accountSettingsDefaultEnableNewRegions string, accountSettingsEnabled string) string {
	// first need to retrieve the existing account settings revision field (rev) before updating account settings.
	return fmt.Sprintf(`
		 data "ibm_cloud_shell_account_settings" "account_settings_before_update" {
			 account_id = "%s"
		 }
 
		 resource "ibm_cloud_shell_account_settings" "cloud_shell_account_settings" {
			 account_id = "%s"
			 rev = data.ibm_cloud_shell_account_settings.account_settings_before_update.rev
			 default_enable_new_features = %s
			 default_enable_new_regions = %s
			 enabled = %s
			 features {
				 enabled = true
				 key = "server.file_manager"
			 }
			 features {
				 enabled = false
				 key = "server.web_preview"
			 }
			 regions {
				 enabled = true
				 key = "eu-de"
			 }
			 regions {
				 enabled = false
				 key = "us-south"
			 }
		 }
 
		 data "ibm_cloud_shell_account_settings" "account_settings_after_update" {
			 account_id = ibm_cloud_shell_account_settings.cloud_shell_account_settings.account_id
		 }
	 `, accountSettingsAccountID, accountSettingsAccountID, accountSettingsDefaultEnableNewFeatures, accountSettingsDefaultEnableNewRegions, accountSettingsEnabled)
}
