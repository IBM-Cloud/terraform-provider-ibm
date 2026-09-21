// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.116.0-df613dbc-20260803-154903
 */

package drautomationservice_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIBMPdrGrsLocationPairsDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMPdrGrsLocationPairsDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_pdr_grs_location_pairs.pdr_grs_location_pairs_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_pdr_grs_location_pairs.pdr_grs_location_pairs_instance", "instance_id"),
					resource.TestCheckResourceAttrSet("data.ibm_pdr_grs_location_pairs.pdr_grs_location_pairs_instance", "location_pairs.%"),
				),
			},
		},
	})
}

func testAccCheckIBMPdrGrsLocationPairsDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		data "ibm_pdr_grs_location_pairs" "pdr_grs_location_pairs_instance" {
			instance_id = "123456d3-1122-3344-b67d-4389b44b7bf9"
			Accept-Language = "en-US"
		}
	`)
}
