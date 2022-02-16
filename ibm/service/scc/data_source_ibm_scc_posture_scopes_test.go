// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package scc_test

import (
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMSccPostureListScopesDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMSccPostureListScopesDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_scc_posture_scopes.list_scopes", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_posture_scopes.list_scopes", "offset"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_posture_scopes.list_scopes", "limit"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_posture_scopes.list_scopes", "total_count"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_posture_scopes.list_scopes", "first.#"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_posture_scopes.list_scopes", "last.#"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_posture_scopes.list_scopes", "scopes.#"),
				),
			},
		},
	})
}

func testAccCheckIBMSccPostureListScopesDataSourceConfigBasic() string {
	return `
		data "ibm_scc_posture_scopes" "list_scopes" {
		}
	`
}
