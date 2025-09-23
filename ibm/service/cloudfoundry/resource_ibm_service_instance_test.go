// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package cloudfoundry_test

import (
	"fmt"
	"strings"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/Mavrickk3/bluemix-go/api/mccp/mccpv2"
)

func TestAccIBMServiceInstance_Basic(t *testing.T) {
	t.Skip()
	var conf mccpv2.ServiceInstanceFields
	serviceName := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))
	updateName := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMServiceInstanceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMServiceInstance_basic(serviceName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMServiceInstanceExists("ibm_service_instance.service", &conf),
					resource.TestCheckResourceAttr("ibm_service_instance.service", "name", serviceName),
					resource.TestCheckResourceAttr("ibm_service_instance.service", "service", "speech_to_text"),
					resource.TestCheckResourceAttr("ibm_service_instance.service", "plan", "lite"),
					resource.TestCheckResourceAttr("ibm_service_instance.service", "tags.#", "2"),
				),
			},
			{
				Config: testAccCheckIBMServiceInstance_updateWithSameName(serviceName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMServiceInstanceExists("ibm_service_instance.service", &conf),
					resource.TestCheckResourceAttr("ibm_service_instance.service", "name", serviceName),
					resource.TestCheckResourceAttr("ibm_service_instance.service", "service", "speech_to_text"),
					resource.TestCheckResourceAttr("ibm_service_instance.service", "plan", "lite"),
					resource.TestCheckResourceAttr("ibm_service_instance.service", "tags.#", "3"),
				),
			},
			{
				Config: testAccCheckIBMServiceInstance_update(updateName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_service_instance.service", "name", updateName),
					resource.TestCheckResourceAttr("ibm_service_instance.service", "service", "speech_to_text"),
					resource.TestCheckResourceAttr("ibm_service_instance.service", "plan", "lite"),
					resource.TestCheckResourceAttr("ibm_service_instance.service", "tags.#", "1"),
				),
			},
			{
				Config: testAccCheckIBMServiceInstance_newServiceType(updateName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_service_instance.service", "name", updateName),
					resource.TestCheckResourceAttr("ibm_service_instance.service", "service", "speech_to_text"),
					resource.TestCheckResourceAttr("ibm_service_instance.service", "plan", "lite"),
					resource.TestCheckResourceAttr("ibm_service_instance.service", "tags.#", "1"),
				),
			},
		},
	})
}

func TestAccIBMServiceInstance_import(t *testing.T) {
	t.Skip()
	var conf mccpv2.ServiceInstanceFields
	serviceName := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))
	resourceName := "ibm_service_instance.service"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMServiceInstanceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMServiceInstance_basic(serviceName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMServiceInstanceExists(resourceName, &conf),
					resource.TestCheckResourceAttr(resourceName, "name", serviceName),
					resource.TestCheckResourceAttr(resourceName, "service", "speech_to_text"),
					resource.TestCheckResourceAttr(resourceName, "plan", "lite"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "2"),
				),
			},
			{
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
	cfClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).MccpAPI()
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

		if err == nil {
			return fmt.Errorf("CF service still exists: %s", rs.Primary.ID)
		} else if !strings.Contains(err.Error(), "404") {
			return fmt.Errorf("[ERROR] Error waiting for CF service (%s) to be destroyed: %s", rs.Primary.ID, err)
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

		cfClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).MccpAPI()
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
		space = "%s"
		org   = "%s"
	  }
	  
	  resource "ibm_service_instance" "service" {
		name       = "%s"
		space_guid = data.ibm_space.spacedata.id
		service    = "speech_to_text"
		plan       = "lite"
		tags       = ["cluster-service", "cluster-bind"]
	  }
	`, acc.CfSpace, acc.CfOrganization, serviceName)
}

func testAccCheckIBMServiceInstance_updateWithSameName(serviceName string) string {
	return fmt.Sprintf(`
	data "ibm_space" "spacedata" {
		space = "%s"
		org   = "%s"
	  }
	  
	  resource "ibm_service_instance" "service" {
		name       = "%s"
		space_guid = data.ibm_space.spacedata.id
		service    = "speech_to_text"
		plan       = "lite"
		tags       = ["cluster-service", "cluster-bind", "db"]
	  }
	`, acc.CfSpace, acc.CfOrganization, serviceName)
}

func testAccCheckIBMServiceInstance_update(updateName string) string {
	return fmt.Sprintf(`
	data "ibm_space" "spacedata" {
		space = "%s"
		org   = "%s"
	  }
	  
	  resource "ibm_service_instance" "service" {
		name       = "%s"
		space_guid = data.ibm_space.spacedata.id
		service    = "speech_to_text"
		plan       = "lite"
		tags       = ["cluster-service"]
	  }
	`, acc.CfSpace, acc.CfOrganization, updateName)
}

func testAccCheckIBMServiceInstance_newServiceType(updateName string) string {
	return fmt.Sprintf(`
	data "ibm_space" "spacedata" {
		space = "%s"
		org   = "%s"
	  }
	  
	  resource "ibm_service_instance" "service" {
		name       = "%s"
		space_guid = data.ibm_space.spacedata.id
		service    = "speech_to_text"
		plan       = "lite"
		tags       = ["cluster-service"]
	  }
	`, acc.CfSpace, acc.CfOrganization, updateName)
}

func TestAccIBMServiceInstance_Discovery_Basic(t *testing.T) {
	t.Skip()
	var conf mccpv2.ServiceInstanceFields
	serviceName := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMServiceInstanceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMServiceInstance_discovery_basic(serviceName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMServiceInstanceExists("ibm_service_instance.service", &conf),
					resource.TestCheckResourceAttr("ibm_service_instance.service", "name", serviceName),
					resource.TestCheckResourceAttr("ibm_service_instance.service", "service", "discovery"),
					resource.TestCheckResourceAttr("ibm_service_instance.service", "plan", "lite"),
					resource.TestCheckResourceAttr("ibm_service_instance.service", "tags.#", "2"),
				),
			},
			{
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
		space = "%s"
		org   = "%s"
	  }
	  
	  resource "ibm_service_instance" "service" {
		name       = "%s"
		space_guid = data.ibm_space.spacedata.id
		service    = "discovery"
		plan       = "lite"
		tags       = ["cluster-service", "cluster-bind"]
	  }
	`, acc.CfSpace, acc.CfOrganization, serviceName)
}

func testAccCheckIBMServiceInstance_discovery_update(serviceName string) string {
	return fmt.Sprintf(`
	data "ibm_space" "spacedata" {
		space = "%s"
		org   = "%s"
	  }
	  
	  resource "ibm_service_instance" "service" {
		name       = "%s"
		space_guid = data.ibm_space.spacedata.id
		service    = "discovery"
		plan       = "lite"
		tags       = ["cluster-service", "cluster-bind", "db"]
	  }
	`, acc.CfSpace, acc.CfOrganization, serviceName)
}
