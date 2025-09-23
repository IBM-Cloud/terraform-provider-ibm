package cis_test

import (
	"fmt"
	"testing"

	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMCisOrigAuthDataSource_basic(t *testing.T) {
	node := "data.ibm_cis_origin_auths.test"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCisOrigAuthDataSourceZoneConfig("test"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(node, "origin_pull_settings_enabled"),
					resource.TestCheckResourceAttrSet(node, "origin_pull_certs_list.0.%"),
				),
			},
			{
				Config: testAccCheckIBMCisOrigAuthDataSourceHostnameConfig("test"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(node, "origin_pull_settings_enabled"),
				),
			},
		},
	})
}

func testAccCheckIBMCisOrigAuthDataSourceZoneConfig(id string) string {
	return testAccCheckIBMCisDomainDataSourceConfigBasic1() + fmt.Sprintf(`
	data "ibm_cis_origin_auths" "%[1]s" {
		cis_id          = data.ibm_cis.cis.id
		domain_id 		= data.ibm_cis_domain.cis_domain.domain_id
	  }
`, id)
}

func testAccCheckIBMCisOrigAuthDataSourceHostnameConfig(id string) string {
	return testAccCheckIBMCisDomainDataSourceConfigBasic1() + fmt.Sprintf(`
	data "ibm_cis_origin_auths" "%[1]s" {
		cis_id          = data.ibm_cis.cis.id
		domain_id 		= data.ibm_cis_domain.cis_domain.domain_id
		request_type    = "per_hostname"
		hostname        = data.ibm_cis_domain.cis_domain.domain
	}
`, id)
}
