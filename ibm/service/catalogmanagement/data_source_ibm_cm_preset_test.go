// Copyright IBM Corp. 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package catalogmanagement_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIBMCmPresetDataSourceSimpleArgs(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCmPresetDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_cm_preset.cm_preset", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_preset.cm_preset", "preset"),
				),
			},
		},
	})
}

func testAccCheckIBMCmPresetDataSourceConfig() string {
	return fmt.Sprintf(`
		data "ibm_cm_preset" "cm_preset" {
			id = "28bece25-9615-448d-b449-204552f47ff6-provider_test_obj@1.0.0"
		}
	`)
}
