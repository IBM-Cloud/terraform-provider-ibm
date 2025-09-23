// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package cis_test

import (
	"fmt"
	"testing"

	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/conns"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/flex"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccIBMCisPageRule_Basic(t *testing.T) {
	var record string
	name := "ibm_cis_page_rule.page_rule"
	resource.Test(t, resource.TestCase{
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCisPageRuleDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCisPageRuleConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMCisPageRuleExists(name, &record),
					resource.TestCheckResourceAttr(
						name, "actions.#", "1"),
					resource.TestCheckResourceAttr(
						name, "targets.#", "1"),
				),
			},
			{
				Config: testAccCheckIBMCisPageRuleConfigUpdate(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMCisPageRuleExists(name, &record),
					resource.TestCheckResourceAttr(
						name, "actions.#", "2"),
					resource.TestCheckResourceAttr(
						name, "targets.#", "1"),
				),
			},
		},
	})
}

func TestAccIBMCisPageRule_Import(t *testing.T) {
	name := "ibm_cis_page_rule.page_rule"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCisPageRuleConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						name, "actions.#", "1"),
					resource.TestCheckResourceAttr(
						name, "targets.#", "1"),
				),
			},
			{
				Config: testAccCheckIBMCisPageRuleConfigUpdate(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						name, "actions.#", "2"),
					resource.TestCheckResourceAttr(
						name, "targets.#", "1"),
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
func testAccCheckIBMCisPageRuleDestroy(s *terraform.State) error {
	cisClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).CisPageRuleClientSession()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_cis_page_rule" {
			continue
		}
		ruleID, zoneID, cisID, _ := flex.ConvertTfToCisThreeVar(rs.Primary.ID)
		cisClient.Crn = core.StringPtr(cisID)
		cisClient.ZoneID = core.StringPtr(zoneID)
		opt := cisClient.NewGetPageRuleOptions(ruleID)
		_, _, err := cisClient.GetPageRule(opt)
		if err == nil {
			return fmt.Errorf("Rule still exists")
		}
	}
	return nil
}

func testAccCheckIBMCisPageRuleExists(n string, tfRecordID *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("[ERROR] Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("[ERROR] No Record ID is set")
		}

		// tfRecord := *tfRecordID
		cisClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).CisPageRuleClientSession()
		if err != nil {
			return err
		}
		ruleID, zoneID, cisID, _ := flex.ConvertTfToCisThreeVar(rs.Primary.ID)
		cisClient.Crn = core.StringPtr(cisID)
		cisClient.ZoneID = core.StringPtr(zoneID)
		opt := cisClient.NewGetPageRuleOptions(ruleID)
		foundRecordPtr, _, err := cisClient.GetPageRule(opt)
		if err != nil {
			return err
		}

		foundRecord := foundRecordPtr.Result
		if *foundRecord.ID != ruleID {
			return fmt.Errorf("Record not found")
		}

		tfRecord := flex.ConvertCisToTfThreeVar(*foundRecord.ID, zoneID, cisID)
		*tfRecordID = tfRecord
		return nil
	}
}

func testAccCheckIBMCisPageRuleConfigBasic() string {
	forwardingURL := fmt.Sprintf("https://%s/*", acc.CisDomainStatic)
	url := fmt.Sprintf("%s/", acc.CisDomainStatic)

	return testAccCheckIBMCisDomainDataSourceConfigBasic1() + fmt.Sprintf(`
	resource "ibm_cis_page_rule" "page_rule" {
		cis_id = data.ibm_cis.cis.id
		domain_id = data.ibm_cis_domain.cis_domain.id
		targets {
		  target = "url"
		  constraint {
			operator = "matches"
			value    = "%[1]s"
		  }
		}
		status   = "active"
		priority = 1
		actions {
			id          = "forwarding_url"
			url         = "%[2]s"
			status_code = 302
		}
	}
	`, url, forwardingURL)
}

func testAccCheckIBMCisPageRuleConfigUpdate() string {
	url := fmt.Sprintf("%s/", acc.CisDomainStatic)
	return testAccCheckIBMCisDomainDataSourceConfigBasic1() + fmt.Sprintf(`
	resource "ibm_cis_page_rule" "page_rule" {
		cis_id = data.ibm_cis.cis.id
		domain_id = data.ibm_cis_domain.cis_domain.id
		targets {
		  target = "url"
		  constraint {
			operator = "matches"
			value    = "%[1]s"
		  }
		}
		status   = "active"
		priority = 1
		actions {
			id = "disable_security"
		}
		actions {
		  id    = "browser_check"
		  value = "on"
		}
	}
	`, url)
}
