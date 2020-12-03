package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccIBMCisSettings_Basic(t *testing.T) {
	name := "ibm_cis_domain_settings." + "test"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheckCis(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCisSettingsConfigBasic3("test", cisDomainStatic),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "waf", "off"),
					resource.TestCheckResourceAttr(name, "min_tls_version", "1.1"),
				),
			},
			{
				Config: testAccCheckCisSettingsConfigBasic1("test", cisDomainStatic),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "waf", "on"),
					resource.TestCheckResourceAttr(name, "ssl", "full"),
					resource.TestCheckResourceAttr(name, "min_tls_version", "1.2"),
				),
			},
			{
				Config: testAccCheckCisSettingsConfigBasic2("test", cisDomainStatic),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "waf", "off"),
					resource.TestCheckResourceAttr(name, "ssl", "flexible"),
					resource.TestCheckResourceAttr(name, "min_tls_version", "1.1"),
				),
			},
			{
				Config: testAccCheckCisSettingsConfigBasic4("test", cisDomainStatic),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "waf", "off"),
					resource.TestCheckResourceAttr(name, "ssl", "flexible"),
				),
			},
		},
	})
}

func TestAccIBMCisSettings_Import(t *testing.T) {
	name := "ibm_cis_domain_settings." + "test"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCisSettingsConfigBasic4("test", cisDomainStatic),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "waf", "off"),
					resource.TestCheckResourceAttr(name, "ssl", "flexible"),
				),
			},
			{
				ResourceName:      name,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckCisSettingsConfigBasic3(id string, cisDomainStatic string) string {
	return testAccCheckIBMCisDomainDataSourceConfigBasic1() + fmt.Sprintf(`
	resource "ibm_cis_domain_settings" "%[1]s" {
		cis_id          = data.ibm_cis.cis.id
		domain_id       = data.ibm_cis_domain.cis_domain.id
		waf = "off"
		min_tls_version = "1.1"
	  }
`, id)
}

func testAccCheckCisSettingsConfigBasic1(id string, cisDomainStatic string) string {
	return testAccCheckIBMCisDomainDataSourceConfigBasic1() + fmt.Sprintf(`
	resource "ibm_cis_domain_settings" "%[1]s" {
		cis_id          = data.ibm_cis.cis.id
		domain_id       = data.ibm_cis_domain.cis_domain.id
		waf             = "on"
		ssl             = "full"
		min_tls_version = "1.2"
	  }
`, id)
}

func testAccCheckCisSettingsConfigBasic2(id string, cisDomainStatic string) string {
	return testAccCheckIBMCisDomainDataSourceConfigBasic1() + fmt.Sprintf(`
	resource "ibm_cis_domain_settings" "%[1]s" {
		cis_id          = data.ibm_cis.cis.id
		domain_id       = data.ibm_cis_domain.cis_domain.id
		waf             = "off"
		ssl             = "flexible"
		min_tls_version = "1.1"
	  }
`, id)
}

func testAccCheckCisSettingsConfigBasic4(id string, cisDomainStatic string) string {
	return testAccCheckIBMCisDomainDataSourceConfigBasic1() + fmt.Sprintf(`
	resource "ibm_cis_domain_settings" "%[1]s" {
		cis_id          = data.ibm_cis.cis.id
		domain_id       = data.ibm_cis_domain.cis_domain.domain_id
		waf                         = "off"
		ssl                         = "flexible"
		min_tls_version             = "1.2"
		cname_flattening            = "flatten_all"
		opportunistic_encryption    = "off"
		automatic_https_rewrites    = "on"
		always_use_https            = "off"
		ipv6                        = "off"
		browser_check               = "off"
		hotlink_protection          = "on"
		http2                       = "on"
		image_load_optimization     = "on"
		image_size_optimization     = "lossless"
		ip_geolocation              = "on"
		origin_error_page_pass_thru = "on"
		pseudo_ipv4              = "off"
		prefetch_preload         = "on"
		response_buffering       = "on"
		script_load_optimization = "on"
		server_side_exclude      = "on"
		tls_client_auth          = "on"
		true_client_ip_header    = "on"
		websockets               = "on"
		challenge_ttl            = 3600
		max_upload               = 300
		cipher                   = ["AES128-GCM-SHA256"]
		minify {
		  css  = "on"
		  js   = "on"
		  html = "on"
		}
		security_header {
		  enabled            = false
		  include_subdomains = true
		  max_age            = 100
		  nosniff            = false
		}
		mobile_redirect {
		  status           = "off"
		}
	  }
`, id)
}
