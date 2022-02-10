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
	webhooksecret := "Zdsfs3e23k223sfsdffdsfsdfanchfnR4TISjOPC_I1U"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckCis(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCisWebhooks_basic("test", acc.CisDomainStatic, webhookname, webhooksurl, webhooksecret),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_cis_webhook.test", "name", webhookname),
					resource.TestCheckResourceAttr("ibm_cis_webhook.test", "url", webhooksurl),
					resource.TestCheckResourceAttr("ibm_cis_webhook.test", "secret", webhooksecret),
				),
			},
		},
	})
}
func testAccCheckCisWebhooks_basic(id, CisDomainStatic, name, url, secret string) string {
	return testAccCheckIBMCisDomainDataSourceConfigBasic1() + fmt.Sprintf(`
	resource "ibm_cis_webhook"  "%[1]s"  {
		cis_id 		= ata.ibm_cis.cis.id
		name 		= "%s"
		url			= "%s"
		secret		=  "%s"
	  }
`, id, acc.CisDomainStatic, name, url, secret)
}
