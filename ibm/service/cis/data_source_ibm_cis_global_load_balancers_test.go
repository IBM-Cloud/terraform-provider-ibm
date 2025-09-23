// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package cis_test

import (
	"testing"

	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMCisGLBDataSource_basic(t *testing.T) {
	node := "data.ibm_cis_global_load_balancers.test"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCisGLBDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(node, "cis_glb.0.id"),
					resource.TestCheckResourceAttrSet(node, "cis_glb.0.glb_id"),
				),
			},
		},
	})
}

func testAccCheckIBMCisGLBDataSourceConfig() string {
	// status filter defaults to empty
	return testAccCheckCisGlbConfigCisDSBasic("test", acc.CisDomainStatic) + `
	data "ibm_cis_global_load_balancers" "test" {
		cis_id     = ibm_cis_global_load_balancer.test.cis_id
		domain_id  = ibm_cis_global_load_balancer.test.domain_id
	  }`
}
