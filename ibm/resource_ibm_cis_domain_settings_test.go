package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccIBMCisSettings_Basic(t *testing.T) {
	// multiple instances of this config would conflict but we only use it once
	//t.Parallel()

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
					resource.TestCheckResourceAttr(name, "min_tls_version", "1.3"),
				),
			},
		},
	})
}

func testAccCheckCisSettingsConfigBasic3(id string, cisDomainStatic string) string {
	return testAccCheckIBMCisDomainDataSourceConfigBasic1() + fmt.Sprintf(`
	resource "ibm_cis_domain_settings" "%[1]s" {
		cis_id          = data.ibm_cis.cis.id
		domain_id       = data.ibm_cis_domain.cis_domain.id
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
		min_tls_version = "1.3"
	  }
`, id)
}
