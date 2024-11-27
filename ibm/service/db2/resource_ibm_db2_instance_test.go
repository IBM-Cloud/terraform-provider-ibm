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
					resource.TestCheckResourceAttr(name, "instance_type", ""),
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
		instance_type     = ""
		high_availability = "no"
		backup_location   = "us"
		tags              = ["one:two"]

		parameters_json   = <<EOF
		{
			"disk_encryption_instance_crn": "none",
			"disk_encryption_key_crn": "none",
			"oracle_compatibility": "no"
		}
		EOF

		timeouts {
			create = "720m"
			update = "30m"
			delete = "30m"
		}
	}

	`, databaseResourceGroup, testName)
}
