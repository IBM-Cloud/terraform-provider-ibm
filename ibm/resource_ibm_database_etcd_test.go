package ibm

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccIBMDatabaseInstance_Etcd_Basic(t *testing.T) {
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
				Config: testAccCheckIBMDatabaseInstance_Etcd_basic(databaseResourceGroup, testName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMDatabaseInstanceExists(name, &databaseInstanceOne),
					resource.TestCheckResourceAttr(name, "name", testName),
					resource.TestCheckResourceAttr(name, "service", "databases-for-etcd"),
					resource.TestCheckResourceAttr(name, "plan", "standard"),
					resource.TestCheckResourceAttr(name, "location", "us-south"),
					resource.TestCheckResourceAttr(name, "adminuser", "root"),
					resource.TestCheckResourceAttr(name, "members_memory_allocation_mb", "3072"),
					resource.TestCheckResourceAttr(name, "members_disk_allocation_mb", "61440"),
					resource.TestCheckResourceAttr(name, "whitelist.#", "1"),
					resource.TestCheckResourceAttr(name, "users.#", "1"),
					resource.TestCheckResourceAttr(name, "connectionstrings.#", "2"),
					resource.TestCheckResourceAttr(name, "connectionstrings.1.name", "root"),
					resource.TestMatchResourceAttr(name, "connectionstrings.1.certname", regexp.MustCompile("[-a-z0-9]*")),
					resource.TestMatchResourceAttr(name, "connectionstrings.1.certbase64", regexp.MustCompile("^(?:[A-Za-z0-9+/]{4})*(?:[A-Za-z0-9+/]{2}==|[A-Za-z0-9+/]{3}=)?$")),
					resource.TestCheckResourceAttr(name, "connectionstrings.0.hosts.#", "1"),
					resource.TestCheckResourceAttr(name, "connectionstrings.0.database", ""),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMDatabaseInstance_Etcd_fullyspecified(databaseResourceGroup, testName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(name, "name", testName),
					resource.TestCheckResourceAttr(name, "service", "databases-for-etcd"),
					resource.TestCheckResourceAttr(name, "plan", "standard"),
					resource.TestCheckResourceAttr(name, "location", "us-south"),
					resource.TestCheckResourceAttr(name, "members_memory_allocation_mb", "6144"),
					resource.TestCheckResourceAttr(name, "members_disk_allocation_mb", "64512"),
					resource.TestCheckResourceAttr(name, "whitelist.#", "2"),
					resource.TestCheckResourceAttr(name, "users.#", "2"),
					resource.TestCheckResourceAttr(name, "connectionstrings.#", "3"),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMDatabaseInstance_Etcd_reduced(databaseResourceGroup, testName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(name, "name", testName),
					resource.TestCheckResourceAttr(name, "service", "databases-for-etcd"),
					resource.TestCheckResourceAttr(name, "plan", "standard"),
					resource.TestCheckResourceAttr(name, "location", "us-south"),
					resource.TestCheckResourceAttr(name, "members_memory_allocation_mb", "3072"),
					resource.TestCheckResourceAttr(name, "members_disk_allocation_mb", "64512"),
					resource.TestCheckResourceAttr(name, "whitelist.#", "0"),
					resource.TestCheckResourceAttr(name, "users.#", "0"),
					resource.TestCheckResourceAttr(name, "connectionstrings.#", "1"),
				),
			},
		},
	})
}

// TestAccIBMDatabaseInstance_CreateAfterManualDestroy not required as tested by resource_instance tests

func TestAccIBMDatabaseInstance_Etcd_import(t *testing.T) {
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
				Config: testAccCheckIBMDatabaseInstance_Etcd_import(databaseResourceGroup, serviceName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMDatabaseInstanceExists(resourceName, &databaseInstanceOne),
					resource.TestCheckResourceAttr(resourceName, "name", serviceName),
					resource.TestCheckResourceAttr(resourceName, "service", "databases-for-etcd"),
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

func testAccCheckIBMDatabaseInstance_Etcd_basic(databaseResourceGroup string, name string) string {
	return fmt.Sprintf(`
				data "ibm_resource_group" "test_acc" {
					is_default = true
				  # name = "%[1]s"
				}

				resource "ibm_database" "%[2]s" {
				  resource_group_id = "${data.ibm_resource_group.test_acc.id}"	
				  name = "%[2]s"	
				  service 			= "databases-for-etcd"
				  plan              = "standard"
				  location          = "us-south"
				  adminpassword     = "password12"
				  members_memory_allocation_mb = 3072
  				  members_disk_allocation_mb   = 61440
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

func testAccCheckIBMDatabaseInstance_Etcd_fullyspecified(databaseResourceGroup string, name string) string {
	return fmt.Sprintf(`
				data "ibm_resource_group" "test_acc" {
				  is_default = true
				  # name = "%[1]s"
				}

				resource "ibm_database" "%[2]s" {
				  resource_group_id = "${data.ibm_resource_group.test_acc.id}"	
				  name = "%[2]s"	
				  service 			= "databases-for-etcd"
				  plan              = "standard"
				  location          = "us-south"
				  adminpassword     = "password12"
				  members_memory_allocation_mb = 6144
  				  members_disk_allocation_mb   = 64512
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

func testAccCheckIBMDatabaseInstance_Etcd_reduced(databaseResourceGroup string, name string) string {
	return fmt.Sprintf(`
				data "ibm_resource_group" "test_acc" {
				  is_default = true
				  # name = "%[1]s"
				}

				resource "ibm_database" "%[2]s" {
				  resource_group_id = "${data.ibm_resource_group.test_acc.id}"	
				  name = "%[2]s"	
				  service 			= "databases-for-etcd"
				  plan              = "standard"
				  location          = "us-south"
				  adminpassword     = "password12"
				  members_memory_allocation_mb = 3072
  				  members_disk_allocation_mb   = 64512

				}`, databaseResourceGroup, name)
}

func testAccCheckIBMDatabaseInstance_Etcd_import(databaseResourceGroup string, name string) string {
	return fmt.Sprintf(`
				data "ibm_resource_group" "test_acc" {
				  is_default = true
				  # name = "%[1]s"
				}

				resource "ibm_database" "%[2]s" {
				  resource_group_id = "${data.ibm_resource_group.test_acc.id}"	
				  name = "%[2]s"	
				  service 			= "databases-for-etcd"
				  plan              = "standard"
				  location          = "us-south"
				}`, databaseResourceGroup, name)
}
