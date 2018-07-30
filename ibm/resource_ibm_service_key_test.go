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

func TestAccIBMServiceKey_Basic(t *testing.T) {
	var conf mccpv2.ServiceKeyFields
	serviceName := fmt.Sprintf("terraform_%d", acctest.RandInt())
	serviceKey := fmt.Sprintf("terraform_%d", acctest.RandInt())

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMServiceKeyDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMServiceKey_basic(serviceName, serviceKey),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMServiceKeyExists("ibm_service_key.serviceKey", &conf),
					resource.TestCheckResourceAttr("ibm_service_key.serviceKey", "name", serviceKey),
					resource.TestCheckResourceAttr("ibm_service_key.serviceKey", "credentials.%", "3"),
				),
			},
		},
	})
}

func TestAccIBMServiceKey_With_Tags(t *testing.T) {
	var conf mccpv2.ServiceKeyFields
	serviceName := fmt.Sprintf("terraform_%d", acctest.RandInt())
	serviceKey := fmt.Sprintf("terraform_%d", acctest.RandInt())

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMServiceKeyDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMServiceKey_with_tags(serviceName, serviceKey),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMServiceKeyExists("ibm_service_key.serviceKey", &conf),
					resource.TestCheckResourceAttr("ibm_service_key.serviceKey", "name", serviceKey),
					resource.TestCheckResourceAttr("ibm_service_key.serviceKey", "tags.#", "1"),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMServiceKey_with_updated_tags(serviceName, serviceKey),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMServiceKeyExists("ibm_service_key.serviceKey", &conf),
					resource.TestCheckResourceAttr("ibm_service_key.serviceKey", "tags.#", "2"),
				),
			},
		},
	})
}

func TestAccIBMServiceKey_Parameters(t *testing.T) {
	var conf mccpv2.ServiceKeyFields
	serviceName := fmt.Sprintf("terraform_%d", acctest.RandInt())
	serviceKey := fmt.Sprintf("terraform_%d", acctest.RandInt())

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMServiceKeyDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMServiceKey_parameters(serviceName, serviceKey),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMServiceKeyExists("ibm_service_key.serviceKey", &conf),
					resource.TestCheckResourceAttr("ibm_service_key.serviceKey", "name", serviceKey),
					resource.TestCheckResourceAttr("ibm_service_key.serviceKey", "parameters.%", "1"),
					resource.TestCheckResourceAttr("ibm_service_key.serviceKey", "credentials.%", "9"),
				),
			},
		},
	})
}

func testAccCheckIBMServiceKeyExists(n string, obj *mccpv2.ServiceKeyFields) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		cfClient, err := testAccProvider.Meta().(ClientSession).MccpAPI()
		if err != nil {
			return err
		}
		serviceKeyGuid := rs.Primary.ID

		serviceKey, err := cfClient.ServiceKeys().Get(serviceKeyGuid)
		if err != nil {
			return err
		}

		*obj = *serviceKey
		return nil
	}
}

func testAccCheckIBMServiceKeyDestroy(s *terraform.State) error {
	cfClient, err := testAccProvider.Meta().(ClientSession).MccpAPI()
	if err != nil {
		return err
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_service_key" {
			continue
		}

		serviceKeyGuid := rs.Primary.ID

		// Try to find the key
		_, err := cfClient.ServiceKeys().Get(serviceKeyGuid)

		if err != nil && !strings.Contains(err.Error(), "404") {
			return fmt.Errorf("Error waiting for CF service key (%s) to be destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}

func testAccCheckIBMServiceKey_basic(serviceName, serviceKey string) string {
	return fmt.Sprintf(`
		
		data "ibm_space" "spacedata" {
			space  = "%s"
			org    = "%s"
		}
		
		resource "ibm_service_instance" "service" {
			name              = "%s"
			space_guid        = "${data.ibm_space.spacedata.id}"
			service           = "speech_to_text"
			plan              = "lite"
			tags               = ["cluster-service","cluster-bind"]
		}

		resource "ibm_service_key" "serviceKey" {
			name = "%s"
			service_instance_guid = "${ibm_service_instance.service.id}"
		}
	`, cfSpace, cfOrganization, serviceName, serviceKey)
}

func testAccCheckIBMServiceKey_with_tags(serviceName, serviceKey string) string {
	return fmt.Sprintf(`
		
		data "ibm_space" "spacedata" {
			space  = "%s"
			org    = "%s"
		}
		
		resource "ibm_service_instance" "service" {
			name              = "%s"
			space_guid        = "${data.ibm_space.spacedata.id}"
			service           = "speech_to_text"
			plan              = "lite"
			tags               = ["cluster-service","cluster-bind"]
		}

		resource "ibm_service_key" "serviceKey" {
			name = "%s"
			service_instance_guid = "${ibm_service_instance.service.id}"
			tags				  = ["one"]	
		}
	`, cfSpace, cfOrganization, serviceName, serviceKey)
}

func testAccCheckIBMServiceKey_with_updated_tags(serviceName, serviceKey string) string {
	return fmt.Sprintf(`
		
		data "ibm_space" "spacedata" {
			space  = "%s"
			org    = "%s"
		}
		
		resource "ibm_service_instance" "service" {
			name              = "%s"
			space_guid        = "${data.ibm_space.spacedata.id}"
			service           = "speech_to_text"
			plan              = "lite"
			tags               = ["cluster-service","cluster-bind"]
		}

		resource "ibm_service_key" "serviceKey" {
			name = "%s"
			service_instance_guid = "${ibm_service_instance.service.id}"
			tags				  = ["one", "two"]	
		}
	`, cfSpace, cfOrganization, serviceName, serviceKey)
}

func testAccCheckIBMServiceKey_parameters(serviceName, serviceKey string) string {
	return fmt.Sprintf(`
		
		data "ibm_space" "spacedata" {
			space  = "%s"
			org    = "%s"
		}
		
		resource "ibm_service_instance" "service" {
			name              = "%s"
			space_guid        = "${data.ibm_space.spacedata.id}"
			service           = "cloud-object-storage"
			plan              = "Lite"
			tags              = ["cluster-service","cluster-bind"]
		}

		resource "ibm_service_key" "serviceKey" {
			name = "%s"
			service_instance_guid = "${ibm_service_instance.service.id}"
			parameters        = {"HMAC" = true}
		}
	`, cfSpace, cfOrganization, serviceName, serviceKey)
}
