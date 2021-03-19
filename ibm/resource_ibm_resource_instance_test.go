// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/IBM-Cloud/bluemix-go/models"
)

func TestAccIBMResourceInstance_Basic(t *testing.T) {
	var conf models.ServiceInstance
	serviceName := fmt.Sprintf("tf-cos-%d", acctest.RandIntRange(10, 100))
	updateName := fmt.Sprintf("tf-kms-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMResourceInstanceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMResourceInstanceBasic(serviceName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMResourceInstanceExists("ibm_resource_instance.instance", conf),
					resource.TestCheckResourceAttr("ibm_resource_instance.instance", "name", serviceName),
					resource.TestCheckResourceAttr("ibm_resource_instance.instance", "service", "cloud-object-storage"),
					resource.TestCheckResourceAttr("ibm_resource_instance.instance", "plan", "standard"),
					resource.TestCheckResourceAttr("ibm_resource_instance.instance", "location", "global"),
				),
			},
			{
				Config: testAccCheckIBMResourceInstanceUpdateWithSameName(serviceName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMResourceInstanceExists("ibm_resource_instance.instance", conf),
					resource.TestCheckResourceAttr("ibm_resource_instance.instance", "name", serviceName),
					resource.TestCheckResourceAttr("ibm_resource_instance.instance", "service", "cloud-object-storage"),
					resource.TestCheckResourceAttr("ibm_resource_instance.instance", "plan", "standard"),
					resource.TestCheckResourceAttr("ibm_resource_instance.instance", "location", "global"),
				),
			},
			{
				Config: testAccCheckIBMResourceInstanceUpdate(updateName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_resource_instance.instance", "name", updateName),
					resource.TestCheckResourceAttr("ibm_resource_instance.instance", "service", "cloud-object-storage"),
					resource.TestCheckResourceAttr("ibm_resource_instance.instance", "plan", "standard"),
					resource.TestCheckResourceAttr("ibm_resource_instance.instance", "location", "global"),
				),
			},
			{
				Config: testAccCheckIBMResourceInstanceNewServiceType(updateName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_resource_instance.instance", "name", updateName),
					resource.TestCheckResourceAttr("ibm_resource_instance.instance", "service", "kms"),
					resource.TestCheckResourceAttr("ibm_resource_instance.instance", "plan", "tiered-pricing"),
					resource.TestCheckResourceAttr("ibm_resource_instance.instance", "location", "us-south"),
				),
			},
		},
	})
}

func TestAccIBMResourceInstance_import(t *testing.T) {
	var conf models.ServiceInstance
	serviceName := fmt.Sprintf("tf-ins-%d", acctest.RandIntRange(10, 100))
	resourceName := "ibm_resource_instance.instance"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMResourceInstanceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMResourceInstanceBasic(serviceName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMResourceInstanceExists(resourceName, conf),
					resource.TestCheckResourceAttr(resourceName, "name", serviceName),
					resource.TestCheckResourceAttr(resourceName, "service", "cloud-object-storage"),
					resource.TestCheckResourceAttr(resourceName, "plan", "standard"),
					resource.TestCheckResourceAttr(resourceName, "location", "global"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"wait_time_minutes", "parameters"},
			},
		},
	})
}

func TestAccIBMResourceInstance_with_serviceendpoints(t *testing.T) {
	var conf models.ServiceInstance
	serviceName := fmt.Sprintf("tf-Pgress-%d", acctest.RandIntRange(10, 100))
	resourceName := "ibm_resource_instance.instance"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMResourceInstanceDestroy,
		Steps: []resource.TestStep{

			{
				Config: testAccCheckIBMResourceInstanceServiceendpoints(serviceName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMResourceInstanceExists(resourceName, conf),
					resource.TestCheckResourceAttr(resourceName, "name", serviceName),
					resource.TestCheckResourceAttr(resourceName, "service", "databases-for-postgresql"),
					resource.TestCheckResourceAttr(resourceName, "plan", "standard"),
					resource.TestCheckResourceAttr(resourceName, "location", "us-south"),
				),
			},
		},
	})
}

