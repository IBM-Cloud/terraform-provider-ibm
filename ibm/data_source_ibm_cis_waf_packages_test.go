package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccIBMCisWAFPackagesDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheckCis(t) },
		Providers: testAccProviders,
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
	return testAccCheckIBMCisDomainDataSourceConfigBasic1() + fmt.Sprintf(`
	data "ibm_cis_waf_packages" "packages" {
		cis_id    = data.ibm_cis.cis.id
		domain_id = data.ibm_cis_domain.cis_domain.id
	  }
	`)
}
