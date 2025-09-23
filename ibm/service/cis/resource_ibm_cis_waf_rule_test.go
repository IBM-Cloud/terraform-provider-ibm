// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package cis_test

import (
	"testing"

	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMCisWAFRule_Basic1(t *testing.T) {
	name := "ibm_cis_waf_rule." + "test"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckCis(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCisWAFRuleConfigBasic1(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "mode", "on"),
				),
			},
			{
				Config: testAccCheckCisWAFRuleConfigBasic2(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "mode", "off"),
				),
			},
		},
	})
}

func TestAccIBMCisWAFRule_Basic2(t *testing.T) {
	name := "ibm_cis_waf_rule." + "test"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckCis(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCisWAFRuleConfigBasic3(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "mode", "simulate"),
				),
			},
			{
				Config: testAccCheckCisWAFRuleConfigBasic4(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "mode", "default"),
				),
			},
		},
	})
}

func TestAccIBMCisWAFRule_Import(t *testing.T) {
	name := "ibm_cis_waf_rule." + "test"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCisWAFRuleConfigBasic1(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "mode", "on"),
				),
			},
			{
				Config: testAccCheckCisWAFRuleConfigBasic2(),
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

func testAccCheckCisWAFRuleConfigBasic1() string {
	return testAccCheckIBMCisDomainDataSourceConfigBasic1() + `
	resource "ibm_cis_waf_rule" "test" {
		cis_id     = data.ibm_cis.cis.id
		domain_id  = data.ibm_cis_domain.cis_domain.id
		package_id = "c504870194831cd12c3fc0284f294abb"
		rule_id    = "100000356"
		mode       = "on"
	  }`
}

func testAccCheckCisWAFRuleConfigBasic2() string {
	return testAccCheckIBMCisDomainDataSourceConfigBasic1() + `
	resource "ibm_cis_waf_rule" "test" {
		cis_id     = data.ibm_cis.cis.id
		domain_id  = data.ibm_cis_domain.cis_domain.id
		package_id = "c504870194831cd12c3fc0284f294abb"
		rule_id    = "100000356"
		mode       = "off"
	  }`
}

func testAccCheckCisWAFRuleConfigBasic3() string {
	return testAccCheckIBMCisDomainDataSourceConfigBasic1() + `
	resource "ibm_cis_waf_rule" "test" {
		cis_id     = data.ibm_cis.cis.id
		domain_id  = data.ibm_cis_domain.cis_domain.id
		package_id = "1e334934fd7ae32ad705667f8c1057aa"
		rule_id    = "100000"
		mode       = "simulate"
	  }`
}

func testAccCheckCisWAFRuleConfigBasic4() string {
	return testAccCheckIBMCisDomainDataSourceConfigBasic1() + `
	resource "ibm_cis_waf_rule" "test" {
		cis_id     = data.ibm_cis.cis.id
		domain_id  = data.ibm_cis_domain.cis_domain.id
		package_id = "1e334934fd7ae32ad705667f8c1057aa"
		rule_id    = "100000"
		mode       = "default"
	  }`
}
