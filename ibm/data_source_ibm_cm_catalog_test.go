// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMCmCatalogDataSource(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCmCatalogDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_cm_catalog.cm_catalog_data", "label"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_catalog.cm_catalog_data", "crn"),
				),
			},
		},
	})
}

func testAccCheckIBMCmCatalogDataSourceConfig() string {
	return fmt.Sprintf(`

		resource "ibm_cm_catalog" "cm_catalog" {
			label = "tf_test_datasource_catalog"
			short_description = "testing terraform provider with catalog"
		}
		
		data "ibm_cm_catalog" "cm_catalog_data" {
			catalog_identifier = ibm_cm_catalog.cm_catalog.id
		}
		`)
}
