// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package cloudshell_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccIBMCloudShellAccountSettingsBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckCloudShell(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCloudShellAccountSettingsConfigBasic(acc.CloudShellAccountID),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_cloud_shell_account_settings.cloud_shell_account_settings", "account_id", acc.CloudShellAccountID),
				),
			},
		},
	})
}

func TestAccIBMCloudShellAccountSettingsAllArgs(t *testing.T) {
	defaultEnableNewFeatures := "false"
	defaultEnableNewRegions := "true"
	enabled := "false"
	featureWebPreview := "false"
	regionJpTok := "false"
	defaultEnableNewFeaturesUpdate := "true"
	defaultEnableNewRegionsUpdate := "false"
	enabledUpdate := "true"
	featureWebPreviewUpdate := "true"
	regionJpTokUpdate := "true"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckCloudShell(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCloudShellAccountSettingsConfig("1", acc.CloudShellAccountID, defaultEnableNewFeatures, defaultEnableNewRegions, enabled, featureWebPreview, regionJpTok),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_cloud_shell_account_settings.cloud_shell_account_settings1", "id", fmt.Sprintf("ac-%s", acc.CloudShellAccountID)),
					resource.TestCheckResourceAttr("ibm_cloud_shell_account_settings.cloud_shell_account_settings1", "account_id", acc.CloudShellAccountID),
					resource.TestCheckResourceAttrSet("ibm_cloud_shell_account_settings.cloud_shell_account_settings1", "rev"),
					resource.TestCheckResourceAttr("ibm_cloud_shell_account_settings.cloud_shell_account_settings1", "default_enable_new_features", defaultEnableNewFeatures),
					resource.TestCheckResourceAttr("ibm_cloud_shell_account_settings.cloud_shell_account_settings1", "default_enable_new_regions", defaultEnableNewRegions),
					resource.TestCheckResourceAttr("ibm_cloud_shell_account_settings.cloud_shell_account_settings1", "enabled", enabled),
					resource.TestCheckResourceAttrSet("ibm_cloud_shell_account_settings.cloud_shell_account_settings1", "features.#"),
					resource.TestCheckResourceAttr("ibm_cloud_shell_account_settings.cloud_shell_account_settings1", "features.1.enabled", featureWebPreview),
					resource.TestCheckResourceAttr("ibm_cloud_shell_account_settings.cloud_shell_account_settings1", "features.1.key", "server.web_preview"),
					resource.TestCheckResourceAttrSet("ibm_cloud_shell_account_settings.cloud_shell_account_settings1", "regions.#"),
					resource.TestCheckResourceAttr("ibm_cloud_shell_account_settings.cloud_shell_account_settings1", "regions.1.enabled", regionJpTok),
					resource.TestCheckResourceAttr("ibm_cloud_shell_account_settings.cloud_shell_account_settings1", "regions.1.key", "jp-tok"),
				),
			},
			{
				Config: testAccCheckIBMCloudShellAccountSettingsConfig("2", acc.CloudShellAccountID, defaultEnableNewFeaturesUpdate, defaultEnableNewRegionsUpdate, enabledUpdate, featureWebPreviewUpdate, regionJpTokUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_cloud_shell_account_settings.cloud_shell_account_settings2", "id", fmt.Sprintf("ac-%s", acc.CloudShellAccountID)),
					resource.TestCheckResourceAttr("ibm_cloud_shell_account_settings.cloud_shell_account_settings2", "account_id", acc.CloudShellAccountID),
					resource.TestCheckResourceAttrSet("ibm_cloud_shell_account_settings.cloud_shell_account_settings2", "rev"),
					resource.TestCheckResourceAttr("ibm_cloud_shell_account_settings.cloud_shell_account_settings2", "default_enable_new_features", defaultEnableNewFeaturesUpdate),
					resource.TestCheckResourceAttr("ibm_cloud_shell_account_settings.cloud_shell_account_settings2", "default_enable_new_regions", defaultEnableNewRegionsUpdate),
					resource.TestCheckResourceAttr("ibm_cloud_shell_account_settings.cloud_shell_account_settings2", "enabled", enabledUpdate),
					resource.TestCheckResourceAttrSet("ibm_cloud_shell_account_settings.cloud_shell_account_settings2", "features.#"),
					resource.TestCheckResourceAttr("ibm_cloud_shell_account_settings.cloud_shell_account_settings2", "features.1.enabled", featureWebPreviewUpdate),
					resource.TestCheckResourceAttr("ibm_cloud_shell_account_settings.cloud_shell_account_settings2", "features.1.key", "server.web_preview"),
					resource.TestCheckResourceAttrSet("ibm_cloud_shell_account_settings.cloud_shell_account_settings2", "regions.#"),
					resource.TestCheckResourceAttr("ibm_cloud_shell_account_settings.cloud_shell_account_settings2", "regions.1.enabled", regionJpTokUpdate),
					resource.TestCheckResourceAttr("ibm_cloud_shell_account_settings.cloud_shell_account_settings2", "regions.1.key", "jp-tok"),
				),
			},
		},
	})
}

func testAccCheckIBMCloudShellAccountSettingsConfigBasic(accountID string) string {
	return fmt.Sprintf(`
	data "ibm_cloud_shell_account_settings" "account_settings" {
		account_id = "%s"
	}

	resource "ibm_cloud_shell_account_settings" "cloud_shell_account_settings" {
		account_id = "%s"
		rev = data.ibm_cloud_shell_account_settings.account_settings.rev
		default_enable_new_features = true
		default_enable_new_regions = true
		enabled = true
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
			enabled = true
			key = "jp-tok"
		}
		regions {
			enabled = false
			key = "us-south"
		}
	}
	`, accountID, accountID)
}

func testAccCheckIBMCloudShellAccountSettingsConfig(suffix, accountID string, defaultEnableNewFeatures string, defaultEnableNewRegions string, enabled string, featureWebPreview string, regionJpTok string) string {
	return fmt.Sprintf(`
	data "ibm_cloud_shell_account_settings" "account_settings%s" {
		account_id = "%s"
	}

	resource "ibm_cloud_shell_account_settings" "cloud_shell_account_settings%s" {
		account_id = "%s"
		rev = data.ibm_cloud_shell_account_settings.account_settings%s.rev
		default_enable_new_features = %s
		default_enable_new_regions = %s
		enabled = %s
		features {
			enabled = true
			key = "server.file_manager"
		}
		features {
			enabled = %s
			key = "server.web_preview"
		}
		regions {
			enabled = true
			key = "eu-de"
		}
		regions {
			enabled = %s
			key = "jp-tok"
		}
		regions {
			enabled = false
			key = "us-south"
		}
	}
	`, suffix, accountID, suffix, accountID, suffix, defaultEnableNewFeatures, defaultEnableNewRegions, enabled, featureWebPreview, regionJpTok)
}
