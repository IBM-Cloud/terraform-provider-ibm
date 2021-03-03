// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccIBMServiceInstanceDataSource_basic(t *testing.T) {
	t.Skip()
	serviceName := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))
	serviceKey := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: setupServiceInstanceConfig(serviceName, serviceKey),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_service_key.servicekey", "credentials.%", "3"),
					resource.TestCheckResourceAttr("ibm_service_instance.service", "service_keys.#", "0"),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMServiceInstanceDataSourceConfig(serviceName, serviceKey),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ibm_service_instance.testacc_ds_service_instance", "name", serviceName),
					resource.TestCheckResourceAttr("data.ibm_service_instance.testacc_ds_service_instance", "service_keys.#", "1"),
					resource.TestCheckResourceAttr("data.ibm_service_instance.testacc_ds_service_instance", "service_keys.0.credentials.%", "3"),
					resource.TestCheckResourceAttr("data.ibm_service_instance.testacc_ds_service_instance", "service_keys.0.name", serviceKey),
				),
			},
		},
	})
}

func setupServiceInstanceConfig(serviceName, serviceKey string) string {
	return fmt.Sprintf(`
	data "ibm_space" "spacedata" {
		org   = "%s"
		space = "%s"
	  }
	  
	  resource "ibm_service_instance" "service" {
		name       = "%s"
		space_guid = data.ibm_space.spacedata.id
		service    = "speech_to_text"
		plan       = "lite"
		tags       = ["cluster-service", "cluster-bind"]
	  }
	  
	  resource "ibm_service_key" "servicekey" {
		name                  = "%s"
		service_instance_guid = ibm_service_instance.service.id
	  }
	
`, cfOrganization, cfSpace, serviceName, serviceKey)

}

func testAccCheckIBMServiceInstanceDataSourceConfig(serviceName, serviceKey string) string {
	return fmt.Sprintf(`
	data "ibm_space" "spacedata" {
		org   = "%s"
		space = "%s"
	  }
	  
	  resource "ibm_service_instance" "service" {
		name       = "%s"
		space_guid = data.ibm_space.spacedata.id
		service    = "speech_to_text"
		plan       = "lite"
		tags       = ["cluster-service", "cluster-bind"]
	  }
	  
	  resource "ibm_service_key" "servicekey" {
		name                  = "%s"
		service_instance_guid = ibm_service_instance.service.id
	  }
	  
	  data "ibm_service_instance" "testacc_ds_service_instance" {
		name       = ibm_service_instance.service.name
		space_guid = data.ibm_space.spacedata.id
	  }
`, cfOrganization, cfSpace, serviceName, serviceKey)

}
