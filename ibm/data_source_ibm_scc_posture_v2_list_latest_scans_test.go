// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMSccPostureV2ListLatestScansDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMSccPostureV2ListLatestScansDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_scc_posture_v2_list_latest_scans.list_latest_scans", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_posture_v2_list_latest_scans.list_latest_scans", "first.#"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_posture_v2_list_latest_scans.list_latest_scans", "last.#"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_posture_v2_list_latest_scans.list_latest_scans", "latest_scans.#"),
				),
			},
		},
	})
}

func testAccCheckIBMSccPostureV2ListLatestScansDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		data "ibm_scc_posture_v2_list_latest_scans" "list_latest_scans" {
		}
	`)
}

