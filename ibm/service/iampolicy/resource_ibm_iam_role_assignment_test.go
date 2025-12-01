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

var (
	roleCreateAction = "cloud-object-storage.bucket.delete_bucket"
	roleDeleteAction = "cloud-object-storage.bucket.delete_backup_policy"
)

func TestAccIBMIAMRoleAssignmentBasic(t *testing.T) {
	var conf iampolicymanagementv1.RoleAssignment
	var name string = fmt.Sprintf("TerraformTemplateTest%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMRoleAssignmentDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMRoleAssignmentConfigBasic(name, acc.TargetAccountId),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMRoleAssignmentExists("ibm_iam_role_assignment.role_assignment", conf),
					resource.TestCheckResourceAttr("ibm_iam_role_template.role_template", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_role_template.role_template", "role.0.actions.0", roleCreateAction),
					resource.TestCheckResourceAttr("ibm_iam_role_assignment.role_assignment", "operation", "apply"),
				),
			},
		},
	})
}

func TestAccIBMIAMRoleAssignmentUpdate(t *testing.T) {
	var conf iampolicymanagementv1.RoleAssignment
	var name string = fmt.Sprintf("TerraformTemplateTest%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMRoleAssignmentDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMRoleAssignmentConfigBasic(name, acc.TargetAccountId),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMRoleAssignmentExists("ibm_iam_role_assignment.role_assignment", conf),
					resource.TestCheckResourceAttr("ibm_iam_role_template.role_template", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_role_template.role_template", "role.0.actions.0", roleCreateAction),
				),
			},
			{
				Config: testAccCheckIBMRoleAssignmentConfigUpdate(name, acc.TargetAccountId),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMRoleAssignmentExists("ibm_iam_role_assignment.role_assignment", conf),
					resource.TestCheckResourceAttr("ibm_iam_role_template_version.template_version", "role.0.actions.0", roleDeleteAction),
					resource.TestCheckResourceAttr("ibm_iam_role_assignment.role_assignment", "operation", "update"),
				),
			},
		},
	})
}

func TestAccIBMIAMRoleAssignmentAccountGroup(t *testing.T) {
	var conf iampolicymanagementv1.RoleAssignment
	var name string = fmt.Sprintf("TerraformTemplateTest%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMRoleAssignmentDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMRoleAssignmentConfigAccountGroup(name, acc.TargetAccountGroupId),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMRoleAssignmentExists("ibm_iam_role_assignment.role_assignment", conf),
					resource.TestCheckResourceAttr("ibm_iam_role_template.role_template", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_role_template.role_template", "role.0.actions.0", roleCreateAction),
					resource.TestCheckResourceAttr("ibm_iam_role_assignment.role_assignment", "operation", "apply"),
				),
			},
		},
	})
}

func testAccCheckIBMRoleAssignmentConfigBasic(name string, targetId string) string {
	return fmt.Sprintf(`
		resource "ibm_iam_role_template" "role_template" {
			name = "%s"
			role {
				name = "TestRoleAssignment"
				display_name = "TestRoleAssignmentDisplay"
				actions = ["cloud-object-storage.bucket.delete_bucket" ]
				service_name="cloud-object-storage"
			} 
			committed=true
		}
		resource "ibm_iam_role_assignment" "role_assignment" {
			target  {
				type = "Account"
				id = "%s"
			}
			templates{
				id = ibm_iam_role_template.role_template.role_template_id 
				version = ibm_iam_role_template.role_template.version
			}
		}`, name, targetId)
}

func testAccCheckIBMRoleAssignmentConfigUpdate(name string, targetId string) string {
	return fmt.Sprintf(`
		resource "ibm_iam_role_template" "role_template" {
			name = "%s"
			role {
				name = "TestRoleAssignment"
				display_name = "TestRoleAssignmentDisplay"
				actions = ["cloud-object-storage.bucket.delete_bucket" ]
				service_name="cloud-object-storage"
			} 
			committed=true
		}
		resource "ibm_iam_role_template_version" "template_version" {
			role_template_id = ibm_iam_role_template.role_template.role_template_id
			role {
				name = "TestRoleAssignmentUpdate"
				display_name = "TestRoleAssignmentDisplay"
				actions = ["cloud-object-storage.bucket.delete_backup_policy" ]
				service_name="cloud-object-storage"
			}
			committed=true
		}
	
		resource "ibm_iam_role_assignment" "role_assignment" {
			target  {
				type = "Account"
				id = "%s"
			}

			templates{
                id = ibm_iam_role_template.role_template.role_template_id
				version = ibm_iam_role_template.role_template.version
			}
			template_version = ibm_iam_role_template_version.template_version.version
		}`, name, targetId)
}

func testAccCheckIBMRoleAssignmentConfigAccountGroup(name string, targetId string) string {
	return fmt.Sprintf(`
		resource "ibm_iam_role_template" "role_template" {
			name = "%s"
			role {
				name = "TestRoleAssignment"
				display_name = "TestRoleAssignmentDisplay"
				actions = ["cloud-object-storage.bucket.delete_bucket" ]
				service_name="cloud-object-storage"
			}
			committed=true
		}
		resource "ibm_iam_role_assignment" "role_assignment" {
			target  {
				type = "AccountGroup"
				id = "%s"
			}
			templates{
				id = ibm_iam_role_template.role_template.role_template_id 
				version = ibm_iam_role_template.role_template.version
			}
		}`, name, targetId)
}

func testAccCheckIBMRoleAssignmentExists(n string, obj iampolicymanagementv1.RoleAssignment) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		iamPolicyManagementClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).IAMPolicyManagementV1API()
		if err != nil {
			return err
		}

		getRoleAssignmentsOptions := &iampolicymanagementv1.GetRoleAssignmentOptions{}

		getRoleAssignmentsOptions.SetAssignmentID(rs.Primary.ID)

		actionControlAssignment, _, err := iamPolicyManagementClient.GetRoleAssignment(getRoleAssignmentsOptions)
		if err != nil {
			return err
		}

		obj = *actionControlAssignment
		return nil
	}
}

func testAccCheckIBMRoleAssignmentDestroy(s *terraform.State) error {
	iamPolicyManagementClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).IAMPolicyManagementV1API()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_iam_role_assignment" {
			continue
		}

		getRoleAssignmentsOptions := &iampolicymanagementv1.GetRoleAssignmentOptions{}

		getRoleAssignmentsOptions.SetAssignmentID(rs.Primary.ID)

		// Try to find the key
		_, response, err := iamPolicyManagementClient.GetRoleAssignment(getRoleAssignmentsOptions)

		if err == nil {
			return fmt.Errorf("role_assignment still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for role_assignment (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
