// Copyright IBM Corp. 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package power_test

import (
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccIBMPIDisasterRecoveryLocationsDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPIDisasterRecoveryLocationsDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_pi_disaster_recovery_locations.testacc_disaster_recovery_locations", "id"),
				),
			},
		},
	})
}

func testAccCheckIBMPIDisasterRecoveryLocationsDataSourceConfig() string {
	return `data "ibm_pi_disaster_recovery_locations" "testacc_disaster_recovery_locations" {}`
}
