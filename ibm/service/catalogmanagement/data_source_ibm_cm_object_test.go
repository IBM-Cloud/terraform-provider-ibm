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

func TestAccIBMCmObjectDataSourceSimpleArgs(t *testing.T) {
	catalogObjectName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	catalogObjectParentID := fmt.Sprintf("tf_parent_id_%d", acctest.RandIntRange(10, 100))
	catalogObjectLabel := fmt.Sprintf("tf_label_%d", acctest.RandIntRange(10, 100))
	catalogObjectShortDescription := fmt.Sprintf("tf_short_description_%d", acctest.RandIntRange(10, 100))
	catalogObjectKind := "preset_configuration"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCmObjectDataSourceConfig(catalogObjectName, catalogObjectParentID, catalogObjectLabel, catalogObjectShortDescription, catalogObjectKind),
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

func testAccCheckIBMCmObjectDataSourceConfig(catalogObjectName string, catalogObjectParentID string, catalogObjectLabel string, catalogObjectShortDescription string, catalogObjectKind string) string {
	return fmt.Sprintf(`
		resource "ibm_cm_catalog" "cm_catalog" {
			label = "tf_provider_obj_test"
			kind = "preset_configuration"
		}

		resource "ibm_cm_object" "cm_object" {
			catalog_id = ibm_cm_catalog.cm_catalog.id
			name = "%s"
			parent_id = "%s"
			label = "%s"
			short_description = "%s"
			kind = "%s"
		}

		data "ibm_cm_object" "cm_object" {
			catalog_id = ibm_cm_object.cm_object.catalog_id
			object_id = ibm_cm_object.cm_object.id
		}
	`, catalogObjectName, catalogObjectParentID, catalogObjectLabel, catalogObjectShortDescription, catalogObjectKind)
}