func TestAccIBMResourceInstance_with_resource_group(t *testing.T) {
	var conf models.ServiceInstance
	serviceName := fmt.Sprintf("tf-cos-%d", acctest.RandIntRange(10, 100))
	resourceName := "ibm_resource_instance.instance"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMResourceInstanceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMResourceInstanceWithResourceGroup(serviceName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMResourceInstanceExists(resourceName, conf),
					resource.TestCheckResourceAttr(resourceName, "name", serviceName),
					resource.TestCheckResourceAttr(resourceName, "service", "cloud-object-storage"),
					resource.TestCheckResourceAttr(resourceName, "plan", "standard"),
					resource.TestCheckResourceAttr(resourceName, "location", "global"),
				),
			},
		},
	})
}

func testAccCheckIBMResourceInstanceDestroy(s *terraform.State) error {
	rsContClient, err := testAccProvider.Meta().(ClientSession).ResourceControllerAPI()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_resource_instance" {
			continue
		}

		instanceID := rs.Primary.ID

		// Try to find the key
		instance, err := rsContClient.ResourceServiceInstance().GetInstance(instanceID)

		if err == nil {
			if !reflect.DeepEqual(instance, models.ServiceInstance{}) && instance.State == "active" {
				return fmt.Errorf("Resource Instance still exists: %s", rs.Primary.ID)
			}
		} else {
			if !strings.Contains(err.Error(), "404") {
				return fmt.Errorf("Error checking if Resource Instance (%s) has been destroyed: %s", rs.Primary.ID, err)
			}
		}
	}

	return nil
}

func testAccCheckIBMResourceInstanceExists(n string, obj models.ServiceInstance) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		rsContClient, err := testAccProvider.Meta().(ClientSession).ResourceControllerAPI()
		if err != nil {
			return err
		}
		instanceID := rs.Primary.ID

		instance, err := rsContClient.ResourceServiceInstance().GetInstance(instanceID)

		if err != nil {
			return err
		}

		obj = instance
		return nil
	}
}

func testAccCheckIBMResourceInstanceBasic(serviceName string) string {
	return fmt.Sprintf(`
		
	resource "ibm_resource_instance" "instance" {
		name     = "%s"
		service  = "cloud-object-storage"
		plan     = "standard"
		location = "global"
		parameters = {
		  "HMAC" = true
		}
	  
		timeouts {
		  create = "15m"
		  update = "15m"
		  delete = "15m"
		}
	  }
	`, serviceName)
}

func testAccCheckIBMResourceInstanceUpdateWithSameName(serviceName string) string {
	return fmt.Sprintf(`
		
	resource "ibm_resource_instance" "instance" {
		name     = "%s"
		service  = "cloud-object-storage"
		plan     = "standard"
		location = "global"
		parameters = {
		  "HMAC" = true
		}
	}
	  
	`, serviceName)
}

func testAccCheckIBMResourceInstanceUpdate(updateName string) string {
	return fmt.Sprintf(`

	resource "ibm_resource_instance" "instance" {
		name     = "%s"
		service  = "cloud-object-storage"
		plan     = "standard"
		location = "global"
		parameters = {
		  "HMAC" = true
		}
	}
	`, updateName)
}

func testAccCheckIBMResourceInstanceNewServiceType(updateName string) string {
	return fmt.Sprintf(`
	resource "ibm_resource_instance" "instance" {
		name     = "%s"
		service  = "kms"
		plan     = "tiered-pricing"
		location = "us-south"
	}
	`, updateName)
}

func testAccCheckIBMResourceInstanceWithResourceGroup(serviceName string) string {
	return fmt.Sprintf(`

	data "ibm_resource_group" "group" {
		is_default=true
	  }
	  
	resource "ibm_resource_instance" "instance" {
		name              = "%s"
		service           = "cloud-object-storage"
		plan              = "standard"
		location          = "global"
		resource_group_id = data.ibm_resource_group.group.id
		parameters = {
		  "HMAC" = true
		}
	}
	`, serviceName)
}

func testAccCheckIBMResourceInstanceServiceendpoints(serviceName string) string {
	return fmt.Sprintf(`
	
	resource "ibm_resource_instance" "instance" {
		name     = "%s"
		location = "us-south"
		service  = "databases-for-postgresql"
		plan     = "standard"
		parameters = {
		  members_memory_allocation_mb = "4096"
		}
	  
		//service_endpoints = "public-and-private"
		timeouts {
		  create = "25m"
		  update = "15m"
		  delete = "15m"
		}
	}
			
	`, serviceName)
}
