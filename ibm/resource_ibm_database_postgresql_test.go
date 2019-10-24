package ibm

import (
	"fmt"
	"regexp"
	"strings"
	"testing"
	"time"

	"github.com/IBM-Cloud/bluemix-go/bmxerror"
	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccIBMDatabaseInstance_Postgres_Basic(t *testing.T) {
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
				Config: testAccCheckIBMDatabaseInstance_Postgres_basic(databaseResourceGroup, testName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMDatabaseInstanceExists(name, &databaseInstanceOne),
					resource.TestCheckResourceAttr(name, "name", testName),
					resource.TestCheckResourceAttr(name, "service", "databases-for-postgresql"),
					resource.TestCheckResourceAttr(name, "plan", "standard"),
					resource.TestCheckResourceAttr(name, "location", "us-south"),
					resource.TestCheckResourceAttr(name, "adminuser", "admin"),
					resource.TestCheckResourceAttr(name, "members_memory_allocation_mb", "2048"),
					resource.TestCheckResourceAttr(name, "members_disk_allocation_mb", "10240"),
					resource.TestCheckResourceAttr(name, "whitelist.#", "1"),
					resource.TestCheckResourceAttr(name, "users.#", "1"),
					resource.TestCheckResourceAttr(name, "connectionstrings.#", "2"),
					resource.TestCheckResourceAttr(name, "connectionstrings.1.name", "admin"),
					resource.TestMatchResourceAttr(name, "connectionstrings.1.certname", regexp.MustCompile("[-a-z0-9]*")),
					resource.TestMatchResourceAttr(name, "connectionstrings.1.certbase64", regexp.MustCompile("^(?:[A-Za-z0-9+/]{4})*(?:[A-Za-z0-9+/]{2}==|[A-Za-z0-9+/]{3}=)?$")),
					resource.TestCheckResourceAttr(name, "tags.#", "1"),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMDatabaseInstance_Postgres_fullyspecified(databaseResourceGroup, testName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMDatabaseInstanceExists(name, &databaseInstanceOne),
					resource.TestCheckResourceAttr(name, "name", testName),
					resource.TestCheckResourceAttr(name, "service", "databases-for-postgresql"),
					resource.TestCheckResourceAttr(name, "plan", "standard"),
					resource.TestCheckResourceAttr(name, "location", "us-south"),
					resource.TestCheckResourceAttr(name, "members_memory_allocation_mb", "4096"),
					resource.TestCheckResourceAttr(name, "members_disk_allocation_mb", "14336"),
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
			resource.TestStep{
				Config: testAccCheckIBMDatabaseInstance_Postgres_reduced(databaseResourceGroup, testName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMDatabaseInstanceExists(name, &databaseInstanceOne),
					resource.TestCheckResourceAttr(name, "name", testName),
					resource.TestCheckResourceAttr(name, "service", "databases-for-postgresql"),
					resource.TestCheckResourceAttr(name, "plan", "standard"),
					resource.TestCheckResourceAttr(name, "location", "us-south"),
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

// TestAccIBMDatabaseInstance_CreateAfterManualDestroy not required as tested by resource_instance tests

func TestAccIBMDatabaseInstance_Postgres_import(t *testing.T) {
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
				Config: testAccCheckIBMDatabaseInstance_Postgres_import(databaseResourceGroup, serviceName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMDatabaseInstanceExists(resourceName, &databaseInstanceOne),
					resource.TestCheckResourceAttr(resourceName, "name", serviceName),
					resource.TestCheckResourceAttr(resourceName, "service", "databases-for-postgresql"),
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

func testAccCheckIBMDatabaseInstanceDestroy(s *terraform.State) error {
	rsContClient, err := testAccProvider.Meta().(ClientSession).ResourceControllerAPI()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_database" {
			continue
		}

		instanceID := rs.Primary.ID

		_, err := rsContClient.ResourceServiceInstance().GetInstance(instanceID)

		if err != nil && !strings.Contains(err.Error(), "404") {
			return fmt.Errorf("Error checking if instance (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}
	return nil
}

func testAccDatabaseInstanceManuallyDelete(tfDatabaseId *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		_ = testAccDatabaseInstanceManuallyDeleteUnwrapped(s, tfDatabaseId)
		return nil
	}
}

func testAccDatabaseInstanceManuallyDeleteUnwrapped(s *terraform.State, tfDatabaseId *string) error {
	rsConClient, err := testAccProvider.Meta().(ClientSession).ResourceControllerAPI()
	if err != nil {
		return err
	}
	instance := *tfDatabaseId
	var instanceId string
	if strings.HasPrefix(instance, "crn") {
		instanceId = instance
	} else {
		_, instanceId, _ = convertTftoCisTwoVar(instance)
	}
	err = rsConClient.ResourceServiceInstance().DeleteInstance(instanceId, true)
	if err != nil {
		return fmt.Errorf("Error deleting resource instance: %s", err)
	}

	_ = &resource.StateChangeConf{
		Pending: []string{databaseInstanceProgressStatus, databaseInstanceInactiveStatus, databaseInstanceSuccessStatus},
		Target:  []string{databaseInstanceRemovedStatus},
		Refresh: func() (interface{}, string, error) {
			instance, err := rsConClient.ResourceServiceInstance().GetInstance(instanceId)
			if err != nil {
				if apiErr, ok := err.(bmxerror.RequestFailure); ok && apiErr.StatusCode() == 404 {
					return instance, databaseInstanceSuccessStatus, nil
				}
				return nil, "", err
			}
			if instance.State == databaseInstanceFailStatus {
				return instance, instance.State, fmt.Errorf("The resource instance %s failed to delete: %v", instanceId, err)
			}
			return instance, instance.State, nil
		},
		Timeout:    90 * time.Second,
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}
	if err != nil {
		return fmt.Errorf(
			"Error waiting for resource instance (%s) to be deleted: %s", instanceId, err)
	}
	return nil
}

func testAccCheckIBMDatabaseInstanceExists(n string, tfDatabaseId *string) resource.TestCheckFunc {

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
			if strings.Contains(err.Error(), "Object not found") ||
				strings.Contains(err.Error(), "status code: 404") {
				*tfDatabaseId = ""
				return nil
			}
			return fmt.Errorf("Error retrieving resource instance: %s", err)
		}
		if strings.Contains(instance.State, "removed") {
			*tfDatabaseId = ""
			return nil
		}

		*tfDatabaseId = instanceID
		return nil
	}
}

func testAccCheckIBMDatabaseInstance_Postgres_basic(databaseResourceGroup string, name string) string {
	return fmt.Sprintf(`
				data "ibm_resource_group" "test_acc" {
				  name = "%[1]s"
				}

				resource "ibm_database" "%[2]s" {
				  resource_group_id = "${data.ibm_resource_group.test_acc.id}"	
				  name = "%[2]s"	
				  service 			= "databases-for-postgresql"
				  plan              = "standard"
				  location          = "us-south"
				  adminpassword     = "password12"
				  members_memory_allocation_mb = 2048
  				  members_disk_allocation_mb   = 10240
  				  tags = ["one:two"]
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

func testAccCheckIBMDatabaseInstance_Postgres_fullyspecified(databaseResourceGroup string, name string) string {
	return fmt.Sprintf(`
				data "ibm_resource_group" "test_acc" {
				  name = "%[1]s"
				}

				resource "ibm_database" "%[2]s" {
				  resource_group_id = "${data.ibm_resource_group.test_acc.id}"	
				  name = "%[2]s"	
				  service 			= "databases-for-postgresql"
				  plan              = "standard"
				  location          = "us-south"
				  adminpassword     = "password12"
				  members_memory_allocation_mb = 4096
  				  members_disk_allocation_mb   = 14336
  				  tags = ["one:two"]
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

func testAccCheckIBMDatabaseInstance_Postgres_reduced(databaseResourceGroup string, name string) string {
	return fmt.Sprintf(`
				data "ibm_resource_group" "test_acc" {
				  name = "%[1]s"
				}

				resource "ibm_database" "%[2]s" {
				  resource_group_id = "${data.ibm_resource_group.test_acc.id}"	
				  name = "%[2]s"	
				  service 			= "databases-for-postgresql"
				  plan              = "standard"
				  location          = "us-south"
				  adminpassword     = "password12"
				  members_memory_allocation_mb = 2048
  				  members_disk_allocation_mb   = 14336
  				  tags = ["one:two"]
				}`, databaseResourceGroup, name)
}

func testAccCheckIBMDatabaseInstance_Postgres_import(databaseResourceGroup string, name string) string {
	return fmt.Sprintf(`
				data "ibm_resource_group" "test_acc" {
				  is_default = true
				  # name = "%[1]s"
				}

				resource "ibm_database" "%[2]s" {
				  resource_group_id = "${data.ibm_resource_group.test_acc.id}"	
				  name = "%[2]s"	
				  service 			= "databases-for-postgresql"
				  plan              = "standard"
				  location          = "us-south"
				}`, databaseResourceGroup, name)
}
