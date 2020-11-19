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
