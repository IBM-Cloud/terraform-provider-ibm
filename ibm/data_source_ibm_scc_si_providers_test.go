// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMSccSiProvidersDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMSccSiProvidersDataSourceConfigBasic(scc_si_account),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_scc_si_providers.scc_si_providers", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_si_providers.scc_si_providers", "providers.#"),
				),
			},
		},
	})
}

func testAccCheckIBMSccSiProvidersDataSourceConfigBasic(accountID string) string {
	return fmt.Sprintf(`
		data "ibm_scc_si_providers" "scc_si_providers" {
			account_id = "%s"
		}
	`, accountID)
}
