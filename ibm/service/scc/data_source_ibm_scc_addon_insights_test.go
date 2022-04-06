// Copyright IBM Corp. 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package scc_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIBMSccAddonInsightsDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMSccAddonInsightsDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_scc_addon_insights.scc_addon_insights", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_addon_insights.scc_addon_insights", "type.#"),
				),
			},
		},
	})
}

func testAccCheckIBMSccAddonInsightsDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		data "ibm_scc_addon_insights" "scc_addon_insights" {
		}
	`)
}
