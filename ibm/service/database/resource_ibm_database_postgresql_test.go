// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package database_test

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"
	"testing"

	"time"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"

	rc "github.com/IBM/platform-services-go-sdk/resourcecontrollerv2"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/IBM-Cloud/bluemix-go/bmxerror"
	"github.com/IBM-Cloud/bluemix-go/models"
)

const (
	databaseInstanceSuccessStatus      = "active"
	databaseInstanceProvisioningStatus = "provisioning"
	databaseInstanceProgressStatus     = "in progress"
	databaseInstanceInactiveStatus     = "inactive"
	databaseInstanceFailStatus         = "failed"
	databaseInstanceRemovedStatus      = "removed"
	databaseInstanceReclamation        = "pending_reclamation"
)

func TestAccIBMDatabaseInstancePostgresBasic(t *testing.T) {
	t.Parallel()
	databaseResourceGroup := "default"
	var databaseInstanceOne string
	rnd := fmt.Sprintf("tf-Pgress-%d", acctest.RandIntRange(10, 100))
	testName := rnd
	name := "ibm_database." + testName

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMDatabaseInstanceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMDatabaseInstancePostgresBasic(databaseResourceGroup, testName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMDatabaseInstanceExists(name, &databaseInstanceOne),
					resource.TestCheckResourceAttr(name, "name", testName),
					resource.TestCheckResourceAttr(name, "service", "databases-for-postgresql"),
					resource.TestCheckResourceAttr(name, "plan", "standard"),
					resource.TestCheckResourceAttr(name, "location", acc.IcdDbRegion),
					resource.TestCheckResourceAttr(name, "adminuser", "admin"),
					resource.TestCheckResourceAttr(name, "members_memory_allocation_mb", "2048"),
					resource.TestCheckResourceAttr(name, "members_disk_allocation_mb", "10240"),
					resource.TestCheckResourceAttr(name, "members_cpu_allocation_count", "0"),
					resource.TestCheckResourceAttr(name, "service_endpoints", "public"),
					resource.TestCheckResourceAttr(name, "whitelist.#", "1"),
					resource.TestCheckResourceAttr(name, "users.#", "1"),
					resource.TestCheckResourceAttr(name, "connectionstrings.#", "2"),
					resource.TestCheckResourceAttr(name, "connectionstrings.1.name", "admin"),
					resource.TestMatchResourceAttr(name, "connectionstrings.1.certname", regexp.MustCompile("[-a-z0-9]*")),
					resource.TestMatchResourceAttr(name, "connectionstrings.1.certbase64", regexp.MustCompile("^(?:[A-Za-z0-9+/]{4})*(?:[A-Za-z0-9+/]{2}==|[A-Za-z0-9+/]{3}=)?$")),
					resource.TestCheckResourceAttr(name, "tags.#", "1"),
				),
			},
			{
				Config: testAccCheckIBMDatabaseInstancePostgresFullyspecified(databaseResourceGroup, testName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMDatabaseInstanceExists(name, &databaseInstanceOne),
					resource.TestCheckResourceAttr(name, "name", testName),
					resource.TestCheckResourceAttr(name, "service", "databases-for-postgresql"),
					resource.TestCheckResourceAttr(name, "plan", "standard"),
					resource.TestCheckResourceAttr(name, "location", acc.IcdDbRegion),
					resource.TestCheckResourceAttr(name, "members_memory_allocation_mb", "4096"),
					resource.TestCheckResourceAttr(name, "members_disk_allocation_mb", "14336"),
					resource.TestCheckResourceAttr(name, "service_endpoints", "public-and-private"),
					resource.TestCheckResourceAttr(name, "whitelist.#", "2"),
					resource.TestCheckResourceAttr(name, "users.#", "2"),
					resource.TestCheckResourceAttr(name, "connectionstrings.#", "3"),
					resource.TestCheckResourceAttr(name, "connectionstrings.2.name", "admin"),
					resource.TestCheckResourceAttr(name, "connectionstrings.0.hosts.#", "1"),
					resource.TestCheckResourceAttr(name, "connectionstrings.0.scheme", "postgres"),
					resource.TestMatchResourceAttr(name, "connectionstrings.0.certname", regexp.MustCompile("[-a-z0-9]*")),
					resource.TestMatchResourceAttr(name, "connectionstrings.0.certbase64", regexp.MustCompile("^(?:[A-Za-z0-9+/]{4})*(?:[A-Za-z0-9+/]{2}==|[A-Za-z0-9+/]{3}=)?$")),
					resource.TestMatchResourceAttr(name, "connectionstrings.0.database", regexp.MustCompile("[-a-z0-9]+")),
					resource.TestCheckResourceAttr(name, "tags.#", "1"),
				),
			},
			{
				Config: testAccCheckIBMDatabaseInstancePostgresReduced(databaseResourceGroup, testName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMDatabaseInstanceExists(name, &databaseInstanceOne),
					resource.TestCheckResourceAttr(name, "name", testName),
					resource.TestCheckResourceAttr(name, "service", "databases-for-postgresql"),
					resource.TestCheckResourceAttr(name, "plan", "standard"),
					resource.TestCheckResourceAttr(name, "location", acc.IcdDbRegion),
					resource.TestCheckResourceAttr(name, "members_memory_allocation_mb", "2048"),
					resource.TestCheckResourceAttr(name, "members_disk_allocation_mb", "14336"),
					resource.TestCheckResourceAttr(name, "whitelist.#", "0"),
					resource.TestCheckResourceAttr(name, "users.#", "0"),
					resource.TestCheckResourceAttr(name, "connectionstrings.#", "1"),
					resource.TestCheckResourceAttr(name, "tags.#", "1"),
				),
			},
			// {
			// 	ResourceName:      name,
			// 	ImportState:       true,
			// 	ImportStateVerify: true,
			// },
		},
	})
}

