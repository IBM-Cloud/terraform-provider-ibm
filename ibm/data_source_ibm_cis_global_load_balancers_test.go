/*
* IBM Confidential
* Object Code Only Source Materials
* 5747-SM3
* (c) Copyright IBM Corp. 2017,2021
*
* The source code for this program is not published or otherwise divested
* of its trade secrets, irrespective of what has been deposited with the
* U.S. Copyright Office.
 */

package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccIBMCisGLBDataSource_basic(t *testing.T) {
	node := "data.ibm_cis_global_load_balancers.test"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
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
	return testAccCheckCisGlbConfigCisDSBasic("test", cisDomainStatic) + fmt.Sprintf(`
	data "ibm_cis_global_load_balancers" "test" {
		cis_id     = ibm_cis_global_load_balancer.test.cis_id
		domain_id  = ibm_cis_global_load_balancer.test.domain_id
	  }`)
}
