package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccIBMDatabaseInstance_Elasticsearch_Basic(t *testing.T) {
	t.Parallel()
	databaseResourceGroup := "default"
	var databaseInstanceOne string
	rnd := fmt.Sprintf("tf_test_acc_%d", acctest.RandInt())
	testName := rnd
	name := "ibm_database." + testName

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMDatabaseInstanceDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMDatabaseInstance_Elasticsearch_basic(databaseResourceGroup, testName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMDatabaseInstanceExists(name, &databaseInstanceOne),
					resource.TestCheckResourceAttr(name, "name", testName),
					resource.TestCheckResourceAttr(name, "service", "databases-for-elasticsearch"),
					resource.TestCheckResourceAttr(name, "plan", "standard"),
					resource.TestCheckResourceAttr(name, "location", "us-south"),
					resource.TestCheckResourceAttr(name, "adminuser", "admin"),
					resource.TestCheckResourceAttr(name, "members_memory_allocation_mb", "3072"),
					resource.TestCheckResourceAttr(name, "members_disk_allocation_mb", "15360"),
					resource.TestCheckResourceAttr(name, "whitelist.#", "1"),
					resource.TestCheckResourceAttr(name, "users.#", "1"),
					resource.TestCheckResourceAttr(name, "connectionstrings.#", "2"),
					resource.TestCheckResourceAttr(name, "connectionstrings.1.name", "admin"),
					resource.TestCheckResourceAttr(name, "connectionstrings.0.hosts.#", "1"),
					resource.TestCheckResourceAttr(name, "connectionstrings.0.database", ""),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMDatabaseInstance_Elasticsearch_fullyspecified(databaseResourceGroup, testName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMDatabaseInstanceExists(name, &databaseInstanceOne),
					resource.TestCheckResourceAttr(name, "name", testName),
					resource.TestCheckResourceAttr(name, "service", "databases-for-elasticsearch"),
					resource.TestCheckResourceAttr(name, "plan", "standard"),
					resource.TestCheckResourceAttr(name, "location", "us-south"),
					resource.TestCheckResourceAttr(name, "members_memory_allocation_mb", "6144"),
					resource.TestCheckResourceAttr(name, "members_disk_allocation_mb", "18432"),
					resource.TestCheckResourceAttr(name, "whitelist.#", "2"),
					resource.TestCheckResourceAttr(name, "users.#", "2"),
					resource.TestCheckResourceAttr(name, "connectionstrings.#", "3"),
					resource.TestCheckResourceAttr(name, "connectionstrings.2.name", "admin"),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMDatabaseInstance_Elasticsearch_reduced(databaseResourceGroup, testName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMDatabaseInstanceExists(name, &databaseInstanceOne),
					resource.TestCheckResourceAttr(name, "name", testName),
					resource.TestCheckResourceAttr(name, "service", "databases-for-elasticsearch"),
					resource.TestCheckResourceAttr(name, "plan", "standard"),
					resource.TestCheckResourceAttr(name, "location", "us-south"),
					resource.TestCheckResourceAttr(name, "members_memory_allocation_mb", "3072"),
					resource.TestCheckResourceAttr(name, "members_disk_allocation_mb", "18432"),
					resource.TestCheckResourceAttr(name, "whitelist.#", "0"),
					resource.TestCheckResourceAttr(name, "users.#", "0"),
					resource.TestCheckResourceAttr(name, "connectionstrings.#", "1"),
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

// TestAccIBMDatabaseInstance_CreateAfterManualDestroy not required as tested by resource_instance tests

func TestAccIBMDatabaseInstance_Elasticsearch_import(t *testing.T) {
	t.Parallel()
	databaseResourceGroup := "default"
	var databaseInstanceOne string
	serviceName := fmt.Sprintf("tf_test_acc_%d", acctest.RandInt())
	//serviceName := "test_acc"
	resourceName := "ibm_database." + serviceName

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMDatabaseInstanceDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMDatabaseInstance_Elasticsearch_import(databaseResourceGroup, serviceName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMDatabaseInstanceExists(resourceName, &databaseInstanceOne),
					resource.TestCheckResourceAttr(resourceName, "name", serviceName),
					resource.TestCheckResourceAttr(resourceName, "service", "databases-for-elasticsearch"),
					resource.TestCheckResourceAttr(resourceName, "plan", "standard"),
					resource.TestCheckResourceAttr(resourceName, "location", "us-south"),
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

// func testAccCheckIBMDatabaseInstanceDestroy(s *terraform.State) etc in resource_ibm_database_postgresql_test.go

func testAccCheckIBMDatabaseInstance_Elasticsearch_basic(databaseResourceGroup string, name string) string {
	return fmt.Sprintf(`
				data "ibm_resource_group" "test_acc" {
				  is_default = true
				  # name = "%[1]s"
				}

				resource "ibm_database" "%[2]s" {
				  resource_group_id = "${data.ibm_resource_group.test_acc.id}"	
				  name = "%[2]s"	
				  service 			= "databases-for-elasticsearch"
				  plan              = "standard"
				  location          = "us-south"
				  adminpassword     = "password12"
				  members_memory_allocation_mb = 3072
  				  members_disk_allocation_mb   = 15360
				  users = {
  				   		name     = "user123"
  					   	password = "password12"
  				  		}
  				  whitelist = {
  						address     = "172.168.1.2/32"
  						description = "desc1"
  						}
				}`, databaseResourceGroup, name)
}

func testAccCheckIBMDatabaseInstance_Elasticsearch_fullyspecified(databaseResourceGroup string, name string) string {
	return fmt.Sprintf(`
				data "ibm_resource_group" "test_acc" {
				  is_default = true
				  # name = "%[1]s"
				}

				resource "ibm_database" "%[2]s" {
				  resource_group_id = "${data.ibm_resource_group.test_acc.id}"	
				  name = "%[2]s"	
				  service 			= "databases-for-elasticsearch"
				  plan              = "standard"
				  location          = "us-south"
				  adminpassword     = "password12"
				  members_memory_allocation_mb = 6144
  				  members_disk_allocation_mb   = 18432
				  users = {
  				   		name     = "user123"
  					   	password = "password12"
  				  		}
  				  users = {
  				   		name     = "user124"
  					   	password = "password12"
  				  		}		
  					whitelist = {
  						address     = "172.168.1.2/32"
  						description = "desc1"
  						}
  					whitelist = {
  						address     = "172.168.1.1/32"
  						description = "desc"
  						}	
				}`, databaseResourceGroup, name)
}

func testAccCheckIBMDatabaseInstance_Elasticsearch_reduced(databaseResourceGroup string, name string) string {
	return fmt.Sprintf(`
				data "ibm_resource_group" "test_acc" {
				  is_default = true
				  # name = "%[1]s"
				}

				resource "ibm_database" "%[2]s" {
				  resource_group_id = "${data.ibm_resource_group.test_acc.id}"	
				  name = "%[2]s"	
				  service 			= "databases-for-elasticsearch"
				  plan              = "standard"
				  location          = "us-south"
				  adminpassword     = "password12"
				  members_memory_allocation_mb = 3072
  				  members_disk_allocation_mb   = 18432

				}`, databaseResourceGroup, name)
}

func testAccCheckIBMDatabaseInstance_Elasticsearch_import(databaseResourceGroup string, name string) string {
	return fmt.Sprintf(`
				data "ibm_resource_group" "test_acc" {
				  is_default = true
				  # name = "%[1]s"
				}

				resource "ibm_database" "%[2]s" {
				  resource_group_id = "${data.ibm_resource_group.test_acc.id}"	
				  name = "%[2]s"	
				  service 			= "databases-for-elasticsearch"
				  plan              = "standard"
				  location          = "us-south"
				}`, databaseResourceGroup, name)
}
