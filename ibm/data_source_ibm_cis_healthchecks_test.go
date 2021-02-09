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

func TestAccIBMCisGLBHealthCheckDataSource_basic(t *testing.T) {
	node := "data.ibm_cis_healthchecks.test"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
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
	return testAccCheckCisHealthcheckConfigFullySpecified("test", cisDomainStatic) + fmt.Sprintf(`
	data "ibm_cis_healthchecks" "test" {
		cis_id     = ibm_cis_healthcheck.health_check.cis_id
	  }`)
}
