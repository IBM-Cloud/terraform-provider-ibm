// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package iampolicy_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

var (
	resourceServiceName        string = "kms"
	updatedResourceServiceName string = "appid"
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

func TestAccIBMIAMPolicyTemplateVersionBasicWithRoleReference(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMPolicyTemplateDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPolicyRoleTemplateVersionConfigBasic(name, serviceName, "iam-identity"),
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

func TestAccIBMIAMPolicyTemplateVersionBasicWithPolcyType(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMPolicyTemplateDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPolicyTemplateVersionConfigBasicWithPolicyType(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMPolicyTemplateExists("ibm_iam_policy_template.policy_template", conf),
					resource.TestCheckResourceAttr("ibm_iam_policy_template_version.template_version", "name", name),
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

func TestAccIBMIAMPolicyTemplateVersionUpdateWithRoleReference(t *testing.T) {
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
				Config: testAccCheckIBMPolicyRoleTemplateVersionUpdate(name, serviceName, "iam-identity"),
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

func TestAccIBMIAMPolicyTemplateVersionS2SBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMPolicyTemplateDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPolicyTemplateVersionConfigS2SBasic(name, sourceServiceName, resourceServiceName, sourceServiceName, updatedResourceServiceName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMPolicyTemplateExists("ibm_iam_policy_template.policy_template", conf),
					resource.TestCheckResourceAttr("ibm_iam_policy_template_version.template_version", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_policy_template.policy_template", "policy.0.resource.0.attributes.0.value", resourceServiceName),
					resource.TestCheckResourceAttr("ibm_iam_policy_template_version.template_version", "policy.0.resource.0.attributes.0.value", updatedResourceServiceName),
				),
			},
		},
	})
}

func TestAccIBMIAMPolicyTemplateVersionS2SUpdate(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMPolicyTemplateDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPolicyTemplateVersionConfigS2SBasic(name, sourceServiceName, resourceServiceName, sourceServiceName, updatedResourceServiceName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMPolicyTemplateExists("ibm_iam_policy_template.policy_template", conf),
					resource.TestCheckResourceAttr("ibm_iam_policy_template_version.template_version", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_policy_template.policy_template", "policy.0.resource.0.attributes.0.value", resourceServiceName),
					resource.TestCheckResourceAttr("ibm_iam_policy_template_version.template_version", "policy.0.resource.0.attributes.0.value", updatedResourceServiceName),
				),
			},
			{
				Config: testAccCheckIBMPolicyTemplateVersionConfigS2SUpdateCommit(name, sourceServiceName, resourceServiceName, sourceServiceName, updatedResourceServiceName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMPolicyTemplateExists("ibm_iam_policy_template.policy_template", conf),
					resource.TestCheckResourceAttr("ibm_iam_policy_template_version.template_version", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_policy_template.policy_template", "policy.0.resource.0.attributes.0.value", resourceServiceName),
					resource.TestCheckResourceAttr("ibm_iam_policy_template_version.template_version", "policy.0.resource.0.attributes.0.value", updatedResourceServiceName),
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

func testAccCheckIBMPolicyRoleTemplateVersionConfigBasic(name string, serviceName string, updatedService string) string {
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
		resource "ibm_iam_role_template" "role_template" {
			name = "TerraformPolicyRoleUpdateTest"
			description = "Create role template and reference in update policy template through Terraform resources"
			role {
			    name = "TerrPolicyRoleCreate"
				display_name = "TestingTerraformPolicyRole"
				actions = ["iam-identity.serviceid.create" ]
				service_name="iam-identity"
			}
			committed = true
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
				role_template_references {
					id = ibm_iam_role_template.role_template.role_template_id
					version = ibm_iam_role_template.role_template.version
				}
			}
			description = "Template version"
		}
	`, name, serviceName, updatedService)
}

func testAccCheckIBMPolicyTemplateVersionConfigBasicWithPolicyType(name string) string {
	return fmt.Sprintf(`
		resource "ibm_iam_policy_template" "policy_template" {
			name = "%s"
			policy {
				type = "access"
				description = "description"
			}
		}

		resource "ibm_iam_policy_template_version" "template_version" {
			template_id = ibm_iam_policy_template.policy_template.template_id
			policy {
				type = "access"
				description = "template description"
			}
			description = "Template version"
		}
	`, name)
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

func testAccCheckIBMPolicyRoleTemplateVersionUpdate(name string, serviceName string, updatedService string) string {
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

		resource "ibm_iam_role_template" "role_template" {
			name = "TerraformPolicyRoleUpdateTest"
			description = "Create role template and reference in update policy template through Terraform resources"
			role {
			    name = "TerrPolicyRoleCreate"
				display_name = "TestingTerraformPolicyRole"
				actions = ["iam-identity.serviceid.create" ]
				service_name="iam-identity"
			}
			committed = true
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
				role_template_references {
					id = ibm_iam_role_template.role_template.role_template_id
					version = ibm_iam_role_template.role_template.version
				}
			}
			description = "Template version"
			committed = true
		}
	`, name, serviceName, updatedService)
}

func testAccCheckIBMPolicyTemplateVersionConfigS2SBasic(name string, sourceServiceName string, resourceServiceName string, versionServiceName string, updatedResourceServiceName string) string {
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

		resource "ibm_iam_policy_template_version" "template_version" {
			template_id = ibm_iam_policy_template.policy_template.template_id
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
	`, name, sourceServiceName, resourceServiceName, versionServiceName, updatedResourceServiceName)
}

func testAccCheckIBMPolicyTemplateVersionConfigS2SUpdateCommit(name string, sourceServiceName string, resourceServiceName string, versionServiceName string, updatedResourceServiceName string) string {
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

		resource "ibm_iam_policy_template_version" "template_version" {
			template_id = ibm_iam_policy_template.policy_template.template_id
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
			committed= true
		}
	`, name, sourceServiceName, resourceServiceName, versionServiceName, updatedResourceServiceName)
}
