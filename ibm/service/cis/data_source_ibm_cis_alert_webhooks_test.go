package cis_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMCisWebhookskDataSource_basic(t *testing.T) {
	node := "data.ibm_cis_webhooks.test"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCisWebhooksDataSourceConfig("test", acc.CisDomainStatic),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(node, "cis_webhooks.0.name"),
					resource.TestCheckResourceAttrSet(node, "cis_webhooks.0.url"),
					resource.TestCheckResourceAttrSet(node, "cis_webhooks.0.type"),
				),
			},
		},
	})
}

func testAccCheckIBMCisWebhooksDataSourceConfig(id, CisDomainStatic string) string {
	return testAccCheckIBMCisDomainDataSourceConfigBasic1() + fmt.Sprintf(`
	data "ibm_cis_webhooks" "%[1]s" {
		cis_id          = data.ibm_cis.cis.id
	  }
`, id, acc.CisDomainStatic)
}
