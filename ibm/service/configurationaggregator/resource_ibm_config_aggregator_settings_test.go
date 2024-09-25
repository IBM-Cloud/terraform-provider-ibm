// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package configurationaggregator_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/IBM/configuration-aggregator-go-sdk/configurationaggregatorv1"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/configurationaggregator"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/stretchr/testify/assert"
)

func TestAccIbmConfigAggregatorSettingsBasic(t *testing.T) {
	var conf configurationaggregatorv1.SettingsResponse

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmConfigAggregatorSettingsDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmConfigAggregatorSettingsConfigBasic(),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmConfigAggregatorSettingsExists("ibm_config_aggregator_settings.config_aggregator_settings_instance", conf),
				),
			},
		},
	})
}

func TestAccIbmConfigAggregatorSettingsAllArgs(t *testing.T) {
	var conf configurationaggregatorv1.SettingsResponse
	resourceCollectionEnabled := "false"
	trustedProfileID := fmt.Sprintf("tf_trusted_profile_id_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmConfigAggregatorSettingsDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmConfigAggregatorSettingsConfig(resourceCollectionEnabled, trustedProfileID),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmConfigAggregatorSettingsExists("ibm_config_aggregator_settings.config_aggregator_settings_instance", conf),
					resource.TestCheckResourceAttr("ibm_config_aggregator_settings.config_aggregator_settings_instance", "resource_collection_enabled", resourceCollectionEnabled),
					resource.TestCheckResourceAttr("ibm_config_aggregator_settings.config_aggregator_settings_instance", "trusted_profile_id", trustedProfileID),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_config_aggregator_settings.config_aggregator_settings",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIbmConfigAggregatorSettingsConfigBasic() string {
	return fmt.Sprintf(`
		resource "ibm_config_aggregator_settings" "config_aggregator_settings_instance" {
		}
	`)
}

func testAccCheckIbmConfigAggregatorSettingsConfig(resourceCollectionEnabled string, trustedProfileID string) string {
	return fmt.Sprintf(`

		resource "ibm_config_aggregator_settings" "config_aggregator_settings_instance" {
			resource_collection_enabled = %s
			trusted_profile_id = "%s"
			regions = "FIXME"
			additional_scope {
				type = "Enterprise"
				enterprise_id = "enterprise_id"
				profile_template {
					id = "ProfileTemplate-adb55769-ae22-4c60-aead-bd1f84f93c57"
					trusted_profile_id = "Profile-6bb60124-8fc3-4d18-b63d-0b99560865d3"
				}
			}
		}
	`, resourceCollectionEnabled, trustedProfileID)
}

func testAccCheckIbmConfigAggregatorSettingsExists(n string, obj configurationaggregatorv1.SettingsResponse) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		_, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		configurationAggregatorClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).ConfigurationAggregatorV1()
		if err != nil {
			return err
		}

		getSettingsOptions := &configurationaggregatorv1.GetSettingsOptions{}

		updateSettings, _, err := configurationAggregatorClient.GetSettings(getSettingsOptions)
		if err != nil {
			return err
		}

		obj = *updateSettings
		return nil
	}
}

