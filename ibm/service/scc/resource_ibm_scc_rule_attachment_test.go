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
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/scc-go-sdk/v3/configurationgovernancev1"
)

func TestAccIBMSccRuleAttachmentBasic(t *testing.T) {
	var conf configurationgovernancev1.RuleAttachment

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMSccRuleAttachmentDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMSccRuleAttachmentConfigBasic(),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMSccRuleAttachmentExists("ibm_scc_rule_attachment.scc_rule_attachment", conf),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_scc_rule_attachment.scc_rule_attachment",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIBMSccRuleAttachmentConfigBasic() string {
	account_id := os.Getenv("SCC_GOVERNANCE_ACCOUNT_ID")
	resource_group_id := os.Getenv("IBM_SCC_RESOURCE_GROUP")
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
				value = "us-south"
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

	resource "ibm_scc_rule_attachment" "scc_rule_attachment" {
		rule_id = ibm_scc_rule.scc_rule.id
		account_id = "%s"
		included_scope {
			note = "note"
			scope_id = "%s"
			scope_type = "account"
		}
		excluded_scopes {
			note = "note"
			scope_id = "%s"
			scope_type = "account.resource_group"
		}
		depends_on = [
			ibm_scc_rule.scc_rule
		]
	}
	`, account_id, account_id, account_id, resource_group_id)
}

func testAccCheckIBMSccRuleAttachmentExists(n string, obj configurationgovernancev1.RuleAttachment) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		configurationGovernanceClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).ConfigurationGovernanceV1()
		if err != nil {
			return err
		}

		getRuleAttachmentOptions := &configurationgovernancev1.GetRuleAttachmentOptions{}
		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		ruleID := parts[0]
		getRuleAttachmentOptions.SetRuleID(ruleID)
		getRuleAttachmentOptions.SetAttachmentID(parts[1])

		ruleAttachment, _, err := configurationGovernanceClient.GetRuleAttachment(getRuleAttachmentOptions)
		if err != nil {
			return err
		}

		if *ruleAttachment.RuleID != ruleID {
			return fmt.Errorf(
				"ibm_scc_rule_attachment.scc_rule_attachment: Attribute 'rule_id' expected %#v, got %#v",
				ruleID,
				ruleAttachment.RuleID,
			)
		}

		obj = *ruleAttachment
		return nil
	}
}

func testAccCheckIBMSccRuleAttachmentDestroy(s *terraform.State) error {
	configurationGovernanceClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).ConfigurationGovernanceV1()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_scc_rule_attachment" {
			continue
		}

		getRuleAttachmentOptions := &configurationgovernancev1.GetRuleAttachmentOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getRuleAttachmentOptions.SetRuleID(parts[0])
		getRuleAttachmentOptions.SetAttachmentID(parts[1])

		// Try to find the key
		_, response, err := configurationGovernanceClient.GetRuleAttachment(getRuleAttachmentOptions)

		if err == nil {
			return fmt.Errorf("scc_rule_attachment still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for scc_rule_attachment (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
