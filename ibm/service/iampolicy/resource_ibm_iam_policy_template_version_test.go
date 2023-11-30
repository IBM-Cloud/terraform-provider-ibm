// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package iampolicy_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIBMIAMPolicyTemplateVersionBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMPolicyTemplateDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPolicyTemplateVersionConfigBasic(name, serviceName, "iam-identity"),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMPolicyTemplateExists("ibm_iam_policy_template.policy_template", conf),
					resource.TestCheckResourceAttr("ibm_iam_policy_template_version.template_version", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_policy_template.policy_template", "policy.0.resource.0.attributes.0.value", serviceName),
					resource.TestCheckResourceAttr("ibm_iam_policy_template_version.template_version", "policy.0.resource.0.attributes.0.value", "iam-identity"),
				),
			},
		},
	})
}

func TestAccIBMIAMPolicyTemplateVersionUpdateCommit(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMPolicyTemplateDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPolicyTemplateVersionConfigBasic(name, serviceName, "iam-identity"),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMPolicyTemplateExists("ibm_iam_policy_template.policy_template", conf),
					resource.TestCheckResourceAttr("ibm_iam_policy_template_version.template_version", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_policy_template.policy_template", "policy.0.resource.0.attributes.0.value", serviceName),
					resource.TestCheckResourceAttr("ibm_iam_policy_template_version.template_version", "policy.0.resource.0.attributes.0.value", "iam-identity"),
				),
			},
			{
				Config: testAccCheckIBMPolicyTemplateVersionUpdateCommit(name, serviceName, "iam-identity"),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMPolicyTemplateExists("ibm_iam_policy_template.policy_template", conf),
					resource.TestCheckResourceAttr("ibm_iam_policy_template_version.template_version", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_policy_template.policy_template", "policy.0.resource.0.attributes.0.value", serviceName),
					resource.TestCheckResourceAttr("ibm_iam_policy_template_version.template_version", "committed", "true"),
				),
			},
		},
	})
}

func testAccCheckIBMPolicyTemplateVersionConfigBasic(name string, serviceName string, updatedService string) string {
	return fmt.Sprintf(`
		resource "ibm_iam_policy_template" "policy_template" {
			name = "%s"
			policy {
				type = "access"
				description = "description"
				resource {
					attributes {
						key = "serviceName"
						operator = "stringEquals"
						value = "%s"
					}
				}
				roles = ["Operator"]
			}
		}

		resource "ibm_iam_policy_template_version" "template_version" {
			template_id = ibm_iam_policy_template.policy_template.template_id
			policy {
				type = "access"
				description = "description"
				resource {
					attributes {
						key = "serviceName"
						operator = "stringEquals"
						value = "%s"
					}
				}
				roles = ["Service ID creator", "Operator"]
			}
			description = "Template version"
		}
	`, name, serviceName, updatedService)
}

func testAccCheckIBMPolicyTemplateVersionUpdateCommit(name string, serviceName string, updatedService string) string {
	return fmt.Sprintf(`
		resource "ibm_iam_policy_template" "policy_template" {
			name = "%s"
			policy {
				type = "access"
				description = "description"
				resource {
					attributes {
						key = "serviceName"
						operator = "stringEquals"
						value = "%s"
					}
				}
				roles = ["Operator"]
			}
		}

		resource "ibm_iam_policy_template_version" "template_version" {
			template_id = ibm_iam_policy_template.policy_template.template_id
			policy {
				type = "access"
				description = "description"
				resource {
					attributes {
						key = "serviceName"
						operator = "stringEquals"
						value = "%s"
					}
				}
				roles = ["Service ID creator", "Operator"]
			}
			description = "Template version"
			committed = true
		}
	`, name, serviceName, updatedService)
}
