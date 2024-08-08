// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package scc_test

import (
	"fmt"
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
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheckScc(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmSccRuleDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIbmSccRuleConfigBasic(acc.SccInstanceID, description),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmSccRuleExists("ibm_scc_rule.scc_rule_instance", conf),
					resource.TestCheckResourceAttr("ibm_scc_rule.scc_rule_instance", "description", description),
				),
			},
			{
				Config: testAccCheckIbmSccRuleConfigBasic(acc.SccInstanceID, descriptionUpdate),
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

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheckScc(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmSccRuleDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIbmSccRuleConfig(acc.SccInstanceID, description, version),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmSccRuleExists("ibm_scc_rule.scc_rule_instance", conf),
					resource.TestCheckResourceAttr("ibm_scc_rule.scc_rule_instance", "description", description),
					resource.TestCheckResourceAttr("ibm_scc_rule.scc_rule_instance", "version", version),
					resource.TestCheckResourceAttr("ibm_scc_rule.scc_rule_instance", "required_config.0.and.0.or.2.value", "[\"0.0.0.0/0\"]"),
				),
			},
			{
				Config: testAccCheckIbmSccRuleConfig(acc.SccInstanceID, descriptionUpdate, versionUpdate),
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

func TestAccIbmSccRuleSubRuleBasic(t *testing.T) {
	var conf securityandcompliancecenterapiv3.Rule
	description := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))
	version := fmt.Sprintf("0.0.%d", acctest.RandIntRange(10, 100))
	descriptionUpdate := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))
	versionUpdate := fmt.Sprintf("0.0.%d", acctest.RandIntRange(2, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheckScc(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmSccRuleDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIbmSccRuleConfigSubRule(acc.SccInstanceID, description, version),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmSccRuleExists("ibm_scc_rule.scc_rule_instance", conf),
					resource.TestCheckResourceAttr("ibm_scc_rule.scc_rule_instance", "description", description),
					resource.TestCheckResourceAttr("ibm_scc_rule.scc_rule_instance", "version", version),
					resource.TestCheckResourceAttr("ibm_scc_rule.scc_rule_instance", "required_config.0.or.1.any_if.0.required_config.0.value", "[\"us-south\",\"us-east\"]"),
				),
			},
			{
				Config: testAccCheckIbmSccRuleConfigSubRule(acc.SccInstanceID, descriptionUpdate, versionUpdate),
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

func TestAccIbmSccRuleWithNumberRC(t *testing.T) {
	var conf securityandcompliancecenterapiv3.Rule
	description := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))
	version := fmt.Sprintf("0.0.%d", acctest.RandIntRange(1, 10))
	descriptionUpdate := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))
	versionUpdate := fmt.Sprintf("0.0.%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheckScc(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmSccRuleDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIbmSccRuleConfigNumVal(acc.SccInstanceID, description, version),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmSccRuleExists("ibm_scc_rule.scc_rule_instance", conf),
					resource.TestCheckResourceAttr("ibm_scc_rule.scc_rule_instance", "description", description),
					resource.TestCheckResourceAttr("ibm_scc_rule.scc_rule_instance", "version", version),
					resource.TestCheckResourceAttr("ibm_scc_rule.scc_rule_instance", "required_config.0.and.0.value", "0"),
				),
			},
			{
				Config: testAccCheckIbmSccRuleConfigNumVal(acc.SccInstanceID, descriptionUpdate, versionUpdate),
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
						description = "description 1"
						property = "storage_class"
						operator = "string_equals"
						value = "smart"
					}
					or {
						description = "description 2"
						property = "storage_class"
						operator = "string_equals"
						value = "cold"
					}
					or {
						description = "description 3"
						property = "firewall.allowed_ip"
						operator = "ips_equals"
						value = jsonencode(["0.0.0.0/0"])
					}
				}
			}
			labels = ["FIXME"]
		}
	`, instanceID, description, version)
}

func testAccCheckIbmSccRuleConfigNumVal(instanceID string, description string, version string) string {
	return fmt.Sprintf(`
		resource "ibm_scc_rule" "scc_rule_instance" {
			instance_id = "%s"
			description = "%s"
			version = "%s"
			target {
				service_name = "is.load-balancer"
				resource_kind = "instance"
				additional_target_attributes {
					name = "profile_family"
					operator = "string_equals"
					value = "application"
				}
			}
			required_config {
				and {
					property = "app_lb_pools_with_multiple_members_count"
                	operator = "num_not_equals"
                	value = "0"
				}
				and {
					property = "app_lb_pools_without_multiple_members_count"
                	operator = "num_not_equals"
                	value = "0"
				}
			}
			labels = ["FIXME"]
		}
	`, instanceID, description, version)
}

func testAccCheckIbmSccRuleConfigSubRule(instanceID string, description string, version string) string {
	return fmt.Sprintf(`

		resource "ibm_scc_rule" "scc_rule_instance" {
			instance_id = "%s"
			description = "%s"
			version = "%s"
			target {
				service_name = "atracker"
				resource_kind = "target"
				reference_name = "this-target"
				additional_target_attributes {
					name = "type"
					operator = "string_equals"
					value = "cloud_object_storage"
				}
			}
			required_config {
				or {
					property = "route_attached"
					operator = "is_false"
                    
				}
				or {
					any_if {
						target {
							service_name = "cloud-object-storage"
							resource_kind = "bucket"
							additional_target_attributes {
								name = "location"
								operator = "strings_in_list"
								value = "$${this-target}.bucket_name"
							}
						}
						required_config {
							property = "location"
							operator = "strings_in_list"
							value = jsonencode(["us-south","us-east"])
						}
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
