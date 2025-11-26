// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package cis_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccIBMCisCustomPage_Basic(t *testing.T) {
	name := "ibm_cis_custom_page." + "test"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckCis(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCisCustomPageConfigBasic1("test", acc.CisDomainStatic),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "page_id", "basic_challenge"),
					resource.TestCheckResourceAttr(name, "url",
						"http://customtest.cis-test-domain.com/index.html"),
				),
			},
			{
				Config: testAccCheckCisCustomPageConfigBasic2("test", acc.CisDomainStatic),
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
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCisCustomPageConfigBasic1("test", acc.CisDomainStatic),
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

func testAccCheckCisCustomPageConfigBasic1(id string, CisDomainStatic string) string {
	return testAccCheckIBMCisDomainDataSourceConfigBasic1() + fmt.Sprintf(`
	resource "ibm_cis_custom_page" "%[1]s" {
		cis_id    = data.ibm_cis.cis.id
		domain_id = data.ibm_cis_domain.cis_domain.domain_id
		page_id   = "basic_challenge"
		url       = "http://customtest.cis-test-domain.com/index.html"
	  }
`, id)
}
func testAccCheckCisCustomPageConfigBasic2(id string, CisDomainStatic string) string {
	return testAccCheckIBMCisDomainDataSourceConfigBasic1() + fmt.Sprintf(`
	resource "ibm_cis_custom_page" "%[1]s" {
		cis_id    = data.ibm_cis.cis.id
		domain_id = data.ibm_cis_domain.cis_domain.domain_id
		page_id   = "basic_challenge"
		url       = ""
	  }
`, id)
}
