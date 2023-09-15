// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package satellite_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"

	"github.com/IBM-Cloud/container-services-go-sdk/kubernetesserviceapiv1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccSatelliteStorageConfiguration_Basic(t *testing.T) {
	config_name := fmt.Sprintf("tf_config_name_%d", acctest.RandIntRange(10, 100))
	location := "satellite-location"
	storage_template_name := "template-name"
	storage_template_version := "template-version"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckSatelliteStorageConfigurationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckSatelliteStorageConfigurationCreate(location, config_name, storage_template_name, storage_template_version),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_satellite_storage_configuration.storage_configuration", "sc.#", "1"),
				),
			},
			{
				Config: testAccCheckSatelliteStorageConfigurationUpdate(location, config_name, storage_template_name, storage_template_version),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_satellite_storage_configuration.storage_configuration", "sc.#", "2"),
				),
			},
			{
				ResourceName:      "ibm_satellite_storage_configuration.storage_configuration",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckSatelliteStorageConfigurationDestroy(s *terraform.State) error {
	satClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).SatelliteClientSession()
	if err != nil {
		return err
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_satellite_storage_configuration" {
			continue
		}

		parts, _ := flex.IdParts(rs.Primary.ID)
		getStorageConfigurationOptions := &kubernetesserviceapiv1.GetStorageConfigurationOptions{
			Name: &parts[1],
		}

		_, _, err = satClient.GetStorageConfiguration(getStorageConfigurationOptions)
		if err == nil {
			return fmt.Errorf("Storage Configuration still exists: %s", rs.Primary.ID)
		}

	}
	return nil
}

func testAccCheckSatelliteStorageConfigurationUpdate(location string, config_name string, storage_template_name string, storage_template_version string) string {
	return fmt.Sprintf(`

	data "ibm_resource_group" "res_group" {
		is_default = true
	}

	resource "ibm_satellite_storage_configuration" "storage_configuration" {
		location = %s
		config_name = %s
		storage_template_name = %s
		storage_template_version = %s
		user_config_parameters = {
			numOfOsd = "2"
		}
		user_secret_parameters = {
			ibm-api-key = "value"
		}
	}
	  
`, location, config_name, storage_template_name, storage_template_version)
}

func testAccCheckSatelliteStorageConfigurationCreate(location string, config_name string, storage_template_name string, storage_template_version string) string {
	return fmt.Sprintf(`

	data "ibm_resource_group" "res_group" {
		is_default = true
	}

	resource "ibm_satellite_storage_configuration" "storage_configuration" {
		location = %s
		config_name = %s
		storage_template_name = %s
		storage_template_version = %s
		user_config_parameters = {
			numOfOsd = "1"
		}
		user_secret_parameters = {
			ibm-api-key = "value"
		}
	}
	  
`, location, config_name, storage_template_name, storage_template_version)
}
