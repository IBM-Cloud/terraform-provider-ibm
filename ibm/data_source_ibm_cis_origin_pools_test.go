/* IBM Confidential
*  Object Code Only Source Materials
*  5747-SM3
*  (c) Copyright IBM Corp. 2017,2021
*
*  The source code for this program is not published or otherwise divested
*  of its trade secrets, irrespective of what has been deposited with the
*  U.S. Copyright Office. */

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