func TestAccIBMDatabaseInstancePostgresNode(t *testing.T) {
	t.Parallel()
	databaseResourceGroup := "default"
	var databaseInstanceOne string
	rnd := fmt.Sprintf("tf-Pgress-%d", acctest.RandIntRange(10, 100))
	testName := rnd
	name := "ibm_database." + testName

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMDatabaseInstanceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMDatabaseInstancePostgresNodeBasic(databaseResourceGroup, testName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMDatabaseInstanceExists(name, &databaseInstanceOne),
					resource.TestCheckResourceAttr(name, "name", testName),
					resource.TestCheckResourceAttr(name, "service", "databases-for-postgresql"),
					resource.TestCheckResourceAttr(name, "plan", "standard"),
					resource.TestCheckResourceAttr(name, "location", acc.IcdDbRegion),
					resource.TestCheckResourceAttr(name, "adminuser", "admin"),
					resource.TestCheckResourceAttr(name, "node_count", "2"),
					resource.TestCheckResourceAttr(name, "node_memory_allocation_mb", "1024"),
					resource.TestCheckResourceAttr(name, "node_disk_allocation_mb", "5120"),
					resource.TestCheckResourceAttr(name, "node_cpu_allocation_count", "3"),
					resource.TestCheckResourceAttr(name, "service_endpoints", "public"),
					resource.TestCheckResourceAttr(name, "whitelist.#", "1"),
					resource.TestCheckResourceAttr(name, "users.#", "1"),
					resource.TestCheckResourceAttr(name, "connectionstrings.#", "2"),
					resource.TestCheckResourceAttr(name, "connectionstrings.1.name", "admin"),
					resource.TestMatchResourceAttr(name, "connectionstrings.1.certname", regexp.MustCompile("[-a-z0-9]*")),
					resource.TestMatchResourceAttr(name, "connectionstrings.1.certbase64", regexp.MustCompile("^(?:[A-Za-z0-9+/]{4})*(?:[A-Za-z0-9+/]{2}==|[A-Za-z0-9+/]{3}=)?$")),
					resource.TestCheckResourceAttr(name, "tags.#", "1"),
				),
			},
			{
				Config: testAccCheckIBMDatabaseInstancePostgresNodeFullyspecified(databaseResourceGroup, testName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMDatabaseInstanceExists(name, &databaseInstanceOne),
					resource.TestCheckResourceAttr(name, "name", testName),
					resource.TestCheckResourceAttr(name, "service", "databases-for-postgresql"),
					resource.TestCheckResourceAttr(name, "plan", "standard"),
					resource.TestCheckResourceAttr(name, "location", acc.IcdDbRegion),
					resource.TestCheckResourceAttr(name, "node_count", "2"),
					resource.TestCheckResourceAttr(name, "node_memory_allocation_mb", "1024"),
					resource.TestCheckResourceAttr(name, "node_disk_allocation_mb", "7168"),
					resource.TestCheckResourceAttr(name, "node_cpu_allocation_count", "3"),
					resource.TestCheckResourceAttr(name, "service_endpoints", "public-and-private"),
					resource.TestCheckResourceAttr(name, "whitelist.#", "2"),
					resource.TestCheckResourceAttr(name, "users.#", "2"),
					resource.TestCheckResourceAttr(name, "connectionstrings.#", "3"),
					resource.TestCheckResourceAttr(name, "connectionstrings.2.name", "admin"),
					resource.TestCheckResourceAttr(name, "connectionstrings.0.hosts.#", "1"),
					resource.TestCheckResourceAttr(name, "connectionstrings.0.scheme", "postgres"),
					resource.TestMatchResourceAttr(name, "connectionstrings.0.certname", regexp.MustCompile("[-a-z0-9]*")),
					resource.TestMatchResourceAttr(name, "connectionstrings.0.certbase64", regexp.MustCompile("^(?:[A-Za-z0-9+/]{4})*(?:[A-Za-z0-9+/]{2}==|[A-Za-z0-9+/]{3}=)?$")),
					resource.TestMatchResourceAttr(name, "connectionstrings.0.database", regexp.MustCompile("[-a-z0-9]+")),
					resource.TestCheckResourceAttr(name, "tags.#", "1"),
				),
			},
			{
				Config: testAccCheckIBMDatabaseInstancePostgresNodeReduced(databaseResourceGroup, testName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMDatabaseInstanceExists(name, &databaseInstanceOne),
					resource.TestCheckResourceAttr(name, "name", testName),
					resource.TestCheckResourceAttr(name, "service", "databases-for-postgresql"),
					resource.TestCheckResourceAttr(name, "plan", "standard"),
					resource.TestCheckResourceAttr(name, "location", acc.IcdDbRegion),
					resource.TestCheckResourceAttr(name, "node_count", "2"),
					resource.TestCheckResourceAttr(name, "node_memory_allocation_mb", "1024"),
					resource.TestCheckResourceAttr(name, "node_disk_allocation_mb", "7168"),
					resource.TestCheckResourceAttr(name, "node_cpu_allocation_count", "3"),
					resource.TestCheckResourceAttr(name, "whitelist.#", "0"),
					resource.TestCheckResourceAttr(name, "users.#", "0"),
					resource.TestCheckResourceAttr(name, "connectionstrings.#", "1"),
					resource.TestCheckResourceAttr(name, "tags.#", "1"),
				),
			},
			{
				Config: testAccCheckIBMDatabaseInstancePostgresNodeScaleOut(databaseResourceGroup, testName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMDatabaseInstanceExists(name, &databaseInstanceOne),
					resource.TestCheckResourceAttr(name, "name", testName),
					resource.TestCheckResourceAttr(name, "service", "databases-for-postgresql"),
					resource.TestCheckResourceAttr(name, "plan", "standard"),
					resource.TestCheckResourceAttr(name, "location", acc.IcdDbRegion),
					resource.TestCheckResourceAttr(name, "node_count", "3"),
					resource.TestCheckResourceAttr(name, "node_memory_allocation_mb", "1024"),
					resource.TestCheckResourceAttr(name, "node_disk_allocation_mb", "7168"),
					resource.TestCheckResourceAttr(name, "node_cpu_allocation_count", "3"),
					resource.TestCheckResourceAttr(name, "whitelist.#", "0"),
					resource.TestCheckResourceAttr(name, "users.#", "0"),
					resource.TestCheckResourceAttr(name, "connectionstrings.#", "1"),
					resource.TestCheckResourceAttr(name, "tags.#", "1"),
				),
			},
			// {
			// 	ResourceName:      name,
			// 	ImportState:       true,
			// 	ImportStateVerify: true,
			// },
		},
	})
}

