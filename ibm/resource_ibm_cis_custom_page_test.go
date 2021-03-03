// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccIBMCisCustomPage_Basic(t *testing.T) {
	name := "ibm_cis_custom_page." + "test"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheckCis(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCisCustomPageConfigBasic1("test", cisDomainStatic),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "page_id", "basic_challenge"),
					resource.TestCheckResourceAttr(name, "url",
						"http://customtest.cis-test-domain.com/index.html"),
				),
			},
			{
				Config: testAccCheckCisCustomPageConfigBasic2("test", cisDomainStatic),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "page_id", "basic_challenge"),
					resource.TestCheckResourceAttr(name, "url", ""),
				),
			},
		},
	})
}

func TestAccIBMCisCustomPage_Import(t *testing.T) {
	name := "ibm_cis_custom_page." + "test"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCisCustomPageConfigBasic1("test", cisDomainStatic),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "page_id", "basic_challenge"),
					resource.TestCheckResourceAttr(name, "url",
						"http://customtest.cis-test-domain.com/index.html"),
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

func testAccCheckCisCustomPageConfigBasic1(id string, cisDomainStatic string) string {
	return testAccCheckIBMCisDomainDataSourceConfigBasic1() + fmt.Sprintf(`
	resource "ibm_cis_custom_page" "%[1]s" {
		cis_id    = data.ibm_cis.cis.id
		domain_id = data.ibm_cis_domain.cis_domain.domain_id
		page_id   = "basic_challenge"
		url       = "http://customtest.cis-test-domain.com/index.html"
	  }
`, id)
}
func testAccCheckCisCustomPageConfigBasic2(id string, cisDomainStatic string) string {
	return testAccCheckIBMCisDomainDataSourceConfigBasic1() + fmt.Sprintf(`
	resource "ibm_cis_custom_page" "%[1]s" {
		cis_id    = data.ibm_cis.cis.id
		domain_id = data.ibm_cis_domain.cis_domain.domain_id
		page_id   = "basic_challenge"
		url       = ""
	  }
`, id)
}
