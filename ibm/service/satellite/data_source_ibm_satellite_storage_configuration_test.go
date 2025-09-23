// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package satellite_test

import (
	"fmt"
	"testing"

	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIbmSatelliteStorageConfigurationDataSourceBasic(t *testing.T) {
	configName := fmt.Sprintf("tf-config-name-%d", acctest.RandIntRange(10, 100))
	location := "satellite-location"
	secretKey := "apikey"
	secretValue := "apikey - value"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIbmSatelliteStorageConfigurationDataSourceConfigBasic(location, configName, secretKey, secretValue),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_satellite_storage_configuration.satellite_storage_configuration", "uuid"),
					resource.TestCheckResourceAttrSet("data.ibm_satellite_storage_configuration.satellite_storage_configuration", "location"),
					resource.TestCheckResourceAttrSet("data.ibm_satellite_storage_configuration.satellite_storage_configuration", "config_name"),
					resource.TestCheckResourceAttrSet("data.ibm_satellite_storage_configuration.satellite_storage_configuration", "config_version"),
					resource.TestCheckResourceAttrSet("data.ibm_satellite_storage_configuration.satellite_storage_configuration", "storage_template_name"),
					resource.TestCheckResourceAttrSet("data.ibm_satellite_storage_configuration.satellite_storage_configuration", "storage_template_version"),
					resource.TestCheckResourceAttrSet("data.ibm_satellite_storage_configuration.satellite_storage_configuration", "user_config_parameters"),
					resource.TestCheckResourceAttrSet("data.ibm_satellite_storage_configuration.satellite_storage_configuration", "user_secret_parameters"),
					resource.TestCheckResourceAttrSet("data.ibm_satellite_storage_configuration.satellite_storage_configuration", "storage_class_parameters"),
				),
			},
		},
	})
}

func testAccCheckIbmSatelliteStorageConfigurationDataSourceConfigBasic(location string, configName string, secretKey string, secretValue string) string {
	return fmt.Sprintf(`
		data "ibm_satellite_storage_configuration" "satellite_storage_configuration" {
			location = "%s"
			config_name = "%s"
			user_secret_parameters = {
				%s : "%s"
		    }
		}
	`, location, configName, secretKey, secretValue)
}
