/**
 * © Copyright IBM Corporation 2020. All Rights Reserved.
 *
 * Licensed under the Mozilla Public License, version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at https://mozilla.org/MPL/2.0/
 */

package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccIBMCisEdgeFunctionsActionsDataSource_basic(t *testing.T) {
	node := "data.ibm_cis_edge_functions_actions.test"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCisEdgeFunctionsActionsDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(node, "cis_edge_functions_actions.0.etag"),
				),
			},
		},
	})
}

func testAccCheckIBMCisEdgeFunctionsActionsDataSourceConfig() string {
	testName := "tf-acctest-basic"
	scriptName := "sample_script"

	return testAccCheckIBMCisEdgeFunctionsActionBasic(testName, scriptName) + fmt.Sprintf(`
	data "ibm_cis_edge_functions_actions" "test" {
		cis_id    = ibm_cis_edge_functions_action.tf-acctest-basic.cis_id
		domain_id = ibm_cis_edge_functions_action.tf-acctest-basic.domain_id
	  }`)
}
