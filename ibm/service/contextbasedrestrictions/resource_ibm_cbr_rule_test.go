// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package contextbasedrestrictions_test

import (
	"fmt"
	"os"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/platform-services-go-sdk/contextbasedrestrictionsv1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccIBMCbrRuleBasic(t *testing.T) {
	var conf contextbasedrestrictionsv1.Rule

	accountID, _ := getTestAccountAndZoneID()
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheckCbr(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCbrRuleDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCbrRuleConfigBasic(accountID),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCbrRuleExists("ibm_cbr_rule.cbr_rule_instance", conf),
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

	accountID, _ := getTestAccountAndZoneID()
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheckCbr(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCbrRuleDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCbrRuleConfig(description, enforcementMode, accountID),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCbrRuleExists("ibm_cbr_rule.cbr_rule_instance", conf),
					resource.TestCheckResourceAttr("ibm_cbr_rule.cbr_rule_instance", "description", description),
					resource.TestCheckResourceAttr("ibm_cbr_rule.cbr_rule_instance", "enforcement_mode", enforcementMode),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMCbrRuleConfigUpdate(descriptionUpdate, enforcementModeUpdate, accountID),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_cbr_rule.cbr_rule_instance", "description", descriptionUpdate),
					resource.TestCheckResourceAttr("ibm_cbr_rule.cbr_rule_instance", "enforcement_mode", enforcementModeUpdate),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_cbr_rule.cbr_rule_instance",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIBMCbrRuleConfigBasic(accountID string) string {
	return fmt.Sprintf(`
		resource "ibm_cbr_zone" "cbr_zone" {
			name = "Test Zone Data Source Config Basic"
			description = "Test Zone Data Source Config Basic"
			account_id = "%s"
			addresses {
				type = "ipRange"
				value = "169.23.22.0-169.23.22.255"
			}
		}

		resource "ibm_cbr_rule" "cbr_rule_instance" {
  			description = "test rule config basic"
  			contexts {
    			attributes {
      				name = "networkZoneId"
      				value = ibm_cbr_zone.cbr_zone.id
    			}
  			}
			resources {
    			attributes {
      				name = "accountId"
      				value = "%s"
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
	`, accountID, accountID)
}

func testAccCheckIBMCbrRuleConfig(description string, enforcementMode string, accountID string) string {
	return fmt.Sprintf(`
		resource "ibm_cbr_zone" "cbr_zone" {
			name = "Test Zone Data Source Config Basic"
			description = "Test Zone Data Source Config Basic"
			account_id = "%s"
			addresses {
				type = "ipRange"
				value = "169.23.22.0-169.23.22.255"
			}
		}

		resource "ibm_cbr_rule" "cbr_rule_instance" {
			description = "%s"
			contexts {
    			attributes {
      				name = "networkZoneId"
      				value = ibm_cbr_zone.cbr_zone.id
    			}
			}
			resources {
    			attributes {
      				name = "accountId"
      				value = "%s"
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
	`, accountID, description, accountID, enforcementMode)
}

func testAccCheckIBMCbrRuleConfigUpdate(description string, enforcementMode string, accountID string) string {
	os.Setenv("IBMCLOUD_CONTEXT_BASED_RESTRICTIONS_ENDPOINT", "https://testing-2-eu-gb.network-policy.test.cloud.ibm.com")
	return fmt.Sprintf(`
		resource "ibm_cbr_zone" "cbr_zone" {
			name = "Test Zone Data Source Config Basic"
			description = "Test Zone Data Source Config Basic"
			account_id = "%s"
			addresses {
				type = "ipRange"
				value = "169.23.22.0-169.23.22.255"
			}
		}

		resource "ibm_cbr_rule" "cbr_rule_instance" {
			description = "%s"
			contexts {
				attributes {
					name = "networkZoneId"
					value = ibm_cbr_zone.cbr_zone.id
				}
			}
			resources {
				attributes {
					name = "serviceName"
					value = "containers-kubernetes"
				}
				attributes {
					name = "accountId"
					value = "%s"
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
	`, accountID, description, accountID, enforcementMode)
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

func TestAccIBMCbrRuleImportThenUpdate(t *testing.T) {
	var zoneID, ruleID, importedEtag string
	description := fmt.Sprintf("tf_import_description_%d", acctest.RandIntRange(10, 100))
	descriptionUpdate := fmt.Sprintf("tf_update_description_%d", acctest.RandIntRange(10, 100))
	enforcementMode := "disabled"
	enforcementModeUpdate := "report"

	accountID, _ := getTestAccountAndZoneID()
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheckCbr(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCbrRuleDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				// Only the zone is managed by Terraform here. The rule is created
				// outside of Terraform in the next step so that it can be imported
				// into the persisted state.
				Config: testAccCheckIBMCbrRuleConfigZoneOnly(accountID),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCbrRuleCaptureID("ibm_cbr_zone.cbr_zone", &zoneID),
				),
			},
			resource.TestStep{
				PreConfig: func() {
					id, err := testAccCreateIBMCbrRuleOutOfBand(zoneID, accountID, description, enforcementMode)
					if err != nil {
						t.Fatalf("Error creating cbr_rule outside of Terraform: %s", err)
					}
					ruleID = id
				},
				Config:       testAccCheckIBMCbrRuleConfig(description, enforcementMode, accountID),
				ResourceName: "ibm_cbr_rule.cbr_rule_instance",
				ImportState:  true,
				ImportStateIdFunc: func(*terraform.State) (string, error) {
					return ruleID, nil
				},
				// Keep the imported state so that the next step updates the
				// imported rule, using the etag that was set during import.
				ImportStatePersist: true,
				ImportStateCheck: func(states []*terraform.InstanceState) error {
					if err := testAccCheckIBMCbrRuleImportedState(states, ruleID, description, enforcementMode); err != nil {
						return err
					}
					importedEtag = states[0].Attributes["etag"]
					return nil
				},
			},
			resource.TestStep{
				Config: testAccCheckIBMCbrRuleConfigUpdate(descriptionUpdate, enforcementModeUpdate, accountID),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrPtr("ibm_cbr_rule.cbr_rule_instance", "id", &ruleID),
					resource.TestCheckResourceAttr("ibm_cbr_rule.cbr_rule_instance", "description", descriptionUpdate),
					resource.TestCheckResourceAttr("ibm_cbr_rule.cbr_rule_instance", "enforcement_mode", enforcementModeUpdate),
					resource.TestCheckResourceAttrWith("ibm_cbr_rule.cbr_rule_instance", "etag", func(value string) error {
						if value == "" {
							return fmt.Errorf("etag is empty after update")
						}
						if value == importedEtag {
							return fmt.Errorf("etag did not change after update of the imported rule: %s", value)
						}
						return nil
					}),
				),
			},
		},
	})
}

