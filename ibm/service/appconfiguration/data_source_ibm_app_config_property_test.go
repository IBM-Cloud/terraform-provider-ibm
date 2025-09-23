// Copyright IBM Corp. 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0
package appconfiguration_test

import (
	"fmt"
	"testing"

	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIbmAppConfigPropertyDataSource(t *testing.T) {
	environmentID := "dev"
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	propertyID := fmt.Sprintf("tf_property_id_%d", acctest.RandIntRange(10, 100))
	propertyType := "BOOLEAN"
	description := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))
	instanceName := fmt.Sprintf("tf_app_config_test_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIbmAppConfigPropertyDataSourceConfigBasic(instanceName, environmentID, name, propertyID, propertyType, description),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_app_config_property.ibm_app_config_property_data1", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_app_config_property.ibm_app_config_property_data1", "environment_id"),
					resource.TestCheckResourceAttrSet("data.ibm_app_config_property.ibm_app_config_property_data1", "property_id"),
					resource.TestCheckResourceAttrSet("data.ibm_app_config_property.ibm_app_config_property_data1", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_app_config_property.ibm_app_config_property_data1", "description"),
					resource.TestCheckResourceAttrSet("data.ibm_app_config_property.ibm_app_config_property_data1", "type"),
					resource.TestCheckResourceAttrSet("data.ibm_app_config_property.ibm_app_config_property_data1", "value"),
					resource.TestCheckResourceAttrSet("data.ibm_app_config_property.ibm_app_config_property_data1", "segment_rules.#"),
					resource.TestCheckResourceAttrSet("data.ibm_app_config_property.ibm_app_config_property_data1", "segment_exists"),
					resource.TestCheckResourceAttrSet("data.ibm_app_config_property.ibm_app_config_property_data1", "created_time"),
					resource.TestCheckResourceAttrSet("data.ibm_app_config_property.ibm_app_config_property_data1", "updated_time"),
					resource.TestCheckResourceAttrSet("data.ibm_app_config_property.ibm_app_config_property_data1", "href"),
				),
			},
		},
	})
}

func testAccCheckIbmAppConfigPropertyDataSourceConfigBasic(instanceName, environmentID, name, propertyID, propertyType, description string) string {
	return fmt.Sprintf(`
	    resource "ibm_resource_instance" "app_config_terraform_test482" {
    		name     = "%s"
    		location = "us-south"
    		service  = "apprapp"
    		plan     = "lite"
    	}
		resource "ibm_app_config_property" "app_config_property" {
		    guid           	= ibm_resource_instance.app_config_terraform_test482.guid
			environment_id = "%s"
			name = "%s"
			property_id = "%s"
			type = "%s"
			value = "true"
			description  = "%s"
		}
		data "ibm_app_config_property" "ibm_app_config_property_data1" {
		    guid           	= ibm_resource_instance.app_config_terraform_test482.guid
			environment_id = ibm_app_config_property.app_config_property.environment_id
			property_id = ibm_app_config_property.app_config_property.property_id
			include = "collections"
		}
	`, instanceName, environmentID, name, propertyID, propertyType, description)
}
