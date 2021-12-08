// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMSccPostureV2ProfileDetailsDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMSccPostureV2ProfileDetailsDataSourceConfigBasic(scc_posture_v2_profile_id),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_scc_posture_v2_profileDetails.profile_details", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_posture_v2_profileDetails.profile_details", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_posture_v2_profileDetails.profile_details", "profile_type"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_posture_v2_profileDetails.profile_details", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_posture_v2_profileDetails.profile_details", "description"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_posture_v2_profileDetails.profile_details", "version"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_posture_v2_profileDetails.profile_details", "created_by"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_posture_v2_profileDetails.profile_details", "modified_by"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_posture_v2_profileDetails.profile_details", "base_profile"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_posture_v2_profileDetails.profile_details", "type"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_posture_v2_profileDetails.profile_details", "no_of_controls"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_posture_v2_profileDetails.profile_details", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_posture_v2_profileDetails.profile_details", "updated_at"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_posture_v2_profileDetails.profile_details", "enabled"),
				),
			},
		},
	})
}

func testAccCheckIBMSccPostureV2ProfileDetailsDataSourceConfigBasic(profileId string) string {
	return fmt.Sprintf(`
		data "ibm_scc_posture_v2_profileDetails" "profile_details" {
			id = "%s"
			profile_type = "4"
		}
	`,profileId)
}

