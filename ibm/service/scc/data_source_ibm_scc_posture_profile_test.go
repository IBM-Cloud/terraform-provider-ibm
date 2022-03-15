// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package scc_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMSccPostureProfileDetailsDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMSccPostureProfileDetailsDataSourceConfigBasic(acc.Scc_posture_profile_id),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_scc_posture_profile.profile_details", "profile_type"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_posture_profile.profile_details", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_posture_profile.profile_details", "description"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_posture_profile.profile_details", "version"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_posture_profile.profile_details", "created_by"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_posture_profile.profile_details", "modified_by"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_posture_profile.profile_details", "base_profile"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_posture_profile.profile_details", "type"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_posture_profile.profile_details", "no_of_controls"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_posture_profile.profile_details", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_posture_profile.profile_details", "updated_at"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_posture_profile.profile_details", "enabled"),
				),
			},
		},
	})
}

func testAccCheckIBMSccPostureProfileDetailsDataSourceConfigBasic(profileId string) string {
	return fmt.Sprintf(`
		data "ibm_scc_posture_profile" "profile_details" {
			profile_id = "%s"
			profile_type = "authored"
		}
	`, profileId)
}
