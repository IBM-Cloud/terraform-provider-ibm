// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package catalogmanagement_test

import (
	"fmt"
	"os"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMCmCatalogDataSource(t *testing.T) {
	ResourceGroupID := os.Getenv("CATMGMT_RESOURCE_GROUP_ID")

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCmCatalogDataSourceConfig(ResourceGroupID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_cm_catalog.cm_catalog_data", "label"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_catalog.cm_catalog_data", "crn"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_catalog.cm_catalog_data", "kind"),
					resource.TestCheckResourceAttrSet("ibm_cm_catalog.cm_catalog", "resource_group_id"),
				),
			},
			{
				Config: testAccCheckIBMCmCatalogDataSourceConfigDefault(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_cm_catalog.cm_catalog_data", "label"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_catalog.cm_catalog_data", "crn"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_catalog.cm_catalog_data", "kind"),
					resource.TestCheckResourceAttrSet("ibm_cm_catalog.cm_catalog", "resource_group_id"),
				),
			},
		},
	})
}

func testAccCheckIBMCmCatalogDataSourceConfig(resourceGroupID string) string {
	return fmt.Sprintf(`

		resource "ibm_cm_catalog" "cm_catalog" {
			label = "tf_test_datasource_catalog"
			short_description = "testing terraform provider with catalog"
			resource_group_id = "%s"
		}
		
		data "ibm_cm_catalog" "cm_catalog_data" {
			catalog_identifier = ibm_cm_catalog.cm_catalog.id
		}
		`, resourceGroupID)
}

func testAccCheckIBMCmCatalogDataSourceConfigDefault() string {
	return `

		resource "ibm_cm_catalog" "cm_catalog" {
			label = "tf_test_datasource_catalog"
			short_description = "testing terraform provider with catalog"
		}
		
		data "ibm_cm_catalog" "cm_catalog_data" {
			catalog_identifier = ibm_cm_catalog.cm_catalog.id
		}
		`
}
