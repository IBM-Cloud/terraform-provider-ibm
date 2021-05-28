// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIbmAppConfigCollectionsDataSourceBasic(t *testing.T) {
	instanceName := fmt.Sprintf("tf_collection_test__%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	collectionID := fmt.Sprintf("tf_collection_id_%d", acctest.RandIntRange(10, 100))
	description := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))
	tags := fmt.Sprintf("tf_tags_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIbmAppConfigCollectionsDataSourceConfigBasic(instanceName, name, collectionID, description, tags),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_app_config_collections.app_config_collection_data1", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_app_config_collections.app_config_collection_data1", "first.#"),
					resource.TestCheckResourceAttrSet("data.ibm_app_config_collections.app_config_collection_data1", "last.#"),
					resource.TestCheckResourceAttrSet("data.ibm_app_config_collections.app_config_collection_data1", "collections.#"),
					resource.TestCheckResourceAttrSet("data.ibm_app_config_collections.app_config_collection_data1", "collections.0.name"),
					resource.TestCheckResourceAttrSet("data.ibm_app_config_collections.app_config_collection_data1", "collections.0.description"),
					resource.TestCheckResourceAttrSet("data.ibm_app_config_collections.app_config_collection_data1", "collections.0.tags"),
					resource.TestCheckResourceAttrSet("data.ibm_app_config_collections.app_config_collection_data1", "collections.0.href"),
					resource.TestCheckResourceAttrSet("data.ibm_app_config_collections.app_config_collection_data1", "collections.0.created_time"),
					resource.TestCheckResourceAttrSet("data.ibm_app_config_collections.app_config_collection_data1", "collections.0.updated_time"),
					resource.TestCheckResourceAttrSet("data.ibm_app_config_collections.app_config_collection_data1", "collections.0.collection_id"),
				),
			},
		},
	})
}

func testAccCheckIbmAppConfigCollectionsDataSourceConfigBasic(instanceName, name, collectionID, description, tags string) string {
	return fmt.Sprintf(`
		 resource "ibm_resource_instance" "app_config_terraform_test462"{
			 name     = "%s"
			 location = "us-south"
			 service  = "apprapp"
			 plan     = "standard"
		 }
		 resource "ibm_app_config_collection" "app_config_collection_resource3" {
			 name          		= "%s"
			 collection_id    = "%s"
			 description      = "%s"
			 tags        			= "%s"
			 guid = ibm_resource_instance.app_config_terraform_test462.guid
		 }
		 data "ibm_app_config_collections" "app_config_collection_data1" {
			 expand 						= true
			 guid 							= ibm_app_config_collection.app_config_collection_resource3.guid
		 }`, instanceName, name, collectionID, description, tags)
}
