// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package db2_test

import (
	"fmt"
	"strings"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"

	rc "github.com/IBM/platform-services-go-sdk/resourcecontrollerv2"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccIBMDb2InstanceBasic(t *testing.T) {
	databaseResourceGroup := "Default"
	var databaseInstanceOne string
	rnd := fmt.Sprintf("tf-db2-%d", acctest.RandIntRange(10, 100))
	testName := rnd
	name := "ibm_db2." + testName

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMDb2InstanceDestroy,
		Steps: []resource.TestStep{
			{

				Config: testAccCheckIBMDb2InstanceBasic(databaseResourceGroup, testName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMDb2InstanceExists(name, &databaseInstanceOne),
					resource.TestCheckResourceAttr(name, "name", testName),
					resource.TestCheckResourceAttr(name, "service", "dashdb-for-transactions"),
					resource.TestCheckResourceAttr(name, "plan", "performance"),
					resource.TestCheckResourceAttr(name, "location", "us-east"),
					resource.TestCheckResourceAttr(name, "service_endpoints", "public-and-private"),
				),
			},
			{

				Config: testAccCheckIBMDb2InstanceFullyspecified(databaseResourceGroup, testName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMDb2InstanceExists(name, &databaseInstanceOne),
					resource.TestCheckResourceAttr(name, "name", testName),
					resource.TestCheckResourceAttr(name, "service", "dashdb-for-transactions"),
					resource.TestCheckResourceAttr(name, "plan", "performance"),
					resource.TestCheckResourceAttr(name, "location", "us-east"),
					resource.TestCheckResourceAttr(name, "service_endpoints", "public-and-private"),
					resource.TestCheckResourceAttr(name, "instance_type", "bx2.4x16"),
					resource.TestCheckResourceAttr(name, "high_availability", "no"),
					resource.TestCheckResourceAttr(name, "backup_location", "us"),
					resource.TestCheckResourceAttr(name, "tags.#", "1"),
				),
			},
		},
	})
}

func testAccCheckIBMDb2InstanceDestroy(s *terraform.State) error {
	rsContClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).ResourceControllerV2API()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_db2" {
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

func testAccCheckIBMDb2InstanceExists(n string, tfDb2ID *string) resource.TestCheckFunc {

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

		instance, response, err := rsContClient.GetResourceInstance(&resourceInstanceGet)

		if err != nil {
			if strings.Contains(err.Error(), "Object not found") ||
				strings.Contains(err.Error(), "status code: 404") {
				*tfDb2ID = ""
				return nil
			}
			return fmt.Errorf("[ERROR] Error retrieving resource instance: %s %s", err, response)
		}
		if strings.Contains(*instance.State, "removed") {
			*tfDb2ID = ""
			return nil
		}

		*tfDb2ID = instanceID

		return nil
	}
}

func testAccCheckIBMDb2InstanceBasic(databaseResourceGroup string, testName string) string {
	return fmt.Sprintf(`
	
    data "ibm_resource_group" "group" {
		name = "%[1]s"
	}

	resource "ibm_db2" "%[2]s" {
		name              = "%[2]s"
		service           = "dashdb-for-transactions"
		plan              = "performance" 
		location          = "us-east"
		resource_group_id = data.ibm_resource_group.group.id
		service_endpoints = "public-and-private"

		timeouts {
			create = "720m"
			update = "30m"
			delete = "30m"
		}
	}
	`, databaseResourceGroup, testName)
}

func testAccCheckIBMDb2InstanceFullyspecified(databaseResourceGroup string, testName string) string {
	return fmt.Sprintf(`
	
    data "ibm_resource_group" "group" {
		name = "%[1]s"
	}

	resource "ibm_db2" "%[2]s" {
		name              = "%[2]s"
		service           = "dashdb-for-transactions"
		plan              = "performance" 
		location          = "us-east"
		resource_group_id = data.ibm_resource_group.group.id
		service_endpoints = "public-and-private"
		instance_type     = "bx2.4x16"
		high_availability = "no"
		backup_location   = "us"
		tags              = ["one:two"]
		disk_encryption_instance_crn = "none"
		disk_encryption_key_crn = "none"
		oracle_compatibility = "no"

		timeouts {
			create = "720m"
			update = "30m"
			delete = "30m"
		}
	}
	`, databaseResourceGroup, testName)
}

func TestAccIBMDb2InstanceAutoscale(t *testing.T) {
	databaseResourceGroup := "Default"
	var databaseInstanceOne string
	rnd := fmt.Sprintf("tf-db2-%d", acctest.RandIntRange(10, 100))
	testName := rnd
	name := "ibm_db2." + testName

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMDb2InstanceDestroy,
		Steps: []resource.TestStep{
			{

				Config: testAccCheckIBMDb2InstanceAutoscale(databaseResourceGroup, testName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMDb2InstanceExists(name, &databaseInstanceOne),
					resource.TestCheckResourceAttr(name, "name", testName),
					resource.TestCheckResourceAttr(name, "service", "dashdb-for-transactions"),
					resource.TestCheckResourceAttr(name, "plan", "performance"),
					resource.TestCheckResourceAttr(name, "location", "us-south"),
					resource.TestCheckResourceAttr(name, "service_endpoints", "public-and-private"),
					resource.TestCheckResourceAttr(name, "autoscale_config.#", "1"),
				),
			},
		},
	})
}

