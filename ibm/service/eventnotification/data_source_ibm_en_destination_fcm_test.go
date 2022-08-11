// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package eventnotification_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMEnFCMDestinationDataSourceBasic(t *testing.T) {
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	instanceName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	description := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMEnFCMDestinationDataSourceConfigBasic(instanceName, name, description),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_en_destination_android.en_destination_data_6", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_en_destination_android.en_destination_data_6", "instance_guid"),
					resource.TestCheckResourceAttrSet("data.ibm_en_destination_android.en_destination_data_6", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_en_destination_android.en_destination_data_6", "description"),
					resource.TestCheckResourceAttrSet("data.ibm_en_destination_android.en_destination_data_6", "type"),
					resource.TestCheckResourceAttrSet("data.ibm_en_destination_android.en_destination_data_6", "updated_at"),
					resource.TestCheckResourceAttrSet("data.ibm_en_destination_android.en_destination_data_6", "destination_id"),
					resource.TestCheckResourceAttrSet("data.ibm_en_destination_android.en_destination_data_6", "subscription_count"),
				),
			},
		},
	})
}

func testAccCheckIBMEnFCMDestinationDataSourceConfigBasic(instanceName, name, description string) string {
	return fmt.Sprintf(`
	resource "ibm_resource_instance" "en_destination_datasource2" {
		name     = "%s"
		location = "us-south"
		plan     = "standard"
		service  = "event-notifications"
	}
	
	resource "ibm_en_destination_android" "en_destination_datasource_4" {
		instance_guid = ibm_resource_instance.en_destination_datasource2.guid
		name        = "%s"
		type        = "push_android"
		description = "%s"
		config {
			params {
				sender_id = "fcm sender id"
				server_key  = "fcm server key"
			}
		}
	}

		data "ibm_en_destination_android" "en_destination_data_6" {
			instance_guid = ibm_resource_instance.en_destination_datasource2.guid
			destination_id = ibm_en_destination_android.en_destination_datasource_4.destination_id
		}
	`, instanceName, name, description)
}
