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

func TestAccIBMCisPageRuleDataSource_basic(t *testing.T) {
	node := "data.ibm_cis_page_rules.test"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCisPageRuleDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(node, "cis_page_rules.0.id"),
					resource.TestCheckResourceAttrSet(node, "cis_page_rules.0.rule_id"),
				),
			},
		},
	})
}

func testAccCheckIBMCisPageRuleDataSourceConfig() string {
	// status filter defaults to empty
	return testAccCheckIBMCisPageRuleConfigBasic() + fmt.Sprintf(`
	data "ibm_cis_page_rules" "test" {
		cis_id     = ibm_cis_page_rule.page_rule.cis_id
		domain_id  = ibm_cis_page_rule.page_rule.domain_id
	  }`)
}
