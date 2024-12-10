// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package resourcecontroller_test

import (
	"fmt"
	"regexp"
	"strings"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"

	rc "github.com/IBM/platform-services-go-sdk/resourcecontrollerv2"
	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccIBMResourceInstanceBasic(t *testing.T) {
	serviceName := fmt.Sprintf("tf-cos-%d", acctest.RandIntRange(10, 100))
	updateName := fmt.Sprintf("tf-kms-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMResourceInstanceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMResourceInstanceBasic(serviceName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMResourceInstanceExists("ibm_resource_instance.instance"),
					resource.TestCheckResourceAttr("ibm_resource_instance.instance", "name", serviceName),
					resource.TestCheckResourceAttr("ibm_resource_instance.instance", "service", "cloud-object-storage"),
					resource.TestCheckResourceAttr("ibm_resource_instance.instance", "plan", "standard"),
					resource.TestCheckResourceAttr("ibm_resource_instance.instance", "location", "global"),
				),
			},
			{
				Config: testAccCheckIBMResourceInstanceUpdateWithSameName(serviceName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMResourceInstanceExists("ibm_resource_instance.instance"),
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

func TestAccIBMResourceInstanceImport(t *testing.T) {
	serviceName := fmt.Sprintf("tf-ins-%d", acctest.RandIntRange(10, 100))
	resourceName := "ibm_resource_instance.instance"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMResourceInstanceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMResourceInstanceBasic(serviceName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMResourceInstanceExists(resourceName),
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

func TestAccIBMResourceInstanceWithServiceendpoints(t *testing.T) {
	serviceName := fmt.Sprintf("tf-Pgress-%d", acctest.RandIntRange(10, 100))
	resourceName := "ibm_resource_instance.instance"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMResourceInstanceDestroy,
		Steps: []resource.TestStep{

			{
				Config: testAccCheckIBMResourceInstanceServiceendpoints(serviceName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMResourceInstanceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", serviceName),
					resource.TestCheckResourceAttr(resourceName, "service", "databases-for-postgresql"),
					resource.TestCheckResourceAttr(resourceName, "plan", "standard"),
					resource.TestCheckResourceAttr(resourceName, "location", "us-south"),
				),
			},
		},
	})
}

func TestAccIBMResourceInstanceWithResourceGroup(t *testing.T) {
	serviceName := fmt.Sprintf("tf-cos-%d", acctest.RandIntRange(10, 100))
	resourceName := "ibm_resource_instance.instance"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMResourceInstanceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMResourceInstanceWithResourceGroup(serviceName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMResourceInstanceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", serviceName),
					resource.TestCheckResourceAttr(resourceName, "service", "cloud-object-storage"),
					resource.TestCheckResourceAttr(resourceName, "plan", "standard"),
					resource.TestCheckResourceAttr(resourceName, "location", "global"),
				),
			},
		},
	})
}

func TestAccIBMCOSResourceInstanceOneRatePlan(t *testing.T) {
	serviceName := fmt.Sprintf("tf-cos-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMResourceInstanceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCOSResourceInstanceOneRatePlan(serviceName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMResourceInstanceExists("ibm_resource_instance.instance"),
					resource.TestCheckResourceAttr("ibm_resource_instance.instance", "name", serviceName),
					resource.TestCheckResourceAttr("ibm_resource_instance.instance", "service", "cloud-object-storage"),
					resource.TestCheckResourceAttr("ibm_resource_instance.instance", "plan", "cos-one-rate-plan"),
					resource.TestCheckResourceAttr("ibm_resource_instance.instance", "location", "global"),
				),
			},
		},
	})
}

func TestAccIBMCOSResourceInstancewithoutplan(t *testing.T) {
	serviceName := fmt.Sprintf("tf-cos-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMResourceInstanceDestroy,
		Steps: []resource.TestStep{
			{
				Config:      testAccCheckIBMCOSResourceInstancewithoutplan(serviceName),
				ExpectError: regexp.MustCompile("Missing required argument"),
			},
		},
	})
}

