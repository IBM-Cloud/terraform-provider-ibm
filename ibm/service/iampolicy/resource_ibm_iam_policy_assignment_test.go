// Copyright IBM Corp. 2024 All Rights Reserved.
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
	"github.com/IBM/platform-services-go-sdk/iampolicymanagementv1"
)

func TestAccIBMPolicyAssignmentBasic(t *testing.T) {
	var conf iampolicymanagementv1.GetPolicyAssignmentResponse
	var name string = fmt.Sprintf("TerraformTemplateTest%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMPolicyAssignmentDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPolicyAssignmentConfigBasic(name, acc.TargetAccountId),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMPolicyAssignmentExists("ibm_iam_policy_assignment.policy_assignment", conf),
					resource.TestCheckResourceAttr("ibm_iam_policy_template.policy_s2s_template", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_policy_template.policy_s2s_template", "policy.0.resource.0.attributes.0.value", resourceServiceName),
					resource.TestCheckResourceAttr("ibm_iam_policy_template.policy_s2s_template", "policy.0.subject.0.attributes.0.value", "compliance"),
				),
			},
			{
				Config: testAccCheckIBMPolicyAssignmentConfigUpdate(name, acc.TargetAccountId),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMPolicyAssignmentExists("ibm_iam_policy_assignment.policy_assignment", conf),
					resource.TestCheckResourceAttr("ibm_iam_policy_template_version.template_version", "policy.0.subject.0.attributes.0.value", "compliance"),
					resource.TestCheckResourceAttr("ibm_iam_policy_template_version.template_version", "policy.0.resource.0.attributes.0.value", updatedResourceServiceName),
				),
			},
		},
	})
}


func TestAccIBMPolicyAssignmentS2SBasic(t *testing.T) {
	var conf iampolicymanagementv1.GetPolicyAssignmentResponse
	var name string = fmt.Sprintf("TerraformTemplateTest%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMPolicyAssignmentDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPolicyAssignmentS2SConfigBasic(name, acc.TargetAccountId),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMPolicyAssignmentExists("ibm_iam_policy_assignment.policy_assignment", conf),
					resource.TestCheckResourceAttr("ibm_iam_policy_template.policy_s2s_template", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_policy_template.policy_s2s_template", "policy.0.resource.0.attributes.0.value", "is"),
					resource.TestCheckResourceAttr("ibm_iam_policy_template.policy_s2s_template", "policy.0.resource.0.attributes.1.value", "true"),
					resource.TestCheckResourceAttr("ibm_iam_policy_template.policy_s2s_template", "policy.0.subject.0.attributes.0.value", "is"),
					resource.TestCheckResourceAttr("ibm_iam_policy_template.policy_s2s_template", "policy.0.subject.0.attributes.1.value", "backup-policy"),
				),
			},
		},
	})
}

func TestAccIBMPolicyAssignmentEnterprise(t *testing.T) {
	var conf iampolicymanagementv1.GetPolicyAssignmentResponse
	var name string = fmt.Sprintf("TerraformTemplateTest%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMPolicyAssignmentDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPolicyAssignmentConfigEnterprise(name, acc.TargetEnterpriseId),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMPolicyAssignmentExists("ibm_iam_policy_assignment.policy_assignment", conf),
					resource.TestCheckResourceAttr("ibm_iam_policy_template.policy_s2s_template", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_policy_template.policy_s2s_template", "policy.0.resource.0.attributes.0.value", resourceServiceName),
					resource.TestCheckResourceAttr("ibm_iam_policy_template.policy_s2s_template", "policy.0.subject.0.attributes.0.value", "compliance"),
				),
			},
		},
	})
}

func testAccCheckIBMPolicyAssignmentConfigBasic(name string, targetId string) string {
	return fmt.Sprintf(`
		resource "ibm_iam_policy_template" "policy_s2s_template" {
			name = "%s"
			policy {
				type = "authorization"
				description = "description"
				resource {
					attributes {
						key = "serviceName"
						operator = "stringEquals"
						value = "kms"
					}
				}
				subject {
					attributes {
						key = "serviceName"
						operator = "stringEquals"
						value = "compliance"
					}
				}
				roles = ["Reader"]
			}
			committed=true
		}
		resource "ibm_iam_policy_assignment" "policy_assignment" {
			version = "1.0"
			target  ={
				type = "Account"
				id = "%s"
			}
			templates{
				id = ibm_iam_policy_template.policy_s2s_template.template_id 
				version = ibm_iam_policy_template.policy_s2s_template.version
			}
		}`, name, targetId)
}

