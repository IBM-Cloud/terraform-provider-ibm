// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package cis_test

import (
	"testing"

	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMCisEdgeFunctionsActionsDataSource_basic(t *testing.T) {
	node := "data.ibm_cis_edge_functions_actions.test"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
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

	return testAccCheckIBMCisEdgeFunctionsActionBasic(testName, scriptName) + `
	data "ibm_cis_edge_functions_actions" "test" {
		cis_id    = ibm_cis_edge_functions_action.tf-acctest-basic.cis_id
		domain_id = ibm_cis_edge_functions_action.tf-acctest-basic.domain_id
	  }`
}