func testAccCheckIBMDb2InstanceAutoscale(databaseResourceGroup string, testName string) string {
	return fmt.Sprintf(`
	
    data "ibm_resource_group" "group" {
		name = "%[1]s"
	}

	resource "ibm_db2" "%[2]s" {
		name              = "%[2]s"
		service           = "dashdb-for-transactions"
		plan              = "performance" 
		location          = "us-south"
		resource_group_id = data.ibm_resource_group.group.id
		service_endpoints = "public-and-private"

		timeouts {
			create = "720m"
			update = "30m"
			delete = "30m"
		}

		autoscale_config  {
    		auto_scaling_enabled = "true"
    		auto_scaling_threshold = "60"
    		auto_scaling_over_time_period = "15"
    		auto_scaling_pause_limit = "70"
   		}
	}
	`, databaseResourceGroup, testName)
}

func TestAccIBMDb2InstanceCustomSetting(t *testing.T) {
	databaseResourceGroup := "Default"
	var databaseInstanceOne string
	rnd := fmt.Sprintf("tf-db2-%d", acctest.RandIntRange(10, 100))
	testName := rnd
	name := "ibm_db2." + testName

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMDb2InstanceDestroy,
		Steps: []resource.TestStep{
			{

				Config: testAccCheckIBMDb2InstanceCustomSetting(databaseResourceGroup, testName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMDb2InstanceExists(name, &databaseInstanceOne),
					resource.TestCheckResourceAttr(name, "name", testName),
					resource.TestCheckResourceAttr(name, "service", "dashdb-for-transactions"),
					resource.TestCheckResourceAttr(name, "plan", "performance"),
					resource.TestCheckResourceAttr(name, "location", "us-south"),
					resource.TestCheckResourceAttr(name, "service_endpoints", "public-and-private"),
					resource.TestCheckResourceAttr(name, "custom_setting_config.#", "1"),
				),
			},
		},
	})
}

func testAccCheckIBMDb2InstanceCustomSetting(databaseResourceGroup string, testName string) string {
	return fmt.Sprintf(`
	
    data "ibm_resource_group" "group" {
		name = "%[1]s"
	}

	resource "ibm_db2" "%[2]s" {
		name              = "%[2]s"
		service           = "dashdb-for-transactions"
		plan              = "performance" 
		location          = "us-south"
		resource_group_id = data.ibm_resource_group.group.id
		service_endpoints = "public-and-private"

		timeouts {
			create = "720m"
			update = "30m"
			delete = "30m"
		}

		custom_setting_config {
    		db {
      			auto_reval = "IMMEDIATE"
    		}
    		dbm {
      			multipartsizemb = "100"
    		}
    		registry {
      		db2_alternate_authz_behaviour = "EXTERNAL_ROUTINE_DBADM"
    	}
  }
	}
	`, databaseResourceGroup, testName)
}

func TestAccIBMDb2InstanceCreateAllowlist(t *testing.T) {
	databaseResourceGroup := "Default"
	var databaseInstanceOne string
	rnd := fmt.Sprintf("tf-db2-%d", acctest.RandIntRange(10, 100))
	testName := rnd
	name := "ibm_db2." + testName

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMDb2InstanceDestroy,
		Steps: []resource.TestStep{
			{

				Config: testAccCheckIBMDb2InstanceCreateAllowlist(databaseResourceGroup, testName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMDb2InstanceExists(name, &databaseInstanceOne),
					resource.TestCheckResourceAttr(name, "name", testName),
					resource.TestCheckResourceAttr(name, "service", "dashdb-for-transactions"),
					resource.TestCheckResourceAttr(name, "plan", "performance-dev"),
					resource.TestCheckResourceAttr(name, "location", "us-east"),
					resource.TestCheckResourceAttr(name, "service_endpoints", "public-and-private"),
					resource.TestCheckResourceAttr(name, "allowlist_config.#", "1"),
				),
			},
		},
	})
}

func testAccCheckIBMDb2InstanceCreateAllowlist(databaseResourceGroup string, testName string) string {
	return fmt.Sprintf(`
	
    data "ibm_resource_group" "group" {
		name = "%[1]s"
	}
	resource "ibm_db2" "%[2]s" {
		name              = "%[2]s"
		service           = "dashdb-for-transactions"
		plan              = "performance-dev" 
		location          = "us-east"
		resource_group_id = data.ibm_resource_group.group.id
		service_endpoints = "public-and-private"
		timeouts {
			create = "720m"
			update = "30m"
			delete = "30m"
		}
		allowlist_config {
		ip_addresses {
			address     = "127.0.0.5"
			description = "A sample IP address 5"
			}
		}
	}
	`, databaseResourceGroup, testName)
}
