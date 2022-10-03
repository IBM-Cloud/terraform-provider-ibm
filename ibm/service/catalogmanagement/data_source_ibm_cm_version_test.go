// Copyright IBM Corp. 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package catalogmanagement_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIBMCmVersionDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCmVersionDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_cm_version.cm_version", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_version.cm_version", "version_loc_id"),
				),
			},
		},
	})
}

func TestAccIBMCmVersionDataSourceAllArgs(t *testing.T) {
	versionName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	versionLabel := fmt.Sprintf("tf_label_%d", acctest.RandIntRange(10, 100))
	versionInstallKind := fmt.Sprintf("tf_install_kind_%d", acctest.RandIntRange(10, 100))
	versionFormatKind := fmt.Sprintf("tf_format_kind_%d", acctest.RandIntRange(10, 100))
	versionProductKind := fmt.Sprintf("tf_product_kind_%d", acctest.RandIntRange(10, 100))
	versionSha := fmt.Sprintf("tf_sha_%d", acctest.RandIntRange(10, 100))
	versionVersion := fmt.Sprintf("tf_version_%d", acctest.RandIntRange(10, 100))
	versionWorkingDirectory := fmt.Sprintf("tf_working_directory_%d", acctest.RandIntRange(10, 100))
	versionZipurl := fmt.Sprintf("tf_zipurl_%d", acctest.RandIntRange(10, 100))
	versionTargetVersion := fmt.Sprintf("tf_target_version_%d", acctest.RandIntRange(10, 100))
	versionIncludeConfig := "true"
	versionIsVsi := "true"
	versionRepotype := fmt.Sprintf("tf_repotype_%d", acctest.RandIntRange(10, 100))
	versionXAuthToken := fmt.Sprintf("tf_x_auth_token_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCmVersionDataSourceConfig(versionName, versionLabel, versionInstallKind, versionFormatKind, versionProductKind, versionSha, versionVersion, versionWorkingDirectory, versionZipurl, versionTargetVersion, versionIncludeConfig, versionIsVsi, versionRepotype, versionXAuthToken),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_cm_version.cm_version", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_version.cm_version", "version_loc_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_version.cm_version", "version_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_version.cm_version", "rev"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_version.cm_version", "crn"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_version.cm_version", "version"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_version.cm_version", "flavor.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_version.cm_version", "sha"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_version.cm_version", "created"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_version.cm_version", "updated"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_version.cm_version", "offering_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_version.cm_version", "catalog_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_version.cm_version", "kind_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_version.cm_version", "tags.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_version.cm_version", "repo_url"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_version.cm_version", "source_url"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_version.cm_version", "tgz_url"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_version.cm_version", "configuration.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_version.cm_version", "configuration.0.key"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_version.cm_version", "configuration.0.type"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_version.cm_version", "configuration.0.default_value"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_version.cm_version", "configuration.0.display_name"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_version.cm_version", "configuration.0.value_constraint"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_version.cm_version", "configuration.0.description"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_version.cm_version", "configuration.0.required"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_version.cm_version", "configuration.0.hidden"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_version.cm_version", "configuration.0.type_metadata"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_version.cm_version", "outputs.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_version.cm_version", "outputs.0.key"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_version.cm_version", "outputs.0.description"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_version.cm_version", "iam_permissions.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_version.cm_version", "iam_permissions.0.service_name"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_version.cm_version", "metadata.%"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_version.cm_version", "validation.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_version.cm_version", "required_resources.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_version.cm_version", "required_resources.0.type"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_version.cm_version", "required_resources.0.value"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_version.cm_version", "single_instance"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_version.cm_version", "install.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_version.cm_version", "pre_install.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_version.cm_version", "pre_install.0.instructions"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_version.cm_version", "pre_install.0.script"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_version.cm_version", "pre_install.0.script_permission"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_version.cm_version", "pre_install.0.delete_script"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_version.cm_version", "pre_install.0.scope"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_version.cm_version", "entitlement.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_version.cm_version", "licenses.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_version.cm_version", "licenses.0.id"),
					resource.TestCheckResourceAttr("data.ibm_cm_version.cm_version", "licenses.0.name", versionName),
					resource.TestCheckResourceAttrSet("data.ibm_cm_version.cm_version", "licenses.0.type"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_version.cm_version", "licenses.0.url"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_version.cm_version", "licenses.0.description"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_version.cm_version", "image_manifest_url"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_version.cm_version", "deprecated"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_version.cm_version", "package_version"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_version.cm_version", "state.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_version.cm_version", "version_locator"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_version.cm_version", "long_description"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_version.cm_version", "long_description_i18n.%"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_version.cm_version", "whitelisted_accounts.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_version.cm_version", "image_pull_key_name"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_version.cm_version", "deprecate_pending.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_version.cm_version", "solution_info.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_version.cm_version", "is_consumable"),
				),
			},
		},
	})
}

func testAccCheckIBMCmVersionDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		resource "ibm_cm_catalog" "cm_catalog" {
		}

		resource "ibm_cm_version" "cm_version" {
			catalog_identifier = ibm_cm_catalog.cm_catalog.id
			offering_id = ibm_cm_offering.cm_offering.offering_id
		}

		data "ibm_cm_version" "cm_version" {
			version_loc_id = "version_loc_id"
		}
	`)
}

func testAccCheckIBMCmVersionDataSourceConfig(versionName string, versionLabel string, versionInstallKind string, versionFormatKind string, versionProductKind string, versionSha string, versionVersion string, versionWorkingDirectory string, versionZipurl string, versionTargetVersion string, versionIncludeConfig string, versionIsVsi string, versionRepotype string, versionXAuthToken string) string {
	return fmt.Sprintf(`
		resource "ibm_cm_catalog" "cm_catalog" {
		}

		resource "ibm_cm_version" "cm_version" {
			catalog_identifier = ibm_cm_catalog.cm_catalog.id
			offering_id = ibm_cm_offering.cm_offering.offering_id
			tags = "FIXME"
			content = "FIXME"
			name = "%s"
			label = "%s"
			install_kind = "%s"
			target_kinds = "FIXME"
			format_kind = "%s"
			product_kind = "%s"
			sha = "%s"
			version = "%s"
			flavor {
				name = "name"
				label = "label"
				label_i18n = { "key": "inner" }
				index = 1
			}
			metadata {
				operating_system {
					dedicated_host_only = true
					vendor = "vendor"
					name = "name"
					href = "href"
					display_name = "display_name"
					family = "family"
					version = "version"
					architecture = "architecture"
				}
				file {
					size = 1
				}
				minimum_provisioned_size = 1
				images {
					id = "id"
					name = "name"
					region = "region"
				}
			}
			working_directory = "%s"
			zipurl = "%s"
			target_version = "%s"
			include_config = %s
			is_vsi = %s
			repotype = "%s"
			x_auth_token = "%s"
		}

		data "ibm_cm_version" "cm_version" {
			version_loc_id = "version_loc_id"
		}
	`, versionName, versionLabel, versionInstallKind, versionFormatKind, versionProductKind, versionSha, versionVersion, versionWorkingDirectory, versionZipurl, versionTargetVersion, versionIncludeConfig, versionIsVsi, versionRepotype, versionXAuthToken)
}