func TestAccIBMDatabaseInstancePostgresGroup(t *testing.T) {
	t.Parallel()
	databaseResourceGroup := "default"
	var databaseInstanceOne string
	rnd := fmt.Sprintf("tf-Pgress-%d", acctest.RandIntRange(10, 100))
	testName := rnd
	name := "ibm_database." + testName

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMDatabaseInstanceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMDatabaseInstancePostgresGroupBasic(databaseResourceGroup, testName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMDatabaseInstanceExists(name, &databaseInstanceOne),
					resource.TestCheckResourceAttr(name, "name", testName),
					resource.TestCheckResourceAttr(name, "service", "databases-for-postgresql"),
					resource.TestCheckResourceAttr(name, "plan", "standard"),
					resource.TestCheckResourceAttr(name, "location", acc.IcdDbRegion),
					resource.TestCheckResourceAttr(name, "adminuser", "admin"),
					resource.TestCheckResourceAttr(name, "groups.0.count", "2"),
					resource.TestCheckResourceAttr(name, "groups.0.memory.0.allocation_mb", "2048"),
					resource.TestCheckResourceAttr(name, "groups.0.disk.0.allocation_mb", "10240"),
					resource.TestCheckResourceAttr(name, "groups.0.cpu.0.allocation_count", "6"),
					resource.TestCheckResourceAttr(name, "service_endpoints", "public"),
					resource.TestCheckResourceAttr(name, "whitelist.#", "1"),
					resource.TestCheckResourceAttr(name, "users.#", "1"),
					resource.TestCheckResourceAttr(name, "connectionstrings.#", "2"),
					resource.TestCheckResourceAttr(name, "connectionstrings.1.name", "admin"),
					resource.TestMatchResourceAttr(name, "connectionstrings.1.certname", regexp.MustCompile("[-a-z0-9]*")),
					resource.TestMatchResourceAttr(name, "connectionstrings.1.certbase64", regexp.MustCompile("^(?:[A-Za-z0-9+/]{4})*(?:[A-Za-z0-9+/]{2}==|[A-Za-z0-9+/]{3}=)?$")),
					resource.TestCheckResourceAttr(name, "tags.#", "1"),
				),
			},
			{
				Config: testAccCheckIBMDatabaseInstancePostgresGroupFullyspecified(databaseResourceGroup, testName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMDatabaseInstanceExists(name, &databaseInstanceOne),
					resource.TestCheckResourceAttr(name, "name", testName),
					resource.TestCheckResourceAttr(name, "service", "databases-for-postgresql"),
					resource.TestCheckResourceAttr(name, "plan", "standard"),
					resource.TestCheckResourceAttr(name, "location", acc.IcdDbRegion),
					resource.TestCheckResourceAttr(name, "groups.0.count", "2"),
					resource.TestCheckResourceAttr(name, "groups.0.memory.0.allocation_mb", "2304"),
					resource.TestCheckResourceAttr(name, "groups.0.disk.0.allocation_mb", "14336"),
					resource.TestCheckResourceAttr(name, "groups.0.cpu.0.allocation_count", "6"),
					resource.TestCheckResourceAttr(name, "service_endpoints", "public-and-private"),
					resource.TestCheckResourceAttr(name, "whitelist.#", "2"),
					resource.TestCheckResourceAttr(name, "users.#", "2"),
					resource.TestCheckResourceAttr(name, "connectionstrings.#", "3"),
					resource.TestCheckResourceAttr(name, "connectionstrings.2.name", "admin"),
					resource.TestCheckResourceAttr(name, "connectionstrings.0.hosts.#", "1"),
					resource.TestCheckResourceAttr(name, "connectionstrings.0.scheme", "postgres"),
					resource.TestMatchResourceAttr(name, "connectionstrings.0.certname", regexp.MustCompile("[-a-z0-9]*")),
					resource.TestMatchResourceAttr(name, "connectionstrings.0.certbase64", regexp.MustCompile("^(?:[A-Za-z0-9+/]{4})*(?:[A-Za-z0-9+/]{2}==|[A-Za-z0-9+/]{3}=)?$")),
					resource.TestMatchResourceAttr(name, "connectionstrings.0.database", regexp.MustCompile("[-a-z0-9]+")),
					resource.TestCheckResourceAttr(name, "tags.#", "1"),
				),
			},
			{
				Config: testAccCheckIBMDatabaseInstancePostgresGroupReduced(databaseResourceGroup, testName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMDatabaseInstanceExists(name, &databaseInstanceOne),
					resource.TestCheckResourceAttr(name, "name", testName),
					resource.TestCheckResourceAttr(name, "service", "databases-for-postgresql"),
					resource.TestCheckResourceAttr(name, "plan", "standard"),
					resource.TestCheckResourceAttr(name, "location", acc.IcdDbRegion),
					resource.TestCheckResourceAttr(name, "groups.0.count", "2"),
					resource.TestCheckResourceAttr(name, "groups.0.memory.0.allocation_mb", "2048"),
					resource.TestCheckResourceAttr(name, "groups.0.disk.0.allocation_mb", "14336"),
					resource.TestCheckResourceAttr(name, "groups.0.cpu.0.allocation_count", "6"),
					resource.TestCheckResourceAttr(name, "whitelist.#", "0"),
					resource.TestCheckResourceAttr(name, "users.#", "0"),
					resource.TestCheckResourceAttr(name, "connectionstrings.#", "1"),
					resource.TestCheckResourceAttr(name, "tags.#", "1"),
				),
			},
			{
				Config: testAccCheckIBMDatabaseInstancePostgresGroupScaleOut(databaseResourceGroup, testName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMDatabaseInstanceExists(name, &databaseInstanceOne),
					resource.TestCheckResourceAttr(name, "name", testName),
					resource.TestCheckResourceAttr(name, "service", "databases-for-postgresql"),
					resource.TestCheckResourceAttr(name, "plan", "standard"),
					resource.TestCheckResourceAttr(name, "location", acc.IcdDbRegion),
					resource.TestCheckResourceAttr(name, "groups.0.count", "3"),
					resource.TestCheckResourceAttr(name, "groups.0.memory.0.allocation_mb", "3072"),
					resource.TestCheckResourceAttr(name, "groups.0.disk.0.allocation_mb", "21504"),
					resource.TestCheckResourceAttr(name, "groups.0.cpu.0.allocation_count", "9"),
					resource.TestCheckResourceAttr(name, "whitelist.#", "0"),
					resource.TestCheckResourceAttr(name, "users.#", "0"),
					resource.TestCheckResourceAttr(name, "connectionstrings.#", "1"),
					resource.TestCheckResourceAttr(name, "tags.#", "1"),
				),
			},
		},
	})
}

