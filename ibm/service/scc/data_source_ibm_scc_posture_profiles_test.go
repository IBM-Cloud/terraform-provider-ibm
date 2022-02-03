// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package scc_test

import (
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMSccPostureListProfilesDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMSccPostureListProfilesDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_scc_posture_profiles.list_profiles", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_posture_profiles.list_profiles", "first.#"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_posture_profiles.list_profiles", "last.#"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_posture_profiles.list_profiles", "profiles.#"),
				),
			},
		},
	})
}

func testAccCheckIBMSccPostureListProfilesDataSourceConfigBasic() string {
	return `
		data "ibm_scc_posture_profiles" "list_profiles" {
		}
	`
}
