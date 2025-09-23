// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0
package appconfiguration_test

import (
	"fmt"
	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

func TestAccIbmAppConfigCollectionsDataSource(t *testing.T) {
	collectionID := fmt.Sprintf("tf_collection_id_%d", acctest.RandIntRange(10, 100))
	instanceName := fmt.Sprintf("tf_app_config_test_%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	description := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))
	tags := "development collections"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIbmAppConfigCollectionsDataSourceConfigBasic(instanceName, name, collectionID, description, tags),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_app_config_collections.app_config_collections_data2", "limit"),
					resource.TestCheckResourceAttrSet("data.ibm_app_config_collections.app_config_collections_data2", "offset"),
					resource.TestCheckResourceAttrSet("data.ibm_app_config_collections.app_config_collections_data2", "total_count"),
					resource.TestCheckResourceAttrSet("data.ibm_app_config_collections.app_config_collections_data2", "collections.#"),
					resource.TestCheckResourceAttr("data.ibm_app_config_collections.app_config_collections_data2", "collections.0.name", name),
					resource.TestCheckResourceAttr("data.ibm_app_config_collections.app_config_collections_data2", "collections.0.collection_id", collectionID),
					resource.TestCheckResourceAttr("data.ibm_app_config_collections.app_config_collections_data2", "collections.0.description", description),
					resource.TestCheckResourceAttr("data.ibm_app_config_collections.app_config_collections_data2", "collections.0.tags", tags),
					resource.TestCheckResourceAttrSet("data.ibm_app_config_collections.app_config_collections_data2", "collections.0.created_time"),
					resource.TestCheckResourceAttrSet("data.ibm_app_config_collections.app_config_collections_data2", "collections.0.updated_time"),
				),
			},
		},
	})
}

func testAccCheckIbmAppConfigCollectionsDataSourceConfigBasic(instanceName, name, collectionID, description, tags string) string {
	return fmt.Sprintf(`
		resource "ibm_resource_instance" "app_config_terraform_test487" {
			name     = "%s"
			location = "us-south"
			service  = "apprapp"
			plan     = "lite"
		}
		
		resource "ibm_app_config_collection" "app_config_collection_resource2" {
			guid           	= ibm_resource_instance.app_config_terraform_test487.guid
			name 			= "%s"
			collection_id   = "%s"
			description		= "%s"
			tags			= "%s"
		}
		
		data "ibm_app_config_collections" "app_config_collections_data2" {
			guid			= ibm_app_config_collection.app_config_collection_resource2.guid
			expand			= "true"
		}
		`, instanceName, name, collectionID, description, tags)
}
