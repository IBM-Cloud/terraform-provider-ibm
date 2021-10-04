// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/IBM/platform-services-go-sdk/contextbasedrestrictionsv1"
)

func TestAccIBMCbrRuleBasic(t *testing.T) {
	var conf contextbasedrestrictionsv1.Rule

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
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
	transactionID := fmt.Sprintf("tf_transaction_id_%d", acctest.RandIntRange(10, 100))
	descriptionUpdate := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))
	transactionIDUpdate := fmt.Sprintf("tf_transaction_id_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMCbrRuleDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCbrRuleConfig(description, transactionID),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCbrRuleExists("ibm_cbr_rule.cbr_rule", conf),
					resource.TestCheckResourceAttr("ibm_cbr_rule.cbr_rule", "description", description),
					resource.TestCheckResourceAttr("ibm_cbr_rule.cbr_rule", "transaction_id", transactionID),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMCbrRuleConfig(descriptionUpdate, transactionIDUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_cbr_rule.cbr_rule", "description", descriptionUpdate),
					resource.TestCheckResourceAttr("ibm_cbr_rule.cbr_rule", "transaction_id", transactionIDUpdate),
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
		}
	`)
}

func testAccCheckIBMCbrRuleConfig(description string, transactionID string) string {
	return fmt.Sprintf(`

		resource "ibm_cbr_rule" "cbr_rule" {
			description = "%s"
			contexts {
				attributes {
					name = "name"
					value = "value"
				}
			}
			resources {
				attributes {
					name = "name"
					value = "value"
					operator = "operator"
				}
				tags {
					name = "name"
					value = "value"
					operator = "operator"
				}
			}
			transaction_id = "%s"
		}
	`, description, transactionID)
}

func testAccCheckIBMCbrRuleExists(n string, obj contextbasedrestrictionsv1.Rule) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		contextBasedRestrictionsClient, err := testAccProvider.Meta().(ClientSession).ContextBasedRestrictionsV1()
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
	contextBasedRestrictionsClient, err := testAccProvider.Meta().(ClientSession).ContextBasedRestrictionsV1()
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
