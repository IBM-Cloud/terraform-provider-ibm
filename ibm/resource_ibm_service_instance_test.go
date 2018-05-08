package ibm

import (
	"fmt"
	"testing"

	"strings"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"

	"github.com/IBM-Cloud/bluemix-go/api/mccp/mccpv2"
)

func TestAccIBMServiceInstance_Basic(t *testing.T) {
	var conf mccpv2.ServiceInstanceFields
	serviceName := fmt.Sprintf("terraform_%d", acctest.RandInt())
	updateName := fmt.Sprintf("terraform_%d", acctest.RandInt())

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMServiceInstanceDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMServiceInstance_basic(serviceName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMServiceInstanceExists("ibm_service_instance.service", &conf),
					resource.TestCheckResourceAttr("ibm_service_instance.service", "name", serviceName),
					resource.TestCheckResourceAttr("ibm_service_instance.service", "service", "cleardb"),
					resource.TestCheckResourceAttr("ibm_service_instance.service", "plan", "cb5"),
					resource.TestCheckResourceAttr("ibm_service_instance.service", "tags.#", "2"),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMServiceInstance_updateWithSameName(serviceName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMServiceInstanceExists("ibm_service_instance.service", &conf),
					resource.TestCheckResourceAttr("ibm_service_instance.service", "name", serviceName),
					resource.TestCheckResourceAttr("ibm_service_instance.service", "service", "cleardb"),
					resource.TestCheckResourceAttr("ibm_service_instance.service", "plan", "cb5"),
					resource.TestCheckResourceAttr("ibm_service_instance.service", "tags.#", "3"),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMServiceInstance_update(updateName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_service_instance.service", "name", updateName),
					resource.TestCheckResourceAttr("ibm_service_instance.service", "service", "cleardb"),
					resource.TestCheckResourceAttr("ibm_service_instance.service", "plan", "cb5"),
					resource.TestCheckResourceAttr("ibm_service_instance.service", "tags.#", "1"),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMServiceInstance_newServiceType(updateName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_service_instance.service", "name", updateName),
					resource.TestCheckResourceAttr("ibm_service_instance.service", "service", "cloudantNoSQLDB"),
					resource.TestCheckResourceAttr("ibm_service_instance.service", "plan", "Lite"),
					resource.TestCheckResourceAttr("ibm_service_instance.service", "tags.#", "1"),
				),
			},
		},
	})
}

