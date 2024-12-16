// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.92.0-af5c89a5-20240617-153232
 */

package configurationaggregator_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/configurationaggregator"
	"github.com/IBM/configuration-aggregator-go-sdk/configurationaggregatorv1"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/stretchr/testify/assert"
)

func TestAccIbmConfigAggregatorSettingsDataSourceBasic(t *testing.T) {
	instanceID := "instance_id"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmConfigAggregatorSettingsDataSourceConfigBasic(instanceID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_config_aggregator_settings.config_aggregator_settings_instance", "id"),
				),
			},
		},
	})
}

func testAccCheckIbmConfigAggregatorSettingsDataSourceConfigBasic(instanceID string) string {
	return fmt.Sprintf(`
		data "ibm_config_aggregator_settings" "config_aggregator_settings_instance" {
		instance_id="%s"
		}
	`, instanceID)
}

func TestDataSourceIbmConfigAggregatorSettingsAdditionalScopeToMap(t *testing.T) {
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

	result, err := configurationaggregator.DataSourceIbmConfigAggregatorSettingsAdditionalScopeToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmConfigAggregatorSettingsProfileTemplateToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["id"] = "ProfileTemplate-adb55769-ae22-4c60-aead-bd1f84f93c57"
		model["trusted_profile_id"] = "Profile-6bb60124-8fc3-4d18-b63d-0b99560865d3"

		assert.Equal(t, result, model)
	}

	model := new(configurationaggregatorv1.ProfileTemplate)
	model.ID = core.StringPtr("ProfileTemplate-adb55769-ae22-4c60-aead-bd1f84f93c57")
	model.TrustedProfileID = core.StringPtr("Profile-6bb60124-8fc3-4d18-b63d-0b99560865d3")

	result, err := configurationaggregator.DataSourceIbmConfigAggregatorSettingsProfileTemplateToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}
