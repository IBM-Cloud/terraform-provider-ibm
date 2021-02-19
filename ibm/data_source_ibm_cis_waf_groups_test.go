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

func TestAccIBMCisWAFGroupsDataSource_Basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCisWAFGroupsDataSourceConfigBasic1(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_cis_waf_groups.waf_groups", "waf_groups.0.group_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cis_waf_groups.waf_groups", "waf_groups.0.modified_rules_count"),
				),
			},
		},
	})
}

func testAccCheckIBMCisWAFGroupsDataSourceConfigBasic1() string {
	return testAccCheckIBMCisDomainDataSourceConfigBasic1() + fmt.Sprintf(`
	data "ibm_cis_waf_groups" "waf_groups" {
		cis_id     = data.ibm_cis.cis.id
		domain_id  = data.ibm_cis_domain.cis_domain.id
		package_id = "c504870194831cd12c3fc0284f294abb"
	}
	`)
}
