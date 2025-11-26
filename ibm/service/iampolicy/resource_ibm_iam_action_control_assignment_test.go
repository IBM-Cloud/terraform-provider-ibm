// Copyright IBM Corp. 2024 All Rights Reserved.
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
	"github.com/IBM/platform-services-go-sdk/iampolicymanagementv1"
)

func TestAccIBMIAMActionControlAssignmentBasic(t *testing.T) {
	var conf iampolicymanagementv1.ActionControlAssignment
	var name string = fmt.Sprintf("TerraformTemplateTest%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMActionControlAssignmentDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMActionControlAssignmentConfigBasic(name, acc.TargetAccountId),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMActionControlAssignmentExists("ibm_iam_action_control_assignment.action_control_assignment", conf),
					resource.TestCheckResourceAttr("ibm_iam_action_control_template.action_control_template", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_action_control_template.action_control_template", "action_control.0.actions.0", actCreateAction),
					resource.TestCheckResourceAttr("ibm_iam_action_control_assignment.action_control_assignment", "operation", "apply"),
				),
			},
		},
	})
}

func TestAccIBMIAMActionControlAssignmentUpdate(t *testing.T) {
	var conf iampolicymanagementv1.ActionControlAssignment
	var name string = fmt.Sprintf("TerraformTemplateTest%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMActionControlAssignmentDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMActionControlAssignmentConfigBasic(name, acc.TargetAccountId),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMActionControlAssignmentExists("ibm_iam_action_control_assignment.action_control_assignment", conf),
					resource.TestCheckResourceAttr("ibm_iam_action_control_template.action_control_template", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_action_control_template.action_control_template", "action_control.0.actions.0", actCreateAction),
				),
			},
			{
				Config: testAccCheckIBMActionControlAssignmentConfigUpdate(name, acc.TargetAccountId),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMActionControlAssignmentExists("ibm_iam_action_control_assignment.action_control_assignment", conf),
					resource.TestCheckResourceAttr("ibm_iam_action_control_template_version.template_version", "action_control.0.actions.0", actDeleteAction),
					resource.TestCheckResourceAttr("ibm_iam_action_control_assignment.action_control_assignment", "operation", "update"),
				),
			},
		},
	})
}

func TestAccIBMIAMActionControlAssignmentAccountGroup(t *testing.T) {
	var conf iampolicymanagementv1.ActionControlAssignment
	var name string = fmt.Sprintf("TerraformTemplateTest%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMActionControlAssignmentDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMActionControlAssignmentConfigAccountGroup(name, acc.TargetAccountGroupId),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMActionControlAssignmentExists("ibm_iam_action_control_assignment.action_control_assignment", conf),
					resource.TestCheckResourceAttr("ibm_iam_action_control_template.action_control_template", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_action_control_template.action_control_template", "action_control.0.actions.0", actCreateAction),
					resource.TestCheckResourceAttr("ibm_iam_action_control_assignment.action_control_assignment", "operation", "apply"),
				),
			},
		},
	})
}

func testAccCheckIBMActionControlAssignmentConfigBasic(name string, targetId string) string {
	return fmt.Sprintf(`
		resource "ibm_iam_action_control_template" "action_control_template" {
			name = "%s"
			action_control {
				actions = ["am-test-service.test.create" ]
				service_name="am-test-service"
			} 
			committed=true
		}
		resource "ibm_iam_action_control_assignment" "action_control_assignment" {
			target  ={
				type = "Account"
				id = "%s"
			}
			templates{
				id = ibm_iam_action_control_template.action_control_template.action_control_template_id 
				version = ibm_iam_action_control_template.action_control_template.version
			}
		}`, name, targetId)
}

func testAccCheckIBMActionControlAssignmentConfigUpdate(name string, targetId string) string {
	return fmt.Sprintf(`
		resource "ibm_iam_action_control_template" "action_control_template" {
			name = "%s"
			action_control {
				actions = ["am-test-service.test.create" ]
				service_name="am-test-service"
			} 
			committed=true
		}
		resource "ibm_iam_action_control_template_version" "template_version" {
			action_control_template_id = ibm_iam_action_control_template.action_control_template.action_control_template_id
			action_control {
				actions = ["am-test-service.test.delete" ]
				service_name="am-test-service"
			}
			committed=true
		}
	
		resource "ibm_iam_action_control_assignment" "action_control_assignment" {
			target  ={
				type = "Account"
				id = "%s"
			}

			templates{
                id = ibm_iam_action_control_template.action_control_template.action_control_template_id
				version = ibm_iam_action_control_template.action_control_template.version
			}
			template_version = ibm_iam_action_control_template_version.template_version.version
		}`, name, targetId)
}

func testAccCheckIBMActionControlAssignmentConfigAccountGroup(name string, targetId string) string {
	return fmt.Sprintf(`
		resource "ibm_iam_action_control_template" "action_control_template" {
			name = "%s"
			action_control {
				actions = ["am-test-service.test.create" ]
				service_name="am-test-service"
			}
			committed=true
		}
		resource "ibm_iam_action_control_assignment" "action_control_assignment" {
			target  ={
				type = "AccountGroup"
				id = "%s"
			}
			templates{
				id = ibm_iam_action_control_template.action_control_template.action_control_template_id 
				version = ibm_iam_action_control_template.action_control_template.version
			}
		}`, name, targetId)
}

func testAccCheckIBMActionControlAssignmentExists(n string, obj iampolicymanagementv1.ActionControlAssignment) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		iamPolicyManagementClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).IAMPolicyManagementV1API()
		if err != nil {
			return err
		}

		getActionControlAssignmentsOptions := &iampolicymanagementv1.GetActionControlAssignmentOptions{}

		getActionControlAssignmentsOptions.SetAssignmentID(rs.Primary.ID)

		actionControlAssignment, _, err := iamPolicyManagementClient.GetActionControlAssignment(getActionControlAssignmentsOptions)
		if err != nil {
			return err
		}

		obj = *actionControlAssignment
		return nil
	}
}

func testAccCheckIBMActionControlAssignmentDestroy(s *terraform.State) error {
	iamPolicyManagementClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).IAMPolicyManagementV1API()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_iam_action_control_assignment" {
			continue
		}

		getActionControlAssignmentsOptions := &iampolicymanagementv1.GetActionControlAssignmentOptions{}

		getActionControlAssignmentsOptions.SetAssignmentID(rs.Primary.ID)

		// Try to find the key
		_, response, err := iamPolicyManagementClient.GetActionControlAssignment(getActionControlAssignmentsOptions)

		if err == nil {
			return fmt.Errorf("action_control_assignment still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for action_control_assignment (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
