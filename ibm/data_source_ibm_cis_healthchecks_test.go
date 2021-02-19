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
