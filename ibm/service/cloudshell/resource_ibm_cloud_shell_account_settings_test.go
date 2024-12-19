// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package cloudshell_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/cloudshell"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/platform-services-go-sdk/ibmcloudshellv1"
	"github.com/stretchr/testify/assert"
)

func TestAccIBMCloudShellAccountSettingsBasic(t *testing.T) {
	var conf ibmcloudshellv1.AccountSettings
	accountID := fmt.Sprintf("tf_account_id_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCloudShellAccountSettingsDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCloudShellAccountSettingsConfigBasic(accountID),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCloudShellAccountSettingsExists("ibm_cloud_shell_account_settings.cloud_shell_account_settings_instance", conf),
					resource.TestCheckResourceAttr("ibm_cloud_shell_account_settings.cloud_shell_account_settings_instance", "account_id", accountID),
				),
			},
		},
	})
}

func TestAccIBMCloudShellAccountSettingsAllArgs(t *testing.T) {
	var conf ibmcloudshellv1.AccountSettings
	accountID := fmt.Sprintf("tf_account_id_%d", acctest.RandIntRange(10, 100))
	rev := fmt.Sprintf("tf_rev_%d", acctest.RandIntRange(10, 100))
	defaultEnableNewFeatures := "false"
	defaultEnableNewRegions := "true"
	enabled := "false"
	revUpdate := fmt.Sprintf("tf_rev_%d", acctest.RandIntRange(10, 100))
	defaultEnableNewFeaturesUpdate := "true"
	defaultEnableNewRegionsUpdate := "false"
	enabledUpdate := "true"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCloudShellAccountSettingsDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCloudShellAccountSettingsConfig(accountID, rev, defaultEnableNewFeatures, defaultEnableNewRegions, enabled),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCloudShellAccountSettingsExists("ibm_cloud_shell_account_settings.cloud_shell_account_settings_instance", conf),
					resource.TestCheckResourceAttr("ibm_cloud_shell_account_settings.cloud_shell_account_settings_instance", "account_id", accountID),
					resource.TestCheckResourceAttr("ibm_cloud_shell_account_settings.cloud_shell_account_settings_instance", "rev", rev),
					resource.TestCheckResourceAttr("ibm_cloud_shell_account_settings.cloud_shell_account_settings_instance", "default_enable_new_features", defaultEnableNewFeatures),
					resource.TestCheckResourceAttr("ibm_cloud_shell_account_settings.cloud_shell_account_settings_instance", "default_enable_new_regions", defaultEnableNewRegions),
					resource.TestCheckResourceAttr("ibm_cloud_shell_account_settings.cloud_shell_account_settings_instance", "enabled", enabled),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMCloudShellAccountSettingsConfig(accountID, revUpdate, defaultEnableNewFeaturesUpdate, defaultEnableNewRegionsUpdate, enabledUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_cloud_shell_account_settings.cloud_shell_account_settings_instance", "account_id", accountID),
					resource.TestCheckResourceAttr("ibm_cloud_shell_account_settings.cloud_shell_account_settings_instance", "rev", revUpdate),
					resource.TestCheckResourceAttr("ibm_cloud_shell_account_settings.cloud_shell_account_settings_instance", "default_enable_new_features", defaultEnableNewFeaturesUpdate),
					resource.TestCheckResourceAttr("ibm_cloud_shell_account_settings.cloud_shell_account_settings_instance", "default_enable_new_regions", defaultEnableNewRegionsUpdate),
					resource.TestCheckResourceAttr("ibm_cloud_shell_account_settings.cloud_shell_account_settings_instance", "enabled", enabledUpdate),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_cloud_shell_account_settings.cloud_shell_account_settings",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIBMCloudShellAccountSettingsConfigBasic(accountID string) string {
	return fmt.Sprintf(`
		resource "ibm_cloud_shell_account_settings" "cloud_shell_account_settings_instance" {
			account_id = "%s"
		}
	`, accountID)
}

func testAccCheckIBMCloudShellAccountSettingsConfig(accountID string, rev string, defaultEnableNewFeatures string, defaultEnableNewRegions string, enabled string) string {
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
	`, accountID, rev, defaultEnableNewFeatures, defaultEnableNewRegions, enabled)
}

func testAccCheckIBMCloudShellAccountSettingsExists(n string, obj ibmcloudshellv1.AccountSettings) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		ibmCloudShellClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).IBMCloudShellV1()
		if err != nil {
			return err
		}

		getAccountSettingsOptions := &ibmcloudshellv1.GetAccountSettingsOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getAccountSettingsOptions.SetAccountID(parts[0])
		getAccountSettingsOptions.SetAccountID(parts[1])

		accountSettings, _, err := ibmCloudShellClient.GetAccountSettings(getAccountSettingsOptions)
		if err != nil {
			return err
		}

		obj = *accountSettings
		return nil
	}
}

func testAccCheckIBMCloudShellAccountSettingsDestroy(s *terraform.State) error {
	ibmCloudShellClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).IBMCloudShellV1()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_cloud_shell_account_settings" {
			continue
		}

		getAccountSettingsOptions := &ibmcloudshellv1.GetAccountSettingsOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getAccountSettingsOptions.SetAccountID(parts[0])
		getAccountSettingsOptions.SetAccountID(parts[1])

		// Try to find the key
		_, response, err := ibmCloudShellClient.GetAccountSettings(getAccountSettingsOptions)

		if err == nil {
			return fmt.Errorf("cloud_shell_account_settings still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for cloud_shell_account_settings (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}

func TestResourceIBMCloudShellAccountSettingsFeatureToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["enabled"] = true
		model["key"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(ibmcloudshellv1.Feature)
	model.Enabled = core.BoolPtr(true)
	model.Key = core.StringPtr("testString")

	result, err := cloudshell.ResourceIBMCloudShellAccountSettingsFeatureToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMCloudShellAccountSettingsRegionSettingToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["enabled"] = true
		model["key"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(ibmcloudshellv1.RegionSetting)
	model.Enabled = core.BoolPtr(true)
	model.Key = core.StringPtr("testString")

	result, err := cloudshell.ResourceIBMCloudShellAccountSettingsRegionSettingToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMCloudShellAccountSettingsMapToFeature(t *testing.T) {
	checkResult := func(result *ibmcloudshellv1.Feature) {
		model := new(ibmcloudshellv1.Feature)
		model.Enabled = core.BoolPtr(true)
		model.Key = core.StringPtr("testString")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["enabled"] = true
	model["key"] = "testString"

	result, err := cloudshell.ResourceIBMCloudShellAccountSettingsMapToFeature(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMCloudShellAccountSettingsMapToRegionSetting(t *testing.T) {
	checkResult := func(result *ibmcloudshellv1.RegionSetting) {
		model := new(ibmcloudshellv1.RegionSetting)
		model.Enabled = core.BoolPtr(true)
		model.Key = core.StringPtr("testString")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["enabled"] = true
	model["key"] = "testString"

	result, err := cloudshell.ResourceIBMCloudShellAccountSettingsMapToRegionSetting(model)
	assert.Nil(t, err)
	checkResult(result)
}
