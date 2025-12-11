// Copyright IBM Corp. 2025 All Rights Reserved.
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
	roleTemplateName string = fmt.Sprintf("TerraformRoleTemplateTest%d", acctest.RandIntRange(10, 100))
	roleName         string = fmt.Sprintf("TerraformRoleTest%d", acctest.RandIntRange(10, 100))
	roleConf         iampolicymanagementv1.RoleTemplate
	displayName      string = "TerraformRoleTest"
)

func TestAccIBMIAMRoleTemplateBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMRoleTemplateDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMRoleTemplateConfigBasic(roleTemplateName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMRoleTemplateExists("ibm_iam_role_template.role_template", roleConf),
					resource.TestCheckResourceAttr("ibm_iam_role_template.role_template", "name", roleTemplateName),
				),
			},
		},
	})
}

func TestAccIBMIAMRoleTemplateBasicUpdateWithRole(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMRoleTemplateDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMRoleTemplateConfigBasic(roleTemplateName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMRoleTemplateExists("ibm_iam_role_template.role_template", roleConf),
					resource.TestCheckResourceAttr("ibm_iam_role_template.role_template", "name", roleTemplateName),
				),
			},
			{
				Config: testAccCheckIBMRoleTemplateConfigWithRole(roleTemplateName, roleName, displayName, roleCreateAction),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMRoleTemplateExists("ibm_iam_role_template.role_template", roleConf),
					resource.TestCheckResourceAttr("ibm_iam_role_template.role_template", "name", roleTemplateName),
					resource.TestCheckResourceAttr("ibm_iam_role_template.role_template", "role.0.actions.0", roleCreateAction),
				),
			},
		},
	})
}

func TestAccIBMIAMRoleTemplateBasicCommit(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMRoleTemplateDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMRoleTemplateConfigBasic(roleTemplateName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMRoleTemplateExists("ibm_iam_role_template.role_template", roleConf),
					resource.TestCheckResourceAttr("ibm_iam_role_template.role_template", "name", roleTemplateName),
				),
			},
			{
				Config: testAccCheckIBMRoleTemplateConfigWithRoleAndCommit("UpdateBasicTemplate", roleName, displayName, roleCreateAction),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMRoleTemplateExists("ibm_iam_role_template.role_template", roleConf),
					resource.TestCheckResourceAttr("ibm_iam_role_template.role_template", "name", "UpdateBasicTemplate"),
					resource.TestCheckResourceAttr("ibm_iam_role_template.role_template", "role.0.actions.0", roleCreateAction),
					resource.TestCheckResourceAttr("ibm_iam_role_template.role_template", "committed", "true"),
				),
			},
		},
	})
}

func TestAccIBMIAMRoleTemplate(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMRoleTemplateDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMRoleTemplateConfigWithRoleAndCommit("RoleTemplate", roleName, displayName, roleCreateAction),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMRoleTemplateExists("ibm_iam_role_template.role_template", roleConf),
					resource.TestCheckResourceAttr("ibm_iam_role_template.role_template", "name", "RoleTemplate"),
					resource.TestCheckResourceAttr("ibm_iam_role_template.role_template", "role.0.actions.0", roleCreateAction),
					resource.TestCheckResourceAttr("ibm_iam_role_template.role_template", "committed", "true"),
				),
			},
		},
	})
}

func testAccCheckIBMRoleTemplateExists(n string, obj iampolicymanagementv1.RoleTemplate) resource.TestCheckFunc {

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

		getRoleTemplateOptions := &iampolicymanagementv1.GetRoleTemplateVersionOptions{
			RoleTemplateID: &parts[0],
			Version:        &parts[1],
		}

		roleTemplate, _, err := iamPolicyManagementClient.GetRoleTemplateVersion(getRoleTemplateOptions)
		if err != nil {
			return err
		}
		obj = *roleTemplate
		return nil
	}
}

func testAccCheckIBMRoleTemplateDestroy(s *terraform.State) error {
	iamPolicyManagementClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).IAMPolicyManagementV1API()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_iam_role_template" {
			continue
		}

		parts, err := flex.IdParts(rs.Primary.ID)
		if err != nil {
			return err
		}

		getRoleTemplateOptions := &iampolicymanagementv1.GetRoleTemplateVersionOptions{
			RoleTemplateID: &parts[0],
			Version:        &parts[1],
		}

		// Try to find the key
		_, response, err := iamPolicyManagementClient.GetRoleTemplateVersion(getRoleTemplateOptions)

		if err == nil {
			return fmt.Errorf("ibm_iam_role_template still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for ibm_iam_role_template (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}

func testAccCheckIBMRoleTemplateConfigBasic(name string) string {
	return fmt.Sprintf(`

		resource "ibm_iam_role_template" "role_template" {
			name = "%s"
			description = "Create Action Control basic templates through Terraform resources"
		}
	`, name)
}

func testAccCheckIBMRoleTemplateConfigWithRole(name string, roleName string, displayName string, role string) string {
	return fmt.Sprintf(`

		resource "ibm_iam_role_template" "role_template" {
			name = "%s"
			description = "Update Action Control basic templates through Terraform resources"
			role {
			    name = "%s"
				display_name = "%s"
				actions = ["%s" ]
				service_name="cloud-object-storage"
			} 
		}
	`, name, roleName, displayName, role)
}

func testAccCheckIBMRoleTemplateConfigWithRoleAndCommit(name string, roleName string, displayName string, role string) string {
	return fmt.Sprintf(`

		resource "ibm_iam_role_template" "role_template" {
			name = "%s"
			description = "Commit Action Control templates through Terraform resources"
			role {
				name = "%s"
				display_name = "%s"
				actions = ["%s" ]
				service_name="cloud-object-storage"
			}
			committed = true
		}
	`, name, roleName, displayName, role)
}
