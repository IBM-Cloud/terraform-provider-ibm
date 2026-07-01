// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.114.3-943fbc81-20260603-173645
 */

package configurationaggregator_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/configurationaggregator"
	"github.com/IBM/configuration-aggregator-go-sdk/configurationaggregatorv1"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/stretchr/testify/assert"
)

func TestAccIbmConfigAggregatorBatchConfigurationsDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmConfigAggregatorBatchConfigurationsDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_config_aggregator_batch_configurations.config_aggregator_batch_configurations_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_config_aggregator_batch_configurations.config_aggregator_batch_configurations_instance", "configs"),
				),
			},
		},
	})
}

func testAccCheckIbmConfigAggregatorBatchConfigurationsDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		data "ibm_config_aggregator_batch_configurations" "config_aggregator_batch_configurations_instance" {
			configs {"resource_crn":"crn:v1:staging:public:cloud-object-storage:global:a/c1d20fee2fe24c42b8ef6583283d2dcf:fc8de30c-43a5-407b-8c61-2b86dd820922:bucket:cos-pra-1"},{"resource_crn":"crn:v1:staging:public:event-notifications:us-south:a/c1d20fee2fe24c42b8ef6583283d2dcf:4e3b5d24-3acc-4352-a98d-788749e1d7da::","service_name":"event-notifications","config_type":["source","topic"]},{"resource_crn":"crn:v1:staging:public:event-notifications:us-south:a/c1d20fee2fe24c42b8ef6583283d2dcf:4e3b5d24-3acc-4352-a98d-788749e1d7db::","type_id":"e3b5d24-3acc-4352-a98d-788749e1d7db"}
		}
	`)
}

func TestDataSourceIbmConfigAggregatorBatchConfigurationsPaginatedPreviousToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["href"] = "testString"
		model["start"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(configurationaggregatorv1.PaginatedPrevious)
	model.Href = core.StringPtr("testString")
	model.Start = core.StringPtr("testString")

	result, err := configurationaggregator.DataSourceIbmConfigAggregatorBatchConfigurationsPaginatedPreviousToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}
