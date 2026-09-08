// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package cloudant_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMCloudantGen2_basic(t *testing.T) {
	resourceName := "ibm_cloudant.instance"
	serviceName := fmt.Sprintf("terraform-test-%s", acctest.RandString(8))
	updateName := fmt.Sprintf("terraform-test-%s", acctest.RandString(8))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCloudantDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCloudantGen2ResourceConfig(serviceName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCloudantExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", serviceName),
					resource.TestCheckResourceAttr(resourceName, "service", "cloudantnosqldb"),
					resource.TestCheckResourceAttr(resourceName, "plan", "standard-gen2"),
					resource.TestCheckResourceAttr(resourceName, "include_data_events", "true"),
					resource.TestCheckResourceAttr(resourceName, "capacity", "1"),
					resource.TestCheckResourceAttr(resourceName, "enable_cors", "true"),
					resource.TestCheckResourceAttr(resourceName, "cors_config.0.allow_credentials", "false"),
				),
			},
			{
				Config: testAccCheckIBMCloudantGen2ResourceUpdateConfig(updateName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", updateName),
					resource.TestCheckResourceAttr(resourceName, "service", "cloudantnosqldb"),
					resource.TestCheckResourceAttr(resourceName, "plan", "standard-gen2"),
					resource.TestCheckResourceAttr(resourceName, "include_data_events", "false"),
					resource.TestCheckResourceAttr(resourceName, "capacity", "2"),
					resource.TestCheckResourceAttr(resourceName, "enable_cors", "false"),
				),
			},
		},
	})
}

func TestAccIBMCloudantGen2_import(t *testing.T) {
	resourceName := "ibm_cloudant.instance"
	serviceName := fmt.Sprintf("terraform-test-%s", acctest.RandString(8))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCloudantDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCloudantGen2ResourceConfigMinimal(serviceName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCloudantExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", serviceName),
					resource.TestCheckResourceAttr(resourceName, "service", "cloudantnosqldb"),
					resource.TestCheckResourceAttr(resourceName, "plan", "standard-gen2"),
					resource.TestCheckResourceAttr(resourceName, "capacity", "1"),
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

func testAccCheckIBMCloudantGen2ResourceConfig(serviceName string) string {
	return fmt.Sprintf(`

	resource "ibm_cloudant" "instance" {
		name                = "%s"
		plan                = "standard-gen2"
		location            = "%s"

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

	`, serviceName, acc.Region())
}

func testAccCheckIBMCloudantGen2ResourceUpdateConfig(serviceName string) string {
	return fmt.Sprintf(`

	resource "ibm_cloudant" "instance" {
		name                = "%s"
		plan                = "standard-gen2"
		location            = "%s"

		include_data_events = false
		capacity            = 2
		enable_cors         = false

		timeouts {
		  create = "15m"
		  update = "15m"
		  delete = "15m"
		}
	  }

	`, serviceName, acc.Region())
}

func testAccCheckIBMCloudantGen2ResourceConfigMinimal(serviceName string) string {
	return fmt.Sprintf(`

	resource "ibm_cloudant" "instance" {
		name     = "%s"
		plan     = "standard-gen2"
		location = "%s"

		timeouts {
		  create = "15m"
		  update = "15m"
		  delete = "15m"
		}
	  }

	`, serviceName, acc.Region())
}
