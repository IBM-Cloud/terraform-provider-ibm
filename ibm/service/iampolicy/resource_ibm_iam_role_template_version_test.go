// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package iampolicy_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var roleVersionName string = fmt.Sprintf("TerraformRoleTemplateVersionTemplateTest%d", acctest.RandIntRange(10, 100))

var basicRoleName string = fmt.Sprintf("TerraformRoleTemplateTest%d", acctest.RandIntRange(10, 100))
var basicRoleTemplateVersionName string = fmt.Sprintf("TerraformRoleVersionTest%d", acctest.RandIntRange(10, 100))

func TestAccIBMIAMRoleVersionTemplate(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMRoleTemplateDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMRoleTemplateVersionTemplateConfig(roleVersionName, basicRoleName, actCreateAction, actDeleteAction),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMRoleTemplateExists("ibm_iam_role_template.role_template", roleConf),
					resource.TestCheckResourceAttr("ibm_iam_role_template.role_template", "name", roleVersionName),
					resource.TestCheckResourceAttr("ibm_iam_role_template_version.role_template_version", "name", roleVersionName),
					resource.TestCheckResourceAttr("ibm_iam_role_template.role_template", "role.0.actions.0", actCreateAction),
					resource.TestCheckResourceAttr("ibm_iam_role_template_version.role_template_version", "role.0.actions.0", actDeleteAction),
				),
			},
		},
	})
}

func TestAccIBMIAMRoleVersionTemplateUpdate(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMRoleTemplateDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMRoleTemplateVersionTemplateConfig(roleVersionName, basicRoleName, actCreateAction, actDeleteAction),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMRoleTemplateExists("ibm_iam_role_template.role_template", roleConf),
					resource.TestCheckResourceAttr("ibm_iam_role_template.role_template", "name", roleVersionName),
					resource.TestCheckResourceAttr("ibm_iam_role_template_version.role_template_version", "name", roleVersionName),
					resource.TestCheckResourceAttr("ibm_iam_role_template.role_template", "role.0.actions.0", actCreateAction),
					resource.TestCheckResourceAttr("ibm_iam_role_template_version.role_template_version", "role.0.actions.0", actDeleteAction),
				),
			},
			{
				Config: testAccCheckIBMRoleTemplateVersionTemplateUpdateConfig(roleVersionName, basicRoleName, actCreateAction, actCreateAction, actDeleteAction, false),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMRoleTemplateExists("ibm_iam_role_template.role_template", roleConf),
					resource.TestCheckResourceAttr("ibm_iam_role_template.role_template", "name", roleVersionName),
					resource.TestCheckResourceAttr("ibm_iam_role_template_version.role_template_version", "name", roleVersionName),
					resource.TestCheckResourceAttr("ibm_iam_role_template.role_template", "role.0.actions.0", actCreateAction),
					resource.TestCheckResourceAttr("ibm_iam_role_template_version.role_template_version", "role.0.actions.#", "2"),
				),
			},
		},
	})
}

func TestAccIBMIAMRoleVersionTemplateUpdateCommit(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMRoleTemplateDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMRoleTemplateVersionTemplateConfig(roleVersionName, basicRoleName, actCreateAction, actDeleteAction),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMRoleTemplateExists("ibm_iam_role_template.role_template", roleConf),
					resource.TestCheckResourceAttr("ibm_iam_role_template.role_template", "name", roleVersionName),
					resource.TestCheckResourceAttr("ibm_iam_role_template_version.role_template_version", "name", roleVersionName),
					resource.TestCheckResourceAttr("ibm_iam_role_template.role_template", "role.0.actions.0", actCreateAction),
					resource.TestCheckResourceAttr("ibm_iam_role_template_version.role_template_version", "role.0.actions.0", actDeleteAction),
				),
			},
			{
				Config: testAccCheckIBMRoleTemplateVersionTemplateUpdateConfig(roleVersionName, basicRoleName, actCreateAction, actCreateAction, actDeleteAction, true),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMRoleTemplateExists("ibm_iam_role_template.role_template", roleConf),
					resource.TestCheckResourceAttr("ibm_iam_role_template.role_template", "name", roleVersionName),
					resource.TestCheckResourceAttr("ibm_iam_role_template_version.role_template_version", "name", roleVersionName),
					resource.TestCheckResourceAttr("ibm_iam_role_template.role_template", "role.0.actions.0", actCreateAction),
					resource.TestCheckResourceAttr("ibm_iam_role_template_version.role_template_version", "role.0.actions.#", "2"),
				),
			},
		},
	})
}

func testAccCheckIBMRoleTemplateVersionTemplateConfig(name string, basicRoleName string, actionControl string, actionControl1 string) string {
	return fmt.Sprintf(`
		resource "ibm_iam_role_template" "role_template" {
			name = "%s"
			description = "Create Role base templates through Terraform resources"
			role {
				name = "%s"
				display_name = "BasicRoleTemplate"
				actions = ["%s"]
				service_name="am-test-service"
			}
		}
		resource "ibm_iam_role_template_version" "role_template_version" {
			role_template_id = ibm_iam_role_template.role_template.role_template_id
			description = "Create Role template versions through Terraform resources"
			role {
				display_name = "BasicRoleTemplateVersion"
				actions = ["%s"]
			}
		}
	`, name, basicRoleName, actionControl, actionControl1)
}

func testAccCheckIBMRoleTemplateVersionTemplateUpdateConfig(name string, basicRoleName string, actionControl string, actionControl1 string, actionControl2 string, committed bool) string {
	return fmt.Sprintf(`
		resource "ibm_iam_role_template" "role_template" {
			name = "%s"
			description = "Create Role template versions through Terraform resources"
			role {
				name = "%s"
				display_name = "UpdateRoleVersionTemplateActions"
				actions = ["%s"]
				service_name="am-test-service"
			}
		}
		resource "ibm_iam_role_template_version" "role_template_version" {
			role_template_id = ibm_iam_role_template.role_template.role_template_id
			description = "Update Role template versions through Terraform resources"
			role {
				display_name = "UpdateRoleVersionTemplateActions"
				actions = ["%s", "%s"]
			}
			committed = %t
		}
	`, name, basicRoleName, actionControl, actionControl1, actionControl2, committed)
}
