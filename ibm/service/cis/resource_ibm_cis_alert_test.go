// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package cis_test

import (
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMCisAlert_Basic(t *testing.T) {

	alertname := "test-alert-policy"
	alertdesc := "Description alert policy"

	alertnameUpdate := "test-alert-policy-update"
	alertdescUpdate := "Description alert policy update"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckCis(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCisAlert_basic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_cis_alert.test", "name", alertname),
					resource.TestCheckResourceAttr("ibm_cis_alert.test", "description", alertdesc),
					resource.TestCheckResourceAttr("ibm_cis_alert.test", "enabled", "true"),
					resource.TestCheckResourceAttrSet("ibm_cis_alert.test", "filters"),
				),
			},
			{
				Config: testAccCheckCisAlert_update(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_cis_alert.test", "name", alertnameUpdate),
					resource.TestCheckResourceAttr("ibm_cis_alert.test", "description", alertdescUpdate),
					resource.TestCheckResourceAttr("ibm_cis_alert.test", "enabled", "true"),
					resource.TestCheckResourceAttrSet("ibm_cis_alert.test", "conditions"),
				),
			},
		},
	})
}
func testAccCheckCisAlert_basic() string {
	return testAccCheckIBMCisDomainDataSourceConfigBasic1() + `
	resource "ibm_cis_webhook"  "test"  {
		cis_id 		= "crn:v1:staging:public:internet-svcs-ci:global:a/01652b251c3ae2787110a995d8db0135:79c0ce9a-f0fd-4c10-ae78-d890aca7a350::"
		name 		= "test-Webhooks"
		url			= "https://hooks.slack.com/services/Ds3fdBFbV/1234568"
		secret		=  "ZaHkAf0iNXNWn8ySUJjTJHkzlanchfnR4TISjOPC_I1U"
	  }
	resource "ibm_cis_alert" "test" {
		depends_on  = [ibm_cis_webhook.test]
		cis_id      = "crn:v1:staging:public:internet-svcs-ci:global:a/01652b251c3ae2787110a995d8db0135:79c0ce9a-f0fd-4c10-ae78-d890aca7a350::"
		name        = "test-alert-policy"
		description = "Description alert policy"
		enabled     = true
		alert_type = "dos_attack_l7"
		mechanisms {
		  email    = ["mynotifications@email.com"]
		  webhooks = [ibm_cis_webhook.test.webhook_id]
		}
		filters =<<FILTERS
  		{}
  		FILTERS
		conditions =<<CONDITIONS
  		{}
  		CONDITIONS
	}
`
}

func testAccCheckCisAlert_update() string {
	return testAccCheckIBMCisDomainDataSourceConfigBasic1() + `
	resource "ibm_cis_webhook"  "test"  {
		cis_id 		= "crn:v1:staging:public:internet-svcs-ci:global:a/01652b251c3ae2787110a995d8db0135:79c0ce9a-f0fd-4c10-ae78-d890aca7a350::"
		name 		= "test-Webhooks"
		url			= "https://hooks.slack.com/services/Ds3fdBFbV/1234568"
		secret		=  "ZaHkAf0iNXNWn8ySUJjTJHkzlanchfnR4TISjOPC_I1U"
	  }
	resource "ibm_cis_alert" "test" {
		depends_on  = [ibm_cis_webhook.test]
		cis_id      = "crn:v1:staging:public:internet-svcs-ci:global:a/01652b251c3ae2787110a995d8db0135:79c0ce9a-f0fd-4c10-ae78-d890aca7a350::"
		name        = "test-alert-policy-update"
		description = "Description alert policy update"
		enabled     = true
		alert_type = "dos_attack_l7"
		mechanisms {
		  email    = ["mynotifications@email.com"]
		  webhooks = [ibm_cis_webhook.test.webhook_id]
		}
		filters =<<FILTERS
  		{}
  		FILTERS
		conditions =<<CONDITIONS
  		{}
  		CONDITIONS
	}
`
}