// TestAccIBMDatabaseInstance_CreateAfterManualDestroy not required as tested by resource_instance tests

func TestAccIBMDatabaseInstancePostgresImport(t *testing.T) {
	t.Parallel()
	databaseResourceGroup := "default"
	var databaseInstanceOne string
	serviceName := fmt.Sprintf("tf-Pgress-%d", acctest.RandIntRange(10, 100))
	//serviceName := "test_acc"
	resourceName := "ibm_database." + serviceName

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMDatabaseInstanceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMDatabaseInstancePostgresImport(databaseResourceGroup, serviceName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMDatabaseInstanceExists(resourceName, &databaseInstanceOne),
					resource.TestCheckResourceAttr(resourceName, "name", serviceName),
					resource.TestCheckResourceAttr(resourceName, "service", "databases-for-postgresql"),
					resource.TestCheckResourceAttr(resourceName, "plan", "standard"),
					resource.TestCheckResourceAttr(resourceName, "location", acc.IcdDbRegion),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"wait_time_minutes", "plan_validation"},
			},
		},
	})
}

func testAccCheckIBMDatabaseInstanceDestroy(s *terraform.State) error {
	rsContClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).ResourceControllerV2API()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_database" {
			continue
		}

		instanceID := rs.Primary.ID

		rsInst := rc.GetResourceInstanceOptions{
			ID: &instanceID,
		}
		instance, response, err := rsContClient.GetResourceInstance(&rsInst)
		if err == nil {
			if !reflect.DeepEqual(instance, models.ServiceInstance{}) && *instance.State == "active" {
				return fmt.Errorf("Database still exists: %s", rs.Primary.ID)
			}
		} else {
			if !strings.Contains(err.Error(), "404") {
				return fmt.Errorf("[ERROR] Error checking if database (%s) has been destroyed: %s %s", rs.Primary.ID, err, response)
			}
		}
	}
	return nil
}

