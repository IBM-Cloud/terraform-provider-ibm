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

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccIBMCisPoolsDataSource_Basic(t *testing.T) {
	node := "data.ibm_cis_origin_pools.test"
	rnd := acctest.RandString(10)
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCisPoolsDataSourceConfig(rnd, cisDomainStatic),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(node, "cis_origin_pools.0.id"),
					resource.TestCheckResourceAttrSet(node, "cis_origin_pools.0.pool_id"),
					resource.TestCheckResourceAttrSet(node, "cis_origin_pools.0.description"),
				),
			},
		},
	})
}

func testAccCheckIBMCisPoolsDataSourceConfig(resourceID, cisDomainStatic string) string {
	return testAccCheckCisPoolConfigCisDSBasic(resourceID, cisDomainStatic) + fmt.Sprintf(`
	data "ibm_cis_origin_pools" "test" {
		cis_id    = ibm_cis_origin_pool.origin_pool.cis_id
	}
	`)
}
