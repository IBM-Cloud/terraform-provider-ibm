// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package power_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccIBMPISAPProfilesDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPISAPProfilesDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_pi_sap_profiles.test", "id"),
				),
			},
		},
	})
}

func testAccCheckIBMPISAPProfilesDataSourceConfig() string {
	return fmt.Sprintf(`
		data "ibm_pi_sap_profiles" "test" {
			pi_cloud_instance_id = "%s"
		}`, acc.Pi_cloud_instance_id)
}

func TestAccIBMPISAPProfilesDataSourceFilters(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPISAPProfilesDataSourceFiltersConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_pi_sap_profiles.test", "id"),
				),
			},
		},
	})
}

func testAccCheckIBMPISAPProfilesDataSourceFiltersConfig() string {
	return fmt.Sprintf(`
		data "ibm_pi_sap_profiles" "test" {
			pi_cloud_instance_id = "%s"
			pi_family_filter     = "balanced"
			pi_prefix_filter     = "bh1"
		}`, acc.Pi_cloud_instance_id)
}
