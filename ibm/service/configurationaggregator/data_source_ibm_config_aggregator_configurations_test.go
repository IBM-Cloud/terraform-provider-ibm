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
