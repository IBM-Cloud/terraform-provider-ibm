// Copyright IBM Corp. 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package catalogmanagement_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIBMCmObjectDataSourceSimpleArgs(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCmObjectDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_cm_object.cm_object", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_object.cm_object", "catalog_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_object.cm_object", "object_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_object.cm_object", "catalog_object_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_object.cm_object", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_object.cm_object", "rev"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_object.cm_object", "crn"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_object.cm_object", "url"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_object.cm_object", "parent_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_object.cm_object", "label"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_object.cm_object", "created"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_object.cm_object", "updated"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_object.cm_object", "short_description"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_object.cm_object", "kind"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_object.cm_object", "catalog_name"),
				),
			},
		},
	})
}

func testAccCheckIBMCmObjectDataSourceConfig() string {
	return fmt.Sprintf(`
		data "ibm_cm_object" "cm_object" {
			catalog_id = "28bece25-9615-448d-b449-204552f47ff6"
			object_id = "28bece25-9615-448d-b449-204552f47ff6-provider_test_obj"
		}
	`)
}
