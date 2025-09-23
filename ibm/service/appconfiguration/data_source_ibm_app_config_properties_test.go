// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package appconfiguration_test

import (
	"fmt"
	"testing"

	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIbmAppConfigPropertiesDataSourceBasic(t *testing.T) {
	propertyType := "BOOLEAN"
	name := fmt.Sprintf("tf_test_%d", acctest.RandIntRange(10, 100))
	propertyID := fmt.Sprintf("tf_property_id_%d", acctest.RandIntRange(10, 100))
	description := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))
	tags := fmt.Sprintf("tf_tags_%d", acctest.RandIntRange(10, 100))
	instanceName := fmt.Sprintf("tf_app_config_test_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIbmAppConfigPropertiesDataSourceConfig(instanceName, name, propertyID, propertyType, description, tags),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_app_config_properties.app_config_properties", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_app_config_properties.app_config_properties", "properties.#"),
					resource.TestCheckResourceAttrSet("data.ibm_app_config_properties.app_config_properties", "properties.0.name"),
					resource.TestCheckResourceAttrSet("data.ibm_app_config_properties.app_config_properties", "properties.0.property_id"),
					resource.TestCheckResourceAttrSet("data.ibm_app_config_properties.app_config_properties", "properties.0.type"),
					resource.TestCheckResourceAttrSet("data.ibm_app_config_properties.app_config_properties", "properties.0.description"),
					resource.TestCheckResourceAttrSet("data.ibm_app_config_properties.app_config_properties", "properties.0.tags"),
					resource.TestCheckResourceAttrSet("data.ibm_app_config_properties.app_config_properties", "total_count"),
				),
			},
		},
	})
}

func testAccCheckIbmAppConfigPropertiesDataSourceConfig(instanceName, name, propertyID, propertyType, description, tags string) string {
	return fmt.Sprintf(`
	resource "ibm_resource_instance" "app_config_terraform_test494" {
		name     = "%s"
		location = "us-south"
		service  = "apprapp"
		plan     = "lite"
	}
		resource "ibm_app_config_property" "app_config_properties_resource21" {
			guid = ibm_resource_instance.app_config_terraform_test494.guid
			environment_id = "dev"
			name = "%s"
			property_id = "%s"
			type = "%s"
			value = "false"
			description = "%s"
			tags = "%s"
		}
		data "ibm_app_config_properties" "app_config_properties" {
			guid = ibm_app_config_property.app_config_properties_resource21.guid
			environment_id = ibm_app_config_property.app_config_properties_resource21.environment_id
			expand = true
		}
	`, instanceName, name, propertyID, propertyType, description, tags)
}
