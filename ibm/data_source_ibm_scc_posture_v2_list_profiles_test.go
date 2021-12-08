// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMSccPostureV2ListProfilesDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMSccPostureV2ListProfilesDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_scc_posture_v2_list_profiles.list_profiles", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_posture_v2_list_profiles.list_profiles", "first.#"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_posture_v2_list_profiles.list_profiles", "last.#"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_posture_v2_list_profiles.list_profiles", "profiles.#"),
				),
			},
		},
	})
}

func testAccCheckIBMSccPostureV2ListProfilesDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		data "ibm_scc_posture_v2_list_profiles" "list_profiles" {
		}
	`)
}

