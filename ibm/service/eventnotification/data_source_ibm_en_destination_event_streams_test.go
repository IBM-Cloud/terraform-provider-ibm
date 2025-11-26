// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package eventnotification_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccIBMEnEventStreamsDestinationDataSourceBasic(t *testing.T) {
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	instanceName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	description := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMEnEventStreamsDestinationDataSourceConfigBasic(instanceName, name, description),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_en_destination_event_streams.en_destination_data_6", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_en_destination_event_streams.en_destination_data_6", "instance_guid"),
					resource.TestCheckResourceAttrSet("data.ibm_en_destination_event_streams.en_destination_data_6", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_en_destination_event_streams.en_destination_data_6", "description"),
					resource.TestCheckResourceAttrSet("data.ibm_en_destination_event_streams.en_destination_data_6", "type"),
					resource.TestCheckResourceAttrSet("data.ibm_en_destination_event_streams.en_destination_data_6", "updated_at"),
					resource.TestCheckResourceAttrSet("data.ibm_en_destination_event_streams.en_destination_data_6", "destination_id"),
					resource.TestCheckResourceAttrSet("data.ibm_en_destination_event_streams.en_destination_data_6", "subscription_count"),
				),
			},
		},
	})
}

func testAccCheckIBMEnEventStreamsDestinationDataSourceConfigBasic(instanceName, name, description string) string {
	return fmt.Sprintf(`
	resource "ibm_resource_instance" "en_destination_datasource2" {
		name     = "%s"
		location = "us-south"
		plan     = "standard"
		service  = "event-notifications"
	}
	
	resource "ibm_en_destination_event_streams" "en_destination_datasource_4" {
		instance_guid = ibm_resource_instance.en_destination_resource.guid
		name        = "%s"
		type        = "event_streams"
		description = "%s"
		config {
			params {
				crn = "crn:v1:staging:public:messagehub:us-south:a/9f007405a9fe4a5d9345fa8c131610c8:a292db6e-af78-4c0b-b3db-7d6794b40aeb::"
				endpoint = "https://n6627w6t7y62cfgd.svc09.us-south.eventstreams.test.cloud.ibm.com"
				topic = "test_demo"
			}
		}
	}

		data "ibm_en_destination_event_streams" "en_destination_data_6" {
			instance_guid = ibm_resource_instance.en_destination_datasource2.guid
			destination_id = ibm_en_destination_event_streams.en_destination_datasource_4.destination_id
		}
	`, instanceName, name, description)
}
