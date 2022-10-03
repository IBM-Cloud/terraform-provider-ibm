// Copyright IBM Corp. 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package catalogmanagement_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM/platform-services-go-sdk/catalogmanagementv1"
)

func TestAccIBMCmCatalogBasic(t *testing.T) {
	var conf catalogmanagementv1.Catalog

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCmCatalogDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCmCatalogConfigBasic(),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCmCatalogExists("ibm_cm_catalog.cm_catalog", conf),
				),
			},
		},
	})
}

func TestAccIBMCmCatalogAllArgs(t *testing.T) {
	var conf catalogmanagementv1.Catalog
	label := fmt.Sprintf("tf_label_%d", acctest.RandIntRange(10, 100))
	shortDescription := fmt.Sprintf("tf_short_description_%d", acctest.RandIntRange(10, 100))
	catalogIconURL := fmt.Sprintf("tf_catalog_icon_url_%d", acctest.RandIntRange(10, 100))
	disabled := "true"
	resourceGroupID := fmt.Sprintf("tf_resource_group_id_%d", acctest.RandIntRange(10, 100))
	owningAccount := fmt.Sprintf("tf_owning_account_%d", acctest.RandIntRange(10, 100))
	kind := fmt.Sprintf("tf_kind_%d", acctest.RandIntRange(10, 100))
	labelUpdate := fmt.Sprintf("tf_label_%d", acctest.RandIntRange(10, 100))
	shortDescriptionUpdate := fmt.Sprintf("tf_short_description_%d", acctest.RandIntRange(10, 100))
	catalogIconURLUpdate := fmt.Sprintf("tf_catalog_icon_url_%d", acctest.RandIntRange(10, 100))
	disabledUpdate := "false"
	resourceGroupIDUpdate := fmt.Sprintf("tf_resource_group_id_%d", acctest.RandIntRange(10, 100))
	owningAccountUpdate := fmt.Sprintf("tf_owning_account_%d", acctest.RandIntRange(10, 100))
	kindUpdate := fmt.Sprintf("tf_kind_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCmCatalogDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCmCatalogConfig(label, shortDescription, catalogIconURL, disabled, resourceGroupID, owningAccount, kind),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCmCatalogExists("ibm_cm_catalog.cm_catalog", conf),
					resource.TestCheckResourceAttr("ibm_cm_catalog.cm_catalog", "label", label),
					resource.TestCheckResourceAttr("ibm_cm_catalog.cm_catalog", "short_description", shortDescription),
					resource.TestCheckResourceAttr("ibm_cm_catalog.cm_catalog", "catalog_icon_url", catalogIconURL),
					resource.TestCheckResourceAttr("ibm_cm_catalog.cm_catalog", "disabled", disabled),
					resource.TestCheckResourceAttr("ibm_cm_catalog.cm_catalog", "resource_group_id", resourceGroupID),
					resource.TestCheckResourceAttr("ibm_cm_catalog.cm_catalog", "owning_account", owningAccount),
					resource.TestCheckResourceAttr("ibm_cm_catalog.cm_catalog", "kind", kind),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMCmCatalogConfig(labelUpdate, shortDescriptionUpdate, catalogIconURLUpdate, disabledUpdate, resourceGroupIDUpdate, owningAccountUpdate, kindUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_cm_catalog.cm_catalog", "label", labelUpdate),
					resource.TestCheckResourceAttr("ibm_cm_catalog.cm_catalog", "short_description", shortDescriptionUpdate),
					resource.TestCheckResourceAttr("ibm_cm_catalog.cm_catalog", "catalog_icon_url", catalogIconURLUpdate),
					resource.TestCheckResourceAttr("ibm_cm_catalog.cm_catalog", "disabled", disabledUpdate),
					resource.TestCheckResourceAttr("ibm_cm_catalog.cm_catalog", "resource_group_id", resourceGroupIDUpdate),
					resource.TestCheckResourceAttr("ibm_cm_catalog.cm_catalog", "owning_account", owningAccountUpdate),
					resource.TestCheckResourceAttr("ibm_cm_catalog.cm_catalog", "kind", kindUpdate),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_cm_catalog.cm_catalog",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIBMCmCatalogConfigBasic() string {
	return fmt.Sprintf(`

		resource "ibm_cm_catalog" "cm_catalog" {
		}
	`)
}

func testAccCheckIBMCmCatalogConfig(label string, shortDescription string, catalogIconURL string, disabled string, resourceGroupID string, owningAccount string, kind string) string {
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
	`, label, shortDescription, catalogIconURL, disabled, resourceGroupID, owningAccount, kind)
}

func testAccCheckIBMCmCatalogExists(n string, obj catalogmanagementv1.Catalog) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		catalogManagementClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).CatalogManagementV1()
		if err != nil {
			return err
		}

		getCatalogOptions := &catalogmanagementv1.GetCatalogOptions{}

		getCatalogOptions.SetCatalogIdentifier(rs.Primary.ID)

		catalog, _, err := catalogManagementClient.GetCatalog(getCatalogOptions)
		if err != nil {
			return err
		}

		obj = *catalog
		return nil
	}
}

func testAccCheckIBMCmCatalogDestroy(s *terraform.State) error {
	catalogManagementClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).CatalogManagementV1()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_cm_catalog" {
			continue
		}

		getCatalogOptions := &catalogmanagementv1.GetCatalogOptions{}

		getCatalogOptions.SetCatalogIdentifier(rs.Primary.ID)

		// Try to find the key
		_, response, err := catalogManagementClient.GetCatalog(getCatalogOptions)

		if err == nil {
			return fmt.Errorf("cm_catalog still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for cm_catalog (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
