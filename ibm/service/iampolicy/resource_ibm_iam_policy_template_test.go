// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package iampolicy_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/platform-services-go-sdk/iampolicymanagementv1"
)

var (
	conf                    iampolicymanagementv1.PolicyTemplate
	name                    string = fmt.Sprintf("TerraformTemplateTest%d", acctest.RandIntRange(10, 100))
	serviceName             string = "is"
	beforeUpdateServiceName string = "conversation"
	updatedServiceName      string = "kms"
	accountID               string = acc.IAMAccountId
)

func TestAccIBMIAMPolicyTemplateBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMPolicyTemplateDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPolicyTemplateConfigBasic(name, accountID, serviceName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMPolicyTemplateExists("ibm_iam_policy_template.policy_template", conf),
					resource.TestCheckResourceAttr("ibm_iam_policy_template.policy_template", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_policy_template.policy_template", "account_id", accountID),
					resource.TestCheckResourceAttr("ibm_iam_policy_template.policy_template", "policy.0.resource.0.attributes.0.value", serviceName),
				),
			},
		},
	})
}

func TestAccIBMIAMPolicyTemplateBasicUpdate(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMPolicyTemplateDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPolicyTemplateConfigBasic(name, accountID, beforeUpdateServiceName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMPolicyTemplateExists("ibm_iam_policy_template.policy_template", conf),
					resource.TestCheckResourceAttr("ibm_iam_policy_template.policy_template", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_policy_template.policy_template", "account_id", accountID),
					resource.TestCheckResourceAttr("ibm_iam_policy_template.policy_template", "policy.0.resource.0.attributes.0.value", beforeUpdateServiceName),
				),
			},
			{
				Config: testAccCheckIBMPolicyTemplateConfigBasicUpdate(name, accountID, updatedServiceName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMPolicyTemplateExists("ibm_iam_policy_template.policy_template", conf),
					resource.TestCheckResourceAttr("ibm_iam_policy_template.policy_template", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_policy_template.policy_template", "account_id", accountID),
					resource.TestCheckResourceAttr("ibm_iam_policy_template.policy_template", "policy.0.resource.0.attributes.0.value", updatedServiceName),
				),
			},
		},
	})
}

func TestAccIBMIAMPolicyTemplateBasicCommit(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMPolicyTemplateDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPolicyTemplateConfigBasicCommit(name, accountID, serviceName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMPolicyTemplateExists("ibm_iam_policy_template.policy_template", conf),
					resource.TestCheckResourceAttr("ibm_iam_policy_template.policy_template", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_policy_template.policy_template", "account_id", accountID),
					resource.TestCheckResourceAttr("ibm_iam_policy_template.policy_template", "policy.0.resource.0.attributes.0.value", serviceName),
					resource.TestCheckResourceAttr("ibm_iam_policy_template.policy_template", "committed", "true"),
				),
			},
		},
	})
}

func testAccCheckIBMPolicyTemplateExists(n string, obj iampolicymanagementv1.PolicyTemplate) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		iamPolicyManagementClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).IAMPolicyManagementV1API()
		if err != nil {
			return err
		}

		parts, err := flex.IdParts(rs.Primary.ID)
		if err != nil {
			return err
		}

		getPolicyTemplateOptions := &iampolicymanagementv1.GetPolicyTemplateVersionOptions{
			PolicyTemplateID: &parts[0],
			Version:          &parts[1],
		}

		policyTemplate, _, err := iamPolicyManagementClient.GetPolicyTemplateVersion(getPolicyTemplateOptions)
		if err != nil {
			return err
		}
		obj = *policyTemplate
		return nil
	}
}

func testAccCheckIBMPolicyTemplateDestroy(s *terraform.State) error {
	iamPolicyManagementClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).IAMPolicyManagementV1API()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_iam_policy_template" {
			continue
		}

		parts, err := flex.IdParts(rs.Primary.ID)
		if err != nil {
			return err
		}

		getPolicyTemplateOptions := &iampolicymanagementv1.GetPolicyTemplateVersionOptions{
			PolicyTemplateID: &parts[0],
			Version:          &parts[1],
		}

		// Try to find the key
		_, response, err := iamPolicyManagementClient.GetPolicyTemplateVersion(getPolicyTemplateOptions)

		if err == nil {
			return fmt.Errorf("policy_template still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for policy_template (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}

func testAccCheckIBMPolicyTemplateConfigBasic(name string, accountID string, serviceName string) string {
	return fmt.Sprintf(`

		resource "ibm_iam_policy_template" "policy_template" {
			name = "%s"
			account_id = "%s"
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
	`, name, accountID, serviceName)
}

func testAccCheckIBMPolicyTemplateConfigBasicUpdate(name string, accountID string, serviceName string) string {
	return fmt.Sprintf(`

		resource "ibm_iam_policy_template" "policy_template" {
			name = "%s"
			account_id = "%s"
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
				roles = ["Operator", "KeyPurge"]
			}
		}
	`, name, accountID, serviceName)
}

func testAccCheckIBMPolicyTemplateConfigBasicCommit(name string, accountID string, serviceName string) string {
	return fmt.Sprintf(`

		resource "ibm_iam_policy_template" "policy_template" {
			name = "%s"
			account_id = "%s"
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
			committed = true
		}
	`, name, accountID, serviceName)
}
