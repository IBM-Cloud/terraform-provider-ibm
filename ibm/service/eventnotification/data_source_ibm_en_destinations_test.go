// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package eventnotification_test

import (
	"fmt"
	"testing"

	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMEnDestinationsDataSourceBasic(t *testing.T) {
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	instanceName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	description := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMEnDestinationDatasourceConfig(instanceName, name, description),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_en_destinations.data_destination_2", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_en_destinations.data_destination_2", "instance_guid"),
					resource.TestCheckResourceAttrSet("data.ibm_en_destinations.data_destination_2", "total_count"),
					resource.TestCheckResourceAttrSet("data.ibm_en_destinations.data_destination_2", "destinations.#"),
					resource.TestCheckResourceAttrSet("data.ibm_en_destinations.data_destination_2", "destinations.0.name"),
					resource.TestCheckResourceAttrSet("data.ibm_en_destinations.data_destination_2", "destinations.0.updated_at"),
					resource.TestCheckResourceAttrSet("data.ibm_en_destinations.data_destination_2", "destinations.0.type"),
					resource.TestCheckResourceAttrSet("data.ibm_en_destinations.data_destination_2", "destinations.0.id"),
					resource.TestCheckResourceAttrSet("data.ibm_en_destinations.data_destination_2", "destinations.0.description"),
				),
			},
		},
	})
}

func testAccCheckIBMEnDestinationDatasourceConfig(instanceName, name, description string) string {
	return fmt.Sprintf(`
	resource "ibm_resource_instance" "en_destination_datasource" {
		name     = "%s"
		location = "us-south"
		plan     = "standard"
		service  = "event-notifications"
	}
	
	resource "ibm_en_destination" "en_destination_datasource_1" {
		instance_guid = ibm_resource_instance.en_destination_datasource.guid
		name        = "%s"
		type        = "webhook"
		description = "%s"
		config {
			params {
				verb = "POST"
				url  = "https://demo.webhook.com"
			}
		}
	}

	data "ibm_en_destinations" "data_destination_2" {
		instance_guid = ibm_resource_instance.en_destination_datasource.guid
	}
	`, instanceName, name, description)
}
