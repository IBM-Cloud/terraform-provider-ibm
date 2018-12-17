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

func TestAccIBMCisInstance_Basic(t *testing.T) {
	var conf models.ServiceInstance
	serviceName := fmt.Sprintf("terraform_%d", acctest.RandInt())

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMCisInstanceDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCisInstance_basic(serviceName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCisInstanceExists("ibm_cis.instance", conf),
					resource.TestCheckResourceAttr("ibm_cis.instance", "name", serviceName),
					resource.TestCheckResourceAttr("ibm_cis.instance", "service", "internet-svcs"),
					resource.TestCheckResourceAttr("ibm_cis.instance", "plan", "standard"),
					resource.TestCheckResourceAttr("ibm_cis.instance", "location", "global"),
				),
			},
		},
	})
}

func TestAccIBMCisInstance_import(t *testing.T) {
	var conf models.ServiceInstance
	serviceName := fmt.Sprintf("terraform_%d", acctest.RandInt())
	resourceName := "ibm_cis.instance"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMCisInstanceDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCisInstance_basic(serviceName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCisInstanceExists(resourceName, conf),
					resource.TestCheckResourceAttr(resourceName, "name", serviceName),
					resource.TestCheckResourceAttr(resourceName, "service", "internet-svcs"),
					resource.TestCheckResourceAttr(resourceName, "plan", "standard"),
					resource.TestCheckResourceAttr(resourceName, "location", "global"),
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

func TestAccIBMCisInstance_with_resource_group(t *testing.T) {
	var conf models.ServiceInstance
	serviceName := fmt.Sprintf("terraform_%d", acctest.RandInt())
	resourceName := "ibm_cis.instance"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMCisInstanceDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCisInstance_with_resource_group(serviceName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCisInstanceExists(resourceName, conf),
					resource.TestCheckResourceAttr(resourceName, "name", serviceName),
					resource.TestCheckResourceAttr(resourceName, "service", "internet-svcs"),
					resource.TestCheckResourceAttr(resourceName, "plan", "standard"),
					resource.TestCheckResourceAttr(resourceName, "location", "global"),
				),
			},
		},
	})
}

func testAccCheckIBMCisInstanceDestroy(s *terraform.State) error {
	rsContClient, err := testAccProvider.Meta().(ClientSession).ResourceControllerAPI()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_cis" {
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

func testAccCheckIBMCisInstanceExists(n string, obj models.ServiceInstance) resource.TestCheckFunc {

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

func testAccCheckIBMCisInstance_basic(serviceName string) string {
	return fmt.Sprintf(`
		
		resource "ibm_cis" "instance" {
			name              = "%s"		
			plan              = "standard"
			location          = "global"
			
			timeouts {
				create = "15m"
				update = "15m"
				delete = "15m"
			  }
		}
	`, serviceName)
}

func testAccCheckIBMCisInstance_with_resource_group(serviceName string) string {
	return fmt.Sprintf(`

		data "ibm_resource_group" "group" {
			name = "default"
		}
		
		resource "ibm_cis" "instance" {
			name              = "%s"		
			plan              = "standard"
			location          = "global"
			resource_group_id = "${data.ibm_resource_group.group.id}"
			
		}
	`, serviceName)
}
