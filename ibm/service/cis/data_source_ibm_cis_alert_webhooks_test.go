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
				Config: testAccCheckIBMCisWebhooksDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(node, "cis_webhooks_list.0.name"),
					resource.TestCheckResourceAttrSet(node, "cis_webhooks_list.0.url"),
					resource.TestCheckResourceAttrSet(node, "cis_webhooks_list.0.type"),
				),
			},
		},
	})
}

func testAccCheckIBMCisWebhooksDataSourceConfig() string {
	return fmt.Sprintf(`
	resource "ibm_cis_webhook"  "test" {
		cis_id 		= "crn:v1:staging:public:internet-svcs-ci:global:a/01652b251c3ae2787110a995d8db0135:79c0ce9a-f0fd-4c10-ae78-d890aca7a350::"
		name 		= "test-Webhooks"
		url			= "https://hooks.slack.com/services/Ds3fdBFbV/1234568"
		secret		=  "ZaHkdsdsadsdfsdfdsfsdffdsfsdfanchfnR4TISjOPC_I1U"
	  }
	data "ibm_cis_webhooks" "test" {
		cis_id = "crn:v1:staging:public:internet-svcs-ci:global:a/01652b251c3ae2787110a995d8db0135:79c0ce9a-f0fd-4c10-ae78-d890aca7a350::"
	  }`)
}