func TestAccIBMServiceInstance_import(t *testing.T) {
	var conf mccpv2.ServiceInstanceFields
	serviceName := fmt.Sprintf("terraform_%d", acctest.RandInt())
	resourceName := "ibm_service_instance.service"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMServiceInstanceDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMServiceInstance_basic(serviceName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMServiceInstanceExists(resourceName, &conf),
					resource.TestCheckResourceAttr(resourceName, "name", serviceName),
					resource.TestCheckResourceAttr(resourceName, "service", "cleardb"),
					resource.TestCheckResourceAttr(resourceName, "plan", "cb5"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "2"),
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

func testAccCheckIBMServiceInstanceDestroy(s *terraform.State) error {
	cfClient, err := testAccProvider.Meta().(ClientSession).MccpAPI()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_service_instance" {
			continue
		}

		serviceGuid := rs.Primary.ID

		// Try to find the key
		_, err := cfClient.ServiceInstances().Get(serviceGuid)

		if err != nil && !strings.Contains(err.Error(), "404") {
			return fmt.Errorf("Error waiting for CF service (%s) to be destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}

func testAccCheckIBMServiceInstanceExists(n string, obj *mccpv2.ServiceInstanceFields) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		cfClient, err := testAccProvider.Meta().(ClientSession).MccpAPI()
		if err != nil {
			return err
		}
		serviceGuid := rs.Primary.ID

		service, err := cfClient.ServiceInstances().Get(serviceGuid)

		if err != nil {
			return err
		}

		*obj = *service
		return nil
	}
}

func testAccCheckIBMServiceInstance_basic(serviceName string) string {
	return fmt.Sprintf(`
		data "ibm_space" "spacedata" {
			space  = "%s"
			org    = "%s"
		}
		
		resource "ibm_service_instance" "service" {
			name              = "%s"
			space_guid        = "${data.ibm_space.spacedata.id}"
			service           = "cleardb"
			plan              = "cb5"
			tags               = ["cluster-service","cluster-bind"]
		}
	`, cfSpace, cfOrganization, serviceName)
}

func testAccCheckIBMServiceInstance_updateWithSameName(serviceName string) string {
	return fmt.Sprintf(`
		data "ibm_space" "spacedata" {
			space  = "%s"
			org    = "%s"
		}
		
		resource "ibm_service_instance" "service" {
			name              = "%s"
			space_guid        = "${data.ibm_space.spacedata.id}"
			service           = "cleardb"
			plan              = "cb5"
			tags               = ["cluster-service","cluster-bind","db"]
		}
	`, cfSpace, cfOrganization, serviceName)
}

func testAccCheckIBMServiceInstance_update(updateName string) string {
	return fmt.Sprintf(`
		data "ibm_space" "spacedata" {
			space  = "%s"
			org    = "%s"
		}

		resource "ibm_service_instance" "service" {
			name              = "%s"
			space_guid        = "${data.ibm_space.spacedata.id}"
			service           = "cleardb"
			plan              = "cb5"
			tags               = ["cluster-service"]
		}
	`, cfSpace, cfOrganization, updateName)
}

func testAccCheckIBMServiceInstance_newServiceType(updateName string) string {
	return fmt.Sprintf(`
		data "ibm_space" "spacedata" {
			space  = "%s"
			org    = "%s"
		}

		resource "ibm_service_instance" "service" {
			name              = "%s"
			space_guid        = "${data.ibm_space.spacedata.id}"
			service           = "cloudantNoSQLDB"
			plan              = "Lite"
			tags               = ["cluster-service"]
		}
	`, cfSpace, cfOrganization, updateName)
}

func TestAccIBMServiceInstance_Discovery_Basic(t *testing.T) {
	var conf mccpv2.ServiceInstanceFields
	serviceName := fmt.Sprintf("terraform_%d", acctest.RandInt())

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMServiceInstanceDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMServiceInstance_discovery_basic(serviceName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMServiceInstanceExists("ibm_service_instance.service", &conf),
					resource.TestCheckResourceAttr("ibm_service_instance.service", "name", serviceName),
					resource.TestCheckResourceAttr("ibm_service_instance.service", "service", "discovery"),
					resource.TestCheckResourceAttr("ibm_service_instance.service", "plan", "lite"),
					resource.TestCheckResourceAttr("ibm_service_instance.service", "tags.#", "2"),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMServiceInstance_discovery_update(serviceName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMServiceInstanceExists("ibm_service_instance.service", &conf),
					resource.TestCheckResourceAttr("ibm_service_instance.service", "name", serviceName),
					resource.TestCheckResourceAttr("ibm_service_instance.service", "service", "discovery"),
					resource.TestCheckResourceAttr("ibm_service_instance.service", "plan", "lite"),
					resource.TestCheckResourceAttr("ibm_service_instance.service", "tags.#", "3"),
				),
			},
		},
	})
}

func testAccCheckIBMServiceInstance_discovery_basic(serviceName string) string {
	return fmt.Sprintf(`
		data "ibm_space" "spacedata" {
			space  = "%s"
			org    = "%s"
		}
		
		resource "ibm_service_instance" "service" {
			name              = "%s"
			space_guid        = "${data.ibm_space.spacedata.id}"
			service           = "discovery"
			plan              = "lite"
			tags               = ["cluster-service","cluster-bind"]
		}
	`, cfSpace, cfOrganization, serviceName)
}

func testAccCheckIBMServiceInstance_discovery_update(serviceName string) string {
	return fmt.Sprintf(`
		data "ibm_space" "spacedata" {
			space  = "%s"
			org    = "%s"
		}
		
		resource "ibm_service_instance" "service" {
			name              = "%s"
			space_guid        = "${data.ibm_space.spacedata.id}"
			service           = "discovery"
			plan              = "lite"
			tags               = ["cluster-service","cluster-bind","db"]
		}
	`, cfSpace, cfOrganization, serviceName)
}
