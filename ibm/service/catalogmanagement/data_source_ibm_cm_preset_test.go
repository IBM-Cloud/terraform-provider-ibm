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

func TestAccIBMCmPresetDataSourceSimpleArgs(t *testing.T) {
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
				Config: testAccCheckIBMCmPresetDataSourceConfig(catalogObjectName, catalogObjectParentID, catalogObjectLabel, catalogObjectShortDescription, catalogObjectKind),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_cm_preset.cm_preset", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_preset.cm_preset", "preset"),
				),
			},
		},
	})
}

func testAccCheckIBMCmPresetDataSourceConfig(catalogObjectName string, catalogObjectParentID string, catalogObjectLabel string, catalogObjectShortDescription string, catalogObjectKind string) string {
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

		data "ibm_cm_preset" "cm_preset" {
			id = "${ibm_cm_object.cm_object.id}:1.0.0"
		}
	`, catalogObjectName, catalogObjectParentID, catalogObjectLabel, catalogObjectShortDescription, catalogObjectKind)
}
