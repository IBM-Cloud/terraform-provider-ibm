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
	actName         string = fmt.Sprintf("TerraformActionControlTemplateTest%d", acctest.RandIntRange(10, 100))
	actServiceName  string = "am-test-service"
	actConf         iampolicymanagementv1.ActionControlTemplate
	actCreateAction string = "am-test-service.test.create"
	actDeleteAction string = "am-test-service.test.delete"
)

func TestAccIBMIAMActionControlBasicTemplate(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMActionControlTemplateDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMActionControlTemplateConfigBasic(actName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMActionControlTemplateExists("ibm_iam_action_control_template.action_control_template", actConf),
					resource.TestCheckResourceAttr("ibm_iam_action_control_template.action_control_template", "name", actName),
				),
			},
		},
	})
}

func TestAccIBMIAMActionControlBasicTemplateUpdate(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMActionControlTemplateDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMActionControlTemplateConfigBasic(actName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMActionControlTemplateExists("ibm_iam_action_control_template.action_control_template", actConf),
					resource.TestCheckResourceAttr("ibm_iam_action_control_template.action_control_template", "name", actName),
				),
			},
			{
				Config: testAccCheckIBMActionControlTemplateConfigBasic("UpdateBasicTemplate"),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMActionControlTemplateExists("ibm_iam_action_control_template.action_control_template", actConf),
					resource.TestCheckResourceAttr("ibm_iam_action_control_template.action_control_template", "name", "UpdateBasicTemplate"),
				),
			},
		},
	})
}

func TestAccIBMIAMActionControlBasicTemplateUpdateWithActionControl(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMActionControlTemplateDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMActionControlTemplateConfigBasic(actName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMActionControlTemplateExists("ibm_iam_action_control_template.action_control_template", actConf),
					resource.TestCheckResourceAttr("ibm_iam_action_control_template.action_control_template", "name", actName),
				),
			},
			{
				Config: testAccCheckIBMActionControlTemplateConfigWithActionControl(actName, actCreateAction),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMActionControlTemplateExists("ibm_iam_action_control_template.action_control_template", actConf),
					resource.TestCheckResourceAttr("ibm_iam_action_control_template.action_control_template", "name", actName),
					resource.TestCheckResourceAttr("ibm_iam_action_control_template.action_control_template", "action_control.0.actions.0", actCreateAction),
				),
			},
		},
	})
}

func TestAccIBMIAMActionControlBasicTemplateCommit(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMActionControlTemplateDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMActionControlTemplateConfigBasic(actName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMActionControlTemplateExists("ibm_iam_action_control_template.action_control_template", actConf),
					resource.TestCheckResourceAttr("ibm_iam_action_control_template.action_control_template", "name", actName),
				),
			},
			{
				Config: testAccCheckIBMActionControlTemplateConfigWithActionControlAndCommit("UpdateBasicTemplate", actCreateAction),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMActionControlTemplateExists("ibm_iam_action_control_template.action_control_template", actConf),
					resource.TestCheckResourceAttr("ibm_iam_action_control_template.action_control_template", "name", "UpdateBasicTemplate"),
					resource.TestCheckResourceAttr("ibm_iam_action_control_template.action_control_template", "action_control.0.actions.0", actCreateAction),
					resource.TestCheckResourceAttr("ibm_iam_action_control_template.action_control_template", "committed", "true"),
				),
			},
		},
	})
}

func TestAccIBMIAMActionControlTemplate(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMActionControlTemplateDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMActionControlTemplateConfigWithActionControlAndCommit("ActionControlTemplate", actCreateAction),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMActionControlTemplateExists("ibm_iam_action_control_template.action_control_template", actConf),
					resource.TestCheckResourceAttr("ibm_iam_action_control_template.action_control_template", "name", "ActionControlTemplate"),
					resource.TestCheckResourceAttr("ibm_iam_action_control_template.action_control_template", "action_control.0.actions.0", actCreateAction),
					resource.TestCheckResourceAttr("ibm_iam_action_control_template.action_control_template", "committed", "true"),
				),
			},
		},
	})
}

func testAccCheckIBMActionControlTemplateExists(n string, obj iampolicymanagementv1.ActionControlTemplate) resource.TestCheckFunc {

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

		getActionControlTemplateOptions := &iampolicymanagementv1.GetActionControlTemplateVersionOptions{
			ActionControlTemplateID: &parts[0],
			Version:                 &parts[1],
		}

		actionControlTemplate, _, err := iamPolicyManagementClient.GetActionControlTemplateVersion(getActionControlTemplateOptions)
		if err != nil {
			return err
		}
		obj = *actionControlTemplate
		return nil
	}
}

func testAccCheckIBMActionControlTemplateDestroy(s *terraform.State) error {
	iamPolicyManagementClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).IAMPolicyManagementV1API()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_iam_action_control_template" {
			continue
		}

		parts, err := flex.IdParts(rs.Primary.ID)
		if err != nil {
			return err
		}

		getActionControlTemplateOptions := &iampolicymanagementv1.GetActionControlTemplateVersionOptions{
			ActionControlTemplateID: &parts[0],
			Version:                 &parts[1],
		}

		// Try to find the key
		_, response, err := iamPolicyManagementClient.GetActionControlTemplateVersion(getActionControlTemplateOptions)

		if err == nil {
			return fmt.Errorf("ibm_iam_action_control_template still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for ibm_iam_action_control_template (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}

func testAccCheckIBMActionControlTemplateConfigBasic(name string) string {
	return fmt.Sprintf(`

		resource "ibm_iam_action_control_template" "action_control_template" {
			name = "%s"
			description = "Base template Testing"
		}
	`, name)
}

func testAccCheckIBMActionControlTemplateConfigWithActionControl(name string, action string) string {
	return fmt.Sprintf(`

		resource "ibm_iam_action_control_template" "action_control_template" {
			name = "%s"
			description = "Base template Testing"
			action_control {
				actions = ["%s" ]
				service_name="am-test-service"
			} 
		}
	`, name, action)
}

func testAccCheckIBMActionControlTemplateConfigWithActionControlAndCommit(name string, action string) string {
	return fmt.Sprintf(`

		resource "ibm_iam_action_control_template" "action_control_template" {
			name = "%s"
			description = "Base template Testing"
			action_control {
				actions = ["%s" ]
				service_name="am-test-service"
			}
			committed = true
		}
	`, name, action)
}
