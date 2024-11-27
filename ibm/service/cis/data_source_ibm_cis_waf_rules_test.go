// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package cis_test

import (
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccIBMCisWAFRulesDataSource_Basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCisWAFRulesDataSourceConfigBasic1(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_cis_waf_rules.rules", "waf_rules.0.package_id"),
				),
			},
		},
	})
}

func testAccCheckIBMCisWAFRulesDataSourceConfigBasic1() string {
	return testAccCheckIBMCisDomainDataSourceConfigBasic1() + `
	data "ibm_cis_waf_rules" "rules" {
		cis_id    = data.ibm_cis.cis.id
		domain_id = data.ibm_cis_domain.cis_domain.id
		package_id = "1e334934fd7ae32ad705667f8c1057aa"
	}
	`
}
