// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package cis_test

import (
	"testing"

	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMCisPoolsDataSource_Basic(t *testing.T) {
	node := "data.ibm_cis_origin_pools.test"
	rnd := acctest.RandString(10)
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCisPoolsDataSourceConfig(rnd, acc.CisDomainStatic),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(node, "cis_origin_pools.0.id"),
					resource.TestCheckResourceAttrSet(node, "cis_origin_pools.0.pool_id"),
					resource.TestCheckResourceAttrSet(node, "cis_origin_pools.0.description"),
				),
			},
		},
	})
}

func testAccCheckIBMCisPoolsDataSourceConfig(resourceID, CisDomainStatic string) string {
	return testAccCheckCisPoolConfigCisDSBasic(resourceID, acc.CisDomainStatic) + `
	data "ibm_cis_origin_pools" "test" {
		cis_id    = ibm_cis_origin_pool.origin_pool.cis_id
	}
	`
}
