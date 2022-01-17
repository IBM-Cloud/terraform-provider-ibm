// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package scc_test

import (
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIbmSccAccountLocationSettingsDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIbmSccAccountLocationSettingsDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_scc_account_settings.scc_account_location_settings", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_account_settings.scc_account_location_settings", "id"),
				),
			},
		},
	})
}

func testAccCheckIbmSccAccountLocationSettingsDataSourceConfigBasic() string {
	return `
		data "ibm_scc_account_settings" "scc_account_location_settings" {
		}
	`
}