func testAccCheckIbmConfigAggregatorSettingsDestroy(s *terraform.State) error {
	configurationAggregatorClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).ConfigurationAggregatorV1()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_config_aggregator_settings" {
			continue
		}

		getSettingsOptions := &configurationaggregatorv1.GetSettingsOptions{}

		// Try to find the key
		_, response, err := configurationAggregatorClient.GetSettings(getSettingsOptions)

		if err == nil {
			return fmt.Errorf("config_aggregator_settings still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for config_aggregator_settings (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}

func TestResourceIbmConfigAggregatorSettingsAdditionalScopeToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		profileTemplateModel := make(map[string]interface{})
		profileTemplateModel["id"] = "ProfileTemplate-adb55769-ae22-4c60-aead-bd1f84f93c57"
		profileTemplateModel["trusted_profile_id"] = "Profile-6bb60124-8fc3-4d18-b63d-0b99560865d3"

		model := make(map[string]interface{})
		model["type"] = "Enterprise"
		model["enterprise_id"] = "testString"
		model["profile_template"] = []map[string]interface{}{profileTemplateModel}

		assert.Equal(t, result, model)
	}

	profileTemplateModel := new(configurationaggregatorv1.ProfileTemplate)
	profileTemplateModel.ID = core.StringPtr("ProfileTemplate-adb55769-ae22-4c60-aead-bd1f84f93c57")
	profileTemplateModel.TrustedProfileID = core.StringPtr("Profile-6bb60124-8fc3-4d18-b63d-0b99560865d3")

	model := new(configurationaggregatorv1.AdditionalScope)
	model.Type = core.StringPtr("Enterprise")
	model.EnterpriseID = core.StringPtr("testString")
	model.ProfileTemplate = profileTemplateModel

	result, err := configurationaggregator.ResourceIbmConfigAggregatorSettingsAdditionalScopeToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmConfigAggregatorSettingsProfileTemplateToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["id"] = "ProfileTemplate-adb55769-ae22-4c60-aead-bd1f84f93c57"
		model["trusted_profile_id"] = "Profile-6bb60124-8fc3-4d18-b63d-0b99560865d3"

		assert.Equal(t, result, model)
	}

	model := new(configurationaggregatorv1.ProfileTemplate)
	model.ID = core.StringPtr("ProfileTemplate-adb55769-ae22-4c60-aead-bd1f84f93c57")
	model.TrustedProfileID = core.StringPtr("Profile-6bb60124-8fc3-4d18-b63d-0b99560865d3")

	result, err := configurationaggregator.ResourceIbmConfigAggregatorSettingsProfileTemplateToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmConfigAggregatorSettingsMapToAdditionalScope(t *testing.T) {
	checkResult := func(result *configurationaggregatorv1.AdditionalScope) {
		profileTemplateModel := new(configurationaggregatorv1.ProfileTemplate)
		profileTemplateModel.ID = core.StringPtr("ProfileTemplate-adb55769-ae22-4c60-aead-bd1f84f93c57")
		profileTemplateModel.TrustedProfileID = core.StringPtr("Profile-6bb60124-8fc3-4d18-b63d-0b99560865d3")

		model := new(configurationaggregatorv1.AdditionalScope)
		model.Type = core.StringPtr("Enterprise")
		model.EnterpriseID = core.StringPtr("testString")
		model.ProfileTemplate = profileTemplateModel

		assert.Equal(t, result, model)
	}

	profileTemplateModel := make(map[string]interface{})
	profileTemplateModel["id"] = "ProfileTemplate-adb55769-ae22-4c60-aead-bd1f84f93c57"
	profileTemplateModel["trusted_profile_id"] = "Profile-6bb60124-8fc3-4d18-b63d-0b99560865d3"

	model := make(map[string]interface{})
	model["type"] = "Enterprise"
	model["enterprise_id"] = "testString"
	model["profile_template"] = []interface{}{profileTemplateModel}

	result, err := configurationaggregator.ResourceIbmConfigAggregatorSettingsMapToAdditionalScope(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmConfigAggregatorSettingsMapToProfileTemplate(t *testing.T) {
	checkResult := func(result *configurationaggregatorv1.ProfileTemplate) {
		model := new(configurationaggregatorv1.ProfileTemplate)
		model.ID = core.StringPtr("ProfileTemplate-adb55769-ae22-4c60-aead-bd1f84f93c57")
		model.TrustedProfileID = core.StringPtr("Profile-6bb60124-8fc3-4d18-b63d-0b99560865d3")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["id"] = "ProfileTemplate-adb55769-ae22-4c60-aead-bd1f84f93c57"
	model["trusted_profile_id"] = "Profile-6bb60124-8fc3-4d18-b63d-0b99560865d3"

	result, err := configurationaggregator.ResourceIbmConfigAggregatorSettingsMapToProfileTemplate(model)
	assert.Nil(t, err)
	checkResult(result)
}
