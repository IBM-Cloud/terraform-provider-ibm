// Copyright IBM Corp. 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package contextbasedrestrictions_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM/platform-services-go-sdk/contextbasedrestrictionsv1"
)

func TestAccIBMCbrRuleBasic(t *testing.T) {
	var conf contextbasedrestrictionsv1.Rule

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCbrRuleDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCbrRuleConfigBasic(),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCbrRuleExists("ibm_cbr_rule.cbr_rule", conf),
				),
			},
		},
	})
}

func TestAccIBMCbrRuleAllArgs(t *testing.T) {
	var conf contextbasedrestrictionsv1.Rule
	description := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))
	enforcementMode := "enabled"
	descriptionUpdate := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))
	enforcementModeUpdate := "report"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCbrRuleDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCbrRuleConfig(description, enforcementMode),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCbrRuleExists("ibm_cbr_rule.cbr_rule", conf),
					resource.TestCheckResourceAttr("ibm_cbr_rule.cbr_rule", "description", description),
					resource.TestCheckResourceAttr("ibm_cbr_rule.cbr_rule", "enforcement_mode", enforcementMode),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMCbrRuleConfig(descriptionUpdate, enforcementModeUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_cbr_rule.cbr_rule", "description", descriptionUpdate),
					resource.TestCheckResourceAttr("ibm_cbr_rule.cbr_rule", "enforcement_mode", enforcementModeUpdate),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_cbr_rule.cbr_rule",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIBMCbrRuleConfigBasic() string {
	return fmt.Sprintf(`

		resource "ibm_cbr_rule" "cbr_rule" {
  			description = "test rule config basic"
  			contexts {
    			attributes {
      				name = "networkZoneId"
      				value = "559052eb8f43302824e7ae490c0281eb"
    			}
  			}
			resources {
    			attributes {
      				name = "accountId"
      				value = "12ab34cd56ef78ab90cd12ef34ab56cd"
    			}
    			attributes {
      				name = "serviceName"
      				value = "user-management"
    			}
    			tags {
      				name     = "tag_name"
      				value    = "tag_value"
    			}
  			}
			enforcement_mode = "disabled"
		}
	`)
}

func testAccCheckIBMCbrRuleConfig(description string, enforcementMode string) string {
	return fmt.Sprintf(`

		resource "ibm_cbr_rule" "cbr_rule" {
			description = "%s"
			contexts {
    			attributes {
      				name = "networkZoneId"
      				value = "559052eb8f43302824e7ae490c0281eb"
    			}
			}
			resources {
    			attributes {
      				name = "accountId"
      				value = "12ab34cd56ef78ab90cd12ef34ab56cd"
    			}
    			attributes {
      				name = "serviceName"
      				value = "containers-kubernetes"
    			}
				tags {
					name = "name"
					value = "value"
					operator = "stringEquals"
				}
			}
			operations {
				api_types {
					api_type_id = "crn:v1:bluemix:public:containers-kubernetes::::api-type:management"
				}
			}
			enforcement_mode = "%s"
		}
	`, description, enforcementMode)
}

func testAccCheckIBMCbrRuleExists(n string, obj contextbasedrestrictionsv1.Rule) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		contextBasedRestrictionsClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).ContextBasedRestrictionsV1()
		if err != nil {
			return err
		}

		getRuleOptions := &contextbasedrestrictionsv1.GetRuleOptions{}

		getRuleOptions.SetRuleID(rs.Primary.ID)

		rule, _, err := contextBasedRestrictionsClient.GetRule(getRuleOptions)
		if err != nil {
			return err
		}

		obj = *rule
		return nil
	}
}

func testAccCheckIBMCbrRuleDestroy(s *terraform.State) error {
	contextBasedRestrictionsClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).ContextBasedRestrictionsV1()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_cbr_rule" {
			continue
		}

		getRuleOptions := &contextbasedrestrictionsv1.GetRuleOptions{}

		getRuleOptions.SetRuleID(rs.Primary.ID)

		// Try to find the key
		_, response, err := contextBasedRestrictionsClient.GetRule(getRuleOptions)

		if err == nil {
			return fmt.Errorf("cbr_rule still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for cbr_rule (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