func TestAccIBMCbrRuleDisappears(t *testing.T) {
	var conf contextbasedrestrictionsv1.Rule

	accountID, _ := getTestAccountAndZoneID()
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheckCbr(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCbrRuleDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCbrRuleConfigBasic(accountID),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCbrRuleExists("ibm_cbr_rule.cbr_rule_instance", conf),
					testAccCheckIBMCbrRuleDisappears("ibm_cbr_rule.cbr_rule_instance"),
				),
				// The rule is deleted outside of Terraform, so the read after apply
				// must drop it from state and the next plan must recreate it.
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func testAccCheckIBMCbrRuleConfigZoneOnly(accountID string) string {
	return fmt.Sprintf(`
		resource "ibm_cbr_zone" "cbr_zone" {
			name = "Test Zone Data Source Config Basic"
			description = "Test Zone Data Source Config Basic"
			account_id = "%s"
			addresses {
				type = "ipRange"
				value = "169.23.22.0-169.23.22.255"
			}
		}
	`, accountID)
}

func testAccCheckIBMCbrRuleCaptureID(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("No ID is set for %s", n)
		}
		*id = rs.Primary.ID
		return nil
	}
}

// testAccCreateIBMCbrRuleOutOfBand creates a rule that matches
// testAccCheckIBMCbrRuleConfig, without going through Terraform.
func testAccCreateIBMCbrRuleOutOfBand(zoneID, accountID, description, enforcementMode string) (string, error) {
	contextBasedRestrictionsClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).ContextBasedRestrictionsV1()
	if err != nil {
		return "", err
	}

	createRuleOptions := &contextbasedrestrictionsv1.CreateRuleOptions{
		Description: core.StringPtr(description),
		Contexts: []contextbasedrestrictionsv1.RuleContext{
			{
				Attributes: []contextbasedrestrictionsv1.RuleContextAttribute{
					{Name: core.StringPtr("networkZoneId"), Value: core.StringPtr(zoneID)},
				},
			},
		},
		Resources: []contextbasedrestrictionsv1.Resource{
			{
				Attributes: []contextbasedrestrictionsv1.ResourceAttribute{
					{Name: core.StringPtr("accountId"), Value: core.StringPtr(accountID)},
					{Name: core.StringPtr("serviceName"), Value: core.StringPtr("containers-kubernetes")},
				},
				Tags: []contextbasedrestrictionsv1.ResourceTagAttribute{
					{Name: core.StringPtr("name"), Value: core.StringPtr("value"), Operator: core.StringPtr("stringEquals")},
				},
			},
		},
		Operations: &contextbasedrestrictionsv1.NewRuleOperations{
			APITypes: []contextbasedrestrictionsv1.NewRuleOperationsAPITypesItem{
				{APITypeID: core.StringPtr("crn:v1:bluemix:public:containers-kubernetes::::api-type:management")},
			},
		},
		EnforcementMode: core.StringPtr(enforcementMode),
	}

	rule, _, err := contextBasedRestrictionsClient.CreateRule(createRuleOptions)
	if err != nil {
		return "", err
	}
	if rule == nil || rule.ID == nil {
		return "", fmt.Errorf("CreateRule returned no rule ID")
	}
	return *rule.ID, nil
}

func testAccCheckIBMCbrRuleImportedState(states []*terraform.InstanceState, ruleID, description, enforcementMode string) error {
	if len(states) != 1 {
		return fmt.Errorf("expected 1 imported state, got %d", len(states))
	}
	state := states[0]
	if state.ID != ruleID {
		return fmt.Errorf("imported ibm_cbr_rule ID is %q, expected %q", state.ID, ruleID)
	}
	expected := map[string]string{
		"description":      description,
		"enforcement_mode": enforcementMode,
	}
	for attr, want := range expected {
		if got := state.Attributes[attr]; got != want {
			return fmt.Errorf("imported ibm_cbr_rule %s is %q, expected %q", attr, got, want)
		}
	}
	if state.Attributes["etag"] == "" {
		return fmt.Errorf("imported ibm_cbr_rule etag is empty")
	}
	return nil
}

func testAccCheckIBMCbrRuleDisappears(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		contextBasedRestrictionsClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).ContextBasedRestrictionsV1()
		if err != nil {
			return err
		}

		deleteRuleOptions := &contextbasedrestrictionsv1.DeleteRuleOptions{}
		deleteRuleOptions.SetRuleID(rs.Primary.ID)

		_, err = contextBasedRestrictionsClient.DeleteRule(deleteRuleOptions)
		return err
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
