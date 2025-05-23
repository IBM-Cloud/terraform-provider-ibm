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

var actVersionName string = fmt.Sprintf("TerraformActionControlTemplateVersionTest%d", acctest.RandIntRange(10, 100))

func TestAccIBMIAMActionControlVersionTemplate(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMActionControlTemplateDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMActionControlTemplateVersionConfig(actName, actCreateAction, actDeleteAction),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMActionControlTemplateExists("ibm_iam_action_control_template.action_control_template", actConf),
					resource.TestCheckResourceAttr("ibm_iam_action_control_template.action_control_template", "name", actName),
					resource.TestCheckResourceAttr("ibm_iam_action_control_template_version.action_control_template_version", "name", actName),
					resource.TestCheckResourceAttr("ibm_iam_action_control_template.action_control_template", "action_control.0.actions.0", actCreateAction),
					resource.TestCheckResourceAttr("ibm_iam_action_control_template_version.action_control_template_version", "action_control.0.actions.0", actDeleteAction),
				),
			},
		},
	})
}
func TestAccIBMIAMActionControlVersionTemplateBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMActionControlTemplateDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMActionControlTemplateVersionConfigBasic(actName, actCreateAction),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMActionControlTemplateExists("ibm_iam_action_control_template.action_control_template", actConf),
					resource.TestCheckResourceAttr("ibm_iam_action_control_template.action_control_template", "name", actName),
					resource.TestCheckResourceAttr("ibm_iam_action_control_template_version.action_control_template_version", "name", actName),
					resource.TestCheckResourceAttr("ibm_iam_action_control_template_version.action_control_template_version", "action_control.0.actions.0", actCreateAction),
				),
			},
		},
	})
}

func TestAccIBMIAMActionControlVersionTemplateUpdate(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMActionControlTemplateDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMActionControlTemplateVersionConfig(actName, actCreateAction, actDeleteAction),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMActionControlTemplateExists("ibm_iam_action_control_template.action_control_template", actConf),
					resource.TestCheckResourceAttr("ibm_iam_action_control_template.action_control_template", "name", actName),
					resource.TestCheckResourceAttr("ibm_iam_action_control_template_version.action_control_template_version", "name", actName),
					resource.TestCheckResourceAttr("ibm_iam_action_control_template.action_control_template", "action_control.0.actions.0", actCreateAction),
					resource.TestCheckResourceAttr("ibm_iam_action_control_template_version.action_control_template_version", "action_control.0.actions.0", actDeleteAction),
				),
			},
			{
				Config: testAccCheckIBMActionControlTemplateVersionUpdateConfig(actName, actCreateAction, actCreateAction, actDeleteAction, false),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMActionControlTemplateExists("ibm_iam_action_control_template.action_control_template", actConf),
					resource.TestCheckResourceAttr("ibm_iam_action_control_template.action_control_template", "name", actName),
					resource.TestCheckResourceAttr("ibm_iam_action_control_template_version.action_control_template_version", "name", actName),
					resource.TestCheckResourceAttr("ibm_iam_action_control_template.action_control_template", "action_control.0.actions.0", actCreateAction),
					resource.TestCheckResourceAttr("ibm_iam_action_control_template_version.action_control_template_version", "action_control.0.actions.#", "2"),
				),
			},
		},
	})
}

func TestAccIBMIAMActionControlVersionTemplateUpdateCommit(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMActionControlTemplateDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMActionControlTemplateVersionConfig(actName, actCreateAction, actDeleteAction),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMActionControlTemplateExists("ibm_iam_action_control_template.action_control_template", actConf),
					resource.TestCheckResourceAttr("ibm_iam_action_control_template.action_control_template", "name", actName),
					resource.TestCheckResourceAttr("ibm_iam_action_control_template_version.action_control_template_version", "name", actName),
					resource.TestCheckResourceAttr("ibm_iam_action_control_template.action_control_template", "action_control.0.actions.0", actCreateAction),
					resource.TestCheckResourceAttr("ibm_iam_action_control_template_version.action_control_template_version", "action_control.0.actions.0", actDeleteAction),
				),
			},
			{
				Config: testAccCheckIBMActionControlTemplateVersionUpdateConfig(actName, actCreateAction, actCreateAction, actDeleteAction, true),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMActionControlTemplateExists("ibm_iam_action_control_template.action_control_template", actConf),
					resource.TestCheckResourceAttr("ibm_iam_action_control_template.action_control_template", "name", actName),
					resource.TestCheckResourceAttr("ibm_iam_action_control_template_version.action_control_template_version", "name", actName),
					resource.TestCheckResourceAttr("ibm_iam_action_control_template.action_control_template", "action_control.0.actions.0", actCreateAction),
					resource.TestCheckResourceAttr("ibm_iam_action_control_template_version.action_control_template_version", "action_control.0.actions.#", "2"),
				),
			},
		},
	})
}

func testAccCheckIBMActionControlTemplateVersionConfig(name string, actionControl string, actionControl1 string) string {
	return fmt.Sprintf(`
		resource "ibm_iam_action_control_template" "action_control_template" {
			name = "%s"
			description = "Create Action Control base templates through Terraform resources"
			action_control {
				actions = ["%s"]
				service_name="am-test-service"
			}
		}
		resource "ibm_iam_action_control_template_version" "action_control_template_version" {
			action_control_template_id = ibm_iam_action_control_template.action_control_template.action_control_template_id
			description = "Create Action Control template versions through Terraform resources"
			action_control {
				actions = ["%s"]
				service_name="am-test-service"
			}
		}
	`, name, actionControl, actionControl1)
}

func testAccCheckIBMActionControlTemplateVersionUpdateConfig(name string, actionControl string, actionControl1 string, actionControl2 string, committed bool) string {
	return fmt.Sprintf(`
		resource "ibm_iam_action_control_template" "action_control_template" {
			name = "%s"
			description = "Create Action Control template versions through Terraform resources"
			action_control {
				actions = ["%s"]
				service_name="am-test-service"
			}
		}
		resource "ibm_iam_action_control_template_version" "action_control_template_version" {
			action_control_template_id = ibm_iam_action_control_template.action_control_template.action_control_template_id
			description = "Update Action Control template versions through Terraform resources"
			action_control {
				actions = ["%s", "%s"]
				service_name="am-test-service"
			}
			committed = %t
		}
	`, name, actionControl, actionControl1, actionControl2, committed)
}

func testAccCheckIBMActionControlTemplateVersionConfigBasic(name string, actionControl string) string {
	return fmt.Sprintf(`
		resource "ibm_iam_action_control_template" "action_control_template" {
			name = "%s"
			description = "Create Action Control basic template through Terraform resources"
		}
		resource "ibm_iam_action_control_template_version" "action_control_template_version" {
			action_control_template_id = ibm_iam_action_control_template.action_control_template.action_control_template_id
			description = "Create Action Control template versions under basic template through Terraform resources"
			action_control {
				actions = ["%s"]
				service_name="am-test-service"
			}
		}
	`, name, actionControl)
}
