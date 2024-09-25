// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.92.0-af5c89a5-20240617-153232
 */

package configurationaggregator_test

import (
	"fmt"
	"testing"

	"github.com/IBM/configuration-aggregator-go-sdk/configurationaggregatorv1"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/configurationaggregator"
	. "github.com/IBM-Cloud/terraform-provider-ibm/ibm/unittest"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/stretchr/testify/assert"
)

func TestAccIbmConfigAggregatorConfigurationsDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmConfigAggregatorConfigurationsDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_config_aggregator_configurations.config_aggregator_configurations_instance", "id"),
				),
			},
		},
	})
}

func testAccCheckIbmConfigAggregatorConfigurationsDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		data "ibm_config_aggregator_configurations" "config_aggregator_configurations_instance" {
			config_type = "config_type"
			service_name = "service_name"
			resource_group_id = "resource_group_id"
			location = "location"
			resource_crn = "resource_crn"
		}
	`)
}

func TestDataSourceIbmConfigAggregatorConfigurationsPaginatedPreviousToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["href"] = "testString"
		model["start"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(configurationaggregatorv1.PaginatedPrevious)
	model.Href = core.StringPtr("testString")
	model.Start = core.StringPtr("testString")

	result, err := configurationaggregator.DataSourceIbmConfigAggregatorConfigurationsPaginatedPreviousToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmConfigAggregatorConfigurationsConfigToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		tagsModel := make(map[string]interface{})
		tagsModel["tag"] = "testString"

		aboutModel := make(map[string]interface{})
		aboutModel["account_id"] = "testString"
		aboutModel["config_type"] = "testString"
		aboutModel["resource_crn"] = "testString"
		aboutModel["resource_group_id"] = "testString"
		aboutModel["service_name"] = "testString"
		aboutModel["resource_name"] = "testString"
		aboutModel["last_config_refresh_time"] = "2019-01-01T12:00:00.000Z"
		aboutModel["location"] = "testString"
		aboutModel["tags"] = []map[string]interface{}{tagsModel}

		configurationModel := make(map[string]interface{})

		model := make(map[string]interface{})
		model["about"] = []map[string]interface{}{aboutModel}
		model["config"] = []map[string]interface{}{configurationModel}

		assert.Equal(t, result, model)
	}

	tagsModel := new(configurationaggregatorv1.Tags)
	tagsModel.Tag = core.StringPtr("testString")

	aboutModel := new(configurationaggregatorv1.About)
	aboutModel.AccountID = core.StringPtr("testString")
	aboutModel.ConfigType = core.StringPtr("testString")
	aboutModel.ResourceCrn = core.StringPtr("testString")
	aboutModel.ResourceGroupID = core.StringPtr("testString")
	aboutModel.ServiceName = core.StringPtr("testString")
	aboutModel.ResourceName = core.StringPtr("testString")
	aboutModel.LastConfigRefreshTime = CreateMockDateTime("2019-01-01T12:00:00.000Z")
	aboutModel.Location = core.StringPtr("testString")
	aboutModel.Tags = tagsModel

	configurationModel := new(configurationaggregatorv1.Configuration)

	model := new(configurationaggregatorv1.Config)
	model.About = aboutModel
	model.Config = configurationModel

	result, err := configurationaggregator.DataSourceIbmConfigAggregatorConfigurationsConfigToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmConfigAggregatorConfigurationsAboutToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		tagsModel := make(map[string]interface{})
		tagsModel["tag"] = "testString"

		model := make(map[string]interface{})
		model["account_id"] = "testString"
		model["config_type"] = "testString"
		model["resource_crn"] = "testString"
		model["resource_group_id"] = "testString"
		model["service_name"] = "testString"
		model["resource_name"] = "testString"
		model["last_config_refresh_time"] = "2019-01-01T12:00:00.000Z"
		model["location"] = "testString"
		model["tags"] = []map[string]interface{}{tagsModel}

		assert.Equal(t, result, model)
	}

	tagsModel := new(configurationaggregatorv1.Tags)
	tagsModel.Tag = core.StringPtr("testString")

	model := new(configurationaggregatorv1.About)
	model.AccountID = core.StringPtr("testString")
	model.ConfigType = core.StringPtr("testString")
	model.ResourceCrn = core.StringPtr("testString")
	model.ResourceGroupID = core.StringPtr("testString")
	model.ServiceName = core.StringPtr("testString")
	model.ResourceName = core.StringPtr("testString")
	model.LastConfigRefreshTime = CreateMockDateTime("2019-01-01T12:00:00.000Z")
	model.Location = core.StringPtr("testString")
	model.Tags = tagsModel

	result, err := configurationaggregator.DataSourceIbmConfigAggregatorConfigurationsAboutToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmConfigAggregatorConfigurationsTagsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["tag"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(configurationaggregatorv1.Tags)
	model.Tag = core.StringPtr("testString")

	result, err := configurationaggregator.DataSourceIbmConfigAggregatorConfigurationsTagsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmConfigAggregatorConfigurationsConfigurationToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})

		assert.Equal(t, result, model)
	}

	model := new(configurationaggregatorv1.Configuration)

	result, err := configurationaggregator.DataSourceIbmConfigAggregatorConfigurationsConfigurationToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}
