// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package cis_test

import (
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccIBMCisWAFPackagesDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckCis(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config:  testAccCheckIBMCisWAFPackagesDataSourceConfig(),
				Destroy: true,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_cis_waf_packages.packages", "waf_packages.0.name"),
				),
			},
		},
	})
}

func testAccCheckIBMCisWAFPackagesDataSourceConfig() string {
	return testAccCheckIBMCisDomainDataSourceConfigBasic1() + `
	data "ibm_cis_waf_packages" "packages" {
		cis_id    = data.ibm_cis.cis.id
		domain_id = data.ibm_cis_domain.cis_domain.id
	  }
	`
}
