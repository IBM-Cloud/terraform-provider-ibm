// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0
package appconfiguration_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccIbmAppConfigCollectionDataSource(t *testing.T) {
	collectionID := fmt.Sprintf("tf_collection_id_%d", acctest.RandIntRange(10, 100))
	instanceName := fmt.Sprintf("tf_app_config_test_%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	description := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))
	tags := "development collection"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIbmAppConfigCollectionDataSourceConfigBasic(instanceName, name, collectionID, description, tags),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_app_config_collection.ibm_app_config_collection_data1", "collection_id"),
					resource.TestCheckResourceAttrSet("data.ibm_app_config_collection.ibm_app_config_collection_data1", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_app_config_collection.ibm_app_config_collection_data1", "expand"),
					resource.TestCheckResourceAttrSet("data.ibm_app_config_collection.ibm_app_config_collection_data1", "description"),
					resource.TestCheckResourceAttrSet("data.ibm_app_config_collection.ibm_app_config_collection_data1", "tags"),
					resource.TestCheckResourceAttrSet("data.ibm_app_config_collection.ibm_app_config_collection_data1", "created_time"),
					resource.TestCheckResourceAttrSet("data.ibm_app_config_collection.ibm_app_config_collection_data1", "updated_time"),
				),
			},
		},
	})
}

func testAccCheckIbmAppConfigCollectionDataSourceConfigBasic(instanceName, name, collectionID, description, tags string) string {
	return fmt.Sprintf(`
		resource "ibm_resource_instance" "app_config_terraform_test487" {
			name     = "%s"
			location = "us-south"
			service  = "apprapp"
			plan     = "lite"
		}
		
		resource "ibm_app_config_collection" "app_config_collection_resource1" {
			guid           	= ibm_resource_instance.app_config_terraform_test487.guid
			name 			= "%s"
			collection_id   = "%s"
			description		= "%s"
			tags			= "%s"
		}
		
		data "ibm_app_config_collection" "ibm_app_config_collection_data1" {
			guid			= ibm_app_config_collection.app_config_collection_resource1.guid
			collection_id	= ibm_app_config_collection.app_config_collection_resource1.collection_id
			expand			= "true"
		}
		`, instanceName, name, collectionID, description, tags)
}
