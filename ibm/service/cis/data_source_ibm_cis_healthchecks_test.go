// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package cis_test

import (
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccIBMCisGLBHealthCheckDataSource_basic(t *testing.T) {
	node := "data.ibm_cis_healthchecks.test"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCisGLBHealthCheckDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(node, "cis_healthchecks.0.id"),
					resource.TestCheckResourceAttrSet(node, "cis_healthchecks.0.monitor_id"),
					resource.TestCheckResourceAttrSet(node, "cis_healthchecks.0.type"),
				),
			},
		},
	})
}

func testAccCheckIBMCisGLBHealthCheckDataSourceConfig() string {
	// status filter defaults to empty
	return testAccCheckCisHealthcheckConfigFullySpecified("test", acc.CisDomainStatic) + `
	data "ibm_cis_healthchecks" "test" {
		cis_id     = ibm_cis_healthcheck.health_check.cis_id
	  }`
}
