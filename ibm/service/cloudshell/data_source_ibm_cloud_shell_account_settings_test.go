// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.98.0-8be2046a-20241205-162752
 */

package cloudshell_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/cloudshell"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/platform-services-go-sdk/ibmcloudshellv1"
	"github.com/stretchr/testify/assert"
)

func TestAccIBMCloudShellAccountSettingsDataSourceBasic(t *testing.T) {
	accountSettingsAccountID := fmt.Sprintf("tf_account_id_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCloudShellAccountSettingsDataSourceConfigBasic(accountSettingsAccountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_cloud_shell_account_settings.cloud_shell_account_settings_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_cloud_shell_account_settings.cloud_shell_account_settings_instance", "account_id"),
				),
			},
		},
	})
}

func TestAccIBMCloudShellAccountSettingsDataSourceAllArgs(t *testing.T) {
	accountSettingsAccountID := fmt.Sprintf("tf_account_id_%d", acctest.RandIntRange(10, 100))
	accountSettingsRev := fmt.Sprintf("tf_rev_%d", acctest.RandIntRange(10, 100))
	accountSettingsDefaultEnableNewFeatures := "false"
	accountSettingsDefaultEnableNewRegions := "true"
	accountSettingsEnabled := "false"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCloudShellAccountSettingsDataSourceConfig(accountSettingsAccountID, accountSettingsRev, accountSettingsDefaultEnableNewFeatures, accountSettingsDefaultEnableNewRegions, accountSettingsEnabled),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_cloud_shell_account_settings.cloud_shell_account_settings_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_cloud_shell_account_settings.cloud_shell_account_settings_instance", "account_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cloud_shell_account_settings.cloud_shell_account_settings_instance", "rev"),
					resource.TestCheckResourceAttrSet("data.ibm_cloud_shell_account_settings.cloud_shell_account_settings_instance", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_cloud_shell_account_settings.cloud_shell_account_settings_instance", "created_by"),
					resource.TestCheckResourceAttrSet("data.ibm_cloud_shell_account_settings.cloud_shell_account_settings_instance", "default_enable_new_features"),
					resource.TestCheckResourceAttrSet("data.ibm_cloud_shell_account_settings.cloud_shell_account_settings_instance", "default_enable_new_regions"),
					resource.TestCheckResourceAttrSet("data.ibm_cloud_shell_account_settings.cloud_shell_account_settings_instance", "enabled"),
					resource.TestCheckResourceAttrSet("data.ibm_cloud_shell_account_settings.cloud_shell_account_settings_instance", "features.#"),
					resource.TestCheckResourceAttr("data.ibm_cloud_shell_account_settings.cloud_shell_account_settings_instance", "features.0.enabled", accountSettingsEnabled),
					resource.TestCheckResourceAttrSet("data.ibm_cloud_shell_account_settings.cloud_shell_account_settings_instance", "features.0.key"),
					resource.TestCheckResourceAttrSet("data.ibm_cloud_shell_account_settings.cloud_shell_account_settings_instance", "regions.#"),
					resource.TestCheckResourceAttr("data.ibm_cloud_shell_account_settings.cloud_shell_account_settings_instance", "regions.0.enabled", accountSettingsEnabled),
					resource.TestCheckResourceAttrSet("data.ibm_cloud_shell_account_settings.cloud_shell_account_settings_instance", "regions.0.key"),
					resource.TestCheckResourceAttrSet("data.ibm_cloud_shell_account_settings.cloud_shell_account_settings_instance", "type"),
					resource.TestCheckResourceAttrSet("data.ibm_cloud_shell_account_settings.cloud_shell_account_settings_instance", "updated_at"),
					resource.TestCheckResourceAttrSet("data.ibm_cloud_shell_account_settings.cloud_shell_account_settings_instance", "updated_by"),
				),
			},
		},
	})
}

func testAccCheckIBMCloudShellAccountSettingsDataSourceConfigBasic(accountSettingsAccountID string) string {
	return fmt.Sprintf(`
		resource "ibm_cloud_shell_account_settings" "cloud_shell_account_settings_instance" {
			account_id = "%s"
		}

		data "ibm_cloud_shell_account_settings" "cloud_shell_account_settings_instance" {
			account_id = ibm_cloud_shell_account_settings.cloud_shell_account_settings_instance.account_id
		}
	`, accountSettingsAccountID)
}

func testAccCheckIBMCloudShellAccountSettingsDataSourceConfig(accountSettingsAccountID string, accountSettingsRev string, accountSettingsDefaultEnableNewFeatures string, accountSettingsDefaultEnableNewRegions string, accountSettingsEnabled string) string {
	return fmt.Sprintf(`
		resource "ibm_cloud_shell_account_settings" "cloud_shell_account_settings_instance" {
			account_id = "%s"
			rev = "%s"
			default_enable_new_features = %s
			default_enable_new_regions = %s
			enabled = %s
			features {
				enabled = true
				key = "key"
			}
			regions {
				enabled = true
				key = "key"
			}
		}

		data "ibm_cloud_shell_account_settings" "cloud_shell_account_settings_instance" {
			account_id = ibm_cloud_shell_account_settings.cloud_shell_account_settings_instance.account_id
		}
	`, accountSettingsAccountID, accountSettingsRev, accountSettingsDefaultEnableNewFeatures, accountSettingsDefaultEnableNewRegions, accountSettingsEnabled)
}

func TestDataSourceIBMCloudShellAccountSettingsFeatureToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["enabled"] = true
		model["key"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(ibmcloudshellv1.Feature)
	model.Enabled = core.BoolPtr(true)
	model.Key = core.StringPtr("testString")

	result, err := cloudshell.DataSourceIBMCloudShellAccountSettingsFeatureToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMCloudShellAccountSettingsRegionSettingToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["enabled"] = true
		model["key"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(ibmcloudshellv1.RegionSetting)
	model.Enabled = core.BoolPtr(true)
	model.Key = core.StringPtr("testString")

	result, err := cloudshell.DataSourceIBMCloudShellAccountSettingsRegionSettingToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}
