// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package iampolicy_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

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
	sourceServiceName       string = "compliance"
)

func TestAccIBMIAMPolicyTemplateBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMPolicyTemplateDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPolicyTemplateConfigBasic(name, serviceName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMPolicyTemplateExists("ibm_iam_policy_template.policy_template", conf),
					resource.TestCheckResourceAttr("ibm_iam_policy_template.policy_template", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_policy_template.policy_template", "policy.0.resource.0.attributes.0.value", serviceName),
				),
			},
		},
	})
}

func TestAccIBMIAMPolicyTemplateBasicS2SUpdateTest(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMPolicyTemplateDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPolicyS2STemplateUpdateConfigBasicTest(name, "Service ID creator", "iam-identity", "secrets-manager"),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMPolicyTemplateExists("ibm_iam_policy_template.policy_template", conf),
					resource.TestCheckResourceAttr("ibm_iam_policy_template.policy_template", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_policy_template.policy_template", "policy.0.resource.0.attributes.0.value", "iam-identity"),
					resource.TestCheckResourceAttr("ibm_iam_policy_template.policy_template", "policy.0.subject.0.attributes.0.value", "secrets-manager"),
				),
			},
			{
				Config: testAccCheckIBMPolicyS2STemplateUpdateConfigBasicTest(name, "Operator", "iam-identity", "secrets-manager"),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMPolicyTemplateExists("ibm_iam_policy_template.policy_template", conf),
					resource.TestCheckResourceAttr("ibm_iam_policy_template.policy_template", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_policy_template.policy_template", "policy.0.resource.0.attributes.0.value", "iam-identity"),
					resource.TestCheckResourceAttr("ibm_iam_policy_template.policy_template", "policy.0.subject.0.attributes.0.value", "secrets-manager"),
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
				Config: testAccCheckIBMPolicyTemplateConfigBasic(name, beforeUpdateServiceName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMPolicyTemplateExists("ibm_iam_policy_template.policy_template", conf),
					resource.TestCheckResourceAttr("ibm_iam_policy_template.policy_template", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_policy_template.policy_template", "policy.0.resource.0.attributes.0.value", beforeUpdateServiceName),
				),
			},
			{
				Config: testAccCheckIBMPolicyTemplateConfigBasicUpdate(name, updatedServiceName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMPolicyTemplateExists("ibm_iam_policy_template.policy_template", conf),
					resource.TestCheckResourceAttr("ibm_iam_policy_template.policy_template", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_policy_template.policy_template", "policy.0.resource.0.attributes.0.value", updatedServiceName),
				),
			},
		},
	})
}

func TestAccIBMIAMPolicyTemplateBasicWithTags(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMPolicyTemplateDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPolicyTemplateConfigBasicWithTags(name, serviceName, "testing_tags"),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMPolicyTemplateExists("ibm_iam_policy_template.policy_template", conf),
					resource.TestCheckResourceAttr("ibm_iam_policy_template.policy_template", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_policy_template.policy_template", "policy.0.resource.0.attributes.0.value", serviceName),
					resource.TestCheckResourceAttr("ibm_iam_policy_template.policy_template", "policy.0.resource.0.tags.0.value", "testing_tags"),
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
				Config: testAccCheckIBMPolicyTemplateConfigBasicCommit(name, serviceName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMPolicyTemplateExists("ibm_iam_policy_template.policy_template", conf),
					resource.TestCheckResourceAttr("ibm_iam_policy_template.policy_template", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_policy_template.policy_template", "policy.0.resource.0.attributes.0.value", serviceName),
					resource.TestCheckResourceAttr("ibm_iam_policy_template.policy_template", "committed", "true"),
				),
			},
		},
	})
}

func TestAccIBMIAMPolicyTemplateBasicCommitWithPolicyType(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMPolicyTemplateDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPolicyTemplateConfigBasicTest(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMPolicyTemplateExists("ibm_iam_policy_template.policy_template", conf),
					resource.TestCheckResourceAttr("ibm_iam_policy_template.policy_template", "name", name),
				),
			},
		},
	})
}

func TestAccIBMIAMPolicyTemplateBasicS2S(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMPolicyTemplateDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPolicyS2STemplateConfigBasic(name, sourceServiceName, updatedServiceName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMPolicyTemplateExists("ibm_iam_policy_template.policy_template", conf),
					resource.TestCheckResourceAttr("ibm_iam_policy_template.policy_template", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_policy_template.policy_template", "policy.0.resource.0.attributes.0.value", updatedServiceName),
					resource.TestCheckResourceAttr("ibm_iam_policy_template.policy_template", "policy.0.subject.0.attributes.0.value", sourceServiceName),
				),
			},
		},
	})
}

func TestAccIBMIAMPolicyTemplateBasicS2STest(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMPolicyTemplateDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPolicyS2STemplateConfigBasicTest("TerraformS2STest", "is", "true", "is", "backup-policy"),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMPolicyTemplateExists("ibm_iam_policy_template.policy_template", conf),
					resource.TestCheckResourceAttr("ibm_iam_policy_template.policy_template", "name", "TerraformS2STest"),
					resource.TestCheckResourceAttr("ibm_iam_policy_template.policy_template", "policy.0.resource.0.attributes.0.value", "is"),
					resource.TestCheckResourceAttr("ibm_iam_policy_template.policy_template", "policy.0.resource.0.attributes.1.value", "true"),
					resource.TestCheckResourceAttr("ibm_iam_policy_template.policy_template", "policy.0.subject.0.attributes.0.value", "is"),
					resource.TestCheckResourceAttr("ibm_iam_policy_template.policy_template", "policy.0.subject.0.attributes.1.value", "backup-policy"),
				),
			},
		},
	})
}