func testAccDatabaseInstanceManuallyDelete(tfDatabaseID *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		_ = testAccDatabaseInstanceManuallyDeleteUnwrapped(s, tfDatabaseID)
		return nil
	}
}

func testAccDatabaseInstanceManuallyDeleteUnwrapped(s *terraform.State, tfDatabaseID *string) error {
	rsConClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).ResourceControllerV2API()
	if err != nil {
		return err
	}
	instance := *tfDatabaseID
	var instanceID string
	if strings.HasPrefix(instance, "crn") {
		instanceID = instance
	} else {
		_, instanceID, _ = flex.ConvertTftoCisTwoVar(instance)
	}
	recursive := true
	deleteReq := rc.DeleteResourceInstanceOptions{
		ID:        &instanceID,
		Recursive: &recursive,
	}
	response, err := rsConClient.DeleteResourceInstance(&deleteReq)
	if err != nil {
		return fmt.Errorf("[ERROR] Error deleting resource instance: %s %s", err, response)
	}

	_ = &resource.StateChangeConf{
		Pending: []string{databaseInstanceProgressStatus, databaseInstanceInactiveStatus, databaseInstanceSuccessStatus},
		Target:  []string{databaseInstanceRemovedStatus},
		Refresh: func() (interface{}, string, error) {
			rsInst := rc.GetResourceInstanceOptions{
				ID: &instanceID,
			}
			instance, response, err := rsConClient.GetResourceInstance(&rsInst)
			if err != nil {
				if apiErr, ok := err.(bmxerror.RequestFailure); ok && apiErr.StatusCode() == 404 {
					return instance, databaseInstanceSuccessStatus, nil
				}
				return nil, "", err
			}
			if *instance.State == databaseInstanceFailStatus {
				return instance, *instance.State, fmt.Errorf("[ERROR] The resource instance %s failed to delete: %v %s", instanceID, err, response)
			}
			return instance, *instance.State, nil
		},
		Timeout:    90 * time.Second,
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}
	if err != nil {
		return fmt.Errorf("[ERROR] Error waiting for resource instance (%s) to be deleted: %s", instanceID, err)
	}
	return nil
}

