// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package cis_test

import (
	"testing"

	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMCisCacheSettingsDataSource_Basic(t *testing.T) {
	node := "data.ibm_cis_cache_settings.test"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCisCacheSettingsDataSourceConfigBasic1("test", acc.CisDomainStatic),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(node, "caching_level.0.value", "simplified"),
					resource.TestCheckResourceAttr(node, "development_mode.0.value", "on"),
					resource.TestCheckResourceAttr(node, "query_string_sort.0.value", "on"),
					resource.TestCheckResourceAttr(node, "serve_stale_content.0.value", "on"),
				),
			},
			{
				Config: testAccCheckCisCacheSettingsDataSourceConfigBasic2("test", acc.CisDomainStatic),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(node, "caching_level.0.value", "aggressive"),
					resource.TestCheckResourceAttr(node, "development_mode.0.value", "off"),
					resource.TestCheckResourceAttr(node, "query_string_sort.0.value", "off"),
					resource.TestCheckResourceAttr(node, "serve_stale_content.0.value", "off"),
				),
			},
		},
	})
}

func testAccCheckCisCacheSettingsDataSourceConfigBasic1(id string, CisDomainStatic string) string {
	return testAccCheckCisCacheSettingsConfigBasic1(id, acc.CisDomainStatic) + `
	  data "ibm_cis_cache_settings" "test" {
		cis_id    = data.ibm_cis.cis.id
		domain_id = data.ibm_cis_domain.cis_domain.domain_id
	  }
`
}
func testAccCheckCisCacheSettingsDataSourceConfigBasic2(id string, CisDomainStatic string) string {
	return testAccCheckCisCacheSettingsConfigBasic2(id, acc.CisDomainStatic) + `
	  data "ibm_cis_cache_settings" "test" {
		cis_id    = data.ibm_cis.cis.id
		domain_id = data.ibm_cis_domain.cis_domain.domain_id
	  }
`
}
