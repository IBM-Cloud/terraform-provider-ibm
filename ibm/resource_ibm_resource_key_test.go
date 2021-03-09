// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/IBM-Cloud/bluemix-go/models"
)

func TestAccIBMResourceKey_Basic(t *testing.T) {
	var conf models.ServiceKey
	resourceName := fmt.Sprintf("tf-cos-%d", acctest.RandIntRange(10, 100))
	resourceKey := fmt.Sprintf("tf-cos-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMResourceKeyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMResourceKeyBasic(resourceName, resourceKey),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMResourceKeyExists("ibm_resource_key.resourceKey", conf),
					resource.TestCheckResourceAttr("ibm_resource_key.resourceKey", "name", resourceKey),
					resource.TestCheckResourceAttr("ibm_resource_key.resourceKey", "credentials.%", "7"),
					resource.TestCheckResourceAttr("ibm_resource_key.resourceKey", "role", "Reader"),
				),
			},
			{
				ResourceName:      "ibm_resource_key.resourceKey",
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"resource_instance_id", "resource_alias_id"},
			},
		},
	})
}

func TestAccIBMResourceKey_With_Tags(t *testing.T) {
	var conf models.ServiceKey
	resourceName := fmt.Sprintf("tf-cos-%d", acctest.RandIntRange(10, 100))
	resourceKey := fmt.Sprintf("tf-cos-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMResourceKeyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMResourceKeyWithTags(resourceName, resourceKey),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMResourceKeyExists("ibm_resource_key.resourceKey", conf),
					resource.TestCheckResourceAttr("ibm_resource_key.resourceKey", "name", resourceKey),
					resource.TestCheckResourceAttr("ibm_resource_key.resourceKey", "role", "Manager"),
					resource.TestCheckResourceAttr("ibm_resource_key.resourceKey", "tags.#", "1"),
				),
			},
			{
				Config: testAccCheckIBMResourceKeyWithUpdatedTags(resourceName, resourceKey),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMResourceKeyExists("ibm_resource_key.resourceKey", conf),
					resource.TestCheckResourceAttr("ibm_resource_key.resourceKey", "tags.#", "2"),
				),
			},
		},
	})
}

func TestAccIBMResourceKey_Parameters(t *testing.T) {
	var conf models.ServiceKey
	resourceName := fmt.Sprintf("tf-cos-%d", acctest.RandIntRange(10, 100))
	resourceKey := fmt.Sprintf("tf-cos-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMResourceKeyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMResourceKeyParameters(resourceName, resourceKey),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMResourceKeyExists("ibm_resource_key.resourceKey", conf),
					resource.TestCheckResourceAttr("ibm_resource_key.resourceKey", "name", resourceKey),
					resource.TestCheckResourceAttr("ibm_resource_key.resourceKey", "role", "Manager"),
					resource.TestCheckResourceAttrSet("ibm_resource_key.resourceKey", "credentials.%"),
				),
			},
		},
	})
}

func TestAccIBMResourceKeyWithCustomRole(t *testing.T) {
	var conf models.ServiceKey
	resourceName := fmt.Sprintf("tf-cos-%d", acctest.RandIntRange(10, 100))
	resourceKey := fmt.Sprintf("tf-cos-%d", acctest.RandIntRange(10, 100))
	crName := fmt.Sprintf("Name%d", acctest.RandIntRange(10, 100))
	displayName := fmt.Sprintf("Disp%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMResourceKeyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMResourceKeyWithCustomRole(resourceName, resourceKey, crName, displayName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMResourceKeyExists("ibm_resource_key.resourceKey", conf),
					resource.TestCheckResourceAttr("ibm_resource_key.resourceKey", "name", resourceKey),
					resource.TestCheckResourceAttr("ibm_resource_key.resourceKey", "credentials.%", "7"),
					resource.TestCheckResourceAttr("ibm_resource_key.resourceKey", "role", crName),
				),
			},
		},
	})
}

