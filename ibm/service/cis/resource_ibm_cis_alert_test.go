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
		cis_id 		=  data.ibm_cis.cis.id
		name 		= "test-Webhooks"
		url			= "https://hooks.slack.com/services/Ds3fdBFbV/1234568"
		secret		=  "fBqWqLTwgx9aXuoOqLwenB6lIIyAYCvHJUowBS54Y3GC"
	  }
	resource "ibm_cis_alert" "test" {
		depends_on  = [ibm_cis_webhook.test]
		cis_id      = data.ibm_cis.cis.id
		name        = "test-alert-policy"
		description = "Description alert policy"
		enabled     = true
		alert_type = "g6_pool_toggle_alert"
		mechanisms {
		  email    = ["mynotifications@email.com"]
		  webhooks = [ibm_cis_webhook.test.webhook_id]
		}
		filters =<<FILTERS
		{
			"enabled": [
			  "true"
			],
			"pool_id": [
			  "9984f902f29adfc9bb8a5e42b7b5c592"
			]
		  }
  		FILTERS
		conditions =<<CONDITIONS
		{
			"and": [
			  {
				"or": [
				  {
					"==": [
					  {
						"var": "pool_id"
					  },
					  "9984f902f29adfc9bb8a5e42b7b5c592"
					]
				  }
				]
			  },
			  {
				"or": [
				  {
					"==": [
					  {
						"var": "enabled"
					  },
					  "true"
					]
				  }
				]
			  }
			]
		  }
		CONDITIONS
	}
`
}

func testAccCheckCisAlert_update() string {
	return testAccCheckIBMCisDomainDataSourceConfigBasic1() + `
	resource "ibm_cis_webhook"  "test"  {
		cis_id 		=  data.ibm_cis.cis.id
		name 		= "test-Webhooks"
		url			= "https://hooks.slack.com/services/Ds3fdBFbV/1234568"
		secret		=  "fBqWqLTwgx9aXuoOqLwenB6lIIyAYCvHJUowBS54Y3GC"
	  }
	resource "ibm_cis_alert" "test" {
		depends_on  = [ibm_cis_webhook.test]
		cis_id      =  data.ibm_cis.cis.id
		name        = "test-alert-policy-update"
		description = "Description alert policy update"
		enabled     = true
		alert_type = "g6_pool_toggle_alert"
		mechanisms {
		  email    = ["mynotifications@email.com"]
		  webhooks = [ibm_cis_webhook.test.webhook_id]
		}
		filters =<<FILTERS
		{
			"enabled": [
			  "true"
			],
			"pool_id": [
			  "9984f902f29adfc9bb8a5e42b7b5c592"
			]
		  }
  		FILTERS
		conditions =<<CONDITIONS
		{
			"and": [
			  {
				"or": [
				  {
					"==": [
					  {
						"var": "pool_id"
					  },
					  "9984f902f29adfc9bb8a5e42b7b5c592"
					]
				  }
				]
			  },
			  {
				"or": [
				  {
					"==": [
					  {
						"var": "enabled"
					  },
					  "true"
					]
				  }
				]
			  }
			]
		  }
		CONDITIONS
	}
`
}
