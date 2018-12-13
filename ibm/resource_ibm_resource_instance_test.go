package ibm

import (
	"fmt"
	"testing"

	"github.com/IBM-Cloud/bluemix-go/models"

	"strings"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccIBMResourceInstance_Basic(t *testing.T) {
	var conf models.ServiceInstance
	serviceName := fmt.Sprintf("terraform_%d", acctest.RandInt())
	updateName := fmt.Sprintf("terraform_%d", acctest.RandInt())

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMResourceInstanceDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMResourceInstance_basic(serviceName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMResourceInstanceExists("ibm_resource_instance.instance", conf),
					resource.TestCheckResourceAttr("ibm_resource_instance.instance", "name", serviceName),
					resource.TestCheckResourceAttr("ibm_resource_instance.instance", "service", "cloud-object-storage"),
					resource.TestCheckResourceAttr("ibm_resource_instance.instance", "plan", "lite"),
					resource.TestCheckResourceAttr("ibm_resource_instance.instance", "location", "global"),
					resource.TestCheckResourceAttr("ibm_resource_instance.instance", "parameters.%", "1"),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMResourceInstance_updateWithSameName(serviceName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMResourceInstanceExists("ibm_resource_instance.instance", conf),
					resource.TestCheckResourceAttr("ibm_resource_instance.instance", "name", serviceName),
					resource.TestCheckResourceAttr("ibm_resource_instance.instance", "service", "cloud-object-storage"),
					resource.TestCheckResourceAttr("ibm_resource_instance.instance", "plan", "lite"),
					resource.TestCheckResourceAttr("ibm_resource_instance.instance", "location", "global"),
					resource.TestCheckResourceAttr("ibm_resource_instance.instance", "parameters.%", "1"),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMResourceInstance_update(updateName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_resource_instance.instance", "name", updateName),
					resource.TestCheckResourceAttr("ibm_resource_instance.instance", "service", "cloud-object-storage"),
					resource.TestCheckResourceAttr("ibm_resource_instance.instance", "plan", "lite"),
					resource.TestCheckResourceAttr("ibm_resource_instance.instance", "location", "global"),
					resource.TestCheckResourceAttr("ibm_resource_instance.instance", "parameters.%", "1"),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMResourceInstance_newServiceType(updateName),
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
	serviceName := fmt.Sprintf("terraform_%d", acctest.RandInt())
	resourceName := "ibm_resource_instance.instance"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMResourceInstanceDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMResourceInstance_basic(serviceName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMResourceInstanceExists(resourceName, conf),
					resource.TestCheckResourceAttr(resourceName, "name", serviceName),
					resource.TestCheckResourceAttr(resourceName, "service", "cloud-object-storage"),
					resource.TestCheckResourceAttr(resourceName, "plan", "lite"),
					resource.TestCheckResourceAttr(resourceName, "location", "global"),
					resource.TestCheckResourceAttr(resourceName, "parameters.%", "1"),
				),
			},
			resource.TestStep{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"wait_time_minutes"},
			},
		},
	})
}

func TestAccIBMResourceInstance_with_resource_group(t *testing.T) {
	var conf models.ServiceInstance
	serviceName := fmt.Sprintf("terraform_%d", acctest.RandInt())
	resourceName := "ibm_resource_instance.instance"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMResourceInstanceDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMResourceInstance_with_resource_group(serviceName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMResourceInstanceExists(resourceName, conf),
					resource.TestCheckResourceAttr(resourceName, "name", serviceName),
					resource.TestCheckResourceAttr(resourceName, "service", "cloud-object-storage"),
					resource.TestCheckResourceAttr(resourceName, "plan", "lite"),
					resource.TestCheckResourceAttr(resourceName, "location", "global"),
					resource.TestCheckResourceAttr(resourceName, "parameters.%", "1"),
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
		_, err := rsContClient.ResourceServiceInstance().GetInstance(instanceID)

		if err != nil && !strings.Contains(err.Error(), "404") {
			return fmt.Errorf("Error waiting for resource instance (%s) to be destroyed: %s", rs.Primary.ID, err)
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

func testAccCheckIBMResourceInstance_basic(serviceName string) string {
	return fmt.Sprintf(`
		
		resource "ibm_resource_instance" "instance" {
			name              = "%s"		
			service           = "cloud-object-storage"
			plan              = "lite"
			location          = "global"
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

func testAccCheckIBMResourceInstance_updateWithSameName(serviceName string) string {
	return fmt.Sprintf(`
		
		resource "ibm_resource_instance" "instance" {
			name              = "%s"
			service           = "cloud-object-storage"
			plan              = "lite"
			location          = "global"
			parameters = {
				"HMAC" = true
			  }
		}
	`, serviceName)
}

func testAccCheckIBMResourceInstance_update(updateName string) string {
	return fmt.Sprintf(`

		resource "ibm_resource_instance" "instance" {
			name              = "%s"		
			service           = "cloud-object-storage"
			plan              = "lite"
			location          = "global"
			parameters = {
				"HMAC" = true
			  }
		}
	`, updateName)
}

func testAccCheckIBMResourceInstance_newServiceType(updateName string) string {
	return fmt.Sprintf(`
		resource "ibm_resource_instance" "instance" {
			name              = "%s"		
			service           = "kms"
			plan              = "tiered-pricing"
			location          = "us-south"
		}
	`, updateName)
}

func testAccCheckIBMResourceInstance_with_resource_group(serviceName string) string {
	return fmt.Sprintf(`

		data "ibm_resource_group" "group" {
			name = "default"
		}
		
		resource "ibm_resource_instance" "instance" {
			name              = "%s"		
			service           = "cloud-object-storage"
			plan              = "lite"
			location          = "global"
			resource_group_id = "${data.ibm_resource_group.group.id}"
			parameters = {
				"HMAC" = true
			  }
		}
	`, serviceName)
}
