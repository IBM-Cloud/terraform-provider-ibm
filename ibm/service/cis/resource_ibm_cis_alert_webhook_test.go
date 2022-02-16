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
	webhooksecret := "ff1d9b80-b51d-4a06-bf67-6752fae1eb74"

	updatewebhookname := "update-test-Webhooks"
	updatewebhooksurl := "https://hooks.slack.com/services/Ds3fdBFbV/1234568"
	updatewebhooksecret := "ff1d9b80-b51d-4a06-bf67-6752fae1eb34"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckCis(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCisWebhooksBasic1("test", acc.CisDomainStatic),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_cis_webhook.test", "name", webhookname),
					resource.TestCheckResourceAttr("ibm_cis_webhook.test", "url", webhooksurl),
					resource.TestCheckResourceAttr("ibm_cis_webhook.test", "secret", webhooksecret),
				),
			},
			{
				Config: testAccCheckCisWebhooksBasic2("test", acc.CisDomainStatic),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_cis_webhook.test", "name", updatewebhookname),
					resource.TestCheckResourceAttr("ibm_cis_webhook.test", "url", updatewebhooksurl),
					resource.TestCheckResourceAttr("ibm_cis_webhook.test", "secret", updatewebhooksecret),
				),
			},
		},
	})
}
func TestAccIBMCisWebhooks_Import(t *testing.T) {
	webhookname := "test-Webhooks"
	webhooksurl := "https://hooks.slack.com/services/Ds3fdBFbV/1234568"
	webhooksecret := "ff1d9b80-b51d-4a06-bf67-6752fae1eb74"
	name := "ibm_cis_webhook.test"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCisWebhooksBasic1("test", acc.CisDomainStatic),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_cis_webhook.test", "name", webhookname),
					resource.TestCheckResourceAttr("ibm_cis_webhook.test", "url", webhooksurl),
					resource.TestCheckResourceAttr("ibm_cis_webhook.test", "secret", webhooksecret),
				),
			},
			{
				ResourceName:      name,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"secret"},
			},
		},
	})
}

func testAccCheckCisWebhooksBasic1(id, CisDomainStatic string) string {
	return testAccCheckIBMCisDomainDataSourceConfigBasic1() + fmt.Sprintf(`
	resource "ibm_cis_webhook"  "%[1]s"  {
		cis_id 		= data.ibm_cis.cis.id
		name 		= "test-Webhooks"
		url			= "https://hooks.slack.com/services/Ds3fdBFbV/1234568"
		secret		=  "ff1d9b80-b51d-4a06-bf67-6752fae1eb74"
	  }
`, id)
}

func testAccCheckCisWebhooksBasic2(id, CisDomainStatic string) string {
	return testAccCheckIBMCisDomainDataSourceConfigBasic1() + fmt.Sprintf(`
	resource "ibm_cis_webhook"  "%[1]s"  {
		cis_id 		= data.ibm_cis.cis.id
		name 		= "update-test-Webhooks"
		url			= "https://hooks.slack.com/services/Ds3fdBFbV/1234568"
		secret		=  "ff1d9b80-b51d-4a06-bf67-6752fae1eb34"
	  }
`, id)
}
