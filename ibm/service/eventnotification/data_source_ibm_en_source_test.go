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

func TestAccIBMEnSourceDataSourceBasic(t *testing.T) {
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	instanceName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	description := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMEnSourceDataSourceConfigBasic(instanceName, name, description),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_en_source.en_source_data_1", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_en_source.en_source_data_1", "instance_guid"),
					resource.TestCheckResourceAttrSet("data.ibm_en_source.en_source_data_1", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_en_source.en_source_data_1", "description"),
					resource.TestCheckResourceAttrSet("data.ibm_en_source.en_source_data_1", "updated_at"),
				),
			},
		},
	})
}

func testAccCheckIBMEnSourceDataSourceConfigBasic(instanceName, name, description string) string {
	return fmt.Sprintf(`
	resource "ibm_resource_instance" "en_source_datasource2" {
		name     = "%s"
		location = "us-south"
		plan     = "standard"
		service  = "event-notifications"
	}
	
	resource "ibm_en_source" "en_source_resource_1" {
		instance_guid = ibm_resource_instance.en_source_resource.guid
		name        = "%s"
		enabled = true
		description = "%s"
	}

		data "ibm_en_source" "en_source_data_1" {
			instance_guid = ibm_resource_instance.en_source_datasource2.guid
			destination_id = ibm_en_source.en_source_datasource_4.destination_id
		}
	`, instanceName, name, description)
}
