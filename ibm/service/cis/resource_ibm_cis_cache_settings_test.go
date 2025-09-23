// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package cis_test

import (
	"fmt"
	"testing"

	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMCisCacheSettings_Basic(t *testing.T) {
	name := "ibm_cis_cache_settings." + "test"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckCis(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCisCacheSettingsConfigBasic1("test", acc.CisDomainStatic),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "caching_level", "simplified"),
					resource.TestCheckResourceAttr(name, "browser_expiration", "7200"),
					resource.TestCheckResourceAttr(name, "development_mode", "on"),
					resource.TestCheckResourceAttr(name, "query_string_sort", "on"),
					resource.TestCheckResourceAttr(name, "serve_stale_content", "on"),
				),
			},
			{
				Config: testAccCheckCisCacheSettingsConfigBasic2("test", acc.CisDomainStatic),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "caching_level", "aggressive"),
					resource.TestCheckResourceAttr(name, "browser_expiration", "14400"),
					resource.TestCheckResourceAttr(name, "development_mode", "off"),
					resource.TestCheckResourceAttr(name, "query_string_sort", "off"),
					resource.TestCheckResourceAttr(name, "serve_stale_content", "off"),
				),
			},
		},
	})
}

func TestAccIBMCisCacheSettings_WithoutPurgeAction(t *testing.T) {
	name := "ibm_cis_cache_settings." + "test"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckCis(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCisCacheSettingsConfigWithoutPurgeAction("test", acc.CisDomainStatic),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "caching_level", "simplified"),
					resource.TestCheckResourceAttr(name, "browser_expiration", "7200"),
					resource.TestCheckResourceAttr(name, "development_mode", "on"),
					resource.TestCheckResourceAttr(name, "query_string_sort", "on"),
					resource.TestCheckResourceAttr(name, "serve_stale_content", "on"),
				),
			},
			{
				Config: testAccCheckCisCacheSettingsConfigBasic2("test", acc.CisDomainStatic),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "caching_level", "aggressive"),
					resource.TestCheckResourceAttr(name, "browser_expiration", "14400"),
					resource.TestCheckResourceAttr(name, "development_mode", "off"),
					resource.TestCheckResourceAttr(name, "query_string_sort", "off"),
					resource.TestCheckResourceAttr(name, "serve_stale_content", "off"),
				),
			},
		},
	})
}

func TestAccIBMCisCacheSettings_Import(t *testing.T) {
	name := "ibm_cis_cache_settings." + "test"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCisCacheSettingsConfigBasic2("test", acc.CisDomainStatic),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "caching_level", "aggressive"),
					resource.TestCheckResourceAttr(name, "browser_expiration", "14400"),
					resource.TestCheckResourceAttr(name, "development_mode", "off"),
					resource.TestCheckResourceAttr(name, "query_string_sort", "off"),
					resource.TestCheckResourceAttr(name, "serve_stale_content", "off"),
				),
			},
			{
				ResourceName:            name,
				ImportState:             true,
				ImportStateVerifyIgnore: []string{"purge"},
			},
		},
	})
}

func testAccCheckCisCacheSettingsConfigBasic1(id string, CisDomainStatic string) string {
	return testAccCheckIBMCisDomainDataSourceConfigBasic1() + fmt.Sprintf(`
	resource "ibm_cis_cache_settings" "%[1]s" {
		cis_id          = data.ibm_cis.cis.id
		domain_id       = data.ibm_cis_domain.cis_domain.domain_id
		caching_level      = "simplified"
		browser_expiration = 7200
		development_mode   = "on"
		query_string_sort  = "on"
		purge_all          = true
		serve_stale_content = "on"
	  }
`, id)
}
func testAccCheckCisCacheSettingsConfigBasic2(id string, CisDomainStatic string) string {
	return testAccCheckIBMCisDomainDataSourceConfigBasic1() + fmt.Sprintf(`
	resource "ibm_cis_cache_settings" "%[1]s" {
		cis_id          = data.ibm_cis.cis.id
		domain_id       = data.ibm_cis_domain.cis_domain.domain_id
		caching_level      = "aggressive"
		browser_expiration = 14400
		development_mode   = "off"
		query_string_sort  = "off"
		purge_by_urls      = ["http://test.com/index.html", "http://example.com/index.html"]
		serve_stale_content = "off"
	  }
`, id)
}

func testAccCheckCisCacheSettingsConfigWithoutPurgeAction(id string, CisDomainStatic string) string {
	return testAccCheckIBMCisDomainDataSourceConfigBasic1() + fmt.Sprintf(`
	resource "ibm_cis_cache_settings" "%[1]s" {
		cis_id          = data.ibm_cis.cis.id
		domain_id       = data.ibm_cis_domain.cis_domain.domain_id
		caching_level      = "simplified"
		browser_expiration = 7200
		development_mode   = "on"
		query_string_sort  = "on"
		serve_stale_content = "on"
	  }
`, id)
}
