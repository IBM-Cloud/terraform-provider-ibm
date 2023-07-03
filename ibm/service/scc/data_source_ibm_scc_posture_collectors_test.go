// Copyright IBM Corp. 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package scc_test

import (
	"fmt"
	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMSccPostureListCollectorsDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMSccPostureListCollectorsDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_scc_posture_collectors.list_collectors", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_posture_collectors.list_collectors", "offset"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_posture_collectors.list_collectors", "limit"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_posture_collectors.list_collectors", "total_count"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_posture_collectors.list_collectors", "first.#"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_posture_collectors.list_collectors", "last.#"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_posture_collectors.list_collectors", "collectors.#"),
				),
			},
		},
	})
}

func testAccCheckIBMSccPostureListCollectorsDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		data "ibm_scc_posture_collectors" "list_collectors" {
		}
	`)
}
