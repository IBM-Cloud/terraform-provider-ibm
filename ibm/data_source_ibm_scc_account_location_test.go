// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIbmSccAccountLocationDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmSccAccountLocationDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_scc_account_location.scc_account_location", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_account_location.scc_account_location", "location_id"),
				),
			},
		},
	})
}

func testAccCheckIbmSccAccountLocationDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		data "ibm_scc_account_location" "scc_account_location" {
			location_id = "us"
		}
	`)
}