func testAccCheckIBMResourceKeyExists(n string, obj models.ServiceKey) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		rsContClient, err := testAccProvider.Meta().(ClientSession).ResourceControllerAPI()
		if err != nil {
			return err
		}
		resourceKeyID := rs.Primary.ID

		resourceKey, err := rsContClient.ResourceServiceKey().GetKey(resourceKeyID)
		if err != nil {
			return err
		}

		obj = resourceKey
		return nil
	}
}

func testAccCheckIBMResourceKeyDestroy(s *terraform.State) error {
	rsContClient, err := testAccProvider.Meta().(ClientSession).ResourceControllerAPI()
	if err != nil {
		return err
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_resource_key" {
			continue
		}

		resourceKeyID := rs.Primary.ID

		// Try to find the key
		key, err := rsContClient.ResourceServiceKey().GetKey(resourceKeyID)

		if err == nil {
			if key.State == "removed" {
				return nil
			}
			return fmt.Errorf("Resource key still exists: %s", rs.Primary.ID)
		} else if !strings.Contains(err.Error(), "404") {
			return fmt.Errorf("Error waiting for resource key (%s) to be destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}

func testAccCheckIBMResourceKeyBasic(resourceName, resourceKey string) string {
	return fmt.Sprintf(`
		
		resource "ibm_resource_instance" "resource" {
			name              = "%s"
			service           = "cloud-object-storage"
			plan              = "standard"
			location          = "global"
		}
		resource "ibm_resource_key" "resourceKey" {
			name = "%s"
			resource_instance_id = ibm_resource_instance.resource.id
			role = "Reader"
		}
	`, resourceName, resourceKey)
}

func testAccCheckIBMResourceKeyWithCustomRole(resourceName, resourceKey, crName, displayName string) string {
	return fmt.Sprintf(`
		
		resource "ibm_resource_instance" "resource" {
			name              = "%s"
			service           = "cloud-object-storage"
			plan              = "standard"
			location          = "global"
		}
		resource "ibm_iam_custom_role" "customrole" {
			name         = "%s"
			display_name = "%s"
			description  = "role for test scenario1"
			service = "cloud-object-storage"
			actions      = ["cloud-object-storage.bucket.get_cors"]
		}
		resource "ibm_resource_key" "resourceKey" {
			name = "%s"
			resource_instance_id = ibm_resource_instance.resource.id
			role = ibm_iam_custom_role.customrole.display_name
		}
	`, resourceName, crName, displayName, resourceKey)
}

func testAccCheckIBMResourceKeyWithTags(resourceName, resourceKey string) string {
	return fmt.Sprintf(`
		
		resource "ibm_resource_instance" "resource" {
			name              = "%s"
			service           = "cloud-object-storage"
			plan              = "standard"
			location          = "global"
		}
		resource "ibm_resource_key" "resourceKey" {
			name = "%s"
			resource_instance_id = ibm_resource_instance.resource.id
			role = "Manager"
			tags				  = ["one"]	
		}
	`, resourceName, resourceKey)
}

func testAccCheckIBMResourceKeyWithUpdatedTags(resourceName, resourceKey string) string {
	return fmt.Sprintf(`
		resource "ibm_resource_instance" "resource" {
			name              = "%s"
			service           = "cloud-object-storage"
			plan              = "standard"
			location          = "global"
		}
		resource "ibm_resource_key" "resourceKey" {
			name = "%s"
			resource_instance_id = ibm_resource_instance.resource.id
			role = "Manager"
			tags				  = ["one", "two"]	
		}
	`, resourceName, resourceKey)
}

func testAccCheckIBMResourceKeyParameters(resourceName, resourceKey string) string {
	return fmt.Sprintf(`
		
		resource "ibm_resource_instance" "resource" {
			name              = "%s"
			service           = "cloud-object-storage"
			plan              = "standard"
			location          = "global"
		}
		resource "ibm_resource_key" "resourceKey" {
			name = "%s"
			resource_instance_id = ibm_resource_instance.resource.id
			parameters        = {"HMAC" = true}
			role = "Manager"
		}
	`, resourceName, resourceKey)
}
