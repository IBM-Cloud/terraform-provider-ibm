// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package iampolicy_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM/platform-services-go-sdk/iampolicymanagementv1"
)

func TestAccIBMIAMPolicyTemplateVersionBasic(t *testing.T) {
	var conf iampolicymanagementv1.PolicyTemplate
	name := fmt.Sprintf("TerraformTemplateTest%d", acctest.RandIntRange(10, 100))
	accountID := "e3aa0adff348470f803d4b6e53d625cf"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMPolicyTemplateDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPolicyTemplateVersionConfigBasic(name, accountID),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMPolicyTemplateExists("ibm_iam_policy_template.policy_template", conf),
					resource.TestCheckResourceAttr("ibm_iam_policy_template_version.template_version", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_policy_template_version.template_version", "account_id", accountID),
				),
			},
		},
	})
}

func testAccCheckIBMPolicyTemplateVersionConfigBasic(name string, accountID string) string {
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
						value = "is"
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
						value = "iam-identity"
					}
				}
				roles = ["Service ID creator", "Operator"]
			}
			description = "Template version"
		}

	`, name, accountID)
}