func testAccCheckIBMDatabaseInstanceExists(n string, tfDatabaseID *string) resource.TestCheckFunc {

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

		rsInst := rc.GetResourceInstanceOptions{
			ID: &instanceID,
		}
		instance, response, err := rsContClient.GetResourceInstance(&rsInst)
		if err != nil {
			if strings.Contains(err.Error(), "Object not found") ||
				strings.Contains(err.Error(), "status code: 404") {
				*tfDatabaseID = ""
				return nil
			}
			return fmt.Errorf("[ERROR] Error retrieving resource instance: %s %s", err, response)
		}
		if strings.Contains(*instance.State, "removed") {
			*tfDatabaseID = ""
			return nil
		}

		*tfDatabaseID = instanceID
		return nil
	}
}

func testAccCheckIBMDatabaseInstancePostgresBasic(databaseResourceGroup string, name string) string {
	return fmt.Sprintf(`
	data "ibm_resource_group" "test_acc" {
		name = "%[1]s"
	}

	resource "ibm_database" "%[2]s" {
		resource_group_id            = data.ibm_resource_group.test_acc.id
		name                         = "%[2]s"
		service                      = "databases-for-postgresql"
		plan                         = "standard"
		location                     = "%[3]s"
		adminpassword                = "password12"
		members_memory_allocation_mb = 2048
		members_disk_allocation_mb   = 10240
		tags                         = ["one:two"]
		users {
			name     = "user123"
			password = "password12"
		}
		whitelist {
			address     = "172.168.1.2/32"
			description = "desc1"
		}
	}
				`, databaseResourceGroup, name, acc.IcdDbRegion)
}

func testAccCheckIBMDatabaseInstancePostgresFullyspecified(databaseResourceGroup string, name string) string {
	return fmt.Sprintf(`
	data "ibm_resource_group" "test_acc" {
		name = "%[1]s"
	}

	resource "ibm_database" "%[2]s" {
		resource_group_id            = data.ibm_resource_group.test_acc.id
		name                         = "%[2]s"
		service                      = "databases-for-postgresql"
		plan                         = "standard"
		location                     = "%[3]s"
		adminpassword                = "password12"
		members_memory_allocation_mb = 4096
		members_disk_allocation_mb   = 14336
		members_cpu_allocation_count = 6
		service_endpoints            = "public-and-private"
		tags                         = ["one:two"]
		users {
			name     = "user123"
			password = "password12"
		}
		users {
			name     = "user124"
			password = "password12"
		}
		whitelist {
			address     = "172.168.1.2/32"
			description = "desc1"
		}
		whitelist {
			address     = "172.168.1.1/32"
			description = "desc"
		}
	}
				`, databaseResourceGroup, name, acc.IcdDbRegion)
}

