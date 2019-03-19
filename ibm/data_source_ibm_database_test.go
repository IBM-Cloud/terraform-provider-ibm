package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccIBMDatabaseDataSource_basic(t *testing.T) {
	//t.Parallel()
	databaseResourceGroup := "default"
	var databaseInstanceOne string
	testName := fmt.Sprintf("tf_test_acc_%s", acctest.RandString(16))
	dataName := "data.ibm_database." + testName
	resourceName := "ibm_database." + testName

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config:  testAccCheckIBMDatabaseDataSourceConfig(databaseResourceGroup, testName),
				Destroy: false,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMDatabaseInstanceExists(resourceName, &databaseInstanceOne),
					testAccCheckIBMDatabaseInstanceExists(dataName, &databaseInstanceOne),
					resource.TestCheckResourceAttr(dataName, "name", testName),
					resource.TestCheckResourceAttr(dataName, "service", "databases-for-postgresql"),
					resource.TestCheckResourceAttr(dataName, "plan", "standard"),
					resource.TestCheckResourceAttr(dataName, "location", "us-south"),
					resource.TestCheckResourceAttr(dataName, "adminuser", "admin"),
					resource.TestCheckResourceAttr(dataName, "members_memory_allocation_mb", "2048"),
					resource.TestCheckResourceAttr(dataName, "members_disk_allocation_mb", "10240"),
					resource.TestCheckResourceAttr(dataName, "whitelist.#", "0"),
					resource.TestCheckResourceAttr(dataName, "connectionstrings.#", "1"),
					resource.TestCheckResourceAttr(dataName, "connectionstrings.0.name", "admin"),
					resource.TestCheckResourceAttr(dataName, "connectionstrings.0.hosts.#", "1"),
					resource.TestCheckResourceAttr(dataName, "connectionstrings.0.scheme", "postgres"),
					resource.TestCheckResourceAttr(dataName, "tags.#", "1"),
				),
			},
		},
	})
}

func testAccCheckIBMDatabaseDataSourceConfig(databaseResourceGroup string, name string) string {
	return fmt.Sprintf(`
				data "ibm_resource_group" "test_acc" {
				  is_default = true
				  # name = "%[1]s"
				}

				data "ibm_database" "%[2]s" {
				  resource_group_id = "${data.ibm_resource_group.test_acc.id}"	
				  name = "${ibm_database.%[2]s.name}"	
				}

				resource "ibm_database" "%[2]s" {
				  resource_group_id = "${data.ibm_resource_group.test_acc.id}"	
				  name = "%[2]s"	
				  service 			= "databases-for-postgresql"
				  plan              = "standard"
				  location          = "us-south"
				  tags = ["one:two"]
				}

				`, databaseResourceGroup, name)
}