func testAccCheckIBMPolicyAssignmentS2SConfigBasic(name string, targetId string) string {
	return fmt.Sprintf(`
		resource "ibm_iam_policy_template" "policy_s2s_template" {
			name = "%s"
			policy {
				type = "authorization"
				description = "Test terraform enterprise S2S"
				resource {
					attributes {
						key = "serviceName"
						operator = "stringEquals"
						value = "is"
					}
					attributes {
						key = "volumeId"
						operator = "stringExists"
						value = "true"
					}
				}
				subject {
					attributes {
						key = "serviceName"
						operator = "stringEquals"
						value = "is"
					}
					attributes {
						key = "resourceType"
						operator = "stringEquals"
						value = "backup-policy"
					}
				}
				roles = ["Operator"]
			}
			committed=true
		}
		resource "ibm_iam_policy_assignment" "policy_assignment" {
			version = "1.0"
			target  ={
				type = "Account"
				id = "%s"
			}
			templates{
				id = ibm_iam_policy_template.policy_s2s_template.template_id 
				version = ibm_iam_policy_template.policy_s2s_template.version
			}
		}
		
	`, name, targetId)
}

func testAccCheckIBMPolicyAssignmentConfigUpdate(name string, targetId string) string {
	return fmt.Sprintf(`
		resource "ibm_iam_policy_template" "policy_s2stemplate" {
			name = "%s"
			policy {
				type = "authorization"
				description = "description"
				resource {
					attributes {
						key = "serviceName"
						operator = "stringEquals"
						value = "kms"
					}
				}
				subject {
					attributes {
						key = "serviceName"
						operator = "stringEquals"
						value = "compliance"
					}
				}
				roles = ["Reader"]
			}
			committed=true
		}
		resource "ibm_iam_policy_template_version" "template_version" {
			template_id = ibm_iam_policy_template.policy_s2s_template.template_id
			policy {
				type = "authorization"
				description = "description"
				resource {
					attributes {
						key = "serviceName"
						operator = "stringEquals"
						value = "appid"
					}
				}
				subject {
					attributes {
						key = "serviceName"
						operator = "stringEquals"
						value = "compliance"
					}
				}
				roles = ["Reader"]
			}
			committed=true
		}
	
		resource "ibm_iam_policy_assignment" "policy_assignment" {
			version = "1.0"
			target  ={
				type = "Account"
				id = "%s"
			}

			templates{
                id = ibm_iam_policy_template_version.template_version.template_id
				version = ibm_iam_policy_template_version.template_version.version
			}
		}`, name, targetId)
}

func testAccCheckIBMPolicyAssignmentConfigEnterprise(name string, targetId string) string {
	return fmt.Sprintf(`
		resource "ibm_iam_policy_template" "policy_s2s_template" {
			name = "%s"
			policy {
				type = "authorization"
				description = "description"
				resource {
					attributes {
						key = "serviceName"
						operator = "stringEquals"
						value = "kms"
					}
				}
				subject {
					attributes {
						key = "serviceName"
						operator = "stringEquals"
						value = "compliance"
					}
				}
				roles = ["Reader"]
			}
			committed=true
		}
		resource "ibm_iam_policy_assignment" "policy_assignment" {
			version = "1.0"
			target  ={
				type = "Enterprise"
				id = "%s"
			}
			templates{
				id = ibm_iam_policy_template.policy_s2s_template.template_id 
				version = ibm_iam_policy_template.policy_s2s_template.version
			}
		}`, name, targetId)
}

func testAccCheckIBMPolicyAssignmentExists(n string, obj iampolicymanagementv1.GetPolicyAssignmentResponse) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		iamPolicyManagementClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).IAMPolicyManagementV1API()
		if err != nil {
			return err
		}

		getPolicyAssignmentsOptions := &iampolicymanagementv1.GetPolicyAssignmentOptions{}

		getPolicyAssignmentsOptions.SetAssignmentID(rs.Primary.ID)
		getPolicyAssignmentsOptions.SetVersion("1.0")

		policyAssignmentV1, _, err := iamPolicyManagementClient.GetPolicyAssignment(getPolicyAssignmentsOptions)
		if err != nil {
			return err
		}

		assignmentDetails := policyAssignmentV1.(*iampolicymanagementv1.GetPolicyAssignmentResponse)

		obj = *assignmentDetails
		return nil
	}
}

func testAccCheckIBMPolicyAssignmentDestroy(s *terraform.State) error {
	iamPolicyManagementClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).IAMPolicyManagementV1API()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_policy_assignment" {
			continue
		}

		getPolicyAssignmentsOptions := &iampolicymanagementv1.GetPolicyAssignmentOptions{}

		getPolicyAssignmentsOptions.SetAssignmentID(rs.Primary.ID)
		getPolicyAssignmentsOptions.SetVersion("1.0")

		// Try to find the key
		_, response, err := iamPolicyManagementClient.GetPolicyAssignment(getPolicyAssignmentsOptions)

		if err == nil {
			return fmt.Errorf("policy_assignment still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for policy_assignment (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
