// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccIBMCisWAFPackage_Basic(t *testing.T) {
	name := "ibm_cis_waf_package." + "test"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheckCis(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCisWAFPackageConfigBasic1("test", cisDomainStatic),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "sensitivity", "low"),
					resource.TestCheckResourceAttr(name, "action_mode", "block"),
				),
			},
			{
				Config: testAccCheckCisWAFPackageConfigBasic2("test", cisDomainStatic),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "sensitivity", "off"),
					resource.TestCheckResourceAttr(name, "action_mode", "challenge"),
				),
			},
		},
	})
}

func TestAccIBMCisWAFPackage_Import(t *testing.T) {
	name := "ibm_cis_waf_package." + "test"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCisWAFPackageConfigBasic2("test", cisDomainStatic),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "sensitivity", "off"),
					resource.TestCheckResourceAttr(name, "action_mode", "challenge"),
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

func testAccCheckCisWAFPackageConfigBasic1(id string, cisDomainStatic string) string {
	return testAccCheckIBMCisDomainDataSourceConfigBasic1() + fmt.Sprintf(`
	resource "ibm_cis_waf_package" "%[1]s" {
		cis_id      = data.ibm_cis.cis.id
		domain_id   = data.ibm_cis_domain.cis_domain.id
		package_id  = "c504870194831cd12c3fc0284f294abb:546"
		sensitivity = "low"
		action_mode = "block"
	  }
`, id)
}
func testAccCheckCisWAFPackageConfigBasic2(id string, cisDomainStatic string) string {
	return testAccCheckIBMCisDomainDataSourceConfigBasic1() + fmt.Sprintf(`
	resource "ibm_cis_waf_package" "%[1]s" {
		cis_id      = data.ibm_cis.cis.id
		domain_id   = data.ibm_cis_domain.cis_domain.id
		package_id  = "c504870194831cd12c3fc0284f294abb"
		sensitivity = "off"
		action_mode = "challenge"
	  }
`, id)
}
