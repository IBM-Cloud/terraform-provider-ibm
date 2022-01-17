// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package cloudant_test

import (
	"fmt"
	"reflect"
	"strings"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/IBM-Cloud/bluemix-go/models"
)

func TestAccIBMCloudant_basic(t *testing.T) {
	var conf models.ServiceInstance
	resourceName := "ibm_cloudant.instance"
	serviceName := fmt.Sprintf("terraform-test-%s", acctest.RandString(8))
	updateName := fmt.Sprintf("terraform-test-%s", acctest.RandString(8))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCloudantDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCloudantResourceConfig(serviceName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCloudantExists(resourceName, conf),
					resource.TestCheckResourceAttr(resourceName, "name", serviceName),
					resource.TestCheckResourceAttr(resourceName, "service", "cloudantnosqldb"),
					resource.TestCheckResourceAttr(resourceName, "plan", "standard"),
					resource.TestCheckResourceAttr(resourceName, "legacy_credentials", "true"),
					resource.TestCheckResourceAttr(resourceName, "include_data_events", "true"),
					resource.TestCheckResourceAttr(resourceName, "capacity", "1"),
					resource.TestCheckResourceAttr(resourceName, "throughput.read", "100"),
					resource.TestCheckResourceAttr(resourceName, "enable_cors", "true"),
					resource.TestCheckResourceAttr(resourceName, "cors_config.0.allow_credentials", "false"),
				),
			},
			{
				Config: testAccCheckIBMCloudantResourceUpdateConfig(updateName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", updateName),
					resource.TestCheckResourceAttr(resourceName, "service", "cloudantnosqldb"),
					resource.TestCheckResourceAttr(resourceName, "plan", "standard"),
					resource.TestCheckResourceAttr(resourceName, "legacy_credentials", "true"),
					resource.TestCheckResourceAttr(resourceName, "include_data_events", "false"),
					resource.TestCheckResourceAttr(resourceName, "capacity", "2"),
					resource.TestCheckResourceAttr(resourceName, "throughput.read", "200"),
					resource.TestCheckResourceAttr(resourceName, "enable_cors", "false"),
				),
			},
		},
	})
}

func TestAccIBMCloudant_import(t *testing.T) {
	var conf models.ServiceInstance
	resourceName := "ibm_cloudant.instance"
	serviceName := fmt.Sprintf("terraform-test-%s", acctest.RandString(8))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCloudantDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCloudantResourceConfigLite(serviceName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCloudantExists(resourceName, conf),
					resource.TestCheckResourceAttr(resourceName, "name", serviceName),
					resource.TestCheckResourceAttr(resourceName, "service", "cloudantnosqldb"),
					resource.TestCheckResourceAttr(resourceName, "plan", "lite"),
					resource.TestCheckResourceAttr(resourceName, "legacy_credentials", "false"),
					resource.TestCheckResourceAttr(resourceName, "capacity", "1"),
					resource.TestCheckResourceAttr(resourceName, "throughput.read", "20"),
					resource.TestCheckResourceAttr(resourceName, "enable_cors", "true"),
					resource.TestCheckResourceAttr(resourceName, "cors_config.0.allow_credentials", "true"),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"parameters"},
			},
		},
	})
}

func testAccCheckIBMCloudantDestroy(s *terraform.State) error {
	rsContClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).ResourceControllerAPI()
	if err != nil {
		return err
	}
	for rsName, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_cloudant" || !strings.HasPrefix(rsName, "terraform-test-") {
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
				return fmt.Errorf("[ERROR] Error checking if Resource Instance (%s) has been destroyed: %s", rs.Primary.ID, err)
			}
		}
	}

	return nil
}

func testAccCheckIBMCloudantExists(resourceName string, obj models.ServiceInstance) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("Not found: %s", resourceName)
		}

		rsContClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).ResourceControllerAPI()
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

func testAccCheckIBMCloudantResourceConfig(serviceName string) string {
	return fmt.Sprintf(`

	resource "ibm_cloudant" "instance" {
		name                = "%s"
		plan                = "standard"
		location            = "us-south"

		legacy_credentials  = true
		include_data_events = true
		capacity            = 1
		enable_cors         = true

		cors_config {
			allow_credentials = false
			origins           = ["https://example.com"]
		}

		timeouts {
		  create = "15m"
		  update = "15m"
		  delete = "15m"
		}
	  }

	`, serviceName)
}

func testAccCheckIBMCloudantResourceUpdateConfig(serviceName string) string {
	return fmt.Sprintf(`

	resource "ibm_cloudant" "instance" {
		name                = "%s"
		plan                = "standard"
		location            = "us-south"

		legacy_credentials  = true
		include_data_events = false
		capacity            = 2
		enable_cors         = false

		timeouts {
		  create = "15m"
		  update = "15m"
		  delete = "15m"
		}
	  }

	`, serviceName)
}

func testAccCheckIBMCloudantResourceConfigLite(serviceName string) string {
	return fmt.Sprintf(`

	resource "ibm_cloudant" "instance" {
		name               = "%s"
		plan               = "lite"
		location           = "us-south"

		timeouts {
		  create = "15m"
		  update = "15m"
		  delete = "15m"
		}
	  }

	`, serviceName)
}
