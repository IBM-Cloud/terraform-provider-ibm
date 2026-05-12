// Copyright IBM Corp. 2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.108.0-56772134-20251111-102802
 */

package drautomationservice_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIBMPdrGetGrsLocationPairsDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMPdrGetGrsLocationPairsDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_pdr_get_grs_location_pairs.pdr_get_grs_location_pairs_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_pdr_get_grs_location_pairs.pdr_get_grs_location_pairs_instance", "instance_id"),
				),
			},
		},
	})
}

func testAccCheckIBMPdrGetGrsLocationPairsDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		data "ibm_pdr_get_grs_location_pairs" "pdr_get_grs_location_pairs_instance" {
			instance_id = "ac645fe5-fba1-4cb3-952e-e1b09fa0df26"
		}
	`)
}
