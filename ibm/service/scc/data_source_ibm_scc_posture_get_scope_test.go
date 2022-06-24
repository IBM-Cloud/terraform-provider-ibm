// Copyright IBM Corp. 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package scc

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMSccPostureGetScopeDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMSccPostureGetScopeDataSourceConfigBasic(scc_posture_scope_id),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_scc_posture_get_scope.get_scope", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_posture_get_scope.get_scope", "name"),
				),
			},
		},
	})
}

func testAccCheckIBMSccPostureGetScopeDataSourceConfigBasic(scopeId string) string {
	return fmt.Sprintf(`
		data "ibm_scc_posture_get_scope" "get_scope" {
			id = "%s"
		}
	`, scopeId)
}