func testAccCheckIBMDatabaseInstancePostgresReduced(databaseResourceGroup string, name string) string {
	return fmt.Sprintf(`
	data "ibm_resource_group" "test_acc" {
		name = "%[1]s"
	}

	resource "ibm_database" "%[2]s" {
		resource_group_id            = data.ibm_resource_group.test_acc.id
		name                         = "%[2]s"
		service                      = "databases-for-postgresql"
		plan                         = "standard"
		location                     = "%[3]s"
		adminpassword                = "password12"
		members_memory_allocation_mb = 2048
		members_disk_allocation_mb   = 14336
		service_endpoints            = "public"
		tags                         = ["one:two"]
	  }
				`, databaseResourceGroup, name, acc.IcdDbRegion)
}

func testAccCheckIBMDatabaseInstancePostgresNodeBasic(databaseResourceGroup string, name string) string {
	return fmt.Sprintf(`
	data "ibm_resource_group" "test_acc" {
		name = "%[1]s"
	}

	resource "ibm_database" "%[2]s" {
		resource_group_id            = data.ibm_resource_group.test_acc.id
		name                         = "%[2]s"
		service                      = "databases-for-postgresql"
		plan                         = "standard"
		location                     = "%[3]s"
		adminpassword                = "password12"
		node_count                   = 2
		node_memory_allocation_mb    = 1024
		node_disk_allocation_mb      = 5120
		node_cpu_allocation_count    = 3
		tags                         = ["one:two"]
		users {
			name     = "user123"
			password = "password12"
		}
		whitelist {
			address     = "172.168.1.2/32"
			description = "desc1"
		}
	}
				`, databaseResourceGroup, name, acc.IcdDbRegion)
}

func testAccCheckIBMDatabaseInstancePostgresNodeFullyspecified(databaseResourceGroup string, name string) string {
	return fmt.Sprintf(`
	data "ibm_resource_group" "test_acc" {
		name = "%[1]s"
	}

	resource "ibm_database" "%[2]s" {
		resource_group_id            = data.ibm_resource_group.test_acc.id
		name                         = "%[2]s"
		service                      = "databases-for-postgresql"
		plan                         = "standard"
		location                     = "%[3]s"
		adminpassword                = "password12"
		node_count                   = 2
		node_memory_allocation_mb    = 1024
		node_disk_allocation_mb      = 7168
		node_cpu_allocation_count    = 3
		service_endpoints            = "public-and-private"
		tags                         = ["one:two"]
		users {
			name     = "user123"
			password = "password12"
		}
		users {
			name     = "user124"
			password = "password12"
		}
		whitelist {
			address     = "172.168.1.2/32"
			description = "desc1"
		}
		whitelist {
			address     = "172.168.1.1/32"
			description = "desc"
		}
	}
				`, databaseResourceGroup, name, acc.IcdDbRegion)
}

func testAccCheckIBMDatabaseInstancePostgresNodeReduced(databaseResourceGroup string, name string) string {
	return fmt.Sprintf(`
	data "ibm_resource_group" "test_acc" {
		name = "%[1]s"
	}

	resource "ibm_database" "%[2]s" {
		resource_group_id            = data.ibm_resource_group.test_acc.id
		name                         = "%[2]s"
		service                      = "databases-for-postgresql"
		plan                         = "standard"
		location                     = "%[3]s"
		adminpassword                = "password12"
		node_count                   = 2
		node_memory_allocation_mb    = 1024
		node_disk_allocation_mb      = 7168
		node_cpu_allocation_count    = 3
		service_endpoints            = "public"
		tags                         = ["one:two"]
	  }
				`, databaseResourceGroup, name, acc.IcdDbRegion)
}
func testAccCheckIBMDatabaseInstancePostgresNodeScaleOut(databaseResourceGroup string, name string) string {
	return fmt.Sprintf(`
	data "ibm_resource_group" "test_acc" {
		name = "%[1]s"
	}

	resource "ibm_database" "%[2]s" {
		resource_group_id            = data.ibm_resource_group.test_acc.id
		name                         = "%[2]s"
		service                      = "databases-for-postgresql"
		plan                         = "standard"
		location                     = "%[3]s"
		adminpassword                = "password12"
		node_count                   = 3
		node_memory_allocation_mb    = 1024
		node_disk_allocation_mb      = 7168
		node_cpu_allocation_count    = 3
		service_endpoints            = "public"
		tags                         = ["one:two"]
	}
				`, databaseResourceGroup, name, acc.IcdDbRegion)
}

