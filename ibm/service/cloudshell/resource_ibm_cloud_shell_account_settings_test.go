// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package cloudshell_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIBMCloudShellAccountSettingsBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckCloudShell(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
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
	regionUsSouth := "false"

	defaultEnableNewFeaturesUpdate := "true"
	defaultEnableNewRegionsUpdate := "false"
	enabledUpdate := "true"
	featureWebPreviewUpdate := "true"
	regionUsSouthUpdate := "true"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckCloudShell(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCloudShellAccountSettingsConfig("1", acc.CloudShellAccountID, defaultEnableNewFeatures, defaultEnableNewRegions, enabled, featureWebPreview, regionUsSouth),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_cloud_shell_account_settings.cloud_shell_account_settings1", "id", acc.CloudShellAccountID),
					resource.TestCheckResourceAttr("ibm_cloud_shell_account_settings.cloud_shell_account_settings1", "account_id", acc.CloudShellAccountID),
					resource.TestCheckResourceAttrSet("ibm_cloud_shell_account_settings.cloud_shell_account_settings1", "rev"),
					resource.TestCheckResourceAttr("ibm_cloud_shell_account_settings.cloud_shell_account_settings1", "default_enable_new_features", defaultEnableNewFeatures),
					resource.TestCheckResourceAttr("ibm_cloud_shell_account_settings.cloud_shell_account_settings1", "default_enable_new_regions", defaultEnableNewRegions),
					resource.TestCheckResourceAttr("ibm_cloud_shell_account_settings.cloud_shell_account_settings1", "enabled", enabled),
					resource.TestCheckResourceAttrSet("ibm_cloud_shell_account_settings.cloud_shell_account_settings1", "features.#"),
					resource.TestCheckResourceAttr("ibm_cloud_shell_account_settings.cloud_shell_account_settings1", "features.1.enabled", featureWebPreview),
					resource.TestCheckResourceAttr("ibm_cloud_shell_account_settings.cloud_shell_account_settings1", "features.1.key", "server.web_preview"),
					resource.TestCheckResourceAttrSet("ibm_cloud_shell_account_settings.cloud_shell_account_settings1", "regions.#"),
					resource.TestCheckResourceAttr("ibm_cloud_shell_account_settings.cloud_shell_account_settings1", "regions.1.enabled", regionUsSouth),
					resource.TestCheckResourceAttr("ibm_cloud_shell_account_settings.cloud_shell_account_settings1", "regions.1.key", "us-south"),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMCloudShellAccountSettingsConfig("2", acc.CloudShellAccountID, defaultEnableNewFeaturesUpdate, defaultEnableNewRegionsUpdate, enabledUpdate, featureWebPreviewUpdate, regionUsSouthUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_cloud_shell_account_settings.cloud_shell_account_settings2", "id", acc.CloudShellAccountID),
					resource.TestCheckResourceAttr("ibm_cloud_shell_account_settings.cloud_shell_account_settings2", "account_id", acc.CloudShellAccountID),
					resource.TestCheckResourceAttrSet("ibm_cloud_shell_account_settings.cloud_shell_account_settings2", "rev"),
					resource.TestCheckResourceAttr("ibm_cloud_shell_account_settings.cloud_shell_account_settings2", "default_enable_new_features", defaultEnableNewFeaturesUpdate),
					resource.TestCheckResourceAttr("ibm_cloud_shell_account_settings.cloud_shell_account_settings2", "default_enable_new_regions", defaultEnableNewRegionsUpdate),
					resource.TestCheckResourceAttr("ibm_cloud_shell_account_settings.cloud_shell_account_settings2", "enabled", enabledUpdate),
					resource.TestCheckResourceAttrSet("ibm_cloud_shell_account_settings.cloud_shell_account_settings2", "features.#"),
					resource.TestCheckResourceAttr("ibm_cloud_shell_account_settings.cloud_shell_account_settings2", "features.1.enabled", featureWebPreviewUpdate),
					resource.TestCheckResourceAttr("ibm_cloud_shell_account_settings.cloud_shell_account_settings2", "features.1.key", "server.web_preview"),
					resource.TestCheckResourceAttrSet("ibm_cloud_shell_account_settings.cloud_shell_account_settings2", "regions.#"),
					resource.TestCheckResourceAttr("ibm_cloud_shell_account_settings.cloud_shell_account_settings2", "regions.1.enabled", regionUsSouthUpdate),
					resource.TestCheckResourceAttr("ibm_cloud_shell_account_settings.cloud_shell_account_settings2", "regions.1.key", "us-south"),
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
			key = "us-south"
		}
	}
	`, accountID, accountID)
}

func testAccCheckIBMCloudShellAccountSettingsConfig(suffix, accountID string, defaultEnableNewFeatures string, defaultEnableNewRegions string, enabled string, featureWebPreview string, regionUsSouth string) string {
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
				key = "us-south"
			}
	}
	`, suffix, accountID, suffix, accountID, suffix, defaultEnableNewFeatures, defaultEnableNewRegions, enabled, featureWebPreview, regionUsSouth)
}
