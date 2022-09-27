// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package appconfiguration_test

import (
	"fmt"
	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"testing"

	"github.com/IBM/appconfiguration-go-admin-sdk/appconfigurationv1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
)

func TestAccIbmIbmAppConfigCollectionBasic(t *testing.T) {
	var conf appconfigurationv1.Collection
	instanceName := fmt.Sprintf("tf_app_config_test_%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	collectionId := fmt.Sprintf("tf_collection_id_%d", acctest.RandIntRange(10, 100))
	description := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))
	descriptionUpdate := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))
	nameUpdate := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	tags := fmt.Sprintf("tags_%d", acctest.RandIntRange(10, 100))
	tagsUpdated := fmt.Sprintf("tags_updated_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmAppConfigCollectionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIbmAppConfigCollectionConfigBasic(instanceName, name, collectionId, description, tags),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmAppConfigCollectionExists("ibm_app_config_collection.ibm_app_config_collection_resource1", conf),
					resource.TestCheckResourceAttrSet("ibm_app_config_collection.ibm_app_config_collection_resource1", "id"),
					resource.TestCheckResourceAttrSet("ibm_app_config_collection.ibm_app_config_collection_resource1", "name"),
					resource.TestCheckResourceAttrSet("ibm_app_config_collection.ibm_app_config_collection_resource1", "tags"),
					resource.TestCheckResourceAttrSet("ibm_app_config_collection.ibm_app_config_collection_resource1", "collection_id"),
					resource.TestCheckResourceAttrSet("ibm_app_config_collection.ibm_app_config_collection_resource1", "description"),
				),
			},
			{
				Config: testAccCheckIbmAppConfigCollectionConfigBasic(instanceName, nameUpdate, collectionId, descriptionUpdate, tagsUpdated),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_app_config_collection.ibm_app_config_collection_resource1", "name", nameUpdate),
					resource.TestCheckResourceAttr("ibm_app_config_collection.ibm_app_config_collection_resource1", "tags", tagsUpdated),
					resource.TestCheckResourceAttr("ibm_app_config_collection.ibm_app_config_collection_resource1", "description", descriptionUpdate),
				),
			},
		},
	})
}

func testAccCheckIbmAppConfigCollectionConfigBasic(name, envName, collectionID, description, tags string) string {
	return fmt.Sprintf(`
		resource "ibm_resource_instance" "app_config_terraform_test456" {
			name     = "%s"
			location = "us-south"
			service  = "apprapp"
			plan     = "lite"
		}
		resource "ibm_app_config_collection" "ibm_app_config_collection_resource1" {
			guid           	    = ibm_resource_instance.app_config_terraform_test456.guid
			name            	= "%s"
			collection_id     	= "%s"
			description    	    = "%s"
			tags    		    = "%s"
		}`, name, envName, collectionID, description, tags)
}

func testAccCheckIbmAppConfigCollectionExists(n string, obj appconfigurationv1.Collection) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}
		parts, err := flex.IdParts(rs.Primary.ID)
		if err != nil {
			return err
		}
		appconfigClient, err := getAppConfigClient(acc.TestAccProvider.Meta(), parts[0])
		if err != nil {
			return err
		}

		options := &appconfigurationv1.GetCollectionOptions{}

		options.SetCollectionID(parts[1])

		result, _, err := appconfigClient.GetCollection(options)
		if err != nil {
			return err
		}

		obj = *result
		return nil
	}
}

func testAccCheckIbmAppConfigCollectionDestroy(s *terraform.State) error {

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_app_config_collection_resource1" {
			continue
		}
		parts, err := flex.IdParts(rs.Primary.ID)
		if err != nil {
			return err
		}
		appconfigClient, err := getAppConfigClient(acc.TestAccProvider.Meta(), parts[0])
		if err != nil {
			return err
		}
		options := &appconfigurationv1.GetCollectionOptions{}

		options.SetCollectionID(parts[1])

		// Try to find the key
		_, response, err := appconfigClient.GetCollection(options)

		if err == nil {
			return fmt.Errorf("Collection still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("[ERROR] Error checking for Collection (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
