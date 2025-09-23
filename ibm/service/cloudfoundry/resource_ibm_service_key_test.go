// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package cloudfoundry_test

import (
	"fmt"
	"strings"
	"testing"

	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/conns"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/Mavrickk3/bluemix-go/api/mccp/mccpv2"
)

func TestAccIBMServiceKey_Basic(t *testing.T) {
	t.Skip()
	var conf mccpv2.ServiceKeyFields
	serviceName := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))
	serviceKey := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMServiceKeyDestroy,
		Steps: []resource.TestStep{
			{
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
	t.Skip()
	var conf mccpv2.ServiceKeyFields
	serviceName := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))
	serviceKey := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMServiceKeyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMServiceKey_with_tags(serviceName, serviceKey),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMServiceKeyExists("ibm_service_key.serviceKey", &conf),
					resource.TestCheckResourceAttr("ibm_service_key.serviceKey", "name", serviceKey),
					resource.TestCheckResourceAttr("ibm_service_key.serviceKey", "tags.#", "1"),
				),
			},
			{
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
	t.Skip()
	var conf mccpv2.ServiceKeyFields
	serviceName := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))
	serviceKey := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMServiceKeyDestroy,
		Steps: []resource.TestStep{
			{
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

		cfClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).MccpAPI()
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
	cfClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).MccpAPI()
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

		if err == nil {
			return fmt.Errorf("CF service key still exists: %s", rs.Primary.ID)
		} else if !strings.Contains(err.Error(), "404") {
			return fmt.Errorf("[ERROR] Error waiting for CF service key (%s) to be destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}

func testAccCheckIBMServiceKey_basic(serviceName, serviceKey string) string {
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
	  
	  resource "ibm_service_key" "serviceKey" {
		name                  = "%s"
		service_instance_guid = ibm_service_instance.service.id
	  }
	`, acc.CfSpace, acc.CfOrganization, serviceName, serviceKey)
}

func testAccCheckIBMServiceKey_with_tags(serviceName, serviceKey string) string {
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
	  
	  resource "ibm_service_key" "serviceKey" {
		name                  = "%s"
		service_instance_guid = ibm_service_instance.service.id
		tags                  = ["one"]
	  }
	  
	`, acc.CfSpace, acc.CfOrganization, serviceName, serviceKey)
}

func testAccCheckIBMServiceKey_with_updated_tags(serviceName, serviceKey string) string {
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
	  
	  resource "ibm_service_key" "serviceKey" {
		name                  = "%s"
		service_instance_guid = ibm_service_instance.service.id
		tags                  = ["one", "two"]
	  }	  
	`, acc.CfSpace, acc.CfOrganization, serviceName, serviceKey)
}

func testAccCheckIBMServiceKey_parameters(serviceName, serviceKey string) string {
	return fmt.Sprintf(`
		
	data "ibm_space" "spacedata" {
		space = "%s"
		org   = "%s"
	  }
	  
	  resource "ibm_service_instance" "service" {
		name       = "%s"
		space_guid = data.ibm_space.spacedata.id
		service    = "cloud-object-storage"
		plan       = "Lite"
		tags       = ["cluster-service", "cluster-bind"]
	  }
	  
	  resource "ibm_service_key" "serviceKey" {
		name                  = "%s"
		service_instance_guid = ibm_service_instance.service.id
		parameters = {
		  "HMAC" = true
		}
	  }
	`, acc.CfSpace, acc.CfOrganization, serviceName, serviceKey)
}
