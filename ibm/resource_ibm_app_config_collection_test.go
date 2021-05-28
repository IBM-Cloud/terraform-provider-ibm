// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/IBM/appconfiguration-go-admin-sdk/appconfigurationv1"
)

func TestAccIbmAppConfigCollectionBasic(t *testing.T) {
	var conf appconfigurationv1.Collection
	instanceName := fmt.Sprintf("tf_collection_test__%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	collectionID := fmt.Sprintf("tf_collection_id_%d", acctest.RandIntRange(10, 100))
	description := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))
	tags := fmt.Sprintf("tf_tags_%d", acctest.RandIntRange(10, 100))
	nameUpdate := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	descriptionUpdate := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))
	tagsUpdate := fmt.Sprintf("tf_tags_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIbmAppConfigCollectionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIbmAppConfigCollectionConfigBasic(instanceName, name, collectionID, description, tags),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmAppConfigCollectionExists("ibm_app_config_collection.app_config_collection_resource1", conf),
					resource.TestCheckResourceAttrSet("ibm_app_config_collection.app_config_collection_resource1", "name"),
					resource.TestCheckResourceAttrSet("ibm_app_config_collection.app_config_collection_resource1", "tags"),
					resource.TestCheckResourceAttrSet("ibm_app_config_collection.app_config_collection_resource1", "description"),
					resource.TestCheckResourceAttrSet("ibm_app_config_collection.app_config_collection_resource1", "collection_id"),
				),
			},
			{
				Config: testAccCheckIbmAppConfigCollectionConfigBasic(instanceName, nameUpdate, collectionID, descriptionUpdate, tagsUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_app_config_collection.app_config_collection_resource1", "name", nameUpdate),
					resource.TestCheckResourceAttr("ibm_app_config_collection.app_config_collection_resource1", "tags", tagsUpdate),
					resource.TestCheckResourceAttr("ibm_app_config_collection.app_config_collection_resource1", "description", descriptionUpdate),
				),
			},
			{
				ResourceName:      "ibm_app_config_collection.app_config_collection_resource1",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIbmAppConfigCollectionConfigBasic(instanceName, name, collectionID, description, tags string) string {
	return fmt.Sprintf(`
		 resource "ibm_resource_instance" "app_config_terraform_test464"{
			 name     = "%s"
			 location = "us-south"
			 service  = "apprapp"
			 plan     = "standard"
		 }
		 resource "ibm_app_config_collection" "app_config_collection_resource1" {
			 name          		= "%s"
			 collection_id   	= "%s"
			 description      = "%s"
			 tags							= "%s"
			 guid 						= ibm_resource_instance.app_config_terraform_test464.guid
		 }`, instanceName, name, collectionID, description, tags)
}

func testAccCheckIbmAppConfigCollectionExists(n string, obj appconfigurationv1.Collection) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}
		parts, err := idParts(rs.Primary.ID)
		if err != nil {
			return err
		}

		appconfigClient, err := getAppConfigClient(testAccProvider.Meta(), parts[0])
		if err != nil {
			return err
		}

		options := &appconfigurationv1.GetCollectionOptions{}

		options.SetCollectionID(parts[1])

		reslut, _, err := appconfigClient.GetCollection(options)
		if err != nil {
			return err
		}

		obj = *reslut
		return nil
	}
}

func testAccCheckIbmAppConfigCollectionDestroy(s *terraform.State) error {
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "app_config_collection_resource1" {
			continue
		}

		parts, err := idParts(rs.Primary.ID)
		if err != nil {
			return err
		}

		appconfigClient, err := getAppConfigClient(testAccProvider.Meta(), parts[0])
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
			return fmt.Errorf("Error checking for Collection (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
