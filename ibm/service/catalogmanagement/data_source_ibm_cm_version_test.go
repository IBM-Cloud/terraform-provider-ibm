// Copyright IBM Corp. 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package catalogmanagement_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIBMCmVersionDataSourceSimpleArgs(t *testing.T) {
	versionZipurl := "https://github.com/IBM-Cloud/terraform-sample/archive/refs/tags/v1.1.0.tar.gz"
	versionTargetVersion := "1.1.1"
	versionIncludeConfig := "true"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCmVersionDataSourceConfig(versionZipurl, versionTargetVersion, versionIncludeConfig),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_cm_version.cm_version", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_version.cm_version", "rev"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_version.cm_version", "crn"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_version.cm_version", "version"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_version.cm_version", "sha"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_version.cm_version", "created"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_version.cm_version", "updated"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_version.cm_version", "offering_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_version.cm_version", "catalog_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_version.cm_version", "kind_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_version.cm_version", "repo_url"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_version.cm_version", "tgz_url"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_version.cm_version", "configuration.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_version.cm_version", "licenses.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_version.cm_version", "state.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_version.cm_version", "version_locator"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_version.cm_version", "long_description"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_version.cm_version", "is_consumable"),
				),
			},
		},
	})
}

func testAccCheckIBMCmVersionDataSourceConfig(versionZipurl string, versionTargetVersion string, versionIncludeConfig string) string {
	return fmt.Sprintf(`
		resource "ibm_cm_catalog" "cm_catalog" {
			label = "data_source_version_test_catalog_label"
			kind = "offering"
		}

		resource "ibm_cm_offering" "cm_offering" {
			catalog_id = ibm_cm_catalog.cm_catalog.id
			label = "test_tf_offering_label_1"
			name = "test_tf_offering_name_1"
			offering_icon_url = "test.url.1"
			tags = ["dev_ops"]
		}

		resource "ibm_cm_version" "cm_version" {
			catalog_id = ibm_cm_catalog.cm_catalog.id
			offering_id = ibm_cm_offering.cm_offering.id
			zipurl = "%s"
			target_version = "%s"
			include_config = %s
			install {}
			configuration {
				default_value = "foo"
				description = "The name to pass to the template."
				key = "name"
				type = "string"
				hidden = false
				required = false
				value_constraints {
					type  = "regex"
					value = "*"
					description = "Invalid name input"
				}
			}
		}

		data "ibm_cm_version" "cm_version" {
			version_loc_id = ibm_cm_version.cm_version.version_locator
		}
	`, versionZipurl, versionTargetVersion, versionIncludeConfig)
}
