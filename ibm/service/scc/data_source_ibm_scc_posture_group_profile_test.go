// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package scc_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMSccPostureGroupProfileDetailsDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMSccPostureGroupProfileDetailsDataSourceConfigBasic(acc.Scc_posture_group_profile_id),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_scc_posture_group_profile.group_profile_details", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_posture_group_profile.group_profile_details", "profile_id"),
				),
			},
		},
	})
}

func testAccCheckIBMSccPostureGroupProfileDetailsDataSourceConfigBasic(groupprofileId string) string {
	return fmt.Sprintf(`
		data "ibm_scc_posture_group_profile" "group_profile_details" {
			profile_id = "%s"
		}
	`, groupprofileId)
}
