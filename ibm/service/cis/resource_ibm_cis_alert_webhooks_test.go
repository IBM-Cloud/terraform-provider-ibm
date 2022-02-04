// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package cis_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMCisWebhooks_Basic(t *testing.T) {

	webhookname := "test-Webhooks"
	webhooksurl := "https://hooks.slack.com/services/Ds3fdBFbV/1234568"
	webhooksecret := "ZaHkdsdsadsdfsdfdsfsdffdsfsdfanchfnR4TISjOPC_I1U"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckCis(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCisWebhooks_basic(webhookname, webhooksurl, webhooksecret),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_cis_webhook.test", "name", webhookname),
					resource.TestCheckResourceAttr("ibm_cis_webhook.test", "url", webhooksurl),
					resource.TestCheckResourceAttr("ibm_cis_webhook.test", "secret", webhooksecret),
				),
			},
		},
	})
}
func testAccCheckCisWebhooks_basic(name, url, secret string) string {
	return fmt.Sprintf(`
	resource "ibm_cis_webhook"  "test" {
		cis_id 		= "crn:v1:staging:public:internet-svcs-ci:global:a/01652b251c3ae2787110a995d8db0135:79c0ce9a-f0fd-4c10-ae78-d890aca7a350::"
		name 		= "%s"
		url			= "%s"
		secret		=  "%s"
	  }
`, name, url, secret)
}
