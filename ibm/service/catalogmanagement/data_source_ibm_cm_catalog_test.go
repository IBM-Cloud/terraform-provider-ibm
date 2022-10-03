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

func TestAccIBMCmCatalogDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCmCatalogDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_cm_catalog.cm_catalog", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_catalog.cm_catalog", "catalog_identifier"),
				),
			},
		},
	})
}

func TestAccIBMCmCatalogDataSourceAllArgs(t *testing.T) {
	catalogLabel := fmt.Sprintf("tf_label_%d", acctest.RandIntRange(10, 100))
	catalogShortDescription := fmt.Sprintf("tf_short_description_%d", acctest.RandIntRange(10, 100))
	catalogCatalogIconURL := fmt.Sprintf("tf_catalog_icon_url_%d", acctest.RandIntRange(10, 100))
	catalogDisabled := "true"
	catalogResourceGroupID := fmt.Sprintf("tf_resource_group_id_%d", acctest.RandIntRange(10, 100))
	catalogOwningAccount := fmt.Sprintf("tf_owning_account_%d", acctest.RandIntRange(10, 100))
	catalogKind := fmt.Sprintf("tf_kind_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCmCatalogDataSourceConfig(catalogLabel, catalogShortDescription, catalogCatalogIconURL, catalogDisabled, catalogResourceGroupID, catalogOwningAccount, catalogKind),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_cm_catalog.cm_catalog", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_catalog.cm_catalog", "catalog_identifier"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_catalog.cm_catalog", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_catalog.cm_catalog", "rev"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_catalog.cm_catalog", "label"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_catalog.cm_catalog", "label_i18n.%"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_catalog.cm_catalog", "short_description"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_catalog.cm_catalog", "short_description_i18n.%"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_catalog.cm_catalog", "catalog_icon_url"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_catalog.cm_catalog", "tags.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_catalog.cm_catalog", "url"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_catalog.cm_catalog", "crn"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_catalog.cm_catalog", "offerings_url"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_catalog.cm_catalog", "features.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_catalog.cm_catalog", "features.0.title"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_catalog.cm_catalog", "features.0.description"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_catalog.cm_catalog", "disabled"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_catalog.cm_catalog", "created"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_catalog.cm_catalog", "updated"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_catalog.cm_catalog", "resource_group_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_catalog.cm_catalog", "owning_account"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_catalog.cm_catalog", "catalog_filters.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_catalog.cm_catalog", "syndication_settings.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_catalog.cm_catalog", "kind"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_catalog.cm_catalog", "metadata.%"),
				),
			},
		},
	})
}

func testAccCheckIBMCmCatalogDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		resource "ibm_cm_catalog" "cm_catalog" {
		}

		data "ibm_cm_catalog" "cm_catalog" {
			catalog_identifier = ibm_cm_catalog.cm_catalog.catalog_id
		}
	`)
}

func testAccCheckIBMCmCatalogDataSourceConfig(catalogLabel string, catalogShortDescription string, catalogCatalogIconURL string, catalogDisabled string, catalogResourceGroupID string, catalogOwningAccount string, catalogKind string) string {
	return fmt.Sprintf(`
		resource "ibm_cm_catalog" "cm_catalog" {
			label = "%s"
			label_i18n = "FIXME"
			short_description = "%s"
			short_description_i18n = "FIXME"
			catalog_icon_url = "%s"
			tags = "FIXME"
			features {
				title = "title"
				title_i18n = { "key": "inner" }
				description = "description"
				description_i18n = { "key": "inner" }
			}
			disabled = %s
			resource_group_id = "%s"
			owning_account = "%s"
			catalog_filters {
				include_all = true
				category_filters = { "key": { example: "object" } }
				id_filters {
					include {
						filter_terms = [ "filter_terms" ]
					}
					exclude {
						filter_terms = [ "filter_terms" ]
					}
				}
			}
			syndication_settings {
				remove_related_components = true
				clusters {
					region = "region"
					id = "id"
					name = "name"
					resource_group_name = "resource_group_name"
					type = "type"
					namespaces = [ "namespaces" ]
					all_namespaces = true
				}
				history {
					namespaces = [ "namespaces" ]
					clusters {
						region = "region"
						id = "id"
						name = "name"
						resource_group_name = "resource_group_name"
						type = "type"
						namespaces = [ "namespaces" ]
						all_namespaces = true
					}
					last_run = "2021-01-31T09:44:12Z"
				}
				authorization {
					token = "token"
					last_run = "2021-01-31T09:44:12Z"
				}
			}
			kind = "%s"
			metadata = "FIXME"
		}

		data "ibm_cm_catalog" "cm_catalog" {
			catalog_identifier = ibm_cm_catalog.cm_catalog.catalog_id
		}
	`, catalogLabel, catalogShortDescription, catalogCatalogIconURL, catalogDisabled, catalogResourceGroupID, catalogOwningAccount, catalogKind)
}
