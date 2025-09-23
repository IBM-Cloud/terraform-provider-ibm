// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package power_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"
)

func TestAccIBMPIStorageTiersDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPIStorageTiersDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_pi_storage_tiers.storage_tiers", "id"),
				),
			},
		},
	})
}

func testAccCheckIBMPIStorageTiersDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		data "ibm_pi_storage_tiers" "storage_tiers" {
			pi_cloud_instance_id = "%s"
		}`, acc.Pi_cloud_instance_id)
}
