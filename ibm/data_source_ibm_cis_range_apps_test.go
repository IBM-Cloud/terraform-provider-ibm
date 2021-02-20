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

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccIBMCisRangeAppsDataSource_basic(t *testing.T) {
	node := "data.ibm_cis_range_apps.test"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCisRangeAppsDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(node, "range_apps.0.id"),
					resource.TestCheckResourceAttrSet(node, "range_apps.0.origin_direct.0"),
				),
			},
		},
	})
}

func testAccCheckIBMCisRangeAppsDataSourceConfig() string {
	return testAccCheckCisRangeAppConfigBasic() + fmt.Sprintf(`
	data "ibm_cis_range_apps" "test" {
		cis_id     = ibm_cis_range_app.app.cis_id
		domain_id  = ibm_cis_range_app.app.domain_id
	}`)
}
