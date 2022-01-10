// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMSccPostureScopeCorrelationDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMSccPostureScopeCorrelationDataSourceConfigBasic(scc_posture_correlation_id),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_scc_posture_scope_correlation.scope_correlation", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_posture_scope_correlation.scope_correlation", "correlation_id"),
				),
			},
		},
	})
}

func testAccCheckIBMSccPostureScopeCorrelationDataSourceConfigBasic(correlationId string) string {
	return fmt.Sprintf(`
		data "ibm_scc_posture_scope_correlation" "scope_correlation" {
			correlation_id = "%s"
		}
	`, correlationId)
}
