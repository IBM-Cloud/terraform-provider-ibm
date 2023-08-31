// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package scc_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM/scc-go-sdk/v5/securityandcompliancecenterapiv3"
)

func TestAccIbmSccRuleBasic(t *testing.T) {
	var conf securityandcompliancecenterapiv3.Rule
	description := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))
	descriptionUpdate := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmSccRuleDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIbmSccRuleConfigBasic(description),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmSccRuleExists("ibm_scc_rule.scc_rule", conf),
					resource.TestCheckResourceAttr("ibm_scc_rule.scc_rule", "description", description),
				),
			},
			{
				Config: testAccCheckIbmSccRuleConfigBasic(descriptionUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_scc_rule.scc_rule", "description", descriptionUpdate),
				),
			},
		},
	})
}

func TestAccIbmSccRuleAllArgs(t *testing.T) {
	var conf securityandcompliancecenterapiv3.Rule
	description := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))
	version := fmt.Sprintf("tf_version_%d", acctest.RandIntRange(10, 100))
	descriptionUpdate := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))
	versionUpdate := fmt.Sprintf("tf_version_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmSccRuleDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIbmSccRuleConfig(description, version),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmSccRuleExists("ibm_scc_rule.scc_rule", conf),
					resource.TestCheckResourceAttr("ibm_scc_rule.scc_rule", "description", description),
					resource.TestCheckResourceAttr("ibm_scc_rule.scc_rule", "version", version),
				),
			},
			{
				Config: testAccCheckIbmSccRuleConfig(descriptionUpdate, versionUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_scc_rule.scc_rule", "description", descriptionUpdate),
					resource.TestCheckResourceAttr("ibm_scc_rule.scc_rule", "version", versionUpdate),
				),
			},
			{
				ResourceName:      "ibm_scc_rule.scc_rule",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIbmSccRuleConfigBasic(description string) string {
	return fmt.Sprintf(`
		resource "ibm_scc_rule" "scc_rule_instance" {
			description = "%s"
			target {
				service_name = "service_name"
				service_display_name = "service_display_name"
				resource_kind = "resource_kind"
				additional_target_attributes {
					name = "name"
					operator = "string_equals"
					value = "value"
				}
			}
			required_config {
				description = "description"
				and {
					description = "description"
					or {
						description = "description"
						property = "property"
						operator = "string_equals"
						value = "anything as a string"
					}
				}
			}
		}
	`, description)
}

func testAccCheckIbmSccRuleConfig(description string, version string) string {
	return fmt.Sprintf(`

		resource "ibm_scc_rule" "scc_rule_instance" {
			description = "%s"
			version = "%s"
			import {
				parameters {
					name = "name"
					display_name = "display_name"
					description = "description"
					type = "string"
				}
			}
			target {
				service_name = "service_name"
				service_display_name = "service_display_name"
				resource_kind = "resource_kind"
				additional_target_attributes {
					name = "name"
					operator = "string_equals"
					value = "value"
				}
			}
			required_config {
				description = "description"
				and {
					description = "description"
					or {
						description = "description"
						property = "property"
						operator = "string_equals"
						value = "anything as a string"
					}
				}
			}
			labels = "FIXME"
		}
	`, description, version)
}

func testAccCheckIbmSccRuleExists(n string, obj securityandcompliancecenterapiv3.Rule) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		configManagerClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).SecurityAndComplianceCenterV3()
		if err != nil {
			return err
		}

		getRuleOptions := &securityandcompliancecenterapiv3.GetRuleOptions{}

		getRuleOptions.SetRuleID(rs.Primary.ID)

		rule, _, err := configManagerClient.GetRule(getRuleOptions)
		if err != nil {
			return err
		}

		obj = *rule
		return nil
	}
}

func testAccCheckIbmSccRuleDestroy(s *terraform.State) error {
	configManagerClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).SecurityAndComplianceCenterV3()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_scc_rule" {
			continue
		}

		getRuleOptions := &securityandcompliancecenterapiv3.GetRuleOptions{}

		getRuleOptions.SetRuleID(rs.Primary.ID)

		// Try to find the key
		_, response, err := configManagerClient.GetRule(getRuleOptions)

		if err == nil {
			return fmt.Errorf("scc_rule still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for scc_rule (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