func testAccCheckIBMDatabaseInstancePostgresGroupBasic(databaseResourceGroup string, name string) string {
	return fmt.Sprintf(`
	data "ibm_resource_group" "test_acc" {
		name = "%[1]s"
	}

	resource "ibm_database" "%[2]s" {
		resource_group_id            = data.ibm_resource_group.test_acc.id
		name                         = "%[2]s"
		service                      = "databases-for-postgresql"
		plan                         = "standard"
		location                     = "%[3]s"
		adminpassword                = "password12"
		tags                         = ["one:two"]
		group {
			group_id = "member"
			members {
				allocation_count = 2
			}
			memory {
				allocation_mb = 1024
			}
			 disk {
				allocation_mb = 5120
			}
			cpu {
				allocation_count = 3
			}
		}
		users {
			name     = "user123"
			password = "password12"
		}
		whitelist {
			address     = "172.168.1.2/32"
			description = "desc1"
		}
	}
				`, databaseResourceGroup, name, acc.IcdDbRegion)
}

func testAccCheckIBMDatabaseInstancePostgresGroupFullyspecified(databaseResourceGroup string, name string) string {
	return fmt.Sprintf(`
	data "ibm_resource_group" "test_acc" {
		name = "%[1]s"
	}

	resource "ibm_database" "%[2]s" {
		resource_group_id            = data.ibm_resource_group.test_acc.id
		name                         = "%[2]s"
		service                      = "databases-for-postgresql"
		plan                         = "standard"
		location                     = "%[3]s"
		adminpassword                = "password12"
		service_endpoints            = "public-and-private"
		tags                         = ["one:two"]
		group {
			group_id = "member"
			members {
				allocation_count = 2
			}
			memory {
				allocation_mb = 1152
			}
			 disk {
				allocation_mb = 7168
			}
			cpu {
				allocation_count = 3
			}
		}
		users {
			name     = "user123"
			password = "password12"
		}
		users {
			name     = "user124"
			password = "password12"
		}
		whitelist {
			address     = "172.168.1.2/32"
			description = "desc1"
		}
		whitelist {
			address     = "172.168.1.1/32"
			description = "desc"
		}
	}
				`, databaseResourceGroup, name, acc.IcdDbRegion)
}

func testAccCheckIBMDatabaseInstancePostgresGroupReduced(databaseResourceGroup string, name string) string {
	return fmt.Sprintf(`
	data "ibm_resource_group" "test_acc" {
		name = "%[1]s"
	}

	resource "ibm_database" "%[2]s" {
		resource_group_id            = data.ibm_resource_group.test_acc.id
		name                         = "%[2]s"
		service                      = "databases-for-postgresql"
		plan                         = "standard"
		location                     = "%[3]s"
		adminpassword                = "password12"
		service_endpoints            = "public"
		tags                         = ["one:two"]
		group {
			group_id = "member"
			members {
				allocation_count = 2
			}
			memory {
				allocation_mb = 1024
			}
			 disk {
				allocation_mb = 7168
			}
			cpu {
				allocation_count = 3
			}
		}
	}
				`, databaseResourceGroup, name, acc.IcdDbRegion)
}

func testAccCheckIBMDatabaseInstancePostgresGroupScaleOut(databaseResourceGroup string, name string) string {
	return fmt.Sprintf(`
	data "ibm_resource_group" "test_acc" {
		name = "%[1]s"
	}

	resource "ibm_database" "%[2]s" {
		resource_group_id            = data.ibm_resource_group.test_acc.id
		name                         = "%[2]s"
		service                      = "databases-for-postgresql"
		plan                         = "standard"
		location                     = "%[3]s"
		adminpassword                = "password12"
		group {
			group_id = "member"
			members {
				allocation_count = 3
			}
			memory {
				allocation_mb = 1024
			}
			 disk {
				allocation_mb = 7168
			}
			cpu {
				allocation_count = 3
			}
		}
		service_endpoints            = "public"
		tags                         = ["one:two"]
	}
				`, databaseResourceGroup, name, acc.IcdDbRegion)
}

func testAccCheckIBMDatabaseInstancePostgresImport(databaseResourceGroup string, name string) string {
	return fmt.Sprintf(`
	data "ibm_resource_group" "test_acc" {
		is_default = true
		# name = "%[1]s"
	}

	resource "ibm_database" "%[2]s" {
		resource_group_id = data.ibm_resource_group.test_acc.id
		name              = "%[2]s"
		service           = "databases-for-postgresql"
		plan              = "standard"
		location          = "%[3]s"
	  }
				`, databaseResourceGroup, name, acc.IcdDbRegion)
}