func TestAccIBMIAMPolicyTemplateBasicS2SUpdate(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMPolicyTemplateDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPolicyS2STemplateConfigBasic(name, sourceServiceName, updatedServiceName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMPolicyTemplateExists("ibm_iam_policy_template.policy_template", conf),
					resource.TestCheckResourceAttr("ibm_iam_policy_template.policy_template", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_policy_template.policy_template", "policy.0.resource.0.attributes.0.value", updatedServiceName),
					resource.TestCheckResourceAttr("ibm_iam_policy_template.policy_template", "policy.0.subject.0.attributes.0.value", sourceServiceName),
				),
			},
			{
				Config: testAccCheckIBMPolicyS2STemplateConfigUpdate(name, sourceServiceName, "appid"),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMPolicyTemplateExists("ibm_iam_policy_template.policy_template", conf),
					resource.TestCheckResourceAttr("ibm_iam_policy_template.policy_template", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_policy_template.policy_template", "policy.0.resource.0.attributes.0.value", "appid"),
					resource.TestCheckResourceAttr("ibm_iam_policy_template.policy_template", "policy.0.subject.0.attributes.0.value", sourceServiceName),
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

func testAccCheckIBMPolicyTemplateConfigBasic(name string, serviceName string) string {
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
	`, name, serviceName)
}

func testAccCheckIBMPolicyS2STemplateUpdateConfigBasicTest(name string, role string, resourceServiceName string, subjectServiceName string) string {
	return fmt.Sprintf(`

		resource "ibm_iam_policy_template" "policy_template" {
    		name        = "%s"
    		policy {
        		description = "description"
        		roles       = [
            		"%s",
        		]
        		type        = "authorization"

        		resource {
            		attributes {
                		key      = "serviceName"
                		operator = "stringEquals"
                		value    = "%s"
            		}
        		}

        		subject {
            		attributes {
                		key      = "serviceName"
                		operator = "stringEquals"
                		value    = "%s"
            		}
        		}
    		}
		}
	`, name, role, resourceServiceName, subjectServiceName)
}

func testAccCheckIBMPolicyS2STemplateConfigBasicTest(name string, sourceServiceName string, volumeId string, serviceName string, resourceServiceName string) string {
	return fmt.Sprintf(`

		resource "ibm_iam_policy_template" "policy_template" {
			name = "%s"
			policy {
				type = "authorization"
				description = "Test terraform enterprise S2S"
				resource {
					attributes {
						key = "serviceName"
						operator = "stringEquals"
						value = "%s"
					}
					attributes {
						key = "volumeId"
						operator = "stringExists"
						value = "%s"
					}
				}
				subject {
					attributes {
						key = "serviceName"
						operator = "stringEquals"
						value = "%s"
					}
					attributes {
						key = "resourceType"
						operator = "stringEquals"
						value = "%s"
					}
				}
				roles = ["Operator"]
			}
			committed=true
		}
	`, name, sourceServiceName, volumeId, serviceName, resourceServiceName)
}

func testAccCheckIBMPolicyS2STemplateConfigBasic(name string, sourceServiceName string, resourceServiceName string) string {
	return fmt.Sprintf(`

		resource "ibm_iam_policy_template" "policy_template" {
			name = "%s"
			policy {
				type = "authorization"
				description = "Test terraform enterprise S2S"
				subject {
					attributes {
						key = "serviceName"
						operator = "stringEquals"
						value = "%s"
					}
				}
				resource {
					attributes {
						key = "serviceName"
						operator = "stringEquals"
						value = "%s"
					}
				}
				roles = ["Reader"]
			}
		}
	`, name, sourceServiceName, resourceServiceName)
}

func testAccCheckIBMPolicyS2STemplateConfigUpdate(name string, sourceServiceName string, resourceServiceName string) string {
	return fmt.Sprintf(`

		resource "ibm_iam_policy_template" "policy_template" {
			name = "%s"
			policy {
				type = "authorization"
				description = "Test terraform enterprise update S2S"
				subject {
					attributes {
						key = "serviceName"
						operator = "stringEquals"
						value = "%s"
					}
				}
				resource {
					attributes {
						key = "serviceName"
						operator = "stringEquals"
						value = "%s"
					}
				}
				roles = ["Reader"]
			}
		}
	`, name, sourceServiceName, resourceServiceName)
}

func testAccCheckIBMPolicyTemplateConfigBasicUpdate(name string, serviceName string) string {
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
				roles = ["Operator", "KeyPurge"]
			}
		}
	`, name, serviceName)
}

func testAccCheckIBMPolicyTemplateConfigBasicCommit(name string, serviceName string) string {
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
			committed = true
		}
	`, name, serviceName)
}

func testAccCheckIBMPolicyTemplateConfigBasicTest(name string) string {
	return fmt.Sprintf(`

		resource "ibm_iam_policy_template" "policy_template" {
			name = "%s"
			policy {
				type = "access"
				description = "description"
			}
		}
	`, name)
}

func testAccCheckIBMPolicyTemplateConfigBasicWithTags(name string, serviceName string, tagValue string) string {
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
					tags {
						key = "terraform"
						operator = "stringEquals"
						value = "%s"
					}
				}
				roles = ["Operator"]
			}
		}
	`, name, serviceName, tagValue)
}
