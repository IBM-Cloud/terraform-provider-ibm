// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIbmResourceBindingsDataSourceBasic(t *testing.T) {
	resourceBindingName := fmt.Sprintf("tf_binding_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheckResourceAliasBinding(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmResourceBindingsDataSourceConfigBasic(resourceBindingName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_resource_bindings.resource_bindings", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_resource_bindings.resource_bindings", "bindings.#"),
					resource.TestCheckResourceAttrSet("data.ibm_resource_bindings.resource_bindings", "bindings.0.id"),
					resource.TestCheckResourceAttrSet("data.ibm_resource_bindings.resource_bindings", "bindings.0.guid"),
					resource.TestCheckResourceAttrSet("data.ibm_resource_bindings.resource_bindings", "bindings.0.url"),
					resource.TestCheckResourceAttrSet("data.ibm_resource_bindings.resource_bindings", "bindings.0.created_at"),
					// resource.TestCheckResourceAttrSet("data.ibm_resource_bindings.resource_bindings", "bindings.0.updated_at"),
					// resource.TestCheckResourceAttrSet("data.ibm_resource_bindings.resource_bindings", "bindings.0.deleted_at"),
					resource.TestCheckResourceAttrSet("data.ibm_resource_bindings.resource_bindings", "bindings.0.created_by"),
					// resource.TestCheckResourceAttrSet("data.ibm_resource_bindings.resource_bindings", "bindings.0.updated_by"),
					// resource.TestCheckResourceAttrSet("data.ibm_resource_bindings.resource_bindings", "bindings.0.deleted_by"),
					resource.TestCheckResourceAttrSet("data.ibm_resource_bindings.resource_bindings", "bindings.0.source_crn"),
					resource.TestCheckResourceAttrSet("data.ibm_resource_bindings.resource_bindings", "bindings.0.target_crn"),
					resource.TestCheckResourceAttrSet("data.ibm_resource_bindings.resource_bindings", "bindings.0.crn"),
					resource.TestCheckResourceAttrSet("data.ibm_resource_bindings.resource_bindings", "bindings.0.region_binding_id"),
					resource.TestCheckResourceAttrSet("data.ibm_resource_bindings.resource_bindings", "bindings.0.region_binding_crn"),
					resource.TestCheckResourceAttr("data.ibm_resource_bindings.resource_bindings", "bindings.0.name", resourceBindingName),
					resource.TestCheckResourceAttrSet("data.ibm_resource_bindings.resource_bindings", "bindings.0.account_id"),
					resource.TestCheckResourceAttrSet("data.ibm_resource_bindings.resource_bindings", "bindings.0.resource_group_id"),
					resource.TestCheckResourceAttrSet("data.ibm_resource_bindings.resource_bindings", "bindings.0.state"),
					resource.TestCheckResourceAttrSet("data.ibm_resource_bindings.resource_bindings", "bindings.0.iam_compatible"),
					resource.TestCheckResourceAttrSet("data.ibm_resource_bindings.resource_bindings", "bindings.0.resource_id"),
					resource.TestCheckResourceAttrSet("data.ibm_resource_bindings.resource_bindings", "bindings.0.migrated"),
					resource.TestCheckResourceAttrSet("data.ibm_resource_bindings.resource_bindings", "bindings.0.resource_alias_url"),
				),
			},
		},
	})
}

func testAccCheckIbmResourceBindingsDataSourceConfigBasic(resourceBindingName string) string {
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
		}
		resource "ibm_resource_binding" "binding" {
			source = ibm_resource_alias.alias.guid
			target = "crn:v1:bluemix:public:cf:us-south:s/${ibm_app.app.space_guid}::cf-application:${ibm_app.app.id}"
			name = "%s"
			role = "Writer"
		}
		data "ibm_resource_bindings" "resource_bindings" {
			depends_on = [ibm_resource_binding.binding]
			name = "%[3]s"
		}
	`, cfOrganization, cfSpace, resourceBindingName)
}
