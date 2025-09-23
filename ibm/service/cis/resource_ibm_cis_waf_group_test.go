// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package cis_test

import (
	"fmt"
	"testing"

	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMCisWAFGroup_Basic(t *testing.T) {
	name := "ibm_cis_waf_group." + "test"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckCis(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCisWAFGroupConfigBasic1("test", acc.CisDomainStatic),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "mode", "on"),
				),
			},
			{
				Config: testAccCheckCisWAFGroupConfigBasic2("test", acc.CisDomainStatic),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "mode", "off"),
				),
			},
			{
				Config: testAccCheckCisWAFGroupConfigBasic1("test", acc.CisDomainStatic),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "mode", "on"),
				),
			},
			{
				Config: testAccCheckCisWAFGroupConfigBasic3("test", acc.CisDomainStatic),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "mode", "on"),
				),
			},
		},
	})
}

func TestAccIBMCisWAFGroup_Import(t *testing.T) {
	name := "ibm_cis_waf_group." + "test"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCisWAFGroupConfigBasic2("test", acc.CisDomainStatic),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "mode", "off"),
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

func testAccCheckCisWAFGroupConfigBasic1(id string, CisDomainStatic string) string {
	return testAccCheckIBMCisDomainDataSourceConfigBasic1() + fmt.Sprintf(`
	resource "ibm_cis_waf_group" "%[1]s" {
		cis_id     = data.ibm_cis.cis.id
		domain_id  = data.ibm_cis_domain.cis_domain.domain_id
		package_id = "c504870194831cd12c3fc0284f294abb"
		group_id   = "3d8fb0c18b5a6ba7682c80e94c7937b2"
		mode       = "on"
	  }
`, id)
}
func testAccCheckCisWAFGroupConfigBasic2(id string, CisDomainStatic string) string {
	return testAccCheckIBMCisDomainDataSourceConfigBasic1() + fmt.Sprintf(`
	resource "ibm_cis_waf_group" "%[1]s" {
		cis_id     = data.ibm_cis.cis.id
		domain_id  = data.ibm_cis_domain.cis_domain.domain_id
		package_id = "c504870194831cd12c3fc0284f294abb"
		group_id   = "3d8fb0c18b5a6ba7682c80e94c7937b2"
		mode       = "off"
	  }
`, id)
}

func testAccCheckCisWAFGroupConfigBasic3(id string, CisDomainStatic string) string {
	return testAccCheckIBMCisDomainDataSourceConfigBasic1() + fmt.Sprintf(`
	resource "ibm_cis_waf_group" "%[1]s" {
		cis_id     = data.ibm_cis.cis.id
		domain_id  = data.ibm_cis_domain.cis_domain.domain_id
		package_id = "c504870194831cd12c3fc0284f294abb"
		group_id   = "3d8fb0c18b5a6ba7682c80e94c7937b2"
		mode       = "on"
		check_mode = true
	  }
`, id)
}
