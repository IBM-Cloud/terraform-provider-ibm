// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.92.0-af5c89a5-20240617-153232
 */

package configurationaggregator_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIbmConfigAggregatorConfigurationsDataSourceBasic(t *testing.T) {
	instanceID := "instance_id"
	var configType = "your-config-type"
	var location = "your-location"
	var resourceCrn = "your-resource-crn"
	var resourceGroupID = "your-resource-group-id"
	var serviceName = "your-service-name"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIbmConfigAggregatorConfigurationsDataSourceConfigBasic(instanceID, configType, location, resourceCrn, resourceGroupID, serviceName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_config_aggregator_configurations.config_aggregator_configurations_instance", "id"),
					resource.TestCheckResourceAttr("data.ibm_config_aggregator_configurations.config_aggregator_configurations_instance", "config_type", configType),
					resource.TestCheckResourceAttr("data.ibm_config_aggregator_configurations.config_aggregator_configurations_instance", "location", location),
					resource.TestCheckResourceAttr("data.ibm_config_aggregator_configurations.config_aggregator_configurations_instance", "resource_crn", resourceCrn),
					resource.TestCheckResourceAttr("data.ibm_config_aggregator_configurations.config_aggregator_configurations_instance", "resource_group_id", resourceGroupID),
					resource.TestCheckResourceAttr("data.ibm_config_aggregator_configurations.config_aggregator_configurations_instance", "service_name", serviceName),
				),
			},
		},
	})
}

func testAccCheckIbmConfigAggregatorConfigurationsDataSourceConfigBasic(instanceID, configType, location, resourceCrn, resourceGroupID, serviceName string) string {
	return fmt.Sprintf(`
		data "ibm_config_aggregator_configurations" "config_aggregator_configurations_instance" {
			instance_id ="%s"
			config_type       = "%s"
			location          = "%s"
			resource_crn      = "%s"
			resource_group_id = "%s"
			service_name      = "%s"
		}
	`, instanceID, configType, location, resourceCrn, resourceGroupID, serviceName)
}
