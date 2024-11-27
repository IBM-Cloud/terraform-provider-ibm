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

func TestAccIBMEnSourcesDataSourceBasic(t *testing.T) {
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	instanceName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	description := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))
	enabled := true
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMEnSourceDatasourceConfig(instanceName, name, description, enabled),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_en_sources.data_source_2", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_en_sources.data_source_2", "instance_guid"),
					resource.TestCheckResourceAttrSet("data.ibm_en_sources.data_source_2", "total_count"),
					resource.TestCheckResourceAttrSet("data.ibm_en_sources.data_source_2", "sources.#"),
					resource.TestCheckResourceAttrSet("data.ibm_en_sources.data_source_2", "sources.0.name"),
					resource.TestCheckResourceAttrSet("data.ibm_en_sources.data_source_2", "sources.0.updated_at"),
					resource.TestCheckResourceAttrSet("data.ibm_en_sources.data_source_2", "sources.0.type"),
					resource.TestCheckResourceAttrSet("data.ibm_en_sources.data_source_2", "sources.0.id"),
					resource.TestCheckResourceAttrSet("data.ibm_en_sourcess.data_source_2", "sources.0.description"),
					resource.TestCheckResourceAttrSet("data.ibm_en_sourcess.data_source_2", "sources.0.enabled"),
				),
			},
		},
	})
}

func testAccCheckIBMEnSourceDatasourceConfig(instanceName, name, description string, enabled bool) string {
	return fmt.Sprintf(`
	resource "ibm_resource_instance" "en_source_datasource" {
		name     = "%s"
		location = "us-south"
		plan     = "standard"
		service  = "event-notifications"
	}
	
	resource "ibm_en_source" "en_source_datasource_1" {
		instance_guid = ibm_resource_instance.en_source_datasource.guid
		name        = "%s"
		description = "%s"
		enabled = %t
	}

	data "ibm_en_sources" "data_source_2" {
		instance_guid = ibm_resource_instance.en_source_datasource.guid
	}
	`, instanceName, name, description, enabled)
}
