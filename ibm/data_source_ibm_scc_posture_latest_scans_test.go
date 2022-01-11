// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMSccPostureListLatestScansDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMSccPostureListLatestScansDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_scc_posture_latest_scans.list_latest_scans", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_posture_latest_scans.list_latest_scans", "first.#"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_posture_latest_scans.list_latest_scans", "last.#"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_posture_latest_scans.list_latest_scans", "latest_scans.#"),
				),
			},
		},
	})
}

func testAccCheckIBMSccPostureListLatestScansDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		data "ibm_scc_posture_latest_scans" "list_latest_scans" {
		}
	`)
}
