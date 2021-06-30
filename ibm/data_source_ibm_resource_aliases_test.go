// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIbmResourceAliasesDataSourceBasic(t *testing.T) {
	resourceAliasName := fmt.Sprintf("tf_alias_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheckResourceAliasBinding(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmResourceAliasesDataSourceConfigBasic(resourceAliasName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_resource_aliases.resource_aliases", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_resource_aliases.resource_aliases", "aliases.#"),
					resource.TestCheckResourceAttr("data.ibm_resource_aliases.resource_aliases", "aliases.0.name", resourceAliasName),
					resource.TestCheckResourceAttrSet("data.ibm_resource_aliases.resource_aliases", "aliases.0.id"),
					resource.TestCheckResourceAttrSet("data.ibm_resource_aliases.resource_aliases", "aliases.0.guid"),
					resource.TestCheckResourceAttrSet("data.ibm_resource_aliases.resource_aliases", "aliases.0.url"),
					resource.TestCheckResourceAttrSet("data.ibm_resource_aliases.resource_aliases", "aliases.0.created_at"),
					// resource.TestCheckResourceAttrSet("data.ibm_resource_aliases.resource_aliases", "aliases.0.updated_at"),
					// resource.TestCheckResourceAttrSet("data.ibm_resource_aliases.resource_aliases", "aliases.0.deleted_at"),
					resource.TestCheckResourceAttrSet("data.ibm_resource_aliases.resource_aliases", "aliases.0.created_by"),
					// resource.TestCheckResourceAttrSet("data.ibm_resource_aliases.resource_aliases", "aliases.0.updated_by"),
					// resource.TestCheckResourceAttrSet("data.ibm_resource_aliases.resource_aliases", "aliases.0.deleted_by"),
					resource.TestCheckResourceAttrSet("data.ibm_resource_aliases.resource_aliases", "aliases.0.resource_instance_id"),
					resource.TestCheckResourceAttrSet("data.ibm_resource_aliases.resource_aliases", "aliases.0.target_crn"),
					resource.TestCheckResourceAttrSet("data.ibm_resource_aliases.resource_aliases", "aliases.0.account_id"),
					resource.TestCheckResourceAttrSet("data.ibm_resource_aliases.resource_aliases", "aliases.0.resource_id"),
					resource.TestCheckResourceAttrSet("data.ibm_resource_aliases.resource_aliases", "aliases.0.resource_group_id"),
					resource.TestCheckResourceAttrSet("data.ibm_resource_aliases.resource_aliases", "aliases.0.crn"),
					resource.TestCheckResourceAttrSet("data.ibm_resource_aliases.resource_aliases", "aliases.0.region_instance_id"),
					resource.TestCheckResourceAttrSet("data.ibm_resource_aliases.resource_aliases", "aliases.0.region_instance_crn"),
					resource.TestCheckResourceAttrSet("data.ibm_resource_aliases.resource_aliases", "aliases.0.state"),
					resource.TestCheckResourceAttrSet("data.ibm_resource_aliases.resource_aliases", "aliases.0.migrated"),
					resource.TestCheckResourceAttrSet("data.ibm_resource_aliases.resource_aliases", "aliases.0.resource_instance_url"),
					resource.TestCheckResourceAttrSet("data.ibm_resource_aliases.resource_aliases", "aliases.0.resource_bindings_url"),
					resource.TestCheckResourceAttrSet("data.ibm_resource_aliases.resource_aliases", "aliases.0.resource_keys_url"),
				),
			},
		},
	})
}

func testAccCheckIbmResourceAliasesDataSourceConfigBasic(resourceAliasName string) string {
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
			name = "%s"
			source = ibm_resource_instance.instance.guid
			target = "crn:v1:bluemix:public:cf:us-south:o/${data.ibm_org.org.id}::cf-space:${data.ibm_space.space.id}"
		}
		data "ibm_resource_aliases" "resource_aliases" {
			depends_on = [ibm_resource_alias.alias]
			name = "%[3]s"
		}
	`, cfOrganization, cfSpace, resourceAliasName)
}
