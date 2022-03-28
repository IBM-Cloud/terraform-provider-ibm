// Copyright IBM Corp. 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package scc_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM/scc-go-sdk/v3/configurationgovernancev1"
)

func TestAccIBMSccRuleBasic(t *testing.T) {
	var conf configurationgovernancev1.Rule

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMSccRuleDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMSccRuleConfigBasic(),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMSccRuleExists("ibm_scc_rule.scc_rule", conf),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_scc_rule.scc_rule",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIBMSccRuleConfigBasic() string {
	// Check if the user has a SCC_GOVERANCE_ACCOUNT_ID
	account_id := os.Getenv("SCC_GOVERNANCE_ACCOUNT_ID")
	return fmt.Sprintf(`

	resource "ibm_scc_rule" "scc_rule" {
		account_id = "%s"
		name = "scc_tf_sample_rule"
		description = "description"
		target {               
			service_name = "cloud-object-storage"
			resource_kind = "bucket"
			additional_target_attributes {
				name = "location"
				value = "us-south"                     
				operator = "string_equals"
			}                         
		}
		labels = ["test1", "test2"]  
		required_config {
			description = "test config"
			or {                                                                                            
				property = "location"                                                                       
				operator = "string_equals"        
				value = "us-west"
			}
			or {                                   
				property = "location"
				operator = "string_equals"
				value = "us-east"
			}
		} 
		enforcement_actions {
			action = "disallow"
		}                                         
	}
	`, account_id)
}

func testAccCheckIBMSccRuleExists(n string, obj configurationgovernancev1.Rule) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		configurationGovernanceClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).ConfigurationGovernanceV1()
		if err != nil {
			return err
		}

		getRuleOptions := &configurationgovernancev1.GetRuleOptions{}

		getRuleOptions.SetRuleID(rs.Primary.ID)

		rule, _, err := configurationGovernanceClient.GetRule(getRuleOptions)
		if err != nil {
			return err
		}

		obj = *rule
		return nil
	}
}

func testAccCheckIBMSccRuleDestroy(s *terraform.State) error {
	configurationGovernanceClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).ConfigurationGovernanceV1()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_scc_rule" {
			continue
		}

		getRuleOptions := &configurationgovernancev1.GetRuleOptions{}

		getRuleOptions.SetRuleID(rs.Primary.ID)

		// Try to find the key
		_, response, err := configurationGovernanceClient.GetRule(getRuleOptions)

		if err == nil {
			return fmt.Errorf("scc_rule still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for scc_rule (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
