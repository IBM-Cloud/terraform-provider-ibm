// Copyright IBM Corp. 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package catalogmanagement_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

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

func TestAccIBMCmCatalogDataSourceSimpleArgs(t *testing.T) {
	catalogLabel := fmt.Sprintf("tf_label_%d", acctest.RandIntRange(10, 100))
	catalogShortDescription := fmt.Sprintf("tf_short_description_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCmCatalogDataSourceConfig(catalogLabel, catalogShortDescription),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_cm_catalog.cm_catalog", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_catalog.cm_catalog", "label"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_catalog.cm_catalog", "short_description"),
				),
			},
		},
	})
}

func testAccCheckIBMCmCatalogDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		resource "ibm_cm_catalog" "cm_catalog" {
			label = "basic-catalog-label-test"
			kind = "offering"
		}

		data "ibm_cm_catalog" "cm_catalog" {
			catalog_identifier = ibm_cm_catalog.cm_catalog.id
		}
	`)
}

func testAccCheckIBMCmCatalogDataSourceConfig(catalogLabel string, catalogShortDescription string) string {
	return fmt.Sprintf(`
		resource "ibm_cm_catalog" "cm_catalog" {
			label = "%s"
			short_description = "%s"
			kind = "offering"
		}

		data "ibm_cm_catalog" "cm_catalog" {
			label = ibm_cm_catalog.cm_catalog.label
		}
	`, catalogLabel, catalogShortDescription)
}
