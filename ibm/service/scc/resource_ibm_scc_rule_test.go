// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package scc_test

import (
	"fmt"
	"os"
	"strings"
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
	descriptionUpdate := description
	instanceID, ok := os.LookupEnv("IBMCLOUD_SCC_INSTANCE_ID")
	if !ok {
		t.Logf("Missing the env var IBMCLOUD_SCC_INSTANCE_ID.")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheckSccInstanceID(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmSccRuleDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIbmSccRuleConfigBasic(instanceID, description),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmSccRuleExists("ibm_scc_rule.scc_rule_instance", conf),
					resource.TestCheckResourceAttr("ibm_scc_rule.scc_rule_instance", "description", description),
				),
			},
			{
				Config: testAccCheckIbmSccRuleConfigBasic(instanceID, descriptionUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_scc_rule.scc_rule_instance", "description", descriptionUpdate),
				),
			},
		},
	})
}

func TestAccIbmSccRuleAllArgs(t *testing.T) {
	var conf securityandcompliancecenterapiv3.Rule
	description := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))
	version := fmt.Sprintf("0.0.%d", acctest.RandIntRange(10, 100))
	descriptionUpdate := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))
	versionUpdate := fmt.Sprintf("0.0.%d", acctest.RandIntRange(2, 100))
	instanceID, ok := os.LookupEnv("IBMCLOUD_SCC_INSTANCE_ID")
	if !ok {
		t.Logf("Missing the env var IBMCLOUD_SCC_INSTANCE_ID.")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheckSccInstanceID(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmSccRuleDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIbmSccRuleConfig(instanceID, description, version),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmSccRuleExists("ibm_scc_rule.scc_rule_instance", conf),
					resource.TestCheckResourceAttr("ibm_scc_rule.scc_rule_instance", "description", description),
					resource.TestCheckResourceAttr("ibm_scc_rule.scc_rule_instance", "version", version),
				),
			},
			{
				Config: testAccCheckIbmSccRuleConfig(instanceID, descriptionUpdate, versionUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_scc_rule.scc_rule_instance", "description", descriptionUpdate),
					resource.TestCheckResourceAttr("ibm_scc_rule.scc_rule_instance", "version", versionUpdate),
				),
			},
			{
				ResourceName:      "ibm_scc_rule.scc_rule_instance",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIbmSccRuleConfigBasic(instanceID string, description string) string {
	return fmt.Sprintf(`
		resource "ibm_scc_rule" "scc_rule_instance" {
			instance_id = "%s"
			description = "%s"
			version = "0.0.1"
			target {
				service_name = "cloud-object-storage"
				resource_kind = "bucket"
				additional_target_attributes {
					name = "location"
					operator = "string_equals"
					value = "us-south"
				}
			}
			required_config {
				and {
					or {
						description = "description"
						property = "storage_class"
						operator = "string_equals"
						value = "smart"
					}
					or {
						description = "description"
						property = "storage_class"
						operator = "string_equals"
						value = "cold"
					}
				}
			}
		}
	`, instanceID, description)
}

func testAccCheckIbmSccRuleConfig(instanceID string, description string, version string) string {
	return fmt.Sprintf(`

		resource "ibm_scc_rule" "scc_rule_instance" {
			instance_id = "%s"
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
				service_name = "cloud-object-storage"
				resource_kind = "bucket"
				additional_target_attributes {
					name = "location"
					operator = "string_equals"
					value = "$${name}"
				}
			}
			required_config {
				and {
					or {
						description = "description"
						property = "storage_class"
						operator = "string_equals"
						value = "smart"
					}
					or {
						description = "description"
						property = "storage_class"
						operator = "string_equals"
						value = "cold"
					}
				}
			}
			labels = ["FIXME"]
		}
	`, instanceID, description, version)
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
		id := strings.Split(rs.Primary.ID, "/")
		getRuleOptions.SetInstanceID(id[0])
		getRuleOptions.SetRuleID(id[1])

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

		id := strings.Split(rs.Primary.ID, "/")
		getRuleOptions.SetInstanceID(id[0])
		getRuleOptions.SetRuleID(id[1])

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
