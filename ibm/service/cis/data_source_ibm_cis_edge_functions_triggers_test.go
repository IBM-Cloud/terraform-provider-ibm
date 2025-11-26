// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package cis_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccIBMCisEdgeFunctionsTriggersDataSource_basic(t *testing.T) {
	node := "data.ibm_cis_edge_functions_triggers.test"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCisEdgeFunctionsTriggersDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(node, "cis_edge_functions_triggers.0.pattern_url"),
				),
			},
		},
	})
}

func testAccCheckIBMCisEdgeFunctionsTriggersDataSourceConfig() string {
	testName := "test"
	scriptName := "sample_script"
	pattern := fmt.Sprintf("example.%s/*", acc.CisDomainStatic)
	return testAccCheckIBMCisEdgeFunctionsTriggerBasic(testName, pattern, scriptName) + `
	data "ibm_cis_edge_functions_triggers" "test" {
		cis_id    = ibm_cis_edge_functions_trigger.test.cis_id
		domain_id = ibm_cis_edge_functions_trigger.test.domain_id
	}`
}
