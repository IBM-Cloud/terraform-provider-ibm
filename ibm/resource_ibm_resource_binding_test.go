// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/IBM/platform-services-go-sdk/resourcecontrollerv2"
)

func TestAccIbmResourceBindingBasic(t *testing.T) {
	var conf resourcecontrollerv2.ResourceBinding
	name := fmt.Sprintf("tf_binding_name_%d", acctest.RandIntRange(10, 100))
	nameUpdate := fmt.Sprintf("tf_binding_nameUpdate_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheckResourceAliasBinding(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIbmResourceBindingDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmResourceBindingConfig(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmResourceBindingExists("ibm_resource_binding.resource_binding", conf),
					resource.TestCheckResourceAttrSet("ibm_resource_binding.resource_binding", "source"),
					resource.TestCheckResourceAttrSet("ibm_resource_binding.resource_binding", "target"),
					resource.TestCheckResourceAttr("ibm_resource_binding.resource_binding", "name", name),
					resource.TestCheckResourceAttr("ibm_resource_binding.resource_binding", "role", "Writer"),
				),
			},
			resource.TestStep{
				Config: testAccCheckIbmResourceBindingConfig(nameUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_resource_binding.resource_binding", "name", nameUpdate),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_resource_binding.resource_binding",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIbmResourceBindingConfig(name string) string {
	return fmt.Sprintf(`
		data "ibm_org" "org" {
			name = "%s"
		}
		data "ibm_space" "space" {
			org = data.ibm_org.org.name
			name = "%s"
		}
		resource "ibm_resource_instance" "instance" {
			name     = "tf_test_service"
			service  = "cloud-object-storage"
			plan     = "lite"
			location = "global"
		}
		resource "ibm_resource_alias" "alias" {
			name = "tf_test_alias"
			source = ibm_resource_instance.instance.guid
			target = "crn:v1:bluemix:public:cf:us-south:o/${data.ibm_org.org.id}::cf-space:${data.ibm_space.space.id}"
		}
		resource "ibm_app" "app" {
			name              = "tf_test_app"
			space_guid        = data.ibm_space.space.id
			app_path          = "test-fixtures/app1.zip"
			wait_time_minutes = 90
			buildpack         = "sdk-for-nodejs"
			// service_instance_guid = [ibm_resource_alias.alias.region_instance_id]
		}
		resource "ibm_resource_binding" "resource_binding" {
			source = ibm_resource_alias.alias.guid
			target = "crn:v1:bluemix:public:cf:us-south:s/${ibm_app.app.space_guid}::cf-application:${ibm_app.app.id}"
			name = "%s"
			role = "Writer"
		}
	`, cfOrganization, cfSpace, name)
}

func testAccCheckIbmResourceBindingExists(n string, obj resourcecontrollerv2.ResourceBinding) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		resourceControllerClient, err := testAccProvider.Meta().(ClientSession).ResourceControllerV2API()
		if err != nil {
			return err
		}

		getResourceBindingOptions := &resourcecontrollerv2.GetResourceBindingOptions{}

		getResourceBindingOptions.SetID(rs.Primary.ID)

		resourceBinding, _, err := resourceControllerClient.GetResourceBinding(getResourceBindingOptions)
		if err != nil {
			return err
		}

		obj = *resourceBinding
		return nil
	}
}

func testAccCheckIbmResourceBindingDestroy(s *terraform.State) error {
	resourceControllerClient, err := testAccProvider.Meta().(ClientSession).ResourceControllerV2API()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_resource_binding" {
			continue
		}

		getResourceBindingOptions := &resourcecontrollerv2.GetResourceBindingOptions{}

		getResourceBindingOptions.SetID(rs.Primary.ID)

		// Try to find the key
		binding, response, err := resourceControllerClient.GetResourceBinding(getResourceBindingOptions)

		if err == nil {
			if *binding.State == "removed" {
				return nil
			}
			return fmt.Errorf("Resource alias still exists: %s with resp code: %s", rs.Primary.ID, response)
		} else if !strings.Contains(err.Error(), "404") {
			return fmt.Errorf("Error waiting for resource key (%s) to be destroyed: %s with resp code: %s", rs.Primary.ID, err, response)
		}
	}

	return nil
}