func TestAccIBMCOSResourceInstanceWithInvalidPlan(t *testing.T) {
	serviceName := fmt.Sprintf("tf-cos-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMResourceInstanceDestroy,
		Steps: []resource.TestStep{
			{
				Config:      testAccCheckIBMCOSResourceInstanceWithInvalidPlan(serviceName),
				ExpectError: regexp.MustCompile("Error retrieving deployment for plan invalidcosplan"),
			},
		},
	})
}

func testAccCheckIBMResourceInstanceDestroy(s *terraform.State) error {
	rsContClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).ResourceControllerV2API()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_resource_instance" {
			continue
		}

		instanceID := rs.Primary.ID
		resourceInstanceGet := rc.GetResourceInstanceOptions{
			ID: &instanceID,
		}

		// Try to find the key
		instance, resp, err := rsContClient.GetResourceInstance(&resourceInstanceGet)

		if err == nil {
			if *instance.State == "active" {
				return fmt.Errorf("Resource Instance still exists: %s", rs.Primary.ID)
			}
		} else {
			if !strings.Contains(err.Error(), "404") {
				return fmt.Errorf("[ERROR] Error checking if Resource Instance (%s) has been destroyed: %s with resp code: %s", rs.Primary.ID, err, resp)
			}
		}
	}

	return nil
}

func testAccCheckIBMResourceInstanceExists(n string) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		rsContClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).ResourceControllerV2API()
		if err != nil {
			return err
		}
		instanceID := rs.Primary.ID
		resourceInstanceGet := rc.GetResourceInstanceOptions{
			ID: &instanceID,
		}

		_, resp, err := rsContClient.GetResourceInstance(&resourceInstanceGet)

		if err != nil {
			return fmt.Errorf("Get resource instance error: %s with resp code: %s", err, resp)
		}

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
func testAccCheckIBMCOSResourceInstanceOneRatePlan(serviceName string) string {
	return fmt.Sprintf(`
		
	resource "ibm_resource_instance" "instance" {
		name     = "%s"
		service  = "cloud-object-storage"
		plan     = "cos-one-rate-plan"
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

func testAccCheckIBMCOSResourceInstancewithoutplan(serviceName string) string {
	return fmt.Sprintf(`

	resource "ibm_resource_instance" "instance" {
		name     = "%s"
		service  = "cloud-object-storage"
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
func testAccCheckIBMCOSResourceInstanceWithInvalidPlan(serviceName string) string {
	return fmt.Sprintf(`

	resource "ibm_resource_instance" "instance" {
		name     = "%s"
		service  = "cloud-object-storage"
		plan     = "invalidcosplan"
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

// /#### Adding new test case for onetime_credentials
func TestAccIBMResourceInstanceWithOnetimeCredentials(t *testing.T) {
	serviceName := fmt.Sprintf("tf-Pgress-%d", acctest.RandIntRange(10, 100))
	resourceName := "ibm_resource_instance.instance"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMResourceInstanceDestroy,
		Steps: []resource.TestStep{

			{
				Config: testAccCheckIBMResourceInstanceOnetimeCredentals(serviceName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMResourceInstanceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", serviceName),
					resource.TestCheckResourceAttr(resourceName, "service", "cloud-object-storage"),
					resource.TestCheckResourceAttr(resourceName, "plan", "lite"),
					resource.TestCheckResourceAttr(resourceName, "location", "global"),
				),
			},
		},
	})
}

func testAccCheckIBMResourceInstanceOnetimeCredentals(serviceName string) string {
	return fmt.Sprintf(`
    
    resource "ibm_resource_instance" "instance" {
        name     = "%s"
        location = "global"
        service  = "cloud-object-storage"
        plan     = "lite"
        parameters = {
          onetime_credentials = true,
        }
      
        
    }
            
    `, serviceName)
}
